# What to Read

> Canonical map of what the AI must read before working on this project.
> Last updated: 2026-08-09T18:21:37Z

## Changelog

- 2026-08-09T18:21:37Z, Memory write: code red refactor and strict absolute path avoidance.

## Before any task (always)

- `version.json`, why: single source of truth for the repository version, backend/frontend sections, and sub-package version tracks. All codebases must import this file for version information.
- `.lovable/memory/index.md`, why: core memory index
- `.lovable/memory/standards/version-source-of-truth.md`, why: mandatory standard for version.json single source of truth, 'inherit' keyword for sub-packages, and release sync workflow
- `.lovable/memory/release-architecture-map.md`, why: architectural map of version propagation, sync pipeline, and release ceremony
- `.lovable/coding-guidelines/coding-guidelines.md`, why: baseline rules and coding standards
- `.lovable/plans/index.md`, why: active roadmap and pending tasks
- `.lovable/strictly-avoid.md`, why: hard constraints and anti-patterns
- `.lovable/question-and-ambiguity/01-new-ambiguity/`, why: open questions

## Before writing code

- `spec/`, why: understand feature specifications

## Before adding a feature

- `spec/`, why: ensure it fits within existing specs

## Before writing a spec

- `spec/01-spec-authoring-guide/`, why: follow authoring format

## Before adding a unit test

- `spec/02-coding-guidelines/`, why: testing conventions

## See also

- Root `readme.md` (must stay in sync with this file)
- .lovable/plans/pending/01-apperror-new-constructors.md
- .lovable/plans/subtasks/01-apperror-new-constructors/01-update-spec.md
- .lovable/plans/subtasks/01-apperror-new-constructors/02-release.md
- .lovable/plans/pending/02-apperror-human-logger-methods.md
- .lovable/plans/subtasks/02-apperror-human-logger-methods/01-update-display-specs.md
- .lovable/plans/subtasks/02-apperror-human-logger-methods/02-release.md
