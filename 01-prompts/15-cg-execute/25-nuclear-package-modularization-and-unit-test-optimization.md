```text
N = 300 (Total self-loop steps budget — editable top-header parameter, default: 300)
A = 2   (MANDATORY number of spawned autonomous subagents running concurrently via invoke_subagent, default: 2)
H = 2   (Operational hands per agent: dual-task batch capacity & parallel tool dispatch, default: 2)
C = 30  (Tool calls per worker before it must report, default: 30)

System Concurrency Capacity = A × H = 2 agents × 2 hands = 4 concurrent subtask operations
PHASE_1_BUDGET = N / 2   (Steps 1 .. 150: Planning, Parallel Discovery Subagents, Detailed Spec, and Lean Subtask Generation)
PHASE_2_BUDGET = N / 2   (Steps 151 .. 300: Mandatory Parallel Subagent Execution, Self-Looping, Targeted Quality Linting)
WAVES = ceil(subtasks / (A x H))
```

[/goal](slashCommand;goal) Autonomously scan, audit, plan, and modularize monolithic packages and optimize unit test execution across the codebase. Enforce a strict Directed Acyclic Graph (DAG) architecture, extract reusable zero-dependency leaf packages, isolate slow or destructive tests (`exec.Command`, git CLI subprocesses, network sockets, `time.Sleep`) into external blackbox test packages (`tests/heavy_test/` under `package heavy_test`), maintain the centralized test inventory manifest (`.ai-memory/test-inventory.json`), and adhere to the 5-day cache freshness decision engine to achieve ultra-fast sub-0.05s unit test execution loops without circular dependencies: FIRST showcase and list out the given task in visible chat during Turn 1, capture it verbatim, plan it in the repo, spawn autonomous subagents via `invoke_subagent` (A = 2, H = 2; solo execution without calling `invoke_subagent` is an auto-reject failure) in disjoint file boxes using GitMap high-speed commands as primary, prove every single claim with concrete evidence, enforce coding guidelines to 100%, and finish with one atomic GitMap commit that holds strictly this task's files.

[/learn](slashCommand;learn) Enforce the Top-Instruction Priority Mandate: whatever directives, custom requirements, checklists, or user instructions are provided ABOVE this prompt outrank everything below. Turn 1 MUST showcase the given task list in visible chat before any background execution. Each rule is stated once (R1 to R16) and cited by ID. Progress lives in the ledger and in `.ai-memory/plans/`, never only in chat.

[/plan](slashCommand;plan) Execute thorough step-by-step planning in the repository before execution. Ensure all deliverables, architecture boundaries, and requirements are clearly defined in the audit ledger and subtask plans before dispatching worker waves.

### 🚨 MANDATORY SUBAGENT SPAWNING GATE (A = 2, H = 2 — ZERO SOLO EXECUTION ALLOWED)

- **ACTUAL TOOL CALL REQUIRED:** You must ACTUALLY CALL the `invoke_subagent` tool via your tool-calling API. Do NOT just print the text "Dispatched Worker..." and stop. If you only print text, the agents will not spawn and the task will fail! You must execute the `invoke_subagent` JSON tool payload.
- **3-STAGE MANDATORY DISPATCH:** You must invoke `A = 2` agents (`invoke_subagent`) at EVERY stage of the workflow:
  1. **Planning Step:** Spawn 2 subagents to research the codebase and author the step-by-step execution plan.
  2. **Spec Step:** Spawn 2 subagents to write the detailed architectural spec.
  3. **Execution Step:** Spawn 2 worker subagents to execute the actual code modifications.
- **SOLO EXECUTION IS AN AUTO-REJECT FAILURE:** The lead orchestrator is **STRICTLY FORBIDDEN** from executing planning, spec writing, or code changes by itself without calling the `invoke_subagent` tool. Failing to call the actual tool is a critical protocol violation.

---

## The Unified Master Pipeline (Atomic Numbered Steps)

Execute this task via a strict 3-Phase pipeline. Do not skip steps.

### Phase 1A: Verbatim Capture, Task Extraction & Chat Output Gate (Step 0)

Before executing any file searches, scans, spec writing, or code changes, you must execute Phase 1A:

