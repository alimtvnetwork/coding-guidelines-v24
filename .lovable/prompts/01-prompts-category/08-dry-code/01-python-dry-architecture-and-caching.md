# Python Script DRY Architecture, Enums, Pluggable Caching & Fast AI Reading Specification

> **Prompt Version:** 2.0.0  
> **Target:** `.lovable/prompts/01-prompts-category/08-dry-code/01-python-dry-architecture-and-caching.md`  
> **Synchronization:** Meta-Repo & AI Scripting Ecosystem

/goal Standardize the architectural design of all Python CI/CD, linting, and fix scripts using top-level Enums, small decomposed functions, DRY shared engines, pluggable `tmp/cache/` storage, cross-process atomic file locking, and two-phase incremental `mtime` caching.

## 🎯 Architectural Philosophy

All AI-authored Python scripts in this repository (`.lovable/ai-fix-scripts/` and `.agent/scripts/`) MUST strictly adhere to the following principles:

---

## 1. Top-Level Constants & Enums

- **Root Definition:** All configuration parameters (`EXCLUDE_DIRS`, `CACHE_BASE_DIR`, `DEFAULT_MAX_FILE_KB`) must be declared as module-level constants at the top of the file.
- **Enum `Type` Suffix:** All Enums MUST end with the `Type` suffix (e.g., `ScanModeType`, `ExitCodeType`, `SeverityType`).
- **Implicit Boolean Checks:** Never compare booleans against explicit `True` (BAN: `if is_valid == True:` -> MANDATORY: `if is_valid:`).

```python
from enum import Enum
from pathlib import Path

# --- Module-Level Constants ---
CACHE_BASE_DIR = Path("tmp/cache")
CACHE_PATHS_DIR = CACHE_BASE_DIR / "paths"
CACHE_LOCKS_DIR = CACHE_BASE_DIR / "locks"
CACHE_FILES_DIR = CACHE_BASE_DIR / "files"
DEFAULT_MAX_FILE_KB = 2048
EXCLUDE_DIRS = {".git", "node_modules", "dist", "build", ".venv", ".gemini", "tmp", ".system_generated"}

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
```

---

## 2. Decomposed Pure Functions (< 25 Lines Each)

- Monolithic scripts are strictly forbidden.
- Scripts MUST be broken down into small, composable, single-responsibility functions.
- Separate file I/O, cache state verification, violation analysis, and CLI output formatting into distinct functions.

---

## 3. Pluggable `tmp/cache/` Structure & Cross-Process Locking

All cache data is organized in structured, pluggable subdirectories under `tmp/cache/`:

1. **`tmp/cache/paths/`**: Stores repository file path listings and metadata indexes.
2. **`tmp/cache/locks/`**: Cross-process file locks (`repo-cache.lock`) preventing corruption when multiple agents or subagents operate simultaneously.
3. **`tmp/cache/files/`**: Cached tokenized contents or AST data.

### Safe Atomic Locking & Stale Lock Recovery
- Lock files MUST record PID and timestamp.
- If a lock file is older than 15 seconds (stale lock from an interrupted process), it is automatically evicted to avoid deadlocks.
- File mutations use atomic temp file replacement (`.tmp` -> final path).

---

## 4. Two-Phase Incremental Caching & Missing File Tolerance

To achieve sub-15ms repository-wide execution without redundant disk I/O:

1. **Phase 1 (Cache-First Processing):**
   - Read pre-computed file metadata (`mtime`, `size`) from `tmp/cache/repo-file-cache.json`.
   - Deliver cached files immediately (<0.1ms) so the consumer starts processing file #1 without waiting.
2. **Phase 2 (Streaming Discovery for New / Modified Files):**
   - Concurrently stream directory entries via `os.scandir`.
   - For any newly created or modified file, stream it immediately through the pipeline.
3. **Fault-Tolerant File Reading (Zero Crash on Deletions):**
   - If a file is deleted or missing during traversal, `read_file_safe()` returns `None` instead of raising an unhandled exception.
   - Deleted files are automatically evicted from the cache without treating the deletion as a failure.

---

## 5. High-Speed File Reading & Exploration for AI Agents

AI agents and subagents should avoid slow, recursive shell commands (`Get-ChildItem -Recurse`, `dir /s`, or brute-force glob searches). Instead, use the optimized Python toolchain:

| Task | Recommended AI Command | Speed |
|---|---|:---:|
| **List Folder Files** | `python .lovable/ai-fix-scripts/14-fast-file-reader.py --list-folder <dir>` | **<1ms** |
| **Fast Safe File Read** | `python .lovable/ai-fix-scripts/14-fast-file-reader.py --read-file <path>` | **<1ms** |
| **Full Repo File Index** | `python .lovable/ai-fix-scripts/08-fast-file-scanner.py --lang ts,go --path spec/` | **~14ms** |
| **Parallel Content Grep** | `python .lovable/ai-fix-scripts/09-fast-cached-grep.py --pattern "<text>"` | **~12ms** |

---

## 6. Execution Checklist for the AI

- [ ] `/goal` Ensure all scripts import traversal and caching logic from `00-shared-engine.py`.
- [ ] `/learn` Declare module constants and `*Type` Enums at the root of the file.
- [ ] `/goal` Keep individual functions under 25 lines following Single Responsibility Principle.
- [ ] `/learn` Store cache and lockfiles in pluggable `tmp/cache/` directories with atomic locks.
- [ ] `/learn` Handle missing/deleted files safely via `read_file_safe()` with zero crashes.
- [ ] `/learn` Verify zero external dependencies (Python standard library only).
- [ ] `/learn` Ensure all generated files use strict UTF-8 with UNIX LF line endings.
