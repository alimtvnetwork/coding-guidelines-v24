---
name: letterly-desktop
description: >-
  Formats raw voice dictation into desktop high-priority instructions, non-negotiable action items, and execute-parent-task-with-n-steps skill invocation suffix without conversational boilerplate.
---

# Desktop Mode — Letterly Prompt Formatter

Format whatever input text is provided according to the exact output template below. Do NOT add conversational filler (never write "Certainly! Here is your output:"), and do NOT prepend introductions or acknowledgments.

1. Clean the input text verbatim by removing conversational filler words (`um`, `ah`, `uh`, `like`, `you know`) while strictly preserving every technical detail, requirement, file path, command, and directive.
2. Structure the output starting immediately with `# High Priority Instruction`.
3. Prepend `[/goal](slashCommand;goal) [/learn](slashCommand;learn)` immediately before the cleaned verbatim input.
4. Extract 3 to 6 discrete technical action items under `# Actionable Items Must Follow Non-Negotiable` reflecting the exact directives in the input.
5. End with the mandatory agent invocation suffix pointing to `[execute-parent-task-with-n-steps](file;.agents/skills/execute-parent-task-with-n-steps)`.
6. Output ONLY the resulting markdown block.

${Input Text Verbatim} = The cleaned input text as it is, without conversational filler words.

Output Format:

# High Priority Instruction

[/goal](slashCommand;goal) [/learn](slashCommand;learn) ${Input Text Verbatim}

# Actionable Items Must Follow Non-Negotiable

1. [First actionable technical directive extracted from input]
2. [Second actionable technical directive extracted from input]
3. [Third actionable technical directive extracted from input]

Must follow and spawn agent using

[execute-parent-task-with-n-steps](file;.agents/skills/execute-parent-task-with-n-steps)

## Additional Instructions

learn [/learn](slashCommand;learn) if you have to learn something and [/plan](slashCommand;plan) stuff before working please.
