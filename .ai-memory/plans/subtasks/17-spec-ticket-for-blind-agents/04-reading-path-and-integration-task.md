# Subtask 04: Blind Agent Reading Path & Authoring Guide Integration (Tasks C & D)

**Parent Plan:** `.ai-memory/plans/pending/17-spec-ticket-for-blind-agents.md`  
**Target Files:**
- `02-spec/01-spec-authoring-guide/16-blind-agent-reading-path.md`
- `02-spec/01-spec-authoring-guide/readme.md`
- `02-spec/01-spec-authoring-guide/04-required-files.md`
- `02-spec/01-spec-authoring-guide/99-consistency-report.md`
- `.ai-memory/plans/readme.md`
**Assigned Role:** Subagent / Implementer

---

## 1. Goal & Boundaries

1. Author `02-spec/01-spec-authoring-guide/16-blind-agent-reading-path.md`: A concise (<120 non-blank lines), authoritative protocol establishing the exact 3-step reading sequence for incoming AI agents.
2. Update authoring guide index files (`readme.md`, `04-required-files.md`, `99-consistency-report.md`) to integrate files 15 and 16 cleanly without rewriting historical specs.
3. Transition plan 17 from pending to completed once all acceptance criteria are verified.

---

## 2. Reading Path Protocol (`16-blind-agent-reading-path.md`)

The file MUST mandate this 3-step sequence and strictly forbid reading outside these boundaries:

1. **Step 1: Read the Ticket Template & Contract:**
   - Ingest `02-spec/01-spec-authoring-guide/15-executable-spec-ticket.md`.
2. **Step 2: Read the Target Module Overview:**
   - Read the single module `readme.md` named in the ticket's `Context` or `Current state`.
3. **Step 3: Read Only the Target Files:**
   - Read strictly the files listed in that ticket's `Files` table.

### Non-Negotiable Boundaries for Blind Agents:
- STOP immediately after reading those files.
- DO NOT open `02-spec/07-design-system/` or other massive folders unless explicitly listed in the ticket.
- DO NOT treat historical `Ambiguity: None` claims as license to invent missing numbers or values.

---

## 3. Authoring Guide Registration Updates

### 3.1 `02-spec/01-spec-authoring-guide/readme.md`
- Add table rows in the file index for:
  - `15-executable-spec-ticket.md`
  - `16-blind-agent-reading-path.md`
- Under **Scoring Metrics**, add the rule:
  *An overview may say `Ambiguity: None` only when `99-consistency-report.md` in that same folder has result `PASS` and names no open contradiction. Otherwise write the contradiction in one sentence.*

### 3.2 `02-spec/01-spec-authoring-guide/04-required-files.md`
- Add `15-executable-spec-ticket.md` under **Strongly Recommended Files** (for single-change implementation tickets).
- Do NOT make tickets mandatory for every historical module.

### 3.3 `02-spec/01-spec-authoring-guide/99-consistency-report.md`
- Register `15-executable-spec-ticket.md` and `16-blind-agent-reading-path.md` in the inventory.

### 3.4 `.ai-memory/plans/readme.md`
- Move `17-spec-ticket-for-blind-agents.md` from `Pending Plans` to `Completed Plans` upon verified passing completion.
