# Subtask 03: Fleet Migration, Catalog Registration & Verification

## 1. Subtask Overview

| Attribute | Details |
| :--- | :--- |
| **Subtask ID** | `03-subtask-migration-and-verification` |
| **Parent Task** | `01-create-prompts-for-100-improvements` |
| **Parent Plan Ref** | `.ai-memory/plans/pending/01-create-prompts-for-100-improvements.md` |
| **Specification Ref** | `02-spec/21-app/01-create-prompts-for-100-improvements/03-migration-and-verification-spec.md` |
| **Status** | `PENDING` |
| **Assigned Role** | Worker 03 / Lead Orchestrator |
| **Plan Motto** | Improved work or less friction, faster work and more pure and effective work |

---

## 2. Objective & Scope

Orchestrate the end-to-end catalog registration, fleet-wide multi-repository synchronization across all 43 connected repositories, cognitive heatmap telemetry audit, and atomic commit push ceremony for the **100 Improvements Prompt Suite** (`30-improvements`) and accompanying dual-runtime skills (`.agents/skills/` and `.cursor/skills/`).

### Owned Deliverables & Files in Scope
- `02-spec/21-app/01-create-prompts-for-100-improvements/03-migration-and-verification-spec.md` (Specification)
- `.ai-memory/plans/subtasks/01-create-prompts-for-100-improvements/03-subtask-migration-and-verification.md` (Subtask Plan)
- `01-prompts/readme.md` (Category 30 Directory Index & Table Entry)
- `01-prompts/30-improvements/readme.md` (Category 30 Group Catalog)
- `.ai-memory/prompts.md` (Operational Memory Prompts Matrix)
- Full fleet synchronization of prompts, skills, shared specs (`02-spec/01-*` .. `02-spec/20-*`), and additive AI scripts across all 43 connected repositories via `03-ai-scripts/38-sync-prompts-skills-scripts.py`.

---

## 3. Compliance & Boundary Matrix

| Invariant / Check Category | Rule Requirement | Verification Method | Pass Threshold |
| :--- | :--- | :--- | :--- |
| **Spec 21 Exclusion** | Zero leakage of `02-spec/21-*` through `02-spec/25-*` into target repositories. | Fleet path scan | 0 leaked files |
| **Additive Scripts** | Add new scripts; never overwrite scripts modified by target repositories. | Engine filter verification | 0 overwritten target scripts |
| **Bump Script Protection** | Version bump scripts (`bump*`, `version.json`) remain untouched. | File hash check | 0 modified bump scripts |
| **Memory & Plan Isolation** | `.ai-memory/plans/` and `.ai-memory/memory/` remain strictly isolated per repository. | Directory scan | 0 overwritten plan files |
| **Archive Exclusion** | `06-archive/` and `06-old-prompts/` are never synchronized downstream. | Target repos scan | 0 archive files in fleet |
| **Strict Relative Paths** | All links and file references use strictly relative Git paths (`02-spec/...`, `.ai-memory/...`). TOTAL BAN on `file:///` and absolute paths. | Relative paths linter | 0 absolute path references |
| **Strict Lowercase Naming** | All filenames and folders adhere to lowercase kebab-case (`^[0-9a-z_.-]+$`). | GitMap file scan | 0 uppercase characters |
| **Cognitive Heatmap Closure** | 100% action outcome closure across all phases. Zero dangling logs. | `gitmap heatmap status` & `gitmap heatmap tree` | 100% closed (`met=true`) |
| **Atomic Push Syntax** | Commit via `gitmap cpf "prompts - add 100 improvements prompt suite and skills"` (hyphen format, no colons in message). | Commit log check | Clean commit + pushed upstream |

---

## 4. Detailed Action Checklist

### Phase 1: Preflight Verification & Deliverables Audit
- [ ] Verify that Task-02 deliverables exist and are non-empty:
  - `01-prompts/30-improvements/01-find-100-improvements-in-ui-level.md`
  - `01-prompts/30-improvements/02-find-100-improvements-in-code-base.md`
  - `01-prompts/30-improvements/03-find-100-improvements-in-whole-system.md`
  - `01-prompts/30-improvements/readme.md`
- [ ] Verify that Task-03 deliverables exist and are symmetrical:
  - .agents/skills/find-100-improvements-in-ui-level/skill.md
  - .agents/skills/find-100-improvements-in-code-base/skill.md
  - .agents/skills/find-100-improvements-in-whole-system/skill.md
  - `.cursor/skills/find-100-improvements-in-ui-level/skill.md`
  - `.cursor/skills/find-100-improvements-in-code-base/skill.md`
  - `.cursor/skills/find-100-improvements-in-whole-system/skill.md`
- [ ] Verify that all application specifications in `02-spec/21-app/01-create-prompts-for-100-improvements/` exist:
  - `01-prompts-architecture-spec.md`
  - `02-skills-suite-spec.md`
  - `03-migration-and-verification-spec.md`
