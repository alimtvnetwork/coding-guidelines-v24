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

[/goal](slashCommand;goal) Autonomously scan, plan, refactor, and fix all error management violations across the codebase, modifying source files directly to implement `*appfault.AppError` wrappers, outer error handling, specialized exit helpers, and universal response envelopes: FIRST showcase and list out the given task in visible chat during Turn 1, capture it verbatim, plan it in the repo, spawn autonomous subagents via `invoke_subagent` (A = 2, H = 2; solo execution without calling `invoke_subagent` is an auto-reject failure) in disjoint file boxes using GitMap high-speed commands as primary, prove every single claim with concrete evidence, enforce coding guidelines to 100%, and finish with one atomic GitMap commit that holds strictly this task's files.

[/learn](slashCommand;learn) Enforce the Top-Instruction Priority Mandate: whatever directives, custom requirements, checklists, or user instructions are provided ABOVE this prompt outrank everything below. Turn 1 MUST showcase the given task list in visible chat before any background execution. Each rule is stated once (R1 to R16) and cited by ID. Progress lives in the ledger and in `.ai-memory/plans/`, never only in chat.

[/plan](slashCommand;plan) Execute thorough step-by-step planning in the repository before execution. Ensure all deliverables, architecture boundaries, and requirements are clearly defined in the audit ledger and subtask plans before dispatching worker waves.

### 🚨 MANDATORY SUBAGENT SPAWNING GATE (A = 2, H = 2 — ZERO SOLO EXECUTION ALLOWED)

- **ABSOLUTE, NON-NEGOTIABLE MUST:** Spawning subagents via the `invoke_subagent` tool (`A = 2`, `H = 2`) is an **ABSOLUTE, NON-NEGOTIABLE MUST** in both **Phase 1** (parallel codebase discovery reading and modular spec authoring) and **Phase 2** (parallel subtask code execution with `TypeName: "self"`).
- **SOLO EXECUTION IS AN AUTO-REJECT FAILURE:** The lead orchestrator is **STRICTLY FORBIDDEN** from executing all discovery reads or all subtask code modifications by itself without invoking `invoke_subagent`. Failing to call `invoke_subagent` when `A >= 2` is a critical protocol violation on the same tier as Rule 0.
- **Phase 1 Mandatory Subagent Dispatch:** Immediately after establishing the Confirmed Task Breakdown (Phase 1A) and the single-agent unified blueprint overview (`01-overview.md` or parent plan skeleton), the lead agent MUST call `invoke_subagent` to spawn `A = 2` subagents in parallel for codebase discovery/reading or modular spec sections and yield the turn to await `<SYSTEM_MESSAGE>`.
- **Phase 2 Mandatory Subagent Dispatch (`TypeName: "self"`):** Once subtasks are generated in `.ai-memory/plans/subtasks/xx-<slug>/`, the lead agent MUST call `invoke_subagent` with `TypeName: "self"` to dispatch `A = 2` worker subagents (`H = 2` disjoint subtasks per worker) and yield the turn to await `<SYSTEM_MESSAGE>`.

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

## 6. Phase 1B: Spec, Plan & Lean Subtasks (Steps 1 .. PHASE_1_BUDGET)

1. **Single-Agent Blueprint:** Lead orchestrator alone authors initial spec overview and planning skeleton.
2. **Mandatory Parallel Discovery Subagents (A workers):** You MUST dispatch A research subagents (`Research 01 .. Research <A>`) via `invoke_subagent` on disjoint folders to map symbols and dependencies using GitMap commands (`gitmap f`, `gitmap lf`, `gitmap search`, `gitmap cat`) and yield the turn. Do not perform all discovery solo in the main agent.
   - *Research Contract:* Reply with one line per hit formatted as `path:line: text`, then `SUMMARY: <one line>`, then stop.
