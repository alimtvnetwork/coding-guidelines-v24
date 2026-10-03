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
**Last Updated:** 2026-10-03
**Status:** Active
**AI Confidence:** Production-Ready

---

## Overview

Consolidated master registry index of testable criteria across all guideline categories, referencing category-specific acceptance criteria registries.

---

## AC-00: Root Style and Sizing Guidelines

| # | Criterion | Category | Source |
|---|-----------|----------|--------|
| AC-CG-STYLE-002 | Canonical Size Tier Enforcement (Function <= 15 lines, File <= 300 lines) | Code Style | [`02-canonical-size-tier.md`](./02-canonical-size-tier.md) |
| AC-CG-STYLE-003 | Core Coding Style and Parameter Rules (Max 3 params, option structs, UTF-8 LF) | Code Style | [`03-coding-style-checklist.md`](./03-coding-style-checklist.md) |
| AC-CG-STYLE-004 | Condensed Review Guide Rules (Zero nesting, no positive/negative mix) | Code Style | [`04-consolidated-review-guide-condensed.md`](./04-consolidated-review-guide-condensed.md) |
| AC-CG-STYLE-005 | Consolidated Master Review Rules (Complete cross-language rule synthesis) | Code Style | [`05-consolidated-review-guide.md`](./05-consolidated-review-guide.md) |

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
| AC-CG-LOG-CROSS | Cross-Language Changelog Conformance | Changelog | [`01-cross-language/98-changelog.md`](./01-cross-language/98-changelog.md) |
| AC-CG-REG-001 | Detailed Cross-Language Criteria Registry (BOOL, STYLE, NAME, TYPE, ARCH, TEST, STATIC) | Registry | [`01-cross-language/97-acceptance-criteria.md`](./01-cross-language/97-acceptance-criteria.md) |

---

## AC-02: TypeScript Standards Registry

See [`02-typescript/97-acceptance-criteria.md`](./02-typescript/97-acceptance-criteria.md) for the complete 16 testable criteria inventory across TypeScript status enums, type-safety plans, discriminated unions, and ESLint enforcement.

| # | Criterion | Source |
|---|-----------|--------|
| AC-CG-TS-000 | TypeScript Guidelines Index Conformance | [`02-typescript/readme.md`](./02-typescript/readme.md) |
| AC-CG-TS-002 | ConnectionStatusEnum Definition and Validation | [`02-typescript/02-connection-status-enum.md`](./02-typescript/02-connection-status-enum.md) |
| AC-CG-TS-003 | EntityStatusEnum Definition and Validation | [`02-typescript/03-entity-status-enum.md`](./02-typescript/03-entity-status-enum.md) |
| AC-CG-TS-004 | ExecutionStatusEnum Definition and Validation | [`02-typescript/04-execution-status-enum.md`](./02-typescript/04-execution-status-enum.md) |
| AC-CG-TS-005 | ExportStatusEnum Definition and Validation | [`02-typescript/05-export-status-enum.md`](./02-typescript/05-export-status-enum.md) |
| AC-CG-TS-006 | HttpMethodEnum Definition and Validation | [`02-typescript/06-http-method-enum.md`](./02-typescript/06-http-method-enum.md) |
| AC-CG-TS-007 | MessageStatusEnum Definition and Validation | [`02-typescript/07-message-status-enum.md`](./02-typescript/07-message-status-enum.md) |
| AC-CG-TS-008 | TypeScript Type Safety Remediation Plan and Any Elimination | [`02-typescript/08-type-safety-remediation-plan.md`](./02-typescript/08-type-safety-remediation-plan.md) |
| AC-CG-TS-009 | TypeScript Standards Reference and Architectural Patterns | [`02-typescript/09-typescript-standards-reference.md`](./02-typescript/09-typescript-standards-reference.md) |
| AC-CG-TS-010 | TypeScript Promise, Async/Await, and Concurrency Patterns | [`02-typescript/10-promise-await-patterns.md`](./02-typescript/10-promise-await-patterns.md) |
| AC-CG-TS-011 | LogLevelEnum Definition and Validation | [`02-typescript/11-log-level-enum.md`](./02-typescript/11-log-level-enum.md) |
| AC-CG-TS-012 | TypeScript ESLint Rules, Type Checking, and Lint Automation | [`02-typescript/12-eslint-enforcement.md`](./02-typescript/12-eslint-enforcement.md) |
| AC-CG-TS-013 | TypeScript Discriminated Unions and Exhaustive Type Narrowing | [`02-typescript/13-discriminated-union-patterns.md`](./02-typescript/13-discriminated-union-patterns.md) |
| AC-CG-TS-014 | TypeScript Enum Runtime Validation and Parse Guard Utilities | [`02-typescript/14-enum-checking-and-validation.md`](./02-typescript/14-enum-checking-and-validation.md) |
| AC-CG-TS-015 | TypeScript State Management, Stores, and Reactive Architecture | [`02-typescript/15-state-management.md`](./02-typescript/15-state-management.md) |
| AC-CG-LOG-TS | TypeScript Standards Changelog Conformance | [`02-typescript/98-changelog.md`](./02-typescript/98-changelog.md) |
| AC-CG-TS-REG-001 | TypeScript Acceptance Criteria Registry Conformance | [`02-typescript/97-acceptance-criteria.md`](./02-typescript/97-acceptance-criteria.md) |

