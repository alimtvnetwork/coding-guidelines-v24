# Architecture Specification: Letterly Golden Prompt Standardization, Retrospective AI Verification Isolation & Execution Priority Governance

> **Spec ID:** `02-spec/21-app/11-gitmap-training-and-letterly-golden-prompt-sync/01-architecture-spec.md`  
> **Status:** APPROVED  
> **Author:** Lead Systems Architect & Spec Subagent 01  
> **Target Subsystems:** `AGENTS.md`, `01-prompts/22-letterly/`, `01-prompts/23-cursor-prompts/`, `.agents/skills/letterly-*/`, `.cursor/skills/letterly-*/`  
> **Reference Artifact:** `assets/screenshots/ai-verification-plan-tag-removal-01.png`  

---

## 1. Executive Summary

This architecture specification establishes strict systemic standards across all voice-dictation (Letterly) prompt formatters, Cursor IDE prompt mirrors, Antigravity skills, and root repository guidelines (`AGENTS.md`).

A forensic analysis of prompt generation defects—documented in `assets/screenshots/ai-verification-plan-tag-removal-01.png`—identified critical governance contradictions:
1. Erroneous placement of the slash command `[/plan](slashCommand;plan)` at Line 1 above top-priority execution headers.
2. Catastrophic misclassification of retrospective AI verification as a planning task.
3. Uncontrolled leakage of retrospective AI verification into generic execution prompts where it was never requested.
4. Divergence between canonical golden prompt patterns and active prompt templates.

This specification codifies `01-prompts/22-letterly/03-execute-n-steps-letterly.md` as the canonical Golden Prompt, restricts retrospective AI verification exclusively to dedicated verification prompts (`09-execute-with-verification-letterly.md` and its mirrors), and establishes rigorous structural standards for planning, execution, and release workflows.

---

## 2. Root Cause Analysis (RCA): Erroneous `[/plan]` Placement & Verification Pollution

### 2.1. Empirical Defect Analysis from Screenshot

The visual evidence captured in `assets/screenshots/ai-verification-plan-tag-removal-01.png` highlights a severe formatting failure in a generated task prompt:

```text
+-----------------------------------------------------------------------------------+
|  [/plan](slashCommand;plan)                                                      |
|  # High Priority Instruction: Retrospective AI Verification Audit                 |
|                                                                                   |
|  Conduct a retrospective quality audit on Task 37 (GitMap Search Subagent         |
|  Enforcement & SQLite Task Schema). Audit specifications in 02-spec/21-app/10-... |
+-----------------------------------------------------------------------------------+
```

In this captured failure:
- The slash command `[/plan](slashCommand;plan)` occupies Line 1, placed directly above `# High Priority Instruction: Retrospective AI Verification Audit`.
- The task itself is an audit and verification mission inspecting completed work, yet it is instructed to begin by running `/plan`.

### 2.2. Four-Part Root Cause Analysis

