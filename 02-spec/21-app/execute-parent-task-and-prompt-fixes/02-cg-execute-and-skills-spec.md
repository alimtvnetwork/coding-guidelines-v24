# Specification: CG Execute Prompts Classification, Migration & Companion Skills Alignment

> **Target:** `02-spec/21-app/execute-parent-task-and-prompt-fixes/02-cg-execute-and-skills-spec.md`  
> **Status:** Authoritative Technical Specification  
> **Version:** 1.0.0  
> **Date:** 2026-10-19  
> **Companion Specs:** `02-spec/21-app/execute-parent-task-and-prompt-fixes/01-architecture-spec.md`  
> **Parent Plan:** `.ai-memory/plans/execute-parent-task-and-prompt-fixes.md`  

---

## 1. Executive Summary & Architectural Scope

This specification formalizes the structural classification, migration mechanics, and companion skills alignment across the entire Coding Guidelines (CG) prompt ecosystem and skills registry. 

Building upon the 6-Part Task Verification Schema established in `02-spec/21-app/execute-parent-task-and-prompt-fixes/01-architecture-spec.md`, this specification defines:
1. **Structural Classification:** Partitioning the 38 CG execute prompts into Group 1 (37 standard preamble-precedence prompts) and Group 2 (Prompt 31 suffix-precedence prompt).
2. **Regex Migration Engine:** A deterministic Python migration architecture using boundary-anchored multiline regular expressions to update Phase 1A narrative instructions and task breakdown templates with zero formatting drift and automated mirroring to `01-prompts/v4/15-cg-execute/`.
3. **Companion Skills Audit & Parity:** A comprehensive mapping and synchronization strategy ensuring complete 1:1 schema parity across `.agents/skills/` (Google Antigravity) and `.cursor/skills/` (Cursor IDE) for the 4 priority execution skills and 36 remaining CG domain skills.
4. **Strict Relative Git Path Hygiene:** Total ban on absolute filesystem paths (`C:\...`, `/home/...`) and `file:///` URIs across all prompts, skills, specs, and manifests.

---

## 2. Structural Classification of the 38 CG Execute Prompts

The repository contains 38 executable prompts under `01-prompts/15-cg-execute/` and their identical counterparts in `01-prompts/v4/15-cg-execute/`. These prompts govern autonomous code refactoring and guideline enforcement. Structurally, they divide into two distinct functional groups based on instruction precedence directionality, task template cardinality, and file layout.

```
                           38 CG Execute Prompts
                                     |
         +---------------------------+---------------------------+
         |                                                       |
      Group 1                                                 Group 2
(37 Standard Prompts)                                    (Prompt 31 Unique)
01-30, 32-38                                            31-cg-execute-in-below-steps
- Above / Preamble Precedence                           - Below / Suffix Precedence
- Dual Task Template (Task-01 & Task-02)                - Single Task Template (Task-01)
- Standard Preflight Exit Anchor                        - Spec/Subtask Exit Anchor
- Preamble Conflict Resolution                          - Trailing Instructions Boundary
```

### 2.1 Group 1: Standard Prompts (37 Prompts: 01–30, 32–38)

Group 1 constitutes 37 of the 38 prompts. These prompts operate on the **Above Precedence / Preamble Precedence** model, where instructions preceding the prompt (user preamble, chat headers, or incoming task descriptions) outrank all lower-level instructions.

#### Key Structural Characteristics
1. **Header Metadata:**
   ```markdown
   > **Top-Instruction Priority Mandate (Above Precedence / Preamble Precedence):**
   > Whatever directives, constraints, checklists, or user instructions are given ABOVE this prompt (including in the user preamble, header blocks, or incoming user request above) are HIGHEST PRIORITY and MUST BE FOLLOWED as strictly NON-NEGOTIABLE.
   ```
2. **Phase 1A Step 1 Narrative:**
   ```markdown
   1. **Top-Instruction Priority Verification:** Whatever directives, constraints, checklists, or instructions are given before this section or prompt (user preamble, header constraints, prior instructions) must be verified as highest priority and non-negotiable.
   ```
3. **Task Template Cardinality:** The Phase 1A template illustrates **two** task items: `Task-01` (`[IN PROGRESS — EXECUTING IMMEDIATELY]`) and `Task-02` (`[QUEUED — EXECUTING NOW WITHOUT USER PROMPT]`), demonstrating proper queuing and status separation for multi-step execution.
4. **Phase 1A Exit Anchor:** Concludes with:
   ```markdown
   Proceeding directly to Preflight & Phase 1B Spec Generation (Active Tool Call Running Below).
   ```
