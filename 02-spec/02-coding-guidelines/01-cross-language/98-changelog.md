# Coding Guidelines — Changelog (AI Execution Prompt)

> **/goal** Maintain an immutable, sequential audit trail of all cross-language architectural revisions, guideline refactorings, and rule additions across 02-spec/02-coding-guidelines/01-cross-language/.
> **/learn** Adhere to Keep-a-Changelog SemVer conventions, document structural migrations across subfolders, and verify all cross-language rule cross-references.

**Version:** 3.2.0
**Last Updated:** 2026-04-16
**Status:** Active
**AI Confidence:** Production-Ready
**Ambiguity:** None

---

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Record every cross-language rule addition, split, or modification under the appropriate SemVer release header.
- [ ] `/learn` Ensure deduplicated rules (e.g. enum specifications in 06-ai-optimization/) maintain canonical cross-references.
- [ ] `/goal` Enforce strict relative git paths for all referenced spec files (e.g., `02-spec/02-coding-guidelines/...`).
- [ ] `/learn` Verify zero syntax or spacing violations via `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/01-cross-language --check-only`.

. **CRITICAL AI INSTRUCTION:** This changelog is an active AI execution record. Any updates to cross-language guidelines must be logged here before concluding the task.

---

All notable changes to the Coding Guidelines specification are documented here.

---

## v3.2.0 — 2026-03-31

### Structural Improvements

#### Changed

- `02-boolean-principles.md` split into subfolder (5 files, max 262 lines)
- `15-master-coding-guidelines.md` split into subfolder (7 files, max 277 lines)
- Deduplicated enum rules — `06-ai-optimization/05-enum-naming-quick-reference.md` is now the single cross-language enum source
- Fixed 229 spacing violations in code examples (R4, R5, R10 rules)
- Fixed all broken anchor links across 48 files
- Updated cross-references to point to new subfolder locations

---

## v3.0.0 — 2026-03-31

### Phase 4 Rules Added to Master Guidelines

#### Changed

- `15-master-coding-guidelines/readme.md` bumped to **v2.0.0**
- Added 7 new sections (§14–§20): Lazy Evaluation, Regex Usage, Code Mutation Avoidance, Null Pointer Safety, Nesting Resolution, Newline Styling, Defer Rules (Go)
- Expanded Quick Checklist with 7 new items covering mutation, regex, lazy eval, defer, nesting, newlines, null safety
- Added cross-references to Phase 4 spec files (16–21) in "How to Use" section

---

## v2.1.0 — 2026-03-11

### Added

- `14-test-naming-and-03-structure.md` — New spec covering test file naming, three-part test function naming convention, table-driven test rules, test helper placement, AAA pattern, test isolation, and integration test boundaries. Applies to Go, TypeScript, and PHP.

---

## v2.0.0 — 2026-03-09

### Global Version Bump

Project-wide major version increment (+1.0.0) applied to all specification files in `03-coding-guidelines/01-cross-language`.

#### Changed

- All spec files received a major version bump and date update to 2026-03-09.
- Part of a global effort spanning ~638 files across all 30+ spec folders, establishing a new project-wide versioning baseline.

---

*Keep this file updated when specs change.*
