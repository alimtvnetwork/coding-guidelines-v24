# 02 — Edge Cases

| # | Scenario | Behavior |
|---|---|---|
| 1 | Run outside a git repo | Exit `2`, message `not a git work tree`. |
| 2 | No `origin` remote | Exit `3`, message `no origin remote configured`. |
| 3 | `origin` points to Bitbucket / Gitea / unknown | Exit `4`, name the host. |
| 4 | `gh` / `glab` not installed | Exit `5`, print install URL for the relevant CLI. |
| 5 | CLI installed but not authenticated | Exit `5`, print `gh auth login` / `glab auth login` hint. |
| 6 | Repo is GitLab `internal` (third state) | Treat as `private` for toggle math; explicit `--visible pub` still works. |
| 7 | `--visible pub` when already public | Exit `0`, print `already public`. |
| 8 | `--visible` with garbage value | Exit `6`, list accepted values. |
| 9 | `private → public` in a non-interactive shell, no `--yes` | Exit `7`, hint to pass `--yes`. |
| 10 | API call succeeds but verification re-read shows unchanged | Exit `8` (provider replication lag or silent failure). |
| 11 | Repo URL uses SSH form `git@github.com:owner/repo.git` | Parse owner/repo from the path; same code path as HTTPS. |
| 12 | `origin` URL has trailing `.git` | Strip before passing to `gh`/`glab`. |
| 13 | Self-hosted GitLab on custom domain | Honored if listed in `VISIBILITY_GITLAB_HOSTS` env (comma-separated). |
| 14 | User passes both `-Visible pub` and `-Visible pri` (PS array) | PowerShell binds last value; documented but not special-cased. |
| 15 | Flag passed to `run.ps1 visibility` that the standalone doesn't recognize | Standalone's exit `6` propagates through the runner unchanged. |
| 16 | `--dry-run` with `--visible` matching current state | Print `[dry-run] already <state>`, exit `0`. |
| 17 | `gh` prompts for `--accept-visibility-change-consequences` | Always pass that flag automatically; never let the underlying CLI prompt. |
| 18 | Network failure mid-call | CLI exits non-zero → we surface stderr + exit `5`. |
