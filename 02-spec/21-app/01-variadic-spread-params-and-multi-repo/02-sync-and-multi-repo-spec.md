# Canonical Specification: Multi-Repository Synchronization & Safe Release Protocol

> **Document Version:** 1.0.0  
> **Status:** APPROVED & ACTIVE  
> **Target Scope:** Meta-Repository (`coding-guidelines`) & 42 Connected Target Repositories  
> **Reference Scripts:** `03-ai-scripts/38-sync-prompts-skills-scripts.py`, `03-ai-scripts/41-audit-all-repos.py`  
> **Reference Plan:** `.ai-memory/plans/subtasks/01-variadic-spread-params-and-multi-repo/03-sync-script-protection.md`  

---

## 1. Executive Summary & Architectural Motivation

The `coding-guidelines` meta-repository serves as the single source of truth for canonical prompts (`01-prompts/`), agent skills (`.agents/skills/`, `.cursor/skills/`), shared automation scripts (`03-ai-scripts/`, `.agents/scripts/`), and global architecture standards (`02-spec/02-coding-guidelines/`).

To propagate governance and prompt tooling across the organization, automated synchronization is performed across 42 downstream repositories. However, synchronization must balance central governance with downstream autonomy. Downstream repositories frequently develop localized scripts, specialized version bump routines, repository-specific application specifications, and private `.ai-memory/` state.

This specification formalizes:
1. **The Four Non-Negotiable Sync Boundaries:** Hard architectural boundaries preventing overwrites of application specs, local AI scripts, custom bump scripts, and repository memory/plans.
2. **The Mandatory Pre-Pull Workflow:** Pre-flight synchronization pulling the latest remote changes across all 42 target repositories before executing backup branches, tags, or file modifications.
3. **Synchronization Guard Implementation:** Exact programmatic guards including `is_protected_memory_or_plan` in `03-ai-scripts/38-sync-prompts-skills-scripts.py`.
4. **Target Repositories Directory:** Comprehensive registry of all 42 connected target repositories with relative workspace locations, categories, and branch topologies.
5. **Dry-Run & Verification Protocol:** Step-by-step verification pipeline guaranteeing zero inadvertent file loss or clobbering.

---

## 2. The Four Non-Negotiable Sync Boundaries

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                      META-REPOSITORY (coding-guidelines)                    │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
                      SYNCHRONIZATION PIPELINE (Script 38)
                                       │
         ┌─────────────────────────────┼─────────────────────────────┐
         ▼                             ▼                             ▼
┌──────────────────┐          ┌──────────────────┐          ┌──────────────────┐
│   01-prompts/    │          │  .agents/skills/ │          │ 02-coding-guide/ │
│  (Clean Mirror)  │          │  (Clean Mirror)  │          │  (Clean Mirror)  │
└──────────────────┘          └──────────────────┘          └──────────────────┘
                                       │
                        STRICT BOUNDARY GUARDS (TOTAL BANS)
                                       │
   ┌───────────────────────┬───────────┴───────────┬───────────────────────┐
   ▼                       ▼                       ▼                       ▼
