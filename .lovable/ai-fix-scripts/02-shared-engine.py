#!/usr/bin/env python3
"""
Shared Core Engine for AI Repository Tooling, CI Fix Scripts & High-Speed Caching
Dual-Platform Engine (100% Native Unix & Windows Support)

Features:
1. Centralized Constants for Encodings, Separators, Tokens, and Configuration Maps.
2. Top-Level Enums (PascalCase class, UPPER_CASE members mirroring string values).
3. Thread-Safe Lazy Regex Registry (Singleton Double-Checked Locking).
4. Dual-Mode Cross-Process Locking:
   - POSIX: Kernel-level `fcntl.flock` (automatic cleanup on process kill/crash).
   - Windows: Atomic `os.O_CREAT | os.O_EXCL` with PID timestamp & stale-lock recovery.
5. Unix Symlink & Cycle Guard: Inode tracking (st_dev, st_ino) preventing infinite recursion.
6. Unix Permission Preservation: Preserves executable bits (chmod +x / st_mode) across atomic writes.
7. Universal Line Ending Normalizer: Aggressively converts CRLF (\\r\\n) and legacy Mac CR (\\r) to clean UNIX LF (\\n).
8. Memory-Safe Chunked Binary Probe: Inspects first 8KB for null-bytes without loading large blobs into RAM.
9. Two-phase incremental mtime-based file streaming (cache-first + parallel scan).
10. Pluggable cache layout in tmp/cache/ (paths, locks, files).
11. Fault-tolerant file reader handling missing/deleted files gracefully (zero crash).
"""

from collections.abc import Generator
from contextlib import contextmanager
from enum import Enum
import json
import os
from pathlib import Path
import re
import sys
import threading
import time
from typing import Any, Callable

# Optional POSIX kernel locking
try:
    import fcntl
    IS_FCNTL_AVAILABLE = True
except ImportError:
    fcntl = None
    IS_FCNTL_AVAILABLE = False

if hasattr(sys.stdout, "reconfigure"):
    sys.stdout.reconfigure(encoding="utf-8")

# --- Centralized Constants for Encodings, Separators & Tokens ---
DEFAULT_ENCODING = "utf-8"
UTF16_ENCODING = "utf-16"
UTF16_LE_ENCODING = "utf-16le"
UTF16_BE_ENCODING = "utf-16be"

LINE_SEPARATOR = "\n"
CARRIAGE_RETURN = "\r"
CRLF_SEPARATOR = "\r\n"
TAB_CHAR = "\t"
PATH_SEPARATOR = "/"
WINDOWS_PATH_SEPARATOR = "\\"

# --- Module-Level Directory & File Constants ---
CACHE_BASE_DIR = Path("tmp/cache")
CACHE_PATHS_DIR = CACHE_BASE_DIR / "paths"
CACHE_LOCKS_DIR = CACHE_BASE_DIR / "locks"
CACHE_FILES_DIR = CACHE_BASE_DIR / "files"
LEGACY_CACHE_FILE = Path("tmp/repo-file-cache.json")
PRIMARY_CACHE_FILE = CACHE_BASE_DIR / "repo-file-cache.json"

DEFAULT_MAX_FILE_KB = 2048
LOCK_TIMEOUT_SECONDS = 5.0
STALE_LOCK_SECONDS = 15.0
MAX_READ_SIZE_BYTES = 20 * 1024 * 1024  # 20MB memory safety cap

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

DEFAULT_TEXT_EXTENSIONS = (
    ".md", ".markdown", ".py", ".ts", ".tsx", ".js", ".jsx",
    ".json", ".yaml", ".yml", ".go", ".php", ".cs", ".sh", ".ps1"
)

DEFAULT_CODE_EXTENSIONS = (
    ".ts", ".tsx", ".js", ".jsx", ".go", ".py", ".php", ".cs"
)

DEFAULT_CLI_EXTENSIONS = (
    ".go", ".ts", ".tsx", ".py", ".php"
)

ALLOWED_LARGE_FILES = {
    "src/data/specTree.json",
    "src\\data\\specTree.json",
    "slides-app/dist.zip",
    "slides-app\\dist.zip",
}

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

# --- Top-Level Enums Following Standard ---
class ScanModeType(str, Enum):
    """Enumeration for file scanning modes."""
    CHECK = "CHECK"
    FIX = "FIX"
    STREAM = "STREAM"

class SeverityType(str, Enum):
    """Enumeration for issue severity levels."""
    BLOCKER = "BLOCKER"
    HIGH = "HIGH"
    WARN = "WARN"
    INFO = "INFO"

class ExitCodeType(int, Enum):
    """Enumeration for application exit codes."""
    SUCCESS = 0
    VIOLATIONS_FOUND = 1
    TOOL_ERROR = 2

