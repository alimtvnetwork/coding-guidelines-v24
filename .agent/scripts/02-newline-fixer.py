#!/usr/bin/env python3
"""
Fast Newline & Trailing Whitespace Fixer
Enforces clean UNIX LF line endings, trims trailing spaces, and ensures a single trailing newline.
"""

import argparse
from pathlib import Path
import sys

# Ensure import from same directory or shared engine
sys.path.insert(0, str(Path(__file__).parent))
try:
    from importlib import import_module
    engine = import_module("00-shared-engine")
    process_repository_files = engine.process_repository_files
    read_file_lf = engine.read_file_lf
    write_file_lf = engine.write_file_lf
    ScanModeType = engine.ScanModeType
    ExitCodeType = engine.ExitCodeType
except Exception:
    ScanModeType = None
    ExitCodeType = None

TEXT_EXTENSIONS = (".md", ".py", ".ts", ".tsx", ".js", ".jsx", ".go", ".json", ".yaml", ".yml", ".sh", ".ps1")

def clean_file_content(content: str) -> str:
    """Strips trailing whitespace per line and guarantees a single final newline."""
    lines = [line.rstrip() for line in content.split("\n")]
    # Remove excessive blank lines at EOF
    while lines and not lines[-1]:
        lines.pop()
    return "\n".join(lines) + "\n"

def process_file_newlines(file_path: Path, is_fix_mode: bool = False) -> tuple[str, bool]:
    """Checks and optionally fixes newlines and trailing whitespace in a single file."""
    try:
        raw = read_file_lf(file_path)
        cleaned = clean_file_content(raw)
        if raw != cleaned:
            if is_fix_mode:
                write_file_lf(file_path, cleaned)
            return (str(file_path), True)
    except Exception:
        pass
    return (str(file_path), False)

def run_newline_auditor(target_dir: str = ".", is_fix_mode: bool = False) -> int:
    """Executes two-phase repository scan to audit/fix newlines."""
    def handler(p: Path):
        path_str, has_issue = process_file_newlines(p, is_fix_mode=is_fix_mode)
        if has_issue:
            return path_str
        return None

    stats = process_repository_files(handler, root_dir=target_dir, extensions=TEXT_EXTENSIONS)
    violations = stats["results"]

    if violations:
        action_word = "Fixed" if is_fix_mode else "Found issues in"
        print(f"\n⚠️ {action_word} {len(violations)} file(s) ({stats['elapsed_ms']:.2f}ms):")
        for v in violations:
            print(f"  ::notice file={v}::{v}")
        if not is_fix_mode:
            return ExitCodeType.VIOLATIONS_FOUND.value if ExitCodeType else 1
    else:
        print(f"✅ All {stats['total_files']} files have clean newlines ({stats['elapsed_ms']:.2f}ms).")

    return ExitCodeType.SUCCESS.value if ExitCodeType else 0

def main():
    parser = argparse.ArgumentParser(description="Fix trailing whitespace and newlines")
    parser.add_argument("path", nargs="?", default=".", help="Root directory to scan")
    parser.add_argument("--fix", action="store_true", help="Auto-fix whitespace issues in-place")
    args = parser.parse_args()

    sys.exit(run_newline_auditor(target_dir=args.path, is_fix_mode=args.fix))

if __name__ == "__main__":
    main()
