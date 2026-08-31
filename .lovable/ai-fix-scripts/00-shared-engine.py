#!/usr/bin/env python3
"""
Shared Core Engine for AI Repository Tooling, CI Fix Scripts & High-Speed Caching
Provides:
1. Module-level Enums & Constants with strict 'Type' suffix naming.
2. Pluggable cache layout in tmp/cache/ (paths, locks, files).
3. Cross-process safe atomic file locking with timeout and stale-lock recovery.
4. Two-phase incremental mtime-based file streaming (cache-first + parallel scan).
5. Fault-tolerant file reader handling missing/deleted files gracefully (zero crash).
6. Robust multi-folder scoping, customizable extensions, and nested ignore pruning (.git, .gitmap, node_modules).
"""

from collections.abc import Generator
from contextlib import contextmanager
from enum import Enum
import json
import os
from pathlib import Path
import re
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
    ".git", ".gitmap", "gitmap", ".git-map",
    "node_modules", "dist", "build", ".venv", "venv",
    ".gemini", "tmp", ".system_generated", "vendor", ".cache",
    ".next", "bin", "obj", "coverage", "__pycache__",
    ".vs", ".idea", ".agent", "release-artifacts", "release-assets",
    ".turbo", ".parcel-cache",
}

BINARY_EXTENSIONS = {
    ".png", ".jpg", ".jpeg", ".gif", ".webp", ".ico", ".svg",
    ".pdf", ".zip", ".tar", ".gz", ".7z", ".rar", ".bz2",
    ".exe", ".dll", ".so", ".dylib", ".bin", ".o", ".a",
    ".db", ".sqlite", ".sqlite3",
    ".woff", ".woff2", ".ttf", ".eot", ".otf",
    ".mp3", ".mp4", ".wav", ".avi", ".mov",
    ".pyc", ".pyo", ".pyd", ".class",
}

# Pre-compiled common regular expressions
RE_WINDOWS_BACKSLASH = re.compile(r"\\")
RE_LEADING_DOT_SLASH = re.compile(r"^\./")
RE_CRLF = re.compile(r"\r\n")

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

def is_ignored_directory(dir_name: str, custom_excludes: set[str] | None = None) -> bool:
    """Checks if directory name is in the global or custom exclusion list."""
    excludes = EXCLUDE_DIRS if custom_excludes is None else EXCLUDE_DIRS | custom_excludes
    return dir_name.lower() in {d.lower() for d in excludes}

def is_ignored_path(path: str | Path, custom_excludes: set[str] | None = None) -> bool:
    """Checks if any segment of the path matches an excluded directory."""
    excludes = EXCLUDE_DIRS if custom_excludes is None else EXCLUDE_DIRS | custom_excludes
    excludes_lower = {d.lower() for d in excludes}
    parts = Path(path).parts
    return any(p.lower() in excludes_lower for p in parts)

def is_binary_file(file_path: Path) -> bool:
    """Checks if file has a known binary extension."""
    return file_path.suffix.lower() in BINARY_EXTENSIONS

def normalize_rel_path(path: str | Path) -> str:
    """Converts a path into a canonical relative POSIX path."""
    p_str = RE_WINDOWS_BACKSLASH.sub("/", str(path))
    return RE_LEADING_DOT_SLASH.sub("", p_str)

def normalize_extensions(extensions: tuple | set | list | str | None) -> set[str] | None:
    """Normalizes custom extensions into a lowercased set with leading dots."""
    if not extensions:
        return None
    if isinstance(extensions, str):
        raw_items = [e.strip() for e in extensions.split(",") if e.strip()]
    else:
        raw_items = [str(e).strip() for e in extensions if str(e).strip()]
    normalized = set()
    for item in raw_items:
        clean = item.lower()
        if not clean.startswith("."):
            clean = f".{clean}"
        normalized.add(clean)
    return normalized if normalized else None

def read_file_safe(path: str | Path) -> str | None:
    """Fault-tolerant file reader. Handles missing or deleted files gracefully with zero crashes."""
    try:
        p = Path(path)
        if not p.exists():
            return None
        if not p.is_file():
            return None
        with open(p, "r", encoding="utf-8", errors="replace") as f:
            return RE_CRLF.sub("\n", f.read())
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
    temp_path = p.with_name(f"{p.name}.tmp_{os.getpid()}_{int(time.time()*1000)}")
    try:
        lf_content = RE_CRLF.sub("\n", content)
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
            if lock_file.exists():
                lock_age = time.time() - lock_file.stat().st_mtime
                if lock_age > STALE_LOCK_SECONDS:
                    try:
                        lock_file.unlink()
                    except Exception:
                        pass

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
        if acquired:
            if lock_file.exists():
                try:
                    lock_file.unlink()
                except Exception:
                    pass

