# Fleet Commit, Push & Pooling Architecture Specification

## 1. Executive Summary & Objective

In multi-repository development environments and autonomous agent swarms, tasks often modify code across several repositories or leaves dirty files uncommitted and unpushed in the workspace root.

This specification establishes three binding invariants across all agent prompts, execution runners, and skills:
1. **Mandatory Fleet Commit & Push Invariant ("No Push = Not Done"):**
   If code or artifacts are not committed to Git and pushed upstream to GitHub (main/master/tracking branch), the task is strictly considered **INCOMPLETE and NOT DONE**. Leaving uncommitted changes, dirty working trees, or unpushed commits means the execution is unfinished and has failed.
2. **Pre-Commit Pooling / Pulling Protocol:**
   Before staging and committing dirty changes across repositories, the fleet orchestrator must execute a safe pull (`git pull origin <branch> --no-rebase` or `gitmap pull <repo>`) to reconcile any upstream changes cleanly and prevent divergence.
3. **Non-Owned & Third-Party Repository Exclusions:**
   The workspace scanner must aggressively detect and ignore repositories not owned by our organization or project fleet. Specifically, tools such as `oh-my-zsh`, `ohmyzsh`, `zsh`, `omis`, dotfile managers, Homebrew repos, and external vendor clones must NEVER be automatically mutated, committed to, or pushed.

## 2. Architectural Boundaries & Non-Negotiable Invariants

- **Strict Relative Git Paths:** All documentation, plans, subtasks, scripts, and links MUST use strictly relative paths starting from the repository root (`02-spec/...`, `.ai-memory/...`, `03-ai-scripts/...`). Absolute filesystem paths and `file:///` URIs are totally banned.
- **GitMap Commit Formatting:** All GitMap commits (`gitmap cpf`, `gitmap cpb`, `gitmap cpr`) MUST format commit messages with hyphens (`"<module> - <summary>"`) and NEVER include colons inside the message argument.
- **Search Primacy:** Code exploration MUST use GitMap commands (`gitmap aum search`, `gitmap find`, `gitmap cat`, `gitmap ps`, `gitmap py`); total ban on `rg`, `ripgrep`, `grep`, `Select-String`.
- **Completion Verification Gate:** A task cannot conclude, mark itself successful, or close its turn without proving that all modified files have been staged, committed, and pushed upstream to GitHub.

## 3. Exclusion Engine for Non-Owned Repositories

The fleet discovery scanner in `03-ai-scripts/49-commit-and-push-all-repos.py` and associated prompts must apply strict negative filters:

### Non-Owned Repository Patterns & Directory Names:
- `oh-my-zsh`, `ohmyzsh`, `.oh-my-zsh`
- `zsh`, `.zsh`, `zsh-autosuggestions`, `zsh-syntax-highlighting`
- `omis`, `oh-my-posh`, `.posh`
- `.dotfiles`, `dotfiles`
- `homebrew`, `brew`
- System-level configs, `.config`, `.local`
- External upstream clones without write permissions

Any discovered `.git` folder matching these patterns in its path parts, repository name, or remote URL is categorized as `SKIPPED (Non-Owned / Third-Party)` and never touched by commit or push routines.

## 4. Pre-Commit Pooling / Pulling Protocol

Before staging changes in any discovered owned repository:
1. Detect the active branch (`git rev-parse --abbrev-ref HEAD`). If detached, skip mutation.
2. If the repository has an upstream remote configured (`origin`), perform pre-commit pooling:
   - Check remote status: `git fetch origin <branch>`.
   - If incoming commits exist, pull non-rebase: `git pull origin <branch> --no-rebase`.
   - If merge conflicts arise during pooling, abort mutation for that repository, mark `BLOCKED (Merge Conflict)`, and alert without crashing other repositories.
3. Stage all modified and untracked files (`git add -A`).
4. Generate atomic conventional commit:
   `chore(sync): automated repository sync and working tree commit (<N> files)`.
5. Push to remote tracking branch:
   `git push origin <branch>` (or `git push -u origin <branch>` if upstream unbound).
6. Verify remote SHA matches local HEAD:
   Confirm push success. If push fails (network error, auth error, protected branch), the repository is marked `FAILED` and the task execution is flagged as **INCOMPLETE**.

## 5. Integration with `execute-parent-task-with-n-steps-v6`

The core V6 execution prompt (`01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`) and skills (`.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`, `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`) must embed this agreement directly into:
- Section 12 (Banned Operations Checklist): Auto-reject when changes are uncommitted or unpushed.
- Section 14 (Anti-Hallucination & Blast Radius): Pre-commit diff proof and post-push verification.
- Section 15 (Final Step Git Commit & Push Mandate): Explicit "No Push = Not Done" clause.

## 6. Delivery Plan
1. Update `03-ai-scripts/49-commit-and-push-all-repos.py` with non-owned filters, pooling, and push verification.
2. Update `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` and both skills.
3. Author canonical prompt `01-prompts/09-commit-and-multi-agent-code-fix/09-commit-and-push-all-repos.md`.
4. Update skills `.agents/skills/commit-and-push-all-repos/skill.md` and `.cursor/skills/commit-and-push-all-repos/skill.md`.
5. Synchronize all prompts and skills across the 43 connected repositories using `03-ai-scripts/38-sync-prompts-skills-scripts.py`.
