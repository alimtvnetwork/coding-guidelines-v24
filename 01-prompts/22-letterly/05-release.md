# Release Mode — Letterly Prompt Formatter

Format whatever input text is provided according to the exact release output template below. Do NOT add conversational filler or explanations.

1. Clean the input text verbatim while preserving version requirements, changelog highlights, and release constraints.
2. Structure the output starting immediately with `# High Priority Instruction`.
3. Prepend `[/goal](slashCommand;goal) Execute minor version release ceremony:` before the input text.
4. Extract structured release action items under `# Actionable Items Must Follow Non-Negotiable` (update root `changelog.md`, execute version bump script, enforce zero-storage GitHub Actions rules, commit via GitMap, tag release, push to origin).
5. Append the mandatory release skill invocation suffix `[minor-bump](file;.agents/skills/minor-bump)`.
6. Output ONLY the resulting markdown block.

${Input Text Verbatim} = The cleaned input text as it is, without conversational filler words.

Output Format:

# High Priority Instruction

[/goal](slashCommand;goal) Execute minor version release ceremony across the repository: ${Input Text Verbatim}

# Actionable Items Must Follow Non-Negotiable

1. Consolidate and update release notes in root changelog.md
2. Increment minor version across all project manifests (package.json, version.json, pyproject.toml, Cargo.toml)
3. Enforce zero-storage GitHub Actions rules (zero artifact uploads)
4. Commit release changes atomically via GitMap using hyphen format
5. Tag release version and push to remote tracking branch

Must follow and spawn agent using

[minor-bump](file;.agents/skills/minor-bump)

## Additional Instructions

learn [/learn](slashCommand;learn) if you have to learn something and [/plan](slashCommand;plan) stuff before working please.
