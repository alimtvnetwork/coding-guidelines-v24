# 23 — Visibility Change

## Purpose

Toggle (or explicitly set) the visibility of the current Git repository on
its hosting provider — **GitHub** or **GitLab** — without leaving the
terminal. Provides matched **PowerShell** (`visibility-change.ps1`) and
**Bash** (`visibility-change.sh`) front-ends, plus a `visibility`
sub-command on the existing `run.ps1` / `run.sh` runner.

## Why

The team frequently flips repos between `public` and `private` during the
release cycle (private while staging, public on cut). The web UIs hide the
toggle in a “Danger Zone” section that requires several clicks and is easy
to mis-click. A single, idempotent CLI invocation is faster, scriptable,
and less error-prone.

## Scope

- **In scope**
  - GitHub repos detected from `origin` (https or ssh forms).
  - GitLab repos detected from `origin` (gitlab.com and self-hosted).
  - Toggle (default), explicit `--visible pub`, explicit `--visible pri`.
  - Confirmation prompt when going `private → public` (skip with `--yes`).
  - `--dry-run` to preview without acting.
  - Integration into the `run.ps1` / `run.sh` sub-command surface.
- **Out of scope**
  - Any provider other than GitHub / GitLab (Bitbucket, Gitea, etc.).
  - Bulk operations across multiple repos.
  - Changing `internal` (GitLab-only third state) — only `public` /
    `private` are exposed.
  - Token storage; we delegate auth to `gh` / `glab` CLIs.

## Files (this folder)

- `00-overview.md` — this file
- `01-spec.md` — normative behavior, flags, exit codes, contracts
- `02-edge-cases.md` — unusual inputs and their handling
- `03-acceptance-criteria.md` — testable assertions
- `04-examples.md` — copy-pasteable invocations
- `plan.md` — phased implementation plan

## Decisions Locked (from approval)

| Question | Choice |
|---|---|
| Auth backend | `gh` / `glab` CLIs (delegated browser login) |
| No-flag default | Toggle current visibility |
| Runner integration | Sub-command + standalone scripts |
| Confirmation | Only on `private → public` (skip with `--yes`) |
