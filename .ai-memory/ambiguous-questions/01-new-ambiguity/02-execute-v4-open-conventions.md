# Execute V4 Open Conventions: Skill File Casing, Screenshot Folder, and Boolean Prefixes

Slug: execute-v4-open-conventions
Status: open
Raised: 2026-10-01
Blocking: the sibling-repo rollout of V4 (follow-up 3 in `.ai-memory/plans/pending/15-execute-parent-task-v4-antigravity.md`)

## Question 1: `skill.md` or `SKILL.md`?

The repo renames every file to lowercase (commit `1c91f6cf` and section 4 of `agents.md`), so skills live at `.agents/skills/<name>/skill.md` and `.cursor/skills/<name>/skill.md`. The Antigravity and Cursor skill docs name the file `SKILL.md`. Windows and default macOS file systems ignore case, but Linux file systems do not, so a lowercase skill may not be discovered there.

Options:

- **A:** Keep lowercase everywhere, and confirm skill discovery on a Linux checkout before rolling V4 out.
- **B:** Allow `SKILL.md` as the one documented exception to the lowercase rule.
- **C:** Keep lowercase in this repo, and let the sync script rename the file to `SKILL.md` in repos that run on Linux.

Impact if guessed wrong: V4 and every other skill can silently fail to load as a slash command on case-sensitive systems.

## Question 2: Where do screenshots go?

V3 and V4 save screenshots to `assets/screenshots/<slug>-<NN>.png`; that folder exists and holds one file from plan 14. Hard Rule 13 requires `assets/<NN-folder>/<NN-file>.<ext>` with two-digit prefixes.

Options:

- **A:** Keep `assets/screenshots/` and record it as an exception to Hard Rule 13.
- **B:** Rename the folder to a numbered one and update V3, V4, and the image link in plan 14.

Impact if guessed wrong: either screenshot links break, or every run keeps violating the asset rule.

## Question 3: Which boolean prefixes are allowed?

The canonical guidelines and V3 allow only `is` and `has`. The hand-written Quick Rules at the top of `.cursorrules` also allow `can`, `should`, `was`, and `will`. V4 defers to `.ai-memory/coding-guidelines.md` and does not restate the rule.

Options:

- **A:** `is` and `has` only, and update the `.cursorrules` Quick Rules to match.
- **B:** Allow the wider list, and update the canonical guidelines and the mirrors.

Impact if guessed wrong: agents that read `.cursorrules` and agents that read the canonical file name booleans differently, and reviews reject code that the other source allows.

## Next Steps for User Review

Pick one option per question. The answers then go into the canonical guidelines (and their mirrors through `scripts/sync-guidelines.mjs`) and into V4, and this note moves out of `01-new-ambiguity/`.
