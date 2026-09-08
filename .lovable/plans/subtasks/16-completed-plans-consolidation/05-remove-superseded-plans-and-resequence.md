# Subtask 05: Remove Superseded Plans and Resequence

> **Plan:** `16-completed-plans-consolidation`  
> **Status:** Completed  
> **Target File:** `.lovable/plans/completed/`, `.lovable/plans/01-index.md`  

---

## Intent

1. Cleanly remove the 31 superseded individual micro-plan files from `.lovable/plans/completed/` via `git rm`.
2. Ensure the remaining consolidated milestones are named strictly monotonically from `01-` to `06-`.
3. Update `.lovable/plans/01-index.md` to list the 6 consolidated milestones.

## Verification

- `ls .lovable/plans/completed/` shows exactly 6 files: `01-` to `06-`.
- `git status` shows all old files staged for deletion and new milestones staged for addition.
