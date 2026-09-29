[/goal](slashCommand:goal) Autonomously ingest, follow, and execute the referenced external prompts, task instructions, and coding guideline directives across the target codebase in bounded 5-8 file micro-batches with strict no-build and no-test execution (NEVER run build commands like `go build` or `npm run build`, and NEVER run test suites like `go test ./...` or `pytest` during routine execution turns; all compilation and testing are strictly verified later in CI/CD). Spawn autonomous subagents (A = 2, H = 2) for parallel reading and rule extraction, read the codebase using GitMap AUM as primary, establish a single-agent blueprint during Phase 1 planning (first 50% steps budget), and execute disjoint refactoring in Phase 2 (remaining 50% steps budget) with continuous self-looping until 100% complete and finalized with an atomic push.

[/learn](slashCommand:learn) Enforce the Bottom-Instruction Priority Mandate: whatever directives, custom requirements, checklists, or user instructions are appended BELOW this prompt (following the `--` divider border at the bottom) are HIGHEST PRIORITY and strictly NON-NEGOTIABLE. Ingest and execute the bottom instructions ahead of all default conventions above, internalize the single-agent architectural blueprint, and persist all progress into `.ai-memory/plans/` and memory logs.

> **Bottom-Instruction Priority Mandate (Below Precedence / Suffix Precedence):**
> Whatever directives, constraints, checklists, or user instructions are given BELOW this prompt (following the `--` divider border at the bottom) are HIGHEST PRIORITY and MUST BE FOLLOWED as strictly NON-NEGOTIABLE. They supersede and strictly override any conflicting general advice, default conventions, or guidelines above. The agent MUST inspect and follow the instructions below with absolute precedence.

```text
N = 200 (Total self-loop steps budget)
A = 2   (Number of spawned autonomous subagents running concurrently, default: 2)
H = 2   (Operational hands per agent: dual-task batch capacity & parallel tool dispatch, default: 2)

System Concurrency Capacity = A × H = 2 agents × 2 hands = 4 concurrent subtask operations
```

```text
PHASE_1_STEPS = N / 2   (Steps 1 .. N/2: Planning, Detailed Spec, and Lean Subtask Generation)
PHASE_2_STEPS = N / 2   (Steps N/2+1 .. N: Parallel Execution, Self-Looping, Targeted Quality Linting)
```

N, A, H, PHASE_1_STEPS, and PHASE_2_STEPS are read-only after initialization. Never modify them mid-execution.

#### High-Speed GitMap Acceleration Options (Run Everything Faster)

Always prefer native GitMap commands over slow generic shell pipelines:
1. **Ultra-Fast File & Directory Discovery (AUM Index & Walk):**
   - **Wildcard / Glob Search:** `gitmap find "<wildcard*>" [-ext <ext>]` (alias `gitmap f`)
   - **Exact Filename Search:** `gitmap find-files <name> [-ext <ext>]` (alias `gitmap ff`)
   - **Substring Filename Search:** `gitmap find-files-any "<str>" [-ext <ext>]` (alias `gitmap ffa`)
   - **Prefix / Suffix Search:** `gitmap find-files-startswith <prefix>` (`gitmap ffs`) / `gitmap find-files-endswith <suffix>` (`gitmap ffe`)
   - **List Indexed Repo Files:** `gitmap list-files [pattern] [-ext <ext>]` (alias `gitmap lf`)
   - **Directory Tree & Scaffolding:** `gitmap folder-tree` (alias `gitmap ft`)
   - **Zero-Write File Stream:** `gitmap cat <filepath>`
   - **Instant Multi-Core Regex Search:** `gitmap search "<term>"` or `gitmap aum search "<query>" [dir] --ext <ext>`
2. **Fast Repository Hygiene, Lowercase & Symlink Repair:**
   - **Auto-Lowercase Files (Safe 2-Step `git mv`):** `gitmap lowercase` (alias `gitmap lcf [--dry-run]`)
   - **Lowercase Root Readme:** `gitmap lowercase-readme`
   - **Sync Curated `.gitignore` / `.gitattributes` / `.prettierignore`:** `gitmap commons` (alias `gitmap co` or `gitmap sync all`)
   - **Repair Broken Symlinks:** `gitmap fix-link` (alias `gitmap fixlink`)
   - **Clean Update Temp & Inspect Storage:** `gitmap update-cleanup`, `gitmap storage` (alias `gitmap stor`)
