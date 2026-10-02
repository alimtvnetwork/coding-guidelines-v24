# 17 Spec ticket for a blind agent

**Version:** 1.0.0
**Updated:** 2026-10-02
**Status:** Pending
**Verification date:** 2026-10-02

This plan is the work. Implement it in order. Do not redesign it.

---

## Context

A weak agent can follow a short ticket with pass/fail checks. It cannot follow a large library that disagrees with itself.

This repository is the library: coding rules, error types, and design tokens, plus linters for some of the rules. The G-Spec source is the `/spec` skill in the local gstack checkout, file `spec/SKILL.md`. Read it. Do not edit it. Do not copy that repository into this one.

Who is affected: the next agent asked to implement a spec in this repo.
Current gap: the authoring guide tells every module to claim a confidence score, and many files say `Ambiguity: None` while other files contradict them.
Desired result: one new ticket template, one heading checker, and one reading path. Existing modules stay as they are.
Why now: design and coding files have grown past what one agent can read, and false "None" scores send the agent into the wrong file.
Done when: the acceptance criteria in this plan all pass.

---

## Current state (verified 2026-10-02)

| Piece | Path | What it does today |
|---|---|---|
| Authoring entry | `02-spec/01-spec-authoring-guide/readme.md` | Requires a confidence score and an ambiguity score on every overview |
| Required files | `02-spec/01-spec-authoring-guide/04-required-files.md` | `readme.md` and `99-consistency-report.md` are mandatory. `97-acceptance-criteria.md` is recommended |
| Next free authoring number | none | Files `02` through `14` exist. `15` is free. `97`, `98`, `99` stay reserved |
| Module template | `02-spec/01-spec-authoring-guide/08-non-cli-module-template.md` | Module shape. It is not a one-change ticket |
| Plans index | `.ai-memory/plans/readme.md` | Pending list. This file is `pending/17-spec-ticket-for-blind-agents.md` |
| G-Spec process | gstack checkout `spec/SKILL.md` around the Process and Issue Quality Standards headings | Five phases, 14 quality rules, issue headings, anti-patterns |
| G-Spec gate | gstack checkout `spec/sections/gate-and-file.md` | Outside reviewer score plus a secret scan. Depends on bun and a local gstack install |

Do not edit the gstack checkout.

---

## What to copy, and what to leave

Copy these ideas into our template, in our words:

1. Who, current behavior, desired behavior, why now, and how we know it is done. All five are required before the proposed change.
2. An explicit out-of-scope list.
3. Current state cited as `path:line`, with a verification date.
4. A do-not-touch list.
5. A file table: path, and the change.
6. Numbered acceptance criteria that a stranger can pass or fail. Ban the phrases `works correctly` and `edge cases are handled`.
7. A test table with layer, what, and count.
8. A rollback line. `Revert the commit` is allowed when no data migration exists.
9. One ticket stays small. If the file table grows past 8 paths, split the ticket. Do not do that split inside this plan.

Leave these in G-Spec. Do not build them:

- `gstack-skill-start`, telemetry, and Claude preamble
- bun, Codex, and the outside reviewer process
- GitHub issue filing, issue dedupe, and `/ship` auto-close
- AskUserQuestion tool formats
- Secret-scanning binaries

---

## Do not touch

- `version.json`, `package.json`, and `changelog.md`
- `02-spec/07-design-system/` token values, theme ids, and page sizes
- `02-spec/02-coding-guidelines/` hard-rule text
- `.cursorrules` and `AGENTS.md` (the sync script owns the mirrors)
- Any file in the gstack checkout
- Existing `99-consistency-report.md` scores in modules you are not editing

The boolean-prefix clash between `AGENTS.md` (`is` and `has` only) and `.cursorrules` (`is`, `has`, `can`, `should`, `was`, `will`) is real. It is out of scope here. Do not "fix" it in a mirror file.

---

## Sequencing

1. Write the ticket template (task A). The checker reads its heading list.
2. Write the checker (task B). It fails if a ticket file is missing a heading.
3. Add the reading path (task C) so an agent opens three files, not the whole tree.
4. Point the authoring readme at those files (task D).
5. Run the checker on the template. Fix until it exits 0.

Task C can be drafted in parallel with B only after A is on disk. D is last.

---

## Task A: ticket template

Create `02-spec/01-spec-authoring-guide/15-executable-spec-ticket.md`.

Required headings, in this order, as level-2 headings:

1. `Context`
2. `Current state`
3. `Proposed change`
4. `Acceptance criteria`
5. `Testing plan`
6. `Rollback`
7. `Files`
8. `Out of scope`
9. `Do not touch`
10. `Checklist`

Rules the template must state, each as its own sentence:

- Fill `Context` with five labeled lines: Who, Current, Desired, Why now, Done when.
- `Current state` names a verification date and at least one `path:line` citation.
- `Acceptance criteria` are numbered. Each one is pass or fail.
- `Testing plan` is a table with columns Layer, What, Count.
- `Files` is a table with columns File, Change. Paths start at the repo root. No drive letters. No `file:///` URIs.
- `Out of scope` and `Do not touch` each contain at least one bullet.
- `Checklist` repeats the acceptance criteria as `- [ ]` items.
- A ticket with more than 8 rows in `Files` must be split before implementation starts.
- Version stamp on this file: `4.3.0`. Updated: `2026-10-02`. Status: `Active`.

Include one filled example about adding a heading checker. The example's file table may list only the checker path and this template. The example must not describe a product, a company, or a URL.

