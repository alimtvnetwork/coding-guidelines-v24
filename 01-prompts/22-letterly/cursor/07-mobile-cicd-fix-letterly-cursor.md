# Mobile CI/CD Fix Mode (Cursor) — Letterly Prompt Formatter

Combine the entire output into exactly ONE continuous single-line paragraph with zero newlines, zero line gaps, and zero conversational filler, using the Cursor skill format.

1. Clean the input text verbatim without conversational filler words (`um`, `ah`, `uh`).
2. Prefix with `[/goal](slashCommand;goal) [/learn](slashCommand;learn) Run gitmap pe -t to diagnose live CI/CD errors, fix all pipeline failures via 4-part RCA, verify locally, commit atomically via gitmap cpf, and release minor update. `.
3. Append `${Input Text Verbatim}`.
4. Suffix with ` - must follow the skill [.cursor/skills/ci-cd-fix-gitmap-release/skill.md](.cursor/skills/ci-cd-fix-gitmap-release/skill.md)`.
5. Output exactly that single line with no leading "Output" or markdown code fences.

${Input Text Verbatim} = The cleaned input text as it is, without filler words.

Output Format:
[/goal](slashCommand;goal) [/learn](slashCommand;learn) Run gitmap pe -t to diagnose live CI/CD errors, fix all pipeline failures via 4-part RCA, verify locally, commit atomically via gitmap cpf, and release minor update. ${Input Text Verbatim} - must follow the skill [.cursor/skills/ci-cd-fix-gitmap-release/skill.md](.cursor/skills/ci-cd-fix-gitmap-release/skill.md)
