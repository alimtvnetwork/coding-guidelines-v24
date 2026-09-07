# Subtask 01: Cross-Platform Temp Directory & Environment Expansion

> **Parent Plan:** `18-modular-filepath-util-and-cross-platform-temp`  
> **Status:** Completed  
> **Bounded Files:**  
> - `04-code/golang/pkg/fileutil/path_temp.go`  
> - `04-code/golang/pkg/fileutil/path_temp_test.go`  
> - `04-code/golang/pkg/fileutil/path_env.go`  
> - `04-code/golang/pkg/fileutil/path_env_test.go`

---

## Instructions

1. **In `04-code/golang/pkg/fileutil/path_temp.go`:**
   - Implement `UserTempDir() StringResult`:
     - Resolves user-scoped temporary directory with multi-tier fallback:
       - Windows: `%TEMP%` -> `%TMP%` -> `%LOCALAPPDATA%\Temp` -> `%USERPROFILE%\AppData\Local\Temp` -> `os.TempDir()`.
       - Linux/Ubuntu: `$TMPDIR` -> `$XDG_RUNTIME_DIR` -> `~/.cache/tmp` (if created) -> `/tmp`.
       - macOS: `$TMPDIR` -> `os.TempDir()` -> `/tmp`.
     - Ensures the directory exists or returns `StringFailure(errtype.IO, err, dir, "failed to resolve temp directory")`.
   - Implement `UserTempPath(subpath ...string) StringResult`:
     - Joins subpaths onto `UserTempDir()`.
   - Implement `CreateTempFile(dir string, pattern string, perm FilePermType) FileResult`:
     - If `dir` is empty, uses `UserTempDir().Data()`.
     - Creates temp file using `os.CreateTemp(dir, pattern)`.
     - Returns `FileSuccess(f)` or `FileFailure(errtype.IO, err, dir, "failed to create temp file")`.
   - Implement `CreateTempDir(dir string, pattern string, perm FilePermType) StringResult`:
     - Uses `os.MkdirTemp(dir, pattern)`.
     - Returns `StringSuccess(d)` or `StringFailure(errtype.IO, err, dir, "failed to create temp dir")`.
   - Implement `TempFile(pattern string) FileResult`:
     - Calls `CreateTempFile("", pattern, FilePermStandard)`.
   - Implement `TempDir(pattern string) StringResult`:
     - Calls `CreateTempDir("", pattern, FilePermStandard)`.
   - Keep all functions <= 15 lines. Blank line after closing brace `}` if followed by code.

2. **In `04-code/golang/pkg/fileutil/path_env.go`:**
   - Implement pure Go environment variable and tilde expansion:
     - `ExpandTilde(path string) StringResult`:
       - Expands `~`, `~/...`, `~\...` using `os.UserHomeDir()` without spawning any shell subprocesses.
     - `ExpandEnv(path string) StringResult`:
       - Expands POSIX `$VAR` and `${VAR}` syntax.
       - Expands Windows `%VAR%` syntax (case-insensitive on Windows, handles `%USERPROFILE%`, `%TEMP%`, `%APPDATA%`, etc.).
       - Safely handles undefined variables (preserves original token or leaves empty).
     - `Expand(path string) StringResult`:
       - Runs `ExpandTilde` followed by `ExpandEnv`.
   - Keep all functions <= 15 lines.

3. **In `04-code/golang/pkg/fileutil/path_temp_test.go` and `path_env_test.go`:**
   - Unit tests covering `UserTempDir()`, `UserTempPath()`, `CreateTempFile()`, `CreateTempDir()`.
   - Unit tests for `$VAR`, `${VAR}`, `%VAR%`, and tilde `~` expansions.
   - Run `go test -C 04-code/golang -v ./pkg/fileutil` to verify.
