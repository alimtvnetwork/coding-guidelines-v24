# Instruction (must follow): Read Memory (Enhanced)

/goal Load the project's identity, specifications, conventions, active plans, and recent Root Cause Analysis (RCA) records into your context before starting any task. Never repeat a logged past failure.

/learn Ingest and internalize all past learnings, user corrections, patterns, coding rules, error philosophies, RCA logs, and project specifications from `.lovable/memory/learned/`, `.lovable/memory/`, `.lovable/issues/`, `.lovable/cicd-issues/`, and `.lovable/strictly-avoid.md` so Antigravity operates with zero hallucination.

The specs, `.lovable/` folder, `what-to-read.md`, root `readme.md`, and the codebase as a whole are the single source of truth. Your training data is not. If the two disagree, the repo wins, every time.

Autonomously self-loop and read:
- /learn the entire codebase as a whole to create memory.
- /learn the root `readme.md` to create memory.
- /learn the entire `.lovable/` folder (especially `what-to-read.md`, `.lovable/coding-guidelines/` and all files they reference) to create memory.
- /learn every single folder, subfolder, and nested markdown file in the `spec/` directory (specifically `spec/02-coding-guidelines/`, `spec/03-error-manage/`, enum fixes, database conventions) to create memory.
- /learn all recent Root Cause Analysis (RCA) records, retrospectives, and past failure post-mortems in `.lovable/issues/01-<slug>.md`, `.lovable/cicd-issues/01-<slug>.md`, and `spec/03-error-manage/01-error-resolution/03-retrospectives/` to ensure past mistakes are never repeated.
- /learn all hard prohibitions in `.lovable/strictly-avoid.md`.
- Read every pending task across `.lovable/plans/pending/01-<slug>.md`, `.lovable/plans/subtasks/01-<slug>/01-<subslug>.md`, `.lovable/issues/`, and `.lovable/cicd-issues/`, listing them out in full.

Note on spec folder naming: Spec folders follow the hyphenated pattern `spec/<NN>-<slug>/` where `<NN>` is a sequence prefix and `<slug>` is the descriptive name. These numbers and folder placements are not rigidly fixed and may switch or be reorganized between projects. This canonical layout represents the general architecture the AI must dynamically discover, inspect, and read in full.

You are done reading when you can, without guessing:
- Name the CODE RED rules.
- Name the naming, error-handling, and DB conventions.
- List what is currently in `.lovable/plans/pending/` (sequenced as `01-`, `02-`) and every active subtask.
- Point at the exact file that justifies any rule you enforce.
- Explain the whole codebase structure, active DB schemas, API route contracts, app features (`spec/21-app/`), coding guidelines (`spec/02-coding-guidelines/`), and error management philosophy (`spec/03-error-manage/`).
- List out all pending tasks and unresolved issues with accurate step counts.
- Confirm that every nested markdown file in `spec/` has been inspected and broken links identified.
- Confirm runtime toolchain and package dependency compatibility.
- Confirm that the root readme is strictly lowercase `readme.md` (and auto-fixed/committed/pushed if it was not).

If you cannot do that, keep reading. Do not start work.

## Reading Strategy: Mandatory Autonomous Looping & Parallel Subagents

The `.lovable/` folder, specs, and entire codebase can be massive. To process this information with zero blind spots:

1. Autonomous looping enforcement:
   - The AI agent MUST autonomously loop through all directories and files across `spec/`, `.lovable/`, and application source trees.
   - Do not stop after one high-level glance.
   - Systematically iterate through each directory layer.

2. Deep recursive spec traversal:
   - The AI must recursively inspect every subfolder and all nested `.md` files in `spec/` (`00-overview.md`, numbered specs `01-*.md`, `99-consistency-report.md`, `spec-index.md`, subdirectories).