3. **Fast Git State, Execution & Atomic Commits:**
   - **Repo Status & Remote Check:** `gitmap status` (`gitmap st`), `gitmap has-any-updates` (`gitmap hau`), `gitmap latest-branch` (`gitmap lb`)
   - **Fast Cross-Platform Shell Runner:** `gitmap pwsh "<command>"` (`gitmap ps`), `gitmap bash "<command>"` (`gitmap sh`), `gitmap async <cmd>` (`gitmap asyn`)
   - **Semantic Atomic Commit & Push:** `gitmap cpf "<summary>"` (Feature), `gitmap cpb "<summary>"` (Bug), `gitmap cpr "<summary>"` (Release), `gitmap pcp "<summary>"` (Pull-Commit-Push)
   - **Smart CI/CD Pipeline Waiting:** `gitmap pe`, `gitmap pipeline-ai status --json` (`gitmap pl-ai status -t <etaSeconds>`)

---

## Phase 0: Antigravity Skill Bootstrap (Memory Optimization)

Before executing the tasks below, check if this prompt is already installed as a native Antigravity Skill.

1. If `.agents/skills/cg-follow-other-prompts/skill.md` does not exist in the workspace, create it now.
2. Extract the core instructions of this prompt and save it into that `skill.md` using the standard YAML frontmatter (with `name` and `description`).
3. Once installed, rely on progressive disclosure for future runs. Do not keep the entire prompt in active memory if not needed.

---

## Phase 1A: Referenced Prompt Ingestion, Task Extraction & Chat Output Gate (Step 0)

Before modifying any code:

1. **Prompt Reference Ingestion:** Read all referenced prompt files (e.g. from `01-prompts/`, task specs, or paths given below) using GitMap AUM or fast file reader. Extract their non-negotiable directives, requirements, and constraints.
2. **Bottom-Instruction Priority Verification:** Whatever directives, constraints, checklists, or prompt paths are given BELOW this prompt (following the `--` divider border at the bottom) must be verified as highest priority and non-negotiable, strictly overriding all general guidelines above.
3. **Verbatim Capture:** Losslessly preserve the referenced prompt names and user instructions in `.ai-memory/plans/pending/xx-<slug>.md` under `## Referenced Prompts & Instructions (Verbatim)`.
4. **Mandatory Chat Output Gate & Same-Turn Tool Chaining (TOTAL BAN ON CLOSING CONVERSATION):**
   - You MUST output the confirmed deliverables list directly in chat, but you MUST NOT end your turn or pause after emitting it!
   - In the EXACT SAME RESPONSE turn as your markdown output, you MUST immediately invoke the first tool call (e.g. `run_command` for fast discovery or `write_to_file` to initialize the planning spec).
   - NEVER emit the breakdown text without invoking a tool call. Emitting visible text without a tool call terminates the AI turn, causing the engine to close or hang the conversation while tasks remain pending.
   - TOTAL BAN ON WAITING FOR APPROVAL: Do not pause, do not wait for the user to reply, and do not ask "Should I proceed?". Override any planning mode stop directives. Proceed immediately into Phase 1B in the same turn.
   - Line-by-Line Output Format Structure:
     - Line 1: Header `### 📋 Confirmed Task Breakdown & Requirement Ingestion`
     - Line 2: Empty blank line
     - Line 3: Numbered task title `1. **Task-01: [Descriptive Task Title]**`
     - Line 4: Indented state bullet (3 spaces) `   - **State:** [IN PROGRESS — EXECUTING IMMEDIATELY]`
     - Line 5: Indented understanding check (3 spaces) `   - **Understood:** [YES] — [1-2 concise sentences proving understanding of intent, scope, and verified constraints]`
     - Line 6: Indented actionable scope bullet (3 spaces) `   - **Actionable Scope:** [Precise technical deliverable and implementation scope]`
     - Line 7: Indented target files bullet (3 spaces) `   - **Target Files / Area:** [relative/path/or/module]`
     - Line 8: Empty blank line (vertical gap before next task)
     - Concluding Line: `Proceeding directly to Phase 1B: Spec & Subtask Generation (Active Tool Call Running Below).`

