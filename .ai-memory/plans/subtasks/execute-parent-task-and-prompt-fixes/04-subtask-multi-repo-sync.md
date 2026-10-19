# Subtask 04: Multi-Repository Fleet Synchronization Protocol

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/execute-parent-task-and-prompt-fixes/04-subtask-multi-repo-sync.md`  
> **Parent Plan:** `.ai-memory/plans/execute-parent-task-and-prompt-fixes.md`  
> **Specification:** `02-spec/21-app/execute-parent-task-and-prompt-fixes/03-migration-and-sync-spec.md`  
> **State:** `COMPLETED`  
> **Assigned Role:** Subagent Worker 04 (Multi-Repo Fleet Sync Lead)  

---

## 1. Task Definition & Verification Contract

- **Task ID:** `04-subtask-multi-repo-sync`
- **State:** `[COMPLETED]`
- **Understood / AI Comprehension:** `[YES] — Fully understand the requirement to propagate canonical prompts, skills, shared specs (01-20), and additive AI scripts across all 43 connected repositories via 03-ai-scripts/38-sync-prompts-skills-scripts.py while strictly enforcing the 5 Non-Negotiable Boundaries and pre-change safety backup protocols.`
- **Azure AI / Model Runtime Context:** `Azure AI / Google Antigravity Swarm (Subagent Worker 04) via GitMap CLI Engine`
- **Intent & Technical Deliverables:**
  - Execute pre-run compilation and dry-run validation of `03-ai-scripts/38-sync-prompts-skills-scripts.py` across the 43 canonical repositories.
  - Execute live fleet synchronization, propagating upgraded V6 prompts, CG execute prompts, companion skills, and `AGENTS.md`.
  - Enforce automated safety backup branch creation (`backup/sync-<timestamp>`) and push for every touched repository.
  - Verify boundary protections across sample repositories (confirming untouched `02-spec/21-*` through `25-*`, preserved bump scripts, isolated `.ai-memory/plans/`, and zero credentials leaked).
- **Negative Boundaries & Invariants:**
  - MUST NOT synchronize or touch `02-spec/21-*` through `02-spec/25-*` in any target repository.
  - MUST NOT overwrite existing target scripts in `03-ai-scripts/` or `.agents/scripts/` (additive-only rule).
  - MUST NOT mutate version bump scripts (`bump*`, `bump_versions.py`, `version.json`).
  - MUST NOT touch target repository `.ai-memory/memory/` or `.ai-memory/plans/`.
  - MUST NOT stage, commit, or push non-owned repositories (`omis`, `oh-my-zsh`, `ohmyzsh`, `zsh`, `oh-my-posh`, `dotfiles`).
  - Worker subagents MUST NOT execute git commits directly outside the synchronization orchestrator.
- **Disjoint File Boundaries:**
  - `Owned Files (Exclusive):`
    - `.ai-memory/plans/subtasks/execute-parent-task-and-prompt-fixes/04-subtask-multi-repo-sync.md`
  - `Read-Only Context Files:`
    - `02-spec/21-app/execute-parent-task-and-prompt-fixes/03-migration-and-sync-spec.md`
    - `.ai-memory/plans/execute-parent-task-and-prompt-fixes.md`
    - `03-ai-scripts/38-sync-prompts-skills-scripts.py`
    - `AGENTS.md`
- **Verification Evidence Contract:**
  - Dry-run validation output confirming 43 valid repository targets and zero non-owned exclusions.
  - Synchronization run log reporting clean completion across connected repositories.
  - Boundary audit sampling at least 3 distinct target repositories (e.g. `gitmap`, `movie-cli`, `letsmarknow`) verifying zero boundary violations.

---

## 2. Detailed Execution Steps

### Step 1: Source Repository Pre-Flight Check
- [ ] Confirm that upstream subtasks (`01-subtask-preflight-and-v6-prompts`, `02-subtask-cg-execute-prompts-migration`, and `03-subtask-companion-skills-parity`) are marked `DONE`.
- [ ] Verify source repository syntax and path hygiene:
  ```bash
  python linter-scripts/check-relative-paths.py
  ```
- [ ] Verify compilation of the synchronization orchestrator:
  ```bash
  python -m py_compile 03-ai-scripts/38-sync-prompts-skills-scripts.py
  ```

### Step 2: Dry-Run Inspection & Scope Verification
- [ ] Execute dry-run mode (or inspect configured target repositories):
  ```bash
  python 03-ai-scripts/38-sync-prompts-skills-scripts.py --dry-run
  ```
- [ ] Confirm target registry validation:
  - 43 canonical repositories identified.
  - Excluded repositories list strictly enforces absence of `omis`, `ohmyzsh`, `zsh`, and system packages.
  - Shared specs list restricted to `02-spec/01-*` through `02-spec/20-*`.

### Step 3: Live Multi-Repository Fleet Synchronization
- [ ] Execute live fleet synchronization:
  ```bash
  python 03-ai-scripts/38-sync-prompts-skills-scripts.py
  ```
- [ ] Monitor execution log for per-repository ceremony:
  - Base branch pull (`git pull origin <base_branch> --no-rebase`).
  - Pre-change backup branch creation and remote push (`backup/sync-<timestamp>`).
  - Asset mirroring with boundary filtering.
  - Base branch commit and push with clean structured message.

### Step 4: Post-Sync Boundary Verification Audit
- [ ] Inspect git status and directory trees in sample target repositories (e.g. `gitmap`, `movie-cli`, `digital-name-card`):
  - Check Boundary 1: Verify `02-spec/21-*` through `02-spec/25-*` directories were NOT touched.
  - Check Boundary 2: Verify existing scripts in `03-ai-scripts/` were NOT overwritten.
  - Check Boundary 3: Verify version bump scripts (`bump*`) remain unmodified.
  - Check Boundary 4: Verify `.ai-memory/plans/` in target repositories remained completely untouched.
  - Check Boundary 5: Verify zero `.env` or credential files were created or staged.

### Step 5: GitMap Telemetry & State Closure
- [ ] Check fleet git status:
  ```bash
  gitmap status
  ```
- [ ] Update heatmap telemetry and transition subtask state from `QUEUED` to `DONE`:
  ```bash
  gitmap heatmap update --id 25 --outcome "43 connected repositories synchronized with 5 boundary protections verified" --met true
  ```

---

## 3. Negative Boundaries & Strict Bans

1. **TOTAL BAN on Absolute Paths:** Strictly relative Git paths only (`02-spec/...`, `01-prompts/...`). Zero `C:\...` or `file:///` URIs.
2. **TOTAL BAN on Mutating Non-Owned Repositories:** Never stage, commit, or push `omis`, `oh-my-zsh`, `ohmyzsh`, `zsh`, `oh-my-posh`, `dotfiles`.
3. **TOTAL BAN on Secrets Synchronization:** Zero `.env` files, tokens, or credentials synchronized across repositories.
4. **TOTAL BAN on Disabling CI/CD:** Target repository workflows must remain fully active and unmodified.
