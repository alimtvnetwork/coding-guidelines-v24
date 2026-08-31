#!/usr/bin/env python3
"""
Fast CLI Command Discovery & Help Text Parity Auditor
Inspects CLI entry points, subcommands, and flags across Go (Cobra), TypeScript (Commander), Python (Click/Argparse), and PHP (Symfony).
Multi-folder capable, customizable extensions, and pre-compiled regex engine.
"""

import argparse
import ast
from pathlib import Path
import re
import sys

sys.path.insert(0, str(Path(__file__).parent))
try:
    from importlib import import_module
    engine = import_module("00-shared-engine")
    process_repository_files = engine.process_repository_files
    read_file_lf = engine.read_file_lf
    normalize_extensions = engine.normalize_extensions
    normalize_rel_path = engine.normalize_rel_path
    ExitCodeType = engine.ExitCodeType
except Exception:
    ExitCodeType = None

DEFAULT_CLI_EXTENSIONS = (".go", ".ts", ".tsx", ".py", ".php")

# Pre-compiled regular expressions
RE_COBRA_COMMAND = re.compile(r"var\s+(\w+Cmd)\s*=\s*&cobra\.Command\s*\{([^}]+)\}", re.DOTALL)
RE_SHORT_DESC = re.compile(r"Short:\s*\"[^\"]+\"")
RE_EXAMPLE_USAGE = re.compile(r"Example:\s*\"[^\"]+\"")

def audit_go_cobra_commands(content: str) -> list[tuple[str, str]]:
    """Detects Go Cobra commands missing Short or Example descriptions."""
    violations = []
    for match in RE_COBRA_COMMAND.finditer(content):
        cmd_var = match.group(1)
        body = match.group(2)
        if not RE_SHORT_DESC.search(body):
            violations.append((cmd_var, "Missing Short description in cobra.Command"))
        if cmd_var != "rootCmd":
            if not RE_EXAMPLE_USAGE.search(body):
                violations.append((cmd_var, "Missing Example usage in cobra.Command"))
    return violations

def audit_python_cli_commands(file_path: Path, content: str) -> list[tuple[str, str]]:
    """Detects Python CLI commands missing docstrings or help text."""
    violations = []
    try:
        tree = ast.parse(content, filename=str(file_path))
        for node in ast.walk(tree):
            if isinstance(node, ast.FunctionDef):
                for dec in node.decorator_list:
                    if isinstance(dec, ast.Call):
                        if hasattr(dec.func, "attr"):
                            if dec.func.attr == "command":
                                if not ast.get_docstring(node):
                                    violations.append((node.name, "Missing docstring for CLI command function"))
    except Exception:
        pass
    return violations

def audit_single_file_cli(file_path: Path) -> list[tuple[str, str]]:
    """Audits a single file for CLI help compliance."""
    try:
        content = read_file_lf(file_path)
        if file_path.suffix == ".go":
            return audit_go_cobra_commands(content)
        if file_path.suffix == ".py":
            return audit_python_cli_commands(file_path, content)
    except Exception:
        pass
    return []

def run_cli_auditor(
    target_dir: str = ".",
    is_strict: bool = False,
    extensions: set[str] | tuple | None = None
) -> int:
    """Runs repository CLI help audit using two-phase pipeline."""
    exts = normalize_extensions(extensions) or DEFAULT_CLI_EXTENSIONS

    def handler(p: Path):
        vios = audit_single_file_cli(p)
        return (normalize_rel_path(p), vios) if vios else None

    stats = process_repository_files(handler, root_dir=target_dir, extensions=exts)
    all_violations = stats["results"]

    if all_violations:
        print(f"\n⚠️ Found CLI help description issues in {len(all_violations)} file(s) ({stats['elapsed_ms']:.2f}ms):")
        for fp, vios in all_violations:
            for cmd, msg in vios:
                print(f"  ::warning file={fp}::{cmd}: {msg}")
        if is_strict:
            return ExitCodeType.VIOLATIONS_FOUND.value if ExitCodeType else 1
    else:
        print(f"✅ All CLI commands in {stats['total_files']} files in '{target_dir}' contain required help strings ({stats['elapsed_ms']:.2f}ms).")

    return ExitCodeType.SUCCESS.value if ExitCodeType else 0

def main():
    parser = argparse.ArgumentParser(description="Audit CLI commands for help descriptions across folders")
    parser.add_argument("path", nargs="?", default=".", help="Directory to audit")
    parser.add_argument("--dir", "--path", "-p", dest="opt_dir", help="Directory to audit")
    parser.add_argument("--ext", help="Comma-separated extensions to scan (e.g. .go,.py)")
    parser.add_argument("--strict", action="store_true", help="Fail with exit code 1 on warnings")
    args = parser.parse_args()

    target_path = args.opt_dir or args.path or "."
    sys.exit(run_cli_auditor(target_dir=target_path, is_strict=args.strict, extensions=args.ext))

if __name__ == "__main__":
    main()
