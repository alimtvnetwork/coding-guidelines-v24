# Specification: Unified Muse Master Onboarding & Execution Prompt (V6 Architecture)

> Category: `21-app`
> Status: Active
> Architecture: Self-Bootstrapping Autonomous Agent Directive & Execution Engine
> Reference: `01-prompts/27-muse-prompts/01-muse-master-prompt.md`
> Canonical Engine: `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`

---

## 1. User Request (Verbatim)

```text
In the coding guideline prompt section, in the Muse prompt, I want you to combine all this prompt to a single one following the execute NS steps V6, something like this. That means it would understand everything. Also, it would have the steps. For example, it should also have these steps properly, installation, everything all together so that it can follow through and create all the chats for each one of them. Does this make sense? Also, try to include this onboarding part somehow to the... Okay, in the master prompt, the coding guideline full URL should be there, how it should start with, and also it should have the GitMap installation command, including that. And these are the steps on top of my head, but also I think you should improve it so that anytime I put that prompt in, the Muse should be able to clone all these repos that we want. It should be also able to understand and clone all these repositories, also following the repo cache, and also create chats for each one of the repositories. And also, basically, GitMap should know all of the repository information by running the GitScan, GitMap scan, and also GitMap login needs to be corrected either using the SSH or from the browser. SSH is something that I would prefer. Write it like this. So first, if I paste it there, if it is already logged in, then it would move to the next steps. That means pulling all these repositories using the GitMap PA, updating the GitMap first, and then pulling all, and then proceeding with other steps. Make sure that none of the repositories has a conflict. Things like that. Do you understand? Do you follow through with this? So merge all of these for the Muse prompt to a single one. Okay, please do that. With a checklist, like the execute kind of steps in V6. Okay? Follow through like that as well. Is it clear?
```

---

## 2. Executive Summary & Architectural Motivation

Previously, category `27-muse-prompts` contained 5 separate, fragmented prompts:
1. `01-muse-master-prompt.md` (Master onboarding prompt v6.1.2)
2. `02-muse-execute-in-a-step.md` (Execute-in-a-step one-shot prompt v1.1.0)
3. `03-muse-master-with-repos.md` (Repo-aware master prompt v1.1.0 with 10 repos)
4. `04-muse-master-with-repos-v2.md` (Repo-aware master prompt v1.3.1 with 14 repos, pull first, 15MB split rule)
5. `05-muse-master-prompt-v2.md` (Master prompt v2 v6.2.2)

Having 5 different prompts created confusion about which prompt to paste into a new Muse session, fragmented crucial operational invariants (such as repository routing, pull-first safeguards, large file split rules, and execution tracking), and lacked an end-to-end automated onboarding pipeline that could bootstrap GitMap, verify authentication (with preferred SSH), execute `gitmap scan`, update GitMap, pull the entire fleet with `gitmap pa`, ensure conflict-free trees, clone missing repositories, set up per-repository side chats, and execute work using the canonical V6 multi-agent engine (`N=300`, `A=2`, `H=2`, `C=30`).

This specification formalizes the **Unified Muse Master Prompt (V6)**, merging all previous prompt behaviors into a single cohesive, self-contained, execute-N-steps-V6 directive.

---

## 3. Core Architecture & Pipeline Phases

The unified prompt is structured into a 10-phase pipeline governed by V6 execution parameters and an actionable checklist (`## 🎯 Actionable Agent Checklist`):

### Phase 0: Parameter Header & Slash Commands
- Parameters: `N = 300`, `A = 2`, `H = 2`, `C = 30`, System Concurrency Capacity = 4 operations (`A × H`).
- Preamble Precedence / Top-Instruction Priority Mandate.
- Slash commands: `[/goal]`, `[/learn]`, `[/plan]`.
- Mandatory Subagent Spawning Gate (`invoke_subagent`, `A = 2, H = 2`, zero solo execution allowed).

### Phase 1: Authentication Bootstrap (SSH Primacy & Fast-Path Bypass)
- Authentication check:
  - Check if already authenticated (`ssh -T git@github.com` or `gh auth status` and `gitmap login --status`).
  - **Fast-Path Bypass:** If already logged in, IMMEDIATELY skip authentication and proceed directly to Phase 2!
  - If not authenticated: Preferred method is **SSH** (`eval "$(ssh-agent -s)"`, `ssh-add`, or configuring SSH public keys with GitHub CLI `gh auth login -p ssh -w`).
  - Fallback method: Browser login (`gitmap login --web` or `gh auth login`).
  - Zero private clones or pushes occur before authentication passes.

### Phase 2: GitMap Installation, Self-Update & Fleet Scan
- **GitMap Installation Command:**
  `eval "$(curl -fsSL https://raw.githubusercontent.com/alimtvnetwork/gitmap-v28/main/install-quick.sh)"`
  Fallback: Clone `https://github.com/alimtvnetwork/gitmap-v28.git` and build.
- **Update GitMap First:**
  Ensure GitMap binary is up to date before touching repos (`gitmap update` or pull/build).
- **Fleet Scan (`gitmap scan`):**
  Run `gitmap scan` across `~/workspace/repos/` so GitMap discovers and registers all repository metadata, local paths, remotes, and branch states in `gitmap.db`.

