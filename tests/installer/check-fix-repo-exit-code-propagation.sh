#!/usr/bin/env bash
# Verifies run.sh (and run.ps1 if pwsh available) propagate fix-repo's
# non-zero exit code unchanged to the caller.
#
# Spec: spec/15-distribution-and-runner/06-fix-repo-forwarding.md
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT

# Codes to test (incl. 1, a mid value, and the highest portable POSIX shell exit)
CODES=(1 2 7 42 99 113 255)

pass=0; fail=0
assert_eq() {
  local got="$1" want="$2" label="$3"
  if [ "$got" = "$want" ]; then
    echo "  ✓ $label (exit=$got)"; pass=$((pass+1))
  else
    echo "  ✗ $label — got=$got want=$want"; fail=$((fail+1))
  fi
}

# ── Bash: run.sh ─────────────────────────────────────────────
echo "── run.sh propagation ──────────────────────────────────"
SH_DIR="$WORK/sh"
mkdir -p "$SH_DIR/scripts"
cp "$REPO_ROOT/run.sh" "$SH_DIR/run.sh"
echo "help" > "$SH_DIR/scripts/runner-help.txt"
cat > "$SH_DIR/fix-repo.sh" <<'SHIM'
#!/usr/bin/env bash
# Echo argv on stdout (proves we ran), then exit with caller-requested code.
printf 'INNER_RAN argc=%d\n' "$#"
exit "${EXIT_CODE:-0}"
SHIM
chmod +x "$SH_DIR/run.sh" "$SH_DIR/fix-repo.sh"

for code in "${CODES[@]}"; do
  set +e
  EXIT_CODE="$code" "$SH_DIR/run.sh" fix-repo --flag "two words" >/dev/null 2>&1
  got=$?
  set -e
  assert_eq "$got" "$code" "run.sh exit propagated"
done

# Sanity: with --debug the preflight also must not swallow the exit code
set +e
EXIT_CODE=37 "$SH_DIR/run.sh" fix-repo --debug --x "y z" >/dev/null 2>&1
got=$?
set -e
assert_eq "$got" "37" "run.sh exit propagated with --debug preflight"

# ── PowerShell: run.ps1 (only if pwsh available) ────────────
if command -v pwsh >/dev/null 2>&1; then
  echo "── run.ps1 propagation ─────────────────────────────────"
  PS_DIR="$WORK/ps"
  mkdir -p "$PS_DIR"
  cp "$REPO_ROOT/run.ps1" "$PS_DIR/run.ps1"
  cat > "$PS_DIR/fix-repo.ps1" <<'SHIM'
$code = 0
if ($env:EXIT_CODE) { $code = [int]$env:EXIT_CODE }
Write-Host "INNER_RAN argc=$($args.Count)"
exit $code
SHIM

  for code in "${CODES[@]}"; do
    set +e
    EXIT_CODE="$code" pwsh -NoProfile -File "$PS_DIR/run.ps1" fix-repo --flag "two words" >/dev/null 2>&1
    got=$?
    set -e
    assert_eq "$got" "$code" "run.ps1 exit propagated"
  done

  set +e
  EXIT_CODE=37 pwsh -NoProfile -File "$PS_DIR/run.ps1" fix-repo --debug --x "y z" >/dev/null 2>&1
  got=$?
  set -e
  assert_eq "$got" "37" "run.ps1 exit propagated with --debug preflight"
else
  echo "── run.ps1 propagation ── SKIPPED (pwsh not on PATH)"
fi

echo
echo "── summary ── pass=$pass fail=$fail"
[ "$fail" -eq 0 ]
