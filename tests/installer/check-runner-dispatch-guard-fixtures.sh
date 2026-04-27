#!/usr/bin/env bash
# tests/installer/check-runner-dispatch-guard-fixtures.sh
#
# Fixture-based tests for linter-scripts/check-runner-dispatch-antipatterns.sh
#
# Strategy: build a temp directory containing synthetic run.sh / run.ps1
# fixtures, then invoke the guard against each fixture by overriding the
# repo root it scans (it's resolved as `dirname "$0"/..`, so we copy the
# guard into <fixture>/linter-scripts/ and execute it from there).
#
# Each case asserts:
#   - exit code (0 = pass, 1 = violations, 2 = malformed)
#   - presence/absence of expected REASON substrings in output
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
GUARD="$REPO_ROOT/linter-scripts/check-runner-dispatch-antipatterns.sh"
[ -f "$GUARD" ] || { echo "❌ guard not found at $GUARD" >&2; exit 2; }

TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT

pass=0; fail=0
record_pass() { pass=$((pass+1)); echo "  ✅ $1"; }
record_fail() { fail=$((fail+1)); echo "  ❌ $1"; echo "$2" | sed 's/^/      /'; }

# ── fixture writers ─────────────────────────────────────────────────
write_clean_sh() {
  cat > "$1" <<'SH'
#!/usr/bin/env bash
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
case "${1:-}" in
  fix-repo)      shift; exec bash "$SCRIPT_DIR/fix-repo.sh" "$@" ;;
  *) echo "noop" ;;
esac
SH
}
write_clean_ps() {
  cat > "$1" <<'PS'
param([string]$Command = "")
switch ($Command.ToLower()) {
  "fix-repo"   { $inner = "x"; & $inner @args; exit $LASTEXITCODE }
  default      { Write-Host "noop" }
}
PS
}

# Build an isolated fake repo containing run.sh, run.ps1 and the guard.
build_fixture() {
  local dir="$1" sh_writer="$2" ps_writer="$3"
  mkdir -p "$dir/linter-scripts"
  "$sh_writer" "$dir/run.sh"
  "$ps_writer" "$dir/run.ps1"
  cp "$GUARD" "$dir/linter-scripts/check-runner-dispatch-antipatterns.sh"
  chmod +x "$dir/linter-scripts/check-runner-dispatch-antipatterns.sh"
}

run_guard() {
  local dir="$1"
  bash "$dir/linter-scripts/check-runner-dispatch-antipatterns.sh" 2>&1
}

# ── assertion helpers ───────────────────────────────────────────────
assert_exit() {
  local name="$1" want="$2" got="$3" out="$4"
  [ "$got" = "$want" ] && { record_pass "$name (exit=$got)"; return; }
  record_fail "$name expected exit=$want got=$got" "$out"
}
assert_contains() {
  local name="$1" needle="$2" out="$3"
  printf '%s' "$out" | grep -qF "$needle" && { record_pass "$name contains '$needle'"; return; }
  record_fail "$name missing '$needle'" "$out"
}
assert_not_contains() {
  local name="$1" needle="$2" out="$3"
  printf '%s' "$out" | grep -qF "$needle" || { record_pass "$name lacks '$needle'"; return; }
  record_fail "$name unexpectedly contains '$needle'" "$out"
}

# ── CASE 1: clean fixtures → exit 0 ─────────────────────────────────
echo "── CASE 1: clean run.sh + run.ps1 should pass"
D1="$TMP/case1"
build_fixture "$D1" write_clean_sh write_clean_ps
out="$(run_guard "$D1")" && rc=0 || rc=$?
assert_exit "case1" 0 "$rc" "$out"
assert_contains "case1" "no anti-patterns found" "$out"

