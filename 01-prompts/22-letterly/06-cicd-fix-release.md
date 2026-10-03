# CI/CD Fix & Release Mode — Letterly Prompt Formatter

Format whatever input text is provided according to the exact CI/CD fix and minor release template below. Do NOT add conversational filler or commentary.

1. Clean the input text verbatim while capturing all failing workflow names, errors, and reproduction steps.
2. Structure the output starting immediately with `# High Priority Instruction`.
3. Prepend `[/goal](slashCommand;goal) [/learn](slashCommand;learn) Diagnose, repair, and verify CI/CD pipeline failures using gitmap pe -t telemetry, apply targeted surgical fixes, commit atomically, and execute minor release ceremony:` before the input text.
4. Extract structured action items under `# Actionable Items Must Follow Non-Negotiable` (run `gitmap pe -t` to watch and capture failing CI/CD logs, conduct 4-part Root Cause Analysis, apply minimal surgical code fixes, verify with local linters without full runner bloat, commit atomically via `gitmap cpf "<module> - <summary>"`, execute minor version release ceremony).
5. Append the mandatory skill invocation suffix `[ci-cd-fix-with-release](file;.agents/skills/ci-cd-fix-with-release)`.
6. Output ONLY the resulting markdown block.

${Input Text Verbatim} = The cleaned input text as it is, without conversational filler words.

Output Format:

# High Priority Instruction

[/goal](slashCommand;goal) [/learn](slashCommand;learn) Diagnose, repair, and verify CI/CD pipeline failures using gitmap pe -t telemetry, apply targeted surgical fixes, commit atomically, and execute minor release ceremony: ${Input Text Verbatim}

# Actionable Items Must Follow Non-Negotiable

1. Inspect live CI/CD pipeline errors and execution timeline using gitmap pe -t (or gitmap pipeline fix)
2. Perform grounded 4-part Root Cause Analysis (RCA) on exact failing step and log
3. Apply surgical code fixes directly resolving the root cause without disabling any CI checks
4. Verify fixes locally with targeted file linters (05-guideline-autofixer.py, check-prompts-loaded.py)
5. Commit atomically via gitmap cpf "<module> - fix pipeline failure"
6. Execute minor version bump and release ceremony

Must follow and spawn agent using

[ci-cd-fix-with-release](file;.agents/skills/ci-cd-fix-with-release)

## Additional Instructions

learn [/learn](slashCommand;learn) if you have to learn something and [/plan](slashCommand;plan) stuff before working please.
