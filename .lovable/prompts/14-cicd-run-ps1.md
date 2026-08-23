# Instruction (must follow): CI/CD Pipeline & Run Script Creation

You are responsible for creating or fixing CI/CD pipelines (e.g., GitHub Actions workflows) and dynamic execution scripts (e.g., `run.ps1`, `run.sh`, `run.config.json`). 

## Generic CI/CD & Automation Guidelines Checklist

This document serves as a strict, universal checklist and specification for setting up and fixing CI/CD pipelines, as well as the standard run scripts (e.g., `run.ps1`/`run.sh`), for any project. Any AI agent operating on DevOps or CI/CD tasks MUST read and follow these guidelines to ensure consistency across different tech stacks and repositories.

### 1. Specification Files to Read & Maintain

Before making any changes to `.github/workflows` or automation scripts, you must read the following architecture documents. These contain the foundational constraints and mechanisms for deployment, automation, and CI/CD pipelines.

#### PowerShell & Orchestration (`spec/11-powershell-integration`)

- [ ] `spec/11-powershell-integration/00-overview.md`
- [ ] `spec/11-powershell-integration/02-script-reference.md`
- [ ] `spec/11-powershell-integration/03-integration-guide.md`

#### CI/CD Pipeline Workflows (`spec/12-cicd-pipeline-workflows`)

- [ ] `spec/12-cicd-pipeline-workflows/00-overview.md`
- [ ] `spec/12-cicd-pipeline-workflows/01-ci-pipeline.md`
- [ ] `spec/12-cicd-pipeline-workflows/02-release-pipeline.md`
- [ ] `spec/12-cicd-pipeline-workflows/04-install-script-generation.md`
- [ ] `spec/12-cicd-pipeline-workflows/05-changelog-integration.md`

#### CLI & Build (`spec/13-generic-cli`)

- [ ] `spec/13-generic-cli/00-overview.md`
- [ ] `spec/13-generic-cli/11-build-deploy.md`
- [ ] `spec/13-generic-cli/18-batch-execution.md`

#### Update Mechanisms (`spec/14-update`)

- [ ] `spec/14-update/04-build-scripts.md`
- [ ] `spec/14-update/17-release-pipeline.md`
- [ ] `spec/14-update/18-install-scripts.md`

#### Release Engineering (`spec/16-generic-release`)

- [ ] `spec/16-generic-release/00-overview.md`
- [ ] `spec/16-generic-release/02-release-pipeline.md`
- [ ] `spec/16-generic-release/03-install-scripts.md`

#### Context / Issue Logging

- [ ] `cicd-issues/<issue-name>.md`
  - Any time a CI/CD pipeline fails, an AI must log the failure here (including error traces and environment context) before attempting a fix.

---

### 2. CI/CD Implementation Checklist

When building or fixing the CI/CD pipelines (e.g., `.github/workflows/ci.yml`), enforce the following steps universally:

- [ ] Dependency Caching: The pipeline MUST cache dependencies based on the project's lockfiles (e.g., `package-lock.json`, `bun.lockb`, `poetry.lock`, `go.sum`). This minimizes build times.
- [ ] Toolchain Matching: Ensure the CI runner uses the exact tool versions specified in the project's configuration (e.g., `.nvmrc`, `.python-version`, or the generic `run.config.json`).
- [ ] Linting & Formatting: Run the project's defined linting and formatting commands first. The pipeline must fail immediately if there are style violations, preventing bad code from proceeding to tests.
- [ ] Type Checking / Static Analysis: If the language supports it (e.g., TypeScript, Python with mypy, Go), run static analysis as a parallel job before or alongside testing.
- [ ] Test Execution: Execute the project's testing suites (unit, integration, e2e) as defined in the configuration file.
- [ ] Artifact Generation: (Optional) If it is a release branch, the pipeline should compile/zip the application using rules defined in the run script and attach it as a release artifact.

---

### 3. Dynamic Script Architecture (e.g., `run.ps1` & `run.config.json`)

To prevent hardcoded commands and ports, the execution architecture MUST be dynamic. The local run script (e.g., `run.ps1` or `run.sh`) must act as a generic orchestrator, while a JSON configuration file serves as the single source of truth for both local development and CI execution.

#### Expected Configuration Structure (Reference: `run.config.json`)

The JSON file should define all services, ports, and lifecycle commands.

```json
{
  "projectName": "Generic Project Name",
  "frontend": {
    "dir": "path/to/frontend",
    "port": 3000
  },
  "backend": {
    "dir": "path/to/backend",
    "port": 8080
  },
  "host": "127.0.0.1",
  "commands": {
    "install": "package-manager install",
    "dev:frontend": "command to start frontend dev server --port {fePort}",
    "dev:backend": "command to start backend server --port {bePort}",
    "build:frontend": "command to build frontend",
    "test:frontend": "command to run frontend tests",
    "test:backend": "command to run backend tests",
    "lint": "command to run linter"
  }
}
```

#### Expected Run Script Implementation Rules (Reference: `run.ps1` / `run.sh`)

When writing or modifying the orchestration scripts, any AI must adhere to the following logic:

1. Parse Configuration First:
   The script must read the configuration file (e.g., `run.config.json`) into memory and extract variables (ports, directories, commands).
2. Dynamic Command Injection:
   Never hardcode commands (e.g., `npm run dev`). Instead, read the command string from the JSON and dynamically substitute any necessary placeholders (like `{fePort}`).
3. Graceful Error Handling & Waiting:
   When orchestrating multiple services, the script must wait for dependencies to become healthy before proceeding. Use generic HTTP or TCP polling mechanisms to ensure a backend is fully online before a frontend attempts to connect to it.
