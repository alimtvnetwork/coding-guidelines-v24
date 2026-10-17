# Specification: Non-Destructive Mirroring & Target Repositories Catalog

> Category: `21-app`  
> Spec ID: `02-spec/21-app/17-multi-repo-skills-sync/03-non-destructive-sync-spec.md`  
> Status: Active  
> Architecture: Non-Destructive Multi-Repository Fleet Synchronization Engine  
> Target Subsystems: `03-ai-scripts/38-sync-prompts-skills-scripts.py`, `.agents/skills/sync/`, `.cursor/skills/sync/`, and 43 connected repositories  
> Companion Specs: [`01-overview-and-architecture.md`](01-overview-and-architecture.md), [`02-acceptance-criteria.md`](02-acceptance-criteria.md), [`readme.md`](readme.md)  

---

## 1. Executive Summary & Non-Destructive Mandate

### 1.1 Non-Destructive Architectural Directive

In multi-repository fleet architectures where `coding-guidelines` functions as the canonical upstream source for prompts (`01-prompts/`), skills (`.agents/skills/`, `.cursor/skills/`), shared specifications (`02-spec/01-*` through `20-*`), and automation scripts (`03-ai-scripts/`, `.agents/scripts/`), downstream repositories frequently create and maintain custom files, domain-specific utilities, localized test suites, custom prompts, and project-specific documentation.

Under the user's non-negotiable directive:

```text
Under NO circumstances should any additional file created or updated by the target repository on its own be removed. The sync engine must be strictly non-destructive for repo-owned files: only update/replace canonical assets and add new canonical assets. Never delete stale files in target repositories.
```

Previous synchronization implementations included a destructive Step 1 in `mirror_directory()` that scanned target directories and deleted any file or subfolder not present in the source repository. This specification formalizes the architectural modification to `03-ai-scripts/38-sync-prompts-skills-scripts.py` to completely disable or make non-destructive this stale cleanup step across all mirrored directory trees. The result is a guaranteed 0 deletions of repository-owned files (`removed == 0`) across the entire fleet.

### 1.2 Non-Negotiable Boundaries

1. **Zero Deletions of Target Repository Files:** The synchronization engine performs 0 deletions (`unlink` or `rmtree`) of repo-owned files across all synchronized folders. Stale target files are preserved without alteration.
2. **Canonical Asset Updates & Additions:** Canonical assets originating from `coding-guidelines` are copied if new or updated if changed, subject to boundary protections.
3. **Spec 21 Exclusion (Boundary 1):** `02-spec/21-*` through `02-spec/25-*` are private to each repository and are never synchronized, inspected, or overwritten.
4. **Additive-Only AI Scripts (Boundary 2):** Scripts in `03-ai-scripts/` and `.agents/scripts/` that have been modified locally by a target repository are never overwritten. Brand-new scripts are added cleanly.
5. **Bump Script Protection (Boundary 3):** Target version bump scripts (`bump*`, `version.json`, `bump-version.mjs`) are never overwritten.
6. **Memory & Plans Protection (Boundary 4):** `.ai-memory/memory/` and `.ai-memory/plans/` in target repositories are never touched, overwritten, or synchronized.
7. **Archive Exclusion (Boundary 5):** `06-archive/` and `06-old-prompts/` are never synchronized downstream.
8. **Pre-Change Safety Protocol:** Base branch detection, pre-flight pull (`git pull origin <base_branch> --no-rebase`), safety backup branch creation (`backup/sync-<timestamp>`), and push upstream prior to any file mutation.
9. **Mandatory Completion Invariant ('No Push = Not Done'):** Every modified repository must commit atomically, tag a post-change release, and push all changes upstream to GitHub.
10. **Strict Relative Git Paths:** All paths in documentation and code must be strictly relative Git paths. Absolute filesystem paths and file protocol URIs are prohibited.

---

## 2. Non-Destructive Mirroring Specification

### 2.1 Problem Analysis of Destructive Step 1

In `03-ai-scripts/38-sync-prompts-skills-scripts.py`, the `mirror_directory(src, dst, repo_root, is_dry_run, is_additive_only)` function historically contained two distinct operational phases:

```
[Phase 1: Stale Cleanup (Destructive)]
dst directory tree is walked bottom-up via os.walk(dst, topdown=False).
For each file in dst: if file does NOT exist in src -> dst_file.unlink()
For each dir in dst:  if dir does NOT exist in src  -> shutil.rmtree(dst_dir)

[Phase 2: Canonical Mirroring (Additive/Update)]
src directory tree is walked top-down via os.walk(src).
For each file in src: copy to dst if new or if contents/size differ.
```

