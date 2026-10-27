# Skills Inventory Reduction Matrix & Deprecation Proposal

## 1. Overview
The `.agents/skills/` and `.cursor/skills/` directories currently contain **108 skills**. Having over a hundred individual prompt wrappers creates high cognitive overhead for agents, bloats the system prompt context, and duplicates documentation.

This audit classifies the entire skill set into:
- **Core Orchestrator & Workflow Skills (KEEP):** Foundational skills actively required for daily development, execution, memory, and testing.
- **Unified Master Skills (CONSOLIDATE):** Broad umbrella skills like `follow-cg-execute`, `release`, and `follow-coding-guidelines`.
- **Candidate Skills for Archival/Consolidation (REDUCE):** Granular prompt wrappers whose instructions are already encapsulated in the master orchestrators.

---

## 2. Proposed Skills Reduction Matrix

### Cluster A: Individual Coding Guideline Wrappers (19 Skills -> Consolidated into `follow-cg-execute`)
These individual skills duplicate the prompts in `01-prompts/02-coding-guidelines/`. Under the new `follow-cg-execute` master orchestrator, all 39 steps are executed sequentially in bounded fast-forwarding mode (zero build/test until step 39). Therefore, maintaining 19 separate skill wrappers is redundant.

| Skill Name | Canonical Replacement | Reason for Archival |
| :--- | :--- | :--- |
| `cg-boolean-and-naming` | `follow-cg-execute` / `follow-coding-guidelines` | Subsumed by step 01 in `follow-cg-execute` |
| `cg-branch-immutability-and-clean-construction` | `follow-cg-execute` | Subsumed by step 02 in `follow-cg-execute` |
| `cg-clean-repo-build-and-caches` | `follow-cg-execute` | Subsumed by step 03 in `follow-cg-execute` |
| `cg-code-dryness-and-library-extraction` | `follow-cg-execute` | Subsumed by step 04 in `follow-cg-execute` |
| `cg-database-conventions` | `follow-cg-execute` | Subsumed by step 05 in `follow-cg-execute` |
| `cg-enum-standards` | `follow-cg-execute` | Subsumed by step 06 in `follow-cg-execute` |
| `cg-error-management` | `follow-cg-execute` | Subsumed by step 07 in `follow-cg-execute` |
| `cg-extract-types` | `follow-cg-execute` | Subsumed by step 08 in `follow-cg-execute` |
| `cg-file-size-and-function-reduction` | `follow-cg-execute` | Subsumed by step 09 in `follow-cg-execute` |
| `cg-golang-pointer-reduction-and-value-semantics` | `follow-cg-execute` | Subsumed by step 10 in `follow-cg-execute` |
| `cg-golang-pointers-and-mutation-reduction` | `follow-cg-execute` | Subsumed by step 11 in `follow-cg-execute` |
| `cg-isolate-os-tests` | `follow-cg-execute` | Subsumed by step 12 in `follow-cg-execute` |
| `cg-map-dispatch-and-handler-lookup` | `follow-cg-execute` | Subsumed by step 13 in `follow-cg-execute` |
| `cg-string-efficiency` | `follow-cg-execute` | Subsumed by step 14 in `follow-cg-execute` |
| `cg-string-normalization-and-equalfoldany` | `follow-cg-execute` | Subsumed by step 15 in `follow-cg-execute` |
| `cg-variadic-and-spread-parameters` | `follow-cg-execute` | Subsumed by step 16 in `follow-cg-execute` |
| `nested-if-and-guard-clauses` | `follow-cg-execute` | Subsumed by step 17 in `follow-cg-execute` |
| `multi-language-enums-and-traits` | `follow-cg-execute` | Subsumed by step 18 in `follow-cg-execute` |
| `lazy-regex-and-pattern-matching` | `follow-cg-execute` | Subsumed by step 19 in `follow-cg-execute` |

### Cluster B: Ideation & 100-Improvements Catalogers (8 Skills -> 1 Unified Skill)
Eight separate skills currently exist for finding 100 improvements across different surfaces.

| Skill Name | Canonical Replacement | Reason for Archival |
| :--- | :--- | :--- |
| `find-100-codebase-improvements` | `find-100-improvements` | Variant of codebase 100-improvements |
| `find-100-system-ux-improvements` | `find-100-improvements` | Variant of whole-system UX improvements |
| `find-100-ui-improvements` | `find-100-improvements` | Variant of UI-level 100-improvements |
| `next-100-code-improvements` | `find-100-improvements` | Duplicate prompt wrapper |
| `next-100-ideas-for-improvements` | `find-100-improvements` | Duplicate prompt wrapper |
| `next-100-ui-improvements` | `find-100-improvements` | Duplicate prompt wrapper |
| `next-100-whole-system-improvements` | `find-100-improvements` | Duplicate prompt wrapper |
| `next-detail-out-100-improvements` | `find-100-improvements` | Duplicate prompt wrapper |

### Cluster C: Ephemeral & Ad-Hoc Migration Skills (2 Skills -> Archive)
Single-project legacy artifacts.

| Skill Name | Canonical Replacement | Reason for Archival |
| :--- | :--- | :--- |
| `movie-cli-migration-and-optimization` | Archive | Specific to one-off `movie-cli-v8` migration |
| `ent-goal-feasibility-and-email-command-integration` | Archive | Specific to single research spike |

---

## 3. Reduction Impact Summary

- Current Skill Count: **108 skills**
- Candidate Reductions (Cluster A + B + C): **29 skills**
- Projected Lean Skill Count: **~79 skills** (over 26% reduction)
- **Preservation Invariant:** All archived skills are moved to `.agents/skills-archived/` and `.cursor/skills-archived/`, preserving git history and recoverable documentation without active context bloat.
