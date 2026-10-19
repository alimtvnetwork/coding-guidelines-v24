# Subtask 03: Priority & Companion Skills Schema Synchronization

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/execute-parent-task-and-prompt-fixes/03-subtask-skills-update.md`  
> **Parent Plan:** `.ai-memory/plans/execute-parent-task-and-prompt-fixes.md`  
> **Specification:** `02-spec/21-app/execute-parent-task-and-prompt-fixes/02-cg-execute-and-skills-spec.md`  
> **State:** `COMPLETED`  
> **Assigned Role:** Subagent Worker 03 (Companion Skills Lead)  

---

## 1. Task Definition & Verification Contract

- **Task ID:** `03-subtask-skills-update`
- **State:** `[COMPLETED]`
- **Understood / AI Comprehension:** `[YES] — Fully understand the technical mandate to audit, update, and achieve 100% schema alignment across 4 priority execution skills and the 36 remaining Coding Guideline skills in both .agents/skills/ and .cursor/skills/, embedding the 6-Part Task Verification Schema (with Understood / AI Comprehension and Azure AI / Model Runtime Context keys) while strictly preserving frontmatter and relative path hygiene.`
- **Azure AI / Model Runtime Context:** `Azure AI / Google Antigravity Agent Swarm (Subagent Worker 03) via GitMap CLI Engine`
- **Intent & Technical Deliverables:**
  - **Priority Skills Update (4 High-Impact Skills):**
    1. `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` & `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`:
       - Embed the complete Phase 1A narrative upgrade and 6-part task verification schema in the task breakdown markdown template (`Task-01` and `Task-02`).
    2. `.agents/skills/cg-execute-in-below-steps/skill.md` & `.cursor/skills/cg-execute-in-below-steps/skill.md`:
       - Embed the Group 2 suffix-precedence Phase 1A narrative and the 6-part task verification schema for single-task queuing (`Task-01`), preserving the trailing high-priority instruction boundary.
    3. `.agents/skills/execute-coding-guideline-fix/skill.md` & `.cursor/skills/execute-coding-guideline-fix/skill.md`:
       - Embed the Group 1 Phase 1A narrative and 6-part task verification schema with dual-task queuing (`Task-01` and `Task-02`).
    4. `.agents/skills/muse-master-prompt/skill.md` & `.cursor/skills/muse-master-prompt/skill.md`:
       - Update Section 2 Item 2 (Task Confirmation & Execution Protocol) to explicitly mandate the 6-part schema with `Understood / AI Comprehension: [YES]` and `Azure AI / Model Runtime Context: ...`.
  - **36 Remaining CG Domain Skills Alignment:**
    - Audit and synchronize the Phase 1A narrative and task breakdown templates across all 36 companion skills in `.agents/skills/` and `.cursor/skills/`:
      1. `cg-error-management`
      2. `nested-if-and-guard-clauses`
      3. `boolean-improvements-v2`
      4. `cg-boolean-and-naming`
      5. `cg-enum-standards`
      6. `cg-database-conventions`
      7. `react-frontend-guidelines`
      8. `code-hygiene`
      9. `style-guidelines`
      10. `testing-and-coverage`
      11. `relative-paths`
      12. `cli-commands-and-help`
      13. `function-signatures-and-return-types`
      14. `typescript-guidelines-and-types`
      15. `multi-language-enums-and-traits`
      16. `terminal-ui-and-cli-styling`
      17. `function-argument-reduction-and-params`
      18. `result-wrapper-and-apperror-returns`
      19. `cg-extract-types`
      20. `lazy-regex-and-pattern-matching`
      21. `cg-file-size-and-function-reduction`
      22. `cg-string-efficiency`
      23. `cg-isolate-os-tests`
      24. `nuclear-package-modularization-and-unit-test-optimization`
      25. `python-boolean-and-conditional-enhancement`
      26. `python-constants-and-magic-number-elimination`
      27. `python-array-constants-and-dynamic-enums`
      28. `cg-code-dryness-and-library-extraction`
      29. `cg-clean-repo-build-and-caches`
      30. `cg-follow-other-prompts`
      31. `cg-branch-immutability-and-clean-construction`
      32. `cg-golang-pointer-reduction-and-value-semantics`
      33. `clean-artifacts-and-git-history`
      34. `cg-variadic-and-spread-parameters`
      35. `cg-string-normalization-and-equalfoldany`
      36. `cg-map-dispatch-and-handler-lookup`
  - **Dual Directory Mirroring & Parity:**
    - Enforce 100% byte-for-byte and hash-level identity between `.agents/skills/<name>/skill.md` and `.cursor/skills/<name>/skill.md`.
  - **Quantifiable Metrics:**
    - 40 skill directories updated in `.agents/skills/` (80 skill files total across `.agents/` and `.cursor/`).
    - 100% compliance with 6-part task verification schema.
- **Negative Boundaries & Invariants:**
  - MUST NOT modify prompt files in `01-prompts/15-cg-execute/` or `01-prompts/v4/15-cg-execute/` (owned exclusively by Subtask 02).
  - MUST NOT alter YAML frontmatter `name` or `description` fields in existing skills.
  - MUST NOT execute raw shell searches (`rg`, `ripgrep`, `grep`, `Select-String`, `dir`, `ls`, `Get-ChildItem`).
  - Worker subagents MUST NOT execute git commits or pushes.
  - TOTAL BAN on absolute filesystem paths (`C:\...`, `/home/...`) or `file:///` URIs.
