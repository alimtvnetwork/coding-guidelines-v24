# Subtask [03]: Release Specs & Companion Skills Synchronization and Verification

Traceability ID: Task-04, Task-05, Task-06
Spec Reference: [02-spec/21-app/12-relative-paths-and-legacy-skills-purge/01-architecture-spec.md](../../../../02-spec/21-app/12-relative-paths-and-legacy-skills-purge/01-architecture-spec.md)
Target Files:
- `.agents/skills/letterly-*/` (all 10 skills)
- `.cursor/skills/letterly-*/` (all 10 skills)
- `02-spec/16-generic-release/readme.md`
- `02-spec/16-generic-release/07-release-metadata.md`
- `01-prompts/17-release-management/02-minor-bump.md`
- `changelog.md`
- `version.json`

Action:
1. Synchronize all 20 companion skills with the updated relative paths mandate.
2. Update release specifications and minor bump prompt to mandate relative paths in changelog.md and release pages.
3. Run linters (`check-relative-paths.py`, `check-prompts-loaded.py`, `check-forbidden-strings.py`).
4. Commit atomically via GitMap cpf and conduct minor release ceremony (`v6.73.0`).

Acceptance Criteria:
- All companion skills reflect relative paths mandate.
- Release specs enforce relative paths in changelogs.
- Linters exit 0.
- Clean release tag `v6.73.0` published.