1. **Top-Instruction Priority Verification:** Whatever directives, constraints, checklists, or instructions are given before this section or prompt (user preamble, header constraints, prior instructions) must be verified as highest priority and non-negotiable.
2. **Showcase Given Task First (Turn 1 Action):** In your VERY FIRST response turn upon receiving the prompt, you MUST output the confirmed task breakdown directly in visible chat. Never execute tools silently without displaying the task breakdown to the user first!
3. **Lossless Verbatim Capture:** Store incoming prompt losslessly under `## User Request (Verbatim)` in canonical spec and parent plan.
4. **Screenshots & Media:** Decode base64/screenshots immediately into `assets/screenshots/<slug>-<NN>.png`. Reference via relative markdown links (`![Screenshot](assets/screenshots/<slug>-<NN>.png)`).
5. **Discrete Deliverables Extraction:** Break down whatever user requirements were given into discrete, actionable items with ordered traceable IDs (`Task-01`, `Task-02`, ...).
6. **Mandatory Same-Turn Tool Chaining (TOTAL BAN ON TURNING OFF):** Emit the breakdown in chat with clean vertical formatting, and in the **EXACT SAME TURN**, invoke your first tool call (e.g. `write_to_file` to initialize ledger/spec, or run preflight). NEVER emit text alone (which ends the turn prematurely), and never ask "Should I proceed?".

```markdown
### 📋 Confirmed Task Breakdown & Requirement Ingestion

1. **Task-01: [Descriptive Task Title]**
   - **State:** `[IN PROGRESS — EXECUTING IMMEDIATELY]`
   - **Understood:** `[YES]` — [1-2 concise sentences proving understanding of intent, scope, and verified constraints]
   - **Actionable Scope:** [Precise technical deliverable and implementation scope]
   - **Target Files / Area:** `[relative/path/or/module]`

2. **Task-02: [Descriptive Task Title]**
   - **State:** `[QUEUED — EXECUTING NOW WITHOUT USER PROMPT]`
   - **Understood:** `[YES]` — [1-2 concise sentences proving understanding of intent, scope, and verified constraints]
   - **Actionable Scope:** [Precise technical deliverable and implementation scope]
   - **Target Files / Area:** `[relative/path/or/module]`

Proceeding directly to Preflight & Phase 1B Spec Generation (Active Tool Call Running Below).
```

---

## 1. Precedence Hierarchy & Scope (Highest First)

1. **User Instructions & Preamble:** Directives and parameters ABOVE this prompt outrank everything below.
2. **Platform Limits:** Native tools, Artifact Review Policy, permission prompts, hooks. Never claim to override them.
3. **Repo Rules:** `AGENTS.md`, `.ai-memory/strictly-avoid.md`, and `coding-guidelines.md`.
4. **This Prompt.**

If sources conflict, follow stricter one and record under `Conflicts:` in ledger.

### Scope Control Rules
- **Turn 1 Task Showcase:** In your very first response turn, you MUST showcase and list out the given task in visible chat. Running tools silently without presenting the task breakdown is strictly banned.
- **Read budget:** Read only requested paths, search hits, and Step 0 context.
- **History read-only:** Never edit past events, changelogs, completed plans, release notes, or `06-old-prompts/` and `19-old-execute-prompts/`.
- **Out-of-scope:** Log under `Follow-ups:` in plan; never fix in this run.
- **Minimal diff:** Change only lines required; never reflow unaffected lines.
- **Indexes:** Update `01-prompts/readme.md` and `.ai-memory/prompts.md` only when adding/modifying prompts. Update `.ai-memory/plans/readme.md` and `02-spec/21-app/readme.md` every run. Register recent completed tasks before push.
- **Mandatory Multi-Agent Partitioning:** Even for tasks touching few files, work MUST be partitioned across A workers (e.g. Worker 01 implements changes, Worker 02 implements verification/linters/companion tests). Solo execution is strictly banned.
- **Finish early:** When all Task-IDs are `DONE`, proceed directly to consolidation.
- **Zero releases:** Never bump versions or edit changelogs unless requested (R10).

---

## 2. Core Operational Rules (Cite by ID)

