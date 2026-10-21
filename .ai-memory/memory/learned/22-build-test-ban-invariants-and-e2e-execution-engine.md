# Learned Protocol: In-Flight Build/Test Ban Invariants & Dedicated E2E Execution Engine

> **File:** `.ai-memory/memory/learned/22-build-test-ban-invariants-and-e2e-execution-engine.md`  
> **Date:** 2026-10-21  
> **Status:** ACTIVE  
> **Classification:** Core Execution Architecture & Linter Standard  

---

## 1. Context & Motivation

During autonomous execution turns, engineering agents frequently experienced token budget exhaustion, latency, and pipeline lockups caused by:
1. **Unrequested In-Flight Builds:** Attempting full project compilation (`go build`, `npm run build`, `vite build`, `cargo build`) every time code or markdown was edited.
2. **Unrequested Full Test Suites:** Running entire test suites (`go test ./...`, `pytest`, `npm test`) during routine turns when only single files or documentation items were modified.
3. **Unrequested CI/CD Pipeline Monitoring:** Committing and pushing changes, then remaining stuck polling or sleeping on remote GitHub Actions pipelines (`gitmap pe -t`, `gh run view`) when the user never asked for CI/CD verification.

---

## 2. The Four Non-Negotiable Invariants

### 2.1 Invariant 1: In-Flight Build Ban
> **"Every time you write code, you don't build the code."**
- Never execute compiler checks or full builds to verify compilation in routine turns.
- Rely on file-scoped targeted linters (`check-relative-paths.py`, `check-forbidden-strings.py`, `guideline-autofixer.py --check-only`).

### 2.2 Invariant 2: Local Test Suite Ban
> **"Never run all tests locally or full build unless explicitly asked."**
- Complete local test suites are strictly banned in routine turns.
- Only explicit user commands lift this prohibition.

### 2.3 Invariant 3: Scope-Restricted E2E Testing
> **"Only do end-to-end testing based on mentioned test cases if explicitly asked."**
- If the user explicitly mentions a specific test case, run only that test case.
- Never extrapolate an unprompted E2E regression flow.

### 2.4 Invariant 4: Opt-In CI/CD Verification
> **"Do CI/CD verification only if asked; if not asked explicitly, don't do it."**
- Routine tasks conclude immediately upon successful upstream push (`gitmap cpf` / `gitmap cpb`).
- Never sleep, poll, or wait for remote pipelines unless the user explicitly requested CI/CD verification.

---

## 3. Dedicated E2E Execution Engine (`08-execute-with-e2e.md`)

When complete test verification and automated release finalization are required, agents operate under `01-prompts/14-execute/08-execute-with-e2e.md` (and companion skill `execute-with-e2e`):
1. **Turn 1 Recitation:** Displays the last 5 commits (`git log -n 5 --oneline`) and inventories discovered test cases alongside the 5 Pillars and confirmed task breakdown.
2. **Mandatory 100% Green Local Test Suite Execution:** Explicitly commands running all local test suites (`go test ./...`, `npm test`, `pytest`) to exit code 0.
3. **Audit Log Inspection:** Verifies `.ai-memory/heatmap.db` and agent task logs (`gitmap agent log`, `gitmap heatmap tree`).
4. **Automated Minor Release Ceremony:** Commands minor version bump (`python 03-ai-scripts/37-bump-version.py --tier minor` or `npm run bump minor`), changelog entry, root readme version pin, and atomic GitMap release push.

---

## 4. Multi-Repository Fleet Synchronization Invariant

All canonical prompt additions, skill creations, and coding guideline updates are synchronized across all 43 connected repositories via `03-ai-scripts/38-sync-prompts-skills-scripts.py`. The 5 Non-Negotiable Boundaries must always be strictly observed:
1. `02-spec/21-*` through `02-spec/25-*` are private to target repos and NEVER touched.
2. AI scripts (`03-ai-scripts/`) are strictly additive; existing scripts modified in target repos are NEVER overwritten.
3. Version bump scripts are immutable per repository.
4. `.ai-memory/memory/` and `.ai-memory/plans/` in target repositories are completely isolated.
5. Zero secrets or credentials leaked across repositories.
