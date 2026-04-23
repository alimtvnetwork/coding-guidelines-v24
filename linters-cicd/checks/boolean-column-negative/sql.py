#!/usr/bin/env python3
"""BOOL-NEG-001 — Forbid Not/No-prefixed boolean DB column names.

Flags any column name matching ``^(Is|Has)(Not|No)[A-Z]`` inside a
``CREATE TABLE`` body. Names like ``IsDisabled``, ``IsInvalid``,
``IsUnverified``, and ``IsUnpublished`` are explicitly **allow-listed**
because they describe positive domain states (see
``spec/04-database-conventions/01-naming-conventions.md`` Rule 2, v3.3.0).

Scope: ``.sql`` files and any file under a path segment named
``migrations``/``migration`` (Go ``embed.FS`` SQL strings are out of
scope for v1 — use a Go-aware plugin if needed).
"""

from __future__ import annotations

import re
import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parents[1]))
from _lib.cli import build_parser
from _lib.sarif import Finding, Rule, SarifRun, emit
from _lib.walker import relpath, walk_files


RULE = Rule(
    id="BOOL-NEG-001",
    name="BooleanColumnNegativePrefix",
    short_description=(
        "Database boolean columns must not use Not/No prefixes "
        "(e.g. IsNotActive, HasNoLicense). Use the positive form or "
        "store one canonical state and derive the inverse in code."
    ),
    help_uri_relative="../04-database-conventions/01-naming-conventions.md",
)

EXTENSIONS = [".sql"]
MIGRATION_HINTS = ("migrations", "migration")

# Match Is/Has + Not/No + UpperCamel (e.g. IsNotActive, HasNoLicense).
NEG_PREFIX_RE = re.compile(r"\b((?:Is|Has)(?:Not|No)[A-Z][A-Za-z0-9]*)\b")

# Allow-listed "Approved Inverse of Positive" names — single-negative roots,
# NOT double negatives. Per Rule 2 + Rule 8 in
# spec/04-database-conventions/01-naming-conventions.md (v3.4.0).
# Note: the NEG_PREFIX_RE above only matches Not/No prefixes, so most of
# these names would never trigger it. They are kept here for auditability
# and for future extensions of the regex (e.g. catching Un-/In-/Dis- roots).
ALLOWLIST = {
    "IsDisabled",
    "IsInvalid",
    "IsIncomplete",
    "IsUnavailable",
    "IsUnread",
    "IsHidden",
    "IsBroken",
    "IsLocked",
    "IsUnpublished",
    "IsUnverified",
}

# Detect CREATE TABLE blocks so we only inspect column definitions.
CREATE_TABLE_RE = re.compile(
    r"CREATE\s+TABLE[^\(]*\((?P<body>.*?)\)\s*;",
    re.IGNORECASE | re.DOTALL,
)


def is_in_scope(path: Path) -> bool:
    if path.suffix.lower() in EXTENSIONS:
        return True
    parts = {p.lower() for p in path.parts}
    return any(hint in parts for hint in MIGRATION_HINTS) and path.suffix.lower() in {".sql"}


def scan(path: Path, root: str) -> list[Finding]:
    text = path.read_text(encoding="utf-8", errors="replace")
    findings: list[Finding] = []
    for block in CREATE_TABLE_RE.finditer(text):
        body = block.group("body")
        body_offset = block.start("body")
        for match in NEG_PREFIX_RE.finditer(body):
            name = match.group(1)
            if name in ALLOWLIST:
                continue
            abs_offset = body_offset + match.start()
            line_no = text.count("\n", 0, abs_offset) + 1
            findings.append(
                Finding(
                    rule_id=RULE.id,
                    level="error",
                    message=(
                        f"Boolean column '{name}' uses a forbidden Not/No prefix. "
                        "Rename to the positive form (e.g. IsActive, HasLicense) "
                        "and derive the inverse as a computed field in code. "
                        "See Rule 2 + Rule 9 in 04-database-conventions/01-naming-conventions.md."
                    ),
                    file_path=relpath(path, root),
                    start_line=line_no,
                )
            )
    return findings


def main() -> int:
    args = build_parser("BOOL-NEG-001 boolean-column-negative (sql)").parse_args()
    run = SarifRun(
        tool_name="coding-guidelines-boolean-column-negative",
        tool_version="1.0.0",
        rules=[RULE],
    )
    for f in walk_files(args.path, EXTENSIONS):
        if not is_in_scope(f):
            continue
        for finding in scan(f, args.path):
            run.add(finding)
    return emit(run, args.format, args.output)


if __name__ == "__main__":
    sys.exit(main())
