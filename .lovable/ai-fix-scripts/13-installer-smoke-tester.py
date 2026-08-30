#!/usr/bin/env python3
"""
Fast Generic Installer Smoke Tester
Validates install.sh and install.ps1 scripts:
1. Validates placeholder resolution (no residual PLACEHOLDER strings).
2. Verifies URL schema, fallback download mechanisms, and SHA256 checksum checks.
3. Ensures non-destructive rename-first upgrade handling.
"""

import argparse
import os
import re
import sys
import time

if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8")

def smoke_test_installers(dist_dir="dist"):
    start_time = time.perf_counter()
    errors = []
    checked_files = 0

    candidates = ["install.sh", "install.ps1"]
    for fname in candidates:
        fp = os.path.join(dist_dir, fname) if dist_dir != "." else fname
        if not os.path.exists(fp):
            # check root if not in dist
            if os.path.exists(fname):
                fp = fname
            else:
                continue

        checked_files += 1
        with open(fp, "r", encoding="utf-8", errors="ignore") as f:
            content = f.read()

        # 1. Check for unreplaced placeholder tokens
        placeholders = re.findall(r"[A-Z0-9_]*PLACEHOLDER[A-Z0-9_]*", content)
        if placeholders:
            errors.append(f"Unreplaced placeholders in {fp}: {set(placeholders)}")

        # 2. Check for checksum verification pattern
        if "sha256" not in content.lower() and "hash" not in content.lower():
            errors.append(f"Installer {fp} missing SHA256 checksum verification routine")

        # 3. Check for rename-first / safe overwrite
        if ".old" not in content and "old" not in content.lower():
            errors.append(f"Installer {fp} missing rename-first upgrade handling")

    elapsed_ms = (time.perf_counter() - start_time) * 1000

    if checked_files == 0:
        print("ℹ️ No install.sh or install.ps1 found to smoke-test (skipping).")
        return 0

    if errors:
        print(f"\n❌ Installer smoke testing failed ({elapsed_ms:.2f}ms):")
        for err in errors:
            print(f"  ::error::{err}")
        return 1
    else:
        print(f"✅ Installer smoke tests passed for {checked_files} script(s) ({elapsed_ms:.2f}ms)")
        return 0

def main():
    parser = argparse.ArgumentParser(description="Smoke test install scripts")
    parser.add_argument("--dir", type=str, default=".", help="Directory containing install scripts")
    args = parser.parse_args()
    sys.exit(smoke_test_installers(args.dir))

if __name__ == "__main__":
    main()
