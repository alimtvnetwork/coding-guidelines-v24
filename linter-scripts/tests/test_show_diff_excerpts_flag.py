"""CLI-level tests for the tri-state ``--show-diff-excerpts`` flag.

The flag decouples excerpt rendering from ``--diff-context``:

* ``--show-diff-excerpts``     forces excerpts ON (even with ctx=0)
* ``--no-show-diff-excerpts``  forces excerpts OFF (even with ctx>0
                               or ``--json-excerpts``)
* unset (auto)                 keeps the historical behaviour
                               (human: ctx>0; JSON: --json-excerpts)

These tests exercise the resolution table end-to-end via ``main()``
to guarantee that:

1. The default JSON schema stays byte-identical when nothing is
   forced (legacy consumers see only file/line/code/message).
2. Forcing ON with ``--diff-context=0`` still produces a usable
   excerpt (effective context falls back to 1, not a bare focus row).
3. Forcing OFF with ``--diff-context=3`` suppresses the excerpt in
   human output AND the ``excerpt`` key in JSON, even when
   ``--json-excerpts`` is also set (force-off wins).
4. The flag does NOT touch ``--suggest-patch`` /
   ``--json-suggest-patch`` — those keep working under their own
   gates so authors can disable noise but still get the fix scaffold.

To stay sandbox-safe (no ``git add``/``git commit`` available), we
monkey-patch the two git-touching helpers (``_resolve_changed_md``
and ``_fetch_diff_excerpts``) with deterministic stand-ins that
return a hand-crafted ``_DiffExcerpts`` derived from the on-disk
spec file. Everything downstream — argparse, gating, rendering —
runs as production code.
"""

from __future__ import annotations

import io
import json
import sys
import tempfile
import unittest
from contextlib import redirect_stdout
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parent))

from conftest_shim import load_placeholder_linter  # noqa: E402

chk = load_placeholder_linter()


class _StubbedGit:
    """Context manager that replaces the two git-touching helpers
    with deterministic stand-ins. The originals are restored on
    exit so concurrent tests in the same process don't interfere.
    """

    def __init__(self, repo_root: Path, rel_path: str,
                 post_lines: dict[int, tuple[str, str]]):
        self.repo_root = repo_root
        self.rel_path = rel_path
        self.post_lines = post_lines
        self._orig_resolve = None
        self._orig_fetch = None

    def __enter__(self):
        self._orig_resolve = chk._resolve_changed_md
        self._orig_fetch = chk._fetch_diff_excerpts
        target = (self.repo_root / self.rel_path).resolve()

        def fake_resolve(repo_root, root, *, diff_base, changed_files):
            return {target}

        # Build a single-hunk _DiffExcerpts spanning every supplied
        # post-state line so the renderer treats them all as inside
        # the hunk (no "nearest hunk" breadcrumb noise in tests).
        min_ln = min(self.post_lines)
        max_ln = max(self.post_lines)
        hunk = chk._Hunk(start=min_ln, end=max_ln)
        excerpt = chk._DiffExcerpts(
            lines=dict(self.post_lines),
            min_line=min_ln,
            max_line=max_ln,
            hunks=(hunk,),
        )

        def fake_fetch(repo_root, diff_base, rel_path, context):
            return excerpt if rel_path == self.rel_path else None

        chk._resolve_changed_md = fake_resolve
        chk._fetch_diff_excerpts = fake_fetch
        return self

    def __exit__(self, *exc):
        chk._resolve_changed_md = self._orig_resolve
        chk._fetch_diff_excerpts = self._orig_fetch


def _run_main(argv: list[str]) -> tuple[int, str]:
    """Invoke main() with stdout captured. Returns (rc, stdout)."""
    buf = io.StringIO()
    with redirect_stdout(buf):
        rc = chk.main(argv)
    return rc, buf.getvalue()


