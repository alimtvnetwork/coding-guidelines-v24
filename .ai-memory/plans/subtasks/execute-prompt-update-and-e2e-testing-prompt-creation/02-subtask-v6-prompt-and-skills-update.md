# Subtask 02: Update V6 Prompts and Skills to Ban Unrequested Builds and Test Suites

## 1. Identity & Context
- **Subtask ID:** `Subtask-02`
- **Parent Task:** `task-20261021023054-execute-prompt-update-and-e2e-testing-prompt-creation`
- **Status:** `COMPLETED`
- **Owner:** `Worker 01`

---

## 2. Deliverables
- [x] Update `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` across 10 touchpoints.
- [x] Update `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md` across 10 touchpoints.
- [x] Update `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` across 10 touchpoints.
- [x] Update `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md` across 10 touchpoints.

---

## 3. Specific Invariants to Inject
- "Every time you write code, you don't build the code."
- "Never run all tests locally or full build unless explicitly asked."
- "Only do end-to-end testing based on mentioned test cases if explicitly asked."
- "Do CI/CD verification only if asked; if not asked explicitly, don't do it."

---

## 4. Verification Criteria
- All 4 files updated with identical wording across the 10 touchpoints.
- Git diff confirms only specified lines are modified without reflowing unaffected sections.
