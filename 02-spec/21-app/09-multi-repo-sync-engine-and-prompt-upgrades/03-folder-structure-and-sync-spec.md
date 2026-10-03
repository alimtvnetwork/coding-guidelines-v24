# Architecture Specification: Directory Structure Migration, Prompt Re-Sequencing & Fleet-Wide Synchronization

> **Spec ID:** `02-spec/21-app/09-multi-repo-sync-engine-and-prompt-upgrades/03-folder-structure-and-sync-spec.md`  
> **Parent Architecture Spec:** `02-spec/21-app/09-multi-repo-sync-engine-and-prompt-upgrades/01-architecture-spec.md`  
> **Companion Spec:** `02-spec/21-app/09-multi-repo-sync-engine-and-prompt-upgrades/02-letterly-and-cursor-spec.md`  
> **Status:** APPROVED  
> **Target Subsystems:** `01-prompts/`, `.ai-memory/prompts.md`, `.agents/skills/sync/`, `.cursor/skills/sync/`, `.agents/skills/ai-verification/`, `.cursor/skills/ai-verification/`, `03-ai-scripts/38-sync-prompts-skills-scripts.py`, and 43 connected repositories

---

## 1. Executive Summary & Architectural Motivation

As the Prompt Architect repository evolves, its canonical prompt library (`01-prompts/`) and multi-repository synchronization engine require architectural restructuring to maintain strict numerical ordering, zero orphan prompts, symmetric IDE skill integration, and deterministic fleet-wide synchronization.

This specification governs five major architectural interventions:

1. **Top-Level Category Promotion & Re-Sequencing:**
   - The Cursor prompt suite previously resided in a nested subfolder `01-prompts/22-letterly/cursor/`. This created structural asymmetry where Cursor IDE templates were treated as subordinate to Letterly rather than a first-class prompt category. Elevating Cursor prompts to `01-prompts/23-cursor-prompts/` requires shifting downstream numerical categories:
     - `01-prompts/22-letterly/cursor/` -> `01-prompts/23-cursor-prompts/`
     - `01-prompts/23-sync/` -> `01-prompts/24-sync/`
     - `01-prompts/24-ai-verification/` -> `01-prompts/25-ai-verification/`

2. **Sequence Prefix Collision Remediation in Sync Category:**
   - Inside `01-prompts/23-sync/`, two prompt files were created sharing the identical numerical prefix `01-`:
     - `01-sync.md` (Full-Fleet 43-Repository Synchronization Engine)
     - `01-sync-other-codebase.md` (Parameter-Driven Downstream Codebase Mirroring)
   - In the re-sequenced `01-prompts/24-sync/` directory, this collision is resolved deterministically:
     - `01-sync.md` retains prefix `01-` as the default full-fleet synchronization prompt.
     - `01-sync-other-codebase.md` is re-sequenced to `02-sync-other-codebase.md` as the secondary parameter-driven synchronization prompt.

3. **Orphan Prompt Index Remediation (`.ai-memory/prompts.md`):**
   - The master prompt catalog `.ai-memory/prompts.md` serves as the on-disk index validated by `linter-scripts/check-prompts-loaded.py`.
   - `01-sync.md` was authored without a corresponding row in `.ai-memory/prompts.md`, causing CI linter failures due to orphan prompt detection.
   - All relocated categories (`23-cursor-prompts/`, `24-sync/`, `25-ai-verification/`) must be updated in `.ai-memory/prompts.md` with complete catalog rows, eliminating all orphan prompts and broken references.

