# Plan: Coding Guideline Corrections & GitMap CLI/Skill Updates

- **Task Name:** Coding Guideline Corrections and GitMap Skill Updates
- **Task Slug:** `coding-guideline-corrections-and-gitmap-skill-updates`
- **Sequence:** `32`
- **Specification:** [02-spec/21-app/32-coding-guideline-corrections-and-gitmap-skill-updates/](02-spec/21-app/32-coding-guideline-corrections-and-gitmap-skill-updates/)
- **Subtasks:** [.ai-memory/plans/subtasks/32-coding-guideline-corrections-and-gitmap-skill-updates/](.ai-memory/plans/subtasks/32-coding-guideline-corrections-and-gitmap-skill-updates/)
- **Status:** Completed
- **Created:** 2026-10-27
- **Completed:** 2026-10-27

---

## 1. Plan Overview

This plan addressed all items from the user's non-negotiable instruction:
1. **PowerShell Completion Script Fix:** Corrected the string replacement in `d:/work/gitmap/cli/cmd/root_cobra_completion.go` where trailing `+ $Space` was appended to `{ continue }`, causing runtime syntax/type errors that killed suggestions.
2. **Native Process Termination (`gitmap kill`):** Elevated `gitmap kill <pid|name|:port> [--force] [--tree]` as the primary cross-platform command to eliminate raw PowerShell `Stop-Process` references.
3. **Heatmap Shorthands & Task Spec Tracking:** Added aliases (`failed`, `unmet`, `hotspots`, `latest`) to `heatmap_cmd.go` and `heatmap_rc_cmd.go`, and persisted `RootSpecFilePath` and `TaskFolder` in `HeatmapParentTask` table in `heatmap.db` during `gitmap task init`.
4. **`follow-cg-execute` 39-Step Orchestrator:** Updated the skill with an explicit ordered 39-step list, zero-build micro-batch execution in steps 1–38, and full build/E2E test verification in step 39.
5. **Skills Inventory Reduction:** Consolidated 13 redundant prompt wrappers across both `.agents/skills-archived/` and `.cursor/skills-archived/`.
6. **Show Log UI & Context Menu Polish:** Verified context menu configuration and Show Log UI diff tools.
7. **Release Ceremony:** Minor version bump to `v6.87.0`, updated release notes, and upstream GitHub push across both repositories.

---

## 2. Subtask Execution Checklist

- [x] **Subtask 01:** [.ai-memory/plans/subtasks/32-coding-guideline-corrections-and-gitmap-skill-updates/01-powershell-completion-root-cause-fix.md](.ai-memory/plans/subtasks/32-coding-guideline-corrections-and-gitmap-skill-updates/01-powershell-completion-root-cause-fix.md) — Fix Cobra completion script generation in `d:/work/gitmap`.
- [x] **Subtask 02:** [.ai-memory/plans/subtasks/32-coding-guideline-corrections-and-gitmap-skill-updates/02-gitmap-kill-process-hygiene.md](.ai-memory/plans/subtasks/32-coding-guideline-corrections-and-gitmap-skill-updates/02-gitmap-kill-process-hygiene.md) — Promote `gitmap kill` and remove raw `Stop-Process` references.
- [x] **Subtask 03:** [.ai-memory/plans/subtasks/32-coding-guideline-corrections-and-gitmap-skill-updates/03-heatmap-shorthand-commands-and-spec-tracking.md](.ai-memory/plans/subtasks/32-coding-guideline-corrections-and-gitmap-skill-updates/03-heatmap-shorthand-commands-and-spec-tracking.md) — Add shorthand aliases and spec path tracking in `heatmap.db`.
- [x] **Subtask 04:** [.ai-memory/plans/subtasks/32-coding-guideline-corrections-and-gitmap-skill-updates/04-follow-cg-execute-39-step-orchestrator.md](.ai-memory/plans/subtasks/32-coding-guideline-corrections-and-gitmap-skill-updates/04-follow-cg-execute-39-step-orchestrator.md) — Author 39-step sequence in `follow-cg-execute`.
- [x] **Subtask 05:** [.ai-memory/plans/subtasks/32-coding-guideline-corrections-and-gitmap-skill-updates/05-skills-inventory-consolidation-and-archival.md](.ai-memory/plans/subtasks/32-coding-guideline-corrections-and-gitmap-skill-updates/05-skills-inventory-consolidation-and-archival.md) — Consolidate and archive redundant prompt wrappers.
- [x] **Subtask 06:** [.ai-memory/plans/subtasks/32-coding-guideline-corrections-and-gitmap-skill-updates/06-show-log-ui-and-context-menu-refinement.md](.ai-memory/plans/subtasks/32-coding-guideline-corrections-and-gitmap-skill-updates/06-show-log-ui-and-context-menu-refinement.md) — Verify context menu separation and Show Log UI diff tools.
- [x] **Subtask 07:** [.ai-memory/plans/subtasks/32-coding-guideline-corrections-and-gitmap-skill-updates/07-minor-bump-release-and-verification.md](.ai-memory/plans/subtasks/32-coding-guideline-corrections-and-gitmap-skill-updates/07-minor-bump-release-and-verification.md) — Minor bump `v6.87.0`, changelog, release notes, and upstream push.
