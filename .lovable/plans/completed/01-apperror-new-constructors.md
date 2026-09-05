# Plan: Go AppError Namespace Constructors

**Status:** COMPLETED

## Acceptance Criteria

1. Provide detailed API definition in the 02-spec/03-error-manage architecture specs for the new Apperror.New.* namespaced creator methods.
2. Define Apperror.New.Error(errortype, error): If error == nil, returns
il. Otherwise wraps the error with the enum variant.
3. Define Apperror.New.UsingErrorMsg(errortype, error, msg): If error == nil, returns
il. Otherwise wraps the error with custom message.
4. Define Apperror.New.UsingMsg(errortype, msg): Creates a new error using the provided message and enum variant.
5. Define Apperror.New.ErrorVar(errortype, error, varname, varvalue): Wraps an error and injects a single key-value into Values (if error == nil, returns
il).
6. Define Apperror.New.ErrorVars(errortype, error, vars): Wraps an error and injects a map of key-values into Values (if error == nil, returns
il).
7. Must include actionable instructions for the AI on how to implement this pattern so that "if no error, nothing should be created" (early return
il).
8. Update
eadme.md, .lovable/what-to-read.md properly.
9. Verify all coding guidelines and formatting.

## Architecture

The Apperror package will expose a global New variable (or struct namespace) providing these constructors, replacing the global package-level functions. The key benefit is preventing bloated if err != nil blocks when the error is being passed directly.

## Subtasks

- 01-update-spec.md: Update 02-spec/03-error-manage/02-error-architecture/06-apperror-package/01-apperror-reference/02-apperror-struct.md and 09-ai-action-plan-apperror-methods.md.
- 02-release.md: Run sync and release v6.27.0.
