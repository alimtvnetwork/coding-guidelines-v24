# Subtask 01: Standardize Coding Guideline Changelog Files with AI Execution Prompt Headers & Actionable Checklists

> **Parent Spec:** [`02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/01-architecture-spec.md`](../../../../02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/01-architecture-spec.md)
> **Status:** `COMPLETED`
> **Traceability ID:** `Task-01-Standardize-Changelog-Checklists`
> **Worker Git Policy:** `STRICT WORKER GIT BAN` (Workers MUST NOT execute git commands; only lead orchestrators commit).
> **Target Files (Disjoint Boundary):**
> 1. `02-spec/02-coding-guidelines/01-cross-language/98-changelog.md`
> 2. `02-spec/02-coding-guidelines/01-cross-language/16-static-analysis/98-changelog.md`
> 3. `02-spec/02-coding-guidelines/02-typescript/98-changelog.md`
> 4. `02-spec/02-coding-guidelines/03-golang/98-changelog.md`
> 5. `02-spec/02-coding-guidelines/04-php/98-changelog.md`
> 6. `02-spec/02-coding-guidelines/05-rust/98-changelog.md`
> 7. `02-spec/02-coding-guidelines/07-csharp/98-changelog.md`

---

## 🎯 Actionable CI/CD & Agent Checklist

- [x] `/goal` Inject the canonical AI Execution Prompt header (`> **/goal**` and `> **/learn**`) into all 7 changelog files.
- [x] `/learn` Inject the standardized `## 🎯 Actionable CI/CD & Agent Checklist` block immediately below the header/metadata in all 7 changelogs.
- [x] `/goal` Preserve 100% of historical release logs, SemVer sections, and change summaries without content loss or line truncation.
- [x] `/learn` Verify zero absolute filesystem paths or `file:///` URIs, and confirm compliance using `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only` (exit code 0).

. **CRITICAL AI INSTRUCTION:** Workers modifying these 7 files must operate within this isolated bounding box and observe the strict worker git ban.

---

## 1. Subtask Objective & Context

Under the global Prompt Architect standard established in [`02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/01-architecture-spec.md`](../../../../02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/01-architecture-spec.md), every document within `02-spec/02-coding-guidelines/` must function as an authoritative, active AI execution prompt.

Currently, the 7 historical changelog files (`98-changelog.md`) across the language modules and subfolders are structured as passive text files. They lack:
1. The top-level AI Execution Prompt header (`> **/goal**` and `> **/learn**`).
2. The actionable checklist (`## 🎯 Actionable CI/CD & Agent Checklist`) with `/goal` and `/learn` directives.
3. The `. **CRITICAL AI INSTRUCTION:**` directive block guiding autonomous agents during changelog updates.

This subtask directs the surgical standardization of all 7 changelog files, bringing them into 100% alignment with the architectural contract without altering historical release notes.

---

## 2. Standardized Header & Checklist Template for Changelogs

Every changelog file must be updated to follow this exact structural template at the top of the file:

```markdown
# [Module Name] — Changelog (AI Execution Prompt)

> **/goal** Maintain an immutable, sequential audit trail of all architectural revisions, version increments, and rule migrations across [module/subfolder] coding guidelines.
> **/learn** Adhere to Keep-a-Changelog SemVer conventions, document deprecations and breaking changes explicitly, and verify all relative file links point to valid specification paths.

**Version:** [Preserve existing version, e.g. 3.2.0]
**Last Updated:** [Preserve existing date, e.g. 2026-04-16]
**Status:** Active
**AI Confidence:** Production-Ready
**Ambiguity:** None

---

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Verify new version entries follow SemVer format with ISO-8601 release dates (`vX.Y.Z — YYYY-MM-DD`).
- [ ] `/learn` Ensure changelog categorizations strictly use standard headings (Added, Changed, Deprecated, Removed, Fixed, Security).
- [ ] `/goal` Verify all cross-referenced guideline paths use strict relative repository paths without broken anchors.
- [ ] `/learn` Confirm clean vertical formatting and zero trailing whitespace via `python 03-ai-scripts/05-guideline-autofixer.py <folder> --check-only`.

. **CRITICAL AI INSTRUCTION:** This changelog is an active AI execution record. Every modification to adjacent specifications must be logged here before task completion.

---
```

