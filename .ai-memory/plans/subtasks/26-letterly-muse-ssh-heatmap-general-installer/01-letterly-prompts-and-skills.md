# Subtask 01: Letterly Prompts, Skills & Linter Re-Sequencing

## Objective
Re-sequence all Letterly formatters (`01-prompts/22-letterly/`, `01-prompts/23-cursor-prompts/`), update companion skills, create `letterly-git-pull-first` skill, enforce mobile single-line output, and update `check-letterly-format.py`.

## Targets
- `01-prompts/22-letterly/01-mobile-letterly.md` through `11-git-pull-first-letterly.md`
- `01-prompts/23-cursor-prompts/01-mobile-letterly-cursor.md` through `10-execute-with-release-letterly-cursor.md`
- `.agents/skills/letterly-*`
- `.cursor/skills/letterly-*`
- `linter-scripts/check-letterly-format.py`
- `linter-scripts/tests/test_check_letterly_format.py`

## Acceptance Criteria
- Item 1 is Git pull via `gitmap pull` (fallback `git pull`) with conflict resolution.
- Item 2 is Write spec under `02-spec/21-app/<slug>/`.
- Item 3 is Search codebase exclusively via GitMap.
- Item 4 is Strictly use relative Git paths.
- Mobile prompt outputs strictly in ONE continuous line with zero newlines.
- `check-letterly-format.py` passes all unit tests.
