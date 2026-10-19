# Subtask 01: Preflight Verification & V6 Master Prompts Schema Upgrade

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/execute-parent-task-and-prompt-fixes/01-subtask-preflight-and-v6-prompts.md`  
> **Parent Plan:** `.ai-memory/plans/execute-parent-task-and-prompt-fixes.md`  
> **Specification:** `02-spec/21-app/execute-parent-task-and-prompt-fixes/01-architecture-spec.md`  
> **State:** `COMPLETED`  
> **Assigned Role:** Subagent Worker 01 (V6 Prompts & Skills Lead)  

---

## 1. Task Definition & Verification Contract

- **Task ID:** `01-subtask-preflight-and-v6-prompts`
- **State:** `[COMPLETED]`
- **Understood / AI Comprehension:** `[YES] — Fully understand the requirement to upgrade Phase 1A narrative instructions and task breakdown templates in V6 master prompts and companion skills to mandate explicit Understood/AI Comprehension verification and Azure AI/Model Runtime context declarations.`
- **Azure AI / Model Runtime Context:** `Azure AI / Google Antigravity Agent Swarm (Subagent Worker 01) via GitMap CLI Engine`
- **Intent & Technical Deliverables:**
  - Update `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` Phase 1A narrative Step 2 to mandate the `Understood / AI Comprehension` key and `Azure AI / Model Runtime Context` key in every task item.
  - Update `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` markdown template under `### 📋 Confirmed Task Breakdown & Requirement Ingestion` to include the full 6-part task verification schema across `Task-01` and `Task-02`.
  - Apply 1:1 identical updates to `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md`.
  - Update companion skill definitions in `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` and `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`.
  - Quantifiable metrics: 4 files modified, 100% field parity with 02-spec/21-app/execute-parent-task-and-prompt-fixes/01-architecture-spec.md.
- **Negative Boundaries & Invariants:**
  - MUST NOT alter core rules R1 through R12 outside Phase 1A narrative and template sections.
  - MUST NOT modify any prompts in `01-prompts/15-cg-execute/` (reserved for Subtask 02).
  - MUST NOT execute raw shell searches (`rg`, `ripgrep`, `grep`, `Select-String`, `dir`, `ls`, `Get-ChildItem`).
  - Worker subagents MUST NOT execute git commits or pushes.
- **Disjoint File Boundaries:**
  - `Owned Files (Exclusive):`
    - `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`
    - `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md`
    - `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`
    - `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`
  - `Read-Only Context Files:`
    - `02-spec/21-app/execute-parent-task-and-prompt-fixes/01-architecture-spec.md`
    - `.ai-memory/plans/execute-parent-task-and-prompt-fixes.md`
    - `AGENTS.md`
- **Verification Evidence Contract:**
  - Structural search via `gitmap aum search "Understood / AI Comprehension" 01-prompts/14-execute -e .md` confirming presence.
  - Structural search via `gitmap aum search "Azure AI / Model Runtime Context" 01-prompts/14-execute -e .md` confirming presence.
  - Relative path audit verifying 0 occurrences of absolute filesystem paths or `file:///` URIs.

---

## 2. Implementation Steps

1. **Step 1: Inspect Target V6 Prompts:** Inspect line ranges in `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` around Phase 1A Step 2 and markdown templates.
2. **Step 2: Update Phase 1A Narrative (Step 2):** Integrate explicit requirement for `Understood / AI Comprehension` and `Azure AI / Model Runtime Context`.
3. **Step 3: Update Markdown Template:** Embed the 6-part task verification schema in the code block under `### 📋 Confirmed Task Breakdown & Requirement Ingestion`.
4. **Step 4: Synchronize to V4 Mirror:** Apply identical changes to `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md`.
5. **Step 5: Synchronize Skills:** Update `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` and `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`.
6. **Step 6: Verification & Telemetry:** Log outcomes to GitMap Cognitive Heatmap and update subtask status to `DONE`.

---

## 3. Execution Log & Verification Evidence

*(To be populated during active execution of file modifications)*

- [ ] V6 Master Prompt narrative updated.
- [ ] V6 Master Prompt template updated.
- [ ] V4 Master Prompt mirrored.
- [ ] Antigravity skill updated.
- [ ] Cursor skill updated.
- [ ] Relative paths verified (0 absolute path violations).