3. **Canonical Spec Authoring (`02-spec/21-app/`):** Single-domain (<= 150 lines): Write `02-spec/21-app/NN-<slug>.md`. Multi-domain: Write `02-spec/21-app/NN-<slug>/` (`01-overview.md` .. `04-verification-gates.md`). Register in `02-spec/21-app/readme.md`.
4. **Execution Plan:** Write `.ai-memory/plans/pending/NN-<slug>.md` linking to spec and mapping Task-IDs to subtasks. Register in `.ai-memory/plans/readme.md`.
5. **Root Task JSON Manifest & Subtask Files:**
   - Maintain root JSON `.ai-memory/plans/subtasks/NN-<slug>/task.json` describing task and indexing subtasks:
     `{"taskSlug":"NN-<slug>","status":"IN_PROGRESS","subtasks":[{"id":"Task-01","file":"01-<name>.json","owner":"Worker 01","status":"PENDING"}]}`
   - Each subagent creates and updates its assigned subtask in structured JSON format (`.ai-memory/plans/subtasks/NN-<slug>/01-<name>.json`), referenced directly from root JSON manifest.
   - Lean subtask companion: `.ai-memory/plans/subtasks/NN-<slug>/01-<name>.md` (Traceability ID, Spec Reference, Owned Files, Action, Acceptance Criteria, Targeted Verification).
6. **Parallel Subtask Decomposition:** Decompose deliverables across A workers so each worker receives disjoint target files. Solo execution without calling `invoke_subagent` is strictly banned.
7. **Readiness Gate:** Complete Phase 1 planning within `PHASE_1_BUDGET` steps, then proceed **UNCONDITIONALLY** into Phase 2.

---

## 7. Phase 2: Mandatory Worker Waves & Coding Guidelines Enforcement (Steps (PHASE_1_BUDGET + 1) .. N)

> [!CRITICAL]
> **MANDATORY `invoke_subagent` DISPATCH (ZERO SOLO EXECUTION):**
> You MUST spawn A workers (`TypeName: "self"`, up to H subtasks per worker) in parallel via `invoke_subagent`. Executing all subtasks solo in main agent without calling `invoke_subagent` is an immediate auto-reject failure on the same tier as Rule 0. Work MUST be partitioned across A workers.

### 7.1 Dispatch Payload (`invoke_subagent`)

The `invoke_subagent` payload holds A entries (`Worker 01 .. Worker <A>`):