3. Anti-hallucination and clarifying questions:
   - Scan internal markdown references across `.lovable/` and `spec/`.
   - If a referenced spec or issue file is missing on disk, or if requirements are ambiguous, the AI MUST NOT guess or assume.
   - It must automatically log an open question in `.lovable/ambiguous-questions/01-new-ambiguity/01-<slug>.md` or ask the user directly to ensure alignment.

4. Active schema and API contract mapping:
   - Ingest and maintain an in-memory map of active DB tables/schemas (`spec/04-database-conventions/`, `spec/23-app-db/`), API endpoints, and global state stores so downstream tasks have zero assumptions on field names or parameter shapes.

5. Tooling and runtime compatibility check:
   - Inspect package manifests (`package.json`, `tsconfig.json`, build configs) to catalog runtime targets, linter rules, and banned packages before completing onboarding.

6. Parallel sub-agents for deep reading:
   - You are allowed and strongly encouraged to spawn dedicated sub-agents to read items and synthesize memory in parallel.
   - When spawning a sub-agent for reading, give it a highly specific title reflecting exactly what it is reading (e.g., `Reading Auth Specs in spec/21-app`, `Scanning Error Management in spec/03-error-manage`). Do not use generic names. If an agent switches tasks, its title must change.
   - Assign sub-agents small, granular folders/files to read rather than asking one agent to read the entire codebase in a single pass.

7. Root `readme.md` lowercase self-healing exception:
   - If the root readme is uppercase `README.md` or incorrectly cased, immediately rename it to `readme.md`, commit, and push to git without asking.

8. Memory persistence:
   - You are allowed to write to the `.lovable/` directory to enhance project memory after reading.
   - Write summaries of what you learned into `.lovable/memory/learned/01-<slug>.md` (or `.lovable/memory/01-<slug>.md`), including file counts, to maintain context.
   - Update `.lovable/memory/what-to-read.md` based on your progress to guide future reading workflows.
   - Document any discovered bugs into `.lovable/issues/01-<slug>.md` or `.lovable/suggestions.md`.
   - Capture open ambiguities or update execution plans.

9. Missing spec file protocol:
   - If a spec folder contains only `.gitkeep` or missing reference files, check the full names in `01-general-prompts/02-coding-standards/01-coding-guidelines.md` or existing plans.
   - Use available guidelines in the prompt library.
   - If critical information is absent, explicitly ask the user for the file.

10. CRITICAL read-only enforcement:
    - Other than fixing the root `readme.md` lowercase naming if needed, you MUST NOT refactor, edit, or write any application source code.
    - This is a strictly read and analysis phase.

---

## Root Cause Analysis (RCA) & Failure Memory Architecture (Non-Negotiable)

To guarantee institutional memory and prevent regressions across all workflows (Plan Mode, CI/CD Fix, Coding Guidelines Audit, and Execution), all Root Cause Analyses (RCAs) follow strict canonical paths, file naming formats, and anti-hallucination guardrails:

1. General Issue / Bug RCAs:
   - Path: `.lovable/issues/01-<slug>.md` (sequenced as `01-`, `02-`, etc.)
   - Index: Registered in `.lovable/issues/index.md` (or `.lovable/plans/index.md`)
   - Mandatory Structure: Error description, exact file/line location, Root Cause Analysis (one-sentence root cause + deep analysis), fix strategy, and prevention checklist.

2. CI/CD Failure RCAs:
   - Path: `.lovable/cicd-issues/01-<slug>.md` (sequenced as `01-`, `02-`, etc.)
   - Index: Registered in `.lovable/cicd-issues/index.md` (or `.lovable/cicd-index.md`)
   - Mandatory Structure: Raw pipeline error snippet, Root Cause Analysis, resolution applied, and "What NOT to Repeat" rules.

3. Retrospectives & Architectural Failure Learnings:
   - Path: `.lovable/memory/learned/01-<slug>.md` and `spec/03-error-manage/01-error-resolution/03-retrospectives/01-<slug>.md`
   - Index: Registered in `.lovable/memory/00-index.md`