---

## AC-03: Golang Standards Registry

See [`03-golang/97-acceptance-criteria.md`](./03-golang/97-acceptance-criteria.md) for the complete criteria inventory across Go standards, enums, and architecture.

| # | Criterion | Source |
|---|-----------|--------|
| AC-CG-GO-000 | Golang Standards Index Conformance | [`03-golang/readme.md`](./03-golang/readme.md) |
| AC-CG-GO-001 | Positive boolean naming and evaluation patterns | [`03-golang/02-boolean-standards.md`](./03-golang/02-boolean-standards.md) |
| AC-CG-GO-003 | HttpMethod enum standard and string literal bans | [`03-golang/03-httpmethod-enum.md`](./03-golang/03-httpmethod-enum.md) |
| AC-CG-GO-005 | Resource defer rules and loop safety | [`03-golang/05-defer-rules.md`](./03-golang/05-defer-rules.md) |
| AC-CG-GO-006 | Go string and slice internals efficiency | [`03-golang/06-string-slice-internals.md`](./03-golang/06-string-slice-internals.md) |
| AC-CG-GO-007 | Code severity taxonomy and fault logging | [`03-golang/07-code-severity-taxonomy.md`](./03-golang/07-code-severity-taxonomy.md) |
| AC-CG-GO-008 | Unified PathUtil and FileUtil specification | [`03-golang/08-pathutil-fileutil-spec.md`](./03-golang/08-pathutil-fileutil-spec.md) |
| AC-CG-GO-009 | Wrapped boolean results and monadic error returns | [`03-golang/09-wrapped-boolean-results.md`](./03-golang/09-wrapped-boolean-results.md) |
| AC-CG-GO-ENUM-000 | Go Enum Specification Registry & Patterns | [`03-golang/01-enum-specification/readme.md`](./03-golang/01-enum-specification/readme.md) |
| AC-CG-GO-REF-000 | Go Standards Reference (Sizing, Types, DB, Naming, Concurrency) | [`03-golang/04-golang-standards-reference/readme.md`](./03-golang/04-golang-standards-reference/readme.md) |
| AC-CG-LOG-GO | Golang Standards Changelog Conformance | [`03-golang/98-changelog.md`](./03-golang/98-changelog.md) |

---

## AC-04: PHP Standards Registry

See [`04-php/97-acceptance-criteria.md`](./04-php/97-acceptance-criteria.md) for the complete criteria inventory across PHP backed enums, response arrays, PSR-4 naming, and standards reference.

| # | Criterion | Source |
|---|-----------|--------|
| AC-CG-PHP-000 | PHP Guidelines Directory Index Conformance | [`04-php/readme.md`](./04-php/readme.md) |
| AC-CG-PHP-001 | Class naming follows WordPress PSR-4 autoloading conventions | [`04-php/readme.md`](./04-php/readme.md) |
| AC-CG-PHP-002 | Database queries use $wpdb prepared statements exclusively | [`04-php/readme.md`](./04-php/readme.md) |
| AC-CG-PHP-003 | Type declarations (parameter + return types) required on all functions | [`04-php/readme.md`](./04-php/readme.md) |
| AC-CG-PHP-004 | Input sanitization and output escaping follow WordPress security standards | [`04-php/readme.md`](./04-php/readme.md) |
| AC-CG-LOG-PHP | PHP Standards Changelog Conformance | [`04-php/98-changelog.md`](./04-php/98-changelog.md) |
| AC-CG-PHP-REG-001 | PHP Acceptance Criteria Registry Conformance | [`04-php/97-acceptance-criteria.md`](./04-php/97-acceptance-criteria.md) |

---

