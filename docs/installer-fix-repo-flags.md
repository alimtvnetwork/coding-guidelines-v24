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

## 5. Related references

- Installer behavior contract:
  [`spec/14-update/27-generic-installer-behavior.md`](../spec/14-update/27-generic-installer-behavior.md)
- `fix-repo` script spec:
  [`spec-authoring/22-fix-repo/01-spec.md`](../spec-authoring/22-fix-repo/01-spec.md)
- Installer test suite:
  [`tests/installer/`](../tests/installer/) — `check-max-fix-repo-logs-flag.sh`,
  `check-run-fix-repo-flag.sh`, `check-log-dir-flag.sh`,
  `check-show-fix-repo-log-flag.sh`, `check-log-header-env.sh`.
