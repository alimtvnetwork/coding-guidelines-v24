#!/usr/bin/env bash
# Verifies `run.sh fix-repo --debug ...`:
#   1. Prints a preflight block to stderr listing received argv.
#   2. Forwards argv to fix-repo.sh UNCHANGED (--debug not consumed).
#   3. A run without --debug prints NO preflight block.
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
WORK="$(mktemp -d)"
trap 'rm -rf "$WORK"' EXIT

cp "$REPO_ROOT/run.sh" "$WORK/run.sh"
mkdir -p "$WORK/scripts"
echo "help" > "$WORK/scripts/runner-help.txt"

# Shim fix-repo.sh that echoes received argv with sentinels
cat > "$WORK/fix-repo.sh" <<'SHIM'
#!/usr/bin/env bash
printf 'INNER_ARGC=%d\n' "$#"
i=0
for a in "$@"; do printf 'INNER_ARG[%d]<<%s>>\n' "$i" "$a"; i=$((i+1)); done
SHIM
chmod +x "$WORK/fix-repo.sh" "$WORK/run.sh"

pass=0; fail=0
assert_contains() {
  local haystack="$1" needle="$2" label="$3"
  if printf '%s' "$haystack" | grep -qF -- "$needle"; then
    echo "  ✓ $label"; pass=$((pass+1))
  else
    echo "  ✗ $label — missing: $needle"; fail=$((fail+1))
  fi
}
assert_absent() {
  local haystack="$1" needle="$2" label="$3"
  if printf '%s' "$haystack" | grep -qF -- "$needle"; then
    echo "  ✗ $label — unexpectedly present: $needle"; fail=$((fail+1))
  else
    echo "  ✓ $label"; pass=$((pass+1))
  fi
}

echo "── case 1: --debug present (with tricky args) ─────────────"
out_file="$WORK/out1.txt"; err_file="$WORK/err1.txt"
"$WORK/run.sh" fix-repo --debug --label "two words" "*.md" >"$out_file" 2>"$err_file"
out="$(cat "$out_file")"; err="$(cat "$err_file")"

assert_contains "$err" "fix-repo preflight" "preflight banner on stderr"
assert_contains "$err" "ARGC=4"             "preflight ARGC=4"
assert_contains "$err" "ARG[0]<<--debug>>"  "preflight ARG[0]"
assert_contains "$err" "ARG[1]<<--label>>"  "preflight ARG[1]"
assert_contains "$err" "ARG[2]<<two words>>" "preflight ARG[2] preserves spaces"
assert_contains "$err" "ARG[3]<<*.md>>"     "preflight ARG[3] preserves glob"

# Inner script must see IDENTICAL argv (--debug not consumed)
assert_contains "$out" "INNER_ARGC=4"             "inner ARGC=4"
assert_contains "$out" "INNER_ARG[0]<<--debug>>"  "inner ARG[0] is --debug"
assert_contains "$out" "INNER_ARG[1]<<--label>>"  "inner ARG[1]"
assert_contains "$out" "INNER_ARG[2]<<two words>>" "inner ARG[2] preserves spaces"
assert_contains "$out" "INNER_ARG[3]<<*.md>>"     "inner ARG[3] preserves glob"

echo "── case 2: no --debug ─────────────────────────────────────"
out_file="$WORK/out2.txt"; err_file="$WORK/err2.txt"
"$WORK/run.sh" fix-repo --label "two words" >"$out_file" 2>"$err_file"
out="$(cat "$out_file")"; err="$(cat "$err_file")"

assert_absent  "$err" "fix-repo preflight" "no preflight without --debug"
assert_contains "$out" "INNER_ARGC=2"              "inner ARGC=2"
assert_contains "$out" "INNER_ARG[1]<<two words>>" "inner preserves spaces without --debug"

echo
echo "── summary ── pass=$pass fail=$fail"
[ "$fail" -eq 0 ]
