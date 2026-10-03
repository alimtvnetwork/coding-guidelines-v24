# Architecture Specification: Letterly & Cursor Prompts Architecture, IDE Syntax Standards & Spec/Plan Workflows

> **Spec ID:** `02-spec/21-app/09-multi-repo-sync-engine-and-prompt-upgrades/02-letterly-and-cursor-spec.md`  
> **Parent Architecture Spec:** `02-spec/21-app/09-multi-repo-sync-engine-and-prompt-upgrades/01-architecture-spec.md`  
> **Status:** APPROVED  
> **Target Subsystems:** `01-prompts/22-letterly/`, `01-prompts/23-cursor-prompts/`, `.agents/skills/letterly-*/`, `.cursor/skills/letterly-*/`

---

## 1. Executive Summary & Architectural Motivation

Voice dictation tools such as Letterly produce raw speech transcripts that must be formatted into rigorous, deterministic prompts for autonomous AI agents. Previously, prompt formatters suffered from several structural and syntactical defects:

1. **Nested Asymmetry:** Cursor prompt variants were nested inside `01-prompts/22-letterly/cursor/`, violating top-level prompt catalog conventions and preventing equal discoverability across environments.
2. **Non-IDE-Detectable Link Syntax:** References to skills and slash commands frequently used inconsistent formats (`@[...]`, bare directory paths, or standard markdown file links) that IDE parsers (Cursor, Antigravity, VSCode) failed to render as interactive clickable chips or capability badges.
3. **Missing Immediate Planning Activation:** Execution prompts failed to activate planning mode from the very first character, allowing agents to occasionally start coding prematurely.
4. **Underspecified Spec and Plan Targets:** Actionable Item 1 generically requested "Write spec and plan first" without enforcing standardized directory targets (`02-spec/21-app/<slug>/` and `.ai-memory/plans/<slug>.md`).
5. **Absence of Retrospective AI Verification:** Tasks could conclude without executing automated quality gates, leaving guideline adherence and CI/CD status unverified.

This specification defines the structural migration, the IDE-detectable syntax standard, the first-line slash command trigger, the strict specification and planning path convention, and the retrospective verification integration for all 10 Letterly and all 10 Cursor prompt formatters.

---

## 2. Directory Layout & Relocation

Letterly and Cursor prompts are organized as first-class peer directories directly under `01-prompts/`:

```text
01-prompts/
├── 22-letterly/                         # Canonical Letterly Voice Formatters (Antigravity / Agent)
│   ├── 01-mobile-letterly.md
│   ├── 02-desktop-letterly.md
│   ├── 03-execute-n-steps-letterly.md
│   ├── 04-plan-letterly.md
│   ├── 05-release-letterly.md
│   ├── 06-cicd-fix-release-letterly.md
│   ├── 07-mobile-cicd-fix-letterly.md
│   ├── 08-run-letterly.md
│   ├── 09-execute-with-verification-letterly.md
│   ├── 10-execute-with-release-letterly.md
│   └── readme.md
│
└── 23-cursor-prompts/                   # Canonical Cursor IDE Prompt Formatters
    ├── 01-mobile-cursor.md
    ├── 02-desktop-cursor.md
    ├── 03-execute-n-steps-cursor.md
    ├── 04-plan-cursor.md
    ├── 05-release-cursor.md
    ├── 06-cicd-fix-release-cursor.md
    ├── 07-mobile-cicd-fix-cursor.md
    ├── 08-run-cursor.md
    ├── 09-execute-with-verification-cursor.md
    ├── 10-execute-with-release-cursor.md
    └── readme.md
```

### Relocation Invariants
- The legacy nested folder `01-prompts/22-letterly/cursor/` is deprecated, moved to `01-prompts/23-cursor-prompts/`, and completely removed.
- Letterly formatters target Antigravity agent skills in `.agents/skills/`.
- Cursor formatters target Cursor IDE skills in `.cursor/skills/`.
- All filenames strictly use lowercase characters and hyphen-separated numbering.

