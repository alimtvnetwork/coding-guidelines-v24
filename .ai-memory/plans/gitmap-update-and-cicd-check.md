# Plan — gitmap-update-and-cicd-check

parent spec: `02-spec/21-app/gitmap-update-and-cicd-check/readme.md`
status: Phase 3 complete (2026-10-18) — all tasks DONE, evidence verified by lead, indexes updated

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

1. gitmap-v28: prior tag v6.570.2 existed while version.json stayed 6.570.1 (version/tag sync slip); 6.571.0 re-syncs. Suggest a pre-release guard asserting version.json == tag. (from C1)
2. No GitHub token / `gh` auth on this VM — releases are tag-based (no release pages, no asset uploads); matches existing convention for these repos. (from C1)
3. Antigravity-Manager: pre-existing version/tag inconsistency NOT fixed (out of scope) — tag v4.226.3 points at a commit whose version.json says 4.226.1; changelog head was v4.226.1 while tags jumped. Owner should reconcile. (from C2)
4. go-email-reader: no CI workflows (pe -t "green" is vacuous); no GitHub Release page for v1.2.0 (no release.yml, gh not authed); NO icon asset exists — confirmed by lead 2026-10-18. If user wants one, use the `logo-and-icon-design` skill. (from C2)
5. repo-cache `.gitmap/last_error.log` documents the known `gitmap llm train` internal pull quirk (`fatal: Cannot fast-forward to multiple branches`, E9000, `llm/llm_train_muse.go:253`) — route to gitmap-v28 track as a CLI defect. (from C2)

## Phase 3 — consolidation (lead, done 2026-10-18)

1. Both ledgers collected; every DONE cites evidence (paths/diffstats/exit codes) — verified.
2. Release tags confirmed on origin via `git ls-remote --tags`: gitmap-v28 v6.571.0, coding-guidelines-v24 v6.79.0, Antigravity-Manager v4.227.0, go-email-reader v1.2.0.
3. gitmap CLI: v6.566.4 installed durably at `~/.local/bin/gitmap-cli/gitmap` + `~/workspace/tools/bin/gitmap` (Go 1.26.9 toolchain at `~/workspace/tools/go`); `gitmap scan ~/workspace/repos` = 9 repos, 13 projects; `gitmap pa` = 9 pulled, 0 conflicts.
4. Indexes updated: `.ai-memory/plans/readme.md`, `02-spec/21-app/readme.md`.
5. Final atomic commit + push of consolidation files.
