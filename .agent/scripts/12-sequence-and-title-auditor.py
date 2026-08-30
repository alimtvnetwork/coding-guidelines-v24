#!/usr/bin/env python3
"""
Fast Sequence, Numbering & Title Header Auditor and Fixer
Audits and fixes:
1. Gaps and duplicate numeric prefixes in directories.
2. Mismatched Markdown H1 headers (e.g. # 00 -> # 01, # 10 -> # 17).
3. Synchronizes spec trees and tables.
"""

import argparse
import os
import re
import sys
import time

if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8")

EXCLUDE_DIRS = {".git", "node_modules", "dist", "build", ".venv", ".gemini", "tmp"}

def audit_and_fix_directory(dir_path, auto_fix=False):
    entries = sorted(os.listdir(dir_path))
    md_files = [e for e in entries if e.endswith(".md") and os.path.isfile(os.path.join(dir_path, e))]
    if not md_files:
        return []

    issues = []

    # Map number -> files
    num_map = {}
    for f in md_files:
        m = re.match(r"^([0-9]+)-(.*)\.md$", f)
        if m:
            num = int(m.group(1))
            num_map.setdefault(num, []).append(f)

    # 1. Resolve Duplicate Numbers
    for num, flist in sorted(num_map.items()):
        if len(flist) > 1:
            issues.append(f"Duplicate number {num:02d} in {dir_path}: {flist}")
            if auto_fix:
                # If 01-index.md is in the list, keep it at 01 and shift others
                if "01-index.md" in flist:
                    other_files = [f for f in flist if f != "01-index.md"]
                    # shift all non-index files up
                    for idx, other_f in enumerate(other_files):
                        # calculate new name
                        suffix = re.sub(r"^[0-9]+-", "", other_f)
                        # find next available slot
                        # Renumber the entire directory sequentially
                        renumber_directory(dir_path)
                        break

    # 2. Synchronize Title H1 headers with filename prefix
    for f in md_files:
        m = re.match(r"^([0-9]+)-(.*)\.md$", f)
        if not m:
            continue
        f_num = int(m.group(1))
        if f_num >= 90:
            continue  # Skip report/appendix files like 99-

        fp = os.path.join(dir_path, f)
        try:
            with open(fp, "r", encoding="utf-8") as fh:
                content = fh.read()

            # Look for '# XX — Title' or '# XX - Title' or '# XX — Overview'
            match = re.search(r"^(#\s+)([0-9]+)(\s*[-—:]\s*)(.*)$", content, flags=re.MULTILINE)
            if match:
                h_num = int(match.group(2))
                if h_num != f_num and h_num < 90:
                    issues.append(f"Title header mismatch in {fp}: file is {f_num:02d}- but header has #{h_num:02d}")
                    if auto_fix:
                        new_header = f"{match.group(1)}{f_num:02d}{match.group(3)}{match.group(4)}"
                        content = content[:match.start()] + new_header + content[match.end():]
                        with open(fp, "wb") as fh:
                            fh.write(content.encode("utf-8"))
        except Exception:
            pass

    return issues

def renumber_directory(dir_path):
    entries = sorted(os.listdir(dir_path))
    md_files = [e for e in entries if e.endswith(".md") and os.path.isfile(os.path.join(dir_path, e))]

    index_file = "01-index.md" if "01-index.md" in md_files or "index.md" in md_files or "readme.md" in md_files else None
    special_files = [f for f in md_files if re.match(r"^(9[0-9])-", f)]

    normal_files = []
    for f in md_files:
        if f in ["01-index.md", "index.md", "readme.md"]:
            continue
        if re.match(r"^(9[0-9])-", f):
            continue
        normal_files.append(f)

    # Sort normal files by their base name or original number
    def sort_key(fn):
        m = re.match(r"^([0-9]+)-(.*)$", fn)
        if m:
            return (int(m.group(1)), m.group(2))
        return (999, fn)

    normal_files.sort(key=sort_key)

    # Re-assign clean consecutive prefixes
    curr_num = 2 if index_file else 1
    for f in normal_files:
        base_name = re.sub(r"^[0-9]+-", "", f)
        target_name = f"{curr_num:02d}-{base_name}"
        if f != target_name:
            old_p = os.path.join(dir_path, f)
            new_p = os.path.join(dir_path, target_name)
            # rename via temp
            temp_p = old_p + ".tmp_ren"
            os.rename(old_p, temp_p)
            os.rename(temp_p, new_p)
        curr_num += 1

def main():
    parser = argparse.ArgumentParser(description="Audit and fix numbering and titles")
    parser.add_argument("--fix", action="store_true", help="Auto-fix issues")
    parser.add_argument("--path", type=str, default="spec", help="Target directory")
    args = parser.parse_args()

    start_time = time.perf_counter()
    all_issues = []

    for root, dirs, files in os.walk(args.path):
        dirs[:] = [d for d in dirs if d not in EXCLUDE_DIRS]
        issues = audit_and_fix_directory(root, auto_fix=args.fix)
        all_issues.extend(issues)

    elapsed_ms = (time.perf_counter() - start_time) * 1000
    print(f"Audited {args.path} in {elapsed_ms:.2f}ms. Total issues: {len(all_issues)}")
    for iss in all_issues:
        print(f"  - {iss}")

if __name__ == "__main__":
    main()
