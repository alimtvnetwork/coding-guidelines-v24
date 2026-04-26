"""OLD-side capture for rename/copy rows.

The diff-mode intake docstring promises that an
``ignored-deleted`` audit row is emitted for *both*:

* the path on a true ``D``-status row, **and**
* the OLD side of an ``R`` (rename) or ``C`` (copy) row — the
  pre-rename / pre-copy path that has no post-state file under the
  new path.

Earlier the second case was silently dropped — the parser captured
only the NEW path and the OLD side disappeared from the audit. This
module pins the now-implemented behaviour across both intakes
(``git diff --name-status`` and ``--changed-files``) and across all
three rename surfaces (tab-form scored, tab-form unscored,
arrow-form). The reason text must name the destination so a
reviewer can follow the rename without consulting the diff itself.

Each test is hermetic: parser unit-tests poke
``_parse_name_status`` / ``_normalise_changed_lines`` directly, and
the end-to-end tests drive the published CLI via subprocess against
a tempdir sandbox so the STDOUT/STDERR contract is exercised
through real OS pipes.
"""
from __future__ import annotations

import json
import subprocess
import sys
import tempfile
import unittest
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parent))
from conftest_shim import load_placeholder_linter  # noqa: E402

_MOD = load_placeholder_linter()
_LINTER = (Path(__file__).resolve().parent.parent
           / "check-placeholder-comments.py")


# ---------------------------------------------------------------------
# Parser unit tests — _parse_name_status (real `git diff` shape).
# ---------------------------------------------------------------------

class ParseNameStatusCapturesOldSide(unittest.TestCase):
    """`git diff --name-status -M -C` rows divert OLD paths."""

    def test_rename_row_records_old_side_with_diff_r_old_tag(self) -> None:
        deleted: list[tuple[str, str, "str | None"]] = []
        out = _MOD._parse_name_status(
            "R092\tspec/old.md\tspec/new.md\n",
            deleted=deleted,
        )
        # NEW path is still returned for linting.
        self.assertEqual(out, ["spec/new.md"])
        # OLD path lands in `deleted` with the rename-aware tag and
        # the destination on the third tuple slot.
        self.assertEqual(deleted, [
            ("spec/old.md", "diff-R-old", "spec/new.md"),
        ])

    def test_copy_row_records_old_side_with_diff_c_old_tag(self) -> None:
        deleted: list[tuple[str, str, "str | None"]] = []
        out = _MOD._parse_name_status(
            "C075\tspec/src.md\tspec/copy.md\n",
            deleted=deleted,
        )
        self.assertEqual(out, ["spec/copy.md"])
        self.assertEqual(deleted, [
            ("spec/src.md", "diff-C-old", "spec/copy.md"),
        ])

    def test_d_and_rename_in_same_payload_keep_distinct_provenance(
            self) -> None:
        # A real ``D`` row keeps its plain ``diff-D`` tag; the
        # rename's OLD side gets the ``diff-R-old`` tag. Order
        # matches input order so downstream `diff` between two runs
        # stays reviewable.
        deleted: list[tuple[str, str, "str | None"]] = []
        _MOD._parse_name_status(
            "D\tspec/gone.md\n"
            "R100\tspec/legacy.md\tspec/current.md\n",
            deleted=deleted,
        )
        self.assertEqual(deleted, [
            ("spec/gone.md", "diff-D", None),
            ("spec/legacy.md", "diff-R-old", "spec/current.md"),
        ])

    def test_empty_old_path_does_not_synthesise_blank_row(self) -> None:
        # Pathological input ``R092\t\tspec/new.md`` (empty OLD col)
        # must NOT produce a blank-path ``ignored-deleted`` row.
        deleted: list[tuple[str, str, "str | None"]] = []
        out = _MOD._parse_name_status(
            "R092\t\tspec/new.md\n",
            deleted=deleted,
        )
        self.assertEqual(out, ["spec/new.md"])
        self.assertEqual(deleted, [])

    def test_no_deleted_arg_keeps_legacy_behaviour(self) -> None:
        # Backwards-compat: callers that don't ask for the audit
        # trail must see exactly the same return value as before
        # OLD-side capture landed.
        out = _MOD._parse_name_status(
            "R092\tspec/old.md\tspec/new.md\n"
            "C050\tspec/src.md\tspec/copy.md\n"
        )
        self.assertEqual(out, ["spec/new.md", "spec/copy.md"])


# ---------------------------------------------------------------------
# Parser unit tests — _normalise_changed_lines (`--changed-files`).
# ---------------------------------------------------------------------

