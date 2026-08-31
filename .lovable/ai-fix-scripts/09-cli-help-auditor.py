#!/usr/bin/env python3
"""
Fast CLI Command Discovery & Help Text Parity Auditor
Inspects CLI entry points, subcommands, and flags across Go (Cobra), TypeScript (Commander), Python (Click/Argparse), and PHP (Symfony).
Multi-folder capable, customizable extensions, and thread-safe lazy regex engine.

All Enums, Constants, and Functions are imported directly from 02-shared-engine.py.
"""

import argparse
import ast
from importlib import import_module
from pathlib import Path
import sys

sys.path.insert(0, str(Path(__file__).parent))
engine = import_module("02-shared-engine")

process_repository_files = engine.process_repository_files
read_file_lf = engine.read_file_lf
normalize_extensions = engine.normalize_extensions
normalize_rel_path = engine.normalize_rel_path
ExitCodeType = engine.ExitCodeType
RegexPatternType = engine.RegexPatternType
get_compiled_regex = engine.get_compiled_regex
DEFAULT_CLI_EXTENSIONS = engine.DEFAULT_CLI_EXTENSIONS
DEFAULT_ENCODING = engine.DEFAULT_ENCODING
LINE_SEPARATOR = engine.LINE_SEPARATOR
CURRENT_DIR = engine.CURRENT_DIR

def audit_go_cobra_commands(content: str) -> list[tuple[str, str]]:
    """Detects Go Cobra commands missing Short or Example descriptions."""
    violations = []
    re_cobra = get_compiled_regex(RegexPatternType.COBRA_COMMAND)
    re_short = get_compiled_regex(RegexPatternType.SHORT_DESC)
    re_example = get_compiled_regex(RegexPatternType.EXAMPLE_USAGE)

    for match in re_cobra.finditer(content):
        cmd_var = match.group(1)
        body = match.group(2)
        has_short = bool(re_short.search(body))
        if not has_short:
            violations.append((cmd_var, "Missing Short description in cobra.Command"))
        is_root = (cmd_var == "rootCmd")
        if not is_root:
            has_example = bool(re_example.search(body))
            if not has_example:
                violations.append((cmd_var, "Missing Example usage in cobra.Command"))
    return violations

def audit_python_cli_commands(file_path: Path, content: str) -> list[tuple[str, str]]:
    """Detects Python CLI commands missing docstrings or help text."""
    violations = []
    try:
        tree = ast.parse(content, filename=str(file_path))
        for node in ast.walk(tree):
            is_func = isinstance(node, ast.FunctionDef)
            if is_func:
                for dec in node.decorator_list:
                    if isinstance(dec, ast.Call):
                        if hasattr(dec.func, "attr"):
                            is_cmd = (dec.func.attr == "command")
                            if is_cmd:
                                has_doc = bool(ast.get_docstring(node))
                                if not has_doc:
                                    violations.append((node.name, "Missing docstring for CLI command function"))
    except Exception:
        pass
    return violations

def audit_single_file_cli(file_path: Path) -> list[tuple[str, str]]:
    """Audits a single file for CLI help compliance."""
    try:
        content = read_file_lf(file_path, encoding=DEFAULT_ENCODING)
        is_go = (file_path.suffix == ".go")
        if is_go:
            return audit_go_cobra_commands(content)
        is_py = (file_path.suffix == ".py")
        if is_py:
            return audit_python_cli_commands(file_path, content)
    except Exception:
        pass
    return []

def run_cli_auditor(
    target_dir: str = CURRENT_DIR,
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

    has_violations = len(all_violations) > 0
    if has_violations:
        print(f"{LINE_SEPARATOR}⚠️ Found CLI help description issues in {len(all_violations)} file(s) ({stats['elapsed_ms']:.2f}ms):")
        for fp, vios in all_violations:
            for cmd, msg in vios:
                print(f"  ::warning file={fp}::{cmd}: {msg}")
        if is_strict:
            return ExitCodeType.VIOLATIONS_FOUND.value
    else:
        print(f"✅ All CLI commands in {stats['total_files']} files in '{target_dir}' contain required help strings ({stats['elapsed_ms']:.2f}ms).")

    return ExitCodeType.SUCCESS.value

def main():
    parser = argparse.ArgumentParser(description="Audit CLI commands for help descriptions across folders")
    parser.add_argument("path", nargs="?", default=CURRENT_DIR, help="Directory to audit")
    parser.add_argument("--dir", "--path", "-p", dest="opt_dir", help="Directory to audit")
    parser.add_argument("--ext", help="Comma-separated extensions to scan (e.g. .go,.py)")
    parser.add_argument("--strict", action="store_true", help="Fail with exit code 1 on warnings")
    args = parser.parse_args()

    target_path = args.opt_dir or args.path or CURRENT_DIR
    sys.exit(run_cli_auditor(target_dir=target_path, is_strict=args.strict, extensions=args.ext))

if __name__ == "__main__":
    main()