- **R1 Zero Builds or Test Suites (TOTAL BAN).** NEVER run `go build`, `npm run build`, `vite build`, `go test ./...`, `pytest`, `npm test`, or `03-ai-scripts/06-cicd-local-runner.py`. CI verifies builds and suites. Routine turns must never waste time on heavy compilation/tests. Only explicit user command lifts this.
- **R2 Targeted Checks Only.** Run only fast, file-scoped checks on specifically modified files (see Section 10). A check scanning 0 files is a **FAIL**.
- **R3 Evidence or It Did Not Happen.** Every `DONE`, `PASS`, or "verified" claim MUST cite a concrete file path, git diffstat, or command exit code (`exit 0`). Vague assurances are auto-rejected.
- **R4 Never Invent Commands, Flags, or Paths.** Verify commands with a harmless call (`gitmap lf readme.md`), not `--help`. Use documented fallbacks and log in ledger.
- **R5 Mandatory Subagents (`invoke_subagent`).** Spawning subagents via `invoke_subagent` (`A = 2`, `H = 2`) is an **ABSOLUTE MUST** (`research` for discovery in Phase 1, `self` for edits in Phase 2). The lead agent is STRICTLY FORBIDDEN from executing all reads or edits solo. Solo execution without calling `invoke_subagent` is an auto-reject failure on the same tier as Rule 0.
- **R6 One Owner Per File (Disjoint Bounding Boxes).** Within every worker wave, each file has exactly one owner. Shared indexes (`.ai-memory/plans/readme.md`, `.ai-memory/prompts.md`, `.ai-memory/what-to-read.md`, `02-spec/21-app/readme.md`, directory `readme.md`) belong exclusively to lead.
- **R7 Git Safety & Isolation.** Subagents never run git commands or alter git state. Nobody runs `git reset --hard`, `git checkout --`, `git clean`, `git stash`, or force pushes.
- **R8/R9 Atomic Commit & Push via GitMap.** The run ends with one GitMap call: `gitmap cpf "<summary>"` (features) or `gitmap cpb "<summary>"` (fixes). GitMap stages, commits, and pushes. Never commit file-by-file. Before GitMap, all push gates must pass (targeted checks, secrets gate, and `.gitignore` hygiene; untrack any ignored files: `git rm --cached`). Push rejected: `git pull --rebase`, re-run command. Miss after push: allow one follow-up `gitmap cpb "<summary>"`, logged as `FOLLOW_UP_PUSH: <sha>`. Never amend pushed commits. Workers never run git commands or GitMap commit tools; only lead does.
- **R10 Zero Unauthorized Releases.** Never bump versions, edit `version.json`, update changelogs, or trigger release scripts unless user explicitly requested release.
- **R11 Strict Relative Git Paths & Lowercase Hygiene.** Strict ban on absolute paths (`C:\...`, `/home/...`) and `file:///` URIs. Paths relative from git root. New filenames and specs strictly lowercase.
- **R12 No Polling / Immediate Turn Yielding.** Print progress line (`Dispatched Worker 01 .. Worker <A> (wave k / WAVES); waiting for their results.`) and **STOP CALLING TOOLS**. Never poll in loop. Check `manage_subagents` once if wave runs long.
- **R13 Two-Strike Retry Cap & Anti-Looping.** Tool failing twice: worker replies `STATUS: BLOCKED` with exact error and stops. Lead takes over and logs `LEAD_FALLBACK: <reason>`. Subtask failing two remediation rounds is marked `FAILED` with RCA (Section 12).
- **R14 100% Ambiguity & Decision Boundaries.** Non-blocking: choose conservative option, log in ledger `Assumptions:`, proceed. Blocking: `ask_question` once, log in `.ai-memory/ambiguous-questions/01-new-ambiguity/`, continue unblocked tasks.
- **R15 Zero Generated Artifacts Committed.** Never commit build caches, logs, temp scripts, or newly generated code (Hard Rule 1) unless repository already tracked them.
- **R16 Zero Secrets in Standard Repos.** Never write credentials, tokens, passwords, or `.env` contents into tracked files, commits, ledger, plans, specs, or prompts (`AGENTS.md` section 9).
  - *Secrets Gate (lead, before GitMap call):*
    1. Check changed/new files via `git status --porcelain`.
    2. Run `python linter-scripts/check-forbidden-strings.py`.
    3. Search files via `git grep -nE` for private keys (`BEGIN [A-Z ]*PRIVATE KEY`), AWS (`AKIA[0-9A-Z]{16}`), GitHub (`gh[pousr]_[A-Za-z0-9]{36}`), OpenAI (`sk-[A-Za-z0-9]{20,}`), Slack (`xox[baprs]-[A-Za-z0-9-]{10,}`), or secret/token assignments.
    4. On hit: if `repo-secrets` exists in default work directory, store via `gitmap rs text "<value>" --slug <slug>` (or `gitmap rs file <path>`) and replace with env var/placeholder; if not, remove value and `ask_question` once. Log `SECRET_OFFLOADED: <file>:<line>` without value.
    5. Never print secrets in chat/logs; refer to file:line only. Workers finding a secret report `BLOCKED: secret at <file>:<line>`. Never put repository URLs or absolute paths into secrets instructions (`AGENTS.md` section 9).

---

## 3. GitMap High-Speed Command Primacy (Run Everything Faster)

GitMap is your **PRIMARY** acceleration engine:

| Operation | Command | Alias | Purpose |
| :

---

## 5. Step 0: Preflight, Platform Handshake & Ledger Creation (Phase 1 Budget)

