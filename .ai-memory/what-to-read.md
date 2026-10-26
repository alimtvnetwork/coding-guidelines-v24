# What to Read

> Canonical map of what the AI must read before working on this project.
> Version: 5.0.0
> Last updated: 2026-10-27T00:00:00Z

## Changelog

- 2026-10-27T00:00:00Z, What-to-Read 5.0.0 Upgrade: restructured into 5-Tier Reading Order (Tier 1: Safety & Invariants, Tier 2: Identity & Baseline Contracts, Tier 3: Institutional Memory, Tier 4: Active Roadmap & Work State, Tier 5: Operational Tooling & Specifications), linked V6 execution prompt (`01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`), Letterly prompt formatters (`01-prompts/22-letterly/`), and Muse universal master prompt (`01-prompts/27-muse-prompts/04-muse-universal-master-prompt.md`), replaced raw git commands with `gitmap status` / `gitmap ps`, fixed dead links to `04-version-source-of-truth.md` and `.ai-memory/ambiguous-questions/`, and converted skill links to IDE syntax `[skill-name](file;.agents/skills/<skill-name>)`.
- 2026-10-02T12:10:00Z, Multi-Repository Sync Guardrails: codified non-negotiable cross-repo sync boundaries forbidding sync of 02-spec/21-* (spec 21), protecting existing modified AI scripts in target repos (additive-only sync), and locking bump scripts from being overwritten (38-sync-prompts-skills-scripts.py, strictly-avoid.md, 18-cross-repository-sync-rules.md).
- 2026-09-30T16:39:00Z, Release v6.49.0 and Execute-Folder Hardening (plan 16): execute prompts 01 to 09 and their six Antigravity skill copies now stage by explicit path, drop the invalid `"Model"` payload field and the undocumented `<SYSTEM_MESSAGE>` tag, run the autofixer with `--check-only` on folders, and commit `go generate` output only when the repo already tracks it; prompts 02 and 06 point to V4.
- 2026-09-30T16:18:00Z, Parent Task N-Steps V4 Prompt (`11-execute-parent-task-with-n-steps-v4.md`): Antigravity-native rewrite of V3 with rules R1 to R15, a capability preflight, a resumable ledger, a corrected `invoke_subagent` payload, explicit-path staging, and evidence-gated checks; plan 15 and thin pointer skills for Antigravity and Cursor.
- 2026-09-30T14:00:00Z, White Blue Theme Design System (`04-white-blue-theme`), Parent Task N-Steps V3 Prompt (`10-execute-parent-task-with-n-steps-v3.md` with `N = 300` top header and mandatory `invoke_subagent` `A = 2, H = 2`), skills sync, and 43-repository backup/release/sync/release orchestration.
- 2026-09-18T19:20:00Z, GitMap Pipeline-AI & Dynamic Waiting Protocol: updated execute prompts, fix-with-rca, ci-cd-fix, and ci-cd-fix-with-release prompts and skills to mandate GitMap pipeline AI integration, targeted failure extraction, and dynamic waiting (-t <sec>) to eliminate credit waste from tight-loop polling.
- 2026-09-13T08:15:00Z, Prompt & Memory Write Upgrade (v2.2.0): updated 01-write-antigravity.md, 03-write-memory.md, and skills to mandate 30-commit git history audit and recent 20-task tracking register in .ai-memory/plans/readme.md.
- 2026-09-13T07:45:00Z, Memory Persistence & Learned Standards: recorded institutional memory for regex centralization, generic dbengine architecture, isDefined positive convention, zero-storage GitHub Actions mandate, and pre-commit guard enforcement.
- 2026-09-13T07:00:00Z, Regex Centralization & Generic DbEngine: lazy regex engine audit, core regex harvesting from 03-aukgo/core, gitmap call-site refactoring, generic redistributable dbengine package (pkg/dbengine), and Python database code generator (35-db-struct-enum-generator.py).
- 2026-09-09T05:30:00Z, Prompt update: added mandatory inspection of last 10 git commits and what-to-read prioritization to read-memory-enhanced prompt and skill.
- 2026-09-09T05:00:00Z, Memory write: conversation log & context wrapper protocol, prompt staging, split SQLite logging, task retention, errcmd streaming, atomic file writes, and ApiManager spec.
- 2026-09-04T17:39:00Z, Memory write: parallel multi-worker CI/CD local runner, selective log filtering, streamwriter contracts, and naming standards.
- 2026-09-04T02:15:00Z, Ingested whole codebase, added write-memory and write-antigravity skills, and recorded 05-codebase-topology-and-skills-architecture.md.
- 2026-08-09T18:21:37Z, Memory write: code red refactor and strict absolute path avoidance.

