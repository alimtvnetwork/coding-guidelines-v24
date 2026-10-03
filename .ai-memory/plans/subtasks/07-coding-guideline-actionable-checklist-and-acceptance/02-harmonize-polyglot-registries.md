# Subtask 02: Harmonize Polyglot Registries & Synchronize Master Acceptance Criteria

> **Parent Spec:** [`02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/02-component-spec.md`](../../../../02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/02-component-spec.md)
> **Status:** `PENDING`
> **Traceability ID:** `Task-07-Subtask-02`
> **Target Files:**
> - `02-spec/02-coding-guidelines/12-python/97-acceptance-criteria.md` (to create)
> - `02-spec/02-coding-guidelines/13-cpp/97-acceptance-criteria.md` (to create)
> - `02-spec/02-coding-guidelines/97-acceptance-criteria.md` (to synchronize)

---

## 1. Subtask Objective & Context

The objective of Task-02 is to close the remaining polyglot acceptance criteria gap across the Prompt Architect specification suite. While TypeScript, Golang, PHP, Rust, and C# possess comprehensive domain registries in their respective directories, Python (`12-python`) and Modern C++ (`13-cpp`) currently lack standalone `97-acceptance-criteria.md` registries despite having core standard specifications (`02-standards.md`).

Furthermore, the master acceptance criteria registry at `02-spec/02-coding-guidelines/97-acceptance-criteria.md` terminates at section `AC-11: Application Architecture & Module Readmes`, omitting sections for Python and C++.

This subtask defines the exact operational sequence to author the missing domain registries, map all Python and C++ rules into structured `Given / When / Then` testable criteria, and integrate them into the master registry index.

---

## 2. Detailed Work Breakdown & Execution Plan

### Step 1: Author `02-spec/02-coding-guidelines/12-python/97-acceptance-criteria.md`

Author a comprehensive domain acceptance criteria registry for Python adhering strictly to Prompt Architect layout standards:
- **Header:** `# Python Standards — Acceptance Criteria Registry (AI Execution Prompt)`
- **Prompt Directives:** `> **/goal**` and `> **/learn**` blocks emphasizing static typing, Pydantic data validation, PEP-8 formatting, and specific exception handling.
- **Actionable Checklist:** 5-point checklist with `/goal` and `/learn` tags.
- **Criteria Inventory Table:**
  - `AC-CG-PY-001`: Python Guidelines Directory Index Conformance (Source: `readme.md`)
  - `AC-CG-PY-002`: Python Coding Standards Conformance (Source: `02-standards.md`)
  - `AC-CG-PY-003`: Python Dynamic Enum & Array Constants Standard (Source: `02-standards.md`)
  - `AC-CG-PY-004`: Python DRY Architecture & Engine Caching Conformance (Source: `02-standards.md`)
  - `AC-CG-PY-REG-001`: Python Acceptance Criteria Registry Conformance (Source: `97-acceptance-criteria.md`)
- **Detailed Specifications:** Exhaustive `Given / When / Then` definitions and executable verification commands for each criterion.

### Step 2: Author `02-spec/02-coding-guidelines/13-cpp/97-acceptance-criteria.md`

Author a comprehensive domain acceptance criteria registry for Modern C++ adhering strictly to Prompt Architect layout standards:
- **Header:** `# Modern C++ Standards — Acceptance Criteria Registry (AI Execution Prompt)`
- **Prompt Directives:** `> **/goal**` and `> **/learn**` blocks emphasizing C++20 baseline, concepts over `enable_if`, RAII memory safety (`std::unique_ptr`/`std::shared_ptr`), Rule of Zero/Five, and FFI boundary safety.
- **Actionable Checklist:** 5-point checklist with `/goal` and `/learn` tags.
- **Criteria Inventory Table:**
  - `AC-CG-CPP-001`: Modern C++ Guidelines Directory Index Conformance (Source: `readme.md`)
  - `AC-CG-CPP-002`: Modern C++ Standards Conformance (Source: `02-standards.md`)
  - `AC-CG-CPP-003`: C++ Memory Safety & RAII Resource Management (Source: `02-standards.md`)
  - `AC-CG-CPP-004`: C++ FFI Boundary Exception Safety & Standard Types (Source: `02-standards.md`)
  - `AC-CG-CPP-REG-001`: Modern C++ Acceptance Criteria Registry Conformance (Source: `97-acceptance-criteria.md`)
- **Detailed Specifications:** Exhaustive `Given / When / Then` definitions and executable verification commands for each criterion.

### Step 3: Synchronize Master Registry `02-spec/02-coding-guidelines/97-acceptance-criteria.md`

Update the master registry to formally index the new domain registries:
1. Append Section `## AC-12: Python Standards Registry`:
   - Summary table referencing `AC-CG-PY-001` through `AC-CG-PY-REG-001` with relative links to [`12-python/97-acceptance-criteria.md`](./12-python/97-acceptance-criteria.md).
2. Append Section `## AC-13: Modern C++ Standards Registry`:
   - Summary table referencing `AC-CG-CPP-001` through `AC-CG-CPP-REG-001` with relative links to [`13-cpp/97-acceptance-criteria.md`](./13-cpp/97-acceptance-criteria.md).
3. Expand Section `## Cross-References`:
   - Add relative links to `./12-python/readme.md` and `./12-python/97-acceptance-criteria.md`.
   - Add relative links to `./13-cpp/readme.md` and `./13-cpp/97-acceptance-criteria.md`.

---

## 3. Structural Templates & Concrete Schemas

### 3.1 Python Registry Schema (`12-python/97-acceptance-criteria.md`)

