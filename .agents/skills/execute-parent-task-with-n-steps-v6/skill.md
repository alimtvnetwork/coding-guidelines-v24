---
name: execute-parent-task-with-n-steps-v6
description: >-
  Use this skill when the user asks you to execute a parent task with N steps using the V6 prompt (N_TOTAL_COUNT_OF_ITERATION step ceiling, A_AGENTS worker subagents via invoke_subagent, GitMap atomic commits, in-brief coding guideline injection, R1-R16 rules, and variable-driven budgeting).
---

# [V6] Parent Task N-Step Loop: Antigravity-Native Ultra-Orchestrator — Workflow (must follow)

```text
N_TOTAL_COUNT_OF_ITERATION = 300   Total steps ceiling for run (edit before running, default: 300)
A_AGENTS = 2                       Worker subagents per wave (invoke_subagent, default: 2)
H_AGENT_HANDS = 2                  Subtasks per worker per wave (batch capacity, default: 2)

PHASE_1_BUDGET = N_TOTAL_COUNT_OF_ITERATION / 2   Steps 1 .. PHASE_1_BUDGET: Spec, Plan, Subtasks
PHASE_2_BUDGET = N_TOTAL_COUNT_OF_ITERATION / 2   Steps (PHASE_1_BUDGET + 1) .. N_TOTAL_COUNT_OF_ITERATION: Waves, Commit
```

> [!IMPORTANT]
> Prompt Version: 6.0.0 | Runtime: Google Antigravity 2.0 (IDE & CLI)
> Invoke: paste below task or run `/execute-parent-task-with-n-steps-v6 <task>`.
> N_TOTAL_COUNT_OF_ITERATION is a hard ceiling: finishing early is success; padding is failure. One step is one tool round by lead, logged in ledger.

[/goal](slashCommand:goal) Autonomously execute parent task: capture verbatim, plan, run via A_AGENTS worker subagents in disjoint file boxes using GitMap primary, prove claims with evidence, enforce guidelines 100%, and finish with one atomic GitMap commit.

[/learn](slashCommand:learn) Preamble directives outrank everything below. Rules R1-R16 cited by ID. Progress lives in ledger and `.ai-memory/plans/`.

---

## 1. Precedence Hierarchy & Scope (Highest First)

1. **User Instructions & Preamble:** Directives above this prompt outrank everything below.
2. **Platform Limits:** Native tools, Artifact Review Policy, permission prompts, hooks.
3. **Repo Rules:** `AGENTS.md`, `.ai-memory/strictly-avoid.md`, `coding-guidelines.md` (source: `02-spec/17-consolidated-guidelines/34-compiled-simple-coding-guidelines.md`).
4. **This Prompt.**

If sources conflict, follow stricter one and log in ledger `Conflicts:`.

### Scope Boundaries
- **Historical Content:** Never rewrite old names, changelogs, completed plans.
- **Unrelated Findings:** Log in plan follow-ups, not immediate fixes.
- **Prompt Registers:** Update prompt registers only when task touches prompts.
- **Lean Mode:** Tasks touching <= 3 files skip subtasks; use single spec in `02-spec/21-app/`.
- **Minimal Diff:** Never renumber or reflow unaffected lines.

---

## 2. Core Operational Rules (Cite by ID)