---

## 3. Detailed File-by-File Transformation Specifications

### 3.1 File 1: `02-spec/02-coding-guidelines/01-cross-language/98-changelog.md`

- **Current State:** Starts with passive title `# Coding Guidelines — Changelog` followed by metadata.
- **Action Required:**
  1. Update Title to `# Coding Guidelines — Changelog (AI Execution Prompt)`.
  2. Inject goal and learn block:
     ```markdown
     > **/goal** Maintain an immutable, sequential audit trail of all cross-language architectural revisions, guideline refactorings, and rule additions across 02-spec/02-coding-guidelines/01-cross-language/.
     > **/learn** Adhere to Keep-a-Changelog SemVer conventions, document structural migrations across subfolders, and verify all cross-language rule cross-references.
     ```
  3. Inject actionable checklist:
     ```markdown
     ## 🎯 Actionable CI/CD & Agent Checklist

     - [ ] `/goal` Record every cross-language rule addition, split, or modification under the appropriate SemVer release header.
     - [ ] `/learn` Ensure deduplicated rules (e.g. enum specifications in 06-ai-optimization/) maintain canonical cross-references.
     - [ ] `/goal` Enforce strict relative git paths for all referenced spec files (e.g., `02-spec/02-coding-guidelines/...`).
     - [ ] `/learn` Verify zero syntax or spacing violations via `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/01-cross-language --check-only`.

     . **CRITICAL AI INSTRUCTION:** This changelog is an active AI execution record. Any updates to cross-language guidelines must be logged here before concluding the task.
     ```
  4. Preserve all existing version entries (`v3.2.0`, `v3.0.0`, `v2.1.0`, `v2.0.0`).

### 3.2 File 2: `02-spec/02-coding-guidelines/01-cross-language/16-static-analysis/98-changelog.md`

- **Current State:** Starts with `# Changelog: Static Analysis & Linter Enforcement`.
- **Action Required:**
  1. Update Title to `# Changelog: Static Analysis & Linter Enforcement (AI Execution Prompt)`.
  2. Inject goal and learn block:
     ```markdown
     > **/goal** Track version revisions, rule matrices, and CI quality gate additions for static analysis and linter enforcement specifications.
     > **/learn** Understand SonarQube rule mappings (S1126, S4144) across 8 supported languages, and ensure CI pipeline templates stay synchronized with guideline standards.
     ```
  3. Inject actionable checklist:
     ```markdown
     ## 🎯 Actionable CI/CD & Agent Checklist

     - [ ] `/goal` Log updates to multi-language static analysis configurations and SonarQube rule matrices.
     - [ ] `/learn` Verify all referenced linter specs (golangci-lint, phpcs/phpstan, stylecop, clippy, ruff, eslint) maintain active relative links.
     - [ ] `/goal` Confirm new entries follow the standard bracketed SemVer format (`## [X.Y.Z] — YYYY-MM-DD`).
     - [ ] `/learn` Audit formatting with `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/01-cross-language/16-static-analysis --check-only`.

     . **CRITICAL AI INSTRUCTION:** This changelog tracks static analysis enforcement standards. Changes to linter rules or pipelines must be recorded here.
     ```
  4. Preserve existing entries (`[1.2.0]`, `[1.1.0]`, `[1.0.0]`).

### 3.3 File 3: `02-spec/02-coding-guidelines/02-typescript/98-changelog.md`

- **Current State:** Starts with `# TypeScript Standards — Changelog`.
- **Action Required:**
  1. Update Title to `# TypeScript Standards — Changelog (AI Execution Prompt)`.
  2. Inject goal and learn block:
     ```markdown
     > **/goal** Record all architectural improvements, strict typing mandates, async patterns, and code hygiene updates across TypeScript coding guidelines.
     > **/learn** Internalize TypeScript CODE RED rules (such as Promise.all() for independent async operations) and ensure all type-safety revisions are documented.
     ```
  3. Inject actionable checklist:
     ```markdown
     ## 🎯 Actionable CI/CD & Agent Checklist

     - [ ] `/goal` Document any additions or updates to TypeScript strict typing, enum conventions, and async guidelines.
     - [ ] `/learn` Maintain explicit references to CODE RED async patterns (`Promise.all()`) and interface encapsulation rules.
     - [ ] `/goal` Ensure all specification file paths use strict relative repository paths.
     - [ ] `/learn` Audit formatting and boolean conventions using `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/02-typescript --check-only`.

     . **CRITICAL AI INSTRUCTION:** Active AI execution record for TypeScript standards. Modifications to TypeScript guidelines must be recorded here.
     ```
  4. Preserve existing entries (`v2.1.0`, `v2.0.0`).

