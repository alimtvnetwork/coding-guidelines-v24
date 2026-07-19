# 31. Compiled Simple Coding Guidelines (AI Blind-Follow)

**Version:** 1.0.0
**Purpose:** One-page checklist any AI can follow without extra context. If a rule here conflicts with a longer spec, the longer spec wins; but this file is enough to write compliant code.

---

## 0. Read These First (in order)

1. This file: `spec/17-consolidated-guidelines/31-compiled-simple-coding-guidelines.md`
2. Canonical size tier (function/file/class caps): `spec/02-coding-guidelines/00-canonical-size-tier.md`
3. Boolean naming: `spec/17-consolidated-guidelines/02-coding-guidelines.md` (Rule 2)
4. Enum standards: `spec/17-consolidated-guidelines/04-enum-standards.md`
5. Error management: `spec/17-consolidated-guidelines/03-error-management.md` + `spec/03-error-manage/`
6. Language-specific: `spec/02-coding-guidelines/{07-csharp,08-go,09-typescript,10-rust,11-python,12-php}/`
7. Mirror for AI editors: `.lovable/coding-guidelines/coding-guidelines.md`

---

## 1. Hard Rules (Zero Tolerance)

1. **Function length**: ≤ 8 lines preferred, ≤ 15 lines hard cap. Skip blanks and comments. Waiver: `// lint-allow: function-length reason="..." max=N`.
2. **No nested `if`**. Flatten with early returns / guard clauses.
3. **Positive, simple `if` conditions**. No `!`, no double negatives. Extract a positively named boolean (`isReady`, `hasAccess`) instead.
4. **Boolean names** start with `is`, `has`, `can`, `should`, `was`, `will`. Never negative (`isNotReady` is forbidden; use `isReady` and invert at the call site).
5. **Narrow types only**. Never `any`, `unknown`, `interface{}`, `object`, `Object`, `dynamic`. Exception: at trust boundaries (catch block, external JSON, third-party lib) narrow immediately with a type guard. `Generic<T>` is the only accepted wide-scope tool.
6. **No swallowed errors**. Every `catch` logs with context AND either rethrows or handles explicitly. Silent `catch {}` is a build-fail.
7. **File size**: ≤ 300 lines any file, ≤ 100 lines any React component (`.tsx`), ≤ 120 lines any class/struct.
8. **No magic strings or numbers**. Use enum, typed constant, or config key. Comparisons must be against a named symbol.
9. **Definitions in dedicated files**. Types, enums, constants, interfaces get their own file, not inline next to first use.
10. **Booleans use `is`/`has` prefix AND positive framing** (reinforces rule 3+4). If you write `if (!isDisabled)`, refactor to `isEnabled`.
11. **DRY is priority #1**. Duplicate logic across two sites means extract it now, not later.
12. **Components stay small and reusable** (React/TS/Vue/etc.). For any feature with 3+ components, produce a Mermaid component diagram first.
13. **Error-manage folder is law**. If `spec/**/error-manage/` exists, every error path follows it exactly.
14. **Immutable-first (Rust-style)**. Assign every variable once at declaration. Never reassign except loop indices. Prefer `const`/`let`/`final`/`val` over `let mut`/`var`. Build result objects with spread/copy, not in-place mutation.
15. **Assets** go to `assets/<NN-folder>/<NN-file>.<ext>` with two-digit sequence prefixes (`assets/01-icons/03-logo.svg`).

---

## 2. Line-Gap and Whitespace Style

1. **One blank line** before every `return`/`throw`, unless it is the only statement in the block.
2. **One blank line** after a closing `}`, unless the next line is another `}`, `else`, `case`, or `catch`.
3. **No double blank lines** anywhere. Ever.
4. **No blank line** immediately after `{` or immediately before `}`.
5. **One blank line** between logical sections inside a function (input validation, core logic, return). But if you need sections, the function is probably too long, refactor first.
6. **One blank line** between top-level declarations (functions, classes, exported constants).
7. **Group imports** with one blank line between groups: stdlib, third-party, first-party absolute, first-party relative. Never mix.
8. **Trailing newline** at end of file. No trailing whitespace on any line.