---

## 3. Strict IDE-Detectable Syntax Standard

To ensure that both Cursor and Antigravity chat interfaces detect, parse, and render interactive UI chips, all generated prompt templates must strictly utilize semicolon-delimited custom scheme syntax:

### 3.1. Skill Reference Syntax
Skills must be referenced using the `file;` scheme pointing directly to the skill directory:

- **Agent Environment:**
  ```markdown
  [<skill-name>](file;.agents/skills/<skill-name>)
  ```
- **Cursor Environment:**
  ```markdown
  [<skill-name>](file;.cursor/skills/<skill-name>)
  ```

*Rules:*
- NEVER use standard URI schemes (`file` triple-slash URI, `http://`, `https://`).
- NEVER reference Windows absolute drive paths (`C:\...`, `D:\...`).
- NEVER append `/skill.md` inside the `file;` target; the IDE token parser matches against skill directory basenames.
- NEVER use bare `@` strings without the markdown link structure.

### 3.2. Slash Command Syntax
Slash commands must be referenced using the `slashCommand;` scheme:

- **Plan Slash Command:**
  ```markdown
  [/plan](slashCommand;plan)
  ```
- **Learn Slash Command:**
  ```markdown
  [/learn](slashCommand;learn)
  ```
- **Goal Slash Command:**
  ```markdown
  [/goal](slashCommand;goal)
  ```

*Rules:*
- The anchor text displays the slash command in square brackets (e.g., `[/plan]`).
- The link target is `slashCommand;<command-name>` without leading slashes.
- Semicolon delimiters are mandatory for IDE protocol recognition.

---

## 4. First-Line Slash Command Planning Trigger

For execution-oriented prompt formatters, the output format MUST immediately open on line 1 with the plan slash command chip:

```markdown
[/plan](slashCommand;plan)
```

### 4.1. Affected Prompts
The leading `[/plan](slashCommand;plan)` is mandatory on character 1 of the output block for:
- `01-prompts/22-letterly/03-execute-n-steps-letterly.md`
- `01-prompts/22-letterly/09-execute-with-verification-letterly.md`
- `01-prompts/22-letterly/10-execute-with-release-letterly.md`
- `01-prompts/23-cursor-prompts/03-execute-n-steps-cursor.md`
- `01-prompts/23-cursor-prompts/09-execute-with-verification-cursor.md`
- `01-prompts/23-cursor-prompts/10-execute-with-release-cursor.md`

### 4.2. Layout Invariant
The resulting formatted prompt starts with:

```markdown
[/plan](slashCommand;plan)
# High Priority Instruction

${Input Text Verbatim}
```

This guarantees that whenever a user pastes the formatted prompt into the chat, the IDE immediately recognizes the task as requiring a structured plan before any file modifications are permitted.

---

## 5. Standardized Actionable Item 1: Specification & Plan Enqueueing

Every structured desktop and execution prompt must enforce rigorous specification authoring and plan queuing before code execution. Actionable Item 1 is strictly standardized as:

```markdown
1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task under .ai-memory/plans/<slug>.md (subtasks under .ai-memory/plans/subtasks/<slug>/)
```

### 5.1. Token Definitions
- `<slug>`: A concise, hyphen-separated kebab-case identifier derived directly from the task's primary objective (e.g., `09-multi-repo-sync-engine-and-prompt-upgrades`).
- `02-spec/21-app/<slug>/`: Target directory for all architectural specifications, interface definitions, and acceptance criteria.
- `.ai-memory/plans/<slug>.md`: Master plan file tracking high-level milestones, phases, and global gates.
- `.ai-memory/plans/subtasks/<slug>/`: Directory containing granular, bounded subtask files executed by individual subagents.

### 5.2. Strict Ban on Premature Execution
No AI agent may touch production source code, run destructive commands, or modify application logic until Actionable Item 1 is completed and persisted to git.

---

## 6. Retrospective AI Verification Mandate

