# Architecture Spec: Multi-Repo Sync Date Format and Branch Pruning

## 1. Overview
The synchronization engine `03-ai-scripts/38-sync-prompts-skills-scripts.py` orchestrates multi-repository asset distribution across connected repositories. This specification details the standardized date representation and automated lifecycle pruning of historical sync branches.

---

## 2. Date Formatting Standard

### 2.1 Display and Log Format (`dd-MMM-yy 12h am/pm`)
All human-readable timestamps, summary logs, commit annotations, and release tags must adhere to:
- **Format**: `dd-MMM-yy hh:mm a`
- **Example**: `28-Oct-26 10:11 am` (or `pm`)
- **Tokens**:
  - `dd`: 2-digit zero-padded day (`%d`).
  - `MMM`: 3-character Titlecase English month (`%b`).
  - `yy`: 2-digit year (`%y`).
  - `hh`: 2-digit 12-hour clock (`%I`).
  - `mm`: 2-digit minute (`%M`).
  - `a`: lowercase 12-hour period indicator (`am` or `pm`).

### 2.2 Branch-Safe Ref Name Format
Git ref formatting mandates zero spaces, colons, or illegal characters in ref names. The branch-safe name uses hyphen separation:
- **Format**: `backup/sync-dd-MMM-yy-hh-mma`
- **Example**: `backup/sync-28-Oct-26-10-11am`
- **Collision Protection**: If `backup/sync-28-Oct-26-10-11am` already exists, append seconds `-%S` or counter `-01`.

---

## 3. Older Sync Branch Pruning Engine

### 3.1 Eligibility Invariants
A branch qualifies for automatic pruning during sync if and only if all of the following conditions are met:
1. **Namespace & Pattern**:
   - Local: `refs/heads/backup/*` with `sync` in the branch name, or `refs/heads/feat/sync*`.
   - Remote: `refs/remotes/origin/backup/*` with `sync` in the branch name, or `refs/remotes/origin/feat/sync*`.
2. **Age Threshold**: The commit timestamp (`%(committerdate:unix)`) is strictly greater than 10 days old (`now - committerdate > 10 * 86400 seconds`).
3. **Safety Exclusions**:
   - Never prune the current checked out branch (`git branch --show-current`).
   - Never prune `main`, `master`, `dev`, `release/*`, or non-sync feature/backup branches.

### 3.2 Pruning Mechanics
- **Local Branch Deletion**:
  ```bash
  git branch -D <local_branch_name>
  ```
- **Remote Branch Deletion**:
  ```bash
  git push origin --delete <remote_branch_name>
  ```
- **Dry-Run & Push-Disabled Awareness**:
  - Under `--dry-run`: Log identified branches without executing deletion.
  - Under `--no-push`: Delete local eligible branches only; skip remote branch deletion.
- **Resilience**: Branch deletion errors (e.g., branch already removed remotely) must be logged and handled gracefully without terminating the sync flow.

---

## 4. Execution Sequence in `sync_repo`

1. **Checkout & Pull**: Switch to `base_branch`, run `git pull origin <base_branch> --no-rebase`.
2. **Prune Stale Branches**: Execute `prune_older_sync_branches(target, max_age_days=10)`.
3. **Daily Sync Guard Check**: Check for prior same-day syncs unless `--force` is specified.
4. **Create Pre-Change Backup**: Create `backup/sync-{format_sync_branch_timestamp()}` and push upstream.
5. **Release Tag Pre-Flight**: Ensure pre-change release tag exists.
6. **Mirror Canonical Assets**: Copy prompts, skills, specs (01-20), and additive AI scripts.
7. **Commit & Post-Change Release Ceremony**: Tag and release with updated date annotations.
