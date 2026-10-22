# Subtask 03: Fleet Migration, Catalog Registration & Verification

## 1. Subtask Overview

| Attribute | Details |
| :--- | :--- |
| **Subtask ID** | `03-subtask-migration-and-verification` |
| **Parent Task** | `02-detail-plans-and-improvements-for-review` |
| **Parent Plan Ref** | `.ai-memory/plans/pending/02-detail-plans-and-improvements-for-review.md` |
| **Specification Ref** | `02-spec/21-app/02-detail-plans-and-improvements-for-review/03-migration-and-verification-spec.md` |
| **Status** | `PENDING` |
| **Assigned Role** | Worker 03 / Lead Orchestrator |
| **Plan Motto** | Improved work or less friction, faster work and more pure and effective work |

---

## 2. Objective & Scope

Orchestrate the end-to-end master catalog updates, linter quality checks, fleet-wide multi-repository synchronization across all 43 connected repositories, cognitive heatmap telemetry audit, and atomic commit push ceremony for the **Detail Plans and Improvements for Review Prompt & Skills Suite** (`01-prompts/30-improvements/04-detail-plans-and-improvements-for-review.md` and accompanying skills in `.agents/skills/detail-plans-and-improvements-for-review/` and `.cursor/skills/detail-plans-and-improvements-for-review/`).

### Owned Deliverables & Files in Scope
- `02-spec/21-app/02-detail-plans-and-improvements-for-review/03-migration-and-verification-spec.md` (Specification)
- `.ai-memory/plans/subtasks/02-detail-plans-and-improvements-for-review/03-subtask-migration-and-verification.md` (Subtask Plan)
- `01-prompts/readme.md` (Category 30 Table Entry update registering prompt 04)
- `01-prompts/30-improvements/readme.md` (Category 30 Group Catalog documenting prompt 04 and review-first gate)
- `.ai-memory/prompts.md` (Operational Memory Prompts Matrix linking prompt 04)
- `.ai-memory/plans/readme.md` (Pending Plans Index linking task 02)
- `02-spec/21-app/readme.md` (Application Domain Specifications Index listing task 02 container)
- Full fleet synchronization of prompts, skills, shared specs (`02-spec/01-*` through `02-spec/20-*`), and additive AI scripts across all 43 connected repositories via `03-ai-scripts/38-sync-prompts-skills-scripts.py`.

---

## 3. Compliance & Boundary Matrix

| Invariant / Check Category | Rule Requirement | Verification Method | Pass Threshold |
| :--- | :--- | :--- | :--- |
| **Review-First Gate** | Prompt 04 and skills must enforce review before any code execution; zero immediate changes. | Spec & prompt inspection | Explicit review gate present |
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
| **Atomic Push Syntax** | Commit via `gitmap cpf "prompts - add detail plans and improvements for review prompt and skills"` (hyphen format, no colons in message). | Commit log check | Clean commit + pushed upstream |

---

## 4. Detailed Action Checklist

### Phase 1: Preflight Verification & Deliverables Audit
- [ ] Verify Task-02 deliverable exists and is non-empty:
  - `01-prompts/30-improvements/04-detail-plans-and-improvements-for-review.md`
- [ ] Verify Task-03 deliverables exist and are symmetrical:
  - `.agents/skills/detail-plans-and-improvements-for-review/skill.md`
  - `.cursor/skills/detail-plans-and-improvements-for-review/skill.md`
- [ ] Verify that Task-01 application specifications exist in `02-spec/21-app/02-detail-plans-and-improvements-for-review/`:
  - `01-architecture-spec.md`
  - `02-component-spec.md`
  - `03-migration-and-verification-spec.md`
- [ ] Verify that Task-01 subtask plans exist in `.ai-memory/plans/subtasks/02-detail-plans-and-improvements-for-review/`:
  - `01-subtask-architecture-and-planning.md`
  - `02-subtask-prompt-and-skills.md`
  - `03-subtask-migration-and-verification.md`
- [ ] Verify review-first and no-immediate-execution invariants are embedded across prompt and skill files.

### Phase 2: Master Catalog Updates & Registration
- [ ] Update `01-prompts/readme.md`:
  - Update Category 30 row in Table of Categories to register prompt `04-detail-plans-and-improvements-for-review.md`.
- [ ] Update `01-prompts/30-improvements/readme.md`:
  - Add prompt 04 row to Directory Index & Catalog table.
  - Document the Review-First & No-Immediate-Execution Gate.
- [ ] Update `.ai-memory/prompts.md`:
  - Add Prompts Matrix entry for prompt `04-detail-plans-and-improvements-for-review.md` under Category `30-improvements`.
- [ ] Update `.ai-memory/plans/readme.md`:
  - Add entry under `## Pending Plans` for `02-detail-plans-and-improvements-for-review.md`.
