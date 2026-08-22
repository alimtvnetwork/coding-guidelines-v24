# Lovable Prompts Index

**Version:** 2.0.0
**Updated:** 2026-08-22

This index is the canonical entry point for the prompts that govern how
Lovable reads, writes, and updates project memory. It is referenced from
`.lovable/coding-guidelines/coding-guidelines.md` as a required read
before generating any code.

The companion linter `linter-scripts/check-prompts-loaded.py` enforces
that this file exists and lists every prompt under `.lovable/prompts/`.
Add an entry here whenever a new prompt file is introduced.

## Required reads

Read these in order before writing or modifying code:

1. `.lovable/coding-guidelines/coding-guidelines.md` - project-wide rules
2. `.lovable/prompts.md` (this file) - prompt index
3. Every prompt referenced in the table below

## Prompts

| # | File | Title | Purpose |
|---|------|-------|---------|
| 01 | [`prompts/01-read-memory-enhanced.md`](./prompts/01-read-memory-enhanced.md) | Read Memory (Enhanced) | Canonical procedure for reading existing memory before acting. |
| 02 | [`prompts/02-write-antigravity.md`](./prompts/02-write-antigravity.md) | Write Antigravity | Write Antigravity procedures. |
| 03 | [`prompts/03-write-memory.md`](./prompts/03-write-memory.md) | Write Memory | End of session write memory. |
| 04 | [`prompts/04-plan-coding-guideline-audit.md`](./prompts/04-plan-coding-guideline-audit.md) | Plan Guideline Audit | Audit code against coding guidelines. |
| 05 | [`prompts/05-execute-coding-guideline-fix.md`](./prompts/05-execute-coding-guideline-fix.md) | Execute Guideline Fix | Execute the fixes from the coding guideline audit. |
| 06 | [`prompts/06-commit-fix.md`](./prompts/06-commit-fix.md) | Commit Fix | Commit and fix routines. |
| 07 | [`prompts/07-commit-fix-v2.md`](./prompts/07-commit-fix-v2.md) | Commit Fix v2 | Refined commit fix routines. |
| 08 | [`prompts/08-release.md`](./prompts/08-release.md) | Release | Instructions for bumping versions and releases. |
| 09 | [`prompts/09-clean-artifacts-and-git-history.md`](./prompts/09-clean-artifacts-and-git-history.md) | Clean Artifacts and History | Repository purge procedures. |

## Maintenance

- One row per file in `.lovable/prompts/`. No orphans, no missing entries.
- The linter accepts a relative link or a bare filename match - keep links
  relative for clickability.
- When superseding a prompt, update the *Purpose* column rather than
  deleting the row, so historical references still resolve.