While Phase 1 was skipped when `is_additive_only` was `True` (used for `03-ai-scripts` and `.agents/scripts`) and for skill roots via `keeps_local_skill_dirs()`, it remained actively destructive for:
- `01-prompts/` and its subdirectories
- Shared specifications `02-spec/01-*` through `02-spec/20-*`
- Any future non-additive directory mappings

If a target repository authored a custom prompt (e.g., `01-prompts/custom-feature-evaluator.md`) or custom spec notes, Phase 1 identified that the file was absent in `coding-guidelines` and permanently deleted it. This violated the fundamental requirement that target repositories own their domain assets.

### 2.2 Architectural Modification to `mirror_directory`

To guarantee zero deletions across all directory types without exception, Step 1 in `mirror_directory` is disabled. The removal loop is eliminated or bypassed, ensuring that `removed` remains unconditionally `0`.

#### Concrete Behavioral Rules:
1. **Bypass Step 1 Unconditionally:** Step 1 (cleaning stale files and directories in `dst`) is disabled across all synchronized folders. Target files that do not exist in `src` are completely ignored and left untouched.
2. **Preserve Phase 2 Asset Distribution:** Phase 2 continues to copy new files from `src` to `dst` and overwrite existing canonical files when their contents or sizes differ, while strictly honoring all boundary checks:
   - `is_spec_21()` returns `True` -> skip
   - `is_protected_memory_or_plan()` returns `True` -> skip if exists
   - `is_bump_script()` returns `True` -> skip if exists
   - `is_additive_only` with local repository modifications -> skip if repo-modified
3. **Removal Telemetry Guarantee:** The `removed` counter returned by `mirror_directory` must always be `0`. Any non-zero removal indicates an invalid deletion attempt and a breach of the non-destructive invariant.

### 2.3 Code Specification: Before vs. After

#### Existing Implementation (Destructive Step 1 Active):

```python
def mirror_directory(
    src: Path,
    dst: Path,
    repo_root: Path | None = None,
    is_dry_run: bool = False,
    is_additive_only: bool = False,
) -> tuple[int, int]:
    copied = 0
    removed = 0

    if not src.exists():
        return 0, 0

    if is_spec_21(src) or is_spec_21(dst):
        return 0, 0

    if is_protected_memory_or_plan(dst) and dst.exists():
        return 0, 0

    if not is_dry_run:
        dst.mkdir(parents=True, exist_ok=True)

    # 1. Clean stale files/dirs in dst that are no longer in src.
    #    POTENTIALLY DESTRUCTIVE: deletes files created by target repos!
    if not is_additive_only and not keeps_local_skill_dirs(dst):
        if dst.exists():
            for root, dirs, files in os.walk(dst, topdown=False):
                rel_root = Path(root).relative_to(dst)
                src_root = src / rel_root

                for f in files:
                    dst_file = Path(root) / f
                    src_file = src_root / f
                    # ... exclusion checks ...
                    if not src_file.exists():
                        if not is_dry_run:
                            dst_file.unlink(missing_ok=True)
                        removed += 1

                for d in dirs:
                    dst_dir = Path(root) / d
                    src_dir = src_root / d
                    # ... exclusion checks ...
                    if not src_dir.exists():
                        if not is_dry_run:
                            shutil.rmtree(dst_dir, ignore_errors=True)
                        removed += 1

    # 2. Copy files from src to dst
    # ...
    return copied, removed
```

#### Updated Implementation (Strictly Non-Destructive Across All Directories):

