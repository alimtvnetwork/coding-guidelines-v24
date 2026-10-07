# Subtask: V6 Prompt & Skill Invariants ('No Push = Not Done')

## Metadata
- **Subtask ID:** `02-v6-prompt-and-skill-invariants`
- **Parent Plan:** `.ai-memory/plans/fleet-commit-and-push.md`
- **Spec Reference:** `02-spec/21-app/fleet-commit-and-push/01-architecture-spec.md`
- **Status:** Pending

## Objective
Embed the mandatory completion agreement into:
1. `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`
2. `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`
3. `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`

## Required Agreement Invariant
"**MANDATORY COMPLETION INVARIANT (NO PUSH = NOT DONE):** If the code is not committed to Git and pushed upstream to GitHub (main/master/tracking branch), the task is strictly considered INCOMPLETE and NOT DONE. Leaving uncommitted dirty changes or unpushed commits means the execution is unfinished and has failed. A task cannot be marked completed or successful until all changes are committed and confirmed pushed to GitHub."

## Acceptance Criteria
- [ ] Section 12, 14, and 15 in `02-execute-parent-task-with-n-steps-v6.md` updated with this invariant.
- [ ] `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` updated.
- [ ] `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md` updated.
- [ ] Zero absolute paths or `file:///` URIs.
