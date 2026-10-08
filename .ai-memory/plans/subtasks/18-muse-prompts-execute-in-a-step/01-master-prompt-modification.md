# Subtask 01 — Master prompt modification

**Task slug:** `18-muse-prompts-execute-in-a-step`
**Worker:** Worker 01
**Wave:** 2 (execution)
**Design authority:** `02-spec/21-app/15-muse-prompts-execute-in-a-step/01-overview-and-acceptance.md` — implement it, do not redesign it.

---

## Objective

Modify the Muse master prompt to fix the named protocol violation `LISTING-WITHOUT-STARTING` (agent prints the confirmed task breakdown but never starts executing). Add an explicit execution-state declaration step ("Are you running or not?" → `RUNNING` + ETA) with 5-minute status pings, plus one Section 5 checklist line. Minimal diff; version stays 6.0.0.

---

## Owned file

`01-prompts/27-muse-prompts/01-muse-master-prompt.md` — and ONLY this file. Never touch any other file.

## Hard boundaries

- NEVER run any git command (no git status/diff/add/commit/push — lead owns git).
- Search ONLY via gitmap (`gitmap aum search`, `gitmap find`, `gitmap cat`). TOTAL BAN on rg, ripgrep, grep, git grep, Select-String, findstr.
- All filenames lowercase. Zero absolute paths in file content (relative repo paths only).
- Minimal diff: no reflow of untouched lines, no formatting churn.

---

## Edit instructions

1. Read Section 4 of the owned file to confirm current anchors (verified 2026-10-09):
   - `### Step 2.5 — SQLite Task DB Initialization & Ledger Preflight` at line 284.
   - `### Step 3 — Multi-agent execution (mandatory for multi-part work)` at line 292.
   - Section 5 checklist line `- [ ] Confirmed task breakdown shown FIRST (Section 4, Step 2)` at line 324.
   - Version header `> Prompt Version: 6.0.0` at line 20 (must remain untouched).

2. Insert the Step 2.6 block AFTER the Step 2.5 subsection and BEFORE the Step 3 heading. Use the draft wording from `02-spec/21-app/15-muse-prompts-execute-in-a-step/01-overview-and-acceptance.md` (section "Master prompt change design"). You may tighten wording but must NOT change semantics: keep 2.6.1 (list without starting; tracking-only tool call), 2.6.2 (same-turn `RUNNING — Task-NN list — ETA ~<time>` declaration with estimate math), 2.6.3 (5-minute pings: current Task-NN, completed/total, elapsed vs ETA, blockers), 2.6.4 (listing-without-starting is the named violation `LISTING-WITHOUT-STARTING`).

3. In Section 5, add exactly one checklist line directly after the "Confirmed task breakdown shown FIRST" line:

   `- [ ] RUNNING declaration with ETA printed after breakdown + 5-minute status pings during execution (Section 4, Step 2.6)`

4. Do NOT bump the version (stays 6.0.0). Do NOT touch the Cursor mirror, any readme, or any index file.

---

## Verification steps (run all before reporting)

1. Re-read the edited Section 4 region and the Section 5 checklist to confirm the insertion sits between Step 2.5 and Step 3 and the wording is intact.
2. `gitmap aum search "Are you running or not" 01-prompts/27-muse-prompts` must return a hit in the owned file.
3. `gitmap aum search "LISTING-WITHOUT-STARTING" 01-prompts/27-muse-prompts` must return a hit in the owned file.
4. `gitmap aum search "5-minute" 01-prompts/27-muse-prompts/01-muse-master-prompt.md` must return a hit.
5. Confirm no absolute paths were added (only relative paths like `01-prompts/27-muse-prompts/...`).
6. Confirm the version header still reads 6.0.0.

---

## Output contract

Reply with exactly this JSON shape (no extra prose):

```json
{
  "task": "18-muse-prompts-execute-in-a-step / subtask 01",
  "status": "DONE",
  "filesChanged": ["01-prompts/27-muse-prompts/01-muse-master-prompt.md"],
  "checks": ["<verification step> -> <pass/fail + evidence>", "..."],
  "acceptance": ["AC1 -> pass", "AC2 -> pass", "AC3 -> pass", "AC4 -> pass", "AC5 -> pass", "AC6 -> pass"],
  "assumptions": ["<any assumption made, or none>"],
  "blockers": ["<any blocker, or none>"]
}
```

Set `"status": "DONE"` only when all verification steps pass. If anything fails, set `"status": "BLOCKED"` and describe the blocker instead of working around it.
