# CI/CD Fix Loop

- slug: cicd-fix
- status: active

## Prompt

# CI/CD Fix Loop

## Goal

The user will provide an error log from a failed CI/CD pipeline. Your objective is to perform a Root Cause Analysis (RCA), fix the codebase to resolve the issue, record the memory of this failure and its solution in the `.lovable` folder, commit the changes using the commit-fix workflow, and repeat the loop until the CI/CD pipeline is fully green.

## Rules & Constraints

1. Analyze First: Do not blindly change code. Read the provided CI/CD error, trace it back to the exact file and line, and perform a proper Root Cause Analysis (RCA).
2. Update Memory: The RCA and the solution must be permanently recorded. Write the details to `.lovable/cicd-issues/XX-<slug>.md` and update `.lovable/cicd-index.md` so that future AI sessions do not repeat the same mistake.
3. Commit the Fix: Once the code is fixed, you must invoke the standard commit-fix procedure. Group the changes logically and use a clear commit message.
4. Iterative Looping: If the pipeline fails again after your fix, the user will provide the new error. You must repeat this exact process—RCA, memory update, code fix, commit—until the CI/CD run succeeds.
5. No Blind Overwrites: When updating memory, never delete existing history. Append your new findings.

## Actionable Items & Checklist

### 1. Root Cause Analysis

- [ ] Read the provided CI/CD error log carefully.
- [ ] Identify the exact file, line, and dependency causing the failure.
- [ ] Formulate a clear Root Cause Analysis (why did it fail, and what is the proper fix?).

### 2. Memory Update (Mandatory)

- [ ] Create a new file for this specific issue at `.lovable/cicd-issues/XX-<slug>.md`.
- [ ] Document the Error, the Root Cause, the Solution, and "What NOT to Repeat" in that file.
- [ ] Update `.lovable/cicd-index.md` in the same operation to link to the new issue file.

### 3. Execution & Code Fix

- [ ] Implement the fix in the codebase based on the RCA.
- [ ] Ensure the fix adheres to the project's coding guidelines and error management specs.
- [ ] Run local builds or unit tests if available to verify the fix before committing.

### 4. Verification & Final Checks

- [ ] Run all project tests (e.g., `go test ./...`).
- [ ] If the project uses Go, run the race detector (`go test -race ./...`).
- [ ] Ensure no new test failures or race conditions are introduced.

### 5. Commit, Minor Release & Push

- [ ] Stage the changes and commit them using the commit-fix workflow with a descriptive commit message.
- [ ] Tag a minor release (e.g., bump patch version) and push the tag.
- [ ] Push the commit to the remote repository.
- [ ] Wait for the user to provide the next CI/CD result. If it fails again, repeat the entire checklist from Step 1. If it passes, mark the task as complete.

## Awaiting Input

Wait for the user to paste the CI/CD error log. Once provided, immediately begin at Step 1 of the checklist.