class RegexPatternType(str, Enum):
    """Enumeration for cached regex pattern identifiers."""
    WINDOWS_BACKSLASH = "WINDOWS_BACKSLASH"
    LEADING_DOT_SLASH = "LEADING_DOT_SLASH"
    CRLF = "CRLF"
    UNIVERSAL_LINE_ENDING = "UNIVERSAL_LINE_ENDING"
    TRAILING_WHITESPACE = "TRAILING_WHITESPACE"
    SEQ_PREFIX = "SEQ_PREFIX"
    UPPERCASE = "UPPERCASE"
    FILE_URI_WIN = "FILE_URI_WIN"
    DRIVE_ABS_WIN = "DRIVE_ABS_WIN"
    REPO_FILE_URI = "REPO_FILE_URI"
    EXPLICIT_DOUBLE_TRUE = "EXPLICIT_DOUBLE_TRUE"
    EXPLICIT_TRIPLE_TRUE = "EXPLICIT_TRIPLE_TRUE"
    EXPLICIT_PYTHON_TRUE = "EXPLICIT_PYTHON_TRUE"
    COMMENT_PREFIX = "COMMENT_PREFIX"
    COBRA_COMMAND = "COBRA_COMMAND"
    SHORT_DESC = "SHORT_DESC"
    EXAMPLE_USAGE = "EXAMPLE_USAGE"
    CHANGELOG_HEADER = "CHANGELOG_HEADER"
    FILE_NUM_PREFIX = "FILE_NUM_PREFIX"
    H1_HEADER = "H1_HEADER"
    PLACEHOLDER_TOKEN = "PLACEHOLDER_TOKEN"
    NON_ALPHANUMERIC = "NON_ALPHANUMERIC"

# Centralized Raw Regex Definitions: Enum -> (Pattern String, Flags)
REGEX_DEFINITIONS: dict[RegexPatternType, tuple[str, int]] = {
    RegexPatternType.WINDOWS_BACKSLASH: (r"\\", 0),
    RegexPatternType.LEADING_DOT_SLASH: (r"^\./", 0),
    RegexPatternType.CRLF: (r"\r\n", 0),
    RegexPatternType.UNIVERSAL_LINE_ENDING: (r"\r\n|\r", 0),
    RegexPatternType.TRAILING_WHITESPACE: (r"[ \t]+$", re.MULTILINE),
    RegexPatternType.SEQ_PREFIX: (r"^([0-9]+)-(.*)$", 0),
    RegexPatternType.UPPERCASE: (r"[A-Z]", 0),
    RegexPatternType.FILE_URI_WIN: (r"file:///[A-Za-z]:/[^\s\)\]\"'>]+", 0),
    RegexPatternType.DRIVE_ABS_WIN: (r"(?<![A-Za-z0-9_])[A-Za-z]:\\[A-Za-z0-9_\\.-]+", 0),
    RegexPatternType.REPO_FILE_URI: (r"file:///[A-Za-z]:/[^/]+/coding-guidelines/([^\s\)\]\"'>]+)", 0),
    RegexPatternType.EXPLICIT_DOUBLE_TRUE: (r"==\s*true\b", re.IGNORECASE),
    RegexPatternType.EXPLICIT_TRIPLE_TRUE: (r"===\s*true\b", re.IGNORECASE),
    RegexPatternType.EXPLICIT_PYTHON_TRUE: (r"==\s*True\b", 0),
    RegexPatternType.COMMENT_PREFIX: (r"^\s*(//|#|\*|/\*)", 0),
    RegexPatternType.COBRA_COMMAND: (r"var\s+(\w+Cmd)\s*=\s*&cobra\.Command\s*\{([^}]+)\}", re.DOTALL),
    RegexPatternType.SHORT_DESC: (r"Short:\s*\"[^\"]+\"", 0),
    RegexPatternType.EXAMPLE_USAGE: (r"Example:\s*\"[^\"]+\"", 0),
    RegexPatternType.CHANGELOG_HEADER: (r"##\s+\[v?([0-9]+\.[0-9]+\.[0-9]+[^\]]*)\]", 0),
    RegexPatternType.FILE_NUM_PREFIX: (r"^([0-9]+)-(.*)\.md$", 0),
    RegexPatternType.H1_HEADER: (r"^(#\s+)([0-9]+)(\s*[-—:]\s*)(.*)$", re.MULTILINE),
    RegexPatternType.PLACEHOLDER_TOKEN: (r"[A-Z0-9_]*PLACEHOLDER[A-Z0-9_]*", 0),
    RegexPatternType.NON_ALPHANUMERIC: (r"[^a-zA-Z0-9_-]+", 0),
}