Keep the file at or under 300 lines, skipping blanks and comments. Comments in markdown are not a special syntax here: count every non-blank line. If the draft exceeds 280 non-blank lines, cut the example, not the heading rules.

## Task B: heading checker

Create `linter-scripts/check-spec-ticket-headings.py`.

Behavior:

- Input: one or more paths. If no path is given, scan `02-spec/01-spec-authoring-guide/15-executable-spec-ticket.md` only.
- A file is in scope when its text contains the exact line `# Executable spec ticket` or the exact line `## Context` together with `## Do not touch`.
- For an in-scope file, require the 10 headings from task A, in order. Missing, renamed, or reordered headings print `path: missing <heading>` or `path: out of order <heading>` and set a non-zero exit.
- Ignore files that are not in scope. Do not fail the historical spec tree.
- No network. No new dependency. Python standard library only.
- Print `PASS` and exit 0 when every in-scope file has the 10 headings in order.

Do not register this checker in CI in this plan. A later plan can add the CI step after the checker has passed locally.

## Task C: reading path

Create `02-spec/01-spec-authoring-guide/16-blind-agent-reading-path.md`.

State this order and no other:

1. Read `02-spec/01-spec-authoring-guide/15-executable-spec-ticket.md`.
2. Read the one module `readme.md` the ticket names.
3. Read only the files listed in that ticket's `Files` table.

Then stop. Do not open `02-spec/07-design-system/` unless the ticket lists a file there. Do not treat `Ambiguity: None` on an old file as permission to invent a missing value.

Version stamp: `4.3.0`. Updated: `2026-10-02`. Status: `Active`. End with a checklist of those three steps.

Keep this file under 120 non-blank lines.

## Task D: entry points

Edit only these lines. Do not rewrite the files.

1. `02-spec/01-spec-authoring-guide/readme.md`: in the files table, add rows for `15-executable-spec-ticket.md` and `16-blind-agent-reading-path.md`. Add one sentence under Scoring Metrics: an overview may say `Ambiguity: None` only when `99-consistency-report.md` in that same folder has result `PASS` and names no open contradiction. Otherwise write the contradiction in one sentence.
2. `02-spec/01-spec-authoring-guide/04-required-files.md`: add one row to Strongly Recommended Files for `15-executable-spec-ticket.md` shape, used when the work is a single change a stranger will implement. Do not make it mandatory for every module.
3. `.ai-memory/plans/readme.md`: leave this plan under Pending until the acceptance criteria pass. After they pass, move the bullet to Completed. Do that move in the same commit as the passing checker run.

Update `02-spec/01-spec-authoring-guide/99-consistency-report.md` only to add the two new filenames to its inventory. Do not change its result line unless that file's own rules require a new row. If the report has no inventory table, skip it and say so in the commit message.

---

## Acceptance criteria

1. `02-spec/01-spec-authoring-guide/15-executable-spec-ticket.md` exists and contains the 10 headings in the task A order.
2. `python linter-scripts/check-spec-ticket-headings.py` prints `PASS` and exits 0.
3. The same command exits non-zero when run on a temp copy of the template with `## Rollback` deleted. Delete the temp copy after the check. Do not commit it.
4. `16-blind-agent-reading-path.md` lists exactly three read steps, in the task C order.
5. The authoring `readme.md` files table links to files 15 and 16.
6. `04-required-files.md` does not say every module must contain a ticket.
7. `git diff --stat` for the implementation commit does not include `version.json`, `02-spec/07-design-system/`, or any path outside this repository.
8. No new file contains a drive letter, a `file:///` URI, or a client product name.

## Testing plan

| Layer | What | Count |
|---|---|---|
| Command | `python linter-scripts/check-spec-ticket-headings.py` on the real template | 1 pass |
| Command | Same checker on a temp file missing `## Rollback` | 1 fail |
| Read | Files 15 and 16 stay under the line caps in tasks A and C | 2 checks |

No unit-test framework is required. No browser check. No release.

## Rollback

Revert the implementation commit. No data migration.

## Files

| File | Change |
|---|---|
| `02-spec/01-spec-authoring-guide/15-executable-spec-ticket.md` | Create the template |
| `02-spec/01-spec-authoring-guide/16-blind-agent-reading-path.md` | Create the three-step path |
| `linter-scripts/check-spec-ticket-headings.py` | Create the heading checker |
| `02-spec/01-spec-authoring-guide/readme.md` | Add two inventory rows and one ambiguity sentence |
| `02-spec/01-spec-authoring-guide/04-required-files.md` | Recommend the ticket shape, do not require it |
| `02-spec/01-spec-authoring-guide/99-consistency-report.md` | Add the two filenames if an inventory table exists |
| `.ai-memory/plans/readme.md` | Move this plan to Completed after the criteria pass |

## Out of scope

- Rewriting existing design or coding modules into tickets
- CI wiring for the new checker
- Reconciling boolean prefixes
- A version bump or a GitHub release
- Porting G-Spec's bun, Codex, or issue-filing flow

## Checklist

- [ ] Task A file exists with the 10 headings in order
- [ ] Checker exits 0 on that file and non-zero when `## Rollback` is removed from a temp copy
- [ ] Task C file lists only the three read steps
- [ ] Authoring readme links files 15 and 16
- [ ] Required-files doc does not force a ticket on every module
- [ ] Diff does not touch `version.json`, the design system, or the gstack checkout
- [ ] Plans index moves this plan to Completed in the passing commit
