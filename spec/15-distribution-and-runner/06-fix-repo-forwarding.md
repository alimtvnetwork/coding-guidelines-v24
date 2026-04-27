# Runner Contract — `fix-repo` Sub-command Forwarding

**Version:** 1.0.0
**Updated:** 2026-04-27
**Status:** Normative
**Companion to:** [02-runner-contract.md](./02-runner-contract.md)
**Inner scripts:** `fix-repo.sh`, `fix-repo.ps1`
**Inner spec:** [`spec-authoring/22-fix-repo/01-spec.md`](../../spec-authoring/22-fix-repo/01-spec.md)

---

## 1. Purpose

Define — exactly once, in one place — how the repo-root runners (`run.sh`,
`run.ps1`) MUST forward arguments to the inner `fix-repo` scripts. This
document is the single source of truth for:

- the user-visible CLI surface of `./run.{sh,ps1} fix-repo …`
- the flag mapping between Bash and PowerShell invocations
- the argument-forwarding semantics (verbatim, no re-quoting, no mutation)
- the conformance tests that gate every PR

Any divergence between `run.sh` and `run.ps1` for `fix-repo` is a **bug**.
Any drift between this document and the actual code is a **bug**.

---

## 2. Sub-command surface

| Invocation                              | Effect                                                       |
|-----------------------------------------|--------------------------------------------------------------|
| `./run.sh fix-repo` / `./run.ps1 fix-repo` | Default mode (`--2` / `-2`). Rewrite last 2 prior versions. |
| `./run.sh fix-repo --3`                 | Rewrite last 3 prior versions.                              |
| `./run.sh fix-repo --5`                 | Rewrite last 5 prior versions.                              |
| `./run.sh fix-repo --all`               | Rewrite every prior version (`v1`..`v(current-1)`).         |
| `./run.sh fix-repo --dry-run`           | Report changes; do not write.                                |
| `./run.sh fix-repo --verbose`           | List every modified file.                                    |

PowerShell uses the same flag tokens (`--2`, `--3`, `--5`, `--all`,
`--dry-run`, `--verbose`). Long-form `--` flags MUST be accepted on both
shells. Single-dash short forms (`-2`, `-DryRun`, etc.) are **not** part
of the contract — `fix-repo.ps1` parses the same `--`-prefixed tokens
as `fix-repo.sh`.

---

## 3. Flag mapping (Bash ↔ PowerShell)

| User types (Bash)   | User types (PowerShell) | Forwarded to inner script as |
|---------------------|-------------------------|------------------------------|
| `--2`               | `--2`                   | `--2`                        |
| `--3`               | `--3`                   | `--3`                        |
| `--5`               | `--5`                   | `--5`                        |
| `--all`             | `--all`                 | `--all`                      |
| `--dry-run`         | `--dry-run`             | `--dry-run`                  |
| `--verbose`         | `--verbose`             | `--verbose`                  |
| (none)              | (none)                  | (none — inner script defaults to `--2`) |

There is **no flag translation**. The runner is a transparent dispatcher.
If a flag is not recognized by the inner script, the inner script emits
the error — never the runner.

---

## 4. Argument-forwarding contract

The runner MUST forward every argument **byte-for-byte** to the inner
script. Specifically:

1. **No re-quoting.** Use the shell's native argv-array forwarding
   (`"$@"` in Bash, `@args` in PowerShell). Never `eval`, never
   `printf %q`, never join-then-split.
2. **No mutation.** The runner MUST NOT inject, drop, reorder, or
   transform any argument the user typed after the `fix-repo`
   sub-command token.
3. **No intermediate copies.** The dispatch path SHOULD be a single
   `exec` (Bash) or `&` invocation (PowerShell) on the original argv.
   Function indirection is allowed only when the function does not
   touch the argv (i.e. it just forwards `"$@"` / `@args` unchanged).
4. **Quoted arguments preserved.** Spaces, glob characters, dashes,
   and embedded equals signs MUST survive forwarding intact.
