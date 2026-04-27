#!/usr/bin/env bash
# =====================================================================
# check-visibility-arg-parsing.sh — Phase 5 (visibility-change)
#
# Verifies exit codes for argument validation, isolated from any real
# git/gh/glab calls. We invoke visibility-change.sh from a tmp dir that
# is NOT a git repo, so the script reaches arg parsing before any
# network/CLI step.
#
# Cases:
#   1. -h / --help          → exit 0
#   2. unknown flag         → exit 6 (EXIT_BAD_FLAG)
#   3. --visible bogus      → exit 6 (EXIT_BAD_FLAG)
#   4. --visible pub (no repo) → exit 2 (EXIT_NOT_A_REPO)
#   5. --visible private (no repo) → exit 2
#   6. (no args, no repo)   → exit 2
# =====================================================================
set -uo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
VIS_SH="$ROOT/visibility-change.sh"
TMPDIR="$(mktemp -d)"
trap 'rm -rf "$TMPDIR"' EXIT
RC=0

run_case() {
    local name="$1" expected="$2"; shift 2
    local actual
    ( cd "$TMPDIR" && bash "$VIS_SH" "$@" ) >/dev/null 2>&1
    actual=$?
    if [ "$actual" -ne "$expected" ]; then
        echo "::error::case '$name' expected exit $expected, got $actual" >&2
        RC=1
    fi
}

run_case "help short"        0 -h
run_case "help long"         0 --help
run_case "unknown flag"      6 --bogus
run_case "bad --visible"     6 --visible nope
run_case "force pub no repo" 2 --visible pub
run_case "force pri no repo" 2 --visible pri
run_case "no args no repo"   2

if [ "$RC" -eq 0 ]; then
    echo "✅ visibility-change.sh argument parsing exit codes correct (6 cases)"
fi
exit "$RC"