BOUNDARY 1              BOUNDARY 2              BOUNDARY 3              BOUNDARY 4
Spec 21 Exclusion       Additive-Only Scripts   Bump Script Guard       Memory & Plans Guard
`02-spec/21-*`          `03-ai-scripts/`        `*bump*`                `.ai-memory/memory/`
TOTAL BAN on sync.      New scripts ONLY.       NEVER overwrite         `.ai-memory/plans/`
Leave target intact.    Never overwrite/delete. target bump script.     NEVER modify/overwrite.
```

### 2.1 Boundary 1: Spec 21 Exclusion (`02-spec/21-*` / `21-app*`)

- **Rule:** `02-spec/21-*` (`02-spec/21-app`, `02-spec/21-app-issues`, `02-spec/21-app-db`, `02-spec/21-app-ui-design-system`) must NEVER be synchronized to target repositories.
- **Architectural Rationale:** Specifications residing under `02-spec/21-*` define domain models, business logic, endpoints, and application specifications unique to specific projects. Propagating meta-repo spec 21 files downstream clobbers child domain models, while pushing child spec 21 upstream clobbers central indices.
- **Enforcement Mechanism:**
  - `EXCLUDE_NAMES` set contains `"21-app"`, `"21-app-issues"`, `"21-app-db"`, `"21-app-ui-design-system"`.
  - Path discriminator `is_spec_21(path: Path) -> bool` intercepts both source and destination paths:
    ```python
    def is_spec_21(path: Path) -> bool:
        norm = str(path).replace("\\", "/").lower()

        if "/21-" in norm:
            return True

        if "spec/21" in norm:
            return True

        for part in path.parts:
            if part.startswith("21-"):
                return True

        return False
    ```

### 2.2 Boundary 2: Additive-Only AI Scripts (`03-ai-scripts/`, `.agents/scripts/`)

- **Rule:** When synchronizing AI automation scripts:
  - **New scripts:** If a script exists in the source meta-repo but is missing from the target repository (`not dst_file.exists()`), copy it cleanly.
  - **Existing scripts:** If a script already exists in the target repository (`dst_file.exists()`), **DO NOT TOUCH OR OVERWRITE IT**.
  - **No deletions:** Stale files existing in the target repo's script directory that are absent from upstream must NEVER be deleted.
- **Architectural Rationale:** Target repositories frequently specialize shared scripts (e.g., custom database credentials, different container runtimes, tailored CI flags). Overwriting existing scripts clobbers working configurations.
- **Enforcement Mechanism:**
  - Directory mapping in `SYNC_DIRS` marks script directories with `is_additive = True`.
  - `mirror_directory` skips stale-file cleanup when `is_additive_only` is true.
  - `copy_single_file` halts early if `is_additive_only` is true and `dst_file.exists()`.

### 2.3 Boundary 3: Version Bump Script Protection (`bump*`)

- **Rule:** Version bump scripts (`bump-version.mjs`, `bump_versions.py`, `37-bump-version.py`, etc.) in target repositories must NEVER be overwritten during sync.
- **Architectural Rationale:** Each repository manages its release lifecycle using distinct package formats (Node `package.json`, Python `pyproject.toml`/`version.json`, Go git tags, PHP `composer.json`). Overwriting a target repository's bump script breaks automated release ceremonies.
- **Enforcement Mechanism:**
  - Evaluated via `is_bump_script(path: Path) -> bool`:
    ```python
    def is_bump_script(path: Path) -> bool:
        name = path.name.lower()
        is_bump = "bump" in name
        is_version = "version" in name or name.startswith("bump")

        if is_bump:
            if is_version:
                return True

        return False
    ```
  - If `is_bump_script(dst_file)` is true and `dst_file.exists()`, synchronization skips the file unconditionally.

### 2.4 Boundary 4: Memory & Plans Protection (`.ai-memory/memory/` and `.ai-memory/plans/`)

- **Rule:** If `.ai-memory/memory/` or `.ai-memory/plans/` exist in a target repository, the synchronization pipeline must **NEVER modify, overwrite, delete, or mirror files into these directories**.
- **Architectural Rationale:** `.ai-memory/` stores local agent memory, task logs, completed execution plans, subtasks, and active transaction history. These records are strictly private to each repository's operational history. Mirroring upstream memory or plans into a child repository corrupts the child's local state and destroys audit trails.
- **Enforcement Mechanism:**
  - Evaluated via `is_protected_memory_or_plan(path: Path) -> bool`:
    ```python
    def is_protected_memory_or_plan(path: Path) -> bool:
        norm = str(path).replace("\\", "/").lower()

        if ".ai-memory/memory" in norm:
            return True

        if ".ai-memory/plans" in norm:
            return True

        for part in path.parts:
            if part in ("memory", "plans"):
                if ".ai-memory" in norm:
                    return True

        return False
    ```
  - If `is_protected_memory_or_plan(dst_file)` is true and `dst_file.exists()`, copy operations abort immediately with zero modifications.
  - Conditional guideline syncing only updates designated top-level index references: `.ai-memory/coding-guidelines.md` and `.ai-memory/prompts.md`.

---

## 3. Mandatory Pre-Pull Workflow Across All 42 Target Repositories

### 3.1 Pre-Flight Synchronization Mandate

Before any backup branches are created, before any tags are placed, and before any files are copied, the pipeline must execute a mandatory `git pull` across all 42 target repositories:

```bash
git checkout <base_branch>
git pull origin <base_branch> --no-rebase
```

### 3.2 Operational Failure Modes Prevented by Pre-Pull

1. **Non-Fast-Forward Push Rejections:** If remote CI/CD or another developer pushed a release commit to `origin/<base_branch>`, attempting to commit and push without pulling first triggers `[rejected - non-fast-forward]`.
2. **Tag Collision & Diverged HEADs:** Release tags generated against stale local HEADs point to outdated commits, stranding releases behind remote changes.
3. **Backup Branch Staleness:** Backup branches created from un-pulled local trees capture an inaccurate pre-change baseline that excludes recent remote contributions.

### 3.3 Pre-Pull Execution Contract

```
Step 1: Check Working Tree Cleanliness (`git status --porcelain`)
        │
        ├── If uncommitted changes exist ──► ABORT repo sync, log error
        │
