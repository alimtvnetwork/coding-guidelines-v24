"""Audit-trail schema validation — every row must carry a `reason`,
and every `ignored-deleted` row's reason must come from the closed
shared vocabulary in `audit_reason_vocab`.

The placeholder linter promises two contracts on its
`--list-changed-files` audit output:

1. **Universal `reason` field.** Every row, regardless of `status`,
   carries a non-empty string `reason`. This is the operator's
   one-line explanation of *why* the path landed in this row's
   bucket, and CI dashboards rely on it being present so they don't
   need a per-status branch just to render a tooltip.

2. **Closed `ignored-deleted` vocabulary.** When `status ==
   "ignored-deleted"`, the row's `reason` must be one of the texts
   produced by `audit_reason_vocab.resolve_deleted_reason()` for a
   tag in `DELETED_SOURCES` — i.e. either a flat template
   (`diff-D`, `changed-files-D`) or a `{new_path}`-substituted
   template (`*-R-old`, `*-C-old`). Free-form reason strings on
   deleted rows would defeat the single-source-of-truth invariant
   the shared vocab module is designed to enforce.

The fallback message (`DELETED_REASON_FALLBACK`) is *intentionally*
treated as a contract violation here: it's a runtime safety net for
parser drift, not a value the shipped linter should ever emit on a
clean payload. A test failure on the fallback means a parser is
producing a tag the vocabulary doesn't know about — exactly the
drift these tests are designed to catch.

Both contracts are validated across all three audit surfaces:
JSON (via `--json`), CSV (via `--similarity-csv -`), and the
plain-text table (parsed loosely by extracting the trailing
reason cell of each row). Each surface runs against the same
fixture so a discrepancy across formats is also caught.
"""
from __future__ import annotations

import csv as _csv
import io
import json
import re
import subprocess
import sys
import tempfile
import unittest
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parent))
from conftest_shim import (  # noqa: E402
    load_audit_reason_vocab,
    load_placeholder_linter,
)

_MOD = load_placeholder_linter()
_VOCAB = load_audit_reason_vocab()
_LINTER = (Path(__file__).resolve().parent.parent
           / "check-placeholder-comments.py")


# ---------------------------------------------------------------------
# Fixture — one payload that exercises every status the linter can
# produce on a `--list-changed-files` audit, so the schema tests
# don't pass vacuously when only `matched` rows are present.
#
# Status coverage (post-dedupe, pre-filter):
#   matched              — `spec/new1.md` (NEW side of rename)
#                          + `spec/copy.md` (NEW side of copy)
#                          + `spec/dest.md` (NEW side of arrow rename)
#                          + `spec/touched.md` (plain modification)
#   ignored-extension    — `spec/notes.txt` (txt not in default allowlist)
#   ignored-out-of-root  — `outside/elsewhere.md` (above --root)
#   ignored-missing      — `spec/ghost.md` (path absent from disk)
#   ignored-deleted      — covers EVERY tag in `DELETED_SOURCES`:
#                            • `spec/gone.md`          → changed-files-D
#                            • `spec/old1.md`          → changed-files-R-old
#                            • `spec/src.md`           → changed-files-C-old
#                            • `spec/orig.md`          → changed-files-R-old
#                          (`diff-*` tags are exercised via the
#                          `_DIFF_PAYLOAD` fixture below.)
# ---------------------------------------------------------------------

_CHANGED_FILES_PAYLOAD = (
    # Plain modify → matched
    "spec/touched.md\n"
    # Plain delete → ignored-deleted, source=changed-files-D
    "D\tspec/gone.md\n"
    # Rename tab form → matched (NEW) + ignored-deleted (OLD,
    # source=changed-files-R-old)
    "R092\tspec/old1.md\tspec/new1.md\n"
    # Copy tab form → matched (NEW) + ignored-deleted (OLD,
    # source=changed-files-C-old)
    "C075\tspec/src.md\tspec/copy.md\n"
    # Rename arrow form → matched (NEW) + ignored-deleted (OLD,
    # source=changed-files-R-old)
    "spec/orig.md => spec/dest.md\n"
    # Wrong extension → ignored-extension
    "spec/notes.txt\n"
    # Outside --root → ignored-out-of-root
    "outside/elsewhere.md\n"
    # On-disk absent → ignored-missing
    "spec/ghost.md\n"
)

