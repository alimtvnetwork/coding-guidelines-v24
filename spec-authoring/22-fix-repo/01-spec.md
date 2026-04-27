# 01 — Specification (Normative)

## 1. Goal

Provide a cross-platform versioned-repo-name replacer that detects
the current repo's base name and version from the configured git
remote, and rewrites prior-version tokens to the current version
across all tracked text files in the repo.

## 2. Detection algorithm

1. **Locate repo root.**
   `git rev-parse --show-toplevel`. Non-zero exit → fail with
   `E_NOT_A_REPO` (exit code `2`).

2. **Read remote URL.** In order:
   1. `git remote get-url origin`
   2. If that fails, first line of `git remote -v` whose 3rd field
      is `(fetch)`.
   3. If both fail → `E_NO_REMOTE` (exit code `3`).

3. **Parse URL into `{Host, Owner, RepoFull}`.**
   Supported forms (all case-sensitive on `RepoFull`):

   | Form | Pattern |
   |------|---------|
   | HTTPS | `https?://<host>[:port]/<owner>/<repo>(.git)?[/...]` |
   | SSH (scp-like) | `git@<host>:<owner>/<repo>(.git)?` |
   | SSH (ssh://) | `ssh://git@<host>[:port]/<owner>/<repo>(.git)?` |

   Strip a single trailing `.git`. Anything past the repo segment
   (e.g. `/tree/main`) is ignored.

4. **Split `RepoFull` into `{RepoBase, CurrentVersion}`.**
   Match the suffix `-v(\d+)$` (case-sensitive `-v`).
   - On match → `RepoBase` = prefix, `CurrentVersion` = integer.
   - On miss → `E_NO_VERSION_SUFFIX` (exit code `4`) with a message
     pointing at this section.

5. **Validate `CurrentVersion`.**
   - `CurrentVersion >= 1` required. `0` or negative → `E_BAD_VERSION` (`5`).
   - `CurrentVersion == 1` AND mode is not `-all` → exit `0`
     immediately with "nothing to replace" (no prior versions
     exist to migrate).

## 3. Flag set (exhaustive)

Exactly one mode flag may be passed. Default = `-2`.

| PowerShell | Bash | Meaning |
|------------|------|---------|
| (none) | (none) | Replace last **2** versions |
| `-2` | `--2` | Replace last **2** versions (explicit) |
| `-3` | `--3` | Replace last **3** versions |
| `-5` | `--5` | Replace last **5** versions |
| `-all` | `--all` | Replace **every** version `1..Current-1` |
| `-DryRun` | `--dry-run` | Report changes; do not write |
| `-Verbose` | `--verbose` | Print every modified file path |

Mode flag set is **closed**: `-4`, `-6`, `-10` etc. are NOT
accepted. Passing two mode flags or an unknown flag → exit `6`
(`E_BAD_FLAG`).

`-DryRun` / `--dry-run` and `-Verbose` / `--verbose` may be
combined freely with any mode flag.

## 4. Target version computation

Let `M` = the mode integer (`2`, `3`, `5`, or `Current-1` for `-all`).

```
TargetVersions = { v ∈ ℤ | max(1, Current - M) <= v <= Current - 1 }
```

Examples (Current = 18):
- default / `-2` → `{16, 17}`
- `-3` → `{15, 16, 17}`
- `-5` → `{13, 14, 15, 16, 17}`
- `-all` → `{1, 2, …, 17}`

Examples (Current = 3):
- default / `-2` → `{1, 2}`
- `-3` → `{1, 2}`  (clamped — never below `1`)
- `-all` → `{1, 2}`

## 5. Replacement rules

1. **Token form:** the literal string `{RepoBase}-v{N}` for each
   `N ∈ TargetVersions`.
2. **Replacement:** the literal string `{RepoBase}-v{CurrentVersion}`.
3. **Match policy:** plain substring, case-sensitive. **No** word
   boundaries are required, BUT a numeric-overflow guard is
   applied: the token MUST NOT be immediately followed by a digit.
   So `coding-guidelines-v17` matches inside
   `coding-guidelines-v17/install.sh` and inside
   `https://github.com/x/coding-guidelines-v17` but does NOT match
   inside `coding-guidelines-v170`.
4. **URL handling:** plain substring everywhere. Host stays
   untouched because the host is not part of the token. Both
   modes behave identically wrt URLs (see
   `.lovable/memory/features/fix-repo-url-handling.md`).
5. **Replacement order:** ascending by `N`. Idempotent because
   the replacement string contains a different version number
   than any token being matched.

## 6. File traversal

1. Use `git ls-files -z` from repo root. This automatically:
   - honors `.gitignore`,
   - skips `.git/`,
   - excludes submodule contents (only the gitlink is listed).
2. For each path, skip if any of:
   - file is a symlink (don't follow, don't rewrite),
   - file is larger than **5 MiB** (configurable constant),
   - first 8192 bytes contain a NUL byte (binary sniff),
   - file extension is in the **always-binary** set:
     `.png .jpg .jpeg .gif .webp .ico .pdf .zip .tar .gz .tgz .bz2 .xz .7z .rar .woff .woff2 .ttf .otf .eot .mp3 .mp4 .mov .wav .ogg .webm .class .jar .so .dylib .dll .exe .pyc`.
3. Read as UTF-8 with **lossless** byte fallback (preserve original
   bytes for any sequence that fails to decode — never lose data).
4. Preserve original line endings and trailing newline presence.

## 7. Output contract

On every run print exactly this header, then per-file results when
`--verbose`, then a summary block:

```
fix-repo  base=<RepoBase>  current=v<N>  mode=<flag>
targets:  v<a>, v<b>, ...
host:     <Host>  owner=<Owner>
```

Summary:

```
scanned: <int> files
changed: <int> files (<int> replacements)
mode:    <write|dry-run>
```

## 8. Exit codes

| Code | Symbol | Meaning |
|------|--------|---------|
| `0` | OK | Success (including dry-run; including "nothing to replace") |
| `2` | E_NOT_A_REPO | `git rev-parse` failed |
| `3` | E_NO_REMOTE | No remote URL found |
| `4` | E_NO_VERSION_SUFFIX | Repo name has no `-vN` suffix |
| `5` | E_BAD_VERSION | `N <= 0` or non-integer |
| `6` | E_BAD_FLAG | Unknown / conflicting flags |
| `7` | E_WRITE_FAILED | At least one file failed to write |

## 9. Non-goals

- Rewriting URLs from one host to another.
- Touching package manifests semantically (they are text and get
  the same plain-substring treatment as everything else).
- Acting on files outside the repo working tree.
- Migrating away from non-`-vN` naming schemes.
- Creating `.bak` files or auto-staging changes — the user's
  rollback path is `git checkout -- .` (confirmed in design review).
