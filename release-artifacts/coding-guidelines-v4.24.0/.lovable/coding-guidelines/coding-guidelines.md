# Coding Guidelines

These rules are mandatory for all code generated in this project. The AI MUST read this file (and all referenced guideline files) before writing or modifying code.

## Files To Read Before Coding

1. This file: `.lovable/coding-guidelines.md`
2. All prompts: `.lovable/prompts/*.*`
3. All specs: `spec/01-app/*.md`
4. Master prompt: `spec/02-master-prompts/04-master-prompt-v4.md`
5. Memory protocol: `.lovable/memory/protocol/01-workflow-rules.md`
6. Decision log: `.lovable/memory/history/01-decisions.md`
7. Plan: `.lovable/memory/workflow/01-plan.md`

## Hard Rules

1. Functions ≤ 8 lines (hard cap; split otherwise).
2. No nested `if` statements.
3. `if` conditions must be positive and simple — no negations, no `!`.
4. Follow Boolean naming guidelines: prefix with `is` or `has`. Never use negative booleans.
5. Use proper, narrow types. Never `any`, `unknown`, `interface{}`, or any wide-range catch-all type. `Generic<T>` is the only exception.
6. No swallowed errors. Every `catch` must log per the project logging guidelines.
7. Files / classes ≤ 80–100 lines max.
8. No magic strings or numbers — use Enums or Constants.
9. Definitions live in their own dedicated files, not inline.
10. Keep code DRY — reusability is the highest-priority concern.
11. React/TypeScript components must be as small and reusable as possible. For multi-component features, plan first and produce a Mermaid component diagram.
12. Use Enums (typed) for any `Type`, `Kind`, `Status`, `Category` field.
13. If a `spec/**/error-manage/` folder exists, every error handler MUST follow those guidelines exactly. No exceptions.

## Data & Schema Rules

1. Tables, types, entities → **PascalCase**.
2. Fields/columns → **camelCase**.
3. JSON keys and values (when project uses JSON) → **PascalCase**.
4. Every primary key: `int auto-increment`, named `{PascalCaseTableName}Id`.
5. `Type` / `Status` / `Category` / `Kind` columns → 1-N or N-M join tables (never inline strings/enums in the row).
6. Use the smallest appropriate integer type for category IDs.
7. Default DB: SQLite. Prefer ORM. Define joins, PK/FK explicitly.
8. Any DB discussion must include a Mermaid ERD.

## Error & Logging

1. Catch → log → rethrow or handle. Never silent.
2. Log level appropriate to severity.
3. Include context (operation name, key inputs) in log messages.

## Important

- These guidelines override convenience. If a rule conflicts with a quick fix, follow the rule.
- When in doubt, ask before writing code.