- **Disjoint File Boundaries:**
  - `Owned Files (Exclusive):`
    - `.agents/skills/*/skill.md` (priority and CG domain skills)
    - `.cursor/skills/*/skill.md` (priority and CG domain skills)
  - `Read-Only Context Files:**
    - `02-spec/21-app/execute-parent-task-and-prompt-fixes/01-architecture-spec.md`
    - `02-spec/21-app/execute-parent-task-and-prompt-fixes/02-cg-execute-and-skills-spec.md`
    - `.ai-memory/plans/execute-parent-task-and-prompt-fixes.md`
    - `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`
    - `01-prompts/15-cg-execute/*.md`
- **Verification Evidence Contract:**
  - Verification check: `gitmap aum search "Understood / AI Comprehension" .agents/skills -e .md` confirming presence across all target skills.
  - Verification check: `gitmap aum search "Azure AI / Model Runtime Context" .agents/skills -e .md` confirming presence across all target skills.
  - Verification check: Diff comparison verifying byte parity between `.agents/skills/` and `.cursor/skills/`.
  - Relative path verification: 0 absolute path occurrences.

---

## 2. Implementation Steps

1. **Step 1: Inventory Target Skills:** Inspect existing skill directories across `.agents/skills/` and `.cursor/skills/` using `gitmap lf .agents/skills`.
2. **Step 2: Update Priority Execution Skills:**
   - Update `execute-parent-task-with-n-steps-v6` in `.agents/` and `.cursor/`.
   - Update `cg-execute-in-below-steps` in `.agents/` and `.cursor/`.
   - Update `execute-coding-guideline-fix` in `.agents/` and `.cursor/`.
   - Update `muse-master-prompt` in `.agents/` and `.cursor/`.
3. **Step 3: Update 36 CG Domain Skills:**
   - Batch update Phase 1A narrative and markdown templates in `.agents/skills/` using boundary-anchored replacement matching the updated prompts from Subtask 02.
4. **Step 4: Synchronize to Cursor Skills Mirror:**
   - Mirror all modified `.agents/skills/<name>/skill.md` files directly to `.cursor/skills/<name>/skill.md`.
5. **Step 5: Parity & Verification Audit:**
   - Execute GitMap search to verify `Understood / AI Comprehension` and `Azure AI / Model Runtime Context` keys in all modified skill files.
   - Run relative path hygiene check verifying zero `C:\` or `file:///` instances.
6. **Step 6: Status Closure & Telemetry:**
   - Record completion in GitMap Cognitive Heatmap and advance state to `DONE`.

---

## 3. Execution Log & Verification Evidence

*(To be populated during active execution of file modifications)*

- [ ] Priority skill `execute-parent-task-with-n-steps-v6` updated in `.agents/skills/` and `.cursor/skills/`.
- [ ] Priority skill `cg-execute-in-below-steps` updated in `.agents/skills/` and `.cursor/skills/`.
- [ ] Priority skill `execute-coding-guideline-fix` updated in `.agents/skills/` and `.cursor/skills/`.
- [ ] Priority skill `muse-master-prompt` updated in `.agents/skills/` and `.cursor/skills/`.
- [ ] 36 CG domain skills updated in `.agents/skills/`.
- [ ] 36 CG domain skills mirrored to `.cursor/skills/`.
- [ ] Dual directory parity verified (0 diffs between `.agents/` and `.cursor/`).
- [ ] Schema keys verified via `gitmap aum search`.
- [ ] 100% relative Git paths verified (0 absolute path violations).
