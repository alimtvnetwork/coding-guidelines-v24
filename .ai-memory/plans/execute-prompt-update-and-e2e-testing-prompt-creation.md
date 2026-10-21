# Master Execution Plan: Execute Prompt Update and E2E Testing Prompt Creation

## User Request (Verbatim)
```text
In the coding guideline, in the execute v6 prompt, make sure to not run all the tests locally or full build unless it is asked. Only do end-to-end testing based on the test case that is mentioned, if it is asked. If not asked, then do not do end-to-end test or CI/CD verification. Do the CI/CD verification only if asked. If it is not asked explicitly, don't do it. Same will repeat for end-to-end test. Also the build, do not build unless it is asked to do the build testing or end-to-end testing. Every time you write code, you don't build the code. This is how the instruction needs to be updated in the execute prompt in v6, and skills. Also, create another prompt close to the execute prompt v6 that will be execute with end-to-end testing. It will run all the tests, all the test suites locally, and also the test case that is mentioned recently. It will list out the recent commits and test cases and execute all of these locally. Make sure that everything is green, the test case is confirmed, it is following every one of the steps, and also check the audit log. That would be execute with E2E prompt name. Should be kept in the execute folder. Make sure that it also mentions to release with a minor bump as soon as everything is green or everything is checked and passed.
```

---

## 1. Executive Summary & Strategy

This mission updates the core autonomous execution protocols of the repository to eliminate compute waste, in-flight build loops, and unrequested CI/CD waiting, while providing a dedicated end-to-end verification and release workflow:
1. **Surgical Invariant Injection in V6:** Update `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` (and its V4 mirror and twin skills) across 10 exact touchpoints to ban unrequested builds, full local test suites, unrequested E2E flows, and post-push CI/CD monitoring.
2. **Author `08-execute-with-e2e.md`:** Create the new dedicated prompt and twin skills (`execute-with-e2e`) requiring Turn 1 commit recitation (`git log -n 5`), recent test case enumeration, full local test suite execution, audit log inspection, and automated minor version bump release.
3. **Register & Distribute:** Update prompt indexes (`01-prompts/readme.md`, `14-execute/readme.md`, `.ai-memory/prompts.md`) and deploy across all 43 connected repositories via `03-ai-scripts/38-sync-prompts-skills-scripts.py`.

---

## 2. Specifications Cross-Reference
- Architecture Spec: `02-spec/21-app/execute-prompt-update-and-e2e-testing-prompt-creation/01-architecture-spec.md`
- Component & Prompt Spec: `02-spec/21-app/execute-prompt-update-and-e2e-testing-prompt-creation/02-component-and-prompt-spec.md`
- Fleet Sync Spec: `02-spec/21-app/execute-prompt-update-and-e2e-testing-prompt-creation/03-fleet-sync-spec.md`

---

## 3. Subtask Decomposition (5 Atomic Phases)

| Subtask ID | Title | Disjoint Owned Files | Status |
| :--- | :--- | :--- | :--- |
| `Subtask-01` | Author Architecture, Component & Fleet Sync Specs and Master Plan | `02-spec/21-app/execute-prompt-update-and-e2e-testing-prompt-creation/*`, `.ai-memory/plans/*` | `[IN PROGRESS]` |
| `Subtask-02` | Update V6 Prompts and Skills to Ban Unrequested Builds and Test Suites | `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`, `01-prompts/v4/...`, `.agents/skills/...`, `.cursor/skills/...` | `[PENDING]` |
| `Subtask-03` | Create `08-execute-with-e2e.md` Prompts with Mandatory E2E and Release Minor Bump | `01-prompts/14-execute/08-execute-with-e2e.md`, `01-prompts/v4/14-execute/08-execute-with-e2e.md` | `[PENDING]` |
| `Subtask-04` | Create `execute-with-e2e` Twin Skills and Update Prompt Registries | `.agents/skills/execute-with-e2e/skill.md`, `.cursor/skills/execute-with-e2e/skill.md`, `01-prompts/readme.md`, `14-execute/readme.md`, `.ai-memory/prompts.md` | `[PENDING]` |
| `Subtask-05` | Validate Linters, Synchronize across 43 Fleet Repositories, and Atomic GitMap Push | Fleet repositories, atomic push via `gitmap cpf` | `[PENDING]` |

---

## 4. Verification & Completion Criteria
- Zero relative path violations (`python linter-scripts/check-relative-paths.py` exits 0).
- Zero forbidden strings or credentials.
- 100% heatmap action closure (`gitmap heatmap status` reports 100% met).
- Atomic commit & push formatted with hyphen (`gitmap cpf "<module> - <summary>"`).
