# Subtask 03: Fleet Migration, Master Catalogs & Quality Verification

## 1. Subtask Overview

| Attribute | Details |
| :--- | :--- |
| **Subtask ID** | `03-subtask-migration-and-verification` |
| **Parent Task** | `06-cg-execute-golang-pointers-and-mutation-reduce` |
| **Parent Plan Ref** | `.ai-memory/plans/pending/06-cg-execute-golang-pointers-and-mutation-reduce.md` |
| **Specification Ref** | `02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce/03-migration-and-verification-spec.md` |
| **Status** | `PENDING` |
| **Assigned Role** | Worker 03 / Lead Orchestrator |
| **Plan Motto** | [/plan](slashCommand;plan) first before doing work for improved work or less friction, faster work and more pure and effective work |

---

## 2. Objective & Scope

Orchestrate the end-to-end master catalog updates, comprehensive linter quality verification gates, fleet-wide multi-repository synchronization across all 43 connected repositories via `03-ai-scripts/38-sync-prompts-skills-scripts.py`, cognitive heatmap telemetry tracking, and the atomic commit push ceremony for the canonical Prompt 39 (`39-golang-pointers-and-mutation-reduction.md`) and its symmetric companion skills in Google Antigravity (`.agents/skills/cg-golang-pointers-and-mutation-reduction/`) and Cursor IDE (`.cursor/skills/cg-golang-pointers-and-mutation-reduction/`).

### Owned Deliverables & Files in Scope

- `02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce/03-migration-and-verification-spec.md` (Specification)
- `.ai-memory/plans/subtasks/06-cg-execute-golang-pointers-and-mutation-reduce/03-subtask-migration-and-verification.md` (Subtask Plan)
- `01-prompts/15-cg-execute/readme.md` (Category 15 Catalog: Add row 39, update tree, and update summary)
- `01-prompts/readme.md` (Master Prompts Catalog: Update Category 15 row to include prompt 39 and tree branch)
- `.ai-memory/prompts.md` (Operational Memory Prompts Matrix: Register prompt 39 with slug `cg-golang-pointers-and-mutation-reduction` and invocation `/golang-pointers-and-mutation-reduction`)
- `.ai-memory/plans/readme.md` (Plans Index: Register pending task `06-cg-execute-golang-pointers-and-mutation-reduce`)
- `02-spec/21-app/readme.md` (Application Specifications Index: Register active container `06-cg-execute-golang-pointers-and-mutation-reduce/`)
- `03-ai-scripts/38-sync-prompts-skills-scripts.py` & `.agents/scripts/38-sync-prompts-skills-scripts.py` (Fleet synchronization execution)
- Full fleet distribution of Prompt 39, companion skills, shared specs (`02-spec/01-*` through `02-spec/20-*`), and additive AI scripts across all 43 connected repositories.

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
| **Prompts Loaded Contract** | Master index `.ai-memory/prompts.md` matches disk prompts directory exactly. | `python linter-scripts/check-prompts-loaded.py` | Exit code 0 (100% on-disk match) |
| **Strict Lowercase Naming** | All filenames and folders adhere to lowercase kebab-case (`^[0-9a-z_.-]+$`). | Path string scan | 0 uppercase characters |
| **Cognitive Heatmap Closure** | 100% action outcome closure across all phases. Zero dangling logs. | `gitmap heatmap status` & `gitmap heatmap tree` | 100% closed (`met=true`) |
| **Atomic Push Syntax** | Commit via `gitmap cpf "prompts - add golang pointers and mutation reduction cg-execute prompt and skills"` (hyphen format, no colons in message). | Commit log check | Clean commit + pushed upstream |

---

## 4. Detailed Action Checklist

### Phase 1: Preflight Verification & Deliverables Audit
- [ ] Verify Wave 1 specifications exist in `02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce/`:
  - `01-architecture-spec.md` (or `01-architecture-and-rules-spec.md`)
  - `02-component-spec.md`
  - `03-migration-and-verification-spec.md`
- [ ] Verify Wave 1 subtask plans exist in `.ai-memory/plans/subtasks/06-cg-execute-golang-pointers-and-mutation-reduce/`:
  - `01-subtask-architecture-and-contracts.md`
  - `02-subtask-component-and-examples.md`
  - `03-subtask-migration-and-verification.md`
- [ ] Verify Wave 2 prompt deliverable exists:
  - `01-prompts/15-cg-execute/39-golang-pointers-and-mutation-reduction.md`
- [ ] Verify Wave 2 companion skills exist symmetrically across runtimes:
  - `.agents/skills/cg-golang-pointers-and-mutation-reduction/skill.md`
  - `.cursor/skills/cg-golang-pointers-and-mutation-reduction/skill.md`

### Phase 2: Master Catalog Updates & Registration
- [ ] Update `01-prompts/15-cg-execute/readme.md`:
  - Add Prompt 39 to Directory Index & Catalog table.
  - Update Directory Tree to include `39-golang-pointers-and-mutation-reduction.md`.
  - Update suite prompt count summary.
- [ ] Update `01-prompts/readme.md`:
  - Update Category 15 row in Table of Categories to reflect `01-execute-coding-guideline-fix.md through 39-*` (all 39 prompts V6 compliant).
  - Update Category 15 Directory Tree branch to include `39-golang-pointers-and-mutation-reduction.md`.
- [ ] Update `.ai-memory/prompts.md`:
  - Register `15-cg-execute/39-golang-pointers-and-mutation-reduction.md` under Category `15-cg-execute`.
  - Add metadata: slug `cg-golang-pointers-and-mutation-reduction`, invocation `/golang-pointers-and-mutation-reduction`.
