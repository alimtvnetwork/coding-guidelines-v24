# Write Memory

**Version:** 1.0.0  
**Updated:** 2026-04-19

> **Purpose:** After completing work or at the end of a session, the AI must persist everything it learned, did, and left undone — so the next AI session can pick up seamlessly with zero context loss.
>
> **When to run:** At the end of every session, after completing a task batch, or when explicitly asked to "update memory", "write memory", or "end memory".

---

## Table of Contents

1. [Core Principle](#core-principle)
2. [Phase 1 — Audit Current State](#phase-1--audit-current-state)
3. [Phase 2 — Update Memory Files](#phase-2--update-memory-files)
4. [Phase 3 — Update Plans & Suggestions](#phase-3--update-plans--suggestions)
5. [Phase 4 — Update Issues](#phase-4--update-issues)
6. [Phase 5 — Consistency Validation](#phase-5--consistency-validation)
7. [File Naming & Structure Rules](#file-naming--structure-rules)
8. [Anti-Corruption Rules](#anti-corruption-rules)

---

## Core Principle

> **The memory system is the project's brain.** If you did something and didn't write it down, it didn't happen. If something is pending and you didn't record it, it will be lost. Write memory as if the next AI has amnesia — because it does.

---

## Phase 1 — Audit Current State

Before writing anything, take inventory:

- **What was done this session?** Every task completed, every file changed, every decision made.
- **What is still pending?** Started-but-unfinished, discussed-but-unstarted, blocked items.
- **What was learned?** New patterns, gotchas, user preferences (explicit or implicit).
- **What went wrong?** Bugs, failed approaches, things to never repeat.

---

## Phase 2 — Update Memory Files

### Target: `.lovable/memory/`

1. **Read the current index** — `.lovable/memory/index.md`. Avoid duplicates.
2. **Update existing memory files** — append to the right section, mark items `[x]` / `✅`, never truncate.
3. **Create new memory files** if needed: `.lovable/memory/<topic>/NN-descriptive-name.md` and update `index.md` in the same operation.
4. **Update workflow state** in `.lovable/memory/workflow/` using these markers:

| Status | Marker |
|--------|--------|
| Done | `✅ Done` |
| In Progress | `🔄 In Progress` |
| Pending | `⏳ Pending` |
| Blocked | `🚫 Blocked — [reason]` |
| Avoid or Skip | `🚫 Blocked — [avoid]` |

---

## Phase 3 — Update Plans & Suggestions

### 3A — Plans (`.lovable/plan.md`)

- Update task statuses (done / in progress / pending).
- Add any new tasks discovered.
- Move fully-complete items to a `## Completed` section in the **same file** (do not delete).
- Single source of truth for the roadmap.

### 3B — Suggestions (`.lovable/suggestions.md`)

Maintain a **single file** with this structure:

```markdown
## Active Suggestions
### [Title]
- **Status:** Pending | In Review | Approved | Rejected
- **Priority:** High | Medium | Low
- **Description:** What and why
- **Added:** [date or session reference]

## Implemented Suggestions
### [Title]
- **Implemented:** [date or session reference]
- **Notes:** Any relevant details about the implementation
```

When a suggestion is implemented, **move** it from Active → Implemented and add notes.

---

## Phase 4 — Update Issues

### 4A — Pending Issues (`.lovable/pending-issues/NN-short-description.md`)

```markdown
# [Issue Title]

## Description
What is broken or unexpected.

## Root Cause
Why it happens. If unknown: "Under investigation."

## Steps to Reproduce
1. Step one
2. Step two
3. Expected vs actual

## Attempted Solutions
- [ ] Approach 1 — [result]
- [ ] Approach 2 — [result]

## Priority
High | Medium | Low

## Blocked By (if applicable)
What dependency or decision is needed.
```

### 4B — Solved Issues (`.lovable/solved-issues/NN-short-description.md`)

When resolved, **move** the file from `pending-issues/` → `solved-issues/` and append:

```markdown
## Solution
What fixed it.

## Iteration Count
How many attempts.

## Learning
What we learned.

## What NOT to Repeat
Specific anti-patterns to avoid.
```

### 4C — Strictly Avoided Patterns (`.lovable/strictly-avoid.md`)

If a solved issue revealed a pattern that must **never** be used again:

```markdown
- **[Pattern Name]:** [Why it's forbidden]. See: `.lovable/solved-issues/NN-filename.md`
```

---

## Phase 5 — Consistency Validation

### 5.1 — Index Integrity
Every file in `.lovable/memory/` (including subfolders) must be listed in `index.md`.

### 5.2 — Cross-Reference Check
- Every `✅ Done` task in `plan.md` should have evidence (memory update / solved issue / code change).
- Every `pending-issues/` item should be reflected in `plan.md` or `suggestions.md` if actionable.
- No file should exist in both `pending-issues/` and `solved-issues/`.

### 5.3 — Orphan Check
- No memory file without an index entry.
- No "Implemented" suggestion without code evidence.
- No solved issue without a `## Solution` section.

### 5.4 — Final Confirmation

Respond with:

```
✅ Memory update complete.

Session Summary:
- Tasks completed: [X]
- Tasks pending: [Y]
- New memory files created: [Z]
- Issues resolved: [N]
- Issues opened: [M]
- Suggestions added: [S]
- Suggestions implemented: [T]

Files modified:
- [list every file touched]

Inconsistencies found and fixed:
- [list any, or "None"]

The next AI session can pick up from: [current state + next logical step]
```

---

## File Naming & Structure Rules

| Rule | Example |
|------|---------|
| All files use numeric prefix (where applicable) | `01-auth-flow.md`, `02-api-design.md` |
| Lowercase, hyphen-separated | `03-error-handling.md` ✅ / `03_Error_Handling.md` ❌ |
| Plans → single file | `.lovable/plan.md` |
| Suggestions → single file | `.lovable/suggestions.md` |
| Pending issues → one file per issue | `.lovable/pending-issues/01-login-crash.md` |
| Solved issues → one file per issue | `.lovable/solved-issues/01-login-crash.md` |
| Memory → grouped by topic | `.lovable/memory/workflow/`, `.lovable/memory/decisions/` |
| Completed plans/suggestions → `## Completed` section in same file | Do NOT create separate `completed/` folders |

### Folder Structure Reference

```
.lovable/
├── overview.md                  # Project summary
├── strictly-avoid.md            # Hard prohibitions (single file)
├── user-preferences             # Communication style
├── plan.md                      # Active roadmap (single file)
├── suggestions.md               # All suggestions (single file)
├── prompt.md                    # Prompt index
├── memory/
│   ├── index.md                 # Index of all memory files
│   ├── workflow/                # Workflow state and progress
│   ├── decisions/               # Architectural decisions
│   └── [topic]/                 # Other topic-specific memory
├── pending-issues/              # Unresolved bugs/issues
├── solved-issues/               # Resolved bugs/issues
└── prompts/                     # AI prompts (read, write, etc.)
```

> ⚠️ **NEVER** create `.lovable/memories/` (with trailing `s`). The correct path is `.lovable/memory/`.

---

## Anti-Corruption Rules

1. **Never delete history** — mark items as done; move to completed sections.
2. **Never overwrite blindly** — read before writing; preserve existing content.
3. **Never leave orphans** — every file indexed; every reference resolves.
4. **Never split what should be unified** — plans & suggestions each live in ONE file.
5. **Never mix states** — an issue cannot be both pending and solved.
6. **Never skip the index update** — if you create a memory file, update `index.md` in the same operation.
7. **Never assume the next AI knows anything** — write as if explaining to a stranger.
8. **Tasks marked to skip or avoid** → put a memory note in `.lovable/memory/avoid/` and reference from `strictly-avoid.md`.

---

*Write-memory prompt — v1.0.0 — 2026-04-19. Must stay in sync with `.lovable/prompts/01-read-prompt.md` (AI Onboarding Protocol).*
