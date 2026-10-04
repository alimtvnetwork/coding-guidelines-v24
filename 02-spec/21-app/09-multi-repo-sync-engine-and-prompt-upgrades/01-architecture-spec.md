# Architecture Specification: Multi-Repository Synchronization Engine & Prompt Archival Architecture

> **Spec ID:** `02-spec/21-app/09-multi-repo-sync-engine-and-prompt-upgrades/01-architecture-spec.md`  
> **Status:** APPROVED  
> **Author:** Lead Systems Architect  
> **Target Subsystems:** `01-prompts/24-sync/`, `06-archive/v4/`, `03-ai-scripts/38-sync-prompts-skills-scripts.py`, `.agents/skills/sync/`, `.cursor/skills/sync/`, and all 43 connected repositories

---

## 1. Problem Statement & Strategic Objectives

1. **Multi-Repository Synchronization Workflow (`01-sync.md`):**  
   The workspace requires a dedicated prompt `01-prompts/24-sync/01-sync.md` and companion skill `sync` (for Agent and Cursor) that encapsulates the exact sequence of 43 connected repositories. This prompt drives automated repository pull, pre-change backup branch creation, canonical asset synchronization, and safe release tagging across all connected codebases.

2. **Prompt Archival (`06-archive/v4/`):**  
   To preserve prompt evolution history, a dedicated `06-archive/v4/` tier must be established alongside existing `v1/`, `v2/`, and `v3/` archives, while strictly ensuring `06-archive/` is excluded from downstream synchronization.

3. **Expanded Shared Specification Synchronization (`02-spec/01-*` to `02-spec/20-*`):**  
   In addition to prompts and skills, shared specifications `02-spec/01-*` through `02-spec/20-*` (such as `01-spec-authoring-guide`, `02-coding-guidelines`, `03-error-manage`, etc.) must synchronize across connected repositories, while `02-spec/21-*` through `02-spec/25-*` remain strictly protected.

4. **The Non-Negotiable Boundary Invariants:**  
   - **Spec 21 Exclusion:** Never touch or synchronize `02-spec/21-*` (client-specific application specs).
   - **Additive-Only AI Scripts:** Add missing scripts, but NEVER overwrite scripts modified by target repositories.
   - **Bump Script Protection:** NEVER overwrite target repository version bump scripts.
   - **Memory & Plans Protection:** NEVER overwrite or touch `.ai-memory/memory/` or `.ai-memory/plans/` in target repositories.

---

## 2. Connected Repositories Inventory (43 Target Repositories)

| # | Relative Folder Path |
| :---: | :--- |
| **01** | `02-prompts/ai-empathy-prompt-tuner` |
| **02** | `alim-cv` |
| **03** | `alim-karim-profile` |
| **04** | `aukgit/alim.karim.profile` |
| **05** | `antigravity-manager` |
| **06** | `cat-my` |
| **07** | `03-aukgo/core` |
| **08** | `digital-name-card` |
| **09** | `presentations-repos/flat-slide-show` |
| **10** | `gitlogger-new` |
| **11** | `gitmap` |
| **12** | `presentations-repos/global-ppt-v1` |
| **13** | `presentations-repos/hiltrax` |
| **14** | `icon-coding-guidelines` |
| **15** | `img-pdf` |
| **16** | `presentations-repos/ki-health-ppt` |
| **17** | `aukgit/kubernetes-training` |
| **18** | `lara-licensing` |
| **19** | `lara-publishing` |
| **20** | `laravel-automation` |
| **21** | `letsmarknow-ui` |
| **22** | `letsmarknow` |
| **23** | `macro-ahk` |
| **24** | `presentations-repos/maid-app-spec-presentation` |
| **25** | `movie-cli` |
| **26** | `03-aukgo/pathhelper` |
| **27** | `presentations-repos/presentation-aug-2026-plans-alim` |
| **28** | `02-prompts/prompts-connect` |
| **29** | `punam-case-studies-v1` |
| **30** | `presentations-repos/rasia-logo` |
| **31** | `scripts-fixer` |
| **32** | `presentations-repos/slides-spec` |
| **33** | `spec-builder` |
| **34** | `web-system/sweet-digs-finder` |
| **35** | `ui-prompts-cat` |
| **36** | `presentations-repos/white-presentation-v1` |
| **37** | `workflowy-ui` |
| **38** | `workflowy` |
| **39** | `wp-exam` |
| **40** | `wp-git-log` |
| **41** | `wp-html-automate` |
| **42** | `wp-link-manager` |
| **43** | `wp-onboarding` |

---

## 3. Synchronization Pipeline Ceremony

For each target repository, execution proceeds in 6 discrete stages:

1. **Pre-flight Pull:** Checkout active base branch and pull remote changes:  
   `git pull origin <base_branch> --no-rebase`
2. **Pre-change Backup Branch:** Create and push dedicated backup branch before changes:  
   `git checkout -b backup/sync-<timestamp>` -> `git push origin backup/sync-<timestamp>`
3. **Return to Base Branch:** Return to base branch:  
   `git checkout <base_branch>`
4. **Controlled Asset Mirroring:**
   - Mirror `01-prompts/` (excluding `06-archive/`).
   - Mirror `.agents/skills/` and `.cursor/skills/`.
   - Add new AI scripts to `03-ai-scripts/` and `.agents/scripts/` (additive only; NEVER overwrite modified scripts).
   - Mirror shared `02-spec/01-*` through `02-spec/20-*` specifications, strictly skipping `02-spec/21-*`.
   - Protect all bump scripts (`bump*`).
   - Protect `.ai-memory/memory` and `.ai-memory/plans`.
5. **Atomic Commit & Push:**
   Commit changes using `gitmap cpf "sync - update prompts, skills, and coding guidelines"`.
6. **Release Tagging:**
   Create post-change release tag and push to remote.

---

## Acceptance Criteria

- [x] `01-prompts/24-sync/01-sync.md` created with V6 continuous execution architecture and full 43-repo sequence.
- [ ] Companion skills `.agents/skills/sync/skill.md` and `.cursor/skills/sync/skill.md` created.
- [ ] `06-archive/v4/` folder created with `readme.md`.
- [ ] `03-ai-scripts/38-sync-prompts-skills-scripts.py` updated to synchronize `02-spec/01-*` through `02-spec/20-*` while protecting `02-spec/21-*`.
- [ ] Memory updated in `.ai-memory/memory/learned/18-cross-repository-sync-rules.md`.
- [ ] All 43 connected repositories pulled, backed up, synchronized, and verified.
- [ ] Working tree in `coding-guidelines` committed cleanly and pushed.