Step 2: Detect Active Base Branch (`detect_base_branch`)
        │
        ├── Returns `main` or `master` (skips ephemeral feature/release branches)
        │
Step 3: Checkout Base Branch (`git checkout <base_branch>`)
        │
Step 4: Execute Fast-Forward Pull (`git pull origin <base_branch> --no-rebase`)
        │
        ├── If pull succeeds ──► Proceed to Backup & Pre-Release stage
        └── If pull fails (network/conflict) ──► Mark FAIL in summary, skip write
```

---

## 4. Target Repositories Directory (42 Connected Codebases)

The 42 target repositories represent the connected ecosystem managed under the parent workspace root:

| # | Repository Slug | Relative Workspace Path | Ecosystem / Tech Stack | Primary Branch | Role in Architecture |
| :- | :--- | :--- | :--- | :--- | :--- |
| 1 | `ai-empathy-prompt-tuner` | `02-prompts/ai-empathy-prompt-tuner` | Python / AI Prompts | `main` | Prompt Optimization Suite |
| 2 | `alim-cv` | `alim-cv` | HTML / Tailwind / TS | `main` | Professional Resume Portal |
| 3 | `alim-karim-profile` | `alim-karim-profile` | Web / React / Vite | `main` | Personal Portfolio |
| 4 | `alim.karim.profile` | `aukgit/alim.karim.profile` | Web / Hugo / Static | `main` | Mirror Profile Architecture |
| 5 | `antigravity-manager` | `antigravity-manager` | TypeScript / Electron | `main` | Antigravity AI Agent Manager |
| 6 | `cat-my` | `cat-my` | Web / TS / React | `main` | Catalog Management Portal |
| 7 | `core` | `03-aukgo/core` | Golang / Core Systems | `main` | High-Performance Go Utilities |
| 8 | `digital-name-card` | `digital-name-card` | React / Vite / CSS | `main` | Interactive Identity Card |
| 9 | `flat-slide-show` | `presentations-repos/flat-slide-show` | React / Slide System | `main` | Presentation Slide Engine |
| 10 | `gitlogger-new` | `gitlogger-new` | Go / CLI / Git Log | `main` | Git History Extraction CLI |
| 11 | `global-ppt-v1` | `presentations-repos/global-ppt-v1` | React / Slide Deck | `main` | Global Presentation Portal |
| 12 | `hiltrax` | `presentations-repos/hiltrax` | React / Design System | `main` | Enterprise Presentation Deck |
| 13 | `icon-coding-guidelines` | `icon-coding-guidelines` | SVG / Design Tokens | `main` | Iconography Asset Store |
| 14 | `img-pdf` | `img-pdf` | Python / Document PDF | `main` | Document & PDF Processing CLI |
| 15 | `ki-health-ppt` | `presentations-repos/ki-health-ppt` | React / Health Deck | `main` | Healthcare Presentation Deck |
| 16 | `kubernetes-training` | `aukgit/kubernetes-training` | K8s / Helm / Docs | `main` | Infrastructure Training Suite |
| 17 | `lara-licensing` | `lara-licensing` | PHP / Laravel / License | `main` | Software Licensing Service |
| 18 | `lara-publishing` | `lara-publishing` | PHP / Laravel / CMS | `main` | Publishing Automation Engine |
| 19 | `laravel-automation` | `laravel-automation` | PHP / Laravel / Tasks | `main` | Laravel Task Orchestrator |
| 20 | `letsmarknow-ui` | `letsmarknow-ui` | React / TS / Tailwind | `main` | Markdown Studio UI |
| 21 | `letsmarknow` | `letsmarknow` | Node / NestJS / Backend | `main` | Markdown Workspace Backend |
| 22 | `macro-ahk` | `macro-ahk` | AutoHotkey / Automation | `main` | Desktop Keyboard Macro Suite |
| 23 | `maid-app-spec-presentation` | `presentations-repos/maid-app-spec-presentation` | React / Presentation | `main` | Domain Spec Presentation Deck |
| 24 | `movie-cli` | `movie-cli` | Golang / SQLite / CLI | `main` | Movie Metadata CLI Tool |
| 25 | `pathhelper` | `03-aukgo/pathhelper` | Golang / Path Resolution | `main` | Cross-Platform Path Library |
| 26 | `presentation-aug-2026-plans-alim` | `presentations-repos/presentation-aug-2026-plans-alim` | React / Slides | `main` | Executive Planning Deck |
| 27 | `prompts-connect` | `02-prompts/prompts-connect` | Python / Prompt Sync | `main` | Prompt Connection Bus |
| 28 | `punam-case-studies-v1` | `punam-case-studies-v1` | React / Case Studies | `main` | Case Study Showcase |
| 29 | `rasia-logo` | `presentations-repos/rasia-logo` | SVG / Vector Branding | `main` | Vector Brand Asset System |
| 30 | `scripts-fixer` | `scripts-fixer` | Python / Tool Repair | `main` | Maintenance Script Fixer |
| 31 | `slides-spec` | `presentations-repos/slides-spec` | React / Slide Core | `main` | Slides Specification Engine |
| 32 | `spec-builder` | `spec-builder` | TypeScript / CLI / Docs | `main` | Spec Generator & Builder |
| 33 | `sweet-digs-finder` | `web-system/sweet-digs-finder` | Web / Real Estate Portal | `main` | Rental Search Platform |
| 34 | `ui-prompts-cat` | `ui-prompts-cat` | Web / Prompt Catalog | `main` | UI Prompt Explorer |
| 35 | `white-presentation-v1` | `presentations-repos/white-presentation-v1` | React / White Blue Theme | `main` | Light-Theme Slide Engine |
| 36 | `workflowy-ui` | `workflowy-ui` | React / Tree Outline UI | `main` | Outliner Frontend Interface |
| 37 | `workflowy` | `workflowy` | Node / Outliner Engine | `main` | Outliner Core Engine |
| 38 | `wp-exam` | `wp-exam` | PHP / WordPress / Quiz | `main` | WP Exam & Quiz Engine |
| 39 | `wp-git-log` | `wp-git-log` | PHP / WordPress / Audit | `main` | WP Git Activity Tracker |
| 40 | `wp-html-automate` | `wp-html-automate` | PHP / WordPress / HTML | `main` | WP HTML Processing Plugin |
| 41 | `wp-link-manager` | `wp-link-manager` | PHP / WordPress / Links | `main` | WP Link Governance Plugin |
| 42 | `wp-onboarding` | `wp-onboarding` | PHP / WordPress / User | `main` | WP Onboarding Automation |

> [!NOTE]
> The primary automation orchestrator repository `gitmap` (`gitmap-v28`) is also connected within the workspace. When synchronizing child repositories, `gitmap` functions either as the central orchestration driver or is synchronized alongside the fleet as repository 43.

---

## 5. Implementation Specification for `38-sync-prompts-skills-scripts.py`

### 5.1 Memory & Plans Protection Function

```python
def is_protected_memory_or_plan(path: Path) -> bool:
    """Check if directory/file belongs to .ai-memory/memory/ or .ai-memory/plans/ which must NEVER be overwritten."""
    norm = str(path).replace("\\", "/").lower()

    if ".ai-memory/memory" in norm:
        return True

    if ".ai-memory/plans" in norm:
        return True

    for part in path.parts:
        if part in ("memory", "plans"):
            if ".ai-memory" in norm:
                return True

    return False
