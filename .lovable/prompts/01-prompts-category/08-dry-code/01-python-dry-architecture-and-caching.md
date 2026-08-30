# Python Script DRY Architecture, Enums & Incremental Caching Specification

> **Prompt Version:** 1.0.0  
> **Target:** `.lovable/prompts/01-prompts-category/08-dry-code/01-python-dry-architecture-and-caching.md`  
> **Synchronization:** Meta-Repo & AI Scripting Ecosystem

/goal Standardize the architectural design of all Python CI/CD, linting, and fix scripts using top-level Enums, small decomposed functions, DRY shared engines, and two-phase incremental `mtime` caching.

## 🎯 Architectural Philosophy

All AI-authored Python scripts in this repository (`.lovable/ai-fix-scripts/` and `.agent/scripts/`) MUST strictly adhere to the following principles:

---

## 1. Top-Level Constants & Enums

- **Root Definition:** All configuration parameters (`EXCLUDE_DIRS`, `CACHE_FILE_PATH`, `DEFAULT_BUFFER_SIZE`) must be declared as module-level constants at the top of the file.
- **Enum `Type` Suffix:** All Enums MUST end with the `Type` suffix (e.g., `ScanModeType`, `ExitCodeType`, `SeverityType`).
- **Implicit Boolean Checks:** Never compare booleans against explicit `True` (BAN: `if is_valid == True:` -> MANDATORY: `if is_valid:`).

```python
from enum import Enum
from pathlib import Path

# --- Module-Level Constants ---
CACHE_FILE_PATH = Path("tmp/repo-file-cache.json")
DEFAULT_MAX_FILE_KB = 1024
EXCLUDE_DIRS = {".git", "node_modules", "dist", "build", ".venv", ".gemini", "tmp"}

# --- Top-Level Enums ---
class ScanModeType(str, Enum):
    CHECK = "check"
    FIX = "fix"
    STREAM = "stream"

class SeverityType(str, Enum):
    BLOCKER = "blocker"
    HIGH = "high"
    WARN = "warn"

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

## 3. Two-Phase Incremental `mtime` Caching & Streaming

To achieve sub-15ms repository-wide execution without redundant disk I/O:

1. **Phase 1 (Cache-First Processing):**
   - Read pre-computed file metadata (`mtime`, `size`, `hash`) from `tmp/repo-file-cache.json`.
   - If cache is warm and file `mtime`/`size` has not changed, skip re-parsing unchanged content or immediately process cached entries.
2. **Phase 2 (Streaming Discovery for New / Modified Files):**
   - Use `os.scandir` to stream directory entries.
   - For any newly created or modified file, yield it immediately through the streaming pipeline so processing begins without waiting for the full walk to finish.
3. **Universal Shared Engine (`00-shared-engine.py`):**
   - All scripts MUST reuse `00-shared-engine.py` for repository traversal, caching, path normalization, and safe UNIX LF writing.

---

## 4. Execution Checklist for the AI

- [ ] `/goal` Ensure all scripts import traversal and caching logic from `00-shared-engine.py`.
- [ ] `/learn` Declare module constants and `*Type` Enums at the root of the file.
- [ ] `/goal` Keep individual functions under 25 lines following Single Responsibility Principle.
- [ ] `/learn` Verify zero external dependencies (Python standard library only).
- [ ] `/learn` Ensure all generated files use strict UTF-8 with UNIX LF line endings.
