# Specification: Execute Parent Task & Prompt Fixes (Understood Gate & Azure AI Runtime Integration)

> **Target:** `02-spec/21-app/execute-parent-task-and-prompt-fixes/01-architecture-spec.md`  
> **Status:** Authoritative Architecture Spec  
> **Version:** 1.0.0  
> **Date:** 2026-10-19  

---

## 1. Executive Summary & Problem Statement

### 1.1 Context
In autonomous multi-agent execution workflows running under the V6 execution protocol (`01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`) and the Coding Guidelines Execution Suite (`01-prompts/15-cg-execute/*.md`), AI agents must exhibit complete cognitive transparency and deterministic environment awareness during their initial Turn-1 interaction.

Two critical omissions were identified in active operational runs:
1. **Missing Understood / AI Comprehension Field:** While Phase 1A narrative instructions mandated an "Extreme Understood Turn-1 Showcase", the markdown task template under `### 📋 Confirmed Task Breakdown & Requirement Ingestion` lacked an explicit, structured field for confirming AI comprehension per task. Consequently, agents following the template literally omitted the comprehension declaration.
2. **Missing Azure AI / Model Runtime Context Field:** Modern engineering swarms operate across diverse model runtimes (Azure AI Foundry, Azure OpenAI endpoints, Google Antigravity, Cursor IDE, local GitMap CLI workers). The task breakdown failed to specify the active model runtime context, leading to loss of runtime provenance and execution environment ambiguity during agent handoffs.

### 1.2 Objective
This specification formalizes the remediation across the entire prompt ecosystem:
1. Establish a rigorous **6-Part Task Verification Schema** (with 7 mandatory keys) for all task breakdown manifests.
2. Update the narrative instructions and markdown templates in `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` and its `v4` mirror.
3. Propagate the updated schema across all 38 prompts in `01-prompts/15-cg-execute/` and `01-prompts/v4/15-cg-execute/`.
4. Update companion skills in `.agents/skills/` and `.cursor/skills/`.
5. Maintain strictly relative Git paths and lowercase file naming conventions with zero exceptions.

---

## 2. Root Cause Analysis (RCA)

### 2.1 Root Cause 1: Narrative vs. Template Divergence in Phase 1A
- **Observation:** In Phase 1A Step 2 of `02-execute-parent-task-with-n-steps-v6.md`, the text instructed:
  > *"Extreme 'Understood' Turn-1 Showcase (Turn 1 Action): In your VERY FIRST response turn upon receiving the prompt, you MUST output a comprehensive, in-depth understanding block directly in visible chat before any background tool execution."*
- **The Defect:** Immediately following this narrative, the reference markdown template for task items defined:
  ```markdown
  1. **Task-01: [Descriptive Task Title]**
     - **State:** `[IN PROGRESS — EXECUTING IMMEDIATELY]`
     - **Intent & Technical Deliverables:**
       ...
     - **Negative Boundaries & Invariants:**
       ...
     - **Disjoint File Boundaries:**
       ...
     - **Verification Evidence Contract:**
       ...
  ```
- **Analysis:** Large language models rely heavily on concrete few-shot template formatting. Because the template omitted an explicit `Understood / AI Comprehension` key, agents copied the structural keys verbatim and omitted the comprehension declaration at the task level.

### 2.2 Root Cause 2: Omission of Cloud & Model Runtime Provenance
- **Observation:** Tasks executed via cloud runners (such as Azure AI Foundry, Azure OpenAI endpoints, or Antigravity agents) provided no structural metadata indicating which runtime environment or model family was executing the subtask.
- **The Defect:** Without an explicit runtime field, multi-agent swarms could not distinguish between local GitMap CLI execution, cloud agent sandboxes, or Azure AI proxy endpoints. This created audit blind spots when diagnosing runtime-specific timeouts, rate limits, or context budget exhaustion.

---

## 3. The 6-Part Task Verification Schema

Every task breakdown emitted during Phase 1A (and persisted in `.ai-memory/plans/` and subtask manifests) MUST strictly adhere to the 6-Part Task Verification Schema:

```markdown
1. **Task-01: [Descriptive Task Title]**
   - **State:** `[IN PROGRESS — EXECUTING IMMEDIATELY]`
   - **Understood / AI Comprehension:** [YES] — [1–2 concise sentences proving deep understanding of intent, scope, and verified constraints; answering did the AI understand the requirements or not]
   - **Azure AI / Model Runtime Context:** [Explicitly document Azure AI, Azure OpenAI, or local model endpoint context if applicable, or state 'N/A — Local CLI/GitMap']
   - **Intent & Technical Deliverables:**
     - [Clear statement of technical goal and deliverable artifacts]
     - [Precise implementation scope: functions, types, schemas, or prompt sections being added/edited]
     - [Quantifiable success metrics: line count, field completeness, schema compliance]
   - **Negative Boundaries & Invariants:**
     - [Explicit list of things this task MUST NOT do or alter]
     - [Untouchable files, reserved packages, or excluded dependencies]
     - [Invariants that must remain preserved after modification]
   - **Disjoint File Boundaries:**
     - `Owned Files (Exclusive):` `[strictly/relative/path/to/file-01.md]`, `[strictly/relative/path/to/file-02.md]`
     - `Read-Only Context Files:` `[strictly/relative/path/to/reference.md]`
   - **Verification Evidence Contract:**
     - [Exact command or check to be executed to prove completion]
     - [Expected exit code: e.g. exit 0]
     - [Required diffstat evidence or structural regex check]
```

### 3.1 Detailed Schema Field Invariants

| Key | Mandatory Format | Invariant Description |
| :--- | :--- | :--- |
| **State** | `[IN PROGRESS ...]` / `[QUEUED ...]` / `[DONE]` | Tracks lifecycle state. Exactly one active task is `IN PROGRESS`; all remaining tasks are `QUEUED`. |
| **Understood / AI Comprehension** | `[YES] — <justification>` | Binary assertion `[YES]` followed by 1–2 sentences summarizing architectural understanding and constraints. `[NO]` triggers immediate clarification. |
| **Azure AI / Model Runtime Context** | `[Environment & Model Context]` | Explicit declaration of runtime environment (e.g., Azure AI Foundry GPT-4o, Azure OpenAI, Antigravity Claude-3.5-Sonnet, Cursor IDE, or `N/A — Local CLI/GitMap`). |
| **Intent & Technical Deliverables** | 3 bullet points | Quantifiable implementation goals, modified structs/types, deliverable filenames. |
| **Negative Boundaries & Invariants** | 3 bullet points | Explicit prohibitions, untouchable files, preserved backwards compatibility rules. |
| **Disjoint File Boundaries** | `Owned Files (Exclusive)` & `Read-Only Context Files` | Strict single-owner guarantee per file to eliminate concurrency write collisions in multi-agent swarms. |
| **Verification Evidence Contract** | Command + Exit Code + Evidence | Grounded verification proof (exit 0, diffstat, linter regex check). Assurances without proof are auto-rejected. |

---

## 4. Narrative & Template Modifications

### 4.1 Phase 1A Narrative Updates (Step 2)
In `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` and `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md`, Section "Phase 1A: Verbatim Capture, Task Extraction & Chat Output Gate (Step 0)" Step 2 is updated:

1. **Explicit Comprehension Requirement:** The text must explicitly mandate the inclusion of `- **Understood / AI Comprehension:** [YES] — ...` in every task item.
2. **Explicit Azure AI Context Requirement:** The text must explicitly mandate `- **Azure AI / Model Runtime Context:** ...` in every task item.
3. **Turn 1 Output Guarantee:** Turn 1 must present both the 5-Pillar Understood Mandate and the 6-Part Task Verification Schema in visible chat while executing initial background tools in the exact same turn.

### 4.2 Markdown Template Updates
The code block under `### 📋 Confirmed Task Breakdown & Requirement Ingestion` must contain the updated template with both `Task-01` and `Task-02` demonstrating the complete 7-key structure.

---

## 5. Scope & Fleet Impact Analysis

### 5.1 Primary Files (Subtask 01)
- `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`
- `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md`
- `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`
- `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`

