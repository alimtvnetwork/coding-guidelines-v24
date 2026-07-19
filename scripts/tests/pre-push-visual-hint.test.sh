#!/usr/bin/env bash
# scripts/tests/pre-push-visual-hint.test.sh
#
# Exercises both branches of scripts/pre-push-visual-hint.sh via
# HINT_ENV_OVERRIDE. Guards against:
#   1. Swapped branches (sandbox path recommending the host command
#      and vice versa) that would leave sandbox users chasing
#      libglib-2.0.so.0 errors after following the hint.
#   2. Unknown-env silent success. The helper must exit non-zero when
#      HINT_ENV_OVERRIDE is neither 'sandbox' nor 'host'.
#
# Run standalone: bash scripts/tests/pre-push-visual-hint.test.sh
# Exit code = number of failed cases.

set -u

HELPER="scripts/pre-push-visual-hint.sh"
failed=0

assert_contains() {
  local name="$1"; local haystack="$2"; local needle="$3"
  if [[ "$haystack" == *"$needle"* ]]; then
    echo "  ok  $name"
  else
    failed=$((failed + 1))
    echo "  FAIL $name"
    echo "       expected substring: $needle"
    echo "       got: $haystack"
  fi
}

assert_not_contains() {
  local name="$1"; local haystack="$2"; local needle="$3"
  if [[ "$haystack" != *"$needle"* ]]; then
    echo "  ok  $name"
  else
    failed=$((failed + 1))
    echo "  FAIL $name (unwanted substring '$needle' present)"
  fi
}

# --- Case: sandbox branch ---------------------------------------------
sandbox_out=$(HINT_ENV_OVERRIDE=sandbox bash "$HELPER")
assert_contains "sandbox branch recommends :sandbox target" \
  "$sandbox_out" "npm run slides:bake-baselines:sandbox"
assert_contains "sandbox branch labels the context correctly" \
  "$sandbox_out" "Lovable sandbox"
assert_not_contains "sandbox branch does NOT recommend the host target" \
  "$sandbox_out" "run slides:bake-baselines  ("

# --- Case: host branch ------------------------------------------------
host_out=$(HINT_ENV_OVERRIDE=host bash "$HELPER")
assert_contains "host branch recommends plain bake target" \
  "$host_out" "npm run slides:bake-baselines  "
assert_contains "host branch labels the context correctly" \
  "$host_out" "local dev / CI runner"
assert_not_contains "host branch does NOT recommend the sandbox target" \
  "$host_out" ":sandbox"

# --- Case: both branches include the workflow-dispatch fallback -------
assert_contains "sandbox branch mentions workflow dispatch fallback" \
  "$sandbox_out" "slides-visual.yml with update_baselines=true"
assert_contains "host branch mentions workflow dispatch fallback" \
  "$host_out" "slides-visual.yml with update_baselines=true"

# --- Case: unknown env exits non-zero ---------------------------------
if HINT_ENV_OVERRIDE=weird bash "$HELPER" >/dev/null 2>&1; then
  failed=$((failed + 1))
  echo "  FAIL unknown env should exit non-zero"
else
  echo "  ok  unknown env exits non-zero"
fi

if (( failed > 0 )); then
  echo ""
  echo "$failed case(s) failed"
  exit 1
fi
echo ""
echo "OK, pre-push-visual-hint self-test passed."
