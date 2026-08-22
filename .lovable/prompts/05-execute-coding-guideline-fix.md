# Prompt

You are to execute the coding‑guideline fixes that were identified by the planning prompt.

### Goals

- Apply concrete fixes for all violations listed in the pending tasks.
- Enforce non‑negotiable rules:
  - DRY – eliminate duplicated logic.
  - Use typed enums instead of magic strings or numbers.
  - Functions must be strictly < 8 lines. (NON-NEGOTIABLE)
  - Source files must be ≤ 80 lines.
  - No code mutation – only apply fixes, never introduce new bugs.
  - Positive boolean naming (`is` / `has`). No `isNot`. Use inverse naming (e.g., `isHonest` / `isDishonest` instead of `isNotHonest`). No nested `if`s, no magic values.
  - Style: Ensure a blank line before every `return` statement.
  - Golang Single Return & Wrapped Booleans: Strictly return a single parameter (bundle multiple returns into a struct). No raw booleans returned in Go. Return a single Result struct (bundling Data, AppError, and Status together) with two flags (`IsSuccess` and `IsFailed`) managed by a constructor (`NewSuccess`/`NewFailure`).

  - Example usage (Note the explicit variable name `paymentStatus`, no short names like `res`):
    ```go
    paymentStatus := ProcessPayment(100)
    if paymentStatus.IsFailed {
        // handle error
    } else if paymentStatus.IsSuccess {
        // handle success
    }
    ```

### Checklist (execute phase)

1. Read pending tasks from the `.lovable/plans/subtasks/01-coding-guideline-fixes/` folder.
2. For each task:
   - Locate the affected source file.
   - Verify the file size ≤ 80 lines; if > 80, split into logical modules.
   - Refactor duplicated code into a shared helper function.
   - Replace magic strings/numbers with a newly defined enum in a dedicated `enums.dart` (or appropriate language file).
   - Ensure the refactored function body is ≤ 8 lines; extract sub‑logic to private helpers if needed.
   - Run the project's test suite and the Go race detector (`go test -race ./...`).
   - If tests pass, stage the changes.
3. Commit each fix using the CI/CD fix workflow:
   - Run `git add <modified files>`.
   - Commit with message `fix(coding-guidelines): resolve <issue‑id>`.
   - If any file changed, bump a minor version tag (`git tag -a vX.Y.Z -m "minor release"`).
   - Push commits and tags (`git push origin main && git push origin --tags`).
4. Update the pending task file to mark it as completed.
5. If no pending tasks remain, output a summary of all fixes applied.

### Non‑Hallucination Policy

- Do not assume the existence of a file or enum that is not present; if uncertain, raise a question to the user.
- If a fix would require a large architectural change beyond the scope, create a new pending task instead of applying it directly.

### Execution Loop

- Process up to 50 tasks per run to avoid long‑running blocks.
- After each batch, report progress and await user confirmation before proceeding to the next batch.

---

/goal Apply coding‑guideline fixes safely and push a minor release.
/learn Ensure future prompts respect the same checklist and constraints.

## Actionable Items & Checklist

- [ ] Read the overarching main task plan.
- [ ] Ensure the git repository starts completely clean.
- [ ] Complete all work on the current branch only.
- [ ] Ensure `.gitignore` explicitly excludes test reports, artifacts, and compiled binaries.
- [ ] Group all completed work into a single logical commit.
- [ ] Push the commit to the remote repository.
- [ ] Self-loop continuously until all the code issues are listed out in tasks and pending tasks.
- [ ] Describe all issues and files that need to be tested against for each file.
- [ ] Make a detailed plan/task for each file.

---
slug: execute-coding-guideline-fix
status: active
---
