# Subtask 02: Component Specification, Schemas & Concrete Examples

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/01-create-prompts-for-100-improvements/02-subtask-component-and-examples.md`  
> **Parent Plan:** `.ai-memory/plans/pending/01-create-prompts-for-100-improvements.md`  
> **Specification:** `02-spec/21-app/01-create-prompts-for-100-improvements/02-component-spec.md`  
> **Status:** `[DONE]`  
> **Assigned Role:** Spec Author 02  

---

## 1. Objective & Scope

Author the exhaustive, grounded component specifications, mathematical models, data schemas, and concrete operational examples for the **100 Improvements Prompt Suite**. 

This subtask delivers `02-spec/21-app/01-create-prompts-for-100-improvements/02-component-spec.md` covering:
1. **UI Level Schema & Tokens:** 10 functional clusters of 10 items (100 total), 4-format color token matrix (`--blue-500: #2563EB`, `--blue-900: #0D2975`, `--violet-500: #822EE8`, etc.), WCAG 2.2 contrast standards, relative luminance formulas (WCAG 2.2 and ITU-R BT.709), button sizing tiers (`sm: 36px`, `md: 44px`, `lg: 52px`), 72px glassmorphic sticky header mechanics (`window.scrollY > 12`), hairline specular lighting (`border-border/40`), safe-region navigation (`14px` padded corridor), and 3D flip card physics (`perspective: 1400px`, `rotateY(180deg)`).
2. **Code Base Schema & Cards:** 10-field specification card schema (Target File Path, File Name, Function/Symbol, Line Numbers, Rule Violated & Citation, Severity, Current Code Context, Violation Analysis, Exact Fix Proposal, Blast Radius & Verification), 5–8 file micro-batching model, and 10 clusters of 10 items (100 total).
3. **Whole System Schema & Blueprint:** 8-part specification schema (Identifier, System UX Impact, Cause, Effect, Zero Harm / Regression Safety, Backward Compatibility, Implementation Blueprint, Binary Verification), and 10 clusters of 10 items (100 total).
4. **Companion Skills Architecture:** Symmetrical Antigravity (`.agents/skills/`) and Cursor (`.cursor/skills/`) skill definitions, YAML frontmatter, IDE link syntax rules (`@[.agents/skills/...]` vs `@[.cursor/skills/...]`), and plan enqueueing invariants.

---

## 2. Context & Guardrail Invariants

- **Strict Relative Git Paths:** All markdown links, file paths, citations, and subtask paths must strictly use relative paths (`02-spec/...`, `.ai-memory/...`, `01-prompts/...`). TOTAL BAN on absolute filesystem paths (`C:\...`, `/home/...`) and `file:///` URIs.
- **Strict Lowercase Naming:** All created and modified files, directories, and paths must adhere to lowercase naming conventions.
- **Zero Builds & Tests (Rule R1):** No builds, compilation, or test suites may be executed.
- **No Git Commands:** Subagents must never run `git add`, `git commit`, `git push`, or `git status` to prevent working tree collisions.
- **Zero Placeholders:** Specifications and subtasks must contain fully resolved file paths, exact code samples, and deterministic criteria. Placeholders like `[N]`, `TODO`, or `...` are strictly prohibited.

---

## 3. Action Checklist

- [x] Analyze master plan requirements in `.ai-memory/plans/pending/01-create-prompts-for-100-improvements.md`.
- [x] Inspect existing design system tokens and formulas in `02-spec/07-design-system/04-white-blue-theme/` and `02-spec/07-design-system/32-slide-color-options.md`.
- [x] Document 4-format color token table (OKLCH, HEX, RGB, HSL) with WCAG 2.2 contrast ratings.
- [x] Formulate mathematical models for WCAG 2.2 gamma linearization contrast and ITU-R BT.709 pill ink luminance ($L > 0.6 \implies$ `#090E18`, else `#FFFFFF`).
- [x] Detail button physics and sizing increments (`sm: 36px`, `md: 44px`, `lg: 52px`, `icon: 44px`) with child isolation and vertical press feedback.
- [x] Detail 72px sticky header physics with `window.scrollY > 12` scroll elevation threshold and glassmorphic backdrop blur.
- [x] Enumerate all 10 clusters of 10 items (100 items total) for the UI-level prompt.
- [x] Construct the 10-field code base specification card schema with 5–8 file micro-batching rules.
- [x] Author concrete, production-grade 10-field code cards for `CRITICAL` (session validator) and `HIGH` (parameter bloat) severities.
- [x] Enumerate all 10 clusters of 10 items (100 items total) for the code base prompt.
- [x] Construct the 8-part whole-system improvement schema with zero harm and backward compatibility guarantees.
- [x] Author concrete 8-part system improvement examples (`SYS-IMP-012` SQLite WAL and `SYS-IMP-045` Zero-Storage CI Diagnostics).
- [x] Enumerate all 10 clusters of 10 items (100 items total) for the whole system prompt.
- [x] Define symmetrical Antigravity and Cursor skills schemas with platform-specific link syntax.
- [x] Author and persist `02-spec/21-app/01-create-prompts-for-100-improvements/02-component-spec.md`.
- [x] Author and persist `.ai-memory/plans/subtasks/01-create-prompts-for-100-improvements/02-subtask-component-and-examples.md`.
- [x] Verify relative path hygiene (zero `file:///` URIs or absolute paths).

