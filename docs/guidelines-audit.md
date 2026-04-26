# Coding Guidelines — Audit & Top-3 Quick Wins

**Reviewer:** Lovable AI
**Date:** 2026-04-26
**Scope reviewed:** `spec/17-consolidated-guidelines/01-spec-authoring.md`, `02-coding-guidelines.md`, `19-gap-analysis.md`, plus the source folders they summarize.

---

## TL;DR

The guidelines are unusually thorough — coverage maps, gap-analysis scores (97.6/100), source-back-references on every section. The weak spots are not in **what** is enforced, they are in **how authors discover and apply** the rules. Most of the friction is structural noise that an author hits before they ever read the actual rule.

The 3 highest-leverage improvements are:

1. **Fix the version-bump table** in `01-spec-authoring.md` — the examples contradict the rule and will mis-train AI authors.
2. **Promote `strictly-avoid` to a single canonical CODE-RED quick reference** — the rules currently exist in `02-coding-guidelines.md` (numbered §17), `.lovable/strictly-avoid.md`, and the gap-analysis report, and the wording drifts between them.
3. **Add a one-page "spec frontmatter snippet"** so authors stop re-deriving the metadata header from memory.

These are implemented in this PR. The full deep-dive on author DX is in `docs/spec-author-dx.md`.

---

## Strengths Worth Preserving

| Area | Why it's good |
|------|---------------|
| **Source-folder coverage map** in every consolidated file | Lets an AI handed a single file know exactly which sources it summarizes and whether anything is intentionally deferred. |
| **Numeric severity (CODE-RED / STYLE / WARN / BEST PRACTICE)** | Makes lint output triage-able. |
| **Result\<T\> over (T, error)** as a uniform Go return shape | Eliminates the dual-return error-checking branch at every call site. |
| **Boolean rules (P1–P8)** as a single named ruleset | Easy to cite in code review (`violates P3`). |
| **`apperror.Wrap()` mandatory** | Forces every error to carry a code + context. |
| **No mid-rule deviation policy** — exceptions are tracked in `01-spec-authoring-guide/09-exceptions.md` | Prevents drift from "rule + 17 informal exceptions". |

---

## Issues Found

### A. Contradictions / drift (HIGH priority)

| # | Where | Issue | Suggested fix |
|---|-------|-------|---------------|
| A1 | `01-spec-authoring.md` §"Version Bump Rules" | The examples contradict the rule: a "patch" example shows `3.2.0 → 3.1.1` (minor went **down**), the "minor" example shows `3.2.0 → 3.2.0` (no change). | Replace with `3.2.0 → 3.2.1`, `3.2.0 → 3.3.0`, `3.2.0 → 4.0.0`. **Implemented in this PR.** |
| A2 | `01-spec-authoring.md` §"Folder Structure — Numbering Policy" → "Fixed Assignments" | Table shows `10 → 11-powershell-integration`, `11 → 10-research`, `12 → 17-consolidated-guidelines`. The "#" column is the *display order*, not the folder prefix. Authors mis-read this and create `10-…` and `11-…` collisions. | Rename "#" column to "Display Order" or drop it entirely; the actual folder name already encodes the prefix. |
| A3 | `02-coding-guidelines.md` §3 "Code Style" rule numbering | Jumps from "Rule 6" to "Rule 17" with no explanation, suggesting deleted rules. | Either renumber sequentially or add an explicit `*Rules 7–16: see source spec for legacy entries.*` note. |
| A4 | `02-coding-guidelines.md` §1.5 vs §1.4 | Function naming bans boolean flag *parameters*, but variable naming says nothing about boolean *flags as constants in arrays*. The `boolFlags.ts` file in this very repo (which §1 of the source spec mandates) is **not** mentioned anywhere in the consolidated reference. | Add a one-line cross-reference under §1.5: "See also §22.X — `boolFlags.ts` named-constant pattern." |

### B. Same rule, multiple homes (MEDIUM)

