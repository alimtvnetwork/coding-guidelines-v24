#!/usr/bin/env python3
"""
Fast Guideline Autofixer Forwarder & Composite Runner
Combines newline/whitespace fixes (02-newline-fixer) and boolean convention auditing (05-naming-autofixer).
Multi-folder capable, customizable extensions, and sub-25ms execution.
"""

import argparse
from pathlib import Path
import sys

sys.path.insert(0, str(Path(__file__).parent))
try:
    from importlib import import_module
    newline_mod = import_module("04-newline-fixer")
    naming_mod = import_module("08-naming-autofixer")
    run_newline_auditor = newline_mod.run_newline_auditor
    run_naming_auditor = naming_mod.run_naming_auditor
except Exception as e:
    print(f"Import error: {e}")
    sys.exit(1)

def run_composite_guideline_autofixer(
    target_dir: str = ".",
    is_fix_mode: bool = True,
    extensions: str | None = None
) -> int:
    """Runs newline normalization and naming convention checks across target directory."""
    print(f"🔧 Running Guideline Autofixer across '{target_dir}'...")
    exit_nl = run_newline_auditor(target_dir=target_dir, is_fix_mode=is_fix_mode, extensions=extensions)
    exit_nm = run_naming_auditor(target_dir=target_dir, extensions=extensions)

    if exit_nl != 0:
        return exit_nl
    if exit_nm != 0:
        return exit_nm
    return 0

def main():
    parser = argparse.ArgumentParser(description="Autofix newlines and verify boolean conventions across folders")
    parser.add_argument("path", nargs="?", default=".", help="Target directory or folder (default: .)")
    parser.add_argument("--path", "-p", dest="opt_path", help="Alternative flag to specify target directory")
    parser.add_argument("--check-only", action="store_true", help="Audit without modifying files")
    parser.add_argument("--ext", help="Comma-separated extensions to scan (e.g. .md,.ts,.py)")
    args = parser.parse_args()

    target_path = args.opt_path or args.path or "."
    is_fix = not args.check_only
    sys.exit(run_composite_guideline_autofixer(target_dir=target_path, is_fix_mode=is_fix, extensions=args.ext))

if __name__ == "__main__":
    main()