5. **Precedence Hierarchy Section 1:**
   ```markdown
   1. **User Instructions & Preamble:** Directives and parameters ABOVE this prompt outrank everything below.
   2. **Platform Limits:** Native tools, Artifact Review Policy, permission prompts, hooks.
   3. **Repo Rules:** agents.md, .ai-memory/strictly-avoid.md, and coding-guidelines.md.
   4. **This Prompt.**
   ```

### 2.2 Group 2: Suffix-Precedence Prompt (1 Prompt: 31-cg-execute-in-below-steps.md)

Group 2 consists exclusively of `01-prompts/15-cg-execute/31-cg-execute-in-below-steps.md` (and its `v4` counterpart). This prompt implements the **Below Precedence / Suffix Precedence** model, designed specifically for workflows where users or parent dispatchers append custom checklists, bug reports, or instruction blocks *after* the prompt text.

#### Key Structural Characteristics
1. **Header Metadata:**
   ```markdown
   > **Bottom-Instruction Priority Mandate (Below Precedence / Suffix Precedence):**
   > Whatever directives, constraints, checklists, or user instructions are given BELOW this prompt (following the `--` divider border at the bottom) are HIGHEST PRIORITY and MUST BE FOLLOWED as strictly NON-NEGOTIABLE.
   ```
2. **Phase 1A Step 1 Narrative:**
   ```markdown
   1. **Bottom-Instruction Priority Verification:** Verify whatever directives, constraints, or task instructions are given BELOW this prompt (following the `--` divider border at the bottom) as highest priority and non-negotiable.
   ```
3. **Task Template Cardinality:** The Phase 1A template illustrates **one** task item (`Task-01: [Descriptive Task Title]`), focusing directly on the single appended instruction batch.
4. **Phase 1A Exit Anchor:** Concludes with:
   ```markdown
   Proceeding directly to Phase 1B: Spec & Subtask Generation (Active Tool Call Running Below).
   ```
5. **Precedence Hierarchy Section 1:**
   ```markdown
   1. **User Instructions & Suffix:** Directives and parameters BELOW this prompt (following the `--` divider border at the bottom) are HIGHEST PRIORITY and strictly NON-NEGOTIABLE.
   2. **Platform Limits:** Native tools, Artifact Review Policy, permission prompts, hooks.
   3. **Repo Rules:** agents.md, .ai-memory/strictly-avoid.md, and coding-guidelines.md.
   4. **This Prompt.**
   ```
6. **Trailing Boundary Marker:** Preserves the mandatory trailing instruction border at the end of the file:
   ```markdown
   ---

   ## 🚨 High Priority Instructions Below (Append Directives, Prompts, or Checklists Below)
   ```

### 2.3 Comprehensive Classification & Skill Mapping Matrix

