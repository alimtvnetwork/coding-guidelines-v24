# Master Execution Plan: Multi-Repository Synchronization & Canonical Mirroring

> **Plan ID:** `sync-prompts-and-repos`  
> **Status:** ACTIVE  
> **Budget:** N = 300 steps  
> **Concurrency:** A = 2, H = 2  

---

## 1. Objectives & Overview

1. Verify and establish the sequenced prompt sync folder `01-prompts/24-sync/` with `01-sync.md` and matching skills for Antigravity (`.agents/skills/sync/`) and Cursor (`.cursor/skills/sync/`).
2. Add and harden the two requested CG-execute prompts (`31-cg-execute-in-below-steps.md` and `32-cg-follow-other-prompts.md`) and the V4 prompt folder structure.
3. Update memory rules and learned boundaries in `.ai-memory/memory/` and `AGENTS.md` regarding Spec 21 exclusion, additive-only AI scripts, and bump script protection.
4. Execute full multi-repository pull, safety backup branching (`backup/sync-<timestamp>`), and safe mirroring across connected repositories using `03-ai-scripts/38-sync-prompts-skills-scripts.py`.

---

## 2. Decomposed Subtask Breakdown

- **Subtask 01 (`Task-01`):** `01-create-sync-folder-and-skills.md` — Ensure `01-prompts/24-sync/01-sync.md` contains the clean repository sequence and sync skills exist for both Antigravity and Cursor.
- **Subtask 02 (`Task-02`):** `02-cg-execute-prompts-and-v4-folder.md` — Validate and finalize `31-cg-execute-in-below-steps.md`, `32-cg-follow-other-prompts.md` in `01-prompts/15-cg-execute/`, and establish the V4 prompt archive/folder.
- **Subtask 03 (`Task-03`):** `03-memory-rules-and-sync-script-protection.md` — Document and persist cross-repo synchronization rules in `.ai-memory/memory/` to guarantee permanent adherence.
- **Subtask 04 (`Task-04`):** `04-pre-pull-backup-and-multi-repo-sync.md` — Pull target repositories, create safety backup branches, and execute synchronization across target repositories.

---

## 3. Subtask Traceability Matrix

| Subtask Code | Title | Owned Files | Status |
| :--- | :--- | :--- | :--- |
| `Task-01` | Sequenced Sync Prompt & Skills | `01-prompts/24-sync/01-sync.md`, `.agents/skills/sync/`, `.cursor/skills/sync/` | IN_PROGRESS |
| `Task-02` | CG Execute & V4 Prompts | `01-prompts/15-cg-execute/`, `01-prompts/v4/`, `06-archive/v4/` | QUEUED |
| `Task-03` | Memory Sync Invariants | `.ai-memory/memory/learned/18-cross-repository-sync-rules.md`, `.ai-memory/what-to-read.md` | QUEUED |
| `Task-04` | Multi-Repo Pull, Backup & Sync | `d:\work\*` connected repositories via `38-sync-prompts-skills-scripts.py` | QUEUED |
