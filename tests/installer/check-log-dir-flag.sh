#!/usr/bin/env bash
# =====================================================================
# check-log-dir-flag.sh — verify --log-dir / -LogDir wiring.
#
# Contract per Bash installer (install.sh, cli-install.sh,
# consolidated-install.sh + every generated bundle *.sh):
#   B1. Declares LOG_DIR sourced from $INSTALL_LOG_DIR (default empty).
#   B2. Parses --log-dir <dir> CLI flag → assigns to LOG_DIR.
#   B3. run_fix_repo() resolves: empty → <target>/.install-logs;
#       absolute → used as-is; relative → joined to target.
#
# Contract per PowerShell installer (install.ps1, cli-install.ps1,
# consolidated-install.ps1):
#   P1. Declares [string]$LogDir = "" parameter.
#   P2. Falls back to $env:INSTALL_LOG_DIR when -LogDir not supplied.
#   P3. Resolves rooted vs joined-to-Target before mkdir.
#
# Generator template (scripts/generate-bundle-installers.mjs):
#   G1. Emits all of B1–B3 + P1–P3.
#
# Functional verification:
#   F1. Source the run_fix_repo block from cli-install.sh, point
#       LOG_DIR at a custom relative path, and assert the timestamped
#       log file lands under <TARGET>/<LOG_DIR>/, not under
#       <TARGET>/.install-logs/.
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

# ── Bash installers (B1–B3) ─────────────────────────────────────────
for f in "${SH_INSTALLERS[@]}"; do
  path="${ROOT}/${f}"
  [[ -f "${path}" ]] || { fail "${f}: missing"; continue; }

  grep -qE 'LOG_DIR="\$\{INSTALL_LOG_DIR:-\}"' "${path}" \
    && pass "${f}: B1 LOG_DIR seeded from INSTALL_LOG_DIR" \
    || fail "${f}: B1 missing LOG_DIR=\${INSTALL_LOG_DIR:-} declaration"

  grep -qE -- '--log-dir\)[[:space:]]*LOG_DIR="\$2"' "${path}" \
    && pass "${f}: B2 parses --log-dir <dir>" \
    || fail "${f}: B2 missing --log-dir CLI parse branch"

  grep -qE 'log_dir="\$\{?LOG_DIR\}?"' "${path}" \
    && grep -qE '\.install-logs' "${path}" \
    && pass "${f}: B3 run_fix_repo resolves LOG_DIR with .install-logs default" \
    || fail "${f}: B3 missing LOG_DIR resolution in run_fix_repo"
done

# ── PowerShell installers (P1–P3) ───────────────────────────────────
for f in "${PS_INSTALLERS[@]}"; do
  path="${ROOT}/${f}"
  [[ -f "${path}" ]] || { fail "${f}: missing"; continue; }

  grep -qE '\[string\]\$LogDir\s*=\s*""' "${path}" \
    && pass "${f}: P1 declares [string]\$LogDir parameter" \
    || fail "${f}: P1 missing [string]\$LogDir = \"\" parameter"

  grep -qE 'LogDir\s*=\s*\$env:INSTALL_LOG_DIR' "${path}" \
    && pass "${f}: P2 reads INSTALL_LOG_DIR env var" \
    || fail "${f}: P2 missing \$env:INSTALL_LOG_DIR fallback"

  grep -qE '\[System\.IO\.Path\]::IsPathRooted\(\$LogDir\)' "${path}" \
    && pass "${f}: P3 resolves rooted vs joined LogDir" \
    || fail "${f}: P3 missing IsPathRooted(\$LogDir) resolution"
done

# ── Generator template (G1) ─────────────────────────────────────────
GEN="${ROOT}/scripts/generate-bundle-installers.mjs"
if [[ -f "${GEN}" ]]; then
  for needle in \
    'INSTALL_LOG_DIR' \
    '[-][-]log-dir' \
    'LOG_DIR="\$2"' \
    '\[string\]\$LogDir' \
    '\[System\.IO\.Path\]::IsPathRooted\(\$LogDir\)'; do
    grep -qE -- "${needle}" "${GEN}" \
      && pass "generator: G1 emits '${needle}'" \
      || fail "generator: G1 missing '${needle}'"
  done
else
  fail "generator: scripts/generate-bundle-installers.mjs missing"
fi