| Index | Prompt Filename | Group | Primary Companion Skill (.agents / .cursor) | Domain Focus |
| :--- | :--- | :--- | :--- | :--- |
| 01 | `01-execute-coding-guideline-fix.md` | Group 1 | `execute-coding-guideline-fix`, `cg-execute` | General coding guideline remediation |
| 02 | `02-error-management.md` | Group 1 | `cg-error-management` | AppError wrappers & fault handling |
| 03 | `03-nested-if-and-guard-clauses.md` | Group 1 | `nested-if-and-guard-clauses` | Guard clauses & early returns |
| 04 | `04-booleans-and-complex-conditions.md` | Group 1 | `boolean-improvements-v2` | Boolean naming & condition decomposition |
| 05 | `05-naming-conventions-and-boolean-prefixes.md` | Group 1 | `cg-boolean-and-naming` | is/has prefixes & naming conventions |
| 06 | `06-constants-and-enums.md` | Group 1 | `cg-enum-standards` | Constant tables & enum types |
| 07 | `07-data-and-schema.md` | Group 1 | `cg-database-conventions` | Database schema & PascalCase tables |
| 08 | `08-react-frontend-guidelines.md` | Group 1 | `react-frontend-guidelines` | React component decomposition & hooks |
| 09 | `09-code-hygiene.md` | Group 1 | `code-hygiene` | Dead code, LF endings, formatting |
| 10 | `10-style-guidelines.md` | Group 1 | `style-guidelines` | Vertical spacing & blank line rules |
| 11 | `11-testing-and-coverage.md` | Group 1 | `testing-and-coverage` | Unit test coverage & table tests |
| 12 | `12-relative-paths.md` | Group 1 | `relative-paths` | Relative Git path enforcement |
| 13 | `13-cli-commands-and-help.md` | Group 1 | `cli-commands-and-help` | CLI argument parsing & help flags |
| 14 | `14-function-signatures-and-return-types.md` | Group 1 | `function-signatures-and-return-types` | Explicit returns & Result envelopes |
| 15 | `15-typescript-guidelines-and-types.md` | Group 1 | `typescript-guidelines-and-types` | TypeScript strictness & discriminated unions |
| 16 | `16-multi-language-enums-and-traits.md` | Group 1 | `multi-language-enums-and-traits` | Cross-language enum standards |
| 17 | `17-terminal-ui-and-cli-styling.md` | Group 1 | `terminal-ui-and-cli-styling` | Terminal formatting & spinners |
| 18 | `18-function-argument-reduction-and-params.md` | Group 1 | `function-argument-reduction-and-params` | Parameter structs for >2-3 arguments |
| 19 | `19-result-wrapper-and-apperror-returns.md` | Group 1 | `result-wrapper-and-apperror-returns` | Monadic Result[T] & *appfault.AppError |
| 20 | `20-extract-generic-types-to-types-go.md` | Group 1 | `cg-extract-types` | Centralized domain struct definitions |
| 21 | `21-lazy-regex-and-pattern-matching.md` | Group 1 | `lazy-regex-and-pattern-matching` | Pre-compiled package regexes |
| 22 | `22-file-size-and-function-reduction.md` | Group 1 | `cg-file-size-and-function-reduction` | 100-line file & 15-line function caps |
| 23 | `23-string-operations-and-efficiency.md` | Group 1 | `cg-string-efficiency` | Zero-allocation string operations |
| 24 | `24-isolate-destructive-os-and-heavy-unit-tests.md` | Group 1 | `cg-isolate-os-tests` | Mocked OS calls & unit test safety |
| 25 | `25-nuclear-package-modularization-and-unit-test-optimization.md` | Group 1 | `nuclear-package-modularization-and-unit-test-optimization` | Monolithic package decomposition |
| 26 | `26-python-boolean-and-conditional-enhancement.md` | Group 1 | `python-boolean-and-conditional-enhancement` | Python conditional optimization |
| 27 | `27-python-constants-and-magic-number-elimination.md` | Group 1 | `python-constants-and-magic-number-elimination` | Python magic numbers & constants |
| 28 | `28-python-array-constants-and-dynamic-enums.md` | Group 1 | `python-array-constants-and-dynamic-enums` | Hardcoded array elimination & Enums |
| 29 | `29-code-dryness-and-library-extraction.md` | Group 1 | `cg-code-dryness-and-library-extraction` | Shared util/framework extraction |
| 30 | `30-clean-repo-build-and-caches.md` | Group 1 | `cg-clean-repo-build-and-caches` | Cache purging & artifact sanitation |
| 31 | `31-cg-execute-in-below-steps.md` | **Group 2** | `cg-execute-in-below-steps` | Suffix-driven instruction execution |
| 32 | `32-cg-follow-other-prompts.md` | Group 1 | `cg-follow-other-prompts` | External prompt ingestion & execution |
| 33 | `33-branch-immutability-and-clean-construction.md` | Group 1 | `cg-branch-immutability-and-clean-construction` | Branch immutability & clean constructors |
| 34 | `34-golang-pointer-reduction-and-value-semantics.md` | Group 1 | `cg-golang-pointer-reduction-and-value-semantics` | Value semantics & pointer reduction |
| 35 | `35-clean-work-artifacts-and-os-caches.md` | Group 1 | `clean-artifacts-and-git-history` | Worktree artifact purging |
| 36 | `36-variadic-and-spread-parameters.md` | Group 1 | `cg-variadic-and-spread-parameters` | Variadic & spread parameter standards |
| 37 | `37-string-normalization-and-equalfoldany.md` | Group 1 | `cg-string-normalization-and-equalfoldany` | EqualFoldAny normalization helpers |
| 38 | `38-map-dispatch-and-handler-lookup.md` | Group 1 | `cg-map-dispatch-and-handler-lookup` | Map dispatch tables over long switches |

