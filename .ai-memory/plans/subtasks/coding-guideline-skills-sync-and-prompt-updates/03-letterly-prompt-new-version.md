# Subtask 03 — letterly prompt new version (W2)

Owner: W2. Owns `01-prompts/22-letterly/` additions only. Do NOT touch `03-ai-scripts/`.

## What to do

1. Read `01-prompts/22-letterly/readme.md` (invariants) and the closest
   existing prompt (e.g. `03-execute-n-steps-letterly.md`) as the base format.
2. Create `01-prompts/22-letterly/11-<slug>-letterly.md` (next sequence number
   11; keep the `-letterly.md` suffix). It formats voice dictation into a task
   prompt whose FIRST step is unmistakable:
   - **Step 1 — Git pull BEFORE any work:** first try `gitmap pull`; if that
     fails, fall back to regular `git pull` in the current repository/folder.
     Plain language, numbered step 1, impossible to miss. No work starts before
     the pull succeeds.
3. Keep the letterly invariants: zero conversational filler, lossless verbatim
   capture, mandatory native skill trigger, standard suffix, relative paths only.
4. Update `01-prompts/22-letterly/readme.md` index table with row 11.

## Rules

- TOTAL BAN on rg/grep/git grep/Select-String → gitmap search tools.
- New file only — never edit the 10 existing prompt files in place.

## Acceptance

- [ ] New `11-*-letterly.md` exists with pull-first as step 1 (`gitmap pull` → `git pull` fallback).
- [ ] `readme.md` index lists it as #11.
