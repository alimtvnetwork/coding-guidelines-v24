# Component Specification: Relative Paths Mandate & Legacy Execute Skills Purge

## 1. Scope of Changes

This specification details the component-level modifications across all affected files in the repository.

---

## 2. Component Inventory & Modifications

### 2.1 AGENTS.md Section 10 Enhancement
In `AGENTS.md` under `## 10. IDE Skill Link Syntax & Prompt Formatter Invariants` -> `Execution Formatter Standards (Letterly & Cursor)`:
- Add explicit bullet mandating strict relative Git paths in Actionable Items:
  `Item 3 of Actionable Items must explicitly mandate strictly relative Git paths (02-spec/..., .ai-memory/..., cmd/...): only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well.`

### 2.2 Letterly Formatters (`01-prompts/22-letterly/`)
In all 10 Letterly formatters, inject the relative path mandate:
- **`01-mobile-letterly.md`:**
  Update rule 5 and output format to include:
  `- strictly use relative git paths only, never add absolute paths during your work (this must be respected on the release page and in release notes as well), must follow the skill [execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)`
- **`02-desktop-letterly.md`:**
  Under `# Actionable Items Must Follow Non-Negotiable`, insert:
  `Item 3 is ALWAYS: 3. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well`
  Followed by discrete technical action items.
- **`03-execute-n-steps-letterly.md` (Golden Prompt):**
  Under `# Actionable Items Must Follow Non-Negotiable`, insert:
  `Item 3 is ALWAYS: 3. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well`
  Followed by Item 4..N sequential directives.
- **`04-plan-letterly.md`:**
  Under `# Actionable Items Must Follow Non-Negotiable`, insert:
  `Item 3 is ALWAYS: 3. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well`
  Followed by Item 4 (task decomposition) and Item 5 (no-build/no-test).
- **`05-release-letterly.md`:**
  Under `# Actionable Items Must Follow Non-Negotiable`, insert:
  `Item 3 is ALWAYS: 3. Strictly use relative Git paths only; only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page, in release notes, changelog.md, manifests, and documentation`
- **`06-cicd-fix-release-letterly.md`:**
  Under `# Actionable Items Must Follow Non-Negotiable`, insert:
  `Item 3 is ALWAYS: 3. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well`
- **`07-mobile-cicd-fix-letterly.md`:**
  Update rule 4 and output format to include:
  `- strictly use relative git paths only, never add absolute paths during your work (this must be respected on the release page and in release notes as well), must follow the skill [ci-cd-fix-gitmap-release](file;.agents/skills/ci-cd-fix-gitmap-release)`
- **`08-run-letterly.md`:**
  Under `# Actionable Items Must Follow Non-Negotiable`, insert:
  `Item 4 is ALWAYS: 4. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well`
- **`09-execute-with-verification-letterly.md`:**
  Under `# Actionable Items Must Follow Non-Negotiable`, insert:
  `Item 3 is ALWAYS: 3. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well`
- **`10-execute-with-release-letterly.md`:**
  Under `# Actionable Items Must Follow Non-Negotiable`, insert:
  `Item 3 is ALWAYS: 3. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well`

### 2.3 Cursor Formatters (`01-prompts/23-cursor-prompts/`)
Mirror the exact same relative paths mandate in all 10 Cursor prompt formatters (01 through 10), substituting `(file;.cursor/skills/...)`.

### 2.4 Companion Skills (`.agents/skills/letterly-*/` and `.cursor/skills/letterly-*/`)
Synchronize all 10 Antigravity companion skills and all 10 Cursor companion skills to reflect the updated instructions and output templates.

### 2.5 Release Specifications and Management Prompts
- **`02-spec/16-generic-release/readme.md`:**
  Add a dedicated subsection under `## Shared Conventions` mandating strict relative Git paths for all release artifacts, changelogs, and manifests.
- **`02-spec/16-generic-release/07-release-metadata.md`:**
  Add section `### Strict Relative Git Paths Mandate in Release Metadata and Changelogs`.
- **`01-prompts/17-release-management/02-minor-bump.md`:**
  Add key invariant mandating strict relative Git paths across `changelog.md`, release notes, manifests, and documentation.

### 2.6 Purge of Legacy Execute Parent Task Skills
Delete the following 4 skill folders from `.agents/skills/`:
1. `.agents/skills/execute-parent-task/`
2. `.agents/skills/execute-parent-task-with-n-steps/`
3. `.agents/skills/parent-task-in-below-steps/`
4. `.agents/skills/parent-task-n-step-loop/`

Delete the following 4 skill folders from `.cursor/skills/`:
1. `.cursor/skills/execute-parent-task/`
2. `.cursor/skills/execute-parent-task-with-n-steps/`
3. `.cursor/skills/parent-task-in-below-steps/`
4. `.cursor/skills/parent-task-n-step-loop/`

Retain ONLY:
- `.agents/skills/execute-parent-task-with-n-steps-v6/`
- `.cursor/skills/execute-parent-task-with-n-steps-v6/`

### 2.7 Cross-Reference Cleanup
Update any lingering references in:
- `01-prompts/13-plan-audit/02-plan-spec-steps-v2.md`
- `.agents/skills/spec-authoring-and-validation/skill.md`
- `.cursor/skills/spec-authoring-and-validation/skill.md`
to ensure all active pointers target `execute-parent-task-with-n-steps-v6`.
