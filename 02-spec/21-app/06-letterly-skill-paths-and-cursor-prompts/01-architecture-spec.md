# Architecture Specification: Letterly Skill Path Normalization & Cursor Skill Format Architecture

> **Specification Version:** 1.0.0  
> **Status:** Approved / Active  
> **Target Package / Directory:** `01-prompts/22-letterly/`, `01-prompts/22-letterly/cursor/`, `.agents/skills/`, `.cursor/skills/`

---

## 1. Executive Summary & Problem Analysis

### 1.1. Root Cause Analysis (RCA) of Prior Path Inconsistencies
Previously, Letterly formatted prompts utilized an ad-hoc link format:
`[<skill-name>](file;.agents/skills/<skill-name>)`
This format contained two fundamental errors:
1. **Invalid URI Scheme:** `file;` with a semicolon is invalid URI syntax that fails to render as an actionable markdown link or interactive file tag.
2. **Missing Terminal File Component:** The path stopped at the directory level (`.agents/skills/<skill-name>`) rather than referencing the concrete specification file (`skill.md`).
3. **Monolithic Agent Coupling:** All prompts unilaterally referenced `.agents/skills/`, making them incompatible with Cursor IDE workflows that locate skills in `.cursor/skills/`.

### 1.2. The Solution
1. **Explicit Folder & File Path Mandate:** All skill references MUST explicitly mention the complete relative path to the folder and file:
   - Agent Environment: `[.agents/skills/<skill-name>/skill.md](.agents/skills/<skill-name>/skill.md)` or `@[.agents/skills/<skill-name>/skill.md](.agents/skills/<skill-name>/skill.md)`
   - Cursor Environment: `[.cursor/skills/<skill-name>/skill.md](.cursor/skills/<skill-name>/skill.md)` or `@[.cursor/skills/<skill-name>/skill.md](.cursor/skills/<skill-name>/skill.md)`
2. **Dedicated Cursor Letterly Suite:** Author a parallel Cursor-optimized suite under `01-prompts/22-letterly/cursor/` referencing `.cursor/skills/<skill-name>/skill.md`.
3. **Bilateral Skill Synchronization:** Update `.agents/skills/` to reference `.agents/skills/` paths and `.cursor/skills/` to reference `.cursor/skills/` paths.

---

## 2. Specification: Letterly Prompt Suite Matrix

### 2.1. Agent Suite (`01-prompts/22-letterly/`)
References agent skills via `.agents/skills/<skill>/skill.md`:

| # | Prompt File | Target Mode | Exact Skill Suffix |
| :---: | :--- | :--- | :--- |
| **01** | `01-mobile-letterly.md` | Mobile Single-Line | `- must follow the skill [.agents/skills/execute-parent-task-with-n-steps-v6/skill.md](.agents/skills/execute-parent-task-with-n-steps-v6/skill.md)` |
| **02** | `02-desktop-letterly.md` | Desktop Structured | `[.agents/skills/execute-parent-task-with-n-steps-v6/skill.md](.agents/skills/execute-parent-task-with-n-steps-v6/skill.md)` |
| **03** | `03-execute-n-steps-letterly.md` | Execute N-Steps | `[.agents/skills/execute-parent-task-with-n-steps-v6/skill.md](.agents/skills/execute-parent-task-with-n-steps-v6/skill.md)` |
| **04** | `04-plan-letterly.md` | Plan Mode | `[.agents/skills/plan-spec-steps-v2/skill.md](.agents/skills/plan-spec-steps-v2/skill.md)` |
| **05** | `05-release-letterly.md` | Release Mode | `[.agents/skills/minor-bump/skill.md](.agents/skills/minor-bump/skill.md)` |
| **06** | `06-cicd-fix-release-letterly.md` | CI/CD Fix & Release | `[.agents/skills/ci-cd-fix-gitmap-release/skill.md](.agents/skills/ci-cd-fix-gitmap-release/skill.md)` |
| **07** | `07-mobile-cicd-fix-letterly.md` | Mobile CI/CD Fix | `- must follow the skill [.agents/skills/ci-cd-fix-gitmap-release/skill.md](.agents/skills/ci-cd-fix-gitmap-release/skill.md)` |

### 2.2. Cursor Suite (`01-prompts/22-letterly/cursor/`)
References Cursor skills via `.cursor/skills/<skill>/skill.md`:

| # | Prompt File | Target Mode | Exact Skill Suffix |
| :---: | :--- | :--- | :--- |
| **01** | `01-mobile-letterly-cursor.md` | Mobile Single-Line | `- must follow the skill [.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md](.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md)` |
| **02** | `02-desktop-letterly-cursor.md` | Desktop Structured | `[.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md](.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md)` |
| **03** | `03-execute-n-steps-letterly-cursor.md` | Execute N-Steps | `[.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md](.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md)` |
| **04** | `04-plan-letterly-cursor.md` | Plan Mode | `[.cursor/skills/plan-spec-steps-v2/skill.md](.cursor/skills/plan-spec-steps-v2/skill.md)` |
| **05** | `05-release-letterly-cursor.md` | Release Mode | `[.cursor/skills/minor-bump/skill.md](.cursor/skills/minor-bump/skill.md)` |
| **06** | `06-cicd-fix-release-letterly-cursor.md` | CI/CD Fix & Release | `[.cursor/skills/ci-cd-fix-gitmap-release/skill.md](.cursor/skills/ci-cd-fix-gitmap-release/skill.md)` |
| **07** | `07-mobile-cicd-fix-letterly-cursor.md` | Mobile CI/CD Fix | `- must follow the skill [.cursor/skills/ci-cd-fix-gitmap-release/skill.md](.cursor/skills/ci-cd-fix-gitmap-release/skill.md)` |

---

## 3. Core Invariants (Non-Negotiable)

1. **Explicit Folder & File Paths:** Always mention the directory AND file name (`.agents/skills/<name>/skill.md` or `.cursor/skills/<name>/skill.md`).
2. **Zero `file;` or `file:///` URIs:** Strict relative markdown link paths only.
3. **Single Paragraph Mobile Format:** All mobile formats output exactly ONE continuous line starting with `# High Priority Instruction: `.
4. **Desktop 5-Block Blueprint:** All desktop prompts start with `# High Priority Instruction`, include `1. Write spec and plan first` under `# Actionable Items`, and conclude with the mandatory skill link.
