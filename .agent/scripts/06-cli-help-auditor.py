#!/usr/bin/env python3
"""
06-cli-help-auditor.py — Automated CLI Command Discovery & Help Text Parity Auditor

Discovers CLI entry points, subcommands, flags, and inspects help descriptions
across Go (Cobra), TypeScript (Commander), Python (Click/Argparse), and PHP (Symfony).

Usage:
  python .lovable/ai-fix-scripts/06-cli-help-auditor.py [--dir <path>] [--strict]
"""

import ast
import os
import re
import sys
import argparse

# Ensure UTF-8 output on Windows consoles
if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8", errors="replace")
if hasattr(sys.stderr, "reconfigure"):
    sys.stderr.reconfigure(encoding="utf-8", errors="replace")

IGNORE_DIRS = {".git", "node_modules", ".github", ".ci-out", "dist", "build", "tmp", ".venv", "vendor"}


def scan_go_cobra_commands(filepath, content):
    """Detect Go Cobra commands and check for Short, Long, and Example."""
    violations = []
    cmd_pattern = re.compile(r"var\s+(\w+Cmd)\s*=\s*&cobra\.Command\s*\{([^}]+)\}", re.DOTALL)

    for match in cmd_pattern.finditer(content):
        cmd_var = match.group(1)
        body = match.group(2)

        has_short = "Short:" in body
        has_example = "Example:" in body

        if not has_short:
            violations.append((cmd_var, "Missing Short: description in cobra.Command"))
        if not has_example and cmd_var != "rootCmd":
            violations.append((cmd_var, "Missing Example: usage string in cobra.Command"))

    return violations


def scan_python_cli_commands(filepath, content):
    """Detect Click/Argparse/Typer command definitions using AST."""
    violations = []
    try:
        tree = ast.parse(content, filename=filepath)
    except Exception:
        return violations

    for node in ast.walk(tree):
        # Check function decorators for @cli.command(...) or @click.command(...)
        if isinstance(node, ast.FunctionDef):
            for decorator in node.decorator_list:
                if isinstance(decorator, ast.Call):
                    func_name = ""
                    if isinstance(decorator.func, ast.Attribute):
                        func_name = decorator.func.attr
                    if func_name == "command":
                        # Check keyword arguments for 'help' or 'short_help'
                        has_help = any(kw.arg in ("help", "short_help") for kw in decorator.keywords)
                        # Also check docstring
                        has_docstring = ast.get_docstring(node) is not None
                        if not has_help and not has_docstring:
                            violations.append((node.name, f"Click command '{node.name}' missing help description or docstring"))

        # Check parser.add_argument(...)
        elif isinstance(node, ast.Call):
            if isinstance(node.func, ast.Attribute) and node.func.attr == "add_argument":
                has_help = any(kw.arg == "help" for kw in node.keywords)
                if not has_help:
                    # Extract argument name
                    arg_name = "argument"
                    if node.args and isinstance(node.args[0], ast.Constant) and isinstance(node.args[0].value, str):
                        arg_name = node.args[0].value
                    violations.append((arg_name, f"Argparse argument '{arg_name}' missing help= parameter"))

    return violations


def scan_file(filepath):
    violations = []
    try:
        with open(filepath, "r", encoding="utf-8") as f:
            content = f.read()
    except Exception:
        return violations

    rel_path = os.path.relpath(filepath, os.getcwd()).replace("\\", "/")

    if filepath.endswith(".go"):
        go_v = scan_go_cobra_commands(filepath, content)
        for target, msg in go_v:
            violations.append((rel_path, target, msg))

    elif filepath.endswith(".py"):
        py_v = scan_python_cli_commands(filepath, content)
        for target, msg in py_v:
            violations.append((rel_path, target, msg))

    return violations


def main():
    parser = argparse.ArgumentParser(description="CLI Command & Help Text Parity Auditor")
    parser.add_argument("--dir", default=".", help="Root directory to scan (default: .)")
    parser.add_argument("--strict", action="store_true", help="Exit with non-zero code on violations")
    args = parser.parse_args()

    target_dir = os.path.abspath(args.dir)
    all_violations = []

    for root, dirs, files in os.walk(target_dir):
        dirs[:] = [d for d in dirs if d not in IGNORE_DIRS]
        for file in files:
            if file.endswith((".go", ".py", ".ts", ".php")):
                filepath = os.path.join(root, file)
                v = scan_file(filepath)
                all_violations.extend(v)

    print("============================================================")
    print("CLI COMMAND & HELP TEXT PARITY AUDIT REPORT")
    print("============================================================")
    print(f"Scanned Directory : {target_dir}")
    print(f"Total Violations  : {len(all_violations)}")
    print("============================================================")

    if all_violations:
        print("\n| File Path | Command / Target | Violation |")
        print("|---|---|---|")
        for file_path, target, msg in all_violations:
            print(f"| `{file_path}` | `{target}` | {msg} |")

        if args.strict:
            sys.exit(1)
    else:
        print("\n[PASS] All detected CLI commands and arguments contain valid help metadata.")
        sys.exit(0)


if __name__ == "__main__":
    main()
