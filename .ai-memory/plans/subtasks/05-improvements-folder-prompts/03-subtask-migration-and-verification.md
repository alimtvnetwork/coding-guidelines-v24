# Subtask 03: Fleet Migration, Master Catalogs & Quality Verification

## 1. Subtask Overview

| Attribute | Details |
| :--- | :--- |
| **Subtask ID** | `03-subtask-migration-and-verification` |
| **Parent Task** | `05-improvements-folder-prompts` |
| **Parent Plan Ref** | `.ai-memory/plans/pending/05-improvements-folder-prompts.md` |
| **Specification Ref** | `02-spec/21-app/05-improvements-folder-prompts/03-migration-and-verification-spec.md` |
| **Status** | `PENDING` |
| **Assigned Role** | Worker 03 / Lead Orchestrator |
| **Plan Motto** | Improved work or less friction, faster work and more pure and effective work |

---

## 2. Objective & Scope

Orchestrate the end-to-end master catalog updates, comprehensive linter quality gates, V6 parent task synchronization, fleet-wide multi-repository synchronization across all 43 connected repositories via `03-ai-scripts/38-sync-prompts-skills-scripts.py`, cognitive heatmap telemetry tracking, and the atomic commit push ceremony for the expanded Category `30-improvements` suite (`06-check-pending-improvements.md` and `07-implement-next-10-improvements.md`) and their symmetric companion skills in Google Antigravity (`.agents/skills/`) and Cursor IDE (`.cursor/skills/`).

### Owned Deliverables & Files in Scope

- `02-spec/21-app/05-improvements-folder-prompts/03-migration-and-verification-spec.md` (Specification)
- `.ai-memory/plans/subtasks/05-improvements-folder-prompts/03-subtask-migration-and-verification.md` (Subtask Plan)
- `01-prompts/readme.md` (Category 30 row update registering Prompts 06 and 07 and expanded 7-prompt scope)
- `01-prompts/30-improvements/readme.md` (Category 30 Group Catalog documenting all 7 prompts, directory tree, and skills)
- `.ai-memory/prompts.md` (Operational Memory Prompts Matrix registering Prompts 06 and 07, slugs, and slash commands)
- `.ai-memory/plans/readme.md` (Pending Plans Index linking task 05)
- `02-spec/21-app/readme.md` (Application Domain Specifications Index listing task 05 container)
- `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` & companion skills in `.agents/skills/execute-parent-task-with-n-steps-v6/` and `.cursor/skills/execute-parent-task-with-n-steps-v6/` (Fast-forwarding and coding guideline checkpoints integration)
- `03-ai-scripts/38-sync-prompts-skills-scripts.py` & `.agents/scripts/38-sync-prompts-skills-scripts.py` (Fleet synchronization execution)
- Full fleet distribution of prompts, skills, shared specs (`02-spec/01-*` through `02-spec/20-*`), and additive AI scripts across all 43 connected repositories.

---

## 3. Compliance & Boundary Matrix

| Invariant / Check Category | Rule Requirement | Verification Method | Pass Threshold |
| :--- | :--- | :--- | :--- |
| **Spec 21 Exclusion** | Zero leakage of `02-spec/21-*` through `02-spec/25-*` into target repositories. | Fleet path scan | 0 leaked files |
| **Additive Scripts** | Add new scripts; never overwrite scripts modified by target repositories. | Engine filter verification | 0 overwritten target scripts |
| **Bump Script Protection** | Version bump scripts (`bump*`, `version.json`) remain untouched. | File hash check | 0 modified bump scripts |
| **Memory & Plan Isolation** | `.ai-memory/plans/` and `.ai-memory/memory/` remain strictly isolated per repository. | Directory scan | 0 overwritten plan files |
| **Archive Exclusion** | `06-archive/` and `06-old-prompts/` are never synchronized downstream. | Target repos scan | 0 archive files in fleet |
| **Strict Relative Paths** | All links and file references use strictly relative Git paths (`02-spec/...`, `.ai-memory/...`). TOTAL BAN on `file:///` and absolute paths. | `python linter-scripts/check-relative-paths.py` | 0 absolute path references |
| **Forbidden Strings** | Zero forbidden patterns, unescaped tokens, or banned template syntax. | `python linter-scripts/check-forbidden-strings.py` | 0 violations |
| **Prompts Loaded Contract** | Master index `.ai-memory/prompts.md` matches disk prompts directory exactly (220 prompts). | `python linter-scripts/check-prompts-loaded.py` | Exit code 0 (100% on-disk match) |
| **Strict Lowercase Naming** | All filenames and folders adhere to lowercase kebab-case (`^[0-9a-z_.-]+$`). | Path string scan | 0 uppercase characters |
| **Cognitive Heatmap Closure** | 100% action outcome closure across all phases. Zero dangling logs. | `gitmap heatmap status` & `gitmap heatmap tree` | 100% closed (`met=true`) |
| **Atomic Push Syntax** | Commit via `gitmap cpf "prompts - add check pending and implement next 10 improvements prompts and skills"` (hyphen format, no colons in message). | Commit log check | Clean commit + pushed upstream |

