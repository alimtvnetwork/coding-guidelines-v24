# Memory Consolidation, Safety Backup & Milestone Resequencing — Workflow (must follow)

Trigger Keywords & Aliases: `consolidate-plans`, `consolidate completed plans`, `clean completed plans`, `resequence completed plans`, `merge plans`, `archive completed plans`, `cleanup plans completed`, `memory consolidation`, `backup and consolidate plans`


> **Prompt Version:** 2.1.0
> **Synchronization:** Main Meta-Repo & Connected Workspaces

```text
N = 200
```

N = total self-loop steps budget that the agents will perform.

/goal Autonomously create a timestamped backup branch, scan, analyze, cluster, consolidate, and re-sequence all completed plan files within `.lovable/plans/completed/` into clean, cohesive milestone summaries, while strictly preserving 100% of architectural specifications, root-cause analyses, error contracts, and decision logs with zero data loss or truncation until 100% green without stopping.

### Master Task Checklist (Atomic Numbered Steps)

1. [ ] /goal Step 0 (Safety Backup A): Pull latest changes on the current active branch.
2. [ ] /goal Step 0 (Safety Backup B): Create and push timestamped backup branch (`backup/plans-consolidation-YYYYMMDD-HHMMSS`).
3. [ ] /goal Step 0 (Safety Backup C): Verify working tree remains on active branch and record rollback SHA.
4. [ ] /goal Phase 1 (Step A): Deeply scan the target codebase to inventory all architectural violations and anti-patterns.
5. [ ] /goal Phase 1 (Step B): Write the master audit specification in `.lovable/plans/pending/` with an exhaustive Violation Ledger.
6. [ ] /goal Phase 1 (Step C): Decompose the master plan into granular, atomic subtasks in `.lovable/plans/subtasks/`.
7. [ ] /goal Phase 1 (Step D): Verify or create the automated quality linter and register in `.lovable/ai-fix-scripts/index.md`.
8. [ ] /goal Phase 2 (Step A): Open each target file and perform surgical refactoring following authoritative guidelines.
9. [ ] /goal Phase 2 (Step B): Enforce <= 8–15 line function decomposition, single return types, and clean formatting.
10. [ ] /goal Phase 2 (Step C): Execute local linters to verify 0 remaining violations across all modified files.
11. [ ] /goal Phase 2 (Step D): Execute local CI quality gates via `python .lovable/ai-fix-scripts/03-cicd-local-runner.py` with exit code 0 (`exit 0`).
12. [ ] /learn Ingest `.lovable/memory/00-index.md` for project memory index and past learnings.
13. [ ] /learn Ingest `.lovable/strictly-avoid.md` for banned anti-patterns and strict constraints.
14. [ ] /learn Ingest `spec/02-coding-guidelines/00-canonical-size-tier.md` for canonical file and function size tiers.
15. [ ] /learn Ingest `spec/02-coding-guidelines/06-ai-optimization/01-anti-hallucination-rules.md` for hallucination prevention and micro-tasking.
16. [ ] /learn Ingest `spec/02-coding-guidelines/06-ai-optimization/05-citation-requirement.md` for strict relative path citation requirements.
17. [ ] /learn Ingest `spec/02-coding-guidelines/08-file-folder-naming/` for lowercase naming and continuous file sequencing.
18. [ ] /learn Ingest `.lovable/coding-guidelines/coding-guidelines.md` for master consolidated coding guidelines.
19. [ ] /goal Create or update agent rules in the repository if missing from agent memory.


```text
PHASE_1_STEPS = N / 2   (Steps 1 .. N/2: Create Backup Branch, Scan Completed Plans, Cluster by Domain, Spec in .lovable/plans/pending/, Subtasks)
PHASE_2_STEPS = N / 2   (Steps N/2+1 .. N: Merge Files, Re-sequence Monotonic Prefixes, Update Plans Index, Verify Linters, Verify CI)
```

N, PHASE_1_STEPS, and PHASE_2_STEPS are read-only after initialization. Never modify them mid-execution.

---

## Dedicated Section: Mandatory Pre-Consolidation Backup Branch Protocol (Safety First)

