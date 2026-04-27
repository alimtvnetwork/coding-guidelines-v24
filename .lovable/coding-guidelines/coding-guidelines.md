# Coding Guidelines

These guidelines apply to all code generated or modified in this project. Read this file fully before writing any code. Also read all prompt files under `.lovable/prompts/*.*` and the index `.lovable/prompts.md`.

## Function & Control Flow

1. Keep functions under 8 lines.
   *Enforcement:* **CODE-RED-005** (build-failing `error`). Any function
   body whose effective line count exceeds 8 fails CI. CODE-RED-004
   remains as a redundant >15-line safety net at the same severity.
   See [`spec/02-coding-guidelines/06-cicd-integration/06-rules-mapping.md`](../../spec/02-coding-guidelines/06-cicd-integration/06-rules-mapping.md)
   for the coordinated tier table.

2. No nested `if` statements.

3. Keep `if` conditions simple — avoid negative conditions; prefer positive, simple terms.

4. Follow the Boolean guidelines below.

## Boolean Guidelines

1. Boolean variables and functions must be prefixed with `is` or `has` (e.g., `isActive`, `hasExpired`).

2. Do not use negative boolean names (e.g., avoid `isNotReady`; use `isReady` instead).

3. Do not use negative conditions in `if` statements when a positive equivalent exists.

## Types

1. Use proper, specific types — never use `any`, `unknown`, `interface{}`, or any wide-range type.

2. Generics (`Generic<T>`) are the only allowed open type form.

3. Define types and interfaces in separate dedicated files, not inline at usage site.

## Enums & Constants

1. No magic strings or numbers — use Enums or Constants.

2. `Type`, `Status`, `Category`, `Kind`, and similar classification fields must be Enums.

3. Define enums and constants in separate files.

## Error Handling

1. Never swallow errors.

2. Every `catch` block must log the error properly per the language-specific logging guidelines.

3. Follow language-specific error management rules.

## File & Class Size

1. No file or class may exceed 80–100 lines maximum.

2. Split larger files into smaller, focused modules.

## Definitions

1. Do not define types, enums, constants, or schemas inline.

2. Place each definition in its own dedicated file, separated by concern.

## Required Reads Before Coding

1. `.lovable/coding-guidelines.md` (this file)

2. `.lovable/prompts.md` (prompt index)

3. All files under `.lovable/prompts/*.*`

4. The relevant spec folder for the current task (e.g., `/spec/YY-app/`, `/spec/xx-app-issues/`)

5. Language-specific guidelines, boolean guidelines, enum guidelines, and error management guidelines referenced in this file.
