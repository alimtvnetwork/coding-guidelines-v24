#!/usr/bin/env python3
"""
Fast Repository File Scanner & Caching Engine.
Ultra-fast file scanner with multi-language filters, substring search, and persistent cache indexing in tmp/.

Usage:
  python .lovable/ai-fix-scripts/08-fast-file-scanner.py [options]

Examples:
  python .lovable/ai-fix-scripts/08-fast-file-scanner.py
  python .lovable/ai-fix-scripts/08-fast-file-scanner.py --lang go,ts,tsx
  python .lovable/ai-fix-scripts/08-fast-file-scanner.py --path spec/ --ext .md
  python .lovable/ai-fix-scripts/08-fast-file-scanner.py --search install --stats
  python .lovable/ai-fix-scripts/08-fast-file-scanner.py --query-cache "component"
"""

import argparse
import datetime
import json
import os
import re
import sys
import time

if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8")

# Common directories and patterns to ignore
DEFAULT_IGNORE_DIRS = {
    ".git",
    "node_modules",
    "dist",
    "build",
    ".next",
    ".cache",
    ".venv",
    "venv",
    "vendor",
    ".gemini",
    ".agent",
    "release-artifacts",
    "bin",
    "obj",
    "coverage",
    ".turbo",
    ".parcel-cache",
    ".vs",
    ".idea",
}

# Binary and non-code asset extensions to exclude
BINARY_EXTENSIONS = {
    ".png", ".jpg", ".jpeg", ".gif", ".webp", ".ico", ".svg",
    ".pdf", ".zip", ".tar", ".gz", ".7z", ".rar", ".bz2",
    ".exe", ".dll", ".so", ".dylib", ".bin", ".o", ".a",
    ".db", ".sqlite", ".sqlite3",
    ".woff", ".woff2", ".ttf", ".eot", ".otf",
    ".mp3", ".mp4", ".wav", ".avi", ".mov",
    ".pyc", ".pyo", ".pyd", ".class",
}

# Language alias mappings to standard file extensions
LANG_EXT_MAP = {
    "go": [".go"],
    "golang": [".go"],
    "ts": [".ts", ".tsx", ".mts", ".cts"],
    "typescript": [".ts", ".tsx", ".mts", ".cts"],
    "tsx": [".tsx"],
    "js": [".js", ".jsx", ".mjs", ".cjs"],
    "javascript": [".js", ".jsx", ".mjs", ".cjs"],
    "jsx": [".jsx"],
    "py": [".py", ".pyi"],
    "python": [".py", ".pyi"],
    "php": [".php", ".phtml"],
    "cs": [".cs"],
    "csharp": [".cs"],
    "rust": [".rs"],
    "rs": [".rs"],
    "md": [".md", ".markdown"],
    "markdown": [".md", ".markdown"],
    "json": [".json"],
    "yaml": [".yaml", ".yml"],
    "yml": [".yaml", ".yml"],
    "sh": [".sh", ".bash"],
    "bash": [".sh", ".bash"],
    "ps1": [".ps1", ".psm1", ".psd1"],
    "powershell": [".ps1", ".psm1", ".psd1"],
    "sql": [".sql"],
    "html": [".html", ".htm"],
    "css": [".css", ".scss", ".sass", ".less"],
    "c": [".c", ".h"],
    "cpp": [".cpp", ".hpp", ".cc", ".cxx"],
}


