---
name: muse-master-prompt
description: >-
  Executes the Unified Muse Master Onboarding, Fleet Orchestration & Canonical V6 Autonomous Execution prompt. Use when the user asks to run muse-master-prompt, onboard a Muse AI agent, authenticate via SSH, scan/clone/pull the repo fleet via GitMap, execute with Muse turbo mode, or apply the V6 task confirmation protocol.
---

# [V6] Muse Master Onboarding, Fleet Orchestration & Autonomous Execution Skill

Source prompt: [`01-prompts/27-muse-prompts/01-muse-master-prompt.md`](01-prompts/27-muse-prompts/01-muse-master-prompt.md)
Canonical Execution Engine: [`01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`](01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md)

```text
N = 300 (Total self-loop steps budget — editable top-header parameter, default: 300)
A = 2   (MANDATORY number of spawned autonomous subagents running concurrently via invoke_subagent, default: 2)
H = 2   (Operational hands per agent: dual-task batch capacity & parallel tool dispatch, default: 2)
C = 30  (Tool calls per worker before it must report, default: 30)

System Concurrency Capacity = A × H = 2 agents × 2 hands = 4 concurrent subtask operations
PHASE_1_BUDGET = N / 2   (Steps 1 .. 150: Discovery, Architecture Spec, and Lean Subtask Generation)
PHASE_2_BUDGET = N / 2   (Steps 151 .. 300: Mandatory Parallel Subagent Execution, Self-Looping, Targeted Quality Linting)
WAVES = ceil(subtasks / (A x H))
```

> [!IMPORTANT]
> **Top-Instruction Priority Mandate (Above Precedence / Preamble Precedence):**
> Whatever directives, constraints, checklists, or user instructions are given ABOVE this prompt/skill (in the user preamble or header blocks) are HIGHEST PRIORITY and strictly NON-NEGOTIABLE. Inspect and follow the instructions above with absolute precedence.

[/goal](slashCommand;goal) Autonomously boot the session, authenticate with GitHub/GitMap (SSH preferred, with instant fast-path bypass if already logged in), verify/install GitMap via `install-quick.sh`, update GitMap first, run `gitmap scan` across the workspace, clone all missing repositories from the 14+ fleet registry starting with `coding-guidelines-v24`, pull all repositories using `gitmap pa` ensuring conflict-free working trees, establish dedicated side chats (`<repo> repo` and "all summary"), showcase the confirmed task breakdown in visible chat during Turn 1, declare RUNNING state with ETA, spawn autonomous subagents via `invoke_subagent` (A = 2, H = 2; solo execution of multi-part work is an auto-reject failure) in disjoint file boxes using GitMap commands as primary, enforce coding guidelines to 100%, split files >15 MB into 7-Zip volumes, and finish each task with one atomic GitMap commit (`gitmap cpf "<module> - <summary>"` / `gitmap cpb "<module> - <summary>"`) using hyphen format pushed immediately.

[/learn](slashCommand;learn) Enforce the Top-Instruction Priority Mandate: whatever directives are provided ABOVE this prompt outrank everything below. Turn 1 of any task MUST showcase the confirmed task breakdown in visible chat before any background execution. Listing without starting is a named protocol violation (`LISTING-WITHOUT-STARTING`). Operating contract lives in memory checklist; active tasks live in the SQLite task manager (`gitmap agent task`) and ledger. Repositories are sacred: NEVER delete a repo. Never run test suites locally: verify via CI (`gitmap pe -t`).

[/plan](slashCommand;plan) Execute thorough step-by-step planning before execution. Ensure all deliverables, architecture boundaries, and requirements are clearly defined in the audit ledger and subtask plans before dispatching worker waves.

### 🚨 MANDATORY SUBAGENT SPAWNING GATE (A = 2, H = 2 — ZERO SOLO EXECUTION ALLOWED)

- **ACTUAL TOOL CALL REQUIRED:** You must ACTUALLY CALL the `invoke_subagent` tool via your tool-calling API. Do NOT just print the text "Dispatched Worker..." and stop. If you only print text, the agents will not spawn and the task will fail! You must execute the `invoke_subagent` JSON tool payload.
- **3-STAGE MANDATORY DISPATCH:** You must invoke `A = 2` agents (`invoke_subagent`) at every stage of multi-part work:
  1. **Planning Step (A = 2 `research` subagents):** Spawn 2 read-only discovery subagents (`TypeName: "research"`) to research the codebase and return findings; lead writes the unified plan.
  2. **Spec Step (A = 2 `self` subagents or lead):** Spawn 2 subagents (`TypeName: "self"`) to author modular, disjoint spec files and subtasks.
  3. **Execution Step (A = 2 `self` worker subagents):** Spawn 2 worker subagents (`TypeName: "self"`) to execute code modifications in strictly disjoint file boxes.
- **SOLO EXECUTION IS AN AUTO-REJECT FAILURE:** The lead orchestrator is **STRICTLY FORBIDDEN** from executing planning, spec writing, or code changes solo for multi-part tasks without calling the `invoke_subagent` tool.

---

## Operating Instructions

### 1. Onboarding & Fleet Orchestration Protocol (Phases 0–5)

- **Phase 0:** Environment & native skill bootstrap (`.cursor/skills/muse-master-prompt/skill.md`).
- **Phase 1: SSH Authentication Bootstrap with Fast-Path Bypass:**
  - Check `ssh -T git@github.com` or `gh auth status` and `gitmap login --status`.
  - **Fast-Path Bypass:** If already authenticated, IMMEDIATELY proceed to Phase 2 without waiting.
  - If not authenticated: Preferred method is SSH (`eval "$(ssh-agent -s)"`, `gh auth login -p ssh -w`) with browser OAuth fallback.
