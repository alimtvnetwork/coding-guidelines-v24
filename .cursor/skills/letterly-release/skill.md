---
name: letterly-release
description: >-
  Formats raw voice dictation into minor version release ceremony instructions starting with git pull, changelog updates, and minor-bump skill invocation suffix.
---

# Release Mode (Cursor) — Letterly Prompt Formatter

Format whatever input text is provided according to the exact minor release template below, following the execute N-steps structure, using the Cursor skill format. Do NOT add conversational filler or commentary.

1. Clean the input text verbatim while strictly capturing version scope, changelog notes, and release constraints.
2. Structure the output starting immediately with `# High Priority Instruction`.
3. Put `${Input Text Verbatim}` directly beneath the high priority header.
4. Construct `# Actionable Items Must Follow Non-Negotiable`:
   - Item 1 is ALWAYS: `1. Git pull BEFORE starting any work — first try \`gitmap pull <repo path>\` (or run \`gitmap pull\` in repository folder); if that fails, fall back to regular \`git pull\` and resolve any merge conflicts so the repository is completely up to date and clean before starting.`
   - Item 2 is ALWAYS: `2. Write spec and plan first adhering to 02-spec/16-generic-release/ and 01-prompts/17-release-management/02-minor-bump.md`
   - Item 3 is ALWAYS: `3. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String`
   - Item 4 is ALWAYS: `4. Strictly use relative Git paths only; only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page, in release notes, changelog.md, manifests, and documentation`
   - Item 5 is ALWAYS: `5. Enforce zero-storage GitHub Actions rules (zero routine artifact uploads)`
   - Item 6 is ALWAYS: `6. Execute minor version bump via python 03-ai-scripts/37-bump-version.py -t minor -s "<summary>"`
   - Item 7 is ALWAYS: `7. Consolidate and update release notes in root changelog.md and manifests`
   - Item 8 is ALWAYS: `8. Commit atomically via gitmap cpf "<module> - release minor version" and push release tag to remote tracking branch`
5. Append the mandatory release skill invocation suffix `[minor-bump](file;.cursor/skills/minor-bump)`.
6. Output ONLY the resulting formatted prompt as plain markdown text. STRICT — NEVER wrap it in a fenced code block: never emit ```markdown, ```plaintext, or any ``` fence at the start or end of the output. `#` headers and inline markdown are correct; the fence around them is forbidden. No conversational filler, no commentary before or after it.
7. Relative Paths Mandate: Only add the relative paths, never add the absolute path during your work; ensure this is respected on the release page, in release notes, changelog.md, manifests, and documentation.
8. Additional Instructions Mandate: Always append the relative paths directive under ## Additional Instructions: '- Only add the relative paths, never add the absolute path during your work; this should be respected on the release page and in release notes as well.'

${Input Text Verbatim} = The cleaned input text as it is, without conversational filler words.

Output Format:

# High Priority Instruction

${Input Text Verbatim}

# Actionable Items Must Follow Non-Negotiable

1. Git pull BEFORE starting any work — first try `gitmap pull <repo path>` (or run `gitmap pull` in repository folder); if that fails, fall back to regular `git pull` and resolve any merge conflicts so the repository is completely up to date and clean before starting.
2. Write spec and plan first adhering to 02-spec/16-generic-release/ and 01-prompts/17-release-management/02-minor-bump.md
3. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String
4. Strictly use relative Git paths only; only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page, in release notes, changelog.md, manifests, and documentation
5. Enforce zero-storage GitHub Actions rules (zero routine artifact uploads)
6. Execute minor version bump via python 03-ai-scripts/37-bump-version.py -t minor -s "<summary>"
7. Consolidate and update release notes in root changelog.md and manifests
8. Commit atomically via gitmap cpf "<module> - release minor version" and push release tag to remote tracking branch

## Must follow and spawn agent using

[minor-bump](file;.cursor/skills/minor-bump)

## Additional Instructions

- [/plan](slashCommand;plan) first before doing the work to reduce the credits.
- [/learn](slashCommand;learn) from [gitmap](file;.cursor/skills/gitmap) skill to leverage GitMap high-speed search, toolchain discovery, and caching.
- Only add the relative paths, never add the absolute path during your work; this should be respected on the release page and in release notes as well.
