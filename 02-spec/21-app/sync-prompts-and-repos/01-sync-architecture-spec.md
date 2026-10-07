# Multi-Repository Synchronization & Canonical Mirroring Architecture Specification

> **Specification:** `02-spec/21-app/sync-prompts-and-repos/01-sync-architecture-spec.md`  
> **Status:** APPROVED  
> **Architecture Version:** 6.0.0  
> **Runtime:** Google Antigravity 2.0 (IDE & CLI)  

---

## 1. User Request (Verbatim)

```text
Create a prompt sync folder with the sequence, and its job is to sync the repositories. Only the repository names will be there. It will just say as sync, zero one sync dot MD file, and create a skill for this. For the agent cursor, deploy all these things to all these repositories. First, pull these repositories, ensure you have the latest, and then take a backup of those repositories because the prompts are going to change. Then return to the previous branch, sync the prompts, and ensure not to touch the above script and other changes in the AI script folder for that repository. If there is no touch on that file, then yes, our one will be prioritized and pushed. The same goes for spec. One to 2120 folder, synchronize from the coding guideline. Do not touch the rest of the things. Be very careful when modifying things that are modified. Remember that.

Add two new prompts in the CG execute. One is for following the other prompts. Also, try to follow the newest prompt to execute the below n steps. Make the prompts similar to this goal prompt, goal writing, how we have done it for the below n tasks, how we have written it. All the CG execute prompts need to be similar to that, and also the execute prompts, and the spec writing audit. Create another prompt folder called V4 and do all these changes. Synchronize the skill and the repositories that you have modified and synchronize, seeing all these repositories do sync the prompts and the skills one more time by creating a backup of the current one and then putting these current prompts and skill.

When synchronizing with other repositories, keep in mind: the spec 21 folder should not sync. If it has updates, leave it as it is. AI scripts, if new, should be put there. But if the existing script is modified by the repository itself, do not touch this. Also, do not modify the bump script because it can be updated using its repo's own stuff. Keep this in mind, and based on that, if you think the script needs to be updated, follow that. Update in your memory so that this never happens. Whenever I say sync, you should follow this through.

Once you have everything, update all these prompts and skills to all these repositories. Before that, make sure to pull each one of the repositories. Then try to sync. Also, similar should go for ai-memory/memory folder and plans if they exist don't modify.
```

---

## 2. Architectural Invariants & Non-Negotiable Boundaries

When synchronizing canonical assets across the connected 43-repository fleet:

1. **Pre-Change Safety Protocol (Pull & Backup):**
   - For every connected repository:
     1. Detect base branch (`main` or `master`).
     2. Pull latest upstream changes: `git pull origin <base_branch> --no-rebase`.
     3. Create and push safety backup branch: `git checkout -b backup/sync-<timestamp>` and `git push origin backup/sync-<timestamp>`.
     4. Return to base branch: `git checkout <base_branch>`.

2. **The 5 Non-Negotiable Boundaries:**
   - **Boundary 1 (Spec 21 Exclusion):** TOTAL BAN on synchronizing or touching `02-spec/21-*` through `02-spec/25-*` (private application domain specs, issues, db, UI design system). Only shared specs `02-spec/01-*` through `02-spec/20-*` are synchronized.
   - **Boundary 2 (Additive-Only AI Scripts):** Brand-new AI scripts (`03-ai-scripts/`, `.agents/scripts/`) that do not exist in the target repository are copied cleanly. Existing scripts modified by the target repository must NEVER be overwritten.
   - **Boundary 3 (Bump Script Protection):** NEVER overwrite version bump scripts or manifests (`bump*`, `bump_versions.py`, `bump-version.mjs`, `version.json`). Each repository maintains custom SemVer logic.
   - **Boundary 4 (Memory & Plans Protection):** NEVER touch, overwrite, or sync `.ai-memory/memory/` or `.ai-memory/plans/` in target repositories.
   - **Boundary 5 (Zero Secrets Leakage):** NEVER synchronize `.env` files, tokens, or credentials across repositories. Secrets reside strictly in `repo-secrets` via `gitmap rs`.

3. **Sequenced Prompt Sync Structure:**
   - `01-prompts/24-sync/01-sync.md`: Sequenced canonical prompt containing the 43 target repositories in execution order.
   - Skills: `.agents/skills/sync/skill.md` and `.cursor/skills/sync/skill.md` deploying to both Antigravity and Cursor environments.

