# Golang Standards — Changelog (AI Execution Prompt)

> **/goal** Track version revisions, package standards, integer enum specifications, and error handling updates across Golang coding guidelines.
> **/learn** Internalize Go-specific requirements: *appfault.AppError return types, zero bare void returns, integer-backed enums with PascalCase serialization, and parameter structs.

**Version:** 3.2.0
**Last Updated:** 2026-04-16
**Status:** Active
**AI Confidence:** Production-Ready
**Ambiguity:** None

---

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Log all Go standard updates, including integer enum rules, appfault error conventions, and struct sizing rules.
- [ ] `/learn` Ensure retrospective references point to valid relative paths (e.g. `02-spec/03-error-manage/...`).
- [ ] `/goal` Enforce proper SemVer numbering and ISO date formats on all new changelog sections.
- [ ] `/learn` Audit formatting with `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/03-golang --check-only`.

. **CRITICAL AI INSTRUCTION:** Active AI execution record for Go standards. Log all changes to Go specifications here.

---

All notable changes to the Golang Standards specification are documented here.

---

## v2.2.0 — 2026-09-01

### Added — Integer-Backed Enums & PascalCase Serialization Standard

- Enforced integer underlying types (`byte`, `uint16`, `uint32`) for all Go enums.
- Mandated PascalCase string outputs (`"Info"`, `"Warn"`, `"Error"`, `"Critical"`, `"Fatal"`) via `String()` and custom JSON/YAML marshaling (`MarshalJSON()`, `UnmarshalJSON()`).
- Documented in retrospective `02-spec/03-error-manage/01-error-resolution/03-retrospectives/07-golang-integer-enums-and-pascal-serialization.md`.

---

## v2.1.0 — 2026-03-31

### Changed

- `04-golang-standards-reference.md` split into subfolder (6 files, max 362 lines — down from 1,280)
- Deduplicated enum content in `05-enums-and-dry.md` — now links to `01-enum-specification/` as canonical source
- Fixed spacing violations in code examples

---

## v2.0.0 — 2026-03-09

### Global Version Bump

Project-wide major version increment (+1.0.0) applied to all specification files in `03-coding-guidelines/03-golang`.

#### Changed

- All spec files received a major version bump and date update to 2026-03-09.
- Part of a global effort spanning ~638 files across all 30+ spec folders, establishing a new project-wide versioning baseline.

---

*Keep this file updated when specs change.*
