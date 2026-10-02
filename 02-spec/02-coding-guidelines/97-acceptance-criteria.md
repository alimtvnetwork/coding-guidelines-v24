# Coding Guidelines — Master Acceptance Criteria Registry (AI Execution Prompt)

> **/goal** Establish a unified, master registry of testable acceptance criteria across all coding guideline domains in `02-spec/02-coding-guidelines/`.
> **/learn** Enforce the canonical criteria taxonomy (`AC-CG-[CATEGORY]-[NUM]`), ensure every criterion maps to a concrete specification file, and verify compliance via automated linters.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Verify each criterion maps 1:1 to an authoritative specification file in `02-spec/02-coding-guidelines/`.
- [ ] `/learn` Ensure all criteria follow the structured `Given / When / Then` verification contract.
- [ ] `/goal` Validate that all verification commands execute cleanly with `Expected: exit 0`.
- [ ] `/learn` Verify 100% relative paths and zero absolute filesystem paths across all guideline files.

. **CRITICAL AI INSTRUCTION:** This specification is an active AI execution directive. All code generated or modified must strictly follow the rules below.

**Version:** 4.0.0
**Last Updated:** 2026-10-02
**Status:** Active
**AI Confidence:** Production-Ready

---

## Overview

Consolidated master registry index of testable criteria across all guideline categories, referencing category-specific acceptance criteria registries.

---

## AC-01: Cross-Language Standards Registry

See [`01-cross-language/97-acceptance-criteria.md`](./01-cross-language/97-acceptance-criteria.md) for the complete 40+ testable criteria inventory across Boolean principles, code style, naming conventions, type safety, and architecture.

| # | Criterion | Category | Source |
|---|-----------|----------|--------|
| AC-CG-BOOL-001 | Boolean principles define naming (`isX`, `hasX`) and affirmative implicit evaluation patterns | Boolean | [`01-cross-language/02-boolean-principles/readme.md`](./01-cross-language/02-boolean-principles/readme.md) |
| AC-CG-TYPE-004 | Casting elimination patterns cover type-safe alternatives to type assertions | Type Safety | [`01-cross-language/04-casting-elimination-patterns.md`](./01-cross-language/04-casting-elimination-patterns.md) |
| AC-CG-STYLE-001 | Code style defines formatting, naming, vertical spacing, and structural conventions | Code Style | [`01-cross-language/04-code-style/readme.md`](./01-cross-language/04-code-style/readme.md) |
| AC-CG-STYLE-002 | Zero nesting, guard clauses, and early return patterns | Code Style | [`01-cross-language/04-code-style/02-braces-and-nesting.md`](./01-cross-language/04-code-style/02-braces-and-nesting.md) |
| AC-CG-ARCH-008 | DRY principles documented with refactoring patterns and modular extraction | Architecture | [`01-cross-language/08-dry-principles.md`](./01-cross-language/08-dry-principles.md) |
| AC-CG-ARCH-006 | Cyclomatic complexity limits defined with enforcement rules (<= 10) | Architecture | [`01-cross-language/06-cyclomatic-complexity.md`](./01-cross-language/06-cyclomatic-complexity.md) |
| AC-CG-NAME-022 | Variable and collection naming conventions (camelCase, plural arrays, map prefixes) | Naming | [`01-cross-language/22-variable-naming-conventions.md`](./01-cross-language/22-variable-naming-conventions.md) |
| AC-CG-REG-001 | Detailed Cross-Language Criteria Registry (BOOL, STYLE, NAME, TYPE, ARCH, TEST, STATIC) | Registry | [`01-cross-language/97-acceptance-criteria.md`](./01-cross-language/97-acceptance-criteria.md) |

---

## AC-02: TypeScript Standards

| # | Criterion | Source |
|---|-----------|--------|
| AC-CG-TS-001 | Connection status enums define all valid states with TypeScript string literals | [`02-typescript/`](./02-typescript/readme.md) |
| AC-CG-TS-002 | Type definitions avoid `any` and use proper generic constraints | [`02-typescript/`](./02-typescript/readme.md) |
| AC-CG-TS-003 | React component patterns follow functional component with hooks style | [`02-typescript/`](./02-typescript/readme.md) |
| AC-CG-TS-004 | State management patterns use Zustand stores with typed selectors | [`02-typescript/`](./02-typescript/readme.md) |

---

## AC-03: Golang Standards

| # | Criterion | Source |
|---|-----------|--------|
| AC-CG-GO-001 | Boolean standards define naming and evaluation patterns per Go idioms | [`03-golang/readme.md`](./03-golang/readme.md) |
| AC-CG-GO-002 | Error handling uses `*appfault.AppError` and monadic `Result[T]` pattern consistently | [`03-golang/readme.md`](./03-golang/readme.md) |
| AC-CG-GO-003 | Defer rules prevent defer within tight loops and enforce explicit resource closing | [`03-golang/05-defer-rules.md`](./03-golang/05-defer-rules.md) |
| AC-CG-GO-004 | Service layer follows interface-based dependency injection | [`03-golang/readme.md`](./03-golang/readme.md) |

---

## AC-04: PHP Standards

| # | Criterion | Source |
|---|-----------|--------|
| AC-CG-PHP-001 | Class naming follows WordPress PSR-4 autoloading conventions | [`04-php/readme.md`](./04-php/readme.md) |
| AC-CG-PHP-002 | Database queries use $wpdb prepared statements exclusively | [`04-php/readme.md`](./04-php/readme.md) |
| AC-CG-PHP-003 | Type declarations (parameter + return types) required on all functions | [`04-php/readme.md`](./04-php/readme.md) |
| AC-CG-PHP-004 | Input sanitization and output escaping follow WordPress security standards | [`04-php/readme.md`](./04-php/readme.md) |

---

## AC-05: Rust Standards

| # | Criterion | Source |
|---|-----------|--------|
| AC-CG-RUST-001 | Naming conventions follow Rust idioms (snake_case for functions, PascalCase for types) | [`05-rust/readme.md`](./05-rust/readme.md) |
| AC-CG-RUST-002 | Error handling uses `Result<T, E>` pattern with custom error types | [`05-rust/readme.md`](./05-rust/readme.md) |
| AC-CG-RUST-003 | Async patterns use tokio runtime with proper cancellation handling | [`05-rust/readme.md`](./05-rust/readme.md) |
| AC-CG-RUST-004 | Memory safety patterns documented for FFI boundaries | [`05-rust/readme.md`](./05-rust/readme.md) |

---

## Cross-References

- [Coding Guidelines Overview](./readme.md)
- [Cross-Language Standards](./01-cross-language/readme.md)
- [Cross-Language Acceptance Criteria Registry](./01-cross-language/97-acceptance-criteria.md)
- [Golang Standards](./03-golang/readme.md)
- [PHP Standards](./04-php/readme.md)
- [Rust Standards](./05-rust/readme.md)
- [C# Standards](./07-csharp/readme.md)
- [AI Optimization](./06-ai-optimization/readme.md)