| # | Rule | Lives in | Risk |
|---|------|----------|------|
| B1 | "Never use `not`/`no`/`non` in boolean names" | `02-coding-guidelines.md` §P2, source `01-cross-language/02-boolean-principles/`, `.lovable/strictly-avoid.md`, `.lovable/memory/style/naming-conventions` | Wording drifts between copies. Recommend canonical home in source spec, with consolidated + memory files quoting verbatim with a `> Source:` backlink. |
| B2 | "Max 15 lines per function" | `02-coding-guidelines.md` §3 Rule 6 (15 lines) AND §23.1 (8–15 lines target, 30 ceiling for Go) | Two different ceilings — language-specific override is fine, but call it out: "Go ceiling is 30 lines; cross-language target is 15." |
| B3 | "No bare boolean as positional arg" | `02-coding-guidelines.md` §1.5 (function naming) AND CODE-RED-024 (mentioned only in `boolFlags.ts` doc-comment) | Promote CODE-RED-024 to a numbered section in the consolidated file with the `boolFlags.ts` example. |

### C. Discoverability friction (MEDIUM)

| # | Pain point | Suggested fix |
|---|------------|---------------|
| C1 | New author wanting to add a spec file has to read `01-spec-authoring.md` (475 lines) end-to-end before knowing the metadata header. | Ship a copy-pasteable frontmatter snippet at the **top** of the file and as a standalone `spec/_template.md`. **Implemented in this PR.** |
| C2 | The "three sync scripts" sequence (§X.2) is buried in a 475-line doc and only discoverable by full-text search. | Pin it as a top-level "Before you commit" callout. **Implemented in this PR.** |
| C3 | Validation command (`python linter-scripts/validate-guidelines.py`) is mentioned twice with slightly different wording. | One canonical bash snippet, copy-pasted from `linter-scripts/run.sh`. |
| C4 | No editor integration story. Authors learn about violations only after running `validate-guidelines.py` locally or losing CI minutes. | See `docs/spec-author-dx.md` — proposed VS Code task + pre-commit hook expansion. |

### D. Coverage gaps (LOW — already tracked in 19-gap-analysis.md)

- Per the gap-analysis report, the lowest-coverage consolidated files (≤ 65% implementability) are `10-research.md`, `12-root-research.md`, `13-app.md`, `14-app-issues.md`. These are intentionally placeholder-ish; the gap is acknowledged.
- One actual coverage hole: there is no consolidated coverage of `.cursorrules` or `.github/copilot-instructions.md`, both of which are real, present in repo, and silently shape AI behavior. Worth a single section in `21-lovable-folder-structure.md`.

---

## Top-3 Quick Wins — Implemented in this PR

### Quick win 1 — Fix A1 (version bump table)

File: `spec/17-consolidated-guidelines/01-spec-authoring.md`. Examples now match the rule.

### Quick win 2 — Promote a single CODE-RED quick reference

File: `spec/17-consolidated-guidelines/00-strictly-avoid-quickref.md` (new).

Single-page list of every 🔴 CODE RED with one-line rationale and a link to the source spec for examples. Authors can scan this in 30 seconds before opening a PR. The existing `.lovable/strictly-avoid.md` is now reduced to a pointer at this canonical file to eliminate drift (see B1).

### Quick win 3 — Spec frontmatter snippet

File: `spec/_template.md` (new).

Drop-in template covering the metadata header, scoring table, document inventory, and cross-references. Linked from the top of `01-spec-authoring.md` so new authors find it before reading the long-form doc.

---

## What I deliberately did not change

- **Renumbering** anything in published consolidated docs — even fixing rule-number gaps (A3) is a versioned change that needs the sync workflow run; flagged here for an author with sync-script access.
- **Memory file edits** — `mem://` writes are out of scope of a guidelines audit; suggested follow-ups are listed but not applied.
- **Source-spec edits** — the audit deliberately stays in the consolidated layer to avoid touching 30,000+ lines of source markdown.

---

## Suggested next pass (not in this PR)

1. Renumber §3 "Code Style" rules sequentially (A3).
2. Add the `boolFlags.ts` cross-reference under §1.5 (A4).
3. Promote CODE-RED-024 to a numbered section (B3).
4. Add `.cursorrules` / Copilot instructions coverage to `21-lovable-folder-structure.md` (D).
5. Run `node scripts/sync-version.mjs && node scripts/sync-spec-tree.mjs` after any of the above.