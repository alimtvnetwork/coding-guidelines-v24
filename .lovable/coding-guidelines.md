# Coding Guidelines

> NOTE: Conflict with folder-level `spec/02-coding-guidelines/` or `spec/coding-guidelines/`. The folder-level spec wins over this file.

- Language/Runtime: TypeScript, PHP, Python, Go.
- Enums: TypeScript string unions are banned. All Enums must end with the `Type` suffix.
- Error Handling: No silent failures. Use explicit boolean states (e.g. `isFail`), do not invert success booleans (e.g. `!isSuccess`).
- Naming Rules: PascalCase everywhere.
- Project Bans: No magic strings or numbers except for loggers. See `.lovable/strictly-avoid.md`.
