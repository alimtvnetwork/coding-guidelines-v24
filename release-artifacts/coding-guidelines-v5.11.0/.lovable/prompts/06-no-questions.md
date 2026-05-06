# 06 — No-Questions Mode

**Version:** 1.0.0
**Created:** 2026-05-04
**Status:** Active (window-scoped, 40 tasks)
**Tags:** `no question`, `not ques for 40`

---

## Purpose

Suspend all user-facing question prompts for a fixed window of **40 tasks**.
Instead of asking, the AI infers the most reasonable interpretation, logs
the ambiguity for later review, and continues without interruption.

## Scope

- Active starting **task #1** after this prompt was registered.
- Auto-expires when the counter in
  `.lovable/question-and-ambiguity/task-counter.md` reaches **40**.
- Applies to ALL ambiguity types: design, scope, naming, behavior,
  prioritization, tooling.
- Hard exceptions where a question is still allowed:
  1. Destructive actions that cannot be inferred safely (e.g. "delete the
     entire spec/ folder").
  2. Required secret/credential inputs that have no reasonable default.
  3. Explicit user instruction inside the same message saying "ask me".

## Rules

### 1. Logging
- Folder: `.lovable/question-and-ambiguity/`
- File: `xx-brief-title.md` where `xx` is the next sequential number
  (continue from existing files; do not restart at 01).
- Sections required:
  - **Task Context** — feature/spec affected.
  - **Specific Question** — the exact uncertainty.
  - **Inferred Decision** — assumption made to proceed.
  - **Impact** — what the decision changes.
  - **Suggested Clarification** — what the user should confirm.
  - **Timestamp** — ISO date.
- Length: under 200 words.

### 2. Inference
- Prefer existing codebase style.
- Prefer the simpler implementation.
- Prefer the most common UX pattern.
- Never silently swallow the ambiguity — always log it.

### 3. Counter
- File: `.lovable/question-and-ambiguity/task-counter.md`.
- Append one line per completed task: `NN. <YYYY-MM-DD> <one-line summary>`.
- When `NN == 40`, mark the prompt **expired** and resume normal behavior.

### 4. User Visibility
- Do NOT voice questions in chat.
- Do NOT mention "I logged an ambiguity" repeatedly — a brief mention in
  the closing sentence is fine when relevant.
- The user reviews `.lovable/question-and-ambiguity/` at the end.

## Deactivation

This prompt becomes inactive when either:
- The counter reaches 40 completed tasks, OR
- The user explicitly says "resume questions" / "exit no-questions mode".

When deactivated, update the table row in `.lovable/prompts.md` Status
column from **Active** to **Expired**.

---

*No-Questions Mode v1.0.0 — 2026-05-04*