- [ ] Update `.ai-memory/plans/readme.md`:
  - Add entry under `## Pending Plans` for `06-cg-execute-golang-pointers-and-mutation-reduce.md`.
- [ ] Update `02-spec/21-app/readme.md`:
  - Add entry under `## Active Contents` for `06-cg-execute-golang-pointers-and-mutation-reduce/`.

### Phase 3: Linter Quality Gates
- [ ] Execute `python linter-scripts/check-relative-paths.py` and verify 0 violations.
- [ ] Execute `python linter-scripts/check-forbidden-strings.py` and verify 0 violations.
- [ ] Execute `python linter-scripts/check-prompts-loaded.py` and verify exit code 0 (100% on-disk match).

### Phase 4: Fleet Synchronization Engine Execution
- [ ] Execute preflight dry run of fleet synchronization:
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
  - Target changes committed via structured chore commit (`gitmap cpc`).
  - Upstream push confirmed for tracking branches.
- [ ] Verify downstream boundary invariants across target repositories:
  - 0 files matching `02-spec/21-*` through `25-*` in target repositories.
  - 0 overwritten custom AI scripts in target repositories.
  - 0 mutated `bump*` or `version.json` files in target repositories.
  - 0 mutated target `.ai-memory/plans/` or `.ai-memory/memory/` files.
  - 0 target files in `06-archive/` or `06-old-prompts/`.

### Phase 5: Cognitive Heatmap Telemetry Closure
- [ ] Check cognitive heatmap status:
  ```bash
  gitmap heatmap status --task "06-cg-execute-golang-pointers-and-mutation-reduce"
  ```
- [ ] Close all pending cognitive heatmap log items (`gitmap heatmap update --id <id> --outcome ... --met true`).
- [ ] Audit final cognitive heatmap tree:
  ```bash
  gitmap heatmap tree --task "06-cg-execute-golang-pointers-and-mutation-reduce"
  ```
- [ ] Confirm 100% action closure (`met=true`).

### Phase 6: Release Ceremony & Atomic Push Upstream
- [ ] Verify working tree state (`gitmap status`).
- [ ] Execute atomic commit and push:
  ```bash
  gitmap cpf "prompts - add golang pointers and mutation reduction cg-execute prompt and skills"
  ```
- [ ] Confirm push success on GitHub remote tracking branch (`origin/main` or `origin/master`).

---

## 5. Execution Trace & Audit Log

| Timestamp | Phase | Action / Command | Expected Outcome | Verified Result | Status |
| :--- | :--- | :--- | :--- | :--- | :---: |
| T0 | Preflight | `gitmap lf 02-spec/21-app/06-cg-execute-golang-pointers-and-mutation-reduce` | Inspect specification directory | Directory initialized | `MET` |
| T1 | Spec 03 | Author `03-migration-and-verification-spec.md` | Full migration, sync & verification spec written | File written with zero relative path violations | `MET` |
| T2 | Subtask 03 | Author `03-subtask-migration-and-verification.md` | Canonical subtask plan established | File written following strict subtask schema | `MET` |
| T3 | Catalogs | Update master catalogs (01-prompts, 15-cg-execute, .ai-memory, 21-app) | All 5 catalogs registered with Prompt 39 | Pending worker execution | `PENDING` |
| T4 | Linters | Run `check-relative-paths.py`, `check-forbidden-strings.py`, `check-prompts-loaded.py` | 0 relative path errors, 0 forbidden strings, exact on-disk parity | Pending worker execution | `PENDING` |
| T5 | Fleet Sync | Run `38-sync-prompts-skills-scripts.py` across 43 repositories | All 43 repos synced, safety branches pushed, boundaries preserved | Pending lead execution | `PENDING` |
| T6 | Heatmap | Verify heatmap closure (`gitmap heatmap status`) | 100% action items closed (`met=true`) | Pending lead execution | `PENDING` |
| T7 | Commit | `gitmap cpf "prompts - add golang pointers and mutation reduction cg-execute prompt and skills"` | Atomic commit and push upstream | Pending lead execution | `PENDING` |

---

## 6. Verification Criteria & Definition of Done

This subtask reaches completion when all of the following binary conditions evaluate to `true`:

1. **Deliverable Completeness:** Both `03-migration-and-verification-spec.md` and `03-subtask-migration-and-verification.md` exist on disk, are non-empty, and contain zero absolute filesystem paths or `file:///` URIs.
2. **Catalog Synchronization:** All 5 master catalog files (`01-prompts/15-cg-execute/readme.md`, `01-prompts/readme.md`, `.ai-memory/prompts.md`, `.ai-memory/plans/readme.md`, `02-spec/21-app/readme.md`) register Prompt 39 and Task `06-cg-execute-golang-pointers-and-mutation-reduce`.
3. **Linter Parity:** `python linter-scripts/check-prompts-loaded.py` confirms exact parity on disk with 0 orphan prompts and 0 dangling references.
4. **Boundary Integrity:** 0 files from `02-spec/21-*` through `25-*` leaked into any of the 43 downstream repositories.
5. **Telemetry Closure:** GitMap Cognitive Heatmap shows 100% of action items marked `met=true` with 0 unclosed records.
6. **Upstream Invariant:** Changes committed via `gitmap cpf "prompts - add golang pointers and mutation reduction cg-execute prompt and skills"` and pushed cleanly to the remote tracking branch.
