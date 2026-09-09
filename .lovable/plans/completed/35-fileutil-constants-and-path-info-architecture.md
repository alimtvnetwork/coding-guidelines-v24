# Plan 35: Fileutil Constants Centralization & .NET-Style PathInfo / FolderInfo Architecture

## Overview
This architectural plan addresses the user's three key observations:
1. **Constant Everywhere:** Eliminate hardcoded string and byte literals across `04-code/golang/pkg/fileutil/` (starting with Windows long path prefix `\\?\` in `cleanLongPath` and expanding to all path prefixes, separators, environment variables, and default permissions).
2. **Path Namespace Clarification:** Provide complete documentation and architectural clarity on how `Path` and `PathWrapper` operate.
3. **.NET-Style PathInfo & FolderInfo:** Introduce rich, stateful navigation models (`FolderInfo`, `FileInfo`, and `PathInfo`) inspired by .NET's `System.IO.DirectoryInfo` and `System.IO.FileInfo`, providing parent navigation, child enumeration, sibling/parent folder inspection, recursive directory walking, and monadic result containers.

## Task-Specific Rule Set (Guarantees & Constraints)
1. **Rule 1 (Zero Hardcoded Magic Strings):** All filesystem tokens, path separators (`/`, `\`), Windows prefixes (`\\?\`, `\\?\UNC\`), home delimiters (`~`), env markers (`$`, `%`), and octal directory modes (`0755`) MUST reference centralized constants in `pkg/fileutil/consts.go`.
2. **Rule 2 (.NET Hierarchy Parity):** `FolderInfo` MUST support `.Parent()`, `.ParentFolderName()`, `.Subfolders()` / `.Directories()`, `.Files()`, `.ParentFolderFiles()`, `.ParentFolderDirectories()`, and recursive directory walking (`.Walk()`, `.WalkFiles()`, `.WalkDirectories()`, `.AllFiles()`, `.AllDirectories()`).
3. **Rule 3 (Bounded Function Size $\le 15$ Lines):** Every function and method in newly created or refactored Go files MUST NOT exceed 15 lines. Helper functions must be extracted for complex logic.
4. **Rule 4 (Implicit Boolean Evaluation):** All booleans MUST use `is` or `has` prefixes (`IsExists()`, `IsDir()`, `IsFile()`) and MUST be evaluated implicitly (`if info.IsExists()`). Explicit comparisons against `true` or mixed-polarity conditions are strictly banned.
5. **Rule 5 (Universal AppError & Monadic Result):** All error states from disk inspection MUST return `*appfault.AppError` wrapped in monadic envelopes (`ResultSlice[T]`, `BoolResult`, `FileInfoResult`).

## Subtask Decomposition

- [01-task-constants-centralization.md](../subtasks/35-fileutil-constants-and-path-info-architecture/01-task-constants-centralization.md): Centralize all path prefixes, separators, environment variable keys, and permissions into `consts.go` and refactor `path_normalize.go`, `path_env.go`, `path_info.go`, `path_temp.go`, `file_path_ops.go`, and `fileutil.go`.
- [02-task-type-aliases-and-results.md](../subtasks/35-fileutil-constants-and-path-info-architecture/02-task-type-aliases-and-results.md): Add `ResultSlice[T]` forwarders and typed `FolderInfoResult` / `FileInfoObjResult` constructors to `types.go` and `results.go`.
- [03-task-folder-info-and-file-info-implementation.md](../subtasks/35-fileutil-constants-and-path-info-architecture/03-task-folder-info-and-file-info-implementation.md): Implement `FolderInfo` (`folder_info.go`), `FileInfo` (`file_info.go`), and universal `PathInfo` (`path_info_obj.go`).
- [04-task-namespace-integration-and-bridges.md](../subtasks/35-fileutil-constants-and-path-info-architecture/04-task-namespace-integration-and-bridges.md): Integrate `FolderInfo` and `FileInfo` into `Path`, `File`, `New`, and `PathWrapper`.
- [05-task-testing-and-verification.md](../subtasks/35-fileutil-constants-and-path-info-architecture/05-task-testing-and-verification.md): Implement comprehensive unit tests for `FolderInfo`, `FileInfo`, constants, and run full CI/CD quality gates.
