#!/usr/bin/env python3
"""CODE-RED-012 - Boolean Return Wrapper. Go."""

from __future__ import annotations
import re
import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parents[1]))
from _lib.cli import build_parser, parse_exclude_paths
from _lib.sarif import Finding, Rule, SarifRun, emit
from _lib.walker import relpath, walk_files

RULE = Rule(
    id="CODE-RED-012",
    name="BooleanReturnWrapper",
    short_description="Do not return raw booleans in multi-return signatures; use a wrapper struct.",
    help_uri_relative="17-consolidated-guidelines/31-compiled-simple-coding-guidelines.md",
)
# Matches `func Name(...) (..., bool, ...)` or `func Name(...) (bool, ...)`
FUNC_MULTI_RET_BOOL_RE = re.compile(r"func\s+(?:[A-Za-z0-9_]+\s*)?(?:\([^)]*\)\s*)?[A-Za-z0-9_]+\s*\([^)]*\)\s*\([^)]*?(?:,[^)]*\bbool\b|\bbool\b[^)]*,)[^)]*?\)\s*\{")

def scan(path: Path, root: str) -> list[Finding]:
    findings: list[Finding] = []
    text = path.read_text(encoding="utf-8", errors="replace")
    for i, raw in enumerate(text.splitlines(), start=1):
        line = raw.split("//", 1)[0]
        if FUNC_MULTI_RET_BOOL_RE.search(line):
            findings.append(
                Finding(
                    rule_id=RULE.id,
                    level="error",
                    message="Multi-return contains a raw boolean. Use a wrapper struct for clarity.",
                    file_path=relpath(path, root),
                    start_line=i,
                )
            )
    return findings

def main() -> int:
    args = build_parser("CODE-RED-012 boolean-return-wrapper (Go)").parse_args()
    _globs = parse_exclude_paths(args.exclude_paths)
    run = SarifRun(tool_name="coding-guidelines-bool-ret-go", tool_version="1.0.0", rules=[RULE])
    for f in walk_files(args.path, [".go"], exclude_globs=_globs):
        for finding in scan(f, args.path):
            run.add(finding)
    return emit(run, args.format, args.output)

if __name__ == "__main__":
    main()