class _Sandbox:
    """Tempdir with the on-disk files the fixture's NEW-side rows
    need to resolve to `matched` (rather than `ignored-missing`)."""

    def __init__(self) -> None:
        self._tmp = tempfile.TemporaryDirectory()
        self.root = Path(self._tmp.name).resolve()
        (self.root / "spec").mkdir()
        (self.root / "outside").mkdir()
        # Files the matched-side rows expect to find on disk.
        for rel in (
            "spec/touched.md",
            "spec/new1.md",
            "spec/copy.md",
            "spec/dest.md",
            "spec/diff_new.md",
            "spec/diff_copy.md",
            # `spec/notes.txt` is created so the path resolves to
            # `ignored-extension` (extension filter trumps presence).
            "spec/notes.txt",
            # `outside/elsewhere.md` is created so the path resolves
            # to `ignored-out-of-root` (root check trumps presence).
            "outside/elsewhere.md",
        ):
            (self.root / rel).write_text("# x\n", encoding="utf-8")
        self.changed = self.root / "changed.txt"
        self.changed.write_text(_CHANGED_FILES_PAYLOAD,
                                encoding="utf-8")

    def __enter__(self) -> "_Sandbox":
        return self

    def __exit__(self, *exc: object) -> None:
        self._tmp.cleanup()

    def run_changed_files(
        self, *extra: str, csv_path: "Path | None" = None,
    ) -> subprocess.CompletedProcess:
        cmd = [
            sys.executable, str(_LINTER),
            "--root", str(self.root / "spec"),
            "--changed-files", str(self.changed),
            "--list-changed-files",
            *extra,
        ]
        return subprocess.run(cmd, capture_output=True, text=True,
                              check=False, cwd=str(self.root))


# ---------------------------------------------------------------------
# Pre-compute the canonical set of valid `ignored-deleted` reason
# strings ONCE so every test can membership-check in O(1). The
# placeholder-substituted reasons depend on the runtime `new_path`
# value, so we pre-render them for the exact destinations the
# fixture uses; for any other `new_path` we also accept the
# canonical templates with `<unknown>` substituted, which the
# vocab module emits when `new_path` is absent.
# ---------------------------------------------------------------------

def _expected_deleted_reasons(
    new_paths_by_source: dict[str, set[str]],
) -> set[str]:
    """Build the closed set of acceptable `ignored-deleted` reason
    strings for the given fixture.

    For each source tag in :data:`audit_reason_vocab.DELETED_SOURCES`:
      * If the template has no placeholder, include the verbatim
        template once.
      * If the template embeds ``{new_path}``, include one rendered
        copy per path in ``new_paths_by_source[source]``, plus the
        ``"<unknown>"`` rendering as a defensive baseline.
    """
    out: set[str] = set()
    for src in _VOCAB.DELETED_SOURCES:
        template = _VOCAB.DELETED_REASON[src]
        if "{new_path}" in template:
            for np in new_paths_by_source.get(src, set()) | {None}:
                out.add(_VOCAB.resolve_deleted_reason(src, new_path=np))
        else:
            out.add(_VOCAB.resolve_deleted_reason(src))
    return out


_FIXTURE_NEW_PATHS = {
    # `--changed-files` payload destinations.
    "changed-files-R-old": {"spec/new1.md", "spec/dest.md"},
    "changed-files-C-old": {"spec/copy.md"},
    # `git diff` payload destinations.
    "diff-R-old": {"spec/diff_new.md"},
    "diff-C-old": {"spec/diff_copy.md"},
}

_VALID_DELETED_REASONS = _expected_deleted_reasons(_FIXTURE_NEW_PATHS)


# ---------------------------------------------------------------------
# JSON surface — `--list-changed-files --json`.
# Strict schema: every row is an object with `path`, `status`,
# `reason` (all strings), no missing keys, no nulls.
# ---------------------------------------------------------------------