```
+-----------------------------------------------------------------------------------+
| 1. SYMPTOM:                                                                       |
|    - `[/plan](slashCommand;plan)` prepended to Line 1 of verification prompt.     |
|    - Standard `execute-n-steps` prompts forced to mandate AI verification.        |
+-----------------------------------------+-----------------------------------------+
                                          |
                                          v
+-----------------------------------------------------------------------------------+
| 2. PROXIMATE CAUSE:                                                               |
|    - `AGENTS.md` Section 10 mandated: "Line 1 of generated execution prompts     |
|      must begin directly with `[/plan](slashCommand;plan)`."                      |
|    - `AGENTS.md` Section 10 mandated: "Final actionable item for execute-n-steps  |
|      must trigger retrospective AI verification."                                 |
+-----------------------------------------+-----------------------------------------+
                                          |
                                          v
+-----------------------------------------------------------------------------------+
| 3. ROOT CAUSE:                                                                    |
|    - Conflation of task execution entry point with interactive slash command tips.|
|    - Category error treating retrospective verification as a planning workflow.   |
|    - Failure to maintain single responsibility: mixing verification into standard|
|      execution formatters rather than isolating to dedicated verification mode.   |
+-----------------------------------------+-----------------------------------------+
                                          |
                                          v
+-----------------------------------------------------------------------------------+
| 4. SYSTEMIC RESOLUTION:                                                           |
|    - Mandate Line 1 ALWAYS starts with `# High Priority Instruction`.             |
|    - Relegate `[/plan]` strictly to bottom `## Additional Instructions`.          |
|    - Strip retrospective AI verification from all 9 non-verification prompts.    |
|    - Enforce golden prompt structure across all 40 prompts and skills.            |
+-----------------------------------------------------------------------------------+
```

#### Part 1: Symptom & Observed Damage
- **Prompt Priority Inversion:** Agents encountering `[/plan](slashCommand;plan)` on Line 1 interpret the slash command as the primary execution directive. Instead of executing the substantive high-priority instructions, agents enter an unnecessary planning interview or produce empty task decomposition skeletons.
- **Workflow Paralysis on Audits:** When retrospective AI verification prompts inherit `[/plan]`, the agent attempts to plan an audit rather than executing immediate verification scripts (`03-ai-scripts/47-retrospective-ai-verification.py`) and inspecting GitMap telemetry.
- **Scope Creep in Standard Execution:** Generic implementation tasks formatted by `execute-n-steps` were forced to append retrospective AI verification as the final action item, adding unwarranted overhead to self-contained tasks.

#### Part 2: Proximate Cause
- In `AGENTS.md` Section 10, the "Execution Formatter Standards" contained two flawed rules:
  1. `Line 1 of generated execution prompts must begin directly with [/plan](slashCommand;plan).`
  2. `Final actionable item for execute-n-steps must trigger retrospective AI verification (47-retrospective-ai-verification.py / ai-verification).`
- These faulty rules propagated into `01-prompts/23-cursor-prompts/03-execute-n-steps-letterly-cursor.md`, `01-prompts/23-cursor-prompts/10-execute-with-release-letterly-cursor.md`, and related skills.

#### Part 3: Root Cause
- **Semantic Mismatch between Verification and Planning:** Retrospective AI verification is an audit and quality gate executed on existing, completed work. Planning is a prospective decomposition performed on unstarted work. Forcing a prospective planning command (`/plan`) onto a retrospective verification audit is logically invalid.
- **Loss of Canonical Golden Standard:** The canonical Letterly prompt `01-prompts/22-letterly/03-execute-n-steps-letterly.md` correctly structured the header on Line 1 and placed `[/plan]` at the bottom under `## Additional Instructions`. However, secondary prompts and mirrors diverged from this golden baseline due to conflicting rules in `AGENTS.md`.

#### Part 4: Systemic Prevention
- Update `AGENTS.md` Section 10 to establish Line 1 header primacy and restrict AI verification strictly to `09-execute-with-verification-letterly.md`.
- Refactor all 10 Letterly prompts, all 10 Cursor prompts, all 10 Antigravity skills, and all 10 Cursor skills to enforce unified invariants.

---

## 3. Core Architectural Invariants

### 3.1. Invariant 1: Top-Header Priority Primacy (Line 1 Mandate)
- **Rule:** Every generated prompt and prompt template MUST begin directly on Line 1 with `# High Priority Instruction` (or `# High Priority Instruction: <Title>` for single-line formatters).
- **TOTAL BAN:** NEVER place slash commands (`[/plan]`, `[/learn]`, `[/goal]`) or conversational preamble text above `# High Priority Instruction`.
- **Rationale:** AI orchestrators and subagents parse input prompts from top to bottom. The top header dictates the operational authority and priority boundary of the entire turn.

### 3.2. Invariant 2: Retrospective AI Verification Isolation
- **Rule:** Retrospective AI verification is strictly an audit and verification quality gate, NOT a planning step.
- **Prohibition:** Verification audit prompts MUST NEVER contain `[/plan](slashCommand;plan)`.
- **Execution Target:** Retrospective verification operates by inspecting git commit history, validating acceptance criteria in `02-spec/21-app/`, verifying coding guidelines via `03-ai-scripts/47-retrospective-ai-verification.py`, and evaluating GitMap CI/CD telemetry (`gitmap pe -t`).

### 3.3. Invariant 3: Strict Restriction of Retrospective AI Verification Scope
- **Rule:** Retrospective AI verification is authorized ONLY inside:
  - `01-prompts/22-letterly/09-execute-with-verification-letterly.md`
  - `01-prompts/23-cursor-prompts/09-execute-with-verification-letterly-cursor.md`
  - `.agents/skills/letterly-execute-with-verification/skill.md`
  - `.cursor/skills/letterly-execute-with-verification/skill.md`