4. **Category Documentation & Skill Pointer Alignment:**
   - Category index README files across all affected folders (`01-prompts/readme.md`, `01-prompts/22-letterly/readme.md`, `01-prompts/23-cursor-prompts/readme.md`, `01-prompts/24-sync/readme.md`, `01-prompts/25-ai-verification/readme.md`) must be synchronized to reflect the new category numbers, file names, and companion skills.
   - Skill definitions (`.agents/skills/sync/skill.md`, `.cursor/skills/sync/skill.md`, `.agents/skills/ai-verification/skill.md`, `.cursor/skills/ai-verification/skill.md`) must update their `Canonical Prompt` metadata to reference the new paths in `01-prompts/24-sync/` and `01-prompts/25-ai-verification/`.

5. **Fleet-Wide Synchronization Execution Rules (43 Repositories):**
   - Full distribution across all 43 connected repositories using `03-ai-scripts/38-sync-prompts-skills-scripts.py`.
   - Enforcement of the 5 Non-Negotiable Boundaries: Spec 21 Exclusion, Additive-Only AI Scripts with repo-modification checks, Bump Script Protection, Memory & Plans Protection, and Zero Secrets Leakage.

---

## 2. Directory Restructuring & Migration Topology

### 2.1 Before vs. After Directory Topology

```text
BEFORE (Asymmetric & Prefix Collisions):
01-prompts/
├── 22-letterly/
│   ├── 01-mobile-letterly.md ... 10-execute-with-release-letterly.md
│   ├── cursor/                          <-- Nested subordinate folder
│   │   ├── 01-mobile-letterly-cursor.md ... 10-execute-with-release-letterly-cursor.md
│   │   └── readme.md
│   └── readme.md
├── 23-sync/
│   ├── 01-sync.md                       <-- Collision: duplicate "01-" prefix
│   ├── 01-sync-other-codebase.md        <-- Collision: duplicate "01-" prefix
│   └── readme.md
└── 24-ai-verification/
    ├── 01-retrospective-ai-verification.md
    └── readme.md

AFTER (Normalized, Resequenced & Linear):
01-prompts/
├── 22-letterly/                         <-- Letterly Voice Formatters (Agent)
│   ├── 01-mobile-letterly.md ... 10-execute-with-release-letterly.md
│   └── readme.md
├── 23-cursor-prompts/                   <-- First-Class Cursor IDE Formatters
│   ├── 01-mobile-cursor.md ... 10-execute-with-release-cursor.md
│   └── readme.md
├── 24-sync/                             <-- Re-sequenced Synchronization Suite
│   ├── 01-sync.md                       <-- Primary Full-Fleet 43-Repo Engine
│   ├── 02-sync-other-codebase.md        <-- Clean "02-" Prefix (Parameter-Driven)
│   └── readme.md
└── 25-ai-verification/                  <-- Re-sequenced Verification Suite
    ├── 01-retrospective-ai-verification.md
    └── readme.md
```

### 2.2 File Relocation & Migration Matrix