class JsonAuditSchema(unittest.TestCase):

    def _load(self, *extra: str) -> list[dict]:
        with _Sandbox() as box:
            proc = box.run_changed_files("--json", *extra)
        self.assertEqual(proc.returncode, 0, msg=proc.stderr)
        # With `--list-changed-files --json`, the placeholder-
        # violation array lands on STDOUT and the changed-file
        # AUDIT JSON lands on STDERR (the audit is the human
        # summary; STDOUT stays reserved for the violations the
        # caller is normally piping into another tool). Parse
        # strictly — a malformed payload is itself a schema
        # violation. The audit JSON is a top-level array starting
        # with `[` on its own line; everything before that is
        # diagnostic text we ignore.
        stderr = proc.stderr
        start = stderr.find("\n[")
        if start == -1 and stderr.lstrip().startswith("["):
            start = stderr.find("[")
        else:
            start = start + 1 if start != -1 else -1
        self.assertNotEqual(
            start, -1,
            f"audit JSON not found on stderr:\n{stderr}",
        )
        # Find the matching closing bracket by scanning to the
        # last `]` on its own line (the array is pretty-printed).
        end = stderr.rfind("\n]")
        self.assertNotEqual(
            end, -1,
            f"audit JSON closing `]` not found on stderr:\n{stderr}",
        )
        data = json.loads(stderr[start:end + 2])
        self.assertIsInstance(data, list)
        return data

    def test_every_row_has_path_status_reason_strings(self) -> None:
        rows = self._load()
        self.assertGreater(len(rows), 0,
                           "fixture must produce at least one row")
        for i, row in enumerate(rows):
            with self.subTest(row_index=i, row=row):
                self.assertIsInstance(row, dict)
                for key in ("path", "status", "reason"):
                    self.assertIn(key, row,
                                  f"row {i} missing required key "
                                  f"`{key}`")
                    self.assertIsInstance(row[key], str,
                                          f"row {i} key `{key}` "
                                          f"must be a string")
                    # Empty strings would technically satisfy
                    # `isinstance(..., str)` but defeat the
                    # contract — operators rely on `reason` being
                    # human-readable, not blank.
                    self.assertNotEqual(row[key].strip(), "",
                                        f"row {i} key `{key}` is "
                                        f"empty / whitespace")

    def test_every_status_is_in_closed_vocabulary(self) -> None:
        rows = self._load()
        for i, row in enumerate(rows):
            with self.subTest(row_index=i, status=row.get("status")):
                self.assertIn(row["status"], _MOD._AUDIT_STATUSES)

    def test_no_unexpected_top_level_keys_in_legacy_mode(self) -> None:
        # Without `--list-changed-files-verbose` and without
        # `--with-similarity` the row schema must be exactly the
        # 3-key shape downstream dashboards close on. A new key
        # leaking in is a backward-compat break.
        rows = self._load()
        for i, row in enumerate(rows):
            with self.subTest(row_index=i):
                self.assertEqual(set(row.keys()),
                                 {"path", "status", "reason"})

    def test_ignored_deleted_reasons_match_vocabulary(self) -> None:
        rows = self._load()
        deleted_rows = [r for r in rows
                        if r["status"] == "ignored-deleted"]
        self.assertGreater(len(deleted_rows), 0,
                           "fixture must produce at least one "
                           "ignored-deleted row")
        for i, row in enumerate(deleted_rows):
            with self.subTest(row_index=i, reason=row["reason"]):
                self.assertIn(
                    row["reason"], _VALID_DELETED_REASONS,
                    f"`ignored-deleted` reason for "
                    f"`{row['path']}` is not in the closed "
                    f"vocabulary produced by "
                    f"`resolve_deleted_reason()` — got "
                    f"{row['reason']!r}",
                )
                # Defensive: a fallback-message leak means a
                # parser is emitting an unknown tag.
                self.assertNotEqual(
                    row["reason"], _VOCAB.DELETED_REASON_FALLBACK,
                    "ignored-deleted row carries the safety-net "
                    "fallback message — a parser must be emitting "
                    "an unknown source tag",
                )

    def test_verbose_mode_source_matches_a_known_vocabulary_tag(
        self,
    ) -> None:
        # With --list-changed-files-verbose the row gains a
        # `source` field. For deleted rows it MUST be in
        # DELETED_SOURCES; for non-deleted rows it must be `None`.
        rows = self._load("--list-changed-files-verbose")
        for i, row in enumerate(rows):
            with self.subTest(row_index=i, row=row):
                self.assertIn("source", row,
                              "verbose mode must include `source`")
                if row["status"] == "ignored-deleted":
                    self.assertIn(row["source"],
                                  _VOCAB.DELETED_SOURCES)
                else:
                    self.assertIsNone(
                        row["source"],
                        f"non-deleted row `{row['path']}` "
                        f"({row['status']}) must have "
                        f"`source: null`",
                    )


# ---------------------------------------------------------------------
# Parser-level coverage for `diff-*` provenance tags.
#
# The shipped CLI only emits `diff-*` tags when it runs `git diff`
# itself, which would force every test to stand up a real
# repository fixture. Instead we drive the same code path that the
# end-to-end JSON / CSV / text tests rely on by feeding a synthetic
# `git diff --name-status -M -C` payload straight into
# `_parse_name_status`, then resolving each captured tuple through
# the shared `resolve_deleted_reason()`. The resulting reason
# strings must satisfy the SAME closed-vocabulary contract the
# end-to-end tests assert on JSON/CSV/text — proving the schema
# guarantee holds independently of which intake produced the
# `ignored-deleted` row.
# ---------------------------------------------------------------------