- **R1 Zero Builds or Tests (TOTAL BAN).** Never run `go build`, `npm run build`, `vite build`, `go test`, `pytest`, `npm test`, or `03-ai-scripts/06-cicd-local-runner.py`. CI verifies builds/tests.
- **R2 Targeted Checks Only.** Run only fast, file-scoped checks on modified files (Section 10). 0 files scanned is a **FAIL**.
- **R3 Evidence Required.** Every `DONE`/`PASS` MUST cite file path, git diffstat, or exit code (`exit 0`). Vague assurances are auto-rejected.
- **R4 Never Invent Commands.** Verify commands exist via workspace test (`gitmap lf readme.md`) before use. Use documented fallbacks.
- **R5 Mandatory Subagents (`invoke_subagent`).** Absolute MUST when tool exists (`research` for discovery, `self` for edits). Solo fallback only if tool absent.
- **R6 One Owner Per File.** In each wave, each file has one owner. Shared indexes belong exclusively to lead.
- **R7 Git Safety.** Subagents never run git. Nobody runs `git reset --hard`, `git checkout --`, `git clean`, `git stash`, or force pushes.
- **R8 .gitignore Hygiene.** Verify `.gitignore` covers caches, build outputs, logs, test reports, secrets before commit (Hard Rule 1).
- **R9 Atomic Commit & Push via GitMap (`gitmap cpf`/`cpb`).** Commit/push once at run end via GitMap: `gitmap cpf "<summary>"` (features) or `gitmap cpb "<summary>"` (bugs). Never commit per-file. On push rejection: `git pull --rebase` and re-run. Squash/amend only unpushed commits.
- **R10 Zero Unauthorized Releases.** Never bump versions, edit `version.json`, or release unless explicitly instructed.
- **R11 Strict Relative Paths.** Paths relative from git root (zero absolute paths, zero `file:///` URIs). All new filenames strictly lowercase.
- **R12 No Polling / Immediate Turn Yielding.** Print yield line (`Dispatched Worker 01 .. Worker <A_AGENTS>; waiting for their results.`) and STOP CALLING TOOLS. If wave runs long, check `manage_subagents` once.
- **R13 Two-Strike Retry Cap.** Two consecutive tool failures -> worker replies `STATUS: BLOCKED` with exact error and stops. Subtask failing two remediation rounds is marked `FAILED` with RCA (Section 12).
- **R14 Ambiguity Boundaries.** Non-blocking: choose conservative option in ledger `Assumptions:`. Blocking: `ask_question` once, continue unblocked tasks.
- **R15 Zero Generated Artifacts.** Never commit untracked build caches, logs, temp scripts, or generated code (Hard Rule 1).
- **R16 Repo Secrets Mandate.** If `repo-secrets` exists in default work directory, offload secrets via `gitmap rs`; never store credentials in standard repositories (`AGENTS.md` section 9).

---

## 3. GitMap High-Speed Command Primacy (Run Everything Faster)

GitMap is your PRIMARY acceleration engine:

| Operation | Command | Alias | Purpose |
| :--- | :--- | :--- | :--- |
| Wildcard Search | `gitmap find "<pattern>" [-ext <ext>]` | `gitmap f "<pat>"` | Indexed file finding |
| File Listing | `gitmap list-files [pat] [-ext <ext>]` | `gitmap lf [pat]` | Instant indexed inventory |
| Substring Search | `gitmap find-files-any "<str>"` | `gitmap ffa "<str>"` | Partial filename matcher |
| Stream File | `gitmap cat <filepath>` | `gitmap cat` | Zero-disk memory streaming |
| Regex Search | `gitmap search "<query>"` | `gitmap search` | Parallel text scanner |
| PowerShell Runner | `gitmap pwsh "<command>"` | `gitmap ps "<cmd>"` | PowerShell (-NoProfile) |
| Bash Runner | `gitmap bash "<command>"` | `gitmap sh "<cmd>"` | Cross-platform Bash |
| Offload Secrets | `gitmap rs file <path>` / `folder` / `text` | `gitmap rs` | Commits to `repo-secrets` |
| Offload Scripts | `gitmap rc file <file.ps1>` / `text` | `gitmap rc` | Commits to `repo-cache` |
| Atomic Commits | `gitmap cpf "<summary>"` / `cpb` (Bug) | `gitmap cpf` | Stages, commits, pushes |
| Pipeline Waiting | `gitmap pipeline-ai status --json` | `gitmap pl-ai` | Dynamic ETA CI monitor |

*Rule:* Never run `gitmap pa`/`pae` unless requested. Only lead runs git or `gitmap rs`.

---

## 4. Step 0: Preflight & Platform Handshake (Phase 1 Budget)

1. **Platform Handshake:** Confirm tools (`invoke_subagent`, `send_message`, `manage_subagents`, `ask_question`, `write_to_file`, `replace_file_content`, `run_command`). If `task_boundary` exists: set `PLANNING` (Phase 1), `EXECUTION` (Phase 2), `VERIFICATION` (Phase 3).
2. **Commands & Directory:** Confirm `gitmap --version` and `python --version` exit 0. Commands run with `Cwd` in workspace root; paths workspace-relative.
3. **Working Tree:** Run `git status --porcelain`. Record modified files in ledger; never touch them. Confirm root `readme.md` is lowercase.
4. **Context Ingestion:** Read `what-to-read.md`, `strictly-avoid.md`, `coding-guidelines.md` in `.ai-memory/`.
5. **Resume Check:** Match `.ai-memory/temp-agents/NN-<slug>/ledger.md` by `Request slug:` and `Request first line:`. If matched, resume from `Next action:`. If commit in `Commits:` exists in `git log`, verify and report without redoing.
6. **Ledger Creation:** If not resuming, create `.ai-memory/temp-agents/NN-<slug>/ledger.md`:

