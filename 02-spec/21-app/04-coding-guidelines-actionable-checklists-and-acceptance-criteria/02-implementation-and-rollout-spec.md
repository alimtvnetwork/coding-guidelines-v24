# Implementation & Rollout Specification: Coding Guidelines Transformation

> **/goal** Coordinate the multi-wave rollout transforming 32 cross-language coding guideline files into prompt-driven actionable checklists and acceptance criteria.
> **/learn** Enforce disjoint file boundaries, bounded 5–8 file micro-batches, parallel subagent dispatch (A = 2, H = 2), and zero-lock execution under execute-parent-task-with-n-steps-v6.

**Version:** 1.0.0  
**Updated:** 2026-10-02  
**Status:** Active  
**AI Confidence:** Production-Ready  
**Ambiguity:** None  

---

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Verify each worker wave touches strictly disjoint, non-overlapping file sets.
- [ ] `/learn` Ban all worker git executions to prevent `.git/index.lock` collisions in shared workspaces.
- [ ] `/goal` Execute targeted linter checks (`python 03-ai-scripts/05-guideline-autofixer.py`) after each wave.
- [ ] `/learn` Verify that all completed subtasks are recorded in the SQLite database and ledger.

---

## 1. Rollout Waves & Worker Allocations

To guarantee safe concurrent execution without file lock collisions, work is partitioned into bounded waves:

