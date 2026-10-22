# Subtask 01: Prompt Execution Improvement & Plan Motto Re-Anchoring

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/coding-guideline-and-prompt-execution-improvement/01-subtask-prompt-execution-improvement.md`  
> **Parent Plan:** `.ai-memory/plans/coding-guideline-and-prompt-execution-improvement.md`  
> **Specification:** `02-spec/21-app/24-coding-guideline-and-prompt-execution-improvement/01-execution-prompt-standards-spec.md`  
> **Status:** `[DONE]`  
> **Assigned Role:** Spec Author 01  

---

## 1. Objective & Scope

Establish comprehensive specifications and concrete implementation blueprints to:
1. Re-anchor the standard `/plan` directive motto across all prompt formatters and templates from credit reduction (`to reduce credits` / `to reduce the credits`) to developer excellence and friction minimization:
   ```markdown
   - [/plan](slashCommand;plan) first before doing work for improved work or less friction.
   ```
2. De-couple upfront planning instructions from execution steps to prevent the recurrent anti-pattern where autonomous agents prematurely halt turns or seek permission after completing Phase 1 plans.
3. Author the authoritative specification `02-spec/21-app/24-coding-guideline-and-prompt-execution-improvement/01-execution-prompt-standards-spec.md` and copy to `02-spec/21-app/coding-guideline-and-prompt-execution-improvement/01-execution-prompt-standards-spec.md`.
4. Provide a full 28-occurrence audit matrix and exact line diffs across Letterly prompts, Cursor mirrors, Muse generic template, V6 execution engine prompt/skill, test fixtures, and historical plans.

---

## 2. File Boundaries & Ownership

### 2.1 Owned Deliverables (Spec Author 01)
- `02-spec/21-app/24-coding-guideline-and-prompt-execution-improvement/01-execution-prompt-standards-spec.md`
- `02-spec/21-app/coding-guideline-and-prompt-execution-improvement/01-execution-prompt-standards-spec.md`
- `.ai-memory/plans/subtasks/coding-guideline-and-prompt-execution-improvement/01-subtask-prompt-execution-improvement.md`

### 2.2 Downstream Execution Boundaries (Worker Subagents & Lead)
- `01-prompts/22-letterly/` (9 desktop/execute files)
- `01-prompts/23-cursor-prompts/` (9 cursor mirror files)
- `01-prompts/27-muse-prompts/02-muse-generic.md`
- `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md`
- `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`
- `linter-scripts/tests/test_check_letterly_format.py`

---

## 3. Action Checklist

### Phase 1: Exploration, Audit & Specification (This Subtask)
- [x] Run GitMap high-speed code search (`gitmap aum search`) to inventory all occurrences of credit reduction language across the codebase.
- [x] Catalog all 28 occurrences spanning active prompts, mirrors, generic templates, skills, test fixtures, and historical documentation.
- [x] Author canonical specification `02-spec/21-app/24-coding-guideline-and-prompt-execution-improvement/01-execution-prompt-standards-spec.md` with:
  - System overview detailing the rationale for shifting from cost-skimping to friction reduction.
  - Complete 28-occurrence audit matrix with line numbers and exact before/after strings.
  - De-coupling rationale separating upfront reasoning decorators from deterministic execution phases.
  - Mechanical before/after diffs for all prompt categories.
  - Verification contracts and quality gates.
- [x] Copy specification cleanly to unnumbered path `02-spec/21-app/coding-guideline-and-prompt-execution-improvement/01-execution-prompt-standards-spec.md`.
- [x] Author this subtask execution plan `.ai-memory/plans/subtasks/coding-guideline-and-prompt-execution-improvement/01-subtask-prompt-execution-improvement.md`.
- [x] Verify strict relative Git path compliance (zero absolute filesystem paths or `file:///` URIs).
- [x] Log intent and closure to GitMap Cognitive Heatmap telemetry (`gitmap heatmap log`, `gitmap heatmap update`).

### Phase 2: Downstream Execution & File Updates (Worker Subtasks)
- [ ] Update 9 desktop Letterly prompt files in `01-prompts/22-letterly/` (02, 03, 04, 05, 06, 08, 09, 10, 11).
- [ ] Update 9 Cursor mirror files in `01-prompts/23-cursor-prompts/` (02, 03, 04, 05, 06, 08, 09, 10, 11).
- [ ] Update Muse generic prompt in `01-prompts/27-muse-prompts/02-muse-generic.md`.
- [ ] Update V6 execution engine prompt in `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md`.
- [ ] Update V6 execution skill in `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`.
- [ ] Update Letterly test fixture in `linter-scripts/tests/test_check_letterly_format.py`.

### Phase 3: Verification & Upstream Push (Lead Agent)
- [ ] Execute pytest suite for Letterly format validation:
  ```bash
  gitmap py -m pytest linter-scripts/tests/test_check_letterly_format.py -v
  ```
- [ ] Execute relative path linter across modified files:
  ```bash
  gitmap py linter-scripts/check-relative-paths.py .
  ```
- [ ] Confirm clean git status with strictly modified task files.
- [ ] Execute single atomic GitMap commit and remote push:
  ```bash
  gitmap cpf "Prompts - modernize plan motto to improved work or less friction"
  ```

---

## 4. Verification & Acceptance Criteria

| Criteria ID | Requirement Description | Verification Method |
|---|---|---|
| AC-SPEC-001 | Canonical spec authored with full 28-occurrence audit matrix | Inspect `02-spec/21-app/24-coding-guideline-and-prompt-execution-improvement/01-execution-prompt-standards-spec.md` Section 3 |
| AC-SPEC-002 | Canonical spec copied to unnumbered application spec folder | Inspect `02-spec/21-app/coding-guideline-and-prompt-execution-improvement/01-execution-prompt-standards-spec.md` |
| AC-SPEC-003 | Subtask plan authored with clear boundaries and execution phases | Inspect `.ai-memory/plans/subtasks/coding-guideline-and-prompt-execution-improvement/01-subtask-prompt-execution-improvement.md` |
| AC-SPEC-004 | Zero absolute filesystem paths or `file:///` URIs in all deliverables | Run `gitmap py linter-scripts/check-relative-paths.py .` |
| AC-SPEC-005 | Strict lowercase naming across all authored files and directory paths | Mechanical filename inspection |
| AC-SPEC-006 | GitMap Cognitive Heatmap actions logged with 100% expectation closure | Run `gitmap heatmap status` |
