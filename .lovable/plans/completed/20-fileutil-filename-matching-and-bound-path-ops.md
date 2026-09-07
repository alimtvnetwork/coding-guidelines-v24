# Master Plan 20: Fileutil Filename Matching, Bound File Path Operations & AI Skill

## 1. Problem Statement & User Mandate
1. **Filename to Struct Name Alignment:**
   - As identified by the user in `04-code/golang/pkg/fileutil/file_append.go`, struct `appendOps` was in a file prefixed with `file_`.
   - User Mandate: Ensure file names match the struct names so developers and AI can immediately find them:
     - `file_append.go` -> `append_ops.go` (and `file_append_test.go` -> `append_ops_test.go`)
     - `file_open.go` -> `open_ops.go`
     - `file_create.go` -> `create_ops.go`
     - `file_read.go` -> `read_ops.go`
     - `file_write.go` -> `write_ops.go`
2. **Bound File Path Struct (`FilePathOps`):**
   - The current operation structs (`openOps`, `createOps`, etc.) are stateless and require the caller to pass `path` on every method call.
   - User Mandate: Provide a bound struct containing file path information so callers don't have to pass the file path or `os.File` handle on every call:
     - Encapsulates three parameters: `workDir string`, `relPath string`, `absPath string`.
     - Primarily operates on `absPath` while maintaining a clear breakdown of `workDir` and `relPath`.
     - **Strict Immutability:** Any changes (`WithWorkDir`, `WithRelPath`, `Join`) clone and return a new instance.
     - **Pre-Flight Directory Safety:** Built-in methods (`EnsureParentDir`, `EnsureFile`, `CreateIfNotExist`) so operations never fail due to missing parent directories.
     - **Zero-Path Bound Operations:** `ReadBytes()`, `ReadString()`, `ReadLines()`, `WriteBytes()`, `WriteString()`, `AppendBytes()`, `AppendString()`, `Open()`, `Create()`, `Delete()`, `Stat()`.
     - **Singleton Integration:** Accessible via `File.Target(path)`, `File.At(workDir, relPath)`, `New.Target(path)`, `New.At(workDir, relPath)`.
3. **AI Skill in README:**
   - Document the fileutil architecture and usage patterns inside `04-code/golang/pkg/fileutil/readme.md` formatted as an AI Skill playbook so any AI model reading the codebase can understand and execute it.

---

## 2. Task-Specific Rule Set
1. **Rule 1 (Matching Names):** Every file defining an operation struct must be named after the snake_case conversion of that struct (e.g., `FilePathOps` in `file_path_ops.go`, `appendOps` in `append_ops.go`).
2. **Rule 2 (Immutability & Safety):** `FilePathOps` methods modifying paths must return a new copy (clone). Pre-flight directory checks must create missing parent directories using `0755` permissions without panicking.
3. **Rule 3 (Canonical Function Size & Formatting):** Every function must strictly adhere to `<= 15 lines`. A blank line must follow every closing brace `}` if followed by code. Blank line before `return` unless sole statement.
4. **Rule 4 (Zero-Panic Policy):** Functions must return concrete result aliases (`BoolResult`, `FileResult`, `BytesResult`, `StringResult`, `LinesResult`, `FileInfoResult`) wrapping `*appfault.AppError`. Never panic or use `log.Fatal`.
5. **Rule 5 (Strict Relative Git Paths):** All links and references in documentation and plans must be relative to repository root with zero absolute paths.

---

## 3. Subtask Breakdown

- `.lovable/plans/subtasks/20-fileutil-filename-matching-and-bound-path-ops/01-task-rename-operation-files-to-match-struct-names.md`
- `.lovable/plans/subtasks/20-fileutil-filename-matching-and-bound-path-ops/02-task-bound-file-path-ops-and-singletons.md`
- `.lovable/plans/subtasks/20-fileutil-filename-matching-and-bound-path-ops/03-task-fileutil-readme-and-ai-skill-guide.md`

---

## 4. Verification Plan
- Unit tests in `04-code/golang/pkg/fileutil/file_path_ops_test.go` and `append_ops_test.go`.
- Full package test: `go test -C 04-code/golang -v ./pkg/fileutil`.
- Repository-wide Go test suite: `go test -C 04-code/golang -count=1 ./...`.
- Local CI runner: `python 03-ai-scripts/06-cicd-local-runner.py --all` (all 31 quality gates).
- Metadata sync: `npm run sync` and `node scripts/sync-check.mjs`.
