# Specification: Mobile Prompt Input Placement & Letterly Format Modernization

## 1. Overview & Objectives

In mobile Letterly prompts, the user's raw voice input text was previously appended *after* boilerplate checklist directives (`do gitmap pull ...; write spec ...; search codebase ...;`). This positioning caused the user's core task instructions to get buried or wasted beneath repetitive boilerplate clauses.

This specification establishes the new input-first placement standard for all mobile Letterly prompts and companion skills:
1. **Immediate Input Placement**:
   - The user's cleaned input text (`${Input Text Verbatim}`) MUST be placed **IMMEDIATELY AFTER** the header colon.
   - Format: `# <task title> - high priority instruction, non-negotiable task (task: <task title>, slug: <task-slug>): ${Input Text Verbatim}; `
2. **Sequential Checklist Following Input**:
   - The execution checklist (pull first, write spec, search codebase via GitMap, relative paths mandate, extracted discrete tasks, and mandatory skill trigger) follows cleanly after `${Input Text Verbatim}; `.
3. **Strict Single-Line Invariant**:
   - Output MUST remain strictly **ONE continuous single-line paragraph** with ZERO newlines, zero line gaps, and zero conversational filler.
4. **Surface Parity**:
   - Enforce across both standard mobile execute formatters (`01-mobile-letterly`) and mobile CI/CD fix formatters (`07-mobile-cicd-fix-letterly`) across Antigravity (`01-prompts/22-letterly/`, `.agents/skills/`) and Cursor IDE (`01-prompts/23-cursor-prompts/`, `.cursor/skills/`).

---

## 2. Standard Mobile Execute Contract

### Output Shape Contract:
```text
[/goal](slashCommand;goal) [/learn](slashCommand;learn) from [gitmap](file;.agents/skills/gitmap) [/plan](slashCommand;plan) # <task title> - high priority instruction, non-negotiable task (task: <task title>, slug: <task-slug>): ${Input Text Verbatim}; do gitmap pull <repo path> (fallback git pull) and resolve conflicts before starting; write spec under 02-spec/21-app/<slug>/ and enqueue plan in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/); search codebase exclusively via GitMap; strictly use relative git paths only; discrete understandable tasks to execute: 1. [task-1], 2. [task-2], 3. [task-3] - strictly use relative git paths only; only add the relative paths, never add the absolute path during your work (this must be respected on the release page and in release notes as well); must follow the skill [execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)
```

### Cursor Mirror Contract:
```text
[/goal](slashCommand;goal) [/learn](slashCommand;learn) from [gitmap](file;.cursor/skills/gitmap) [/plan](slashCommand;plan) # <task title> - high priority instruction, non-negotiable task (task: <task title>, slug: <task-slug>): ${Input Text Verbatim}; do gitmap pull <repo path> (fallback git pull) and resolve conflicts before starting; write spec under 02-spec/21-app/<slug>/ and enqueue plan in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/); search codebase exclusively via GitMap; strictly use relative git paths only; discrete understandable tasks to execute: 1. [task-1], 2. [task-2], 3. [task-3] - strictly use relative git paths only; only add the relative paths, never add the absolute path during your work (this must be respected on the release page and in release notes as well); must follow the skill [execute-parent-task-with-n-steps-v6](file;.cursor/skills/execute-parent-task-with-n-steps-v6)
```

---

## 3. Mobile CI/CD Fix Contract

### Output Shape Contract:
```text
[/goal](slashCommand;goal) [/learn](slashCommand;learn) from [gitmap](file;.agents/skills/gitmap) [/plan](slashCommand;plan) # <task title> - high priority instruction, non-negotiable task (task: <task title>, slug: <task-slug>): ${Input Text Verbatim}; first try gitmap pull <repo-path> (fallback git pull) and resolve conflicts before starting; run gitmap pe -t to diagnose live CI/CD errors, fix all pipeline failures via 4-part RCA, verify locally, commit atomically via gitmap cpf, and release minor update. discrete understandable tasks to execute: 1. diagnose CI/CD via gitmap pe -t, 2. apply 4-part RCA fix, 3. verify locally, 4. atomic commit via gitmap cpf, 5. minor release - strictly use relative git paths only; only add the relative paths, never add the absolute path during your work (this must be respected on the release page and in release notes as well), must follow the skill [ci-cd-fix-gitmap-release](file;.agents/skills/ci-cd-fix-gitmap-release)
```

### Cursor Mirror Contract:
```text
[/goal](slashCommand;goal) [/learn](slashCommand;learn) from [gitmap](file;.cursor/skills/gitmap) [/plan](slashCommand;plan) # <task title> - high priority instruction, non-negotiable task (task: <task title>, slug: <task-slug>): ${Input Text Verbatim}; first try gitmap pull <repo-path> (fallback git pull) and resolve conflicts before starting; run gitmap pe -t to diagnose live CI/CD errors, fix all pipeline failures via 4-part RCA, verify locally, commit atomically via gitmap cpf, and release minor update. discrete understandable tasks to execute: 1. diagnose CI/CD via gitmap pe -t, 2. apply 4-part RCA fix, 3. verify locally, 4. atomic commit via gitmap cpf, 5. minor release - strictly use relative git paths only; only add the relative paths, never add the absolute path during your work (this must be respected on the release page and in release notes as well), must follow the skill [ci-cd-fix-gitmap-release](file;.cursor/skills/ci-cd-fix-gitmap-release)
```

---

## 4. Acceptance Criteria

| Criteria ID | Requirement | Verification Method |
|---|---|---|
| AC-MOB-001 | `${Input Text Verbatim}` positioned immediately after `# ... (task: ..., slug: ...): ` | Inspect Step 6 and Output Format in `01-mobile-letterly.md` |
| AC-MOB-002 | Boilerplate directives follow `${Input Text Verbatim}; ` | Inspect Step 7 and Output Format in `01-mobile-letterly.md` |
| AC-MOB-003 | Cursor mirror has identical input placement and `.cursor/` skill links | Inspect `01-mobile-letterly-cursor.md` |
| AC-MOB-004 | Mobile CI/CD fix prompt has `${Input Text Verbatim}; ` right after colon | Inspect `07-mobile-cicd-fix-letterly.md` and cursor mirror |
| AC-MOB-005 | Companion skills `.agents/skills/` match prompt formatters verbatim | Inspect `letterly-mobile` and `letterly-mobile-cicd-fix` skills |
| AC-MOB-006 | Cursor companion skills match prompt formatters verbatim | Inspect `.cursor/skills/` counterparts |
| AC-MOB-007 | Zero line breaks in single-line output formats | Mechanical inspection of output format lines |
| AC-MOB-008 | Zero absolute file paths in specs, plans, and prompt files | Run `python linter-scripts/check-relative-paths.py .` |
