# 01 — Spec (Normative)

## 1. Files

| Path | Role |
|---|---|
| `visibility-change.ps1` | Repo-root standalone script (PowerShell) |
| `visibility-change.sh`  | Repo-root standalone script (Bash) |
| `run.ps1` (existing)    | Adds `visibility` sub-command, forwards flags |
| `run.sh`  (existing)    | Adds `visibility` sub-command, forwards flags |

Both standalone scripts live at the repo root for parity with
`fix-repo.ps1` / `fix-repo.sh`.

## 2. Flag Surface

Identical between PowerShell and Bash; only the sigil differs.

| PowerShell | Bash | Meaning |
|---|---|---|
| `-Visible pub` | `--visible pub` | Force visibility to **public** |
| `-Visible pri` | `--visible pri` | Force visibility to **private** |
| (none) | (none) | **Toggle** current visibility |
| `-Yes` | `--yes` / `-y` | Skip the `private → public` confirmation prompt |
| `-DryRun` | `--dry-run` | Print what would change; perform no API calls |
| `-Help` / `-h` | `--help` / `-h` | Print help and exit `0` |

`-Visible` accepted values (case-insensitive): `pub`, `public`, `pri`,
`private`. Anything else → exit `6` (Bad flag).

## 3. Provider Detection

1. `git remote get-url origin` → URL string.
2. Match host against:
   - `github.com` or `ssh.github.com` → **GitHub**
   - `gitlab.com` or any host whose path contains `/gitlab/` API
     reachable, or matched by an explicit allow-list env
     `VISIBILITY_GITLAB_HOSTS` (comma-separated) → **GitLab**
3. No match → exit `4` (Unsupported provider) with a one-line message
   naming the detected host.

## 4. Auth Backend

- GitHub → invoke `gh repo edit <owner>/<repo> --visibility public|private --accept-visibility-change-consequences`.
- GitLab → invoke `glab repo edit <owner>/<repo> --visibility public|private`.
- Current visibility read via `gh repo view --json visibility -q .visibility`
  / `glab repo view -F json` and parsed.
- If the CLI is not on `PATH` → exit `5` with the install URL for that CLI.
- If the CLI returns an auth error → re-print its stderr and exit `5`.
  The user is expected to run `gh auth login` / `glab auth login` (which
  themselves handle the browser flow). We **do not** wrap that flow.

## 5. Toggle Logic

```
current = read_current_visibility()
target  = (
    args.Visible if args.Visible
    else "private" if current == "public"
    else "public"
)
if current == target: print "already <target>", exit 0   # idempotent no-op
if target == "public" and not args.Yes: confirm_or_exit_7()
if args.DryRun: print "would change <current> → <target>", exit 0
apply(target)
verify(target)            # re-read and assert
print "changed <current> → <target>"
exit 0
```

## 6. Confirmation Prompt

Triggered **only** when `current=private` and `target=public` and `--yes`
not set. Prompt text:

```
⚠  About to make <owner>/<repo> PUBLIC on <provider>.
   URL: <html_url>
   Type 'yes' to continue, anything else aborts:
```

Non-interactive stdin (e.g. piped) → exit `7` (Confirmation required) with
hint to pass `--yes`.

## 7. Runner Integration

`run.ps1 visibility [flags]` and `./run.sh visibility [flags]` forward all
flags verbatim to the standalone script. The runner help table gains:

```
visibility   toggle GitHub/GitLab repo visibility (--visible pub|pri)
```

The standalone script is the source of truth; the runner is a thin
dispatcher (≤ 5 lines added per runner).

## 8. Exit Codes

| Code | Meaning |
|---|---|
| 0 | Success (changed, already-target, or dry-run printed) |
| 2 | Not inside a Git work tree |
| 3 | No `origin` remote configured |
| 4 | Unsupported provider (host not GitHub/GitLab) |
| 5 | Auth/CLI failure (missing CLI, not logged in, API error) |
| 6 | Bad flag value |
| 7 | Confirmation required but stdin not interactive |
| 8 | Verification failed (apply returned 0 but visibility unchanged) |

## 9. Output Contract

- Single human-readable line on success: `visibility: <old> → <new> (<provider>)`.
- Errors → stderr, exit non-zero. No stack traces; one line per error.
- `--dry-run` prefixes with `[dry-run]`.
- Help text is plain ASCII, ≤ 80 cols.

## 10. CODE RED Compliance

- Each script ≤ 300 lines total; functions 8–15 lines.
- Zero nested conditionals — guard-clause style.
- All booleans positively named (`HasOrigin`, `IsGitHub`, `IsDryRun`).
- Errors never swallowed: every `try` has a `catch` that logs + exits.
- Max 2 boolean operands per expression.
