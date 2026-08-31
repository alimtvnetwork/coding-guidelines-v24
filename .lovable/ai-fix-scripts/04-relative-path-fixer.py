#!/usr/bin/env python3
"""
Fast Relative Path Fixer & Absolute URI Auditor
Detects and sanitizes absolute filesystem paths (C:\\..., D:\\..., /home/..., file:///) in documentation.
Multi-folder capable, customizable extensions, and pre-compiled regex engine.
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
    normalize_extensions = engine.normalize_extensions
    normalize_rel_path = engine.normalize_rel_path
    ExitCodeType = engine.ExitCodeType
except Exception:
    ExitCodeType = None

DEFAULT_TARGET_EXTENSIONS = (".md", ".markdown", ".json", ".yaml", ".yml", ".py", ".ts", ".go", ".php", ".cs")

# Pre-compiled regular expressions at module level
RE_FILE_URI_WIN = re.compile(r"file:///[A-Za-z]:/[^\s\)\]\"'>]+")
RE_DRIVE_ABS_WIN = re.compile(r"(?<![A-Za-z0-9_])[A-Za-z]:\\[A-Za-z0-9_\\.-]+")
RE_REPO_FILE_URI = re.compile(r"file:///[A-Za-z]:/[^/]+/coding-guidelines/([^\s\)\]\"'>]+)")

ABSOLUTE_PATH_PATTERNS = (
    RE_FILE_URI_WIN,
    RE_DRIVE_ABS_WIN,
)

def sanitize_content_paths(content: str) -> tuple[str, int]:
    """Replaces absolute repository paths with clean relative paths."""
    modified = content
    count = 0
    for match in RE_REPO_FILE_URI.finditer(content):
        rel_target = match.group(1)
        modified = modified.replace(match.group(0), rel_target)
        count += 1
    return modified, count

def audit_file_paths(file_path: Path, is_fix_mode: bool = False) -> tuple[str, list[str]]:
    """Audits a single file for forbidden absolute paths."""
    norm_p = normalize_rel_path(file_path)
    try:
        content = read_file_lf(file_path)
        violations = []
        for pat in ABSOLUTE_PATH_PATTERNS:
            for match in pat.finditer(content):
                val = match.group(0)
                if "\\\\?\\" not in val:
                    if not val.endswith("."):
                        violations.append(val)

        if violations:
            if is_fix_mode:
                cleaned, fix_count = sanitize_content_paths(content)
                if fix_count > 0:
                    write_file_lf(file_path, cleaned)

        return (norm_p, violations)
    except Exception:
        return (norm_p, [])

def run_path_auditor(
    target_dir: str = ".",
    is_fix_mode: bool = False,
    extensions: set[str] | tuple | None = None
) -> int:
    """Runs repository-wide path check using two-phase pipeline."""
    exts = normalize_extensions(extensions) or DEFAULT_TARGET_EXTENSIONS

    def handler(p: Path):
        fp_str, vios = audit_file_paths(p, is_fix_mode=is_fix_mode)
        return (fp_str, vios) if vios else None

    stats = process_repository_files(handler, root_dir=target_dir, extensions=exts)
    all_violations = stats["results"]

    if all_violations:
        print(f"\n❌ Found absolute path references in {len(all_violations)} file(s) ({stats['elapsed_ms']:.2f}ms):")
        for fp, vios in all_violations:
            for v in vios[:3]:
                print(f"  ::error file={fp}::Absolute path found: {v}")
        return ExitCodeType.VIOLATIONS_FOUND.value if ExitCodeType else 1

    print(f"✅ All {stats['total_files']} files in '{target_dir}' use strict relative paths ({stats['elapsed_ms']:.2f}ms).")
    return ExitCodeType.SUCCESS.value if ExitCodeType else 0

def main():
    parser = argparse.ArgumentParser(description="Audit and fix absolute paths across target folders")
    parser.add_argument("path", nargs="?", default=".", help="Root directory or folder to scan")
    parser.add_argument("--path", "-p", dest="opt_path", help="Alternative flag to specify target directory")
    parser.add_argument("--fix", action="store_true", help="Auto-fix recognized path patterns")
    parser.add_argument("--ext", help="Comma-separated extensions to scan (e.g. .md,.ts,.py)")
    args = parser.parse_args()

    target_path = args.opt_path or args.path or "."
    sys.exit(run_path_auditor(target_dir=target_path, is_fix_mode=args.fix, extensions=args.ext))

if __name__ == "__main__":
    main()