_DIFF_NAME_STATUS_PAYLOAD = (
    "D\tspec/diff_gone.md\n"
    "R087\tspec/diff_old.md\tspec/diff_new.md\n"
    "C063\tspec/diff_src.md\tspec/diff_copy.md\n"
)


class DiffParserDeletedReasonVocabulary(unittest.TestCase):
    """`diff-*` reasons resolved by the parser must satisfy the
    same closed-vocabulary contract as `changed-files-*` reasons
    resolved by the renderer."""

    def test_diff_d_r_c_old_reasons_are_in_canonical_set(self) -> None:
        deleted: list[tuple[str, str, "str | None"]] = []
        # Drive the same parser the production CLI uses on a real
        # `git diff --name-status` payload. The parser populates
        # `deleted` with `(path, source, new_path | None)` tuples.
        _MOD._parse_name_status(_DIFF_NAME_STATUS_PAYLOAD,
                                deleted=deleted)
        # Pre-condition: at least one tuple per `diff-*` tag the
        # fixture exercises (D, R-old, C-old). Without this, the
        # vocabulary assertion below could pass vacuously.
        observed_sources = {src for _, src, _ in deleted}
        self.assertEqual(observed_sources,
                         {"diff-D", "diff-R-old", "diff-C-old"})

        # Every resolved reason must land in the canonical set,
        # exactly as the JSON/CSV/text schema tests demand on the
        # end-to-end surfaces.
        for path, source, new_path in deleted:
            with self.subTest(path=path, source=source):
                reason = _VOCAB.resolve_deleted_reason(
                    source, new_path=new_path)
                self.assertIn(reason, _VALID_DELETED_REASONS,
                              f"`diff-*` reason for `{path}` "
                              f"({source}) outside vocabulary: "
                              f"{reason!r}")
                self.assertNotEqual(
                    reason, _VOCAB.DELETED_REASON_FALLBACK,
                    "parser must never produce a tag that drops "
                    "to the fallback message")


# ---------------------------------------------------------------------
# CSV surface — `--similarity-csv -` writes to STDOUT in CSV
# dialect. Header is fixed: path,status,reason,kind,score,old_path.
# Same two contracts apply.
# ---------------------------------------------------------------------

class CsvAuditSchema(unittest.TestCase):

    def _load(self) -> list[dict]:
        with _Sandbox() as box:
            proc = box.run_changed_files(
                "--similarity-csv", "-",
                # CSV needs `--with-similarity` to populate the
                # similarity columns, but the `path/status/reason`
                # contract holds either way.
                "--with-similarity",
            )
        self.assertEqual(proc.returncode, 0, msg=proc.stderr)
        # The CSV is appended to STDOUT after the human summary.
        # Find the header row by scanning for the canonical first
        # column name; everything from there to EOF is the CSV.
        lines = proc.stdout.splitlines()
        header_idx = next(
            (i for i, line in enumerate(lines)
             if line.startswith("path,status,reason,")),
            None,
        )
        self.assertIsNotNone(
            header_idx,
            f"CSV header not found in stdout:\n{proc.stdout}",
        )
        csv_text = "\n".join(lines[header_idx:])
        reader = _csv.DictReader(io.StringIO(csv_text))
        rows = list(reader)
        self.assertGreater(len(rows), 0,
                           "CSV must contain at least one data row")
        return rows

    def test_every_csv_row_has_non_empty_reason(self) -> None:
        for i, row in enumerate(self._load()):
            with self.subTest(row_index=i, row=row):
                self.assertIn("reason", row)
                self.assertIsNotNone(row["reason"])
                self.assertNotEqual(row["reason"].strip(), "",
                                    f"CSV row {i} has empty "
                                    f"`reason`")

    def test_csv_ignored_deleted_reasons_match_vocabulary(
        self,
    ) -> None:
        rows = self._load()
        deleted = [r for r in rows
                   if r["status"] == "ignored-deleted"]
        self.assertGreater(len(deleted), 0)
        for i, row in enumerate(deleted):
            with self.subTest(row_index=i, reason=row["reason"]):
                self.assertIn(
                    row["reason"], _VALID_DELETED_REASONS,
                    f"CSV `ignored-deleted` reason for "
                    f"`{row['path']}` outside closed vocabulary",
                )


# ---------------------------------------------------------------------
# Text surface — the human-readable table also surfaces a `reason`
# cell. Parsed loosely by splitting on 2+ spaces; a stricter
# parser would couple the test to column-padding logic that
# already has its own tests elsewhere.
# ---------------------------------------------------------------------