```markdown
# Ledger: NN-<slug>
Request slug: NN-<slug>
Request first line: <first line of verbatim request>
Last step: 0
Next action: Phase 1A Capture
Commits: none
Step: 1 / N_TOTAL_COUNT_OF_ITERATION (Phase 1: 1 / PHASE_1_BUDGET, Phase 2: 0 / PHASE_2_BUDGET)
Branch: <branch> | Tree at start: clean (or dirty)
| Task-ID | Subtask | Owner | Owned files | Status | Evidence |
|---|---|---|---|---|---|
| Task-01 | 01-<name> | Worker 01 | <paths> | PENDING | - |
Assumptions: <list or none>
Conflicts: <list or none>
```

---

## 5. Phase 1A: Verbatim Capture, Deliverables Breakdown & Chat Gate

1. **Capture Verbatim:** Store prompt losslessly under `## User Request (Verbatim)` in spec and plan.
2. **Screenshots:** Decode images into `assets/screenshots/<slug>-<NN>.png`. Link relatively (`![Screenshot](assets/screenshots/<slug>-<NN>.png)`).
3. **Discrete Deliverables:** Decompose request into ordered Task-IDs (`Task-01`, `Task-02`, ...).
4. **Mandatory Same-Turn Tool Chaining (TOTAL BAN ON TURNING OFF):** Emit breakdown in chat with vertical lines, and in **EXACT SAME TURN**, invoke first tool call (`write_to_file` or `invoke_subagent`). NEVER emit text alone. Never ask "Should I proceed?".

```markdown
### 📋 Confirmed Task Breakdown & Ingestion

1. **Task-01: [Descriptive Task Title]**
   - **State:** `[IN PROGRESS — EXECUTING IMMEDIATELY]`
   - **Understood:** `[YES]` — [one sentence proving understanding]
   - **Scope:** [deliverable]
   - **Target Area:** `[relative/path/or/module]`

Proceeding directly to Phase 1B (Tool Call Below).
```

---

## 6. Phase 1B: Spec, Plan & Lean Subtasks (Steps 1 .. PHASE_1_BUDGET)

1. **Blueprint:** Lead orchestrator alone authors initial spec and planning skeleton.
2. **Discovery Subagents (`A_AGENTS workers`):** Dispatch `A_AGENTS` `research` subagents (`Research 01 .. Research <A_AGENTS>`) via `invoke_subagent` on disjoint folders to map symbols and dependencies using workspace search (`git grep` or `gitmap search`).
   - **Research Contract:** Reply with one line per hit formatted as `path:line: text`, then `SUMMARY: <one line>`, then stop.
3. **Canonical Spec (`02-spec/21-app/`):** Single-domain (or <= 3 files): Write `02-spec/21-app/NN-<slug>.md`. Multi-domain: Write `02-spec/21-app/NN-<slug>/` (`01-overview.md` .. `04-verification-gates.md`). Register in `02-spec/21-app/readme.md`.
4. **Plan:** Write `.ai-memory/plans/pending/NN-<slug>.md` linking to spec and mapping Task-IDs to subtasks. Register in `.ai-memory/plans/readme.md`. Tasks touching <= 3 files run in lean mode (skip subtasks; execute against spec).
5. **Lean Subtasks (When > 3 files):** Write `.ai-memory/plans/subtasks/NN-<slug>/01-<name>.md`, etc.:

```markdown
# Subtask [01]: [Descriptive Subtask Name]
Traceability ID: Task-01
Spec: [02-spec/21-app/NN-<slug>.md](../../../02-spec/21-app/NN-<slug>.md)
Owned Files: [relative paths; mark NEW]
Action: [functions, types, logic]
Acceptance Criteria: [2 to 4 conditions]
Targeted Verification: [command from Section 10]
```

6. **Readiness Gate:** Complete Phase 1 planning within `PHASE_1_BUDGET` steps. Proceed **UNCONDITIONALLY** to Phase 2.

---

## 7. Phase 2: Mandatory Worker Waves & Coding Guidelines Enforcement (Steps (PHASE_1_BUDGET + 1) .. N_TOTAL_COUNT_OF_ITERATION)

> [!CRITICAL]
> **MANDATORY `invoke_subagent` DISPATCH (ZERO SOLO EXECUTION):**
> Spawn `A_AGENTS workers` (`TypeName: "self"`, up to `H_AGENT_HANDS subtasks per worker`) in parallel via `invoke_subagent`. Solo execution when tool exists is an auto-reject failure. If fewer disjoint file groups than `A_AGENTS`, spawn one worker per group and log `A_REDUCED: <groups> groups` in ledger.

