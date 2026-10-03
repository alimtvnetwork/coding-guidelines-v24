---
name: letterly-execute-with-verification
description: >-
  Formats raw voice dictation into execute N-steps structure with a mandatory final action item to run retrospective AI verification upon task completion, using the cursor skill format.
---

# Execute with Verification Mode (Cursor) — Letterly Prompt Formatter

Format whatever input text is provided according to the exact execution with retrospective verification template below, following the execute N-steps structure and using the Cursor skill format. Do NOT add conversational filler or commentary.

1. Capture and clean the input text verbatim, stripping verbal filler words (`um`, `ah`, `uh`, `like`) while preserving every technical directive, parameter, flag, and file path.
2. Structure the output starting immediately with `# High Priority Instruction`.
3. Output `${Input Text Verbatim}` directly beneath the header.
4. Construct `# Actionable Items Must Follow Non-Negotiable`:
   - Item 1 is ALWAYS: `1. Write spec and plan first`
   - Item 2..N are sequential, discrete technical directives extracted from the input.
   - Final Item is ALWAYS: `Run retrospective AI verification prompt/script (01-retrospective-ai-verification.md / 03-ai-scripts/47-retrospective-ai-verification.py) to audit specs, touched files, code quality, and CI/CD status upon task completion`
5. Append the mandatory agent invocation suffix pointing to `@[.cursor/skills/execute-parent-task-with-n-steps-v6]`.
6. Output ONLY the resulting formatted markdown block.

${Input Text Verbatim} = The cleaned input text as it is, without conversational filler words.

Output Format:

# High Priority Instruction

${Input Text Verbatim}

# Actionable Items Must Follow Non-Negotiable

1. Write spec and plan first
2. [Second actionable technical directive extracted from input]
3. [Third actionable technical directive extracted from input]
4. Run retrospective AI verification prompt/script (01-retrospective-ai-verification.md / 03-ai-scripts/47-retrospective-ai-verification.py) to audit specs, touched files, code quality, and CI/CD status upon task completion

Must follow and spawn agent using

@[.cursor/skills/execute-parent-task-with-n-steps-v6]

## Additional Instructions

learn [/learn](slashCommand;learn) if you have to learn something and [/plan](slashCommand;plan) stuff before working please.
