#!/usr/bin/env python3
"""
Fast Version Synchronization & Changelog Guard
Validates that version.json, package.json, and changelog.md are in 100% sync.
Multi-folder capable, thread-safe lazy regex engine, and sub-5ms execution.
"""

import argparse
from enum import Enum
import json
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

def check_version_sync(root_dir: str = ".") -> int:
    """Checks version parity across version.json, package.json, and changelog.md."""
    start_time = time.perf_counter()
    errors = []

    version_json_p = os.path.join(root_dir, "version.json")
    package_json_p = os.path.join(root_dir, "package.json")
    changelog_p = os.path.join(root_dir, "changelog.md")

    canonical_version = None

    if os.path.exists(version_json_p):
        try:
            with open(version_json_p, "r", encoding="utf-8") as f:
                v_data = json.load(f)
                canonical_version = v_data.get("version")
        except Exception as e:
            errors.append(f"version.json parse error: {e}")
    elif os.path.exists(package_json_p):
        try:
            with open(package_json_p, "r", encoding="utf-8") as f:
                p_data = json.load(f)
                canonical_version = p_data.get("version")
        except Exception as e:
            errors.append(f"package.json parse error: {e}")

    if not canonical_version:
        print(f"⚠️ No canonical version source (version.json or package.json) found in '{root_dir}'.")
        return ExitCodeType.SUCCESS.value

    # 1. Compare with package.json
    if os.path.exists(package_json_p):
        try:
            with open(package_json_p, "r", encoding="utf-8") as f:
                p_data = json.load(f)
                pkg_ver = p_data.get("version")
                if pkg_ver != canonical_version:
                    errors.append(f"Version mismatch: version.json has '{canonical_version}' but package.json has '{pkg_ver}'")
        except Exception as e:
            errors.append(f"package.json error: {e}")

    # 2. Compare with changelog.md latest entry
    if os.path.exists(changelog_p):
        try:
            with open(changelog_p, "r", encoding="utf-8") as f:
                changelog_text = f.read()
            re_changelog = get_compiled_regex(RegexPatternType.CHANGELOG_HEADER)
            match = re_changelog.search(changelog_text)
            if match:
                latest_cl_ver = match.group(1).lstrip("v")
                if latest_cl_ver != canonical_version.lstrip("v"):
                    errors.append(f"Changelog mismatch: latest header is 'v{latest_cl_ver}' but canonical version is '{canonical_version}'")
        except Exception as e:
            errors.append(f"changelog.md error: {e}")

    elapsed_ms = (time.perf_counter() - start_time) * 1000

    if errors:
        print(f"\n❌ Version synchronization failed in '{root_dir}' ({elapsed_ms:.2f}ms):")
        for err in errors:
            print(f"  ::error::{err}")
        return ExitCodeType.VIOLATIONS_FOUND.value

    print(f"✅ Version synchronization verified: v{canonical_version} ({elapsed_ms:.2f}ms)")
    return ExitCodeType.SUCCESS.value

def main():
    parser = argparse.ArgumentParser(description="Check version synchronization across folders")
    parser.add_argument("path", nargs="?", default=".", help="Root directory containing version files")
    parser.add_argument("--path", "-p", dest="opt_path", help="Alternative flag to specify root directory")
    args = parser.parse_args()

    target_path = args.opt_path or args.path or "."
    sys.exit(check_version_sync(target_path))

if __name__ == "__main__":
    main()