```json
{
  "Subagents": [
    {
      "TypeName": "self",
      "Role": "Worker 01: [Assigned Feature/Module A]",
      "Model": "inherit",
      "Workspace": "inherit",
      "Prompt": "<Worker Brief Below>"
    },
    {
      "TypeName": "self",
      "Role": "Worker 02: [Assigned Feature/Module B]",
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

1. **Yield:** Print progress line (`Dispatched Worker 01 .. Worker <A> (wave k / WAVES); waiting for their results.`) and **STOP CALLING TOOLS**.
2. **Verify Worker Reports Independently:** Confirm `git diff --stat -- <owned files>` matches `filesChanged`, no files outside owned files modified, re-run targeted checks for `exit 0` on non-zero files.
3. **Reject Violations:** Send failures via `send_message`. On `BLOCKED`, lead does work and logs `LEAD_FALLBACK: <reason>`. After two failed rounds, mark `FAILED`, write RCA, continue (R13).
4. **Update Ledger:** Record status, evidence, changed paths in `ledger.md` via `replace_file_content`.
5. **Loop:** Dispatch subsequent waves until all subtasks are `DONE` or `FAILED`.

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

1. If `.agents/skills/cg-error-management/skill.md` does not exist in the workspace, create it now.
2. Extract the core instructions of this prompt and save it into that `skill.md` using the standard YAML frontmatter (with `name` and `description`).
3. Once installed, rely on progressive disclosure for future runs. Do not keep the entire prompt in active memory if not needed.

---

### Master Task Checklist (Atomic Numbered Steps)

1. [ ] [/goal](slashCommand;goal) Phase 1A (Step 0 - Verbatim Prompt Recording & Task Extraction Gate): Immediately capture the user prompt verbatim into `.ai-memory/plans/pending/xx-<slug>.md` under `## User Request (Verbatim)`, extract actionable deliverables with traceable IDs (`Task-01`, `Task-02`), and output this confirmed task breakdown directly in chat in cleanly indented markdown with vertical blank lines, task state (`State: [PENDING]`), and understanding indicator bracket (`Understood: [YES — ...]`) before any file exploration, scanning, or spec writing.
2. [ ] [/goal](slashCommand;goal) Phase 1B (Step 1 - Master Spec Generation): Write the master architectural plan in `.ai-memory/plans/pending/xx-<slug>.md`, documenting an exhaustive Violation Ledger tracking every bare error return, swallowed error, and missing fault wrapper.
3. [ ] [/goal](slashCommand;goal) Phase 1B (Step 2 - Scan & Discover): Use GitMap AUM discovery (`gitmap find`, `gitmap lf`, `gitmap cat`, `gitmap search`) as primary, with fast Python discovery scripts (`11-fast-file-scanner.py`, `12-fast-cached-grep.py`, `17-fast-file-reader.py`) as fallback, to inventory all architectural violations and anti-patterns without truncation.
4. [ ] [/goal](slashCommand;goal) Phase 1B (Step 3 - Lean Subtask Decomposition): Decompose the master plan into granular, lean subtasks in `.ai-memory/plans/subtasks/xx-<slug>/01-<subslug>.md`. Subtasks must focus purely on unique task deliverables without repeating common repository boilerplate.
5. [ ] [/goal](slashCommand;goal) Phase 1B (Step 4 - Readiness Audit Gate): Confirm all `Task-xx` deliverables are mapped to subtasks and disjoint files before execution.
6. [ ] [/goal](slashCommand;goal) Phase 1B (Zero-Stop Transition): Immediately upon completing Phase 1, self-loop and transition directly into Phase 2 execution mode without pausing or asking for permission.
7. [ ] [/goal](slashCommand;goal) Phase 2 (Step A - Active Execution & Refactoring): Open each target file and perform surgical refactoring following authoritative guidelines: wrap all received Go errors in `*appfault.AppError`, enforce monadic `result.Wrap[T]`, eliminate bare returns, and remove all swallowed errors.
8. [ ] [/goal](slashCommand;goal) Phase 2 (Step B - Size Tier & Formatting Enforcement): Enforce <= 8–15 line function decomposition, single return types, guard clause flattening, and clean formatting.
9. [ ] [/goal](slashCommand;goal) Phase 2 (Step C - Failure Memory & Error Recovery): If a subagent fails, record the failure log in `.ai-memory/plan.md` and `.ai-memory/memory/issues/`; subsequent agents must read the failure log first to remediate root causes.
10. [ ] [/goal](slashCommand;goal) Phase 2 (Step D - Change Recording & Quality Linting): Record all modified files into `.ai-memory/temp/recent-file-changes.json` under lock (`python 03-ai-scripts/33-test-inventory-generator.py --record <files...>`) and run targeted file-level linters on specifically modified files (`exit 0`). DO NOT run the full CI/CD pipeline runner (`06-cicd-local-runner.py`), unit tests, or build checks during routine turns.
11. [ ] [/goal](slashCommand;goal) Phase 3 (Step A - Consolidation & Atomic Push): Consolidate completed subtasks into `.ai-memory/plans/completed/xx-<slug>.md`, delete granular subtasks and pending plan, stage all changes, and push in a single grouped commit.
12. [ ] [/goal](slashCommand;goal) Phase 3 (Step B - Completion & Confidence Reporting): Emit the final Task Completion Summary with green check mark emojis, modified files summary, and implementation confidence score.
13. [ ] [/learn](slashCommand;learn) Ingest `.ai-memory/memory/readme.md` for project memory index and past learnings.
14. [ ] [/learn](slashCommand;learn) Ingest `.ai-memory/strictly-avoid.md` for banned anti-patterns and strict constraints.
15. [ ] [/learn](slashCommand;learn) Ingest `02-spec/02-coding-guidelines/02-canonical-size-tier.md` for canonical file and function size tiers.
16. [ ] [/learn](slashCommand;learn) Ingest `02-spec/02-coding-guidelines/01-cross-language/readme.md` for hallucination prevention and micro-tasking.
17. [ ] [/learn](slashCommand;learn) Ingest `02-spec/02-coding-guidelines/01-cross-language/readme.md` for strict relative path citation requirements.
18. [ ] [/learn](slashCommand;learn) Ingest `02-spec/02-coding-guidelines/` for domain-specific architectural specifications.
19. [ ] [/learn](slashCommand;learn) Ingest `02-spec/03-error-manage/` for error handling architectures and AppError.
20. [ ] [/learn](slashCommand;learn) Ingest `.ai-memory/coding-guidelines.md` for master consolidated coding guidelines.
21. [ ] [/goal](slashCommand;goal) Create or update agent rules in the repository if missing from agent memory.