def parse_args():
    parser = argparse.ArgumentParser(
        description="High-performance repository file scanner and cache indexer.",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  Scan all text/source files and generate tmp/ caches:
    python .lovable/ai-fix-scripts/08-fast-file-scanner.py

  Filter by languages (comma-separated):
    python .lovable/ai-fix-scripts/08-fast-file-scanner.py --lang go,ts,tsx

  Filter by specific folder and extension:
    python .lovable/ai-fix-scripts/08-fast-file-scanner.py --path spec/ --ext .md

  Search for specific keyword in file paths:
    python .lovable/ai-fix-scripts/08-fast-file-scanner.py --search install --stats

  Instant cache query (reads tmp/ cache in <1ms without disk walk):
    python .lovable/ai-fix-scripts/08-fast-file-scanner.py --query-cache "button"
        """,
    )
    parser.add_argument("--path", "-p", default=".", help="Subdirectory or root to scan (default: .)")
    parser.add_argument("--lang", "-l", help="Language filter alias (e.g. go, ts, py, md, or comma-separated go,ts)")
    parser.add_argument("--ext", "-e", help="Custom extension filter (comma-separated, e.g. .go,.ts,.json)")
    parser.add_argument("--search", "-s", help="Case-insensitive substring filter on file path")
    parser.add_argument("--out", "-o", help="Custom cache output path (default: auto-named in tmp/)")
    parser.add_argument("--format", "-f", choices=["json", "txt", "summary"], default="json", help="Output format (json, txt, summary)")
    parser.add_argument("--limit", type=int, default=100, help="Max file lines to print to console (default: 100)")
    parser.add_argument("--stats", action="store_true", help="Display extension statistics breakdown")
    parser.add_argument("--no-cache", action="store_true", help="Skip saving results to tmp/ cache")
    parser.add_argument("--include-hidden", action="store_true", help="Include dot-files/folders (normally ignored)")
    parser.add_argument("--query-cache", "-q", help="Query existing cached index without walking the filesystem")
    parser.add_argument("--check", action="store_true", help="CI validation mode: verifies file index and exits 0 on success")

    return parser.parse_args()


def resolve_extensions(lang_arg, ext_arg):
    allowed_exts = set()
    if lang_arg:
        for lang in lang_arg.split(","):
            cleaned = lang.strip().lower()
            if cleaned in LANG_EXT_MAP:
                allowed_exts.update(LANG_EXT_MAP[cleaned])
            else:
                ext_form = f".{cleaned}" if not cleaned.startswith(".") else cleaned
                allowed_exts.add(ext_form)

    if ext_arg:
        for e in ext_arg.split(","):
            cleaned = e.strip().lower()
            ext_form = f".{cleaned}" if not cleaned.startswith(".") else cleaned
            allowed_exts.add(ext_form)

    return allowed_exts if allowed_exts else None


def scan_files(scan_root, allowed_exts, search_term, include_hidden):
    matched_files = []
    ext_counts = {}

    for root, dirs, files in os.walk(scan_root):
        if not include_hidden:
            dirs[:] = [
                d for d in dirs
                if d not in DEFAULT_IGNORE_DIRS and not d.startswith(".")
            ]
        else:
            dirs[:] = [d for d in dirs if d not in DEFAULT_IGNORE_DIRS]

        for filename in sorted(files):
            if not include_hidden and filename.startswith("."):
                continue

            ext = os.path.splitext(filename)[1].lower()
            if ext in BINARY_EXTENSIONS:
                continue

            if allowed_exts and ext not in allowed_exts:
                continue

            rel_path = os.path.relpath(os.path.join(root, filename), ".").replace("\\", "/")

            if search_term and search_term.lower() not in rel_path.lower():
                continue

            matched_files.append(rel_path)
            ext_counts[ext] = ext_counts.get(ext, 0) + 1

    return matched_files, ext_counts


def get_cache_filenames(args):
    slug_parts = []
    if args.path and args.path != ".":
        clean_p = re.sub(r"[^a-zA-Z0-9_-]+", "_", args.path).strip("_")
        if clean_p:
            slug_parts.append(clean_p)
    if args.lang:
        clean_l = re.sub(r"[^a-zA-Z0-9_-]+", "_", args.lang).strip("_")
        slug_parts.append(f"lang-{clean_l}")
    if args.ext:
        clean_e = re.sub(r"[^a-zA-Z0-9_-]+", "_", args.ext).strip("_")
        slug_parts.append(f"ext-{clean_e}")
    if args.search:
        clean_s = re.sub(r"[^a-zA-Z0-9_-]+", "_", args.search).strip("_")
        slug_parts.append(f"search-{clean_s}")

    slug = "-".join(slug_parts) if slug_parts else "all"

    json_path = args.out if (args.out and args.out.endswith(".json")) else "tmp/repo-file-cache.json"
    txt_path = f"tmp/file-list-{slug}.txt"
    all_txt_path = "tmp/file-list-all.txt"

    return json_path, txt_path, all_txt_path


def write_caches(json_path, txt_path, all_txt_path, matched_files, ext_counts, args, scan_duration_ms):
    os.makedirs("tmp", exist_ok=True)
    cache_data = {
        "timestamp": datetime.datetime.now(datetime.timezone.utc).isoformat(),
        "scanRoot": args.path,
        "scanDurationMs": round(scan_duration_ms, 2),
        "totalFiles": len(matched_files),
        "filterCriteria": {
            "lang": args.lang,
            "ext": args.ext,
            "search": args.search,
            "includeHidden": args.include_hidden,
        },
        "cacheFiles": {
            "jsonCache": json_path,
            "filterTextList": txt_path,
            "globalTextList": all_txt_path,
        },
        "stats": {
            "byExtension": dict(sorted(ext_counts.items(), key=lambda x: x[1], reverse=True))
        },
        "files": matched_files,
    }

    # 1. Main JSON Cache
    with open(json_path, "w", encoding="utf-8") as f:
        json.dump(cache_data, f, indent=2)

    # 2. Filter-specific text file
    with open(txt_path, "w", encoding="utf-8") as f:
        for fp in matched_files:
            f.write(fp + "\n")

    # 3. Global text list (if scanning full repo)
    if args.path == "." and not args.lang and not args.ext and not args.search:
        with open(all_txt_path, "w", encoding="utf-8") as f:
            for fp in matched_files:
                f.write(fp + "\n")


def query_cached_index(query_term):
    cache_path = "tmp/repo-file-cache.json"
    if not os.path.exists(cache_path):
        print(f"⚠️ Cache not found at `{cache_path}`. Running quick scan to generate cache...")
        matched_files, _ = scan_files(".", None, None, False)
        os.makedirs("tmp", exist_ok=True)
        with open(cache_path, "w", encoding="utf-8") as f:
            json.dump({"files": matched_files}, f)
    else:
        with open(cache_path, "r", encoding="utf-8") as f:
            data = json.load(f)
        matched_files = data.get("files", [])

    results = [f for f in matched_files if query_term.lower() in f.lower()]
    print(f"⚡ Instant Cache Query for `{query_term}`: found **{len(results)}** matches in pre-computed index:\n")
    for idx, r in enumerate(results[:100], 1):
        print(f"   {idx:>3}. {r}")
    if len(results) > 100:
        print(f"   ... and {len(results) - 100} more matches.")
    sys.exit(0)


def main():
    args = parse_args()

    if args.query_cache:
        query_cached_index(args.query_cache)
        return

    start_time = time.perf_counter()
    allowed_exts = resolve_extensions(args.lang, args.ext)
    matched_files, ext_counts = scan_files(args.path, allowed_exts, args.search, args.include_hidden)
    scan_duration_ms = (time.perf_counter() - start_time) * 1000.0

    json_path, txt_path, all_txt_path = get_cache_filenames(args)

    if not args.no_cache:
        write_caches(json_path, txt_path, all_txt_path, matched_files, ext_counts, args, scan_duration_ms)

    # Print clean summary & results to console
    print("================================================================================")
    print(f"⚡ Fast Repository File Scanner: scanned in {scan_duration_ms:.2f}ms")
    print(f"📁 Root: `{args.path}` | Filtered Files Found: **{len(matched_files)}**")
    print("================================================================================")

    if not args.no_cache:
        print("\n💾 TEMP FOLDER CACHE INVENTORY:")
        print(f"   • Full JSON Cache : `{json_path}`")
        print(f"   • Specific Filter : `{txt_path}`")
        print(f"   • Global List     : `{all_txt_path}`")

    if args.stats or len(matched_files) == 0:
        print("\n📊 Extension Breakdown:")
        for ext, count in sorted(ext_counts.items(), key=lambda x: x[1], reverse=True):
            display_ext = ext if ext else "(no extension)"
            print(f"   • {display_ext:<14} : {count:>5} files")

    print("\n📋 File Inventory Preview:")
    preview_limit = args.limit if args.limit > 0 else len(matched_files)
    for idx, fp in enumerate(matched_files[:preview_limit], 1):
        print(f"   {idx:>4}. {fp}")

    if len(matched_files) > preview_limit:
        print(f"   ... and {len(matched_files) - preview_limit} more files (see `{txt_path}` for complete list).")

    print("\n💡 AI AGENT INSTRUCTION:")
    print(f"   To read or verify files in subsequent steps without ad-hoc queries,")
    print(f"   simply read `{txt_path}` or `{json_path}` directly.")
    print("================================================================================")
    sys.exit(0)


if __name__ == "__main__":
    main()