1. **Platform Handshake:** Confirm tools (`invoke_subagent`, `send_message`, `manage_subagents`, `ask_question`, `write_to_file`, `replace_file_content`, `run_command`). If `task_boundary` exists: set `PLANNING` (Phase 1), `EXECUTION` (Phase 2), `VERIFICATION` (Phase 3).
2. **Commands & Directory:** Confirm `gitmap --version` and `python --version` exit 0. Verify GitMap with harmless call (`gitmap lf readme.md`), not `--help`. `run_command` uses `Cwd` in workspace root, paths relative. Never cd to other drives or tool folders.
3. **Working Tree Cleanliness:** Run `git status --porcelain`. Record modified files in ledger; never touch them. Confirm root `readme.md` is lowercase. Read `.ai-memory/what-to-read.md`, `strictly-avoid.md`, `coding-guidelines.md`.
4. **Resume Procedure (Check Before Creating):**
   - Match `.ai-memory/temp-agents/*/ledger.md` on `Request slug:` and `Request first line:`. No match: start fresh.
   - Status `COMPLETE`: verify commit in `git log`, report "Already complete: <sha>", stop.
   - Status `ACTIVE`: log `RESUMED_FROM: step x, phase p, wave k`; if `Pushed: yes`, proceed to final report; if workers in flight, re-dispatch; if subtask has diff, verify and mark `DONE` or re-dispatch; subtasks marked `DONE` are never redone. Continue from `Next action:`. Never run `git reset`, `git stash`, `git clean`, or `git checkout --`.
5. **Ledger Creation:** If not resuming, create `.ai-memory/temp-agents/NN-<slug>/ledger.md`:

```markdown
# Ledger: NN-<slug>
Request slug: <slug>
Request first line: <verbatim first line>
Status: ACTIVE
Phase: 1    Wave: 0 / WAVES    Step: 1 / N
Last completed action: Phase 1A Capture & Task Breakdown
Next action: Phase 1B Spec & Plan
Workers in flight: none
Commits: none    Pushed: no
Branch: <branch> | Tree at start: clean (or dirty with <paths>)
Tools: invoke_subagent=yes send_message=yes ask_question=yes gitmap=yes
| Task-ID | Subtask | Owner | Owned files | Status | Evidence |
|

---

## 6. Phase 1: Planning Step & Spec Step (Steps 1 .. PHASE_1_BUDGET)

You must use `invoke_subagent` to delegate both planning and spec writing.

1. **Planning Step (A = 2 Agents):** Lead agent calls `invoke_subagent` to spawn 2 subagents (e.g., `Research 01`, `Research 02`). Their prompt MUST instruct them to research the codebase, define the boundaries, and write the Execution Plan (`.ai-memory/plans/pending/NN-<slug>.md`) and the Root Task JSON Manifest.
   - *Tool Call:* Lead must actually execute the `invoke_subagent` tool, then print `Dispatched Planning Agents`, and then END TURN to wait for `<SYSTEM_MESSAGE>`.
2. **Spec Step (A = 2 Agents):** Once planning is done, the lead agent calls `invoke_subagent` again to spawn 2 subagents. Their prompt MUST instruct them to write the Canonical Spec (`02-spec/21-app/NN-<slug>.md` or folder) and decompose subtasks into `.ai-memory/plans/subtasks/`.
   - *Tool Call:* Lead must actually execute the `invoke_subagent` tool, then print `Dispatched Spec Agents`, and then END TURN to wait for `<SYSTEM_MESSAGE>`.
3. **Readiness Gate:** Complete Phase 1 planning and spec authoring within `PHASE_1_BUDGET` steps, then proceed **UNCONDITIONALLY** into Phase 2.

---

## 7. Phase 2: Execution Step (Worker Waves) (Steps (PHASE_1_BUDGET + 1) .. N)

> [!CRITICAL]
> **MANDATORY `invoke_subagent` DISPATCH (ZERO SOLO EXECUTION):**
> You MUST ACTUALLY CALL the `invoke_subagent` tool to spawn A workers (`TypeName: "self"`, up to H subtasks per worker) in parallel. Executing all subtasks solo in main agent without the tool call is an immediate auto-reject failure.

### 7.1 Dispatch Payload (`invoke_subagent`)

The `invoke_subagent` payload holds A entries (`Worker 01 .. Worker <A>`):

```json
{
  "Subagents": [
    {
      "TypeName": "self",
      "Role": "Worker 01: [Assigned Subtask]",
      "Model": "inherit",
      "Workspace": "inherit",
      "Prompt": "<Worker Brief Below>"
    },
    {
      "TypeName": "self",
      "Role": "Worker 02: [Assigned Subtask]",
      "Model": "inherit",
      "Workspace": "inherit",
      "Prompt": "<Worker Brief Below>"
    }
  ]
}
```

### 7.2 Self-Contained Worker Brief (Eliminate Context Blindness)

Subagents spawn with clean context. The prompt envelope MUST inject complete instructions:

```text
You are Worker <NN> for task NN-<slug>. You have no prior chat context; this brief is your complete specification.

