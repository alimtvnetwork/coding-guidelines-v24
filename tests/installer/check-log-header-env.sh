#!/usr/bin/env bash
# =====================================================================
# check-log-header-env.sh — verify fix-repo log header captures key
# install-environment details (os, shell, uname, cwd) for debugging.
#
# Static contract per Bash installer:
#   B1. Header emits `# os:`     line sourced from `uname -s`.
#   B2. Header emits `# shell:`  line referencing $BASH_VERSION.
#   B3. Header emits `# uname:`  line sourced from `uname -a`.
#   B4. Header emits `# cwd:`    line sourced from `pwd`.
#
# Static contract per PowerShell installer:
#   P1. Header emits `# os:`     via RuntimeInformation::OSDescription.
#   P2. Header emits `# shell:`  via $PSVersionTable.PSEdition + Version.
#   P3. Header emits `# uname:`  via RuntimeInformation OS + Architecture.
#   P4. Header emits `# cwd:`    via (Get-Location).Path.
#
# Generator template (G1) emits B1–B4 + P1–P4 so all generated bundle
# installers stay conformant.
#
# Functional verification (F1):
#   Source the run_fix_repo block from cli-install.sh against a sandbox
#   target with a fake fix-repo.sh that returns 0. Read the produced
#   log file and assert every new header line is present AND non-empty.
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

# ── Bash header static checks ───────────────────────────────────────
for f in "${SH_INSTALLERS[@]}"; do
  path="${ROOT}/${f}"
  [[ -f "${path}" ]] || { fail "${f}: missing"; continue; }

  grep -qE '"# os: .*uname -s' "${path}" \
    && pass "${f}: B1 header captures os via uname -s" \
    || fail "${f}: B1 missing 'uname -s' in '# os:' header line"

  grep -qE '"# shell:.*BASH_VERSION' "${path}" \
    && pass "${f}: B2 header captures shell via \$BASH_VERSION" \
    || fail "${f}: B2 missing \$BASH_VERSION in '# shell:' header line"

  grep -qE '"# uname:.*uname -a' "${path}" \
    && pass "${f}: B3 header captures full uname -a" \
    || fail "${f}: B3 missing 'uname -a' in '# uname:' header line"

  grep -qE '"# cwd: .*\$\(pwd\)' "${path}" \
    && pass "${f}: B4 header captures cwd via \$(pwd)" \
    || fail "${f}: B4 missing \$(pwd) in '# cwd:' header line"
done

# ── PowerShell header static checks ─────────────────────────────────
for f in "${PS_INSTALLERS[@]}"; do
  path="${ROOT}/${f}"
  [[ -f "${path}" ]] || { fail "${f}: missing"; continue; }

  grep -qE '"# os:.*RuntimeInformation.*OSDescription' "${path}" \
    && pass "${f}: P1 header captures os via OSDescription" \
    || fail "${f}: P1 missing OSDescription in '# os:' header line"

  grep -qE '"# shell:.*PSVersionTable.*PSEdition' "${path}" \
    && pass "${f}: P2 header captures shell via \$PSVersionTable" \
    || fail "${f}: P2 missing \$PSVersionTable in '# shell:' header line"

  grep -qE '"# uname:.*OSArchitecture' "${path}" \
    && pass "${f}: P3 header captures arch via OSArchitecture" \
    || fail "${f}: P3 missing OSArchitecture in '# uname:' header line"

  grep -qE '"# cwd:.*Get-Location' "${path}" \
    && pass "${f}: P4 header captures cwd via Get-Location" \
    || fail "${f}: P4 missing Get-Location in '# cwd:' header line"
done

# ── Generator template (G1) ─────────────────────────────────────────
GEN="${ROOT}/scripts/generate-bundle-installers.mjs"
if [[ -f "${GEN}" ]]; then
  for needle in \
    '"# os: .*uname -s' \
    '"# shell:.*BASH_VERSION' \
    '"# uname:.*uname -a' \
    '"# cwd: .*\$\(pwd\)' \
    '"# os:.*OSDescription' \
    '"# shell:.*PSVersionTable' \
    '"# uname:.*OSArchitecture' \
    '"# cwd:.*Get-Location'; do
    grep -qE -- "${needle}" "${GEN}" \
      && pass "generator: G1 emits '${needle}'" \
      || fail "generator: G1 missing '${needle}'"
  done
else
  fail "generator: scripts/generate-bundle-installers.mjs missing"
fi

# ── F1: real log header populated end-to-end ────────────────────────
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
  echo 'LOG_DIR=""'
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

if ! bash "${DRIVER}" </dev/null > "${TMP}/out" 2> "${TMP}/err"; then
  fail "F1 driver exited non-zero"
  sed 's/^/      /' "${TMP}/err" >&2
else
  log="$(ls "${TARGET}/.install-logs"/fix-repo-*.log 2>/dev/null | head -1)"
  if [[ -z "${log}" ]]; then
    fail "F1 no log file produced"
  else
    # Each header line must (a) exist and (b) have a non-empty value.
    check_field() {
      local label="$1" field="$2"
      local line; line="$(grep -E "^# ${field}:" "${log}" | head -1)"
      if [[ -z "${line}" ]]; then
        fail "F1 header missing '# ${field}:' line"
        return
      fi
      local val="${line#*: }"
      val="${val#"${val%%[![:space:]]*}"}"
      if [[ -z "${val}" || "${val}" == "unknown" ]]; then
        fail "F1 '# ${field}:' present but empty/unknown — got: ${line}"
        return
      fi
      pass "F1 ${label}: ${line}"
    }
    check_field "os field"    "os"
    check_field "shell field" "shell"
    check_field "uname field" "uname"
    check_field "cwd field"   "cwd"
  fi
fi

printf '\n  PASS=%d  FAIL=%d\n' "${PASS}" "${FAIL}"
[[ "${RC}" -eq 0 ]] && echo "✅ fix-repo log header captures install env (os/shell/uname/cwd)"
exit "${RC}"
