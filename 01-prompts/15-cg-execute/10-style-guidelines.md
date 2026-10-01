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

[/goal](slashCommand:goal) Autonomously inventory, scan, partition, refactor, and fix all coding style, vertical newline spacing, blank line before `if`, blank line after `}`, blank line before `return`, blank lines around parameter struct instantiations and sequential function invocations, nested `if` elimination, function length (<= 8–15 lines), file size (<= 100 lines), LF line endings (`\n`), and UTF-8 (no BOM) encoding across ALL source files in the repository in bounded micro-batches of 5–8 files per subtask, running a continuous 2-agent unstoppable self-loop: FIRST showcase and list out the given task in visible chat during Turn 1, capture it verbatim, plan it in the repo, spawn autonomous subagents via `invoke_subagent` (A = 2, H = 2; solo execution without calling `invoke_subagent` is an auto-reject failure) in disjoint file boxes using GitMap high-speed commands as primary, prove every single claim with concrete evidence, enforce coding guidelines to 100%, and finish with one atomic GitMap commit that holds strictly this task's files.

[/learn](slashCommand:learn) Enforce the Top-Instruction Priority Mandate: whatever directives, custom requirements, checklists, or user instructions are provided ABOVE this prompt outrank everything below. Turn 1 MUST showcase the given task list in visible chat before any background execution. Each rule is stated once (R1 to R16) and cited by ID. Progress lives in the ledger and in `.ai-memory/plans/`, never only in chat.

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

## AI Diagnostic Guide: Why Newline Violations are Missed in Functions & How to Fix Them

> [!IMPORTANT]
> **CRITICAL FAILURE ROOT CAUSE (PREVENTING PREMATURE COMPLETION):**
>
> In past runs, AI agents failed to fix function newlines because they:
>
> 1. **Only ran surface linters** (e.g. checking file-level EOF newlines) without opening individual function bodies.
> 2. **Wrote a planning markdown file and immediately declared completion** without editing a single `.go` or `.ts` file.
> 3. **Attempted to process the entire codebase in one prompt**, causing context exhaustion and truncated file edits.
>
> **THE MANDATORY REMEDY:**
>
> - You MUST partition the full list of codebase files into **batches of 5–8 files each**.
> - Subagents MUST open and edit every single file in their assigned batch line-by-line.
> - The master orchestrator MUST continuously self-loop across all batches until every single batch in `.ai-memory/plans/subtasks/` is completed.

---

---

## Dedicated Section: Comprehensive Coding Style, Line-Gaps & Anti-Pattern Gallery (Zero Tolerance)

Proper vertical spacing and code hygiene are essential for readability and automated analysis. Dense, squeezed code without blank lines around control flow statements leads to missed edge cases, obscured invariants, and severe cognitive fatigue.

---

### Rule 1: Mandatory Blank Line BEFORE Control Structures (`if`, `for`, `switch`, `while`, `try`)

Whenever a control structure (`if`, `for`, `switch`, `while`, `try`) is preceded by **any statement** (variable declaration, assignment, method call, channel receive, or loop), there **MUST be exactly one blank line before the control structure**.

*Exception:* If the control structure is the **very first line** of a function body or immediately follows an opening brace `{`, no blank line is required before it.

#### 1a. Go: Blank Line Before `if`

```go
// ❌ WRONG: Squeezed variable declaration / map lookup directly against if
func ProcessUser(id string) error {
    user, isFound := userCache.Get(id)
    if !isFound {
        return ErrUserNotFound
    }
    config, isLoaded := loadConfig()
    if !isLoaded {
        return ErrConfigMissing
    }
    return executeUser(user, config)
}

// ✅ CORRECT: Clean blank line before each if statement, after each closing brace, and before return
func ProcessUser(id string) error {
    user, isFound := userCache.Get(id)

    if !isFound {
        return ErrUserNotFound
    }

    config, isLoaded := loadConfig()

    if !isLoaded {
        return ErrConfigMissing
    }

    return executeUser(user, config)
}
```

---

#### 1b. TypeScript / React: Blank Line Before `if`

```typescript
// ❌ WRONG: Squeezed variable declarations and function calls against if
function getFormattedPrice(item: Item): string {
    const rawPrice = calculateBasePrice(item);
    const hasDiscount = item.discountPercent > 0;
    if (hasDiscount) {
        return applyDiscount(rawPrice, item.discountPercent);
    }
    const formatted = formatCurrency(rawPrice);
    return formatted;
}

// ✅ CORRECT: Clean blank line before if and before final return
function getFormattedPrice(item: Item): string {
    const rawPrice = calculateBasePrice(item);
    const hasDiscount = item.discountPercent > 0;

    if (hasDiscount) {
        return applyDiscount(rawPrice, item.discountPercent);
    }

    const formatted = formatCurrency(rawPrice);

    return formatted;
}
```

