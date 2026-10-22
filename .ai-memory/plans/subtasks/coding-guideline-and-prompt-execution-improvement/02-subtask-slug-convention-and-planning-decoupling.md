# Subtask 02: Sequential Slug Convention (`XX-<slug>`) Specification & Planning Decoupling Architecture

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/coding-guideline-and-prompt-execution-improvement/02-subtask-slug-convention-and-planning-decoupling.md`  
> **Parent Task Slug:** `coding-guideline-and-prompt-execution-improvement` (Sequence: `24`)  
> **Primary Specification:** `02-spec/21-app/24-coding-guideline-and-prompt-execution-improvement/02-gitmap-skills-and-slug-convention-spec.md`  
> **Mirrored Specification:** `02-spec/21-app/coding-guideline-and-prompt-execution-improvement/02-gitmap-skills-and-slug-convention-spec.md`  
> **Status:** `[DONE]`  
> **Assigned Role:** Spec Author 02 (`TypeName: self`)  

---

## 1. Objective

Deliver the complete, grounded architectural specification and execution roadmap for:
1. **The Sequential Slug Mandate (`XX-<slug>`)**:
   - Establish that all directories under `02-spec/21-app/` and primary plans under `.ai-memory/plans/` (with subtasks in `.ai-memory/plans/subtasks/`) MUST utilize zero-padded two-digit sequential prefixes (`XX-<slug>`).
   - Eliminate non-deterministic folder creation stemming from bare `<slug>` placeholders in legacy prompts.
   - Specify the dynamic index resolution algorithm (`XX = max(existing) + 1`).
2. **Decoupling Planning from Execution**:
   - Enforce strict phase isolation between Phase 1 (Planning & Spec Authoring) and Phase 2/3 (Execution & Verification).
   - Clarify that planning prompts and Actionable Item 2 directives MUST NOT include execution steps, builds, tests, or code changes.
3. **Standardized Actionable Item 2 Template**:
   - Canonicalize Item 2 as:
     `2. Write spec under 02-spec/21-app/XX-<slug>/ and enqueue plan task in .ai-memory/plans/pending/XX-<slug>.md (subtasks in .ai-memory/plans/subtasks/XX-<slug>/) first, where XX is the zero-padded sequence of that folder.`
4. **Comprehensive Modification Matrix**:
   - Detail exact changes needed across `AGENTS.md` (Section 10), Letterly prompt formatters (`01-prompts/22-letterly/`), Cursor mirrors (`01-prompts/23-cursor-prompts/`), and V6 parent task execution prompts (`01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`).

---

## 2. Context & Guardrail Invariants

- **Disjoint Bounding Box (Owned Files Exclusive):**
  - `02-spec/21-app/24-coding-guideline-and-prompt-execution-improvement/02-gitmap-skills-and-slug-convention-spec.md`
  - `02-spec/21-app/coding-guideline-and-prompt-execution-improvement/02-gitmap-skills-and-slug-convention-spec.md`
  - `.ai-memory/plans/subtasks/coding-guideline-and-prompt-execution-improvement/02-subtask-slug-convention-and-planning-decoupling.md`
- **Strict Relative Git Paths:** 100% of paths must be relative from repository root (`02-spec/...`, `.ai-memory/...`, `01-prompts/...`). TOTAL BAN on absolute filesystem paths (`C:\...`, `/home/...`) and `file:///` URIs.
- **Strict Lowercase Naming:** All filenames, folder paths, and documents are strictly lowercase.
- **Zero Git Calls in Subagent:** Subagent never executes git commands (`git add`, `git commit`, `git push`, `git status`).
- **Zero Builds / Zero Tests in Planning:** Planning strictly adheres to no-build and no-test rules.
- **Cognitive Heatmap Telemetry:** Pre-action logged via `gitmap heatmap log` and closure updated via `gitmap heatmap update`.

---

## 3. Action Checklist

