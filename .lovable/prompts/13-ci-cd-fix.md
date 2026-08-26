# Instruction (must follow): CI/CD Fix Loop (with RCA & End Aliasing)

Trigger Keywords & Aliases: `fix with RCA`, `fix`, `fix, fix`, `CI/CD fix`, `cicd fix`

/goal The user will provide an error log from a failed CI/CD pipeline or build run. Your objective is to ingest recent Root Cause Analysis (RCA) records from `.lovable/cicd-issues/`, `.lovable/issues/`, and `spec/03-error-manage/` before touching code, perform a surgical RCA on the failed build/test, fix the codebase to resolve the issue, record the memory of this failure in `.lovable/cicd-issues/01-<slug>.md` and `.lovable/strictly-avoid.md`, commit using the commit-fix workflow, and repeat the loop until the CI/CD pipeline is 100% green.

/learn Read past Root Cause Analyses (RCAs) from `.lovable/cicd-issues/` and `.lovable/issues/` so that past mistakes, regression patterns, and build traps are never repeated.

## Rules & Constraints (Non-Negotiable)

1. Analyze First & Read Past RCAs: Do not blindly change code. First read recent RCAs in `.lovable/cicd-issues/` and `.lovable/strictly-avoid.md`. Then trace the provided CI/CD error to the exact file, line, and dependency, performing a complete Root Cause Analysis (RCA).
2. Update Memory & Avoid List: The RCA and solution must be permanently recorded. Write the details to `.lovable/cicd-issues/01-<slug>.md` (sequenced as `01-`, `02-`, etc.) and register it in `.lovable/cicd-issues/index.md` (or `.lovable/cicd-index.md`). If a new forbidden pattern is identified, append it to `.lovable/strictly-avoid.md`.
3. Commit the Fix: Once the code is fixed, invoke the standard commit-fix procedure. Group changes logically with a clean, descriptive commit message (`fix(ci): <description>`).
4. Iterative Looping: If the pipeline fails again after your fix, the user will provide the new error. You must repeat this exact process—RCA, memory update, code fix, verification, commit, push—until the CI/CD run succeeds.
5. No Blind Overwrites: When updating memory, never delete or truncate existing history. Always append.
6. Anti-Hallucination Contract: If the cause is ambiguous or missing from logs, stop and ask clarifying questions instead of guessing.

## Actionable Items & Checklist

### 1. Pre-Flight & Past RCA Ingestion

- [ ] /learn past failure patterns in `.lovable/cicd-issues/` and `.lovable/strictly-avoid.md`.
- [ ] Read the provided CI/CD error log carefully.
- [ ] Identify the exact file, line, function, and dependency causing the failure.
- [ ] Formulate a one-sentence Root Cause Analysis followed by the full causal chain.

### 2. Memory Update (Mandatory)

- [ ] Create a new issue file at `.lovable/cicd-issues/01-<slug>.md`.
- [ ] Document: Error Summary, Root Cause Analysis, Solution Applied, and "What NOT to Repeat".
- [ ] Update `.lovable/cicd-issues/index.md` in the same operation.
- [ ] If a hard rule was broken, append a one-line prohibition to `.lovable/strictly-avoid.md`.

### 3. Execution & Code Fix

- [ ] Implement the minimal correct fix in the codebase based strictly on the RCA.
- [ ] Ensure the fix adheres to `spec/02-coding-guidelines/` and `spec/03-error-manage/`.
- [ ] Run local builds, linters, or unit tests if available to verify the fix before committing.

### 4. Verification & Final Checks

- [ ] Run all project tests (e.g., `go test ./...`, `npm test`, or equivalent test suites).
- [ ] If the project uses Go, run the race detector (`go test -race ./...`).
- [ ] Verify zero regressions, no swallowed errors, and no negative boolean anti-patterns.

### 5. Commit, Minor Release & Push

- [ ] Stage changes and commit using the commit-fix workflow (`fix(ci): <issue-slug>`).
- [ ] Tag a minor release or bump version if required by project release rules.
- [ ] Push the commit and tags to the remote repository.
- [ ] Await the next CI/CD result. If green, mark as completed; if failed, loop back to Step 1.

## Awaiting Input

Wait for the user to trigger with any alias (`fix with RCA`, `fix`, `fix, fix`, `CI/CD fix`) or paste the CI/CD error log. Once provided, immediately begin at Step 1.

---

## Metadata

- slug: cicd-fix
- status: active


## End of Tunnel Release & Version Bump (Mandatory)

When EVERYTHING is completely finished and fixed (at the very end of the tunnel), you MUST trigger a release.
- You must bump the MINOR version.
- You must focus on the `version.json` file as the source of truth for the release.
- If you do not know how to cut a release for this specific repository, or if `version.json` is missing/unclear, you must either search the repository for release instructions or explicitly ask the user for help. Do not guess.
- You MUST strictly exclude all test files (e.g., `*test*`, `*.spec.*`) from version scanning and modification, as they contain mock data.
- You must create and maintain `.lovable/memory/release-architecture-map.md` documenting exactly how releases work in the repository. Ensure it is enqueued in `what-to-read.md` and linked in the root `readme.md`.

