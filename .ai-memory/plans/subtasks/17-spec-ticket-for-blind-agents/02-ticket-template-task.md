# Subtask 02: Executable Spec Ticket Template Specification (Task A)

**Parent Plan:** `.ai-memory/plans/pending/17-spec-ticket-for-blind-agents.md`  
**Target File:** `02-spec/01-spec-authoring-guide/15-executable-spec-ticket.md`  
**Assigned Role:** Subagent / Implementer

---

## 1. Goal & Boundaries

Author `02-spec/01-spec-authoring-guide/15-executable-spec-ticket.md` as the authoritative template for single-change executable specification tickets.

### Hard Constraints:
- Version stamp: `4.3.0` | Updated: `2026-10-02` | Status: `Active`.
- Must contain EXACTLY 10 Level-2 headings in this exact sequence:
  1. `## Context`
  2. `## Current state`
  3. `## Proposed change`
  4. `## Acceptance criteria`
  5. `## Testing plan`
  6. `## Rollback`
  7. `## Files`
  8. `## Out of scope`
  9. `## Do not touch`
  10. `## Checklist`
- Must enforce line cap: under 300 non-blank lines total (aim for ~200–260 lines).
- Must include a concrete, filled example (such as adding a heading checker) that avoids company names, client names, or external URLs.
- Ban the subjective phrases `works correctly` and `edge cases are handled`.

---

## 2. Heading Specifications & Required Rules

### Heading 1: `## Context`
Must declare five labeled items:
- `Who:` User role, internal developer, automated system, or AI agent affected.
- `Current:` Ground-truth verified current state.
- `Desired:` Exact outcome behavior.
- `Why now:` Immediate driver (bug, regression, blocking task, technical debt).
- `Done when:` Observable, measurable condition of completion.

### Heading 2: `## Current state`
- Must include a verification date (e.g. `2026-10-02`).
- Must provide at least one concrete `path:line` citation.

### Heading 3: `## Proposed change`
- Concise description of the implementation.
- Technical architecture, schema, or algorithm.

### Heading 4: `## Acceptance criteria`
- Numbered list of pass/fail criteria.
- Banned phrases: `works correctly`, `edge cases are handled`.
- Every item must be verifiable by a stranger or blind AI agent.

### Heading 5: `## Testing plan`
- Markdown table with three mandatory columns:
  `| Layer | What | Count |`
- Specifies tests at command, unit, integration, or file-read layers.

### Heading 6: `## Rollback`
- Explicit revert strategy.
- If stateless: `Revert the commit; no database migration or persistent data loss.`
- If schema: explicit down migration or schema restore step.

### Heading 7: `## Files`
- Markdown table with two mandatory columns:
  `| File | Change |`
- Strict relative git paths from repository root.
- No drive letters (e.g. `C:`, `D:`), no `file:///` URIs.
- Maximum 8 rows. If more than 8 files, ticket MUST be split into child tickets.

### Heading 8: `## Out of scope`
- At least one explicit bullet item preventing scope creep.

### Heading 9: `## Do not touch`
- At least one explicit bullet listing sensitive files, configs, or modules that must not be modified.

### Heading 10: `## Checklist`
- Mirrors every numbered acceptance criterion as a `- [ ]` checkbox item.

---

## 3. Verification & Acceptance Gate

1. File exists at `02-spec/01-spec-authoring-guide/15-executable-spec-ticket.md`.
2. All 10 headings appear in exact order.
3. Total non-blank line count is <= 300.
4. Linter `linter-scripts/check-spec-ticket-headings.py` passes with exit code 0.