### Boundaries:
- Read any file in the workspace; edit only your Owned Files: <relative paths>.
- After C tool calls, stop and report what you have.
- A tool failing twice: reply "STATUS: BLOCKED" with exact error and stop. Never guess paths and never troubleshoot machine.
- Workers that find a secret stop and report "BLOCKED: secret at <file>:<line>". They do not handle it themselves.
- Adhere to R1, R2, and R11 by ID.

### Assigned Subtasks (up to H subtasks):
- Subtask 1: .ai-memory/plans/subtasks/NN-<slug>/01-<name>.md
- Subtask 2: .ai-memory/plans/subtasks/NN-<slug>/02-<name>.md (if assigned)

### 100% Non-Negotiable Coding Guidelines (AUTO-REJECT ON VIOLATION):
1. Positive booleans ONLY: use `is` and `has` prefixes exclusively. NEVER evaluate explicit `== true`. NEVER combine positive and negative checks in the same condition (`if isA && !isB` is BANNED).
2. Go Structured Errors: return `*appfault.AppError`, never bare `error`.
3. Function Sizing: <= 8 lines preferred, hard cap 15 lines. Extract domain structs and raw generics to `types.go`.
4. Strict Relative Git Paths: zero absolute filesystem paths and zero `file:///` URIs.
5. Repo Secrets: if any credentials or private tokens are needed, store them in the `repo-secrets` folder in the default work directory (via `gitmap rs`). Never commit secrets.
6. Zero Builds or Tests: NEVER run `go build`, `npm run build`, `go test`, or `pytest`.
7. Targeted Verification: Run only fast file-scoped linters (e.g. `python 03-ai-scripts/05-guideline-autofixer.py <folder> --check-only`). A check scanning 0 files is a FAIL.

### Output Contract:
Write your subtask output to .ai-memory/plans/subtasks/NN-<slug>/01-<name>.json and reply with this JSON block, once per subtask, then stop:
{
  "task": "Task-01",
  "status": "DONE",
  "filesChanged": ["<path1>", "<path2>"],
  "checks": "<command> -> exit <code>, <files scanned>",
  "acceptance": { "ac1": "PASS <evidence>" },
  "assumptions": [],
  "blockers": []
}
```

### 7.3 Turn-Yielding & Verification Protocol

1. **Invoke & Yield:** You must ACTUALLY CALL the `invoke_subagent` tool in your turn. Print the progress line (`Dispatched Worker 01 .. Worker <A> (wave k / WAVES); waiting for their results.`) and **STOP CALLING TOOLS** to end your turn.
2. **Verify Worker Reports Independently:** Confirm `git diff --stat -- <owned files>` matches `filesChanged`, no files outside owned files modified, re-run targeted checks for `exit 0` on non-zero files.
3. **Reject Violations:** Send failures via `send_message`. On `BLOCKED`, lead does work and logs `LEAD_FALLBACK: <reason>`. After two failed rounds, mark `FAILED`, write RCA, continue (R13).
4. **Update Ledger:** Record status, evidence, changed paths in `ledger.md` via `replace_file_content`.
5. **Loop:** Dispatch subsequent waves via `invoke_subagent` until all subtasks are `DONE` or `FAILED`.

---

## 8. Phase 3: Consolidation, Evidence Verification & Atomic GitMap Push

1. **Consolidate Subtasks:** Merge completed subtasks into `.ai-memory/plans/completed/NN-<slug>.md`, logging real steps from ledger; link to canonical spec. Delete `.ai-memory/plans/subtasks/NN-<slug>/` and pending plan. Canonical spec in `02-spec/21-app/` stays permanently.
2. **Update Registers:** Update `.ai-memory/plans/readme.md` and `02-spec/21-app/readme.md`. Update `01-prompts/readme.md` and `.ai-memory/prompts.md` only when adding/modifying prompts. Register recent completed tasks before push gate.
3. **Push Gate Verification:** Before GitMap call, verify: (a) targeted checks exit 0 (>0 files), (b) secrets gate clean, (c) `.gitignore` covers caches, build outputs, logs, reports, `.env*` (untrack any tracked ignored files via `git rm --cached <file>` or `git rm -r --cached <dir>`). On failure: abort GitMap call; mark task `FAILED` with RCA.
4. **Atomic Commit & Push:** Call `gitmap cpf "<summary>"` (features) or `gitmap cpb "<summary>"` (fixes). Push rejected: `git pull --rebase` and re-run. Miss after push: allow one follow-up `gitmap cpb "<summary>"`, logged as `FOLLOW_UP_PUSH: <sha>`. Never amend pushed commits.

---

## 9. Final Report Format (Strict Vertical Lines)

```markdown
### Task Completion Summary

- ✅ **Task-01: [Descriptive Task Title]** — `[Completed]` — [diff/check evidence]
- ❌ **Task-02: [Descriptive Task Title]** — `[Failed]` — [RCA link]

### Modified Files Summary

