# Subtask 04: Mandatory Pre-Pull Sweep, Dry-Run Verification, and Multi-Repository Synchronization

> **Task ID:** `01-variadic-spread-params-and-multi-repo-subtask-04`  
> **Parent Task:** `01-variadic-spread-params-and-multi-repo`  
> **Status:** READY FOR EXECUTION  
> **Target Scripts:** `03-ai-scripts/38-sync-prompts-skills-scripts.py`, `03-ai-scripts/41-audit-all-repos.py`  
> **Specification Reference:** [02-spec/21-app/01-variadic-spread-params-and-multi-repo/02-sync-and-multi-repo-spec.md](../../../02-spec/21-app/01-variadic-spread-params-and-multi-repo/02-sync-and-multi-repo-spec.md)  
> **Prerequisite:** [Subtask 03: Multi-Repository Synchronization Script Protection Guards & Memory Safety](03-sync-script-protection.md)  

---

## 1. Executive Summary & Objective

This subtask operationalizes the multi-repository synchronization workflow across all 42 target repositories connected to the `coding-guidelines` meta-repository. The workflow guarantees that all remote updates are pulled before any local git modifications occur, validates that the four non-negotiable boundaries are strictly respected during a dry run, executes live synchronization with automated backup and release ceremonies, and audits post-sync parity.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│ PHASE 1: Mandatory Pre-Pull Sweep                                           │
│          Run `git pull origin <base_branch> --no-rebase` across 42 targets │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
┌──────────────────────────────────────▼──────────────────────────────────────┐
│ PHASE 2: Fleet Dry-Run Verification                                         │
│          Execute `python 38-sync-*.py --dry-run` and inspect boundary stats │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
┌──────────────────────────────────────▼──────────────────────────────────────┐
│ PHASE 3: Single-Repo Smoke Test                                             │
│          Test against `movie-cli` or similar standalone repository          │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
┌──────────────────────────────────────▼──────────────────────────────────────┐
│ PHASE 4: Full Multi-Repository Synchronization                              │
│          Execute live sync with 6 parallel workers                          │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
┌──────────────────────────────────────▼──────────────────────────────────────┐
│ PHASE 5: Post-Sync Verification & Fleet Audit                               │
│          Execute audit script and verify clean trees & release tags         │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 2. Target Repositories Directory (42 Connected Codebases)

The 42 target repositories configured for synchronization reside under the parent workspace root:

| # | Repository Slug | Relative Workspace Directory | Ecosystem | Base Branch | Primary Function |
| :- | :--- | :--- | :--- | :--- | :--- |
| 1 | `ai-empathy-prompt-tuner` | `02-prompts/ai-empathy-prompt-tuner` | Python | `main` | Prompt Engineering & Tuning |
| 2 | `alim-cv` | `alim-cv` | Web/TS | `main` | Resume Portal |
| 3 | `alim-karim-profile` | `alim-karim-profile` | React | `main` | Personal Portfolio |
| 4 | `alim.karim.profile` | `aukgit/alim.karim.profile` | Hugo | `main` | Profile Static Site |
| 5 | `antigravity-manager` | `antigravity-manager` | TypeScript | `main` | Antigravity AI Orchestrator |
| 6 | `cat-my` | `cat-my` | React | `main` | Catalog Application |
| 7 | `core` | `03-aukgo/core` | Golang | `main` | Core Go Utilities & Engine |
| 8 | `digital-name-card` | `digital-name-card` | React | `main` | Digital Business Card |
| 9 | `flat-slide-show` | `presentations-repos/flat-slide-show` | React | `main` | Slide Presentation Engine |
| 10 | `gitlogger-new` | `gitlogger-new` | Golang | `main` | Git Activity Logger CLI |
| 11 | `global-ppt-v1` | `presentations-repos/global-ppt-v1` | React | `main` | Global Presentation Portal |
| 12 | `hiltrax` | `presentations-repos/hiltrax` | React | `main` | Enterprise Presentation Deck |
| 13 | `icon-coding-guidelines` | `icon-coding-guidelines` | SVG | `main` | Iconography Design Tokens |
| 14 | `img-pdf` | `img-pdf` | Python | `main` | Document Image & PDF CLI |
| 15 | `ki-health-ppt` | `presentations-repos/ki-health-ppt` | React | `main` | Healthcare Presentation Deck |
| 16 | `kubernetes-training` | `aukgit/kubernetes-training` | K8s | `main` | Kubernetes Training Docs |
| 17 | `lara-licensing` | `lara-licensing` | PHP | `main` | License Management Service |
| 18 | `lara-publishing` | `lara-publishing` | PHP | `main` | Publishing Automation Engine |
| 19 | `laravel-automation` | `laravel-automation` | PHP | `main` | Task & Worker Automation |
| 20 | `letsmarknow-ui` | `letsmarknow-ui` | React | `main` | Markdown Studio UI |
| 21 | `letsmarknow` | `letsmarknow` | Node | `main` | Markdown Studio Backend |
| 22 | `macro-ahk` | `macro-ahk` | AHK | `main` | Desktop Keyboard Macro Suite |
| 23 | `maid-app-spec-presentation` | `presentations-repos/maid-app-spec-presentation` | React | `main` | Domain Presentation Deck |
| 24 | `movie-cli` | `movie-cli` | Golang | `main` | Movie Metadata Search CLI |
| 25 | `pathhelper` | `03-aukgo/pathhelper` | Golang | `main` | Cross-Platform Path Library |
| 26 | `presentation-aug-2026-plans-alim` | `presentations-repos/presentation-aug-2026-plans-alim` | React | `main` | Executive Slide Presentation |
| 27 | `prompts-connect` | `02-prompts/prompts-connect` | Python | `main` | Prompt Connection Bus |
| 28 | `punam-case-studies-v1` | `punam-case-studies-v1` | React | `main` | Case Studies Showcase |
| 29 | `rasia-logo` | `presentations-repos/rasia-logo` | SVG | `main` | Vector Brand Asset Store |
| 30 | `scripts-fixer` | `scripts-fixer` | Python | `main` | Tool & Script Repair Suite |
| 31 | `slides-spec` | `presentations-repos/slides-spec` | React | `main` | Presentation Slide Specs |
| 32 | `spec-builder` | `spec-builder` | TypeScript | `main` | Specification Builder CLI |
| 33 | `sweet-digs-finder` | `web-system/sweet-digs-finder` | React | `main` | Real Estate Search Portal |
| 34 | `ui-prompts-cat` | `ui-prompts-cat` | Web | `main` | UI Prompt Explorer |
| 35 | `white-presentation-v1` | `presentations-repos/white-presentation-v1` | React | `main` | White Theme Slide System |
| 36 | `workflowy-ui` | `workflowy-ui` | React | `main` | Outliner Tree Frontend |
| 37 | `workflowy` | `workflowy` | Node | `main` | Outliner Backend Service |
| 38 | `wp-exam` | `wp-exam` | PHP | `main` | WP Exam & Quiz Engine |
| 39 | `wp-git-log` | `wp-git-log` | PHP | `main` | WP Git Activity Tracker |
| 40 | `wp-html-automate` | `wp-html-automate` | PHP | `main` | WP HTML Processing Plugin |
| 41 | `wp-link-manager` | `wp-link-manager` | PHP | `main` | WP Link Governance Plugin |
| 42 | `wp-onboarding` | `wp-onboarding` | PHP | `main` | WP Onboarding Workflow |

> [!NOTE]
> The fleet automation orchestrator `gitmap` (`gitmap-v28`) is also connected. When running synchronization, `gitmap` is handled either as the root orchestration driver or synchronized alongside the target fleet as repository 43.

---

## 3. Phase 1: Mandatory Pre-Pull Sweep Protocol

### 3.1 Pre-Pull Rationale & Guarantees

Executing `git pull` across all 42 target repositories before synchronization prevents:
1. **Push Rejections:** Remote branches receiving updates from CI workflows or other developers cause `[rejected - non-fast-forward]` if not integrated.
2. **Tag Inconsistencies:** Releases created on stale branches tag commits that do not include the latest changes on `origin`.
3. **Backup Drift:** Backup branches must reflect the exact current state of the remote repository at synchronization time.

### 3.2 Pre-Pull Command & Scripted Automation

A pre-pull sweep is automatically initiated by `03-ai-scripts/38-sync-prompts-skills-scripts.py` on line 496 during the checkout phase:

```bash
# Executed per repository on its base branch
git checkout <base_branch>
git pull origin <base_branch> --no-rebase
```

For explicit pre-flight verification across the entire fleet before launching the sync tool:

```python
# Autonomous pre-pull verification snippet
for repo_dir in TARGET_REPOS:
    if not (repo_dir / ".git").exists():
        continue
    base = detect_base_branch(repo_dir)
    code, out, err = run_cmd(f"git pull origin {base} --no-rebase", repo_dir)
    print(f"[{'OK' if code == 0 else 'ERR'}] {repo_dir.name:<30} -> {out or err}")
```

---

## 4. Phase 2: Fleet Dry-Run Verification

### 4.1 Invocation

Execute the synchronizer in preview mode:

```bash
python 03-ai-scripts/38-sync-prompts-skills-scripts.py --dry-run
```

### 4.2 Verification Invariants Checklist

