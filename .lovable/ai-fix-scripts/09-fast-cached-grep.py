#!/usr/bin/env python3
"""
Fast Cached Content Grepper & Pattern Matcher.
Leverages tmp/repo-file-cache.json to search file contents at high speed in parallel.

Usage:
  python .lovable/ai-fix-scripts/09-fast-cached-grep.py --pattern <text> [options]

Examples:
  python .lovable/ai-fix-scripts/09-fast-cached-grep.py --pattern "AppError"
  python .lovable/ai-fix-scripts/09-fast-cached-grep.py --pattern "interface.*Type" --regex --lang ts,tsx
  python .lovable/ai-fix-scripts/09-fast-cached-grep.py --pattern "TODO" --path spec/
"""

import argparse
import datetime
import json
import os
import re
import sys
import time
from concurrent.futures import ThreadPoolExecutor

if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8")


def parse_args():
    parser = argparse.ArgumentParser(
        description="High-performance cached content grepper and pattern searcher.",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  Search literal string across all cached files:
    python .lovable/ai-fix-scripts/09-fast-cached-grep.py --pattern "AppError"

  Regex search across TypeScript files:
    python .lovable/ai-fix-scripts/09-fast-cached-grep.py --pattern "enum .*Type" --regex --lang ts,tsx

  Search within specific subdirectory:
    python .lovable/ai-fix-scripts/09-fast-cached-grep.py --pattern "strictly-avoid" --path spec/

  Save findings to structured JSON in tmp/:
    python .lovable/ai-fix-scripts/09-fast-cached-grep.py --pattern "deprecated" --out tmp/grep-deprecated.json
        """,
    )
    parser.add_argument("--pattern", "-p", required=True, help="Search string or regex pattern")
    parser.add_argument("--regex", "-r", action="store_true", help="Treat pattern as regular expression")
    parser.add_argument("--case-sensitive", "-c", action="store_true", help="Perform case-sensitive matching")
    parser.add_argument("--lang", "-l", help="Language filter alias (e.g. go, ts, py, md)")
    parser.add_argument("--path", default=".", help="Subdirectory or root to scope search (default: .)")
    parser.add_argument("--out", "-o", default="tmp/grep-results.json", help="Output results file path (default: tmp/grep-results.json)")
    parser.add_argument("--limit", type=int, default=50, help="Max match results to print to console (default: 50)")
    parser.add_argument("--max-per-file", type=int, default=10, help="Max matches per file (default: 10)")
    parser.add_argument("--no-cache", action="store_true", help="Skip saving results to tmp/ cache")

    return parser.parse_args()


def load_file_list(scope_path, lang_filter):
    cache_path = "tmp/repo-file-cache.json"
    if os.path.exists(cache_path):
        try:
            with open(cache_path, "r", encoding="utf-8") as f:
                data = json.load(f)
            files = data.get("files", [])
            if files:
                # Apply path filter if specified
                if scope_path and scope_path != ".":
                    clean_p = scope_path.replace("\\", "/").rstrip("/") + "/"
                    files = [f for f in files if f.startswith(clean_p) or f == scope_path]
                return files
        except Exception:
            pass

    # Fallback to direct walk if cache missing
    all_files = []
    EXCLUDES = {".git", "node_modules", "dist", "build", ".next", ".cache", ".venv", "vendor", ".agent", ".gemini", "tmp"}
    for root, dirs, fnames in os.walk(scope_path):
        dirs[:] = [d for d in dirs if d not in EXCLUDES and not d.startswith(".")]
        for fn in fnames:
            if not fn.startswith("."):
                all_files.append(os.path.relpath(os.path.join(root, fn), ".").replace("\\", "/"))
    return all_files


def grep_single_file(filepath, matcher, is_regex, max_per_file):
    matches = []
    if not os.path.isfile(filepath):
        return filepath, matches

    try:
        with open(filepath, "r", encoding="utf-8", errors="replace") as f:
            for line_no, line in enumerate(f, 1):
                clean_line = line.rstrip("\r\n")
                found = False
                if is_regex:
                    if matcher.search(clean_line):
                        found = True
                else:
                    target = clean_line if (isinstance(matcher, str) and matcher != matcher.lower()) else clean_line.lower()
                    if matcher in target:
                        found = True

                if found:
                    matches.append({
                        "line": line_no,
                        "content": clean_line.strip()
                    })
                    if len(matches) >= max_per_file:
                        break
    except Exception:
        pass

    return filepath, matches


def main():
    args = parse_args()
    start_time = time.perf_counter()

    flags = 0 if args.case_sensitive else re.IGNORECASE
    matcher = re.compile(args.pattern, flags) if args.regex else (args.pattern if args.case_sensitive else args.pattern.lower())

    files_to_search = load_file_list(args.path, args.lang)
    results = {}
    total_matches = 0

    with ThreadPoolExecutor(max_workers=8) as executor:
        futures = [
            executor.submit(grep_single_file, fp, matcher, args.regex, args.max_per_file)
            for fp in files_to_search
        ]
        for future in futures:
            fp, file_matches = future.result()
            if file_matches:
                results[fp] = file_matches
                total_matches += len(file_matches)

    duration_ms = (time.perf_counter() - start_time) * 1000.0

    if not args.no_cache:
        os.makedirs(os.path.dirname(args.out), exist_ok=True)
        cache_output = {
            "timestamp": datetime.datetime.now(datetime.timezone.utc).isoformat(),
            "durationMs": round(duration_ms, 2),
            "pattern": args.pattern,
            "isRegex": args.regex,
            "totalMatches": total_matches,
            "totalFilesMatched": len(results),
            "results": results
        }
        with open(args.out, "w", encoding="utf-8") as f:
            json.dump(cache_output, f, indent=2)

    print("================================================================================")
    print(f"⚡ Fast Cached Grep: searched {len(files_to_search)} files in {duration_ms:.2f}ms")
    print(f"🔍 Pattern: `{args.pattern}` | Matches: **{total_matches}** in **{len(results)}** files")
    print("================================================================================")

    if not args.no_cache:
        print(f"💾 Results saved to: `{args.out}`\n")

    printed = 0
    for fp, matches in sorted(results.items()):
        print(f"📄 {fp} ({len(matches)} match{'es' if len(matches) > 1 else ''}):")
        for m in matches:
            print(f"   L{m['line']:<4}: {m['content']}")
            printed += 1
            if printed >= args.limit:
                break
        if printed >= args.limit:
            break

    if total_matches > printed:
        print(f"\n... and {total_matches - printed} more matches (see `{args.out}` for full output).")

    print("\n💡 AI AGENT INSTRUCTION:")
    print(f"   Read `{args.out}` for direct structured access to all match positions.")
    print("================================================================================")
    sys.exit(0)


if __name__ == "__main__":
    main()
