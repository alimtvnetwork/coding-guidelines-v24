# 15 — Muse Prompts Execute-in-a-Step (Part 1: Overview + Master Prompt Change)

**Version:** 1.0.0
**Date:** 2026-10-09
**Status:** Spec
**Parent plan:** `.ai-memory/plans/pending/18-muse-prompts-execute-in-a-step.md`
**Subtask (Worker 01):** `.ai-memory/plans/subtasks/18-muse-prompts-execute-in-a-step/01-master-prompt-modification.md`
**Companion spec:** `02-spec/21-app/15-muse-prompts-execute-in-a-step/02-new-prompt-and-template-changes.md`

This spec part covers the master-prompt modification only. The new standalone execute-in-a-step prompt and the Letterly template changes are covered in spec part 2 (Spec Agent 2's area).

---

## The observed failure mode

Named protocol violation: **LISTING-WITHOUT-STARTING**.

The agent prints the confirmed task breakdown (Task-01 … Task-NN, each `Understood: [YES]`) and then the turn ends without any execution beginning. The breakdown exists, but no work-doing tool call is ever made, no subagent is spawned, and the user is left asking "Are you running or not?"

The existing master prompt (`01-prompts/27-muse-prompts/01-muse-master-prompt.md`, V6) mandates breakdown-first and same-turn tool chaining in Section 4, but it never requires the agent to:

1. explicitly declare its execution state right after the breakdown (answer "Are you running or not?" in the open), or
2. give a time approximation (ETA) for the whole task, or
3. ping back with status updates at a fixed cadence while the work runs.

This spec fixes that gap with one new step and one checklist line. No semantics of the existing protocol are removed or weakened.

---

## Master prompt change design

Target file: `01-prompts/27-muse-prompts/01-muse-master-prompt.md` (version stays 6.0.0 — no version bump this run).

Insert a new **Step 2.6** in Section 4, placed immediately after the Step 2.5 subsection (`### Step 2.5 — SQLite Task DB Initialization & Ledger Preflight`, currently at line 284) and immediately before Step 3 (`### Step 3 — Multi-agent execution (mandatory for multi-part work)`, currently at line 292).

### Draft wording for Step 2.6 (Worker 01 may tighten wording, never semantics)

```md
### Step 2.6 — Execution-state declaration ("Are you running or not?") and 5-minute status pings

2.6.1 — The breakdown lists tasks WITHOUT starting work first: the listing completes before any work-doing tool call; the same-turn tool call only initializes tracking (SQLite task DB / ledger / preflight checks).

2.6.2 — Execution-state declaration: immediately after the breakdown, in the same turn, print an explicit line answering "Are you running or not?" in the form:

`RUNNING — Task-01, Task-02, Task-03 — ETA ~45 min`

Include a time approximation for the whole task. Show the estimate math, e.g. research ~10 min + execution ~25 min + verification and push ~10 min = ~45 min.

2.6.3 — Every 5 minutes during execution, ping a status update: current Task-NN, completed/total, elapsed vs ETA, and any blockers.

2.6.4 — Listing without starting is NOT stopping: the turn that shows the breakdown MUST also start execution (mandatory same-turn chaining). Listing-but-never-starting is a named protocol violation: `LISTING-WITHOUT-STARTING`.
```

### Section 5 checklist addition

Add exactly one new checklist line in Section 5, directly after the existing line:

`- [ ] Confirmed task breakdown shown FIRST (Section 4, Step 2)` (currently at line 324)

New line:

`- [ ] RUNNING declaration with ETA printed after breakdown + 5-minute status pings during execution (Section 4, Step 2.6)`

---

## Acceptance criteria

- [ ] AC1: The master prompt contains an explicit execution-state declaration step that answers "Are you running or not?" with `RUNNING` plus a time approximation (ETA) for the whole task.
- [ ] AC2: The master prompt contains a 5-minute status-ping rule (current Task-NN, completed/total, elapsed vs ETA, blockers).
- [ ] AC3: The diff touches only the new Step 2.6 block and the one new Section 5 checklist line — minimal diff, no reflow of untouched lines.
- [ ] AC4: The version header still reads 6.0.0 (no version bump).
- [ ] AC5: Zero absolute paths in new/modified content; all filenames lowercase.
- [ ] AC6: The named protocol violation `LISTING-WITHOUT-STARTING` is present and defined.

---

## Out of scope

- Version bump of the master prompt (stays 6.0.0).
- Cursor mirror `01-prompts/23-cursor-prompts/02-desktop-letterly-cursor.md` — not touched this run.
- Index updates (27 readme, 01-prompts readme, prompts.md matrix, 21-app readme registry, plans readme) — lead-owned in Wave 3.
- The new standalone execute-in-a-step prompt file (`01-prompts/27-muse-prompts/02-muse-execute-in-a-step.md`) and the Letterly template changes — spec part 2, Spec Agent 2's area.
