#!/usr/bin/env bash
# =====================================================================
# check-fix-repo-missing-diagnostic.sh
#
# Asserts that when fix-repo.sh is missing, run.sh:
#   - exits with the documented non-zero code (4),
#   - prints the attempted path,
#   - prints SCRIPT_DIR and the current working directory,
#   - emits a hint pointing at the spec.
# =====================================================================
set -uo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
TMP="$(mktemp -d)"
trap 'rm -rf "$TMP"' EXIT

# Stage a runner copy with NO fix-repo.sh sibling.
cp "$ROOT/run.sh" "$TMP/run.sh"
chmod +x "$TMP/run.sh"
# Stub linter-scripts to keep accidental code paths inert.
mkdir -p "$TMP/linter-scripts"
echo '#!/usr/bin/env bash' > "$TMP/linter-scripts/run.sh"
chmod +x "$TMP/linter-scripts/run.sh"

PASS=0; FAIL=0
ok()  { printf '  ✅ %s\n' "$*"; PASS=$((PASS + 1)); }
bad() { printf '  ❌ %s\n' "$*" >&2; FAIL=$((FAIL + 1)); }

# Run from a distinctive cwd so we can prove it appears in the diagnostic.
CWD_MARKER="$TMP/some-other-dir"
mkdir -p "$CWD_MARKER"

set +e
out="$(cd "$CWD_MARKER" && bash "$TMP/run.sh" fix-repo --dry-run 2>&1)"
rc=$?
set -e

[ "$rc" -eq 4 ]                                  && ok "exit code is 4 (EXIT_FIX_REPO_MISSING)" || bad "exit code was $rc, expected 4"
grep -qF 'fix-repo: inner script is missing' <<<"$out" && ok "header line present"               || bad "missing header line"
grep -qF "$TMP/fix-repo.sh"                   <<<"$out" && ok "attempted path printed"            || bad "attempted path missing from diagnostic"
grep -qF "SCRIPT_DIR     : $TMP"              <<<"$out" && ok "SCRIPT_DIR printed"                || bad "SCRIPT_DIR missing"
grep -qF "working dir    : $CWD_MARKER"       <<<"$out" && ok "current working dir printed"       || bad "PWD missing from diagnostic"
grep -qF 'spec-authoring/22-fix-repo'         <<<"$out" && ok "spec hint present"                 || bad "spec hint missing"

printf '\n  PASS=%d FAIL=%d\n' "$PASS" "$FAIL"
[ "$FAIL" -eq 0 ]