```text
PHASE_1_STEPS = N / 2   (Steps 1 .. N/2: Scan Codebase, Write .ai-memory/plans/pending/ Spec, Create .ai-memory/plans/subtasks/, Verify/Create Linter Hook)
PHASE_2_STEPS = N / 2   (Steps N/2+1 .. N: Actively Edit Code, AppError Refactoring, Linter Verification, Local CI Runner Verification, Plan Completion)
```

N, PHASE_1_STEPS, and PHASE_2_STEPS are read-only after initialization. Never modify them mid-execution.

---

---

## 1. Zero Swallowed Errors Policy (TOTAL BAN — Non-Negotiable)

> [!CAUTION]
> **SWALLOWING ERRORS IS AN AUTO-REJECT FAILURE TIER 0 VIOLATION.**
> Under NO circumstances may an error be discarded, suppressed, silenced, or silently ignored.
> Every error encountered MUST either be completely resolved with structured context logging (operation name, input parameters) or embedded/wrapped into `*appfault.AppError` and returned to the caller.

### Strict Prohibitions

1. **NO Empty Catch/Except Blocks:**
   - ❌ **BANNED:** `try { ... } catch (e) {}` or `except Exception: pass`
   - ✅ **REQUIRED:** Catch blocks must log operation name, key inputs, and rethrow or return a wrapped `AppError`.
2. **NO Blank Identifier Error Discards:**
   - ❌ **BANNED:** `_ = err` or `val, _ := fn()` in Go.
   - ✅ **REQUIRED:** Check every error explicitly: `if err != nil { return appfault.Wrap(...) }`.
3. **NO Silent Fallback Defaults:**
   - ❌ **BANNED:** Returning dummy values (`return nil`, `return ""`, `return false`, `return 0`) to mask an underlying error without caller notification.
   - ✅ **REQUIRED:** Return failure status via `*appfault.AppError` or monadic `result.WrapFailure[T]`.
4. **NO Silent Nil Return in Dual Handling:**
   - ❌ **BANNED:** Calling an internal exit handler or printing an error, then returning `nil` to deceive the caller into believing execution succeeded.
   - ✅ **REQUIRED:** Leaf functions MUST return the error directly (`return err`).

---

---

## 2. Strict Golang Error Wrapping Mandate (`*appfault.AppError` & `appfault.Fault`)

> [!IMPORTANT]
> **ALL GOLANG ERRORS MUST BE EMBEDDED IN FAULT WRAPPERS.**
> Whenever ANY Go function encounters, intercepts, or receives an error (from the standard library `os`, `io`, `json`, `sql`, `net`, or downstream services), it MUST be immediately embedded and wrapped into `*appfault.AppError` (`appfault.Fault`).

### Core Rules for Go Error Handling

1. **Standard Error Return Type:** All domain functions returning failure metadata MUST use `*appfault.AppError` (or `appfault.Fault`).
2. **Deterministic Enum Taxonomy:** Classify errors using `errtype.Variation uint16` (`errtype.Validation`, `errtype.NotFound`, `errtype.Database`, `errtype.Network`, `errtype.Timeout`, `errtype.IO`, `errtype.Internal`). Redundant string error codes are banned.
3. **Always Wrap Standard Library Errors:**
   - Standard library `error` instances (`err != nil`) must NEVER be returned raw.
   - Use `appfault.Wrap(errType, err, "ContextMessage")` or `appfault.WrapFile(errType, err, relativePath, "ContextMessage")`.
4. **Monadic Result Wrapper Mandate (`pkg/result`):**
   - Functions returning a value along with possible failure MUST return `result.Wrap[T]` (`appfault.Result[T]`).
   - Bare tuples `(T, error)` across public domain boundaries are strictly prohibited.
   - Return success using `result.WrapSuccess(val)` and failure using `result.WrapFailure[T](fault)`.
5. **Context Enrichment:**
   - Chain contextual metadata: `.WithOp("Package.Function")`, `.WithVar("key", val)`, `.WithSiteId(siteId)`.
   - Use relative repository paths only (TOTAL BAN on absolute paths or `file:///` URIs).