# Match a row whose first column is one of the known statuses,
# capture status / path / reason. The ``reason`` group is greedy
# to EOL because the reason itself can contain backticks, colons
# and spaces — we trust the leading status anchor.
_TEXT_ROW_RE = re.compile(
    r"^\s*"
    r"(?P<status>matched|ignored-extension|ignored-out-of-root|"
    r"ignored-missing|ignored-deleted)\s{2,}"
    r"(?P<path>\S+)\s{2,}"
    r"(?P<reason>.+\S)\s*$"
)


class TextTableAuditSchema(unittest.TestCase):

    def _parse(self) -> list[dict]:
        with _Sandbox() as box:
            proc = box.run_changed_files()
        self.assertEqual(proc.returncode, 0, msg=proc.stderr)
        # The audit table is on STDERR (it's the human summary; the
        # JSON/CSV surfaces are STDOUT). We only need the rows that
        # match the regex — the header / footer don't.
        rows: list[dict] = []
        for line in proc.stderr.splitlines():
            m = _TEXT_ROW_RE.match(line)
            if m:
                rows.append(m.groupdict())
        self.assertGreater(
            len(rows), 0,
            f"no audit rows parsed from stderr:\n{proc.stderr}",
        )
        return rows

    def test_every_text_row_has_a_reason_cell(self) -> None:
        for i, row in enumerate(self._parse()):
            with self.subTest(row_index=i, row=row):
                # The regex already requires a non-empty `reason`
                # group (`.+\S`), so a row that parsed at all
                # already has a reason cell. Defensive re-assert
                # so a future loosening of the regex breaks the
                # test loudly.
                self.assertTrue(row["reason"])
                self.assertNotEqual(row["reason"].strip(), "")

    def test_text_ignored_deleted_reasons_match_vocabulary(
        self,
    ) -> None:
        deleted = [r for r in self._parse()
                   if r["status"] == "ignored-deleted"]
        self.assertGreater(len(deleted), 0)
        for i, row in enumerate(deleted):
            with self.subTest(row_index=i, reason=row["reason"]):
                self.assertIn(
                    row["reason"], _VALID_DELETED_REASONS,
                    f"text-table `ignored-deleted` reason for "
                    f"`{row['path']}` outside closed vocabulary",
                )


# ---------------------------------------------------------------------
# Cross-surface invariant — the set of ignored-deleted reasons
# observed on the JSON surface must be byte-identical to the set
# on CSV and on the text table. A drift between formats would
# silently break operators who pivot between surfaces.
# ---------------------------------------------------------------------

class CrossSurfaceReasonParity(unittest.TestCase):

    def test_json_csv_text_agree_on_ignored_deleted_reasons(
        self,
    ) -> None:
        with _Sandbox() as box:
            json_proc = box.run_changed_files("--json")
            csv_proc = box.run_changed_files(
                "--similarity-csv", "-", "--with-similarity")
            text_proc = box.run_changed_files()
        for proc in (json_proc, csv_proc, text_proc):
            self.assertEqual(proc.returncode, 0, msg=proc.stderr)

        # JSON.
        json_rows = [
            r for r in json.loads(json_proc.stdout)
            if r["status"] == "ignored-deleted"
        ]
        json_reasons = {(r["path"], r["reason"]) for r in json_rows}

        # CSV — same parser as `CsvAuditSchema._load`, inlined to
        # keep this test self-contained.
        csv_lines = csv_proc.stdout.splitlines()
        header_idx = next(
            i for i, line in enumerate(csv_lines)
            if line.startswith("path,status,reason,")
        )
        csv_reader = _csv.DictReader(
            io.StringIO("\n".join(csv_lines[header_idx:])))
        csv_reasons = {
            (r["path"], r["reason"]) for r in csv_reader
            if r["status"] == "ignored-deleted"
        }

        # Text — same regex as `TextTableAuditSchema._parse`.
        text_reasons = set()
        for line in text_proc.stderr.splitlines():
            m = _TEXT_ROW_RE.match(line)
            if m and m.group("status") == "ignored-deleted":
                text_reasons.add((m.group("path"), m.group("reason")))

        self.assertEqual(
            json_reasons, csv_reasons,
            "JSON vs CSV ignored-deleted (path, reason) sets differ",
        )
        self.assertEqual(
            json_reasons, text_reasons,
            "JSON vs text-table ignored-deleted (path, reason) "
            "sets differ",
        )


if __name__ == "__main__":
    unittest.main()
