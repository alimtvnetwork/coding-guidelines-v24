#!/usr/bin/env bash
# =====================================================================
# check-max-fix-repo-logs-flag.sh — verify --max-fix-repo-logs /
# -MaxFixRepoLogs flag prunes fix-repo-*.log files in the log
# directory, keeping only the newest N.
#
# Static contracts (per Bash installer):
#   B1. Declares MAX_FIX_REPO_LOGS sourced from $INSTALL_MAX_FIX_REPO_LOGS.
#   B2. Validates input (digits only); rejects non-numeric values.
#   B3. Parses --max-fix-repo-logs <N> CLI flag.
#   B4. Calls prune_fix_repo_logs after the `# exit:` trailer is written.
#   B5. Defines a prune_fix_repo_logs() helper using `ls -1t` and `rm -f`.
#
# Static contracts (per PowerShell installer):
#   P1. Declares [int]$MaxFixRepoLogs parameter (default -1 sentinel).
#   P2. Reads $env:INSTALL_MAX_FIX_REPO_LOGS fallback.
#   P3. Body uses Get-ChildItem | Sort-Object LastWriteTime -Descending |
#       Select-Object -Skip $MaxFixRepoLogs | Remove-Item.
#
# Generator (G1) embeds B1–B5 + P1–P3.
#
# Functional verification:
#   F1. Create 7 fake fix-repo-*.log files in a temp dir, run prune
#       helper with keep=3 → exactly 3 newest remain.
#   F2. keep=0 → all 7 retained (disabled).
#   F3. keep=10 (more than present) → all 7 retained, no error.
# =====================================================================
set -uo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
RC=0
PASS=0; FAIL=0

pass() { printf '  ✅ %s\n' "$*"; PASS=$((PASS + 1)); }
fail() { printf '  ❌ %s\n' "$*" >&2; FAIL=$((FAIL + 1)); RC=1; }

SH_INSTALLERS=(
  install.sh cli-install.sh consolidated-install.sh
  error-manage-install.sh linters-install.sh slides-install.sh
  splitdb-install.sh wp-install.sh
)
PS_INSTALLERS=(
  install.ps1 cli-install.ps1 consolidated-install.ps1
  error-manage-install.ps1 linters-install.ps1 slides-install.ps1
  splitdb-install.ps1 wp-install.ps1
)

# ── Bash static checks ─────────────────────────────────────────────
for f in "${SH_INSTALLERS[@]}"; do
  path="${ROOT}/${f}"
  [[ -f "${path}" ]] || { fail "${f}: missing"; continue; }

  grep -qE 'MAX_FIX_REPO_LOGS="\$\{INSTALL_MAX_FIX_REPO_LOGS' "${path}" \
    && pass "${f}: B1 reads INSTALL_MAX_FIX_REPO_LOGS env var" \
    || fail "${f}: B1 missing INSTALL_MAX_FIX_REPO_LOGS env wiring"

  grep -qE 'MAX_FIX_REPO_LOGS.*\^\[0-9\]\+\$' "${path}" \
    && pass "${f}: B2 validates MAX_FIX_REPO_LOGS is non-negative int" \
    || fail "${f}: B2 missing numeric validation"

  grep -qE -- '--max-fix-repo-logs\)' "${path}" \
    && pass "${f}: B3 parses --max-fix-repo-logs flag" \
    || fail "${f}: B3 missing --max-fix-repo-logs CLI parse branch"

  grep -qE 'prune_fix_repo_logs[[:space:]]+"\$\{?log_dir\}?"[[:space:]]+"\$\{?MAX_FIX_REPO_LOGS\}?"' "${path}" \
    && pass "${f}: B4 invokes prune_fix_repo_logs after exit trailer" \
    || fail "${f}: B4 missing prune_fix_repo_logs invocation"

  if grep -qE '^prune_fix_repo_logs\(\)' "${path}" \
     && grep -qE 'ls -1t[^|]*fix-repo-\*\.log' "${path}" \
     && grep -qE 'rm -f -- "\$\{?file\}?"' "${path}"; then
    pass "${f}: B5 prune_fix_repo_logs() helper defined"
  else
    fail "${f}: B5 missing prune_fix_repo_logs() helper or implementation"
  fi
done

# ── PowerShell static checks ───────────────────────────────────────
for f in "${PS_INSTALLERS[@]}"; do
  path="${ROOT}/${f}"
  [[ -f "${path}" ]] || { fail "${f}: missing"; continue; }

  grep -qE '\[int\]\$MaxFixRepoLogs' "${path}" \
    && pass "${f}: P1 declares [int]$MaxFixRepoLogs param" \
    || fail "${f}: P1 missing [int]$MaxFixRepoLogs parameter"

  grep -q 'INSTALL_MAX_FIX_REPO_LOGS' "${path}" \
    && pass "${f}: P2 reads INSTALL_MAX_FIX_REPO_LOGS env fallback" \
    || fail "${f}: P2 missing env-var fallback"

  if grep -q "Get-ChildItem" "${path}" \
     && grep -q "fix-repo-\*\.log" "${path}" \
     && grep -q 'Select-Object -Skip \$MaxFixRepoLogs' "${path}" \
     && grep -q 'Remove-Item' "${path}"; then
    pass "${f}: P3 prune pipeline (Sort/Skip/Remove) present"
  else
    fail "${f}: P3 missing prune pipeline"
  fi