---

## 4. Detailed Action Checklist

### Phase 1: Preflight Verification & Deliverables Audit
- [ ] Verify Task-02 prompt deliverables exist and are populated with complete instructions:
  - `01-prompts/30-improvements/06-check-pending-improvements.md` exists.
  - `01-prompts/30-improvements/07-implement-next-10-improvements.md` exists.
- [ ] Verify Task-03 skill deliverables exist symmetrically across runtimes:
  - `.agents/skills/check-pending-improvements/skill.md` exists.
  - `.cursor/skills/check-pending-improvements/skill.md` exists.
  - `.agents/skills/implement-next-10-improvements/skill.md` exists.
  - `.cursor/skills/implement-next-10-improvements/skill.md` exists.
- [ ] Verify that Task-01 application specifications exist in `02-spec/21-app/05-improvements-folder-prompts/`:
  - `01-architecture-spec.md`
  - `02-component-spec.md`
  - `03-migration-and-verification-spec.md`
- [ ] Verify that Task-01 subtask plans exist in `.ai-memory/plans/subtasks/05-improvements-folder-prompts/`:
  - `01-subtask-architecture-and-contracts.md`
  - `02-subtask-component-and-examples.md`
  - `03-subtask-migration-and-verification.md`

### Phase 2: Master Catalog Updates & Registration
- [ ] Update `01-prompts/30-improvements/readme.md`:
  - Add Prompts 06 and 07 to Directory Index & Catalog table.
  - Update Directory Tree to show all 7 prompt files (`01` through `07`).
  - Add companion skill mapping entries for Antigravity and Cursor runtimes.
- [ ] Update `01-prompts/readme.md`:
  - Update Category 30 row in Table of Categories to list all 7 prompt files and expanded capability summary.
  - Update Category 30 Directory Tree branch to include `06-check-pending-improvements.md` and `07-implement-next-10-improvements.md`.
- [ ] Update `.ai-memory/prompts.md`:
  - Register `06-check-pending-improvements.md` and `07-implement-next-10-improvements.md` under Category `30-improvements`.
  - Add metadata: slugs `check-pending-improvements` and `implement-next-10-improvements`, slash commands `/check-pending-improvements` and `/implement-next-10-improvements`.
- [ ] Update `.ai-memory/plans/readme.md`:
  - Add entry under `## Pending Plans` for `05-improvements-folder-prompts.md`.
- [ ] Update `02-spec/21-app/readme.md`:
  - Add entry under `## Active Contents` for `05-improvements-folder-prompts/`.

### Phase 3: V6 Parent Task Enhancement & Symmetrical Skill Update
- [ ] Update `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`:
  - Embed Checkpoint 1: Fast-forward implementation protocol (`NO-BUILD` / `NO-TEST` during code modification turns).
  - Embed Checkpoint 2: Strict coding guidelines (`*appfault.AppError`, `Result[T]`, parameter structs, string constants/enums).
  - Embed Checkpoint 3: Function and file sizing limits (8–15 line functions, 150–300 line files).
  - Embed Checkpoint 4: Dedicated post-implementation separated end-to-end testing and building verification gate.
  - Embed Checkpoint 5: Repeatable "next" execution cadence.
- [ ] Update `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` symmetrically.
- [ ] Update `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md` symmetrically.

### Phase 4: Linter Quality Gates
- [ ] Execute `python linter-scripts/check-relative-paths.py` and verify 0 violations.
- [ ] Execute `python linter-scripts/check-forbidden-strings.py` and verify 0 violations.
- [ ] Execute `python linter-scripts/check-prompts-loaded.py` and verify exit code 0 (220 prompts indexed on disk).

### Phase 5: Fleet Synchronization Engine Execution
- [ ] Dry-run synchronization preflight:
  ```bash
  python 03-ai-scripts/38-sync-prompts-skills-scripts.py --dry-run
  ```
- [ ] Execute live fleet synchronization across all 43 connected repositories:
  ```bash
  python 03-ai-scripts/38-sync-prompts-skills-scripts.py
  ```
- [ ] Verify safety protocol:
  - Base branch pulled cleanly (`git pull origin <base_branch> --no-rebase`).
  - Safety backup branch created and pushed (`backup/sync-<timestamp>`).
  - Changes committed via structured chore commit (`gitmap cpc`).
  - Upstream push confirmed for tracking branches.