```

### 5.2 Integration in `copy_single_file`

```python
def copy_single_file(
    src_file: Path,
    dst_file: Path,
    is_dry_run: bool = False,
    is_additive_only: bool = False,
) -> int:
    """Copy a single file respecting the 4 non-negotiable boundaries."""
    if not src_file.exists():
        return 0

    # Boundary 1: Spec 21 Exclusion
    if is_spec_21(src_file) or is_spec_21(dst_file):
        return 0

    # Boundary 4: Memory & Plans Protection
    if is_protected_memory_or_plan(dst_file):
        if dst_file.exists():
            return 0

    # Boundary 3: Bump Script Protection
    if is_bump_script(dst_file):
        if dst_file.exists():
            return 0

    # Boundary 2: Additive-Only AI Scripts
    if is_additive_only:
        if dst_file.exists():
            return 0

    is_copy_needed = False

    if not dst_file.exists():
        is_copy_needed = True
    else:
        try:
            if src_file.stat().st_size != dst_file.stat().st_size:
                is_copy_needed = True
            elif src_file.read_bytes() != dst_file.read_bytes():
                is_copy_needed = True
        except Exception:
            is_copy_needed = True

    if is_copy_needed:
        if not is_dry_run:
            dst_file.parent.mkdir(parents=True, exist_ok=True)
            shutil.copy2(src_file, dst_file)

        return 1

    return 0
