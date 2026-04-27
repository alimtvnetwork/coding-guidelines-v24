# Installer post-fix-repo flags — log retention & rollback

This doc covers the installer flags that govern what happens **after**
`fix-repo` runs (whether triggered automatically by `--run-fix-repo` or
implicitly by an opt-in flow):

- `--max-fix-repo-logs N` (env: `INSTALL_MAX_FIX_REPO_LOGS`)
- `--rollback-on-fix-repo-failure`
- `--full-rollback`
- `--show-fix-repo-log`
- `--log-dir DIR` (env: `INSTALL_LOG_DIR`)

All flags are honoured by every installer in the repo: `install.{sh,ps1}`,
`release-install.{sh,ps1}`, and the bundle installers
(`cli-install.*`, `consolidated-install.*`, `error-manage-install.*`,
`linters-install.*`, `slides-install.*`, `splitdb-install.*`,
`wp-install.*`).

---

## 1. Where logs live

Every `fix-repo` invocation writes a timestamped log:

```
<DEST>/.install-logs/fix-repo-YYYYMMDDTHHMMSSZ.log
```

The directory is overridable with `--log-dir DIR` (or env
`INSTALL_LOG_DIR`). Relative paths are resolved against `<DEST>`;
absolute paths are used as-is.

The log header captures: start time, script path, dest, OS, shell,
`uname -a`, cwd, and the rollback flag values for the run. The footer
records the exit code and finish time.

---

## 1a. Glossary — flag ↔ env-var ↔ precedence

The table below is the **single source of truth** for how each setting is named in Bash, in PowerShell, and as an environment variable — plus what wins when more than one is set. Read this first; the rest of the doc assumes these mappings.

| Setting (semantic)                  | Bash flag                          | PowerShell flag               | Env var                       | Type / accepted values                | Default     | Precedence (highest → lowest) |
|-------------------------------------|------------------------------------|-------------------------------|-------------------------------|----------------------------------------|-------------|-------------------------------|
| Log directory                       | `--log-dir DIR`                    | `-LogDir DIR`                 | `INSTALL_LOG_DIR`             | path (absolute or relative to `<DEST>`) | `<DEST>/.install-logs` | CLI flag → env var → default |
| Max retained `fix-repo` logs        | `--max-fix-repo-logs N`            | `-MaxFixRepoLogs N`           | `INSTALL_MAX_FIX_REPO_LOGS`   | non-negative integer (`0` = disabled)  | `0` (keep all) | CLI flag → env var → default |
| Auto-run `fix-repo` after install   | `--run-fix-repo`                   | `-RunFixRepo`                 | —                             | boolean switch                         | `false`     | CLI flag only                 |
| Print log to stdout after run       | `--show-fix-repo-log`              | `-ShowFixRepoLog`             | —                             | boolean switch                         | `false`     | CLI flag only                 |
| Revert `fix-repo` edits on failure  | `--rollback-on-fix-repo-failure`   | `-RollbackOnFixRepoFailure`   | —                             | boolean switch                         | `false`     | CLI flag only                 |
| Full rollback (edits + new files)   | `--full-rollback`                  | `-FullRollback`               | —                             | boolean switch (implies the row above) | `false`     | CLI flag only                 |
| Install destination                 | `--dest DIR`                       | `-Dest DIR`                   | —                             | path                                   | `./` or per-bundle | CLI flag only           |

### Precedence rules in plain English