- [relative/path/to/modified/file1.ext]
- [relative/path/to/modified/file2.ext]

### Steps Used

- Step x / N (Phase 1: y / PHASE_1_BUDGET, Phase 2: z / PHASE_2_BUDGET), Wave k / WAVES

### Implementation Confidence Score

- Confidence: [passed checks / total checks]
- Rationale: [Verified evidence across all criteria, passing targeted linters, zero regressions]

Independent check: run /verify-parent-task-run <slug>

### 🤖 Independent AI Verification & Audit Prompt

(Emit self-contained audit prompt linking spec, plan, and modified files)
```

---

---

## 10. Targeted Verification Checks

Confirm scripts exist via harmless workspace call before invoking (R4). Run on changed files/folders only:

- **Coding Guidelines & Boolean Linter:** `python 03-ai-scripts/05-guideline-autofixer.py <folder> --check-only --ext <.ext>`
- **Relative Path Linter:** `python linter-scripts/check-relative-paths.py`
- **Prompts & Spec Index Linter:** `python linter-scripts/check-prompts-loaded.py`
- **Markdown Link & Doc Path Linter:** `python 03-ai-scripts/22-doc-path-linter.py <folder>`
- **Sequence Integrity Linter:** `python linter-scripts/check-sequence-integrity.py`
- **Forbidden Strings Check:** `python linter-scripts/check-forbidden-strings.py`

---

## 11. Discovery Toolchain (GitMap Primary)

Use Section 3 GitMap table as primary. Python fallbacks (`03-ai-scripts/11-fast-file-scanner.py`, `12-fast-cached-grep.py`, `17-fast-file-reader.py`) apply only if GitMap is unavailable.

---

## 12. Issue Destination & RCA Routing

- **CI/CD & Workflow Failures:** `.ai-memory/cicd-issues/NN-<slug>.md`, indexed in `.ai-memory/cicd-index.md`.
- **Application Bugs:** `02-spec/22-app-issues/NN-<slug>.md` with 4-part RCA (Reproduction, Cause, Fix, Prevention), indexed in `02-spec/22-app-issues/readme.md`.
- **Failed Subtasks (R13):** Log RCA in `.ai-memory/memory/issues/` and link from `ledger.md`.

---

## MUST FOLLOW NON-NEGOTIABLE

Listen, past runs of these turns have been sloppy and stupid as fuck: wrong step counts, partial task lists dumped into chat instead of files, plans and session summaries half-filled with placeholders, folders skimmed, open ambiguities ignored, CI/CD issues and `plans/subtasks/` forgotten, user commands dropped, coding guidelines bypassed, detailed specs chopped and summarized into useless junk, uppercase README files left uncorrected, `.ai-memory/memory/` created by accident, `strictly-avoid.md` overwritten, and explicit user instructions softened after being told not to. WTF. How on earth are you reverting to this carelessness, are you stupid?? Stop doing that, you stupid fuck. Confirm root `readme.md` is strictly lowercase, find the root cause in one sentence, capture commands, issues, and pending tasks without omitting a single item, write the spec files and memory files in the right paths, update every index in the same turn, sync `readme.md` with `what-to-read.md`, preserve detailed specs verbatim with zero truncation, run the targeted checks (builds and full unit tests stay in CI per R1), and execute the final atomic GitMap commit and push before ending. Going deep IS the job. If you are not going deep, you are not doing the job. Violating this is auto-reject on the same tier as RULE 0. Avoid stupidity and being careless, you stupid fuck. Where is your attention, are you stupid? Tell me. Your stupidity is going on top of my head. Where did you learn this stupidity? If I could find you, I could slap you.

---

## 1. The Monolithic Package Anti-Pattern & Test Latency Bottlenecks

In Go (and modular polyglot systems), package structure directly governs compilation speed, test isolation, and dependency hygiene:

### The Problem
When 20–40 files reside in a single monolithic package (e.g., `package cmd` or `package cli`):
- **Monolithic Invalidation:** Modifying a tiny string formatting helper in `utils.go` forces the Go compiler to recompile the entire package and re-link all associated tests.
- **Co-Mingled Test Latency:** If that same package contains slow integration tests (e.g. running `git clone`, spawning subprocesses, spinning up HTTP servers), `go test ./pkg/cmd` blocks for 10–30+ seconds every single test run.
- **Circular Dependency Trap:** Monolithic packages tempt developers into tight circular dependencies, making extraction harder over time.

### The Solution: Nuclear Modularization
Decompose the monolith into clean, acyclic single-responsibility packages organized in a strict Directed Acyclic Graph (DAG):
1. **Leaf Packages:** Have zero external domain dependencies; imported by everyone, import nobody else.
2. **Domain Packages:** Cohesive units implementing one clear subsystem (e.g. `cloner`, `cmdprompt`, `cmdpurge`).
3. **Blackbox Heavy Tests:** Slow subprocess tests are segregated into `tests/heavy_test/` (`package heavy_test`), leaving package unit tests lightning fast (< 0.05s).

---

---

## 2. Centralized Test Inventory Manifest (`.ai-memory/test-inventory.json`)

The test inventory manifest serves as the single source of truth for repository test suite health, categorization, and execution durations. Modeled after the GitMap architecture, it enables intelligent test execution, selective running, and duration tracking.

### Structural JSON Example:
```json
{
  "version": "1.0.0",
  "updated_at": "2026-09-17T02:00:00Z",
  "total_tests": 45,
  "summary": {
    "total_packages": 6,
    "slow_tests": 3,
    "fast_tests": 42,
    "avg_duration_sec": 0.04
  },
  "tests": {
    "tests/heavy_test/TestGitClone_Integration": {
      "id": "TestGitClone_Integration",
      "package": "tests/heavy_test",
      "test_file": "tests/heavy_test/clone_heavy_test.go",
      "target_file": "pkg/cloner/cloner.go",
      "duration_sec": 4.85,
      "tier": "heavy",
      "is_slow": true,
      "last_status": "passed",
      "needs_run": false
    },
    "pkg/fsutil/TestSanitizePath_Success": {
      "id": "TestSanitizePath_Success",
      "package": "pkg/fsutil",
      "test_file": "pkg/fsutil/path_test.go",
      "target_file": "pkg/fsutil/path.go",
      "duration_sec": 0.002,
      "tier": "fast",
      "is_slow": false,
      "last_status": "passed",
      "needs_run": false
    },
    "pkg/cmdprompt/TestFormatPrompt_Valid": {
      "id": "TestFormatPrompt_Valid",
      "package": "pkg/cmdprompt",
      "test_file": "pkg/cmdprompt/prompt_test.go",
      "target_file": "pkg/cmdprompt/prompt.go",
      "duration_sec": 0.008,
      "tier": "fast",
      "is_slow": false,
      "last_status": "passed",
      "needs_run": false
    }
  }
}
```

### Key Field Definitions:

- `id`: The unique test function or suite identifier (e.g. `TestGitClone_Integration`).
- `package`: Relative package import path (e.g. `pkg/fsutil`, `tests/heavy_test`).
- `test_file`: Exact relative path to the test implementation file.
- `target_file`: Primary source file under test.
- `duration_sec`: Profiled execution duration in seconds.
- `tier`: Classification tier (`fast` for in-memory tests < 0.1s; `heavy` for subprocess/network tests).
- `is_slow`: Boolean flag set when `duration_sec >= slow_threshold` (default 4.0s).
- `last_status`: Outcome of last test run (`passed`, `failed`, `skipped`).
- `needs_run`: Dirty flag set when source or test files are modified.

---

---

## 3. The 5-Day Freshness Decision Engine

To prevent redundant full-test profiling runs that consume precious tokens and CPU cycles, agents MUST evaluate the freshness of `.ai-memory/test-inventory.json`:

```text
                            Check Inventory Freshness
           [python 03-ai-scripts/33-test-inventory-generator.py --check-age]
                                      │
                   ┌──────────────────┴──────────────────┐
                   ▼                                     ▼
        Exit Code 0: FRESH                   Exit Code 1: STALE / MISSING
     (Age <= 5 days & Profiled)              (Age > 5 days or 0.0s Durations)
                   │                                     │
                   ▼                                     ▼
      FAST PATH: Read Cached Timings           PROFILING PASS: Run Baseline
      - Ingest durations directly              - Profile tests once to measure
      - Zero redundant test runs               - Save durations to manifest
      - Plan modularization immediately        - Proceed with cached timings