| Source Path (Old) | Target Path (New) | Transformation Action | Rationale |
| :--- | :--- | :--- | :--- |
| `01-prompts/22-letterly/cursor/01-mobile-letterly-cursor.md` | `01-prompts/23-cursor-prompts/01-mobile-cursor.md` | Move & Rename | Promote to top-level category; simplify file slug |
| `01-prompts/22-letterly/cursor/02-desktop-letterly-cursor.md` | `01-prompts/23-cursor-prompts/02-desktop-cursor.md` | Move & Rename | Promote to top-level category; simplify file slug |
| `01-prompts/22-letterly/cursor/03-execute-n-steps-letterly-cursor.md` | `01-prompts/23-cursor-prompts/03-execute-n-steps-cursor.md` | Move & Rename | Promote to top-level category; simplify file slug |
| `01-prompts/22-letterly/cursor/04-plan-letterly-cursor.md` | `01-prompts/23-cursor-prompts/04-plan-cursor.md` | Move & Rename | Promote to top-level category; simplify file slug |
| `01-prompts/22-letterly/cursor/05-release-letterly-cursor.md` | `01-prompts/23-cursor-prompts/05-release-cursor.md` | Move & Rename | Promote to top-level category; simplify file slug |
| `01-prompts/22-letterly/cursor/06-cicd-fix-release-letterly-cursor.md` | `01-prompts/23-cursor-prompts/06-cicd-fix-release-cursor.md` | Move & Rename | Promote to top-level category; simplify file slug |
| `01-prompts/22-letterly/cursor/07-mobile-cicd-fix-letterly-cursor.md` | `01-prompts/23-cursor-prompts/07-mobile-cicd-fix-cursor.md` | Move & Rename | Promote to top-level category; simplify file slug |
| `01-prompts/22-letterly/cursor/08-run-letterly-cursor.md` | `01-prompts/23-cursor-prompts/08-run-cursor.md` | Move & Rename | Promote to top-level category; simplify file slug |
| `01-prompts/22-letterly/cursor/09-execute-with-verification-letterly-cursor.md` | `01-prompts/23-cursor-prompts/09-execute-with-verification-cursor.md` | Move & Rename | Promote to top-level category; simplify file slug |
| `01-prompts/22-letterly/cursor/10-execute-with-release-letterly-cursor.md` | `01-prompts/23-cursor-prompts/10-execute-with-release-cursor.md` | Move & Rename | Promote to top-level category; simplify file slug |
| `01-prompts/22-letterly/cursor/readme.md` | `01-prompts/23-cursor-prompts/readme.md` | Move & Rewrite Header | Update category identifier to `23-cursor-prompts` |
| `01-prompts/23-sync/01-sync.md` | `01-prompts/24-sync/01-sync.md` | Directory Shift | Accommodate `23-cursor-prompts/`; maintain `01-` prefix |
| `01-prompts/23-sync/01-sync-other-codebase.md` | `01-prompts/24-sync/02-sync-other-codebase.md` | Directory Shift & Renumber | Resolve prefix collision; establish secondary sequence |
| `01-prompts/23-sync/readme.md` | `01-prompts/24-sync/readme.md` | Directory Shift & Rewrite Header | Update category identifier to `24-sync` |
| `01-prompts/24-ai-verification/01-retrospective-ai-verification.md` | `01-prompts/25-ai-verification/01-retrospective-ai-verification.md` | Directory Shift | Accommodate shifted `24-sync/` category |
| `01-prompts/24-ai-verification/readme.md` | `01-prompts/25-ai-verification/readme.md` | Directory Shift & Rewrite Header | Update category identifier to `25-ai-verification` |

---

## 3. Orphan Prompt Resolution & Master Index (`.ai-memory/prompts.md`)

### 3.1 Root Cause of Linter Failure
The prompt index verifier `linter-scripts/check-prompts-loaded.py` scans all `.md` files residing under `01-prompts/` and compares them against the table rows in `.ai-memory/prompts.md`. Any on-disk prompt absent from the table triggers an orphan prompt error:

```text
🔎 Prompts on disk: 196
❌ Orphan prompts (not referenced by the index) (1):
   • 23-sync/01-sync.md
```

