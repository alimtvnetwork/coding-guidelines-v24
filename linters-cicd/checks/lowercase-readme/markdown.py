#!/usr/bin/env python3
"""CODE-RED-011 - Lowercase Readme. Markdown. Flags any readme files that are not strictly lowercase."""

from __future__ import annotations
import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parents[1]))
from _lib.cli import build_parser, parse_exclude_paths
from _lib.sarif import Finding, Rule, SarifRun, emit
from _lib.walker import relpath, walk_files

RULE = Rule(
    id="CODE-RED-011",
    name="LowercaseReadme",
    short_description="All readme files (root and subdirectories) must be named exactly 'readme.md'.",
    help_uri_relative="01-cross-language/08-file-folder-naming/00-overview.md",
)

def scan_file(path: Path, root: str) -> list[Finding]:
    findings = []
    # Check if this is a readme file
    if path.name.lower() == "readme.md":
        # Check if it deviates from strict lowercase
        if path.name != "readme.md":
            findings.append(
                Finding(
                    rule_id=RULE.id,
                    level="error",
                    message=f"Readme file '{path.name}' must be strictly lowercase 'readme.md'.",
                    file_path=relpath(path, root),
                    start_line=1,
                )
            )
    return findings

def main() -> int:
    args = build_parser("CODE-RED-011 lowercase-readme (Markdown)").parse_args()
    _globs = parse_exclude_paths(args.exclude_paths)
    run = SarifRun(tool_name="coding-guidelines-lowercase-readme-md", tool_version="1.0.0", rules=[RULE])
    for f in walk_files(args.path, [".md"], exclude_globs=_globs):
        for finding in scan_file(f, args.path):
            run.add(finding)
    return emit(run, args.format, args.output)

if __name__ == "__main__":
    sys.exit(main())
