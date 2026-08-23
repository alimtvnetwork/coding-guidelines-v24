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
| 01 | [`prompts/01-unified-ai-prompt-v4.md`](./prompts/01-unified-ai-prompt-v4.md) | Unified AI Prompt v4 | Core unified AI workflow prompt. |
| 02 | [`prompts/02-next-steps.md`](./prompts/02-next-steps.md) | Next Steps | Determining the immediate next steps. |
| 03 | [`prompts/03-pending-tasks.md`](./prompts/03-pending-tasks.md) | Pending Tasks | Managing pending tasks. |
| 04 | [`prompts/04-plan-steps.md`](./prompts/04-plan-steps.md) | Plan Steps | Planning workflow steps. |
| 05 | [`prompts/05-read-memory-enhanced.md`](./prompts/05-read-memory-enhanced.md) | Read Memory (Enhanced) | Canonical procedure for reading existing memory before acting. |
| 06 | [`prompts/06-write-antigravity.md`](./prompts/06-write-antigravity.md) | Write Antigravity | Write Antigravity procedures. |
| 07 | [`prompts/07-write-memory.md`](./prompts/07-write-memory.md) | Write Memory | End of session write memory. |
| 08 | [`prompts/08-plan-coding-guideline-audit.md`](./prompts/08-plan-coding-guideline-audit.md) | Plan Guideline Audit | Audit code against coding guidelines. |
| 09 | [`prompts/09-execute-coding-guideline-fix.md`](./prompts/09-execute-coding-guideline-fix.md) | Execute Guideline Fix | Execute the fixes from the coding guideline audit. |
| 10 | [`prompts/10-execute-batched-loop.md`](./prompts/10-execute-batched-loop.md) | Execute Batched Loop | Guidelines for executing tasks in batched loops. |
| 11 | [`prompts/11-inventory-pending-tasks.md`](./prompts/11-inventory-pending-tasks.md) | Inventory Pending Tasks | Take stock of pending items. |
| 12 | [`prompts/12-fix-subtask-naming-convention.md`](./prompts/12-fix-subtask-naming-convention.md) | Fix Subtask Naming Convention | Enforce naming conventions on subtasks. |
| 13 | [`prompts/13-ci-cd-fix.md`](./prompts/13-ci-cd-fix.md) | CI/CD Fix | Guidelines for resolving CI/CD pipeline issues. |
| 14 | [`prompts/14-cicd-run-ps1.md`](./prompts/14-cicd-run-ps1.md) | CICD Run PS1 | Details on executing CI/CD PowerShell run scripts. |
| 15 | [`prompts/15-commit-fix.md`](./prompts/15-commit-fix.md) | Commit Fix | Commit and fix routines. |
| 16 | [`prompts/16-commit-fix-v2.md`](./prompts/16-commit-fix-v2.md) | Commit Fix v2 | Refined commit fix routines. |
| 17 | [`prompts/17-clean-artifacts-and-git-history.md`](./prompts/17-clean-artifacts-and-git-history.md) | Clean Artifacts and History | Repository purge procedures. |
| 18 | [`prompts/18-release.md`](./prompts/18-release.md) | Release | Instructions for bumping versions and releases. |
| 19 | [`prompts/19-plan-spec-steps-v2.md`](./prompts/19-plan-spec-steps-v2.md) | Plan Spec Steps v2 | Instructions for planning spec steps (v2). |
| 20 | [`prompts/20-audit-app-spec.md`](./prompts/20-audit-app-spec.md) | Audit App Spec | Instructions for auditing app specs. |

## Maintenance

- One row per file in `.lovable/prompts/`. No orphans, no missing entries.
- The linter accepts a relative link or a bare filename match - keep links
  relative for clickability.
- When superseding a prompt, update the *Purpose* column rather than
  deleting the row, so historical references still resolve.
