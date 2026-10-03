# Architecture Specification: AI Verification Engine & Letterly Extensions

> **Spec ID:** `02-spec/21-app/08-ai-verification-and-letterly-extensions/01-architecture-spec.md`  
> **Status:** APPROVED  
> **Author:** Lead Systems Architect  
> **Target Subsystems:** `01-prompts/22-letterly/`, `01-prompts/24-ai-verification/`, `03-ai-scripts/`, `.agents/skills/`, `.cursor/skills/`

---

## 1. Problem Statement & Strategic Rationale

1. **Letterly Skill Trigger Failure:**  
   Standard markdown hyperlinks `[label](path)` do not trigger active skill autocompletion and agent dispatch in Google Antigravity or Cursor chat input interfaces. Active execution requires the explicit `@` token mention format: `@[.agents/skills/<skill-name>]` for Antigravity, and `@[.cursor/skills/<skill-name>]` for Cursor. An explicit rule exception is required for Letterly prompt formatters so they output the native trigger syntax rather than hyperlinked file paths.

2. **Lack of Retrospective AI Verification:**  
   When autonomous multi-agent pipelines complete complex parent tasks (spanning 2–3 subtasks over 30–40 minutes), there is currently no dedicated retrospective verification mechanism that discovers recent specs in `02-spec/21-app/`, audits touched files for coding guidelines compliance, validates git commit hygiene, and verifies live CI/CD status via GitMap telemetry.

3. **Missing Letterly Lifecycle Formatter Extensions:**  
   Users operating via voice dictation require dedicated formatters for:
   - **Run Execution:** Direct command and pre-flight dependency installation via `run.ps1` / `run.sh` with `@[.agents/skills/run]`.
   - **Execute with Retrospective Verification:** Standard N-step execution that explicitly mandates running the retrospective AI verification audit upon completion.
   - **Execute with Release Ceremony:** Standard N-step execution that executes the task, verifies CI/CD, and directly initiates a minor version bump release ceremony.

---

## 2. Technical Architecture & Invariants

### 2.1. Letterly Skill Path Exception Specification
In all Letterly prompt output templates (`01-prompts/22-letterly/` and `01-prompts/22-letterly/cursor/`), the skill invocation line MUST strictly format as:
- **Antigravity (Agent Suite):** `@[.agents/skills/<skill-name>]`
- **Cursor Suite:** `@[.cursor/skills/<skill-name>]`

This replaces prior markdown hyperlinks (`[...](...)`) specifically for the Letterly formatter templates to ensure native IDE skill resolution.

### 2.2. Retrospective AI Verification Engine (`24-ai-verification`)
- **Location:** `01-prompts/24-ai-verification/01-retrospective-ai-verification.md`
- **Architecture:** Full V6 Parent Task N-Steps specification (`N = 300, A = 2, H = 2, C = 30`).
- **Core Workflow:**
  1. **Phase 1 (Retrospective Discovery):** Inspect the last 2–3 git commits or changes made in the last 30–40 minutes. Identify generated specifications under `02-spec/21-app/`, completed subtasks in `.ai-memory/plans/`, and all modified code/spec files.
  2. **Phase 2 (Multilateral Audit):** Run targeted linters on touched files (implicit boolean checks, vertical line spacing, strict relative paths). Verify CI/CD pipeline health via GitMap telemetry (`gitmap pe -t`).
  3. **Phase 3 (Retrospective Scorecard):** Output a structured retrospective scorecard summarizing spec fidelity, code hygiene, and pipeline status.

### 2.3. Retrospective Verification Python Automation Script
- **Location:** `03-ai-scripts/47-retrospective-ai-verification.py`
- **Capabilities:**
  - Fast, dependency-free Python script querying git log for commits within the last `--since-minutes` (default: 34m) or `--commits` (default: 3).
  - Discovers touched files and validates adherence to acceptance criteria, relative paths, and zero-storage CI rules.
  - Generates clear console output and optional JSON/Markdown reports.

### 2.4. Expanded Letterly Suite Matrix
| # | Prompt File | Target Mode | Exact Skill Trigger |
| :---: | :--- | :--- | :--- |
| **01** | `01-mobile-letterly.md` | Mobile Single-Line | ` - must follow the skill @[.agents/skills/execute-parent-task-with-n-steps-v6]` |
| **02** | `02-desktop-letterly.md` | Desktop Structured | `@[.agents/skills/execute-parent-task-with-n-steps-v6]` |
| **03** | `03-execute-n-steps-letterly.md` | Execute N-Steps | `@[.agents/skills/execute-parent-task-with-n-steps-v6]` |
| **04** | `04-plan-letterly.md` | Planning Spec | `@[.agents/skills/plan-spec-steps-v2]` |
| **05** | `05-release-letterly.md` | Minor Release | `@[.agents/skills/minor-bump]` |
| **06** | `06-cicd-fix-release-letterly.md` | CI/CD Fix & Release | `@[.agents/skills/ci-cd-fix-gitmap-release]` |
| **07** | `07-mobile-cicd-fix-letterly.md` | Mobile CI/CD Fix | ` - must follow the skill @[.agents/skills/ci-cd-fix-gitmap-release]` |
| **08** | `08-run-letterly.md` | Run Script | `@[.agents/skills/run]` |
| **09** | `09-execute-with-verification-letterly.md` | Execute + AI Verify | `@[.agents/skills/execute-parent-task-with-n-steps-v6]` |
| **10** | `10-execute-with-release-letterly.md` | Execute + Release | `@[.agents/skills/minor-bump]` |

*(Identical mirrored suite for Cursor in `01-prompts/22-letterly/cursor/` referencing `@[.cursor/skills/<name>]`)*

---

## Acceptance Criteria

1. [ ] All 7 existing Agent Letterly prompts updated to use `@[.agents/skills/<name>]`.
2. [ ] All 7 existing Cursor Letterly prompts updated to use `@[.cursor/skills/<name>]`.
3. [ ] All 7 companion skills in `.agents/skills/` and `.cursor/skills/` updated accordingly.
4. [ ] Three new Letterly prompts added to both Agent and Cursor suites: `08-run`, `09-execute-with-verification`, `10-execute-with-release`.
5. [ ] Three companion skills created in `.agents/skills/` and `.cursor/skills/`: `letterly-run`, `letterly-execute-with-verification`, `letterly-execute-with-release`.
6. [ ] `01-prompts/24-ai-verification/` created with `readme.md` and `01-retrospective-ai-verification.md` adhering to V6 architecture.
7. [ ] Automation script `03-ai-scripts/47-retrospective-ai-verification.py` created and functional.
8. [ ] Companion skills `.agents/skills/ai-verification/skill.md` and `.cursor/skills/ai-verification/skill.md` created.
9. [ ] All prompt index checks (`check-prompts-loaded.py`) and relative paths linters (`check-relative-paths.py`) PASS cleanly.