## AC-05: Rust Standards Registry

See [`05-rust/97-acceptance-criteria.md`](./05-rust/97-acceptance-criteria.md) for the complete criteria inventory across Rust naming, error handling, async Tokio, memory safety, and FFI boundaries.

| # | Criterion | Source |
|---|-----------|--------|
| AC-CG-RUST-000 | Rust Guidelines Directory Index Conformance | [`05-rust/readme.md`](./05-rust/readme.md) |
| AC-CG-RUST-001 | Naming conventions follow Rust idioms (snake_case for functions, PascalCase for types) | [`05-rust/readme.md`](./05-rust/readme.md) |
| AC-CG-RUST-002 | Error handling uses `Result<T, E>` pattern with custom error types | [`05-rust/readme.md`](./05-rust/readme.md) |
| AC-CG-RUST-003 | Async patterns use tokio runtime with proper cancellation handling | [`05-rust/readme.md`](./05-rust/readme.md) |
| AC-CG-RUST-004 | Memory safety patterns documented for FFI boundaries | [`05-rust/readme.md`](./05-rust/readme.md) |
| AC-CG-LOG-RUST | Rust Standards Changelog Conformance | [`05-rust/98-changelog.md`](./05-rust/98-changelog.md) |
| AC-CG-RUST-REG-001 | Rust Acceptance Criteria Registry Conformance | [`05-rust/97-acceptance-criteria.md`](./05-rust/97-acceptance-criteria.md) |

---

## AC-06: AI Optimization Registry

See [`06-ai-optimization/97-acceptance-criteria.md`](./06-ai-optimization/97-acceptance-criteria.md) for criteria covering anti-hallucination, citation requirements, quick reference checklists, and agent memory lifecycles.

| # | Criterion | Source |
|---|-----------|--------|
| AC-CG-AI-000 | AI Optimization Guidelines Index Conformance | [`06-ai-optimization/readme.md`](./06-ai-optimization/readme.md) |
| AC-CG-AI-002 | Anti-hallucination verification and source attribution | [`06-ai-optimization/02-anti-hallucination-rules.md`](./06-ai-optimization/02-anti-hallucination-rules.md) |
| AC-CG-AI-006 | Strict relative path citation and markdown link verification | [`06-ai-optimization/06-citation-requirement.md`](./06-ai-optimization/06-citation-requirement.md) |
| AC-CG-AI-009 | Agent memory lifecycle, persistent knowledge, and index updates | [`06-ai-optimization/09-agent-memory-lifecycle.md`](./06-ai-optimization/09-agent-memory-lifecycle.md) |
| AC-CG-AI-REG-001 | AI Optimization Acceptance Criteria Registry Conformance | [`06-ai-optimization/97-acceptance-criteria.md`](./06-ai-optimization/97-acceptance-criteria.md) |

---

## AC-07: CI/CD Integration Registry

See [`06-cicd-integration/97-acceptance-criteria.md`](./06-cicd-integration/97-acceptance-criteria.md) and [`06-cicd-integration/08-fix-repo-and-installers/97-acceptance-criteria.md`](./06-cicd-integration/08-fix-repo-and-installers/97-acceptance-criteria.md) for SARIF, plugin contracts, fix repo automation, and quality gates.

| # | Criterion | Source |
|---|-----------|--------|
| AC-CG-CICD-000 | CI/CD Integration Directory Index Conformance | [`06-cicd-integration/readme.md`](./06-cicd-integration/readme.md) |
| AC-CG-CICD-002 | SARIF output contract and schema compliance | [`06-cicd-integration/02-sarif-contract.md`](./06-cicd-integration/02-sarif-contract.md) |
| AC-CG-FIX-000 | Fix Repo & Installer Specifications Registry | [`06-cicd-integration/08-fix-repo-and-installers/readme.md`](./06-cicd-integration/08-fix-repo-and-installers/readme.md) |
| AC-CG-CICD-REG-001 | CI/CD Acceptance Criteria Registry Conformance | [`06-cicd-integration/97-acceptance-criteria.md`](./06-cicd-integration/97-acceptance-criteria.md) |

---

## AC-08: C# Standards Registry

See [`07-csharp/97-acceptance-criteria.md`](./07-csharp/97-acceptance-criteria.md) for C# naming conventions, method design, error handling, and type safety standards.

