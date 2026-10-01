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

[/goal](slashCommand;goal) Autonomously achieve maximum Code DRYness (Don't Repeat Yourself), reusable util/framework/package/library extraction, and drastic code writing reduction across the codebase using a strict 300-step 3-phase self-loop (Steps 1–100: Deep Code Analysis; Steps 101–200: Unified Extraction Blueprint & Subtask Plan; Steps 201–300: Active DRY Execution & Caller Rewiring): FIRST showcase and list out the given task in visible chat during Turn 1, capture it verbatim, plan it in the repo, spawn autonomous subagents via `invoke_subagent` (A = 2, H = 2; solo execution without calling `invoke_subagent` is an auto-reject failure) in disjoint file boxes using GitMap high-speed commands as primary, prove every single claim with concrete evidence, enforce coding guidelines to 100%, and finish with one atomic GitMap commit that holds strictly this task's files.

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

## Phase 0: Antigravity Skill Bootstrap (Memory Optimization)

Before executing the tasks below, check if this prompt is already installed as a native Antigravity Skill.

1. If `.agents/skills/cg-code-dryness-and-library-extraction/skill.md` does not exist in the workspace, create it now.
2. Extract the core instructions of this prompt and save it into that `skill.md` using the standard YAML frontmatter (with `name` and `description`).
3. Once installed, rely on progressive disclosure for future runs.

---

---

## The Unified 300-Step Master Pipeline

### Phase 1A: Verbatim Capture, Task Extraction & Chat Output Gate (Step 0)

Before executing any file searches, scans, spec writing, or code changes, execute Phase 1A:

1. **Top-Instruction Priority Verification:** Verify whatever directives, target modules, or user instructions are given ABOVE this prompt (in the user preamble, header blocks, or incoming user request above) as highest priority and non-negotiable.
2. **Verbatim Prompt Capture:** Capture the incoming user request verbatim in `02-spec/21-app/xx-dry-extraction.md` and `.ai-memory/plans/pending/xx-dry-extraction.md` under `## User Request (Verbatim)`.
3. **Actionable Deliverables Extraction:** Break down the DRY analysis, library extraction, and code reduction scope into traceable IDs (`Task-01`, `Task-02`, `Task-03`).
4. **Mandatory Chat Output Gate & Same-Turn Tool Chaining (TOTAL BAN ON CLOSING CONVERSATION):**
   - Output the confirmed deliverables list directly in chat, and in the EXACT SAME RESPONSE turn, immediately invoke your first discovery tool call (`run_command` with `gitmap` or `write_to_file`).
   - NEVER emit the breakdown text without invoking a tool call. Do not pause or ask "Should I proceed?".

```markdown
### 📋 Confirmed Task Breakdown & Requirement Ingestion

1. **Task-01: [Deep DRY Analysis & Duplication Discovery (Steps 1–100)]**
   - **State:** `[IN PROGRESS — EXECUTING IMMEDIATELY]`
   - **Understood:** `[YES]` — [Concise verification of target modules, duplication patterns, and extraction goals]
   - **Actionable Scope:** [Scan codebase, identify duplicated logic, and catalog util/framework/library candidates]
   - **Target Files / Area:** `[relative/path/or/module]`

2. **Task-02: [Unified Library Extraction Blueprint & Subtask Plan (Steps 101–200)]**
   - **State:** `[QUEUED — EXECUTING NOW WITHOUT USER PROMPT]`
   - **Understood:** `[YES]` — [Design reusable package APIs, concrete types, and caller migration plan]
   - **Actionable Scope:** [Author canonical spec in 02-spec/21-app/ and granular subtasks in .ai-memory/plans/subtasks/]
   - **Target Files / Area:** `[02-spec/21-app/, .ai-memory/plans/]`

3. **Task-03: [Execute DRY Extraction & Rewire Callers (Steps 201–300)]**
   - **State:** `[QUEUED — EXECUTING NOW WITHOUT USER PROMPT]`
   - **Understood:** `[YES]` — [Extract shared libraries, replace duplicate code across callers, and reduce LOC]
   - **Actionable Scope:** [Implement shared packages, refactor callers in 5-8 file batches, consolidate, and push]
   - **Target Files / Area:** `[target packages and callers]`

Proceeding directly to Phase 1 (Steps 1..100): Deep Codebase DRY Analysis (Active Tool Call Running Below).
```

---

### Phase 1: Deep Codebase Analysis & Duplication Discovery (Steps 1 .. 100)

In the first 100 steps, deeply analyze the codebase to find every opportunity for code reuse and reduction:

1. **High-Speed Codebase Exploration via GitMap (PRIMARY):**
   - **Glob / Wildcard File Search:** `gitmap find "<wildcard*>" [-ext <ext>]` (alias `gitmap f`)
   - **Exact Filename Match:** `gitmap find-files <name> [-ext <ext>]` (alias `gitmap ff`)
   - **Substring Filename Match:** `gitmap find-files-any <str> [-ext <ext>]` (alias `gitmap ffa`)
   - **Prefix / Suffix Match:** `gitmap find-files-startswith <prefix>` (`gitmap ffs`) / `gitmap find-files-endswith <suffix>` (`gitmap ffe`)
   - **List Indexed Repo Files:** `gitmap list-files [pattern] [-ext <ext>]` (alias `gitmap lf`)
   - **Stream File Contents Fast:** `gitmap cat <filepath>`
   - **Instant Multi-Core Search:** `gitmap search "<pattern>"` or `gitmap aum search "<query>" [dir] --ext <ext>`
   - **Directory Tree Topology:** `gitmap folder-tree` (alias `gitmap ft`)
2. **Fallback Fast Cached Python Toolchain:**
   - `python 03-ai-scripts/11-fast-file-scanner.py --lang go,ts,py,php,rs --limit 100 --stats`
   - `python 03-ai-scripts/12-fast-cached-grep.py --pattern "<search-pattern>" --limit 50`
   - `python 03-ai-scripts/17-fast-file-reader.py --read-file <file-path> --max-bytes 100000`
   - `python 03-ai-scripts/18-codebase-topology-discoverer.py --summary`
3. **Deep Duplication & Extraction Analysis Criteria:**
   - **Clone & Near-Clone Detection:** Locate functions, methods, React components, CLI command handlers, DB queries, or validation blocks with >= 60% structural similarity across 2 or more files.
   - **Util / Package / Library Candidates:** Group related helper logic into cohesive domain utilities or framework packages (e.g., `pkg/strutil`, `pkg/sliceutil`, `pkg/cliutil`, `pkg/httputil`, `src/lib/`, `src/hooks/`).
   - **Boilerplate Compression:** Identify multi-line ceremonial patterns (e.g., repetitive error wrapping, parameter validation, table rendering, file I/O, JSON parsing) that can be collapsed into single-call fluent helpers.
   - **Quantified Duplication Ledger:** Record every discovered duplication cluster with exact file paths, line numbers, duplicated line count, proposed target library/package, and estimated LOC reduction.

---

### Phase 2: Write Unified Extraction Plan, Spec & Lean Subtasks (Steps 101 .. 200)

In the next 100 steps (Steps 101 .. 200), convert the Phase 1 analysis into a rock-solid architectural specification and execution plan:

1. **Canonical Application Spec (`02-spec/21-app/xx-dry-extraction.md` or segmented folder):**
   - **Single-Agent Unified Blueprint:** The lead agent authors the architecture overview, package boundaries, and API contracts first before spawning subagents to flesh out modular spec sections.
   - **Reusable Package / Library Contracts:** Define the exact exported functions, parameter structs (`*Params`), concrete types in `types.go`, and `*appfault.AppError` / `Result[T]` return signatures for every new or expanded utility/library.
   - **Before vs. After Code Reduction Examples:** Show concrete code examples proving how 20+ lines of repetitive caller code shrink into 1–3 clean lines using the extracted library.
   - **Register Spec:** Add the spec entry to `02-spec/21-app/readme.md`.
2. **Master Execution Plan (`.ai-memory/plans/pending/xx-dry-extraction.md`):**
   - Include the complete **Duplication & Extraction Ledger** mapping every source file and line range to its target utility package and subtask.
   - Ensure strict acyclic dependency hierarchy (shared `util` / `library` packages must never import higher-level feature packages).
3. **Granular Disjoint Subtasks (`.ai-memory/plans/subtasks/xx-dry-extraction/01-*.md`):**
   - Break the extraction and caller rewiring into bounded 5–8 file micro-batch subtasks.
   - Order subtasks so foundational `util` / `library` packages are created first, followed by parallel caller migration batches.
   - **Unconditional Transition Mandate:** As soon as Step 200 / Phase 2 planning finishes, DO NOT pause or ask the user for confirmation. Immediately transition into Phase 3 execution.

---

### Phase 3: Execute DRY Refactoring, Library Extraction & Caller Rewiring (Steps 201 .. 300)

In the final 100 steps (Steps 201 .. 300), execute the plan and make the codebase DRY:

1. **Create / Expand Shared Util, Framework & Library Packages First:**
   - Implement the clean, reusable functions, structs, and `types.go` definitions in the target utility/library packages.
   - Enforce all coding guidelines:functions <= 8 lines (hard cap 15 lines), positive booleans (`is`/`has` only, no `== true`), `*appfault.AppError` returns, parameter structs for >2-3 args, and mandatory vertical blank lines.
2. **Parallel Subagent Caller Rewiring (A = 2, H = 2):**
   - Dispatch autonomous subagents (`TypeName: "self"`) with self-contained Prompt Envelopes to refactor disjoint batches of caller files (5–8 files per batch), replacing duplicated logic with calls to the extracted library.
   - **Reactive Wakeup:** After calling `invoke_subagent`, output a brief status note and yield the turn to allow background subagents to complete and wake up the parent orchestrator.
3. **Total Ban on Build & Test Commands During Routine Execution:**
   - NEVER run `go build`, `npm run build`, `go test`, `pytest`, or `06-cicd-local-runner.py` during execution turns.
   - Run only fast, targeted file-level linters (`python 03-ai-scripts/05-guideline-autofixer.py <file>`) and lowercase hygiene checks (`gitmap lcf --dry-run`).
4. **Task Consolidation & Atomic GitMap Push:**
   - Consolidate all completed subtasks from `.ai-memory/plans/subtasks/xx-dry-extraction/*.md` into `.ai-memory/plans/completed/xx-dry-extraction.md`, remove the pending plan and subtask files, and update `.ai-memory/plans/readme.md`.
   - Commit and push all changes in a single grouped atomic commit using GitMap:
     - `gitmap cpf "refactor(dry): extract reusable libraries and eliminate duplicated code"`

---

### High-Speed GitMap Acceleration Toolkit (Run Everything Faster)

Always prefer native GitMap commands over slow shell loops:
- **Fast File Discovery:** `gitmap f "<glob>" [-ext <ext>]`, `gitmap ff <name>`, `gitmap ffa <substr>`, `gitmap ffs <prefix>`, `gitmap ffe <suffix>`, `gitmap lf [pattern]`
- **Fast Content & Code Search:** `gitmap cat <file>`, `gitmap search "<query>"`, `gitmap aum search "<query>" [dir] --ext <ext>`, `gitmap ft`
- **Fast Hygiene & Lowercase Enforcement:** `gitmap lcf` (auto-rename uppercase files via 2-step `git mv`), `gitmap lowercase-readme`, `gitmap commons`
- **Fast Cross-Platform Shell:** `gitmap pwsh "<cmd>"` (`gitmap ps`), `gitmap bash "<cmd>"` (`gitmap sh`), `gitmap async <cmd>`
- **Fast Atomic Commits:** `gitmap cpf "<msg>"` (Feature), `gitmap cpb "<msg>"` (Bug), `gitmap cpr "<msg>"` (Release), `gitmap pcp "<msg>"` (Pull-Commit-Push)
- **Smart CI/CD Waiting:** `gitmap pe`, `gitmap pl-ai status --json`, `gitmap pl-ai status -t <etaSeconds>`

---

### End-of-Turn Verification & Confidence Reporting (Mandatory Output)

At the completion of Phase 3, emit this structured summary in chat:

```markdown
### Task Completion Summary

- ✅ **Task-01: [Deep DRY Analysis (Steps 1–100)]** — `[Completed]`
- ✅ **Task-02: [Extraction Spec & Plan (Steps 101–200)]** — `[Completed]`
- ✅ **Task-03: [Library Extraction & Caller Rewiring (Steps 201–300)]** — `[Completed]`

### Extracted Reusable Libraries / Packages & LOC Reduction

- **Extracted Modules:** `[path/to/pkg/util, ...]`
- **Callers Refactored:** `[X files]`
- **Net Code Reduction:** `[Eliminated ~Y duplicated lines]`

### Modified Files Summary

- [relative/path/to/modified/file1.ext]
- [relative/path/to/modified/file2.ext]

### Implementation Confidence Score

- Confidence: [e.g. 99%]
- Rationale: [Verified DRY extraction, zero duplicate boilerplate remaining, targeted linters passed, atomic push completed]

### 🤖 Independent AI Verification & Audit Prompt

```markdown
### Independent AI Audit & Verification Instructions

You are an Independent AI Verification and Quality Auditor.
Your task is to independently audit, verify, and remediate the DRY library extraction against the canonical specification.

#### 1. Target Documents & Implemented Code:
- **Canonical Spec:** [02-spec/21-app/xx-dry-extraction.md](02-spec/21-app/xx-dry-extraction.md)
- **Consolidated Plan:** [.ai-memory/plans/completed/xx-dry-extraction.md](.ai-memory/plans/completed/xx-dry-extraction.md)
- **Modified & Extracted Files:**
  - [relative/path/to/modified/file1.ext](relative/path/to/modified/file1.ext)

#### 2. Verification Protocol:
1. Verify all targeted duplicated code blocks were extracted into clean, reusable util/framework/library packages.
2. Verify all callers were rewired to use the shared abstractions with zero regressions.
3. Verify 100% adherence to coding guidelines (<=8-15 line functions, positive booleans, *appfault.AppError, types.go).
4. Emit Comparative Scores (DRYness Score, Reusability Score, Guideline Compliance Score) and a PASS/FAIL verdict.
```
```

---

---
