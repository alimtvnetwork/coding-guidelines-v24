#!/usr/bin/env python3
"""
Fast Repository File Size & Blob Guard
Scans all tracked repository files to ensure no accidental massive binary files exceed thresholds.
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
EXCLUDE_DIRS = {".git", "node_modules", "dist", "build", ".venv", ".gemini", "tmp", ".system_generated"}
ALLOWED_LARGE_FILES = {
    "src/data/specTree.json",
    "src\\data\\specTree.json",
    "slides-app/dist.zip",
    "slides-app\\dist.zip",
}

def is_allowed_large_file(file_path: str) -> bool:
    """Checks if file is on the explicit waiver list for large generated assets."""
    norm = file_path.replace("\\", "/").lstrip("./")
    return norm in {f.replace("\\", "/") for f in ALLOWED_LARGE_FILES}

def audit_file_sizes(max_kb: int = DEFAULT_MAX_KB, target_dir: str = ".") -> int:
    """Scans all files and checks sizes against threshold."""
    start_time = time.perf_counter()
    violations = []
    total_files = 0
    total_bytes = 0
    max_bytes = max_kb * 1024

    for root, dirs, files in os.walk(target_dir):
        dirs[:] = [d for d in dirs if d not in EXCLUDE_DIRS]
        for f in files:
            fp = os.path.join(root, f)
            if is_allowed_large_file(fp):
                continue
            try:
                sz = os.path.getsize(fp)
                total_files += 1
                total_bytes += sz
                if sz > max_bytes:
                    violations.append((fp, sz))
            except Exception:
                pass

    elapsed_ms = (time.perf_counter() - start_time) * 1000
    print(f"📊 Scanned {total_files:,} files ({total_bytes / (1024*1024):.2f} MB) in {elapsed_ms:.2f}ms")

    if violations:
        print(f"\n❌ Found {len(violations)} oversized file(s) exceeding {max_kb} KB:")
        for fp, sz in sorted(violations, key=lambda x: x[1], reverse=True):
            print(f"  ::error file={fp}::{fp} ({sz / 1024:.1f} KB > {max_kb} KB)")
        return ExitCodeType.VIOLATIONS_FOUND.value

    print(f"✅ All files within {max_kb} KB limit.")
    return ExitCodeType.SUCCESS.value

def main():
    if hasattr(sys.stdout, "reconfigure"):
        sys.stdout.reconfigure(encoding="utf-8")
    parser = argparse.ArgumentParser(description="Audit repository file sizes")
    parser.add_argument("--max-kb", type=int, default=DEFAULT_MAX_KB, help="Maximum allowed file size in KB")
    parser.add_argument("--path", type=str, default=".", help="Root path to audit")
    args = parser.parse_args()

    sys.exit(audit_file_sizes(max_kb=args.max_kb, target_dir=args.path))

if __name__ == "__main__":
    main()