### 5.2 CG Execute Prompt Suite (Subtask 02)
All 38 prompt files under `01-prompts/15-cg-execute/` and their 38 counterparts in `01-prompts/v4/15-cg-execute/` (76 total prompt files) will receive the upgraded task breakdown schema, ensuring absolute uniformity across:
- `01-cg-execute-clean-repo-build-and-caches.md`
- `02-cg-execute-nested-if-and-guard-clauses.md`
- `03-cg-execute-style-guidelines.md`
- `04-cg-execute-cli-commands-and-help.md`
- `05-cg-execute-multi-language-enums-and-traits.md`
- `06-cg-execute-lazy-regex-and-pattern-matching.md`
- `07-cg-execute-code-hygiene.md`
- `08-cg-execute-function-signatures-and-return-types.md`
- `09-cg-execute-function-argument-reduction-and-params.md`
- `10-cg-execute-python-constants-and-magic-number-elimination.md`
- `11-cg-execute-python-boolean-and-conditional-enhancement.md`
- `12-cg-execute-python-array-constants-and-dynamic-enums.md`
- `13-cg-execute-relative-paths.md`
- `14-cg-execute-result-wrapper-and-apperror-returns.md`
- `15-cg-execute-react-frontend-guidelines.md`
- `16-cg-execute-testing-and-coverage.md`
- `17-cg-execute-typescript-guidelines-and-types.md`
- `18-cg-execute-boolean-improvements-v2.md`
- `19-cg-execute-commit-fix-v2.md`
- `20-cg-execute-nuclear-package-modularization-and-unit-test-optimization.md`
- `21-cg-execute-in-below-steps.md`
- `22-cg-execute-extract-types.md`
- `23-cg-execute-follow-other-prompts.md`
- `24-cg-execute-clean-artifacts-and-git-history.md`
- `25-cg-execute-database-conventions.md`
- `26-cg-execute-enum-standards.md`
- `27-cg-execute-boolean-and-naming.md`
- `28-cg-execute-error-management.md`
- `29-cg-execute-isolate-os-tests.md`
- `30-cg-execute-string-normalization-and-equalfoldany.md`
- `31-cg-execute-branch-immutability-and-clean-construction.md`
- `32-cg-execute-variadic-and-spread-parameters.md`
- `33-cg-execute-map-dispatch-and-handler-lookup.md`
- `34-cg-execute-string-efficiency.md`
- `35-cg-execute-file-size-and-function-reduction.md`
- `36-cg-execute-golang-pointer-reduction-and-value-semantics.md`
- `37-cg-execute-code-dryness-and-library-extraction.md`
- `38-cg-execute-smart-test-runner-and-inventory.md`

### 5.3 Downstream Skill Synchronization (Subtask 03 & 04)
- Skill files mirroring these execution standards will be updated across `.agents/skills/` and `.cursor/skills/`.
- Updates will be propagated across 43 repositories via `03-ai-scripts/38-sync-prompts-skills-scripts.py`.

---

## 6. Architectural Invariants & Negative Guardrails

1. **Strict Relative Git Paths (TOTAL BAN on Absolute Paths):**
   - All links, references, and paths in specifications, subtasks, prompts, and notes must be strictly relative Git paths (e.g. `02-spec/...`, `01-prompts/...`, `.ai-memory/...`).
   - Zero absolute filesystem paths (`C:\...`, `/home/...`) or `file:///` URIs.
2. **Strict Lowercase Naming:**
   - All newly authored or modified files must use lowercase filenames (`01-architecture-spec.md`, `skill.md`).
3. **No Shell Search Tools (GitMap Primacy):**
   - Discovery and inspections must use GitMap (`gitmap find`, `gitmap aum search`, `gitmap lf`, `gitmap cat`). Total ban on `Get-ChildItem`, `dir`, `ls`, `rg`, `ripgrep`, `grep`, `Select-String`.
4. **Git Isolation for Worker Subagents:**
   - Worker subagents never run git commands to prevent `.git/index.lock` collisions.
5. **GitMap Commit Formatting:**
   - All commits must use `gitmap cpf "<module> - <summary>"` with a hyphen separator and no colons inside the message argument.

---

## 7. Acceptance Criteria & Verification Gates

- [ ] **Architecture Spec Authoring:** `02-spec/21-app/execute-parent-task-and-prompt-fixes/01-architecture-spec.md` is complete with RCA, 6-part schema, narrative instructions, and fleet impact analysis.
- [ ] **Subtask Manifest Authoring:** Subtasks `01-subtask-preflight-and-v6-prompts.md` and `02-subtask-cg-execute-prompts-migration.md` are authored under `.ai-memory/plans/subtasks/execute-parent-task-and-prompt-fixes/` with exact schema compliance.
- [ ] **Relative Path Compliance:** 100% relative Git paths; zero absolute paths or `file:///` URIs.
- [ ] **Lowercase Compliance:** 100% lowercase filenames.
- [ ] **Cognitive Heatmap Logging:** All specification authoring actions logged in `gitmap heatmap log` and verified via `gitmap heatmap status`.
