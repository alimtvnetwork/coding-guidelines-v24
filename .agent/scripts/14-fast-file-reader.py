#!/usr/bin/env python3
"""
Fast Repository File Reader, Finder & Pluggable Cache Accessor
Provides AI agents with sub-millisecond file listing, reading, and searching using pluggable tmp/cache/.
Multi-folder capable, customizable extensions, and pre-compiled regex engine.

Usage:
  python .lovable/ai-fix-scripts/14-fast-file-reader.py --list-folder <path> [--ext .md,.ts]
  python .lovable/ai-fix-scripts/14-fast-file-reader.py --read-file <path>
  python .lovable/ai-fix-scripts/14-fast-file-reader.py --search-pattern <text> [--path <dir>] [--ext .md,.ts]
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
    load_repo_cache = engine.load_repo_cache
    read_file_safe = engine.read_file_safe
    normalize_rel_path = engine.normalize_rel_path
    normalize_extensions = engine.normalize_extensions
    process_repository_files = engine.process_repository_files
    ExitCodeType = engine.ExitCodeType
except Exception:
    ExitCodeType = None

def list_folder_files(target_dir: str = ".", extensions: tuple | set | None = None) -> list[str]:
    """Lists files in target folder quickly using cache-first streaming."""
    norm_target = normalize_rel_path(target_dir).rstrip("/")
    if norm_target == ".":
        norm_target = ""

    exts = normalize_extensions(extensions)
    files_found = []

    def handler(p: Path):
        p_rel = normalize_rel_path(p)
        if not norm_target:
            files_found.append(p_rel)
        elif p_rel.startswith(norm_target + "/"):
            files_found.append(p_rel)
        elif p_rel == norm_target:
            files_found.append(p_rel)
        return None

    process_repository_files(handler, root_dir=target_dir or ".", extensions=exts)
    return sorted(files_found)

def read_target_file(file_path: str) -> tuple[str | None, float]:
    """Reads file content with fault tolerance in sub-millisecond time."""
    start = time.perf_counter()
    content = read_file_safe(file_path)
    elapsed_ms = (time.perf_counter() - start) * 1000
    return content, elapsed_ms

def search_files_by_pattern(
    pattern: str,
    target_dir: str = ".",
    extensions: tuple | set | None = None
) -> list[str]:
    """Searches for pattern in file names/paths using pre-compiled regex."""
    search_re = re.compile(re.escape(pattern), re.IGNORECASE)
    all_files = list_folder_files(target_dir=target_dir, extensions=extensions)
    return [f for f in all_files if search_re.search(f)]

def main():
    if hasattr(sys.stdout, "reconfigure"):
        sys.stdout.reconfigure(encoding="utf-8")

    parser = argparse.ArgumentParser(description="Fast AI file reader & directory explorer across folders")
    parser.add_argument("--list-folder", "-l", help="List all files in directory")
    parser.add_argument("--read-file", "-r", help="Read file contents safely")
    parser.add_argument("--search-pattern", "-s", help="Search file paths by pattern")
    parser.add_argument("--path", "-p", default=".", help="Target directory (default: .)")
    parser.add_argument("--ext", help="Filter extensions comma-separated (e.g. .md,.ts,.py)")
    args = parser.parse_args()

    ext_set = normalize_extensions(args.ext)

    if args.read_file:
        content, elapsed_ms = read_target_file(args.read_file)
        if content is None:
            print(f"❌ File not found or unreadable: {args.read_file}")
            sys.exit(ExitCodeType.VIOLATIONS_FOUND.value if ExitCodeType else 1)
        print(f"📖 Read {args.read_file} ({len(content)} chars) in {elapsed_ms:.2f}ms:\n")
        print(content)
        sys.exit(ExitCodeType.SUCCESS.value if ExitCodeType else 0)

    if args.search_pattern:
        start = time.perf_counter()
        matches = search_files_by_pattern(args.search_pattern, target_dir=args.path, extensions=ext_set)
        elapsed_ms = (time.perf_counter() - start) * 1000
        print(f"🔍 Found {len(matches)} matching file(s) for '{args.search_pattern}' in '{args.path}' in {elapsed_ms:.2f}ms:")
        for m in matches:
            print(f"  {m}")
        sys.exit(ExitCodeType.SUCCESS.value if ExitCodeType else 0)

    if args.list_folder:
        start = time.perf_counter()
        files = list_folder_files(target_dir=args.list_folder, extensions=ext_set)
        elapsed_ms = (time.perf_counter() - start) * 1000
        print(f"📁 Found {len(files)} file(s) in '{args.list_folder}' in {elapsed_ms:.2f}ms:")
        for f in files:
            print(f"  {f}")
        sys.exit(ExitCodeType.SUCCESS.value if ExitCodeType else 0)

    parser.print_help()

if __name__ == "__main__":
    main()
