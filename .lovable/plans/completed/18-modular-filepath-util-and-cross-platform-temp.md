# Master Architectural Specification: Modular Filepath Utilities & Cross-Platform Temp

> **Plan Identifier:** `18-modular-filepath-util-and-cross-platform-temp`  
> **Status:** Completed  
> **Created:** 2026-09-07  
> **Scope:** `04-code/golang/pkg/fileutil/`

---

## 1. Problem Statement & User Rationale

1. **Cross-Platform Temp Directory Resolution:**
   - In legacy `pathhelper`, temp resolution was broken and fragile:
     - On Windows, it checked only `os.Getenv("Temp")` with zero fallbacks, and `LocalTempPath()` pointed to `%APPDATA%` (Roaming) rather than `%LOCALAPPDATA%`, generating invalid paths (`Roaming\local\temp`).
     - On Unix/Linux, inverted logic generated `/home/<user>/TMPDIR`. Furthermore, it ignored Linux `$XDG_RUNTIME_DIR` (`/run/user/<uid>`) and Darwin's per-user temp directories.
     - Worst of all, `UserPath()` panicked if `os.UserHomeDir()` failed, violating our Zero-Panic Policy.
   - User Mandate:
     > *"Make this, uh, temp path to Windows that would expand or use the Windows, uh, temp, I believe that would be the user/temp somewhere, right? So that would be the directory. But also for Ubuntu and other OS also use the temp directory for that user. Um, so that needs to be coming from one path actually."*
2. **Environment Variable & Tilde Expansion:**
   - Legacy `pathhelper` failed to expand standard Windows `%VAR%` syntax (only checked `%VAR` or `%{VAR}`) and spawned slow shell subprocesses (`sh -c "cd && pwd"`, `dscl`, `getent`) for tilde expansion.
   - User Mandate:
     > *"fix the environment path. You can get some idea and put it into your file path util."*
3. **Modular File Breakdown & Namespace Constructors:**
   - Instead of huge monolithic files, break the filepath functionality into multiple small, focused files and group API methods using namespace structs and constructors:
     > *"Can you please divide the file path things into multiple smaller files and group the API methods so that it's eas- easier to navigate... multiple files or multiple, let's say, namespace type thing that we do with the constructor."*

---

## 2. Task-Specific Rules & Invariants (Auto-Reject on Violation)

1. **Rule 1 (Strict Zero-Panic & Clean Error Envelopes):** No function in `pkg/fileutil/` may call `panic()`, `log.Fatal()`, or unhandled error panics. All failures must return `*appfault.AppError` or concrete results (`StringResult`, `FileResult`, `BoolResult`).
2. **Rule 2 (Multi-Tier Cross-Platform Temp Hierarchy):**
   - Windows: `%TEMP%` -> `%TMP%` -> `%LOCALAPPDATA%\Temp` -> `%USERPROFILE%\AppData\Local\Temp` -> `os.TempDir()`.
   - Ubuntu/Linux: `$TMPDIR` -> `$XDG_RUNTIME_DIR` -> `~/.cache/tmp` -> `/tmp`.
   - macOS: `$TMPDIR` -> `os.TempDir()` -> `/tmp`.
3. **Rule 3 (Pure Go Tokenizer for Env & Tilde Expansion):** Support `$VAR`, `${VAR}`, Windows `%VAR%`, and tilde `~` expansion entirely in pure Go without spawning shell subprocesses (`sh -c`, `cmd.exe`).
4. **Rule 4 (Namespace Grouping & Fluent Constructor):**
   - Global singleton namespace: `var Path = pathNamespace{}` providing sub-namespaces `Path.Temp.*`, `Path.Env.*`, `Path.Norm.*`, `Path.Info.*`, and `Path.Join(...)`.
   - Fluent constructor: `NewPath(raw string) *PathWrapper` chaining normalization, expansion, joining, and file I/O operations.
5. **Rule 5 (Coding Guidelines Adherence):**
   - Canonical function size <= 15 lines.
   - Mandatory blank line after closing brace `}` if followed by code.
   - Mandatory blank line before `return` unless sole statement.
   - Implicit boolean evaluation only (`if ok`, never `== true` or mixed polarity).
   - Strict relative git paths in all specs and comments.

---

## 3. Subtask Decomposition

| Subtask | Title | Bounded Files |
|---|---|---|
| `01-task-cross-platform-temp-and-env-expansion.md` | Cross-Platform Temp & Env Expansion | `04-code/golang/pkg/fileutil/path_temp.go`, `path_temp_test.go`, `path_env.go`, `path_env_test.go` |
| `02-task-path-normalize-info-and-namespace-builder.md` | Normalization, Inspection & Namespace Fluent Builder | `04-code/golang/pkg/fileutil/path_normalize.go`, `path_normalize_test.go`, `path_info.go`, `path_info_test.go`, `path_namespace.go`, `path_namespace_test.go`, `readme.md` |

---

## 4. Verification Plan

1. `go test -C 04-code/golang -v ./pkg/fileutil`
2. `go test -C 04-code/golang -count=1 ./...` across all 17 Go packages
3. `node scripts/sync-check.mjs`
4. `python 03-ai-scripts/06-cicd-local-runner.py --all` (all 31 gates green)
