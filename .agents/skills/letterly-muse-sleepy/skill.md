---
name: letterly-muse-sleepy
description: >-
  Formats raw voice dictation into Muse "I am sleepy" mode for autonomous execution with repo pull first, closest-choice ambiguity resolution, morning review log, no-build/no-test default, and pragmatic commits.
---

# Muse Sleepy Mode ("I am Sleepy") — Letterly Prompt Formatter

Format whatever input text is provided according to the exact high-priority execution template below without conversational filler or commentary.

1. Capture and clean the input text verbatim, stripping verbal filler words (`um`, `ah`, `uh`, `like`, leading `Okay`, excess `so`) while strictly preserving every technical directive, parameter, flag, and file path.
2. Derive `<task title>` as task subject (e.g. `SEO Writing and Folder Structure Instructions`). Output line 1: `# <task title>: high priority instruction, non-negotiable task`.
3. Output `${Input Text Verbatim}` directly beneath the title.
4. Output task name and slug subheaders:
   - `## task: <task title>`
   - `## slug: <task-slug>` (lowercase-hyphenated; when authoring directories and plan files, format as XX-<task-slug> where XX is the zero-padded sequence of that folder).
5. Construct `# Actionable Items Must Follow Non-Negotiable`:
   - Item 1: `1. Git pull BEFORE starting any work — first try \`gitmap pull <repo path>\` (or run \`gitmap pull\` in repository folder); if that fails, fall back to regular \`git pull\` and resolve any merge conflicts so the repository is completely up to date and clean before starting.`
   - Item 2: `2. Write spec under 02-spec/21-app/XX-<slug>/ and enqueue plan task in .ai-memory/plans/pending/XX-<slug>.md (subtasks in .ai-memory/plans/subtasks/XX-<slug>/) first, where XX is the zero-padded sequence of that folder. Planning is strictly dedicated to architecture, specification, boundaries, and task decomposition (NO-BUILD, NO-TEST, NO-CODE-CHANGES) with execution decoupled to subsequent phases.`
   - Item 3: `3. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String`
   - Item 4: `4. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well`
   - Item 5: `5. "I am sleepy" Autopilot & Ambiguity Protocol — the operator is going to sleep. Never halt, stall, or prompt for permission. If any ambiguity or underspecified requirement arises, follow the closest sensible solution possible to keep progress moving. Document all ambiguities, assumptions, and batch any questions in .ai-memory/ambiguity/XX-<slug>.md (or .ai-memory/memory/ambiguities.md) so the operator can review and resolve them together in the morning. If confusion or merge conflict arises, resolve cleanly and commit.`
   - Item 6: `6. Pragmatic commits & logical grouping via GitMap (forget the "perfect commit") — write code as strongly as possible adhering to coding guidelines and error management, but do NOT chase "perfect code" or the "perfect commit". Code does not need to build cleanly on every intermediate commit. Bundle related changes into logical commits via GitMap (gitmap cpf "<module> - <summary>", gitmap cpc "<module> - <summary>", gitmap commit "<message>"). Multiple subagents can commit their files together. Never commit file-by-file.`
   - Item 7: `7. Strict Zero Build/Test Default — in between commits, strictly DO NOT check the build or run tests. ONLY run build, test suites, or E2E tests IF explicitly requested in the prompt! If CI/CD verification is asked during the prompt, verify via gitmap pe -t until green (and execute minor release if requested). If CI/CD or build is not asked, skip build and test steps entirely.`
   - Items 8..N: Discrete technical directives extracted from input text.
   - Penultimate Item (conditional): If CI/CD fix is requested in prompt, `Verify live CI/CD pipeline health via gitmap pe -t until green`
   - Final Item (conditional): If release is requested in prompt, `Execute minor version bump release ceremony via python 03-ai-scripts/37-bump-version.py -t minor -s "<summary>", update changelog.md, commit atomically via gitmap cpr, tag release, and push to remote tracking branch`
6. Append mandatory agent invocation suffix pointing to `[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)`.
7. Output plain markdown text only. Never wrap in fenced code blocks (no ```markdown or ``` fences).
8. Relative Paths: Only add relative paths (`02-spec/...`, `.ai-memory/...`, `cmd/...`); never add absolute paths.

${Input Text Verbatim} = Cleaned input text without verbal filler words.

