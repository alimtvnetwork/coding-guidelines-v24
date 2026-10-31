# Master Audit Ledger: Coding Guideline Corrections & GitMap Skill Updates (Phase 2)

- **Task Slug:** `coding-guideline-corrections-and-gitmap-skill-updates`
- **Sequence:** `32`
- **Specification Directory:** `02-spec/21-app/32-coding-guideline-corrections-and-gitmap-skill-updates/`
- **Plan File:** `.ai-memory/plans/completed/32-coding-guideline-corrections-and-gitmap-skill-updates.md`
- **Subtasks Directory:** `.ai-memory/plans/subtasks/32-coding-guideline-corrections-and-gitmap-skill-updates/`
- **Status:** Completed
- **Created:** 2026-10-27
- **Completed:** 2026-10-27

---

## 1. Executive Summary & Root Cause Matrix

| Component | Symptom / Requirement | Root Cause | Target Fix | Verification |
| :--- | :--- | :--- | :--- | :--- |
| **PowerShell Completion Script** | Tab suggestions vanish after initial completion attempt; PSReadLine error. | In `cli/cmd/root_cobra_completion.go`, string replacement placed `+ $Space` after `{ continue }`, evaluating `{ continue } + $Space` which throws `ScriptBlock does not contain method op_Addition`. | Disambiguated string replacement in `root_cobra_completion.go`: replaced `$CompletionText = $($comp.Name \| ... ) + $Space` first to append `$Space` cleanly to string variable. | Verified in `completions.ps1`: 0 hits for `continue } + $Space`; tab completion operational with zero exceptions. |
| **Process Termination (`gitmap kill`)** | Inconsistent process management; raw PowerShell `Stop-Process` mentioned in skills. | Legacy skill documentation referenced `gitmap pwsh "Stop-Process -Id <pid> -Force"` rather than the native, safe, cross-platform `gitmap kill <pid|name|:port> [--force] [--tree]`. | Updated `gitmap/skill.md`, Cursor mirrors, and internal guidelines to mandate `gitmap kill <pid|name|:port> [--force] [--tree]`. | Verified in skills and helpdoc: total ban on raw shell process termination codified. |
| **Heatmap Query Shorthands** | Multi-line SQLite queries required for common operational audits. | Missing route aliases (`failed`, `unmet`, `hotspots`, `latest`) and lack of shorthand routing in `rc heatmap`. | Added route aliases in `heatmap_cmd.go` and `heatmap_rc_cmd.go` for `failures`, `hotfiles`, `recent`, and `tasks`. | Verified runtime: `gitmap heatmap recent` (6.26ms), `gitmap heatmap failed` (0.00ms), `gitmap heatmap hotfiles` (1.44ms), `gitmap heatmap tasks` (0.00ms). |
| **Heatmap Spec File & Folder Tracking** | `HeatmapParentTask` table in `heatmap.db` lacked runtime insertion API. | `store/heatmap_store.go` lacked `RecordParentTask` method, so only `RcHeatmapStore` tracked tasks. | Added `RecordParentTask` with upsert in `store/heatmap_store.go`, exported `OpenHeatmapDBAtPath`, and wired `gitmap task init` to register in both `agent-task.db` and `heatmap.db`. | Verified runtime: `RootSpecFilePath` and `TaskFolder` recorded and retrievable via `gitmap heatmap query` in 0.00ms. |
| **Follow CG Execute Orchestration** | 39 coding guideline prompts were fragmented across individual skills without a single chained runner. | Users needed a unified skill that executes all 39 prompts sequentially via "next" with zero-build fast-forwarding until the 39th step. | Authored `follow-cg-execute` with explicit 39-step ordered inventory, zero-build execution in steps 1–38, and terminal E2E verification in step 39. | Verified in `.agents/skills/follow-cg-execute/skill.md` and `.cursor/skills/follow-cg-execute/skill.md`. |
| **Skills Inventory Reduction** | Over 120 skills populated `.agents/skills`, including duplicate prompt wrappers. | Redundant release wrappers (`major-bump`, `minor-bump`, `patch-bump`, `version-bump-docs`), logo wrappers, theme wrappers, and sync wrappers bloated the skill directory. | Consolidated wrappers into canonical parent skills (`release`, `logo-and-icon-design`, `react-ui-theming-design`, `prompts-and-skills-sync`, `check-pending-improvements`) and moved 13 wrappers to `skills-archived/`. | Verified in both repos: 13 redundant skills moved to `skills-archived/`. |

---

## 2. Traceability & Action Item Verification

- [x] **Item 1:** Git pull executed across both repositories before making changes.
- [x] **Item 2:** Spec and plans written with strictly relative paths and lowercase filenames.
- [x] **Item 3:** Fix PowerShell completion script generation in `d:/work/gitmap/cli/cmd/root_cobra_completion.go`.
- [x] **Item 4:** Promote native `gitmap kill` across all skills and guidelines.
- [x] **Item 5:** Add shorthand query commands and aliases to `gitmap heatmap` and `gitmap rc heatmap`.
- [x] **Item 6:** Wire spec tracking (`RootSpecFilePath`, `TaskFolder`) into `HeatmapParentTask` during `gitmap task init`.
- [x] **Item 7:** Update `follow-cg-execute` with 39-step sequence, zero-build intermediate passes, and step 39 E2E test verification.
- [x] **Item 8:** Consolidate redundant prompt wrappers into parent skills and archive deprecated folders.
- [x] **Item 9:** Rebuild `gitmap.exe`, run unit tests, and deploy active binary.
- [x] **Item 10:** Execute minor bump release ceremony (`v6.87.0`), tag, and push upstream to GitHub.
