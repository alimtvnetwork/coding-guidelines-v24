# Lovable Prompts Index

**Version:** 1.0.0
**Updated:** 2026-04-27

This index is the canonical entry point for the prompts that govern how
Lovable reads, writes, and updates project memory. It is referenced from
`.lovable/coding-guidelines/coding-guidelines.md` as a required read
before generating any code.

The companion linter `linter-scripts/check-prompts-loaded.py` enforces
that this file exists and lists every prompt under `.lovable/prompts/`.
Add an entry here whenever a new prompt file is introduced.

## Required reads

Read these in order before writing or modifying code:

1. `.lovable/coding-guidelines/coding-guidelines.md` — project-wide rules
2. `.lovable/prompts.md` (this file) — prompt index
3. Every prompt referenced in the table below

## Prompts

| # | File | Title | Purpose |
|---|------|-------|---------|
| 01 | [`prompts/01-read-prompt.md`](./prompts/01-read-prompt.md) | Read Memory Prompt | Canonical procedure for reading existing memory before acting. |
| 02 | [`prompts/02-write-prompt.md`](./prompts/02-write-prompt.md) | Write Memory Prompt | Earlier write-memory procedure (v3.3.0); kept for reference. |
| 03 | [`prompts/03-write-prompt.md`](./prompts/03-write-prompt.md) | Write Memory | Intermediate write-memory revision (v1.0.0). |
| 04 | [`prompts/04-write-prompt.md`](./prompts/04-write-prompt.md) | Write Memory (v2.0.0) | Superseded write-memory revision; retained for diff history. |
| 05 | [`prompts/05-write-prompt.md`](./prompts/05-write-prompt.md) | Write Memory — Canonical v3.0.0 | **Active** write/end/update memory procedure. |

## Maintenance

- One row per file in `.lovable/prompts/`. No orphans, no missing entries.
- The linter accepts a relative link or a bare filename match — keep links
  relative for clickability.
- When superseding a prompt, update the *Purpose* column rather than
  deleting the row, so historical references still resolve.