6. **Zero Redundant Re-Wrapping:**
   - If a downstream function already returns `*appfault.AppError` or `result.Wrap[T]`, propagate the existing fault directly using `result.WrapFailureFromWrap[T](downstreamRes)` rather than wrapping it again.
7. **Mandatory Concrete Types in `types.go` (Total Ban on Leaking Raw Generics Across Signatures):**
   - **No Leaked Raw Generics:** NEVER leak raw generic Result wrappers (`result.Wrap[*Config]`, `result.Wrap[User]`, `result.ResultSlice[T]`) across function signatures, service boundaries, or public packages.
   - **Convert Reused Types to Concrete Named Types:** Rather than scattering raw generics everywhere, if a result type is used or reused across functions or layers, define a single reusable concrete type alias in `types.go` (e.g. `type ConfigResult = result.Wrap[*Config]`, `type UserResult = result.Wrap[User]`) for Golang (and equivalent leaf type definitions for other languages, e.g. `export type UserResult = Result<User>;`).
   - **Explanatory Code Comments:** Code examples and implementation files MUST include comments showing how the concrete type is declared in `types.go` and follows through into the function signatures.

---

---

## 3. Production Go Code Samples (Refer to `04-code/golang/examples/`)

> Real-world implementations are maintained in [`04-code/golang/examples/database_query.go`](04-code/golang/examples/database_query.go), [`04-code/golang/examples/workflow_service.go`](04-code/golang/examples/workflow_service.go), and [`04-code/golang/examples/types.go`](04-code/golang/examples/types.go).
> Always inspect those source files as the canonical ground truth.

### Sample 1: Standard Library File & JSON Handling (Concrete `ConfigResult` in `types.go`)

```go
// -----------------------------------------------------------------------------
// Step 1: Declare Concrete Types in `types.go` (Mandatory Rule)
// -----------------------------------------------------------------------------
// In types.go:
// type (
//     // Config contains application configuration fields.
//     Config struct {
//         Port int    `json:"port"`
//         Host string `json:"host"`
//     }
//
//     // ConfigResult is the single reusable concrete result envelope for *Config.
//     // RULE: Convert raw generic result.Wrap[*Config] to an explicit concrete type
//     // in types.go so all signatures and callers share the exact same definition!
//     ConfigResult = result.Wrap[*Config]
// )
// -----------------------------------------------------------------------------

// ❌ FORBIDDEN: Bare error returns, uninformative errors.New, swallowed errors, missing blank lines
func LoadConfig(path string) (*Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        _ = err // ❌ SWALLOWED ERROR
        return nil, err // ❌ Missing blank line before return, bare error return
    }

    var cfg Config
    if err := json.Unmarshal(data, &cfg); err != nil { // ❌ Semicolon in if
        return nil, errors.New("invalid json") // ❌ BARE ERROR WITHOUT CAUSE
    }

    return &cfg, nil
}

// ✅ REQUIRED: Strict appfault wrapping with errtype, concrete ConfigResult from types.go, mandatory blank lines, flat ifs
func LoadConfig(path string) ConfigResult {
    if path == "" {
        fault := appfault.New(errtype.Validation, "config path cannot be empty").
            WithOp("config.LoadConfig")

        return result.WrapFailure[*Config](fault)
    }

    data, err := os.ReadFile(path)

    if err != nil {
        fault := appfault.WrapFile(errtype.IO, err, path, "failed to read configuration file").
            WithOp("config.LoadConfig")

        return result.WrapFailure[*Config](fault)
    }

    var cfg Config
    err = json.Unmarshal(data, &cfg)

    if err != nil {
        fault := appfault.Wrap(errtype.Validation, err, "failed to parse configuration json").
            WithOp("config.LoadConfig").
            WithVar("path", path)

        return result.WrapFailure[*Config](fault)
    }

    return result.WrapSuccess(&cfg)
}
```

### Sample 2: Database Query with Result Monad (Concrete `UserResult` in `types.go`)