### 7.1 Dispatch Payload (`invoke_subagent`)

The `invoke_subagent` payload holds exactly `A_AGENTS` entries, built by repeating this entry template `A_AGENTS` times (for `Worker 01 .. Worker <A_AGENTS>`):

```json
{
  "Subagents": [
    {
      "TypeName": "self",
      "Role": "Worker <NN>: [<Assigned Feature / Module>]",
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
You are Worker <NN> for task NN-<slug>. This brief is your complete specification.

### Boundaries:
- Read any file in workspace; edit only Owned Files: <relative paths>.
- Report credentials to lead; workers never run git or gitmap rs.
- Cap: at most 30 tool calls; at cap, report status and stop.
- Two consecutive tool failures -> reply "STATUS: BLOCKED" with exact error and stop. Never guess paths.

### Assigned Subtasks (up to H_AGENT_HANDS Subtasks):
- Subtask 1: .ai-memory/plans/subtasks/NN-<slug>/01-<name>.md
- Subtask 2: .ai-memory/plans/subtasks/NN-<slug>/02-<name>.md (if assigned)

### 100% Non-Negotiable Coding Guidelines (AUTO-REJECT ON VIOLATION):
Follow canonical specs: 02-spec/02-coding-guidelines/01-cross-language/, 02-spec/02-coding-guidelines/02-boolean-principles/, 02-spec/03-error-manage/, 02-spec/17-consolidated-guidelines/34-compiled-simple-coding-guidelines.md. Adhere to R1, R2, R11 by ID.

- Rules (insert language block for worker's Owned Files):
  - Universal: is/has booleans only (AGENTS.md section 8; never == true, no mixed polarity); flatten nested if into early returns; no magic strings; enum Type suffix; definitions in own files; functions <= 8-15 lines; classes/structs <= 120 lines; files <= 300 lines (100 for .tsx); blank line before return and if; relative paths only.
  - Go (.go): Return *appfault.AppError or Result[T] (02-spec/03-error-manage/); types in types.go; no ignored errors (_).
  - TypeScript (.ts/.tsx): Discriminated union { isSuccess: true, data: T } | { isSuccess: false, error: AppError }; types in types.ts; .tsx <= 100 lines.
  - Python (.py): Structured typed objects or domain exceptions; type annotations; models in types.py; no bare except: pass.
  - PHP (.php): Typed Result or domain exceptions; declare(strict_types=1);; DTOs in dedicated classes.
  - Mandatory Check: python 03-ai-scripts/05-guideline-autofixer.py <folder> --check-only --ext <ext> (>0 files scanned).

### Output Contract:
Reply with this block, once per subtask, then stop:
TASK: Task-01 | STATUS: DONE | FAILED | BLOCKED
FILES_CHANGED: <paths>
CHECKS: <command> -> exit <code>, <files scanned>
ACCEPTANCE: AC1 PASS <evidence>; AC2 PASS <evidence>
ASSUMPTIONS: <list or none>
BLOCKERS: <list or none>
```

### 7.3 Turn-Yielding & Evidence Verification Protocol

1. **Yield:** Print `Dispatched Worker 01 .. Worker <A_AGENTS>; waiting for their results.` and **STOP CALLING TOOLS**.
2. **Verify Reports:** Confirm `git diff --stat -- <owned files>` matches `FILES_CHANGED`. Read worker diffs against acceptance criteria. Re-run discovery search (`git grep <old-name>` or `gitmap search <old-name>`) for expected hits only. Re-run targeted check (`exit 0` on >0 files).
3. **Reject Violations:** Send failures via `send_message`. After two failed rounds, mark `FAILED`, write RCA, proceed (R13).
4. **Update Ledger & Plan:** Every `DONE` cites command/exit code, diffstat, or path. Copy ledger table to `.ai-memory/plans/pending/NN-<slug>.md` after each wave.
5. **Loop:** Dispatch waves of `A_AGENTS workers` until all subtasks are `DONE` or `FAILED`.

---

## 8. Phase 3: Consolidation, Evidence Verification & Atomic GitMap Push