---

#### 1c. Python: Blank Line Before `if`

```python
# ❌ WRONG: Assignment directly followed by if without blank line
def fetch_user_profile(user_id: str) -> Profile:
    auth_token = get_session_token()
    is_valid_token = verify_token(auth_token)
    if not is_valid_token:
        raise UnauthorizedError()
    user_record = db.find_user(user_id)
    if user_record is None:
        raise NotFoundError()
    return Profile.from_record(user_record)

# ✅ CORRECT: Clean blank line before if and before returns
def fetch_user_profile(user_id: str) -> Profile:
    auth_token = get_session_token()
    is_valid_token = verify_token(auth_token)

    if not is_valid_token:
        raise UnauthorizedError()

    user_record = db.find_user(user_id)

    if user_record is None:
        raise NotFoundError()

    return Profile.from_record(user_record)
```

---

#### 1d. PHP: Blank Line Before `if` and `foreach`

```php
// ❌ WRONG: Squeezed statements before if and foreach
$result = $this->apiRequest($agentId, HttpMethodType::Post->value, $endpoint);
if (is_wp_error($result)) {
    return $result;
}
$items = $this->fetchItems();
foreach ($items as $item) {
    $this->process($item);
}

// ✅ CORRECT: Separated with blank lines before control structures
$result = $this->apiRequest($agentId, HttpMethodType::Post->value, $endpoint);

if (is_wp_error($result)) {
    return $result;
}

$items = $this->fetchItems();

foreach ($items as $item) {
    $this->process($item);
}
```

---

### Rule 2: Mandatory Blank Line AFTER Closing Brace `}` When Followed by Code

Whenever a closing brace `}` (from an `if`, `for`, `switch`, `while`, or `try/catch` block) is followed by further executable code or another statement, there **MUST be exactly one blank line after `}`**.

*Exception:* No blank line is needed when `}` is followed by another closing `}`, `else`, `catch`, `finally`, or the end of a function body.

#### 2a. Go: Blank Line After `}` Following Control Flow

```go
// ❌ WRONG: Closing brace squeezed against next statement
func ExecuteStep(step Step) error {
    if err := step.Validate(); err != nil {
        return err
    }
    result, err := step.Run()
    if err != nil {
        return err
    }
    return saveResult(result)
}

// ✅ CORRECT: Clean blank line after every closing brace
func ExecuteStep(step Step) error {
    if err := step.Validate(); err != nil {
        return err
    }

    result, err := step.Run()
    if err != nil {
        return err
    }

    return saveResult(result)
}
```

---

#### 2b. TypeScript: Blank Line After Loops & Try/Catch

```typescript
// ❌ WRONG: Loop and try/catch squeezed against subsequent logic
for (const item of items) {
    processed.push(transform(item));
}
const result = merge(processed);

try {
    saveToStorage(result);
} catch (error) {
    logger.error(error);
}
cleanup();

// ✅ CORRECT: Blank line after each closing brace
for (const item of items) {
    processed.push(transform(item));
}

const result = merge(processed);

try {
    saveToStorage(result);
} catch (error) {
    logger.error(error);
}

cleanup();
```

---

#### 2c. Go: Blank Line After Multiline Map/Slice Literals & Loops with Boolean Extraction

```go
// ❌ FORBIDDEN (Unacceptable): Squeezed loops against map literals, inline conditional assignments, and missing blank lines between if blocks
func ValidateDoubleExtensionFormats(targetPath string) *appfault.AppError {
    cases := map[string]Format{
        "archive.tar.gz":  FormatTarGz,
        "archive.tgz":     FormatTarGz,
        "archive.tar.bz2": FormatTarBz2,
        "archive.tbz2":    FormatTarBz2,
        "archive.tar.xz":  FormatTarXz,
        "archive.txz":     FormatTarXz,
        "archive.tar.zst": FormatTarZst,
        "archive.tzst":    FormatTarZst,
    }
    for path, expectedFormat := range cases {
        if got := FormatFromPath(path); got != expectedFormat {
            return appfault.New(errtype.Validation, "mismatch")
        }
    }
    return nil
}

func ValidateExtensionRoundTrip(formats []Format) *appfault.AppError {
    for _, f := range formats {
        ext := f.Extension()
        if len(ext) == 0 {
            return appfault.New(errtype.Validation, "empty extension")
        }
        if got := FormatFromPath("sample" + ext); got != f {
            return appfault.New(errtype.Validation, "unmatched format")
        }
    }
    return nil
}

// ✅ REQUIRED (Right Practice): Blank line after map literal closing brace, blank line before loops, blank line before if, blank line after closing brace, and extracted affirmative booleans
func ValidateDoubleExtensionFormats(targetPath string) *appfault.AppError {
    cases := map[string]Format{
        "archive.tar.gz":  FormatTarGz,
        "archive.tgz":     FormatTarGz,
        "archive.tar.bz2": FormatTarBz2,
        "archive.tbz2":    FormatTarBz2,
        "archive.tar.xz":  FormatTarXz,
        "archive.txz":     FormatTarXz,
        "archive.tar.zst": FormatTarZst,
        "archive.tzst":    FormatTarZst,
    }

    for path, expectedFormat := range cases {
        resolvedFormat := FormatFromPath(path)
        isFormatMismatch := resolvedFormat != expectedFormat

        if isFormatMismatch {
            return appfault.New(
                errtype.Validation,
                "format mismatch detected",
            ).WithOp("ValidateDoubleExtensionFormats")
        }
    }

    return nil
}

func ValidateExtensionRoundTrip(formats []Format) *appfault.AppError {
    for _, f := range formats {
        ext := f.Extension()
        isEmptyExtension := len(ext) == 0

        if isEmptyExtension {
            return appfault.New(
                errtype.Validation,
                "extension returned empty string",
            ).WithOp("ValidateExtensionRoundTrip")
        }

        got := FormatFromPath("sample" + ext)
        isSampleUnmatchFile := got != f

        if isSampleUnmatchFile {
            return appfault.New(
                errtype.Validation,
                "unmatched sample file format",
            ).WithOp("ValidateExtensionRoundTrip")
        }
    }

    return nil
}
```

