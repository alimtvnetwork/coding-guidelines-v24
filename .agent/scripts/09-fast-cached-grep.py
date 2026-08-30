#!/usr/bin/env python3
"""
Fast Cached Content Grepper & Pattern Matcher
Leverages tmp/repo-file-cache.json and concurrent threads to search file contents at sub-15ms speeds.
"""

import argparse
from concurrent.futures import ThreadPoolExecutor
from enum import Enum
import json
from pathlib import Path
import re
import sys
import time

sys.path.insert(0, str(Path(__file__).parent))
try:
    from importlib import import_module
    engine = import_module("00-shared-engine")
    load_repo_cache = engine.load_repo_cache
    read_file_lf = engine.read_file_lf
    ExitCodeType = engine.ExitCodeType
except Exception:
    ExitCodeType = None

LANG_EXTENSION_MAP = {
    "go": {".go"},
    "ts": {".ts", ".tsx"},
    "tsx": {".tsx"},
    "js": {".js", ".jsx"},
    "jsx": {".jsx"},
    "py": {".py"},
    "md": {".md"},
    "json": {".json"},
    "yaml": {".yaml", ".yml"},
    "php": {".php"},
    "cs": {".cs"},
}

def resolve_language_extensions(lang_str: str) -> set[str]:
    """Resolves language aliases into extension sets."""
    if not lang_str:
        return set()
    exts = set()
    for l in lang_str.lower().split(","):
        l = l.strip()
        if l in LANG_EXTENSION_MAP:
            exts.update(LANG_EXTENSION_MAP[l])
        elif l.startswith("."):
            exts.add(l)
        elif l:
            exts.add(f".{l}")
    return exts

def search_single_file(
    file_path: Path,
    regex_obj: re.Pattern,
    max_per_file: int = 10
) -> list[dict]:
    """Searches a single file and returns matching line details."""
    matches = []
    try:
        content = read_file_lf(file_path)
        for idx, line in enumerate(content.split("\n"), start=1):
            if regex_obj.search(line):
                matches.append({
                    "line": idx,
                    "content": line.strip()
                })
                if len(matches) >= max_per_file:
                    break
    except Exception:
        pass
    return matches

def execute_parallel_grep(
    files: list[Path],
    pattern: str,
    is_regex: bool = False,
    is_case_sensitive: bool = False,
    max_per_file: int = 10
) -> tuple[dict[str, list[dict]], float]:
    """Runs parallel multi-threaded content grep across target files."""
    start_time = time.perf_counter()
    flags = 0 if is_case_sensitive else re.IGNORECASE
    regex_pattern = pattern if is_regex else re.escape(pattern)
    compiled_re = re.compile(regex_pattern, flags)

    results = {}
    with ThreadPoolExecutor(max_workers=8) as executor:
        futures = {executor.submit(search_single_file, fp, compiled_re, max_per_file): fp for fp in files}
        for fut in futures:
            fp = futures[fut]
            file_matches = fut.result()
            if file_matches:
                results[str(fp).replace("\\", "/")] = file_matches

    elapsed_ms = (time.perf_counter() - start_time) * 1000
    return results, elapsed_ms

def main():
    if hasattr(sys.stdout, "reconfigure"):
        sys.stdout.reconfigure(encoding="utf-8")

    parser = argparse.ArgumentParser(description="Fast cached grep tool")
    parser.add_argument("--pattern", "-p", required=True, help="Search string or regex")
    parser.add_argument("--regex", "-r", action="store_true", help="Regex mode")
    parser.add_argument("--case-sensitive", "-c", action="store_true", help="Case sensitive")
    parser.add_argument("--lang", "-l", help="Language filter (e.g. go, ts, py, md)")
    parser.add_argument("--path", default=".", help="Search root")
    parser.add_argument("--limit", type=int, default=50, help="Max results to display")
    parser.add_argument("--max-per-file", type=int, default=10, help="Max matches per file")
    args = parser.parse_args()

    cache_data = load_repo_cache()
    cached_files = cache_data.get("files", [])
    exts = resolve_language_extensions(args.lang)

    target_files = []
    for f in cached_files:
        p = Path(f)
        if args.path != "." and not str(p).replace("\\", "/").startswith(args.path.replace("\\", "/")):
            continue
        if exts and p.suffix.lower() not in exts:
            continue
        target_files.append(p)

    results, elapsed_ms = execute_parallel_grep(
        target_files,
        pattern=args.pattern,
        is_regex=args.regex,
        is_case_sensitive=args.case_sensitive,
        max_per_file=args.max_per_file
    )

    total_matches = sum(len(m) for m in results.values())
    print(f"🔍 Grepped {len(target_files)} cached files in {elapsed_ms:.2f}ms — Found {total_matches} match(es) in {len(results)} file(s):\n")

    printed = 0
    for fp, matches in results.items():
        if printed >= args.limit:
            break
        print(f"📄 {fp}:")
        for m in matches:
            if printed >= args.limit:
                break
            print(f"  Line {m['line']:>4}: {m['content']}")
            printed += 1

    sys.exit(0)

if __name__ == "__main__":
    main()
