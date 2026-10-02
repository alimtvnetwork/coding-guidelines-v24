# Subtask 03: Heading Checker Linter Implementation (Task B)

**Parent Plan:** `.ai-memory/plans/pending/17-spec-ticket-for-blind-agents.md`  
**Target File:** `linter-scripts/check-spec-ticket-headings.py`  
**Assigned Role:** Subagent / Implementer

---

## 1. Goal & Boundaries

Implement a lightweight, standalone Python 3 script (`linter-scripts/check-spec-ticket-headings.py`) that mechanically parses and verifies the 10 mandatory level-2 headings of any executable spec ticket.

### Hard Constraints:
- Pure Python 3 standard library only (`sys`, `re`, `pathlib`, `argparse`).
- No external packages (zero pip dependencies).
- Fast execution (<100ms).
- Must print `PASS` and exit with code 0 on valid tickets.
- Must exit with non-zero exit code (1) and print clear diagnostic errors when headings are missing, renamed, or re-ordered.

---

## 2. In-Scope Detection & Heading Validation Logic

### 2.1 File Scope Determination
A file is determined to be an executable spec ticket if:
1. It contains the exact line `# Executable spec ticket`, OR
2. It contains both `## Context` and `## Do not touch`.

Files not meeting this threshold are skipped silently (to avoid failing the historical 750+ spec files).

### 2.2 Mandatory Headings Sequence
```python
REQUIRED_HEADINGS = [
    "## Context",
    "## Current state",
    "## Proposed change",
    "## Acceptance criteria",
    "## Testing plan",
    "## Rollback",
    "## Files",
    "## Out of scope",
    "## Do not touch",
    "## Checklist",
]
```

### 2.3 Diagnostic Reporting
- If a heading is missing:
  `print(f"{filepath}: missing heading '{heading}'")`
- If headings appear out of order:
  `print(f"{filepath}: heading '{heading}' appeared out of order")`
- If duplicate headings appear:
  `print(f"{filepath}: duplicate heading '{heading}'")`

---

## 3. Verification Protocol

1. **Positive Test:**
   Run checker on `02-spec/01-spec-authoring-guide/15-executable-spec-ticket.md`:
   ```bash
   python linter-scripts/check-spec-ticket-headings.py
   # Expect: PASS, exit code 0
   ```
2. **Negative Test (Mutation Check):**
   Create a temporary file with `## Rollback` removed, run checker:
   ```bash
   python linter-scripts/check-spec-ticket-headings.py temp-ticket.md
   # Expect: temp-ticket.md: missing heading '## Rollback', exit code 1
   ```
   Immediately clean up and remove `temp-ticket.md`.
