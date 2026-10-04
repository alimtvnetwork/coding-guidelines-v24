# Subtask [01]: Inject Strict Relative Git Paths Mandate into Letterly & Cursor Prompts

Traceability ID: Task-03
Spec Reference: [02-spec/21-app/12-relative-paths-and-legacy-skills-purge/01-architecture-spec.md](../../../../02-spec/21-app/12-relative-paths-and-legacy-skills-purge/01-architecture-spec.md)
Target Files:
- `AGENTS.md`
- `01-prompts/22-letterly/01-mobile-letterly.md` through `10-execute-with-release-letterly.md`
- `01-prompts/23-cursor-prompts/01-mobile-letterly-cursor.md` through `10-execute-with-release-letterly-cursor.md`

Action:
1. Update `AGENTS.md` Section 10 to include Item 3 relative paths requirement in Letterly & Cursor action items.
2. Update all 10 Letterly formatters in `01-prompts/22-letterly/` to require strict relative Git paths and ban absolute paths and `file:///` URIs.
3. Update all 10 Cursor formatters in `01-prompts/23-cursor-prompts/` to mirror the relative paths requirement with `(file;.cursor/skills/...)` syntax.

Acceptance Criteria:
- All 10 Letterly prompts explicitly mandate relative paths.
- All 10 Cursor prompts explicitly mandate relative paths.
- `AGENTS.md` Section 10 specifies the relative paths invariant.
- Linters pass (`python linter-scripts/check-prompts-loaded.py`).
