# Coding Guidelines Actionable Checklists & Acceptance Criteria Standard (AI Execution Prompt)

> **/goal** Master and enforce the architectural standards, specifications, and CI/CD validation rules for Coding Guidelines Actionable Checklists and Acceptance Criteria.
> **/learn** Read the sequentially ordered specification files in this directory, follow the actionable CI/CD checklist, and apply mandatory rules before generating code.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Verify every coding guideline specification contains an AI Execution Prompt header and Actionable Checklist.
- [ ] `/learn` Adhere strictly to the 4-part anatomical contract across all files in `02-spec/02-coding-guidelines/`.
- [ ] `/goal` Verify every coding guideline specification ends with a testable Gherkin acceptance criteria block.
- [ ] `/learn` Run all local verification linters via `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only`.

. **CRITICAL AI INSTRUCTION:** This `readme.md` file is the primary entry point for this specification module. AI agents MUST read this file first before exploring other files in this folder.

**Version:** 1.0.0  
**Updated:** 2026-10-03  
**Status:** Active  
**AI Confidence:** Production-Ready  
**Ambiguity:** None  

---

## Overview

This specification directory defines the complete architectural contract, component structure, and rollout plan for transforming all coding guidelines across the meta-repository into active, testable **AI execution prompts**.

Every guideline file is structured as:
1. **AI Execution Prompt Header** (`> **/goal**` and `> **/learn**`)
2. **Actionable CI/CD & Agent Checklist** (`## 🎯 Actionable CI/CD & Agent Checklist`)
3. **Core Specification Body** (Numbered rules, ❌ FORBIDDEN vs ✅ REQUIRED code examples)
4. **Verification & Acceptance Criteria** (Canonical ID `AC-CG-*`, `Given / When / Then`, explicit verification command, and `Expected: exit 0`)

---

## Directory Index

| File | Title | Description |
|:---|:---|:---|
| [`01-architecture-spec.md`](./01-architecture-spec.md) | Architecture Specification | 4-part anatomy, style, boolean, and naming rules |
| [`02-component-spec.md`](./02-component-spec.md) | Component Specification | Acceptance criteria registries and verification engine |

---

## Verification & Acceptance Criteria

_Auto-generated section — see `02-spec/02-coding-guidelines/97-acceptance-criteria.md` for the full criteria index._

### AC-APP-CG-001: Architecture Specification Module Conformance

**Given** The specification files in `02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/`.  
**When** Audited against Prompt Architect specification rules.  
**Then** All files feature valid AI execution headers, actionable checklists, and testable acceptance criteria with 0 absolute paths.

**Verification command:**
```bash
python 03-ai-scripts/05-guideline-autofixer.py 02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance --check-only
```
**Expected:** exit 0. Zero violations.
