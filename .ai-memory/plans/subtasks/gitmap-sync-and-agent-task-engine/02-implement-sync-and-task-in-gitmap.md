# Subtask 02: Implement `gitmap sync` & `gitmap task` in Go

> **Subtask ID:** `Task-02`  
> **Parent Plan:** `.ai-memory/plans/gitmap-sync-and-agent-task-engine.md`  
> **Status:** PENDING  
> **Assigned Agent Role:** `GoEngineWorker`

---

## 1. Objectives
1. Implement `gitmap sync` command in `../gitmap`:
   - Built-in default 43 fleet repositories registry.
   - `--projects <json|path>` flag supporting JSON manifest of paths, URLs, and folders.
   - High-speed goroutine worker pool with `--workers` flag.
   - Pre-pull and safety backup branch creation (`backup/sync-<timestamp>`).
   - Strict boundary protections (Spec 21-25 exclusion, Additive scripts, Bump script protection, Memory & Plans protection).
   - Fast file copying and hashing.
   - Version bump and release tagging (with `--dry-run`, `--no-push`, `--no-release`).
2. Implement `gitmap task` (or `gitmap agent task`) command in `../gitmap`:
   - `init`, `add`, `claim`, `complete`, `status`, and `schema` subcommands.
   - High-speed direct SQLite operations without Python interpreter overhead.
3. Verify compilation and run unit tests.

---

## 2. Deliverables
- Working Go implementation of `gitmap sync` and `gitmap task`.
- Successful binary compilation and test passing in `../gitmap`.
