#!/usr/bin/env python3
"""CODE-RED-010 - Enum suffix. Rust. Flags enums that don't end with _type."""

from __future__ import annotations
import re
import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parents[1]))
from _lib.cli import build_parser, parse_exclude_paths
from _lib.sarif import Finding, Rule, SarifRun, emit
from _lib.walker import relpath, walk_files

RULE = Rule(
    id="CODE-RED-010",
    name="EnumSuffix",
    short_description="Enums must be named with a '_type' suffix.",
    help_uri_relative="17-consolidated-guidelines/04-enum-standards.md",
)
ENUM_RE = re.compile(r'\benum\s+([A-Za-z0-9_]+)\s*(?:\{|$)')

def scan_file(path: Path, root: str) -> list[Finding]:
    text = path.read_text(encoding="utf-8", errors="replace")
    findings = []
    for i, raw in enumerate(text.splitlines(), start=1):
        line = raw.strip()
        m = ENUM_RE.search(line)
        if m:
            name = m.group(1)
            # The user mentioned underscore for Python, and that it could go for Rust too.
            # We enforce `_type` or `Type` depending on case convention, but the user requested underscore.
            if not name.endswith("_type"):
                findings.append(
                    Finding(
                        rule_id=RULE.id,
                        level="error",
                        message=f"Enum '{name}' must end with '_type' suffix.",
                        file_path=relpath(path, root),
                        start_line=i,
                    )
                )
    return findings

def main() -> int:
    args = build_parser("CODE-RED-010 enum-suffix (Rust)").parse_args()
    _globs = parse_exclude_paths(args.exclude_paths)
    run = SarifRun(tool_name="coding-guidelines-enum-suffix-rust", tool_version="1.0.0", rules=[RULE])
    for f in walk_files(args.path, [".rs"], exclude_globs=_globs):
        for finding in scan_file(f, args.path):
            run.add(finding)
    return emit(run, args.format, args.output)

if __name__ == "__main__":
    sys.exit(main())
