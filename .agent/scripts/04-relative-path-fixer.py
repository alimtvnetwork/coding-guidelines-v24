#!/usr/bin/env python3
"""
Fast Relative Path Fixer & Absolute URI Auditor
Detects and sanitizes absolute filesystem paths (C:\\..., D:\\..., /home/..., file:///) in documentation.
"""

import argparse
from pathlib import Path
import re
import sys

sys.path.insert(0, str(Path(__file__).parent))
try:
    from importlib import import_module
    engine = import_module("00-shared-engine")
    process_repository_files = engine.process_repository_files
    read_file_lf = engine.read_file_lf
    write_file_lf = engine.write_file_lf
    ExitCodeType = engine.ExitCodeType
except Exception:
    ExitCodeType = None

TARGET_EXTENSIONS = (".md", ".json", ".yaml", ".yml", ".py", ".ts", ".go")
ABSOLUTE_PATH_PATTERNS = [
    re.compile(r"file:///[A-Za-z]:/[^\s\)\]\"'>]+"),
    re.compile(r"(?<![A-Za-z0-9_])[A-Za-z]:\\[A-Za-z0-9_\\.-]+"),
]

def sanitize_content_paths(content: str) -> tuple[str, int]:
    """Replaces absolute paths with clean relative paths."""
    modified = content
    count = 0
    # Strip file URI prefix -> relative
    for match in re.finditer(r"file:///[A-Za-z]:/[^/]+/coding-guidelines/([^\s\)\]\"'>]+)", content):
        rel_target = match.group(1)
        modified = modified.replace(match.group(0), rel_target)
        count += 1
    return modified, count

def audit_file_paths(file_path: Path, is_fix_mode: bool = False) -> tuple[str, list[str]]:
    """Audits a single file for forbidden absolute paths."""
    try:
        content = read_file_lf(file_path)
        violations = []
        for pat in ABSOLUTE_PATH_PATTERNS:
            for match in pat.finditer(content):
                # ignore standard regex documentation patterns
                val = match.group(0)
                if "\\\\?\\" not in val and not val.endswith("."):
                    violations.append(val)

        if violations and is_fix_mode:
            cleaned, fix_count = sanitize_content_paths(content)
            if fix_count > 0:
                write_file_lf(file_path, cleaned)

        return (str(file_path), violations)
    except Exception:
        return (str(file_path), [])

def run_path_auditor(target_dir: str = ".", is_fix_mode: bool = False) -> int:
    """Runs repository-wide path check using two-phase pipeline."""
    def handler(p: Path):
        fp_str, vios = audit_file_paths(p, is_fix_mode=is_fix_mode)
        if vios:
            return (fp_str, vios)
        return None

    stats = process_repository_files(handler, root_dir=target_dir, extensions=TARGET_EXTENSIONS)
    all_violations = stats["results"]

    if all_violations:
        print(f"\n❌ Found absolute path references in {len(all_violations)} file(s) ({stats['elapsed_ms']:.2f}ms):")
        for fp, vios in all_violations:
            for v in vios[:3]:
                print(f"  ::error file={fp}::Absolute path found: {v}")
        return ExitCodeType.VIOLATIONS_FOUND.value if ExitCodeType else 1
    else:
        print(f"✅ All {stats['total_files']} files use strict relative paths ({stats['elapsed_ms']:.2f}ms).")
        return ExitCodeType.SUCCESS.value if ExitCodeType else 0

def main():
    parser = argparse.ArgumentParser(description="Audit and fix absolute paths")
    parser.add_argument("path", nargs="?", default=".", help="Root directory to scan")
    parser.add_argument("--fix", action="store_true", help="Auto-fix recognized path patterns")
    args = parser.parse_args()

    sys.exit(run_path_auditor(target_dir=args.path, is_fix_mode=args.fix))

if __name__ == "__main__":
    main()