To prevent silent guideline degradation and unverified regressions, retrospective verification is embedded directly into execution workflows.

### 6.1. Execute N-Steps Prompts (`03-execute-n-steps-*`)
The final actionable item in `# Actionable Items Must Follow Non-Negotiable` must explicitly mandate running retrospective verification:

- **Letterly (`03-execute-n-steps-letterly.md`):**
  ```markdown
  N. Run retrospective AI verification prompt/script (01-retrospective-ai-verification.md / 03-ai-scripts/47-retrospective-ai-verification.py) or skill [ai-verification](file;.agents/skills/ai-verification) to audit specs, touched files, code quality, and CI/CD status upon task completion
  ```
- **Cursor (`03-execute-n-steps-cursor.md`):**
  ```markdown
  N. Run retrospective AI verification prompt/script (01-retrospective-ai-verification.md / 03-ai-scripts/47-retrospective-ai-verification.py) or skill [ai-verification](file;.cursor/skills/ai-verification) to audit specs, touched files, code quality, and CI/CD status upon task completion
  ```

### 6.2. Execute with Verification Prompts (`09-execute-with-verification-*`)
Prompts explicitly configured for verification mode must mandate execution via `execute-parent-task-with-n-steps-v6` AND embed the verification skill:

- **Letterly (`09-execute-with-verification-letterly.md`):**
  ```markdown
  Must follow and spawn agent using

  [execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6) and [ai-verification](file;.agents/skills/ai-verification)
  ```
- **Cursor (`09-execute-with-verification-cursor.md`):**
  ```markdown
  Must follow and spawn agent using

  [execute-parent-task-with-n-steps-v6](file;.cursor/skills/execute-parent-task-with-n-steps-v6) and [ai-verification](file;.cursor/skills/ai-verification)
  ```

---

## 7. Complete Prompt Catalog & Specification Matrix

The following matrix defines the exact behavior, skill suffixes, and requirements for all 10 prompts across both environments:

| # | Letterly File (`01-prompts/22-letterly/`) | Cursor File (`01-prompts/23-cursor-prompts/`) | Target Mode | First Line | Actionable Item 1 | Suffix Skill Link |
| :---: | :--- | :--- | :--- | :--- | :--- | :--- |
| **01** | `01-mobile-letterly.md` | `01-mobile-cursor.md` | Single-line mobile directive | `# High Priority Instruction: ` | N/A (single continuous line) | `[execute-parent-task-with-n-steps-v6](file;...)` |
| **02** | `02-desktop-letterly.md` | `02-desktop-cursor.md` | Structured desktop execution | `# High Priority Instruction` | Spec & plan path with `<slug>` | `[execute-parent-task-with-n-steps-v6](file;...)` |
| **03** | `03-execute-n-steps-letterly.md` | `03-execute-n-steps-cursor.md` | N-steps loop with verification | `[/plan](slashCommand;plan)` | Spec & plan path with `<slug>` | `[execute-parent-task-with-n-steps-v6](file;...)` + Retrospective AI Verification |
| **04** | `04-plan-letterly.md` | `04-plan-cursor.md` | Specification planning V2 | `# High Priority Instruction` | `02-spec/21-app/<slug>/` & `.ai-memory/plans/subtasks/<slug>/` | `[plan-spec-steps-v2](file;...)` |
| **05** | `05-release-letterly.md` | `05-release-cursor.md` | Minor release ceremony | `# High Priority Instruction` | Spec & plan path with `<slug>` | `[minor-bump](file;...)` |
| **06** | `06-cicd-fix-release-letterly.md` | `06-cicd-fix-release-cursor.md` | CI/CD RCA, fix & release | `# High Priority Instruction` | Spec & plan path with `<slug>` | `[ci-cd-fix-gitmap-release](file;...)` |
| **07** | `07-mobile-cicd-fix-letterly.md` | `07-mobile-cicd-fix-cursor.md` | Single-line mobile CI/CD fix | `# High Priority Instruction: ` | N/A (single continuous line) | `[ci-cd-fix-gitmap-release](file;...)` |
| **08** | `08-run-letterly.md` | `08-run-cursor.md` | Run script orchestration | `# High Priority Instruction` | Inspect `run.config.json` | `[run](file;...)` |
| **09** | `09-execute-with-verification-letterly.md` | `09-execute-with-verification-cursor.md` | Execution with dual verification | `[/plan](slashCommand;plan)` | Spec & plan path with `<slug>` | `[execute-parent-task-with-n-steps-v6](file;...)` AND `[ai-verification](file;...)` |
| **10** | `10-execute-with-release-letterly.md` | `10-execute-with-release-cursor.md` | Execution with auto minor release | `[/plan](slashCommand;plan)` | Spec & plan path with `<slug>` | `[minor-bump](file;...)` |

