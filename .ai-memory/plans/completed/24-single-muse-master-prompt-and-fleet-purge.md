# Execution Plan: Single Muse Master Prompt & Fleet-Wide Deprecated Prompt Purge

> Plan ID: `24-single-muse-master-prompt-and-fleet-purge`
> Specification: `02-spec/21-app/19-single-muse-master-prompt-and-fleet-purge/`
> Status: Completed
> Execution Engine: V6 Parameter-Driven Engine (`N=300`, `A=2`, `H=2`, `C=30`)

---

## 1. Objectives & Scope

1. **Delete Redundant Prompts**: Permanently delete `02-muse-execute-in-a-step.md`, `03-muse-master-with-repos.md`, `04-muse-master-with-repos-v2.md`, and `05-muse-master-prompt-v2.md` from `01-prompts/27-muse-prompts/`.
2. **Single Canonical Master Prompt**: Maintain only `01-muse-master-prompt.md` and `readme.md` in `01-prompts/27-muse-prompts/`.
3. **Reference Harmonization**: Update all references in `01-prompts/readme.md`, `.ai-memory/prompts.md`, `.agents/skills/muse-master-prompt/skill.md`, `.cursor/skills/muse-master-prompt/skill.md`, `01-mastery-bootstrap-prompt.md`, and related specs.
4. **Automated Fleet Purge**: Enhance `03-ai-scripts/38-sync-prompts-skills-scripts.py` and `.agents/scripts/38-sync-prompts-skills-scripts.py` to automatically unlink deprecated Muse prompts across all 43 repositories.
5. **Full Fleet Deployment**: Run synchronization across all 43 connected repositories, ensuring all target repositories have the deprecated prompts purged and are committed and pushed upstream.

---

## 2. Checklist & Quality Gates

- [x] Subtask 1: Delete deprecated Muse prompts (`02`, `03`, `04`, `05`) and update `01-muse-master-prompt.md` and `readme.md`.
- [x] Subtask 2: Update skills and references across codebase (`muse-master-prompt` skills, `01-prompts/readme.md`, `.ai-memory/prompts.md`, specs).
- [x] Subtask 3: Update `38-sync-prompts-skills-scripts.py` with automatic unlinking of deprecated Muse prompts.
- [x] Subtask 4: Author specifications in `02-spec/21-app/19-single-muse-master-prompt-and-fleet-purge/` and register in app spec index.
- [x] Subtask 5: Regenerate sync-managed assets (`npm run sync:fix`).
- [x] Subtask 6: Run linters (relative paths, forbidden strings, secret scan).
- [x] Subtask 7: Commit and push changes in `coding-guidelines`.
- [x] Subtask 8: Execute fleet-wide synchronization across all 43 connected repositories and verify completion.
