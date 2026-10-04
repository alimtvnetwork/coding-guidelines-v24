# Master Execution Plan: Relative Paths Mandate & Legacy Execute Skills Purge

Spec Reference: [02-spec/21-app/12-relative-paths-and-legacy-skills-purge/01-architecture-spec.md](../../../02-spec/21-app/12-relative-paths-and-legacy-skills-purge/01-architecture-spec.md)
Component Spec: [02-spec/21-app/12-relative-paths-and-legacy-skills-purge/02-component-spec.md](../../../02-spec/21-app/12-relative-paths-and-legacy-skills-purge/02-component-spec.md)

## User Request (Verbatim)

```text
In all these latterly prompts, can you please add one more line to only add the relative paths, never add the absolute path during your work. Make sure you include that, and also this should be respected in the release page as well. Mention that. Okay. I hope you can respect this and make a final improvements on all the prompts and also seeing the skills and make sure that the skill, like the execute parent task in n steps. So this one, all the previous versions, I want you to remove from the skills. Remember that and confirm that this is really applied properly.
```

## Actionable Deliverables Breakdown

| ID | Title | Assigned Role | Target Scope | Status |
| :--- | :--- | :--- | :--- | :--- |
| `Task-01` | Author Architecture & Component Specs & Plan Files | Lead Architect | `02-spec/21-app/12-*/`, `.ai-memory/plans/` | `DONE` |
| `Task-02` | Purge Legacy Execute Skills & Clean Repo Cross-References | Worker 01 | `.agents/skills/`, `.cursor/skills/`, references | `DONE` |
| `Task-03` | Inject Strict Relative Git Paths Mandate into Letterly & Cursor Prompts | Worker 02 | `AGENTS.md`, `01-prompts/22-letterly/`, `01-prompts/23-cursor-prompts/` | `DONE` |
| `Task-04` | Synchronize All Letterly & Cursor Companion Skills with Relative Paths | Worker 03 | `.agents/skills/letterly-*/`, `.cursor/skills/letterly-*/` | `DONE` |
| `Task-05` | Update Release Specs and Management Prompts with Relative Paths Mandate | Worker 04 | `02-spec/16-generic-release/`, `01-prompts/17-release-management/` | `DONE` |
| `Task-06` | Lint Verification, Atomic GitMap Commit, and Minor Release Ceremony | Release Orchestrator | Linters, GitMap cpf, tag `v6.73.0` | `DONE` |

## Subtask Mapping
- [Task-01 Subtasks](subtasks/12-relative-paths-and-legacy-skills-purge/01-relative-paths-prompts.md)
- [Task-02 Subtasks](subtasks/12-relative-paths-and-legacy-skills-purge/02-purge-legacy-skills.md)
- [Task-03 Subtasks](subtasks/12-relative-paths-and-legacy-skills-purge/03-release-specs-and-verification.md)
