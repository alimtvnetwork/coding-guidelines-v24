#!/usr/bin/env python3
"""
Fast Repository File Reader, Finder & Pluggable Cache Accessor
Provides AI agents with sub-millisecond file listing, reading, and searching using pluggable tmp/cache/.

Usage:
  python .lovable/ai-fix-scripts/14-fast-file-reader.py --list-folder <path>
  python .lovable/ai-fix-scripts/14-fast-file-reader.py --read-file <path>
  python .lovable/ai-fix-scripts/14-fast-file-reader.py --search-pattern <text> [--ext .md,.ts]
"""

import argparse
from pathlib import Path
import sys
import time

sys.path.insert(0, str(Path(__file__).parent))
try:
    from importlib import import_module
    engine = import_module("00-shared-engine")
    load_repo_cache = engine.load_repo_cache
    read_file_safe = engine.read_file_safe
    normalize_rel_path = engine.normalize_rel_path
    process_repository_files = engine.process_repository_files
    ExitCodeType = engine.ExitCodeType
except Exception:
    ExitCodeType = None

def list_folder_files(target_dir: str = ".", extensions: tuple = None) -> list[str]:
    """Lists files in target folder quickly using cache-first streaming."""
    norm_target = normalize_rel_path(target_dir).rstrip("/")
    if norm_target == ".":
        norm_target = ""

    files_found = []
    def handler(p: Path):
        p_rel = normalize_rel_path(p)
        if not norm_target or p_rel.startswith(norm_target + "/") or p_rel == norm_target:
            files_found.append(p_rel)
        return None

    process_repository_files(handler, root_dir=target_dir or ".", extensions=extensions)
    return sorted(files_found)

def read_target_file(file_path: str) -> tuple[str | None, float]:
    """Reads file content with fault tolerance in sub-millisecond time."""
    start = time.perf_counter()
    content = read_file_safe(file_path)
    elapsed_ms = (time.perf_counter() - start) * 1000
    return content, elapsed_ms

def main():
    if hasattr(sys.stdout, "reconfigure"):
        sys.stdout.reconfigure(encoding="utf-8")

    parser = argparse.ArgumentParser(description="Fast AI file reader & directory explorer")
    parser.add_argument("--list-folder", "-l", help="List all files in directory")
    parser.add_argument("--read-file", "-r", help="Read file contents safely")
    parser.add_argument("--ext", help="Filter extensions comma-separated (e.g. .md,.ts)")
    args = parser.parse_args()

    ext_tuple = tuple(e.strip().lower() for e in args.ext.split(",")) if args.ext else None

    if args.read_file:
        content, elapsed_ms = read_target_file(args.read_file)
        if content is None:
            print(f"❌ File not found or unreadable: {args.read_file}")
            sys.exit(ExitCodeType.VIOLATIONS_FOUND.value if ExitCodeType else 1)
        print(f"📖 Read {args.read_file} ({len(content)} chars) in {elapsed_ms:.2f}ms:\n")
        print(content)
        sys.exit(ExitCodeType.SUCCESS.value if ExitCodeType else 0)

    if args.list_folder:
        start = time.perf_counter()
        files = list_folder_files(target_dir=args.list_folder, extensions=ext_tuple)
        elapsed_ms = (time.perf_counter() - start) * 1000
        print(f"📁 Found {len(files)} file(s) in '{args.list_folder}' in {elapsed_ms:.2f}ms:")
        for f in files:
            print(f"  {f}")
        sys.exit(ExitCodeType.SUCCESS.value if ExitCodeType else 0)

    parser.print_help()

if __name__ == "__main__":
    main()