| # | Criterion | Source |
|---|-----------|--------|
| AC-CG-CS-000 | C# Standards Directory Index Conformance | [`07-csharp/readme.md`](./07-csharp/readme.md) |
| AC-CG-CS-002 | C# Naming Conventions and PascalCase Types | [`07-csharp/02-naming-and-conventions.md`](./07-csharp/02-naming-and-conventions.md) |
| AC-CG-CS-003 | Method Design, Parameter Limits, and Pure Expressions | [`07-csharp/03-method-design.md`](./07-csharp/03-method-design.md) |
| AC-CG-CS-004 | Structured Exceptions and Zero Swallowed Faults | [`07-csharp/04-error-handling.md`](./07-csharp/04-error-handling.md) |
| AC-CG-LOG-CS | C# Standards Changelog Conformance | [`07-csharp/98-changelog.md`](./07-csharp/98-changelog.md) |
| AC-CG-CS-REG-001 | C# Acceptance Criteria Registry Conformance | [`07-csharp/97-acceptance-criteria.md`](./07-csharp/97-acceptance-criteria.md) |

---

## AC-09: File & Folder Naming Registry

See [`08-file-folder-naming/97-acceptance-criteria.md`](./08-file-folder-naming/97-acceptance-criteria.md) for cross-language kebab-case and lowercase directory hygiene.

| # | Criterion | Source |
|---|-----------|--------|
| AC-CG-FILE-000 | File & Folder Naming Directory Index Conformance | [`08-file-folder-naming/readme.md`](./08-file-folder-naming/readme.md) |
| AC-CG-FILE-002 | Cross-language lowercase kebab-case naming standard | [`08-file-folder-naming/02-cross-language.md`](./08-file-folder-naming/02-cross-language.md) |
| AC-CG-FILE-REG-001 | File Naming Acceptance Criteria Registry Conformance | [`08-file-folder-naming/97-acceptance-criteria.md`](./08-file-folder-naming/97-acceptance-criteria.md) |

---

## AC-10: Security Guidelines Registry

See [`11-security/97-acceptance-criteria.md`](./11-security/97-acceptance-criteria.md) for JWT lifecycles, encryption standards, OWASP mitigation, secret vaulting, and dependency pinning.

| # | Criterion | Source |
|---|-----------|--------|
| AC-CG-SEC-000 | Security Guidelines Directory Index & Policy Conformance | [`11-security/readme.md`](./11-security/readme.md) |
| AC-CG-SEC-001 | JWT Token Lifecycle & HttpOnly Cookie Storage | [`11-security/02-jwt-standards.md`](./11-security/02-jwt-standards.md) |
| AC-CG-SEC-002 | Encryption at Rest & Key Derivation Standards | [`11-security/03-encryption-standards.md`](./11-security/03-encryption-standards.md) |
| AC-CG-SEC-003 | OWASP Top 10 Mitigation Controls & Input Validation | [`11-security/04-owasp-top-10.md`](./11-security/04-owasp-top-10.md) |
| AC-CG-SEC-004 | Zero Secrets in Source Control & Environment Vaulting | [`11-security/05-secret-management.md`](./11-security/05-secret-management.md) |
| AC-CG-SEC-AXIOS-000 | Axios Client Security Overview | [`11-security/01-axios-version-control/readme.md`](./11-security/01-axios-version-control/readme.md) |
| AC-CG-SEC-AXIOS-001 | Strict Axios Version Pinning & Dependency Locking | [`11-security/01-axios-version-control/02-implementation-rules.md`](./11-security/01-axios-version-control/02-implementation-rules.md) |
| AC-CG-SEC-AXIOS-002 | CVE Remediation & Supply Chain Security Verification | [`11-security/01-axios-version-control/03-security-notes.md`](./11-security/01-axios-version-control/03-security-notes.md) |
| AC-CG-SEC-REG-001 | Security Acceptance Criteria Registry Conformance | [`11-security/97-acceptance-criteria.md`](./11-security/97-acceptance-criteria.md) |

---

## AC-11: Application Architecture & Module Readmes

| # | Criterion | Source |
|---|-----------|--------|
| AC-CG-APP-001 | Application Specifications Index & Navigation | [`21-app/readme.md`](./21-app/readme.md) |
| AC-CG-APP-002 | Non-CI/CD Application Issues & Bug Catalog Index | [`22-app-issues/readme.md`](./22-app-issues/readme.md) |
| AC-CG-APP-003 | Application Database Standards & Split SQLite Architecture | [`23-app-db/readme.md`](./23-app-db/readme.md) |
| AC-CG-APP-004 | Application UI/UX Design System Specifications | [`24-app-ui-design-system/readme.md`](./24-app-ui-design-system/readme.md) |
| AC-CG-CONSISTENCY-001 | Guideline Consistency & Structure Conformance | [`99-consistency-report.md`](./99-consistency-report.md) |

