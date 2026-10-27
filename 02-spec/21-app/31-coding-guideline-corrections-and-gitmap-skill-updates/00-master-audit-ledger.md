# Master Audit Ledger: Coding Guideline Corrections & GitMap Skill Updates

> **Specification Sequence:** 31  
> **Slug:** `coding-guideline-corrections-and-gitmap-skill-updates`  
> **Date:** 2026-10-27  
> **Status:** Active / Specification Phase  

---

## 1. Verbatim User Request & Core Objectives

```text
In the coding guideline, a few things need to be corrected:
1. New Skill: follow-coding-guidelines skill in lowercase MD format with comprehensive checklist, learning loop (/learn), and guidance for AI.
2. GitMap Skill Overhaul:
   - Fix uppercase module names in examples (e.g. SKILLS, CLONER, DOCS, RELEASE) to Title Case (Skills, Cloner, Docs) or lowercase.
   - Eliminate repetitive sections and redundant commit tables.
   - Purge bloat: remove gitmap disk, os storage, os dev-clean, VS Code repair, starbar client list, cluster worker, cluster status.
   - Retain cargo commands (cargo install, cargo run) and process control.
   - Create native gitmap process stop/kill command with graceful + forced termination.
3. Skills Inventory & Reduction:
   - Identify obsolete/redundant skills (Letterly prompts, Muse prompts, prompt-library, python array/boolean/constants duplicates, python-dry-caching).
   - Propose follow-cg-execute skill consolidating individual cg-execute prompts.
   - Catalog skills to reduce and discuss prior to removal.
4. GitMap CLI & Show Log UI Improvements:
   - gitmap clone should resolve short repo names (e.g. coding-guidelines -> alimtvnetwork/coding-guidelines-v24).
   - Fix Show Log crash when executed in non-git directory (cmdlogviewer/server.go:85 exit status 128 stack trace).
   - In terminal/shell, Show Log should handle non-git roots gracefully.
   - Show Log UI enhancements: GitMap logo favicon, compact icon buttons with hover tooltips, searchable branch combo, file structure diff comparison, Beyond Compare / VS Code integration, and settings modal.
   - Heatmap query shortcuts: shorten verbose SQL queries into dedicated subcommands; add root spec file and folder tracking to task tables.
```

---

## 2. The 5 Mandatory Pillars

1. **Pillar 1: Specification & Planning Rigor (Spec 31 & Plan 31)**
   Author architectural contracts and subtask decomposition before code execution. Maintain strictly relative Git paths (`02-spec/...`, `.ai-memory/...`).

2. **Pillar 2: Skill Modernization & Casing Hygiene (`follow-coding-guidelines` & `gitmap`)**
   Author `follow-coding-guidelines` skill in lowercase MD file. Eliminate ALL-CAPS module tags (`SKILLS`, `CLONER`, `DOCS`, `RELEASE`) in `gitmap/skill.md`, enforcing Title Case (`Skills`, `Cloner`, `Docs`) and lowercase (`skills`, `cloner`, `docs`). Purge deprecated sections (cluster status, starbar, vscode repair).

3. **Pillar 3: Skill Directory Reduction & De-Duplication**
   Inventory all skills across `.agents/skills/` and `.cursor/skills/`. Catalog Letterly, Muse, duplicate Python, and fragmented CG skills for archiving/consolidation into `follow-cg-execute`.

4. **Pillar 4: GitMap CLI Resilience & Git Clone Automation**
   Implement short-repo resolution in `gitmap clone`. Prevent exit 128 stack trace crashes in Show Log when run outside a Git repo. Add native process kill command (`gitmap proc kill`). Shorten Heatmap SQL queries into subcommands and track root spec paths.

5. **Pillar 5: Show Log UI Modernization & Professional Ergonomics**
   Inject GitMap logo icon, compact window buttons with hover tooltips, searchable branch combo, file comparison via Beyond Compare / VS Code, and configurable settings.

---

## 3. 6-Part Verification Schema

1. **V1: Relative Path Hygiene:** Every link, file reference, and spec path uses strictly relative Git root paths.
2. **V2: Lowercase Filename Mandate:** All new skills, specs, and scripts strictly lowercase (`readme.md`, `skill.md`).
3. **V3: Module Casing & Pseudo-Module Elimination:** No ALL-CAPS module tags or `commit - ` examples in `gitmap/skill.md`.
4. **V4: GitMap Non-Git Root Guard:** `gitmap show-log` executed in a non-git directory returns a graceful error message without stack traces.
5. **V5: Unit Test & Build Verification:** `go test -v ./...` in `gitmap` passes 100%; binary builds and deploys cleanly.
6. **V6: Clean Git Status & Completion Invariant:** Both repositories committed and pushed upstream to GitHub tracking branches.
