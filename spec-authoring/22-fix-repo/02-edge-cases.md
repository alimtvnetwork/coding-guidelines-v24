# 02 — Edge Cases

Each row lists a situation the script MUST handle and the
expected behavior. Implementation tests SHOULD cover each row.

| # | Situation | Expected behavior |
|---|-----------|-------------------|
| 1 | No git remote configured | Exit `3` (`E_NO_REMOTE`), no files touched. |
| 2 | Multiple remotes (`origin`, `upstream`) | Use `origin`. If absent, first `(fetch)` row. |
| 3 | SSH remote (`git@github.com:owner/repo-v9.git`) | Parsed identically to HTTPS. `RepoBase=repo`, `Current=9`. |
| 4 | `ssh://` form with port | Port stripped during parse; same result. |
| 5 | Self-hosted git (`https://git.acme.internal/x/y-v3`) | Works. Host is opaque. |
| 6 | Repo name without `-vN` | Exit `4` (`E_NO_VERSION_SUFFIX`), no files touched. |
| 7 | Repo name `-v0` | Exit `5` (`E_BAD_VERSION`). |
| 8 | Repo name `-v01` (leading zero) | Parses as `1`. Replaces as `v1` form too. Output uses canonical `vN` (no leading zeros). |
| 9 | `Current == 1`, default mode | Print "nothing to replace", exit `0`. |
| 10 | `Current == 1`, `-all` | Same — empty target set, exit `0`. |
| 11 | `-3` when `Current == 2` | Targets clamped to `{1}`. |
| 12 | CRLF line endings | Preserved byte-for-byte. |
| 13 | UTF-8 BOM at file start | Preserved. |
| 14 | Mixed line endings inside one file | Preserved exactly. |
| 15 | Symlink to a tracked text file | Skipped (don't follow, don't rewrite). |
| 16 | File larger than 5 MiB | Skipped. Logged in verbose mode. |
| 17 | Binary file by extension | Skipped. |
| 18 | Binary file by NUL-byte sniff | Skipped even if extension would otherwise pass. |
| 19 | File inside `.git/` | Never enumerated (`git ls-files` excludes it). |
| 20 | Submodule contents | Never enumerated (`git ls-files` lists only the gitlink). |
| 21 | `-V17` (uppercase V) inside text | NOT replaced (case-sensitive). |
| 22 | Token immediately followed by digit (`v170`) | NOT replaced (numeric-overflow guard). |
| 23 | Token immediately followed by letter (`v17a`) | Replaced — letter is fine, only trailing digit is the foot-gun. |
| 24 | Token inside a URL | Replaced (identical to plain text). Host preserved. |
| 25 | Same token appearing N times in one file | All replaced, counted in summary. |
| 26 | File whose only change would be inside a string literal in code | Replaced — script is text-level, has no language awareness. **This is intentional**; users requested it. |
| 27 | Read-only file | Write attempt fails → exit `7`, but only AFTER finishing other files. Failures are reported, not silently swallowed. |
| 28 | `--dry-run --verbose` | Prints every would-change file + summary, no writes, exit `0`. |
| 29 | Two mode flags (`-2 -3`) | Exit `6` (`E_BAD_FLAG`) before scanning. |
| 30 | Run inside a worktree | Treated as a regular repo (`git rev-parse` resolves). |
| 31 | Run with a detached HEAD | No effect — script doesn't care about HEAD. |
| 32 | Token appears in the script's own source | Replaced (intentional — the script's docs may carry the token). To suppress, exclude in caller via `.gitattributes` or `.gitignore`. |
| 33 | NFC vs NFD unicode in `RepoBase` | Compared as raw bytes, no normalization. Matches what git stores. |
