# Specification: Multi-Repository Skills Synchronization & Non-Destructive Fleet Deployment

> Category: `21-app`  
> Spec ID: `02-spec/21-app/17-multi-repo-skills-sync/01-overview-and-architecture.md`  
> Status: Active  
> Architecture: Non-Destructive Multi-Repository Fleet Synchronization Engine  
> Target Subsystems: `03-ai-scripts/38-sync-prompts-skills-scripts.py`, `.agents/skills/sync/`, `.agents/skills/muse-master-prompt/`, `.agents/skills/gitmap/`, `01-prompts/27-muse-prompts/`, and all 43 registered repositories  
> Companion Spec: `02-spec/21-app/17-multi-repo-skills-sync/02-acceptance-criteria.md`  

---

## 1. User Directive & Non-Negotiable Invariant

```text
Under NO circumstances should any additional file created or updated by the target repository on its own be removed. The sync engine must be strictly non-destructive for repo-owned files: only update/replace canonical assets and add new canonical assets. Never delete stale files in target repositories.
```

### 1.1 Non-Negotiable Invariants

1. **Strict Non-Destructive Invariant for Repo-Owned Files:**  
   Target repositories frequently contain custom tools, domain scripts, bespoke test suites, local configurations, or custom documentation created directly inside target trees. The synchronization engine MUST NEVER delete, unlink, purge, or prune any file or folder created or maintained by the target repository.
2. **Canonical Assets Update & Addition Only:**  
   The synchronization engine ONLY updates or replaces canonical assets originating from `coding-guidelines`, and adds newly created canonical assets. Stale files in target directories are left intact.
3. **The 5 Non-Negotiable Boundaries:**
   - **Boundary 1 (Spec 21 Exclusion):** NEVER synchronize, copy, or touch `02-spec/21-*` through `02-spec/25-*` (private application domain specs, issues, db schemas, and UI design systems). Only shared specifications `02-spec/01-*` through `02-spec/20-*` are synchronized.
   - **Boundary 2 (Additive-Only AI Scripts):** Brand-new AI scripts (`03-ai-scripts/`, `.agents/scripts/`) that do not exist in the target repository are copied cleanly. Existing scripts modified by the target repository must NEVER be overwritten.
   - **Boundary 3 (Bump Script Protection):** NEVER overwrite version bump scripts or manifests (`bump*`, `bump_versions.py`, `bump-version.mjs`, `version.json`). Each repository maintains custom SemVer targets and release logic.
   - **Boundary 4 (Memory & Plans Protection):** NEVER modify, sync, or mirror `.ai-memory/memory/` or `.ai-memory/plans/` in target repositories. Target repositories own their operational memory logs and execution plans.
   - **Boundary 5 (Archive Exclusion):** NEVER synchronize `06-archive/` or `06-old-prompts/` downstream.
4. **Pre-Change Safety Protocol (Pull & Backup):**  
   Every target repository must pull latest changes on its base branch (`git pull origin <base_branch> --no-rebase`), create and push a remote safety backup branch (`backup/sync-<timestamp>`), and confirm pre-release state before modifying files.
5. **Mandatory Completion Invariant ('No Push = Not Done'):**  
   All synchronized changes must be committed atomically (`feat(sync) - ...`), tagged with post-change release versions, and pushed upstream to GitHub. Uncommitted changes or unpushed branches constitute an incomplete task.
6. **Strict Relative Git Paths Mandate:**  
   All paths in specifications, scripts, commits, release notes, and documentation must be strictly relative Git paths starting from the repository root (e.g., `02-spec/21-app/...`, `03-ai-scripts/...`). Absolute filesystem paths and `file:///` URIs are totally banned.

---

## 2. Fleet Overview & Synchronization Scope

The fleet comprises 43 registered repositories connected to the central prompt and skill architecture. The synchronization operation distributes updated canonical assets across the entire fleet to ensure consistent agent capabilities, LLM training protocols, coding guidelines, and tool automation.

