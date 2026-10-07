# Subtask: Script & Pooling Enhancements

## Metadata
- **Subtask ID:** `01-script-and-pooling-enhancements`
- **Parent Plan:** `.ai-memory/plans/fleet-commit-and-push.md`
- **Spec Reference:** `02-spec/21-app/fleet-commit-and-push/01-architecture-spec.md`
- **Status:** Pending

## Objective
Update `03-ai-scripts/49-commit-and-push-all-repos.py` to:
1. Exclude non-owned repositories: `oh-my-zsh`, `ohmyzsh`, `zsh`, `omis`, `oh-my-posh`, `dotfiles`, etc.
2. Add pooling / pulling support: before staging dirty changes, execute `git pull origin <branch> --no-rebase` (with `--no-pull` flag if skipped).
3. Validate upstream push: confirm remote tracking push, verify remote SHA matches local HEAD.
4. Enforce "No Push = Not Done": if changes cannot be pushed to remote, exit with non-zero status code and flag task as incomplete.

## Acceptance Criteria
- [ ] Non-owned repositories filtered out in discovery.
- [ ] Pulling / pooling integrated prior to staging.
- [ ] Push verification confirms remote synchronization.
- [ ] Self-tests pass via `python 03-ai-scripts/49-commit-and-push-all-repos.py --self-test`.
