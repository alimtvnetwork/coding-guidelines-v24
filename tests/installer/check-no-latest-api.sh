#!/usr/bin/env bash
# =====================================================================
# check-no-latest-api.sh — static guard test.
#
# Asserts the property from spec/14-update/25-release-pinned-installer.md:
#   release-install.sh and release-install.ps1 MUST NOT call the GitHub
#   "releases/latest" API (or any api.github.com endpoint that resolves
#   "latest"). Comments and help text mentioning "/releases/latest" are
#   permitted; *executable* references are not.
#
# Strategy:
#   - Strip every shell line whose first non-whitespace char is `#`.
#   - Strip every PowerShell line that is a comment (`#`, `<# ... #>`).
#   - Grep the residue for the forbidden tokens.
#
# Exit: 0 clean, 1 violation found.
# =====================================================================
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
SH="$REPO_ROOT/release-install.sh"
PS="$REPO_ROOT/release-install.ps1"
RC=0

check_sh() {
  local file="$1" residue
  residue="$(grep -vE '^[[:space:]]*#' "$file" || true)"
  if grep -qE '(releases/latest|api\.github\.com)' <<<"$residue"; then
    printf '  ❌ %s: executable reference to releases/latest or api.github.com\n' "$file" >&2
    grep -nE '(releases/latest|api\.github\.com)' <<<"$residue" | sed 's/^/      /' >&2
    RC=1
  else
    printf '  ✅ %s: no executable call to releases/latest\n' "$file"
  fi
}

check_ps() {
  local file="$1" residue
  # Drop block comments <# ... #> then line comments starting with #.
  residue="$(awk '
    /<#/ { inblock=1 }
    inblock { if ($0 ~ /#>/) { inblock=0 }; next }
    /^[[:space:]]*#/ { next }
    { print }
  ' "$file")"
  if grep -qE '(releases/latest|api\.github\.com)' <<<"$residue"; then
    printf '  ❌ %s: executable reference to releases/latest or api.github.com\n' "$file" >&2
    grep -nE '(releases/latest|api\.github\.com)' <<<"$residue" | sed 's/^/      /' >&2
    RC=1
  else
    printf '  ✅ %s: no executable call to releases/latest\n' "$file"
  fi
}

printf '\nT3: release-install.* MUST NOT call api.github.com/.../releases/latest\n'
[[ -f "$SH" ]] || { printf '  ❌ missing %s\n' "$SH" >&2; exit 1; }
[[ -f "$PS" ]] || { printf '  ❌ missing %s\n' "$PS" >&2; exit 1; }
check_sh "$SH"
check_ps "$PS"

exit $RC
