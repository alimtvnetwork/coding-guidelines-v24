# Acceptance Criteria: Multi-Repository Skills Synchronization & Non-Destructive Fleet Deployment

> Category: `21-app`  
> Spec ID: `02-spec/21-app/17-multi-repo-skills-sync/02-acceptance-criteria.md`  
> Status: Active  
> Architecture: Non-Destructive Multi-Repository Fleet Synchronization Engine  
> Target Subsystems: `03-ai-scripts/38-sync-prompts-skills-scripts.py`, `.agents/skills/sync/`, `.agents/skills/muse-master-prompt/`, `.agents/skills/gitmap/`, `01-prompts/27-muse-prompts/`, and all 43 registered repositories  
> Companion Spec: `02-spec/21-app/17-multi-repo-skills-sync/01-overview-and-architecture.md`  

---

## 1. Acceptance Criteria Register

### AC-SYNC-001: Comprehensive Fleet Coverage (43 Repositories)
- **Given** The synchronization script `03-ai-scripts/38-sync-prompts-skills-scripts.py` is executed across the default workspace target.
- **When** Iterating through all registered repositories.
- **Then** Every one of the 43 registered repositories is visited, processed, synchronized, and verified without skipping any active repository.

### AC-SYNC-002: Zero Deletions of Repo-Owned Files (Strict Non-Destructive Invariant)
- **Given** Target repositories contain custom files, domain tools, local test suites, or files that do not exist in `coding-guidelines`.
- **When** Canonical assets (`01-prompts/`, `.agents/skills/`, `.cursor/skills/`, `02-spec/01-*` to `20-*`, `03-ai-scripts/`, `.agents/scripts/`) are mirrored.
- **Then** The synchronization engine performs ZERO unlinks (`dst_file.unlink()`) and ZERO directory tree deletions (`shutil.rmtree()`) on repo-owned files. The removal counter is strictly 0 (`removed == 0`), preserving all local files intact.

### AC-SYNC-003: Muse Master Prompt & Skills Synchronization
- **Given** `01-prompts/27-muse-prompts/01-muse-master-prompt.md`, `.agents/skills/muse-master-prompt/`, and `.cursor/skills/muse-master-prompt/` are updated in `coding-guidelines`.
- **When** Fleet synchronization runs.
- **Then** All 43 target repositories receive the latest unified Muse master prompt and skill files, ensuring uniform V6 autonomous execution and onboarding capabilities across all codebases.

### AC-SYNC-004: GitMap LLM Training Skills Synchronization
- **Given** `.agents/skills/gitmap/` and `.cursor/skills/gitmap/` include updated instructions for `gitmap llm train muse --days 7`, Split-DB architecture, and SSH cluster delegation.
- **When** Fleet synchronization runs.
- **Then** The updated GitMap skills are deployed to all target repositories without removing any repository-specific custom skills that may exist in `.agents/skills/` or `.cursor/skills/`.

### AC-SYNC-005: Shared Specifications Synchronization (`02-spec/01-*` to `02-spec/20-*`)
- **Given** Shared specifications `02-spec/01-*` through `02-spec/20-*` define foundational coding guidelines, error management, and architecture rules.
- **When** Fleet synchronization runs.
- **Then** All shared specifications from 01 through 20 are mirrored to target repositories, while `02-spec/21-*` through `02-spec/25-*` remain strictly excluded.

### AC-SYNC-006: Non-Negotiable Boundaries Protection
- **Given** Target repositories contain domain application specs, bump scripts, and operational memory.
- **When** File copying decisions are evaluated by `copy_single_file()` and `mirror_directory()`.
- **Then**:
  1. `02-spec/21-*` through `02-spec/25-*` are never touched or written.
  2. Bump scripts (`bump*`, `bump-version.mjs`, `version.json`) are preserved without overwriting.
  3. `.ai-memory/memory/` and `.ai-memory/plans/` are protected from overwrite or mutation.
  4. `06-archive/` and `06-old-prompts/` are excluded from synchronization.

### AC-SYNC-007: Pre-Flight Pull and Safety Backup Branch Creation
- **Given** A target repository is about to receive updates.
- **When** The pre-change protocol executes.
- **Then**:
  1. The base branch is detected and pulled (`git pull origin <base_branch> --no-rebase`).
  2. A dedicated timestamped backup branch is created and pushed upstream (`git checkout -b backup/sync-<timestamp>` -> `git push -u origin backup/sync-<timestamp>`).
  3. The repository returns to `<base_branch>` before applying file changes.

### AC-SYNC-008: Atomic Commits with Hyphen Message Formatting
- **Given** Canonical assets have been mirrored into a target repository.
- **When** Changes are detected via `git status --porcelain`.
- **Then** Changes are staged (`git add -A`) and committed atomically using the hyphen-separated message format `feat(sync) - sync canonical prompts, skills, gitmap agent standards, and purge deprecated 46 task manager` without colons in the subject argument.

