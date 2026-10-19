# Subtask 02: Coding Guidelines Execution Prompts Mass Migration & Schema Upgrade

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/execute-parent-task-and-prompt-fixes/02-subtask-cg-execute-prompts-migration.md`  
> **Parent Plan:** `.ai-memory/plans/execute-parent-task-and-prompt-fixes.md`  
> **Specification:** `02-spec/21-app/execute-parent-task-and-prompt-fixes/01-architecture-spec.md`  
> **State:** `COMPLETED`  
> **Assigned Role:** Subagent Worker 02 (CG Execute Suite Lead)  

---

## 1. Task Definition & Verification Contract

- **Task ID:** `02-subtask-cg-execute-prompts-migration`
- **State:** `[COMPLETED]`
- **Understood / AI Comprehension:** `[YES] — Fully understand the requirement to upgrade all 38 prompts in 01-prompts/15-cg-execute/ and their 38 mirrors in 01-prompts/v4/15-cg-execute/ (76 total prompt files) to embed the 6-part task verification schema (Understood/AI Comprehension and Azure AI/Model Runtime Context) in their Phase 1A task breakdown templates.`
- **Azure AI / Model Runtime Context:** `Azure AI / Google Antigravity Agent Swarm (Subagent Worker 02) via GitMap CLI Engine`
- **Intent & Technical Deliverables:**
  - Ingest the full list of 38 prompts under `01-prompts/15-cg-execute/` and their corresponding files in `01-prompts/v4/15-cg-execute/`.
  - Surgically upgrade each prompt's Phase 1A task breakdown template under `### 📋 Confirmed Task Breakdown & Requirement Ingestion` to include:
    - `- **Understood / AI Comprehension:** [YES] — [1–2 concise sentences proving deep understanding of intent, scope, and verified constraints; answering did the AI understand the requirements or not]`
    - `- **Azure AI / Model Runtime Context:** [Explicitly document Azure AI, Azure OpenAI, or local model endpoint context if applicable, or state 'N/A — Local CLI/GitMap']`
  - Ensure both `Task-01` and `Task-02` template instances in each prompt adhere to the 6-part schema.
  - Deliverable scope: Exactly 76 prompt files upgraded with 100% schema consistency.
- **Negative Boundaries & Invariants:**
  - MUST NOT alter the prompt-specific coding guideline rules or domain instructions in sections 2 through 12 of each prompt.
  - MUST NOT touch files owned by Subtask 01 (`01-prompts/14-execute/` and skills).
  - MUST NOT execute raw shell search tools (`rg`, `ripgrep`, `grep`, `Select-String`, `dir`, `ls`, `Get-ChildItem`).
  - Worker subagents MUST NOT execute git commands.
- **Disjoint File Boundaries:**
  - `Owned Files (Exclusive):`
    - `01-prompts/15-cg-execute/*.md` (38 prompts: `01-cg-execute-clean-repo-build-and-caches.md` through `38-cg-execute-smart-test-runner-and-inventory.md`)
    - `01-prompts/v4/15-cg-execute/*.md` (38 mirrored prompts)
  - `Read-Only Context Files:**`
    - `02-spec/21-app/execute-parent-task-and-prompt-fixes/01-architecture-spec.md`
    - `.ai-memory/plans/execute-parent-task-and-prompt-fixes.md`
    - `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`
- **Verification Evidence Contract:**
  - Python verification script or GitMap streaming search checking all 76 files for `Understood / AI Comprehension` key.
  - Python verification script or GitMap streaming search checking all 76 files for `Azure AI / Model Runtime Context` key.
  - Relative path check ensuring 0 absolute paths or `file:///` URIs across all 76 modified files.

---

## 2. Inventory of Target Prompts (38 Prompts / 76 Files)

| Prompt ID | Base Filename | Target Paths |
| :--- | :--- | :--- |
| 01 | `01-cg-execute-clean-repo-build-and-caches.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 02 | `02-cg-execute-nested-if-and-guard-clauses.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 03 | `03-cg-execute-style-guidelines.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 04 | `04-cg-execute-cli-commands-and-help.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 05 | `05-cg-execute-multi-language-enums-and-traits.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 06 | `06-cg-execute-lazy-regex-and-pattern-matching.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 07 | `07-cg-execute-code-hygiene.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 08 | `08-cg-execute-function-signatures-and-return-types.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 09 | `09-cg-execute-function-argument-reduction-and-params.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 10 | `10-cg-execute-python-constants-and-magic-number-elimination.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 11 | `11-cg-execute-python-boolean-and-conditional-enhancement.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 12 | `12-cg-execute-python-array-constants-and-dynamic-enums.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 13 | `13-cg-execute-relative-paths.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 14 | `14-cg-execute-result-wrapper-and-apperror-returns.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 15 | `15-cg-execute-react-frontend-guidelines.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 16 | `16-cg-execute-testing-and-coverage.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 17 | `17-cg-execute-typescript-guidelines-and-types.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 18 | `18-cg-execute-boolean-improvements-v2.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 19 | `19-cg-execute-commit-fix-v2.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 20 | `20-cg-execute-nuclear-package-modularization-and-unit-test-optimization.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 21 | `21-cg-execute-in-below-steps.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 22 | `22-cg-execute-extract-types.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 23 | `23-cg-execute-follow-other-prompts.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 24 | `24-cg-execute-clean-artifacts-and-git-history.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 25 | `25-cg-execute-database-conventions.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 26 | `26-cg-execute-enum-standards.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 27 | `27-cg-execute-boolean-and-naming.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 28 | `28-cg-execute-error-management.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 29 | `29-cg-execute-isolate-os-tests.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 30 | `30-cg-execute-string-normalization-and-equalfoldany.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 31 | `31-cg-execute-branch-immutability-and-clean-construction.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 32 | `32-cg-execute-variadic-and-spread-parameters.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 33 | `33-cg-execute-map-dispatch-and-handler-lookup.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 34 | `34-cg-execute-string-efficiency.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 35 | `35-cg-execute-file-size-and-function-reduction.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 36 | `36-cg-execute-golang-pointer-reduction-and-value-semantics.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 37 | `37-cg-execute-code-dryness-and-library-extraction.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |
| 38 | `38-cg-execute-smart-test-runner-and-inventory.md` | `01-prompts/15-cg-execute/` & `01-prompts/v4/15-cg-execute/` |

---

## 3. Migration Mechanics

1. **Step 1: Automated Detection:** Inspect template blocks matching `### 📋 Confirmed Task Breakdown & Requirement Ingestion` across all 38 prompts in `01-prompts/15-cg-execute/`.
2. **Step 2: Insertion of Missing Keys:** Inject the two required keys directly beneath `- **State:** ...` for each task item.
3. **Step 3: Verification Script:** Run an automated verification script to assert that every prompt contains the required keys.
4. **Step 4: Mirror Synchronization:** Propagate verified updates to `01-prompts/v4/15-cg-execute/`.
5. **Step 5: Completion & Status Update:** Mark subtask as `DONE` and register results in `.ai-memory/plans/execute-parent-task-and-prompt-fixes.md`.

---

## 4. Execution Trace & Proof

*(To be populated during active execution)*

- [ ] All 38 prompts in `01-prompts/15-cg-execute/` upgraded.
- [ ] All 38 prompts in `01-prompts/v4/15-cg-execute/` mirrored.
- [ ] Verification script confirms 76/76 files compliant.
- [ ] Zero absolute paths or `file:///` URIs.
