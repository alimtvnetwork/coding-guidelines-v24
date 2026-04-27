# Plan — 23-visibility-change

Phased delivery. Each `next` from the user advances one phase.
Implementation does not begin until Phase 2.

## Phase 1 — Spec + plan authoring  ✅

- [x] `00-overview.md`
- [x] `01-spec.md`
- [x] `02-edge-cases.md`
- [x] `03-acceptance-criteria.md`
- [x] `04-examples.md`
- [x] `plan.md` (this file)
- [x] Open questions resolved (auth backend, default, runner, confirm).

Outputs are documents only. No source code touched yet.

## Phase 2 — `visibility-change.ps1` (PowerShell)

- [ ] Repo-root file `visibility-change.ps1`.
- [ ] Param block: `[string]$Visible`, `[switch]$Yes`, `[switch]$DryRun`, `[switch]$Help`.
- [ ] Helper functions (each ≤ 15 lines, positive booleans):
      `Get-OriginUrl`, `Resolve-Provider`, `Resolve-OwnerRepo`,
      `Get-CurrentVisibility`, `Resolve-Target`, `Confirm-PublicChange`,
      `Invoke-VisibilityChange`, `Test-Verified`, `Show-Help`.
- [ ] Exit codes per `01-spec.md §8`.
- [ ] No swallowed errors; every external call checks `$LASTEXITCODE`.

## Phase 3 — `visibility-change.sh` (Bash)

- [ ] Repo-root file `visibility-change.sh`, `chmod +x`.
- [ ] `#!/usr/bin/env bash`, `set -euo pipefail`.
- [ ] Same flag surface, same exit codes, same one-line success output.
- [ ] Bash helpers (≤ 15 lines each), parity with PS function names.

## Phase 4 — Runner wiring ✅

- [x] `run.ps1`: added `visibility` to switch dispatch; forwards `$args` to `.\visibility-change.ps1` via `Invoke-Visibility`.
- [x] `run.sh`:  added `visibility` case; forwards `"$@"` to `./visibility-change.sh` via `invoke_visibility`.
- [x] Updated help tables in both runners (visibility command + flag forwarding block).

## Phase 5 — Tests + sync ✅

- [x] Unit tests added under `tests/installer/`:
      - `check-visibility-runner-wiring.sh` — help advertises sub-command, dispatch branches present in both runners.
      - `check-visibility-arg-parsing.sh` — exit codes 0/2/6 across 6 cases (help, unknown flag, bad `--visible`, no-repo with/without forced target).
      - `check-visibility-provider-detect.sh` — `resolve_provider` / `resolve_owner_repo` across 10 URL shapes (HTTPS, SCP-SSH, SSH-URL, with/without `.git`, GitHub, gitlab.com, self-hosted allow-listed via `VISIBILITY_GITLAB_HOSTS`, denied self-hosted, Bitbucket rejected, empty URL).
- [x] All 3 test scripts pass locally.
- [x] `npm run sync` clean (version 4.24.0, spec tree 614 files / 85 folders / 28 top-level entries).
- [x] `node scripts/sync-spec-tree.mjs` regenerated `src/data/specTree.json`; `spec-authoring/` is tracked as a top-level node (consistent with `22-fix-repo`; leaf folders are not individually indexed by design).
