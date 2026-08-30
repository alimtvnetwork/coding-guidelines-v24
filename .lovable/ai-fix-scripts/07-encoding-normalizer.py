#!/usr/bin/env python3
"""
Fast UTF-8 & UNIX LF Encoding Normalizer
Recursively audits and standardizes all text files to UTF-8 without BOM and strict UNIX LF (\\n).
"""

import argparse
from pathlib import Path
import sys

sys.path.insert(0, str(Path(__file__).parent))
try:
    from importlib import import_module
    engine = import_module("00-shared-engine")
    process_repository_files = engine.process_repository_files
    write_file_lf = engine.write_file_lf
    is_binary_file = engine.is_binary_file
    ExitCodeType = engine.ExitCodeType
except Exception:
    ExitCodeType = None

def normalize_single_file(file_path: Path, is_fix_mode: bool = False) -> tuple[str, bool]:
    """Audits and converts CRLF/BOM in a file to clean UTF-8 LF."""
    if is_binary_file(file_path):
        return (str(file_path), False)
    try:
        with open(file_path, "rb") as f:
            raw_bytes = f.read()

        # Check BOM or CRLF
        has_bom = raw_bytes.startswith(b"\xef\xbb\xbf")
        has_crlf = b"\r\n" in raw_bytes

        if has_bom or has_crlf:
            if is_fix_mode:
                text = raw_bytes.decode("utf-8-sig", errors="replace")
                write_file_lf(file_path, text)
            return (str(file_path), True)
    except Exception:
        pass
    return (str(file_path), False)

def run_encoding_normalizer(target_dir: str = ".", is_fix_mode: bool = False) -> int:
    """Runs repository encoding check and normalizer."""
    def handler(p: Path):
        fp_str, changed = normalize_single_file(p, is_fix_mode=is_fix_mode)
        if changed:
            return fp_str
        return None

    stats = process_repository_files(handler, root_dir=target_dir)
    affected = stats["results"]

    if affected:
        action_verb = "Normalized" if is_fix_mode else "Found CRLF/BOM in"
        print(f"\n⚠️ {action_verb} {len(affected)} file(s) ({stats['elapsed_ms']:.2f}ms):")
        for f in affected[:5]:
            print(f"  ::notice file={f}::{f}")
        if not is_fix_mode:
            return ExitCodeType.VIOLATIONS_FOUND.value if ExitCodeType else 1
    else:
        print(f"✅ All {stats['total_files']} files normalized to UTF-8 LF ({stats['elapsed_ms']:.2f}ms).")

    return ExitCodeType.SUCCESS.value if ExitCodeType else 0

def main():
    parser = argparse.ArgumentParser(description="Normalize files to UTF-8 UNIX LF")
    parser.add_argument("path", nargs="?", default=".", help="Root directory")
    parser.add_argument("--fix", action="store_true", help="Fix CRLF and BOM in-place")
    args = parser.parse_args()

    sys.exit(run_encoding_normalizer(target_dir=args.path, is_fix_mode=args.fix))

if __name__ == "__main__":
    main()
