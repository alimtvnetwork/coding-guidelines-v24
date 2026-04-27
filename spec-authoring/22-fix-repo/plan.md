# Plan — 22-fix-repo

Phased delivery. Each `next` from the user advances one phase.
Implementation does not begin until Phase 2.

## Phase 1 — Spec + plan authoring  ✅

- [x] `00-overview.md`
- [x] `01-spec.md`
- [x] `02-edge-cases.md`
- [x] `03-acceptance-criteria.md`
- [x] `04-examples.md`
- [x] `plan.md` (this file)
- [x] Memory rule saved: `.lovable/memory/features/fix-repo-url-handling.md`
- [x] Open questions resolved (URL handling, flag set, case + scope, backups).

Outputs are documents only. No source code touched yet.

## Phase 2 — `fix-repo.ps1` (PowerShell)

- [ ] Repo-root file `fix-repo.ps1`.
- [ ] Param block: `[switch]$All`, `[switch]$DryRun`, `[switch]$Verbose`,
      and discrete `[switch]$P2 / $P3 / $P5` aliased to `-2 / -3 / -5`
      (PowerShell can't accept literal `-2` as a parameter name without
      help — implement via `args` parsing if cleaner; see notes).
- [ ] Helper functions (each ≤ 8 lines, positive booleans):
      `Get-RemoteUrl`, `Resolve-RepoIdentity`, `Get-TargetVersions`,
      `Test-IsBinaryFile`, `Test-IsSkippedPath`, `Invoke-FileRewrite`,
      `Write-Summary`.
- [ ] Exit codes per `01-spec.md §8`.
- [ ] No swallowed errors — every `try` has a `catch` that logs + sets the failure flag.
- [ ] Splits into sibling `.ps1` files under `scripts/fix-repo/` if total exceeds 100 lines (per file-size rule).

## Phase 3 — `fix-repo.sh` (Bash)

- [ ] Repo-root file `fix-repo.sh`, `chmod +x`, `#!/usr/bin/env bash`, `set -euo pipefail`.
- [ ] Same flag surface, same exit codes, same outputs.
- [ ] Helpers as small bash functions (≤ 8 lines each).
- [ ] Use `git ls-files -z` + `while IFS= read -r -d ''` to handle paths with spaces.
- [ ] NUL-byte sniff via `head -c 8192 | grep -qzv $'\\0'` style check.
- [ ] Splits into `scripts/fix-repo/*.sh` if needed.

## Phase 4 — One-time v16/v17 → v18 sweep  ✅

- [x] v17 already migrated last session (541 occurrences, 176 files).
- [x] Confirmed zero remaining `coding-guidelines-v(16|17)` matches via `rg`.
- [x] No `npm run sync` needed — no derived files changed (zero rewrites this phase).

## Phase 5 — Install / setup docs  ✅

- [x] Added "Repo version migration — fix-repo" subsection to `readme.md` under Full-Repo Install Scripts.
- [x] 6 examples shown across both shells (default, `-3`/`--3` dry-run, `-all`/`--all` verbose).
- [x] Cross-linked to `spec-authoring/22-fix-repo/01-spec.md`.
- [x] `npm run sync` skipped — readme is a source file, no derived mirrors affected.

## Phase 6 — Verification + dry-run capture  ✅

- [x] Fixture created at `linters-cicd/tests/fixtures/fix-repo/README.md`.
- [x] Functional matrix passed: default / `--3` / `--all` / dry-run all behave per spec.
- [x] Error matrix passed: `--bogus` → exit 6, no-git-repo → exit 2.
- [x] Three real bugs found and fixed (`.git` strip, null-byte detector, perl dependency); details in `05-verification-log.md`.
- [x] Test suite: 151 passed + 1 pre-existing baseline failure — no regressions.
- [x] All phases ✅ in this file.

## Notes / risks

- PowerShell parameter parsing of bare numeric flags (`-3`, `-5`) is
  awkward. If the natural `[switch]$3` form is rejected by the parser,
  fall back to splatting `$args` and detecting `'-2' / '-3' / '-5' / '-all'`
  manually before the `param()` block consumes them. Keep the user-visible
  surface (`.\fix-repo.ps1 -3`) unchanged regardless of the
  implementation strategy.
- `release-artifacts/coding-guidelines-vX.Y.Z/` snapshots will be
  rewritten on every run because their internal text mentions the
  versioned token. This matches the user's prior decision (Apr 2026)
  to include release artifacts in the rename. To opt out later, add
  `release-artifacts/` to `.gitattributes` with an `export-ignore`-style
  marker and teach the script to honor it; out of scope for this
  initial delivery.
