# Subtask 02: Spec 21 Sanitization & Guidelines Realignment

## Task Details
- **Parent Task:** `25-coding-guidelines-and-folder-structure-instructions`
- **Subtask Code:** `Subtask-02`
- **Slug:** `02-spec-21-sanitization`
- **Status:** PENDING
- **Plan Motto:** Improved work and less friction

## Owned Files & Scope
- `02-spec/21-app/` directory contents in `coding-guidelines`
- `02-spec/21-app/readme.md` (to update as container documentation)

## Actionable Steps
1. [ ] Verify that all coding guidelines rules represented in `02-spec/21-app/` exist in `02-spec/02-coding-guidelines/`, `01-prompts/15-cg-execute/`, and `.agents/skills/`.
2. [ ] Delete the 44 misplaced milestone subdirectories from `02-spec/21-app/`.
3. [ ] Delete the 7 misplaced root markdown files from `02-spec/21-app/`.
4. [ ] Preserve `02-spec/21-app/.gitkeep`.
5. [ ] Rewrite `02-spec/21-app/readme.md` with explicit instructions that `21-app` is strictly for application domain specifications in projects that build software applications, never for meta-guidelines or prompts.
6. [ ] Keep `02-spec/21-app/25-coding-guidelines-and-folder-structure-instructions/` intact as the active task specification container.

## Acceptance Criteria
- `gitmap lf 02-spec/21-app` contains strictly `readme.md`, `.gitkeep`, and `25-coding-guidelines-and-folder-structure-instructions/`.
- No orphan guidelines or temporary execution files remain in `02-spec/21-app/`.