| Wave | Subtask ID | Worker | Assigned Files (Disjoint Bounding Box) | Focus Area |
|:---|:---|:---|:---|:---|
| **Wave 1** | Task-02A | Worker 01 | `01-cross-language/02-boolean-principles/readme.md`<br>`01-cross-language/02-boolean-principles/02-naming-prefixes.md`<br>`01-cross-language/02-boolean-principles/03-guards-and-extraction.md`<br>`01-cross-language/02-boolean-principles/04-parameters-and-conditions.md`<br>`01-cross-language/02-boolean-principles/05-quick-reference.md`<br>`01-cross-language/02-boolean-principles/06-exemptions-and-api.md` | Boolean Principles Subfolder |
| **Wave 1** | Task-02B | Worker 02 | `01-cross-language/12-no-negatives.md`<br>`01-cross-language/24-boolean-flag-methods.md` | Positive Guards & Flag Methods |
| **Wave 2** | Task-03A | Worker 01 | `01-cross-language/04-code-style/readme.md`<br>`01-cross-language/04-code-style/02-braces-and-nesting.md`<br>`01-cross-language/04-code-style/03-conditions-and-extraction.md`<br>`01-cross-language/04-code-style/04-blank-lines-and-spacing.md`<br>`01-cross-language/04-code-style/05-function-and-type-size.md` | Code Style (Part 1) |
| **Wave 2** | Task-03B | Worker 02 | `01-cross-language/04-code-style/06-multi-line-formatting.md`<br>`01-cross-language/04-code-style/07-comments-and-documentation.md`<br>`01-cross-language/04-code-style/08-checklist.md`<br>`01-cross-language/21-newline-styling-examples.md` | Code Style (Part 2) |
| **Wave 3** | Task-04A | Worker 01 | `01-cross-language/10-function-naming.md`<br>`01-cross-language/11-key-naming-pascalcase.md`<br>`01-cross-language/22-variable-naming-conventions.md` | Function & Variable Naming |
| **Wave 3** | Task-04B | Worker 02 | `01-cross-language/07-database-naming.md`<br>`01-cross-language/28-slug-conventions.md` | Database Naming & Slugs |
| **Wave 4** | Task-05A | Worker 01 | `01-cross-language/13-strict-typing.md`<br>`01-cross-language/18-code-mutation-avoidance.md`<br>`01-cross-language/32-branch-immutability-and-clean-construction.md` | Type Safety & Immutability |
| **Wave 4** | Task-05B | Worker 02 | `01-cross-language/33-variadic-and-spread-parameters.md`<br>`01-cross-language/34-string-normalization-and-equalfoldany.md` | Variadic & String Standard |
| **Wave 5** | Task-06A | Worker 01 | `01-cross-language/03-casting-elimination-patterns.md`<br>`01-cross-language/06-cyclomatic-complexity.md`<br>`01-cross-language/08-dry-principles.md` | Casting, Complexity & DRY |
| **Wave 5** | Task-06B | Worker 02 | `01-cross-language/16-lazy-evaluation-patterns.md`<br>`01-cross-language/19-null-pointer-safety.md`<br>`01-cross-language/20-nesting-resolution-patterns.md`<br>`01-cross-language/23-solid-principles.md` | Patterns, Safety & SOLID |
| **Wave 6** | Task-07 | Lead | `01-cross-language/97-acceptance-criteria.md`<br>`97-acceptance-criteria.md`<br>`01-cross-language/readme.md`<br>`02-coding-guidelines/readme.md` | Central Registries & Push Gate |
| **Wave 7** | Task-08A/B | Workers 01 & 02 | `03-golang/readme.md`, `02`, `03`, `05`, `06`, `07`, `08`, `09` | Go Core Coding Standards |
| **Wave 8** | Task-11A/B | Workers 01 & 02 | Root Guidelines (`02`, `03`, `04`, `05`) & `03-golang/01-enum-specification/*` | Root Style Guidelines & Go Enum Specs |
| **Wave 9** | Task-12A/B | Workers 01 & 02 | `03-golang/04-golang-standards-reference/*` & `02-typescript/readme.md`, `02..07`, `11` | Go Reference Specs & TS Status Enums |
| **Wave 10** | Task-13A/B | Workers 01 & 02 | `02-typescript/08..10`, `12..15` & `02-typescript/97-acceptance-criteria.md` | TypeScript Advanced Specs & TS Registry |
| **Wave 11** | Task-14A/B | Workers 01 & 02 | `04-php/*` & `04-php/07-php-standards-reference/*` | PHP Core Standards & Reference Specs |
| **Wave 12** | Task-15A/B | Workers 01 & 02 | `05-rust/*` & `07-csharp/*` | Rust & C# Coding Guidelines |
| **Wave 13** | Task-16A/B | Workers 01 & 02 | `06-ai-optimization/*` | AI Optimization Guidelines & Memory Lifecycles |
| **Wave 14** | Task-17A/B | Workers 01 & 02 | `06-cicd-integration/*` & `08-fix-repo-and-installers/*` | CI/CD Integration & Automated Fix Installers |
| **Wave 15** | Task-18A/B | Workers 01 & 02 | `08-file-folder-naming/*` & Polyglot Guidelines (`09`, `10`, `12`, `13`) | File Naming & Polyglot Guidelines |
| **Wave 16** | Task-19A/B | Workers 01 & 02 | `11-security/*` & `21-app/`..`24-app-ui/` Module Readmes | Security Standards & Module Readmes |
| **Wave 17** | Task-20A/B | Workers 01 & 02 | `11-security/97-acceptance-criteria.md`, CI/CD FAQ/Troubleshoot, Cross-Language Readme | Security AC Registry & Documentation |
| **Wave 18** | Task-21 | Lead | Master Acceptance Criteria Registry & Repository Gates | Master 97-acceptance-criteria.md & Final Verification |

---


## 2. Verification & Acceptance Criteria

### AC-CG-ROLL-001: Rollout Partitioning & Non-Overlapping Files

**Given** The multi-wave worker task allocation matrix.  
**When** Subagents are spawned to edit guideline files.  
**Then** Zero file path overlaps occur between concurrent workers, preventing git index collisions and file lock contention.

**Verification command:**
```bash
python 03-ai-scripts/46-agent-sqlite-task-manager.py status --db .ai-memory/temp-agents/04-coding-guidelines-actionable-checklists/agent-task.db
```
**Expected:** exit 0. All subtasks completed cleanly.

---

## 3. Related Specifications

- [`01-architecture-spec.md`](./01-architecture-spec.md) — Architectural standard and contract
- [`02-spec/02-coding-guidelines/01-cross-language/readme.md`](../../02-coding-guidelines/01-cross-language/readme.md) — Target guideline folder