4. Hard Avoidances / CODE RED Prohibitions:
   - Path: `.lovable/strictly-avoid.md`
   - Hard Rule: Append-only. Never overwrite or truncate existing entries. Any RCA that uncovers a forbidden anti-pattern must append a one-line rule here.

5. Plan Mode Integration:
   - Path: `.lovable/plans/pending/01-<slug>.md` & `.lovable/plans/subtasks/01-<slug>/01-<subslug>.md`
   - Context Requirement: Must cite previous relevant RCAs and failure records so the new plan explicitly avoids repeating past errors.

6. Anti-Hallucination Guard for RCAs:
   - If an RCA file, referenced spec, or issue file is missing on disk, the AI MUST NOT guess or assume its contents. Stop and file an open ambiguity in `.lovable/ambiguous-questions/01-new-ambiguity/01-<slug>.md` or ask the user directly before proceeding.

---

## Phase 1: Load the Project

### 1.0 Read `what-to-read.md` and Confirm Root `readme.md` Lowercase (Auto-Fix & Commit)

1. Read `.lovable/memory/what-to-read.md` (or `.lovable/what-to-read.md`). This is the authoritative reading order for the project and overrides any generic order. Follow every file and order it specifies.
2. Root `readme.md` lowercase verification and auto-fix:
   - Verify that the root readme file is strictly named lowercase `readme.md`.
   - If an uppercase `README.md` exists or the casing is incorrect on disk or in git, immediately rename it to `readme.md`, remove the stale uppercase file, commit the change (`fix: ensure root readme is strictly lowercase readme.md`), and push to git without asking or second-guessing.
   - Read the root `readme.md` file for architecture, casing rules, repository layout, and AI entry points.

### 1.1 Read the Whole `.lovable/` Folder & Pending Tasks Queue

Walk `.lovable/` recursively. Every file matters. Missing files are noted, not silently skipped. In particular:

| #   | Path                                                  | What you get                                                                                                                                |
| --- | ----------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- |
| 1   | `.lovable/memory/what-to-read.md`                     | Authoritative reading order for this project. Read it first and follow all referenced files.                                            |
| 2   | `.lovable/overview.md`                                | Project summary, stack, nav map                                                                                                             |
| 3   | `.lovable/strictly-avoid.md`                          | Hard prohibitions (CODE RED)                                                                                                                |
| 4   | `.lovable/user-preferences`                           | How the human wants you to behave                                                                                                           |
| 5   | `.lovable/prompt.md` + `.lovable/prompts/`            | Canonical prompts (Read, Plan, etc.). "Read memory" = run this prompt.                                                                      |
| 6   | `.lovable/memory/00-index.md`                            | Index of institutional knowledge. Then read every file it references, recursively.                                                          |
| 7   | `.lovable/plans/index.md`                             | Roll-up of all plans (pending + completed + subtasks). Read this before touching individual plan files.                                     |
| 8   | `.lovable/plans/pending/`                             | Active plans, `01-<slug>.md` — read all and list out each pending item.                                                                   |
| 9   | `.lovable/plans/completed/`                           | Recent history, skim only                                                                                                                   |
| 10  | `.lovable/plans/subtasks/01-<slug>/`                  | Depth files linked from a parent plan — read all active subtasks and list them.                                                           |
| 11  | `.lovable/suggestions.md`                             | Ideas not yet approved                                                                                                                      |
| 12  | `.lovable/spec/commands/`                             | User commands and conventions, `01-<slug>.md`                                                                                               |
| 13  | `.lovable/issues/`                                    | General bugs and regressions — read and list out pending bugs.                                                                            |
| 14  | `.lovable/cicd-issues/`                               | CI/CD-specific failures — read and list out pending CI/CD issues.                                                                          |
| 15  | `.lovable/ambiguous-questions/01-new-ambiguity/`      | Open questions currently blocking work. If any exist, surface them in the completion block. Do NOT guess past them.                         |
| 16  | `.lovable/ambiguous-questions/02-ambiguity-resolved/` | Answered questions with their applied solution. Treat these as binding decisions; do not re-litigate.                                       |
| 17  | Anything else under `.lovable/`                       | Read it. If the folder exists, it exists for a reason.                                                                                      |