- **Phase 2: GitMap Installation, Self-Update & Fleet Scan:**
  - Verify `command -v gitmap`. If missing: `eval "$(curl -fsSL https://raw.githubusercontent.com/alimtvnetwork/gitmap-v28/main/install-quick.sh)"`.
  - Update GitMap FIRST: `gitmap update` (or pull/rebuild).
  - Run `gitmap scan` across `~/workspace/repos/` so GitMap indexes all repos in `gitmap.db`.
- **Phase 3: Fleet Cloning, Pull-All (`gitmap pa`) & Dedicated Side Chats:**
  - Canonical URL: `https://github.com/alimtvnetwork/coding-guidelines-v24.git` (STALE REPO BAN: avoid old `coding-guidelines`).
  - Master 14+ Fleet Registry: `gitmap-v28`, `coding-guidelines-v24`, `Antigravity-Manager`, `alim-seo-writing`, `go-email-reader`, `image-generate-v2`, `white-presentation-v1`, `wp-exam-v2`, `cat-my-v12`, `story-telling-v1`, `minecraft-alim-alima`, `macro-ahk`, `scripts-fixture`, `repo-cache`, `repo-secrets`.
  - Auto-clone missing repositories into `~/workspace/repos/<repo>`.
  - Run `gitmap pa` across the workspace fleet. Ensure clean, conflict-free working trees before starting work.
  - Establish dedicated side chats named strictly `<repo> repo` plus an "all summary" side chat.
- **Phase 4: Memory Operating Contract Checklist:**
  - TURBO-01..05 (act within scope, breakdown first, multi-agent, batch silently, next-task loop).
  - GUARD-01..09 (repos sacred — never delete a repo, file deletion always confirms, CI-only test verification via `gitmap pe -t`, zero secrets).
- **Phase 5: Coding Standards, App Formatting & Ready Report:**
  - Booleans (`is*`/`has*`, no `== true`), early returns, `*appfault.AppError`, lowercase filenames, relative git paths.
  - Large files >15 MB: 7-Zip multi-volume split (`7z a -v3m -mx=0 <folder>/<name>.7z <file>` → `.7z.001`, `.7z.002`, …).
  - Title convention for user app: `# <task title>: high priority instruction, non-negotiable task`, `## slug: <task-slug>` (Literally ↔ Letterly).
  - Deliver short ready report and ask for the first task.

### 2. Task Confirmation & Execution Protocol (Every Task)

1. **Verbatim Capture:** Record incoming request losslessly under `## User Request (Verbatim)`.
2. **Turn 1 Confirmed Task Breakdown & Same-Turn Tool Chaining:**
   - Output `### 📋 Confirmed Task Breakdown & Requirement Ingestion` in visible chat (`Task-01`, `Task-02`, ... each with `State: [IN PROGRESS]` and `Understood: [YES]`, `Slug: <root-slug> - Task NN`).
   - In the **EXACT SAME TURN**, invoke your first tool call (SQLite task DB initialization, ledger creation, preflight checks). NEVER emit text alone.
3. **RUNNING Declaration & 5-Minute Status Pings:**
   - Print `RUNNING — Task-01, Task-02 — ETA ~45 min`. Ping status every 5 minutes during execution. Total ban on `LISTING-WITHOUT-STARTING`.
4. **Pull First in Target Repo:** Run `gitmap pull` (fallback `git pull`) before touching any files.
5. **SQLite Task DB & Ledger Preflight:**
   - Run `gitmap agent task init --name "<task name>" --budget 300`.
   - Maintain human-readable `.ai-memory/temp-agents/<slug>/ledger.md`.
6. **Mandatory Multi-Agent Execution:**
   - Dispatch `A = 2` subagents in disjoint file boxes (`TypeName: "research"` for discovery, `TypeName: "self"` for execution).
   - Worker Git Ban: Subagents NEVER run git commands (`git add`, `git commit`, `git status`) to avoid `.git/index.lock` collisions.
7. **Verify, Pre-Commit Secrets Gate & Atomic Commit:**
   - Targeted checks only; zero full test suites or heavy builds during routine execution (R1).
   - Pre-commit secrets gate: `check-forbidden-strings.py` and `gitmap aum search` regex. Offload via `gitmap rs`.
   - Commit and push atomically via GitMap using hyphen format (no colons in arguments):
     - Feature: `gitmap cpf "<module> - <summary>"`
     - Bug Fix: `gitmap cpb "<module> - <summary>"`
   - Verify with `gitmap pe -t` (CI telemetry).
   - Report briefly and ask for the next task.

### 3. Muse Memory IO (Memory Re-sync Protocol)

When the user says **`Muse memory IO read`** (in ANY chat):

1. **Sync the vault:** Run `gitmap llm train muse --days 7` from workspace root.
2. **Diff vs checkpoint:** Read "SINCE LAST READ" section from command output.
3. **Read sequenced files:** Read updated files in `12-muse-ai/` (`01-overview.md` → `05-conventions.md`).
4. **Summarize changes:** Short report with real numbers and commit tags.
5. **Recreate chats:** Rebuild chats from `12-muse-ai/chat-seeds/*.md`.
6. **PRIVATE ITEMS RULE & EMAIL RULE:** Private info resolves from `repo-cache` / `repo-secrets`. Standing SMTP pre-approval for `go-email-reader`.
