# Subtask Plan: Directory Restructuring, Prompt Re-Sequencing & Fleet-Wide Synchronization

> **Subtask Plan ID:** `.ai-memory/plans/subtasks/12-multi-repo-sync-engine-and-prompt-upgrades/03-sync-and-verification-subtask.md`  
> **Parent Specification:** `02-spec/21-app/09-multi-repo-sync-engine-and-prompt-upgrades/03-folder-structure-and-sync-spec.md`  
> **Parent Architecture Spec:** `02-spec/21-app/09-multi-repo-sync-engine-and-prompt-upgrades/01-architecture-spec.md`  
> **Status:** READY  
> **Target Subsystems:** `01-prompts/`, `.ai-memory/prompts.md`, `.agents/skills/sync/`, `.cursor/skills/sync/`, `.agents/skills/ai-verification/`, `.cursor/skills/ai-verification/`, `03-ai-scripts/38-sync-prompts-skills-scripts.py`, 43 connected repositories

---

## 1. Overview & Execution Scope

This subtask plan details the deterministic execution steps required to migrate directory structures, resolve prompt prefix collisions, eliminate orphan prompts, synchronize category READMEs, align companion skill pointers, and execute synchronization across all 43 connected target repositories.

Execution is organized into 7 discrete, bounded phases:

```text
┌─────────────────────────────────────────────────────────────────────────────┐
│                       7-PHASE EXECUTION ROADMAP                             │
├─────────────────────────────────────────────────────────────────────────────┤
│ Phase 1: Directory Restructuring & Sequential Renaming                      │
│ Phase 2: Prompt Index Updates & Orphan Prompt Remediation (.ai-memory/)     │
│ Phase 3: Category README Overhaul & Cross-Reference Alignment               │
│ Phase 4: Companion Skills Pointer Synchronization (.agents & .cursor)       │
│ Phase 5: Automated Linters & Pre-Flight Dry-Run Validation                  │
│ Phase 6: Fleet-Wide Synchronization Across 43 Repositories                  │
│ Phase 7: Retrospective AI Verification & Clean Atomic GitMap Commit         │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 2. Phase 1: Directory Restructuring & Sequential Renaming

### Task 1.1: Verify Cursor Prompts Promotion
- Confirm relocation of `01-prompts/22-letterly/cursor/` to `01-prompts/23-cursor-prompts/` (coordinated with Subtask 02).
- Verify all 10 Cursor prompt files are present under `01-prompts/23-cursor-prompts/`:
  - `01-mobile-cursor.md` through `10-execute-with-release-cursor.md`.
- Verify legacy folder `01-prompts/22-letterly/cursor/` is completely removed.

### Task 1.2: Relocate Sync Category (`01-prompts/23-sync/` -> `01-prompts/24-sync/`)
- Create target directory: `01-prompts/24-sync/`.
- Move contents from `01-prompts/23-sync/` to `01-prompts/24-sync/`.
- Remove empty source directory `01-prompts/23-sync/`.

### Task 1.3: Clean Up Prefix Collisions Inside `01-prompts/24-sync/`
- Maintain `01-sync.md` as the primary full-fleet 43-repo synchronization engine:
  - Path: `01-prompts/24-sync/01-sync.md`.
- Rename `01-sync-other-codebase.md` to `02-sync-other-codebase.md`:
  - Path: `01-prompts/24-sync/02-sync-other-codebase.md`.
- Verify no duplicate numerical prefixes exist within `01-prompts/24-sync/`.

### Task 1.4: Relocate AI Verification Category (`01-prompts/24-ai-verification/` -> `01-prompts/25-ai-verification/`)
- Create target directory: `01-prompts/25-ai-verification/`.
- Move contents from `01-prompts/24-ai-verification/` to `01-prompts/25-ai-verification/`:
  - `01-retrospective-ai-verification.md`
  - `readme.md`
- Remove empty source directory `01-prompts/24-ai-verification/`.

### Task 1.5: Validate Directory Hierarchy
- Confirm clean directory sequence in `01-prompts/`:
  - `22-letterly/`
  - `23-cursor-prompts/`
  - `24-sync/`
  - `25-ai-verification/`

---

## 3. Phase 2: Prompt Index Updates & Orphan Prompt Remediation

### Task 2.1: Update Master Index (`.ai-memory/prompts.md`)
Modify `.ai-memory/prompts.md` to catalog all re-sequenced and relocated prompts:
1. Replace `22-letterly/cursor/*` rows with `23-cursor-prompts/*`:
   ```markdown
   | `23-cursor-prompts` | [`23-cursor-prompts/01-mobile-cursor.md`](../01-prompts/23-cursor-prompts/01-mobile-cursor.md) | Mobile Single-Line Mode (Cursor) — Letterly Prompt Formatter |
   | `23-cursor-prompts` | [`23-cursor-prompts/02-desktop-cursor.md`](../01-prompts/23-cursor-prompts/02-desktop-cursor.md) | Desktop Structured Mode (Cursor) — Letterly Prompt Formatter |
   | `23-cursor-prompts` | [`23-cursor-prompts/03-execute-n-steps-cursor.md`](../01-prompts/23-cursor-prompts/03-execute-n-steps-cursor.md) | Execute N-Steps Mode (Cursor) — Letterly Prompt Formatter |
   | `23-cursor-prompts` | [`23-cursor-prompts/04-plan-cursor.md`](../01-prompts/23-cursor-prompts/04-plan-cursor.md) | Planning Spec Mode (Cursor) — Letterly Prompt Formatter |
   | `23-cursor-prompts` | [`23-cursor-prompts/05-release-cursor.md`](../01-prompts/23-cursor-prompts/05-release-cursor.md) | Release Mode (Cursor) — Letterly Prompt Formatter |
   | `23-cursor-prompts` | [`23-cursor-prompts/06-cicd-fix-release-cursor.md`](../01-prompts/23-cursor-prompts/06-cicd-fix-release-cursor.md) | CI/CD Fix & Release Mode (Cursor) — Letterly Prompt Formatter |
   | `23-cursor-prompts` | [`23-cursor-prompts/07-mobile-cicd-fix-cursor.md`](../01-prompts/23-cursor-prompts/07-mobile-cicd-fix-cursor.md) | Mobile CI/CD Fix Mode (Cursor) — Letterly Prompt Formatter |
   | `23-cursor-prompts` | [`23-cursor-prompts/08-run-cursor.md`](../01-prompts/23-cursor-prompts/08-run-cursor.md) | Run Mode (Cursor) — Letterly Prompt Formatter |
   | `23-cursor-prompts` | [`23-cursor-prompts/09-execute-with-verification-cursor.md`](../01-prompts/23-cursor-prompts/09-execute-with-verification-cursor.md) | Execute with Verification Mode (Cursor) — Letterly Prompt Formatter |
   | `23-cursor-prompts` | [`23-cursor-prompts/10-execute-with-release-cursor.md`](../01-prompts/23-cursor-prompts/10-execute-with-release-cursor.md) | Execute with Release Mode (Cursor) — Letterly Prompt Formatter |
   | `23-cursor-prompts` | [`23-cursor-prompts/readme.md`](../01-prompts/23-cursor-prompts/readme.md) | Cursor IDE Prompt Formatters (`23-cursor-prompts`) |
   ```
2. Insert missing `01-sync.md` row and update `24-sync` entries:
   ```markdown
   | `24-sync` | [`24-sync/01-sync.md`](../01-prompts/24-sync/01-sync.md) | [V6] Full-Fleet Multi-Repository Synchronization & Canonical Mirroring Engine — Workflow (must follow) |
   | `24-sync` | [`24-sync/02-sync-other-codebase.md`](../01-prompts/24-sync/02-sync-other-codebase.md) | [V6] Multi-Repository Synchronization & Downstream Codebase Mirroring — Workflow (must follow) |
   | `24-sync` | [`24-sync/readme.md`](../01-prompts/24-sync/readme.md) | Multi-Repository Synchronization Prompts (`24-sync`) — Index & Catalog |
   ```
3. Update `25-ai-verification` entries:
   ```markdown
   | `25-ai-verification` | [`25-ai-verification/01-retrospective-ai-verification.md`](../01-prompts/25-ai-verification/01-retrospective-ai-verification.md) | Retrospective AI Verification & Code Quality Audit — Canonical V6 Workflow (must follow) |
   | `25-ai-verification` | [`25-ai-verification/readme.md`](../01-prompts/25-ai-verification/readme.md) | AI Verification Prompts (`25-ai-verification`) |
   ```

### Task 2.2: Verify Prompt Index Completeness
- Run `python linter-scripts/check-prompts-loaded.py`.
- Verify exit code is 0 with 0 orphan prompts and 0 dangling references detected.

---

## 4. Phase 3: Category README Overhaul & Cross-References

### Task 3.1: Update `01-prompts/readme.md`
- Update directory index code block to show categories 22 through 25:
  ```text
  ├── 22-letterly/
  ├── 23-cursor-prompts/
  ├── 24-sync/
  └── 25-ai-verification/
  ```
- Update Table of Categories:
  - Add `23-cursor-prompts/` row.
  - Update `24-sync/` row with `01-sync.md` and `02-sync-other-codebase.md`.
  - Update `25-ai-verification/` row with `01-retrospective-ai-verification.md`.

### Task 3.2: Update `01-prompts/22-letterly/readme.md`
- Remove references to nested `cursor/` subfolder.
- Add cross-reference section pointing to `../23-cursor-prompts/` for Cursor IDE users.

### Task 3.3: Verify `01-prompts/23-cursor-prompts/readme.md`
- Ensure header reads `# Cursor IDE Prompt Formatters (23-cursor-prompts)`.
- Verify table links to `01-mobile-cursor.md` through `10-execute-with-release-cursor.md`.

### Task 3.4: Update `01-prompts/24-sync/readme.md`
- Update header to `# Multi-Repository Synchronization Prompts (24-sync) — Index & Catalog`.
- Update metadata block to `Category: 24-sync`.
- Update catalog table:
  - `01-sync.md`: Full-Fleet (43 Repos) multi-agent autonomous synchronization.
  - `02-sync-other-codebase.md`: Parameter-Driven (Cross-Repo) dynamic synchronization.

### Task 3.5: Update `01-prompts/25-ai-verification/readme.md`
- Update header to `# AI Verification Prompts (25-ai-verification)`.
- Update metadata block to `Category: 25-ai-verification`.
- Verify prompt link to `01-retrospective-ai-verification.md`.

---

## 5. Phase 4: Companion Skills Pointer Synchronization

### Task 4.1: Update `.agents/skills/sync/skill.md`
- Update canonical prompt metadata:
  ```markdown
  **Canonical Prompt:** `01-prompts/24-sync/01-sync.md`
  ```

### Task 4.2: Update `.cursor/skills/sync/skill.md`
- Update canonical prompt metadata:
  ```markdown
  **Canonical Prompt:** `01-prompts/24-sync/01-sync.md`
  ```

### Task 4.3: Update `.agents/skills/ai-verification/skill.md`
- Add canonical prompt metadata:
  ```markdown
  **Canonical Prompt:** `01-prompts/25-ai-verification/01-retrospective-ai-verification.md`
  ```

### Task 4.4: Update `.cursor/skills/ai-verification/skill.md`
- Add canonical prompt metadata:
  ```markdown
  **Canonical Prompt:** `01-prompts/25-ai-verification/01-retrospective-ai-verification.md`
  ```

---

## 6. Phase 5: Automated Linters & Pre-Flight Dry-Run Validation

### Task 5.1: Run Prompt Index Linter
```bash
python linter-scripts/check-prompts-loaded.py
```
- Assert: Exit code 0, 0 orphans, 0 dangling references.

### Task 5.2: Run Relative Path Checker
```bash
python linter-scripts/check-relative-paths.py
```
- Assert: Zero absolute paths or file protocol URIs.

### Task 5.3: Run Pre-Flight Dry-Run of Fleet Synchronizer
```bash
python 03-ai-scripts/38-sync-prompts-skills-scripts.py --dry-run
```
- Assert: All 43 connected repositories discovered cleanly.
- Assert: 5 Non-Negotiable Boundaries verified:
  - `02-spec/21-*` through `25-*` excluded.
  - `bump*` scripts protected.
  - Additive-only AI scripts guarded against overwriting repo-modified utilities.
  - `.ai-memory/memory/` and `.ai-memory/plans/` excluded.
  - `repo-secrets` excluded.

---

## 7. Phase 6: Fleet Synchronization Across 43 Repositories

### Task 6.1: Execute Live Fleet Synchronization
```bash
python 03-ai-scripts/38-sync-prompts-skills-scripts.py
```
- Execute the 6-stage synchronization ceremony across all 43 repositories:
  1. Pre-flight pull on default base branch.
  2. Create and push safety backup branch (`backup/sync-<timestamp>`).
  3. Return to base branch.
  4. Mirror canonical prompts (`01-prompts/`), skills, shared specs (`02-spec/01-*` through `02-spec/20-*`), and additive AI scripts.
  5. Stage and atomic commit via `gitmap cpf`.
  6. Push commit and create release tags.

### Task 6.2: Validate Target Repository Integrity
- Verify that every target repository in `../` has a clean working tree.
- Verify backup branches exist on remotes (`git branch -r | grep backup/sync-`).
- Verify no `02-spec/21-*` files were created or modified downstream.

---

## 8. Phase 7: Retrospective AI Verification & Clean Atomic Commit

### Task 8.1: Run Retrospective Quality Gate
```bash
python 03-ai-scripts/47-retrospective-ai-verification.py --tasks 3 --since-minutes 34
```
- Verify: Acceptance criteria present in new specs.
- Verify: Relative path hygiene intact.
- Verify: Implicit boolean standards enforced.
- Verify: GitMap pipeline telemetry green via `gitmap pe -t`.

### Task 8.2: Atomic Stage, Commit & Push
```bash
gitmap cpf "sync - folder restructuring, prompt resequencing, and fleet sync"
```

---

## 9. Verification & Validation Commands

| Command | Purpose | Expected Outcome |
| :--- | :--- | :--- |
| `python linter-scripts/check-prompts-loaded.py` | Prompt Index Completeness | Exit code 0; 0 orphan prompts |
| `python linter-scripts/check-relative-paths.py` | Relative Path Hygiene | Zero absolute paths or file protocol URIs |
| `python 03-ai-scripts/38-sync-prompts-skills-scripts.py --dry-run` | Fleet Pre-flight | 43 repositories validated |
| `python 03-ai-scripts/38-sync-prompts-skills-scripts.py` | Fleet Execution | 43/43 repositories synchronized |
| `python 03-ai-scripts/47-retrospective-ai-verification.py` | Quality Gate | 100% checks passed |

---

## 10. Subtask Acceptance Checklist

- [ ] `01-prompts/23-cursor-prompts/` verified and legacy `01-prompts/22-letterly/cursor/` removed.
- [ ] `01-prompts/23-sync/` moved to `01-prompts/24-sync/`.
- [ ] Inside `01-prompts/24-sync/`, prefix collision resolved to `01-sync.md` and `02-sync-other-codebase.md`.
- [ ] `01-prompts/24-ai-verification/` moved to `01-prompts/25-ai-verification/`.
- [ ] `.ai-memory/prompts.md` updated with `01-sync.md` row and updated paths for categories 23, 24, and 25.
- [ ] `python linter-scripts/check-prompts-loaded.py` passes with exit code 0.
- [ ] Category READMEs updated for `01-prompts/readme.md`, `22-letterly`, `23-cursor-prompts`, `24-sync`, `25-ai-verification`.
- [ ] Skills `.agents/skills/sync` and `.cursor/skills/sync` reference `01-prompts/24-sync/01-sync.md`.
- [ ] Skills `.agents/skills/ai-verification` and `.cursor/skills/ai-verification` reference `01-prompts/25-ai-verification/`.
- [ ] All 43 connected repositories pulled, backed up, synchronized, and verified via `38-sync-prompts-skills-scripts.py`.
- [ ] All 5 Non-Negotiable Boundaries verified intact across all target repositories.
- [ ] Clean working tree committed atomically via `gitmap cpf`.