```
+-----------------------------------------------------------------------------------+
|                        Source Repository: coding-guidelines                       |
+-----------------------------------------------------------------------------------+
  |
  |-- 01-prompts/27-muse-prompts/ (Unified Muse Master Prompt V6)
  |-- .agents/skills/muse-master-prompt/ & .cursor/skills/muse-master-prompt/
  |-- .agents/skills/gitmap/ & .cursor/skills/gitmap/ (LLM Training, Split-DB, SSH)
  |-- 02-spec/01-* through 02-spec/20-* (Shared Architectural & Coding Specs)
  |-- 03-ai-scripts/ & .agents/scripts/ (Additive Automation Scripts)
  |
  +----------> [Non-Destructive Synchronization Engine: 38-sync-prompts-skills-scripts.py]
                 |
                 |-- (1) git pull origin <base_branch> --no-rebase
                 |-- (2) git checkout -b backup/sync-<timestamp> && git push
                 |-- (3) git checkout <base_branch>
                 |-- (4) ensure_pre_release (tag & branch)
                 |-- (5) Non-destructive asset mirror (0 deletions of repo files)
                 |-- (6) git commit -m "feat(sync) - ..." && git push
                 |-- (7) Post-change release ceremony (bump version, tag, merge, push)
                 v
+-----------------------------------------------------------------------------------+
|                     Fleet of 43 Registered Target Repositories                     |
+-----------------------------------------------------------------------------------+
```

### 2.1 Core Asset Payloads Synchronized

1. **Unified Muse Master Prompt & Skills (`01-prompts/27-muse-prompts/` & `muse-master-prompt`):**  
   Consolidated V6 autonomous execution prompt merging onboarding, SSH authentication fast-path, GitMap installation and update, `gitmap scan`, fleet clone/pull (`gitmap pa`), conflict prevention, side-chat architecture (`<repo> repo` + `all summary`), 15MB 7-Zip split rule, V6 task confirmation protocol, and Muse Memory IO re-sync.
2. **GitMap LLM Training & Autonomous Agent Skills (`gitmap`):**  
   Updated GitMap skills integrating `gitmap llm train muse --days 7`, three-tier SQLite Split-DB architecture (`gitmap.db`, `installation.db`, `repodb/pipeline.db`), SSH cluster delegation, and zero-storage CI telemetry (`gitmap pe -t`).
3. **Shared Specifications (`02-spec/01-*` through `02-spec/20-*`):**  
   Enterprise coding guidelines, error management (`*appfault.AppError`), boolean principles (positive-only, implicit evaluations, no mixed polarity), DRY library extraction, database split-db schemas, and spec authoring guides.
4. **Additive AI Scripts (`03-ai-scripts/` & `.agents/scripts/`):**  
   Cross-platform Python automation scripts, test inventory runners, linter tools, and CI diagnostic helpers, copied in additive mode to protect repo modifications.

---

## 3. Registered Fleet Inventory (43 Repositories)

The 43 registered target repositories span prompt libraries, core packages, CLI tools, web platforms, and slide systems:

| # | Repository Identifier | Relative Workspace Path | Primary Domain |
| :---: | :--- | :--- | :--- |
| **01** | `ai-empathy-prompt-tuner` | `02-prompts/ai-empathy-prompt-tuner` | Prompt Engineering |
| **02** | `alim-cv` | `alim-cv` | Career & CV Platform |
| **03** | `alim-karim-profile` | `alim-karim-profile` | Identity & Profile |
| **04** | `aukgit/alim.karim.profile` | `aukgit/alim.karim.profile` | Profile Infrastructure |
| **05** | `antigravity-manager` | `antigravity-manager` | AGY Tooling & Management |
| **06** | `cat-my` | `cat-my` | Content Aggregation Tool |
| **07** | `03-aukgo/core` | `03-aukgo/core` | Go Enterprise Core Package |
| **08** | `digital-name-card` | `digital-name-card` | Digital Identity System |
| **09** | `flat-slide-show` | `presentations-repos/flat-slide-show` | Slide Deck & Presentations |
| **10** | `gitlogger-new` | `gitlogger-new` | Git Activity Logging & Telemetry |
| **11** | `gitmap` | `gitmap` | GitMap CLI, AUM & Fleet Engine |
| **12** | `global-ppt-v1` | `presentations-repos/global-ppt-v1` | Global Slide Decks |
| **13** | `hiltrax` | `presentations-repos/hiltrax` | Presentation Framework |
| **14** | `icon-coding-guidelines` | `icon-coding-guidelines` | Design & Icon Guidelines |
| **15** | `img-pdf` | `img-pdf` | Document & PDF Processing |
| **16** | `ki-health-ppt` | `presentations-repos/ki-health-ppt` | Health Domain Slide Deck |
| **17** | `aukgit/kubernetes-training` | `aukgit/kubernetes-training` | Cloud & K8s Architecture |
| **18** | `lara-licensing` | `lara-licensing` | Licensing Architecture |
| **19** | `lara-publishing` | `lara-publishing` | Content Publishing Platform |
| **20** | `laravel-automation` | `laravel-automation` | Backend Automation Framework |
| **21** | `letsmarknow-ui` | `letsmarknow-ui` | Markdown UI Frontend |
| **22** | `letsmarknow` | `letsmarknow` | Markdown Engine Backend |
| **23** | `macro-ahk` | `macro-ahk` | Windows Automation Macros |
| **24** | `maid-app-spec-presentation` | `presentations-repos/maid-app-spec-presentation` | Presentation Specifications |
| **25** | `movie-cli` | `movie-cli` | Media Management CLI |
| **26** | `03-aukgo/pathhelper` | `03-aukgo/pathhelper` | Cross-Platform Path Library |
| **27** | `presentation-aug-2026-plans-alim` | `presentations-repos/presentation-aug-2026-plans-alim` | Strategic Roadmap Slides |
| **28** | `prompts-connect` | `02-prompts/prompts-connect` | Prompt Gateway Architecture |
| **29** | `punam-case-studies-v1` | `punam-case-studies-v1` | Case Studies & Content |
| **30** | `rasia-logo` | `presentations-repos/rasia-logo` | Design System & Branding |
| **31** | `scripts-fixer` | `scripts-fixer` | Script Remediation Engine |
| **32** | `slides-spec` | `presentations-repos/slides-spec` | Slides Specification Engine |
| **33** | `spec-builder` | `spec-builder` | Automated Spec Synthesis |
| **34** | `sweet-digs-finder` | `web-system/sweet-digs-finder` | Property Search Engine |
| **35** | `ui-prompts-cat` | `ui-prompts-cat` | UI Prompt Categorization |
| **36** | `white-presentation-v1` | `presentations-repos/white-presentation-v1` | High-Authority White Presentations |
| **37** | `workflowy-ui` | `workflowy-ui` | Outline UI Interface |
| **38** | `workflowy` | `workflowy` | Outline Architecture Core |
| **39** | `wp-exam` | `wp-exam` | WordPress Assessment System |
| **40** | `wp-git-log` | `wp-git-log` | WordPress Git Log Plugin |
| **41** | `wp-html-automate` | `wp-html-automate` | HTML Automation Framework |
| **42** | `wp-link-manager` | `wp-link-manager` | Link Management Subsystem |
| **43** | `wp-onboarding` | `wp-onboarding` | WordPress Onboarding Workflow |

---

## 4. Architecture of `03-ai-scripts/38-sync-prompts-skills-scripts.py`

The synchronization script `03-ai-scripts/38-sync-prompts-skills-scripts.py` implements the non-destructive multi-repository release ceremony.

### 4.1 Seven-Stage Per-Repository Execution Flow

For each repository in the fleet, the script executes 7 sequential stages:

