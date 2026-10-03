---
name: letterly-execute-n-steps
description: >-
  Formats raw voice dictation into high-priority instructions, action items starting with write plan and spec, and execute-parent-task-with-n-steps-v6 skill invocation suffix.
---

# Execute N-Steps — Letterly Prompt Formatter

Format whatever input text is provided according to the exact high-priority execution template below. Do NOT add conversational filler or commentary (never write "Certainly! Here is your output:").

1. Capture and clean the input text verbatim, stripping verbal filler words (`um`, `ah`, `uh`, `like`) while preserving every technical directive, parameter, flag, and file path.
2. Structure the output starting immediately with `# High Priority Instruction`.
3. Output `${Input Text Verbatim}` directly beneath the header.
4. Construct `# Actionable Items Must Follow Non-Negotiable`:
   - Item 1 is ALWAYS: `1. Write a /plan [/plan](slashCommand;plan) and spec first`
   - Item 2..N are sequential, discrete technical directives extracted from the input.
5. Append the mandatory agent invocation suffix pointing to `[.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md](.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md)`.
6. Output ONLY the resulting formatted markdown block.

${Input Text Verbatim} = The cleaned input text as it is, without conversational filler words.

Output Format:

# High Priority Instruction

${Input Text Verbatim}

# Actionable Items Must Follow Non-Negotiable

1. Write a /plan [/plan](slashCommand;plan) and spec first
2. [Second actionable technical directive extracted from input]
3. [Third actionable technical directive extracted from input]

Must follow and spawn agent using

[.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md](.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md)

## Additional Instructions

learn [/learn](slashCommand;learn) if you have to learn something and [/plan](slashCommand;plan) stuff before working please.
