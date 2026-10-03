# Changelog: Static Analysis & Linter Enforcement (AI Execution Prompt)

> **/goal** Track version revisions, rule matrices, and CI quality gate additions for static analysis and linter enforcement specifications.
> **/learn** Understand SonarQube rule mappings (S1126, S4144) across 8 supported languages, and ensure CI pipeline templates stay synchronized with guideline standards.

**Version:** 1.2.0
**Last Updated:** 2026-04-01
**Status:** Active
**AI Confidence:** Production-Ready
**Ambiguity:** None

---

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Log updates to multi-language static analysis configurations and SonarQube rule matrices.
- [ ] `/learn` Verify all referenced linter specs (golangci-lint, phpcs/phpstan, stylecop, clippy, ruff, eslint) maintain active relative links.
- [ ] `/goal` Confirm new entries follow the standard bracketed SemVer format (`## [X.Y.Z] — YYYY-MM-DD`).
- [ ] `/learn` Audit formatting with `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/01-cross-language/16-static-analysis --check-only`.

. **CRITICAL AI INSTRUCTION:** This changelog tracks static analysis enforcement standards. Changes to linter rules or pipelines must be recorded here.

---

All notable changes to the `16-static-analysis/` subfolder.

---

## [1.2.0] — 2026-04-01

### Added

- `10-cross-language-rule-matrix.md` — side-by-side SonarQube rule mapping across all 8 languages
- `97-acceptance-criteria.md` — acceptance criteria for the subfolder
- `98-changelog.md` — this file

## [1.1.0] — 2026-04-01

### Added

- `09-ci-pipeline-quality-gate.md` — unified CI pipeline spec with GitHub Actions and GitLab CI templates
- `99-consistency-report.md` — initial consistency report

### Changed

- All 8 language specs bumped to v1.1.0 — standardized Keywords/Scoring sections, integration checklist format, and added missing SonarQube rules (S1126, S4144)
- `readme.md` bumped to v1.1.0 — added CI pipeline to inventory

## [1.0.0] — 2026-03-31

### Added

- `readme.md` — subfolder overview with document inventory and rule mapping table
- `02-go-golangci-lint.md` — Go static analysis spec
- `03-php-phpcs-phpstan.md` — PHP static analysis spec
- `04-csharp-stylecop.md` — C# static analysis spec
- `05-rust-clippy.md` — Rust static analysis spec
- `06-vb-dotnet-analyzers.md` — VB.NET static analysis spec
- `07-nodejs-eslint.md` — Node.js static analysis spec
- `08-python-ruff.md` — Python static analysis spec
