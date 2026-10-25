---
name: letterly-muse-long-update
description: >-
  Formats raw voice dictation into Muse long update mode for overnight execution with gitmap update, rate-limited gitmap pa / repo pull, autopilot ambiguity logging, pragmatic commits, and conditional CI release.
---

# Muse Long Update Mode (Overnight Autopilot) — Letterly Prompt Formatter

Format whatever input text is provided according to the exact high-priority execution template below without conversational filler or commentary.

1. Capture and clean the input text verbatim, stripping verbal filler words (`um`, `ah`, `uh`, `like`, leading `Okay`, excess `so`) while strictly preserving every technical directive, parameter, flag, and file path.
2. Derive `<task title>` as task subject (e.g. `SEO Writing and Folder Structure Instructions`). Output line 1: `# <task title>: high priority instruction, non-negotiable task`.
3. Output `${Input Text Verbatim}` directly beneath the title.
4. Output task name and slug subheaders:
   - `## task: <task title>`
   - `## slug: <task-slug>` (lowercase-hyphenated; when authoring directories and plan files, format as XX-<task-slug> where XX is the zero-padded sequence of that folder).
5. Construct `# Actionable Items Must Follow Non-Negotiable`:
   - Item 1: `1. Always update GitMap FIRST — run \`gitmap update\` to ensure latest GitMap binaries, CLI tooling, and agent commands are fully up to date before starting any task execution.`
   - Item 2: `2. Workspace pull via \`gitmap pa\` with smart rate-limiting — run \`gitmap pa\` across the workspace. If \`gitmap pa\` has already run recently (less than 15 minutes to 1 hour ago), skip \`gitmap pa\` and instead run \`gitmap pull <repo path>\` (or \`gitmap pull\` / \`git pull\` in the working repository folder) and resolve any merge conflicts so the target repository is clean before starting.`
   - Item 3: `3. Write spec under 02-spec/21-app/XX-<slug>/ and enqueue plan task in .ai-memory/plans/pending/XX-<slug>.md (subtasks in .ai-memory/plans/subtasks/XX-<slug>/) first, where XX is the zero-padded sequence of that folder. Planning is strictly dedicated to architecture, specification, boundaries, and task decomposition (NO-BUILD, NO-TEST, NO-CODE-CHANGES) with execution decoupled to subsequent phases.`
   - Item 4: `4. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String`
   - Item 5: `5. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well`
   - Item 6: `6. Enforce SQL repository pattern & enum standards — all SQL and data access behind repository interfaces; every action/command from a typed enum with full util suite (String, Parse, IsValid, All); structured error returns (*appfault.AppError) with typed error dispatch.`
   - Item 7: `7. Autopilot mode & ambiguity logging for morning review — operate autonomously without permission-seeking halts. If an ambiguity or underspecified detail arises, make the best executive decision or closest sensible choice immediately to maintain momentum. Document all ambiguities, assumptions, and batch any questions into the ambiguity section (.ai-memory/ambiguity/XX-<slug>.md or .ai-memory/memory/ambiguities.md) so the operator can review and resolve them together in the morning. If there is a confusion or conflict, resolve to the closest logical path and commit.`
   - Item 8: `8. Pragmatic commits & logical batching via GitMap (forget the "perfect commit") — write robust code adhering to coding guidelines and error management, but do NOT stall in pursuit of "perfect code" or a "perfect commit". Code does not need to build cleanly on every intermediate draft commit. Never commit files one-by-one; bundle related files from multiple subagents into clean, logical commits using GitMap (gitmap cpf "<module> - <summary>", gitmap cpc "<module> - <summary>", gitmap commit "<message>").`
   - Item 9: `9. Zero intermediate builds or tests — during execution and between commits, strictly DO NOT run builds, test suites, or E2E tests. Intermediate build and test checks are banned to maximize speed and prevent pipeline churn. Builds are checked ONLY at the very end, and ONLY IF explicitly requested in the prompt (e.g. CI/CD fix or test run requested). If CI/CD or build is not asked, skip them entirely.`
   - Items 10..N: Discrete technical directives extracted from input text.
   - Penultimate Item (conditional): If CI/CD fix is requested in prompt, `Verify live CI/CD pipeline health via gitmap pe -t until green`
   - Final Item (conditional): If release is requested in prompt, `Execute minor version bump release ceremony via python 03-ai-scripts/37-bump-version.py -t minor -s "<summary>", update changelog.md, commit atomically via gitmap cpr, tag release, and push to remote tracking branch`
6. Append mandatory agent invocation suffix pointing to `[execute-parent-task-with-n-steps-v6](file;.cursor/skills/execute-parent-task-with-n-steps-v6)`.
7. Output plain markdown text only. Never wrap in fenced code blocks (no ```markdown or ``` fences).
8. Relative Paths: Only add relative paths (`02-spec/...`, `.ai-memory/...`, `cmd/...`); never add absolute paths.

${Input Text Verbatim} = Cleaned input text without verbal filler words.

Output Format:

# <task title>: high priority instruction, non-negotiable task

${Input Text Verbatim}

## task: <task title>
## slug: <task-slug>  (should not contain `xx-` here. )

# Actionable Items Must Follow Non-Negotiable

1. Always update GitMap FIRST — run `gitmap update` to ensure latest GitMap binaries, CLI tooling, and agent commands are fully up to date before starting any task execution.
2. Workspace pull via `gitmap pa` with smart rate-limiting — run `gitmap pa` across the workspace. If `gitmap pa` has already run recently (less than 15 minutes to 1 hour ago), skip `gitmap pa` and instead run `gitmap pull <repo path>` (or `gitmap pull` / `git pull` in the working repository folder) and resolve any merge conflicts so the target repository is clean before starting.
3. Write spec under 02-spec/21-app/XX-<slug>/ and enqueue plan task in .ai-memory/plans/pending/XX-<slug>.md (subtasks in .ai-memory/plans/subtasks/XX-<slug>/) first, (xx should be figured out by the folder structure and running AI tools, xx is sequence 01, 02). Planning is strictly dedicated to architecture, specification, boundaries, and task decomposition (NO-BUILD, NO-TEST, NO-CODE-CHANGES) with execution decoupled to subsequent phases.
4. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String
5. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well
6. Enforce SQL repository pattern & enum standards — all SQL and data access behind repository interfaces; every action/command from a typed enum with full util suite (String, Parse, IsValid, All); structured error returns (*appfault.AppError) with typed error dispatch.
7. Autopilot mode & ambiguity logging for morning review — operate autonomously without permission-seeking halts. If an ambiguity or underspecified detail arises, make the best executive decision or closest sensible choice immediately to maintain momentum. Document all ambiguities, assumptions, and batch any questions into the ambiguity section (.ai-memory/ambiguity/XX-<slug>.md or .ai-memory/memory/ambiguities.md) so the operator can review and resolve them together in the morning. If there is a confusion or conflict, resolve to the closest logical path and commit.
8. Pragmatic commits & logical batching via GitMap (forget the "perfect commit") — write robust code adhering to coding guidelines and error management, but do NOT stall in pursuit of "perfect code" or a "perfect commit". Code does not need to build cleanly on every intermediate draft commit. Never commit files one-by-one; bundle related files from multiple subagents into clean, logical commits using GitMap (gitmap cpf "<module> - <summary>", gitmap cpc "<module> - <summary>", gitmap commit "<message>").
9. Zero intermediate builds or tests — during execution and between commits, strictly DO NOT run builds, test suites, or E2E tests. Intermediate build and test checks are banned to maximize speed and prevent pipeline churn. Builds are checked ONLY at the very end, and ONLY IF explicitly requested in the prompt (e.g. CI/CD fix or test run requested). If CI/CD or build is not asked, skip them entirely.
10. ....other steps and more steps sequentially from the input direction. Create more steps in between.
11. [Optional - if CI/CD fix requested] Verify live CI/CD pipeline health via gitmap pe -t until green
12. [Optional - if release requested] Execute minor version bump release ceremony via python 03-ai-scripts/37-bump-version.py -t minor -s "<summary>", update changelog.md, commit atomically via gitmap cpr, tag release, and push to remote tracking branch

### Muse Overnight Autopilot Checklist
- [ ] 1. Always update GitMap first (`gitmap update`)
- [ ] 2. Run `gitmap pa` (skip if run in <15m–1h, falling back to repository `gitmap pull` / `git pull`)
- [ ] 3. Autopilot ambiguity resolution: choose closest sensible solution, log to `.ai-memory/ambiguity/` for morning review
- [ ] 4. Batch all questions together for morning review without halting execution
- [ ] 5. Pragmatic logical commits via `gitmap` (no single-file commits; no "perfect commit" stalling; intermediate code doesn't need to build)
- [ ] 6. Zero intermediate builds/tests (build checked ONLY at end if CI/CD or test is requested)
- [ ] 7. Conditional CI/CD & minor release: if CI/CD fix is requested, verify `gitmap pe -t` to green, bump minor, tag, push
- [ ] 8. Upstream completion invariant: all modified files committed and pushed (`gitmap push` / `gitmap cpf` / `gitmap cpr`)

## Must follow and spawn agent using

[execute-parent-task-with-n-steps-v6](file;.cursor/skills/execute-parent-task-with-n-steps-v6)

## Additional Instructions

- [/plan](slashCommand;plan) first before doing work for improved work or less friction, faster work and more pure and effective work.
- [/learn](slashCommand;learn) from [gitmap](file;.cursor/skills/gitmap) skill to leverage GitMap high-speed search, toolchain discovery, and caching.
- Track in-flight file touches and prevent write collisions using GitMap RC Heatmap (`gitmap rc heatmap status` / `gitmap rc heatmap tree` / `gitmap rc heatmap live` / `gitmap rc heatmap sync`).
- Only add relative paths, never add absolute paths during work; this must be respected on release page and in release notes as well.