### 3.2 Index Remediation Contract
To satisfy `linter-scripts/check-prompts-loaded.py` and achieve 100% prompt discoverability, the trailing category rows of `.ai-memory/prompts.md` must be updated to the following canonical entries:

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
| `22-letterly` | [`22-letterly/readme.md`](../01-prompts/22-letterly/readme.md) | Letterly Prompt Formatters (`22-letterly`) |
| `24-sync` | [`24-sync/01-sync.md`](../01-prompts/24-sync/01-sync.md) | [V6] Full-Fleet Multi-Repository Synchronization & Canonical Mirroring Engine — Workflow (must follow) |
| `24-sync` | [`24-sync/02-sync-other-codebase.md`](../01-prompts/24-sync/02-sync-other-codebase.md) | [V6] Multi-Repository Synchronization & Downstream Codebase Mirroring — Workflow (must follow) |
| `24-sync` | [`24-sync/readme.md`](../01-prompts/24-sync/readme.md) | Multi-Repository Synchronization Prompts (`24-sync`) — Index & Catalog |
| `25-ai-verification` | [`25-ai-verification/01-retrospective-ai-verification.md`](../01-prompts/25-ai-verification/01-retrospective-ai-verification.md) | Retrospective AI Verification & Code Quality Audit — Canonical V6 Workflow (must follow) |
| `25-ai-verification` | [`25-ai-verification/readme.md`](../01-prompts/25-ai-verification/readme.md) | AI Verification Prompts (`25-ai-verification`) |
| `.` | [`readme.md`](../01-prompts/readme.md) | Prompt Architect: Canonical AI Prompts Library |
```

---

## 4. Category Readme Synchronization Specifications

### 4.1 Master Index: `01-prompts/readme.md`
- **Total Categories:** Expanded from 25 to 26 canonical categories (`00` through `25`).
- **Directory Tree:** Must display `22-letterly/`, `23-cursor-prompts/`, `24-sync/`, and `25-ai-verification/`.
- **Table of Categories Update:**
  - Category `22`: `22-letterly/` — `01-mobile-letterly.md` through `10-execute-with-release-letterly.md`
  - Category `23`: `23-cursor-prompts/` — `01-mobile-cursor.md` through `10-execute-with-release-cursor.md`
  - Category `24`: `24-sync/` — `01-sync.md`, `02-sync-other-codebase.md`
  - Category `25`: `25-ai-verification/` — `01-retrospective-ai-verification.md`

### 4.2 Letterly Formatters Index: `01-prompts/22-letterly/readme.md`
- Remove the subordinate `cursor/` subfolder reference.
- Add an explicit cross-reference to `../23-cursor-prompts/` for Cursor IDE environments.
- Maintain the table of 10 Antigravity letterly formatters.

### 4.3 Cursor Formatters Index: `01-prompts/23-cursor-prompts/readme.md`
- Header: `# Cursor IDE Prompt Formatters (23-cursor-prompts)`.
- Index table documenting `01-mobile-cursor.md` through `10-execute-with-release-cursor.md`.
- Explicitly map companion skills to `.cursor/skills/letterly-*/`.

### 4.4 Synchronization Suite Index: `01-prompts/24-sync/readme.md`
- Header: `# Multi-Repository Synchronization Prompts (24-sync) — Index & Catalog`.
- Metadata block: `Category: 24-sync`.
- Index table:
  - `01-sync.md`: Full-Fleet (43 Repos) multi-agent autonomous synchronization.
  - `02-sync-other-codebase.md`: Parameter-Driven (Cross-Repo) dynamic synchronization.

### 4.5 AI Verification Suite Index: `01-prompts/25-ai-verification/readme.md`
- Header: `# AI Verification Prompts (25-ai-verification)`.
- Metadata block: `Category: 25-ai-verification`.
- Index table documenting `01-retrospective-ai-verification.md` and companion skill `ai-verification`.

---

## 5. Companion Skills Link Alignment

Companion skills in `.agents/skills/` and `.cursor/skills/` maintain metadata linking to their canonical prompts. These links must be updated to reference relocated paths:

| Skill File | Field to Update | Old Value | New Value |
| :--- | :--- | :--- | :--- |
| `.agents/skills/sync/skill.md` | `**Canonical Prompt:**` | `01-prompts/23-sync/01-sync.md` | `01-prompts/24-sync/01-sync.md` |
| `.cursor/skills/sync/skill.md` | `**Canonical Prompt:**` | `01-prompts/23-sync/01-sync.md` | `01-prompts/24-sync/01-sync.md` |
| `.agents/skills/ai-verification/skill.md` | Header Metadata | *Implicit / Unlinked* | `**Canonical Prompt:** `01-prompts/25-ai-verification/01-retrospective-ai-verification.md`` |
| `.cursor/skills/ai-verification/skill.md` | Header Metadata | *Implicit / Unlinked* | `**Canonical Prompt:** `01-prompts/25-ai-verification/01-retrospective-ai-verification.md`` |