# ── CASE 2: run.sh uses eval + "$*" → exit 1, both reasons reported
echo "── CASE 2: run.sh with eval and \"\$*\" should fail"
D2="$TMP/case2"
write_dirty_sh_eval() {
  cat > "$1" <<'SH'
#!/usr/bin/env bash
case "${1:-}" in
  fix-repo)  shift; eval bash "./fix-repo.sh" "$*" ;;
esac
SH
}
build_fixture "$D2" write_dirty_sh_eval write_clean_ps
out="$(run_guard "$D2")" && rc=0 || rc=$?
assert_exit "case2" 1 "$rc" "$out"
assert_contains "case2" 'eval-based dispatch' "$out"
assert_contains "case2" 'collapses argv into one string' "$out"
assert_contains "case2" 'must use `exec ... fix-repo.sh "$@"`' "$out"
assert_contains "case2" 'run.sh:' "$out"

# ── CASE 3: run.ps1 missing exit $LASTEXITCODE → exit 1
echo "── CASE 3: run.ps1 without exit \$LASTEXITCODE should fail"
D3="$TMP/case3"
write_dirty_ps_noexit() {
  cat > "$1" <<'PS'
param([string]$Command = "")
switch ($Command.ToLower()) {
  "fix-repo"   { $inner = "x"; & $inner @args }
}
PS
}
build_fixture "$D3" write_clean_sh write_dirty_ps_noexit
out="$(run_guard "$D3")" && rc=0 || rc=$?
assert_exit "case3" 1 "$rc" "$out"
assert_contains "case3" 'must end with `exit $LASTEXITCODE`' "$out"
assert_contains "case3" 'run.ps1:' "$out"

# ── CASE 4: run.ps1 uses Invoke-Expression "$args" → exit 1
echo "── CASE 4: run.ps1 with Invoke-Expression and \"\$args\" should fail"
D4="$TMP/case4"
write_dirty_ps_iex() {
  cat > "$1" <<'PS'
param([string]$Command = "")
switch ($Command.ToLower()) {
  "fix-repo"   { Invoke-Expression "$args"; exit $LASTEXITCODE }
}
PS
}
build_fixture "$D4" write_clean_sh write_dirty_ps_iex
out="$(run_guard "$D4")" && rc=0 || rc=$?
assert_exit "case4" 1 "$rc" "$out"
assert_contains "case4" 'Invoke-Expression on argv' "$out"
assert_contains "case4" '"$args" interpolation flattens argv' "$out"
assert_contains "case4" 'must invoke inner with @args splatting' "$out"

# ── CASE 5: missing fix-repo arm in run.sh → exit 2 (malformed)
echo "── CASE 5: run.sh without fix-repo arm should error 2"
D5="$TMP/case5"
write_no_arm_sh() { cat > "$1" <<'SH'
#!/usr/bin/env bash
echo "no dispatch here"
SH
}
build_fixture "$D5" write_no_arm_sh write_clean_ps
out="$(run_guard "$D5")" && rc=0 || rc=$?
assert_exit "case5" 2 "$rc" "$out"
assert_contains "case5" "no 'fix-repo)' case arm found" "$out"

# ── CASE 6: false-positive guard — eval/exit elsewhere in file is OK
echo "── CASE 6: eval and exit \$LASTEXITCODE outside fix-repo arm should NOT trip"
D6="$TMP/case6"
write_clean_sh_with_eval_elsewhere() {
  cat > "$1" <<'SH'
#!/usr/bin/env bash
helper() { eval "echo $1"; }   # eval in unrelated helper
case "${1:-}" in
  fix-repo)  shift; exec bash "./fix-repo.sh" "$@" ;;
esac
SH
}
write_clean_ps_with_exit_elsewhere() {
  cat > "$1" <<'PS'
param([string]$Command = "")
function Invoke-Lint { & "x"; exit $LASTEXITCODE }   # unrelated branch
switch ($Command.ToLower()) {
  "lint"       { Invoke-Lint }
  "fix-repo"   { & "y" @args; exit $LASTEXITCODE }
}
PS
}
build_fixture "$D6" write_clean_sh_with_eval_elsewhere write_clean_ps_with_exit_elsewhere
out="$(run_guard "$D6")" && rc=0 || rc=$?
assert_exit "case6" 0 "$rc" "$out"
assert_not_contains "case6" 'FORBIDDEN' "$out"
assert_not_contains "case6" 'MISSING' "$out"

