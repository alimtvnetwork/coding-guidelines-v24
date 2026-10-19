# Plan: GitMap RC Heatmap Integration and Agent Task Management

- **Task Slug:** `gitmap-rc-heatmap-integration-and-agent-task-management`
- **Spec:** `02-spec/21-app/gitmap-rc-heatmap-integration-and-agent-task-management/01-architecture-spec.md`
- **Goal:** Enhance execute v6, letterly, cd-execute, and muse prompts with comprehensive GitMap RC Heatmap options, task deduplication/existence reporting, prompt MD5 hash verification, roadmap linkage, and skill synchronization across the repository, concluding with a minor version bump release.

## Execution Subtasks

1. [x] **Subtask 01:** Architecture spec in `02-spec/21-app/gitmap-rc-heatmap-integration-and-agent-task-management/` and plan setup.
2. [ ] **Subtask 02:** Enhance `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` with RC heatmap commands, auto-creation of RC structure, task existence/resume diagnostics, prompt MD5 hash verification, and roadmap loop overview.
3. [ ] **Subtask 03:** Update Letterly (`01-prompts/22-letterly/09-execute-with-verification-letterly.md`, `03-execute-n-steps-letterly.md`), CD/CG-execute, and Muse prompts with RC heatmap and agent continuation commands.
4. [ ] **Subtask 04:** Synchronize skills across `.agents/skills/` and `.cursor/skills/` (`execute-parent-task-with-n-steps-v6`, `gitmap`, `muse-master-prompt`).
5. [ ] **Subtask 05:** Verification, regression check, and minor version bump release ceremony (`v6.81.0`).
