# Plan: Coding Guidelines Actionable Checklists & Acceptance Criteria Standard (AI Execution Prompt)

> **/goal** Transform and formalize all coding guidelines across the repository into active, testable AI execution prompts featuring standardized agent checklists and testable Gherkin acceptance criteria.
> **/learn** Enforce the 4-part architectural anatomy, multi-agent wave partitioning (A = 2, H = 2), disjoint file boxes, worker git ban, and automated verification gates.

## 🎯 Actionable CI/CD & Agent Checklist

- [x] `/goal` Review the coding guidelines in `02-spec/02-coding-guidelines/readme.md` and audit all 223 markdown files.
- [x] `/learn` Create architectural specifications in `02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/`.
- [x] `/goal` Standardize the 7 historical changelog files with AI Execution Prompt headers and Actionable Checklists.
- [x] `/learn` Author standalone acceptance criteria registries for Python and C++ and synchronize `97-acceptance-criteria.md`.

. **CRITICAL AI INSTRUCTION:** This completed plan records the architecture, execution proof, and verification evidence for Task-01 and Task-02.

**Version:** 1.0.0  
**Updated:** 2026-10-03  
**Status:** Completed  
**AI Confidence:** Production-Ready  
**Ambiguity:** None  

---

## 1. Executive Summary

This plan executed the complete transformation of coding guidelines into authoritative, active AI execution prompts. Every file now strictly fulfills the 4-part architectural contract:
1. **AI Execution Prompt Header** (`> **/goal**` and `> **/learn**`)
2. **Actionable CI/CD & Agent Checklist** (`## 🎯 Actionable CI/CD & Agent Checklist` with `/goal` and `/learn`)
3. **Core Specification Body** (Numbered rules, ❌ FORBIDDEN vs ✅ REQUIRED code examples)
4. **Verification & Acceptance Criteria** (Canonical ID `AC-CG-*`, `Given / When / Then`, explicit verification command, and `Expected: exit 0`)

---

## 2. Deliverables & Files Modified

### 2.1 Architectural Specifications (`02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/`)
- `readme.md`: Module overview, directory index, and quality gate.
- `01-architecture-spec.md`: Full architectural specification detailing the 4-part anatomy, style guidelines, Boolean guidelines, naming conventions, and verbatim user prompt.
- `02-component-spec.md`: Component specification detailing the acceptance criteria registry taxonomy and automated verification engine.

### 2.2 Standardized Changelog Files (7/7 Upgraded)
1. `02-spec/02-coding-guidelines/01-cross-language/98-changelog.md`
2. `02-spec/02-coding-guidelines/01-cross-language/16-static-analysis/98-changelog.md`
3. `02-spec/02-coding-guidelines/02-typescript/98-changelog.md`
4. `02-spec/02-coding-guidelines/03-golang/98-changelog.md`
5. `02-spec/02-coding-guidelines/04-php/98-changelog.md`
6. `02-spec/02-coding-guidelines/05-rust/98-changelog.md`
7. `02-spec/02-coding-guidelines/07-csharp/98-changelog.md`

### 2.3 Polyglot Acceptance Criteria Registries & Master Sync
- `02-spec/02-coding-guidelines/12-python/97-acceptance-criteria.md`: Created registry mapping `AC-CG-PY-001` through `AC-CG-PY-REG-001`.
- `02-spec/02-coding-guidelines/13-cpp/97-acceptance-criteria.md`: Created registry mapping `AC-CG-CPP-001` through `AC-CG-CPP-REG-001`.
- `02-spec/02-coding-guidelines/97-acceptance-criteria.md`: Synchronized with `AC-12: Python Standards Registry` and `AC-13: Modern C++ Standards Registry`.

---

## 3. Verification Evidence

1. **Guideline Autofixer Validation:**
   ```bash
   python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only
   ```
   *Result:* Exit code 0. All 225 files have clean newlines and conform to implicit boolean rules.

2. **Relative Path Validation:**
   ```bash
   python linter-scripts/check-relative-paths.py
   ```
   *Result:* Exit code 0. Zero absolute filesystem paths or `file:///` URIs across 3502 tracked files.

3. **Multi-Agent Task Tracking:**
   ```bash
   python 03-ai-scripts/46-agent-sqlite-task-manager.py status --db .ai-memory/temp-agents/07-coding-guideline-actionable-checklist-and-acceptance/agent-task.db
   ```
   *Result:* Exit code 0. 100% completed across all subtasks without a single error or collision.
