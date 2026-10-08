# 20 Muse prompts execute-in-a-step update — COMPLETED

**Version:** 1.0.0
**Completed:** 2026-10-09
**Status:** Completed
**Parent Plan:** `.ai-memory/plans/pending/18-muse-prompts-execute-in-a-step.md` (kept for traceability)
**Canonical Spec:** `02-spec/21-app/15-muse-prompts-execute-in-a-step/`
**Subtasks:** `.ai-memory/plans/subtasks/18-muse-prompts-execute-in-a-step/`

---

## What was done

Fixed the observed failure mode "Muse lists the task but never starts" (`LISTING-WITHOUT-STARTING`) across three deliverables, executed via the V6 pipeline (research wave → spec wave → execution wave, A=2/H=2, disjoint file boxes), and pushed as one atomic commit.

1. **Master prompt** (`01-prompts/27-muse-prompts/01-muse-master-prompt.md`, still 6.0.0): new **Step 2.6** in Section 4 — breakdown lists tasks WITHOUT starting work first; immediately after, the agent answers "Are you running or not?" with `RUNNING — <tasks> — ETA ~<time>` (time approximation with estimate math); every 5 minutes a status ping (current task, done/total, elapsed vs ETA, blockers); listing-but-never-starting is a named protocol violation. One line added to the Section 5 checklist. Minimal diff, no reflow.
2. **New prompt** (`01-prompts/27-muse-prompts/02-muse-execute-in-a-step.md`, v1.0.0): self-contained execute-in-a-step prompt — breakdown-first, RUNNING+ETA declaration, A=2/H=2 disjoint execution, 5-minute pings, commit+push completion, final output of both prompts as MD code blocks for copying to Literally; embeds the five fixed actionable items and the `# <task title>` → `# High Priority Instruction — non-negotiable task` → `slug: <task-slug>` title convention; notes the Literally↔Letterly app-name mapping (no repo relabeling).
3. **Letterly desktop template** (`01-prompts/22-letterly/02-desktop-letterly.md`): output now starts with `# <task title>`, then `# High Priority Instruction — non-negotiable task`, then `slug: <task-slug>`; actionable items 1–3 kept verbatim, new fixed item 4 (`Use `gitmap` AI agents to enter data`) and item 5 (`Task completion includes committing and pushing to Git`), extracted items renumbered from 6.

## Indexes updated

- `01-prompts/27-muse-prompts/readme.md` — catalog row + tree for the new prompt.
- `01-prompts/readme.md` — category **27** registered (table row + tree; was stale at 26).
- `.ai-memory/prompts.md` — Prompts Matrix row for the new prompt.
- `02-spec/21-app/readme.md` — registry row for `15-muse-prompts-execute-in-a-step/`.
- `.ai-memory/plans/readme.md` — pending bullet (this register entry on completion).

## Verification (all before push)

- Relative-path linter: PASS — no absolute paths or `file:///` URIs across 3641 tracked files.
- Prompts-loaded linter: PASS — index in sync with `01-prompts/` (258 prompts on disk).
- Secrets gate: `check-forbidden-strings.py` PASS; `gitmap aum search` secret-regex on all touched paths: 0 hits.
- Worker reports verified independently against the diff (Worker 01: Step 2.6 at lines 292–304, checklist line 339, version still 6.0.0; Worker 02: new 60-line prompt, template header/title/slug/items verified by direct read).
- One atomic `gitmap cpf` push to main; zero intermediate commits.

## Decisions logged

- No version bump on the master prompt (stays 6.0.0) — user asked to modify/improve, not release (R10).
- Cursor mirror `01-prompts/23-cursor-prompts/02-desktop-letterly-cursor.md` untouched (follow-up).
- Pending plan + subtask files KEPT (not deleted per skill consolidation) — user's strict never-delete rule outranks the skill's delete step; this completed file is the consolidation record.
