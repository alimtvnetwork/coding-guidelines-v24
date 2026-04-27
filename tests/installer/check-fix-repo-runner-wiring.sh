#!/usr/bin/env bash
# =====================================================================
# check-fix-repo-runner-wiring.sh — runner wiring for fix-repo
#
# Asserts:
#   1. ./run.sh help advertises the `fix-repo` sub-command.
#   2. run.sh has a 'fix-repo)' case branch.
#   3. fix-repo.sh exists and parses cleanly under bash -n.
#   4. run.ps1 has a 'fix-repo' switch case.
#   5. fix-repo.ps1 exists at repo root.
#
# Spec: spec-authoring/22-fix-repo/01-spec.md
# =====================================================================
set -uo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
RUN_SH="$ROOT/run.sh"
RUN_PS1="$ROOT/run.ps1"
FR_SH="$ROOT/fix-repo.sh"
FR_PS1="$ROOT/fix-repo.ps1"
RC=0

if [ ! -f "$RUN_SH" ]; then
    echo "::error::run.sh missing at $RUN_SH" >&2
    exit 1
fi

if ! "$RUN_SH" help 2>&1 | grep -q "fix-repo"; then
    echo "::error::./run.sh help does not advertise the 'fix-repo' sub-command" >&2
    RC=1
fi

if ! grep -qE "^\s*fix-repo\)" "$RUN_SH"; then
    echo "::error::run.sh has no 'fix-repo)' dispatch branch" >&2
    RC=1
fi

if [ ! -f "$FR_SH" ]; then
    echo "::error::fix-repo.sh missing at $FR_SH" >&2
    RC=1
elif ! bash -n "$FR_SH" 2>/dev/null; then
    echo "::error::fix-repo.sh has bash syntax errors" >&2
    RC=1
fi

if [ -f "$RUN_PS1" ] && ! grep -q '"fix-repo"' "$RUN_PS1"; then
    echo "::error::run.ps1 has no 'fix-repo' switch case" >&2
    RC=1
fi

if [ ! -f "$FR_PS1" ]; then
    echo "::error::fix-repo.ps1 missing at $FR_PS1" >&2
    RC=1
fi

if [ "$RC" -eq 0 ]; then
    echo "✅ fix-repo sub-command wired in run.sh and run.ps1"
fi
exit "$RC"