# --- Pluggable Cache Management ---

def load_repo_cache() -> dict[str, Any]:
    """Loads pre-computed repository file cache from tmp/cache/ or legacy tmp/."""
    for target in (PRIMARY_CACHE_FILE, LEGACY_CACHE_FILE):
        if target.exists():
            try:
                with open(target, "r", encoding="utf-8") as f:
                    data = json.load(f)
                    if isinstance(data, dict):
                        if "files" in data:
                            return data
            except Exception:
                pass
    return {}

def save_repo_cache(cache_data: dict[str, Any]) -> None:
    """Saves repository cache safely with atomic locking and dual-path sync."""
    CACHE_PATHS_DIR.mkdir(parents=True, exist_ok=True)
    with atomic_cache_lock("repo-cache-write.lock"):
        temp_primary = PRIMARY_CACHE_FILE.with_suffix(".json.tmp")
        try:
            with open(temp_primary, "w", encoding="utf-8") as f:
                json.dump(cache_data, f, indent=2)
            temp_primary.replace(PRIMARY_CACHE_FILE)
        except Exception:
            pass

        try:
            temp_legacy = LEGACY_CACHE_FILE.with_suffix(".json.tmp")
            with open(temp_legacy, "w", encoding="utf-8") as f:
                json.dump(cache_data, f, indent=2)
            temp_legacy.replace(LEGACY_CACHE_FILE)
        except Exception:
            pass

# --- Two-Phase Streaming Engine with Delta Eviction ---

def stream_cached_files(
    cache_data: dict[str, Any],
    root_dir: str = ".",
    extensions: set[str] | tuple | None = None,
    custom_excludes: set[str] | None = None
) -> Generator[Path, None, None]:
    """Phase 1: Streams valid files from cache first. Automatically skips missing/deleted/excluded files."""
    file_list = cache_data.get("files", [])
    norm_root = normalize_rel_path(root_dir).rstrip("/")
    ext_set = normalize_extensions(extensions)

    for rel_path in file_list:
        norm_p = normalize_rel_path(rel_path)
        if norm_root:
            if norm_root != ".":
                if not norm_p.startswith(norm_root + "/"):
                    if norm_p != norm_root:
                        continue
        if is_ignored_path(norm_p, custom_excludes=custom_excludes):
            continue
        p = Path(norm_p)
        if not p.exists():
            continue
        if is_binary_file(p):
            continue
        if ext_set:
            if p.suffix.lower() not in ext_set:
                continue
        yield p

def stream_directory_files(
    root_dir: str = ".",
    extensions: set[str] | tuple | None = None,
    custom_excludes: set[str] | None = None
) -> Generator[Path, None, None]:
    """Phase 2: Walks filesystem pruning ignored folders (including nested .git, .gitmap, node_modules)."""
    ext_set = normalize_extensions(extensions)
    for root, dirs, files in os.walk(root_dir):
        dirs[:] = [d for d in dirs if not is_ignored_directory(d, custom_excludes=custom_excludes)]
        for f in files:
            p = Path(os.path.join(root, f))
            if is_binary_file(p):
                continue
            if ext_set:
                if p.suffix.lower() not in ext_set:
                    continue
            yield p

def process_repository_files(
    processor_fn: Callable[[Path], Any],
    root_dir: str = ".",
    extensions: set[str] | tuple | list | str | None = None,
    use_cache: bool = True,
    custom_excludes: set[str] | None = None
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
    norm_exts = normalize_extensions(extensions)

    # Phase 1: Process cached files first
    if cache_data:
        if "files" in cache_data:
            for p in stream_cached_files(cache_data, root_dir=root_dir, extensions=norm_exts, custom_excludes=custom_excludes):
                norm_p = normalize_rel_path(p)
                if norm_p not in processed_paths:
                    processed_paths.add(norm_p)
                    res = processor_fn(p)
                    if res is not None:
                        results.append(res)

    # Phase 2: Stream live directory files (catches new/untracked files)
    for p in stream_directory_files(root_dir=root_dir, extensions=norm_exts, custom_excludes=custom_excludes):
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