## 5-Tier Reading Order

Before starting any task, all AI agents MUST read and ingest the codebase in this strict sequence:

### Tier 1: Safety & Invariants
Foundational guardrails, negative constraints, and non-negotiable boundaries that apply to all actions.

- `.ai-memory/strictly-avoid.md`, why: universal hard prohibitions, CODE RED constraints, absolute path ban, non-owned repository exclusion, and CI/CD protection rules.
- `.ai-memory/memory/learned/18-cross-repository-sync-rules.md`, why: cross-repo sync boundaries: spec 21 exclusion (`02-spec/21-*` through `02-spec/25-*`), additive-only AI scripts, and immutable bump script protection.
- `02-spec/02-coding-guidelines/`, why: universal cross-language standards, positive boolean rules (`is` / `has`), no bare void in Go (`*appfault.AppError` / `Result[T]`), parameter structs, and vertical line gaps.

### Tier 2: Identity & Baseline Contracts
Repository architecture, single source of truth for versioning, and foundational engineering specifications.

- `version.json`, why: single source of truth for repository SemVer version, backend/frontend sections, and sub-package version tracks.
- `.ai-memory/memory/standards/04-version-source-of-truth.md`, why: mandatory standard for version.json single source of truth, 'inherit' keyword for sub-packages, and release sync workflow.
- `.ai-memory/memory/learned/01-project-context-and-guidelines.md`, why: canonical learned memory of repo identity, CODE RED rules, coding guidelines, error philosophy, and active plans.
- `.ai-memory/coding-guidelines.md`, why: baseline rules, coding standards summary, and core behavioral invariants.
- `.ai-memory/folder-structure.md`, why: authoritative canonical layout of `.ai-memory/`, `02-spec/`, and top-level directory hierarchy.
- `02-spec/01-spec-authoring-guide/`, why: specification authoring format, required files, folder structure, and validation rules.

### Tier 3: Institutional Memory
Historical architectural decisions, learned conventions, and deep systems engineering patterns.

- `.ai-memory/memory/readme.md`, why: master institutional memory index and navigation table.
- `.ai-memory/memory/learned/02-database-conventions.md`, why: SQLite Split DB, PascalCase tables, {Table}Id PKs, and positive booleans.
- `.ai-memory/memory/learned/03-enum-standards.md`, why: cross-language enum generator workflow and naming rules.
- `.ai-memory/memory/learned/03-parallel-cicd-runner-and-log-filtering.md`, why: parallel local runner concurrency, duration tracking, and log suppression standard.
- `.ai-memory/memory/learned/04-streamwriter-contracts-and-naming-standards.md`, why: streamwriter contracts, reentrant locker, monadic Bytes[T], JsonResult multi-source ingestion, boolean prefixes, and Id naming standard.
- `.ai-memory/memory/learned/05-codebase-topology-and-skills-architecture.md`, why: comprehensive topology ingestion, AI Python scripts catalog, skills inventory, and CI/CD quality gate enforcement.
- `.ai-memory/memory/learned/07-split-sqlite-logging-and-task-db-migration.md`, why: split SQLite logging architecture and task DB migration contracts.
- `.ai-memory/memory/learned/08-task-retention-streaming-atomic-apimanager.md`, why: task retention, live line streaming, atomic file writes, and ApiManager spec.
- `.ai-memory/memory/learned/09-conversation-log-and-context-wrapper-protocol.md`, why: conversation log persistence protocol and prompt staging boundary.
- `.ai-memory/memory/learned/13-regex-centralization-dbengine-and-isdefined-standard.md`, why: canonical regex harvesting from 03-aukgo/core, redistributable dbengine package, isDefined standard, and zero-storage GitHub Actions mandate.
- `.ai-memory/memory/learned/14-write-prompts-git-audit-and-recent-tasks-register.md`, why: mandatory 30-commit git history audit before memory authoring, compact 20-task recent completion register in plans index, and standardized 19-box verification checklist.
- `.ai-memory/memory/learned/19-gitmap-cicd-healing-and-letterly-integration.md`, why: GitMap pipeline-ai integration, dynamic waiting, failure extraction, and Letterly golden prompt standards.
- `.ai-memory/memory/learned/20-letterly-execute-n-steps-and-run-scripts.md`, why: Letterly prompt formatters, mobile one-liner mode, and run script orchestration.

