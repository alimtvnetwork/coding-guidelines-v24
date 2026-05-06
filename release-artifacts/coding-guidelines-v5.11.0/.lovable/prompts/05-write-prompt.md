# Write Memory — Canonical v3.0.0

**Trigger:** "write memory" / "end memory" / "update memory"
**Updated:** 2026-04-27
**Supersedes:** `04-write-prompt.md` (v2.0.0)

---

## Purpose

After completing work or at the end of a session, the AI must persist everything it learned, did, and left undone — so the next AI session can pick up seamlessly with zero context loss.

> **The memory system is the project's brain.** If you did something and didn't write it down, it didn't happen. If something is pending and you didn't record it, it will be lost. Write memory as if the next AI has amnesia — because it does.

---

## Phase 1 — Audit Current State

- What was done this session? (features, fixes, refactors, every file touched)
- What is still pending? (started-not-finished, discussed-not-started, blockers)
- What was learned? (patterns, gotchas, user preferences expressed)
- What went wrong? (bugs, failed approaches, things to never repeat)

---

## Phase 2 — Update Memory Files (`.lovable/memory/`)

1. Read `.lovable/memory/index.md` first — never duplicate existing files.
2. Update existing memory files in-place; preserve all unrelated entries.
3. Create new files only when no existing file fits — naming `xx-descriptive-name.md` lowercase-hyphenated. **Immediately** add to `index.md` in the same operation.
4. Update workflow state with markers: `✅ Done` · `🔄 In Progress` · `⏳ Pending` · `🚫 Blocked — [reason/avoid]`.

---

## Phase 3 — Update Plans & Suggestions

- **Plans → `.lovable/plan.md`** (single file): update statuses, add new tasks, move fully complete items to `## Completed` section at the bottom (never delete).
- **Suggestions → `.lovable/suggestions.md`** (single file). Schema:
  - `## Active Suggestions` with Status/Priority/Description/Added.
  - `## Implemented Suggestions` with Implemented date + notes.
  - When implemented, move from Active to Implemented and add evidence reference.

---

## Phase 4 — Update Issues

- **Pending** → `.lovable/pending-issues/xx-short-description.md` with: Description / Root Cause / Steps / Attempted Solutions / Priority / Blocked By.
- **Solved** → move to `.lovable/solved-issues/xx-…md` and append: Solution / Iteration Count / Learning / What NOT to Repeat.
- **Strictly avoided** patterns → append to `.lovable/strictly-avoid.md` referencing the solved issue file.

---

## Phase 5 — Consistency Validation

- Index integrity: every file in `.lovable/memory/**` listed in `index.md`.
- Cross-reference: every `✅ Done` plan item has evidence; every actionable pending issue is in `plan.md` or `suggestions.md`; no file is in both `pending-issues/` and `solved-issues/`.
- Orphan check: no memory file without index entry; no implemented suggestion without code evidence; no solved issue missing a `## Solution` section.

### Final confirmation block

```
✅ Memory update complete.
Session Summary:
- Tasks completed: [X]
- Tasks pending: [Y]
- New memory files created: [Z]
- Issues resolved: [N] / opened: [M]
- Suggestions added: [S] / implemented: [T]

Files modified:
- [list every file touched]

Inconsistencies found and fixed: [list or "None"]
The next AI session can pick up from: [current state + next logical step]
```

---

## Additional rules (carried forward + new in v3)

9. Lovable suggestions live under `.lovable/memory/suggestions/` with an index file when they are tracked separately from `.lovable/suggestions.md`.
10. **CI/CD issues** → `.lovable/cicd-issues/xx-issue-name.md` (sequence from `01`). Maintain `.lovable/cicd-index.md` summary table. Never duplicate.
11. When a *recent* large spec is provided in chat, persist it verbatim into the file system AND record a memory pointer so the next AI can find it.
12. **Tasks the user told us to skip or avoid** → `.lovable/memory/avoid/xx-title.md` and reference from `index.md` Core or Memories.
13. **All markdown files** in `.lovable/` use lowercase-hyphenated names with numeric prefix where ordering matters.
14. **Never** create `.lovable/memories/` (with trailing `s`). The correct path is `.lovable/memory/`.
15. Restructure folders to match the canonical layout if drift is detected.

### Canonical folder layout

```
.lovable/
├── overview.md
├── strictly-avoid.md
├── plan.md
├── suggestions.md
├── prompt.md                # index of prompts
├── prompts/                 # xx-*-prompt.md
├── memory/
│   ├── index.md             # MUST list every memory file
│   ├── avoid/               # user-stated prohibitions
│   ├── workflow/
│   ├── sessions/
│   ├── constraints/
│   ├── features/
│   ├── issues/
│   ├── project/
│   ├── standards/
│   ├── done/
│   └── suggestions/
├── pending-issues/          # xx-name.md
├── solved-issues/           # xx-name.md
├── cicd-index.md
└── cicd-issues/             # xx-name.md
```

---

## Anti-corruption rules

1. Never delete history — mark done, move to completed sections.
2. Never overwrite blindly — read before write.
3. Never leave orphans — every file must be indexed; every reference must resolve.
4. Never split what should be unified — `plan.md` and `suggestions.md` each live in ONE file.
5. Never mix states — issues cannot be both pending and solved.
6. Never skip the index update.
7. Never assume the next AI knows anything.

---

*Write-memory prompt — v3.0.0 — 2026-04-27*