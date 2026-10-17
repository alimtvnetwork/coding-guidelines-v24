# Mobile CI/CD Fix Mode — Letterly Prompt Formatter

Combine the entire output into exactly ONE continuous single-line paragraph with zero newlines, zero line gaps, and zero conversational filler.

1. Clean the input text verbatim without conversational filler words (`um`, `ah`, `uh`).
2. Derive `<task title>` from the input text — it is the SUBJECT of the task (e.g. `CI/CD Pipeline Fix`). NEVER use the literal words `High Priority Instruction` as the title. Derive `<task-slug>` as the lowercase-hyphenated task title (e.g. `cicd-pipeline-fix`).
3. Prefix with `[/goal](slashCommand;goal) [/learn](slashCommand;learn) # <task title>: high priority instruction, non-negotiable task (task: <task title>, slug: <task-slug>) - first try gitmap pull <repo-path> (fallback git pull) to resolve conflicts before starting. Run gitmap pe -t to diagnose live CI/CD errors, fix all pipeline failures via 4-part RCA, verify locally, commit atomically via gitmap cpf, and release minor update. `.
4. Append `${Input Text Verbatim}`.
5. Suffix with ` - strictly use relative git paths only; only add the relative paths, never add the absolute path during your work (this must be respected on the release page and in release notes as well), must follow the skill [ci-cd-fix-gitmap-release](file;.agents/skills/ci-cd-fix-gitmap-release)`.
6. Output exactly that single continuous paragraph with ZERO newlines, ZERO line breaks, and NO leading "Output" text.
7. Relative Paths Mandate: Only add the relative paths, never add the absolute path during your work; ensure this is respected on the release page and in release notes as well.

${Input Text Verbatim} = The cleaned input text as it is, without filler words.

Output Format:
[/goal](slashCommand;goal) [/learn](slashCommand;learn) # <task title>: high priority instruction, non-negotiable task (task: <task title>, slug: <task-slug>) - first try gitmap pull <repo-path> (fallback git pull) to resolve conflicts before starting. Run gitmap pe -t to diagnose live CI/CD errors, fix all pipeline failures via 4-part RCA, verify locally, commit atomically via gitmap cpf, and release minor update. ${Input Text Verbatim} - strictly use relative git paths only; only add the relative paths, never add the absolute path during your work (this must be respected on the release page and in release notes as well), must follow the skill [ci-cd-fix-gitmap-release](file;.agents/skills/ci-cd-fix-gitmap-release)