---

## 3. Migration Strategy & Python Regex Engine

To migrate all 38 prompts and ensure zero unintended formatting drift, modifications are executed programmatically via Python using boundary-anchored multiline regular expressions.

### 3.1 Pitfalls of Naive Replacements
- **String replacement (`str.replace`) failures:** Minor whitespace differences, newline character variations (`\r\n` vs `\n`), or prior localized edits across the 38 prompts make global string replacements brittle and error-prone.
- **Line-by-line streaming hazards:** Prompts contain nested code fences (` ```markdown ` inside markdown). Line-by-line parsers easily lose track of block nesting, leading to truncated sections or corrupted code blocks.
- **Accidental reflow:** Markdown formatters can inadvertently alter line breaks in unaffected prompt sections, causing massive, noisy diffs that obscure the architectural change.

### 3.2 Boundary-Anchored Multiline Regex Architecture

The migration engine isolates Phase 1A using rigid outer boundary anchors while replacing the entire block atomically:

```python
# Boundary Anchor Pattern for Phase 1A
PHASE_1A_PATTERN = re.compile(
    r"(### Phase 1A: Verbatim Capture, Task Extraction & Chat Output Gate \(Step 0\)\n\n)"
    r"([\s\S]*?)"
    r"(\n```\n\n---\n\n## 1\. Precedence Hierarchy)",
    re.MULTILINE
)
```

#### Match Invariants
1. **Start Anchor:** `### Phase 1A: Verbatim Capture, Task Extraction & Chat Output Gate (Step 0)\n\n`
2. **Block Interior:** Encompasses all narrative steps (1–6) and the enclosed ` ```markdown ... ``` ` block.
3. **End Anchor:** `\n```\n\n---\n\n## 1. Precedence Hierarchy`

### 3.3 Target Replacement Payloads

#### Group 1 Replacement Payload (37 Prompts)
```markdown
### Phase 1A: Verbatim Capture, Task Extraction & Chat Output Gate (Step 0)

Before executing any file searches, scans, spec writing, or code changes, you must execute Phase 1A:

1. **Top-Instruction Priority Verification:** Whatever directives, constraints, checklists, or instructions are given before this section or prompt (user preamble, header constraints, prior instructions) must be verified as highest priority and non-negotiable.
2. **Extreme "Understood" Turn-1 Showcase (Turn 1 Action):** In your VERY FIRST response turn upon receiving the prompt, you MUST output a comprehensive, in-depth understanding block directly in visible chat before any background tool execution. Any generic boilerplate or superficial 1-sentence summaries (e.g. `Understood: [YES] — I will do this task`) are protocol auto-reject failures. Turn 1 MUST display the top-level mandate containing all 5 Mandatory Pillars followed by the confirmed task breakdown using the 6-Part Task Verification Schema (with explicit `Understood / AI Comprehension` verification and `Azure AI / Model Runtime Context` declaration):
   * **Pillar 1: Strategic Intent & Systems Reasoning:** Articulate overarching architectural goal, user intent, operational rationale, and systems integration mechanics. Explain *why* this change is necessary and how downstream components interact.
   * **Pillar 2: Explicit Scope Boundaries:** Strict two-column partition defining In-Scope vs. strictly Excluded or Deferred items.
   * **Pillar 3: Non-Negotiable Constraints & Negative Guardrails Ingestion:** Explicit recitation of active constraints: raw search ban (`rg`, `ripgrep`, `Select-String`, `git grep`), worker git ban (no git commands in subagents), zero builds/full test suites (CI verification only), strict relative Git paths, lowercase naming, zero secrets in repo (`repo-secrets` offloading), GitMap hyphen-separated commit format without colons.
   * **Pillar 4: Blast Radius & Dependency Topology:** Map impacted subsystems, files, downstream callers, and prompt/skill indexes with explicit risk containment strategies.
   * **Pillar 5: Multi-Agent Concurrency Topology:** Detail swarm structure (`A = 3`, `H = 2`, Capacity = 6), disjoint file assignments, and immediate dynamic respawn.
3. **Lossless Verbatim Capture:** Store incoming prompt losslessly under `## User Request (Verbatim)` in canonical spec and parent plan.
4. **Screenshots & Media:** Decode base64/screenshots immediately into `assets/screenshots/<slug>-<NN>.png`. Reference via relative markdown links (`![Screenshot](assets/screenshots/<slug>-<NN>.png)`).
5. **Discrete Deliverables Extraction:** Break down whatever user requirements were given into discrete, actionable items with ordered traceable IDs (`Task-01`, `Task-02`, ...).
6. **Mandatory Same-Turn Tool Chaining (TOTAL BAN ON TURNING OFF):** Emit the comprehension mandate and breakdown in chat with clean vertical formatting, and in the **EXACT SAME TURN**, invoke your first tool call (e.g. `write_to_file` to initialize ledger/spec, or run preflight). NEVER emit text alone (which ends the turn prematurely), and never ask "Should I proceed?".

```markdown
### 🧠 Deep Comprehension & Understood Mandate

#### Pillar 1: Strategic Intent & Systems Reasoning
- [Articulate overarching architectural goal, user intent, operational rationale, and systems integration mechanics. Explain why this change is necessary and how downstream components interact.]

#### Pillar 2: Explicit Scope Boundaries (In-Scope vs. Excluded / Deferred)
- **In-Scope:** [Strict enumeration of files, modules, and changes executed in this turn.]
- **Excluded / Deferred:** [Explicitly deferred items, preserved untouched areas, and follow-up backlog.]

#### Pillar 3: Non-Negotiable Constraints & Negative Guardrails Ingestion
- [Explicit recitation of active constraints: raw search ban (no rg/ripgrep/Select-String), worker git ban (no git commands in subagents), zero builds/full test suites (CI verification only), strict relative Git paths, lowercase naming, zero secrets in repo (repo-secrets offloading), GitMap hyphen-separated commit format without colons.]

#### Pillar 4: Blast Radius & Dependency Topology
- **Impacted Subsystems & Files:** [Files, downstream callers, and prompt/skill indexes affected.]
- **Containment & Risk Mitigation:** [File lock prevention via disjoint ownership, backward compatibility invariants.]

#### Pillar 5: Multi-Agent Concurrency Topology & Wave Orchestration
- **Swarm Structure:** Lead Orchestrator (Root Node) -> Branch Coordinators (Subtask Owners) -> Leaf Workers & Dynamic Joining Assistants
- **Capacity:** A = 3 agents × H = 2 hands = 6 concurrent operations, WAVES = ceil(subtasks / 6)
- **Disjoint Bounding Boxes:** Zero file ownership overlap; dynamic immediate respawn on worker completion (100% duty cycle pinned at 3 active subagents)

### 📋 Confirmed Task Breakdown & Requirement Ingestion

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

2. **Task-02: [Descriptive Task Title]**
   - **State:** `[QUEUED — EXECUTING NOW WITHOUT USER PROMPT]`
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
     - `Owned Files (Exclusive):` `[strictly/relative/path/to/file-03.md]`
     - `Read-Only Context Files:` `[strictly/relative/path/to/reference.md]`
   - **Verification Evidence Contract:**
     - [Exact command or check to be executed to prove completion]
     - [Expected exit code: e.g. exit 0]
     - [Required diffstat evidence or structural regex check]

Proceeding directly to Preflight & Phase 1B Spec Generation (Active Tool Call Running Below).
```
```

#### Group 2 Replacement Payload (Prompt 31)
```markdown
### Phase 1A: Verbatim Capture, Task Extraction & Chat Output Gate (Step 0)

1. **Bottom-Instruction Priority Verification:** Verify whatever directives, constraints, or task instructions are given BELOW this prompt (following the `--` divider border at the bottom) as highest priority and non-negotiable.
2. **Extreme "Understood" Turn-1 Showcase (Turn 1 Action):** In your VERY FIRST response turn upon receiving the prompt, you MUST output a comprehensive, in-depth understanding block directly in visible chat before any background tool execution. Any generic boilerplate or superficial 1-sentence summaries are protocol auto-reject failures. Turn 1 MUST display the top-level mandate containing all 5 Mandatory Pillars followed by the confirmed task breakdown using the 6-Part Task Verification Schema (with explicit `Understood / AI Comprehension` verification and `Azure AI / Model Runtime Context` declaration):
   * **Pillar 1: Strategic Intent & Systems Reasoning:** Articulate overarching architectural goal, user intent, operational rationale, and systems integration mechanics.
   * **Pillar 2: Explicit Scope Boundaries:** Strict two-column partition defining In-Scope vs. strictly Excluded or Deferred items.
   * **Pillar 3: Non-Negotiable Constraints & Negative Guardrails Ingestion:** Explicit recitation of active constraints: raw search ban (`rg`, `ripgrep`, `Select-String`, `git grep`), worker git ban (no git commands in subagents), zero builds/full test suites, strict relative Git paths, lowercase naming, zero secrets in repo (`repo-secrets` offloading), GitMap hyphen-separated commit format without colons.
   * **Pillar 4: Blast Radius & Dependency Topology:** Map impacted subsystems, files, downstream callers, and prompt/skill indexes with explicit risk containment strategies.
   * **Pillar 5: Multi-Agent Concurrency Topology:** Detail swarm structure (`A = 3`, `H = 2`, Capacity = 6), disjoint file assignments, and immediate dynamic respawn.
3. **Lossless Verbatim Capture:** Store incoming prompt losslessly under `## User Request (Verbatim)` in canonical spec and parent plan.
4. **Mandatory Same-Turn Tool Chaining (TOTAL BAN ON TURNING OFF):** Emit the comprehension mandate and breakdown in chat with clean vertical formatting, and in the **EXACT SAME TURN**, invoke your first tool call. Never pause or ask "Should I proceed?".

```markdown
### 🧠 Deep Comprehension & Understood Mandate

#### Pillar 1: Strategic Intent & Systems Reasoning
- [Articulate overarching architectural goal, user intent, operational rationale, and systems integration mechanics. Explain why this change is necessary.]

#### Pillar 2: Explicit Scope Boundaries (In-Scope vs. Excluded / Deferred)
- **In-Scope:** [Strict enumeration of files, modules, and changes executed in this turn.]
- **Excluded / Deferred:** [Explicitly deferred items, preserved untouched areas, and follow-up backlog.]

#### Pillar 3: Non-Negotiable Constraints & Negative Guardrails Ingestion
- [Explicit recitation of active constraints: raw search ban (no rg/ripgrep/Select-String), worker git ban (no git commands in subagents), zero builds/full test suites (CI verification only), strict relative Git paths, lowercase naming, zero secrets in repo (repo-secrets offloading), GitMap hyphen-separated commit format without colons.]

#### Pillar 4: Blast Radius & Dependency Topology
- **Impacted Subsystems & Files:** [Files, downstream callers, and prompt/skill indexes affected.]
- **Containment & Risk Mitigation:** [File lock prevention via disjoint ownership, backward compatibility invariants.]

#### Pillar 5: Multi-Agent Concurrency Topology & Wave Orchestration
- **Swarm Structure:** Lead Orchestrator (Root Node) -> Branch Coordinators (Subtask Owners) -> Leaf Workers & Dynamic Joining Assistants
- **Capacity:** A = 3 agents × H = 2 hands = 6 concurrent operations, WAVES = ceil(subtasks / 6)
- **Disjoint Bounding Boxes:** Zero file ownership overlap; dynamic immediate respawn on worker completion (100% duty cycle pinned at 3 active subagents)

### 📋 Confirmed Task Breakdown & Requirement Ingestion

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
   - **Disjoint File Boundaries:**
     - `Owned Files (Exclusive):` `[strictly/relative/path/to/file-01.md]`
     - `Read-Only Context Files:` `[strictly/relative/path/to/reference.md]`
   - **Verification Evidence Contract:**
     - [Exact command or check to be executed to prove completion: e.g. exit 0]

Proceeding directly to Phase 1B: Spec & Subtask Generation (Active Tool Call Running Below).
```
```

### 3.4 Automated V4 Mirroring Invariant
Whenever a prompt in `01-prompts/15-cg-execute/` is modified, the migration engine immediately writes a byte-for-byte identical copy to `01-prompts/v4/15-cg-execute/<filename>`. This ensures that downstream tools querying either directory receive identical instruction sets.

---

## 4. Companion Skills Audit & Alignment Strategy

The repository supports two parallel skill ecosystems:
1. **Google Antigravity SDK Skills:** Located under `.agents/skills/<skill-name>/skill.md`.
2. **Cursor IDE Skills:** Located under `.cursor/skills/<skill-name>/skill.md`.

Each skill directory contains a single `skill.md` file featuring YAML frontmatter (`name`, `description`) followed by the full execution instructions.

### 4.1 Priority Execution Skills (4 High-Impact Skills)

Four priority skills require immediate, meticulous alignment:

1. **`execute-parent-task-with-n-steps-v6`:**
   - The master autonomous execution skill.
   - Updates: Full Phase 1A narrative upgrade and 6-part task verification schema in the task breakdown template.
2. **`cg-execute-in-below-steps`:**
   - Mirrors Prompt 31 (Group 2 suffix-precedence).
   - Updates: Ingests `Understood / AI Comprehension` and `Azure AI / Model Runtime Context` keys while preserving suffix-precedence directives and trailing instruction boundaries.
3. **`execute-coding-guideline-fix`:**
   - The flagship guideline fixing skill mirroring Prompt 01.
   - Updates: Ingests the 6-part task verification schema with dual-task queuing examples (`Task-01` and `Task-02`).
4. **`muse-master-prompt`:**
   - Governs fleet orchestration and multi-agent coordination.
   - Updates: Section 2 Item 2 (Task Confirmation & Execution Protocol) is updated to mandate the 6-part schema, explicitly citing `Understood / AI Comprehension: [YES]` and `Azure AI / Model Runtime Context: ...`.

### 4.2 The 36 Remaining CG Domain Skills

The remaining 36 CG domain skills directly correspond to the individual guideline rules (error management, booleans, style, naming, react, typescript, etc.). These skills must be audited and aligned to guarantee that any subagent dispatched under a specific guideline skill enforces identical Turn-1 understanding and runtime declarations.

#### Skill Synchronization Protocol
1. **Frontmatter Integrity:** Preserve all existing YAML frontmatter fields (`name`, `description`).
2. **Instruction Parity:** Synchronize the Phase 1A narrative and task breakdown template directly from the corresponding prompt in `01-prompts/15-cg-execute/`.
3. **Dual Directory Parity:** After updating `.agents/skills/<skill-name>/skill.md`, verify that `.cursor/skills/<skill-name>/skill.md` is updated identically. A hash check (`SHA256`) or byte comparison must report zero diffs.
4. **Index Hygiene:** Verify that `.agents/skills/` and `.cursor/skills/` maintain identical folder inventories.

---

## 5. Architectural Invariants & Negative Guardrails

1. **Strict Relative Git Paths (TOTAL BAN on Absolute Paths):**
   - All links, references, and citations in specifications, plans, subtasks, prompts, and skills MUST be strictly relative Git paths (e.g. `02-spec/21-app/...`, `01-prompts/15-cg-execute/...`, `.agents/skills/...`).
   - TOTAL BAN on absolute filesystem paths (`C:\...`, `/home/...`, `/tmp/...`) and `file:///` URIs.
2. **Strict Lowercase File Naming:**
   - All newly created files and folders must strictly use lowercase alphanumeric characters and hyphens (e.g. `02-cg-execute-and-skills-spec.md`, `03-subtask-skills-update.md`).
3. **GitMap Primacy (TOTAL BAN on Shell Search Tools):**
   - Codebase discovery and verification must use GitMap commands (`gitmap find`, `gitmap aum search`, `gitmap lf`, `gitmap cat`). Total ban on `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `dir`, `ls`, and all variants of `Get-ChildItem`.
4. **Worker Git Ban:**
   - Worker subagents never run git commands to prevent `.git/index.lock` collisions.
5. **GitMap Commit Formatting:**
   - All commit messages must use `gitmap cpf "<module> - <summary>"` with a hyphen separator and no colons inside the message argument.

---

## 6. Verification Gates & Acceptance Criteria

- [ ] **Structural Classification Verified:** Group 1 (37 prompts) and Group 2 (Prompt 31) are fully documented with explicit rationale, narrative differences, and template cardinality.
- [ ] **Migration Architecture Documented:** The boundary-anchored multiline regex strategy and target payloads are detailed with zero ambiguity.
- [ ] **Companion Skills Strategy Grounded:** Technical deliverables for updating both priority skills and the 36 remaining CG skills across `.agents/skills/` and `.cursor/skills/` are clearly defined.
- [ ] **Strict Relative Path Compliance:** 100% relative Git paths throughout the specification.
- [ ] **Cognitive Heatmap Logged:** Action recorded in `gitmap heatmap log` and verified via `gitmap heatmap status`.
