# Architecture Specification: Coding Guidelines Actionable Checklists & Testable Acceptance Criteria Standard

> **/goal** Transform all coding guidelines across the repository from static reference documentation into authoritative, actionable AI execution prompts equipped with standardized agent checklists at the top and formal, testable acceptance criteria at the bottom.
> **/learn** Master the structural standards of Prompt Architect specifications, enforce the mandatory top-level prompt block (`> **/goal**` and `> **/learn**`), the actionable CI/CD and agent checklist (`## 🎯 Actionable CI/CD & Agent Checklist`), and testable verification criteria (`Given / When / Then` with automated verification commands).

**Version:** 1.0.0  
**Updated:** 2026-10-02  
**Status:** Active  
**AI Confidence:** Production-Ready  
**Ambiguity:** None  

---

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Verify every coding guideline document begins with the standard AI action prompt header (`/goal` and `/learn`).
- [ ] `/learn` Verify every coding guideline document contains an immediately actionable `## 🎯 Actionable CI/CD & Agent Checklist` covering core rules.
- [ ] `/goal` Verify every coding guideline document ends with a testable `## Verification & Acceptance Criteria` section using `Given / When / Then` and verification commands.
- [ ] `/learn` Validate 100% relative paths and zero absolute filesystem paths across all updated specifications via `python linter-scripts/check-relative-paths.py`.

. **CRITICAL AI INSTRUCTION:** This specification defines the mandatory structural contract for all coding guidelines in `02-spec/02-coding-guidelines/`. AI agents authoring or modifying guideline files MUST enforce this exact layout.

---

## 1. User Request (Verbatim)

```text
# High Priority Instruction

Can you please go into the coding guideline and see the coding guideline, how it is written? It's not written as an action or prompt on every file. Try to have from the README.md file inside the coding guideline folder, like a spec and then coding guideline. Try to have every one of the file try to mention as actionable items that the AI agent must follow. Okay. As a checklist. And what is acceptance criteria at the end that needs to be mentioned on each one of the specs for coding guideline inside, like the style guideline, Boolean guideline, naming convention, things like that. Okay. Do you understand? Can you please follow through? Do you have any question and confusion?

# Actionable Items Must Follow Non-Negotiable

1. Review the coding guidelines in the README.md file within the coding guideline folder.
2. Create a specification and coding guideline document.
3. Ensure each file includes actionable items for AI agents as a checklist.
4. Define acceptance criteria for each spec, including style guidelines, Boolean guidelines, and naming conventions.

Must follow and spawn agent using 

@[.agents/skills/execute-parent-task-with-n-steps-v6]

## Additional Instructions

learn /learn if you have to learn something and /plan stuff before working please./plan
```

---

## 2. Motivation & Architectural Deficiency

Historically, files in `02-spec/02-coding-guidelines/` were authored as passive reference texts. While rich in explanation and code examples, this structure presented major deficiencies when consumed by autonomous AI agents:

1. **Lack of Immediate Actionability:** An AI agent scanning a file had to read hundreds of lines of prose to deduce the mandatory operational constraints.
2. **Missing Quality Gates:** Individual guideline files lacked unambiguous, testable acceptance criteria at the end, making automated verification subjective and difficult to evaluate in CI/CD.
3. **Inconsistency with Module Readmes:** While `02-spec/02-coding-guidelines/readme.md` provided a high-level `## 🎯 Actionable CI/CD & Agent Checklist` and `## Verification` block, the individual guideline chapters (`02-boolean-principles/*`, `04-code-style/*`, `22-variable-naming-conventions.md`, etc.) did not inherit this disciplined standard.

---

## 3. The Standard Specification Contract

Every specification file in `02-spec/02-coding-guidelines/` MUST strictly adhere to the following 4-part architectural anatomy:

```
┌────────────────────────────────────────────────────────┐
│ 1. AI Execution Prompt Header                          │
│    - Title with (AI Execution Prompt) suffix           │
│    - > **/goal** [Concrete goal]                       │
│    - > **/learn** [Key mental models & anti-patterns]  │
├────────────────────────────────────────────────────────┤
│ 2. Actionable CI/CD & Agent Checklist                  │
│    - ## 🎯 Actionable CI/CD & Agent Checklist          │
│    - Checkboxes with /goal and /learn directives       │
│    - . CRITICAL AI INSTRUCTION directive block         │
├────────────────────────────────────────────────────────┤
│ 3. Core Specification Body                             │
│    - Metadata block (Version, Updated, Status, etc.)   │
│    - Numbered rules, architectural explanations        │
│    - ❌ FORBIDDEN vs ✅ REQUIRED code examples        │
├────────────────────────────────────────────────────────┤
│ 4. Verification & Acceptance Criteria                  │
│    - ## Verification & Acceptance Criteria             │
│    - AC-CG-[CAT]-[NNN] with Given / When / Then        │
│    - Verification command and expected exit code 0     │
└────────────────────────────────────────────────────────┘
```

### 3.1 Standard Header & Checklist Schema

```markdown
# [Guideline Title] (AI Execution Prompt)

> **/goal** [Concise actionable goal statement defining the exact standard to enforce]
> **/learn** [Key references, anti-patterns, and mental models required before applying changes]

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` [Mandatory rule check 1]
- [ ] `/learn` [Anti-pattern check / mental model 2]
- [ ] `/goal` [Mandatory rule check 3]
- [ ] `/learn` [Verification / compliance check 4]

. **CRITICAL AI INSTRUCTION:** This specification is an active AI execution directive. All code generated or modified must strictly follow the rules below.
```

### 3.2 Standard Verification & Acceptance Criteria Schema

```markdown
## Verification & Acceptance Criteria

_Auto-generated section — see `02-spec/02-coding-guidelines/97-acceptance-criteria.md` for the full criteria index._

### AC-CG-[CATEGORY]-[NUM]: [Descriptive Title]

**Given** [Precondition or codebase context]
**When** [Target file or component is analyzed by linters/CI]
**Then** [Expected deterministic outcome: zero violations, exit code 0]

**Verification command:**
```bash
python 03-ai-scripts/05-guideline-autofixer.py <folder> --check-only
```
**Expected:** exit 0. Zero violations detected.
```

---

## 4. Acceptance Criteria Taxonomy & Categorization

To maintain full traceability across the repository, acceptance criteria are tagged using category-specific identifiers:

| Category Prefix | Domain / Scope | Target Directory |
|:---|:---|:---|
| `AC-CG-BOOL-` | Boolean principles, affirmative naming, implicit checks, no mixed polarity | `01-cross-language/02-boolean-principles/`, `12-no-negatives.md`, `24-boolean-flag-methods.md` |
| `AC-CG-STYLE-` | Braces, nesting, blank line vertical spacing, function and file sizing | `01-cross-language/04-code-style/`, `21-newline-styling-examples.md` |
| `AC-CG-NAME-` | Function naming, variable naming, collections, map names, PascalCase keys | `01-cross-language/10-`, `11-`, `22-`, `07-`, `28-` |
| `AC-CG-TYPE-` | Strict typing, casting elimination, immutability, parameter structs | `01-cross-language/03-`, `13-`, `18-`, `32-`, `33-`, `34-` |
| `AC-CG-ARCH-` | Cyclomatic complexity, DRY principles, SOLID, lazy evaluation, null safety | `01-cross-language/06-`, `08-`, `16-`, `19-`, `20-`, `23-` |

---

## 5. Verification & Acceptance Criteria

### AC-CG-ARCH-001: Canonical Architecture Spec Conformance

**Given** The Prompt Architect specification tree under `02-spec/21-app/`.  
**When** Linters inspect `02-spec/21-app/04-coding-guidelines-actionable-checklists-and-acceptance-criteria/01-architecture-spec.md`.  
**Then** The specification captures the verbatim user prompt, defines the standard AI header and checklist, specifies testable acceptance criteria format, and maintains 100% relative paths.

**Verification command:**
```bash
python linter-scripts/check-relative-paths.py
```
**Expected:** exit 0. Zero absolute paths or `file:///` URIs.

---

## 6. Related Specifications

- [`02-spec/02-coding-guidelines/readme.md`](../../02-coding-guidelines/readme.md) — Root coding guidelines index
- [`02-spec/02-coding-guidelines/01-cross-language/readme.md`](../../02-coding-guidelines/01-cross-language/readme.md) — Cross-language coding guidelines
- [`02-spec/01-spec-authoring-guide/readme.md`](../../01-spec-authoring-guide/readme.md) — Spec authoring guide
