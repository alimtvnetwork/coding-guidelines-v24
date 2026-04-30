#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
FILES=(
  "$ROOT/install.ps1"
  "$ROOT/install.sh"
  "$ROOT/release-artifacts/coding-guidelines-v4.24.0/install.ps1"
  "$ROOT/release-artifacts/coding-guidelines-v4.24.0/install.sh"
)

grep -q 'ProbeVersion = 19' "${FILES[0]}"
grep -q 'PROBE_VERSION_FALLBACK=19' "${FILES[1]}"
grep -q 'ProbeVersion = 19' "${FILES[2]}"
grep -q 'PROBE_VERSION_FALLBACK=19' "${FILES[3]}"

for file in "${FILES[@]}"; do
  grep -q 'INSTALLER FAILED — diagnostic report' "$file"
  ! grep -Eq 'ProbeVersion = 14|PROBE_VERSION_FALLBACK=14' "$file"
done