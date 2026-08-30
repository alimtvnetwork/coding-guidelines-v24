# RCA & Boolean Fix Plan

**Status:** PENDING

## Root Cause Analysis (RCA)

**Issue 1: AI writing if hasMatchOrTitle == true { instead of if hasMatchOrTitle {**
* **Root Cause:** The guidelines heavily penalize implicit negation (!variable). To fix this, earlier prompts/resolutions instructed AIs to use explicit === false (e.g., in .lovable/solved-issues/01-readme-boolean-negatives.md). The AI model generalized this "explicit boolean evaluation" heuristic to positive checks as well, assuming that if implicit negative (!) is bad, implicit positive (if variable) might also be ambiguous, leading to the hallucinated and redundant == true pattern.
* **Resolution:** We must explicitly BAN == true and === true across all languages. Positive booleans MUST ALWAYS be evaluated implicitly (if isValid {, not if isValid == true {).

**Issue 2: gitmap && go generate drift CI Failure**
* **Root Cause:** The AI modified Go constants (likely enum definitions) but failed to run the corresponding go generate ./... command. When the code was pushed, the CI pipeline detected that the generated files (stringer, etc.) were out of sync with the constants, causing the pipeline to fail with: Generated files are out of sync with constants. Run 'cd gitmap && go generate ./...' locally and commit the result.
* **Resolution:** Update the CI/CD and Prompt guidelines to explicitly instruct the AI: "If you modify Go constants or enums, you MUST run go generate ./... in the relevant directory (e.g., gitmap) and commit the resulting generated files to prevent CI drift."

## Subtasks
- **Subtask 1:** Update .lovable/strictly-avoid.md with a TOTAL BAN on == true / === true.
- **Subtask 2:** Update spec/02-coding-guidelines/01-cross-language/02-boolean-principles/01-index.md to formally document this.
- **Subtask 3:** Update spec/02-coding-guidelines/03-golang/02-boolean-standards.md to cover this and the go generate requirement.
- **Subtask 4:** Update .lovable/prompts/ (e.g., 15-commit-fix.md and 10-execute-batched-loop.md) to explicitly mention go generate after constant changes and ban == true.
