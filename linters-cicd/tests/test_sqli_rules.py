#!/usr/bin/env python3
"""Tests for SQLI-RAW-001, SQLI-RAW-002, and SQLI-ORDER-001.

Locks behaviour for the three new injection-vector linters introduced
to back the WP-plugin micro-ORM gap analysis:

- SQLI-RAW-001 — rawExecute() concatenation / interpolation
- SQLI-RAW-002 — whereRaw() without strict ? / :param placeholders
- SQLI-ORDER-001 — orderBy/groupBy with non-literal identifiers
"""
from __future__ import annotations

import sys
import tempfile
import unittest
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
sys.path.insert(0, str(ROOT / "checks"))

# rawExecute
sys.path.insert(0, str(ROOT / "checks" / "sqli-raw-execute"))
from _shared import is_unsafe_first_arg as raw_unsafe  # noqa: E402

# whereRaw
sys.path.insert(0, str(ROOT / "checks" / "sqli-where-raw"))
from _shared import (  # noqa: E402
    diagnose_where_raw, has_placeholders, second_arg_present,
)

# orderBy / groupBy
sys.path.insert(0, str(ROOT / "checks" / "sqli-order-group-by"))
from _shared import is_safe_identifier_arg  # noqa: E402


class TestRawExecuteUnsafeArg(unittest.TestCase):
    def test_php_concatenation_flagged(self) -> None:
        self.assertIsNotNone(raw_unsafe('"SELECT * FROM x WHERE id=" . $id'))

    def test_php_interpolation_flagged(self) -> None:
        self.assertIsNotNone(raw_unsafe('"SELECT * FROM x WHERE id=$id"'))

    def test_ts_template_literal_flagged(self) -> None:
        self.assertIsNotNone(raw_unsafe('`SELECT * FROM x WHERE id=${id}`'))

    def test_ts_concat_flagged(self) -> None:
        self.assertIsNotNone(raw_unsafe('"SELECT * FROM x WHERE id=" + id'))

    def test_sprintf_flagged(self) -> None:
        self.assertIsNotNone(raw_unsafe('sprintf("SELECT * FROM %s", $t)'))

    def test_safe_literal_passes(self) -> None:
        self.assertIsNone(raw_unsafe('"SELECT * FROM x WHERE id = :id"'))
        self.assertIsNone(raw_unsafe("'SELECT * FROM x WHERE id = ?'"))


class TestWhereRawDiagnosis(unittest.TestCase):
    def test_interp_is_error(self) -> None:
        reason, level = diagnose_where_raw('"status = $status"')
        self.assertEqual(level, "error")
        self.assertIn("interpolation", reason)

    def test_concat_is_error(self) -> None:
        reason, level = diagnose_where_raw('"status = " . $status')
        self.assertEqual(level, "error")

    def test_placeholder_literal_is_clean(self) -> None:
        reason, _ = diagnose_where_raw("'status = :status'")
        self.assertIsNone(reason)

    def test_has_placeholders_detection(self) -> None:
        self.assertTrue(has_placeholders("'a = ? and b = :name'"))
        self.assertFalse(has_placeholders("'a = b'"))

    def test_second_arg_detection(self) -> None:
        # ->whereRaw('a = ?', [$x])  → second arg present
        self.assertTrue(second_arg_present(", [$x])", 0))
        self.assertFalse(second_arg_present(")", 0))
        self.assertFalse(second_arg_present(", [])", 0))


class TestOrderByIdentifierSafety(unittest.TestCase):
    def test_string_literal_safe(self) -> None:
        for arg in ("'CreatedAt'", '"CreatedAt"', "`CreatedAt`", "'users.id'"):
            with self.subTest(arg=arg):
                self.assertTrue(is_safe_identifier_arg(arg))

    def test_allowlist_lookup_safe(self) -> None:
        self.assertTrue(is_safe_identifier_arg("ALLOWED_COLUMNS['sort']"))
        self.assertTrue(is_safe_identifier_arg("$allowed[$key]"))
        self.assertTrue(is_safe_identifier_arg("COLS.createdAt"))

    def test_bare_variable_unsafe(self) -> None:
        for arg in ("$_GET['sort']", "req.query.sort", "$sort", "userInput"):
            with self.subTest(arg=arg):
                self.assertFalse(is_safe_identifier_arg(arg))

    def test_concat_unsafe(self) -> None:
        self.assertFalse(is_safe_identifier_arg('"users." . $col'))


class TestEndToEndPHPFixture(unittest.TestCase):
    """Run each php.py against a tiny fixture and check exit codes."""

    def _run(self, check_path: Path, source: str) -> int:
        import subprocess
        with tempfile.TemporaryDirectory() as td:
            (Path(td) / "demo.php").write_text(source, encoding="utf-8")
            out = Path(td) / "out.sarif"
            rc = subprocess.call([
                sys.executable, str(check_path),
                "--path", td, "--format", "sarif", "--output", str(out),
            ])
            return rc

    def test_raw_execute_flags_concat(self) -> None:
        check = ROOT / "checks" / "sqli-raw-execute" / "php.py"
        rc = self._run(check, '<?php $r = Orm::rawExecute("SELECT * FROM x WHERE id=" . $id);')
        self.assertEqual(rc, 1)

    def test_where_raw_flags_interp(self) -> None:
        check = ROOT / "checks" / "sqli-where-raw" / "php.py"
        rc = self._run(check, '<?php $q->whereRaw("status = $status");')
        self.assertEqual(rc, 1)

    def test_order_by_flags_variable(self) -> None:
        check = ROOT / "checks" / "sqli-order-group-by" / "php.py"
        rc = self._run(check, '<?php $q->orderBy($_GET["sort"], "asc");')
        self.assertEqual(rc, 1)

    def test_clean_code_passes(self) -> None:
        clean = (
            "<?php\n"
            "$rows = Orm::rawExecute('SELECT * FROM x WHERE id = :id', [':id' => $id]);\n"
            "$q->whereRaw('status = ?', [$status]);\n"
            "$q->orderBy('CreatedAt', 'desc');\n"
        )
        for sub in ("sqli-raw-execute", "sqli-where-raw", "sqli-order-group-by"):
            with self.subTest(sub=sub):
                rc = self._run(ROOT / "checks" / sub / "php.py", clean)
                self.assertEqual(rc, 0)


if __name__ == "__main__":
    unittest.main()
