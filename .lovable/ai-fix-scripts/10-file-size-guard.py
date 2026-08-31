#!/usr/bin/env python3
"""
Fast Repository File Size & Blob Guard
Scans tracked repository files to ensure no accidental massive binary files exceed thresholds.
Multi-folder capable, customizable extensions, and nested ignore pruning (.git, .gitmap, node_modules).
"""

import argparse
from enum import Enum
import os
from pathlib import Path
import sys
import time

# --- Top-Level Enums & Constants ---
class ExitCodeType(int, Enum):
    SUCCESS = 0
    VIOLATIONS_FOUND = 1

DEFAULT_MAX_KB = 2048  # 2 MB general threshold
EXCLUDE_DIRS = {
    ".git", ".gitmap", "gitmap", ".git-map",
    "node_modules", "dist", "build", ".venv", "venv",
    ".gemini", "tmp", ".system_generated", "release-artifacts", "release-assets",
    "vendor", ".cache", ".next", "bin", "obj", "coverage", "__pycache__", ".vs", ".idea",
}

ALLOWED_LARGE_FILES = {
    "src/data/specTree.json",
    "src\\data\\specTree.json",
    "slides-app/dist.zip",
    "slides-app\\dist.zip",
}

def is_allowed_large_file(file_path: str) -> bool:
    """Checks if file is on the explicit waiver list for large generated assets."""
    norm = file_path.replace("\\", "/").lstrip("./")
    return norm in {f.replace("\\", "/").lstrip("./") for f in ALLOWED_LARGE_FILES}

def is_ignored_dir(dir_name: str) -> bool:
    """Checks if directory name is in the global ignore list."""
    return dir_name.lower() in {d.lower() for d in EXCLUDE_DIRS}

def audit_file_sizes(
    max_kb: int = DEFAULT_MAX_KB,
    target_dir: str = ".",
    allowed_exts: set[str] | None = None
) -> int:
    """Scans files and checks sizes against threshold across target directory."""
    start_time = time.perf_counter()
    violations = []
    total_files = 0
    total_bytes = 0
    max_bytes = max_kb * 1024

    for root, dirs, files in os.walk(target_dir):
        dirs[:] = [d for d in dirs if not is_ignored_dir(d)]
        for f in files:
            fp = os.path.join(root, f)
            if is_allowed_large_file(fp):
                continue
            if allowed_exts:
                ext = os.path.splitext(f)[1].lower()
                if ext not in allowed_exts:
                    continue
            try:
                sz = os.path.getsize(fp)
                total_files += 1
                total_bytes += sz
                if sz > max_bytes:
                    violations.append((fp.replace("\\", "/"), sz))
            except Exception:
                pass

    elapsed_ms = (time.perf_counter() - start_time) * 1000
    print(f"📊 Scanned {total_files:,} files in '{target_dir}' ({total_bytes / (1024*1024):.2f} MB) in {elapsed_ms:.2f}ms")

    if violations:
        print(f"\n❌ Found {len(violations)} oversized file(s) exceeding {max_kb} KB:")
        for fp, sz in sorted(violations, key=lambda x: x[1], reverse=True):
            print(f"  ::error file={fp}::{fp} ({sz / 1024:.1f} KB > {max_kb} KB)")
        return ExitCodeType.VIOLATIONS_FOUND.value

    print(f"✅ All files in '{target_dir}' within {max_kb} KB limit.")
    return ExitCodeType.SUCCESS.value

def main():
    if hasattr(sys.stdout, "reconfigure"):
        sys.stdout.reconfigure(encoding="utf-8")
    parser = argparse.ArgumentParser(description="Audit repository file sizes across folders")
    parser.add_argument("--max-kb", type=int, default=DEFAULT_MAX_KB, help="Maximum allowed file size in KB")
    parser.add_argument("--path", "-p", default=".", help="Root path or folder to audit")
    parser.add_argument("--ext", help="Optional comma-separated extension filter (e.g. .json,.zip)")
    args = parser.parse_args()

    exts = {e.strip().lower() if e.strip().startswith(".") else f".{e.strip().lower()}" for e in args.ext.split(",") if e.strip()} if args.ext else None
    sys.exit(audit_file_sizes(max_kb=args.max_kb, target_dir=args.path, allowed_exts=exts))

if __name__ == "__main__":
    main()
