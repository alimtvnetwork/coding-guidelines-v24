#!/usr/bin/env python3
"""CODE-RED-013 - Inverted Complex Conditions. Go."""

from __future__ import annotations
import re
import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parents[1]))
from _lib.cli import build_parser, parse_exclude_paths
from _lib.sarif import Finding, Rule, SarifRun, emit
from _lib.walker import relpath, walk_files

RULE = Rule(
    id="CODE-RED-013",
    name="InvertedComplexConditions",
    short_description="Do not use NOT (!) on complex conditions. Assign to a boolean variable or use De Morgan's laws.",
    help_uri_relative="17-consolidated-guidelines/31-compiled-simple-coding-guidelines.md",
)
INVERTED_COMPLEX_RE = re.compile(r"if\s+!\s*\([^)]+(&&|\|\|)[^)]+\)")

def scan(path: Path, root: str) -> list[Finding]:
    findings: list[Finding] = []
    text = path.read_text(encoding="utf-8", errors="replace")
    for i, raw in enumerate(text.splitlines(), start=1):
        line = raw.split("//", 1)[0]
        if INVERTED_COMPLEX_RE.search(line):
            findings.append(
                Finding(
                    rule_id=RULE.id,
                    level="error",
                    message="Inverted complex condition found. Simplify logic or assign to a well-named boolean.",
                    file_path=relpath(path, root),
                    start_line=i,
                )
            )
    return findings

def main() -> int:
    args = build_parser("CODE-RED-013 inverted-complex-conditions (Go)").parse_args()
    _globs = parse_exclude_paths(args.exclude_paths)
    run = SarifRun(tool_name="coding-guidelines-invert-cond-go", tool_version="1.0.0", rules=[RULE])
    for f in walk_files(args.path, [".go"], exclude_globs=_globs):
        for finding in scan(f, args.path):
            run.add(finding)
    return emit(run, args.format, args.output)

if __name__ == "__main__":
    main()
