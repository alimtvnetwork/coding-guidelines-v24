#!/usr/bin/env bash
# =====================================================================
# check-show-fix-repo-log-flag.sh — verify --show-fix-repo-log /
# -ShowFixRepoLog flag prints the latest fix-repo log to stdout.
#
# Static contracts (per Bash installer):
#   B1. Declares SHOW_FIX_REPO_LOG sourced from $INSTALL_SHOW_FIX_REPO_LOG.
#   B2. Parses --show-fix-repo-log → SHOW_FIX_REPO_LOG=true.
#   B3. run_fix_repo() prints log via `cat "$log_file"` gated by
#       $SHOW_FIX_REPO_LOG, AFTER writing the `# exit:` trailer.
#
# Static contracts (per PowerShell installer):
#   P1. Declares [switch]$ShowFixRepoLog parameter.
#   P2. Reads $env:INSTALL_SHOW_FIX_REPO_LOG fallback.
#   P3. Body emits Get-Content of the log gated by $ShowFixRepoLog.
#
# Generator (G1) embeds B1–B3 + P1–P3.
#
# Functional verification:
#   F1. SUCCESS path — flag enabled, fix-repo exits 0: stdout contains
#       the log header lines (`# fix-repo log`, `# exit: 0`).
#   F2. FAILURE path — fix-repo exits non-zero: log STILL printed, then
#       the driver exits 5.
#   F3. DEFAULT — flag NOT set: log NOT echoed to stdout (privacy).
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

  grep -qE 'SHOW_FIX_REPO_LOG="\$\{INSTALL_SHOW_FIX_REPO_LOG' "${path}" \
    && pass "${f}: B1 reads INSTALL_SHOW_FIX_REPO_LOG env var" \
    || fail "${f}: B1 missing INSTALL_SHOW_FIX_REPO_LOG env wiring"

  grep -qE -- '--show-fix-repo-log\)[[:space:]]*SHOW_FIX_REPO_LOG=true' "${path}" \
    && pass "${f}: B2 parses --show-fix-repo-log flag" \
    || fail "${f}: B2 missing --show-fix-repo-log CLI parse branch"

  # B3: gated `cat "$log_file"` after the # exit: trailer is written.
  if grep -qE '\$\{?SHOW_FIX_REPO_LOG\}?[[:space:];]' "${path}" \
       && grep -qE 'cat "\$\{?log_file\}?"' "${path}"; then
    pass "${f}: B3 run_fix_repo cats log file when flag set"
  else
    fail "${f}: B3 missing 'cat \$log_file' gated by \$SHOW_FIX_REPO_LOG"
  fi
done

# ── PowerShell static checks ───────────────────────────────────────
for f in "${PS_INSTALLERS[@]}"; do
  path="${ROOT}/${f}"
  [[ -f "${path}" ]] || { fail "${f}: missing"; continue; }

  grep -qE '\[switch\]\$ShowFixRepoLog' "${path}" \
    && pass "${f}: P1 declares [switch]\$ShowFixRepoLog" \
    || fail "${f}: P1 missing [switch]\$ShowFixRepoLog parameter"

  grep -q 'INSTALL_SHOW_FIX_REPO_LOG' "${path}" \
    && pass "${f}: P2 reads INSTALL_SHOW_FIX_REPO_LOG env var" \
    || fail "${f}: P2 missing INSTALL_SHOW_FIX_REPO_LOG env wiring"

  if grep -qE 'if[[:space:]]*\(\s*\$ShowFixRepoLog\s*\)' "${path}" \
       && grep -qE 'Get-Content[[:space:]]+-LiteralPath[[:space:]]+\$logFile' "${path}"; then
    pass "${f}: P3 emits Get-Content of logFile gated by \$ShowFixRepoLog"
  else
    fail "${f}: P3 missing Get-Content gated by \$ShowFixRepoLog"
  fi
done

# ── Generator (G1) ─────────────────────────────────────────────────
GEN="${ROOT}/scripts/generate-bundle-installers.mjs"
if [[ -f "${GEN}" ]]; then
  for needle in \
    'INSTALL_SHOW_FIX_REPO_LOG' \
    '[-][-]show-fix-repo-log' \
    'SHOW_FIX_REPO_LOG=true' \
    '\$\{?SHOW_FIX_REPO_LOG\}?' \
    'cat "\\?\$\{?log_file\}?"' \
    '\[switch\]\$ShowFixRepoLog' \
    'if \(\$ShowFixRepoLog\)' \
    'Get-Content -LiteralPath \$logFile'; do
    grep -qE -- "${needle}" "${GEN}" \
      && pass "generator: G1 emits '${needle}'" \
      || fail "generator: G1 missing '${needle}'"
  done
