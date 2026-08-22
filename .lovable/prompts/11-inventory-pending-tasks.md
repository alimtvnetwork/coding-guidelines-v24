# Instruction (must follow): Inventory Pending Tasks (Read-Only Scan & 3-Agent Dispatch Proposal)

> This instruction provides guidelines and directives for inventory pending tasks (read-only scan & 3-agent dispatch proposal).

/goal Perform a strictly read-only scan of the entire repository, `spec/`, and `.lovable/` directory to compile a comprehensive, deduplicated inventory of every pending task, subtask, unresolved issue, and open requirement structured into Execution Waves.

CRITICAL CONSTRAINT: This prompt is strictly for inventorying, structuring, and sequencing pending work. It MUST NOT execute code modifications, build changes, or launch the execution loop. Batch execution is handled by dedicated execution prompts.

---

## Folder Structure Reference (What and Where to Read)

To ensure zero blind spots, the AI must systematically inspect the following authoritative folder structure:

```text
.lovable/
  memory/
    index.md                                    # Master memory index
    what-to-read.md                             # Authoritative reading order
    specs/                                      # Verbatim user directives & requirements
    learned/                                    # Institutional knowledge & learnings
    workflow/                                   # Workflow state markers
  plans/
    index.md                                    # Roll-up index of all plans (pending & completed)
    pending/XX-<slug>.md                        # Active, pending parent execution plans
    completed/XX-<slug>.md                      # Historical completed plans (reference only)
    subtasks/XX-<slug>/SS-<subslug>.md          # Granular subtasks linked to parent plans
  issues/                                       # General pending bug reports
  pending-issues/                               # Active issues queue
  cicd-issues/                                  # CI/CD specific failures & blockers
  ambiguous-questions/
    01-new-ambiguity/XX-<slug>.md               # Blocking & non-blocking open questions
    02-ambiguity-resolved/XX-<slug>.md          # Binding resolved decisions
  suggestions.md                                # Suggestions tracker
  prompts.md                                    # Canonical prompt registry
spec/
  01-spec-authoring-guide/                      # Spec authoring conventions & required files
  02-coding-guidelines/                         # Coding standards & zero-tolerance rules
  03-error-manage/                              # Error management & logging philosophy
  04-database-conventions/                      # Database schema & query rules
  21-app/                                       # App specification & feature requirements
  (note: numbers in spec/<NN>-<slug>/ can switch between projects; dynamically traverse all nested .md files)
readme.md                                       # Root repository guide (strictly lowercase)
```

---

## Step 1: Deep Inspection & Deduplication Protocol

1. Scan Every Pending Source:
   - Open and read `.lovable/plans/index.md`, `.lovable/plans/pending/`, and all `.lovable/plans/subtasks/` files with `Status:` not `completed`.
   - Open and read all files in `.lovable/issues/`, `.lovable/pending-issues/`, and `.lovable/cicd-issues/`.
   - Open and read all open questions in `.lovable/ambiguous-questions/01-new-ambiguity/`.
   - Open and read unfulfilled directives in `.lovable/memory/specs/` and `spec/21-app/`.
   - Open and read active suggestions in `.lovable/suggestions.md`.
2. Deduplicate Across Sources:
   - If a feature is referenced across a spec, a plan, and an issue, consolidate it into ONE primary task with cross-references to all origin files.
3. Step-Count Rubric & Decomposition Alert:
   - Trivial change (1 file, single edit): 1 step.
   - Small change (1-2 files, 1 verification step): 2-3 steps.
   - Standard task (multi-file, logic + UI/backend, test): 4-7 steps.
   - Cross-cutting task (schema + API + UI + full tests): 8-15 steps.
   - Automatic Subtask Decomposition Alert: When a pending task exceeds 7 steps, flag it with `[DECOMPOSITION REQUIRED]` to split it into `.lovable/plans/subtasks/XX-<slug>/` before entering the execution queue.
4. Ambiguity Impact Severity Scoring:
   - High Blast Radius: Blocks multiple core plans or schemas.
   - Medium Blast Radius: Blocks a single isolated feature.
   - Low Blast Radius: Non-blocking cosmetic or detail refinement.

---

## Step 2: Refined Output Format (Structured by Execution Waves)

Present the inventory to the user in this exact markdown structure:

```markdown

# Pending Tasks Inventory

## Executive Summary

- Total Pending Tasks: [N]
- Blocking Ambiguities: [K]
- Pending Plans: [P]
- Pending Issues & CI/CD: [I]
- Unimplemented Spec Scope: [U]
- Sources Scanned: [.lovable/plans/pending/, .lovable/plans/subtasks/, .lovable/issues/, .lovable/cicd-issues/, .lovable/ambiguous-questions/01-new-ambiguity/, spec/]

---

## Execution Waves

### Wave 1: Independent Foundations (DB Schemas, Wrappers, Core Models)

- Can run immediately in parallel across 3 subagents (disjoint files).

#### 1.1 [Task Title]

- Source File(s): `[path/to/file.md]`
- Type: `Plan` | `Issue` | `CI/CD` | `Spec-Scope`
- Status: `Pending` | `In-Progress`
- Estimated Steps: `[N] steps`
- Dependencies: `None`
- Outcome / Intent: [Clear 1-2 sentence description of what "done" looks like]

### Wave 2: Business Logic, Services & Endpoints

- Requires Wave 1 foundations to complete before execution.

#### 2.1 [Task Title]

- Source File(s): `[path/to/file.md]`
- Dependencies: `Wave 1 Task #[X]`
- Estimated Steps: `[N] steps`
- Outcome / Intent: [Description]

### Wave 3: UI Components, Views & Documentation

- Requires Wave 2 business services.

#### 3.1 [Task Title]

- Source File(s): `[path/to/file.md]`
- Dependencies: `Wave 2 Task #[Y]`
- Estimated Steps: `[N] steps`
- Outcome / Intent: [Description]

---

## Blocking Ambiguities (Ranked by Blast Radius)

- [HIGH/MED/LOW] `[slug]`: [Question summary] — *Blocks: Task #[X], Task #[Y]*
```

---

## Step 3: Follow-Up Execution Proposal (Strictly Positive Question Framing)

At the very end of the inventory report, the AI must ask the user whether to trigger execution using the standard 3-agent parallel loop:

```text
Would you like to start the continuous self-loop to execute and resolve these pending tasks in parallel using a maximum of 3 concurrent subagents? (Yes / No)
```

*(If the user answers Yes, the next turn will invoke `execute-robust-loop` or `execute-batched-loop` to begin execution).*

---

## Checklist: What This Instruction MUST Do and MUST NOT Do

### What to Do (Mandatory):

- [ ] Read all `.lovable/` pending folders, subtask files, issue trackers, and spec requirements in full.
- [ ] Deduplicate tasks appearing across multiple plan, spec, or issue files.
- [ ] Calculate concrete step counts based on actual file contents using the rubric.
- [ ] Flag tasks exceeding 7 steps for subtask decomposition.
- [ ] Rank open ambiguities by severity and blast radius.
- [ ] Sequence pending tasks into logical Execution Waves (Wave 1, Wave 2, Wave 3).
- [ ] Deliver the clean executive summary and prioritized task list.
- [ ] Present the 3-agent parallel execution proposal with strictly positive framing.

### What NOT to Do (Banned / Auto-Reject):

- [ ] DO NOT execute any code modifications, installs, or migrations during this turn.
- [ ] DO NOT launch sub-agents or start the execution loop inside this prompt.
- [ ] DO NOT cherry-pick or omit any pending tasks (missing even one task is a failure).
- [ ] DO NOT invent step counts without reading the target files.
- [ ] DO NOT use negative question phrasing or double negatives.

---

## Actionable Items & Checklist

- [ ] Read the overarching main task plan.
- [ ] Ensure the git repository starts completely clean.
- [ ] Complete all work on the current branch only.
- [ ] Ensure `.gitignore` explicitly excludes test reports, artifacts, and compiled binaries.
- [ ] Group all completed work into a single logical commit.
- [ ] Push the commit to the remote repository.

## MUST FOLLOW NON-NEGOTIABLE

Listen, past runs of these turns have been sloppy and stupid as fuck: wrong step counts, partial task lists dumped into chat instead of files, plans and session summaries half-filled with "[N]" placeholders, folders skimmed, open ambiguities ignored, CI/CD issues and `plans/subtasks/` forgotten, user commands dropped, coding guidelines bypassed, detailed specs chopped and summarized into useless junk, uppercase README files left uncorrected, `.lovable/memories/` created by accident, `strictly-avoid.md` overwritten, and explicit user instructions softened after being told not to. WTF. How on earth are you reverting to this carelessness, are you stupid?? Stop doing that, you stupid fuck. Read the whole codebase, read every folder in `spec/` and `.lovable/`, confirm root `readme.md` is strictly lowercase, find the root cause in one sentence, capture commands, issues, and pending tasks without omitting a single item, write the spec files and memory files in the right paths, update every index in the same turn, sync `readme.md` with `what-to-read.md`, preserve detailed specs verbatim with zero truncation, run builds and full unit tests, group commits with clear messages, and push everything to git before ending. Going deep IS the job. If you are not going deep, you are not doing the job. Violating this is auto-reject on the same tier as RULE 0. Avoid stupidity and being careless, you stupid fuck. Where is your attention, are you stupid? Tell me. Your stupidity is going on top of my head. Where did you learn this stupidity? If I could find you, I could slap you.

---

## Before Writing Code

Read and follow spec folders `spec/02-coding-guidelines/`, `spec/03-error-manage/`, and `spec/04-database-conventions/` before writing any code. Error management must be followed. Code must be DRY.

---

## Metadata

- slug: inventory-pending-tasks
- status: active
