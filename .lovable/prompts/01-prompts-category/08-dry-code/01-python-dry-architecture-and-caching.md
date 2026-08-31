# Python Script DRY Architecture, Enums, Pluggable Caching & Fast AI Reading Specification

> **Prompt Version:** 2.3.0  
> **Target:** `.lovable/prompts/01-prompts-category/08-dry-code/01-python-dry-architecture-and-caching.md`  
> **Synchronization:** Meta-Repo & AI Scripting Ecosystem

/goal Standardize the architectural design of all Python CI/CD, linting, and fix scripts using top-level Enums with PascalCase members, centralized configuration maps, thread-safe lazy regex compilation with double-checked locking, small decomposed functions, DRY shared engines, multi-folder scoping, customizable extensions, pluggable `tmp/cache/` storage, cross-process atomic file locking, and two-phase incremental `mtime` caching.

## 🎯 Architectural Philosophy

All AI-authored Python scripts in this repository (`.lovable/ai-fix-scripts/` and `.agent/scripts/`) MUST strictly adhere to the following principles:

---

## 1. Top-Level Constants, Enums (PascalCase Members) & Thread-Safe Lazy Regex Registry

- **Root Definition:** All configuration parameters (`EXCLUDE_DIRS`, `BINARY_EXTENSIONS`, `DEFAULT_TEXT_EXTENSIONS`, `DEFAULT_CODE_EXTENSIONS`, `DEFAULT_CLI_EXTENSIONS`, `CACHE_BASE_DIR`, `DEFAULT_MAX_FILE_KB`, `ALLOWED_LARGE_FILES`) must be centralized in `00-shared-engine.py`.
- **Enum `Type` Suffix & PascalCase Members:** All Enums MUST end with the `Type` suffix (e.g., `ScanModeType`, `ExitCodeType`, `SeverityType`, `RegexPatternType`), and all Enum variable members MUST strictly use **PascalCase** (e.g., `WindowsBackslash`, `LeadingDotSlash`, `Success`, `ViolationsFound`, `Blocker`).
- **Implicit Boolean Checks:** Never compare booleans against explicit `True` (BAN: `if is_valid == True:` -> MANDATORY: `if is_valid:`).
- **Thread-Safe Lazy Regex Compilation (Zero Startup Overhead):**
  - Store raw regex pattern string definitions + compilation flags in a central `REGEX_DEFINITIONS: dict[RegexPatternType, tuple[str, int]]` map in `00-shared-engine.py`.
  - Use `RegexRegistry.get(pattern_type)` or `get_compiled_regex(pattern_type)` with double-checked `threading.Lock()` memoization.
  - Regex patterns are compiled on first demand, then served in `O(1)` time for all subsequent lookups across threads.
- **Ignore Pruning:** Ensure `EXCLUDE_DIRS` covers `.gitmap`, `.git`, `node_modules`, `dist`, `build`, `.venv`, `.gemini`, `tmp`, `.system_generated`, and `release-artifacts` at all subtree depths.

