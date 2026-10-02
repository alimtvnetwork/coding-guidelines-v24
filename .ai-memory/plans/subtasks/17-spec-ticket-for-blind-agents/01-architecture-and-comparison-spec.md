# Subtask 01: Architecture Specification & G-Spec Comparative Analysis

**Parent Plan:** `.ai-memory/plans/pending/17-spec-ticket-for-blind-agents.md`  
**Target Files:**
- `02-spec/21-app/04-gspec-adaptation/01-architecture-spec.md`
- `.ai-memory/plans/subtasks/17-spec-ticket-for-blind-agents/01-architecture-and-comparison-spec.md`

---

## 1. Context & Executive Assessment

### 1.1 What Coding Guidelines Got Right (Unprecedented Precision)
The specifications in `coding-guidelines` represent an extraordinarily high degree of mechanical precision and architectural rigor:
1. **Zero Degrees of Freedom for Hallucination:** By mandating strict return types (`*appfault.AppError`, `Result[T]`), positive boolean prefixes (`is`, `has`), zero explicit `== true`, and 100-line file caps, LLMs are constrained to deterministic code generation.
2. **Polyglot Consistency:** Equal standards across Go, TypeScript, PHP, Rust, and Python ensure that cross-stack developers and AI agents do not invent idiosyncratic patterns.
3. **Automated Mechanical Enforcement:** Hundreds of dedicated AST linters in `linters-cicd/` and `linter-scripts/` provide immediate feedback without expensive full test suites.
4. **Split-DB & Positive Schema Architecture:** Eliminating NULL ambiguity and negative columns in SQLite schemas creates bulletproof data integrity.

### 1.2 The Current Bottleneck: Cognitive Overload for Blind Agents
Despite its code-level excellence, the specification library faces an operational challenge when consumed by fresh or "blind" AI agents:
1. **Library Exhaustion:** 750+ files and 160,000+ lines across 25 folders cause agents to get lost in context windows.
2. **False Confidence Signals:** Many historical overview files declare `Ambiguity: None` even when subtle cross-file contradictions exist, leading agents into dead ends.
3. **Absence of Single-Change Execution Tickets:** Modules describe complete system architectures (e.g. whole design systems or worker topologies) rather than bounded, single-task execution units.

---

## 2. G-Spec (`gstack/spec`) vs `coding-guidelines`: Direct Comparison

| Dimension | G-Spec (`gstack/spec`) | Coding Guidelines (`02-spec/`) | The Synthesis (What We Are Building) |
|:---|:---|:---|:---|
| **Entry Point** | 5-Phase Interrogation loop (`Who`, `Current`, `Desired`, `Why now`, `Done when`) | 25 multi-file spec folders with historical overviews | Bounded 3-step reading path (`16-blind-agent-reading-path.md`) |
| **Execution Unit** | Single GitHub issue / backlog ticket with 14 quality rules | Broad modular specs (e.g., `08-non-cli-module-template.md`) | Standalone Executable Spec Ticket (`15-executable-spec-ticket.md`) |
| **Scope Bounding** | Hard rule: max 8 files per ticket, split if larger; explicit `Out of scope` | Micro-batching rules (5–8 files) scattered across execution prompts | Standardized `Files` table (max 8 rows) + mandatory `Out of scope` section |
| **Verification** | Pass/fail criteria; ban on "works correctly" / "edge cases handled" | Acceptance criteria in `97-acceptance-criteria.md` | Strict binary pass/fail criteria + 10-heading automated linter |
| **Safety Gate** | "What's working well / Do not touch" list + rollback strategy | CODE RED rules, `.ai-memory/strictly-avoid.md` | Mandatory `Do not touch` + explicit `Rollback` section in every ticket |
| **Language Standards** | High-level prose guidelines, no AST linters | Deep AST linters (booleans, line lengths, returns, SQLite) | Maintain 100% of our code linters; adopt G-Spec's ticket structure |
| **External Runtime** | Bun, Codex quality gate, Claude worktree spawn, telemetry | Pure Python 3 standard library, native GitMap orchestration | Zero external runtime dependencies; standard library Python only |

---

## 3. What to Adopt from G-Spec

1. **The 10-Heading Ticket Contract:**
   - `Context` (with the 5 Whys: Who, Current, Desired, Why now, Done when)
   - `Current state` (with `path:line` citations and verification date)
   - `Proposed change`
   - `Acceptance criteria` (numbered, binary pass/fail, zero vague adjectives)
   - `Testing plan` (table with Layer, What, Count)
   - `Rollback` (exact revert strategy)
   - `Files` (table with File, Change; max 8 files)
   - `Out of scope` (at least one bullet)
   - `Do not touch` (at least one bullet)
   - `Checklist` (mirrors acceptance criteria as `- [ ]` checkboxes)

2. **Automated Mechanical Heading Linter:**
   - `linter-scripts/check-spec-ticket-headings.py` validating that every ticket contains all 10 headings in exact order.

3. **Curated 3-Step Reading Path:**
   - `16-blind-agent-reading-path.md` instructing blind agents to read ONLY:
     1. The ticket template (`15-executable-spec-ticket.md`).
     2. The single module `readme.md` referenced in the ticket.
     3. The specific files listed in the ticket's `Files` table.

---

## 4. What to Exclude (Deliberately Kept Out)

- **GStack Preambles & Telemetry:** No `gstack-skill-start`, `gstack-skill-end`, or analytics logging.
- **External Model Dependencies:** No dependency on Bun, Codex, or proprietary LLM scoring APIs.
- **GitHub Issue Tracker Couplings:** No forced dependency on `gh issue create` or issue deduping; tickets live as version-controlled markdown files.
- **Worktree Orchestration Scripts:** GitMap and native git branches remain the authoritative isolation mechanisms.