- [ ] Update `02-spec/21-app/readme.md`:
  - Add entry under `## Active Contents` for `02-detail-plans-and-improvements-for-review/`.

### Phase 3: Linter Verification Suite Execution
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

### Phase 4: Fleet Synchronization Dry-Run & Live Execution
- [ ] Execute dry-run to preview mirrored files and verify 5 boundary exclusions across all 43 repositories:
  ```bash
  python 03-ai-scripts/38-sync-prompts-skills-scripts.py --dry-run
  ```
- [ ] Execute live multi-repository synchronization:
  ```bash
  python 03-ai-scripts/38-sync-prompts-skills-scripts.py
  ```
- [ ] Confirm clean run with 0 errors across all 43 target repositories:
  - Base branch pull succeeded without conflicts (`git pull origin <base_branch> --no-rebase`).
  - Safety backup branch (`backup/sync-<timestamp>`) created and pushed to GitHub for each repository.
  - Canonical assets mirrored cleanly with boundary filters active.
  - Changes committed on base branch via `gitmap cpc` and pushed upstream.

### Phase 5: Downstream Repository Boundary & Leak Verification
- [ ] Verify Spec 21 Boundary: Ensure zero `02-spec/21-*` through `02-spec/25-*` files were copied to any target repository.
- [ ] Verify Bump Script Boundary: Ensure `version.json` and `bump*` scripts remain un-mutated across all target repositories.
- [ ] Verify Memory & Plan Boundary: Ensure target repositories' `.ai-memory/plans/` and `.ai-memory/memory/` remain intact.
- [ ] Verify Archive Boundary: Ensure zero `06-archive/` or `06-old-prompts/` files exist in any target repository.
- [ ] Verify Additive Scripts Boundary: Ensure target repositories maintain existing custom scripts.

### Phase 6: GitMap Cognitive Heatmap Telemetry Verification
- [ ] Run GitMap Cognitive Heatmap status inspection:
  ```bash
  gitmap heatmap status --task "02-detail-plans-and-improvements-for-review"
  ```
- [ ] Run GitMap Cognitive Heatmap tree inspection:
  ```bash
  gitmap heatmap tree --task "02-detail-plans-and-improvements-for-review"
  ```
- [ ] Verify 100% action closure:
  - Every file creation, edit, and audit action marked closed (`met=true`).
  - Zero open, unverified, or dangling actions remaining.

### Phase 7: Final Atomic Semantic Commit & Upstream Push
- [ ] Check working tree status:
  - Verify zero unstaged or uncommitted files (`gitmap status` or inspection).
  - Verify zero temporary files or secrets leaked into repository.
- [ ] Execute atomic commit and push upstream following `AGENTS.md` Section 10 and 13:
  ```bash
  gitmap cpf "prompts - add detail plans and improvements for review prompt and skills"
  ```
- [ ] Verify upstream tracking branch:
  - Confirm `Feature: prompts - add detail plans and improvements for review prompt and skills` is committed.
  - Confirm changes are successfully pushed upstream to GitHub.
  - Confirm working tree is completely clean ('No Push = Not Done' satisfied).

---

## 5. Verification Criteria & Sign-Off Checklist

- [ ] All three application specifications exist in `02-spec/21-app/02-detail-plans-and-improvements-for-review/`.
- [ ] All three subtask plans exist in `.ai-memory/plans/subtasks/02-detail-plans-and-improvements-for-review/`.
- [ ] Zero absolute paths (`C:\`, `/home/`, `file:///`) across all authored specifications and plans.
- [ ] Zero uppercase characters in any authored filenames or directory paths.
- [ ] Canonical prompt `01-prompts/30-improvements/04-detail-plans-and-improvements-for-review.md` verified.
- [ ] Both agent and Cursor skills in `.agents/skills/` and `.cursor/skills/` verified.
- [ ] Review-First & No-Immediate-Execution gate verified in prompt and skills.
- [ ] Master catalogs (`01-prompts/readme.md`, `01-prompts/30-improvements/readme.md`, `.ai-memory/prompts.md`, `.ai-memory/plans/readme.md`, `02-spec/21-app/readme.md`) updated and verified.
- [ ] Linters (`check-relative-paths.py`, `check-forbidden-strings.py`, `check-prompts-loaded.py`) executed and passed cleanly.
- [ ] Fleet synchronization executed across all 43 connected repositories with 0 errors.
- [ ] The 5 Non-Negotiable Boundaries verified across all downstream repositories.
- [ ] Cognitive Heatmap reports 100% action closure (`met=true`).
- [ ] Final atomic commit pushed upstream via `gitmap cpf "prompts - add detail plans and improvements for review prompt and skills"`.
