# C# Coding Standards — Changelog (AI Execution Prompt)

> **/goal** Maintain an immutable record of C# language standards, nullable reference type rules, async patterns, and StyleCop linter enforcement.
> **/learn** Master C# conventions: PascalCase methods, `I` prefix interfaces, boolean flag splitting, record types for DTOs, and pattern matching.

**Version:** 3.2.0
**Last Updated:** 2026-04-16
**Status:** Active
**AI Confidence:** Production-Ready
**Ambiguity:** None

---

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Document all additions to C# coding standards, cross-language integrations, and StyleCop analyzer rules.
- [ ] `/learn` Verify cross-references to cross-language guidelines (`01-cross-language/24-`, `25-`) and AI checklists remain valid relative paths.
- [ ] `/goal` Ensure bracketed SemVer formatting (`## [X.Y.Z] — YYYY-MM-DD`) is strictly maintained.
- [ ] `/learn` Audit formatting with `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/07-csharp --check-only`.

. **CRITICAL AI INSTRUCTION:** Active AI execution record for C# standards. Document all updates to C# guidelines here.

---

## [1.0.0] — 2026-04-02

### Added

- `readme.md` — C# coding standards overview with cross-references
- `01-naming-and-conventions.md` — PascalCase methods, `I` prefix interfaces, abbreviation casing, boolean naming
- `02-method-design.md` — Boolean flag splitting, function size limits, async patterns, LINQ usage
- `03-error-handling.md` — Specific exception catching, guard clauses, nullable reference types
- `04-type-safety.md` — Generics over object, pattern matching, records for DTOs, no magic strings
- `97-acceptance-criteria.md` — 30+ testable checks across 7 acceptance categories
- `99-consistency-report.md` — Initial health report (A+)

### Cross-Language Integration

- Added C# examples to `01-cross-language/24-boolean-flag-methods.md`
- Added C# examples to `01-cross-language/25-generic-return-types.md`
- Added 6 C#-specific checks to `06-ai-optimization/02-ai-quick-reference-checklist.md`
- Added C# column to README key standards table

---

## Verification & Acceptance Criteria

_Auto-generated section — see `02-spec/02-coding-guidelines/07-csharp/97-acceptance-criteria.md` for the full criteria index._

### AC-CG-LOG-CS: C# Standards Changelog Conformance

**Given** The C# standards changelog file in `02-spec/02-coding-guidelines/07-csharp/98-changelog.md`.
**When** Audited by repository linters and guideline autofixers.
**Then** The file strictly adheres to the 4-part prompt anatomy, preserves C# StyleCop, async, and nullable reference type release notes, and passes linter checks with exit code 0.

**Verification command:**
```bash
python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/07-csharp --check-only
```
**Expected:** exit 0. Zero violations detected.