1. **CLI flag always wins.** If you pass `--max-fix-repo-logs 10` *and* export `INSTALL_MAX_FIX_REPO_LOGS=99`, the installer uses `10`. The env var is consulted **only when the CLI flag is absent**.
2. **Env vars only exist for the two value-typed settings** above (`INSTALL_LOG_DIR`, `INSTALL_MAX_FIX_REPO_LOGS`). Boolean switches like `--run-fix-repo` or `--rollback-on-fix-repo-failure` have **no env-var fallback** — pass them on the command line every time. This is intentional: a stale exported env var should never silently arm rollback or auto-execute `fix-repo`.
3. **Empty string ≠ unset.** `INSTALL_LOG_DIR=""` is treated as unset (the installer falls back to the default). To force an empty value, you must explicitly pass `--log-dir ""` (which is then treated as relative path → `<DEST>/`).
4. **Negative or non-integer `--max-fix-repo-logs` values do not crash the run.** They produce a `Log pruning: SKIPPED (… is not a non-negative integer)` line and pruning is bypassed for that run (see §5.2).
5. **Naming convention** — env var names are derived mechanically: `INSTALL_` + the flag in `SCREAMING_SNAKE_CASE` with hyphens replaced by underscores. This makes them grep-able and predictable, e.g. `--max-fix-repo-logs` ↔ `INSTALL_MAX_FIX_REPO_LOGS`.

### Worked precedence example

```bash
export INSTALL_LOG_DIR=/var/log/install
export INSTALL_MAX_FIX_REPO_LOGS=50

./install.sh \
  --run-fix-repo \
  --max-fix-repo-logs 10           # CLI wins → keep 10, ignore env's 50
                                   # --log-dir not passed → env wins → /var/log/install
```

Effective settings for that run:

```
log dir:               /var/log/install        (source: env INSTALL_LOG_DIR)
max-fix-repo-logs:     10                      (source: CLI --max-fix-repo-logs)
run-fix-repo:          true                    (source: CLI --run-fix-repo)
rollback-on-failure:   false                   (source: default — no env fallback exists)
```

---

## 1b. Verifying CLI-vs-env precedence (copy-paste recipes)

These recipes are **assertion-style smoke tests** — each one sets a deliberately conflicting env var, runs the installer in `--dry-run` mode, and `grep`s the output to confirm the CLI flag won. Safe to run in any working directory; nothing is installed (`--dry-run` short-circuits all writes), and `--run-fix-repo` is **not** passed, so no `fix-repo` execution or rollback can occur.

> ✅ Each block exits **0** when the precedence rule holds; non-zero means the installer regressed and would silently honour the env var.

### A. `--log-dir` overrides `INSTALL_LOG_DIR`

**🐧 macOS · Linux · Bash**

```bash
INSTALL_LOG_DIR=/tmp/from-env \
  ./install.sh --dry-run --log-dir /tmp/from-cli 2>&1 \
  | tee /tmp/precedence-log-dir.out \
  | grep -E 'log[- ]dir.*=.*from-cli' \
  && ! grep -q 'from-env' /tmp/precedence-log-dir.out \
  && echo "✅ --log-dir wins over INSTALL_LOG_DIR" \
  || { echo "❌ precedence regressed"; exit 1; }
```

**🪟 Windows · PowerShell**

```powershell
$env:INSTALL_LOG_DIR = 'C:\Temp\from-env'
$out = .\install.ps1 -DryRun -LogDir 'C:\Temp\from-cli' 2>&1 | Out-String
$cliWon = $out -match 'from-cli'
$envLeaked = $out -match 'from-env'
if ($cliWon -and -not $envLeaked) {
    Write-Host '✅ -LogDir wins over $env:INSTALL_LOG_DIR' -ForegroundColor Green
} else {
    Write-Host '❌ precedence regressed' -ForegroundColor Red; exit 1
}
Remove-Item Env:INSTALL_LOG_DIR
```

### B. `--max-fix-repo-logs` overrides `INSTALL_MAX_FIX_REPO_LOGS`

This one runs the installer for real (otherwise the pruning step is skipped) but in a **throwaway temp directory** so nothing on your machine is touched. The log line we grep is the deterministic `Log pruning: --max-fix-repo-logs=…` summary added in §2.

**🐧 macOS · Linux · Bash**