# --- Thread-Safe Lazy Regex Registry ---
class RegexRegistry:
    """Thread-safe lazy-compiling regex registry with double-checked locking."""
    _cache: dict[RegexPatternType, re.Pattern] = {}
    _lock = threading.Lock()

    @classmethod
    def get(cls, pattern_type: RegexPatternType) -> re.Pattern:
        """Lazily compiles and returns the cached re.Pattern object."""
        if pattern_type in cls._cache:
            return cls._cache[pattern_type]

        with cls._lock:
            if pattern_type not in cls._cache:
                if pattern_type not in REGEX_DEFINITIONS:
                    raise KeyError(f"Pattern type '{pattern_type}' is not registered in REGEX_DEFINITIONS")
                raw_pattern, flags = REGEX_DEFINITIONS[pattern_type]
                cls._cache[pattern_type] = re.compile(raw_pattern, flags)
            return cls._cache[pattern_type]

    @classmethod
    def get_group(cls, *pattern_types: RegexPatternType) -> tuple[re.Pattern, ...]:
        """Lazily retrieves a tuple of compiled re.Pattern objects."""
        return tuple(cls.get(pt) for pt in pattern_types)

def get_compiled_regex(pattern_type: RegexPatternType) -> re.Pattern:
    """Convenience functional accessor for RegexRegistry.get."""
    return RegexRegistry.get(pattern_type)

def get_compiled_regex_group(*pattern_types: RegexPatternType) -> tuple[re.Pattern, ...]:
    """Convenience functional accessor for RegexRegistry.get_group."""
    return RegexRegistry.get_group(*pattern_types)

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
    """
    Checks if file is binary by extension or by memory-safe 8KB chunk probing for null bytes.
    Avoids loading full large files into RAM.
    """
    if file_path.suffix.lower() in BINARY_EXTENSIONS:
        return True
    try:
        if file_path.is_file():
            with open(file_path, "rb") as f:
                chunk = f.read(8192)
                if b"\x00" in chunk:
                    return True
    except Exception:
        pass
    return False

def is_allowed_large_file(file_path: str | Path) -> bool:
    """Checks if file is on the explicit waiver list for large generated assets."""
    norm = normalize_rel_path(file_path).lstrip("./")
    return norm in {normalize_rel_path(f).lstrip("./") for f in ALLOWED_LARGE_FILES}

def normalize_rel_path(path: str | Path) -> str:
    """Converts a path into a canonical relative POSIX path using PATH_SEPARATOR."""
    re_slash = get_compiled_regex(RegexPatternType.WINDOWS_BACKSLASH)
    re_lead = get_compiled_regex(RegexPatternType.LEADING_DOT_SLASH)
    p_str = re_slash.sub(PATH_SEPARATOR, str(path))
    return re_lead.sub("", p_str)

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

def read_file_safe(
    path: str | Path,
    max_bytes: int = MAX_READ_SIZE_BYTES,
    encoding: str = DEFAULT_ENCODING
) -> str | None:
    """
    Memory-safe and fault-tolerant file reader.
    Handles missing or deleted files gracefully with zero crashes.
    Normalizes both CRLF (\\r\\n) and legacy Mac CR (\\r) to strict UNIX LF (\\n).
    """
    try:
        p = Path(path)
        if not p.exists():
            return None
        if not p.is_file():
            return None
        re_univ_nl = get_compiled_regex(RegexPatternType.UNIVERSAL_LINE_ENDING)
        with open(p, "r", encoding=encoding, errors="replace") as f:
            raw_text = f.read(max_bytes)
            return re_univ_nl.sub(LINE_SEPARATOR, raw_text)
    except (FileNotFoundError, PermissionError, OSError):
        return None

def read_file_lf(path: str | Path, encoding: str = DEFAULT_ENCODING) -> str:
    """Reads a text file ensuring strict UNIX LF. Returns empty string if file does not exist."""
    content = read_file_safe(path, encoding=encoding)
    return content if content is not None else ""

def write_file_lf(
    path: str | Path,
    content: str,
    encoding: str = DEFAULT_ENCODING
) -> bool:
    """
    Atomic write ensuring strict UNIX LF line endings and Unix permission preservation.
    Encodes file according to the specified encoding constant (default: utf-8 without BOM).
    Preserves original executable bits (chmod +x / st_mode) on Linux/macOS.
    """
    p = Path(path)
    p.parent.mkdir(parents=True, exist_ok=True)
    temp_path = p.with_name(f"{p.name}.tmp_{os.getpid()}_{int(time.time()*1000)}")

    # Preserve executable permissions if original file exists on Unix
    original_mode = None
    if p.exists():
        try:
            original_mode = p.stat().st_mode
        except Exception:
            pass

    try:
        re_univ_nl = get_compiled_regex(RegexPatternType.UNIVERSAL_LINE_ENDING)
        lf_content = re_univ_nl.sub(LINE_SEPARATOR, content)
        with open(temp_path, "wb") as f:
            f.write(lf_content.encode(encoding))

        if original_mode is not None:
            try:
                os.chmod(temp_path, original_mode)
            except Exception:
                pass

        temp_path.replace(p)
        return True
    except Exception:
        if temp_path.exists():
            try:
                temp_path.unlink()
            except Exception:
                pass
        return False