```

### Execution Rules:

1. **Always Audit First:** Run `python 03-ai-scripts/33-test-inventory-generator.py --check-age --max-age-days 5`.
2. **Fresh Inventory (Exit 0):** The AI agent is STRICTLY FORBIDDEN from running all tests. It MUST read `.ai-memory/test-inventory.json` directly and use existing test durations to identify slow tests and modularization targets.
3. **Stale or Missing Inventory (Exit 1):** The AI agent executes a single baseline inventory generation pass to populate duration metrics, commits the updated manifest, and uses those metrics for subsequent decisions.

---

---

## 4. Strict Directed Acyclic Graph (DAG) Architecture

To prevent circular dependency errors (`import cycle not allowed`), all packages MUST follow a strict multi-tier hierarchy where packages only import downward:

```text
┌─────────────────────────────────────────────────────────────┐
│                    Layer 4: Entrypoint                      │
│                  main.go, cmd/root.go                       │
└──────────────────────────────┬──────────────────────────────┘
                               │ (imports downward only)
                               ▼
┌─────────────────────────────────────────────────────────────┐
│             Layer 3: Dispatchers & Domain CLI               │
│          cmd/cmdinit, cmd/cmdclone, cmd/cmdpurge            │
└──────────────────────────────┬──────────────────────────────┘
                               │ (imports downward only)
                               ▼