- **Total Ban on Standard Formatters:** Retrospective AI verification MUST NOT appear in:
  - Standard Execute: `03-execute-n-steps-letterly*`
  - Desktop: `02-desktop-letterly*`
  - Mobile: `01-mobile-letterly*`, `07-mobile-cicd-fix-letterly*`
  - Plan: `04-plan-letterly*`
  - Release: `05-release-letterly*`, `10-execute-with-release-letterly*`
  - Run: `08-run-letterly*`
  - CI/CD Fix: `06-cicd-fix-release-letterly*`

### 3.4. Invariant 4: Slash Command Placement Standard
- **Rule:** Where slash commands are helpful (such as `[/plan]` for complex implementation tasks), they MUST reside strictly under `## Additional Instructions` at the very bottom of the prompt template.
- **Syntax:** Slash commands must use the IDE-detectable syntax:
  - `[/plan](slashCommand;plan)`
  - `[/learn](slashCommand;learn)`
  - `[/goal](slashCommand;goal)`

---

## 4. Canonical Golden Prompt Specification

`01-prompts/22-letterly/03-execute-n-steps-letterly.md` serves as the authoritative Golden Prompt governing all execution prompt formatters.

### 4.1. Structural Breakdown

```markdown
# Execute N-Steps — Letterly Prompt Formatter

Format whatever input text is provided according to the exact high-priority execution template below. Do NOT add conversational filler or commentary (never write "Certainly! Here is your output:").

1. Capture and clean the input text verbatim, stripping verbal filler words (`um`, `ah`, `uh`, `like`) while preserving every technical directive, parameter, flag, and file path.
2. Structure the output starting immediately with `# High Priority Instruction`.
3. Output `${Input Text Verbatim}` directly beneath the header.
4. Construct `# Actionable Items Must Follow Non-Negotiable`:
   - Item 1 is ALWAYS: `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/) first`
   - Item 2 is ALWAYS: `2. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String`
   - Item 3..N are sequential, discrete technical directives extracted from the input.
5. Append the mandatory agent invocation suffix pointing to `[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)`.
6. Make sure all the action items are listed and nothing pending.

${Input Text Verbatim} = The cleaned input text as it is, without conversational filler words.

Output Format:

# High Priority Instruction

${Input Text Verbatim}

# Actionable Items Must Follow Non-Negotiable

1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/) first
2. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String
3. ....other steps and more steps sequentially from the input direction. Create more steps in between.

## Must follow and spawn agent using

[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)

## Additional Instructions

[/plan](slashCommand;plan) First before doing the work to reduce the credits.
```

### 4.2. Invariant Checklist for Golden Prompt Compliance
- [x] Header begins on Line 1: `# High Priority Instruction`
- [x] Item 1 enforces relative paths for spec (`02-spec/21-app/<slug>/`) and plan (`.ai-memory/plans/<slug>.md`)
- [x] Item 2 mandates GitMap search primacy and explicitly bans `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`
- [x] Zero retrospective AI verification in action items
- [x] Agent invocation suffix targets `execute-parent-task-with-n-steps-v6`
- [x] `[/plan]` located strictly in `## Additional Instructions` at the bottom

---

## 5. Domain-Specific Prompt Specifications

### 5.1. Plan Prompt Specification (`04-plan-letterly.md` & Mirrors)

The planning prompt formatter guides architects in authoring specifications and plan files before execution commences.

#### Required Actionable Items:
1. **Item 1:** `1. Write the plan and architectural spec first under 02-spec/21-app/<slug>/ adhering to 02-spec/01-spec-authoring-guide/`
2. **Item 2:** `2. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String`
3. **Item 3:** `3. Write the first step on task decomposition and bounded subtasks under .ai-memory/plans/subtasks/<slug>/`
4. **Item 4:** `4. Enforce strict no-build and no-test rules throughout the planning phase`
5. **Item 5..N:** Discrete architectural, scoping, and acceptance criteria directives extracted from input.

#### Skill Suffix & Additional Instructions:
- Antigravity: `[plan-spec-steps-v2](file;.agents/skills/plan-spec-steps-v2)`
- Cursor: `[plan-spec-steps-v2](file;.cursor/skills/plan-spec-steps-v2)`
- Bottom section:
  ```markdown
  ## Additional Instructions

  [/plan](slashCommand;plan) first before doing the work to reduce the credits.
  ```

