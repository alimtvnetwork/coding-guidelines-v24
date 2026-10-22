# Master Plan: 05-improvements-folder-prompts

## 1. Task Information
- **Task Title:** Improvements Folder Prompts: Check Pending & Implement Next 10 Improvements
- **Task Slug:** `05-improvements-folder-prompts`
- **Status:** `ACTIVE`
- **Preamble Priority Mandate:** Top-Instruction Priority verified and active.
- **Rules Invariants:**
  - Strictly relative Git paths only (`02-spec/...`, `.ai-memory/...`, `01-prompts/...`).
  - Strict lowercase filenames across all files and folders.
  - Symmetrical skills across `.agents/skills/` and `.cursor/skills/`.
  - Zero raw grep/shell searches; GitMap AUM commands exclusively.
  - Fast-forwarding implementation without executing tests or builds during implementation turns.
  - Separate end-to-end testing and building at the conclusion.
  - Strict coding guidelines: error management (`*appfault.AppError`), monadic results, parameter structs, string constants/enums, function sizing (8–15 lines), file sizing (150–300 lines).
  - Single atomic commit via `gitmap cpf` using hyphen separator without colons.

---

## 2. User Request (Verbatim)

```text
In the improvements folder, can you please do new prompts? One is the check for improvements that would show us how many improvements are pending. Another prompt is implement the improvements, let's say next 10, and repeat the process as I say next and give the same prompt. And every time you implement, you do not execute any test or build. So this would be fast-forwarding. At the end, we will run a separate end-to-end testing and building. But make sure that you follow the coding guideline during the implementation. Error manage, make sure no errors are orphaned. It should actually follow the app error. It should also return wrapped error result if that is required multiple parameters. Any string should be used as a constant or from enum logical steps. The function lines needs to be less than eight lines, eight to 15. Make sure the file size is not more than 150 to 300 lines. So follow this setting. You can include those so that the code quality gets better, and also try to add some of the checkpoints into the execute prompts in V6 parent task, that prompt. Enhance gets a skill and prompt and these two prompts and then sync with other repositories. Do you understand?
```

---

## 3. Scope & Deliverables

1. **New Canonical Prompts in Category 30 (`01-prompts/30-improvements/`):**
   - **`06-check-pending-improvements.md`:** Scans `.ai-memory/plans/pending/`, `.ai-memory/plans/subtasks/`, and active improvement roadmaps. Displays a rich, structured dashboard of pending improvements, breakdown by category (UI, Codebase, Whole System, Detailing, Ideas), priority tiers, readiness scorecards, and exact pending count. Slash command: `/check-pending-improvements`.
   - **`07-implement-next-10-improvements.md`:** Fast-forward implementation engine that consumes the next 10 pending improvements from the backlog without intermediate test/build runs (`NO-BUILD`, `NO-TEST`). Enforces strict coding guidelines (`*appfault.AppError`, `Result[T]`, parameter structs, string constants/enums, 8-15 line functions, 150-300 line files). Repeatable on command (`next`). Defers compilation/test verification to a dedicated end-to-end testing phase. Slash command: `/implement-next-10-improvements`.
2. **Companion Skills (Antigravity & Cursor):**
   - `.agents/skills/check-pending-improvements/skill.md`
   - `.cursor/skills/check-pending-improvements/skill.md`
   - `.agents/skills/implement-next-10-improvements/skill.md`
   - `.cursor/skills/implement-next-10-improvements/skill.md`
3. **Enhance V6 Parent Task Prompt & Skills:**
   - Update `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` and companion skill `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` (and `.cursor/skills/`) to embed fast-forwarding micro-batch checkpoints, strict coding guideline parameters, and the separated end-to-end verification gate.
4. **Master Catalog & Registry Updates:**
   - `01-prompts/30-improvements/readme.md` (Catalog table expanded to 7 prompts)
   - `01-prompts/readme.md` (Category 30 row updated)
   - `.ai-memory/prompts.md` (Register prompts 06 and 07, slugs, and slash commands)
   - `.ai-memory/plans/readme.md` & `02-spec/21-app/readme.md` (Register task 05)
5. **Multi-Repository Fleet Synchronization:**
   - Propagate new prompts and skills across all 43 fleet repositories via `38-sync-prompts-skills-scripts.py`.

---

## 4. Subtask Allocation

| Task ID | Title | Assigned Role | Status | Owned Files |
|:---|:---|:---|:---:|:---|
| `Task-01` | Specification Authoring & Subtasks | Spec Authors 01, 02, 03 | `COMPLETED` | `02-spec/21-app/05-improvements-folder-prompts/*`, `.ai-memory/plans/subtasks/05-improvements-folder-prompts/*` |
| `Task-02` | Author Prompts 06 & 07 and Update Category Readme | Worker 01 | `COMPLETED` | `01-prompts/30-improvements/06-check-pending-improvements.md`, `01-prompts/30-improvements/07-implement-next-10-improvements.md`, `01-prompts/30-improvements/readme.md` |
| `Task-03` | Author Symmetric Companion Skills | Worker 02 | `COMPLETED` | `.agents/skills/check-pending-improvements/*`, `.cursor/skills/check-pending-improvements/*`, `.agents/skills/implement-next-10-improvements/*`, `.cursor/skills/implement-next-10-improvements/*` |
| `Task-04` | Enhance V6 Parent Task Prompt & Skills | Worker 03 | `COMPLETED` | `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`, `.agents/skills/execute-parent-task-with-n-steps-v6/*`, `.cursor/skills/execute-parent-task-with-n-steps-v6/*` |
| `Task-05` | Update Master Catalogs, Readmes & Registries | Worker 03 / Lead | `COMPLETED` | `01-prompts/readme.md`, `.ai-memory/prompts.md`, `.ai-memory/plans/readme.md`, `02-spec/21-app/readme.md` |
| `Task-06` | Fleet Sync & Atomic Push | Lead Orchestrator | `COMPLETED` | `03-ai-scripts/38-sync-prompts-skills-scripts.py`, all 43 repositories |
