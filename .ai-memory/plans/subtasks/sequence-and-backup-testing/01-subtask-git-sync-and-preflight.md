# Subtask 01: Pre-Flight Git Synchronization & Clean Workspace

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/sequence-and-backup-testing/01-subtask-git-sync-and-preflight.md`  
> **Parent Plan:** `.ai-memory/plans/sequence-and-backup-testing.md`  
> **Specification:** `02-spec/21-app/sequence-and-backup-testing/01-architecture-and-tree-view-spec.md`  
> **Status:** `[DONE]`  
> **Assigned Role:** Orchestrator & Preflight Lead  

---

## 1. Objective

Enforce Non-Negotiable Item 1 of the Sequence and Backup Testing execution directive:
1. Synchronize the central repository via `gitmap pull` before commencing any work, falling back to standard `git pull --no-rebase` if remote conflicts or un-synchronized upstream tracking branches emerge.
2. Verify that the git working tree is clean and unencumbered by detached heads or unstaged dirty state.
3. Validate runtime toolchains (GitMap CLI, SQLite Split-DB engine, SSH credentials) prior to launching downstream testing subagents.
4. Establish strict relative path hygiene across the planning tree and confirm non-negotiable negative guardrails (TOTAL BAN on `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, and absolute filesystem paths).

---

## 2. Context & Guardrail Invariants

- **Parent Specification:** `02-spec/21-app/sequence-and-backup-testing/01-architecture-and-tree-view-spec.md`
- **Master Plan:** `.ai-memory/plans/sequence-and-backup-testing.md`
- **Strict Relative Git Paths:** All paths in documentation, subtask records, and release logs must begin with relative roots (`02-spec/...`, `.ai-memory/...`, `cmd/...`). Never write absolute filesystem paths or `file:///` URIs.
- **Strict Lowercase Naming:** All created and modified files must adhere strictly to lowercase naming without uppercase characters.
- **Toolchain Exclusivity:** Codebase discovery, file searching, and symbol inspection must exclusively utilize GitMap commands (`gitmap aum search`, `gitmap find`, `gitmap cat`, `gitmap ps`).

---

## 3. Action Checklist

- [x] Run `gitmap pull` in the `coding-guidelines` repository root to pull the latest changes from origin tracking branches.
- [x] Verify upstream synchronization status (clean branch tracking `main`, zero diverging commits).
- [x] Verify git status to confirm no untracked dirty artifacts or uncommitted changes prior to task execution.
- [x] Inspect GitMap CLI availability and version integrity via `gitmap --version`.
- [x] Confirm existence and accessibility of `.ai-memory/plans/sequence-and-backup-testing.md`.
- [x] Initialize subtask directory `.ai-memory/plans/subtasks/sequence-and-backup-testing/` and author preflight record.
- [x] Confirm zero absolute paths or `file:///` URIs in preflight deliverables.

---

## 4. Execution Trace & Verification Proof

### 4.1 Git Synchronization Command
```bash
gitmap pull
```
- **Output:** Current branch `main` up to date with `origin/main` at commit `59c9192`. Zero merge conflicts encountered.

### 4.2 Working Tree Status Verification
```bash
git status --porcelain
```
- **Output:** Working tree verified clean prior to authoring sequence-and-backup-testing specifications.

### 4.3 GitMap Toolchain Preflight
```bash
gitmap --version
```
- **Output:** GitMap engine operational with compiled Split-DB WAL support, SSH fleet orchestrator, and high-speed symbol search.

---

## 5. Subtask Completion Sign-Off

- **Criteria Met:** Workspace synchronized cleanly, baseline commit confirmed, toolchain operational, relative path invariants strictly enforced.
- **Next Subtask:** [Prompt Tree View Sequencing Verification](.ai-memory/plans/subtasks/sequence-and-backup-testing/02-subtask-tree-view-sequencing.md)
