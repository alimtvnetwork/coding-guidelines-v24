# Plan: Coding Guidelines Prompt, Actionable Checklists & Acceptance Audit (AI Execution Prompt)

> **/goal** Execute an end-to-end verification and hardening of coding guidelines in `02-spec/02-coding-guidelines/`, ensuring 100% adherence to active AI prompt headers, actionable agent checklists, and testable acceptance criteria across Style, Boolean, and Naming specifications.
> **/learn** Master the 4-part architectural anatomy, multi-agent wave partitioning (A = 2, H = 2), worker git ban, and automated zero-regression quality gates.

## 🎯 Actionable CI/CD & Agent Checklist

- [x] `/goal` Verify all 199 files in `02-spec/02-coding-guidelines/` have the active AI prompt header `(AI Execution Prompt)` and `> **/goal**` / `> **/learn**`.
- [x] `/learn` Verify every file includes `## 🎯 Actionable CI/CD & Agent Checklist` with interactive checkboxes and `. **CRITICAL AI INSTRUCTION:**`.
- [x] `/goal` Verify all style, boolean, and naming specs conclude with testable `## Verification & Acceptance Criteria` (`AC-CG-*`) in Given/When/Then format.
- [x] `/learn` Execute targeted linters (`03-ai-scripts/05-guideline-autofixer.py`, `linter-scripts/check-relative-paths.py`) ensuring exit code 0.

. **CRITICAL AI INSTRUCTION:** This completed plan records the architecture, execution proof, and verification evidence for the coding guidelines actionable checklist and acceptance criteria audit and hardening run.

**Version:** 1.1.0  
**Updated:** 2026-10-03  
**Status:** Completed  
**AI Confidence:** Production-Ready  
**Ambiguity:** None  

---

## 1. Executive Summary

This plan audited and hardened coding guidelines across `02-spec/02-coding-guidelines/` to guarantee that every document functions as an active AI execution prompt. Every guideline adheres to the standardized 4-part anatomy:
1. **AI Execution Prompt Header** (`> **/goal**` and `> **/learn**`)
2. **Actionable CI/CD & Agent Checklist** (`## 🎯 Actionable CI/CD & Agent Checklist` with `/goal` and `/learn`)
3. **Core Specification Body** (Numbered rules, ❌ FORBIDDEN vs ✅ REQUIRED code comparisons)
4. **Verification & Acceptance Criteria** (Canonical ID `AC-CG-*`, `Given / When / Then`, explicit verification command, and `Expected: exit 0`)

---

## 2. Deliverables & Files Modified

### 2.1 Boolean Principles Index Hardening
- `02-spec/02-coding-guidelines/01-cross-language/02-boolean-principles/readme.md`: Cleaned up Document Inventory table, fixed link target to `02-naming-prefixes.md`, and eliminated duplicated report entries.

### 2.2 Naming and Style Guidelines Hardening
- `02-spec/02-coding-guidelines/01-cross-language/22-variable-naming-conventions.md`: Added explicit Rule 6 banning generic garbage identifiers (`data`, `temp`, `obj`, `val`), complete with polyglot code examples in TypeScript, Go, and PHP, checklist updates, and harmonized acceptance criteria.
- `02-spec/02-coding-guidelines/03-coding-style-checklist.md`: Verified parameter limits (≤3), vertical line gaps, and `AC-CG-ROOT-003` criteria.

---

## 3. Verification Evidence

1. **Guideline Autofixer Validation:**
   ```bash
   python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only
   ```
   *Result:* Exit code 0. All 199 files conform to clean newlines and implicit boolean rules.

2. **Relative Path Linter:**
   ```bash
   python linter-scripts/check-relative-paths.py
   ```
   *Result:* Exit code 0. Zero absolute filesystem paths across repository files.

3. **Multi-Agent Task Tracking:**
   ```bash
   python 03-ai-scripts/46-agent-sqlite-task-manager.py status --db .ai-memory/temp-agents/11-11-guideline-prompt-checklist-acceptance-audit/agent-task.db
   ```
   *Result:* Exit code 0. 100% completed across all subtasks without errors.