- [ ] Verify downstream boundary invariants across target repositories:
  - 0 files matching `02-spec/21-*` through `25-*` in target repositories.
  - 0 overwritten custom AI scripts in target repositories.
  - 0 mutated `bump*` or `version.json` files in target repositories.
  - 0 mutated target `.ai-memory/plans/` or `.ai-memory/memory/` files.
  - 0 target files in `06-archive/` or `06-old-prompts/`.

### Phase 6: Cognitive Heatmap Telemetry Closure
- [ ] Check cognitive heatmap status:
  ```bash
  gitmap heatmap status --task "05-improvements-folder-prompts"
  ```
- [ ] Close all pending cognitive heatmap log items (`gitmap heatmap update --id <id> --outcome ... --met true`).
- [ ] Audit final cognitive heatmap tree:
  ```bash
  gitmap heatmap tree --task "05-improvements-folder-prompts"
  ```
- [ ] Confirm 100% action closure (`met=true`).

### Phase 7: Release Ceremony & Atomic Push Upstream
- [ ] Verify working tree state (`gitmap status`).
- [ ] Execute atomic commit and push:
  ```bash
  gitmap cpf "prompts - add check pending and implement next 10 improvements prompts and skills"
  ```
- [ ] Confirm push success on GitHub remote tracking branch (`origin/main` or `origin/master`).

---

## 5. Execution Trace & Audit Log

| Timestamp | Phase | Action / Command | Expected Outcome | Verified Result | Status |
| :--- | :--- | :--- | :--- | :--- | :---: |
| T0 | Preflight | `gitmap lf 02-spec/21-app/05-improvements-folder-prompts` | Inspect specification directory | Directory initialized | `MET` |
| T1 | Spec 03 | Author `03-migration-and-verification-spec.md` | Full migration, sync & verification spec written | File written with zero relative path violations | `MET` |
| T2 | Subtask 03 | Author `03-subtask-migration-and-verification.md` | Canonical subtask plan established | File written following strict subtask schema | `MET` |
| T3 | Catalogs | Update master catalogs (01-prompts, 30-improvements, .ai-memory, 21-app) | All 5 catalogs registered with Prompts 06 and 07 | Pending worker execution | `PENDING` |
| T4 | V6 Parent Sync | Enhance `02-execute-parent-task-with-n-steps-v6.md` & skills | Fast-forwarding and coding guideline checkpoints embedded | Pending worker execution | `PENDING` |
| T5 | Linters | Run `check-relative-paths.py`, `check-forbidden-strings.py`, `check-prompts-loaded.py` | 0 relative path errors, 0 forbidden strings, 220 prompts match | Pending worker execution | `PENDING` |
| T6 | Fleet Sync | Run `38-sync-prompts-skills-scripts.py` across 43 repositories | All 43 repos synced, safety branches pushed, boundaries preserved | Pending lead execution | `PENDING` |
| T7 | Heatmap | Verify heatmap closure (`gitmap heatmap status`) | 100% action items closed (`met=true`) | Pending lead execution | `PENDING` |
| T8 | Commit | `gitmap cpf "prompts - add check pending and implement next 10 improvements prompts and skills"` | Atomic commit and push upstream | Pending lead execution | `PENDING` |

---

## 6. Verification Criteria & Definition of Done

This subtask reaches completion when all of the following binary conditions evaluate to `true`:

1. **Deliverable Completeness:** Both `03-migration-and-verification-spec.md` and `03-subtask-migration-and-verification.md` exist on disk, are non-empty, and contain zero absolute filesystem paths or `file:///` URIs.
2. **Catalog Synchronization:** All 5 master catalog files (`01-prompts/30-improvements/readme.md`, `01-prompts/readme.md`, `.ai-memory/prompts.md`, `.ai-memory/plans/readme.md`, `02-spec/21-app/readme.md`) register Prompts 06 and 07, task 05, and the full 7-prompt suite.
3. **V6 Cross-Prompt Integration:** `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` and its companion skills embed the fast-forwarding checkpoints (`NO-BUILD`/`NO-TEST`), coding guidelines, sizing bounds, and dedicated end-to-end testing gate.
4. **Linter Parity:** `python linter-scripts/check-prompts-loaded.py` confirms exactly 220 prompt files indexed on disk with 0 orphan prompts and 0 dangling references.
5. **Boundary Integrity:** 0 files from `02-spec/21-*` through `25-*` leaked into any of the 43 downstream repositories.
6. **Telemetry Closure:** GitMap Cognitive Heatmap shows 100% of action items marked `met=true` with 0 unclosed records.
7. **Upstream Invariant:** Changes committed via `gitmap cpf "prompts - add check pending and implement next 10 improvements prompts and skills"` and pushed cleanly to the remote tracking branch.