- [ ] Verify that all subtask plans in `.ai-memory/plans/subtasks/01-create-prompts-for-100-improvements/` exist:
  - `01-subtask-prompts-suite.md`
  - `02-subtask-skills-suite.md`
  - `03-subtask-migration-and-verification.md`

### Phase 2: Catalog Registration & Integrity Verification
- [ ] Update `01-prompts/readme.md`:
  - Add Category `30-improvements/` to directory index tree.
  - Add Category `30` row to Table of Categories with descriptions and key prompts.
- [ ] Verify `01-prompts/30-improvements/readme.md`:
  - Ensure canonical entry point structure conforming to Rule 13 (unadorned `readme.md`, lowercase).
  - Document all three prompts with execution context, triggers, and pending task output path.
- [ ] Update `.ai-memory/prompts.md`:
  - Add Prompts Matrix entries for all three prompts and group readme with relative links.
- [ ] Run prompt catalog integrity validation:
  ```bash
  python linter-scripts/check-prompts-loaded.py
  ```

### Phase 3: Fleet Synchronization Dry-Run & Live Execution
- [ ] Execute dry-run to preview mirrored files and verify 5 boundary exclusions across all 43 repositories:
  ```bash
  python 03-ai-scripts/38-sync-prompts-skills-scripts.py --dry-run
  ```
- [ ] Execute live multi-repository synchronization:
  ```bash
  python 03-ai-scripts/38-sync-prompts-skills-scripts.py
  ```
- [ ] Confirm clean run with 0 errors across all 43 target repositories:
  - Base branch pull succeeded without conflicts.
  - Backup branch (`backup/sync-<timestamp>`) created and pushed to GitHub for each repository.
  - Assets mirrored cleanly with boundary filters active.
  - Changes committed on base branch via `gitmap cpc` and pushed upstream.

### Phase 4: Downstream Repository Boundary & Leak Verification
- [ ] Verify Spec 21 Boundary: Ensure zero `02-spec/21-*` files were copied to any downstream target repository.
- [ ] Verify Bump Script Boundary: Ensure `version.json` and `bump*` scripts remain un-mutated in all target repositories.
- [ ] Verify Memory & Plan Boundary: Ensure target repositories' `.ai-memory/plans/` and `.ai-memory/memory/` remain intact.
- [ ] Verify Archive Boundary: Ensure zero `06-archive/` or `06-old-prompts/` files exist in any target repository.

### Phase 5: GitMap Cognitive Heatmap Telemetry Verification
- [ ] Run GitMap Cognitive Heatmap status inspection:
  ```bash
  gitmap heatmap status --task "01-create-prompts-for-100-improvements"
  ```
- [ ] Run GitMap Cognitive Heatmap tree inspection:
  ```bash
  gitmap heatmap tree --task "01-create-prompts-for-100-improvements"
  ```
- [ ] Verify 100% action closure:
  - Every file creation, edit, and audit action marked closed (`met=true`).
  - Zero open, unverified, or dangling actions remaining.

### Phase 6: Final Atomic Semantic Commit & Upstream Push
- [ ] Check working tree status:
  - Verify zero unstaged or uncommitted files.
  - Verify zero temporary files or secrets leaked into repository.
- [ ] Execute atomic commit and push upstream following `AGENTS.md` Section 10 and 13:
  ```bash
  gitmap cpf "prompts - add 100 improvements prompt suite and skills"
  ```
- [ ] Verify upstream tracking branch:
  - Confirm `Feature: prompts - add 100 improvements prompt suite and skills` is committed.
  - Confirm changes are successfully pushed upstream to GitHub.
  - Confirm working tree is completely clean ('No Push = Not Done' satisfied).

---

## 5. Verification Criteria & Sign-Off Checklist

- [ ] All three application specifications exist in `02-spec/21-app/01-create-prompts-for-100-improvements/`.
- [ ] All three subtask plans exist in `.ai-memory/plans/subtasks/01-create-prompts-for-100-improvements/`.
- [ ] Zero absolute paths (`C:\`, `/home/`, `file:///`) across all authored specifications and plans.
- [ ] Zero uppercase characters in any authored filenames or directory paths.
- [ ] All three canonical prompts in `01-prompts/30-improvements/` verified.
- [ ] All six agent and Cursor skills in `.agents/skills/` and `.cursor/skills/` verified.
- [ ] Master catalogs (`01-prompts/readme.md`, `01-prompts/30-improvements/readme.md`, `.ai-memory/prompts.md`) registered.
- [ ] Fleet synchronization executed across all 43 connected repositories with 0 errors.
- [ ] The 5 Non-Negotiable Boundaries verified across all downstream repositories.
- [ ] Cognitive Heatmap reports 100% action closure (`met=true`).
- [ ] Final atomic commit pushed upstream via `gitmap cpf "prompts - add 100 improvements prompt suite and skills"`.
