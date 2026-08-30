#!/usr/bin/env python3
"""
Shared Core Engine for AI Repository Tooling & CI Fix Scripts
Provides:
1. Module-level Enums & Constants with strict 'Type' suffix naming.
2. Fast two-phase incremental mtime-based file streaming (cache-first + parallel scan).
3. Universal path normalization, safe UTF-8 LF file I/O, and small decomposed utility functions.
"""

from collections.abc import Generator
from enum import Enum
import json
import os
from pathlib import Path
import sys
import time
from typing import Any, Callable

if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8")

# --- Module-Level Constants ---
CACHE_DIR = Path("tmp")
CACHE_FILE_PATH = CACHE_DIR / "repo-file-cache.json"
DEFAULT_MAX_FILE_KB = 1024
EXCLUDE_DIRS = {".git", "node_modules", "dist", "build", ".venv", ".gemini", "tmp", ".system_generated"}
BINARY_EXTENSIONS = {
    ".png", ".jpg", ".jpeg", ".gif", ".webp", ".ico", ".svg",
    ".pdf", ".zip", ".tar", ".gz", ".7z", ".rar", ".bz2",
    ".exe", ".dll", ".so", ".dylib", ".bin", ".o", ".a",
    ".db", ".sqlite", ".sqlite3",
    ".woff", ".woff2", ".ttf", ".eot", ".otf",
}

# --- Top-Level Enums ---
class ScanModeType(str, Enum):
    CHECK = "check"
    FIX = "fix"
    STREAM = "stream"

class SeverityType(str, Enum):
    BLOCKER = "blocker"
    HIGH = "high"
    WARN = "warn"
    INFO = "info"

class ExitCodeType(int, Enum):
    SUCCESS = 0
    VIOLATIONS_FOUND = 1
    TOOL_ERROR = 2

# --- Path & File Utility Functions ---

def is_ignored_directory(dir_name: str) -> bool:
    """Checks if directory name is in the global exclusion list."""
    return dir_name in EXCLUDE_DIRS

def is_binary_file(file_path: Path) -> bool:
    """Checks if file has a known binary extension."""
    return file_path.suffix.lower() in BINARY_EXTENSIONS

def normalize_rel_path(path: str | Path) -> str:
    """Converts a path into a canonical relative POSIX path."""
    p_str = str(path).replace("\\", "/")
    if p_str.startswith("./"):
        p_str = p_str[2:]
    return p_str

def read_file_lf(path: str | Path) -> str:
    """Reads a text file as UTF-8, stripping any carriage returns to ensure strict UNIX LF."""
    with open(path, "r", encoding="utf-8", errors="replace") as f:
        return f.read().replace("\r\n", "\n")

def write_file_lf(path: str | Path, content: str) -> bool:
    """Writes text content ensuring UTF-8 encoding and strict UNIX LF line endings."""
    lf_content = content.replace("\r\n", "\n")
    with open(path, "wb") as f:
        f.write(lf_content.encode("utf-8"))
    return True

# --- Cache Management Functions ---

def load_repo_cache() -> dict[str, Any]:
    """Loads pre-computed repository file cache from tmp/."""
    if not CACHE_FILE_PATH.exists():
        return {}
    try:
        with open(CACHE_FILE_PATH, "r", encoding="utf-8") as f:
            return json.load(f)
    except Exception:
        return {}

def save_repo_cache(cache_data: dict[str, Any]) -> None:
    """Saves repository cache safely into tmp/."""
    CACHE_DIR.mkdir(parents=True, exist_ok=True)
    with open(CACHE_FILE_PATH, "w", encoding="utf-8") as f:
        json.dump(cache_data, f, indent=2)

def get_file_metadata(path: Path) -> dict[str, Any]:
    """Retrieves file stat metadata for cache validation."""
    st = path.stat()
    return {
        "mtime": st.st_mtime,
        "size": st.st_size,
    }

# --- Two-Phase Streaming Engine ---

def stream_cached_files(cache_data: dict[str, Any], extensions: tuple = None) -> Generator[Path, None, None]:
    """Phase 1: Streams files known in the cache first for immediate processing."""
    file_list = cache_data.get("files", [])
    for rel_path in file_list:
        p = Path(rel_path)
        if not p.exists() or is_binary_file(p):
            continue
        if extensions and p.suffix.lower() not in extensions:
            continue
        yield p

def stream_directory_files(root_dir: str = ".", extensions: tuple = None) -> Generator[Path, None, None]:
    """Phase 2: Walks filesystem using os.scandir, streaming entries as they are discovered."""
    for root, dirs, files in os.walk(root_dir):
        dirs[:] = [d for d in dirs if not is_ignored_directory(d)]
        for f in files:
            p = Path(os.path.join(root, f))
            if is_binary_file(p):
                continue
            if extensions and p.suffix.lower() not in extensions:
                continue
            yield p

def process_repository_files(
    processor_fn: Callable[[Path], Any],
    root_dir: str = ".",
    extensions: tuple = None,
    use_cache: bool = True
) -> dict[str, Any]:
    """
    Two-Phase Universal Pipeline:
    1. Starts immediately with cached files if available.
    2. Streams and discovers new / modified files on disk.
    3. Executes processor_fn on each unique file and aggregates statistics.
    """
    start_time = time.perf_counter()
    cache_data = load_repo_cache() if use_cache else {}
    processed_paths: set[str] = set()
    results = []

    # Phase 1: Process cached files first
    if cache_data and "files" in cache_data:
        for p in stream_cached_files(cache_data, extensions=extensions):
            norm_p = normalize_rel_path(p)
            if norm_p not in processed_paths:
                processed_paths.add(norm_p)
                res = processor_fn(p)
                if res is not None:
                    results.append(res)

    # Phase 2: Stream live directory files (catches new/untracked files)
    for p in stream_directory_files(root_dir=root_dir, extensions=extensions):
        norm_p = normalize_rel_path(p)
        if norm_p not in processed_paths:
            processed_paths.add(norm_p)
            res = processor_fn(p)
            if res is not None:
                results.append(res)

    elapsed_ms = (time.perf_counter() - start_time) * 1000

    return {
        "total_files": len(processed_paths),
        "results": results,
        "elapsed_ms": elapsed_ms,
    }