### 5.2. Release Prompt Specification (`05-release-letterly.md` & Mirrors)

The release prompt formatter automates minor semantic version bumps and release hygiene.

#### Required Actionable Items:
1. **Item 1:** `1. Write spec and plan first adhering to 02-spec/16-generic-release/ and 01-prompts/17-release-management/02-minor-bump.md`
2. **Item 2:** `2. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String`
3. **Item 3:** `3. Enforce zero-storage GitHub Actions rules (zero routine artifact uploads)`
4. **Item 4:** `4. Execute minor version bump via python 03-ai-scripts/37-bump-version.py -t minor -s "<summary>"`
5. **Item 5:** `5. Consolidate and update release notes in root changelog.md and manifests`
6. **Item 6:** `6. Commit atomically via gitmap cpf and push release tag to remote tracking branch`

#### Skill Suffix & Additional Instructions:
- Antigravity: `[minor-bump](file;.agents/skills/minor-bump)`
- Cursor: `[minor-bump](file;.cursor/skills/minor-bump)`
- Bottom section:
  ```markdown
  ## Additional Instructions

  [/plan](slashCommand;plan) first before doing the work to reduce the credits.
  ```

### 5.3. Execute with Verification Specification (`09-execute-with-verification-letterly.md` & Mirrors)

The ONLY execution formatter authorized to embed retrospective AI verification.

#### Required Actionable Items:
1. **Item 1:** `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/) first`
2. **Item 2:** `2. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String`
3. **Item 3..N:** Sequential discrete directives extracted from input.
4. **Final Item:** `Run retrospective AI verification prompt/script (01-prompts/25-ai-verification/01-retrospective-ai-verification.md / 03-ai-scripts/47-retrospective-ai-verification.py) or skill [ai-verification](file;.agents/skills/ai-verification) to audit specs, touched files, code quality, and CI/CD status upon task completion`

#### Skill Suffix:
- Antigravity: `[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)`
- Cursor: `[execute-parent-task-with-n-steps-v6](file;.cursor/skills/execute-parent-task-with-n-steps-v6)`

### 5.4. Execute with Release Specification (`10-execute-with-release-letterly.md` & Mirrors)

Enforces complete task execution followed directly by minor release ceremony.

#### Required Actionable Items:
1. **Item 1:** `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/) first`
2. **Item 2:** `2. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String`
3. **Item 3..N:** Discrete technical directives extracted from input.
4. **Penultimate Item:** `Verify live CI/CD pipeline health via gitmap pe -t until green`
5. **Final Item:** `Execute minor version bump release ceremony via python 03-ai-scripts/37-bump-version.py -t minor -s "<summary>", update changelog.md, commit atomically via gitmap cpf, tag release, and push to remote tracking branch`

#### Header & Suffix Governance:
- Line 1 must begin immediately with `# High Priority Instruction` (stripped leading `[/plan]` tag).
- Suffix targets `[minor-bump]`.

---

## 6. Multi-Platform Prompt & Skill Inventory Matrix

| # | Prompt Role | Letterly Prompt (`01-prompts/22-letterly/`) | Cursor Prompt (`01-prompts/23-cursor-prompts/`) | Antigravity Skill (`.agents/skills/`) | Cursor Skill (`.cursor/skills/`) | Has AI Verification | Line 1 Header |
| :- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 01 | Mobile Execute | `01-mobile-letterly.md` | `01-mobile-letterly-cursor.md` | `letterly-mobile` | `letterly-mobile` | NO | `# High Priority Instruction: ` |
| 02 | Desktop Execute | `02-desktop-letterly.md` | `02-desktop-letterly-cursor.md` | `letterly-desktop` | `letterly-desktop` | NO | `# High Priority Instruction` |
| 03 | Execute N-Steps (Golden) | `03-execute-n-steps-letterly.md` | `03-execute-n-steps-letterly-cursor.md` | `letterly-execute-n-steps` | `letterly-execute-n-steps` | NO | `# High Priority Instruction` |
| 04 | Plan Mode | `04-plan-letterly.md` | `04-plan-letterly-cursor.md` | `letterly-plan` | `letterly-plan` | NO | `# High Priority Instruction` |
| 05 | Release Mode | `05-release-letterly.md` | `05-release-letterly-cursor.md` | `letterly-release` | `letterly-release` | NO | `# High Priority Instruction` |
| 06 | CI/CD Fix & Release | `06-cicd-fix-release-letterly.md` | `06-cicd-fix-release-letterly-cursor.md` | `letterly-cicd-fix-release` | `letterly-cicd-fix-release` | NO | `# High Priority Instruction` |
| 07 | Mobile CI/CD Fix | `07-mobile-cicd-fix-letterly.md` | `07-mobile-cicd-fix-letterly-cursor.md` | `letterly-mobile-cicd-fix` | `letterly-mobile-cicd-fix` | NO | `[/goal] [/learn] Run gitmap pe -t...` |
| 08 | Run Mode | `08-run-letterly.md` | `08-run-letterly-cursor.md` | `letterly-run` | `letterly-run` | NO | `# High Priority Instruction` |
| 09 | Execute with Verification | `09-execute-with-verification-letterly.md` | `09-execute-with-verification-letterly-cursor.md` | `letterly-execute-with-verification` | `letterly-execute-with-verification` | **YES** | `# High Priority Instruction` |
| 10 | Execute with Release | `10-execute-with-release-letterly.md` | `10-execute-with-release-letterly-cursor.md` | `letterly-execute-with-release` | `letterly-execute-with-release` | NO | `# High Priority Instruction` |