```bash
TMP="$(mktemp -d)"; trap 'rm -rf "$TMP"' EXIT
( cd "$TMP" && git init -q && git commit --allow-empty -m baseline -q )

INSTALL_MAX_FIX_REPO_LOGS=99 \
  ./install.sh \
    --dest "$TMP" \
    --run-fix-repo \
    --max-fix-repo-logs 7 \
    --log-dir "$TMP/.install-logs" \
    2>&1 \
  | tee "$TMP/precedence-max.out" \
  | grep -E 'Log pruning:.*--max-fix-repo-logs=7\b' \
  && ! grep -qE 'max-fix-repo-logs=99\b' "$TMP/precedence-max.out" \
  && echo "✅ --max-fix-repo-logs wins over INSTALL_MAX_FIX_REPO_LOGS" \
  || { echo "❌ precedence regressed"; exit 1; }
```

**🪟 Windows · PowerShell**

```powershell
$tmp = Join-Path $env:TEMP ("install-precedence-" + [guid]::NewGuid())
New-Item -ItemType Directory -Path $tmp | Out-Null
Push-Location $tmp; git init -q; git commit --allow-empty -m baseline -q; Pop-Location

$env:INSTALL_MAX_FIX_REPO_LOGS = '99'
$out = .\install.ps1 `
  -Dest $tmp `
  -RunFixRepo `
  -MaxFixRepoLogs 7 `
  -LogDir (Join-Path $tmp '.install-logs') 2>&1 | Out-String

$cliWon    = $out -match 'Log pruning:.*-MaxFixRepoLogs=7\b|--max-fix-repo-logs=7\b'
$envLeaked = $out -match 'max-fix-repo-logs=99\b|MaxFixRepoLogs=99\b'
Remove-Item Env:INSTALL_MAX_FIX_REPO_LOGS
Remove-Item -LiteralPath $tmp -Recurse -Force

if ($cliWon -and -not $envLeaked) {
    Write-Host '✅ -MaxFixRepoLogs wins over $env:INSTALL_MAX_FIX_REPO_LOGS' -ForegroundColor Green
} else {
    Write-Host '❌ precedence regressed' -ForegroundColor Red; exit 1
}
```

### C. Sanity-check the env-var fallback (no CLI flag passed)

When the CLI flag is **absent**, the env var must be picked up. This confirms the fallback half of the precedence chain.

**🐧 macOS · Linux · Bash**

```bash
INSTALL_LOG_DIR=/tmp/from-env \
  ./install.sh --dry-run 2>&1 \
  | grep -E 'log[- ]dir.*=.*from-env' \
  && echo "✅ INSTALL_LOG_DIR honoured when --log-dir omitted" \
  || { echo "❌ env fallback broken"; exit 1; }
```

**🪟 Windows · PowerShell**

```powershell
$env:INSTALL_LOG_DIR = 'C:\Temp\from-env'
$out = .\install.ps1 -DryRun 2>&1 | Out-String
Remove-Item Env:INSTALL_LOG_DIR
if ($out -match 'from-env') {
    Write-Host '✅ $env:INSTALL_LOG_DIR honoured when -LogDir omitted' -ForegroundColor Green
} else {
    Write-Host '❌ env fallback broken' -ForegroundColor Red; exit 1
}
```

> 💡 If any recipe fails, the installer is no longer honouring the precedence rules from §1a. File an issue and attach the captured `*.out` file (Bash) or `$out` (PowerShell) — both contain the full installer banner that names the source of every effective setting.

---

## 2. `--max-fix-repo-logs N` (log retention)

After each `fix-repo` run, the installer prunes
`fix-repo-*.log` files in the log directory, keeping only the **N
newest**. `N` must be a non-negative integer.

| Value      | Behaviour                                                       |
|------------|------------------------------------------------------------------|
| unset / 0  | Pruning **disabled** — every log is kept forever.                |
| `N ≥ 1`    | Keep the newest `N` logs; older ones are deleted.                |
| negative   | Treated as invalid → pruning **skipped** with a clear warning.   |

Equivalent flag dialects:

```bash
# Bash
./install.sh --run-fix-repo --max-fix-repo-logs 5
INSTALL_MAX_FIX_REPO_LOGS=5 ./install.sh --run-fix-repo
```

```powershell
# PowerShell
.\install.ps1 -RunFixRepo -MaxFixRepoLogs 5
$env:INSTALL_MAX_FIX_REPO_LOGS = '5'; .\install.ps1 -RunFixRepo
```

