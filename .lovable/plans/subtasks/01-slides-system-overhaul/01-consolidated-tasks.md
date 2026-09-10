# Subtask: Slides System Overhaul & Repo Audit (Consolidated)

> **Parent Plan:** `.lovable/plans/pending/02-slides-system-overhaul.md`  
> **Status:** IN_PROGRESS  

## 1. Scope & Execution Ledger

| Step | Scope | Description | Status |
|:---:|---|---|:---:|
| 1 | Repo-Wide Audit | Audit guidelines SSOT drift, clean up duplicate rules in `.cursorrules` and specs | PENDING |
| 2 | Slides 70-Task Backlog | Modernize slides UI, navigation, accessibility, and presenter notes | PENDING |
| 3 | Slides Distribution | Package slides zip into releases and link in readme | COMPLETED |
| 4 | Guideline Sync CI | Script `scripts/sync-guidelines.mjs` and CI drift prevention | COMPLETED |

## 2. Core Implementation Requirements

- Preserve all cross-references to `02-spec/21-app/` and `.lovable/coding-guidelines.md`.
- Enforce strict relative Git paths across all slide decks and scripts.
- Verify function sizing <= 15 lines and implicit booleans across modified scripts.