---

## 6. Fleet-Wide Synchronization Engine & Non-Negotiable Boundaries

The synchronization engine `03-ai-scripts/38-sync-prompts-skills-scripts.py` synchronizes canonical prompts, skills, shared specs (`02-spec/01-*` through `02-spec/20-*`), and additive AI scripts across all 43 connected repositories.

### 6.1 The 5 Non-Negotiable Boundaries

```text
┌─────────────────────────────────────────────────────────────────────────────┐
│                    5 NON-NEGOTIABLE SYNCHRONIZATION BOUNDARIES              │
├─────────────────────────────────────────────────────────────────────────────┤
│ 1. Spec 21 Exclusion      │ NEVER touch 02-spec/21-* through 25-*.         │
│ 2. Additive-Only Scripts  │ Add missing scripts; diff/preserve modified.   │
│ 3. Bump Script Guard      │ NEVER overwrite target repo bump scripts.       │
│ 4. Memory & Plans Safe    │ NEVER touch .ai-memory/memory/ or plans/.       │
│ 5. Zero Secrets Mandate   │ NEVER copy credentials, .env, or private keys.  │
└─────────────────────────────────────────────────────────────────────────────┘
```

1. **Spec 21 Exclusion (`02-spec/21-*` through `02-spec/25-*`):**
   - Application domain specifications, client data models, and local audit matrices belong exclusively to target repositories.
   - Synchronizing or touching these folders in downstream repos is strictly forbidden.

2. **Additive-Only AI Scripts (`03-ai-scripts/`, `.agents/scripts/`):**
   - New utility scripts absent from target repositories are safely copied.
   - For existing scripts present in target repositories, inspect git log history: if the script was modified by the child repository (commit message does not contain `sync`), preserve the local modifications and do not overwrite.

3. **Bump Script Protection (`bump*`):**
   - Version bumping scripts (`bump-version.mjs`, `bump_versions.py`, `37-bump-version.py`, `version.json`) contain repository-specific hooks, targets, and tag patterns. They must never be overwritten.

4. **Memory & Plans Protection (`.ai-memory/memory/`, `.ai-memory/plans/`):**
   - Target repositories maintain private execution memory, sprint logs, and active plans. These directories must remain 100% isolated.

5. **Zero Secrets Mandate (`repo-secrets` / `rs`):**
   - API tokens, environment files (`.env*`), and credentials must never be synchronized across repositories.

### 6.2 The 6-Stage Synchronization Ceremony

For each of the 43 target repositories, the synchronization script executes the following deterministic ceremony:

1. **Pre-Flight Pull:**
   - Detect repository default base branch (`main`, `master`, or active branch).
   - Fetch and pull upstream changes: `git pull origin <base_branch> --no-rebase`.
2. **Safety Backup Branch:**
   - Create and immediately push a safety backup branch before touching any files:
     `git checkout -b backup/sync-<timestamp>` -> `git push origin backup/sync-<timestamp>`.
3. **Return to Base Branch:**
   - Switch back to the working base branch: `git checkout <base_branch>`.
4. **Controlled Asset Mirroring:**
   - Mirror `01-prompts/` (excluding `06-archive/`).
   - Mirror `.agents/skills/` and `.cursor/skills/`.
   - Additive-only mirror of `03-ai-scripts/` and `.agents/scripts/`.
   - Mirror shared specifications `02-spec/01-*` through `02-spec/20-*` (strictly skipping `02-spec/21-*` through `25-*`).
5. **Atomic Commit & Push:**
   - Stage and commit synchronized assets: `gitmap cpf "sync - update prompts, skills, shared specs, and guidelines"`.
   - Push commit to remote tracking branch: `git push origin <base_branch>`.