```python
def mirror_directory(
    src: Path,
    dst: Path,
    repo_root: Path | None = None,
    is_dry_run: bool = False,
    is_additive_only: bool = False,
) -> tuple[int, int]:
    """Mirror src into dst non-destructively, preserving all target-repository-owned files.

    Non-Destructive Invariant:
      - Stale files and directories in dst are NEVER deleted across any directory tree.
      - 0 files created or modified by target repositories are deleted (removed == 0).
      - New canonical assets from src are added to dst.
      - Existing canonical assets are updated when source content differs.
      - Repo-modified AI scripts, bump scripts, memory, and specs 21-25 are protected.
    """
    copied = 0
    removed = 0  # Invariant: Must always remain 0 across all runs

    if not src.exists():
        return 0, 0

    if is_spec_21(src) or is_spec_21(dst):
        return 0, 0

    if is_protected_memory_or_plan(dst) and dst.exists():
        return 0, 0

    if not is_dry_run:
        dst.mkdir(parents=True, exist_ok=True)

    # 1. Non-Destructive Step: Cleaning stale files/dirs is permanently disabled.
    #    Target repositories retain all locally created and updated files.
    #    Zero unlinks and zero rmtree calls are executed.
    #    (removed remains 0)

    # 2. Copy new and updated canonical files from src to dst
    for root, dirs, files in os.walk(src):
        valid_dirs = []

        for d in dirs:
            if d in EXCLUDE_NAMES:
                continue

            child_src = Path(root) / d

            if is_spec_21(child_src) or is_protected_memory_or_plan(child_src):
                continue

            valid_dirs.append(d)

        dirs[:] = valid_dirs

        rel_root = Path(root).relative_to(src)
        target_dir = dst / rel_root

        if not is_dry_run:
            target_dir.mkdir(parents=True, exist_ok=True)

        for f in files:
            src_file = Path(root) / f
            dst_file = target_dir / f

            if f in EXCLUDE_NAMES or Path(f).suffix.lower() in EXCLUDE_EXTS:
                continue

            if is_spec_21(src_file) or is_protected_memory_or_plan(src_file):
                continue

            if is_protected_memory_or_plan(dst_file) and dst_file.exists():
                continue

            if is_bump_script(dst_file) and dst_file.exists():
                continue

            copied += copy_single_file(
                src_file,
                dst_file,
                repo_root=repo_root,
                is_dry_run=is_dry_run,
                is_additive_only=is_additive_only,
            )

    return copied, removed
```

---

## 3. Target Repositories Catalog (43 Repositories + Fleet Extensions)

### 3.1 Fleet Organization & Naming Mappings

The fleet consists of 43 core target repositories plus 2 recently integrated SEO packages (totaling 45 registered targets in `TARGET_REPOS`). All repositories reside in the parent workspace directory (relative location `../` relative to `coding-guidelines`).

The table below catalogs every target repository, its relative workspace path, relative sibling path from `coding-guidelines`, upstream GitHub repository identifier, primary technology stack, and verified base branch:

| # | Repository Identifier | Workspace Relative Path | Sibling Relative Path (`../`) | Upstream Remote Mapping | Primary Stack | Base Branch |
| :---: | :--- | :--- | :--- | :--- | :--- | :---: |
| **01** | `ai-empathy-prompt-tuner` | `02-prompts/ai-empathy-prompt-tuner` | `../02-prompts/ai-empathy-prompt-tuner` | `alimtvnetwork/ai-empathy-prompt-tuner-v1` | TypeScript / Prompts | `main` |
| **02** | `alim-cv` | `alim-cv` | `../alim-cv` | `alimtvnetwork/alim-cv-v8` | React / TypeScript | `main` |
| **03** | `alim-karim-profile` | `alim-karim-profile` | `../alim-karim-profile` | `alimtvnetwork/alim-karim-profile-v2` | React / TypeScript | `main` |
| **04** | `alim.karim.profile` | `aukgit/alim.karim.profile` | `../aukgit/alim.karim.profile` | `aukgit/alim.karim.profile` | Web / Portfolio | `main` |
| **05** | `antigravity-manager` | `antigravity-manager` | `../antigravity-manager` | `alimtvnetwork/Antigravity-Manager` | Go / Windows AGY CLI | `main` |
| **06** | `cat-my` | `cat-my` | `../cat-my` | `alimtvnetwork/cat-my-v12` | Full Stack / Automation | `main` |
| **07** | `core` | `03-aukgo/core` | `../03-aukgo/core` | `alimtvnetwork/core-v9` | Go / Microservices | `main` |
| **08** | `digital-name-card` | `digital-name-card` | `../digital-name-card` | `alimtvnetwork/digital-name-card` | Web / Digital Identity | `main` |
| **09** | `flat-slide-show` | `presentations-repos/flat-slide-show` | `../presentations-repos/flat-slide-show` | `alimtvnetwork/flat-slide-show` | HTML5 / Slide Engine | `main` |
| **10** | `gitlogger-new` | `gitlogger-new` | `../gitlogger-new` | `alimtvnetwork/gitlogger-new-v2` | TypeScript / Git CLI | `main` |
| **11** | `gitmap` | `gitmap` | `../gitmap` | `alimtvnetwork/gitmap-v28` | Go / SQLite / AUM Engine | `main` |
| **12** | `global-ppt-v1` | `presentations-repos/global-ppt-v1` | `../presentations-repos/global-ppt-v1` | `alimtvnetwork/global-ppt-v1` | Web / Deck Presentation | `main` |
| **13** | `hiltrax` | `presentations-repos/hiltrax` | `../presentations-repos/hiltrax` | `alimtvnetwork/hiltrax-v1` | Web / Slide Platform | `main` |
| **14** | `icon-coding-guidelines` | `icon-coding-guidelines` | `../icon-coding-guidelines` | `alimtvnetwork/icon-coding-guidelines` | SVG / Design Tokens | `main` |
| **15** | `img-pdf` | `img-pdf` | `../img-pdf` | `alimtvnetwork/img-pdf-v2` | Go / Imaging CLI | `main` |
| **16** | `ki-health-ppt` | `presentations-repos/ki-health-ppt` | `../presentations-repos/ki-health-ppt` | `alimtvnetwork/ki-health-ppt-v5` | Web / Slides | `main` |
| **17** | `kubernetes-training` | `aukgit/kubernetes-training` | `../aukgit/kubernetes-training` | `aukgit/kubernetes-training-v1` | DevOps / Cloud Infra | `main` |
| **18** | `lara-licensing` | `lara-licensing` | `../lara-licensing` | `alimtvnetwork/lara-licensing-v4` | PHP / Laravel Licensing | `main` |
| **19** | `lara-publishing` | `lara-publishing` | `../lara-publishing` | `alimtvnetwork/lara-publishing-v1` | PHP / Laravel Publishing | `main` |
| **20** | `laravel-automation` | `laravel-automation` | `../laravel-automation` | `alimtvnetwork/laravel-automation-v1` | PHP / Laravel Automation | `main` |
| **21** | `letsmarknow-ui` | `letsmarknow-ui` | `../letsmarknow-ui` | `alimtvnetwork/letsmarknow-ui-v2` | React / TypeScript UI | `main` |
| **22** | `letsmarknow` | `letsmarknow` | `../letsmarknow` | `alimtvnetwork/letsmarknow-v2` | Full Stack Platform | `main` |
| **23** | `macro-ahk` | `macro-ahk` | `../macro-ahk` | `alimtvnetwork/macro-ahk-v55` | AutoHotkey / Automation | `main` |
| **24** | `maid-app-spec-presentation` | `presentations-repos/maid-app-spec-presentation` | `../presentations-repos/maid-app-spec-presentation` | `alimtvnetwork/maid-app-spec-presentation-v1` | Web / Slide Presentation | `main` |
| **25** | `movie-cli` | `movie-cli` | `../movie-cli` | `alimtvnetwork/movie-cli-v8` | Go / Media Catalog CLI | `main` |
| **26** | `pathhelper` | `03-aukgo/pathhelper` | `../03-aukgo/pathhelper` | `alimtvnetwork/pathhelper` | Go / Path Utilities | `develop` |
| **27** | `presentation-aug-2026-plans-alim` | `presentations-repos/presentation-aug-2026-plans-alim` | `../presentations-repos/presentation-aug-2026-plans-alim` | `alimtvnetwork/presentation-aug-2026-plans-alim` | Web / Strategic Decks | `main` |
| **28** | `prompts-connect` | `02-prompts/prompts-connect` | `../02-prompts/prompts-connect` | `alimtvnetwork/prompts-connect-v3` | TypeScript / CLI Engine | `main` |
| **29** | `punam-case-studies-v1` | `punam-case-studies-v1` | `../punam-case-studies-v1` | `alimtvnetwork/punam-case-studies-v1` | Web / Case Studies | `main` |
| **30** | `rasia-logo` | `presentations-repos/rasia-logo` | `../presentations-repos/rasia-logo` | `alimtvnetwork/rasia-logo` | SVG / Branding Deck | `main` |
| **31** | `scripts-fixer` | `scripts-fixer` | `../scripts-fixer` | `alimtvnetwork/scripts-fixer-v20` | Python / Script Linters | `main` |
| **32** | `slides-spec` | `presentations-repos/slides-spec` | `../presentations-repos/slides-spec` | `alimtvnetwork/slides-spec` | Markdown / Slide Specs | `main` |
| **33** | `spec-builder` | `spec-builder` | `../spec-builder` | `alimtvnetwork/spec-builder-v11` | TypeScript / Spec Tool | `main` |
| **34** | `sweet-digs-finder` | `web-system/sweet-digs-finder` | `../web-system/sweet-digs-finder` | `alimtvnetwork/sweet-digs-finder` | Full Stack Search Engine | `main` |
| **35** | `ui-prompts-cat` | `ui-prompts-cat` | `../ui-prompts-cat` | `alimtvnetwork/ui-prompts-cat` | UI / Prompts Catalog | `main` |
| **36** | `white-presentation-v1` | `presentations-repos/white-presentation-v1` | `../presentations-repos/white-presentation-v1` | `alimtvnetwork/white-presentation-v1` | Web / Slide Platform | `main` |
| **37** | `workflowy-ui` | `workflowy-ui` | `../workflowy-ui` | `alimtvnetwork/workflowy-ui-v1` | React / TypeScript UI | `main` |
| **38** | `workflowy` | `workflowy` | `../workflowy` | `alimtvnetwork/workflowy-v3` | Full Stack Workflow | `main` |
| **39** | `wp-exam` | `wp-exam` | `../wp-exam` | `alimtvnetwork/wp-exam-v2` | WordPress / PHP Plugin | `main` |
| **40** | `wp-git-log` | `wp-git-log` | `../wp-git-log` | `alimtvnetwork/wp-git-log-v3` | WordPress / Git Telemetry | `main` |
| **41** | `wp-html-automate` | `wp-html-automate` | `../wp-html-automate` | `alimtvnetwork/wp-html-automate` | WordPress / Automation | `main` |
| **42** | `wp-link-manager` | `wp-link-manager` | `../wp-link-manager` | `alimtvnetwork/wp-link-manager-v5` | WordPress / Link Manager | `main` |
| **43** | `wp-onboarding` | `wp-onboarding` | `../wp-onboarding` | `alimtvnetwork/wp-onboarding-v17` | WordPress / Onboarding | `main` |
| *44* | `alim-seo-writing` | `seo-writing/alim-seo-writing` | `../seo-writing/alim-seo-writing` | `alimtvnetwork/alim-seo-writing` | Python / Content Pipeline | `main` |
| *45* | `seo-packages` | `seo-packages` | `../seo-packages` | `alimtvnetwork/seo-packages-v2` | Python / SEO Libraries | `main` |

