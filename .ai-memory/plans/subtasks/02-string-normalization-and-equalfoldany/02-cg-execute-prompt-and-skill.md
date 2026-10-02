# Subtask 02: Author CG-Execute Prompt 37, Native Antigravity Skill & Catalog Integrations for String Normalization and EqualFoldAny

> **Parent Spec:** [`02-spec/21-app/02-string-normalization-and-equalfoldany/02-prompt-and-skill-spec.md`](../../../../02-spec/21-app/02-string-normalization-and-equalfoldany/02-prompt-and-skill-spec.md)  
> **Status:** `PENDING`  
> **Traceability ID:** `Task-02`  
> **Target Files:**
> - `01-prompts/15-cg-execute/37-string-normalization-and-equalfoldany.md` (to create)
> - `.agents/skills/cg-string-normalization-and-equalfoldany/skill.md` (to create)
> - `01-prompts/15-cg-execute/readme.md` (to update index table)
> - `01-prompts/readme.md` (to update master prompt catalog)
> - `.ai-memory/prompts.md` (to update prompt registry)

---

## 1. Subtask Objective & Context

This subtask governs authoring the guideline execution tooling and prompt infrastructure for string normalization and case-insensitive comparison optimization, establishing Prompt 37 and its corresponding native Antigravity skill:

1. **Prompt 37 (`37-string-normalization-and-equalfoldany.md`):** Complete autonomous execution prompt for auditing and refactoring repetitive string manipulation, manual trimming, and chained equality checks (`EqualFold` / `ToLower` OR-chains) across polyglot codebases into concise, zero-allocation `strutil` helper calls.
2. **Native Skill (`cg-string-normalization-and-equalfoldany`):** Corresponding Antigravity skill in `.agents/skills/` for IDE and CLI runtime invocation.
3. **Search First Directive Implementation:** Rigorous mandate ensuring all refactoring passes check for existing string utilities (`pkg/strutil/strutil.go` in Go, `strutil.ts` in TS, `strutil.rs` in Rust) before modifying call sites or creating duplicate logic.
4. **Catalog Registries Updates:** Registering Prompt 37 across `01-prompts/15-cg-execute/readme.md`, `01-prompts/readme.md`, and `.ai-memory/prompts.md`.

---

## 2. Deliverable Specifications

### 2.1 Deliverable A: Prompt 37 Specification (`37-string-normalization-and-equalfoldany.md`)

- **Placement:** `01-prompts/15-cg-execute/37-string-normalization-and-equalfoldany.md`
- **Header Format:** Strict V6 Parameter Header:
  ```text
  N = 300 (Total self-loop steps budget — editable top-header parameter, default: 300)
  A = 2   (MANDATORY number of spawned autonomous subagents running concurrently via invoke_subagent, default: 2)
  H = 2   (Operational hands per agent: dual-task batch capacity & parallel tool dispatch, default: 2)
  C = 30  (Tool calls per worker before it must report, default: 30)

  System Concurrency Capacity = A × H = 2 agents × 2 hands = 4 concurrent subtask operations
  PHASE_1_BUDGET = N / 2   (Steps 1 .. 150: Planning, Parallel Discovery Subagents, Detailed Spec, and Lean Subtask Generation)
  PHASE_2_BUDGET = N / 2   (Steps 151 .. 300: Mandatory Parallel Subagent Execution, Self-Looping, Targeted Quality Linting)
  WAVES = ceil(subtasks / (A x H))
  ```
- **Slash Commands:** Semicolon format: `[/goal](slashCommand;goal)`, `[/learn](slashCommand;learn)`, `[/plan](slashCommand;plan)`.
- **Mandatory Subagent Spawning Gate:** Mandatory `invoke_subagent` tool calls across all stages (`A = 2, H = 2`). Zero solo execution permitted.
- **Search First Directive:**
  - Mandatory check for `pkg/strutil/strutil.go`, `04-code/golang/pkg/strutil/strutil.go`, or repo string utilities using `gitmap find-files-any "strutil"`.
  - Zero tolerance for duplicated local helper functions in consumer packages.
- **Code Reduction Transformation Rules:**
  - *Transformation 1:* Replace manual trimming + `strings.EqualFold(target, "a") || strings.EqualFold(target, "b")` with `strutil.EqualFoldAnyTrim(target, "a", "b")`.
  - *Transformation 2:* Replace untrimmed `strings.EqualFold` OR-chains with `strutil.EqualFoldAny(target, "a", "b")`.
  - *Transformation 3:* Replace `strings.ToLower(strings.TrimSpace(target)) == "val"` with `strutil.EqualFoldAnyTrim(target, "val")` or `strutil.NormalizeLowerTrim(target)`.
  - *Transformation 4:* Eliminate temporary allocations in loops; hoist static target strings.
  - *Transformation 5:* Guard clauses, early returns, positive booleans (`is`/`has` only), and vertical blank lines.
- **AST Discovery Patterns using GitMap:**
  - Chained EqualFold: `gitmap aum search -r "EqualFold\(.*\|\|.*EqualFold\(" [dir] -e .go`
  - Trim + EqualFold: `gitmap aum search -r "TrimSpace\(.*EqualFold" [dir] -e .go`
  - ToLower Equality: `gitmap aum search -r "strings\.ToLower\(.*==.*\|\|" [dir] -e .go`
  - TypeScript: `gitmap aum search -r "\.toLowerCase\(\)\s*===\s*.*\|\|" [dir] -e .ts`
  - Rust: `gitmap aum search -r "eq_ignore_ascii_case\(.*\|\|" [dir] -e .rs`