# ── F1: custom --log-dir takes effect at runtime ────────────────────
TMP="$(mktemp -d)"
trap 'rm -rf "${TMP}"' EXIT
TARGET="${TMP}/sandbox"
mkdir -p "${TARGET}"
cat > "${TARGET}/fix-repo.sh" <<'EOF'
#!/usr/bin/env bash
echo "fix-repo OK"
exit 0
EOF
chmod +x "${TARGET}/fix-repo.sh"

DRIVER="${TMP}/driver.sh"
{
  echo '#!/usr/bin/env bash'
  echo 'set -uo pipefail'
  echo "TARGET='${TARGET}'"
  echo 'ASSUME_YES=true'
  echo 'ROLLBACK_ON_FIX_FAIL=false'
  echo 'FULL_ROLLBACK=false'
  echo 'BUNDLE_MAPPING=""'
  echo 'BUNDLE_TOP_LEVEL_FILES=""'
  echo 'PRE_FIX_REPO_HEAD=""'
  echo 'LOG_DIR="my-custom-logs"'   # relative → joined to TARGET
  echo 'SHOW_FIX_REPO_LOG=false'
  echo 'MAX_FIX_REPO_LOGS=0'
  awk '
    /^confirm_fix_repo\(\)/      {capture=1}
    /^snapshot_pre_fix_repo\(\)/ {capture=1}
    /^perform_rollback\(\)/      {capture=1}
    /^run_fix_repo\(\)/          {capture=1}
    capture {print}
    capture && /^\}$/ {capture=0; print ""}
  ' "${ROOT}/cli-install.sh"
  echo 'run_fix_repo'
} > "${DRIVER}"

if bash "${DRIVER}" </dev/null > "${TMP}/out" 2> "${TMP}/err"; then
  if compgen -G "${TARGET}/my-custom-logs/fix-repo-*.log" >/dev/null; then
    pass "F1 relative --log-dir lands under <target>/my-custom-logs/"
  else
    fail "F1 expected log under ${TARGET}/my-custom-logs/, none found"
    ls -laR "${TARGET}" >&2
  fi
  if [[ -d "${TARGET}/.install-logs" ]]; then
    fail "F1 default .install-logs dir was created when --log-dir was set"
  else
    pass "F1 default .install-logs dir was NOT created (correctly bypassed)"
  fi
else
  fail "F1 driver exited non-zero"
  sed 's/^/      /' "${TMP}/err" >&2
fi

# Absolute LOG_DIR variant.
ABS="${TMP}/abs-logs"
DRIVER2="${TMP}/driver2.sh"
{
  echo '#!/usr/bin/env bash'
  echo 'set -uo pipefail'
  echo "TARGET='${TARGET}'"
  echo 'ASSUME_YES=true'
  echo 'ROLLBACK_ON_FIX_FAIL=false'
  echo 'FULL_ROLLBACK=false'
  echo 'BUNDLE_MAPPING=""'
  echo 'BUNDLE_TOP_LEVEL_FILES=""'
  echo 'PRE_FIX_REPO_HEAD=""'
  echo "LOG_DIR='${ABS}'"
  echo 'SHOW_FIX_REPO_LOG=false'
  echo 'MAX_FIX_REPO_LOGS=0'
  awk '
    /^confirm_fix_repo\(\)/      {capture=1}
    /^snapshot_pre_fix_repo\(\)/ {capture=1}
    /^perform_rollback\(\)/      {capture=1}
    /^run_fix_repo\(\)/          {capture=1}
    capture {print}
    capture && /^\}$/ {capture=0; print ""}
  ' "${ROOT}/cli-install.sh"
  echo 'run_fix_repo'
} > "${DRIVER2}"

if bash "${DRIVER2}" </dev/null > "${TMP}/out2" 2> "${TMP}/err2"; then
  if compgen -G "${ABS}/fix-repo-*.log" >/dev/null; then
    pass "F1 absolute --log-dir lands at the absolute path verbatim"
  else
    fail "F1 expected log at ${ABS}/, none found"
  fi
else
  fail "F1 absolute-path driver exited non-zero"
  sed 's/^/      /' "${TMP}/err2" >&2
fi

printf '\n  PASS=%d  FAIL=%d\n' "${PASS}" "${FAIL}"
[[ "${RC}" -eq 0 ]] && echo "✅ --log-dir / -LogDir wired correctly across installers"
exit "${RC}"
