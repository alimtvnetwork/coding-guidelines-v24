---
name: follow-cg-execute
description: Unified orchestrator for executing repository-wide coding guideline compliance suites (cg-*), batching refactors across 5-8 files with subagents, GitMap high-speed search, and zero-build/no-test execution.
---

# Follow Coding Guidelines Execution Suite (`follow-cg-execute`)

## Overview

`follow-cg-execute` is the master orchestration skill designed to execute and enforce all 39 specialized Coding Guideline (`cg-*`) compliance suites across the repository.

Instead of running dozens of isolated, fragmented guideline scripts, `follow-cg-execute` dynamically sequences audits, code refactoring passes, and verification gates in bounded 5–8 file micro-batches.

Always run `[/plan](slashCommand;plan)` first before starting any multi-file refactoring pass, and invoke `[/learn](slashCommand;learn)` from `@[.agents/skills/follow-coding-guidelines]` to maintain strict guideline adherence.

---

## The 39 Core Coding Guideline Domains Orchestrated

`follow-cg-execute` coordinates passes across all 10 canonical guideline families:

1. **Boolean Conventions & Nesting:**
   - `cg-boolean-and-naming`: Positive boolean prefixes (`is`, `has`), implicit evaluation, total ban on `== true`.
   - `nested-if-and-guard-clauses`: Early returns, guard clauses, flattening nested `if` statements.
   - `cg-branch-immutability-and-clean-construction`: Copy-on-write branch immutability, pure constructor returns.

2. **Error Handling & Return Types:**
   - `cg-error-management`: Structured error return type `*appfault.AppError`, leaf error sovereignty.
   - `result-wrapper-and-apperror-returns`: Monadic `Result[T]` containers eliminating multi-value return tuples.
   - `function-signatures-and-return-types`: Signature simplification, parameter structs for >2-3 arguments.

3. **Sizing & Structural Deconstruction:**
   - `cg-file-size-and-function-reduction`: 8–15 line functions max, 150–300 line files max (React UI 100 max).
   - `function-argument-reduction-and-params`: Parameter struct refactoring for oversized parameter lists.
   - `cg-extract-types`: Centralizing domain types, structs, and interfaces into dedicated `types.go` files.

4. **Code Dryness & Performance:**
   - `cg-code-dryness-and-library-extraction`: Eliminating repetitive duplication into reusable packages.
   - `cg-map-dispatch-and-handler-lookup`: Replacing long `switch` and `if-else` chains with map dispatch tables.
   - `cg-string-efficiency`: Zero-allocation string comparisons, hoisting string transformations.
   - `cg-string-normalization-and-equalfoldany`: Standardizing `strutil.EqualFoldAny` comparisons.
   - `lazy-regex-and-pattern-matching`: Pre-compiled lazy regular expressions in dedicated packages.

5. **Type Systems & Pointer Semantics:**
   - `cg-golang-pointer-reduction-and-value-semantics`: Value semantics for small structs, reducing `*T` indirection.
   - `cg-golang-pointers-and-mutation-reduction`: Eliminating in-place mutations, pure helper returns.
   - `cg-variadic-and-spread-parameters`: Enforcing variadic parameters over artificial slice wrapping.
   - `cg-enum-standards`: PascalCase enums with `Type` suffixes across all languages.
   - `typescript-guidelines-and-types`: Discriminated unions, no `any`, strongly-typed error results.

6. **Database & Architecture:**
   - `cg-database-conventions`: Split-DB SQLite three-tier architecture, PascalCase tables.
   - `sql-repository-pattern-and-enum-actions`: Repository pattern encapsulation, enum-driven action routing.

7. **Code Hygiene & Visual Rhythm:**
   - `style-guidelines`: Mandatory vertical spacing (blank before `if`, after `}`, before `return`).
   - `code-hygiene`: LF line endings, trailing whitespace elimination, redundant `.gitkeep` removal.
   - `relative-paths`: Strict relative Git paths, total ban on absolute paths or `file:///` URIs.

8. **Testing & OS Isolation:**
   - `cg-isolate-os-tests`: Strict mocking of system calls, ensuring tests never trigger OS shutdown/reboots.
   - `temp-end-to-end-tests`: Decoupled local integration tests skipped by default in CI/CD.

9. **Terminal UI & Presentation:**
   - `terminal-ui-and-cli-styling`: Lipgloss color tokens, unified CLI help menus, and status formatters.
   - `cli-commands-and-help`: Complete CLI subcommand registration and help coverage.

10. **Build & Cache Governance:**
    - `cg-clean-repo-build-and-caches`: Purging ephemeral build caches, lowercase file hygiene.
    - `zero-storage-actions-purge`: Zero artifact uploads to maintain 0.0 GB GitHub Actions storage.

---

## 4-Phase Micro-Batch Execution Strategy

When tasked with codebase-wide guideline compliance:

### Phase 1: Rapid Codebase Inventory & Gap Analysis (No Code Changes)
- Scan target directories using high-speed GitMap streaming:
  `gitmap aum search "<pattern>" <dir> -e <ext>`
  `gitmap lf <dir>`
- Catalog exact file paths, line ranges, and violation types into a gap matrix in `02-spec/21-app/`.
- Decompose the remediation roadmap into 5–8 file micro-tasks.

### Phase 2: Bounded Micro-Batch Execution (Fast-Forward Mode)
- Execute refactorings strictly in bounded chunks of 5–8 files.
- Apply surgical edits adhering strictly to whitespace gaps, parameter structs, and Result returns.
- Maintain Zero Harm: never alter core business logic or introduce behavioral regressions.
- Strict No-Build / No-Test rule during micro-batching iterations.

### Phase 3: Cognitive Heatmap Telemetry Verification
- Record file leases and updates using GitMap heatmap:
  `gitmap heatmap log --task <slug> --action EDIT --file <path> ...`
  `gitmap heatmap update --id <id> --outcome "..." --met true`
- Verify 100% action closure via `gitmap heatmap status --task <slug>`.

### Phase 4: Atomic Commit & Upstream Delivery
- Intermediate commits: `gitmap commit "wip - <scope>"`
- Final delivery commit: `gitmap cpf "<Module> - <summary>"` (e.g. `gitmap cpf "Guidelines - refactor booleans and parameter structs"`)
- Ensure all changes are pushed upstream ('No Push = Not Done').