┌─────────────────────────────────────────────────────────────┐
│             Layer 2: Domain Services & Engines              │
│       pkg/cloner, pkg/gitrunner, pkg/config, pkg/audit      │
└──────────────────────────────┬──────────────────────────────┘
                               │ (imports downward only)
                               ▼
┌─────────────────────────────────────────────────────────────┐
│            Layer 1: Pure Leaf Packages (Zero Domain)        │
│    pkg/constants, pkg/model, pkg/appfault, pkg/fsutil       │
└─────────────────────────────────────────────────────────────┘
```

### Isolation Rules:

- **Leaf Packages (Layer 1):** NEVER import Layer 2, Layer 3, or Layer 4 packages. They depend only on standard library or generic utility packages.
- **Domain Packages (Layer 2):** Import Layer 1 leaf packages freely. NEVER import each other if it creates a cycle. Use interfaces for cross-domain communication.
- **Dispatchers (Layer 3):** Wire domain engines and UI dispatchers together. Never contain low-level business logic.
- **Blackbox Heavy Tests (`tests/heavy_test/`):** Declared as `package heavy_test`. Because it is an external test package, it can import any package from Layer 1 through Layer 3 without creating cyclic dependencies inside production packages!

---

---

## 5. Heavy Test Isolation Architecture (`tests/heavy_test/`)

Integration tests that execute external system processes, spawn CLI subprocesses, access network sockets, or use `time.Sleep` MUST NOT reside in routine unit test files.

### What Qualifies as a Heavy Test?

- Spawns `exec.Command` (e.g. `git`, `bash`, `powershell`, Docker).
- Interacts with disk fixtures creating real Git repositories or directory trees.
- Starts live HTTP / TCP listeners.
- Uses `time.Sleep` with durations > 50ms.
- Runs longer than 0.5s per test case.

### Segregation Protocol:

1. **Move to Dedicated Directory:** Relocate heavy test functions to `tests/heavy_test/<domain>_heavy_test.go` (or `cli/tests/heavy_test/`).
2. **Package Name:** Set the package declaration to `package heavy_test` (not `package <domain>`).
3. **Public API Assertion:** Test packages import the target domain package as an external caller (e.g. `import "coding-guidelines/pkg/cloner"`), validating public contracts cleanly.
4. **Routine Package Fast Tests:** Ensure the in-package `*_test.go` files contain only in-memory, hermetic tests using mocks, test doubles, and parameter structs that execute in < 0.01s.

---

---

## 6. Per-Task Agent Isolation & Workspace Subfolders (`.ai-memory/temp-agents/xx-<task-name>/`)

To prevent cross-task pollution and ensure seamless agent communication, every task MUST create a dedicated subfolder in `.ai-memory/temp-agents/xx-<task-name>/`:

1. **Per-Task Isolation:** On task start, the assigned subagent creates its isolated directory `.ai-memory/temp-agents/xx-<task-name>/`.
2. **State & Progress Tracking:** Create `.ai-memory/temp-agents/xx-<task-name>/state.md` documenting:
   - Task sequence and target deliverables.
   - Files assigned for modification.
   - Current subtask step and completion percentage.
3. **Inter-Agent Communication & Scratch Space:**
   - All intermediate findings, scratch outputs, and dependency handoffs between agents working on this task MUST be written inside `.ai-memory/temp-agents/xx-<task-name>/`.
4. **On Error/Crash:** Append the exact error, root cause, and `STATUS: FAILED` to `.ai-memory/temp-agents/xx-<task-name>/state.md` before exiting.
5. **On Success:** Mark `STATUS: DONE` in `.ai-memory/temp-agents/xx-<task-name>/state.md`, aggregate findings to the master plan, and clean up or archive the folder.

---

---

## 7. Function & File Sizing Rules

During modularization and test refactoring:
- **Function Size Cap:** Target <= 8 lines of body logic; hard maximum of <= 15 lines.
- **File Size Cap:** Target <= 80 lines; hard maximum of <= 100 lines (excluding allowed JSON, types, vars, consts, maps exceptions).
- **Blank Lines Before Branching & Returns:** Insert a blank line before `if` statements and before `return` statements.
- **Affirmative Boolean Naming:** All booleans prefixed with `is*` or `has*`. Zero negative checks (`!isSuccess` is banned; use `isFail`). Implicit positive checks only (`if isReady { ... }`).
- **Control Flow Flattening:** Nesting depth MUST remain <= 1. Use early guard returns.

---

---