The CLI flag wins over the env var when both are set.

### Output you'll see

```
  ▸ Log pruning: --max-fix-repo-logs=5 | found=8 kept=5 pruned=3 dir=/repo/.install-logs
```

Other states are surfaced explicitly so CI logs are unambiguous:

```
  ▸ Log pruning: DISABLED (--max-fix-repo-logs=0)
  ▸ Log pruning: SKIPPED (log dir not found: …; --max-fix-repo-logs=5)
  ▸ Log pruning: SKIPPED (--max-fix-repo-logs=-1 is not a non-negative integer)
```

---

## 3. Rollback flags

Two opt-in flags control what happens when `fix-repo` exits non-zero:

| Flag (Bash)                          | Flag (PowerShell)             | Effect                                                                                  |
|--------------------------------------|-------------------------------|------------------------------------------------------------------------------------------|
| `--rollback-on-fix-repo-failure`     | `-RollbackOnFixRepoFailure`   | On failure: `git -C <DEST> checkout -- .` (revert tracked-file edits made by fix-repo). |
| `--full-rollback`                    | `-FullRollback`               | Implies the above **and** removes files this install run created + restores backups.    |

`--full-rollback` is a **superset** — it implicitly enables
`--rollback-on-fix-repo-failure`. Specifying both is harmless.

When rollback is armed, the installer prints:

```
Rollback armed: HEAD=<sha>, full-rollback=on
```

### Decision matrix on `fix-repo` outcome

| `fix-repo` exit | rollback flag   | Installer prints                                              | Installer exit |
|-----------------|-----------------|---------------------------------------------------------------|----------------|
| `0`             | any             | `Rollback: not needed (fix-repo succeeded; flags: …)`         | `0`            |
| non-zero        | none            | `Rollback: NOT TRIGGERED (--rollback-on-fix-repo-failure=false  --full-rollback=false)` | `5` |
| non-zero        | on-failure only | `═══ ROLLBACK TRIGGERED (fix-repo failed) ═══` → tracked files reverted | `5` |
| non-zero        | full            | Same as above **plus** new files removed and overwritten files restored from backup | `5` |

The decision line always names the flag values, so a `grep` on the log
tells you exactly which behaviour was selected.

### Pre-conditions

- Rollback only works when `<DEST>` is a git repo. If it isn't, the
  installer prints a warning and silently disables the flag — it does
  **not** fail the install.
- `--full-rollback` relies on bookkeeping populated during the install
  step; running it against a `--dry-run` install is a no-op.

---

## 4. Interaction between pruning and rollback

Pruning runs **after** `fix-repo` finishes but **before** any rollback
decision, so:

1. The newly-written log of the failing run is **always preserved** —
   it's the newest file and therefore inside the keep-window.
2. A rollback won't undo log pruning. If you need forensic history,
   set `--max-fix-repo-logs 0` (the default).
3. If you set `--max-fix-repo-logs 1`, only the failing run's log
   survives — the previous successful runs are deleted before the
   rollback runs. Combine with `--show-fix-repo-log` to dump it to
   stdout for CI capture.

A safe CI recipe:

```bash
./install.sh \
  --run-fix-repo \
  --rollback-on-fix-repo-failure \
  --max-fix-repo-logs 20 \
  --show-fix-repo-log \
  --log-dir "$RUNNER_TEMP/install-logs"
```

```powershell
.\install.ps1 `
  -RunFixRepo `
  -RollbackOnFixRepoFailure `
  -MaxFixRepoLogs 20 `
  -ShowFixRepoLog `
  -LogDir "$env:RUNNER_TEMP\install-logs"