# --- Dual-Platform Cross-Process Locking Mechanism ---

@contextmanager
def atomic_cache_lock(lock_name: str = "repo-cache.lock", timeout: float = LOCK_TIMEOUT_SECONDS):
    """
    Dual-platform cross-process lock:
    - On Unix: Uses kernel-level `fcntl.flock` (automatic cleanup on crash/SIGKILL).
    - On Windows: Uses `os.O_CREAT | os.O_EXCL` with PID timestamp & stale lock eviction (>15s).
    """
    CACHE_LOCKS_DIR.mkdir(parents=True, exist_ok=True)
    lock_file = CACHE_LOCKS_DIR / lock_name
    start_time = time.time()
    is_acquired = False
    lock_fd = None

    if IS_FCNTL_AVAILABLE:
        # Native Unix POSIX flock
        try:
            lock_fd = open(lock_file, "w")
            while time.time() - start_time < timeout:
                try:
                    fcntl.flock(lock_fd, fcntl.LOCK_EX | fcntl.LOCK_NB)
                    is_acquired = True
                    break
                except (BlockingIOError, OSError):
                    time.sleep(0.02)
        except Exception:
            is_acquired = False

        try:
            yield is_acquired
        finally:
            if lock_fd is not None:
                try:
                    fcntl.flock(lock_fd, fcntl.LOCK_UN)
                    lock_fd.close()
                except Exception:
                    pass
    else:
        # Windows native O_EXCL with stale lock eviction
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
                os.write(fd, f"pid={os.getpid()}{LINE_SEPARATOR}time={time.time()}".encode(DEFAULT_ENCODING))
                os.close(fd)
                is_acquired = True
                break
            except FileExistsError:
                time.sleep(0.02)
            except Exception:
                break

        try:
            yield is_acquired
        finally:
            if is_acquired:
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
                with open(target, "r", encoding=DEFAULT_ENCODING) as f:
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
            with open(temp_primary, "w", encoding=DEFAULT_ENCODING) as f:
                json.dump(cache_data, f, indent=2)
            temp_primary.replace(PRIMARY_CACHE_FILE)
        except Exception:
            pass

        try:
            temp_legacy = LEGACY_CACHE_FILE.with_suffix(".json.tmp")
            with open(temp_legacy, "w", encoding=DEFAULT_ENCODING) as f:
                json.dump(cache_data, f, indent=2)
            temp_legacy.replace(LEGACY_CACHE_FILE)
        except Exception:
            pass

# --- Two-Phase Streaming Engine with Inode Cycle Protection ---

def stream_cached_files(
    cache_data: dict[str, Any],
    root_dir: str = ".",
    extensions: set[str] | tuple | None = None,
    custom_excludes: set[str] | None = None
) -> Generator[Path, None, None]:
    """Phase 1: Streams valid files from cache first. Automatically skips missing/deleted/excluded files."""
    file_list = cache_data.get("files", [])
    norm_root = normalize_rel_path(root_dir).rstrip(PATH_SEPARATOR)
    ext_set = normalize_extensions(extensions)

    for rel_path in file_list:
        norm_p = normalize_rel_path(rel_path)
        if norm_root:
            if norm_root != ".":
                if not norm_p.startswith(norm_root + PATH_SEPARATOR):
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
    """
    Phase 2: Walks filesystem pruning ignored folders (including nested .git, .gitmap, node_modules).
    Guarded with visited inode tracking (st_dev, st_ino) on Unix to prevent symlink recursion cycles.
    """
    ext_set = normalize_extensions(extensions)
    visited_inodes: set[tuple[int, int]] = set()

    for root, dirs, files in os.walk(root_dir, followlinks=False):
        # Prevent Unix symlink recursion loops
        try:
            st = os.stat(root)
            inode_key = (st.st_dev, st.st_ino)
            if inode_key in visited_inodes:
                dirs[:] = []
                continue
            visited_inodes.add(inode_key)
        except Exception:
            pass

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
    is_use_cache: bool = True,
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
    cache_data = load_repo_cache() if is_use_cache else {}
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