```go
// -----------------------------------------------------------------------------
// Step 1: Declare Concrete Types in `types.go` (Mandatory Rule)
// -----------------------------------------------------------------------------
// In types.go:
// type (
//     // User represents the persistent user entity model.
//     User struct {
//         Id    int64  `json:"id"`
//         Name  string `json:"name"`
//         Email string `json:"email"`
//     }
//
//     // UserResult is the canonical single reusable result envelope for User.
//     // RULE: Convert raw generic result.Wrap[User] to a concrete named type in types.go.
//     // Never leak raw generic parameters across package boundaries and service signatures.
//     UserResult = result.Wrap[User]
// )
// -----------------------------------------------------------------------------

// ❌ FORBIDDEN: Combined if with semicolon, nested if, error-type branching, missing blank lines
func (r *UserRepository) FindUser(ctx context.Context, id int64) (*User, error) {
    row := r.db.QueryRowContext(ctx, "SELECT name, email FROM users WHERE id = ?", id)
    var user User

    if err := row.Scan(&user.Name, &user.Email); err != nil { // ❌ Combined semicolon if
        if errors.Is(err, sql.ErrNoRows) { // ❌ FORBIDDEN: Nested if and error-type branching
            return nil, nil // ❌ Deceptive swallowed error
        }
        return nil, err // ❌ Missing blank line before return
    }

    user.Id = id
    return &user, nil // ❌ Missing blank line before return
}

// ✅ REQUIRED: Flat if guard, direct error typing without branching, concrete UserResult from types.go, blank lines before return
func (r *UserRepository) FindUser(ctx context.Context, id int64) UserResult {
    if id <= 0 {
        fault := appfault.New(errtype.Validation, "user id must be positive").
            WithOp("UserRepository.FindUser").
            WithVar("id", id)

        return result.WrapFailure[User](fault)
    }

    row := r.db.QueryRowContext(ctx, "SELECT name, email FROM users WHERE id = ?", id)
    var user User

    err := row.Scan(&user.Name, &user.Email)

    if err != nil {
        // Direct error typing: select the error type reflecting this layer (errtype.Database)
        // Attach the ID and variables directly. Never branch on error types or nest ifs!
        fault := appfault.Wrap(errtype.Database, err, "failed to scan user row from database").
            WithOp("UserRepository.FindUser").
            WithVar("id", id)

        return result.WrapFailure[User](fault)
    }

    user.Id = id

    return result.WrapSuccess(user)
}
```

### Sample 3: Propagating Errors Across Boundaries (Concrete `UserResult` Across Services)

```go
// ✅ REQUIRED: Propagate downstream Fault directly without redundant nested wrapping, using concrete UserResult
func (s *UserService) ActivateUser(ctx context.Context, userId int64) UserResult {
    userRes := s.repo.FindUser(ctx, userId)

    if userRes.IsFailed() {
        s.log.LogError(userRes.Fault())

        // Propagate existing Fault directly with zero re-wrapping
        return result.WrapFailureFromWrap[User](userRes)
    }

    user := userRes.Value()
    user.IsActive = true

    updateRes := s.repo.UpdateUser(ctx, user)

    if updateRes.IsFailed() {
        s.log.LogError(updateRes.Fault())

        return result.WrapFailureFromWrap[User](updateRes)
    }

    return result.WrapSuccess(user)
}
```

---

---

## 4. Dedicated Section: Error Return Contract & Outer Handling Principle (Zero Dual-Handling)

A function that declares an error or result return type MUST return the actual error instance directly to the caller. It MUST NEVER invoke an exit handler, terminate the process, or panic internally and then return `nil`.

### Why Dual-Handling & Internal Exit Is Forbidden

1. **Broken Caller Sovereignty:** When a leaf function handles its own exit internally and returns `nil`, the caller is deceived into believing the operation succeeded.
2. **Impossible Testability:** Unit tests cannot assert returned error types or values if the helper function kills the process or handles errors internally.
3. **Dual Execution Hazards:** Calling an exit handler inside a helper while returning a result creates race conditions, partial database mutations, and skipped resource cleanups.

### Mandatory Outer Handling Pattern

- **Leaf/Service Functions:** Construct or wrap `*appfault.AppError` and return it.
- **Top-Level Root Dispatcher / HTTP Router:** Only the outer controller handles the error, decides the exit code via `ExitCodeType` enum, and writes the Universal Response Envelope:

```go
// ✅ REQUIRED: Top-level caller handles the error and exit
func MainCommandDispatcher(args []string) {
    res := ExecuteOperation(args)
    if res.IsFailed() {
        exitHandler.HandleValidationError(res.Fault())
        return
    }

    exitHandler.HandleSuccess()
}
```