5. **Exit code propagation.** The runner MUST exit with the inner
   script's exit code. Bash uses `exec`; PowerShell uses `exit
   $LASTEXITCODE`.

### 4.1 Bash reference implementation

```bash
fix-repo)
  _assert_fix_repo_present
  shift
  exec bash "$SCRIPT_DIR/fix-repo.sh" "$@"
  ;;
```

`exec` is required: it replaces the runner process with the inner
script, guaranteeing the exit code is the inner script's exit code and
that no further re-quoting can occur.

### 4.2 PowerShell reference implementation

```powershell
function Invoke-FixRepo {
    $inner = Join-Path $PSScriptRoot "fix-repo.ps1"
    if (-not (Test-Path $inner)) { Write-Host "❌ Cannot find $inner" -ForegroundColor Red; exit 1 }
    & $inner @args
    exit $LASTEXITCODE
}

# dispatch:
"fix-repo" { Invoke-FixRepo @args }
```

`@args` (splatting) preserves argv as an array — equivalent to Bash's
`"$@"`. Both `Invoke-FixRepo @args` (call site) and `& $inner @args`
(inside the function) MUST splat — never interpolate into a string.

---

## 5. Help text

Both runners MUST advertise `fix-repo` in their help output (`./run.sh
help`, `./run.ps1 help`) with at least:

```
fix-repo     rewrite prior versioned-repo-name tokens to current

Fix-repo flags forwarded to fix-repo.{sh,ps1}:
  --2 | --3 | --5 | --all   how many prior versions to rewrite (default: --2)
  --dry-run                 report changes; do not write
  --verbose                 list every modified file
```

The help text source-of-truth lives in `scripts/runner-help.txt`
(Bash) and `scripts/runner-help.ps.txt` (PowerShell). Both files MUST
list the same flag set in the same order.

---

## 6. Conformance tests

Every PR that touches `run.sh`, `run.ps1`, `fix-repo.sh`, `fix-repo.ps1`,
or this document MUST keep the following tests green:

| Test                                               | Asserts                                                   |
|----------------------------------------------------|-----------------------------------------------------------|
| `tests/installer/check-fix-repo-runner-wiring.sh`  | Both runners advertise and dispatch the `fix-repo` token. |
| `tests/installer/check-fix-repo-arg-forwarding.sh` | `run.sh` forwards every flag verbatim (shim-based).       |
| `tests/installer/check-fix-repo-url-rewrite.sh`    | End-to-end inner-script behavior (URL rewrite + guard).   |

These run automatically inside `bash tests/installer/run-tests.sh` via
the dynamic `check-*.sh` discovery loop.

---

## 7. Forbidden patterns

The following are **bugs** if found in either runner's `fix-repo`
dispatch path:

- ❌ `exec bash "$inner" $*` — unquoted `$*` re-splits on `IFS`.
- ❌ `eval bash "$inner" "$@"` — re-parses every argument.
- ❌ `bash "$inner" "$(printf '%s ' "$@")"` — joins then re-splits.
- ❌ Building a single string in PowerShell: `& $inner ($args -join ' ')`.
- ❌ Conditional flag rewriting (e.g. mapping `-DryRun` → `--dry-run`
  inside the runner). Such mapping belongs in the inner script if
  needed at all.

---

## 8. Change-control

Any change to the flag set, mapping, or forwarding semantics requires:

1. Updating §2, §3, or §4 of this document **first**.
2. Updating both runners and `fix-repo.{sh,ps1}` in the same PR.
3. Updating the help-text files in `scripts/runner-help{,.ps}.txt`.
4. Updating or adding conformance tests in `tests/installer/`.

The sub-command spec ([`spec-authoring/22-fix-repo/01-spec.md`](../../spec-authoring/22-fix-repo/01-spec.md))
remains the authority on inner-script behavior; this document is the
authority only on **how the runners reach it**.
