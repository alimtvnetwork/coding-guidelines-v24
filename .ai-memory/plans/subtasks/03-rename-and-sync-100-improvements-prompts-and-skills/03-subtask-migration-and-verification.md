# Subtask 03: Fleet Migration, Deprecation Purge & Quality Verification

## 1. Subtask Overview

| Attribute | Details |
| :--- | :--- |
| **Subtask ID** | `03-subtask-migration-and-verification` |
| **Parent Task** | `03-rename-and-sync-100-improvements-prompts-and-skills` |
| **Parent Plan Ref** | `.ai-memory/plans/pending/03-rename-and-sync-100-improvements-prompts-and-skills.md` |
| **Specification Ref** | `02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills/03-migration-and-verification-spec.md` |
| **Status** | `PENDING` |
| **Assigned Role** | Worker 03 / Lead Orchestrator |
| **Plan Motto** | Improved work or less friction, faster work and more pure and effective work |

---

## 2. Objective & Scope

Orchestrate the end-to-end master catalog updates, automated deprecation purge engine enhancements in `03-ai-scripts/38-sync-prompts-skills-scripts.py`, comprehensive linter quality gates, fleet-wide multi-repository synchronization across all 43 connected repositories, cognitive heatmap telemetry audit, and atomic commit push ceremony for the **Renamed 100 Improvements Prompts & Skills Suite** (`next-100-ui-improvements`, `next-100-code-improvements`, `next-100-whole-system-improvements`, and `next-detail-out-100-improvements`).

### Owned Deliverables & Files in Scope
- `02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills/03-migration-and-verification-spec.md` (Specification)
- `.ai-memory/plans/subtasks/03-rename-and-sync-100-improvements-prompts-and-skills/03-subtask-migration-and-verification.md` (Subtask Plan)
- `03-ai-scripts/38-sync-prompts-skills-scripts.py` (Purge logic for legacy prompts and skills)
- `.agents/scripts/38-sync-prompts-skills-scripts.py` (Mirrored purge logic for legacy prompts and skills)
- `01-prompts/readme.md` (Category 30 row update registering renamed prompts suite)
- `01-prompts/30-improvements/readme.md` (Category 30 Group Catalog documenting renamed suite)
- `.ai-memory/prompts.md` (Operational Memory Prompts Matrix linking renamed prompts suite)
- `.ai-memory/plans/readme.md` (Pending Plans Index linking task 03)
- `02-spec/21-app/readme.md` (Application Domain Specifications Index listing task 03 container)
- Full fleet purge and synchronization of prompts, skills, shared specs (`02-spec/01-*` through `02-spec/20-*`), and additive AI scripts across all 43 connected repositories.

---

## 3. Compliance & Boundary Matrix

| Invariant / Check Category | Rule Requirement | Verification Method | Pass Threshold |
| :--- | :--- | :--- | :--- |
| **Deprecation Purge** | Legacy prompts (`01-find-100-*`, `04-detail-plans-*`) and legacy skills (`find-100-*`, `detail-plans-*`) must be purged from target repos before mirroring. | Target repo filesystem scan | 0 legacy files/folders in fleet |
| **Symmetrical Renaming** | Renaming must apply identically across prompts (`01-prompts/30-improvements/`), agent skills (`.agents/skills/`), and Cursor skills (`.cursor/skills/`). | GitMap file check | 4 prompts + 8 skills verified |
| **Spec 21 Exclusion** | Zero leakage of `02-spec/21-*` through `02-spec/25-*` into target repositories. | Fleet path scan | 0 leaked files |
| **Additive Scripts** | Add new scripts; never overwrite scripts modified by target repositories. | Engine filter verification | 0 overwritten target scripts |
| **Bump Script Protection** | Version bump scripts (`bump*`, `version.json`) remain untouched. | File hash check | 0 modified bump scripts |
| **Memory & Plan Isolation** | `.ai-memory/plans/` and `.ai-memory/memory/` remain strictly isolated per repository. | Directory scan | 0 overwritten plan files |
| **Archive Exclusion** | `06-archive/` and `06-old-prompts/` are never synchronized downstream. | Target repos scan | 0 archive files in fleet |
| **Strict Relative Paths** | All links and file references use strictly relative Git paths (`02-spec/...`, `.ai-memory/...`). TOTAL BAN on `file:///` and absolute paths. | `linter-scripts/check-relative-paths.py` | 0 absolute path references |
| **Forbidden Strings** | Zero forbidden patterns, unescaped tokens, or banned syntax. | `linter-scripts/check-forbidden-strings.py` | 0 violations |
| **Prompts Loaded Contract** | Master index `.ai-memory/prompts.md` matches disk prompts directory exactly. | `linter-scripts/check-prompts-loaded.py` | Exit code 0 (100% match) |
| **Strict Lowercase Naming** | All filenames and folders adhere to lowercase kebab-case (`^[0-9a-z_.-]+$`). | GitMap file scan | 0 uppercase characters |
| **Cognitive Heatmap Closure** | 100% action outcome closure across all phases. Zero dangling logs. | `gitmap heatmap status` & `gitmap heatmap tree` | 100% closed (`met=true`) |
| **Atomic Push Syntax** | Commit via `gitmap cpf "prompts - rename 100 improvements suite to next-100 and sync fleet"` (hyphen format, no colons in message). | Commit log check | Clean commit + pushed upstream |