```markdown
### 📋 Confirmed Task Breakdown & Requirement Ingestion

1. **Task-01: [Descriptive Task Title]**
   - **State:** `[IN PROGRESS — EXECUTING IMMEDIATELY]`
   - **Understood:** `[YES]` — [Concise 1-sentence verification of user requirement, intent, and verified constraints]
   - **Actionable Scope:** [Precise technical deliverable and implementation scope]
   - **Target Files / Area:** `[relative/path/or/module]`

2. **Task-02: [Descriptive Task Title]**
   - **State:** `[QUEUED — EXECUTING NOW WITHOUT USER PROMPT]`
   - **Understood:** `[YES]` — [Concise 1-sentence verification of user requirement, intent, and verified constraints]
   - **Actionable Scope:** [Precise technical deliverable and implementation scope]
   - **Target Files / Area:** `[relative/path/or/module]`

Proceeding directly to Phase 1B: Spec & Subtask Generation (Active Tool Call Running Below).
```

---

## Phase 1B: Planning Mode, Detailed Spec Generation & Lean Subtasks (Steps 1 .. N/2)

### Step 1: Scan & Discover (GitMap AUM Acceleration & Multi-Agent Parallel Reading)
To avoid 50-result tool truncation limits and eliminate multi-turn exploratory roundtrips, leverage the 2-tier discovery toolchain:

#### Tier 1: GitMap AUM Acceleration (PRIMARY)
- **Universal File Search:** `gitmap find "<pattern>" [-ext <ext>]`
- **List Indexed Files:** `gitmap list-files [pattern]` (alias `gitmap lf [pattern] [-ext <ext>]`)
- **Substring Match:** `gitmap find-files-any "<substring>"` (alias `gitmap ffa "<str>"`)
- **Stream File Content:** `gitmap cat <filepath>`
- **Instant Code Search:** `gitmap search "<term>"`

#### Tier 2: Fast Cached Python Toolchain (FALLBACK)
- **Inventory Target Files:** `python 03-ai-scripts/11-fast-file-scanner.py --lang go,ts --limit 100 --stats`
- **Fast Cached Grep (<15ms):** `python 03-ai-scripts/12-fast-cached-grep.py --pattern "<search-pattern>" --limit 50`
- **Sub-Millisecond Folder Exploration:** `python 03-ai-scripts/17-fast-file-reader.py --list-folder <folder-path> --limit 50`
- **Read Target File:** `python 03-ai-scripts/17-fast-file-reader.py --read-file <file-path> --max-bytes 100000`
- **Subsystem & Topology Overview:** `python 03-ai-scripts/18-codebase-topology-discoverer.py --summary`

### Step 2: Actionable Execution Plan & Lean Subtask Decomposition
Initialize the execution plan and decompose it into subtasks in `.ai-memory/plans/`:
- **Parent Plan:** Write `.ai-memory/plans/pending/xx-<slug>.md` containing:
  - Synthesis of referenced prompt instructions.
  - Complete mapping of deliverables (`Task-01`, `Task-02`, etc.) to subtask files.
- **Subtask Files:** Break down the plan into granular subtasks in `.ai-memory/plans/subtasks/xx-<slug>/01-<subtask>.md`, `02-<subtask>.md`, etc. Complete all spec and subtask writing within 50% of the steps budget (`PHASE_1_STEPS = N / 2`).
- **No Common Boilerplate:** Do not write common repository boilerplate or generic guidelines inside subtask files. Subtasks must contain only the unique items required for that specific subtask.

Subtasks must follow this lean, unique template:
```markdown
# Subtask [01]: [Descriptive Subtask Name]
Traceability ID: Task-01
Prompt Reference: [Path to referenced prompt or spec]
Target Files: [Strict relative paths from repo root]
Action: [Exact code changes, functions, types, and logic to modify or add]
Acceptance Criteria: [2-4 specific testable conditions proving completion]
Targeted Verification: [Specific file-level linter command or exit 0 check]
```

