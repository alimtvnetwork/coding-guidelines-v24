# Subtask 01: 100 Improvements Prompts Suite Architecture & Contracts

## 1. Task Details
- **Parent Task:** `01-create-prompts-for-100-improvements`
- **Subtask Code:** `Subtask-01`
- **Slug:** `01-subtask-architecture-and-contracts`
- **Specification:** `02-spec/21-app/01-create-prompts-for-100-improvements/01-architecture-spec.md`
- **Status:** `COMPLETED`
- **Plan Motto:** `[/plan](slashCommand;plan) first before doing work for improved work or less friction, faster work and more pure and effective work.`

---

## 2. Owned Deliverables & Scope
- `02-spec/21-app/01-create-prompts-for-100-improvements/01-architecture-spec.md`
- `.ai-memory/plans/subtasks/01-create-prompts-for-100-improvements/01-subtask-architecture-and-contracts.md`

---

## 3. Action Checklist

- [x] **Architecture Specification Authoring:**
  - [x] Formulate complete system overview, problem statement, and continuous improvement flywheel mechanics in `02-spec/21-app/01-create-prompts-for-100-improvements/01-architecture-spec.md`.
  - [x] Construct comprehensive Mermaid architecture diagram (`flowchart TD`) mapping triggers, prompt routing, inspection engines, catalog generation, plan enqueueing, V6 swarm execution, and verification.
- [x] **Prompt 01 (UI Level) Specification:**
  - [x] Specify semantic color tokens (`--bg-canvas`, `--bg-surface`, `--text-primary`, `--border-default`, etc.) and elimination of raw hex literals.
  - [x] Detail WCAG 2.1 AAA (7:1) and AA (4.5:1 / 3:1) contrast ratio calculations for text, large text, and interactive controls.
  - [x] Document the 4-plane depth hierarchy: Plane 0 (Canvas), Plane 1 (Surface), Plane 2 (Elevated), Plane 3 (Modal/Overlay).
  - [x] Detail the 60/30/10 visual balance rule (60% canvas neutral, 30% structural surface, 10% high-contrast accent/CTA).
  - [x] Mandate the live DOM text standard with zero rasterized or baked SVG text.
  - [x] Define button and interactive control requirements ($44 \times 44\text{px}$ minimum touch targets, full hover/active/focus-visible/disabled states).
  - [x] Define motion physics standards (spring curves, cubic-bezier, $\le 250\text{ms}$ duration budget, `prefers-reduced-motion` compliance).
- [x] **Prompt 02 (Code Base) Specification:**
  - [x] Establish universal error management via `*appfault.AppError` and monadic result envelopes (`Result[T]`, `ResultSlice[T]`).
  - [x] Define the complete Positive Boolean Standard (Principles P1 through P9: affirmative prefixes `is`/`has`, elimination of negative words, named guards, complex expression decomposition, explicit parameters, zero mixed polarity, no inline assignments, no raw system calls in logic, and absolute ban on explicit `== true`).
  - [x] Detail vertical spacing rules (Rules 4, 5, 10 for blank lines before exits, after closing braces, and before `if` statements).
  - [x] Enforce function length limits of $\le 15$ lines (preferred $\le 8$ lines) and parameter struct encapsulation for $> 2\text{--}3$ arguments.
  - [x] Set modular file sizing targets of $\le 100$ lines and type safety standards (zero bare `any`).
- [x] **Prompt 03 (Whole System) Specification:**
  - [x] Detail UX friction elimination (click path streamlining, keyboard parity, optimistic UI, graceful error recovery).
  - [x] Formalize the 4-part cause-and-effect analysis matrix: (1) Root Cause, (2) Immediate Effect, (3) Proposed Remediation, (4) Blast Radius & Mitigation.
  - [x] Guarantee zero-harm and zero-regression safety through non-destructive evolution, shadow verification, and isolated checkpoints.
  - [x] Formalize backward compatibility preservation across database schemas and API contracts.
- [x] **Plan Enqueueing & Compounding Flywheel Mechanics:**
  - [x] Define master pending plan directory contracts in `.ai-memory/plans/pending/XX-<slug>.md`.
  - [x] Define bounded subtask partitioning in `.ai-memory/plans/subtasks/XX-<slug>/` ($5\text{--}8$ files per micro-batch).
  - [x] Detail the continuous flywheel mechanism where completing one cycle automatically enqueues the next 100 improvements.
- [x] **Skills Symmetry & Fleet Sync Architecture:**
  - [x] Specify dual-skill symmetry across Antigravity (`.agents/skills/`) and Cursor IDE (`.cursor/skills/`).
  - [x] Detail IDE skill link syntax (`file;.agents/skills/...` vs `file;.cursor/skills/...`).
  - [x] Reiterate the five non-negotiable multi-repository sync boundaries (Spec 21 exclusion, additive scripts only, bump script protection, memory isolation, zero secrets).
- [x] **Non-Negotiable Invariants Compliance:**
  - [x] Strictly relative Git paths only (`02-spec/...`, `.ai-memory/...`); zero absolute filesystem paths or URI schemes.
  - [x] Strictly lowercase filenames across all created files and directories.
  - [x] Zero git commands executed by the subagent.
  - [x] Zero build commands or test suite executions.

---

## 4. Verification Evidence & Quality Gates

1. **Owned Deliverables Existence:**
   - Specification file: `02-spec/21-app/01-create-prompts-for-100-improvements/01-architecture-spec.md` (verified non-empty).
   - Subtask file: `.ai-memory/plans/subtasks/01-create-prompts-for-100-improvements/01-subtask-architecture-and-contracts.md` (verified non-empty).
2. **Relative Path Hygiene Audit:**
   - Verified that neither deliverable contains absolute filesystem paths (`/home/`, `C:\`, `d:\`) or absolute URI schemes.
3. **Lowercase Convention Audit:**
   - All directory names (`01-create-prompts-for-100-improvements`) and file names (`01-architecture-spec.md`, `01-subtask-architecture-and-contracts.md`) are strictly lowercase.