### Phase 3: Repository Fleet Orchestration, Clone, Pull-All (`gitmap pa`) & Conflict Prevention
- **Starting Point & Canonical Repositories:**
  - Coding Guidelines canonical URL: `https://github.com/alimtvnetwork/coding-guidelines-v24.git`
  - Explicit warning: `coding-guidelines` is old/stale (2026-03-31); all prompts and guidelines live in `coding-guidelines-v24`.
  - GitMap canonical URL: `https://github.com/alimtvnetwork/gitmap-v28.git`.
  - Repo Cache URL: `https://github.com/alimtvnetwork/repo-cache.git` (`repo-cache/NN-<slug>`, sequence numbers, 7-day memory sync).
  - Repo Secrets URL: `https://github.com/alimtvnetwork/repo-secrets.git` (private credentials via `gitmap rs`).
- **Fleet Cloning:**
  Ensure all 14+ repositories exist in `~/workspace/repos/`. Clone any missing ones using `gitmap clone <url> ~/workspace/repos/<repo>`.
- **Pull-All (`gitmap pa`) & Conflict Prevention:**
  - Run `gitmap pa` across the workspace fleet.
  - Inspect repository statuses for uncommitted changes or divergent branches.
  - Guarantee zero merge conflicts before starting work (`git pull --ff-only` or clean non-rebasing sync).
- **Side Chat Architecture:**
  - Create/manage dedicated side chats named strictly `<repo> repo` for every repository in the fleet.
  - Dedicated **"all summary"** side chat answering "What is the all summary?" with multi-repo status recaps.

### Phase 4: Memory Operating Contract Checklist & Turbo Matrix
- Checklists for TURBO-01 through TURBO-05 and GUARD-01 through GUARD-09.
- Sacred repo rule: NEVER delete any repository or wipe folders. File deletion always requires explicit confirmation.
- CI-only verification: Never run test suites locally; verify pushes via `gitmap pe -t` and inspect failures via `gitmap pe`.
- Turbo mode ask/don't-ask decision boundary matrix.

### Phase 5: Coding Standards, App Formatting & Large File Gate
- Cross-language coding guidelines: booleans (`is*`/`has*`, no `== true`, no mixed polarity), early returns/guard clauses, `*appfault.AppError`, lowercase filenames, strict relative git paths.
- Large files (>15 MB): 7-Zip multi-volume split (`7z a -v3m -mx=0 <folder>/<name>.7z <file>` -> `.7z.001`, `.7z.002`, ... + `README.txt`). Total ban on raw large files and corrupt `zip -s`.
- Title convention for user app: `# <task title>: high priority instruction, non-negotiable task`, `## slug: <task-slug>` (mapping: Literally app ↔ Letterly repo template).
- UI/UX design tokens: Ubuntu typography, tokenized variables, responsive state matrices.

### Phase 6: Task Confirmation & Multi-Agent Execution Protocol (V6 Engine)
- Step 1: Lossless verbatim capture under `## User Request (Verbatim)`.
- Step 2: Confirmed task breakdown FIRST in visible chat (`Task-01`, `Task-02`, ... each with `Understood: [YES]`, `Slug: <root-slug> - Task NN`).
- Step 2.5: SQLite task manager initialization via `gitmap agent task init --name "<task>" --budget 300` + ledger at `.ai-memory/temp-agents/<slug>/ledger.md`.
- Step 2.6: Mandatory RUNNING declaration with ETA (`RUNNING — Task-01, Task-02 — ETA ~45 min`) and 5-minute status pings. Total ban on `LISTING-WITHOUT-STARTING`.
- Step 2.7: Pull first in active repo before touching files (`gitmap pull` / `git pull`).
- Step 3: Mandatory subagent dispatch (`invoke_subagent`, `A = 2, H = 2`) in disjoint file boxes. Worker Git Ban.
- Understanding check-ins: Prove understanding in 1-2 lines on "okay?", "clear?", "understood?".

### Phase 7: Verification, Secrets Gate, Atomic Push & CI Telemetry
- Fast targeted file checks only (R1 zero build/test ban).
- Pre-commit secrets gate: `check-forbidden-strings.py` and `gitmap aum search` regex -> offload to `repo-secrets` via `gitmap rs`.
- GitMap atomic commit + immediate push using hyphen format (no colons in arguments):
  - `gitmap cpf "<module> - <summary>"`
  - `gitmap cpb "<module> - <summary>"`
- CI Telemetry: `gitmap pe -t` watches pipeline to completion; `gitmap pe` inspects errors if red.

### Phase 8: Muse Memory IO (Memory Re-sync Protocol)
- Trigger: `Muse memory IO read`.
- Run `gitmap llm train muse --days 7` from workspace root.
- Read sequenced digest files (`01-overview.md` -> `05-conventions.md`).
- Recreate recent chats from `12-muse-ai/chat-seeds/*.md`.
- PRIVATE ITEMS RULE & EMAIL RULE (standing SMTP pre-approval).

### Phase 9: Master Verification Checklist & Handoff
- Verification gates V-01 through V-14.
- Ready report + zero-stop transition to first task.