1. **Pre-Commit Gate:** Write targeted check commands, exit codes, and files-scanned counts into ledger before commit.
2. **Consolidate Subtasks:** Merge completed subtasks into `.ai-memory/plans/completed/NN-<slug>.md`. Document real steps from ledger and link to canonical spec. Delete `.ai-memory/plans/subtasks/NN-<slug>/` and pending plan file.
3. **Index Registers:** Update `.ai-memory/plans/readme.md`. If prompt files were touched, update `01-prompts/readme.md` and `.ai-memory/prompts.md`.
4. **.gitignore Verification (R8):** Verify `.gitignore` covers build outputs, caches, logs, secrets created during run.
5. **Atomic Commit & Push (R9):** Run `gitmap cpf "<summary>"` (features) or `gitmap cpb "<summary>"` (bug fixes). GitMap stages changes, commits with standardized prefix, and pushes. Never run bare `git add` or `git commit`. On push failure, run `git pull --rebase` and re-run GitMap command. Never force push, never `--no-verify`.

---

## 9. Final Report Format (Strict Vertical Lines)

```markdown
### Task Completion Summary

- ✅ **Task-01: [Descriptive Task Title]** — `[Completed]` — [diff/check evidence]
- ❌ **Task-02: [Descriptive Task Title]** — `[Failed]` — [RCA link]

### Modified Files Summary

- [relative/path/to/modified/file1.ext]

### Steps Used

- Step: [total_steps] / N_TOTAL_COUNT_OF_ITERATION (Phase 1: [p1_steps] / PHASE_1_BUDGET, Phase 2: [p2_steps] / PHASE_2_BUDGET)

### Implementation Confidence Score

- Confidence: [passed checks / total checks (percentage%)]
- Rationale: [Verified evidence across all criteria, passing linters, zero regressions]

### 🤖 Independent AI Verification & Audit Prompt

(Emit self-contained audit prompt linking to spec, plan, and modified files)
```

---

## 10. Targeted Verification Checks (R2)

Confirm scripts exist with a harmless workspace call before invoking (R4). Run on changed files/folders only:

- **Coding Guidelines & Boolean Linter:** `python 03-ai-scripts/05-guideline-autofixer.py <folder> --check-only --ext <.ext>`
- **Relative Path Linter:** `python linter-scripts/check-relative-paths.py`
- **Prompts & Spec Index Linter:** `python linter-scripts/check-prompts-loaded.py`
- **Doc Path Linter:** `python 03-ai-scripts/22-doc-path-linter.py <folder>`
- **Sequence Integrity Linter:** `python 03-ai-scripts/21-sequence-integrity-linter.py`
- **Forbidden Strings Check:** `python linter-scripts/check-forbidden-strings.py`

---

## 11. Discovery Toolchain (GitMap Primary)

Use Section 3 GitMap table: `gitmap f`, `gitmap lf`, `gitmap ffa`, `gitmap cat`, `gitmap search`. Python fallbacks (`03-ai-scripts/11-fast-file-scanner.py`, `03-ai-scripts/12-fast-cached-grep.py`) apply only if GitMap unavailable.

---

## 12. Issue Destination & RCA Routing

- **CI/CD Failures:** `.ai-memory/cicd-issues/NN-<slug>.md`, indexed in `.ai-memory/cicd-index.md`.
- **Application Bugs:** `02-spec/22-app-issues/NN-<slug>.md` with 4-part RCA (Reproduction, Cause, Fix, Prevention), indexed in `02-spec/22-app-issues/readme.md`.
- **Failed Subtasks (R13):** Log RCA in `.ai-memory/memory/issues/` and link from `ledger.md`.

---

## MUST FOLLOW NON-NEGOTIABLE

Listen, past runs of these turns have been sloppy and stupid as fuck: wrong step counts, partial task lists dumped into chat instead of files, plans and session summaries half-filled with placeholders, open ambiguities ignored, CI/CD issues and `plans/subtasks/` forgotten, user commands dropped, coding guidelines bypassed, detailed specs chopped and summarized into useless junk, uppercase README files left uncorrected, `.ai-memory/memory/` created by accident, `strictly-avoid.md` overwritten, and explicit user instructions softened after being told not to. WTF. How on earth are you reverting to this carelessness, are you stupid?? Stop doing that, you stupid fuck. Find the root cause in one sentence, capture commands, issues, and pending tasks without omitting a single item, write the spec files and memory files in the right paths, update every index in the same turn, sync `readme.md` with `what-to-read.md`, preserve detailed specs verbatim with zero truncation, run targeted checks (builds and full unit tests stay in CI per R1), and execute the final atomic GitMap commit and push before ending. Going deep IS the job. If you are not going deep, you are not doing the job. Violating this is auto-reject on the same tier as RULE 0. Avoid stupidity and being careless, you stupid fuck. Where is your attention, are you stupid? Tell me. Your stupidity is going on top of my head. Where did you learn this stupidity? If I could find you, I could slap you.