```

### 5.3 Integration in `mirror_directory`

```python
def mirror_directory(
    src: Path,
    dst: Path,
    is_dry_run: bool = False,
    is_additive_only: bool = False,
) -> tuple[int, int]:
    copied = 0
    removed = 0

    if not src.exists():
        return 0, 0

    if is_spec_21(src) or is_spec_21(dst):
        return 0, 0

    if is_protected_memory_or_plan(dst):
        return 0, 0

    if not is_dry_run:
        dst.mkdir(parents=True, exist_ok=True)

    # 1. Clean stale files (skipped in additive mode or protected paths)
    if not is_additive_only:
        if dst.exists():
            for root, dirs, files in os.walk(dst, topdown=False):
                rel_root = Path(root).relative_to(dst)
                src_root = src / rel_root

                for f in files:
                    dst_file = Path(root) / f
                    src_file = src_root / f
                    is_excluded = (
                        f in EXCLUDE_NAMES
                        or dst_file.suffix.lower() in EXCLUDE_EXTS
                        or is_spec_21(dst_file)
                        or is_protected_memory_or_plan(dst_file)
                    )

                    if is_excluded:
                        continue

                    if not src_file.exists():
                        if not is_dry_run:
                            dst_file.unlink(missing_ok=True)
                        removed += 1

                for d in dirs:
                    dst_dir = Path(root) / d
                    src_dir = src_root / d

                    if d in EXCLUDE_NAMES or is_spec_21(dst_dir) or is_protected_memory_or_plan(dst_dir):
                        continue

                    if not src_dir.exists():
                        if not is_dry_run:
                            shutil.rmtree(dst_dir, ignore_errors=True)
                        removed += 1

    # 2. Copy files from src to dst
    for root, dirs, files in os.walk(src):
        dirs[:] = [
            d for d in dirs
            if d not in EXCLUDE_NAMES
            and not is_spec_21(Path(root) / d)
            and not is_protected_memory_or_plan(Path(root) / d)
        ]

        rel_root = Path(root).relative_to(src)
        target_dir = dst / rel_root

        if not is_dry_run:
            target_dir.mkdir(parents=True, exist_ok=True)

        for f in files:
            src_file = Path(root) / f
            is_excluded = (
                f in EXCLUDE_NAMES
                or Path(f).suffix.lower() in EXCLUDE_EXTS
                or is_spec_21(src_file)
                or is_protected_memory_or_plan(src_file)
            )

            if is_excluded:
                continue

            dst_file = target_dir / f

            if is_bump_script(dst_file) and dst_file.exists():
                continue

            if is_protected_memory_or_plan(dst_file) and dst_file.exists():
                continue

            copied += copy_single_file(
                src_file,
                dst_file,
                is_dry_run=is_dry_run,
                is_additive_only=is_additive_only,
            )

    return copied, removed
```

---

## 6. Dry-Run Verification Protocol & Execution Plan

### 6.1 Multi-Stage Execution Lifecycle

The complete execution lifecycle consists of five strict sequential stages:

```
┌─────────────────────────────────────────────────────────────────────────────┐
│ STAGE 1: Pre-Pull Sweep across 42 repositories                              │
│          Verify clean git status & fast-forward latest remote commits       │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
┌──────────────────────────────────────▼──────────────────────────────────────┐
│ STAGE 2: Dry-Run Verification (`--dry-run`)                                 │
│          Verify 0 boundary violations across all 42 targets                 │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
┌──────────────────────────────────────▼──────────────────────────────────────┐
│ STAGE 3: Single-Repo Smoke Test (`--repo movie-cli --dry-run`)              │
│          Inspect detailed diff, file counts, and boundary adherence         │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
┌──────────────────────────────────────▼──────────────────────────────────────┐
│ STAGE 4: Full Multi-Repo Synchronization (`--workers 6`)                    │
│          Backup branch -> Pre-release tag -> Mirror -> Commit -> Post-rel   │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
┌──────────────────────────────────────▼──────────────────────────────────────┐
│ STAGE 5: Post-Sync Verification Audit (`41-audit-all-repos.py`)             │
│          Confirm clean trees, valid tags, prompt parity, skills parity      │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 6.2 Pre-Flight Audit Checks

Before initiating live sync, verify:
1. `coding-guidelines` working tree is clean and on `main`.
2. All target repositories exist at their expected filesystem locations.
3. Python runner script compiles cleanly with zero syntax warnings.
4. SSH/Git credentials are authenticated for remote pushes (`origin`).

### 6.3 Rollback & Safety Guarantees

If an individual repository fails during live synchronization:
- The per-repository backup branch (`backup/sync-<timestamp>`) captures the exact pre-modification state.
- The pre-release tag (`v<pre-ver>`) preserves tag parity.
- Failed repositories do not interrupt the remaining pool of target repositories.
- Restoring any repository is a single atomic checkout: `git checkout backup/sync-<timestamp> && git branch -D <feature_branch>`.
