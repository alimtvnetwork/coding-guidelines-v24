# 03 — Acceptance Criteria

A correct implementation satisfies every item below. Items are
phrased as binary checks suitable for a manual or automated
verification pass.

## Files exist

- [ ] `spec-authoring/22-fix-repo/00-overview.md`
- [ ] `spec-authoring/22-fix-repo/01-spec.md`
- [ ] `spec-authoring/22-fix-repo/02-edge-cases.md`
- [ ] `spec-authoring/22-fix-repo/03-acceptance-criteria.md`
- [ ] `spec-authoring/22-fix-repo/04-examples.md`
- [ ] `spec-authoring/22-fix-repo/plan.md`
- [ ] `fix-repo.ps1` at repo root
- [ ] `fix-repo.sh` at repo root, `chmod +x`

## Detection

- [ ] Reads `git remote get-url origin` and parses HTTPS, SSH-scp, and `ssh://` forms.
- [ ] Strips trailing `.git`.
- [ ] Extracts `RepoBase` and `CurrentVersion` from a `-vN` suffix.
- [ ] Exits with the documented code (2/3/4/5) for each detection failure.

## Flags

- [ ] No flag → mode `-2`.
- [ ] `-2`, `-3`, `-5`, `-all` accepted; `-4`, `-6`, etc. rejected with exit `6`.
- [ ] `-DryRun` (PS) / `--dry-run` (sh) suppress writes.
- [ ] `-Verbose` (PS) / `--verbose` (sh) print per-file lines.
- [ ] Two mode flags → exit `6`.

## Replacement

- [ ] Token = `{RepoBase}-v{N}`. Replacement = `{RepoBase}-v{Current}`.
- [ ] Case-sensitive `-v`.
- [ ] Numeric-overflow guard: `coding-guidelines-v170` is NOT touched.
- [ ] URL occurrences ARE rewritten (host preserved).
- [ ] Same content writes are NOT issued (only files that would actually change get touched).
- [ ] Idempotent: a second run after the first changes 0 files.

## Traversal

- [ ] Uses `git ls-files -z` from repo root.
- [ ] Skips symlinks, files >5 MiB, binary-extension files, NUL-byte-sniffed files.
- [ ] Honors `.gitignore` automatically (because of `git ls-files`).

## Code quality (ties to coding-guidelines)

- [ ] Every function ≤ 8 effective lines (CODE-RED-005).
- [ ] No nested `if`. Guard-and-return style.
- [ ] Boolean variables / functions prefixed `Is` / `Has` (PS) or `is_` / `has_` (sh) where applicable.
- [ ] No swallowed errors. Every failed read/write is logged AND reflected in the exit code.
- [ ] Constants (token prefix, byte-sniff limit, max file size, exit codes) are named — no magic numbers in branching code.
- [ ] Files ≤ 100 lines (split helpers into sibling files / functions if necessary).

## Docs

- [ ] `readme.md` install / setup section references both scripts.
- [ ] At least one `--dry-run` example shown for each script.
