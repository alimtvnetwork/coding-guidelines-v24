#!/usr/bin/env python3
"""CODE-RED-012 - Boolean Return Wrapper. Python."""

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
    short_description="Do not return raw booleans in tuple multi-returns; use a dataclass or TypedDict.",
    help_uri_relative="17-consolidated-guidelines/31-compiled-simple-coding-guidelines.md",
)
# Matches `def name(...) -> tuple[..., bool, ...]:` or `Tuple[..., bool, ...]`
FUNC_MULTI_RET_BOOL_RE = re.compile(r"def\s+[A-Za-z0-9_]+\s*\([^)]*\)\s*->\s*(?:[Tt]uple)\s*\[[^\]]*?(?:,[^\]]*\bbool\b|\bbool\b[^\]]*,)[^\]]*?\]\s*:")

def scan(path: Path, root: str) -> list[Finding]:
    findings: list[Finding] = []
    text = path.read_text(encoding="utf-8", errors="replace")
    for i, raw in enumerate(text.splitlines(), start=1):
        line = raw.split("#", 1)[0]
        if FUNC_MULTI_RET_BOOL_RE.search(line):
            findings.append(
                Finding(
                    rule_id=RULE.id,
                    level="error",
                    message="Tuple multi-return contains a raw boolean. Use a wrapper class for clarity.",
                    file_path=relpath(path, root),
                    start_line=i,
                )
            )
    return findings

def main() -> int:
    args = build_parser("CODE-RED-012 boolean-return-wrapper (Python)").parse_args()
    _globs = parse_exclude_paths(args.exclude_paths)
    run = SarifRun(tool_name="coding-guidelines-bool-ret-py", tool_version="1.0.0", rules=[RULE])
    for f in walk_files(args.path, [".py"], exclude_globs=_globs):
        for finding in scan(f, args.path):
            run.add(finding)
    return emit(run, args.format, args.output)

if __name__ == "__main__":
    main()