```
[Stage 1: Pre-Flight Pull]
  git checkout <base_branch>
  git pull origin <base_branch> --no-rebase
       |
[Stage 2: Safety Backup Branch]
  git checkout -b backup/sync-<timestamp>
  git push -u origin backup/sync-<timestamp>
  git checkout <base_branch>
       |
[Stage 3: Pre-Release Baseline Tag]
  ensure_pre_release -> v<pre_ver> tag + branch push
       |
[Stage 4: Non-Destructive Asset Mirroring]
  - Copy new & updated canonical assets
  - ZERO deletions of repo-owned files
  - Apply Boundaries 1, 2, 3, 4, 5
       |
[Stage 5: Status Check & Atomic Commit]
  git status --porcelain
  git add -A
  git commit -m "feat(sync) - ..."
  git push -u origin <base_branch>
       |
[Stage 6: Post-Change Release Ceremony]
  Bump patch version -> release/v<next_ver>
  git tag -a v<next_ver> -m "Release v<next_ver> - ..."
  git push -u origin release/v<next_ver>
  git push origin v<next_ver>
       |
[Stage 7: Base Branch Merge & Final Upstream Push]
  git checkout <base_branch>
  git merge release/v<next_ver> -m "chore(release) - merge v<next_ver> [skip ci]"
  git push origin <base_branch>
```

### 4.2 Non-Destructive Mirroring Architecture

The critical architectural difference in the non-destructive sync engine is the removal of stale file deletion:

```python
def mirror_directory(
    src: Path,
    dst: Path,
    repo_root: Path | None = None,
    is_dry_run: bool = False,
    is_additive_only: bool = False,
) -> tuple[int, int]:
    """Mirror src into dst by updating existing canonical files and adding new ones.
    
    CRITICAL NON-DESTRUCTIVE INVARIANT:
    Stale files in dst that do not exist in src are NEVER deleted.
    Repo-owned files, local modifications, and bespoke tools in target
    repositories are preserved with zero unlinks and zero rmtree calls.
    """
    copied = 0
    removed = 0  # Strictly 0 in non-destructive mode

    if not src.exists():
        return 0, 0

    if is_spec_21(src) or is_spec_21(dst):
        return 0, 0

    if is_protected_memory_or_plan(dst):
        if dst.exists():
            return 0, 0

    if not is_dry_run:
        dst.mkdir(parents=True, exist_ok=True)

    # COPY & UPDATE ONLY: Traverse src and copy to dst
    for root, dirs, files in os.walk(src):
        # Filter excluded directories (spec 21, memory/plans, pycache, etc.)
        ...
        for f in files:
            src_file = Path(root) / f
            dst_file = target_dir / f
            copied += copy_single_file(src_file, dst_file, repo_root, is_dry_run, is_additive_only)

    return copied, removed
```

### 4.3 Additive AI Scripts Safeguard

When copying files in `03-ai-scripts/` and `.agents/scripts/`:
1. If the script does not exist in the target repository, it is copied as a new file.
2. If the script exists and has local modifications committed by the target repository (commit message does not contain `sync`), the file is preserved and NEVER overwritten.
3. If the script exists and was previously synced without target repo modification, it is refreshed to the latest canonical version.

### 4.4 Boundary Enforcement Summary

| Boundary | Target Path Pattern | Action on Collision |
| :--- | :--- | :--- |
| **Spec 21 Exclusion** | `02-spec/21-*` through `02-spec/25-*` | Never read from source; never write to target |
| **Additive AI Scripts** | `03-ai-scripts/`, `.agents/scripts/` | Add new; preserve target-modified scripts |
| **Bump Script Protection** | `bump*`, `bump-version.*`, `version.json` | Skip copying if target file already exists |
| **Memory/Plans Protection** | `.ai-memory/memory/`, `.ai-memory/plans/` | Never touch, overwrite, or delete in target |
| **Archive Exclusion** | `06-archive/`, `06-old-prompts/` | Skip source; purge legacy folders in target |
| **Non-Destructive Invariant** | Any target repository file | ZERO deletion of repo-owned files |

---

## 5. Concurrency & Execution Controls

- **Default Workers:** 6 parallel worker threads via `concurrent.futures.ThreadPoolExecutor`.
- **Target Filtering:** Run on a single repository via `--repo <name>` or all 43 repositories by default.
- **Dry-Run Capability:** Preview all additions and updates with `--dry-run` to verify 0 deletions before applying changes.
- **Local-Only Testing:** Skip upstream pushes using `--no-push` for local validation.
