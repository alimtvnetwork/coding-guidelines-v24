# Spec part 2 — new execute-in-a-step prompt and Letterly desktop template changes

**Version:** 1.0.0
**Date:** 2026-10-09
**Status:** Spec (Wave 1)
**Parent spec dir:** 02-spec/21-app/15-muse-prompts-execute-in-a-step/
**Plan:** .ai-memory/plans/pending/18-muse-prompts-execute-in-a-step.md

Part 1 (overview, failure mode, master-prompt change design) lives in `01-overview-and-acceptance.md` in the same dir — Spec Agent 1's file. This file is self-contained for Worker 02's box: the new prompt file and the Letterly desktop template change.

---

## New prompt design

Target file: `01-prompts/27-muse-prompts/02-muse-execute-in-a-step.md` (CREATE).

The new prompt is a self-contained, paste-ready prompt for Muse. A fresh Muse session with zero prior context must be able to read it, paste the task after it, and execute correctly. No cross-file references, no assumed knowledge of the repo.

Required section outline (Worker 02 writes the full markdown following this outline):

1. **Header.** Title line, version 1.0.0, date 2026-10-09, and a purpose line: "One-shot execution prompt — turns a listed task into running work instead of a dead breakdown list."
2. **The failure mode this fixes.** "Muse's issue": the agent lists the confirmed task breakdown and then never starts executing. It must never stop at the listing.
3. **The protocol (all six points, verbatim wording for the declaration formats):**
   1. Confirmed task breakdown listing FIRST, without starting work. List all tasks (Task-01, Task-02, …) with `Understood: [YES]` per task, then start executing — same turn, no pause.
   2. Immediately after the breakdown, explicitly declare execution state, answering "Are you running or not?" with `RUNNING — <task list> — ETA ~<time>` including a time approximation to complete the task.
   3. Execute immediately with multiple concurrent agents (A=2, H=2) on disjoint file boxes — same turn, no waiting for a user signal.
   4. Every 5 minutes during execution, ping with a task status update: current task, done/total, elapsed vs ETA, blockers.
   5. Task completion includes committing and pushing to Git (atomic commit, immediate push, hyphen-format message).
   6. Final output: both the enhanced Literally prompt and the Muse prompt as MD code blocks for copying to Literally.
4. **Fixed actionable items embedded in the prompt** (verbatim):
   1. `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/) first`
   2. `2. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String`
   3. `3. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well`
   4. `4. Use `gitmap` AI agents to enter data`
   5. `5. Task completion includes committing and pushing to Git`
5. **Title convention.** When formatting a task for the user's app: write `# <task title>` FIRST, then the line "high priority instruction, non-negotiable task", then a `slug: <task-slug>` line (lowercase-hyphenated slug).
6. **Note.** The user's phone app is called "Literally"; the repo's template family is named "Letterly" — same format. No repo relabeling (naming mapping only, per the plan assumptions).

---

## Letterly desktop template change design

Target file: `01-prompts/22-letterly/02-desktop-letterly.md` (EDIT, minimal diff — only the changes below).

### Change 1 — instruction 2

OLD (replace):
`Structure the output starting immediately with `# High Priority Instruction`.`

NEW (exact replacement):
`Structure the output starting with `# <task title>`, then `# High Priority Instruction — non-negotiable task`, then a `slug: <task-slug>` line.`

### Change 2 — instruction 4 fixed items

Instruction 4's fixed list becomes items 1–5, followed by the extracted items from the input text (renumbered from 6). Items 1–3 stay VERBATIM — character for character:

- `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/) first`
- `2. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String`
- `3. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well`
- new fixed item 4: `Use `gitmap` AI agents to enter data`
- new fixed item 5: `Task completion includes committing and pushing to Git`

### Change 3 — Output Format section

Mirrors the new header order. The full Output Format block becomes:

```
# <task title>

# High Priority Instruction — non-negotiable task

slug: <task-slug>

${Input Text Verbatim}

# Actionable Items Must Follow Non-Negotiable

1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/) first
2. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String
3. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well
4. Use `gitmap` AI agents to enter data
5. Task completion includes committing and pushing to Git
6. [extracted actionable technical directive from input]

## Must follow and spawn agent using

[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)

## Additional Instructions

- [/plan](slashCommand;plan) first before doing the work to reduce the credits.
- [/learn](slashCommand;learn) from [gitmap](file;.agents/skills/gitmap) skill to leverage GitMap high-speed search, toolchain discovery, and caching.
- Only add the relative paths, never add the absolute path during your work; this should be respected on the release page and in release notes as well.
```

Everything else in the file (instructions 1, 3, 5–8, the `${Input Text Verbatim}` definition) stays unchanged.

---

## Acceptance criteria

- [ ] AC2-1: `01-prompts/27-muse-prompts/02-muse-execute-in-a-step.md` exists, is self-contained (paste-ready, no prior context needed), and covers all six protocol points.
- [ ] AC2-2: The new prompt includes the title-before-header convention (`# <task title>` before the "high priority instruction, non-negotiable task" line) plus the `slug: <task-slug>` line.
- [ ] AC2-3: The new prompt embeds all five fixed actionable items verbatim and carries the Literally↔Letterly mapping note.
- [ ] AC3-1: `01-prompts/22-letterly/02-desktop-letterly.md` header order is `# <task title>`, then `# High Priority Instruction — non-negotiable task`, then the `slug: <task-slug>` line.
- [ ] AC3-2: Template items 1–3 are verbatim unchanged; new fixed items 4 (gitmap AI agents for data entry) and 5 (commit+push completion) are present.
- [ ] AC3-3: Zero absolute paths in new/modified content; all filenames lowercase; no `rg`/`grep`/`git grep`/`Select-String` used.

## Out of scope

- Cursor mirror `01-prompts/23-cursor-prompts/02-desktop-letterly-cursor.md` — follow-up, not this run.
- Index updates (`01-prompts/27-muse-prompts/readme.md`, `01-prompts/readme.md`, `.ai-memory/prompts.md`, `02-spec/21-app/readme.md`, `.ai-memory/plans/readme.md`) — lead-owned in Wave 3.
- Master prompt edit (`01-muse-master-prompt.md`) — Spec Agent 1's area (spec part 1 + Worker 01).