6. **Release Tagging & Telemetry:**
   - Create post-synchronization release tag and push to remote.

---

## 7. Connected Repositories Directory (43 Target Repositories)

| # | Repository Name | Relative Path (`../`) | Primary Stack | Base Branch |
| :---: | :--- | :--- | :--- | :--- |
| **01** | `ai-empathy-prompt-tuner` | `../02-prompts/ai-empathy-prompt-tuner` | TypeScript / Prompts | `main` |
| **02** | `alim-cv` | `../alim-cv` | React / TypeScript | `main` |
| **03** | `alim-karim-profile` | `../alim-karim-profile` | React / TypeScript | `main` |
| **04** | `alim.karim.profile` | `../aukgit/alim.karim.profile` | Web / Portfolio | `master` |
| **05** | `antigravity-manager` | `../antigravity-manager` | Go / CLI / Windows | `main` |
| **06** | `cat-my` | `../cat-my` | Full Stack / Automation | `main` |
| **07** | `core` | `../03-aukgo/core` | Go / Microservices | `main` |
| **08** | `digital-name-card` | `../digital-name-card` | Web / Identity | `main` |
| **09** | `flat-slide-show` | `../presentations-repos/flat-slide-show` | Web / Presentations | `main` |
| **10** | `gitlogger-new` | `../gitlogger-new` | TypeScript / CLI | `main` |
| **11** | `gitmap` | `../gitmap` | Go / SQLite / AUM CLI | `main` |
| **12** | `global-ppt-v1` | `../presentations-repos/global-ppt-v1` | Web / Slides | `main` |
| **13** | `hiltrax` | `../presentations-repos/hiltrax` | Web / Slides | `main` |
| **14** | `icon-coding-guidelines` | `../icon-coding-guidelines` | SVG / Design System | `main` |
| **15** | `img-pdf` | `../img-pdf` | Go / CLI / Imaging | `main` |
| **16** | `ki-health-ppt` | `../presentations-repos/ki-health-ppt` | Web / Slides | `main` |
| **17** | `kubernetes-training` | `../aukgit/kubernetes-training` | Cloud / DevOps | `main` |
| **18** | `lara-licensing` | `../lara-licensing` | PHP / Laravel | `main` |
| **19** | `lara-publishing` | `../lara-publishing` | PHP / Laravel | `main` |
| **20** | `laravel-automation` | `../laravel-automation` | PHP / Automation | `main` |
| **21** | `letsmarknow-ui` | `../letsmarknow-ui` | React / TypeScript | `main` |
| **22** | `letsmarknow` | `../letsmarknow` | Full Stack / App | `main` |
| **23** | `macro-ahk` | `../macro-ahk` | AHK / Automation | `main` |
| **24** | `maid-app-spec-presentation` | `../presentations-repos/maid-app-spec-presentation` | Web / Slides | `main` |
| **25** | `movie-cli` | `../movie-cli` | Go / CLI / Media | `main` |
| **26** | `pathhelper` | `../03-aukgo/pathhelper` | Go / Paths | `main` |
| **27** | `presentation-aug-2026-plans-alim` | `../presentations-repos/presentation-aug-2026-plans-alim` | Web / Slides | `main` |
| **28** | `prompts-connect` | `../02-prompts/prompts-connect` | TypeScript / CLI | `main` |
| **29** | `punam-case-studies-v1` | `../punam-case-studies-v1` | Web / Content | `main` |
| **30** | `rasia-logo` | `../presentations-repos/rasia-logo` | SVG / Branding | `main` |
| **31** | `scripts-fixer` | `../scripts-fixer` | Python / Automation | `main` |
| **32** | `slides-spec` | `../presentations-repos/slides-spec` | Markdown / Specs | `main` |
| **33** | `spec-builder` | `../spec-builder` | TypeScript / Specs | `main` |
| **34** | `sweet-digs-finder` | `../web-system/sweet-digs-finder` | Full Stack / Search | `main` |
| **35** | `ui-prompts-cat` | `../ui-prompts-cat` | UI / Prompts | `main` |
| **36** | `white-presentation-v1` | `../presentations-repos/white-presentation-v1` | Web / Slides | `main` |
| **37** | `workflowy-ui` | `../workflowy-ui` | React / TypeScript | `main` |
| **38** | `workflowy` | `../workflowy` | Full Stack / App | `main` |
| **39** | `wp-exam` | `../wp-exam` | WordPress / PHP | `main` |
| **40** | `wp-git-log` | `../wp-git-log` | WordPress / Git | `main` |
| **41** | `wp-html-automate` | `../wp-html-automate` | WordPress / Automation | `main` |
| **42** | `wp-link-manager` | `../wp-link-manager` | WordPress / Links | `main` |
| **43** | `wp-onboarding` | `../wp-onboarding` | WordPress / Onboarding | `main` |