---

## AC-12: Python Standards Registry

See [`12-python/97-acceptance-criteria.md`](./12-python/97-acceptance-criteria.md) for the complete criteria inventory across Python static typing, Pydantic data validation, PEP-8 compliance, and specific exceptions.

| # | Criterion | Source |
|---|-----------|--------|
| AC-CG-PY-001 | Python Guidelines Directory Index Conformance | [`12-python/readme.md`](./12-python/readme.md) |
| AC-CG-PY-002 | Python Coding Standards Conformance | [`12-python/02-standards.md`](./12-python/02-standards.md) |
| AC-CG-PY-003 | Python Dynamic Enum & Array Constants Standard | [`12-python/02-standards.md`](./12-python/02-standards.md) |
| AC-CG-PY-004 | Python DRY Architecture & Engine Caching Conformance | [`12-python/02-standards.md`](./12-python/02-standards.md) |
| AC-CG-PY-REG-001 | Python Acceptance Criteria Registry Conformance | [`12-python/97-acceptance-criteria.md`](./12-python/97-acceptance-criteria.md) |

---

## AC-13: Modern C++ Standards Registry

See [`13-cpp/97-acceptance-criteria.md`](./13-cpp/97-acceptance-criteria.md) for the complete criteria inventory across C++20 baseline, concepts, RAII smart pointers, Rule of Zero/Five, and FFI exception boundaries.

| # | Criterion | Source |
|---|-----------|--------|
| AC-CG-CPP-001 | Modern C++ Guidelines Directory Index Conformance | [`13-cpp/readme.md`](./13-cpp/readme.md) |
| AC-CG-CPP-002 | Modern C++ Standards Conformance | [`13-cpp/02-standards.md`](./13-cpp/02-standards.md) |
| AC-CG-CPP-003 | C++ Memory Safety & RAII Resource Management | [`13-cpp/02-standards.md`](./13-cpp/02-standards.md) |
| AC-CG-CPP-004 | C++ FFI Boundary Exception Safety & Standard Types | [`13-cpp/02-standards.md`](./13-cpp/02-standards.md) |
| AC-CG-CPP-REG-001 | Modern C++ Acceptance Criteria Registry Conformance | [`13-cpp/97-acceptance-criteria.md`](./13-cpp/97-acceptance-criteria.md) |

---

## Cross-References

- [Coding Guidelines Overview](./readme.md)
- [Cross-Language Standards](./01-cross-language/readme.md)
- [Cross-Language Acceptance Criteria Registry](./01-cross-language/97-acceptance-criteria.md)
- [TypeScript Standards](./02-typescript/readme.md)
- [TypeScript Acceptance Criteria Registry](./02-typescript/97-acceptance-criteria.md)
- [Golang Standards](./03-golang/readme.md)
- [Golang Acceptance Criteria Registry](./03-golang/97-acceptance-criteria.md)
- [PHP Standards](./04-php/readme.md)
- [PHP Acceptance Criteria Registry](./04-php/97-acceptance-criteria.md)
- [Rust Standards](./05-rust/readme.md)
- [Rust Acceptance Criteria Registry](./05-rust/97-acceptance-criteria.md)
- [AI Optimization Standards](./06-ai-optimization/readme.md)
- [AI Optimization Acceptance Criteria Registry](./06-ai-optimization/97-acceptance-criteria.md)
- [CI/CD Integration Standards](./06-cicd-integration/readme.md)
- [CI/CD Acceptance Criteria Registry](./06-cicd-integration/97-acceptance-criteria.md)
- [C# Standards](./07-csharp/readme.md)
- [C# Acceptance Criteria Registry](./07-csharp/97-acceptance-criteria.md)
- [File & Folder Naming](./08-file-folder-naming/readme.md)
- [File & Folder Naming Acceptance Criteria Registry](./08-file-folder-naming/97-acceptance-criteria.md)
- [Security Guidelines](./11-security/readme.md)
- [Security Acceptance Criteria Registry](./11-security/97-acceptance-criteria.md)
- [Python Standards](./12-python/readme.md)
- [Python Acceptance Criteria Registry](./12-python/97-acceptance-criteria.md)
- [Modern C++ Standards](./13-cpp/readme.md)
- [Modern C++ Acceptance Criteria Registry](./13-cpp/97-acceptance-criteria.md)
