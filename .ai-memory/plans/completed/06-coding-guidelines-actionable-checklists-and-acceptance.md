# Plan: Coding Guidelines Actionable Checklists & Acceptance Criteria Standard (AI Execution Prompt)

> **/goal** Transform and formalize all coding guidelines across the repository into active, testable AI execution prompts featuring standardized agent checklists and testable Gherkin acceptance criteria.
> **/learn** Enforce the 4-part architectural anatomy, multi-agent wave partitioning (A = 2, H = 2), disjoint file boxes, worker git ban, and automated verification gates.

## 🎯 Actionable CI/CD & Agent Checklist

- [x] `/goal` Review the coding guidelines in `02-spec/02-coding-guidelines/readme.md` and audit all specification files.
- [x] `/learn` Create architectural specifications in `02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/`.
- [x] `/goal` Standardize the 7 historical changelog files with AI Execution Prompt headers and Actionable Checklists.
- [x] `/learn` Author standalone acceptance criteria registries for Python and C++ and synchronize `97-acceptance-criteria.md`.
- [x] `/goal` Fix heading blank line spacing and renumber duplicate Principle 9 in `01-cross-language/02-boolean-principles/04-parameters-and-conditions.md`.
- [x] `/learn` Harmonize master acceptance criteria registry with runnable verification commands and resolve root style ID collisions.

. **CRITICAL AI INSTRUCTION:** This completed plan records the architecture, execution proof, and verification evidence for the coding guidelines actionable checklist and acceptance criteria standardization.

**Version:** 1.1.0  
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

### 2.2 Standardized Polyglot Registries & Master Registry Harmonization
- `02-spec/02-coding-guidelines/12-python/97-acceptance-criteria.md`: Created registry mapping `AC-CG-PY-001` through `AC-CG-PY-REG-001`.
- `02-spec/02-coding-guidelines/13-cpp/97-acceptance-criteria.md`: Created registry mapping `AC-CG-CPP-001` through `AC-CG-CPP-REG-001`.
- `02-spec/02-coding-guidelines/97-acceptance-criteria.md`: Added `Verification Command` column across all tables, added `09-powershell-integration` and `10-research` sections, and resolved ID collisions.

### 2.3 Style & Boolean Guideline Standardization
- `02-canonical-size-tier.md`: Retagged criterion to `AC-CG-ROOT-002`.
- `03-coding-style-checklist.md`: Retagged criterion to `AC-CG-ROOT-003`.
- `04-consolidated-review-guide-condensed.md`: Retagged criterion to `AC-CG-ROOT-004`.
- `05-consolidated-review-guide.md`: Retagged criterion to `AC-CG-ROOT-005`.
- `01-cross-language/02-boolean-principles/04-parameters-and-conditions.md`: Renumbered duplicate Principle 9 to Principle 12.
- `01-cross-language/34-string-normalization-and-equalfoldany.md`: Normalized markdown heading blank line spacing.
- `06-ai-optimization/97-acceptance-criteria.md`: Normalized markdown heading blank line spacing.
- `01-cross-language/97-acceptance-criteria.md`: Cleaned up dead links.

---

## 3. Verification Evidence

1. **Guideline Autofixer Validation:**
   ```bash
   python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only
   ```
   *Result:* Exit code 0. All 198 files have clean newlines and conform to implicit boolean rules.

2. **Relative Path Validation:**
   ```bash
   python linter-scripts/check-relative-paths.py
   ```
   *Result:* Exit code 0. Zero absolute filesystem paths or `file:///` URIs across 3448 tracked files.

3. **Multi-Agent Task Tracking:**
   ```bash
   python 03-ai-scripts/46-agent-sqlite-task-manager.py status --db .ai-memory/temp-agents/07-coding-guideline-actionable-checklist-and-acceptance/agent-task.db
   ```
   *Result:* Exit code 0. 100% completed across all subtasks (Task-01 through Task-04) without errors.