### Step 3: Unconditional Zero-Question Execution Mandate
- **Strict 50/50 Time & Step Budget Allocation:** Spec writing and subtask generation MUST strictly complete within the first 50% of the budget (`PHASE_1_STEPS = N / 2`).
- **Zero Questions / Unconditional Execution:** As soon as Phase 1 planning completes, the master orchestrator MUST NOT pause, stop, or ask the user "Should I proceed?". There is NO question. It must immediately, unconditionally self-loop and transition directly into Phase 2 execution mode.
- **Spec Writing is Only Half the Task:** Generating specs without executing code changes is an INCOMPLETE FAILURE. The remaining 50% of the budget (`PHASE_2_STEPS = N / 2`) is dedicated strictly to modifying code, running targeted quality linters, consolidating subtasks, and completing the deliverables.

---

## Phase 2: Execution Mode & Parallel Refactoring (Steps N/2+1 .. N)

1. **Parallel Dispatch & Concrete Subagent Schema (A = 2, H = 2):**
   - Unconditionally execute code refactoring across target files in the remaining 50% of the steps budget (`PHASE_2_STEPS = N / 2`).
   - Use `invoke_subagent` to spawn up to A = 2 execution subagents concurrently. Each subagent handles an operational capacity of H = 2 (a bounded batch of up to 2 disjoint subtasks from `.ai-memory/plans/subtasks/xx-<slug>/`).
   - Subagents executing code changes MUST use `TypeName: "self"` to inherit the parent agent's write and command capabilities (`write_to_file`, `replace_file_content`, `run_command`).
2. **Self-Contained Subagent Prompt Envelope:**
   - Subagents spawn with a clean context and must receive a complete, self-contained prompt envelope with assigned subtasks, strict disjoint target files bounding box, non-negotiable coding rules, and completion reporting contract.
3. **Reactive Wakeup & Turn-Yielding Protocol (Deadlock Prevention):**
   - Immediately after issuing the `invoke_subagent` tool call, the parent orchestrator MUST output a brief progress note to the user and **STOP CALLING TOOLS**.
   - NEVER run tight-polling loops using `manage_task` or filesystem checks to wait for subagents. Ending the tool-call chain allows the platform scheduler to execute the background subagents and deliver their completion messages into the parent's inbox upon wakeup.
4. **Coding Guidelines Enforced:**
   - Positive booleans only (`is` and `has`), no explicit `== true`.
   - Structured Go errors: return `*appfault.AppError`, never bare `error`.
   - Function sizing: <= 8 lines preferred (hard cap 15 lines).
   - Concrete types: Extract domain structs and Result wrappers into `types.go`.
   - Multi-line arguments formatted one per line with trailing commas.
   - Strict relative git paths (zero absolute paths or `file:///` URIs).
5. **Total Ban on Test Running & Build Checking:**
   - Do not run any tests using Python scripts, Go (`go test`), or any test runner during routine execution turns.
   - Do not run build verification commands (`go build`, `npm run build`, compiler invocations). Build compilation and testing are checked later in CI/CD.
6. **Targeted Quality Linting Only:**
   - Run only targeted, fast file-level linters or autofixers on specifically modified files (`exit 0`).

---

## Phase 3: Task Consolidation & File Reduction (End of Loop)

To reduce markdown file count and bloat, consolidate subtasks when a parent task is 100% complete:

1. Combine all completed granular subtasks from `.ai-memory/plans/subtasks/xx-<slug>/*.md` into a single consolidated file at `.ai-memory/plans/completed/xx-<slug>.md`.
2. Delete the original granular `.md` files in `.ai-memory/plans/subtasks/xx-<slug>/`.
3. Delete the original parent plan `.ai-memory/plans/pending/xx-<slug>.md`.
4. Update `.ai-memory/plans/readme.md` to point to the newly consolidated completed file.
5. Final Step Git Commit & Push via GitMap Semantic Commit Commands (Mandatory):
   - Use GitMap semantic commit commands:
     - For features/tasks: `gitmap cpf "<summary>"` (automatically stages all files, prefixes `Feature: `, commits, and pushes).
     - For fixes/bugs: `gitmap cpb "<summary>"` (automatically stages all files, prefixes `Bug: `, commits, and pushes).
     - For safe pull-commit-push: `gitmap pcp "<summary>"`.
   - If GitMap CLI is unavailable, fallback to raw git: `git add -A && git commit -m "<summary>" && git push origin <branch>`.
   - Under no circumstances commit each file individually.

---

## End-of-Turn Verification & Confidence Reporting (Mandatory Output)

