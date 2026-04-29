#!/usr/bin/env bash
# =====================================================================
# check-bundle-installers-drift.sh
#
# Pre-test guard. Re-runs the bundle generator in-memory and asserts
# every committed `<bundle>-install.{sh,ps1}` at the repo root matches
# the generator output. On drift, prints a unified diff and exits 1
# so the rest of the test suite isn't testing stale files.
#
# Wraps: scripts/check-bundle-installers-drift.mjs
# =====================================================================
set -uo pipefail

HERE="$(cd "$(dirname "$0")" && pwd)"
REPO_ROOT="$(cd "$HERE/../.." && pwd)"

if ! command -v node >/dev/null 2>&1; then
  echo "❌ node is required for the bundle drift check" >&2
  exit 1
fi

node "$REPO_ROOT/scripts/check-bundle-installers-drift.mjs"
