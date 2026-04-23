# Memory Index

**Updated:** 2026-04-23 (session 2 — readme CODE-RED walkthrough)

# Project Memory

## Core
- 🔴 CODE RED: Never swallow errors. Zero-nesting (no nested if). Max 2 operands. Positively named guard functions.
- 🔴 CODE RED: Strict metrics: functions 8-15 lines, files < 300 lines, React components < 100 lines.
- 🔴 NEVER sync 01-app, 02-app-issues, 03-general, 03-tasks, or 12-consolidated-guidelines from upstream sibling repos. All maintained locally.
- 🔴 SKIP from spec audits: `21-app`, `22-app-issues`, `23-app-database`, `24-app-design-system-and-ui` are intentional stubs. Never write 97/99 files for them, never demote to `_drafts/`, exclude from corpus averages.
- Naming: PascalCase for all internal IDs, DB, JSON, Types. Exceptions: Rust uses snake_case identifiers.
- DB Schema: PascalCase naming. PKs are `{TableName}Id` (INTEGER PRIMARY KEY AUTOINCREMENT). No UUIDs.
- DB Booleans: forbidden `Not`/`No` prefixes; Approved Inverses (`IsDisabled`, `IsInvalid`, `IsIncomplete`, `IsUnavailable`, `IsUnread`, `IsHidden`, `IsBroken`, `IsLocked`, `IsUnpublished`, `IsUnverified`) allowed. Inverses derived in code via Rule 9 codegen.
- DB Descriptive columns: entity tables MUST have `Description TEXT NULL`; transactional tables MUST have `Notes TEXT NULL` and `Comments TEXT NULL`.
- Workflow: Spec-First (`spec/`) and Issue-First (`03-issues/`).
- Repo: `alimtvnetwork/coding-guidelines-v15`. Install scripts live at repo root (`install.ps1` / `install.sh`).
- `.lovable/` structure: single-file convention — `plan.md`, `suggestions.md`, `strictly-avoid.md` each hold their full history. Never create per-task folders.
- Execution: Break complex requests into discrete tasks. Wait for "next" prompt to continue.
- 🔴 INSTALLERS: Two modes only. PINNED (any `--version`/baked tag/`releases/download/<tag>/` URL) = exact tag, NO `releases/latest`, NO `main` fallback, NO cross-repo jump, 404→exit 3. IMPLICIT = latest release → V→V+20 parallel discovery → main (with ⚠️ warning). Spec: `spec/14-update/27-generic-installer-behavior.md`.

## Memory Files

### Workflow
- [Plan Tracker](workflow/01-plan-tracker.md) — Mirrors `.lovable/plan.md`; chronological tracker of completed and pending tasks.

### Suggestions
- [Suggestions Tracker](suggestions/01-suggestions-tracker.md) — Mirrors `.lovable/suggestions.md`; tracks active and implemented suggestions.

### Project
- [Naming Compliance Issues](project/naming-compliance-issues.md)
- [Phase 2 — Content Overlap Audit](project/phase2-content-overlap-audit.md)
- [Phase 3 — Consolidated Structure Design](project/phase3-consolidated-structure-design.md)
- [v2.2 Error Spec Changes](project/v2.2-error-spec-changes.md)
- [README Bundle Installers + GIFs (v3.55.0)](project/03-readme-bundle-installers.md)

### Features
- [Visual Rendering System](features/visual-rendering-system.md)

### Issues
- [Nested Code Fence Rendering](issues/nested-code-fence-rendering.md)

### Done
- [Coding Guidelines Consolidation Plan](done/coding-guidelines-consolidation-plan.md)

### Sessions
- [2026-04-19 — Performance + Boolean Naming + Schema Descriptive Columns](sessions/2026-04-19-perf-boolean-naming-schema.md)
- [2026-04-19 — Distribution & Runner Spec + Slides Sub-Command](sessions/2026-04-19-distribution-runner-slides.md)
- [2026-04-23 — Quickstart + Zero CI/CD Violations](sessions/2026-04-23-quickstart-and-zero-violations.md)
- [2026-04-23 — README CODE-RED Walkthrough + Spec References](sessions/2026-04-23-readme-code-red-walkthrough.md)

### Avoid
- [Avoid Per-Task Folders](avoid/01-avoid-per-task-folders.md) — Why `.lovable/` uses single-file convention; what was removed.

## Memory Pointers (mem://)
- [Axios Pinning](mem://constraints/axios-version-pinning) — Exact pinned versions only (1.14.0/0.30.3). Blocked: 1.14.1, 0.30.4.
- [Database Architecture](mem://architecture/database-schema) — PascalCase naming, no UUIDs, Vw prefixes for views.
- [Error Handling](mem://architecture/error-handling) — `apperror` package, explicit file/path logging required.
- [PowerShell Style](mem://style/powershell-naming) — lowercase-kebab-case files, PascalCase Verb-Noun functions.
- [Development Workflow](mem://processes/development-workflow) — Spec-first workflow, linter enforcement, clean docs.
- [React ForwardRef Warning](mem://constraints/react-app-forwardref-warning) — Ignore lovable.js App.tsx ref console warning.
- [Code Red Guidelines](mem://standards/code-red-guidelines) — Full rules for zero-nesting, booleans, metrics.
- [Standards Enforcement](mem://processes/automated-standards-enforcement) — linter-scripts validation requirements.
- [Naming Conventions](mem://style/naming-conventions) — Zero-Underscore policy, full uppercase acronyms.
- [Caching Policy](mem://architecture/caching-policy) — Explicit TTL, deterministic keys, invalidate on mutation.
- [Nested Code Fences](mem://issues/nested-code-fence-data-corruption) — 4-backtick fences required for nested markdown blocks.
- [TypeScript Patterns](mem://standards/typescript-patterns) — Named interfaces for unions, TypedAction, explicit types.
- [Enum Standards](mem://standards/enum-standards) — Cross-language PascalCase enums, strict parsing methods.
- [Split Database](mem://architecture/split-database) — Root, App, Session hierarchical SQLite with WAL and Casbin.
- [Seedable Config](mem://architecture/seedable-configuration) — SemVer GORM merge of config.seed.json.
- [Self Update Arch](mem://features/self-update-architecture) — Rename-first deployment, atomicity with latest.json.
- [Release-Pinned Installer](mem://features/release-pinned-installer) — Additive release-install.ps1/sh; pins exact version, no latest lookup, no cross-repo drift.
- [Installer Behavior (Generic)](mem://standards/installer-behavior) — Two-mode (PINNED vs IMPLICIT) cross-repo installer contract; V→V+20 parallel discovery; canonical spec at `spec/14-update/27-generic-installer-behavior.md` for sharing with any AI.
- [Doc Standards](mem://project/documentation-standards) — Mandatory numeric folders (01-20 Core, 21+ App), JSON tree syncing.
- [Author Attribution](mem://project/author-attribution) — Md. Alim Ul Karim, Riseup Asia LLC, SEO/footer requirements.
- [Avoid Upstream Sync](mem://constraints/avoid-app-sync) — NEVER sync app/app-issues/general/tasks/consolidated-guidelines from upstream sibling repos.
- [Skip Stub Spec Folders](mem://constraints/skip-stub-spec-folders) — 21-app, 22-app-issues, 23-app-database, 24-app-design-system-and-ui excluded from audit/remediation.
