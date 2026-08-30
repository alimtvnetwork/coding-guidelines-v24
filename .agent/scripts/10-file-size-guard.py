#!/usr/bin/env python3
"""
Fast Repository File Size & Blob Guard
Scans all tracked repository files to ensure no oversized binary blobs or massive files exceed the configurable threshold (default 1MB).
"""

import argparse
import os
import sys
import time

if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8")

DEFAULT_MAX_KB = 1024  # 1 MB threshold
EXCLUDE_DIRS = {".git", "node_modules", "dist", "build", ".venv", ".gemini", "tmp"}

def audit_file_sizes(max_kb=DEFAULT_MAX_KB, target_dir="."):
    start_time = time.perf_counter()
    violations = []
    total_files = 0
    total_bytes = 0

    max_bytes = max_kb * 1024

    for root, dirs, files in os.walk(target_dir):
        dirs[:] = [d for d in dirs if d not in EXCLUDE_DIRS]
        for f in files:
            fp = os.path.join(root, f)
            try:
                size = os.path.getsize(fp)
                total_files += 1
                total_bytes += size
                if size > max_bytes:
                    violations.append((fp, size))
            except Exception:
                pass

    elapsed_ms = (time.perf_counter() - start_time) * 1000

    print(f"📊 Scanned {total_files:,} files ({total_bytes / (1024*1024):.2f} MB) in {elapsed_ms:.2f}ms")

    if violations:
        print(f"\n❌ Found {len(violations)} oversized file(s) exceeding {max_kb} KB:")
        for fp, sz in sorted(violations, key=lambda x: x[1], reverse=True):
            print(f"  ::error file={fp}::{fp} ({sz / 1024:.1f} KB > {max_kb} KB)")
        return 1
    else:
        print(f"✅ All files within {max_kb} KB limit.")
        return 0

def main():
    parser = argparse.ArgumentParser(description="Audit repository file sizes")
    parser.add_argument("--max-kb", type=int, default=DEFAULT_MAX_KB, help="Maximum allowed file size in KB")
    parser.add_argument("--path", type=str, default=".", help="Root path to audit")
    args = parser.parse_args()

    sys.exit(audit_file_sizes(max_kb=args.max_kb, target_dir=args.path))

if __name__ == "__main__":
    main()
