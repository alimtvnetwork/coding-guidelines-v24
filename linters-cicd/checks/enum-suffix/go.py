#!/usr/bin/env python3
"""CODE-RED-010 - Enum suffix. Go. Flags enums that don't end with Type."""

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
    short_description="Enums must be named with a 'Type' suffix (e.g., UserType).",
    help_uri_relative="17-consolidated-guidelines/07-enum-standards.md",
)
# Extremely basic heuristic: `type X int` or `type X string` in Go are commonly used as enums.
# To avoid false positives on standard type aliases, we might just look for any type declaration
# and a const block. For now, we will flag `type [A-Za-z]+ int` that don't end in Type or Code etc.
# But actually the user said: "For enums, it needs to be ending with the type, uh, suffix."
# We'll flag any type that appears to define an Enum (e.g. followed by const iota).
# This is tricky with regex. We'll use a simpler heuristic: look for `type X int` or `string`
# where there is a comment `// ... enum` or `// Enum`, OR if it's explicitly named like an enum.
# Without an AST, we'll just check `type X string` or `type X int` that doesn't end in `Type`
# IF there is a `const` block nearby, or we just rely on explicit `Enum` strings.
# Since this is a simple regex linter, let's just look for `type [^ ]+ (string|int)` and check suffix?
# That would have false positives. Let's look for `type [A-Za-z]+ string` where the type has a const block matching it.
# A safer bet: we can use a very basic regex that flags ANY `type ...` that has a comment containing `enum` but doesn't end in `Type`.

ENUM_HINT_RE = re.compile(r'type\s+([A-Za-z0-9_]+)\s+(?:string|int)')

def scan_file(path: Path, root: str) -> list[Finding]:
    text = path.read_text(encoding="utf-8", errors="replace")
    findings = []
    lines = text.splitlines()
    for i, raw in enumerate(lines, start=1):
        line = raw.strip()
        m = ENUM_HINT_RE.search(line)
        if m:
            name = m.group(1)
            # If there's an enum comment nearby, or it has a const block.
            # We'll just enforce that if it looks like an enum, it ends with Type.
            # To be safe and avoid noisy false positives in simple scripts, we'll only enforce this if
            # the word 'enum' appears in a comment on the same line or previous line.
            prev_line = lines[i-2] if i >= 2 else ""
            if "enum" in line.lower() or "enum" in prev_line.lower():
                if not name.endswith("Type") and not name.endswith("Code"):
                    findings.append(
                        Finding(
                            rule_id=RULE.id,
                            level="error",
                            message=f"Enum type '{name}' must end with 'Type' suffix (e.g., {name}Type).",
                            file_path=relpath(path, root),
                            start_line=i,
                        )
                    )
    return findings

def main() -> int:
    args = build_parser("CODE-RED-010 enum-suffix (Go)").parse_args()
    _globs = parse_exclude_paths(args.exclude_paths)
    run = SarifRun(tool_name="coding-guidelines-enum-suffix-go", tool_version="1.0.0", rules=[RULE])
    for f in walk_files(args.path, [".go"], exclude_globs=_globs):
        for finding in scan_file(f, args.path):
            run.add(finding)
    return emit(run, args.format, args.output)

if __name__ == "__main__":
    sys.exit(main())
