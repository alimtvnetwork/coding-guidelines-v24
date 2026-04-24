#!/usr/bin/env bash
# =====================================================================
# check-install-folders-config.sh — Task 10 E2E (offline)
#
# Asserts install-config.json declares exactly the four canonical folders
# (spec, linters, linter-scripts, linters-cicd) so a fresh `install.sh`
# run pulls all of them. Offline — does not hit GitHub.
#
# Spec: spec/15-distribution-and-runner/04-install-config.md
# =====================================================================
set -uo pipefail

CFG="$(dirname "$0")/../../install-config.json"
EXPECTED=(spec linters linter-scripts linters-cicd)
RC=0

if [ ! -f "$CFG" ]; then
    echo "::error::install-config.json missing at $CFG" >&2
    exit 1
fi

for folder in "${EXPECTED[@]}"; do
    if ! grep -qF "\"$folder\"" "$CFG"; then
        echo "::error::install-config.json missing folder '$folder'" >&2
        RC=1
    fi
done

if [ "$RC" -eq 0 ]; then
    echo "✅ install-config.json declares all 4 canonical folders"
fi
exit "$RC"
