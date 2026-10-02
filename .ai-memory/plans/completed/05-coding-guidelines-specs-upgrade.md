# Completed Plan: Coding Guidelines Specs Upgrade & Slide Presentation Specs Alignment

> **/goal** Transform all remaining specification files in `02-spec/02-coding-guidelines/` (consistency reports across all language modules) and `spec-coding-guideline/` (internal slide presentation guidelines) into prompt-driven specifications with actionable checklists and testable acceptance criteria.
> **/learn** Enforce the uniform 4-part specification contract: AI Execution Prompt header, Actionable CI/CD & Agent Checklist, Core Rules with code examples, and testable Acceptance Criteria (`Given / When / Then`) with deterministic verification commands.

**Version:** 1.0.0  
**Completed:** 2026-10-02  
**Status:** Completed  
**AI Confidence:** Production-Ready  
**Canonical Architecture Spec:** [`02-spec/21-app/04-coding-guidelines-actionable-checklists-and-acceptance-criteria/01-architecture-spec.md`](../../02-spec/21-app/04-coding-guidelines-actionable-checklists-and-acceptance-criteria/01-architecture-spec.md)  
**Task DB Ledger:** `.ai-memory/temp-agents/05-coding-guidelines-specs-upgrade/ledger.md`  

---

## 🎯 Completed Deliverables & Evidence

1. **Internal Slide Presentation Specs (`spec-coding-guideline/`):**
   - Transformed all 9 specification files into active AI execution prompts:
     - `spec-coding-guideline/readme.md` (Index and slide system summary, `AC-CG-SLIDE-001`)
     - `spec-coding-guideline/02-architecture.md` (Store and component architecture, `AC-CG-SLIDE-002`)
     - `spec-coding-guideline/03-slide-authoring.md` (Slide authoring conventions, `AC-CG-SLIDE-003`)
     - `spec-coding-guideline/04-design-tokens.md` (Theme and color tokens, `AC-CG-SLIDE-004`)
     - `spec-coding-guideline/05-animation-primitives.md` (Animation standards, `AC-CG-SLIDE-005`)
     - `spec-coding-guideline/06-curriculum.md` (Curriculum mapping, `AC-CG-SLIDE-006`)
     - `spec-coding-guideline/07-build-and-zip-pipeline.md` (Packaging specs, `AC-CG-SLIDE-007`)
     - `spec-coding-guideline/08-gif-generation.md` (Preview GIF pipeline, `AC-CG-SLIDE-008`)
     - `spec-coding-guideline/09-quality-and-offline.md` (Quality gates and offline viewers, `AC-CG-SLIDE-009`)
   - Each file includes `## 🎯 Actionable CI/CD & Agent Checklist` and `## Verification & Acceptance Criteria`.

2. **Module Consistency Reports (`02-spec/02-coding-guidelines/**/99-consistency-report.md`):**
   - Upgraded all 19 module consistency reports to active AI prompt standards with testable criteria:
     - `01-cross-language/99-consistency-report.md`
     - `01-cross-language/02-boolean-principles/99-consistency-report.md`
     - `01-cross-language/04-code-style/99-consistency-report.md`
     - `01-cross-language/15-master-coding-guidelines/99-consistency-report.md`
     - `01-cross-language/16-static-analysis/99-consistency-report.md`
     - `02-typescript/99-consistency-report.md`
     - `03-golang/99-consistency-report.md`
     - `03-golang/01-enum-specification/99-consistency-report.md`
     - `03-golang/04-golang-standards-reference/99-consistency-report.md`
     - `04-php/99-consistency-report.md`
     - `04-php/07-php-standards-reference/99-consistency-report.md`
     - `05-rust/99-consistency-report.md`
     - `06-ai-optimization/99-consistency-report.md`
     - `07-csharp/99-consistency-report.md`
     - `08-file-folder-naming/99-consistency-report.md`
     - `11-security/99-consistency-report.md`
     - `11-security/01-axios-version-control/99-consistency-report.md`
     - `12-python/99-consistency-report.md`
     - `13-cpp/99-consistency-report.md`

3. **Newline and Whitespace Normalization:**
   - All 223 files in `02-spec/02-coding-guidelines/` and 9 files in `spec-coding-guideline/` verified to use clean Unix LF line endings without trailing whitespace.
   - `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only` passes with exit code 0.
   - `python 03-ai-scripts/05-guideline-autofixer.py spec-coding-guideline --check-only` passes with exit code 0.

4. **Relative Path and Zero-Storage Verification:**
   - `python linter-scripts/check-relative-paths.py` verified 0 absolute paths or file:/// URIs across all 3500+ tracked repository files.

---

## 📋 Subtask Execution Breakdown

| Subtask ID | Task Code | Assigned Agent | Status | Verified Evidence |
|:---|:---|:---|:---|:---|
| 1 | Task-01 | Lead | DONE | Verified architecture spec and rollout spec in `02-spec/21-app/` |
| 2 | Task-02A | Worker 01 | DONE | Upgraded `spec-coding-guideline/02..05` with AI headers and checklists |
| 3 | Task-02B | Worker 02 | DONE | Upgraded `spec-coding-guideline/06..09` with AI headers and checklists |
| 4 | Task-03A | Worker 01 | DONE | Upgraded Cross-Language, TS, and Go consistency reports |
| 5 | Task-03B | Worker 02 | DONE | Upgraded PHP, Rust, AI, C#, Polyglot consistency reports; normalized newlines |
| 6 | Task-04 | Lead | DONE | Verified all 223 guidelines + 9 slide specs with targeted linters (exit 0) |

---

## 🔒 Verification & Acceptance Criteria

### AC-CG-CONSISTENCY-MASTER: Master Coding Guidelines Upgrade Conformance

**Given** All specification and guideline files in `02-spec/02-coding-guidelines/` and `spec-coding-guideline/`.  
**When** Linters and CI/CD quality gates audit the entire documentation tree.  
**Then** Every file includes an active prompt header, actionable checklist, testable acceptance criteria, clean Unix LF endings, and 100% relative paths.

**Verification command:**
```bash
python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only && python linter-scripts/check-relative-paths.py
```
**Expected:** exit 0. Zero violations detected.