During dry-run inspection, verify:
- [ ] **Spec 21 Exclusion:** 0 files under `02-spec/21-*` scheduled for copy into any target repository.
- [ ] **Additive AI Scripts:** Existing scripts in target `03-ai-scripts/` or `.agents/scripts/` marked unchanged; only genuinely new scripts scheduled for addition.
- [ ] **Bump Script Protection:** Zero version bump scripts in target repositories scheduled for overwrite.
- [ ] **Memory & Plans Protection:** Zero files under `.ai-memory/memory/` or `.ai-memory/plans/` scheduled for modification in target repositories.
- [ ] **Prompts & Skills Parity:** Canonical prompts (`01-prompts/`) and skills (`.agents/skills/`, `.cursor/skills/`) properly reflected in planned additions/updates.
- [ ] **Clean Zero-Error Summary:** Output table reports status `OK` across all accessible target repositories.

---

## 5. Phase 3: Single-Repo Smoke Test

Before triggering fleet-wide execution, run an isolated smoke test against a standalone repository:

```bash
python 03-ai-scripts/38-sync-prompts-skills-scripts.py --repo movie-cli --dry-run
```

Confirm:
1. Target directory correctly located.
2. Base branch correctly resolved (`main`).
3. Pre-release tag computed accurately.
4. Planned file copy delta matches expected prompt and skill file additions.

---

## 6. Phase 4: Full Multi-Repository Synchronization Execution

### 6.1 Live Execution Invocation

```bash
python 03-ai-scripts/38-sync-prompts-skills-scripts.py --workers 6
```

### 6.2 Per-Repository Release Ceremony Order

For each repository in the fleet, the script executes the following sequential steps:
1. **Detect & Pull:** Checkout base branch (`main`) and fast-forward pull latest commits (`--no-rebase`).
2. **Safety Backup Branch:** Create and push `backup/sync-<timestamp>` from current HEAD.
3. **Pre-Release Tag & Branch:** Ensure `vX.Y.Z` tag and `release/vX.Y.Z` branch exist on origin.
4. **Purge Archived Artifacts:** Remove any legacy `06-old-prompts` directory if found in target.
5. **Enforce Boundary-Guarded Sync:** Mirror `01-prompts/`, `.agents/skills/`, `.cursor/skills/`, additive `03-ai-scripts/`, and conditional guideline files while strictly enforcing all 4 boundary guards.
6. **Commit on Base Branch:** Commit mirrored changes with message `feat(sync): sync v6 prompts, sqlite task manager, skills, and coding guidelines` and push.
7. **Post-Release Ceremony:** Increment patch version, create `release/v<next_ver>` branch and `v<next_ver>` tag, update repo version file (if configured), push tag/branch, and merge back into base branch with `[skip ci]`.

---

## 7. Phase 5: Post-Sync Verification & Fleet Audit

### 7.1 Multi-Repository Audit Invocation

Execute the fleet audit script to verify repository invariants:

```bash
python 03-ai-scripts/41-audit-all-repos.py
```

### 7.2 Post-Sync Invariants

| Invariant | Expected Standard | Verification Method |
| :--- | :--- | :--- |
| **Branch Topology** | Active branch is strictly `main` | `git branch --show-current` |
| **Working Tree** | Completely clean, zero uncommitted files | `git status --porcelain` is empty |
| **Tag Parity** | New SemVer release tag created and pushed | `git tag --points-at HEAD` |
| **Remote Sync** | Local and origin HEADs match | `git status` reports up to date with origin |
| **Prompt Parity** | `01-prompts/` matches meta-repo prompt count | File count matching |
| **Skill Parity** | `.agents/skills/` matches meta-repo skill count | File count matching |
| **Boundary Integrity** | 0 changes in `21-*`, local scripts, bump scripts, memory/plans | Diff inspection against pre-sync backup |

---

## 8. Rollback & Emergency Contingency Procedures

If any repository fails during live execution:
1. **Identify the Failure:** The summary table flags failed repositories with status `FAIL` and exact error details.
2. **Isolate the Repository:** Other worker threads continue without interruption.
3. **Restore from Backup Branch:**
   ```bash
   git checkout <base_branch>
   git reset --hard backup/sync-<timestamp>
   git push origin <base_branch> --force-with-lease
   ```
4. **Clean Ephemeral Branches:**
   ```bash
   git branch -D release/v<failed_tag>
   git push origin --delete release/v<failed_tag>
   ```

---

## 9. Definition of Done & Quality Gates

- [ ] Pre-pull sweep executed across all 42 target repositories without merge conflicts.
- [ ] Dry-run completed with zero boundary violations verified.
- [ ] Single-repo smoke test validated clean operation.
- [ ] Live synchronization executed across fleet with 6 parallel workers.
- [ ] Backup branches and release tags verified on origin remotes.
- [ ] Post-sync audit via `03-ai-scripts/41-audit-all-repos.py` passes cleanly.
- [ ] Zero absolute paths or file URI schemes present in any authored plans or specifications.
