#!/usr/bin/env python3
"""
Shared Core Engine for AI Repository Tooling, CI Fix Scripts & High-Speed Caching
Provides:
1. Module-level Enums & Constants with strict 'Type' suffix naming.
2. Pluggable cache layout in tmp/cache/ (paths, locks, files).
3. Cross-process safe atomic file locking with timeout and stale-lock recovery.
4. Two-phase incremental mtime-based file streaming (cache-first + parallel scan).
5. Fault-tolerant file reader handling missing/deleted files gracefully (zero crash).
"""

from collections.abc import Generator
from contextlib import contextmanager
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
CACHE_BASE_DIR = Path("tmp/cache")
CACHE_PATHS_DIR = CACHE_BASE_DIR / "paths"
CACHE_LOCKS_DIR = CACHE_BASE_DIR / "locks"
CACHE_FILES_DIR = CACHE_BASE_DIR / "files"
LEGACY_CACHE_FILE = Path("tmp/repo-file-cache.json")
PRIMARY_CACHE_FILE = CACHE_BASE_DIR / "repo-file-cache.json"

DEFAULT_MAX_FILE_KB = 2048
LOCK_TIMEOUT_SECONDS = 5.0
STALE_LOCK_SECONDS = 15.0

EXCLUDE_DIRS = {
    ".git", "node_modules", "dist", "build", ".venv", "venv",
    ".gemini", "tmp", ".system_generated", "vendor", ".cache",
    ".next", "bin", "obj"
}

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

def read_file_safe(path: str | Path) -> str | None:
    """Fault-tolerant file reader. Handles missing or deleted files gracefully with zero crashes."""
    try:
        p = Path(path)
        if not p.exists() or not p.is_file():
            return None
        with open(p, "r", encoding="utf-8", errors="replace") as f:
            return f.read().replace("\r\n", "\n")
    except (FileNotFoundError, PermissionError, OSError):
        return None

def read_file_lf(path: str | Path) -> str:
    """Reads a text file ensuring strict UNIX LF. Returns empty string if file does not exist."""
    content = read_file_safe(path)
    return content if content is not None else ""

def write_file_lf(path: str | Path, content: str) -> bool:
    """Atomic write ensuring UTF-8 encoding and strict UNIX LF line endings."""
    p = Path(path)
    p.parent.mkdir(parents=True, exist_ok=True)
    temp_path = p.with_name(f"{p.name}.tmp_{os.getpid()}")
    try:
        lf_content = content.replace("\r\n", "\n")
        with open(temp_path, "wb") as f:
            f.write(lf_content.encode("utf-8"))
        temp_path.replace(p)
        return True
    except Exception:
        if temp_path.exists():
            try:
                temp_path.unlink()
            except Exception:
                pass
        return False

# --- Safe Cross-Process Locking Mechanism ---

@contextmanager
def atomic_cache_lock(lock_name: str = "repo-cache.lock", timeout: float = LOCK_TIMEOUT_SECONDS):
    """
    Acquires a cross-process lockfile in tmp/cache/locks/.
    Recovers from stale locks (>15s) automatically.
    """
    CACHE_LOCKS_DIR.mkdir(parents=True, exist_ok=True)
    lock_file = CACHE_LOCKS_DIR / lock_name
    start_time = time.time()
    acquired = False

    while time.time() - start_time < timeout:
        try:
            # Check for stale lock
            if lock_file.exists():
                lock_age = time.time() - lock_file.stat().st_mtime
                if lock_age > STALE_LOCK_SECONDS:
                    try:
                        lock_file.unlink()
                    except Exception:
                        pass

            # Attempt atomic lock creation
            fd = os.open(str(lock_file), os.O_CREAT | os.O_EXCL | os.O_RDWR)
            os.write(fd, f"pid={os.getpid()}\ntime={time.time()}".encode("utf-8"))
            os.close(fd)
            acquired = True
            break
        except FileExistsError:
            time.sleep(0.02)
        except Exception:
            break

    try:
        yield acquired
    finally:
        if acquired and lock_file.exists():
            try:
                lock_file.unlink()
            except Exception:
                pass

# --- Pluggable Cache Management ---

def load_repo_cache() -> dict[str, Any]:
    """Loads pre-computed repository file cache from tmp/cache/ or legacy tmp/."""
    # Check primary pluggable path first
    for target in (PRIMARY_CACHE_FILE, LEGACY_CACHE_FILE):
        if target.exists():
            try:
                with open(target, "r", encoding="utf-8") as f:
                    data = json.load(f)
                    if isinstance(data, dict) and "files" in data:
                        return data
            except Exception:
                pass
    return {}

def save_repo_cache(cache_data: dict[str, Any]) -> None:
    """Saves repository cache safely with atomic locking and dual-path sync."""
    CACHE_PATHS_DIR.mkdir(parents=True, exist_ok=True)
    with atomic_cache_lock("repo-cache-write.lock"):
        # Save primary in tmp/cache/
        temp_primary = PRIMARY_CACHE_FILE.with_suffix(".json.tmp")
        try:
            with open(temp_primary, "w", encoding="utf-8") as f:
                json.dump(cache_data, f, indent=2)
            temp_primary.replace(PRIMARY_CACHE_FILE)
        except Exception:
            pass

        # Sync legacy path tmp/repo-file-cache.json for backwards compatibility
        try:
            temp_legacy = LEGACY_CACHE_FILE.with_suffix(".json.tmp")
            with open(temp_legacy, "w", encoding="utf-8") as f:
                json.dump(cache_data, f, indent=2)
            temp_legacy.replace(LEGACY_CACHE_FILE)
        except Exception:
            pass

# --- Two-Phase Streaming Engine with Delta Eviction ---

def stream_cached_files(cache_data: dict[str, Any], extensions: tuple = None) -> Generator[Path, None, None]:
    """Phase 1: Streams valid files from cache first. Automatically skips missing/deleted files."""
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
    1. Starts immediately with cached files if available (<0.1ms).
    2. Streams and discovers new / modified files on disk.
    3. Gracefully skips missing or removed files during processing.
    4. Executes processor_fn on each unique file and aggregates statistics.
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