---

## 3. Error Management (One-Liner Digest)

Full spec: `spec/03-error-manage/` and `spec/17-consolidated-guidelines/03-error-management.md`.

- **Never swallow**: every `catch` logs with operation name + key inputs, then rethrows or returns a typed error.
- **Wrap, do not lose**: `apperror.Wrap(err, "operation", context)` (Go), `throw new AppError(cause, {op, ctx})` (TS). Original stack must survive.
- **Typed errors only**: no `throw "string"`, no `panic("msg")`. Use `AppError`/`apperror.Result[T]` with a registered code.
- **Registered codes**: every user-visible error has an entry in `spec/03-error-manage/03-error-code-registry/`. No ad-hoc codes.
- **Universal response envelope**: backend APIs return `{ data, errors[], meta }`. Frontend parses via a single `parseEnvelope()` helper.
- **Log level matches severity**: `debug` (trace), `info` (lifecycle), `warn` (recoverable), `error` (user-visible failure), `fatal` (process-exit only).
- **Context on every log**: operation name, request/session id, key input values (never secrets, never PII beyond user id).
- **Verify both directions**: before declaring an integration works, curl the backend and inspect the frontend detection logic. Do not trust one side.
- **Retrospective on repeats**: if the same error class hits twice, add a retrospective under `spec/03-error-manage/01-error-resolution/03-retrospectives/`.
- **Frontend**: all captured errors flow through the global error store, then the error modal. No per-component alert boxes.

---

## 4. Data and Schema Rules

1. Tables, types, entities: **PascalCase**.
2. Fields/columns: **camelCase**.
3. JSON keys: **PascalCase** (project convention).
4. Primary key: `int auto-increment`, named `{TableName}Id`. No UUIDs.
5. `Type`/`Status`/`Category`/`Kind` columns: use a 1-N or N-M join table with a registered enum. Never a free-form string column.
6. Every entity/ref table: `Description TEXT NULL`. Every transactional table: `Notes TEXT NULL` + `Comments TEXT NULL`. All nullable, no `DEFAULT`. Join tables exempt.
7. Default DB is SQLite. Prefer ORM. Define joins, PK, FK explicitly.
8. Any DB-touching PR includes a Mermaid ERD.

---

## 5. Language-Specific One-Liners

- **Go**: `apperror.Result[T]` not `(T, error)`. `apperror.New`/`Wrap` not `fmt.Errorf`. `type X byte` + `iota` not string enums.
- **TypeScript**: `Promise.all` for independent async, never sequential `await`. No `any`. `readonly` on interface fields by default.
- **Rust**: prefer `Result<T, E>` with `thiserror`. `let` not `let mut` unless the mutation is the point.
- **PHP**: enum comparison via `->isEqual()`, never `===`.
- **PowerShell**: `Verb-Noun` PascalCase function names, `lowercase-kebab-case` filenames.
- **C#**: `PascalCase` methods and properties, `_camelCase` private fields, `I`-prefix interfaces.

---

## 6. Workflow

1. Read `spec/03-error-manage/` before writing any error path.
2. Read the relevant `spec/NN-*` module before implementing a feature it covers.
3. If you cannot find the spec, ask, do not invent.
4. Plan multi-file features with a Mermaid component/flow diagram first.
5. Run the linters locally: `linters-cicd/` covers file-length, function-length, boolean naming, magic strings, enum standards.
6. Every commit that changes public behavior updates `CHANGELOG.md` + the relevant `spec/**/98-changelog.md`.

---

## 7. Change Protocol for This File

- Edit here first.
- Update the mirror `.lovable/coding-guidelines/coding-guidelines.md` in the same commit.
- Bump the version at the top.
- Do not add rules that contradict the canonical size tier or the error-manage spec, patch those first.
