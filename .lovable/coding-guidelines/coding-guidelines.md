# Coding Guidelines (AI Blind-Follow Mirror)

Mirror of `spec/17-consolidated-guidelines/31-compiled-simple-coding-guidelines.md`. If the two disagree, the spec file wins and this mirror must be patched in the same commit.

The AI MUST read this file AND the referenced sources before writing or modifying code.

## Files To Read Before Coding

1. This file: `.lovable/coding-guidelines/coding-guidelines.md`
2. Compiled spec: `spec/17-consolidated-guidelines/31-compiled-simple-coding-guidelines.md`
3. Canonical size tier: `spec/02-coding-guidelines/00-canonical-size-tier.md`
4. Full coding guidelines: `spec/17-consolidated-guidelines/02-coding-guidelines.md`
5. Error management: `spec/17-consolidated-guidelines/03-error-management.md` + `spec/03-error-manage/`
6. Enum standards: `spec/17-consolidated-guidelines/04-enum-standards.md`
7. Language-specific: `spec/02-coding-guidelines/{07-csharp,08-go,09-typescript,10-rust,11-python,12-php}/`
8. App specs: `spec/01-app/*.md`
9. Memory protocol: `.lovable/memory/protocol/01-workflow-rules.md`

## Hard Rules

1. **Function length**: ≤ 8 lines preferred, ≤ 15 lines hard cap (skip blanks and comments). Waiver: `// lint-allow: function-length reason="..." max=N`.
2. **No nested `if`**. Flatten with early returns.
3. **Positive, simple `if`**. No `!`, no double negatives. Extract a positively named boolean.
4. **Boolean names**: `is`/`has`/`can`/`should`/`was`/`will` prefix, positive framing only.
5. **Narrow types**. No `any`, `unknown`, `interface{}`, `object`, `dynamic`. Exception: trust boundaries (catch, external JSON, third-party) narrow immediately with a type guard. `Generic<T>` is the only wide-scope tool.
6. **No swallowed errors**. Every `catch` logs with context, then rethrows or handles.
7. **File size**: ≤ 300 lines any file, ≤ 100 lines any React `.tsx`, ≤ 120 lines any class/struct.
8. **No magic strings or numbers**. Use enum or typed constant.
9. **Definitions in dedicated files**. Types, enums, constants get their own file.
10. **DRY is priority #1**. Duplicate = extract now.
11. **Small reusable components**. 3+ components in a feature: produce a Mermaid component diagram first.
12. **Immutable-first (Rust-style)**. Assign once. No reassignment except loop indices. Prefer `const`/`let`/`final`.
13. **Error-manage folder is law**. If `spec/**/error-manage/` exists, follow it exactly.
14. **Assets** go to `assets/<NN-folder>/<NN-file>.<ext>` with two-digit sequence prefixes.

## Line-Gap and Whitespace Style

1. One blank line before every `return`/`throw`, unless it is the only statement in the block.
2. One blank line after `}`, unless the next line is `}`, `else`, `case`, or `catch`.
3. Never two blank lines in a row.
4. No blank line right after `{` or right before `}`.
5. One blank line between top-level declarations.
6. Grouped imports with one blank line between groups: stdlib, third-party, first-party absolute, first-party relative.
7. Trailing newline at EOF. No trailing whitespace.

## Error Management (One-Liner Digest)

Full spec: `spec/03-error-manage/`.

- **Never swallow**: every `catch` logs (operation + key inputs), then rethrows or returns a typed error.
- **Wrap, do not lose**: `apperror.Wrap(err, "op", ctx)` (Go) / `throw new AppError(cause, {op, ctx})` (TS). Preserve the original stack.
- **Typed errors only**: no `throw "string"`, no bare `panic`. Use `AppError`/`apperror.Result[T]` with a registered code.
- **Registered codes**: user-visible errors live in `spec/03-error-manage/03-error-code-registry/`. No ad-hoc codes.
- **Universal envelope**: backend returns `{ data, errors[], meta }`. Frontend parses via one `parseEnvelope()` helper.
- **Log level = severity**: `debug` trace, `info` lifecycle, `warn` recoverable, `error` user-visible, `fatal` process-exit only.
- **Context on every log**: operation, request/session id, key inputs. Never secrets or PII beyond user id.
- **Verify both directions**: curl the backend AND inspect frontend detection logic before claiming a fix.
- **Retrospective on repeats**: same error class twice = add file under `spec/03-error-manage/01-error-resolution/03-retrospectives/`.
- **Frontend**: errors flow through the global error store to the error modal. No per-component alert boxes.

## Data and Schema Rules

1. Tables/types/entities: PascalCase. Fields/columns: camelCase. JSON keys: PascalCase.
2. PK: `int auto-increment`, named `{TableName}Id`. No UUIDs.
3. `Type`/`Status`/`Category`/`Kind` columns: 1-N or N-M join table with a registered enum. Never a free-form string column.
4. Entity/ref tables need `Description TEXT NULL`. Transactional tables need `Notes` + `Comments TEXT NULL`. All nullable, no `DEFAULT`. Join tables exempt.
5. Default DB SQLite. Prefer ORM. Define joins, PK, FK explicitly.
6. Any DB-touching PR includes a Mermaid ERD.

## Language One-Liners

- **Go**: `apperror.Result[T]` not `(T, error)`. `apperror.New`/`Wrap` not `fmt.Errorf`. `type X byte`+`iota` not string enums.
- **TypeScript**: `Promise.all` for independent async. No `any`. `readonly` on interface fields by default.
- **Rust**: `Result<T,E>` + `thiserror`. `let`, not `let mut`, unless mutation is the point.
- **PHP**: enum comparison via `->isEqual()`, never `===`.
- **PowerShell**: `Verb-Noun` PascalCase functions, `lowercase-kebab-case` files.
- **C#**: PascalCase methods/properties, `_camelCase` private fields, `I`-prefix interfaces.

## Important

- Rules override convenience. If a rule conflicts with a quick fix, follow the rule.
- When in doubt, ask before writing code.
- Do not invent specs. If the source is missing, ask.