---

## 4. Detailed Action Checklist

### Phase 1: Preflight Verification & Deliverables Audit
- [ ] Verify Task-02 prompt deliverables exist, are non-empty, and legacy prompts are deleted:
  - `01-prompts/30-improvements/01-next-100-ui-improvements.md` exists.
  - `01-prompts/30-improvements/02-next-100-code-improvements.md` exists.
  - `01-prompts/30-improvements/03-next-100-whole-system-improvements.md` exists.
  - `01-prompts/30-improvements/04-next-detail-out-100-improvements.md` exists.
  - Confirm old prompts (`01-find-100-*`, `04-detail-plans-*`) are deleted from `01-prompts/30-improvements/`.
- [ ] Verify Task-03 skill deliverables exist symmetrically across runtimes, and legacy skills are deleted:
  - `.agents/skills/next-100-ui-improvements/skill.md` exists.
  - `.agents/skills/next-100-code-improvements/skill.md` exists.
  - `.agents/skills/next-100-whole-system-improvements/skill.md` exists.
  - `.agents/skills/next-detail-out-100-improvements/skill.md` exists.
  - `.cursor/skills/next-100-ui-improvements/skill.md` exists.
  - `.cursor/skills/next-100-code-improvements/skill.md` exists.
  - `.cursor/skills/next-100-whole-system-improvements/skill.md` exists.
  - `.cursor/skills/next-detail-out-100-improvements/skill.md` exists.
  - Confirm legacy skills (`find-100-*`, `detail-plans-*`) are deleted from `.agents/skills/` and `.cursor/skills/`.
- [ ] Verify that Task-01 application specifications exist in `02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills/`:
  - `01-architecture-spec.md`
  - `02-component-spec.md`
  - `03-migration-and-verification-spec.md`
- [ ] Verify that Task-01 subtask plans exist in `.ai-memory/plans/subtasks/03-rename-and-sync-100-improvements-prompts-and-skills/`:
  - `01-subtask-architecture-and-planning.md`
  - `02-subtask-prompt-and-skills.md`
  - `03-subtask-migration-and-verification.md`

### Phase 2: Master Catalog Updates & Registration
- [ ] Update `01-prompts/readme.md`:
  - Update Category 30 row in Table of Categories to register the renamed suite (`01-next-100-ui-improvements.md`, `02-next-100-code-improvements.md`, `03-next-100-whole-system-improvements.md`, `04-next-detail-out-100-improvements.md`).
- [ ] Update `01-prompts/30-improvements/readme.md`:
  - Update Directory Index & Catalog table with renamed files, versions, and descriptions.
  - Document the unified naming conventions and the review-first gate in prompt 04.
- [ ] Update `.ai-memory/prompts.md`:
  - Replace Category `30-improvements` entries with the four renamed prompt entries.
- [ ] Update `.ai-memory/plans/readme.md`:
  - Add entry under `## Pending Plans` for `03-rename-and-sync-100-improvements-prompts-and-skills.md`.
- [ ] Update `02-spec/21-app/readme.md`:
  - Add entry under `## Active Contents` for `03-rename-and-sync-100-improvements-prompts-and-skills/`.

### Phase 3: Purge Mechanics Implementation in Synchronization Engine
- [ ] Update `03-ai-scripts/38-sync-prompts-skills-scripts.py` (and `.agents/scripts/38-sync-prompts-skills-scripts.py`):
  - Add `DEPRECATED_IMPROVEMENT_PROMPTS` list covering the four old prompt paths.
  - Add `DEPRECATED_IMPROVEMENT_SKILLS` list covering the four old agent skills and four old cursor skills.
  - Insert target purge loop inside `sync_repo` before mirroring assets.
  - Ensure purge logic is wrapped in defensive try/except and respects `--dry-run`.

### Phase 4: Linter Verification Suite Execution
- [ ] Run relative paths linter:
  ```bash
  python linter-scripts/check-relative-paths.py
  ```
  *Requirement:* 0 absolute path or `file:///` violations.
- [ ] Run forbidden strings linter:
  ```bash
  python linter-scripts/check-forbidden-strings.py
  ```
  *Requirement:* 0 forbidden string violations.
- [ ] Run prompts index verification linter:
  ```bash
  python linter-scripts/check-prompts-loaded.py
  ```
  *Requirement:* Exit code 0 (all prompt files on disk indexed in `.ai-memory/prompts.md`).

### Phase 5: Fleet Synchronization Dry-Run & Live Execution
- [ ] Execute dry-run to preview purge and mirror operations across all 43 repositories:
  ```bash
  python 03-ai-scripts/38-sync-prompts-skills-scripts.py --dry-run
  ```
  *Requirement:* Confirm planned purge operations for deprecated prompts and skills in downstream repos.