---

---

## 6. Authoritative Spec Files Checklist (Non-Negotiable Action Items)

You MUST read, follow, and mechanically verify every single specification file below before and during execution:

- [ ] **`02-spec/02-coding-guidelines/02-canonical-size-tier.md`**
  - **Why:** Universal size limits across all languages.
  - **How:** Functions <= 8 lines preferred (hard cap 15 lines). Files <= 100 lines coding max (recommended <= 80 lines). Zero line-compression cheating.
- [ ] **`02-spec/02-coding-guidelines/06-ai-optimization/readme.md`**
  - **Why:** Comprehensive catalog of forbidden vs required generation patterns.
  - **How:** Strictly follow AH-N1 to AH-T2 rules. Zero ghost diffs, zero truncation stubs (`// ...`), zero unverified claims.
- [ ] **`02-spec/02-coding-guidelines/06-ai-optimization/06-citation-requirement.md`**
  - **Why:** Grounded rule enforcement and traceability.
  - **How:** Cite authoritative spec files for every code modification made.
- [ ] **`02-spec/02-coding-guidelines/01-cross-language/04-code-style/02-braces-and-nesting.md`**
  - **Why:** Absolute zero tolerance for nested conditionals.
  - **How:** Flatten all nested `if` statements with guard clauses and early returns.
- [ ] **`02-spec/03-error-manage/readme.md`**
  - **Why:** Authoritative error management foundation across all services.
  - **How:** Never swallow errors; every `catch` logs with operation name and key inputs, then rethrows or returns a typed error.
- [ ] **`02-spec/03-error-manage/02-error-architecture/02-error-handling-reference.md`**
  - **Why:** Universal cross-stack `AppError` and `AppException` structure.
  - **How:** Implement `appfault.Wrap(errType, err, "OpName")` in Go, `throw new AppError(cause, { op, ctx })` in TS, and `AppException` in C#/PHP; preserve the root cause and causal stack.
- [ ] **`02-spec/03-error-manage/02-error-architecture/03-go-delegation-fix.md`**
  - **Why:** Prevents nil pointer panics and raw error leaks in Go routines.
  - **How:** Never delegate errors to uninitialized handlers; use explicit, typed error delegation channels with mutex guards.
- [ ] **`02-spec/03-error-manage/02-error-architecture/readme.md`**
  - **Why:** Standardized error severity and UI feedback mapping.
  - **How:** Map log levels strictly: `debug` (trace), `info` (lifecycle), `warn` (recoverable/amber), `error` (user-visible failure/red), `fatal` (process exit).
- [ ] **`02-spec/03-error-manage/02-error-architecture/05-response-envelope/readme.md`**
  - **Why:** Universal API response contract across all endpoints.
  - **How:** Every HTTP/RPC response MUST return the standard envelope: `{ "data": T, "errors": [AppError], "meta": Meta }`. Never return raw un-enveloped error text.
- [ ] **`02-spec/03-error-manage/03-error-code-registry/readme.md`**
  - **Why:** Stable error code registry and catalog.
  - **How:** All error codes must be registered constants (`errtype.Variation`). No ad-hoc string literals invented at the throw site.

---

---

## 7. Mandatory Linter & Targeted Verification Checklist

Code standards must be mechanically enforced by automated linters. You MUST verify or create the linter and connect it to CI:

- [ ] **Linter Script Identification:** Check if `linter-scripts/check-error-management.py` exists in the repository.
- [ ] **Auto-Create Linter if Missing:** If no dedicated error linter exists, create `linter-scripts/check-error-management.py` that AST-scans for:
  1. Internal exit handler invocations in non-main functions.
  2. Empty `catch` or `except` blocks (swallowed errors).
  3. Bare un-wrapped error returns (`return err` instead of `appfault.Wrap`).
  4. Bare panics/hard exits (`panic()`, `process.exit()`, `os.Exit()`).
  5. Non-standard API responses lacking the `{ data, errors, meta }` envelope.
- [ ] **Local Linter Command:** Execute and verify the linter locally on modified files:
  ```bash
  python linter-scripts/check-error-management.py
  ```

---

---

## Metadata

- slug: cg-error-management
- priority: high
- status: active