```markdown
# Python Standards — Acceptance Criteria Registry (AI Execution Prompt)

> **/goal** Provide a consolidated, traceable registry of testable acceptance criteria across all Python coding guideline specifications in `12-python/`.
> **/learn** Enforce the canonical criteria taxonomy (`AC-CG-PY-[NUM]`), PEP-8 compliance, Black formatting (max 100 chars), explicit static typing, Pydantic data schemas, specific exception handling, and verify compliance using targeted linters.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Verify each criterion maps 1:1 to an authoritative specification file in `12-python/`.
- [ ] `/learn` Ensure all criteria follow the structured `Given / When / Then` verification contract.
- [ ] `/goal` Validate that all verification commands execute cleanly with `Expected: exit 0`.
- [ ] `/learn` Verify 100% relative paths and zero absolute filesystem paths.
- [ ] `/goal` Verify compliance using `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/12-python --check-only`.

. **CRITICAL AI INSTRUCTION:** This specification is an active AI execution directive. All code generated or modified must strictly follow the rules below.

**Version:** 4.0.0
**Last Updated:** 2026-10-03
**Status:** Active
**AI Confidence:** Production-Ready
**Ambiguity:** None

---

## 1. Python Criteria Inventory (`AC-CG-PY-`)

| ID | Title | Authoritative Specification | Verification Command |
|:---|:---|:---|:---|
| `AC-CG-PY-001` | Python Guidelines Directory Index Conformance | [`readme.md`](readme.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/12-python --check-only` |
| `AC-CG-PY-002` | Python Coding Standards Conformance | [`02-standards.md`](02-standards.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/12-python --check-only` |
| `AC-CG-PY-003` | Python Dynamic Enum & Array Constants Standard | [`02-standards.md`](02-standards.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/12-python --check-only` |
| `AC-CG-PY-004` | Python DRY Architecture & Engine Caching Conformance | [`02-standards.md`](02-standards.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/12-python --check-only` |
| `AC-CG-PY-REG-001` | Python Acceptance Criteria Registry Conformance | [`97-acceptance-criteria.md`](97-acceptance-criteria.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/12-python --check-only` |
```

### 3.2 Modern C++ Registry Schema (`13-cpp/97-acceptance-criteria.md`)

```markdown
# Modern C++ Standards — Acceptance Criteria Registry (AI Execution Prompt)

> **/goal** Provide a consolidated, traceable registry of testable acceptance criteria across all Modern C++ coding guideline specifications in `13-cpp/`.
> **/learn** Enforce the canonical criteria taxonomy (`AC-CG-CPP-[NUM]`), modern C++20 standard baseline, concepts over enable_if, smart pointer memory safety (std::unique_ptr/std::shared_ptr), Rule of Zero/Five, PascalCase types, snake_case functions, and FFI exception containment.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Verify each criterion maps 1:1 to an authoritative specification file in `13-cpp/`.
- [ ] `/learn` Ensure all criteria follow the structured `Given / When / Then` verification contract.
- [ ] `/goal` Validate that all verification commands execute cleanly with `Expected: exit 0`.
- [ ] `/learn` Verify 100% relative paths and zero absolute filesystem paths.
- [ ] `/goal` Verify compliance using `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/13-cpp --check-only`.

. **CRITICAL AI INSTRUCTION:** This specification is an active AI execution directive. All code generated or modified must strictly follow the rules below.

**Version:** 4.0.0
**Last Updated:** 2026-10-03
**Status:** Active
**AI Confidence:** Production-Ready
**Ambiguity:** None

---

## 1. Modern C++ Criteria Inventory (`AC-CG-CPP-`)

| ID | Title | Authoritative Specification | Verification Command |
|:---|:---|:---|:---|
| `AC-CG-CPP-001` | Modern C++ Guidelines Directory Index Conformance | [`readme.md`](readme.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/13-cpp --check-only` |
| `AC-CG-CPP-002` | Modern C++ Standards Conformance | [`02-standards.md`](02-standards.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/13-cpp --check-only` |
| `AC-CG-CPP-003` | C++ Memory Safety & RAII Resource Management | [`02-standards.md`](02-standards.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/13-cpp --check-only` |
| `AC-CG-CPP-004` | C++ FFI Boundary Exception Safety & Standard Types | [`02-standards.md`](02-standards.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/13-cpp --check-only` |
| `AC-CG-CPP-REG-001` | Modern C++ Acceptance Criteria Registry Conformance | [`97-acceptance-criteria.md`](97-acceptance-criteria.md) | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/13-cpp --check-only` |
```

---

## 4. Verification & Quality Gates

Upon executing this subtask, the worker agent must verify compliance through the following checks:

### Step 4.1: Guideline Autofixer Verification
Execute the guideline autofixer in audit mode across the coding guidelines:
```bash
python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only
```
**Expected:** exit 0. Zero newline formatting or boolean convention violations detected.

### Step 4.2: Relative Path & Link Integrity Check
Verify that all generated files contain 100% relative repository paths:
- No absolute filesystem paths (`/absolute/...`, `C:\...`).
- No `file:///` URIs.
- All markdown cross-links resolve properly relative to their location.

---

## 5. Non-Negotiable Operational Constraints

1. **Worker Git Ban:** The worker agent executing this subtask MUST NOT run any `git` commands (`git add`, `git commit`, `git push`, `git checkout`, etc.). All git operations are strictly managed by the parent lead orchestrator.
2. **Disjoint Bounding Box:** The worker must touch ONLY the three designated target files:
   - `02-spec/02-coding-guidelines/12-python/97-acceptance-criteria.md`
   - `02-spec/02-coding-guidelines/13-cpp/97-acceptance-criteria.md`
   - `02-spec/02-coding-guidelines/97-acceptance-criteria.md`
3. **Strict Lowercase File Naming:** Filenames must remain strictly lowercase with hyphen separators.