Output Format:

# <task title>: high priority instruction, non-negotiable task

${Input Text Verbatim}

## task: <task title>
## slug: <task-slug>  (should not contain `xx-` here. )

# Actionable Items Must Follow Non-Negotiable

1. Git pull BEFORE starting any work — first try `gitmap pull <repo path>` (or run `gitmap pull` in repository folder); if that fails, fall back to regular `git pull` and resolve any merge conflicts so the repository is completely up to date and clean before starting.
2. Write spec under 02-spec/21-app/XX-<slug>/ and enqueue plan task in .ai-memory/plans/pending/XX-<slug>.md (subtasks in .ai-memory/plans/subtasks/XX-<slug>/) first, (xx should be figured out by the folder structure and running AI tools, xx is sequence 01, 02). Planning is strictly dedicated to architecture, specification, boundaries, and task decomposition (NO-BUILD, NO-TEST, NO-CODE-CHANGES) with execution decoupled to subsequent phases.
3. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String
4. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well
5. "I am sleepy" Autopilot & Ambiguity Protocol — the operator is going to sleep. Never halt, stall, or prompt for permission. If any ambiguity or underspecified requirement arises, follow the closest sensible solution possible to keep progress moving. Document all ambiguities, assumptions, and batch any questions in .ai-memory/ambiguity/XX-<slug>.md (or .ai-memory/memory/ambiguities.md) so the operator can review and resolve them together in the morning. If confusion or merge conflict arises, resolve cleanly and commit.
6. Pragmatic commits & logical grouping via GitMap (forget the "perfect commit") — write code as strongly as possible adhering to coding guidelines and error management, but do NOT chase "perfect code" or the "perfect commit". Code does not need to build cleanly on every intermediate commit. Bundle related changes into logical commits via GitMap (gitmap cpf "<module> - <summary>", gitmap cpc "<module> - <summary>", gitmap commit "<message>"). Multiple subagents can commit their files together. Never commit file-by-file.
7. Strict Zero Build/Test Default — in between commits, strictly DO NOT check the build or run tests. ONLY run build, test suites, or E2E tests IF explicitly requested in the prompt! If CI/CD verification is asked during the prompt, verify via gitmap pe -t until green (and execute minor release if requested). If CI/CD or build is not asked, skip build and test steps entirely.
8. ....other steps and more steps sequentially from the input direction. Create more steps in between.
9. [Optional - if CI/CD fix requested] Verify live CI/CD pipeline health via gitmap pe -t until green
10. [Optional - if release requested] Execute minor version bump release ceremony via python 03-ai-scripts/37-bump-version.py -t minor -s "<summary>", update changelog.md, commit atomically via gitmap cpr, tag release, and push to remote tracking branch

### "I am sleepy" Autopilot Checklist
- [ ] 1. Repository pulled and clean before work (`gitmap pull` / `git pull`)
- [ ] 2. "I am sleepy" autopilot active: followed closest sensible solution for all ambiguities without waking or prompting the user
- [ ] 3. Ambiguities, assumptions, and batch questions recorded in `.ai-memory/ambiguity/` for morning clearance
- [ ] 4. Pragmatic logical commits via `gitmap`: grouped files together; no single-file commits; no "perfect commit" paralysis
- [ ] 5. Zero build/test default: intermediate builds and tests skipped; build/E2E executed ONLY if explicitly asked in prompt
- [ ] 6. Conditional CI/CD: pipeline verified via `gitmap pe -t` (and released) ONLY if explicitly requested in prompt
- [ ] 7. Every code change committed and pushed upstream (`gitmap push` / `gitmap cpf`) — No Push = Not Done

## Must follow and spawn agent using

[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)

## Additional Instructions

- [/plan](slashCommand;plan) first before doing work for improved work or less friction, faster work and more pure and effective work.
- [/learn](slashCommand;learn) from [gitmap](file;.agents/skills/gitmap) skill to leverage GitMap high-speed search, toolchain discovery, and caching.
- Track in-flight file touches and prevent write collisions using GitMap RC Heatmap (`gitmap rc heatmap status` / `gitmap rc heatmap tree` / `gitmap rc heatmap live` / `gitmap rc heatmap sync`).
- Only add relative paths, never add absolute paths during work; this must be respected on release page and in release notes as well.
