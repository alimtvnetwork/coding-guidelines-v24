# Plan: AppError Human and Logger Methods

**Status:** COMPLETED

## Acceptance Criteria

1. Update 02-spec/03-error-manage to define a distinct section for "Human and Logger Output Methods".
2. Add a method HumanString() (or similar) that retrieves the human-friendly version of the error (i.e. the DisplayError or a safe fallback).
3. Add a method LogFields() or LogMap() that serializes the AppError into a structured map suitable for JSON loggers (Zap, Logrus).
4. Clearly separate these from the developer console string (ConsoleString() or FullString()).
5. Ensure these methods "can be accessed from the logger or how we want to log" as explicitly requested.
6. Provide actionable steps for the AI to implement these new display methods.
7. Perform the release cycle (v6.28.0).

## Subtasks

- 01-update-display-specs.md: Rewrite the 2.3 Display Methods section in 02-apperror-struct.md and append to the AI action plan.
- 02-release-6.28.0.md: Release script execution and readme sync.
