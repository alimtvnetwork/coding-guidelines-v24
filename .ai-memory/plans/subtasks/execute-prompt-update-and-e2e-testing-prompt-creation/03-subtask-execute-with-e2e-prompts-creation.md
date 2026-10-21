# Subtask 03: Create 08-execute-with-e2e.md Prompts with Mandatory E2E and Release Minor Bump

## 1. Identity & Context
- **Subtask ID:** `Subtask-03`
- **Parent Task:** `task-20261021023054-execute-prompt-update-and-e2e-testing-prompt-creation`
- **Status:** `DONE`
- **Owner:** `Worker 02`

---

## 2. Deliverables
- [x] Author `01-prompts/14-execute/08-execute-with-e2e.md` based on canonical V6 architecture.
- [x] Mirror to `01-prompts/v4/14-execute/08-execute-with-e2e.md`.

---

## 3. Key Prompt Capabilities
- **Mandatory Local Testing:** Runs all test suites locally (`go test ./...`, `pytest`, `npm test`) with 100% green exit code 0 required.
- **Turn 1 Recitation:** Displays `git log -n 5 --oneline` and inventories recent test cases and scenarios.
- **Audit Log Verification:** Reviews `.ai-memory/heatmap.db` and SQLite agent task logs (`gitmap agent log`, `gitmap heatmap tree`).
- **Release Ceremony Trigger:** Commands minor version bump (`npm run bump minor` or `python 03-ai-scripts/37-bump-version.py --tier minor`) and pushes via GitMap as soon as everything is green.

---

## 4. Verification Criteria
- Files exist and adhere to strictly lowercase naming.
- Strict relative Git paths only.
- Tested and verified format against V6 standard.
