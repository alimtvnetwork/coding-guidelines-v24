# Subtask 01: Architecture Specification & Contracts for 100 Improvements Suite Renaming

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/03-rename-and-sync-100-improvements-prompts-and-skills/01-subtask-architecture-and-contracts.md`  
> **Parent Plan:** `.ai-memory/plans/pending/03-rename-and-sync-100-improvements-prompts-and-skills.md`  
> **Specification:** `02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills/01-architecture-spec.md`  
> **Status:** `[DONE]`  
> **Assigned Role:** Spec Author 01 (Lead Spec Architect)  

---

## 1. Objective

Architect and document the complete specification and operational contract for renaming the 100 Improvements prompts and skills suite across the repository and connected 43-repo fleet. This subtask establishes:
1. The architectural rationale for shifting from static `find-100-*` naming to continuous flywheel `next-100-*` naming.
2. The exact 1:1 mapping matrix for the 4 prompts in `01-prompts/30-improvements/`.
3. Symmetrical mirroring across Antigravity IDE skills (`.agents/skills/`) and Cursor IDE skills (`.cursor/skills/`).
4. The deletion protocol for legacy prompt and skill paths in the source repository.
5. The purge manifest and synchronization mechanics for propagating changes cleanly to 43 fleet repositories without leaving dead legacy files.
6. Documentation, index, and catalog parity requirements across `.ai-memory/prompts.md`, `01-prompts/readme.md`, and related files.

---

## 2. Context & Guardrail Invariants

- **Parent Master Plan:** `.ai-memory/plans/pending/03-rename-and-sync-100-improvements-prompts-and-skills.md`
- **Application Specification:** `02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills/01-architecture-spec.md`
- **Strict Relative Git Paths:** All file references, documentation links, and subtask citations must strictly use relative paths (`02-spec/...`, `.ai-memory/...`, `01-prompts/...`). TOTAL BAN on absolute filesystem paths and `file:///` URIs.
- **Strict Lowercase Naming:** All paths, folders, and filenames must strictly use lowercase characters.
- **No Git Commands:** Subagents and workers must never execute git commands (`git add`, `git commit`, `git push`, `git status`).
- **No Build / No Test Rule:** Verification must not trigger system builds or test suites.
- **GitMap Exclusivity:** File discovery and searches must strictly use GitMap commands (`gitmap aum search`, `gitmap find`, `gitmap lf`, `gitmap cat`).

---

## 3. Owned Deliverables

| Deliverable Path | Description | Status |
| :--- | :--- | :---: |
| `02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills/01-architecture-spec.md` | Primary architecture specification covering rationale, naming matrix, Mermaid diagram, pillar capabilities, legacy purge, and catalog alignment. | `COMPLETED` |
| `.ai-memory/plans/subtasks/03-rename-and-sync-100-improvements-prompts-and-skills/01-subtask-architecture-and-contracts.md` | Subtask contract, action checklist, verification trace, and downstream worker handoff directives. | `COMPLETED` |

---

## 4. Renaming Mapping Reference

| Asset Type | Legacy Path / Slug | New Path / Slug |
| :--- | :--- | :--- |
| **Prompt 01** | `01-prompts/30-improvements/01-find-100-improvements-in-ui-level.md` | `01-prompts/30-improvements/01-next-100-ui-improvements.md` |
| **Prompt 02** | `01-prompts/30-improvements/02-find-100-improvements-in-code-base.md` | `01-prompts/30-improvements/02-next-100-code-improvements.md` |
| **Prompt 03** | `01-prompts/30-improvements/03-find-100-improvements-in-whole-system.md` | `01-prompts/30-improvements/03-next-100-whole-system-improvements.md` |
| **Prompt 04** | `01-prompts/30-improvements/04-detail-plans-and-improvements-for-review.md` | `01-prompts/30-improvements/04-next-detail-out-100-improvements.md` |
| **Antigravity Skill 01** | `.agents/skills/find-100-improvements-in-ui-level/` | `.agents/skills/next-100-ui-improvements/` |
| **Antigravity Skill 02** | `.agents/skills/find-100-improvements-in-code-base/` | `.agents/skills/next-100-code-improvements/` |
| **Antigravity Skill 03** | `.agents/skills/find-100-improvements-in-whole-system/` | `.agents/skills/next-100-whole-system-improvements/` |
| **Antigravity Skill 04** | `.agents/skills/detail-plans-and-improvements-for-review/` | `.agents/skills/next-detail-out-100-improvements/` |
| **Cursor Skill 01** | `.cursor/skills/find-100-improvements-in-ui-level/` | `.cursor/skills/next-100-ui-improvements/` |
| **Cursor Skill 02** | `.cursor/skills/find-100-improvements-in-code-base/` | `.cursor/skills/next-100-code-improvements/` |
| **Cursor Skill 03** | `.cursor/skills/find-100-improvements-in-whole-system/` | `.cursor/skills/next-100-whole-system-improvements/` |
| **Cursor Skill 04** | `.cursor/skills/detail-plans-and-improvements-for-review/` | `.cursor/skills/next-detail-out-100-improvements/` |

