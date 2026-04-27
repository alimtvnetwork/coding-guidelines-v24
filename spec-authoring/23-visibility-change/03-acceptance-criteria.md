# 03 — Acceptance Criteria

A feature is accepted when **all** of the following are true.

## Functional

- [ ] `visibility-change.ps1` and `visibility-change.sh` exist at repo root, both executable.
- [ ] `--help` / `-Help` prints usage and exits `0` on both.
- [ ] No-flag invocation in a GitHub repo toggles visibility (verified by re-reading via `gh repo view`).
- [ ] No-flag invocation in a GitLab repo toggles visibility (verified by re-reading via `glab repo view`).
- [ ] `-Visible pub` / `--visible pub` forces public; `-Visible pri` / `--visible pri` forces private.
- [ ] `--dry-run` performs **zero** mutating API calls (verified by `gh --debug` log inspection).
- [ ] `private → public` without `--yes` in a TTY prompts; declining aborts with non-zero exit.
- [ ] `private → public` without `--yes` with stdin piped exits `7` immediately.
- [ ] All exit codes in `01-spec.md §8` reachable via fixtures.

## Runner Integration

- [ ] `./run.sh visibility --help` prints the same help as `./visibility-change.sh --help`.
- [ ] `.\run.ps1 visibility -Visible pub` is byte-equivalent in effect to `.\visibility-change.ps1 -Visible pub`.
- [ ] `./run.sh help` lists the new `visibility` sub-command.
- [ ] `.\run.ps1 help` lists the new `visibility` sub-command.

## Code Quality

- [ ] Both standalone scripts ≤ 300 lines.
- [ ] Every function ≤ 15 lines.
- [ ] Zero nested conditionals (validated by grep for `if.*\n.*if` patterns in review).
- [ ] All booleans use positive names.
- [ ] No swallowed errors — every `try`/`catch` (or `||`) logs and exits.
- [ ] Linter (`linter-scripts/run.sh` + `run.ps1`) passes against both scripts.

## Sync

- [ ] `npm run sync` runs clean after the change.
- [ ] `spec/` (canonical mirror of `spec-authoring/`) updated if the project's sync workflow promotes 23-visibility-change.
- [ ] Health-score JSON regenerated.