```

---

## 4a. GitHub Actions snippet (drop-in)

Runs the installer with `--run-fix-repo`, prunes to the newest 20 logs, and uploads the entire log directory as a build artifact named `fix-repo-logs-*` — even when the job fails, so you can post-mortem the failing run.

````yaml
# .github/workflows/install-and-fix-repo.yml
name: Install + fix-repo

on:
  push:
    branches: [main]
  workflow_dispatch:

permissions:
  contents: read

jobs:
  install-and-fix-repo:
    name: Install and run fix-repo (with log pruning)
    runs-on: ubuntu-latest
    env:
      INSTALL_LOG_DIR: ${{ runner.temp }}/install-logs
      INSTALL_MAX_FIX_REPO_LOGS: '20'
    steps:
      - name: Checkout
        uses: actions/checkout@v6

      - name: Run installer + fix-repo (Bash)
        if: runner.os != 'Windows'
        run: |
          set -euo pipefail
          curl -fsSL https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v18/main/install.sh \
            | bash -s -- \
                --run-fix-repo \
                --rollback-on-fix-repo-failure \
                --max-fix-repo-logs "$INSTALL_MAX_FIX_REPO_LOGS" \
                --show-fix-repo-log \
                --log-dir "$INSTALL_LOG_DIR"

      - name: Run installer + fix-repo (PowerShell)
        if: runner.os == 'Windows'
        shell: pwsh
        run: |
          $ErrorActionPreference = 'Stop'
          irm https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v18/main/install.ps1 `
            -OutFile install.ps1
          .\install.ps1 `
            -RunFixRepo `
            -RollbackOnFixRepoFailure `
            -MaxFixRepoLogs ([int]$env:INSTALL_MAX_FIX_REPO_LOGS) `
            -ShowFixRepoLog `
            -LogDir $env:INSTALL_LOG_DIR

      - name: Upload fix-repo logs
        if: always()        # capture logs on success AND failure
        uses: actions/upload-artifact@v4
        with:
          name: fix-repo-logs-${{ runner.os }}-${{ github.run_id }}
          path: ${{ runner.temp }}/install-logs/fix-repo-*.log
          if-no-files-found: warn
          retention-days: 14
````

### What this snippet guarantees

- **Reproducible log location** — `INSTALL_LOG_DIR` is set once and reused by both the installer and the upload step, so the artifact path can never drift.
- **Bounded disk use** — `--max-fix-repo-logs 20` keeps history short on long-running self-hosted runners; on hosted runners (fresh VM each run) only the current run's log exists, which still uploads cleanly.
- **Always-uploaded forensics** — `if: always()` runs the upload step even after a non-zero installer exit, and pruning runs **before** rollback (see §4) so the failing run's log is guaranteed to survive into the artifact.
- **Cross-platform parity** — same flag semantics in Bash and PowerShell; add `windows-latest` to a matrix without changing anything else.

---

## 5. Troubleshooting

Each failure case below lists the **trigger**, the **exact log line(s) to grep for**, and the **fix**. All log lines are emitted on stdout *and* written into the per-run `fix-repo-*.log` header — so a failed CI run's uploaded artifact is enough to diagnose every case.

### 5.1 Destination is not a git repo (rollback silently disabled)

**Trigger:** you passed `--rollback-on-fix-repo-failure` (or `--full-rollback`) but `<DEST>` has no `.git/` directory — e.g. you installed into a fresh folder created by `--dest ./pristine`.

**What you'll see:**

```
⚠️  --rollback-on-fix-repo-failure: ./pristine is not a git repo; rollback disabled.
```

The installer **does not fail** — rollback is the opt-in safety net, and the absence of git means there's nothing to revert to. You will *not* see `Rollback armed: HEAD=…` later in the log; instead, on a `fix-repo` failure you'll see:

```
Rollback: NOT TRIGGERED (--rollback-on-fix-repo-failure=false  --full-rollback=false)
```

(Both flags read `false` because the installer flipped them internally after the git-repo check failed.)

**Fix — pick one:**

- Run the installer against an already-cloned repo (`git clone … && ./install.sh --dest <repo>`).
- `git init && git add -A && git commit -m baseline` inside `<DEST>` *before* the install so there's a `HEAD` to roll back to.
- Drop the rollback flags if you genuinely want a non-git install (e.g. extracting into a Docker image layer).

---

### 5.2 Invalid `--max-fix-repo-logs` value

**Trigger:** you passed a non-integer (`--max-fix-repo-logs abc`) or a negative number (`--max-fix-repo-logs -1`), or set `INSTALL_MAX_FIX_REPO_LOGS` to such a value.

**What you'll see (Bash):**

```
  ▸ Log pruning: SKIPPED (--max-fix-repo-logs=-1 is not a non-negative integer)