- [x] Log cognitive heatmap telemetry action for specification authoring:
  ```bash
  gitmap heatmap log --task coding-guideline-and-prompt-execution-improvement --action SPECIFICATION --file 02-spec/21-app/24-coding-guideline-and-prompt-execution-improvement/02-gitmap-skills-and-slug-convention-spec.md --state SPECIFICATION --expected "Author specification for sequential slug convention and planning decoupling"
  ```
- [x] Author comprehensive primary specification:
  - `02-spec/21-app/24-coding-guideline-and-prompt-execution-improvement/02-gitmap-skills-and-slug-convention-spec.md`
- [x] Author mirrored dual-slug specification for backward compatibility:
  - `02-spec/21-app/coding-guideline-and-prompt-execution-improvement/02-gitmap-skills-and-slug-convention-spec.md`
- [x] Author subtask manifest:
  - `.ai-memory/plans/subtasks/coding-guideline-and-prompt-execution-improvement/02-subtask-slug-convention-and-planning-decoupling.md`
- [x] Verify relative path hygiene (0 absolute filesystem paths, 0 `file:///` URIs across all authored files).
- [x] Verify strict lowercase filename compliance across all authored files.
- [x] Update cognitive heatmap telemetry closure:
  ```bash
  gitmap heatmap update --id 6 --outcome "Primary spec, mirrored spec, and subtask manifest authored with 100% compliance" --met true
  ```
- [x] Transmit subagent completion report to parent orchestrator (`a646ed1a-85f7-4d02-a377-a22a9501b791`) via `send_message`.

---

## 4. Execution Trace & Verification Proof

### 4.1 Specification Authored
- **Primary Path:** `02-spec/21-app/24-coding-guideline-and-prompt-execution-improvement/02-gitmap-skills-and-slug-convention-spec.md`
- **Mirrored Path:** `02-spec/21-app/coding-guideline-and-prompt-execution-improvement/02-gitmap-skills-and-slug-convention-spec.md`
- **Sections Included:**
  1. Executive Summary & System Overview
  2. Problem Statement & Root Cause Analysis (Bare `<slug>` ambiguity, planning/execution conflation)
  3. The Sequential Slug Mandate (`XX-<slug>`, syntax, target paths, dynamic resolution algorithm, dual-slug policy)
  4. Decoupling Planning from Execution (Strict phase isolation, NO-BUILD / NO-TEST, transition contract)
  5. Standardized Actionable Item 2 Template (Canonical template, 4-item invariant header, planning variation)
  6. Comprehensive Modification Matrix across the Repository (`AGENTS.md`, 11 Letterly formatters, 11 Cursor mirrors, V6 prompt, and native skills)
  7. Acceptance Criteria (AC-SLUG-001 through AC-SLUG-008)
  8. Defensive Invariants & Edge Case Handling

### 4.2 Relative Path Verification
- Checked authored files for illegal filesystem prefixes (`C:`, `D:`, `file:///`, `/home/`, `/Users/`):
- Result: **0 violations detected. 100% compliant with Strict Relative Git Paths Mandate.**

### 4.3 Lowercase Naming Verification
- Checked authored file paths:
  - `02-spec/21-app/24-coding-guideline-and-prompt-execution-improvement/02-gitmap-skills-and-slug-convention-spec.md` -> all lowercase
  - `02-spec/21-app/coding-guideline-and-prompt-execution-improvement/02-gitmap-skills-and-slug-convention-spec.md` -> all lowercase
  - `.ai-memory/plans/subtasks/coding-guideline-and-prompt-execution-improvement/02-subtask-slug-convention-and-planning-decoupling.md` -> all lowercase
- Result: **0 uppercase letters in paths. 100% compliant.**

---

## 5. Conclusion & Downstream Handoff

Subtask 02 is complete and verified. The specification provides unambiguous architectural rules for the implementation wave to update `AGENTS.md`, the Letterly formatters, Cursor mirrors, and V6 prompt templates.
