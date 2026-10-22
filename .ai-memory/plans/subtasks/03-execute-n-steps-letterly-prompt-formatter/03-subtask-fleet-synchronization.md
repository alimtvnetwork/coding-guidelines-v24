# Subtask 03: Fleet Synchronization & Verification

## Task Details
- **Parent Task:** `03-execute-n-steps-letterly-prompt-formatter`
- **Subtask Code:** `Subtask-03`
- **Slug:** `03-fleet-synchronization`
- **Status:** PENDING
- **Plan Motto:** Improved work or less friction, faster work and more pure and effective work

## Owned Files & Scope
- `03-ai-scripts/38-sync-prompts-skills-scripts.py`
- All 43 connected repositories in workspace

## Actionable Steps
1. [ ] Run `python linter-scripts/check-prompts-loaded.py` and `python linter-scripts/check-forbidden-strings.py` in `coding-guidelines`.
2. [ ] Run `python 03-ai-scripts/38-sync-prompts-skills-scripts.py` across all 43 repositories.
3. [ ] Verify that `v4` and `02-spec/21-*` remain strictly excluded across all synchronized repositories.
4. [ ] Stage, commit via `gitmap cpc`, and push all updated repositories to their upstream tracking branches.
5. [ ] Perform final atomic commit in `coding-guidelines`: `gitmap cpf "letterly - update prompts and skills to gold standard and synchronize fleet"`.

## Acceptance Criteria
- 100% of target repositories synchronized and pushed upstream.
- Zero drift and zero lint errors.
