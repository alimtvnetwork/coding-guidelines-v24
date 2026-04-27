#!/usr/bin/env bash
# =====================================================================
# check-visibility-runner-wiring.sh — Phase 5 (visibility-change)
#
# Asserts:
#   1. ./run.sh help advertises the `visibility` sub-command.
#   2. run.sh has a 'visibility)' case branch.
#   3. visibility-change.sh exists and is executable / parseable.
#   4. run.ps1 has a 'visibility' switch case.
#
# Spec: spec-authoring/23-visibility-change/01-spec.md §7
# =====================================================================
set -uo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
RUN_SH="$ROOT/run.sh"
RUN_PS1="$ROOT/run.ps1"
VIS_SH="$ROOT/visibility-change.sh"
VIS_PS1="$ROOT/visibility-change.ps1"
RC=0

if [ ! -f "$RUN_SH" ]; then
    echo "::error::run.sh missing at $RUN_SH" >&2
    exit 1
fi

if ! "$RUN_SH" help 2>&1 | grep -q "visibility"; then
    echo "::error::./run.sh help does not advertise the 'visibility' sub-command" >&2
    RC=1
fi

if ! grep -qE "^\s*visibility\)" "$RUN_SH"; then
    echo "::error::run.sh has no 'visibility)' dispatch branch" >&2
    RC=1
fi

if [ ! -f "$VIS_SH" ]; then
    echo "::error::visibility-change.sh missing at $VIS_SH" >&2
    RC=1
elif ! bash -n "$VIS_SH" 2>/dev/null; then
    echo "::error::visibility-change.sh has bash syntax errors" >&2
    RC=1
fi

if [ -f "$RUN_PS1" ] && ! grep -q '"visibility"' "$RUN_PS1"; then
    echo "::error::run.ps1 has no 'visibility' switch case" >&2
    RC=1
fi

if [ ! -f "$VIS_PS1" ]; then
    echo "::error::visibility-change.ps1 missing at $VIS_PS1" >&2
    RC=1
fi

if [ "$RC" -eq 0 ]; then
    echo "✅ visibility sub-command wired in run.sh and run.ps1"
fi
exit "$RC"
