# Plan — gitmap-update-and-cicd-check

parent spec: `02-spec/21-app/gitmap-update-and-cicd-check/readme.md`
status: Phase 2 dispatched (2026-10-18)

## Phase 1 — plan + spec (lead, done 2026-10-18)

- /plan: breakdown extracted, routing assumption stated, partition fixed.
- /learn: gitmap skill command matrix absorbed (scan, pa, pe -t, cpf/cpb/cpr, colon ban, search-only-via-gitmap).
- Preflight finding: gitmap binary missing on VM → Task-00 starts with reinstall/upgrade from gitmap-v28.
- This file + parent spec written and committed atomically before dispatch.

## Phase 2 — execution waves

WAVES = ceil(subtasks / (A x H)); A=2 coordinators, disjoint boxes:

| Wave | Owner | Tasks | Subtask ledger |
|---|---|---|---|
| 1 | Coordinator-01 (gitmap-track) | Task-00 (gitmap reinstall, scan, pa), Task-01 (gitmap-v28: pull, CI green, fix, minor bump, release), Task-05a (alim-seo-writing, story-telling-v1: pull, CI green, fix-only) | `.ai-memory/plans/subtasks/gitmap-update-and-cicd-check/c1-gitmap-track.md` |
| 1 | Coordinator-02 (apps-track) | Task-02 (Antigravity-Manager: pull, CI green, fix, minor bump, release), Task-03 (coding-guidelines-v24: pull, CI green, fix, minor bump, release), Task-04 (go-email-reader: clone via SSH, pull, CI green, fix, minor bump, release, GER icon check), Task-05b (cat-my-v12, repo-cache, repo-secrets: pull, CI green, fix-only) | `.ai-memory/plans/subtasks/gitmap-update-and-cicd-check/c2-apps-track.md` |

Each coordinator runs `execute-parent-task-with-n-steps-v6` inside its box (spawn 2 workers per stage, research → spec → execution, one atomic commit+push per task).

Per-repo release checklist (Tasks 01–04): pull clean → `gitmap pe -t` green (wait) → fix failures → minor-bump skill → release skill → release published with relative paths only in notes.

Stop conditions: SSH/API auth failure, CI red unfixable after RCA, release blocked on missing approval → log under Follow-ups, report; do NOT retry blindly.

## Phase 3 — consolidation (lead)

1. Collect both ledgers; verify each DONE cites evidence (path/diffstat/exit 0).
2. Update indexes: `.ai-memory/plans/readme.md`, `02-spec/21-app/readme.md`.
3. Final atomic commit + push of index/spec updates in coding-guidelines-v24.
4. Report to user: per-repo CI status, versions bumped, release links, GER icon confirmation.

## Follow-ups

(none yet)
