#!/usr/bin/env python3
"""CODE-RED-001 - No nested `if`. Python implementation (AST-based)."""

from __future__ import annotations
import ast
import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).resolve().parents[1]))
from _lib.cli import build_parser, parse_exclude_paths
from _lib.sarif import Finding, Rule, SarifRun, emit
from _lib.walker import relpath, walk_files

RULE = Rule(
    id="CODE-RED-001",
    name="NoNestedIf",
    short_description="Nested if statements are forbidden - flatten with guard clauses.",
    help_uri_relative="01-cross-language/04-code-style/00-overview.md",
)

class IfVisitor(ast.NodeVisitor):
    def __init__(self, path: Path, root: str):
        self.path = path
        self.root = root
        self.depth = 0
        self.findings: list[Finding] = []

    def visit_If(self, node: ast.If) -> None:
        self.depth += 1
        if self.depth >= 2:
            self.findings.append(
                Finding(
                    rule_id=RULE.id,
                    level="error",
                    message=f"Nested if at depth {self.depth} - extract guard clauses.",
                    file_path=relpath(self.path, self.root),
                    start_line=node.lineno,
                )
            )

        for child in node.body:
            self.visit(child)

        for child in node.orelse:
            if isinstance(child, ast.If):
                # This is an elif, do not increase depth
                self.depth -= 1
                self.visit(child)
                self.depth += 1
            else:
                self.visit(child)

        self.depth -= 1

def scan(path: Path, root: str) -> list[Finding]:
    text = path.read_text(encoding="utf-8", errors="replace")
    try:
        tree = ast.parse(text, filename=path.name)
    except SyntaxError:
        return []

    visitor = IfVisitor(path, root)
    visitor.visit(tree)

    # Filter out elifs that are naturally nested in the AST
    # Wait, simple ast approach: 'elif' is tricky because it's stored in orelse.
    # To fix 'elif' depth:
    return visitor.findings

def main() -> int:
    args = build_parser("CODE-RED-001 nested-if (Python)").parse_args()
    _globs = parse_exclude_paths(args.exclude_paths)
    run = SarifRun(tool_name="coding-guidelines-nested-if-py", tool_version="1.0.0", rules=[RULE])
    for f in walk_files(args.path, [".py"], exclude_globs=_globs):
        for finding in scan(f, args.path):
            run.add(finding)
    return emit(run, args.format, args.output)

if __name__ == "__main__":
    main()