```python
import re
import threading
from enum import Enum
from pathlib import Path

# --- Top-Level Enums with PascalCase Members ---
class ScanModeType(str, Enum):
    Check = "check"
    Fix = "fix"
    Stream = "stream"

class SeverityType(str, Enum):
    Blocker = "blocker"
    High = "high"
    Warn = "warn"
    Info = "info"

class ExitCodeType(int, Enum):
    Success = 0
    ViolationsFound = 1
    ToolError = 2

class RegexPatternType(str, Enum):
    WindowsBackslash = "windows_backslash"
    LeadingDotSlash = "leading_dot_slash"
    Crlf = "crlf"
    TrailingWhitespace = "trailing_whitespace"
    SeqPrefix = "seq_prefix"
    Uppercase = "uppercase"
    FileUriWin = "file_uri_win"
    DriveAbsWin = "drive_abs_win"
    RepoFileUri = "repo_file_uri"
    ExplicitDoubleTrue = "explicit_double_true"
    ExplicitTripleTrue = "explicit_triple_true"
    ExplicitPythonTrue = "explicit_python_true"
    CommentPrefix = "comment_prefix"
    CobraCommand = "cobra_command"
    ShortDesc = "short_desc"
    ExampleUsage = "example_usage"
    ChangelogHeader = "changelog_header"
    FileNumPrefix = "file_num_prefix"
    H1Header = "h1_header"
    PlaceholderToken = "placeholder_token"
    NonAlphanumeric = "non_alphanumeric"

# Centralized Raw Regex Definitions: Enum -> (Pattern String, Flags)
REGEX_DEFINITIONS: dict[RegexPatternType, tuple[str, int]] = {
    RegexPatternType.WindowsBackslash: (r"\\", 0),
    RegexPatternType.LeadingDotSlash: (r"^\./", 0),
    RegexPatternType.Crlf: (r"\r\n", 0),
    RegexPatternType.TrailingWhitespace: (r"[ \t]+$", re.MULTILINE),
    RegexPatternType.SeqPrefix: (r"^([0-9]+)-(.*)$", 0),
    RegexPatternType.Uppercase: (r"[A-Z]", 0),
    RegexPatternType.FileUriWin: (r"file:///[A-Za-z]:/[^\s\)\]\"'>]+", 0),
    RegexPatternType.DriveAbsWin: (r"(?<![A-Za-z0-9_])[A-Za-z]:\\[A-Za-z0-9_\\.-]+", 0),
    RegexPatternType.RepoFileUri: (r"file:///[A-Za-z]:/[^/]+/coding-guidelines/([^\s\)\]\"'>]+)", 0),
    RegexPatternType.ExplicitDoubleTrue: (r"==\s*true\b", re.IGNORECASE),
    RegexPatternType.ExplicitTripleTrue: (r"===\s*true\b", re.IGNORECASE),
    RegexPatternType.ExplicitPythonTrue: (r"==\s*True\b", 0),
    RegexPatternType.CommentPrefix: (r"^\s*(//|#|\*|/\*)", 0),
    RegexPatternType.CobraCommand: (r"var\s+(\w+Cmd)\s*=\s*&cobra\.Command\s*\{([^}]+)\}", re.DOTALL),
    RegexPatternType.ShortDesc: (r"Short:\s*\"[^\"]+\"", 0),
    RegexPatternType.ExampleUsage: (r"Example:\s*\"[^\"]+\"", 0),
    RegexPatternType.ChangelogHeader: (r"##\s+\[v?([0-9]+\.[0-9]+\.[0-9]+[^\]]*)\]", 0),
    RegexPatternType.FileNumPrefix: (r"^([0-9]+)-(.*)\.md$", 0),
    RegexPatternType.H1Header: (r"^(#\s+)([0-9]+)(\s*[-—:]\s*)(.*)$", re.MULTILINE),
    RegexPatternType.PlaceholderToken: (r"[A-Z0-9_]*PLACEHOLDER[A-Z0-9_]*", 0),
    RegexPatternType.NonAlphanumeric: (r"[^a-zA-Z0-9_-]+", 0),
}

# --- Thread-Safe Lazy Regex Registry ---
class RegexRegistry:
    _cache: dict[RegexPatternType, re.Pattern] = {}
    _lock = threading.Lock()

    @classmethod
    def get(cls, pattern_type: RegexPatternType) -> re.Pattern:
        if pattern_type in cls._cache:
            return cls._cache[pattern_type]
        with cls._lock:
            if pattern_type not in cls._cache:
                raw_pattern, flags = REGEX_DEFINITIONS[pattern_type]
                cls._cache[pattern_type] = re.compile(raw_pattern, flags)
            return cls._cache[pattern_type]
```

---

## 2. Multi-Folder Scoping & Customizable Extensions

- **Multi-Folder Capability:** All scripts MUST accept a target directory (`<path>` or `--path` / `--dir`) allowing them to run on the full repository, specific submodules, or individual feature folders.
- **Customizable File Extensions:** Supported file extensions must be customizable via `--ext` or function parameters (`extensions=...`), with robust lowercasing and leading dot normalization.
- **Nested Directory Pruning:** When using `os.walk`, prune `dirs[:]` dynamically so nested `.git`, `.gitmap`, and `node_modules` folders inside subprojects are skipped instantly.

