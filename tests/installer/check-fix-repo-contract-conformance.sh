#!/usr/bin/env bash
# =====================================================================
# check-fix-repo-contract-conformance.sh
#
# Asserts that run.sh and run.ps1 conform to the shared runner-contract
# document at spec/15-distribution-and-runner/06-fix-repo-forwarding.md.
#
# Checks:
#   1. The contract document exists and references both runners.
#   2. run.sh dispatches fix-repo via `exec bash "$SCRIPT_DIR/fix-repo.sh" "$@"`
#      (verbatim argv forwarding, no re-quoting).
#   3. run.ps1 dispatches fix-repo via `& $inner @args` then propagates
#      $LASTEXITCODE.
#   4. None of the forbidden anti-patterns from §7 appear in either
#      runner's fix-repo path.
#   5. Help-text source files list the same fix-repo flag tokens.
# =====================================================================
set -uo pipefail

HERE="$(cd "$(dirname "$0")" && pwd)"
REPO_ROOT="$(cd "$HERE/../.." && pwd)"

CONTRACT="$REPO_ROOT/spec/15-distribution-and-runner/06-fix-repo-forwarding.md"
RUN_SH="$REPO_ROOT/run.sh"
RUN_PS1="$REPO_ROOT/run.ps1"
HELP_SH="$REPO_ROOT/scripts/runner-help.txt"
HELP_PS="$REPO_ROOT/scripts/runner-help.ps.txt"

PASS=0; FAIL=0
pass() { printf '  ✅ %s\n' "$*"; PASS=$((PASS + 1)); }
fail() { printf '  ❌ %s\n' "$*" >&2; FAIL=$((FAIL + 1)); }

assert_file() {
  [ -f "$1" ] && pass "exists: $(basename "$1")" || fail "missing: $1"
}

assert_contains() {
  local file="$1" needle="$2" label="$3"
  grep -qF -- "$needle" "$file" && pass "$label" || fail "$label (not found in $(basename "$file"))"
}

assert_absent() {
  local file="$1" needle="$2" label="$3"
  grep -qF -- "$needle" "$file" && fail "$label (forbidden pattern found in $(basename "$file"))" || pass "$label"
}

# 1. Contract document
assert_file "$CONTRACT"
assert_contains "$CONTRACT" "run.sh" "contract references run.sh"
assert_contains "$CONTRACT" "run.ps1" "contract references run.ps1"

# 2. run.sh reference dispatch
assert_contains "$RUN_SH" 'exec bash "$SCRIPT_DIR/fix-repo.sh" "$@"' \
  "run.sh dispatches fix-repo via exec + verbatim \$@"

# 3. run.ps1 reference dispatch
assert_contains "$RUN_PS1" '& $inner @args' \
  "run.ps1 splats @args (no string interpolation)"
assert_contains "$RUN_PS1" 'exit $LASTEXITCODE' \
  "run.ps1 propagates inner exit code"

# 4. Forbidden anti-patterns (§7)
assert_absent "$RUN_SH" 'eval bash' "no eval-based dispatch in run.sh"
assert_absent "$RUN_SH" '$inner $*'  "no unquoted \$* in run.sh"
assert_absent "$RUN_PS1" '$args -join' "no @args join-then-split in run.ps1"

# 5. Help-text parity — both files must list every fix-repo flag token
for token in '--2' '--3' '--5' '--all' '--dry-run' '--verbose'; do
  assert_contains "$HELP_SH" "$token" "runner-help.txt lists $token"
  assert_contains "$HELP_PS" "$token" "runner-help.ps.txt lists $token"
done

printf '\n  Total: PASS=%d FAIL=%d\n' "$PASS" "$FAIL"
[ "$FAIL" -eq 0 ]