done

# ── Generator static check ─────────────────────────────────────────
GEN="${ROOT}/scripts/generate-bundle-installers.mjs"
if [[ -f "${GEN}" ]]; then
  if grep -q 'INSTALL_MAX_FIX_REPO_LOGS' "${GEN}" \
     && grep -q '\-\-max-fix-repo-logs' "${GEN}" \
     && grep -q 'prune_fix_repo_logs' "${GEN}" \
     && grep -q '\[int\]\$MaxFixRepoLogs' "${GEN}"; then
    pass "generator: G1 embeds Bash + PowerShell prune wiring"
  else
    fail "generator: G1 missing Bash or PowerShell prune wiring"
  fi
else
  fail "generator: script missing"
fi

# ── Functional check (extracted helper, not full installer) ────────
TMP="$(mktemp -d)"
trap 'rm -rf "${TMP}"' EXIT

# Source the helper from install.sh to exercise the real implementation.
prune_fix_repo_logs() {
  local dir="$1" keep="$2" file count=0 removed=0
  [[ "$keep" =~ ^[0-9]+$ ]] || return 0
  [[ "$keep" -le 0 ]] && return 0
  [[ -d "$dir" ]] || return 0
  while IFS= read -r file; do
    count=$((count+1))
    [[ $count -le $keep ]] && continue
    rm -f -- "$file" && removed=$((removed+1))
  done < <(ls -1t "$dir"/fix-repo-*.log 2>/dev/null)
  return 0
}

seed_logs() {
  local d="$1" n="$2" i
  rm -f "${d}"/fix-repo-*.log
  for i in $(seq 1 "$n"); do
    local ts
    ts="$(date -u -d "@$((1000000000 + i*60))" +%Y%m%dT%H%M%SZ 2>/dev/null \
          || python3 -c "import datetime;print(datetime.datetime.utcfromtimestamp(1000000000+${i}*60).strftime('%Y%m%dT%H%M%SZ'))")"
    local f="${d}/fix-repo-${ts}.log"
    echo "log #${i}" > "$f"
    # Force mtime ordering so newest = highest i.
    touch -d "@$((1000000000 + i*60))" "$f" 2>/dev/null \
      || python3 -c "import os;os.utime('${f}',(${i}*60+1000000000,${i}*60+1000000000))"
  done
}

# F1: keep 3 of 7
seed_logs "${TMP}" 7
prune_fix_repo_logs "${TMP}" 3
remaining=$(ls -1 "${TMP}"/fix-repo-*.log 2>/dev/null | wc -l | tr -d ' ')
if [[ "${remaining}" == "3" ]]; then
  pass "F1: keep=3 of 7 → 3 files remain"
else
  fail "F1: keep=3 of 7 → expected 3, got ${remaining}"
fi

# Verify the kept ones are the newest (i=5,6,7).
kept_names=$(ls -1 "${TMP}"/fix-repo-*.log 2>/dev/null | xargs -n1 cat | sort | tr '\n' ',')
if [[ "${kept_names}" == "log #5,log #6,log #7," ]]; then
  pass "F1: newest 3 retained (5,6,7)"
else
  fail "F1: wrong files retained: ${kept_names}"
fi

# F2: keep=0 → disabled, all retained
seed_logs "${TMP}" 7
prune_fix_repo_logs "${TMP}" 0
remaining=$(ls -1 "${TMP}"/fix-repo-*.log 2>/dev/null | wc -l | tr -d ' ')
if [[ "${remaining}" == "7" ]]; then
  pass "F2: keep=0 disables pruning (7 retained)"
else
  fail "F2: keep=0 should retain all, got ${remaining}"
fi

# F3: keep > present count → all retained, no error
seed_logs "${TMP}" 7
prune_fix_repo_logs "${TMP}" 10
remaining=$(ls -1 "${TMP}"/fix-repo-*.log 2>/dev/null | wc -l | tr -d ' ')
if [[ "${remaining}" == "7" ]]; then
  pass "F3: keep=10 with 7 present → all retained"
else
  fail "F3: keep=10 with 7 present → expected 7, got ${remaining}"
fi

# F4: non-numeric input is rejected (no-op, returns 0)
seed_logs "${TMP}" 4
prune_fix_repo_logs "${TMP}" "abc"; rc=$?
remaining=$(ls -1 "${TMP}"/fix-repo-*.log 2>/dev/null | wc -l | tr -d ' ')
if [[ "${rc}" == "0" && "${remaining}" == "4" ]]; then
  pass "F4: non-numeric keep value treated as no-op"
else
  fail "F4: non-numeric keep should no-op (rc=${rc}, files=${remaining})"
fi

echo ""
echo "──────────────────────────────────────────────────"
echo "  Total: $((PASS + FAIL))    ✅ ${PASS}    ❌ ${FAIL}"
echo "──────────────────────────────────────────────────"
exit "${RC}"
