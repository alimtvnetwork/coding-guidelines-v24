# 18 Muse prompts execute-in-a-step update

**Version:** 1.0.0
**Updated:** 2026-10-09
**Status:** Pending
**Verification date:** 2026-10-09

This plan is the work. Implement it in order. Do not redesign it.

---

## User Request (Verbatim)

> Do a pull in the coding guideline. Modify and improve the Muse prompt, and create an execute-in-a-step type prompt for Muse using Muse's issues. Muse lists out the task but does not start the task. It should always say, "Are you running or not?" and provide a time approximation to complete the task. Every five minutes, it should ping and update the task status. Update the Literally prompt by writing the title of the task before "high priority instruction, non-negotiable task" and include the task slug. Add new fixed points into the actionable items, such as using `gitmap` AI agents to enter data. Ensure task completion includes committing and pushing to Git. Enhance the Literally prompt and create a Muse prompt in the Muse folder. Ask Muse to output both prompts as MD code blocks for copying to Literally.

Actionable items 1–9 and the execute-parent-task-with-n-steps-v6 invocation requirement are captured in the spec (`02-spec/21-app/15-muse-prompts-execute-in-a-step/`).

**Parent Plan:** `.ai-memory/plans/pending/18-muse-prompts-execute-in-a-step.md` (this file)
**Canonical Spec:** `02-spec/21-app/15-muse-prompts-execute-in-a-step/`
**Subtasks:** `.ai-memory/plans/subtasks/18-muse-prompts-execute-in-a-step/`

---

## Context

The user dictates tasks into their phone app ("Literally"; the repo's template family is named "Letterly" — same format, naming mismatch logged as a non-blocking assumption, no repo relabeling without the user's word). The pasted format matches `01-prompts/22-letterly/02-desktop-letterly.md`.

Observed failure mode ("Muse's issues"): the agent lists the confirmed task breakdown but never starts executing. The fix codified by this task:

1. Breakdown lists tasks WITHOUT starting work first (list first, then start — same turn).
2. Immediately after the breakdown, the agent explicitly declares execution state, answering "Are you running or not?" with `RUNNING` plus a time approximation (ETA) for the whole task.
3. Every 5 minutes during execution, the agent pings with a task status update.
4. Task completion includes committing and pushing to Git (already the standing rule; now a fixed, visible checklist item).

Who is affected: every future Muse session booted from the master prompt, and every task formatted through the desktop Letterly template.
Desired result: (a) master prompt gains the execution-state declaration + 5-minute ping protocol, (b) a new standalone `02-muse-execute-in-a-step.md` prompt in the Muse folder, (c) the desktop Letterly template gains title-before-header, task slug, and new fixed actionable items (gitmap AI agents for data entry; commit+push completion).
Done when: the acceptance criteria below all pass and the single atomic `gitmap cpf` push lands on main.

---

## Current state (verified 2026-10-09)

| Piece | Path | What it does today |
|---|---|---|
| Master prompt | `01-prompts/27-muse-prompts/01-muse-master-prompt.md` | V6, 420 lines. Section 4 mandates breakdown-first + same-turn tool chaining, but never requires an explicit "RUNNING + ETA" declaration or periodic status pings |
| Muse folder | `01-prompts/27-muse-prompts/` | Only `01-muse-master-prompt.md` + `readme.md`. No execute-in-a-step prompt exists |
| Letterly desktop template | `01-prompts/22-letterly/02-desktop-letterly.md` | Output starts at `# High Priority Instruction`; items 1–3 fixed (spec-first, gitmap-only search, relative paths); no task-title line, no slug, no gitmap-agents or commit+push fixed items |
| Spec dir (next free) | `02-spec/21-app/` | Entries 01–14 used; next free sequential is `15` |
| Pending plan (next free) | `.ai-memory/plans/pending/` | `02,04,05,09,11,13,15,16,17` used; next free is `18` |
| Subtasks (next free) | `.ai-memory/plans/subtasks/` | Max `17`; next free is `18` |
| Prompt catalog | `01-prompts/27-muse-prompts/readme.md` | Table row + tree per file; currently 1 row (master prompt, 6.0.0) |
| Categories index | `01-prompts/readme.md` | STALE: ends at `26-gitmap/`; category 27 never registered (fix while touching prompts) |
| Prompts matrix | `.ai-memory/prompts.md` | One row per prompt file; master prompt at lines 213–214 |

---

## What changes (file table)

