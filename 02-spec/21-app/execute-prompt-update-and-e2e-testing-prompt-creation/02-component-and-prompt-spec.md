# Component & Prompt Specification: Execute V6 Updates & Execute with E2E Prompt

## 1. Scope of File Modifications

This specification defines the exact component anatomy, section layouts, and line-level changes across two functional groups:
1. **Group 1: Existing V6 Prompts and Skills (Surgical Updates)**
   - `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`
   - `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md`
   - `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`
   - `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`
2. **Group 2: New E2E Execution Prompts, Skills and Registry Updates**
   - `01-prompts/14-execute/08-execute-with-e2e.md`
   - `01-prompts/v4/14-execute/08-execute-with-e2e.md`
   - `.agents/skills/execute-with-e2e/skill.md`
   - `.cursor/skills/execute-with-e2e/skill.md`
   - `01-prompts/readme.md`
   - `01-prompts/14-execute/readme.md`
   - `01-prompts/v4/14-execute/readme.md`
   - `.ai-memory/prompts.md`

---

## 2. Group 1: V6 Prompt & Skill Modifications (10 Surgical Touchpoints)

All 4 target files in Group 1 share identical logical structures. The following 10 touchpoints must be updated consistently across all four files:

### Touchpoint 1: Pillar 3 Recitation (Phase 1A Step 2)
- **Target Phrase:** `zero builds/full test suites (CI verification only)`
- **Replacement:** `zero builds/tests (never run all tests locally or full build unless explicitly asked; every time code is written, do not build the code; only do end-to-end testing based on mentioned test cases if explicitly asked; do CI/CD verification only if asked, otherwise do not do it)`

### Touchpoint 2: Pillar 3 Markdown Template (Phase 1A Step 2)
- **Target Phrase:** `zero builds/full test suites (CI verification only)`
- **Replacement:** `zero builds/tests (never run all tests locally or full build unless explicitly asked; every time you write code, do not build the code; only do end-to-end testing based on mentioned test cases if explicitly asked; do CI/CD verification only if asked, otherwise do not do it)`

### Touchpoint 3: Scope Control Rules (Multi-Agent Partitioning)
- **Target Phrase:** `Worker 02 implements verification/linters/companion tests`
- **Replacement:** `Worker 02 implements verification/linters/memory docs`

### Touchpoint 4: Core Operational Rule R1
- **Target Phrase:** `- **R1 Zero Builds or Test Suites (TOTAL BAN).** NEVER run ... CI verifies builds and suites. Routine turns must never waste time on heavy compilation/tests. Only explicit user command lifts this.`
- **Replacement:** `- **R1 Zero Builds, Test Suites & CI/CD Verification Ban (TOTAL BAN).** NEVER run `go build`, `npm run build`, `vite build`, `cargo build`, `go test ./...`, `pytest`, `npm test`, or `03-ai-scripts/06-cicd-local-runner.py`. Every time you write code, you do not build the code. Never run all tests locally or full build unless explicitly asked by the user. Only do end-to-end (E2E) testing based on mentioned test cases if explicitly asked. Do CI/CD verification (e.g. `gitmap pe`, `gitmap pipeline-ai`, GitHub Actions) only if asked; if not asked explicitly, don't do it. Routine turns must never waste steps or time on compilation, local test suites, E2E flows, or remote CI/CD monitoring unless explicitly commanded.`

### Touchpoint 5: High-Speed Command Table (Pipeline Waiting Alias)
- **Target Phrase:** `Non-polling dynamic ETA CI/CD monitor`
- **Replacement:** `Non-polling dynamic ETA CI/CD monitor (use ONLY if CI/CD verification is explicitly asked)`

### Touchpoint 6: Dispatch Payload Examples (Workers 01, 02, 03)
- **Target Phrase:** `- Targeted Verification: Run only fast file-scoped linters (e.g. check-relative-paths.py, guideline-autofixer.py --check-only). Zero build or test suites.`
- **Replacement:** `- Targeted Verification: Run only fast file-scoped linters (e.g. check-relative-paths.py, guideline-autofixer.py --check-only). Every time you write code, do not build the code. Never run all tests locally or full build unless explicitly asked. Only do end-to-end testing based on mentioned test cases if explicitly asked. Do CI/CD verification only if asked.`

