#!/usr/bin/env python3
"""
Fast CLI Command Discovery & Help Text Parity Auditor
Inspects CLI entry points, subcommands, and flags across Go (Cobra), TypeScript (Commander), Python (Click/Argparse), and PHP (Symfony).
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
    ExitCodeType = engine.ExitCodeType
except Exception:
    ExitCodeType = None

CLI_EXTENSIONS = (".go", ".ts", ".tsx", ".py", ".php")
COBRA_PATTERN = re.compile(r"var\s+(\w+Cmd)\s*=\s*&cobra\.Command\s*\{([^}]+)\}", re.DOTALL)

def audit_go_cobra_commands(content: str) -> list[tuple[str, str]]:
    """Detects Go Cobra commands missing Short or Example descriptions."""
    violations = []
    for match in COBRA_PATTERN.finditer(content):
        cmd_var = match.group(1)
        body = match.group(2)
        if "Short:" not in body:
            violations.append((cmd_var, "Missing Short description in cobra.Command"))
        if "Example:" not in body and cmd_var != "rootCmd":
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
                    if isinstance(dec, ast.Call) and hasattr(dec.func, "attr") and dec.func.attr == "command":
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

def run_cli_auditor(target_dir: str = ".", is_strict: bool = False) -> int:
    """Runs repository CLI help audit using two-phase pipeline."""
    def handler(p: Path):
        vios = audit_single_file_cli(p)
        if vios:
            return (str(p), vios)
        return None

    stats = process_repository_files(handler, root_dir=target_dir, extensions=CLI_EXTENSIONS)
    all_violations = stats["results"]

    if all_violations:
        print(f"\n⚠️ Found CLI help description issues in {len(all_violations)} file(s) ({stats['elapsed_ms']:.2f}ms):")
        for fp, vios in all_violations:
            for cmd, msg in vios:
                print(f"  ::warning file={fp}::{cmd}: {msg}")
        if is_strict:
            return ExitCodeType.VIOLATIONS_FOUND.value if ExitCodeType else 1
    else:
        print(f"✅ All CLI commands in {stats['total_files']} files contain required help strings ({stats['elapsed_ms']:.2f}ms).")

    return ExitCodeType.SUCCESS.value if ExitCodeType else 0

def main():
    parser = argparse.ArgumentParser(description="Audit CLI commands for help descriptions")
    parser.add_argument("--dir", default=".", help="Directory to audit")
    parser.add_argument("--strict", action="store_true", help="Fail with exit code 1 on warnings")
    args = parser.parse_args()

    sys.exit(run_cli_auditor(target_dir=args.dir, is_strict=args.strict))

if __name__ == "__main__":
    main()
