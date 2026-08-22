#!/usr/bin/env python3
"""CODE-RED-014 - Strict Boolean Conditions. PHP."""

from __future__ import annotations
import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parents[1]))
from _lib.cli import build_parser, parse_exclude_paths
from _lib.sarif import Finding, Rule, SarifRun, emit
from _lib.walker import relpath, walk_files

RULE = Rule(
    id="CODE-RED-014",
    name="StrictBooleanConditions",
    short_description="No mixed logical operators, max one join, no mixed polarity.",
    help_uri_relative="01-cross-language/02-boolean-principles/00-overview.md",
)

def scan(path: Path, root: str) -> list[Finding]:
    findings: list[Finding] = []
    text = path.read_text(encoding="utf-8", errors="replace")
    for i, raw in enumerate(text.splitlines(), start=1):
        line = raw.split("//", 1)[0]
        if "{" in line and "if " in line:
            start_idx = line.find("if ") + 2
            end_idx = line.rfind("{")
            cond = line[start_idx:end_idx].strip()
            if cond.startswith("("): cond = cond[1:]
            if cond.endswith(")"): cond = cond[:-1]
            
            and_count = cond.count("&&") + cond.count(" and ")
            or_count = cond.count("||") + cond.count(" or ")
            
            if and_count + or_count > 1:
                findings.append(Finding(rule_id=RULE.id, level="error", message="Max one conditional join allowed (e.g. no A && B && C or A && B || C).", file_path=relpath(path, root), start_line=i))
            elif and_count + or_count == 1:
                op = "&&" if "&&" in cond else ("||" if "||" in cond else (" and " if " and " in cond else " or "))
                left, right = cond.split(op, 1)
                left_neg = "!" in left
                right_neg = "!" in right
                if left_neg != right_neg:
                    findings.append(Finding(rule_id=RULE.id, level="error", message="Mixed polarity in condition. Both sides must be positive.", file_path=relpath(path, root), start_line=i))
    return findings

def main() -> int:
    args = build_parser("CODE-RED-014 strict-boolean-conditions (PHP)").parse_args()
    _globs = parse_exclude_paths(args.exclude_paths)
    run = SarifRun(tool_name="coding-guidelines-strict-bool-php", tool_version="1.0.0", rules=[RULE])
    for f in walk_files(args.path, [".php"], exclude_globs=_globs):
        for finding in scan(f, args.path):
            run.add(finding)
    return emit(run, args.format, args.output)

if __name__ == "__main__":
    main()