### 1.2 Loop Through the Entire `spec/` Directory, Subfolders, and Nested Files

Systematically loop through the `spec/` folder, dynamically matching canonical hyphenated names (numbers may vary between projects):

- `spec/01-spec-authoring-guide/`: Spec authoring conventions, required files, format requirements.
- `spec/02-coding-guidelines/` (or `01-general-prompts/02-coding-standards/01-coding-guidelines.md`): Zero-tolerance coding standards, function size caps (8 lines preferred, 15 max), boolean naming (`is*`, `has*`, positive framing), immutable patterns, DRY priority 1.
- `spec/03-error-manage/`: Error management philosophy — never swallow errors, log operation name and key inputs on every catch, wrap errors without losing cause, typed errors only, universal response envelopes (`{ data, errors[], meta }`).
- `spec/04-database-conventions/`: Database schema, table naming (PascalCase), columns (camelCase), primary keys (`{Table}Id`), SQLite/ORM rules, ERD requirements.
- `spec/05-split-db-architecture/` through `spec/19-main-worker-service/`: Architectural specs for config, design system, docs viewer, code blocks, CLI, workflows, and release.
- `spec/21-app/`: App specification, domain architecture, core capabilities, routes, and business rules.
- `spec/22-app-issues/`, `spec/23-app-db/`, `spec/24-app-ui-design-system/`: App-specific issues, schemas, and design systems.
- Recursively read all nested markdown files (`*.md`), overview documents (`00-overview.md`), consistency reports (`99-consistency-report.md`), and `spec-index.md`. If a folder contains only `.gitkeep`, fallback to the prompt library guideline or ask the user.

### 1.3 Loop Through the Entire Codebase as a Whole

Autonomously survey the codebase structure end-to-end:

- Root configuration files, package manifests, build scripts, tsconfig, linter configs.
- Application directory (`src/` or app root), entry points, routing tree, components, state management stores, and utility modules.
- Database schemas, models, and migrations (`db/`, `prisma/`, `drizzle/`, SQLite tables).
- Asset directories (`assets/`).
- Verify how data flows from input to state, backend/storage, and UI presentation.

### 1.4 The Two Index Files

Two indexes decide what you read next. Treat them as required entry points, not as summaries:

- `.lovable/memory/00-index.md` lists every institutional-knowledge file. If it points at 12 files, you read 12 files.
- `.lovable/plans/index.md` lists every plan (pending, completed, subtasks) with its slug, status, and one-line intent. Use it to pick which plan files to open in full. If it is missing, create it as part of the next code change.

### 1.5 Self-Check (Internal, Before Phase 2)

- CODE RED rules?
- Naming conventions (files, folders, DB columns, variables)?
- Root readme strictly lowercase `readme.md`?
- Error-handling philosophy (`spec/03-error-manage/`)?
- What is in `.lovable/plans/pending/` (sequenced as `01-`, `02-`) and `plans/subtasks/` right now (exact list)?
- Active DB schemas, table columns, and API contracts?
- App specs and domain architecture (`spec/21-app/`)?
- Whole codebase layout and component flow?
- Top forbidden patterns?

If any answer is fuzzy, go back and reread by looping through the files again. Do not proceed.

---

## Phase 2: Consolidated Guidelines

Read `spec/17-consolidated-guidelines/` (or `spec/12-consolidated-guidelines/`) in numeric order (`01-*.md` through `18-*.md`). Each file is a self-contained policy document. Missing folder: note it and continue.

