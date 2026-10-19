# Subtask 05: Multi-Repository Fleet Synchronization Protocol

## 1. Subtask Overview

| Attribute | Details |
| :--- | :--- |
| **Subtask ID** | `05-subtask-multi-repo-sync` |
| **Parent Task** | `gitmap-commands-agent-spawning-heatmap-sync` |
| **Specification Ref** | `02-spec/21-app/gitmap-commands-agent-spawning-heatmap-sync/03-heatmap-and-sync-spec.md` |
| **Owned Files** | `.ai-memory/plans/subtasks/gitmap-commands-agent-spawning-heatmap-sync/05-subtask-multi-repo-sync.md` |
| **Status** | `COMPLETED` |

---

## 2. Objective & Scope

Synchronize all canonical assets updated in `coding-guidelines`—specifically `AGENTS.md`, updated prompt templates in `01-prompts/`, Antigravity and Cursor skills in `.agents/skills/` and `.cursor/skills/`, shared specifications in `02-spec/01-*` through `02-spec/20-*`, and additive AI scripts in `03-ai-scripts/` and `.agents/scripts/`—across all 43 connected repositories in the workspace using `03-ai-scripts/38-sync-prompts-skills-scripts.py`.

Strict adherence to the 5 Non-Negotiable Boundaries is mandatory to ensure private application specifications, bump scripts, memory/plans directories, and secrets remain untouched.

---

## 3. The 5 Non-Negotiable Boundaries Verification Matrix

| Boundary | Rule | Verification Check |
| :--- | :--- | :--- |
| **1. Spec 21 Exclusion** | NEVER sync `02-spec/21-*` through `02-spec/25-*` | Confirm script excludes private app specs |
| **2. Additive AI Scripts** | Add new scripts only; never overwrite modified target scripts | Verify additive-only flag on `03-ai-scripts` and `.agents/scripts` |
| **3. Bump Script Protection** | NEVER overwrite version bump scripts (`bump*`) | Ensure immutable bump script protection remains active |
| **4. Memory & Plans Isolation** | NEVER touch `.ai-memory/memory/` or `.ai-memory/plans/` | Verify total exclusion of target repo `.ai-memory/` |
| **5. Zero Secrets Leakage** | NEVER synchronize `.env` or credential files | Offload all secrets strictly to `repo-secrets` via `gitmap rs` |

---

## 4. Detailed Execution Steps

### Step 1: Pre-Flight Source Validation
- [ ] Verify that source repository `coding-guidelines` is clean of uncommitted changes.
- [ ] Confirm all local linters and syntax checks pass:
  ```bash
  python 03-ai-scripts/07-relative-paths-linter.py
  ```
- [ ] Verify that `03-ai-scripts/38-sync-prompts-skills-scripts.py` has executable permissions and valid syntax:
  ```bash
  python -m py_compile 03-ai-scripts/38-sync-prompts-skills-scripts.py
  ```

### Step 2: Dry-Run Verification
- [ ] Execute `python 03-ai-scripts/38-sync-prompts-skills-scripts.py --dry-run` (or inspect target repo list) to confirm target paths:
  - 43 canonical repositories identified in workspace.
  - Non-owned repositories (`omis`, `oh-my-zsh`, `dotfiles`) strictly excluded.
  - Shared specs list restricted to `01-spec-authoring-guide` through `20-*`.

### Step 3: Live Multi-Repository Fleet Synchronization
- [ ] Run the fleet synchronization script across all 43 connected repositories:
  ```bash
  python 03-ai-scripts/38-sync-prompts-skills-scripts.py
  ```
- [ ] Verify backup branch creation per target repository:
  - Automatic creation and push of `backup/sync-<timestamp>`.
  - Base branch checkout preserved.
  - Canonical assets mirrored cleanly.

### Step 4: Boundary & Git Hygiene Audit
- [ ] Sample check at least 3 distinct target repositories (e.g. `gitmap`, `movie-cli`, `punam-case-studies-v1`):
  - Confirm `AGENTS.md` reflects updated Section 12, 15, and heatmap rules.
  - Confirm `02-spec/21-*` is completely untouched in target repositories.
  - Confirm target repository `.ai-memory/plans/` is intact without contamination.
  - Confirm bump scripts (`bump_versions.py` or `bump.ps1`) were NOT modified.

### Step 5: Post-Sync Status Telemetry
- [ ] Inspect git status and pipeline status across repositories:
  ```bash
  gitmap status
  ```
- [ ] Log completion in GitMap heatmap telemetry:
  ```bash
  gitmap heatmap log --task "gitmap-commands-agent-spawning-heatmap-sync" --action ANALYZE --file "03-ai-scripts/38-sync-prompts-skills-scripts.py" --state EXECUTION --expected "Sync 43 repositories"
  gitmap heatmap update --id <id> --outcome "43 repos synced with 5 boundary protections verified" --met true
  ```

---

## 5. Negative Boundaries & Strict Bans

1. **TOTAL BAN on Absolute Paths:** No absolute filesystem paths (`C:\...`, `/home/...`) or `file:///` URIs in any synced files or sync logs.
2. **TOTAL BAN on Mutating Non-Owned Repositories:** Never touch or stage `omis`, `ohmyzsh`, `zsh`, `oh-my-posh`, `dotfiles`, or runtime package directories.
3. **TOTAL BAN on Secrets Synchronization:** Zero `.env` files or API keys synchronized.
4. **TOTAL BAN on Disabling CI/CD:** Workflows and pipeline actions in target repositories must remain fully enabled and unbypassed.