- **3-Phase Pipeline:**
  - *Phase 0:* Bootstrap `.agents/skills/cg-string-normalization-and-equalfoldany/skill.md`.
  - *Phase 1A:* Verbatim capture, task extraction, Turn 1 visible task breakdown.
  - *Phase 1B:* GitMap discovery, Violation Ledger in `.ai-memory/plans/pending/`, subtasks in `.ai-memory/plans/subtasks/`.
  - *Phase 2:* Parallel worker execution in 5–8 file micro-batches. Refactor AST violations using `strutil` helpers.
  - *Phase 3:* Subtask consolidation into `.ai-memory/plans/completed/`, atomic push with `gitmap cpf "<module> - <summary>"`.
- **Core Operational Rules (R1 to R16 Cited by ID):**
  - R1: Zero builds or test suites.
  - R2: Targeted checks only.
  - R3: Evidence or it did not happen.
  - R4: Never invent commands, flags, or paths.
  - R5: Mandatory subagents (`invoke_subagent`, A=2, H=2).
  - R6: One owner per file (disjoint bounding boxes).
  - R7: Git safety & isolation.
  - R8/R9: Atomic commit & push via GitMap (`gitmap cpf "<module> - <summary>"`).
  - R10: Zero unauthorized releases.
  - R11: Strict relative git paths & lowercase hygiene.
  - R12: No polling / immediate turn yielding.
  - R13: Two-strike retry cap & anti-looping.
  - R14: 100% ambiguity & decision boundaries.
  - R15: Zero generated artifacts committed.
  - R16: Zero secrets in standard repos (`repo-secrets` offload via `gitmap rs`).

### 2.2 Deliverable B: Native Antigravity Skill (`cg-string-normalization-and-equalfoldany`)

- **Placement:** `.agents/skills/cg-string-normalization-and-equalfoldany/skill.md`
- **YAML Frontmatter:**
  ```yaml
  ---
  name: cg-string-normalization-and-equalfoldany
  description: Autonomously audits, refactors, and standardizes string normalization, case-insensitive comparisons, and chained equality into reusable strutil helpers (EqualFoldAny, EqualFoldAnyTrim) across Go, TypeScript, and Rust.
  ---
  ```
- **Contents:**
  - Primary trigger keywords and aliases (`cg-string-normalization`, `cg-equalfoldany`, `strutil-normalization`, `audit-string-comparisons`).
  - Search First protocol for `pkg/strutil/strutil.go`.
  - Step-by-step execution workflow (Discovery -> Planning -> Micro-Batched Refactoring -> Verification).
  - Polyglot code templates and Good vs Bad comparisons for Go, TypeScript, and Rust.
  - AST discovery commands via GitMap.
  - Quality verification checklist (targeted linters, zero build/test runs).

### 2.3 Deliverable C: Registries & Catalog Updates

- **`01-prompts/15-cg-execute/readme.md`:** Add row 37 to the Prompts Catalog & Execution Order table:
  - Sequence: `37`
  - Filename: `37-string-normalization-and-equalfoldany.md`
  - Purpose: `String Normalization, EqualFoldAny & Code Reduction`
- **`01-prompts/readme.md`:** Register prompt 37 under section 15.
- **`.ai-memory/prompts.md`:** Register prompt 37 with aliases and detailed summary.

---

## 3. Execution Step-by-Step Checklist

- [ ] **Step 1:** Author `01-prompts/15-cg-execute/37-string-normalization-and-equalfoldany.md` with complete V6 parameter header (`N = 300, A = 2, H = 2, C = 30`), semicolon slash commands, Search First Directive, AST patterns, code reduction rules, and R1–R16 rules cited by ID.
- [ ] **Step 2:** Author `.agents/skills/cg-string-normalization-and-equalfoldany/skill.md` with valid YAML frontmatter, polyglot templates, and GitMap AST discovery commands.
- [ ] **Step 3:** Update `01-prompts/15-cg-execute/readme.md` catalog table with prompt 37.
- [ ] **Step 4:** Update `01-prompts/readme.md` master catalog with prompt 37.
- [ ] **Step 5:** Update `.ai-memory/prompts.md` prompt registry with prompt 37.
- [ ] **Step 6:** Run sequence integrity, prompt index, and relative path linters to verify clean, zero-drift integration.

---

## 4. Verification & Acceptance Criteria

- **AC-CG-037-A:** Prompt `37-string-normalization-and-equalfoldany.md` exists, defines `N = 300, A = 2, H = 2, C = 30`, includes semicolon slash commands, and contains zero broken links.
- **AC-CG-037-B:** Skill `.agents/skills/cg-string-normalization-and-equalfoldany/skill.md` exists with valid YAML metadata, polyglot templates (Go, TS, Rust), and GitMap search patterns.
- **AC-CG-037-C:** Search First Directive for `pkg/strutil/strutil.go` is prominently documented and enforced in both prompt and skill.
- **AC-CG-037-D:** Catalogs (`01-prompts/15-cg-execute/readme.md`, `01-prompts/readme.md`, `.ai-memory/prompts.md`) register sequence #37 cleanly.
- **AC-CG-037-E:** Strict relative paths enforced; zero absolute filesystem paths or `file:///` URIs present.

**Verification commands:**

```bash
python linter-scripts/check-prompts-loaded.py
python linter-scripts/check-relative-paths.py
python 03-ai-scripts/21-sequence-integrity-linter.py 01-prompts/15-cg-execute
```

**Expected Result:** Exit code 0 across all verification linters.
