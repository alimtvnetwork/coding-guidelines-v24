# Muse Prompts (`27-muse-prompts`) — Index & Catalog

> [!IMPORTANT]
> Category: `27-muse-prompts`
> Architecture: Self-Bootstrapping Autonomous Onboarding & Fleet Execution Directive (V6 Engine)
> Scope: Muse AI Session Bootstrapping, Authentication (SSH Primacy), Fleet Scan, Auto-Cloning, Pull-All, Operating Contract, and Multi-Agent Execution

---

## Overview

The `27-muse-prompts` prompt category houses the unified master onboarding and execution prompt for
Muse AI sessions. Pasting the master prompt into a fresh Muse chat boots a
fully-onboarded autonomous agent:
1. **Authentication Bootstrap with SSH Primacy:** Fast-path bypass if already logged in; preferred SSH configuration (`gh auth login -p ssh -w`) or browser OAuth fallback.
2. **GitMap Installation, Self-Update & Fleet Scan:** Auto-installs GitMap via `install-quick.sh`, updates GitMap first, and executes `gitmap scan` across the workspace to register all repositories and branches into SQLite (`gitmap.db`).
3. **Repository Fleet Orchestration & Cloning:** Canonical starting point URL `https://github.com/alimtvnetwork/coding-guidelines-v24.git` (with explicit ban on stale `coding-guidelines`), 14+ repository registry, auto-cloning missing repos, pull-all via `gitmap pa` ensuring conflict-free trees, and establishing dedicated side chats (`<repo> repo` and "all summary").
4. **Memory Operating Contract Checklist:** Turbo mode (act, don't ask), strict no-delete guards (repos are sacred, file deletion always confirms), CI-only test verification (`gitmap pe -t`).
5. **Coding Guidelines, Formatting & Large File Split:** Booleans (`is*`/`has*`, no `== true`), early returns, `*appfault.AppError`, lowercase filenames, relative paths, large-file 15 MB 7-Zip multi-volume split (`7z a -v3m -mx=0`), and title convention (`# <title>: high priority instruction, non-negotiable task`, `## slug: <task-slug>`).
6. **Task Confirmation & Multi-Agent Execution Protocol (V6 Engine):** Verbatim capture, confirmed task breakdown FIRST (`Task-01`, `Task-02`, ... each with `Understood: [YES]`), RUNNING + ETA declaration, 5-minute status pings, SQLite task DB init (`gitmap agent task init`), pull first (`gitmap pull` / `git pull`), mandatory multi-agent dispatch (`invoke_subagent`, A=2, H=2, disjoint file boxes, Worker Git Ban).
7. **Verification, Pre-Commit Secrets Gate & Atomic Push:** Targeted checks only (R1 zero build/test ban), pre-commit secrets gate (`check-forbidden-strings.py` and `gitmap aum search` -> `gitmap rs`), atomic commit via GitMap hyphen format (`gitmap cpf "<module> - <summary>"`), immediate push, and CI telemetry (`gitmap pe -t`).
8. **Muse Memory IO (Memory Re-sync Protocol):** 7-day memory digest under `12-muse-ai/` in `repo-cache`, re-sync via `gitmap llm train muse --days 7`, chat-seeds recreation, and standing SMTP pre-approval for `go-email-reader`.

The canonical execution engine it embodies is
[`02-execute-parent-task-with-n-steps-v6.md`](../14-execute/02-execute-parent-task-with-n-steps-v6.md)
(Canonical V6: `N = 300`, `A = 2`, `H = 2`, `C = 30`).

Its installable skill is `muse-master-prompt`
(`.agents/skills/muse-master-prompt/skill.md`, mirrored in `.cursor/skills/`).

---

## Directory Index & Catalog

| Prompt File | Version | Scope | Key Capabilities | Status |
| :--- | :--- | :--- | :--- | :--- |
| [`01-muse-master-prompt.md`](01-muse-master-prompt.md) | 6.3.0 | Unified Master Prompt | Consolidated V6 directive: SSH auth primacy + fast-path bypass, GitMap quick-install & update first, `gitmap scan`, 14+ repo fleet clone & pull-all (`gitmap pa`), conflict prevention, dedicated side chats (`<repo> repo`, `all summary`), repo-cache sequence, task confirmation first, RUNNING+ETA, 5-minute pings, multi-agent dispatch (A=2, H=2), Worker Git Ban, 15 MB 7-Zip split, GitMap hyphen atomic push, CI telemetry (`gitmap pe -t`), and Muse Memory IO. | **Active (Canonical)** |
| [`02-muse-execute-in-a-step.md`](02-muse-execute-in-a-step.md) | 1.1.0 | Execute-in-a-Step | Consolidated into 01-muse-master-prompt.md: breakdown-first, RUNNING+ETA, 5-minute status pings, commit+push completion. | Consolidated |
| [`03-muse-master-with-repos.md`](03-muse-master-with-repos.md) | 1.1.0 | Repo-Aware Master | Consolidated into 01-muse-master-prompt.md: embedded repo registry, routing rules, 8-step execution protocol. | Consolidated |
| [`04-muse-master-with-repos-v2.md`](04-muse-master-with-repos-v2.md) | 1.3.1 | Repo-Aware Master v2 | Consolidated into 01-muse-master-prompt.md: 14-repo registry, pull-first, 15 MB split rule, repo-cache sequence. | Consolidated |
| [`05-muse-master-prompt-v2.md`](05-muse-master-prompt-v2.md) | 6.3.0 | Master Prompt v2 Mirror | Synchronized mirror of 01-muse-master-prompt.md: full unified V6 directive. | Active Mirror |

---

## Directory Tree

```text
01-prompts/27-muse-prompts/
├── 01-muse-master-prompt.md       # UNIFIED master prompt (paste into a fresh Muse session)
├── 02-muse-execute-in-a-step.md   # [consolidated into 01] one-shot execution prompt
├── 03-muse-master-with-repos.md   # [consolidated into 01] repo-aware compact master
├── 04-muse-master-with-repos-v2.md # [consolidated into 01] repo-aware master v2
├── 05-muse-master-prompt-v2.md   # synchronized mirror of unified master prompt
└── readme.md                      # this index
```