---

## 8. Quality Gates & Linter Verification Protocol

The directory structure migration and synchronization engine are verified via four automated gates:

```bash
# 1. Verify prompt index completeness and zero orphans
python linter-scripts/check-prompts-loaded.py

# 2. Verify relative path hygiene (zero absolute paths or file protocol URIs)
python linter-scripts/check-relative-paths.py

# 3. Dry-run synchronization across all 43 repositories
python 03-ai-scripts/38-sync-prompts-skills-scripts.py --dry-run

# 4. Retrospective AI verification
python 03-ai-scripts/47-retrospective-ai-verification.py --tasks 3 --since-minutes 34
```

---

## 9. Binary Acceptance Criteria

| # | Acceptance Criterion | Verification Command / Metric | Status |
| :---: | :--- | :--- | :---: |
| **AC-01** | `01-prompts/22-letterly/cursor/` is moved to `01-prompts/23-cursor-prompts/` and removed from `22-letterly/` | `Test-Path 01-prompts/22-letterly/cursor` returns False | [ ] |
| **AC-02** | `01-prompts/23-sync/` is re-sequenced to `01-prompts/24-sync/` | `Test-Path 01-prompts/24-sync` returns True | [ ] |
| **AC-03** | Inside `24-sync/`, prefix collision is resolved to `01-sync.md` and `02-sync-other-codebase.md` | `Test-Path 01-prompts/24-sync/01-sync.md` & `02-sync-other-codebase.md` return True | [ ] |
| **AC-04** | `01-prompts/24-ai-verification/` is re-sequenced to `01-prompts/25-ai-verification/` | `Test-Path 01-prompts/25-ai-verification/01-retrospective-ai-verification.md` returns True | [ ] |
| **AC-05** | `.ai-memory/prompts.md` includes `01-sync.md` row and updated paths for categories 23, 24, and 25 | `python linter-scripts/check-prompts-loaded.py` exits 0 with 0 orphan prompts | [ ] |
| **AC-06** | Category READMEs updated for `01-prompts/readme.md`, `22-letterly`, `23-cursor-prompts`, `24-sync`, `25-ai-verification` | Manual inspection & heading check | [ ] |
| **AC-07** | Skills `.agents/skills/sync` and `.cursor/skills/sync` reference `01-prompts/24-sync/01-sync.md` | Skill markdown contains updated canonical prompt path | [ ] |
| **AC-08** | Skills `.agents/skills/ai-verification` and `.cursor/skills/ai-verification` reference `01-prompts/25-ai-verification/` | Skill markdown contains updated prompt reference | [ ] |
| **AC-09** | All 43 connected repositories pulled, backed up, synchronized, and verified via `38-sync-prompts-skills-scripts.py` | Sync execution summary reports 43/43 successful repositories | [ ] |
| **AC-10** | Strict relative git paths only throughout all modified specifications and plans | `python linter-scripts/check-relative-paths.py` reports zero violations | [ ] |
