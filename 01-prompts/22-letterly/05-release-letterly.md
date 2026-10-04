# Release Mode — Letterly Prompt Formatter

Format whatever input text is provided according to the exact minor release template below, following the execute N-steps structure. Do NOT add conversational filler or commentary.

1. Clean the input text verbatim while strictly capturing version scope, changelog notes, and release constraints.
2. Structure the output starting immediately with `# High Priority Instruction`.
3. Put `${Input Text Verbatim}` directly beneath the high priority header.
4. Construct `# Actionable Items Must Follow Non-Negotiable`:
   - Item 1 is ALWAYS: `1. Write spec and plan first adhering to 02-spec/16-generic-release/ and 01-prompts/17-release-management/02-minor-bump.md`
   - Item 2 is ALWAYS: `2. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String`
   - Item 3 is ALWAYS: `3. Enforce zero-storage GitHub Actions rules (zero routine artifact uploads)`
   - Item 4 is ALWAYS: `4. Execute minor version bump via python 03-ai-scripts/37-bump-version.py -t minor -s "<summary>"`
   - Item 5 is ALWAYS: `5. Consolidate and update release notes in root changelog.md and manifests`
   - Item 6 is ALWAYS: `6. Commit atomically via gitmap cpf "<module> - release minor version" and push release tag to remote tracking branch`
5. Append the mandatory release skill invocation suffix `[minor-bump](file;.agents/skills/minor-bump)`.
6. Output ONLY the resulting formatted markdown block.

${Input Text Verbatim} = The cleaned input text as it is, without conversational filler words.

Output Format:

# High Priority Instruction

${Input Text Verbatim}

# Actionable Items Must Follow Non-Negotiable

1. Write spec and plan first adhering to 02-spec/16-generic-release/ and 01-prompts/17-release-management/02-minor-bump.md
2. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String
3. Enforce zero-storage GitHub Actions rules (zero routine artifact uploads)
4. Execute minor version bump via python 03-ai-scripts/37-bump-version.py -t minor -s "<summary>"
5. Consolidate and update release notes in root changelog.md and manifests
6. Commit atomically via gitmap cpf "<module> - release minor version" and push release tag to remote tracking branch

## Must follow and spawn agent using

[minor-bump](file;.agents/skills/minor-bump)

## Additional Instructions

- [/plan](slashCommand;plan) first before doing the work to reduce the credits.
- [/learn](slashCommand;learn) from [gitmap](file;.agents/skills/gitmap) skill to leverage GitMap high-speed search, toolchain discovery, and caching.