---

#### 2d. Go: Blank Lines Around Struct Instantiations & Sequential Function Invocations

When instantiating a parameter struct or invoking a multi-line function, there **MUST be a blank line before the invocation** (if preceded by assignments or statements) and **MUST be a blank line after the invocation closing brace `}`** before subsequent statements, `if` conditions, or other function calls.

```go
// ❌ FORBIDDEN (Unacceptable): Squeezing variable assignments, multiline struct invocations, and following if statements without vertical line gaps
func PrintIdentityBlock(cwd string) {
    fmt.Println(" " + constants.ColorCyan + "Identity Block" + constants.ColorReset)
    src := getSourceDirectory()
    emitIdentityRows(IdentityRowParams{
        Dir:            src,
        RepoOverride:   buildRepo,
        BranchOverride: buildBranch,
        ShaOverride:    buildCommit,
    })
    if len(buildDate) > 0 {
        fmt.Printf(" Built: %s\n", buildDate)
    }
    emitIdentityRows(IdentityRowParams{
        Dir: cwd,
    })
    fmt.Println()
}

// ✅ REQUIRED (Right Practice): Clean blank lines before and after multiline struct calls, separating discrete execution stages
func PrintIdentityBlock(cwd string) {
    fmt.Println(" " + constants.ColorCyan + "Identity Block" + constants.ColorReset)

    src := getSourceDirectory()

    emitIdentityRows(IdentityRowParams{
        Dir:            src,
        RepoOverride:   buildRepo,
        BranchOverride: buildBranch,
        ShaOverride:    buildCommit,
    })

    if len(buildDate) > 0 {
        fmt.Printf(" Built: %s\n", buildDate)
    }

    emitIdentityRows(IdentityRowParams{
        Dir: cwd,
    })

    fmt.Println()
}
```

---

### Rule 3: Mandatory Blank Line BEFORE `return`, `throw`, `raise`, `yield`

In multi-line functions and blocks, there **MUST be a blank line before `return` / `throw` / `raise` / `yield`**.

*Exception:* Single-statement function body (`func GetId() string { return c.Id }`) or when `return` is the immediate first statement inside a block.

```go
// ❌ WRONG: Return squeezed directly under statements
func CalculateTotal(items []Item, taxRate float64) float64 {
    subtotal := computeSubtotal(items)
    tax := subtotal * taxRate
    total := subtotal + tax
    return total
}

// ✅ CORRECT: Blank line before final return
func CalculateTotal(items []Item, taxRate float64) float64 {
    subtotal := computeSubtotal(items)
    tax := subtotal * taxRate
    total := subtotal + tax

    return total
}
```

---

### Rule 4: Zero Clumping of Consecutive Guard Clauses

When multiple guard clauses follow one another sequentially, **each guard clause MUST be separated by a blank line** after its closing brace `}`. Never clump or stack guard clauses together without vertical breathing room.

```go
// ❌ WRONG: Clumped guard clauses with zero spacing
func ValidateOrder(order *Order) error {
    if order == nil {
        return ErrNilOrder
    }
    if !order.HasItems() {
        return ErrEmptyOrder
    }
    if order.TotalAmount <= 0 {
        return ErrInvalidAmount
    }
    return nil
}

// ✅ CORRECT: Vertical breathing room between discrete guard clauses
func ValidateOrder(order *Order) error {
    if order == nil {
        return ErrNilOrder
    }

    if !order.HasItems() {
        return ErrEmptyOrder
    }

    if order.TotalAmount <= 0 {
        return ErrInvalidAmount
    }

    return nil
}
```