### Touchpoint 7: Self-Contained Worker Brief Guideline 6
- **Target Phrase:** `6. Zero Builds or Tests: NEVER run `go build`, `npm run build`, `go test`, or `pytest`.`
- **Replacement:** `6. Zero Builds, Tests & CI/CD Verification: Every time you write code, you do not build the code. NEVER run `go build`, `npm run build`, `vite build`, `go test`, or `pytest`. Never run all tests locally or full build unless explicitly asked. Only do end-to-end testing based on mentioned test cases if explicitly asked. Do CI/CD verification only if asked; if not asked explicitly, don't do it.`

### Touchpoint 8: Phase 3 Push Gate Verification & Immediate Completion
- **Target Phrase:** Add verification condition in Step 3 and CI/CD directive in Step 4.
- **Replacement:**
  - Step 3: `Before GitMap call, verify: (a) targeted checks exit 0 (>0 files; every time you write code, do not build the code; never run all tests locally or full build unless explicitly asked; only do end-to-end testing based on mentioned test cases if explicitly asked), ...`
  - Step 4: `... Never amend pushed commits. **CI/CD Verification Mandate:** Do CI/CD verification (e.g. gitmap pe, gitmap pe -t, gitmap pipeline-ai, GitHub Actions) ONLY if asked; if not asked explicitly, don't do it. Conclude immediately upon successful GitMap push without waiting for remote CI/CD runs.`

### Touchpoint 9: Banned Operations Checklist Items
- **NO LOCAL TESTS:** `Never run all tests locally or full test suites (06-cicd-local-runner.py, pytest, npm test, go test ./...) unless explicitly asked. Only do end-to-end (E2E) testing based on mentioned test cases if explicitly asked by the user.`
- **NO BUILD CHECKING:** `Every time you write code, you do not build the code. Never run build commands (go build, npm run build, vite build, cargo build, compiler checks) after writing code or to verify compilation. Never run a full build unless explicitly asked.`
- **NO UNREQUESTED CI/CD VERIFICATION:** `Do CI/CD verification only if asked; if not asked explicitly, don't do it (never monitor, wait for, or verify remote pipelines via gitmap pe, gitmap pe -t, or GitHub Actions after push unless explicitly commanded). If CI/CD verification is explicitly requested, never query or loop rapidly...`

### Touchpoint 10: Anti-Hallucination Checklist Item
- **Target Phrase:** `Continuous Loop Maintained: Continuous self-loop executed until 100% complete without running banned test/build commands.`
- **Replacement:** `Continuous Loop Maintained: Continuous self-loop executed until 100% complete without running banned test/build commands (every time code is written, do not build the code; never run all tests locally or full build unless explicitly asked; only do end-to-end testing if explicitly asked; do CI/CD verification only if asked).`

---

## 3. Group 2: The New `08-execute-with-e2e.md` Specification

### 3.1 Metadata & Skill YAML
- **Title:** Execute with End-to-End (E2E) Testing & Automated Minor Release
- **File:** `01-prompts/14-execute/08-execute-with-e2e.md`
- **Mirror:** `01-prompts/v4/14-execute/08-execute-with-e2e.md`
- **Skill Name:** `execute-with-e2e`
- **Description:** Autonomously execute end-to-end verification, run all local test suites to 100% green pass, inspect recent git commits and test cases, verify audit logs, and finalize with an automated minor version bump release ceremony.

### 3.2 Key Differences from Standard V6
1. **Rule R1 Inverted to Mandatory Local Verification:**
   - Instead of banning `go test`, `pytest`, `npm test`, Rule R1 explicitly mandates running all relevant test suites locally to 100% green exit code 0.
2. **Turn 1 Mandatory Commit & Test Case Recitation:**
   - In Turn 1, alongside the 5 Pillars, the agent MUST run and display:
     - Recent git commits: `git log -n 5 --oneline`
     - Recent test cases: list of test files, test fixtures, and E2E scenarios discovered in the repo.
3. **Audit Log Inspection:**
   - Mandatory inspection of `.ai-memory/heatmap.db` and SQLite task engine (`gitmap agent log`, `gitmap heatmap tree`) to guarantee zero unaccounted actions.
4. **Automated Minor Version Bump Release Ceremony:**
   - As soon as all tests are green and audits pass, the prompt commands:
     - Minor version bump: `python 03-ai-scripts/37-bump-version.py --tier minor` or `npm run bump minor`.
     - Changelog update and root readme version pin.
     - Release commit and push via `gitmap cpr` or `gitmap cpf`.