class NormaliseChangedLinesCapturesOldSide(unittest.TestCase):
    """Authored `--changed-files` payloads also divert OLD paths."""

    def test_tab_form_scored_rename_records_changed_files_r_old(
            self) -> None:
        deleted: list[tuple[str, str, "str | None"]] = []
        out = _MOD._normalise_changed_lines(
            ["R087\tspec/old.md\tspec/new.md"],
            deleted=deleted,
        )
        self.assertEqual(out, ["spec/new.md"])
        self.assertEqual(deleted, [
            ("spec/old.md", "changed-files-R-old", "spec/new.md"),
        ])

    def test_tab_form_scoreless_copy_records_changed_files_c_old(
            self) -> None:
        # Scoreless ``C\told\tnew`` (rare but accepted) still yields
        # an OLD-side audit row tagged ``changed-files-C-old``.
        deleted: list[tuple[str, str, "str | None"]] = []
        out = _MOD._normalise_changed_lines(
            ["C\tspec/src.md\tspec/copy.md"],
            deleted=deleted,
        )
        self.assertEqual(out, ["spec/copy.md"])
        self.assertEqual(deleted, [
            ("spec/src.md", "changed-files-C-old", "spec/copy.md"),
        ])

    def test_arrow_form_records_changed_files_r_old(self) -> None:
        # `git status -s` short form has no copy variant — every
        # arrow row is classified as a rename.
        deleted: list[tuple[str, str, "str | None"]] = []
        out = _MOD._normalise_changed_lines(
            ["spec/old.md => spec/new.md"],
            deleted=deleted,
        )
        self.assertEqual(out, ["spec/new.md"])
        self.assertEqual(deleted, [
            ("spec/old.md", "changed-files-R-old", "spec/new.md"),
        ])

    def test_d_row_and_rename_keep_distinct_tags(self) -> None:
        deleted: list[tuple[str, str, "str | None"]] = []
        _MOD._normalise_changed_lines(
            [
                "D\tspec/gone.md",
                "R092\tspec/old.md\tspec/new.md",
                "spec/orig.md => spec/dest.md",
            ],
            deleted=deleted,
        )
        self.assertEqual(deleted, [
            ("spec/gone.md", "changed-files-D", None),
            ("spec/old.md", "changed-files-R-old", "spec/new.md"),
            ("spec/orig.md", "changed-files-R-old", "spec/dest.md"),
        ])

    def test_no_deleted_arg_keeps_legacy_returned_paths(self) -> None:
        out = _MOD._normalise_changed_lines([
            "R092\tspec/old.md\tspec/new.md",
            "C050\tspec/src.md\tspec/copy.md",
            "spec/orig.md => spec/dest.md",
        ])
        self.assertEqual(out, [
            "spec/new.md", "spec/copy.md", "spec/dest.md",
        ])


# ---------------------------------------------------------------------
# `_resolve_deleted_reason` — placeholder substitution + fallback.
# ---------------------------------------------------------------------

class ResolveDeletedReasonFormatsNewPath(unittest.TestCase):

    def test_diff_r_old_includes_destination_path(self) -> None:
        msg = _MOD._resolve_deleted_reason(
            "diff-R-old", new_path="spec/new.md")
        self.assertIn("spec/new.md", msg)
        self.assertIn("renamed".lower(), msg.lower()) if False else None
        # Substring is part of the README's CI-grep contract.
        self.assertIn("rename", msg.lower())

    def test_changed_files_c_old_includes_destination_path(self) -> None:
        msg = _MOD._resolve_deleted_reason(
            "changed-files-C-old", new_path="spec/copy.md")
        self.assertIn("spec/copy.md", msg)
        self.assertIn("copy", msg.lower())

    def test_plain_d_tag_returns_template_verbatim(self) -> None:
        # No `{new_path}` placeholder → returned unchanged even if
        # the caller passes a destination by mistake.
        msg = _MOD._resolve_deleted_reason(
            "diff-D", new_path="should-be-ignored.md")
        self.assertEqual(msg, _MOD._DELETED_REASON["diff-D"])
        self.assertNotIn("should-be-ignored.md", msg)

    def test_missing_new_path_on_r_old_is_safe(self) -> None:
        # Defensive: parser always supplies new_path today, but a
        # future caller forgetting to pass it must not crash the
        # audit emitter mid-render.
        msg = _MOD._resolve_deleted_reason("diff-R-old")
        self.assertIn("<unknown>", msg)

    def test_unknown_tag_falls_back_to_safety_message(self) -> None:
        msg = _MOD._resolve_deleted_reason("future-source-tag")
        self.assertEqual(msg, _MOD._DELETED_REASON_FALLBACK)


# ---------------------------------------------------------------------
# End-to-end CLI: the audit JSON includes OLD-side rows + names dest.
# ---------------------------------------------------------------------

