# Worker 02 — new execute-in-a-step prompt + Letterly desktop template

**Version:** 1.0.0
**Date:** 2026-10-09
**Status:** Ready for Wave 2 execution
**Plan:** .ai-memory/plans/pending/18-muse-prompts-execute-in-a-step.md
**Full design:** 02-spec/21-app/15-muse-prompts-execute-in-a-step/02-new-prompt-and-template-changes.md (read this first — it is the authoritative design)

## Objective

Create the new standalone Muse execute-in-a-step prompt and apply the title/slug/actionable-items changes to the Letterly desktop template. Nothing else.

## Owned files (disjoint from Worker 01 — touch nothing else)

1. CREATE: `01-prompts/27-muse-prompts/02-muse-execute-in-a-step.md`
2. EDIT: `01-prompts/22-letterly/02-desktop-letterly.md`

Hard boundaries: NEVER run any git command (no add/commit/push — the lead handles Wave 3). Search ONLY via gitmap (`gitmap aum search`, `gitmap find`, `gitmap cat`) — TOTAL BAN on rg, ripgrep, grep, git grep, Select-String. All filenames lowercase. Zero absolute paths in file content.

## Per-file instructions

### File A — `01-prompts/27-muse-prompts/02-muse-execute-in-a-step.md` (CREATE)

Write complete, paste-ready markdown following the spec section `## New prompt design` in `02-spec/21-app/15-muse-prompts-execute-in-a-step/02-new-prompt-and-template-changes.md`. It must contain, in order:

1. Header (title, version 1.0.0, date 2026-10-09) + purpose line: "One-shot execution prompt — turns a listed task into running work instead of a dead breakdown list."
2. The failure mode it fixes: Muse's issue — the agent lists the confirmed task breakdown and then never starts executing; it must never stop at the listing.
3. The protocol, all six points:
   - (1) confirmed task breakdown listing FIRST without starting work (list-then-start, same turn, `Understood: [YES]` per task);
   - (2) immediately after the breakdown, answer "Are you running or not?" with `RUNNING — <task list> — ETA ~<time>` including a time approximation;
   - (3) execute immediately with multiple concurrent agents (A=2, H=2) on disjoint file boxes, same turn;
   - (4) status ping every 5 minutes: current task, done/total, elapsed vs ETA, blockers;
   - (5) task completion includes committing and pushing to Git;
   - (6) final output: both the enhanced Literally prompt and the Muse prompt as MD code blocks for copying to Literally.
4. The five fixed actionable items VERBATIM:
   - `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/) first`
   - `2. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String`
   - `3. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well`
   - `4. Use `gitmap` AI agents to enter data`
   - `5. Task completion includes committing and pushing to Git`
5. The title convention: write `# <task title>` FIRST, then the line "high priority instruction, non-negotiable task", then a `slug: <task-slug>` line.
6. The note: the user's phone app is called "Literally"; the repo's template family is named "Letterly" — same format; no repo relabeling.

The prompt must be self-contained: a fresh Muse session with zero prior context can paste it with a task and execute correctly.

### File B — `01-prompts/22-letterly/02-desktop-letterly.md` (EDIT, minimal diff)

Apply exactly the three changes from spec section `## Letterly desktop template change design`:

1. Instruction 2 becomes: `Structure the output starting with `# <task title>`, then `# High Priority Instruction — non-negotiable task`, then a `slug: <task-slug>` line.`
2. Instruction 4's fixed list becomes items 1–5: items 1–3 VERBATIM unchanged (see the five items above), new fixed item 4 `Use `gitmap` AI agents to enter data`, new fixed item 5 `Task completion includes committing and pushing to Git`, followed by the extracted items from the input (renumbered from 6).
3. The Output Format section mirrors the new header order (see the full block in the spec).

Do not touch instructions 1, 3, 5–8, the `${Input Text Verbatim}` definition, or anything else in the file.

## Verification steps (do all)

1. Re-read both files in full after writing.
2. `gitmap aum search "non-negotiable task" 01-prompts/22-letterly/02-desktop-letterly.md` — must hit the new header line.
3. Confirm the title-before-header order in the template: `# <task title>` line comes before `# High Priority Instruction — non-negotiable task`, which comes before the `slug: <task-slug>` line.
4. Confirm items 1–3 in the template are byte-identical to the verbatim items above.
5. Confirm zero absolute paths in new/modified content (relative paths like `02-spec/...` only — never absolute filesystem paths).

## Output contract

Reply with JSON (and nothing else besides a one-line human summary):

```json
{
  "task": "02-new-prompt-and-letterly-template",
  "status": "DONE",
  "filesChanged": ["01-prompts/27-muse-prompts/02-muse-execute-in-a-step.md", "01-prompts/22-letterly/02-desktop-letterly.md"],
  "checks": ["re-read both files", "gitmap aum search non-negotiable task hit", "title-before-header order confirmed", "items 1-3 verbatim", "zero absolute paths"],
  "acceptance": ["AC2-1 new prompt exists, self-contained, six protocol points", "AC2-2 title-before-header + slug line", "AC2-3 five fixed items verbatim + Literally/Letterly note", "AC3-1 template header order", "AC3-2 items 1-3 verbatim, items 4-5 present", "AC3-3 zero absolute paths, lowercase, no banned search tools"],
  "assumptions": ["Literally == Letterly template family; no repo relabeling"],
  "blockers": []
}
```