4. **CG-Execute Prompts & V4 Folder:**
   - `01-prompts/15-cg-execute/31-cg-execute-in-below-steps.md`: Canonical goal-first, subagent A=2 H=2 execution for appended below instructions.
   - `01-prompts/15-cg-execute/32-cg-follow-other-prompts.md`: Ingesting and executing referenced external prompts and guideline directives.
   - `01-prompts/v4/` and `06-archive/v4/`: Dedicated V4 prompt tier snapshots capturing the historical evolution.

---

## 3. Connected Repositories Inventory (43 Target Repositories)

| # | Repository Name | Relative Path | Default Branch |
| :---: | :--- | :--- | :--- |
| 01 | `ai-empathy-prompt-tuner` | `../02-prompts/ai-empathy-prompt-tuner` | `main` |
| 02 | `alim-cv` | `../alim-cv` | `main` |
| 03 | `alim-karim-profile` | `../alim-karim-profile` | `main` |
| 04 | `alim.karim.profile` | `../aukgit/alim.karim.profile` | `main` |
| 05 | `antigravity-manager` | `../antigravity-manager` | `main` |
| 06 | `cat-my` | `../cat-my` | `main` |
| 07 | `core` | `../03-aukgo/core` | `main` |
| 08 | `digital-name-card` | `../digital-name-card` | `main` |
| 09 | `flat-slide-show` | `../presentations-repos/flat-slide-show` | `main` |
| 10 | `gitlogger-new` | `../gitlogger-new` | `main` |
| 11 | `gitmap` | `../gitmap` | `main` |
| 12 | `global-ppt-v1` | `../presentations-repos/global-ppt-v1` | `main` |
| 13 | `hiltrax` | `../presentations-repos/hiltrax` | `main` |
| 14 | `icon-coding-guidelines` | `../icon-coding-guidelines` | `main` |
| 15 | `img-pdf` | `../img-pdf` | `main` |
| 16 | `ki-health-ppt` | `../presentations-repos/ki-health-ppt` | `main` |
| 17 | `kubernetes-training` | `../aukgit/kubernetes-training` | `main` |
| 18 | `lara-licensing` | `../lara-licensing` | `main` |
| 19 | `lara-publishing` | `../lara-publishing` | `main` |
| 20 | `laravel-automation` | `../laravel-automation` | `main` |
| 21 | `letsmarknow-ui` | `../letsmarknow-ui` | `main` |
| 22 | `letsmarknow` | `../letsmarknow` | `main` |
| 23 | `macro-ahk` | `../macro-ahk` | `main` |
| 24 | `maid-app-spec-presentation` | `../presentations-repos/maid-app-spec-presentation` | `main` |
| 25 | `movie-cli` | `../movie-cli` | `main` |
| 26 | `pathhelper` | `../03-aukgo/pathhelper` | `develop` |
| 27 | `presentation-aug-2026-plans-alim` | `../presentations-repos/presentation-aug-2026-plans-alim` | `main` |
| 28 | `prompts-connect` | `../02-prompts/prompts-connect` | `main` |
| 29 | `punam-case-studies-v1` | `../punam-case-studies-v1` | `main` |
| 30 | `rasia-logo` | `../presentations-repos/rasia-logo` | `main` |
| 31 | `scripts-fixer` | `../scripts-fixer` | `main` |
| 32 | `slides-spec` | `../presentations-repos/slides-spec` | `main` |
| 33 | `spec-builder` | `../spec-builder` | `main` |
| 34 | `sweet-digs-finder` | `../web-system/sweet-digs-finder` | `main` |
| 35 | `ui-prompts-cat` | `../ui-prompts-cat` | `main` |
| 36 | `white-presentation-v1` | `../presentations-repos/white-presentation-v1` | `main` |
| 37 | `workflowy-ui` | `../workflowy-ui` | `main` |
| 38 | `workflowy` | `../workflowy` | `main` |
| 39 | `wp-exam` | `../wp-exam` | `main` |
| 40 | `wp-git-log` | `../wp-git-log` | `main` |
| 41 | `wp-html-automate` | `../wp-html-automate` | `main` |
| 42 | `wp-link-manager` | `../wp-link-manager` | `main` |
| 43 | `wp-onboarding` | `../wp-onboarding` | `main` |

---

## 4. Verification & Quality Gates

1. Every repository is pulled prior to sync (`git pull origin <branch> --no-rebase`).
2. A safety backup branch (`backup/sync-<timestamp>`) is created and pushed prior to mirroring.
3. Additive AI script checks strictly ensure existing repository scripts modified by target repos are never overwritten.
4. Bump scripts (`bump*`, `version.json`) remain 100% untouched.
5. All target repos commit and push atomically via GitMap (`gitmap cpf "<module> - <summary>"`).