# ── CASE 7: PS [string]::Join + [string]$args + cmd /c → all caught
echo "── CASE 7: PS array-to-string conversions should fail"
D7="$TMP/case7"
write_dirty_ps_join() {
  cat > "$1" <<'PS'
param([string]$Command = "")
switch ($Command.ToLower()) {
  "fix-repo"   { $s = [string]::Join(' ', $args); cmd /c "x $s"; & $inner @args; exit $LASTEXITCODE }
}
PS
}
build_fixture "$D7" write_clean_sh write_dirty_ps_join
out="$(run_guard "$D7")" && rc=0 || rc=$?
assert_exit "case7" 1 "$rc" "$out"
assert_contains "case7" '[string]::Join on $args' "$out"
assert_contains "case7" '`cmd /c` wrapper re-parses argv' "$out"

# ── CASE 8: SH printf %q join, IFS mutation, sh -c → all caught
echo "── CASE 8: SH printf/IFS/sh -c anti-patterns should fail"
D8="$TMP/case8"
write_dirty_sh_more() {
  cat > "$1" <<'SH'
#!/usr/bin/env bash
case "${1:-}" in
  fix-repo)  shift; IFS=, ; sh -c "./fix-repo.sh $(printf '%q ' "$@")" ;;
esac
SH
}
build_fixture "$D8" write_dirty_sh_more write_clean_ps
out="$(run_guard "$D8")" && rc=0 || rc=$?
assert_exit "case8" 1 "$rc" "$out"
assert_contains "case8" "sh -c" "$out"
assert_contains "case8" "printf %q/%s rebuilds argv" "$out"
assert_contains "case8" "mutating IFS" "$out"
assert_contains "case8" 'command substitution on "$@"' "$out"

# ── CASE 9: PS $args.ToString() + Start-Job → caught
echo "── CASE 9: PS \$args.ToString and Start-Job should fail"
D9="$TMP/case9"
write_dirty_ps_tostring() {
  cat > "$1" <<'PS'
param([string]$Command = "")
switch ($Command.ToLower()) {
  "fix-repo"   { $s = $args.ToString(); Start-Job { & "x" $using:s } | Out-Null; & $inner @args; exit $LASTEXITCODE }
}
PS
}
build_fixture "$D9" write_clean_sh write_dirty_ps_tostring
out="$(run_guard "$D9")" && rc=0 || rc=$?
assert_exit "case9" 1 "$rc" "$out"
assert_contains "case9" '$args.ToString() flattens' "$out"
assert_contains "case9" 'Start-Job detaches' "$out"
# ── CASE 10: whitespace/indent variants in run.sh fix-repo arm
echo "── CASE 10: SH whitespace + tab indentation should remain clean"
D10="$TMP/case10"
write_sh_tabs_and_extra_spaces() {
  printf '%s\n' \
    '#!/usr/bin/env bash' \
    'case "${1:-}" in' \
    $'\t  fix-repo)\t   shift   ;     exec   bash    "./fix-repo.sh"     "$@"   ;;' \
    'esac' > "$1"
}
build_fixture "$D10" write_sh_tabs_and_extra_spaces write_clean_ps
out="$(run_guard "$D10")" && rc=0 || rc=$?
assert_exit "case10" 0 "$rc" "$out"
assert_contains "case10" 'no anti-patterns found' "$out"