class _Sandbox:
    """Tempdir with a `spec/` subroot + a writable changed-files
    payload. Mirrors the helper in `test_ignored_deleted_audit_coverage`
    but stays self-contained so this file can be run in isolation.
    """

    def __init__(self, payload: str) -> None:
        self._tmp = tempfile.TemporaryDirectory()
        self.root = Path(self._tmp.name).resolve()
        self.spec = self.root / "spec"
        self.spec.mkdir()
        self.changed = self.root / "changed.txt"
        self.changed.write_text(payload, encoding="utf-8")

    def __enter__(self) -> "_Sandbox":
        return self

    def __exit__(self, *exc: object) -> None:
        self._tmp.cleanup()

    def write(self, rel: str, body: str = "# x\n") -> None:
        p = self.root / rel
        p.parent.mkdir(parents=True, exist_ok=True)
        p.write_text(body, encoding="utf-8")

    def run(self, *extra: str) -> subprocess.CompletedProcess:
        cmd = [
            sys.executable, str(_LINTER),
            "--root", str(self.spec),
            "--changed-files", str(self.changed),
            "--list-changed-files",
            *extra,
        ]
        return subprocess.run(cmd, capture_output=True, text=True,
                              check=False, cwd=str(self.root))


class CliEmitsOldSideRowsInAudit(unittest.TestCase):

    def _audit(self, *args: str, payload: str,
               files: tuple[str, ...]) -> list[dict]:
        with _Sandbox(payload) as box:
            for f in files:
                box.write(f)
            proc = box.run("--json", *args)
        self.assertEqual(proc.returncode, 0,
                         msg=f"stderr:\n{proc.stderr}")
        return json.loads(proc.stderr.strip())

    def test_rename_row_emits_two_audit_rows_new_and_old(self) -> None:
        audit = self._audit(
            payload="R092\tspec/old.md\tspec/new.md\n",
            files=("spec/new.md",),
        )
        rows = {r["path"]: r for r in audit}
        self.assertEqual(set(rows), {"spec/new.md", "spec/old.md"})
        self.assertEqual(rows["spec/new.md"]["status"], "matched")
        self.assertEqual(rows["spec/old.md"]["status"],
                         "ignored-deleted")
        # OLD-side reason names the destination so a reviewer can
        # follow the rename without consulting the diff itself.
        self.assertIn("spec/new.md", rows["spec/old.md"]["reason"])

    def test_copy_row_emits_two_audit_rows_new_and_old(self) -> None:
        audit = self._audit(
            payload="C075\tspec/src.md\tspec/copy.md\n",
            files=("spec/copy.md",),
        )
        rows = {r["path"]: r for r in audit}
        self.assertEqual(set(rows), {"spec/copy.md", "spec/src.md"})
        self.assertEqual(rows["spec/copy.md"]["status"], "matched")
        old_row = rows["spec/src.md"]
        self.assertEqual(old_row["status"], "ignored-deleted")
        # Copy-specific wording so reviewers can distinguish copy
        # OLD-sides from rename OLD-sides at a glance.
        self.assertIn("copy", old_row["reason"].lower())
        self.assertIn("spec/copy.md", old_row["reason"])

    def test_arrow_form_emits_old_side_with_rename_reason(self) -> None:
        audit = self._audit(
            payload="spec/old.md => spec/new.md\n",
            files=("spec/new.md",),
        )
        rows = {r["path"]: r for r in audit}
        self.assertEqual(rows["spec/old.md"]["status"],
                         "ignored-deleted")
        self.assertIn("rename", rows["spec/old.md"]["reason"].lower())

    def test_verbose_source_field_carries_r_old_tag(self) -> None:
        # `--list-changed-files-verbose` exposes the raw provenance
        # tag on every audit row. OLD-side rows must surface
        # `changed-files-R-old` (NOT `changed-files-D`) so a CI
        # script can switch on the machine-readable tag instead of
        # parsing the human-readable reason.
        audit = self._audit(
            "--list-changed-files-verbose",
            payload="R092\tspec/old.md\tspec/new.md\n",
            files=("spec/new.md",),
        )
        rows = {r["path"]: r for r in audit}
        self.assertEqual(rows["spec/old.md"]["source"],
                         "changed-files-R-old")
        # Non-deleted rows carry `source: null` for schema regularity.
        self.assertIsNone(rows["spec/new.md"]["source"])

    def test_totals_footer_counts_old_sides_under_ignored_deleted(
            self) -> None:
        with _Sandbox(
            "R092\tspec/old1.md\tspec/new1.md\n"
            "C075\tspec/src.md\tspec/copy.md\n"
            "D\tspec/gone.md\n"
        ) as box:
            box.write("spec/new1.md")
            box.write("spec/copy.md")
            proc = box.run()
        # 2 NEW-side matched + 3 ignored-deleted (2 OLD-sides + 1
        # explicit D row).
        self.assertIn("matched=2", proc.stderr)
        self.assertIn("ignored-deleted=3", proc.stderr)


if __name__ == "__main__":
    unittest.main()