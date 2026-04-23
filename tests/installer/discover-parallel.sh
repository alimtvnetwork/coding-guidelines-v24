#!/usr/bin/env bash
# =====================================================================
# discover-parallel.sh — V→V+20 parallel cross-repo discovery probe.
#
# This is the canonical implementation of the discovery algorithm
# specified in spec/14-update/27-generic-installer-behavior.md §6.
# The generator will inline this logic into every implicit-mode
# installer in Phase 2 of .lovable/plans/installer-behavior-rollout.md;
# until then, the test harness exercises it directly so we have a
# regression target.
#
# Usage: discover-parallel.sh <repo-base> <current-major>
# Echoes the highest existing K in [V+1, V+20] (one per line,
# space-separated probe results in $DISCOVER_LOG for inspection),
# or exits 1 if no successor repo is reachable.
#
# Probes use `curl -I` against `https://github.com/<repo-base>-v<K>`,
# 5s per probe, fired in parallel with bash `&` + `wait`, deadline
# enforced by the outer `timeout` wrapper (10s).
# =====================================================================
set -euo pipefail

REPO_BASE="${1:?repo-base required}"
CURRENT_V="${2:?current major required}"
LOOKAHEAD="${LOOKAHEAD:-20}"
: "${DISCOVER_LOG:=/dev/null}"

probe_dir="$(mktemp -d)"
trap 'rm -rf "$probe_dir"' EXIT

probe_one() {
  local k="$1"
  local url="https://github.com/${REPO_BASE}-v${k}"
  if curl -fsSI --max-time 5 -o /dev/null "$url"; then
    printf '%s\n' "$k" > "$probe_dir/$k"
  fi
  printf 'PROBE\t%s\t%s\n' "$k" "$url" >> "$DISCOVER_LOG"
}

START_TS="$(date +%s)"
for k in $(seq $((CURRENT_V + 1)) $((CURRENT_V + LOOKAHEAD))); do
  probe_one "$k" &
done
wait
END_TS="$(date +%s)"
printf 'ELAPSED\t%s\n' "$((END_TS - START_TS))" >> "$DISCOVER_LOG"

# Highest K wins.
highest=""
for f in "$probe_dir"/*; do
  [[ -e "$f" ]] || continue
  k="$(basename "$f")"
  if [[ -z "$highest" || "$k" -gt "$highest" ]]; then
    highest="$k"
  fi
done

if [[ -z "$highest" ]]; then
  exit 1
fi
echo "$highest"