### Tier 4: Active Roadmap & Work State
Current repository status, task backlogs, bounded micro-tasks, and active ambiguities.

- `gitmap status` / `gitmap ps "git log -n 30 --oneline"`, why: inspect the last 30 commits to understand recent file changes, directives, lessons learned, and the latest repository state via GitMap commands before starting any task.
- `.ai-memory/plans/readme.md`, why: active roadmap, pending tasks, and recent completed tasks register (last 20 completed tasks/plans).
- `.ai-memory/plans/pending/`, why: active parent plans and pending milestone backlogs.
- `.ai-memory/plans/subtasks/`, why: bounded micro-tasks grouped by parent slug (`XX-<parent-slug>/`).
- `.ai-memory/ambiguous-questions/01-new-ambiguity/`, why: open questions and pending ambiguities requiring clarification.

### Tier 5: Operational Tooling & Specifications
Autonomous execution prompts, prompt formatters, scripts catalog, and domain-specific specifications.

- `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`, why: canonical V6 autonomous execution engine (Tree Swarm Hierarchy, Brain roadmap, dynamic agent join, SQLite task manager, and GitMap atomic commits).
- `01-prompts/22-letterly/`, why: Letterly voice dictation prompt formatters (mobile, desktop, execute-n-steps, plan, release, cicd-fix, run, muse modes).
- `01-prompts/27-muse-prompts/04-muse-universal-master-prompt.md`, why: unified Muse master prompt for onboarding, fleet orchestration, and autonomous V6 execution.
- `03-ai-scripts/readme.md`, why: master catalog, syntax, and operational guidelines for the Python automation toolchain and local CI runners.
- `02-spec/`, why: complete specification hierarchy (Tiers 01–20 core fundamentals and Tiers 21–25 application domain specs).
- `02-spec/25-spec-audits/`, why: specification audit scorecards, benchmark comparisons, and remediation matrices.

## See Also & Primary Skills

- Root `readme.md` (must stay in sync with this file)
- `.ai-memory/plans/readme.md`
- [execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)
- [muse-master-prompt](file;.agents/skills/muse-master-prompt)
- [letterly-execute-n-steps](file;.agents/skills/letterly-execute-n-steps)
- [letterly-mobile](file;.agents/skills/letterly-mobile)
- [gitmap-pipeline-waiting-prompts](file;.agents/skills/gitmap-pipeline-waiting-prompts)
- [ci-cd-fix-with-release](file;.agents/skills/ci-cd-fix-with-release)
- [ai-memory-consolidate-reduce](file;.agents/skills/ai-memory-consolidate-reduce)
- [ai-verification](file;.agents/skills/ai-verification)
- [coding-guidelines](file;.agents/skills/coding-guidelines)
- [canonical-folder-structure](file;.agents/skills/canonical-folder-structure)
