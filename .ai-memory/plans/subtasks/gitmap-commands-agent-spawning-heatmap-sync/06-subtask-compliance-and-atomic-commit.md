# Subtask 06: Compliance Audit, Linter Verification & Atomic Push

## 1. Subtask Overview

| Attribute | Details |
| :--- | :--- |
| **Subtask ID** | `06-subtask-compliance-and-atomic-commit` |
| **Parent Task** | `gitmap-commands-agent-spawning-heatmap-sync` |
| **Specification Ref** | `02-spec/21-app/gitmap-commands-agent-spawning-heatmap-sync/03-heatmap-and-sync-spec.md` |
| **Owned Files** | `.ai-memory/plans/subtasks/gitmap-commands-agent-spawning-heatmap-sync/06-subtask-compliance-and-atomic-commit.md` |
| **Status** | `COMPLETED` |

---

## 2. Objective & Scope

Conduct a rigorous, multi-faceted compliance audit across all modified and newly authored assets in `coding-guidelines`. Verify strict relative Git path hygiene, lowercase filename enforcement, 100% cognitive heatmap telemetry closure, and execute the final atomic semantic commit and push upstream to GitHub following the mandatory completion invariant ('No Push = Not Done').

---

## 3. Compliance Verification Matrix

| Check Category | Invariant Rule | Verification Method | Pass Threshold |
| :--- | :--- | :--- | :--- |
| **Relative Git Paths** | Strict relative paths (`02-spec/...`, `.ai-memory/...`). TOTAL BAN on `C:\`, `/home/`, `file:///`. | `python 03-ai-scripts/07-relative-paths-linter.py` | 0 violations |
| **Lowercase Filenames** | All files must match `^[0-9a-z_.-]+$` with zero uppercase letters. | GitMap file path scan | 0 violations |
| **Boolean Standards** | Implicit boolean evaluation only (`if isReady`). TOTAL BAN on `== true`. No mixed polarity. | Source code inspection | 0 violations |
| **Cognitive Heatmap** | 100% action outcome closure across all phases. Zero dangling logs. | `gitmap heatmap status` & `gitmap heatmap tree` | All actions closed |
| **Commit Formatting** | Structured prefix with hyphen: `gitmap cpf "<module> - <summary>"`. No colons in argument. | GitMap commit CLI | Strict syntax |
| **Completion Invariant** | Changes committed and pushed upstream to tracking branch. | `git status` / `gitmap push` | Working tree clean |

---

## 4. Detailed Execution Steps

### Step 1: Strict Relative Git Path Linter Audit
- [ ] Run the relative paths validation script:
  ```bash
  python 03-ai-scripts/07-relative-paths-linter.py
  ```
- [ ] Search specifically across newly created and modified files for prohibited absolute patterns:
  - Check for `C:\` or `file:///` in:
    - `02-spec/21-app/gitmap-commands-agent-spawning-heatmap-sync/*.md`
    - `.ai-memory/plans/subtasks/gitmap-commands-agent-spawning-heatmap-sync/*.md`
    - `AGENTS.md`
    - `01-prompts/14-execute/*`
    - `.agents/skills/*`
    - `.cursor/skills/*`
- [ ] Remediate any detected path violations immediately using strictly relative Git paths.

### Step 2: Lowercase Filename Audit
- [ ] Scan repository for any newly introduced uppercase characters in filenames:
  ```bash
  gitmap find "*[A-Z]*"
  ```
- [ ] Ensure all authored spec files and subtask plans use strictly lowercase naming:
  - `03-heatmap-and-sync-spec.md`
  - `05-subtask-multi-repo-sync.md`
  - `06-subtask-compliance-and-atomic-commit.md`

### Step 3: GitMap Cognitive Heatmap Telemetry Closure
- [ ] Execute Gate 3 verification commands:
  ```bash
  gitmap heatmap status --task "gitmap-commands-agent-spawning-heatmap-sync"
  gitmap heatmap tree --task "gitmap-commands-agent-spawning-heatmap-sync"
  ```
- [ ] Inspect the output:
  - Confirm every logged file action has an updated outcome.
  - Verify zero open or unverified actions remaining in the task tree.
  - Output summary report to console.

### Step 4: Working Tree Sanitization & Git Status Inspection
- [ ] Verify clean status and inspect staged vs unstaged delta:
  ```bash
  git status
  ```
- [ ] Confirm no secrets, temporary files, or unneeded artifacts are present.
- [ ] Verify that only relevant files are staged:
  - `AGENTS.md`
  - `02-spec/21-app/gitmap-commands-agent-spawning-heatmap-sync/*.md`
  - `.ai-memory/plans/gitmap-commands-agent-spawning-heatmap-sync.md`
  - `.ai-memory/plans/subtasks/gitmap-commands-agent-spawning-heatmap-sync/*.md`
  - `01-prompts/14-execute/*`
  - `.agents/skills/execute-parent-task-with-n-steps-v6/*`
  - `.cursor/skills/execute-parent-task-with-n-steps-v6/*`
  - `.agents/skills/gitmap/*`
  - `.cursor/skills/gitmap/*`

### Step 5: Final Semantic Auto-Pushing Commit
- [ ] Execute atomic commit and push upstream following `AGENTS.md` Section 10 and 13:
  ```bash
  gitmap cpf "skills - enforce gitmap commands a=3 spawning heatmap gates and sync"
  ```
- [ ] Verify that:
  - Prefix `Feature: ` is automatically prepended by GitMap.
  - Upstream branch is cleanly updated on GitHub.
  - `git status` reports working tree clean.

---

## 5. Negative Boundaries & Invariants

1. **TOTAL BAN on Colons in Commit Arguments:** Never execute `gitmap cpf "module: summary"`. Format strictly as `gitmap cpf "<module> - <summary>"`.
2. **TOTAL BAN on Leaving Unpushed Commits:** Under Section 13 ('No Push = Not Done'), the task is strictly considered incomplete until upstream remote push is verified.
3. **TOTAL BAN on Rewriting Published History:** No `git push --force`, no rebasing pushed commits (Lovable and multi-agent safety).
4. **TOTAL BAN on Disabling CI/CD:** Never comment out or disable GitHub Actions workflows or linters.
