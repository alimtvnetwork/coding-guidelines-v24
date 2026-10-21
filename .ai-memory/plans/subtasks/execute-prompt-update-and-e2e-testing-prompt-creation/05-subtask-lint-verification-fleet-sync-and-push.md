# Subtask 05: Validate Linters, Synchronize across 43 Fleet Repositories, and Atomic GitMap Push

## 1. Identity & Context
- **Subtask ID:** `Subtask-05`
- **Parent Task:** `task-20261021023054-execute-prompt-update-and-e2e-testing-prompt-creation`
- **Status:** `PENDING`
- **Owner:** `Lead Orchestrator`

---

## 2. Deliverables
- [ ] Run relative path check: `python linter-scripts/check-relative-paths.py` (0 errors).
- [ ] Run forbidden strings check: `python linter-scripts/check-forbidden-strings.py`.
- [ ] Perform dry-run and live fleet sync via `python 03-ai-scripts/38-sync-prompts-skills-scripts.py`.
- [ ] Inspect heatmap closure: `gitmap heatmap status` and `gitmap heatmap tree`.
- [ ] Perform atomic commit and push: `gitmap cpf "prompts - update v6 test build ban and create execute with e2e prompt"`.

---

## 3. Verification Criteria
- Zero relative path violations.
- All 43 fleet repositories synchronized according to the 5 boundary rules.
- Git working tree clean after push.