---

## 3. Decomposed Pure Functions (< 25 Lines Each)

- Monolithic scripts are strictly forbidden.
- Scripts MUST be broken down into small, composable, single-responsibility functions.
- Separate file I/O, cache state verification, violation analysis, and CLI output formatting into distinct functions.

---

## 4. Pluggable `tmp/cache/` Structure & Cross-Process Locking

All cache data is organized in structured, pluggable subdirectories under `tmp/cache/`:

1. **`tmp/cache/paths/`**: Stores repository file path listings and metadata indexes.
2. **`tmp/cache/locks/`**: Cross-process file locks (`repo-cache.lock`) preventing corruption when multiple agents or subagents operate simultaneously.
3. **`tmp/cache/files/`**: Cached tokenized contents or AST data.

### Safe Atomic Locking & Stale Lock Recovery
- Lock files MUST record PID and timestamp.
- If a lock file is older than 15 seconds (stale lock from an interrupted process), it is automatically evicted to avoid deadlocks.
- File mutations use atomic temp file replacement (`.tmp` -> final path).

---

## 5. Two-Phase Incremental Caching & Missing File Tolerance

To achieve sub-15ms repository-wide execution without redundant disk I/O:

1. **Phase 1 (Cache-First Processing):**
   - Read pre-computed file metadata (`mtime`, `size`) from `tmp/cache/repo-file-cache.json`.
   - Deliver cached files immediately (<0.1ms) so the consumer starts processing file #1 without waiting.
2. **Phase 2 (Streaming Discovery for New / Modified Files):**
   - Concurrently stream directory entries via `os.walk` / `os.scandir`.
   - For any newly created or modified file, stream it immediately through the pipeline.
3. **Fault-Tolerant File Reading (Zero Crash on Deletions):**
   - If a file is deleted or missing during traversal, `read_file_safe()` returns `None` instead of raising an unhandled exception.
   - Deleted files are automatically evicted from the cache without treating the deletion as a failure.

---

## 6. High-Speed File Reading & Exploration for AI Agents

AI agents and subagents should avoid slow, recursive shell commands (`Get-ChildItem -Recurse`, `dir /s`, or brute-force glob searches). Instead, use the optimized Python toolchain:

| Task | Recommended AI Command | Speed |
|---|---|:---:|
| **List Folder Files** | `python .lovable/ai-fix-scripts/14-fast-file-reader.py --list-folder <dir>` | **<1ms** |
| **Fast Safe File Read** | `python .lovable/ai-fix-scripts/14-fast-file-reader.py --read-file <path>` | **<1ms** |
| **Search File Paths** | `python .lovable/ai-fix-scripts/14-fast-file-reader.py --search-pattern "<term>"` | **<2ms** |
| **Full Repo File Index** | `python .lovable/ai-fix-scripts/08-fast-file-scanner.py --lang ts,go --path spec/` | **~14ms** |
| **Parallel Content Grep** | `python .lovable/ai-fix-scripts/09-fast-cached-grep.py --pattern "<text>"` | **~12ms** |
| **File Manipulation CLI** | `python .lovable/ai-fix-scripts/01-file-manipulator.py <cmd> <dir>` | **~15ms** |

---

## 7. Cross-Platform Python Mandate for CI/CD & Codegen Checks (Ban on `.sh` in CI)

- **Portability Contract:** All CI/CD checks, codegen determinism verifiers, fixture regenerators, and linter runners MUST be implemented as pure, cross-platform Python (`.py`) scripts using standard library modules (`pathlib`, `subprocess`, `difflib`, `tempfile`).
- **TOTAL BAN on `.sh` in CI Workflows:** Shell scripts (`.sh`) fail or require bash emulation on Windows/PowerShell runner environments. CI workflow steps and `package.json` lifecycle scripts MUST invoke Python scripts (`python3 ... .py`) directly.
- **Legacy Migration:** Whenever a legacy `.sh` check is encountered in CI or tests, migrate its core logic to a standalone Python script adhering to these standards, leaving only a lightweight shell forwarder if backward compatibility is required.