### AC-SYNC-009: Post-Change Release Ceremony, Tagging, and Version Bump
- **Given** A target repository has committed synchronization changes.
- **When** The post-change release ceremony runs.
- **Then**:
  1. A patch version bump is calculated (`bump_patch_version(current_ver)`).
  2. A release branch `release/v<next_ver>` is created.
  3. Version manifests (`version.json`, `package.json`, or via bump scripts) are updated.
  4. An annotated Git tag `v<next_ver>` is created and pushed to origin (`git push origin v<next_ver>`).
  5. The release branch is merged back into the base branch (`git merge release/v<next_ver> -m "chore(release) - merge v<next_ver> [skip ci]"`).

### AC-SYNC-010: Upstream Push & Clean Working Trees ('No Push = Not Done')
- **Given** Synchronization and release operations are complete.
- **When** Inspecting the final repository status.
- **Then**:
  1. All local commits, release branches, and release tags are pushed to remote origin (`git push origin <base_branch>`).
  2. The working tree is completely clean (`git status --porcelain` returns empty).
  3. Zero uncommitted or unpushed changes remain.

---

## 2. Fleet Acceptance Criteria Matrix (All 43 Repositories)

The following matrix defines the mandatory verification criteria for all 43 registered repositories:

| # | Repository Identifier | Base Branch | Pre-Flight Pull | Safety Backup Branch | Mirrored Assets | Repo Deletions | Commit Format | Post Tag | Upstream Pushed | Working Tree |
| :---: | :--- | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| **01** | `02-prompts/ai-empathy-prompt-tuner` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **02** | `alim-cv` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **03** | `alim-karim-profile` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **04** | `aukgit/alim.karim.profile` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **05** | `antigravity-manager` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **06** | `cat-my` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **07** | `03-aukgo/core` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **08** | `digital-name-card` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **09** | `presentations-repos/flat-slide-show` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **10** | `gitlogger-new` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **11** | `gitmap` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **12** | `presentations-repos/global-ppt-v1` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **13** | `presentations-repos/hiltrax` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **14** | `icon-coding-guidelines` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **15** | `img-pdf` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **16** | `presentations-repos/ki-health-ppt` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **17** | `aukgit/kubernetes-training` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **18** | `lara-licensing` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **19** | `lara-publishing` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **20** | `laravel-automation` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **21** | `letsmarknow-ui` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **22** | `letsmarknow` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **23** | `macro-ahk` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **24** | `presentations-repos/maid-app-spec-presentation` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **25** | `movie-cli` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **26** | `03-aukgo/pathhelper` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **27** | `presentations-repos/presentation-aug-2026-plans-alim` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **28** | `02-prompts/prompts-connect` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **29** | `punam-case-studies-v1` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **30** | `presentations-repos/rasia-logo` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **31** | `scripts-fixer` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **32** | `presentations-repos/slides-spec` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **33** | `spec-builder` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **34** | `web-system/sweet-digs-finder` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **35** | `ui-prompts-cat` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **36** | `presentations-repos/white-presentation-v1` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **37** | `workflowy-ui` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **38** | `workflowy` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **39** | `wp-exam` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **40** | `wp-git-log` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **41** | `wp-html-automate` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **42** | `wp-link-manager` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |
| **43** | `wp-onboarding` | `main` | Required | `backup/sync-*` | Prompts, Skills, Specs | **0** | `feat(sync) - ...` | `v*` | Yes | Clean |

---

## 3. Boundary Verification Checklist

- [ ] **B-01 (Spec 21 Exclusion):** Source scan ignores all `02-spec/21-*` through `02-spec/25-*` directories.
- [ ] **B-02 (Additive AI Scripts):** `03-ai-scripts/` and `.agents/scripts/` files present in target repository with custom commit messages are never overwritten.
- [ ] **B-03 (Bump Script Protection):** Any file matching `bump*`, `bump_versions.py`, `bump-version.mjs`, or `version.json` in the target repository is untouched.
- [ ] **B-04 (Memory & Plans Protection):** Target repositories retain 100% of `.ai-memory/memory/` and `.ai-memory/plans/` without alteration.
- [ ] **B-05 (Archive Exclusion):** `06-archive/` and `06-old-prompts/` are excluded from the sync payload.
- [ ] **B-06 (Strict Non-Destructive Invariant):** Zero files deleted from target repositories (`removed == 0`).
- [ ] **B-07 (Relative Git Paths):** All specifications, scripts, and logs use strictly relative paths starting from the repository root.
- [ ] **B-08 (Upstream Completion):** Changes are committed, tagged, and pushed to remote origin across all 43 repositories.