4. CI Mode Toggle:
   The script must accept a `-CI` (or `--ci`) flag. When running in CI mode:
   - Skip launching local browsers or interactive shells.
   - Run the build, lint, and test commands from the JSON instead of the dev server commands.
   - Exit with code `1` immediately if any step fails.
5. Process Cleanup (Trap/Finally):
   Ensure all spawned processes are captured in a job/PID array. The script must aggressively kill these processes in the `finally {}` or `trap` block on exit. There should be no zombie processes left blocking ports.

---

## 4. Commit Fix Actionable Items & Checklist (Non-negotiable)

After completing the pipeline and run script creation, you MUST follow this checklist for committing and verifying the changes.

### 4.1. Pre-flight & Planning

- [ ] Ensure the git repository starts completely clean. If dirty, commit, stash, or fix git issues before writing any new code.
- [ ] Read the overarching main task plan from `.lovable/plans/pending/XX-<slug>.md` to understand what needs to be executed.
- [ ] Derive the `<slug>` from the plan filename itself (e.g., plan file `03-auth-refactor.md` → slug is `03-auth-refactor`). Never invent a slug.
- [ ] Confirm subtask files exist under `.lovable/plans/subtasks/XX-<slug>/SS-<subslug>.md` for each step that needs parallel execution. Create them if missing, following the plan prompt structure.
- [ ] Ensure the plan is highly extensive, explicitly detailing where and how to make changes so sub-agents can easily execute tasks (Non-negotiable).
- [ ] Direct subtasks to `.lovable/plans/subtasks/01-<slug>/01-<subslug>.md` and update plans in `.lovable/plans/pending/01-<slug>.md`. Do not write randomly into `.lovable`.
- [ ] Read the memory files and the spec folder coding guidelines + error manage guidelines before touching code.
- [ ] Anti-Hallucination: If referenced files are missing or ambiguous, stop and ask clarifying questions.

### 4.2. Ruthless Management & Subtask Looping

- [ ] Map out the subtasks from the big plan and spawn sub-agents for all independent tasks simultaneously (MAXIMUM 2 sub-agents concurrently to avoid RAM and caching issues).
- [ ] Each sub-agent may only run a MAXIMUM of 2-3 async operations at a time.
- [ ] Enforce lifecycle: sub-agent reads subtask file → marks `🔄 In Progress` → works → marks `✅ Done` with file list and summary → updates parent plan step → signals completion.
- [ ] Track queue state by counting total subtasks spawned vs. total `✅ Done` entries in the subtask files. Proceed to commit only when the counts match.
- [ ] If a sub-agent fails to update its status file or gives garbage, kill it immediately and restart it.

### 4.3. Root Cause

- [ ] Find the root cause of the problem first, before applying any fix.
- [ ] Record the root cause strictly into the `.lovable` memory structure per the write protocols.

### 4.4. File System Writes & Main Agent Commit

- [ ] Sub-agents write to the file system and update their task entries. They do NOT commit.
- [ ] Wait until all sub-agents have signaled completion and updated their subtask files under `.lovable/plans/subtasks/`.
- [ ] Ensure `.gitignore` explicitly excludes test reports, test data, artifacts, and compiled binaries (Non-negotiable).
- [ ] RED FLAG: Verify absolutely NO test results or binaries are staged before making the commit.
- [ ] Group all completed work into a single logical commit.
- [ ] If issues arise during the commit, fix them immediately and retry.
- [ ] Push the commit to the remote repository. Pushing is non-negotiable.

### 4.5. Code Standards (non-negotiable)

- [ ] Follow the code review guidelines from the aspect folder.
- [ ] Ensure every try-catch block explicitly logs the error according to the error manage folder.
- [ ] Create a query wrapper for PHP/Python/TS that handles automatic failure logging, so logging is not scattered.
- [ ] Use explicit `isFail` properties; NEVER use inverted success checks (use `response.isFail`, not `!response.isSuccess`).
- [ ] Remove all magic strings and magic numbers unless used directly for logging — and state that logger exception in the typing.
- [ ] Replace TypeScript string union types (e.g. `"pass" | "fail" | "fallback"`) with Enums.
- [ ] Ensure every Enum name ends with the `Type` suffix (e.g. `StatusType`, never `Status` or `Status7`).
- [ ] Ensure all Enum values are written in PascalCase (e.g., `enum StatusType { ActiveState = "ACTIVE" }`), avoiding `_camelCase` or `camelCase`, unless the specific language (like Rust) conventionally dictates otherwise.
- [ ] Reuse constants — never duplicate them. Code must always be DRY; never repeat code. This is high priority.

### 4.6. End-of-Loop Final Verification (Once only, at the very end)

- [ ] Check the full build. Fix every build failure, commit, and push.
- [ ] Run all unit tests. Fix every failing test, commit, and push.
- [ ] Check CI/CD status and ensure pipelines pass.
- [ ] Audit that coding guidelines from the aspect folder and error manage folder have been followed across all changed files.
- [ ] Finish the job only when everything is green, pushed, and fully verified.

## Actionable Items & Checklist

- [ ] Read the overarching main task plan.
- [ ] Ensure the git repository starts completely clean.
- [ ] Complete all work on the current branch only.
- [ ] Ensure `.gitignore` explicitly excludes test reports, artifacts, and compiled binaries.
- [ ] Group all completed work into a single logical commit.
- [ ] Push the commit to the remote repository.

---

## Metadata

- slug: cicd-run-ps1
- status: active