---

## 7. `AGENTS.md` Section 10 Refactoring Specification

To resolve the root cause permanently, `AGENTS.md` Section 10 must be updated with the following precise text:

```markdown
- **Execution Formatter Standards (Letterly & Cursor):**
  - Line 1 of generated execution prompts must begin directly with `# High Priority Instruction`.
  - Slash command suggestions such as `[/plan](slashCommand;plan)` belong strictly under `## Additional Instructions` at the bottom of execution prompts, and are strictly prohibited on verification audit prompts.
  - Item 1 of Actionable Items must explicitly mandate writing specs in `02-spec/21-app/<slug>/` and enqueueing plan tasks in `.ai-memory/plans/<slug>.md` (with subtasks in `.ai-memory/plans/subtasks/<slug>/`).
  - Item 2 of Actionable Items must explicitly mandate searching codebase exclusively via GitMap (`gitmap aum search`, `gitmap find`, `gitmap cat`, `gitmap ps`) with a total ban on `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`.
  - Retrospective AI verification is an audit and verification gate, NOT a planning step. Only `09-execute-with-verification-letterly.md` (and its mirrors) contains retrospective AI verification; standard `execute-n-steps` formatters MUST NOT mandate retrospective AI verification.
  - Execution with verification must embed both `execute-parent-task-with-n-steps-v6` and `ai-verification`.
```

---

## 8. Acceptance Criteria

### 8.1. Root Cause Resolution & Invariants
- [ ] `assets/screenshots/ai-verification-plan-tag-removal-01.png` defect is comprehensively analyzed with a 4-part RCA in this specification.
- [ ] Retrospective AI verification is established as strictly an audit gate with zero `[/plan]` prepending.
- [ ] Line 1 header primacy is mandated across all execution prompt formatters.

### 8.2. Prompt Suite Consistency
- [ ] All 10 prompts under `01-prompts/22-letterly/` comply with the Golden Prompt standard.
- [ ] All 10 prompts under `01-prompts/23-cursor-prompts/` comply with the Golden Prompt standard.
- [ ] Retrospective AI verification is present ONLY in prompt 09 and its mirrors; stripped from all other 9 prompts.
- [ ] Plan prompts cite `02-spec/01-spec-authoring-guide/` in Item 1 and GitMap search in Item 2.
- [ ] Release prompts cite `02-spec/16-generic-release/` and `01-prompts/17-release-management/02-minor-bump.md` in Item 1 and GitMap search in Item 2.

### 8.3. Companion Skills Parity
- [ ] All 10 skills in `.agents/skills/letterly-*/` match their canonical `01-prompts/22-letterly/` counterparts.
- [ ] All 10 skills in `.cursor/skills/letterly-*/` match their canonical `01-prompts/23-cursor-prompts/` counterparts.

### 8.4. Repository Guidelines
- [ ] `AGENTS.md` Section 10 updated to eliminate erroneous `[/plan]` on Line 1 and remove forced AI verification from generic `execute-n-steps`.
- [ ] All references adhere strictly to relative git paths with zero absolute paths or `file:///` URIs.