| # | Path | Change |
|---|---|---|
| 1 | `01-prompts/27-muse-prompts/01-muse-master-prompt.md` | Add `Step 2.6` in Section 4: execution-state declaration ("Are you running or not?" → `RUNNING` + time approximation) and 5-minute status pings; add one line to the Section 5 checklist. Minimal diff — no reflow, version stays 6.0.0 |
| 2 | `01-prompts/27-muse-prompts/02-muse-execute-in-a-step.md` | NEW standalone execute-in-a-step prompt: breakdown-first, RUNNING+ETA declaration, A=2/H=2 execution, 5-min pings, commit+push completion, final output of both prompts as MD code blocks for copying to Literally |
| 3 | `01-prompts/22-letterly/02-desktop-letterly.md` | Template instructions + output format: `# <task title>` line BEFORE `# High Priority Instruction — non-negotiable task`, `slug: <task-slug>` line, new fixed items 4 (`Use `gitmap` AI agents to enter data`) and 5 (`Task completion includes committing and pushing to Git`) |
| 4 | `02-spec/21-app/15-muse-prompts-execute-in-a-step/01-overview-and-acceptance.md` | Spec part 1 (overview, failure mode, master-prompt change design, acceptance) |
| 5 | `02-spec/21-app/15-muse-prompts-execute-in-a-step/02-new-prompt-and-template-changes.md` | Spec part 2 (new prompt design, letterly template change design, acceptance) |
| 6 | `.ai-memory/plans/subtasks/18-muse-prompts-execute-in-a-step/01-master-prompt-modification.md` | Worker 01 subtask |
| 7 | `.ai-memory/plans/subtasks/18-muse-prompts-execute-in-a-step/02-new-prompt-and-letterly-template.md` | Worker 02 subtask |

Lead-owned (not in any worker box): this plan file, the run ledger, and ALL index updates — `01-prompts/27-muse-prompts/readme.md` (new row + tree line), `01-prompts/readme.md` (register category 27), `.ai-memory/prompts.md` (new matrix row), `02-spec/21-app/readme.md` (register `15-…`), `.ai-memory/plans/readme.md` (pending bullet now, recent-tasks register on completion).

---

## Execution waves

- **Wave 0 (done):** Research 01 (catalog/numbering conventions) + Research 02 (Literally template location, master-prompt references). Both delivered.
- **Wave 1 (spec, A=2 `self`):** Spec Agent 1 → files 4 + 6. Spec Agent 2 → files 5 + 7. Strictly disjoint — no shared files.
- **Wave 2 (execution, A=2 `self`):** Worker 01 → file 1. Worker 02 → files 2 + 3. Strictly disjoint — no shared files. Workers NEVER run git commands.
- **Wave 3 (lead):** index updates → targeted checks (relative-path linter, prompts-loaded linter on touched files) → secrets gate (`linter-scripts/check-forbidden-strings.py` + `gitmap aum search` secret regex) → ONE atomic `gitmap cpf "muse-prompts - execute-in-a-step prompt and letterly template update"` → push. Then final reply with both prompts as MD code blocks.

---

## Acceptance criteria

- [ ] AC1: `01-muse-master-prompt.md` contains an explicit execution-state declaration step answering "Are you running or not?" with RUNNING + time approximation, and a 5-minute status-ping rule; diff touches only the new step + one checklist line; version header still 6.0.0.
- [ ] AC2: `02-muse-execute-in-a-step.md` exists in `01-prompts/27-muse-prompts/`, is self-contained (no prior context needed), and covers: breakdown-first, RUNNING+ETA declaration, A=2/H=2 disjoint execution, 5-minute pings, commit+push completion, and final output of both prompts as MD code blocks for copying to Literally.
- [ ] AC3: `02-desktop-letterly.md` output format starts with `# <task title>`, then `# High Priority Instruction — non-negotiable task`, then a `slug: <task-slug>` line; actionable items list has new fixed items 4 (gitmap AI agents for data entry) and 5 (commit+push completion); items 1–3 verbatim unchanged.
- [ ] AC4: All five index files updated (27 readme, 01-prompts readme incl. category 27 row, prompts.md matrix row, 21-app readme registry row, plans readme pending bullet).
- [ ] AC5: Targeted checks exit 0 on touched files; secrets gate clean; exactly ONE atomic commit pushed to main (hyphen-format `gitmap cpf` message, no colons).
- [ ] AC6: Zero absolute paths in new/modified content; all filenames lowercase; no `rg`/`grep`/`git grep`/`Select-String` used anywhere in the run.

---

## Assumptions

- "Literally" (user's app) == "Letterly" (repo template family). No repo relabeling; the new prompt notes the mapping.
- No version bump (stays 6.0.0) — user asked to modify/improve, not release (R10).
- The Cursor mirror `01-prompts/23-cursor-prompts/02-desktop-letterly-cursor.md` is NOT touched this run.

## Conflicts

None. User preamble (top-instruction priority) outranks all skill defaults; the skill's zero-intermediate-commit rule and the user's commit+push completion requirement agree.

## Follow-ups (out of scope — do not fix this run)

- Mirror the letterly template changes to `01-prompts/23-cursor-prompts/02-desktop-letterly-cursor.md`.
- `02-spec/21-app/readme.md` Contents table omits 4 existing dirs; `.ai-memory/plans/readme.md` Pending list omits pending `05-…` and `17-…` (pre-existing staleness, flagged by Research 01).
- Confirm with the user whether the repo's "Letterly" naming should ever be relabeled to "Literally".
