# Subtask 05: Compliance Verification, Linter Audit & Atomic Commit Push

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/execute-parent-task-and-prompt-fixes/05-subtask-verification-and-atomic-commit.md`  
> **Parent Plan:** `.ai-memory/plans/execute-parent-task-and-prompt-fixes.md`  
> **Specification:** `02-spec/21-app/execute-parent-task-and-prompt-fixes/03-migration-and-sync-spec.md`  
> **State:** `COMPLETED`  
> **Assigned Role:** Subagent Worker 05 (Verification & Release Lead)  

---

## 1. Task Definition & Verification Contract

- **Task ID:** `05-subtask-verification-and-atomic-commit`
- **State:** `[COMPLETED]`
- **Understood / AI Comprehension:** `[YES] — Fully understand the final verification gate and release ceremony requirements, mandating zero relative path violations, strictly lowercase filenames, 100% closed cognitive heatmap telemetry, clean working tree sanitization, and executing the mandatory atomic semantic auto-pushing commit (gitmap cpf) with hyphen formatting and remote push verification.`
- **Azure AI / Model Runtime Context:** `Azure AI / Google Antigravity Swarm (Subagent Worker 05) via GitMap CLI Engine`
- **Intent & Technical Deliverables:**
  - Execute pre-commit relative path audit via `python linter-scripts/check-relative-paths.py` verifying 0 absolute filesystem paths and 0 `file:///` URIs.
  - Audit all modified and authored files for lowercase filename compliance.
  - Execute Gate 3 telemetry verification via `gitmap heatmap status` and `gitmap heatmap tree` ensuring 100% action outcome closure.
  - Sanitize working tree and execute final semantic auto-pushing commit via `gitmap cpf`.
  - Verify remote push upstream to GitHub, fulfilling the mandatory completion invariant ('No Push = Not Done').
- **Negative Boundaries & Invariants:**
  - MUST NOT include colons inside the commit message argument (format strictly as `gitmap cpf "<module> - <summary>"`).
  - MUST NOT leave uncommitted local changes or unpushed local commits upon conclusion.
  - MUST NOT force push (`git push --force`) or amend published commits.
  - MUST NOT disable any CI/CD workflows, GitHub Actions, or validation scripts.
  - MUST NOT execute raw shell searches (`rg`, `ripgrep`, `grep`, `Select-String`, `dir`, `ls`, `Get-ChildItem`).
- **Disjoint File Boundaries:**
  - `Owned Files (Exclusive):`
    - `.ai-memory/plans/subtasks/execute-parent-task-and-prompt-fixes/05-subtask-verification-and-atomic-commit.md`
  - `Read-Only Context Files:**`
    - `02-spec/21-app/execute-parent-task-and-prompt-fixes/03-migration-and-sync-spec.md`
    - `.ai-memory/plans/execute-parent-task-and-prompt-fixes.md`
    - `AGENTS.md`
- **Verification Evidence Contract:**
  - `python linter-scripts/check-relative-paths.py` returns exit code 0 with 0 violations.
  - `gitmap heatmap status --task "execute-parent-task-and-prompt-fixes"` reports 100% met actions.
  - `gitmap heatmap tree --task "execute-parent-task-and-prompt-fixes"` reports zero unclosed actions.
  - `git status` reports working tree clean with upstream tracking branch fully synced.

---

## 2. Detailed Execution Steps

### Step 1: Strict Relative Git Path Linter Audit
- [ ] Run the relative paths validation script:
  ```bash
  python linter-scripts/check-relative-paths.py
  ```
- [ ] Verify that all newly authored specifications, plans, and prompt modifications contain zero absolute paths (`C:\...`, `/home/...`) and zero `file:///` URIs.
- [ ] If any violation is found, resolve immediately using strictly relative Git paths.

### Step 2: Lowercase Filename Compliance Audit
- [ ] Scan newly authored files in the task directory for uppercase letters:
  ```bash
  gitmap find "*[A-Z]*" 02-spec/21-app/execute-parent-task-and-prompt-fixes
  gitmap find "*[A-Z]*" .ai-memory/plans/subtasks/execute-parent-task-and-prompt-fixes
  ```
- [ ] Verify 0 uppercase filenames found.

### Step 3: GitMap Cognitive Heatmap Telemetry 100% Closure
- [ ] Run the telemetry audit commands:
  ```bash
  gitmap heatmap status --task "execute-parent-task-and-prompt-fixes"
  gitmap heatmap tree --task "execute-parent-task-and-prompt-fixes"
  ```
- [ ] Verify that:
  - 100% of recorded actions across all subtasks are marked with `met = true`.
  - Concrete outcome descriptions are recorded for each step.
  - Zero pending or unverified actions remain in the hierarchy.

### Step 4: Working Tree Sanitization & Git Status Inspection
- [ ] Inspect git status to confirm all modified and created files are accounted for:
  ```bash
  git status
  ```
- [ ] Confirm no secrets, temporary files, `.env` files, or test scratch artifacts are present.
- [ ] Stage all modified assets:
  ```bash
  git add -A
  ```

### Step 5: Final Semantic Auto-Pushing Commit & Remote Push
- [ ] Execute the atomic semantic auto-pushing commit following Section 10 of `AGENTS.md`:
  ```bash
  gitmap cpf "prompts - add understood gate and azure ai context to v6 prompts and sync fleet"
  ```
- [ ] Verify that GitMap automatically prepends `Feature: ` prefix.
- [ ] Verify that the remote push completes successfully with exit code 0.

### Step 6: Completion Invariant Verification ('No Push = Not Done')
- [ ] Inspect git status to verify clean working tree and up-to-date tracking branch:
  ```bash
  git status
  ```
- [ ] Confirm that `git status` output confirms `Your branch is up to date with 'origin/main'` and `nothing to commit, working tree clean`.
- [ ] Update heatmap action to mark final completion.

---

## 3. Negative Boundaries & Strict Bans

1. **TOTAL BAN on Colons in Commit Arguments:** Never execute `gitmap cpf "prompts: add ..."`. Commit message argument must strictly use hyphen separation (`gitmap cpf "<module> - <summary>"`).
2. **TOTAL BAN on Leaving Unpushed Commits:** Under Section 13 of `AGENTS.md`, any task with uncommitted changes or unpushed local commits is strictly incomplete. Remote push verification is mandatory.
3. **TOTAL BAN on Rewriting Git History:** No `git push --force`, no rebasing or amending published commits.
4. **TOTAL BAN on Disabling CI/CD:** Never comment out, bypass, or delete CI/CD steps or linters.
