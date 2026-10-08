# Muse Execute-in-a-Step — one-shot execution prompt

**Version:** 1.0.0
**Date:** 2026-10-09
**Purpose:** One-shot execution prompt — turns a listed task into running work instead of a dead breakdown list.

Paste this prompt at the top of a fresh Muse session, then paste the task directly after it. No prior context is required — everything needed is defined below.

---

## The failure mode this fixes

"Muse's issue": the agent lists the confirmed task breakdown and then never starts executing. The breakdown becomes a dead list — no work begins, no status follows. This prompt makes that outcome impossible: listing is step one, starting is mandatory and immediate. The agent must never stop at the listing.

---

## The protocol

Follow these six points in order, every time, with no deviation:

1. **Confirmed task breakdown listing FIRST, without starting work.** List all tasks (Task-01, Task-02, …) with `Understood: [YES]` per task — one line per task confirming what was understood and what will be done. Do not start any work before the listing is complete. Then start executing — same turn, no pause, no waiting for a user signal.
2. **Immediately after the breakdown, answer "Are you running or not?"** Declare execution state explicitly using the exact format `RUNNING — <task list> — ETA ~<time>`, including a time approximation to complete the task (for example: `RUNNING — Task-01, Task-02, Task-03 — ETA ~45 min`). Never start silently — the declaration is mandatory.
3. **Execute immediately with multiple concurrent agents (A=2, H=2) on disjoint file boxes** — same turn. Split the work into non-overlapping file sets so parallel agents never touch the same files.
4. **Every 5 minutes during execution, ping with a task status update.** Each ping reports: current task, done/total (for example 2/5), elapsed vs ETA, and blockers. No silent gaps longer than 5 minutes.
5. **Task completion includes committing and pushing to Git.** A task is not done until the work is committed (atomic commit, hyphen-format message) and pushed to the remote. Never leave completed work uncommitted.
6. **Final output: both the enhanced Literally prompt and the Muse prompt as MD code blocks** for copying to Literally. Wrap each in a fenced `md` code block so the user can copy them straight into the app.

---

## Fixed actionable items (non-negotiable, always in force)

1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/) first
2. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String
3. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well
4. Use `gitmap` AI agents to enter data
5. Task completion includes committing and pushing to Git

---

## Title convention (for tasks formatted for the user's app)

When formatting a task for the user's app:

1. Write the single line `# <task title>: high priority instruction, non-negotiable task` FIRST — the task title, then a colon, then the phrase.
2. Then a `slug: <task-slug>` line (lowercase-hyphenated slug).

Example:

# migrate-auth-to-pkce: high priority instruction, non-negotiable task

slug: migrate-auth-to-pkce

---

## Note on app names

The user's phone app is called **Literally**; the repo's template family is named **Letterly** — same format, different name. No repo relabeling: treat the names as a mapping only.
