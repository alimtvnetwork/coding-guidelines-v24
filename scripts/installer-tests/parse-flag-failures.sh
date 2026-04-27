#!/usr/bin/env bash
# =====================================================================
# parse-flag-failures.sh
#
# Parse the output of `tests/installer/run-tests.sh` and surface which
# *flag-related* assertions failed. A check is considered flag-related
# if its name contains "flag" OR "rollback" OR "log" OR "fix-repo".
#
# Usage:
#   bash scripts/installer-tests/parse-flag-failures.sh <test-output-file>
#
# Emits to stdout:
#   - Markdown summary of flag-related failures (and overall pass/fail).
# Exits 0 always — the workflow decides the job status from the
# original test run's exit code.
# =====================================================================
set -uo pipefail

input="${1:-/dev/stdin}"
[[ -r "$input" ]] || { echo "parse-flag-failures: cannot read $input" >&2; exit 0; }

content="$(cat "$input")"

pass_count="$(printf '%s\n' "$content" | grep -E '^  PASS: ' | tail -1 | sed -E 's/.*PASS: ([0-9]+).*/\1/')"
fail_count="$(printf '%s\n' "$content" | grep -E '  FAIL: ' | tail -1 | sed -E 's/.*FAIL: ([0-9]+).*/\1/')"
pass_count="${pass_count:-0}"
fail_count="${fail_count:-0}"

echo "## Installer test suite"
echo ""
echo "- **Passed:** ${pass_count}"
echo "- **Failed:** ${fail_count}"
echo ""

# Collect every ❌ line.
mapfile -t all_failures < <(printf '%s\n' "$content" | grep -E '^  ❌ ' | sed -E 's/^  ❌ //')

if [[ "${#all_failures[@]}" -eq 0 ]]; then
  echo "✅ No failing assertions."
  exit 0
fi

echo "### ❌ Failing assertions (all)"
echo ""
for line in "${all_failures[@]}"; do
  echo "- ${line}"
done
echo ""

# Filter flag-related failures.
flag_failures=()
for line in "${all_failures[@]}"; do
  case "$line" in
    *flag*|*rollback*|*log*|*fix-repo*|*Rollback*|*Log*|*FixRepo*|*prune*|*Prune*)
      flag_failures+=("$line") ;;
  esac
done

echo "### 🚩 Flag-related failures (${#flag_failures[@]})"
echo ""
if [[ "${#flag_failures[@]}" -eq 0 ]]; then
  echo "_None — failures are unrelated to installer flags._"
else
  for line in "${flag_failures[@]}"; do
    echo "- ${line}"
  done
fi