### 3.4 File 4: `02-spec/02-coding-guidelines/03-golang/98-changelog.md`

- **Current State:** Starts with `# Golang Standards — Changelog`.
- **Action Required:**
  1. Update Title to `# Golang Standards — Changelog (AI Execution Prompt)`.
  2. Inject goal and learn block:
     ```markdown
     > **/goal** Track version revisions, package standards, integer enum specifications, and error handling updates across Golang coding guidelines.
     > **/learn** Internalize Go-specific requirements: *appfault.AppError return types, zero bare void returns, integer-backed enums with PascalCase serialization, and parameter structs.
     ```
  3. Inject actionable checklist:
     ```markdown
     ## 🎯 Actionable CI/CD & Agent Checklist

     - [ ] `/goal` Log all Go standard updates, including integer enum rules, appfault error conventions, and struct sizing rules.
     - [ ] `/learn` Ensure retrospective references point to valid relative paths (e.g. `02-spec/03-error-manage/...`).
     - [ ] `/goal` Enforce proper SemVer numbering and ISO date formats on all new changelog sections.
     - [ ] `/learn` Audit formatting with `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/03-golang --check-only`.

     . **CRITICAL AI INSTRUCTION:** Active AI execution record for Go standards. Log all changes to Go specifications here.
     ```
  4. Preserve existing entries (`v2.2.0`, `v2.1.0`, `v2.0.0`).

### 3.5 File 5: `02-spec/02-coding-guidelines/04-php/98-changelog.md`

- **Current State:** Starts with `# PHP Standards — Changelog`.
- **Action Required:**
  1. Update Title to `# PHP Standards — Changelog (AI Execution Prompt)`.
  2. Inject goal and learn block:
     ```markdown
     > **/goal** Record all architectural revisions, strict typing standards, and static analysis integrations across PHP coding guidelines.
     > **/learn** Internalize modern PHP standards: strict types (`declare(strict_types=1)`), typed properties, enum standards, and spacing requirements.
     ```
  3. Inject actionable checklist:
     ```markdown
     ## 🎯 Actionable CI/CD & Agent Checklist

     - [ ] `/goal` Document updates to PHP coding guidelines, decomposition of reference specs, and code example fixes.
     - [ ] `/learn` Ensure all relative links to PHP subfolders (`07-php-standards-reference/`) resolve correctly.
     - [ ] `/goal` Verify new version entries conform to project-wide SemVer conventions.
     - [ ] `/learn` Audit formatting with `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/04-php --check-only`.

     . **CRITICAL AI INSTRUCTION:** Active AI execution record for PHP standards. Log all revisions to PHP guidelines here.
     ```
  4. Preserve existing entries (`v2.1.0`, `v2.0.0`).

### 3.6 File 6: `02-spec/02-coding-guidelines/05-rust/98-changelog.md`