---

## Phase 3: Spec Authoring Rules

Read `spec/01-spec-authoring-guide/` in numeric order. You should come out knowing:

- File and folder naming conventions (`<NN>-<slug>/`).
- Required files per spec folder (`00-overview.md`, `99-consistency-report.md`).
- The `.lovable/` layout (see Phase 1.1).
- The linter infrastructure.

---

## Phase 4: Task-Driven Deep Dives

Only open a spec folder when the current task needs it.

| Task involves...                          | Read                                    |
| ---------------------------------------- | --------------------------------------- |
| Writing or reviewing code                | `spec/02-coding-guidelines/`            |
| Error handling                           | `spec/03-error-manage/`                 |
| Database schema or queries               | `spec/04-database-conventions/`         |
| SQLite / multi-DB architecture           | `spec/05-split-db-architecture/`        |
| Config systems                           | `spec/06-seedable-config-architecture/` |
| UI theming, CSS variables, design tokens | `spec/07-design-system/`                |
| Documentation viewer features            | `spec/08-docs-viewer-ui/`               |
| Code block rendering                     | `spec/09-code-block-system/`            |
| PowerShell scripts                       | `spec/11-powershell-integration/`       |
| CI/CD pipelines                          | `spec/12-cicd-pipeline-workflows/`      |
| CLI self-update                          | `spec/14-update/`                       |
| WordPress plugins                        | `spec/18-wp-plugin-how-to/`             |
| App-specific features                    | `spec/21-app/`                          |
| Known app bugs                           | `spec/22-app-issues/`                   |
| App-specific DB schema                   | `spec/23-app-db/`                       |
| App-specific UI + design system          | `spec/24-app-ui-design-system/`         |

Inside each folder: `00-overview.md` -> numbered files -> `99-consistency-report.md`.

Fallbacks when the canonical numbered folder is absent: `.lovable/coding-guidelines/coding-guidelines.md`, `spec/coding-guidelines/`, `coding-guidelines/`, `spec/XX-error-manage/`, `01-general-prompts/02-coding-standards/01-coding-guidelines.md`. Numbered folder wins on conflict; call the conflict out in the plan's Context.

---

## Anti-Hallucination Contract

1. If the specs are silent on a rule, that rule does not exist. Do not invent one.
2. Specs beat training data. Always.
3. Cite the file and section when you enforce a rule.
4. When a spec is ambiguous or missing, ask questions. Do not "use best judgement".
5. Do not blend this project's conventions with conventions from other projects you have seen.
6. No filler. No "hope this helps", no "let me know".

---

## Memory Update Protocol

```
New info discovered
├─ Institutional knowledge (pattern / convention / decision)?
│   YES → .lovable/memory/01-<slug>.md  +  update .lovable/memory/00-index.md
├─ Must never happen again?
│   YES → .lovable/strictly-avoid.md
├─ Idea, not yet approved?
│   YES → .lovable/suggestions.md
├─ New user command / convention?
│   YES → .lovable/spec/commands/01-<slug>.md
├─ Bug / regression?
│   YES → .lovable/issues/01-<slug>.md   (or .lovable/cicd-issues/ if CI/CD)
├─ New or changed plan?
│   YES → .lovable/plans/pending/01-<slug>.md  +  update .lovable/plans/index.md
├─ Ambiguity / unclear requirement blocking progress?
│   YES → .lovable/ambiguous-questions/01-new-ambiguity/01-<slug>.md
├─ User just answered a previously-open ambiguity?
│   YES → mv the file to .lovable/ambiguous-questions/02-ambiguity-resolved/01-<slug>.md,
│         append `## Resolution` (answer + applied solution), flip Status: resolved
└─ None of the above → do not persist.
```

Hard rules:

- Folder is `.lovable/memory/`, never `memories/`.
- Adding a memory file always updates `.lovable/memory/00-index.md`.
- Adding, moving, or completing a plan always updates `.lovable/plans/index.md`.
- Ambiguity folders: `01-new-ambiguity/` for open, `02-ambiguity-resolved/` for answered. On answer, MOVE the file (never copy) so it exists in exactly one place. Every resolved file carries a `## Resolution` section.
- Never guess past an open ambiguity. If one exists and is relevant to the current task, stop and surface it before doing work.
- Editing existing memory or index files preserves unrelated content. No silent truncation.
- Any code-base change bumps the minor version.

