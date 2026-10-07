# Subtask: Fleet Sync & Verification

## Metadata
- **Subtask ID:** `03-fleet-sync-and-verification`
- **Parent Plan:** `.ai-memory/plans/fleet-commit-and-push.md`
- **Spec Reference:** `02-spec/21-app/fleet-commit-and-push/01-architecture-spec.md`
- **Status:** Pending

## Objective
Update the canonical commit-and-push prompt and skills:
1. `01-prompts/09-commit-and-multi-agent-code-fix/09-commit-and-push-all-repos.md`
2. `.agents/skills/commit-and-push-all-repos/skill.md`
3. `.cursor/skills/commit-and-push-all-repos/skill.md`

Commit changes locally in `coding-guidelines` using GitMap:
`gitmap cpf "fleet - enforce commit push completion and non-owned repo exclusions"`

Run multi-repository sync across the 43 connected repositories using `03-ai-scripts/38-sync-prompts-skills-scripts.py`.

## Acceptance Criteria
- [ ] Prompt and skills updated with pooling, non-owned repo exclusion, and 'No Push = Not Done' invariant.
- [ ] Coding-guidelines committed and pushed cleanly.
- [ ] Multi-repo sync script executed with pre-pull and backup branches.
- [ ] All 43 connected repositories successfully synchronized.
