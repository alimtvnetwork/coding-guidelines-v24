# 22-fix-repo — Versioned Repo-Name Replacer

Cross-platform tooling that rewrites prior versioned-repo-name tokens
(e.g. `coding-guidelines-v17`) to the **current** version detected
from the configured git remote. Two interchangeable implementations:

- `fix-repo.ps1` (PowerShell, Windows / pwsh-on-anything)
- `fix-repo.sh` (Bash, macOS / Linux / WSL / Git-Bash)

Both scripts share one specification, one detection algorithm, one
flag set, one exit-code contract, and one set of edge-case
guarantees. If the two implementations ever diverge, treat it as a
P1 bug and bring them back into alignment with this folder.

## Files in this folder

| File | Purpose |
|------|---------|
| `00-overview.md` | This file. Top-level map. |
| `01-spec.md` | Full normative spec — detection, flags, replacement, traversal, exit codes. |
| `02-edge-cases.md` | Enumerated edge cases with expected behavior for each. |
| `03-acceptance-criteria.md` | Checkable list used to verify the implementation. |
| `04-examples.md` | Worked CLI examples + expected stdout. |
| `plan.md` | Phased implementation plan (Phase 1–6). |

## Quick summary

1. Detect `RepoBase` and `CurrentVersion` from `git remote get-url origin`.
2. Choose target prior versions:
   - default → last 2 (`Current-2`, `Current-1`)
   - `-3` → last 3
   - `-5` → last 5
   - `-all` → every version `1` through `Current-1`
3. For every file returned by `git ls-files` that is not binary
   (NUL-byte sniff in first 8KB), replace each token
   `{RepoBase}-v{N}` with `{RepoBase}-v{Current}` literally —
   including inside URLs. Host and rest-of-path are preserved
   because they are not part of the token.
4. Print summary; exit `0` on success (including dry-run with no
   changes), non-zero on detection or write failure.

## Cross-references

- Coding guidelines (function size, booleans, error handling):
  `.lovable/coding-guidelines/coding-guidelines.md`
- URL-handling rule (saved to memory):
  `.lovable/memory/features/fix-repo-url-handling.md`
- The original verbatim brief that drove this spec was archived
  in chat history; the spec below is the binding artifact.