```

For non-numeric input, the Bash installer rejects it at flag-parse time with:

```
❌ --max-fix-repo-logs requires a non-negative integer (got: abc)
```

— and exits 1 *before* installing anything. PowerShell's `[int]` parameter binding rejects non-numeric input with PS's standard `ParameterBindingException`; negative ints reach the installer body and produce the same `SKIPPED` line.

**Important:** even when pruning is skipped, the `fix-repo` step still runs normally. The only consequence is that **no logs are deleted** — older `fix-repo-*.log` files accumulate.

**Fix:** pass a non-negative integer (`0` to disable, `N ≥ 1` to keep that many). The CLI flag overrides any env var, so `--max-fix-repo-logs 10` will win over a stale `INSTALL_MAX_FIX_REPO_LOGS=-1` in your shell.

---

### 5.3 `--full-rollback` combined with `--dry-run` (no-op rollback)

**Trigger:** you passed `--dry-run` (PS: `-DryRun`) **and** `--full-rollback`. Because `--dry-run` short-circuits every file mutation, the install bookkeeping arrays (`INSTALLED_NEW`, `INSTALLED_BACKUPS`) stay empty, and the `fix-repo` auto-run step is skipped entirely:

```
if ! $DRY_RUN && $RUN_FIX_REPO; then run_fix_repo; fi
```

**What you'll see:** the dry-run summary at the bottom prints

```
⚠️  DRY-RUN — no changes written
```

…and there is **no** `Rollback armed: HEAD=…` line, **no** `fix-repo-*.log` file, and **no** `═══ ROLLBACK TRIGGERED ═══` line, no matter what fails. If you were expecting rollback to "preview" what it would restore, you won't get it — by design, dry-run never writes the snapshot a rollback would need.

**Fix — pick one based on what you actually want:**

| You want…                                         | Run this                                                                                  |
|---------------------------------------------------|--------------------------------------------------------------------------------------------|
| Preview which files the install would touch       | `./install.sh --dry-run` (drop `--full-rollback`; it's irrelevant here)                    |
| Real install + rollback safety net                | `./install.sh --run-fix-repo --full-rollback` (omit `--dry-run`)                           |
| Smoke-test rollback end-to-end before production  | Run the real install into a throwaway dir: `./install.sh --dest /tmp/smoke --run-fix-repo --full-rollback` |

---

### 5.4 Quick-grep cheat sheet

When you have an uploaded `fix-repo-*.log` artifact, these greps answer "what happened?" in one pass:

```bash
# Did rollback fire? Which mode?
grep -E '^(═══ ROLLBACK TRIGGERED|Rollback: (NOT TRIGGERED|not needed|armed|complete))' fix-repo-*.log

# Did pruning run? Which mode? How many removed?
grep -E '^  ▸ Log pruning:' fix-repo-*.log

# Why was rollback disabled?
grep -E 'is not a git repo; rollback disabled' fix-repo-*.log

# fix-repo final exit code
grep -E '^# exit:' fix-repo-*.log
```

---

## 6. Related references

- Installer behavior contract:
  [`spec/14-update/27-generic-installer-behavior.md`](../spec/14-update/27-generic-installer-behavior.md)
- `fix-repo` script spec:
  [`spec-authoring/22-fix-repo/01-spec.md`](../spec-authoring/22-fix-repo/01-spec.md)
- Installer test suite:
  [`tests/installer/`](../tests/installer/) — `check-max-fix-repo-logs-flag.sh`,
  `check-run-fix-repo-flag.sh`, `check-log-dir-flag.sh`,
  `check-show-fix-repo-log-flag.sh`, `check-log-header-env.sh`.