---

## 4. Execution Trace & Deliverables Summary

### 4.1 Primary Deliverable: `02-spec/21-app/01-create-prompts-for-100-improvements/02-component-spec.md`
- **Section 1: Executive Summary & Design System Grounding:** Anchors the 100 improvements engine in repository design specifications and coding rules.
- **Section 2: UI-Level 100 Improvements:**
  - 4-Format Color Palette Matrix: Complete specification of `--blue-50` through `--blue-900`, `--violet-500`, `--neutral-900`, `--neutral-200`, and `--paper`.
  - Mathematical Formulations: WCAG 2.2 sRGB linearization + relative luminance contrast ratio, and ITU-R BT.709 relative luminance formula.
  - Button Geometry: 4 height tiers (`sm: 36px`, `md: 44px`, `lg: 52px`, `icon: 44px`), `active:translate-y-[1px]`, child isolation `[&>*]:relative [&>*]:z-10`, and bottom-up pointer fill.
  - Sticky Header Physics: `72px` fixed desktop height, `window.scrollY > 12` trigger, `420ms cubic-bezier(0.16, 1, 0.3, 1)` transition, and `14px` safe-region corridor.
  - Complete 100 UI Item Enumeration: Organized into 10 clusters of 10 items (UI-001 through UI-100).
- **Section 3: Code Base 100 Improvements:**
  - 10-Field Card Schema: Standard structure covering Target File Path, File Name, Function/Symbol, Line Numbers, Rule Violated & Citation, Severity, Current Code Context, Violation Analysis, Exact Fix Proposal, Blast Radius & Verification.
  - 5–8 File Micro-Batching: Sequential isolation model with local flat commits (`gitmap commit`) and remote semantic commits (`gitmap cpf`).
  - Production-Grade Examples: Full cards for `CB-001` (`*appfault.AppError` and boolean check in `validator.go`) and `CB-022` (parameter struct in `runner.go`).
  - Complete 100 Code Base Item Enumeration: Organized into 10 clusters of 10 items (CB-001 through CB-100).
- **Section 4: Whole System 100 Improvements:**
  - 8-Part Schema: Complete structure covering Identifier, System UX Impact, Cause, Effect, Zero Harm / Regression Safety, Backward Compatibility, Implementation Blueprint, Binary Verification.
  - Production-Grade Examples: `SYS-IMP-012` (SQLite WAL concurrency engine) and `SYS-IMP-045` (Zero-Storage CI diagnostics).
  - Complete 100 Whole System Item Enumeration: Organized into 10 clusters of 10 items (SYS-001 through SYS-100).
- **Section 5: Companion Antigravity & Cursor Skills Architecture:**
  - Standard `.agents/skills/` and `.cursor/skills/` directory layouts.
  - Complete YAML frontmatter and instruction specifications.
  - IDE skill link syntax invariants (`@[.agents/skills/...]` vs `@[.cursor/skills/...]`).
- **Section 6: Plan Enqueueing & Backlog Invariant:**
  - Mandates enqueuing all discovered items into `.ai-memory/plans/pending/XX-<slug>.md` and `.ai-memory/plans/subtasks/XX-<slug>/`.
  - Zero placeholder rule.
- **Section 7: Verification Protocol:** Matrix of testable assertions and verification methods.

---

## 5. Binary Verification Criteria

| Check | Target | Command / Proof | Expected Outcome | Result |
|:---|:---|:---|:---|:---:|
| Spec 02 Exists | `02-spec/21-app/01-create-prompts-for-100-improvements/02-component-spec.md` | `gitmap lf 02-spec/21-app/01-create-prompts-for-100-improvements` | File present, >500 lines | PASS |
| Subtask 02 Exists | `.ai-memory/plans/subtasks/01-create-prompts-for-100-improvements/02-subtask-component-and-examples.md` | `gitmap lf .ai-memory/plans/subtasks/01-create-prompts-for-100-improvements` | File present, non-empty | PASS |
| Absolute Path Ban | Both files | `gitmap aum search "file:///" 02-spec/21-app/01-create-prompts-for-100-improvements` | 0 hits | PASS |
| Absolute Path Ban | Both files | `gitmap aum search "file:///" .ai-memory/plans/subtasks/01-create-prompts-for-100-improvements` | 0 hits | PASS |
| Windows Drive Ban | Both files | `gitmap aum search "C:\\" 02-spec/21-app/01-create-prompts-for-100-improvements` | 0 hits (except descriptive bans) | PASS |
| Lowercase Naming | File paths | Inspect filenames | Strictly lowercase | PASS |
| Zero Placeholders | Both files | Inspect content | 0 instances of `[N]`, `TODO`, or `...` | PASS |