---

### Rule 5: Zero Nested `if` Statements (Mandatory Flattening to Depth 0)

Conditionals MUST NEVER exceed depth 1 (i.e., **no nested `if` statements inside another `if` block**). Flatten all branching logic using early guard returns or discrete helper functions.

---

### Rule 6: No Multi-Statement Lines & No Inline Compound Condition Cramming (Multi-Line Separation)

1. **No Semicolon Packing:** Never compress multiple statements onto a single line using semicolons (`a = 1; b = 2; return a + b`). Each statement MUST occupy its own line.
2. **Total Ban on Inline Compound Assignments (`if init; cond`):** NEVER cram variable declarations, type assertions, or multi-part boolean checks into the `if` header (e.g. `if v, isString := rawMap[key].(string); isString && len(v) > 0 {`).
3. **Multi-Line Statement Separation:**
   - Execute variable assignments / lookups on their own dedicated line.
   - Evaluate and assign the boolean condition to an affirmative variable (`is*` or `has*`) on its own dedicated line *before* the `if` statement.
   - Maintain vertical breathing room (blank line before `if`).
   - The `if` condition itself must be dead simple, checking **one single variable**.
4. **Zero Magic Strings & Constant Returns:** Never use raw string literals (`"unknown"`, `"pending"`, `"failed"`) as return or fallback values. Define named constants (`VersionUnknown = "unknown"`) and return constants directly. Merge related lookup strings into constants and package-level slices (`versionKeys`), eliminating inline slice allocations.

#### Canonical Example: What NOT to Do vs What to Do

```go
// ❌ BANNED ANTI-PATTERN:
// 1. Cramming type assertion assignment and compound condition into one line.
// 2. Hardcoding magic strings ("Version", "version", "unknown") inline.
// 3. Returning raw fallback literal instead of a defined constant.
func extractVersionValue(rawMap map[string]interface{}) string {
    for _, key := range []string{"Version", "version"} {
        if v, isString := rawMap[key].(string); isString && len(v) > 0 {
            return v
        }
    }

    return "unknown"
}

// ✅ MANDATORY CLEAN PATTERN:
// 1. Zero magic strings: extract lookup keys and defaults into constants.
// 2. Merge repeated/related strings into reusable collections (versionKeys).
// 3. Assignment on its own dedicated line.
// 4. Affirmative boolean (hasContent) pre-evaluated BEFORE the if statement.
// 5. Clean vertical breathing room (blank line before if).
// 6. Dead-simple if statement evaluating exactly ONE variable.
// 7. Return defined constant (VersionUnknown) instead of raw magic string literal.
const (
    VersionUnknown  = "unknown"
    versionKeyUpper = "Version"
    versionKeyLower = "version"
)

var versionKeys = []string{versionKeyUpper, versionKeyLower}

func extractVersionValue(rawMap map[string]interface{}) string {
    for _, key := range versionKeys {
        v, isString := rawMap[key].(string)
        hasContent := isString && len(v) > 0

        if hasContent {
            return v
        }
    }

    return VersionUnknown
}
```

---

### Rule 7: Universal File Hygiene, Line Endings (LF `\n` Only) & Encoding (UTF-8 No BOM)

1. **Unix LF (`\n`) Line Endings Only:** Every file MUST use Unix LF (`\n`). Total ban on Windows CRLF (`\r\n`).
2. **Strict UTF-8 Encoding (NO BOM):** Save all files in UTF-8 without BOM.
3. **Mandatory Single Trailing Newline at EOF:** Exactly one newline at the end of every file.

---

---

## Mandatory Linter & CI/CD Integration

1. **Linter Scripts:** `linter-scripts/check-function-lengths.py`, `linter-scripts/check-mws-error-codes.py`, `linter-scripts/check-newline-styling.py`
2. **Local Run Command:** `python linter-scripts/check-function-lengths.py`
3. **Autofixer Command:** `python 03-ai-scripts/05-guideline-autofixer.py <file>`
4. **CI/CD Integration (`.github/workflows/ci.yml`):**
   ```yaml
   - name: Validate Newline Styling & Function Lengths
     run: |
       python linter-scripts/check-function-lengths.py
       python linter-scripts/check-mws-error-codes.py
       python linter-scripts/check-newline-styling.py
   ```
5. **Runner Registration (`03-ai-scripts/06-cicd-local-runner.py`):**
   ```python
   JOBS = {
       "Newline Styling Check": [sys.executable, "linter-scripts/check-newline-styling.py"],
       "Function Lengths Check": [sys.executable, "linter-scripts/check-function-lengths.py"],
       "Error Codes Check": [sys.executable, "linter-scripts/check-mws-error-codes.py"],
   }
   ```

---

---
