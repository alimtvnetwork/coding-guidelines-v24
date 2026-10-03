---
name: letterly-plan
description: >-
  Formats raw voice dictation into high-priority architecture and planning instructions, actionable spec decomposition items, and plan-spec-steps-v2 skill invocation suffix.
---

# Plan Mode — Letterly Prompt Formatter

Format whatever input text is provided according to the exact planning output template below. Do NOT add conversational boilerplate, greetings, or acknowledgments.

1. Clean the input text verbatim by removing filler words (`um`, `ah`, `uh`) while preserving all feature requirements, architectural goals, and technical scope.
2. Structure the output starting immediately with `# High Priority Instruction`.
3. Prepend `[/plan](slashCommand;plan)` immediately before the cleaned verbatim input.
4. Extract structured planning action items under `# Actionable Items Must Follow Non-Negotiable` (author comprehensive architectural specification in `02-spec/21-app/<slug>/`, decompose into lean subtask plans in `.ai-memory/plans/subtasks/<slug>/`, enforce zero builds and zero tests during planning).
5. Append the mandatory planning skill invocation suffix `[plan-spec-steps-v2](file;.agents/skills/plan-spec-steps-v2)`.
6. Output ONLY the resulting markdown block.

${Input Text Verbatim} = The cleaned input text as it is, without conversational filler words.

Output Format:

# High Priority Instruction

[/plan](slashCommand;plan) ${Input Text Verbatim}

# Actionable Items Must Follow Non-Negotiable

1. Author comprehensive architectural specification under 02-spec/21-app/<slug>/
2. Decompose technical implementation into lean, bounded subtask plans under .ai-memory/plans/subtasks/<slug>/
3. Enforce strict no-build and no-test execution throughout the planning phase
4. Define binary acceptance criteria and verification commands for every subtask

Must follow and spawn agent using

[plan-spec-steps-v2](file;.agents/skills/plan-spec-steps-v2)

## Additional Instructions

learn [/learn](slashCommand;learn) if you have to learn something and [/plan](slashCommand;plan) stuff before working please.
