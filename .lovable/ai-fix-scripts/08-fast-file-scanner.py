#!/usr/bin/env python3
"""
Fast Repository File Scanner & Caching Engine.
Ultra-fast file scanner with multi-language filters, globbing, and persistent cache indexing in tmp/.

Usage:
  python .lovable/ai-fix-scripts/08-fast-file-scanner.py [options]

Examples:
  python .lovable/ai-fix-scripts/08-fast-file-scanner.py
  python .lovable/ai-fix-scripts/08-fast-file-scanner.py --lang go,ts,tsx
  python .lovable/ai-fix-scripts/08-fast-file-scanner.py --path spec/ --ext .md
  python .lovable/ai-fix-scripts/08-fast-file-scanner.py --search install --stats
"""

import argparse
import datetime
import json
import os
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
  Scan all text/source files:
    python .lovable/ai-fix-scripts/08-fast-file-scanner.py

  Filter by languages (comma-separated):
    python .lovable/ai-fix-scripts/08-fast-file-scanner.py --lang go,ts,tsx

  Filter by specific folder and extension:
    python .lovable/ai-fix-scripts/08-fast-file-scanner.py --path spec/ --ext .md

  Search for specific keyword in file paths:
    python .lovable/ai-fix-scripts/08-fast-file-scanner.py --search install --stats

  Output plaintext file list directly to tmp/:
    python .lovable/ai-fix-scripts/08-fast-file-scanner.py --format txt --out tmp/files.txt
        """,
    )
    parser.add_argument("--path", "-p", default=".", help="Subdirectory or root to scan (default: .)")
    parser.add_argument("--lang", "-l", help="Language filter alias (e.g. go, ts, py, md, or comma-separated go,ts)")
    parser.add_argument("--ext", "-e", help="Custom extension filter (comma-separated, e.g. .go,.ts,.json)")
    parser.add_argument("--search", "-s", help="Case-insensitive substring filter on file path")
    parser.add_argument("--out", "-o", default="tmp/repo-file-cache.json", help="Cache output path (default: tmp/repo-file-cache.json)")
    parser.add_argument("--format", "-f", choices=["json", "txt", "summary"], default="json", help="Output format (json, txt, summary)")
    parser.add_argument("--limit", type=int, default=100, help="Max file lines to print to console (default: 100)")
    parser.add_argument("--stats", action="store_true", help="Display extension statistics breakdown")
    parser.add_argument("--no-cache", action="store_true", help="Skip saving results to tmp/ cache")
    parser.add_argument("--include-hidden", action="store_true", help="Include dot-files/folders (normally ignored)")

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
        # Prune ignored directories in-place
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


def write_cache(cache_path, matched_files, ext_counts, args, scan_duration_ms):
    os.makedirs(os.path.dirname(cache_path), exist_ok=True)
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
        "stats": {
            "byExtension": dict(sorted(ext_counts.items(), key=lambda x: x[1], reverse=True))
        },
        "files": matched_files,
    }

    if cache_path.endswith(".json"):
        with open(cache_path, "w", encoding="utf-8") as f:
            json.dump(cache_data, f, indent=2)
    else:
        with open(cache_path, "w", encoding="utf-8") as f:
            for fp in matched_files:
                f.write(fp + "\n")

    # Also automatically write a plain text list in tmp/ for instant shell/script reading
    plain_txt_path = "tmp/repo-file-list.txt"
    try:
        with open(plain_txt_path, "w", encoding="utf-8") as f:
            for fp in matched_files:
                f.write(fp + "\n")
    except Exception:
        pass


def main():
    args = parse_args()
    start_time = time.perf_counter()

    allowed_exts = resolve_extensions(args.lang, args.ext)
    matched_files, ext_counts = scan_files(args.path, allowed_exts, args.search, args.include_hidden)
    scan_duration_ms = (time.perf_counter() - start_time) * 1000.0

    if not args.no_cache:
        write_cache(args.out, matched_files, ext_counts, args, scan_duration_ms)

    # Print summary & results to console
    print(f"⚡ Fast File Scanner: scanned in {scan_duration_ms:.2f}ms")
    print(f"📁 Root: `{args.path}` | Matched Files: **{len(matched_files)}**")
    if not args.no_cache:
        print(f"💾 Cache saved to: `{args.out}` (and `tmp/repo-file-list.txt`)")

    if args.stats or len(matched_files) == 0:
        print("\n📊 Extension Breakdown:")
        for ext, count in sorted(ext_counts.items(), key=lambda x: x[1], reverse=True):
            display_ext = ext if ext else "(no extension)"
            print(f"   • {display_ext:<12} : {count:>5} files")

    print("\n📋 File Inventory Preview:")
    preview_limit = args.limit if args.limit > 0 else len(matched_files)
    for idx, fp in enumerate(matched_files[:preview_limit], 1):
        print(f"   {idx:>4}. {fp}")

    if len(matched_files) > preview_limit:
        print(f"   ... and {len(matched_files) - preview_limit} more files (see `{args.out}` for full list).")

    print("\n💡 Tip for AI Agents: Read `tmp/repo-file-cache.json` or `tmp/repo-file-list.txt` for instant cached file access.")
    sys.exit(0)


if __name__ == "__main__":
    main()