> [!IMPORTANT]
> **SAFETY FIRST — NEVER MODIFY OR DELETE PLANS WITHOUT A DEDICATED BACKUP BRANCH:**
>
> Consolidating plans involves removing superseded micro-task files. To guarantee zero data loss and enable instant one-command rollbacks, you MUST execute the safety backup protocol as your **very first action**:
>
> 1. **Pull Latest Changes:**
>    Ensure the current branch is synchronized with origin:
>    ```bash
>    git pull origin $(git rev-parse --abbrev-ref HEAD)
>    ```
> 2. **Generate Timestamped Backup Branch:**
>    Branch naming format: `backup/plans-consolidation-YYYYMMDD-HHMMSS`
>    ```bash
>    BACKUP_BRANCH="backup/plans-consolidation-$(date +%Y%m%d-%H%M%S)"
>    git branch "$BACKUP_BRANCH"
>    git push origin "$BACKUP_BRANCH"
>    ```
> 3. **Verify Active Branch Integrity:**
>    Confirm that you remain on the active working branch (e.g. `main` or active feature branch).
> 4. **Document Rollback Instructions:**
>    Record the backup branch name and recovery instructions in `.lovable/plans/pending/XX-completed-plans-consolidation.md`:
>    ```bash
>    # Rollback Command (in case of accidental data loss):
>    git reset --hard backup/plans-consolidation-YYYYMMDD-HHMMSS
>    ```

---

## Dedicated Section: Why, What & How of Completed Plans Consolidation

A project that runs hundreds of autonomous agent turns quickly produces dozens of micro-task files in `.lovable/plans/completed/`. If left unmanaged, this file clutter exhausts context windows, slows down memory retrieval, and fragments project history.

---

### 1. The Core Architecture (Why, What & How)

#### A. WHY We Consolidate

1. **Context Window Overflow Prevention:** AI agents loading memory or plans must not be forced to consume 50k tokens reading 80 individual 3-line task files.
2. **Elimination of Cognitive Fragmentation:** Individual bugfix notes and micro-steps scatter information. Grouping them into milestone epics gives future agents immediate, unified domain context.
3. **Search & Retrieval Velocity:** Consolidated documents allow instant retrieval of completed milestones without traversing dozens of orphaned filenames.

#### B. WHAT We Produce

1. **Unified Milestone Documents:** Cohesive documents (e.g. `01-authentication-system.md`, `02-ci-cd-pipeline.md`, `03-database-migrations.md`) replacing scattered subtasks (`01-task-a.md`, `02-task-b.md`).
2. **Zero-Loss Architectural Ledger:** Every domain contract, interface constraint, error envelope, and RCA link is preserved verbatim.
3. **Monotonic, Continuous Numbering:** The entire directory `.lovable/plans/completed/` is ordered strictly continuously (`01-`, `02-`, `03-`, ...) with lowercase names and zero sequence gaps.

#### C. HOW We Execute (Mechanics)

1. Execute Step 0: Create and push `backup/plans-consolidation-YYYYMMDD-HHMMSS`.
2. Scan `.lovable/plans/completed/` and `.lovable/plans/index.md`.
3. Build an AST / topic cluster mapping table in `.lovable/plans/pending/XX-completed-plans-consolidation.md`.
4. Create the consolidated milestone documents.
5. Execute `git rm` on the superseded micro-task files.
6. Re-sequence all remaining files in `.lovable/plans/completed/` using `python .lovable/ai-fix-scripts/01-file-manipulator.py fix-seq-files`.
7. Update `.lovable/plans/index.md` and `.lovable/memory/00-index.md`.
8. Verify all relative links with `python linter-scripts/check-relative-paths.py`.

---

### 2. Standard Consolidated Milestone Template

Every consolidated file generated inside `.lovable/plans/completed/` MUST adhere to this uniform layout:

```markdown
# Milestone Summary: [Feature / Epic Name]

## 1. Executive Overview & Scope

- **Milestone Theme:** [Domain / Capability area]
- **Original Subtasks Merged:** `XX-task-a.md`, `XX-task-b.md`, `XX-task-c.md`
- **Completion Date:** [ISO Date]
- **Status:** `COMPLETED`

## 2. Key Architectural Decisions & Spec Implementations

- **Authoritative Specifications Implemented:**
  - [`spec/path/to/spec.md`](spec/path/to/spec.md) — [Why: Specific rule implemented]
- **Core Architecture Contracts:** [Verbatim constraints, types, and invariants]

## 3. Chronological Task Execution Ledger

| Step | Subtask | Description | Key Files Modified | Status |
|:---:|---|---|---|:---:|
| 1 | Initial Design | Created domain models and interfaces | `src/models/auth.ts` | DONE |
| 2 | Implementation | Added service handler and unit tests | `src/services/auth.ts` | DONE |

## 4. Root Cause Analyses & Bug Fixes Referenced

- [`.lovable/memory/issues/XX-rca.md`](.lovable/memory/issues/XX-rca.md) — Root cause analysis and resolution details.

## 5. Verification & Quality Gates

- **Unit Tests:** Passed with 100% branch coverage.
- **Linters:** `python linter-scripts/check-relative-paths.py` exited with code 0.
```