- [ ] Execute live multi-repository synchronization:
  ```bash
  python 03-ai-scripts/38-sync-prompts-skills-scripts.py
  ```
- [ ] Confirm clean run with 0 errors across all 43 target repositories:
  - Base branch pull succeeded without conflicts (`git pull origin <base_branch> --no-rebase`).
  - Safety backup branch (`backup/sync-<timestamp>`) created and pushed to GitHub for each repository.
  - Deprecated prompts and skills purged cleanly from target repositories.
  - Renamed canonical prompts and skills mirrored cleanly.
  - Changes committed on base branch via `gitmap cpc` and pushed upstream.

### Phase 6: Downstream Repository Boundary & Purge Verification
- [ ] Verify Deprecation Purge: Ensure zero legacy `01-find-100-*`, `04-detail-plans-*`, or old skill folders exist in any target repo.
- [ ] Verify Renamed Assets: Ensure new `01-next-100-*`, `04-next-detail-*`, and `next-*` skills exist in all target repos.
- [ ] Verify Spec 21 Boundary: Ensure zero `02-spec/21-*` through `02-spec/25-*` files were copied to any target repository.
- [ ] Verify Bump Script Boundary: Ensure `version.json` and `bump*` scripts remain un-mutated across all target repositories.
- [ ] Verify Memory & Plan Boundary: Ensure target repositories' `.ai-memory/plans/` and `.ai-memory/memory/` remain intact.
- [ ] Verify Archive Boundary: Ensure zero `06-archive/` or `06-old-prompts/` files exist in any target repository.
- [ ] Verify Additive Scripts Boundary: Ensure target repositories maintain existing custom scripts.

### Phase 7: GitMap Cognitive Heatmap Telemetry Verification & Final Commit
- [ ] Run GitMap Cognitive Heatmap status inspection:
  ```bash
  gitmap heatmap status --task "03-rename-and-sync-100-improvements-prompts-and-skills"
  ```
- [ ] Run GitMap Cognitive Heatmap tree inspection:
  ```bash
  gitmap heatmap tree --task "03-rename-and-sync-100-improvements-prompts-and-skills"
  ```
- [ ] Verify 100% action closure:
  - Every file creation, edit, purge, and audit action marked closed (`met=true`).
  - Zero open, unverified, or dangling actions remaining.
- [ ] Check working tree status:
  - Verify zero unstaged or uncommitted files (`gitmap status`).
  - Verify zero temporary files or secrets leaked into repository.
- [ ] Execute atomic commit and push upstream following `AGENTS.md` Section 10 and 13:
  ```bash
  gitmap cpf "prompts - rename 100 improvements suite to next-100 and sync fleet"
  ```
- [ ] Verify upstream tracking branch:
  - Confirm `Feature: prompts - rename 100 improvements suite to next-100 and sync fleet` is committed.
  - Confirm changes are successfully pushed upstream to GitHub.
  - Confirm working tree is completely clean ('No Push = Not Done' satisfied).

---

## 5. Verification Criteria & Sign-Off Checklist

- [ ] All three application specifications exist in `02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills/`.
- [ ] All three subtask plans exist in `.ai-memory/plans/subtasks/03-rename-and-sync-100-improvements-prompts-and-skills/`.
- [ ] Zero absolute paths (`C:\`, `/home/`, `file:///`) across all authored specifications and plans.
- [ ] Zero uppercase characters in any authored filenames or directory paths.
- [ ] Four renamed canonical prompts verified in `01-prompts/30-improvements/`.
- [ ] Legacy prompt files verified deleted from `01-prompts/30-improvements/`.
- [ ] Eight renamed skills verified across `.agents/skills/` and `.cursor/skills/`.
- [ ] Legacy skill folders verified deleted from `.agents/skills/` and `.cursor/skills/`.
- [ ] Purge mechanics implemented and verified in `03-ai-scripts/38-sync-prompts-skills-scripts.py` and `.agents/scripts/38-sync-prompts-skills-scripts.py`.
- [ ] Master catalogs (`01-prompts/readme.md`, `01-prompts/30-improvements/readme.md`, `.ai-memory/prompts.md`, `.ai-memory/plans/readme.md`, `02-spec/21-app/readme.md`) updated and verified.
- [ ] Linters (`check-relative-paths.py`, `check-forbidden-strings.py`, `check-prompts-loaded.py`) executed and passed cleanly.
- [ ] Fleet synchronization executed across all 43 connected repositories with 0 errors.
- [ ] Target repositories verified free of legacy prompts and legacy skills.
- [ ] The 5 Non-Negotiable Boundaries verified across all downstream repositories.
- [ ] Cognitive Heatmap reports 100% action closure (`met=true`).
- [ ] Final atomic commit pushed upstream via `gitmap cpf "prompts - rename 100 improvements suite to next-100 and sync fleet"`.