class ShowDiffExcerptsResolution(unittest.TestCase):
    """Walk the resolution table for human and JSON output."""

    def setUp(self) -> None:
        self._tmp = tempfile.TemporaryDirectory()
        self.repo = Path(self._tmp.name)
        self.spec = self.repo / "spec"
        self.spec.mkdir()
        # A bare TODO trips P-001 ("missing imperative reason") on
        # line 5 of demo.md. The post-state lines we feed to the
        # stub mirror the on-disk file so render() lines up.
        (self.spec / "demo.md").write_text(
            "# Demo\n"          # 1
            "\n"                # 2
            "Body line.\n"      # 3
            "\n"                # 4
            "<!-- TODO -->\n",  # 5  ← violation
            encoding="utf-8",
        )
        self.post = {
            1: (" ", "# Demo"),
            2: (" ", ""),
            3: (" ", "Body line."),
            4: ("+", ""),
            5: ("+", "<!-- TODO -->"),
        }

    def tearDown(self) -> None:
        self._tmp.cleanup()

    def _argv(self, *extra: str) -> list[str]:
        return [
            "--root", str(self.spec),
            "--repo-root", str(self.repo),
            "--diff-base", "STUB",
            *extra,
        ]

    # ---- Auto (unset) preserves historical behaviour --------------

    def test_auto_human_with_ctx_emits_excerpt(self) -> None:
        with _StubbedGit(self.repo, "spec/demo.md", self.post):
            rc, out = _run_main(self._argv("--diff-context", "2"))
        self.assertEqual(rc, 1)
        # Excerpt rendered → focus marker (►) shows up under the
        # violation row in the indented human block.
        self.assertIn("►", out, f"missing focus marker in:\n{out}")
        self.assertIn("<!-- TODO -->", out)

    def test_auto_human_with_ctx_zero_omits_excerpt(self) -> None:
        with _StubbedGit(self.repo, "spec/demo.md", self.post):
            rc, out = _run_main(self._argv("--diff-context", "0"))
        self.assertEqual(rc, 1)
        self.assertNotIn("►", out,
                         f"excerpt leaked at ctx=0 (auto):\n{out}")

    def test_auto_json_omits_excerpt_field(self) -> None:
        with _StubbedGit(self.repo, "spec/demo.md", self.post):
            rc, out = _run_main(self._argv("--json"))
        self.assertEqual(rc, 1)
        payload = json.loads(out)
        self.assertTrue(payload, "expected at least one violation")
        for v in payload:
            self.assertEqual(set(v), {"file", "line", "code", "message"},
                             f"legacy schema leaked extras: {v}")

    # ---- Force ON ------------------------------------------------

    def test_force_on_overrides_ctx_zero_in_human(self) -> None:
        with _StubbedGit(self.repo, "spec/demo.md", self.post):
            rc, out = _run_main(self._argv(
                "--diff-context", "0",
                "--show-diff-excerpts",
            ))
        self.assertEqual(rc, 1)
        # Effective context falls back to 1, so we get a focus row
        # PLUS at least one neighbour (line 4 above or no line below
        # since the violation is the last post-state row in this fixture).
        self.assertIn("►", out,
                      f"force-on did not produce excerpt at ctx=0:\n{out}")

    def test_force_on_implies_json_excerpts_field(self) -> None:
        with _StubbedGit(self.repo, "spec/demo.md", self.post):
            rc, out = _run_main(self._argv(
                "--json",
                "--show-diff-excerpts",
            ))
        self.assertEqual(rc, 1)
        payload = json.loads(out)
        with_excerpt = [v for v in payload if "excerpt" in v]
        self.assertTrue(with_excerpt,
                        f"force-on did not add excerpt to JSON:\n{out}")
        # Each excerpt row must carry the documented schema.
        for row in with_excerpt[0]["excerpt"]:
            self.assertEqual(
                set(row), {"line", "kind", "text", "focus", "nearest"},
                f"unexpected excerpt row schema: {row}",
            )

    # ---- Force OFF -----------------------------------------------

    def test_force_off_suppresses_human_excerpt_with_ctx(self) -> None:
        with _StubbedGit(self.repo, "spec/demo.md", self.post):
            rc, out = _run_main(self._argv(
                "--diff-context", "3",
                "--no-show-diff-excerpts",
            ))
        self.assertEqual(rc, 1)
        self.assertNotIn("►", out,
                         f"excerpt leaked despite force-off:\n{out}")
        # The violation header itself still prints — only the
        # excerpt block is suppressed.
        self.assertIn("[P-001]", out)

    def test_force_off_beats_json_excerpts_flag(self) -> None:
        # ``--json-excerpts`` would normally enable the field; the
        # force-off override must win.
        with _StubbedGit(self.repo, "spec/demo.md", self.post):
            rc, out = _run_main(self._argv(
                "--json",
                "--json-excerpts",
                "--no-show-diff-excerpts",
            ))
        self.assertEqual(rc, 1)
        payload = json.loads(out)
        for v in payload:
            self.assertNotIn("excerpt", v,
                             f"force-off lost to --json-excerpts: {v}")

    # ---- Independence from --suggest-patch -----------------------

    def test_force_off_keeps_suggest_patch_intact(self) -> None:
        with _StubbedGit(self.repo, "spec/demo.md", self.post):
            rc, out = _run_main(self._argv(
                "--diff-context", "3",
                "--no-show-diff-excerpts",
                "--suggest-patch",
            ))
        self.assertEqual(rc, 1)
        # No excerpt focus marker (force-off honored)…
        self.assertNotIn("►", out,
                         f"excerpt leaked alongside suggest-patch:\n{out}")
        # …but the suggested-patch fences and TODO marker still print.
        self.assertIn("BEGIN SUGGESTED PATCH", out,
                      f"suggest-patch fence missing:\n{out}")
        self.assertIn("TODO(P-001)", out,
                      f"rule-specific TODO missing:\n{out}")

    def test_force_off_keeps_json_suggest_patch_intact(self) -> None:
        with _StubbedGit(self.repo, "spec/demo.md", self.post):
            rc, out = _run_main(self._argv(
                "--json",
                "--no-show-diff-excerpts",
                "--json-suggest-patch",
            ))
        self.assertEqual(rc, 1)
        payload = json.loads(out)
        with_patch = [v for v in payload if "suggested_patch" in v]
        self.assertTrue(with_patch,
                        f"force-off broke --json-suggest-patch:\n{out}")
        for v in payload:
            self.assertNotIn("excerpt", v,
                             f"excerpt leaked into JSON: {v}")


if __name__ == "__main__":
    unittest.main(verbosity=2)
