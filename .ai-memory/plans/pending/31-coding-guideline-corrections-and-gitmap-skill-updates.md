# Plan 31: Coding Guideline Corrections & GitMap Skill Updates

> **Status:** Completed  
> **Sequence ID:** 31  
> **Slug:** `coding-guideline-corrections-and-gitmap-skill-updates`  
> **Specification:** `02-spec/21-app/31-coding-guideline-corrections-and-gitmap-skill-updates/`  
> **Subtasks Directory:** `.ai-memory/plans/subtasks/31-coding-guideline-corrections-and-gitmap-skill-updates/`  
> **Date:** 2026-10-27  

---

## 1. Executive Summary & Core Objectives
1. Author new skill `follow-coding-guidelines` with comprehensive checklist and learning loop.
2. Clean and overhaul `gitmap` skills in `.agents/skills/gitmap/skill.md` (and `.cursor/skills/`):
   - Normalize uppercase module examples (`SKILLS`, `CLONER`, `DOCS`, `RELEASE`) to Title Case (`Skills`, `Cloner`, `Docs`) or lowercase.
   - Remove bloated commands (`disk`, `os storage`, `os dev-clean`, `vscode repair`, `starbar client ls`, `cluster worker/status`).
   - Retain Cargo suite and process controls.
3. Inventory and catalog skills reduction: Letterly, Muse, duplicate Python helpers, and fragmented CG skills.
4. GitMap CLI and Show Log UI Enhancements (`d:\work\gitmap`):
   - Short repo name resolution for `gitmap clone`.
   - Prevent exit 128 crash when Show Log is opened in non-git directory (`cmdlogviewer/server.go:85`).
   - Add native process kill command (`gitmap proc kill <pid> [--force]`).
   - Show Log UI: GitMap logo favicon, compact icon buttons with hover tooltips, searchable branch combo, file structure diff comparison, Beyond Compare / VS Code integration, and settings modal.
5. Verification, compilation, and minor release ceremony.

---

## 2. Subtasks Roadmap

| Subtask ID | File | Description | Target | Status |
| :--- | :--- | :--- | :--- | :--- |
| **Subtask-31.1** | `01-spec-and-plans.md` | Author canonical specifications and pending subtask roadmap in `02-spec/21-app/31-coding-guideline-corrections-and-gitmap-skill-updates/` and `.ai-memory/plans/pending/31-coding-guideline-corrections-and-gitmap-skill-updates.md`. | `02-spec/`, `.ai-memory/` | Completed |
| **Subtask-31.2** | `02-follow-coding-guideline-skill.md` | Create new `follow-coding-guidelines` skill in `.agents/skills/follow-coding-guidelines/skill.md` with checklist, learn directives, and lowercase filename. | `.agents/skills/` | Completed |
| **Subtask-31.3** | `03-gitmap-skill-sanitization.md` | Overhaul `.agents/skills/gitmap/skill.md`: normalize module casing in examples, eliminate duplicate commit tables, and remove deprecated bloat. | `.agents/skills/gitmap/` | Completed |
| **Subtask-31.4** | `04-skills-inventory-and-reduction.md` | Inventory skills to archive/reduce (Letterly, Muse, duplicate Python helpers, fragmented CG skills) and design `follow-cg-execute`. | `.agents/skills/` | Completed |
| **Subtask-31.5** | `05-gitmap-cli-and-showlog-ui.md` | GitMap engine updates: short repo clone resolution, non-git Show Log crash fix, native process kill, and Show Log UI modernization (icon, compact buttons, searchable branch, Beyond Compare diff). | `d:\work\gitmap` | Completed |
| **Subtask-31.6** | `06-verification-and-release.md` | Full unit testing, binary compilation & deployment, dry-run verification, and minor release ceremony. | Multi-repo | Completed |

---

## 3. Verification & Acceptance Criteria
- [x] New `follow-coding-guidelines` skill created in lowercase `skill.md` with complete 12-rule checklist and `/learn` loop.
- [x] Zero ALL-CAPS module tags (`SKILLS`, `CLONER`, `DOCS`, `RELEASE`) in `gitmap/skill.md` examples.
- [x] Bloat removed from `gitmap/skill.md` (`disk`, `os dev-clean`, `vscode repair`, `starbar`, `cluster worker`).
- [x] GitMap clone resolves bare repo names (`gitmap clone coding-guidelines`).
- [x] Show Log in non-git directory returns clean error instead of exit 128 stack trace.
- [x] Show Log UI features GitMap logo favicon, compact icon buttons, searchable branch combo, and diff comparison.
- [x] All unit tests pass 100% and minor release ceremony completed.
