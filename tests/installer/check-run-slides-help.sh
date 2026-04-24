#!/usr/bin/env bash
# =====================================================================
# check-run-slides-help.sh — Task 09 E2E (offline)
#
# Asserts `./run.sh help` advertises the `slides` sub-command and
# `./run.sh slides --help` would route to slides-app build (we only
# verify the sub-command dispatch table, not the full git-pull build,
# which requires a clean checkout and network).
#
# Spec: spec/15-distribution-and-runner/02-runner-contract.md
# =====================================================================
set -uo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
RUN_SH="$ROOT/run.sh"
RC=0

if [ ! -x "$RUN_SH" ]; then
    echo "::error::run.sh missing or not executable at $RUN_SH" >&2
    exit 1
fi

# 1. Help table mentions `slides`
if ! "$RUN_SH" help 2>&1 | grep -q "slides"; then
    echo "::error::./run.sh help does not advertise the 'slides' sub-command" >&2
    RC=1
fi

# 2. The slides dispatch branch exists in the script source
if ! grep -qE "^\s*slides\)|case .* in.*slides" "$RUN_SH"; then
    echo "::error::run.sh has no 'slides)' dispatch branch" >&2
    RC=1
fi

# 3. slides-app/ exists and has package.json (target of the sub-command)
if [ ! -f "$ROOT/slides-app/package.json" ]; then
    echo "::error::slides-app/package.json missing — slides sub-command would fail on a clean checkout" >&2
    RC=1
fi

if [ "$RC" -eq 0 ]; then
    echo "✅ run.sh slides sub-command wired correctly"
fi
exit "$RC"