---

## 5. Action Checklist

- [x] Inspect existing prompt suite in `01-prompts/30-improvements/` and master plan in `.ai-memory/plans/pending/03-rename-and-sync-100-improvements-prompts-and-skills.md`.
- [x] Inspect skill structures in `.agents/skills/` and `.cursor/skills/` to establish naming conventions and frontmatter contracts.
- [x] Author comprehensive architecture specification `02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills/01-architecture-spec.md`.
- [x] Formulate rationale for the continuous improvement flywheel (`next-100` vs static `find-100`).
- [x] Define exact naming matrix for all 4 prompts, 4 Antigravity skills, and 4 Cursor skills.
- [x] Design Mermaid architecture diagram (`flowchart TD`) illustrating prompt router, 4 pillars, skill mirrors, memory persistence, and fleet sync.
- [x] Detail technical disciplines for all 4 pillars (UI tokens/depth/contrast, codebase guidelines/AppError/booleans, whole-system UX friction, and multi-step plan detailing with review-first gating).
- [x] Define local deletion list for legacy files in `coding-guidelines`.
- [x] Define fleet purge manifest for target repositories during sync execution.
- [x] Document catalog and linter reconciliation targets (`01-prompts/readme.md`, `.ai-memory/prompts.md`, `check-prompts-loaded.py`).
- [x] Author subtask specification `.ai-memory/plans/subtasks/03-rename-and-sync-100-improvements-prompts-and-skills/01-subtask-architecture-and-contracts.md`.
- [x] Verify zero absolute paths or `file:///` URIs across all authored deliverables.
- [x] Verify strict lowercase filenames across all created files.

---

## 6. Verification Trace & Evidence

### 6.1 Specification File Presence
- `02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills/01-architecture-spec.md` written cleanly.
- Contains 7 structured sections, Mermaid diagram, detailed matrix, and non-negotiable invariants.

### 6.2 Subtask Contract Presence
- `.ai-memory/plans/subtasks/03-rename-and-sync-100-improvements-prompts-and-skills/01-subtask-architecture-and-contracts.md` written cleanly.
- Contains metadata block, objective, context invariants, owned deliverables, renaming mapping, action checklist, and downstream handoff directives.

### 6.3 Invariant Check
- Relative paths: Verified 100% relative paths (`02-spec/...`, `.ai-memory/...`, `01-prompts/...`).
- Lowercase check: Verified all characters in file paths and names are strictly lowercase.
- Zero Git execution: Confirmed zero git commands executed.
- Zero build execution: Confirmed zero builds or test runs executed.

---

## 7. Downstream Worker Handoff Directives

The architectural foundation and operational contracts are fully established. Downstream tasks may proceed in disjoint bounding boxes:
- **Task-02 (Worker 01):** Rename and update the 4 prompts in `01-prompts/30-improvements/` and remove legacy prompt files.
- **Task-03 (Worker 02):** Rename and update the 8 skills across `.agents/skills/` and `.cursor/skills/` and remove legacy skill folders.
- **Task-04 (Worker 03 / Lead):** Update catalogs (`.ai-memory/prompts.md`, `01-prompts/readme.md`), update fleet sync scripts with the purge manifest, run fleet synchronization across 43 repositories, verify linters, and execute atomic commit and push.