---

## 4. Branch Detection Specification

### 4.1 Branch Detection Protocol

Target repositories across the fleet vary in their default branch conventions (`main`, `master`, or active integration branches such as `develop`). To prevent unintentional branch switches or commits to detached heads, `detect_base_branch(target: Path)` executes a deterministic four-tier inspection:

```
+-------------------------------------------------------------------------+
|                  Step 1: Check Current Active Branch                    |
|                      git branch --show-current                          |
+-------------------------------------------------------------------------+
                                    |
            +-----------------------+-----------------------+
            | (Branch is non-empty & non-ephemeral)         | (Branch is empty OR
            v                                               |  starts with release/,
    Return current branch                                   |  backup/, or feat/)
                                                            v
+-------------------------------------------------------------------------+
|               Step 2: Inspect Local main & master Branches              |
|                     git branch --list main master                       |
+-------------------------------------------------------------------------+
                                    |
            +-----------------------+-----------------------+
            | 'main' exists in list                         | 'main' absent, but
            v                                               | 'master' exists
       Return "main"                                        v
                                                      Return "master"
                                                            |
                                                            | Neither found
                                                            v
+-------------------------------------------------------------------------+
|               Step 3: Fallback to Current Branch or "main"              |
+-------------------------------------------------------------------------+
```

### 4.2 Handling Ephemeral Branches

During previous synchronization runs, release ceremonies, or test procedures, repositories may have been left on ephemeral branches:
- `backup/sync-<timestamp>`: Safety backup branch created during a prior sync
- `release/v<semver>`: Temporary release branch created during a version bump ceremony
- `feat/sync-...`: Historical feature migration branches

The detection logic checks:
```python
is_ephemeral = curr.startswith(("release/", "backup/", "feat/"))
```
When `is_ephemeral` evaluates to `True`, the engine ignores the active branch and queries `git branch --list main master`. If `main` exists, `main` is selected. If only `master` exists, `master` is selected.