# ── CASE 11: many args with quotes/spaces in inner invocation comments
echo "── CASE 11: SH with rich preflight + many helper calls should remain clean"
D11="$TMP/case11"
write_sh_rich_preflight() {
  cat > "$1" <<'SH'
#!/usr/bin/env bash
case "${1:-}" in
  fix-repo)  _assert; shift; _preflight "$@"; _trace "arg1 arg2 arg3"; exec bash "./fix-repo.sh" "$@" ;;
esac
SH
}
build_fixture "$D11" write_sh_rich_preflight write_clean_ps
out="$(run_guard "$D11")" && rc=0 || rc=$?
assert_exit "case11" 0 "$rc" "$out"
assert_contains "case11" 'no anti-patterns found' "$out"

# ── CASE 12: PS arm with extra whitespace + multiple statements should pass
echo "── CASE 12: PS arm with extra whitespace and statements should remain clean"
D12="$TMP/case12"
write_ps_extra_ws() {
  printf '%s\n' \
    'param([string]$Command = "")' \
    'switch ($Command.ToLower()) {' \
    $'  "fix-repo"     {   $inner = "x" ;  Write-Host "go" ;   & $inner   @args  ;   exit   $LASTEXITCODE   }' \
    '}' > "$1"
}
build_fixture "$D12" write_clean_sh write_ps_extra_ws
out="$(run_guard "$D12")" && rc=0 || rc=$?
assert_exit "case12" 0 "$rc" "$out"
assert_contains "case12" 'no anti-patterns found' "$out"

# ── CASE 13: SH arm preceded by leading-spaces-only (no tabs) and trailing comment
echo "── CASE 13: SH with trailing comment after dispatch should remain clean"
D13="$TMP/case13"
write_sh_trailing_comment() {
  cat > "$1" <<'SH'
#!/usr/bin/env bash
case "${1:-}" in
    fix-repo)  shift; exec bash "./fix-repo.sh" "$@" ;;   # forward verbatim
esac
SH
}
build_fixture "$D13" write_sh_trailing_comment write_clean_ps
out="$(run_guard "$D13")" && rc=0 || rc=$?
assert_exit "case13" 0 "$rc" "$out"
assert_contains "case13" 'no anti-patterns found' "$out"

# ── CASE 14: PS arm with ToLower() switch case + many helper calls
echo "── CASE 14: PS arm with multiple helper calls should remain clean"
D14="$TMP/case14"
write_ps_many_helpers() {
  cat > "$1" <<'PS'
param([string]$Command = "")
switch ($Command.ToLower()) {
  "fix-repo"   { $i = Resolve-Inner; Write-Preflight $args; Write-Trace "x y z"; & $i @args; exit $LASTEXITCODE }
}
PS
}
build_fixture "$D14" write_clean_sh write_ps_many_helpers
out="$(run_guard "$D14")" && rc=0 || rc=$?
assert_exit "case14" 0 "$rc" "$out"
assert_contains "case14" 'no anti-patterns found' "$out"

# ── CASE 15: stable region detection — line numbers reported correctly
echo "── CASE 15: dispatch arm at non-trivial line number should report that line"
D15="$TMP/case15"
write_sh_with_padding() {
  {
    printf '#!/usr/bin/env bash\n'
    for i in $(seq 1 20); do printf '# pad line %d\n' "$i"; done
    printf 'case "${1:-}" in\n'
    printf '  fix-repo)  shift; exec bash "./fix-repo.sh" "$@" ;;\n'
    printf 'esac\n'
  } > "$1"
}
build_fixture "$D15" write_sh_with_padding write_clean_ps
out="$(run_guard "$D15")" && rc=0 || rc=$?
assert_exit "case15" 0 "$rc" "$out"
# fix-repo arm sits at line 23 (1 shebang + 20 pad + 1 case + 1)
assert_contains "case15" 'run.sh:23' "$out"


echo
echo "════════════════════════════════════════════════════════════"
echo "  fixture suite: $pass passed, $fail failed"
echo "════════════════════════════════════════════════════════════"
[ "$fail" -eq 0 ] || exit 1
