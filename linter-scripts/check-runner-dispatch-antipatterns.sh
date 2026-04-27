#!/usr/bin/env bash
# linter-scripts/check-runner-dispatch-antipatterns.sh
#
# Grep-based guard. Fails CI if run.sh or run.ps1 reintroduce forbidden
# dispatch anti-patterns in the `fix-repo` branch specifically.
#
# Spec: spec/15-distribution-and-runner/06-fix-repo-forwarding.md
#
# Strategy: extract ONLY the fix-repo dispatch region from each runner,
# then assert presence of required patterns and absence of forbidden
# patterns within that region. This avoids false positives from
# unrelated code (e.g. `eval` in helpers, or `exit $LASTEXITCODE` in
# the visibility branch).
#
# run.sh region        : the single `fix-repo)` case-arm line.
# run.ps1 region       : the `"fix-repo"` switch arm line.
#
# Exit codes:
#   0 — clean
#   1 — at least one anti-pattern found (or required pattern missing)
#   2 — a required runner file is missing, or region not extractable
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
SH="$REPO_ROOT/run.sh"
PS="$REPO_ROOT/run.ps1"

violations=0

# ── region extractors ───────────────────────────────────────────────
# run.sh: the case dispatch is a single line beginning with `fix-repo)`.
extract_sh_region() {
  grep -nE '^[[:space:]]*fix-repo\)' "$SH" || true
}

# run.ps1: the switch arm is a single line beginning with `"fix-repo"`.
extract_ps_region() {
  grep -nE '^[[:space:]]*"fix-repo"[[:space:]]*\{' "$PS" || true
}

forbid_in_region() {
  local label="$1" region="$2" pattern="$3" reason="$4"
  local hits
  hits="$(printf '%s\n' "$region" | grep -nE "$pattern" || true)"
  [ -z "$hits" ] && return 0
  echo "❌ $label: forbidden pattern in fix-repo dispatch — $reason"
  echo "   regex : $pattern"
  printf '   hit   : %s\n' "$hits"
  violations=$((violations + 1))
}

require_in_region() {
  local label="$1" region="$2" pattern="$3" reason="$4"
  printf '%s\n' "$region" | grep -qE "$pattern" && return 0
  echo "❌ $label: missing required pattern in fix-repo dispatch — $reason"
  echo "   regex : $pattern"
  violations=$((violations + 1))
}

[ -f "$SH" ] || { echo "::error::run.sh missing at $SH" >&2; exit 2; }
[ -f "$PS" ] || { echo "::error::run.ps1 missing at $PS" >&2; exit 2; }

SH_REGION="$(extract_sh_region)"
PS_REGION="$(extract_ps_region)"

[ -n "$SH_REGION" ] || { echo "::error::run.sh: no 'fix-repo)' case arm found" >&2; exit 2; }
[ -n "$PS_REGION" ] || { echo "::error::run.ps1: no '\"fix-repo\"' switch arm found" >&2; exit 2; }

echo "── run.sh fix-repo dispatch guard ──────────────────────"
echo "   region: $SH_REGION"

# Forbidden inside the fix-repo arm:
forbid_in_region "run.sh" "$SH_REGION" '"\$\*"' \
  '"$*" collapses argv into one string; use "$@"'
forbid_in_region "run.sh" "$SH_REGION" '\beval\b' \
  'eval-based dispatch breaks argv quoting'
forbid_in_region "run.sh" "$SH_REGION" '\bbash[[:space:]]+-c\b' \
  '`bash -c "..."` wrapper loses argv boundaries and may mask exit codes'
forbid_in_region "run.sh" "$SH_REGION" 'fix-repo\.sh[^|]*\|[^|&]' \
  'piping fix-repo.sh output masks its exit code'
forbid_in_region "run.sh" "$SH_REGION" '\$\{@:[0-9]+\}' \
  '${@:N} slicing drops original argv; forward "$@" verbatim'

# Required inside the fix-repo arm:
require_in_region "run.sh" "$SH_REGION" '\bexec\b[^#]*fix-repo\.sh[^#]*"\$@"' \
  'must use `exec ... fix-repo.sh "$@"` to forward argv and exit code'

echo "── run.ps1 fix-repo dispatch guard ─────────────────────"
echo "   region: $PS_REGION"

# Forbidden inside the fix-repo arm:
forbid_in_region "run.ps1" "$PS_REGION" '\$args[[:space:]]*-join' \
  '$args -join collapses argv into a single string'
forbid_in_region "run.ps1" "$PS_REGION" '"\$args"' \
  '"$args" interpolation flattens argv; use @args splatting'
forbid_in_region "run.ps1" "$PS_REGION" '\bInvoke-Expression\b' \
  'Invoke-Expression on argv breaks quoting and is unsafe'
forbid_in_region "run.ps1" "$PS_REGION" 'Start-Process\b' \
  'Start-Process detaches the child; cannot propagate $LASTEXITCODE'

# Required inside the fix-repo arm:
require_in_region "run.ps1" "$PS_REGION" '@args' \
  'must invoke inner with @args splatting (preserves argv)'
require_in_region "run.ps1" "$PS_REGION" 'exit[[:space:]]+\$LASTEXITCODE' \
  'must end with `exit $LASTEXITCODE` to propagate exit code'

echo
if [ "$violations" -eq 0 ]; then
  echo "✅ runner dispatch guard: no anti-patterns found"
  exit 0
fi
echo "❌ runner dispatch guard: $violations violation(s)"
exit 1
