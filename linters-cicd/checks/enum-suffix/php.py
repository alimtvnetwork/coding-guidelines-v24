#!/usr/bin/env python3
"""CODE-RED-010 - Enum suffix. PHP."""

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
    short_description="Enums must be named with a 'Type' suffix (e.g., StatusType).",
    help_uri_relative="17-consolidated-guidelines/04-enum-standards.md",
)
ENUM_RE = re.compile(r"^\s*enum\s+([A-Za-z0-9_]+)\b")

def scan(path: Path, root: str) -> list[Finding]:
    findings: list[Finding] = []
    text = path.read_text(encoding="utf-8", errors="replace")
    for i, raw in enumerate(text.splitlines(), start=1):
        line = raw.split("//", 1)[0]
        match = ENUM_RE.search(line)
        if match:
            enum_name = match.group(1)
            # Check if ends with Type or _type or _TYPE
            if not enum_name.endswith("Type") and not enum_name.lower().endswith("_type"):
                findings.append(
                    Finding(
                        rule_id=RULE.id,
                        level="error",
                        message=f"Enum '{enum_name}' must end with 'Type' (e.g., '{enum_name}Type').",
                        file_path=relpath(path, root),
                        start_line=i,
                    )
                )
    return findings

def main() -> int:
    args = build_parser("CODE-RED-010 enum-suffix (PHP)").parse_args()
    _globs = parse_exclude_paths(args.exclude_paths)
    run = SarifRun(tool_name="coding-guidelines-enum-suffix-php", tool_version="1.0.0", rules=[RULE])
    for f in walk_files(args.path, [".php"], exclude_globs=_globs):
        for finding in scan(f, args.path):
            run.add(finding)
    return emit(run, args.format, args.output)

if __name__ == "__main__":
    main()
