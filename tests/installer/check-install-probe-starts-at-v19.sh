#!/usr/bin/env bash
# =====================================================================
# check-install-probe-starts-at-v19.sh
#
# Asserts the installer probe-version floor is pinned to v19 (not the
# pre-renumber v14) in BOTH the root installers and the latest baked
# release artifact pair. The release-artifact path is derived from
# version.json so this test never goes stale on bumps.
# =====================================================================
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"

CURRENT_VERSION="$(node -e 'process.stdout.write(require("./version.json").version)')"
RELEASE_DIR="$ROOT/release-artifacts/coding-guidelines-v${CURRENT_VERSION}"

FILES=(
  "$ROOT/install.ps1"
  "$ROOT/install.sh"
)
if [[ -d "$RELEASE_DIR" ]]; then
  [[ -f "$RELEASE_DIR/install.ps1" ]] && FILES+=("$RELEASE_DIR/install.ps1")
  [[ -f "$RELEASE_DIR/install.sh" ]] && FILES+=("$RELEASE_DIR/install.sh")
fi

for file in "${FILES[@]}"; do
  case "$file" in
    *.ps1) grep -q 'ProbeVersion = 19' "$file" ;;
    *.sh)  grep -q 'PROBE_VERSION_FALLBACK=19' "$file" ;;
  esac
  grep -q 'INSTALLER FAILED — diagnostic report' "$file"
  ! grep -Eq 'ProbeVersion = 14|PROBE_VERSION_FALLBACK=14' "$file"
done
