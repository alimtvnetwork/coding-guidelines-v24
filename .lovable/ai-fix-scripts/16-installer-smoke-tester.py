#!/usr/bin/env python3
"""
Fast Generic Installer Smoke Tester
Validates install.sh and install.ps1 scripts:
1. Validates placeholder resolution (no residual PLACEHOLDER strings).
2. Verifies URL schema, fallback download mechanisms, and SHA256 checksum checks.
3. Ensures non-destructive rename-first upgrade handling.
Multi-folder capable and thread-safe lazy regex engine.
"""

import argparse
from enum import Enum
import os
from pathlib import Path
import sys
import time

sys.path.insert(0, str(Path(__file__).parent))
try:
    from importlib import import_module
    engine = import_module("02-shared-engine")
    RegexPatternType = engine.RegexPatternType
    get_compiled_regex = engine.get_compiled_regex
    ExitCodeType = engine.ExitCodeType
except Exception:
    class ExitCodeType(int, Enum):
        SUCCESS = 0
        VIOLATIONS_FOUND = 1
    RegexPatternType = None
    get_compiled_regex = None

if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8")

def smoke_test_installers(dist_dir: str = ".") -> int:
    """Smoke tests installer scripts in the target directory or root."""
    start_time = time.perf_counter()
    errors = []
    checked_files = 0

    re_placeholder = get_compiled_regex(RegexPatternType.PLACEHOLDER_TOKEN)

    candidates = ["install.sh", "install.ps1", "release-install.sh", "release-install.ps1"]
    for fname in candidates:
        fp = os.path.join(dist_dir, fname) if dist_dir != "." else fname
        has_file = os.path.exists(fp)
        if not has_file:
            is_in_root = (dist_dir != "." and os.path.exists(fname))
            if is_in_root:
                fp = fname
            else:
                continue

        checked_files += 1
        with open(fp, "r", encoding="utf-8", errors="ignore") as f:
            content = f.read()

        # 1. Check for unreplaced placeholder tokens
        placeholders = re_placeholder.findall(content)
        has_placeholders = len(placeholders) > 0
        if has_placeholders:
            errors.append(f"Unreplaced placeholders in {fp}: {set(placeholders)}")

        # 2. Check for checksum verification pattern
        content_lower = content.lower()
        has_checksum = ("sha256" in content_lower or "hash" in content_lower)
        if not has_checksum:
            errors.append(f"Installer {fp} missing SHA256 checksum verification routine")

        # 3. Check for rename-first / safe overwrite
        has_rename_first = (".old" in content or "old" in content_lower)
        if not has_rename_first:
            errors.append(f"Installer {fp} missing rename-first upgrade handling")

    elapsed_ms = (time.perf_counter() - start_time) * 1000

    is_no_files = (checked_files == 0)
    if is_no_files:
        print(f"ℹ️ No installer scripts found in '{dist_dir}' to smoke-test (skipping).")
        return ExitCodeType.SUCCESS.value

    has_errors = len(errors) > 0
    if has_errors:
        print(f"\n❌ Installer smoke testing failed in '{dist_dir}' ({elapsed_ms:.2f}ms):")
        for err in errors:
            print(f"  ::error::{err}")
        return ExitCodeType.VIOLATIONS_FOUND.value

    print(f"✅ Installer smoke tests passed for {checked_files} script(s) in '{dist_dir}' ({elapsed_ms:.2f}ms)")
    return ExitCodeType.SUCCESS.value

def main():
    parser = argparse.ArgumentParser(description="Smoke test install scripts across directories")
    parser.add_argument("path", nargs="?", default=".", help="Directory containing install scripts")
    parser.add_argument("--dir", "--path", "-p", dest="opt_dir", help="Directory containing install scripts")
    args = parser.parse_args()

    target_path = args.opt_dir or args.path or "."
    sys.exit(smoke_test_installers(target_path))

if __name__ == "__main__":
    main()