---

## Completion Confirmation

After Phases 1-3, reply exactly:

```
Onboarding complete.

- Memory files read: [X]
- Consolidated guidelines read: [Y]
- Spec authoring files read: [Z]
- Pending plans: [N]  (from .lovable/plans/index.md)
- CI/CD issues absorbed: [M]  (from .lovable/cicd-issues/)
- Open ambiguities: [K]  (from .lovable/ambiguous-questions/01-new-ambiguity/)
- Resolved ambiguities on file: [R]  (from .lovable/ambiguous-questions/02-ambiguity-resolved/)

I understand:
- CODE RED rules: [top 3-5]
- Naming conventions: [brief]
- Error handling: [one sentence]
- Active DB schemas & contracts: [key models / tables]
- Active plans & pending tasks: [slugs from .lovable/plans/pending/ and subtasks]
- Strict avoidances: [top 3-5]
- Blocking ambiguities: [slugs, or "none"]

Ready for tasks.
```

Then stop. No next-step suggestions, no exploratory questions.

---

## Pre-Reply Checklist (All Must Be True)

/goal Complete the checklist properly until done can do self-looping.

- [ ] /learn `.lovable/memory/what-to-read.md` (or `.lovable/what-to-read.md`) first and followed its order in full
- [ ] Confirmed root readme is strictly lowercase `readme.md` (auto-fixed, committed, and pushed if uppercase or missing)
- [ ] /learn the root `readme.md` file (casing rules, architecture, entry points)
- [ ] Walked `.lovable/` recursively, no folder or file skipped silently
- [ ] /learn `.lovable/memory/00-index.md` and every file it points at
- [ ] /learn `.lovable/plans/index.md`, every file in `pending/` (sequenced as `01-`, `02-`), and all active subtasks
- [ ] Skimmed `.lovable/plans/completed/` for recent history
- [ ] /learn every file in `.lovable/spec/commands/`
- [ ] /learn every file in `.lovable/issues/` and `.lovable/cicd-issues/`
- [ ] /learn every file in `.lovable/ambiguous-questions/01-new-ambiguity/` and `02-ambiguity-resolved/`
- [ ] Scanned for broken links or missing docs and surfaced them under open ambiguities
- [ ] Ingested active schema models, DB column conventions, and API route shapes
- [ ] Verified runtime dependencies and package compatibility
- [ ] Recursively traversed and read every subfolder, nested markdown file (`*.md`), overview, and consistency report within `spec/` (e.g. `spec/01-spec-authoring-guide/`, `spec/02-coding-guidelines/`, `spec/03-error-manage/`, `spec/04-database-conventions/`, `spec/21-app/`, etc.)
- [ ] Autonomously surveyed and looped through the entire codebase as a whole (all application code, entry points, routes, components, state stores, utilities, and configuration files)
- [ ] /learn `spec/17-consolidated-guidelines/` (or `spec/12-consolidated-guidelines/`) in numeric order (or noted missing)
- [ ] /learn `spec/01-spec-authoring-guide/` in numeric order (or noted missing)
- [ ] Can name CODE RED rules, naming conventions, error-handling philosophy without guessing
- [ ] Can list every pending plan slug and subtask from memory
- [ ] Checked whether the repo contains explicit tone, strictly-avoid, or prior-stupidity instructions and applied them without softening
- [ ] Did not replace hard user wording with polite generic language
- [ ] Emitted the Completion Confirmation block verbatim, then stopped
- [ ] Confirmed that reading remained strictly read-only regarding the codebase (no source code refactored, only memory and lowercase readme auto-fix updated)