---

## 8. Companion Skills Synchronization (`.agents/skills/` & `.cursor/skills/`)

Every prompt formatter has a matching companion skill across both `.agents/skills/` and `.cursor/skills/`:

1. `letterly-mobile`
2. `letterly-desktop`
3. `letterly-execute-n-steps`
4. `letterly-plan`
5. `letterly-release`
6. `letterly-cicd-fix-release`
7. `letterly-mobile-cicd-fix`
8. `letterly-run`
9. `letterly-execute-with-verification`
10. `letterly-execute-with-release`

### 8.1. Skill Template Parity
- `.agents/skills/letterly-*/skill.md` must embed the exact template from `01-prompts/22-letterly/` and reference `.agents/skills/` paths using `file;.agents/skills/<skill-name>`.
- `.cursor/skills/letterly-*/skill.md` must embed the exact template from `01-prompts/23-cursor-prompts/` and reference `.cursor/skills/` paths using `file;.cursor/skills/<skill-name>`.
- YAML frontmatter must strictly include `name` and lowercase, multiline `description`.

---

## 9. Binary Acceptance Criteria

| Criteria ID | Category | Requirement Description | Verification Method | Status |
| :---: | :--- | :--- | :--- | :---: |
| **AC-01** | Directory | Letterly prompts reside at `01-prompts/22-letterly/` and Cursor prompts at `01-prompts/23-cursor-prompts/`. | `Test-Path` check | PASS |
| **AC-02** | Deprecation | Legacy nested folder `01-prompts/22-letterly/cursor/` is completely purged. | `Test-Path -Not` check | PASS |
| **AC-03** | Syntax | All skill references use `[<skill>](file;.agents/skills/<skill>)` or `[<skill>](file;.cursor/skills/<skill>)`. | Regex scanner | PASS |
| **AC-04** | Syntax | All slash commands use `[/plan](slashCommand;plan)` or `[/learn](slashCommand;learn)`. | Regex scanner | PASS |
| **AC-05** | First Line | `03-execute-n-steps-*`, `09-execute-with-verification-*`, and `10-execute-with-release-*` begin output with `[/plan](slashCommand;plan)`. | Header scanner | PASS |
| **AC-06** | Actionable Item 1 | Actionable Item 1 enforces spec in `02-spec/21-app/<slug>/` and plan in `.ai-memory/plans/<slug>.md`. | AST / Text audit | PASS |
| **AC-07** | Verification | `03-execute-n-steps-*` includes retrospective AI verification as the final actionable item. | Line match | PASS |
| **AC-08** | Verification | `09-execute-with-verification-*` mandates `execute-parent-task-with-n-steps-v6` AND embeds `ai-verification`. | AST / Text audit | PASS |
| **AC-09** | Companion Skills | All 10 skills in `.agents/skills/letterly-*` and 10 in `.cursor/skills/letterly-*` are fully synchronized. | File diff inspection | PASS |
| **AC-10** | Path Hygiene | Zero absolute paths, zero file triple-slash URIs, and zero uppercase file names across all touched artifacts. | Linter check (`07-relative-path-fixer.py`) | PASS |