else
  fail "generator: scripts/generate-bundle-installers.mjs missing"
fi

# ── Functional driver builder (extracts cli-install.sh helpers) ────
TMP="$(mktemp -d)"
trap 'rm -rf "${TMP}"' EXIT

build_driver() {
  local target="$1" show="$2" exitcode="$3" out="$4"
  cat > "${target}/fix-repo.sh" <<EOF
#!/usr/bin/env bash
echo "synthetic fix-repo body line"
exit ${exitcode}
EOF
  chmod +x "${target}/fix-repo.sh"
  {
    echo '#!/usr/bin/env bash'
    echo 'set -uo pipefail'
    echo "TARGET='${target}'"
    echo 'ASSUME_YES=true'
    echo 'ROLLBACK_ON_FIX_FAIL=false'
    echo 'FULL_ROLLBACK=false'
    echo 'BUNDLE_MAPPING=""'
    echo 'BUNDLE_TOP_LEVEL_FILES=""'
    echo 'PRE_FIX_REPO_HEAD=""'
    echo 'LOG_DIR=""'
    echo "SHOW_FIX_REPO_LOG=${show}"
    awk '
      /^confirm_fix_repo\(\)/      {capture=1}
      /^snapshot_pre_fix_repo\(\)/ {capture=1}
      /^perform_rollback\(\)/      {capture=1}
      /^run_fix_repo\(\)/          {capture=1}
      capture {print}
      capture && /^\}$/ {capture=0; print ""}
    ' "${ROOT}/cli-install.sh"
    echo 'run_fix_repo'
  } > "${out}"
}

# F1 — success + flag → stdout contains log
T1="${TMP}/f1"; mkdir -p "${T1}"
D1="${TMP}/d1.sh"
build_driver "${T1}" true 0 "${D1}"
bash "${D1}" </dev/null > "${TMP}/f1.out" 2> "${TMP}/f1.err"
rc1=$?
if [[ ${rc1} -eq 0 ]] \
     && grep -q '# fix-repo log' "${TMP}/f1.out" \
     && grep -q '# exit: 0' "${TMP}/f1.out" \
     && grep -q 'end of log' "${TMP}/f1.out"; then
  pass "F1 success path prints log to stdout when flag set"
else
  fail "F1 success path did not print log (rc=${rc1})"
  sed 's/^/      out: /' "${TMP}/f1.out" >&2
fi

# F2 — failure + flag → log STILL printed, then exit 5
T2="${TMP}/f2"; mkdir -p "${T2}"
D2="${TMP}/d2.sh"
build_driver "${T2}" true 7 "${D2}"
bash "${D2}" </dev/null > "${TMP}/f2.out" 2> "${TMP}/f2.err"
rc2=$?
if [[ ${rc2} -eq 5 ]] \
     && grep -q '# exit: 7' "${TMP}/f2.out" \
     && grep -q 'end of log' "${TMP}/f2.out"; then
  pass "F2 failure path still prints log, then exits 5"
else
  fail "F2 failure path missing log or wrong exit (rc=${rc2})"
  sed 's/^/      out: /' "${TMP}/f2.out" >&2
  sed 's/^/      err: /' "${TMP}/f2.err" >&2
fi

# F3 — flag NOT set → log NOT echoed (silence by default)
T3="${TMP}/f3"; mkdir -p "${T3}"
D3="${TMP}/d3.sh"
build_driver "${T3}" false 0 "${D3}"
bash "${D3}" </dev/null > "${TMP}/f3.out" 2> "${TMP}/f3.err"
rc3=$?
if [[ ${rc3} -eq 0 ]] \
     && ! grep -q 'end of log' "${TMP}/f3.out" \
     && ! grep -q '# fix-repo log' "${TMP}/f3.out"; then
  pass "F3 default (flag off) does NOT echo log to stdout"
else
  fail "F3 default-off leaked log to stdout (rc=${rc3})"
  sed 's/^/      out: /' "${TMP}/f3.out" >&2
fi

printf '\n  PASS=%d  FAIL=%d\n' "${PASS}" "${FAIL}"
[[ "${RC}" -eq 0 ]] && echo "✅ --show-fix-repo-log / -ShowFixRepoLog wired correctly"
exit "${RC}"
