# Subtask 03: Verify, Commit & Deploy Fleet Synchronization

> **/goal** Verify relative path hygiene, execute atomic GitMap commit, and deploy non-destructively across all 43 connected repositories via `38-sync-prompts-skills-scripts.py`.
> **/learn** Ensure zero repo-owned custom files are deleted (`-0`), and verify clean git status on all repositories.

## Canonical Reference

- Specification: [`02-spec/21-app/18-gitmap-agent-task-engine-and-python-sqlite-purge/02-acceptance-criteria.md`](../../../../02-spec/21-app/18-gitmap-agent-task-engine-and-python-sqlite-purge/02-acceptance-criteria.md)

## Action Checklist

- [ ] Run `python linter-scripts/check-relative-paths.py` and confirm 0 errors.
- [ ] Run `gitmap aum search "46-agent" 01-prompts` and confirm 0 hits.
- [ ] Stage and commit changes in `coding-guidelines` via `gitmap cpf "agent-task - replace legacy python sqlite with native gitmap agent commands"` and push to `origin/main`.
- [ ] Execute `python 03-ai-scripts/38-sync-prompts-skills-scripts.py` across all 43 connected repositories.
- [ ] Confirm non-destructive deployment: zero deletions of repo-owned custom files (`-0`), all repos cleanly committed and pushed.