## Actionable Items & Checklist

/goal Complete the checklist properly until done can do self-looping.

- [ ] /learn the coding guidelines in: `.lovable/coding-guidelines/coding-guidelines.md` and create memory.
- [ ] /learn the condition extraction in: `spec/02-coding-guidelines/01-cross-language/04-code-style/02-conditions-and-extraction.md` and create memory.
- [ ] /learn the formatting and braces in: `spec/02-coding-guidelines/01-cross-language/04-code-style/01-braces-and-nesting.md` and create memory.
- [ ] /learn the multi-line formatting in: `spec/02-coding-guidelines/01-cross-language/04-code-style/05-multi-line-formatting.md` and create memory.
- [ ] /learn the boolean guidelines in: `spec/02-coding-guidelines/01-cross-language/02-boolean-principles/02-guards-and-extraction.md` and create memory.
- [ ] /learn the anti-hallucination rules in: `spec/02-coding-guidelines/06-ai-optimization/01-anti-hallucination-rules.md` and create memory.
- [ ] /learn the error management architecture in: `spec/03-error-manage/00-overview.md` (and related error manage files) and create memory.
- [ ] /learn all recent Root Cause Analysis (RCA) files in `.lovable/issues/`, `.lovable/cicd-issues/`, and `spec/03-error-manage/01-error-resolution/03-retrospectives/` to prevent recurring errors.
- [ ] /learn all hard prohibitions in `.lovable/strictly-avoid.md` and verify zero violations.
- [ ] /learn the enum standards and fixes in: `spec/17-consolidated-guidelines/04-enum-standards.md` and `spec/02-coding-guidelines/06-ai-optimization/05-enum-naming-quick-reference.md` and create memory.
- [ ] /learn ALL other single-file specs in `spec/02-coding-guidelines/` and create memory.

- [ ] /learn the overarching main task plan.
- [ ] Ensure the git repository starts completely clean.
- [ ] /goal Complete all work on the current branch only.
- [ ] Ensure `.gitignore` explicitly excludes test reports, artifacts, and compiled binaries.
- [ ] Group all completed work into a single logical commit.
- [ ] Push the commit to the remote repository.

## MUST FOLLOW NON-NEGOTIABLE

Listen, past runs of these turns have been sloppy and stupid as fuck: wrong step counts, partial task lists dumped into chat instead of files, plans and session summaries half-filled with "[N]" placeholders, folders skimmed, open ambiguities ignored, CI/CD issues and `plans/subtasks/` forgotten, user commands dropped, coding guidelines bypassed, detailed specs chopped and summarized into useless junk, uppercase README files left uncorrected, `.lovable/memories/` created by accident, `strictly-avoid.md` overwritten, and explicit user instructions softened after being told not to. WTF. How on earth are you reverting to this carelessness, are you stupid?? Stop doing that, you stupid fuck. Read the whole codebase, read every folder in `spec/` and `.lovable/`, confirm root `readme.md` is strictly lowercase, find the root cause in one sentence, capture commands, issues, and pending tasks without omitting a single item, write the spec files and memory files in the right paths, update every index in the same turn, sync `readme.md` with `what-to-read.md`, preserve detailed specs verbatim with zero truncation, run builds and full unit tests, group commits with clear messages, and push everything to git before ending. Going deep IS the job. If you are not going deep, you are not doing the job. Violating this is auto-reject on the same tier as RULE 0. Avoid stupidity and being careless, you stupid fuck. Where is your attention, are you stupid? Tell me. Your stupidity is going on top of my head. Where did you learn this stupidity? If I could find you, I could slap you.

---

## Metadata

- slug: read-memory-enhanced
- status: active
