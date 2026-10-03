# Learned Protocol: IDE Skill Link Syntax & Plan Enqueueing Protocol

> **File:** `.ai-memory/memory/learned/21-ide-skill-links-and-plan-enqueueing-protocol.md`  
> **Date:** 2026-10-03  
> **Status:** ACTIVE  
> **Classification:** Rule & Skill Invariant  

---

## 1. Context & Motivation

During execution across desktop and mobile IDEs, standard markdown links (e.g. `[.agents/skills/my-skill/skill.md](.agents/skills/my-skill/skill.md)`) or shorthand mentions (e.g. `@[.agents/skills/my-skill]`) caused failures in native IDE skill parser detection.

The IDE parser detects skills and renders clickable execution pills only when formatted using the `file;` URI scheme prefix:
- Antigravity agents: `[<skill-name>](file;.agents/skills/<skill-name>)`
- Cursor IDE: `[<skill-name>](file;.cursor/skills/<skill-name>)`
- Slash commands: `[/plan](slashCommand;plan)` and `[/learn](slashCommand;learn)`

---

## 2. Execution Formatter Invariants (Letterly & Cursor)

Every execution prompt formatter (`execute-n-steps`, `execute-with-verification`, `execute-with-release`) must enforce:
1. **Line 1 `/plan` Command:** The generated output must start immediately with `[/plan](slashCommand;plan)` before `# High Priority Instruction`.
2. **Standardized Task Enqueueing Path:** Item 1 of Actionable Items must explicitly mandate:
   `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/) first`
3. **Retrospective AI Verification:**
   - In `execute-n-steps`: The final actionable item must mandate running retrospective AI verification (`01-retrospective-ai-verification.md` / `47-retrospective-ai-verification.py`) or skill `ai-verification`.
   - In `execute-with-verification`: Must mandate following `execute-parent-task-with-n-steps-v6` AND embed the verification skill (`[ai-verification](file;.agents/skills/ai-verification)` / `[ai-verification](file;.cursor/skills/ai-verification)`).

---

## 3. GitMap Commit Formatting

In GitMap commands (`gitmap cpf`, `gitmap cpb`, `gitmap cpr`):
- GitMap automatically provides `Feature: ` or `Bug: ` (with the colon).
- Providing colons in the CLI message argument results in double-colon defects.
- Commit messages MUST use the hyphen separator format:
  - `gitmap cpf "<module> - <summary>"`
  - `gitmap cpb "<module> - <summary>"`

---

## 4. Directory Structure Symmetry

- `01-prompts/22-letterly/`: Letterly prompts for Antigravity agents (`file;.agents/skills/...`).
- `01-prompts/23-cursor-prompts/`: Cursor-targeted prompts (`file;.cursor/skills/...`).
- `01-prompts/24-sync/`: Fleet-wide multi-repository synchronization engine.
- `01-prompts/25-ai-verification/`: Retrospective AI verification and quality auditing.
