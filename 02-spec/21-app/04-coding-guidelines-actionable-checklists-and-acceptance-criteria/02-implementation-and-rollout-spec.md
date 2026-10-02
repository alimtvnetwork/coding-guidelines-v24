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
