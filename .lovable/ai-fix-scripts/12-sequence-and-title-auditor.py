#!/usr/bin/env python3
"""
Fast Sequence, Numbering & Title Header Auditor and Fixer
Audits and fixes:
1. Gaps and duplicate numeric prefixes in directories.
2. Mismatched Markdown H1 headers (e.g. # 00 -> # 01, # 10 -> # 17).
Multi-folder capable, customizable extensions, and pre-compiled regex engine.
"""

import argparse
from pathlib import Path
import re
import sys
import time

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

# Pre-compiled regular expressions
RE_FILE_NUM_PREFIX = re.compile(r"^([0-9]+)-(.*)\.md$")
RE_H1_HEADER = re.compile(r"^(#\s+)([0-9]+)(\s*[-—:]\s*)(.*)$", flags=re.MULTILINE)

def fix_single_file_title_header(file_path: Path, is_fix_mode: bool = False) -> tuple[str, bool]:
    """Checks and fixes H1 header number to match filename prefix."""
    norm_p = normalize_rel_path(file_path)
    m = RE_FILE_NUM_PREFIX.match(file_path.name)
    if not m:
        return (norm_p, False)

    f_num = int(m.group(1))
    if f_num >= 90:
        return (norm_p, False)

    try:
        content = read_file_lf(file_path)
        if not content:
            return (norm_p, False)
        match = RE_H1_HEADER.search(content)
        if match:
            h_num = int(match.group(2))
            if h_num != f_num:
                if h_num < 90:
                    if is_fix_mode:
                        new_header = f"{match.group(1)}{f_num:02d}{match.group(3)}{match.group(4)}"
                        updated = content[:match.start()] + new_header + content[match.end():]
                        write_file_lf(file_path, updated)
                    return (f"{norm_p}: header #{h_num:02d} != file prefix {f_num:02d}-", True)
    except Exception:
        pass
    return (norm_p, False)

def run_sequence_title_auditor(
    target_dir: str = "spec",
    is_fix_mode: bool = False,
    extensions: set[str] | tuple | None = None
) -> int:
    """Audits title header numbers across markdown specifications in target directory."""
    exts = normalize_extensions(extensions) or (".md", ".markdown")

    def handler(p: Path):
        msg, has_mismatch = fix_single_file_title_header(p, is_fix_mode=is_fix_mode)
        return msg if has_mismatch else None

    stats = process_repository_files(handler, root_dir=target_dir, extensions=exts)
    mismatches = stats["results"]

    if mismatches:
        action_word = "Fixed" if is_fix_mode else "Found header mismatches in"
        print(f"\n⚠️ {action_word} {len(mismatches)} file(s) ({stats['elapsed_ms']:.2f}ms):")
        for msg in mismatches:
            print(f"  ::notice::{msg}")
        if not is_fix_mode:
            return ExitCodeType.VIOLATIONS_FOUND.value if ExitCodeType else 1
    else:
        print(f"✅ All {stats['total_files']} markdown files in '{target_dir}' have synchronized # H1 titles ({stats['elapsed_ms']:.2f}ms).")

    return ExitCodeType.SUCCESS.value if ExitCodeType else 0

def main():
    if hasattr(sys.stdout, "reconfigure"):
        sys.stdout.reconfigure(encoding="utf-8")
    parser = argparse.ArgumentParser(description="Audit and fix numbering and titles across folders")
    parser.add_argument("path", nargs="?", default="spec", help="Target directory (default: spec)")
    parser.add_argument("--path", "-p", dest="opt_path", help="Alternative flag to specify target directory")
    parser.add_argument("--ext", help="Comma-separated extensions to check (e.g. .md)")
    parser.add_argument("--fix", action="store_true", help="Auto-fix issues in-place")
    args = parser.parse_args()

    target_path = args.opt_path or args.path or "spec"
    sys.exit(run_sequence_title_auditor(target_dir=target_path, is_fix_mode=args.fix, extensions=args.ext))

if __name__ == "__main__":
    main()
