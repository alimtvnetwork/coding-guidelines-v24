# What to Read

> Canonical map of what the AI must read before working on this project.
> Last updated: 2026-09-04T17:39:00Z

## Changelog

- 2026-09-04T17:39:00Z, Memory write: parallel multi-worker CI/CD local runner, selective log filtering, streamwriter contracts, and naming standards.
- 2026-09-04T02:15:00Z, Ingested whole codebase, added write-memory and write-antigravity skills, and recorded 05-codebase-topology-and-skills-architecture.md.
- 2026-08-09T18:21:37Z, Memory write: code red refactor and strict absolute path avoidance.

## Before any task (always)

- `version.json`, why: single source of truth for the repository version, backend/frontend sections, and sub-package version tracks. All codebases must import this file for version information.
- `.lovable/memory/01-index.md`, why: core memory index
- `.lovable/memory/learned/01-project-context-and-guidelines.md`, why: canonical learned memory of repo identity, CODE RED rules, coding guidelines, error philosophy, and active plans
- `.lovable/memory/learned/03-parallel-cicd-runner-and-log-filtering.md`, why: parallel local runner concurrency, duration tracking, and log suppression standard
- `.lovable/memory/learned/04-streamwriter-contracts-and-naming-standards.md`, why: streamwriter contracts, reentrant locker, monadic Bytes[T], JsonResult multi-source ingestion, boolean prefixes, and Id naming standard
- `.lovable/memory/learned/05-codebase-topology-and-skills-architecture.md`, why: comprehensive topology ingestion, 28 AI Python scripts catalog, Antigravity skills inventory, and CI/CD quality gate enforcement
- `.lovable/memory/standards/version-source-of-truth.md`, why: mandatory standard for version.json single source of truth, 'inherit' keyword for sub-packages, and release sync workflow
- `.lovable/memory/01-index.md`, why: architectural map of version propagation, sync pipeline, and release ceremony
- `.lovable/coding-guidelines.md`, why: baseline rules and coding standards
- `.lovable/plans/01-index.md`, why: active roadmap and pending tasks
- `.lovable/strictly-avoid.md`, why: hard constraints and anti-patterns
- `.lovable/question-and-ambiguity/01-new-ambiguity/`, why: open questions


## Before writing code

- `spec/`, why: understand feature specifications

## Before adding a feature

- `spec/`, why: ensure it fits within existing specs

## Before writing a spec

- `02-spec/01-spec-authoring-guide/`, why: follow authoring format

## Before adding a unit test

- `02-spec/02-coding-guidelines/`, why: testing conventions

## See also

- Root `readme.md` (must stay in sync with this file)
- .lovable/plans/01-index.md
- .lovable/plans/pending/02-slides-system-overhaul.md
- .lovable/plans/pending/04-guideline-prompt-and-installer-upgrade.md
- .lovable/plans/pending/09-update-prompts-and-release.md
- .lovable/plans/pending/11-code-red-refactor-remediation.md
- .lovable/plans/completed/01-apperror-new-constructors.md
- .lovable/plans/completed/03-apperror-human-logger-methods.md
- .lovable/plans/completed/05-rename-overviews-and-installer-json.md
- .lovable/plans/completed/06-fix-encoding.md
- .lovable/plans/completed/07-trailing-newlines-and-ai-scripts.md
- .lovable/plans/completed/08-lowercase-changelog.md
- .lovable/plans/completed/10-rca-and-boolean-fix.md
- .lovable/plans/completed/12-prompt-architect-version-tracking.md
- .lovable/plans/completed/13-cicd-pipeline-consolidation-and-owner-review.md
