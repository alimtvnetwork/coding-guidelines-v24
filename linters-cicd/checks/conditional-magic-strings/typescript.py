#!/usr/bin/env python3
"""CODE-RED-009 - Conditional magic strings. TypeScript. Flags magic strings used in if/else if conditions."""

from __future__ import annotations
import re
import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parents[1]))
from _lib.cli import build_parser, parse_exclude_paths
from _lib.sarif import Finding, Rule, SarifRun, emit
from _lib.walker import relpath, walk_files

RULE = Rule(
    id="CODE-RED-009",
    name="ConditionalMagicStrings",
    short_description="Magic string literals must not be used in conditional (if/else if) comparisons.",
    help_uri_relative="17-consolidated-guidelines/31-compiled-simple-coding-guidelines.md",
)
# Matches `if ` followed by anything with `==` or `!=` or `===` or `!==` and a string literal
CONDITIONAL_MAGIC_RE = re.compile(r'\bif\s*\([^)]*(?:==|!=|===|!==)\s*[\'"`][^\'"`]+[\'"`]|\bif\s*\([^)]*[\'"`][^\'"`]+[\'"`]\s*(?:==|!=|===|!==)')

def scan_file(path: Path, root: str) -> list[Finding]:
    text = path.read_text(encoding="utf-8", errors="replace")
    findings = []
    for i, raw in enumerate(text.splitlines(), start=1):
        line = raw.split("//", 1)[0].strip()
        if CONDITIONAL_MAGIC_RE.search(line):
            findings.append(
                Finding(
                    rule_id=RULE.id,
                    level="error",
                    message="Conditional comparison uses a magic string. Use an enum or typed constant instead.",
                    file_path=relpath(path, root),
                    start_line=i,
                )
            )
    return findings

def main() -> int:
    args = build_parser("CODE-RED-009 conditional-magic-strings (TypeScript)").parse_args()
    _globs = parse_exclude_paths(args.exclude_paths)
    run = SarifRun(tool_name="coding-guidelines-conditional-magic-strings-ts", tool_version="1.0.0", rules=[RULE])
    for f in walk_files(args.path, [".ts", ".tsx"], exclude_globs=_globs):
        for finding in scan_file(f, args.path):
            run.add(finding)
    return emit(run, args.format, args.output)

if __name__ == "__main__":
    sys.exit(main())
