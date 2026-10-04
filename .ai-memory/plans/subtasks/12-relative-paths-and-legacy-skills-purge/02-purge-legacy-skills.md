# Subtask [02]: Purge Legacy Execute Skills & Clean Repository Cross-References

Traceability ID: Task-02
Spec Reference: [02-spec/21-app/12-relative-paths-and-legacy-skills-purge/01-architecture-spec.md](../../../../02-spec/21-app/12-relative-paths-and-legacy-skills-purge/01-architecture-spec.md)
Target Files:
- `.agents/skills/execute-parent-task/`
- `.agents/skills/execute-parent-task-with-n-steps/`
- `.agents/skills/parent-task-in-below-steps/`
- `.agents/skills/parent-task-n-step-loop/`
- `.cursor/skills/execute-parent-task/`
- `.cursor/skills/execute-parent-task-with-n-steps/`
- `.cursor/skills/parent-task-in-below-steps/`
- `.cursor/skills/parent-task-n-step-loop/`
- `01-prompts/13-plan-audit/02-plan-spec-steps-v2.md`
- `.agents/skills/spec-authoring-and-validation/skill.md`
- `.cursor/skills/spec-authoring-and-validation/skill.md`

Action:
1. Delete the 4 legacy skill directories from `.agents/skills/`.
2. Delete the 4 legacy skill directories from `.cursor/skills/`.
3. Update lingering pointers in `02-plan-spec-steps-v2.md`, `.agents/skills/spec-authoring-and-validation/skill.md`, and `.cursor/skills/spec-authoring-and-validation/skill.md` to reference `execute-parent-task-with-n-steps-v6`.

Acceptance Criteria:
- No legacy execute parent task skills remain in `.agents/skills/` or `.cursor/skills/` except `execute-parent-task-with-n-steps-v6`.
- Zero broken skill references exist across repository files.
