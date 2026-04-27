#!/usr/bin/env bash
# linter-scripts/check-runner-dispatch-antipatterns.sh
#
# Grep-based guard. Fails CI if run.sh or run.ps1 reintroduce forbidden
# dispatch anti-patterns that historically broke argv forwarding or exit
# code propagation for sub-commands like `fix-repo`.
#
# Spec: spec/15-distribution-and-runner/06-fix-repo-forwarding.md
#
# Forbidden in run.sh fix-repo dispatch:
#   - "$@" expanded into a quoted string (loses argv boundaries)
#   - eval-based dispatch
#   - "${@:2}" / shift-then-rebuild patterns that drop the original argv
#   - missing `exec` (would leave a wrapper shell that can mask exit codes)
#
# Forbidden in run.ps1 fix-repo dispatch:
#   - String-joined argv (-join, "$args" interpolation)
#   - Invoke-Expression on argv
#   - Missing `exit $LASTEXITCODE` after the inner call
#
# Exit codes:
#   0 — clean
#   1 — at least one anti-pattern found
#   2 — a required runner file is missing
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
SH="$REPO_ROOT/run.sh"
PS="$REPO_ROOT/run.ps1"

violations=0
report() {
  local file="$1" pattern="$2" reason="$3"
  local hits
  hits="$(grep -nE "$pattern" "$file" || true)"
  [ -z "$hits" ] && return 0
  echo "❌ $(basename "$file"): forbidden pattern — $reason"
  echo "   regex : $pattern"
  printf '   hit   : %s\n' "$hits"
  violations=$((violations + 1))
}

assert_present() {
  local file="$1" pattern="$2" reason="$3"
  grep -qE "$pattern" "$file" && return 0
  echo "❌ $(basename "$file"): missing required pattern — $reason"
  echo "   regex : $pattern"
  violations=$((violations + 1))
}

[ -f "$SH" ] || { echo "::error::run.sh missing at $SH" >&2; exit 2; }
[ -f "$PS" ] || { echo "::error::run.ps1 missing at $PS" >&2; exit 2; }

echo "── run.sh dispatch guard ───────────────────────────────"

# 1. Quoted string expansion of "$@" inside fix-repo dispatch line
report "$SH" '\bfix-repo\b.*"\$\*"' \
  '"$*" collapses argv into one string; use "$@"'

# 2. eval / Invoke-Expression style dispatch
report "$SH" '\beval\b.*fix-repo' \
  'eval-based dispatch breaks argv quoting'

# 3. Sub-shell wrapping that swallows the exit code (e.g. `(...)` or `bash -c "..."`)
report "$SH" 'fix-repo\).*bash -c[[:space:]]+"' \
  '`bash -c "..."` wrapper loses argv boundaries and may mask exit codes'

# 4. Pipe after the inner call (would mask exit code via PIPESTATUS handling)
report "$SH" 'fix-repo\.sh.*\|[^|]' \
  'piping fix-repo.sh output masks its exit code'

# 5. Required: exec-style invocation present in fix-repo dispatch
assert_present "$SH" '\bfix-repo\)[^#]*\bexec\b[^#]*fix-repo\.sh[^#]*"\$@"' \
  'fix-repo dispatch must use `exec ... fix-repo.sh "$@"` to forward argv and exit code'

echo "── run.ps1 dispatch guard ──────────────────────────────"

# 1. -join on $args
report "$PS" '\$args[[:space:]]*-join' \
  '$args -join collapses argv into a single string'

# 2. Interpolated "$args" (treats argv as one string)
report "$PS" '"\$args"' \
  '"$args" interpolation flattens argv; use @args splatting'

# 3. Invoke-Expression on argv
report "$PS" 'Invoke-Expression[[:space:]]+.*\$args' \
  'Invoke-Expression on argv breaks quoting and is unsafe'

# 4. Required: splatted call AND exit propagation in fix-repo branch
assert_present "$PS" '"fix-repo".*@args' \
  'fix-repo branch must invoke inner with @args splatting (preserves argv)'
assert_present "$PS" '"fix-repo".*exit[[:space:]]+\$LASTEXITCODE' \
  'fix-repo branch must end with `exit $LASTEXITCODE` to propagate exit code'

echo
if [ "$violations" -eq 0 ]; then
  echo "✅ runner dispatch guard: no anti-patterns found"
  exit 0
fi
echo "❌ runner dispatch guard: $violations violation(s)"
exit 1
