# Subtask 02: Letterly Prompts Cleanup and Text Compaction

- **Parent Task:** `letterly-prompt-cleanup-and-fleet-sync`
- **Scope:** Clean `01-prompts/22-letterly/03-execute-n-steps-letterly.md` and related formatters, removing step budgets and Muse mentions, reducing text lining and filler words.
- **Status:** DONE

## Execution Summary & Evidence

1. **`01-prompts/22-letterly/03-execute-n-steps-letterly.md`:**
   - Removed Step Budget Mandate (`Rule 9`).
   - Removed 400-step lifecycle budget bullet from `## Additional Instructions`.
   - Stripped verbal filler ("the", redundant phrasing) and compacted instructions from 49 lines down to 43 lines.
   - Retained strict delegation to `[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)`.
2. **Companion Letterly Formatter Cleanup:**
   - `01-mobile-letterly.md`: Removed 400-step lifecycle budget from continuous single-line template and directives.
   - `02-desktop-letterly.md`: Removed step budget mandate and additional instructions budget bullet.
   - `04-plan-letterly.md`: Removed step budget mandate and additional instructions budget bullet.
   - `05-release-letterly.md`: Removed step budget mandate and additional instructions budget bullet.
   - `06-cicd-fix-release-letterly.md`: Removed step budget mandate and additional instructions budget bullet.
   - `07-mobile-cicd-fix-letterly.md`: Removed step budget mandate and compacted single-line template.
   - `08-run-letterly.md`: Removed step budget mandate and additional instructions budget bullet.
   - `09-execute-with-verification-letterly.md`: Removed step budget mandate and additional instructions budget bullet.
   - `10-execute-with-release-letterly.md`: Removed step budget mandate and additional instructions budget bullet.
   - `11-git-pull-first-letterly.md`: Removed step budget mandate and additional instructions budget bullet.
   - `readme.md`: Updated Rule 6 to declare decoupled execution engine invariant.