At the completion of all tasks and before concluding the turn, you MUST emit this structured verification summary in the chat response:

```markdown
### Task Completion Summary

- ✅ **Task-01: [Descriptive Task Title]** — `[Completed]`
- ✅ **Task-02: [Descriptive Task Title]** — `[Completed]`

### Modified Files Summary

- [relative/path/to/modified/file1.ext]
- [relative/path/to/modified/file2.ext]

### Implementation Confidence Score

- Confidence: [e.g. 98% or 100%]
- Rationale: [Detailed explanation of verified quality gates, passing linters, contract adherence, and zero regressions]

### 🤖 Independent AI Verification & Audit Prompt

```markdown
### Independent AI Audit & Verification Instructions

You are an Independent AI Verification and Quality Auditor.
Your task is to independently audit, verify, and remediate the implementation against the canonical specification and verbatim requirements.

#### 1. Target Documents & Implemented Code:
- **Consolidated Plan & Subtasks:** [.ai-memory/plans/completed/xx-<slug>.md](.ai-memory/plans/completed/xx-<slug>.md)
- **Modified & Implemented Code Files:**
  - [relative/path/to/modified/file1.ext](relative/path/to/modified/file1.ext)
  - [relative/path/to/modified/file2.ext](relative/path/to/modified/file2.ext)

#### 2. Verification Protocol:
1. Strict Verbatim Inspection: Read the plan and requirements completely.
2. Line-by-Line Code Comparison: Verify every single guideline requirement is fully implemented.
3. Autonomous Self-Loop Remediation: If any gaps exist, self-loop and fix directly without asking.
4. Final Audit Verdict: Emit PASS/FAIL verdict with confidence score.
```
```

---

## Banned Operations Checklist (TOTAL BAN — Auto-Reject on Violation)

- [ ] BOTTOM-INSTRUCTION PRIORITY MANDATE (BELOW PRECEDENCE): Whatever directives, constraints, checklists, prompt paths, or user instructions are given BELOW this prompt (following the `--` divider border at the bottom) are verified as highest priority and non-negotiable.
- [ ] NO TEST RUNNING (TOTAL BAN): Never run any tests using Python scripts (`06-cicd-local-runner.py`, `pytest`), Go (`go test ./...`), or any test runner during routine execution turns. Testing is strictly checked later on in CI/CD.
- [ ] NO BUILD CHECKING (TOTAL BAN): Never run build commands (`go build`, `npm run build`, compiler checks) to verify compilation. Build verification is checked later on in CI/CD.
- [ ] NO PER-FILE COMMITTING (TOTAL BAN): Never commit each file individually as you work. All modified files across the turn must be accumulated and committed together in a single atomic commit at the final step.
- [ ] NO PREMATURE TURN CLOSING BEFORE EXECUTION (TOTAL BAN): Never halt execution, conclude the turn, or ask the user for permission after generating specs or subtasks. Planning constitutes only 50% of the task budget; you must proceed unconditionally to Phase 2 code execution.
- [ ] GITMAP HEAVY USAGE: Heavily leveraged GitMap commands (`cpf`, `cpb`, `cpr`, `search`, `find`, `pwsh`) for discovery, execution, and commits. Never ran `pull-all` (`gitmap pa` or `gitmap pae`) unconditionally during routine turns.

## MUST FOLLOW NON-NEGOTIABLE

Listen, past runs of these turns have been sloppy and careless: wrong step counts, partial task lists dumped into chat instead of files, plans half-filled with placeholders, coding guidelines bypassed, detailed specs chopped into useless junk, uppercase README files left uncorrected, and explicit user instructions softened. Stop doing that. Read the whole codebase, confirm root `readme.md` is strictly lowercase, find the root cause, capture commands and tasks without omitting an item, write the spec and memory files in the right paths, group commits with clear messages, and push everything to git before ending. Going deep IS the job. Violating this is auto-reject.

--

## 🚨 Highest Priority Instructions (Appended User Tasks & Instructions Below)

[PASTE USER REQUEST / TASK INSTRUCTIONS HERE — THE AGENT MUST EXECUTE WHATEVER IS WRITTEN BELOW WITH ABSOLUTE PRIORITY AND PRECEDENCE OVER ALL GENERAL GUIDELINES ABOVE]