- **Current State:** Starts with `# Changelog: Rust Standards`.
- **Action Required:**
  1. Update Title to `# Changelog: Rust Standards (AI Execution Prompt)`.
  2. Inject goal and learn block:
     ```markdown
     > **/goal** Maintain a precise audit trail of memory safety rules, async patterns, Clippy linter standards, and error handling updates across Rust guidelines.
     > **/learn** Master Rust-specific conventions: Result/Option idiomatic returns, ownership/borrowing rules, affirmative boolean flags, and Clippy pedantic compliance.
     ```
  3. Inject actionable checklist:
     ```markdown
     ## 🎯 Actionable CI/CD & Agent Checklist

     - [ ] `/goal` Record updates to Rust guidelines, AI confidence scoring, and Clippy integration rules.
     - [ ] `/learn` Ensure version numbers match root Rust standards and maintain consistent bracketed SemVer formatting (`## [X.Y.Z] — YYYY-MM-DD`).
     - [ ] `/goal` Verify all cross-references to adjacent Rust specs are valid relative paths.
     - [ ] `/learn` Audit formatting with `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/05-rust --check-only`.

     . **CRITICAL AI INSTRUCTION:** Active AI execution record for Rust standards. Document all updates to Rust guidelines here.
     ```
  4. Preserve existing entries (`[1.1.0]`, `[1.0.0]`).

### 3.7 File 7: `02-spec/02-coding-guidelines/07-csharp/98-changelog.md`

- **Current State:** Starts with `# C# Coding Standards — Changelog`.
- **Action Required:**
  1. Update Title to `# C# Coding Standards — Changelog (AI Execution Prompt)`.
  2. Inject goal and learn block:
     ```markdown
     > **/goal** Maintain an immutable record of C# language standards, nullable reference type rules, async patterns, and StyleCop linter enforcement.
     > **/learn** Master C# conventions: PascalCase methods, `I` prefix interfaces, boolean flag splitting, record types for DTOs, and pattern matching.
     ```
  3. Inject actionable checklist:
     ```markdown
     ## 🎯 Actionable CI/CD & Agent Checklist

     - [ ] `/goal` Document all additions to C# coding standards, cross-language integrations, and StyleCop analyzer rules.
     - [ ] `/learn` Verify cross-references to cross-language guidelines (`01-cross-language/24-`, `25-`) and AI checklists remain valid relative paths.
     - [ ] `/goal` Ensure bracketed SemVer formatting (`## [X.Y.Z] — YYYY-MM-DD`) is strictly maintained.
     - [ ] `/learn` Audit formatting with `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/07-csharp --check-only`.

     . **CRITICAL AI INSTRUCTION:** Active AI execution record for C# standards. Document all updates to C# guidelines here.
     ```
  4. Preserve existing entries (`[1.0.0]`).

---

## 4. Non-Negotiable Operational Rules

1. **Strict Relative Git Paths Mandate (TOTAL BAN on Absolute Paths / `file:///` URIs):**
   - Every file path, link, or reference MUST use strict relative paths from the git root or document directory.
   - NEVER write `/absolute/path/...`, `C:\...`, or `file:///...`.
2. **Worker Git Ban:**
   - Worker agents MUST NOT run `git add`, `git commit`, `git push`, or any other git commands. Commits and branch management are strictly handled by the lead orchestrator.
3. **Preservation of Historical Integrity:**
   - No historical changelog entries, version headers, dates, or bullet points may be deleted, truncated, or summarized. Only the header, metadata, and checklist are added.
4. **Clean UNIX LF & Trailing Whitespace Hygiene:**
   - All files must use clean UNIX LF line endings and end with a single trailing newline.
   - Zero trailing whitespace on any line.

---

## 5. Verification & Acceptance Criteria

### AC-CG-LOG-001: Changelog AI Execution Prompt Standard Conformance

**Given** The 7 changelog files across `02-spec/02-coding-guidelines/`.
**When** Inspected by static analyzers and guideline linters.
**Then** Each file begins with the `(AI Execution Prompt)` title suffix, contains valid `> **/goal**` and `> **/learn**` blocks, features the 4-item `## 🎯 Actionable CI/CD & Agent Checklist`, retains all original version history, and has zero absolute filesystem paths.

**Verification command:**
```bash
python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only
```
**Expected:** exit 0. Zero violations detected.

---

## 6. Related Specifications

- [`02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/01-architecture-spec.md`](../../../../02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/01-architecture-spec.md) — Parent Architecture Specification
- [`02-spec/02-coding-guidelines/readme.md`](../../../../02-spec/02-coding-guidelines/readme.md) — Coding Guidelines Index
- [`02-spec/02-coding-guidelines/01-cross-language/readme.md`](../../../../02-spec/02-coding-guidelines/01-cross-language/readme.md) — Cross-Language Guidelines
