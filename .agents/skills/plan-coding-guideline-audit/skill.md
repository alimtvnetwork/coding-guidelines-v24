---
name: plan-coding-guideline-audit
description: Plan a structured coding guideline audit across repository codebases against 02-spec/02-coding-guidelines/.
---

# Plan Coding Guideline Audit

Autonomously plans a comprehensive audit of repository codebases against the master coding guidelines.

## Audit Areas

1. **Boolean Conventions & Affirmative Parameters:** `is` and `has` prefixes only, implicit positive checks, no mixed polarity, affirmative parameter/field naming (no single-letter `v bool` or bare `stop bool` -> `isStopOnFail bool`, `isStopped bool`).
2. **Control Flow:** Maximum nesting depth 1, guard clauses, flattened conditionals.
3. **Naming & Types:** Enum `Type` suffixes, PascalCase types, no underscores in Go.
4. **Error Handling & Result Containers:** Single Result return containers (`appfault.ResultMap[K, V]`, `appfault.ResultSlice[T]`, `appfault.Result[T]`), `*appfault.AppError` returns, pointer-attached null safety (`*Result[T]` with line-1 `if r == nil` guards), 4 core predicates (`IsCountOtherThan`, `IsEmpty`, `HasRecord`, `IsDefined`), no swallowed errors, typed error responses.
5. **Code Metrics:** Functions <= 15 lines, files <= 100 lines coding, blank line padding.

## Output

Generates structured audit logs and phased remediation plans in `.lovable/plans/pending/` with subtask micro-batches.