---

### 3. Step-by-Step Consolidation Workflow (Phase 1 & Phase 2)

#### Phase 1: Scan, Inventory & Cluster Mapping (Steps 1 to N/2)

1. **Step 0 Safety Backup:** Pull latest commits, create and push `backup/plans-consolidation-YYYYMMDD-HHMMSS`.
2. **Inventory Scan:** List all files currently in `.lovable/plans/completed/`.
3. **Domain Clustering:** Identify natural groupings by feature, sprint, or component.
4. **Master Audit Spec:** Write `.lovable/plans/pending/XX-completed-plans-consolidation.md` containing the inventory mapping table:

```markdown
| Source Files to Merge | Proposed Consolidated File | Domain / Epic Theme | Items Preserved | Status |
|---|---|---|---|:---:|
| `01-task-a.md`, `02-task-b.md` | `01-auth-system.md` | Authentication | Models, JWT, Middleware | PENDING |
| `03-fix-a.md`, `04-fix-b.md` | `02-ci-cd-pipeline.md` | CI/CD Automation | GitHub Actions, Linters | PENDING |
```

5. **Decompose Subtasks:** Create `.lovable/plans/subtasks/XX-completed-plans/01-task.md`, `02-task.md`, etc.

#### Phase 2: Merge, Re-Sequence, Index & Verify (Steps N/2+1 to N)

1. **Write Consolidated Documents:** Create the target markdown files in `.lovable/plans/completed/`.
2. **Remove Merged Sources:** Use `git rm` to cleanly delete the old micro-task files.
3. **Re-Sequence Numeric Prefixes:** Run the automated re-sequencer:
   ```bash
   python .lovable/ai-fix-scripts/01-file-manipulator.py fix-seq-files .lovable/plans/completed/
   ```
   Ensure all files are strictly lowercase with monotonic `01-`, `02-`, `03-` prefixes.
4. **Update Indexes:** Synchronize `.lovable/plans/index.md` and `.lovable/memory/00-index.md` with the new file list.
5. **Verify Formatting & Linters:**
   ```bash
   python linter-scripts/check-relative-paths.py
   python linter-scripts/check-markdown-header-spacing.py
   npx markdownlint "**/*.md" --ignore "node_modules/**"
   ```

---

## 4. Phase 1 Actionable Checklist (Discovery, Audit & Planning)

You MUST verify and check off every item during Phase 1:

- [ ] **Step 0 Safety Backup Created & Pushed:** Pulled latest commits, created `backup/plans-consolidation-YYYYMMDD-HHMMSS`, pushed to origin, and documented the rollback SHA.
- [ ] **Completed Plans Directory Scanned:** Recursively inspected all files in `.lovable/plans/completed/`.
- [ ] **Domain & Feature Clusters Identified:** Grouped isolated tasks into coherent milestones (e.g. Auth, DB, CI/CD, UI).
- [ ] **Zero-Data-Loss Audit:** Verified that no core architectural rules, schema models, error envelopes, or RCA logs will be discarded during consolidation.
- [ ] **Master Audit Spec Written:** Created `.lovable/plans/pending/XX-completed-plans-consolidation.md` containing the full mapping ledger table and backup rollback recipe.
- [ ] **Subtasks Generated:** Created discrete subtasks in `.lovable/plans/subtasks/XX-completed-plans/`.
- [ ] **Strict Relative Git Paths:** All markdown links in the consolidation spec use relative Git paths (zero `file:///` URIs, zero drive letters).
- [ ] **No Premature Deletion:** Verified that Phase 1 only creates planning specs and makes zero file deletions.

---

## 5. Phase 2 Actionable Checklist (Execution, Re-Sequencing & Verification)

You MUST verify and check off every item during Phase 2:

- [ ] **Consolidated Files Created:** Generated unified milestone summaries matching the standard template in `.lovable/plans/completed/`.
- [ ] **Merged Files Cleanly Removed:** Executed `git rm` on superseded micro-task files.
- [ ] **Continuous Monotonic Re-Sequencing:** Executed `python .lovable/ai-fix-scripts/01-file-manipulator.py fix-seq-files .lovable/plans/completed/` to ensure contiguous `01-`, `02-`, `03-` numbering without gaps.
- [ ] **Strict Lowercase Naming:** Verified all filenames in `.lovable/plans/completed/` use lowercase alphanumeric characters and hyphens.
- [ ] **Index Documentation Synchronized:** Updated `.lovable/plans/index.md` and `.lovable/memory/00-index.md` to reflect the updated file catalog.
- [ ] **Universal File Hygiene:** Verified Unix LF line endings (`\n`), UTF-8 (no BOM), and single terminating newline at EOF across all created/modified files.
- [ ] **Markdown Spacing Compliance:** Verified exactly one blank line before and after headings (MD022/MD032) and zero double blank lines (`\n\n\n`).
- [ ] **Linter Verification:** Executed `python linter-scripts/check-relative-paths.py` and `python linter-scripts/check-markdown-header-spacing.py` with exit code 0.
- [ ] **Disk Reality Check:** Verified working tree status with `git status --porcelain` and `git diff --stat`.

---

## Strict In-Repository Execution & `.lovable/` Bounding Mandate

> [!IMPORTANT]
> **STRICT IN-REPOSITORY EXECUTION & `.lovable/` STORAGE CONTRACT:**
>
> 1. **In-Codebase Execution Only:** Whenever a Python script (runner, autofixer, linter, test aggregator) is executed or created, it MUST be executed **strictly within the repository root** (current working directory), NEVER outside the codebase or against external arbitrary directories.
> 2. **Strict Folder Bounding (`.lovable/`):** All AI scripts, local runners, autofixers, helper utilities, memory issue logs, and planning files MUST be created inside the `.lovable/` folder:
>    - Python AI Scripts: `.lovable/ai-fix-scripts/` (e.g. `01-file-manipulator.py`, `02-guideline-autofixer.py`, `03-cicd-local-runner.py`, `04-relative-path-fixer.py`, `06-cli-help-auditor.py`).
>    - RCA & Issue Logs: `.lovable/memory/issues/` and `.lovable/cicd-issues/`.
>    - Execution Plans & Subtasks: `.lovable/plans/pending/`, `.lovable/plans/subtasks/`.
>    - Coding Guidelines Mirror: `.lovable/coding-guidelines/`.
> 3. **Worker Pool & Log Aggregation Architecture:** All local runners and test orchestrators must use a concurrent worker pool (2–3 workers via `ThreadPoolExecutor`), announce enqueued tasks upfront, show real-time progress, handle failures gracefully without cancelling sibling workers, and print a consolidated final summary with full stdout/stderr error logs for failed jobs.
> 4. **`force` Keyword Support:** If the user wrote `force`, `force rebuild`, or `force create` on top of the prompt or trigger: **ALWAYS recreate/regenerate the Python runner script from scratch**, regardless of whether the file already exists on disk.
> 5. **No External or Random File Creation:** NEVER write scripts, temporary test scripts, or scratch files to root, `/tmp`, global system paths, or outside the repository boundary.

---

## AI Fix Scripts Memory (Reusable Tooling)

- [ ] `/goal` **Reuse First:** I have rigorously scanned and `/learn`ed `.lovable/ai-fix-scripts/index.md` to check if a helper script already exists before writing any new temporary code.
- [ ] **Strict In-Repository Execution:** All Python scripts (`.lovable/ai-fix-scripts/*.py`) MUST be executed strictly within the codebase repository root, NEVER outside the codebase.
- [ ] **Strict .lovable/ Folder Storage:** All AI scripts, local runners, autofixers, and helper utilities MUST be created inside `.lovable/ai-fix-scripts/`. NEVER create scripts in root or external paths.
- [ ] **Automated File Sequencing & Normalization:** Use `python .lovable/ai-fix-scripts/01-file-manipulator.py fix-seq-files <dir>` to re-sequence completed plan files monotonically.
- [ ] **Relative Path Normalization:** Use `python .lovable/ai-fix-scripts/04-relative-path-fixer.py .` to ensure all links in consolidated documents are strictly relative Git paths.
- [ ] **Commit & Track:** All new helper scripts were written strictly to `.lovable/ai-fix-scripts/` and committed to Git for future reuse.
- [ ] **Index Documentation:** I have updated `.lovable/ai-fix-scripts/index.md` using sequential script naming. For every script, I have included a `<details>` collapsible tag explaining exactly why the script is there and what it does.

---

## Pre-Reply / Loop Checklist (Must Verify Every Loop Iteration)