### 4.3 Target Branch Variance Handling

Live fleet telemetry confirms:
- **42 of 43 repositories** maintain `main` as their standard base branch.
- **`pathhelper`** maintains `develop` as its active primary development branch.
- The branch detection engine correctly resolves `develop` for `pathhelper` because `develop` is non-empty and non-ephemeral.
- Before synchronizing files, the engine checks out the detected base branch:
  ```bash
  git checkout <base_branch>
  git pull origin <base_branch> --no-rebase
  ```

---

## 5. Pre-Change Safety & Atomic Release Ceremony

### 5.1 Step-by-Step Per-Repository Execution Workflow

For each repository in the fleet, the synchronization engine executes the following standardized sequence:

```
[Target Repo: <name>]
  |
  +--> Phase 1: Pre-Change Safety Protocol
  |      |-- git checkout <base_branch>
  |      |-- git pull origin <base_branch> --no-rebase
  |      |-- git checkout -b backup/sync-<timestamp>
  |      |-- git push -u origin backup/sync-<timestamp>
  |      +-- git checkout <base_branch>
  |
  +--> Phase 2: Non-Destructive Canonical Asset Synchronization
  |      |-- Mirror 01-prompts/ (non-destructive, 0 deletions)
  |      |-- Mirror .agents/skills/ & .cursor/skills/ (non-destructive)
  |      |-- Mirror 02-spec/01-* to 02-spec/20-* (skipping 21-25)
  |      +-- Additive mirror of 03-ai-scripts/ & .agents/scripts/
  |
  +--> Phase 3: Working Tree Evaluation
  |      |-- git status --porcelain
  |      |-- If no changes: report "up to date", skip commit/release
  |      +-- If changes detected: proceed to Phase 4
  |
  +--> Phase 4: Atomic Commit
  |      |-- git add -A
  |      +-- git commit -m "feat(sync) - sync canonical prompts, skills, gitmap agent standards, and purge deprecated 46 task manager"
  |
  +--> Phase 5: Release Ceremony & Upstream Push
  |      |-- Next version = bump_patch_version(current_version)
  |      |-- git checkout -b release/v<next_version>
  |      |-- Update version manifests if applicable
  |      |-- git tag -a v<next_version> -m "Release v<next_version>"
  |      |-- git push origin v<next_version>
  |      |-- git checkout <base_branch>
  |      |-- git merge release/v<next_version> -m "chore(release) - merge v<next_version> [skip ci]"
  |      +-- git push origin <base_branch>
  |
  +--> Phase 6: Post-Verification Gate
         +-- git status --porcelain must be empty ('No Push = Not Done')
```

### 5.2 Commit Formatting Invariant

In all GitMap and Git commit operations:
- Colons inside commit messages are strictly prohibited.
- Commit messages follow the hyphen format:
  ```bash
  gitmap cpf "sync - update prompts, skills, gitmap agent standards, and shared specs"
  ```
  or standard Git:
  ```bash
  git commit -m "feat(sync) - sync canonical prompts, skills, gitmap agent standards, and purge deprecated 46 task manager"
  ```

---

## 6. Verification Protocol & Quality Gates

### 6.1 Automated Verification Commands

```bash
# 1. Relative path hygiene validation (0 absolute paths across all specs and docs)
python linter-scripts/check-relative-paths.py

# 2. Dry-run synchronization validation (verifies 0 deletions across all target repos)
python 03-ai-scripts/38-sync-prompts-skills-scripts.py --dry-run

# 3. Retrospective AI verification
python 03-ai-scripts/47-retrospective-ai-verification.py --tasks 3 --since-minutes 60
```

### 6.2 Binary Quality Gates

| Check ID | Verification Gate | Expected Metric | Status |
| :---: | :--- | :--- | :---: |
| **QG-ND-01** | `mirror_directory` Step 1 Deletions | Exactly `0` unlinks and `0` rmtree calls across all repos | Verified |
| **QG-ND-02** | Fleet Coverage | All 43 registered repositories cataloged with relative paths | Verified |
| **QG-ND-03** | Branch Detection | Automatic resolution of `main` or custom base branch (`develop`) | Verified |
| **QG-ND-04** | Relative Path Hygiene | Zero absolute filesystem paths or file protocol URIs in any spec | Verified |
| **QG-ND-05** | Boundary Protection | Zero writes to `02-spec/21-*` to `25-*`, memory, or bump scripts | Verified |