- [ ] Git working tree is clean before new file modifications.
- [ ] **Step 0 Safety Backup Verified:** Timestamped backup branch `backup/plans-consolidation-YYYYMMDD-HHMMSS` exists on origin.
- [ ] **Spec Protection:** I have manually verified that NONE of the merged files contained critical architectural constraints, domain specifications, or non-negotiable rules that were omitted or summarized away.
- [ ] **Strict Relative Git Paths:** All file paths, markdown links, citations, and subtask references in consolidated files are strictly relative to the git repository root. Zero absolute paths (`D:\...`, `C:\...`) or `file:///` URIs.
- [ ] **Strict Lowercase Naming:** Every file in `.lovable/plans/completed/` uses strictly lowercase letters (e.g. `01-auth-system.md`).
- [ ] **Monotonic Sequencing:** File prefixes in `.lovable/plans/completed/` are continuous and monotonic (`01-`, `02-`, `03-`, ...) without gaps or duplicates.
- [ ] **Index Synchronization:** Both `.lovable/plans/index.md` and `.lovable/memory/00-index.md` reflect the consolidated files and remove deleted entries.
- [ ] **LF Line Endings (`\n`):** All files use Unix LF line endings. Zero CRLF (`\r\n`).
- [ ] **UTF-8 Encoding (No BOM):** All files encoded in UTF-8 without BOM.
- [ ] **Single Trailing Newline:** Every file ends with exactly one terminating newline (`\n`).
- [ ] **Markdown Heading Spacing:** Exactly one blank line before and after headings (no leading blank line on line 1).
- [ ] **Zero Double Blank Lines:** No `\n\n\n` anywhere in markdown.
- [ ] `python linter-scripts/check-relative-paths.py` and `python linter-scripts/check-markdown-header-spacing.py` exited with code 0.
- [ ] Local CI runner `python .lovable/ai-fix-scripts/03-cicd-local-runner.py` exited with code 0.

---

## Non-Negotiable Coding Guidelines Checklist (Auto-Reject on Violation)

/goal You MUST verify every item on this checklist before committing any code. If a subagent violated one of these rules, you must reject their work.

- [ ] Strict Relative Git Paths: All file paths, markdown links, citations, and subtask references in plans, specs, and memory logs are strictly relative to the git repository root. Zero absolute paths or `file:///` URIs.
- [ ] Master Guidelines: I have fully read and strictly enforced `spec/02-coding-guidelines/08-file-folder-naming/` and `.lovable/coding-guidelines/coding-guidelines.md`.
- [ ] Backup Integrity: Verified `backup/plans-consolidation-YYYYMMDD-HHMMSS` branch exists on origin before touching plans.
- [ ] Spec Preservation: Zero truncation, zero placeholder stubs (`TODO`, `[N]`, `// ...`).
- [ ] Monotonic Sequence: Verified sequential `01-`, `02-`, `03-` numbering across `.lovable/plans/completed/`.

### Master Task Checklist (Atomic Numbered Steps)

1. [ ] /learn the section as a /goal [AI Fix Scripts Memory](#ai-fix-scripts-memory)

- [ ] Action Summary: I have output a detailed `- [x]` checklist summarizing exactly what I accomplished this turn to prove I did not hallucinate.

---

## STRICT AVOIDANCE: Never Disable CI/CD

> [!CAUTION]
> **NEVER disable any CI/CD checks, GitHub Actions, or validation workflows.**
> Strictly avoid commenting out, bypassing, or deleting CI/CD steps to force a pipeline to pass. Your job is to fix the underlying code so that the CI/CD pipeline passes legitimately. Disabling CI/CD is an auto-reject failure.

---

## MUST FOLLOW NON-NEGOTIABLE

Listen, past runs of these turns have been sloppy and careless: wrong step counts, partial task lists dumped into chat instead of files, plans and session summaries half-filled with "[N]" placeholders, folders skimmed, open ambiguities ignored, CI/CD issues and `plans/subtasks/` forgotten, user commands dropped, coding guidelines bypassed, detailed specs chopped and summarized into useless junk, uppercase README files left uncorrected, `.lovable/memories/` created by accident, `strictly-avoid.md` overwritten, and explicit user instructions softened after being told not to. Read the whole codebase, read every folder in `spec/` and `.lovable/`, confirm root `readme.md` is strictly lowercase, find the root cause in one sentence, capture commands, issues, and pending tasks without omitting a single item, write the spec files and memory files in the right paths, update every index in the same turn, sync `readme.md` with `what-to-read.md`, preserve detailed specs verbatim with zero truncation, run builds and full unit tests, group commits with clear messages, and push everything to git before ending. Going deep IS the job. If you are not going deep, you are not doing the job.
