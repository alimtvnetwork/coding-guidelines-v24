#!/usr/bin/env bash
# =====================================================================
# check-run-fix-repo-flag.sh — verify --run-fix-repo / -RunFixRepo wiring
# across every installer variant.
#
# Per-installer contract (Bash *.sh):
#   B1. Declares RUN_FIX_REPO sourced from $INSTALL_RUN_FIX_REPO.
#   B2. Parses the --run-fix-repo CLI flag → sets RUN_FIX_REPO=true.
#   B3. Documents the flag in --help / usage text.
#   B4. Final dispatch line `${RUN_FIX_REPO} && run_fix_repo` exists.
#   B5. run_fix_repo() picks fix-repo.ps1 on MINGW/MSYS/CYGWIN, else
#       fix-repo.sh — i.e. invokes the correct script for the platform.
#
# Per-installer contract (PowerShell *.ps1):
#   P1. Declares `[switch]$RunFixRepo` parameter.
#   P2. Reads `INSTALL_RUN_FIX_REPO` env var as alternate enabler.
#   P3. Final dispatch `if ($RunFixRepo) { Invoke-FixRepo }` exists.
#   P4. Invoke-FixRepo references fix-repo.ps1 (Windows-native runner).
#
# Per-installer contract (generator template):
#   G1. scripts/generate-bundle-installers.mjs embeds B1–B5 + P1–P4
#       so newly generated bundle installers stay conformant.
#
# Functional verification (one happy-path slice):
#   F1. Synthesize a minimal sandbox with a fake fix-repo.sh that
#       writes a marker file. Source the run_fix_repo block from
#       cli-install.sh and invoke it. Marker MUST appear → proves the
#       flag actually executes the right script end-to-end.
#
# Spec: spec-authoring/22-fix-repo/01-spec.md (auto-run option)
# =====================================================================
set -uo pipefail

ROOT="$(cd "$(dirname "$0")/../.." && pwd)"
RC=0
PASS=0; FAIL=0

pass() { printf '  ✅ %s\n' "$*"; PASS=$((PASS + 1)); }
fail() { printf '  ❌ %s\n' "$*" >&2; FAIL=$((FAIL + 1)); RC=1; }

SH_INSTALLERS=(
  install.sh
  cli-install.sh
  consolidated-install.sh
  error-manage-install.sh
  linters-install.sh
  slides-install.sh
  splitdb-install.sh
  wp-install.sh
)
PS_INSTALLERS=(
  install.ps1
  cli-install.ps1
  consolidated-install.ps1
  error-manage-install.ps1
  linters-install.ps1
  slides-install.ps1
  splitdb-install.ps1
  wp-install.ps1
)

# ── Bash installers (B1–B5) ─────────────────────────────────────────
for f in "${SH_INSTALLERS[@]}"; do
  path="${ROOT}/${f}"
  if [[ ! -f "${path}" ]]; then
    fail "${f}: file missing"
    continue
  fi

  grep -q 'RUN_FIX_REPO="\${INSTALL_RUN_FIX_REPO' "${path}" \
    && pass "${f}: B1 reads INSTALL_RUN_FIX_REPO env var" \
    || fail "${f}: B1 missing INSTALL_RUN_FIX_REPO env wiring"

  grep -qE -- '--run-fix-repo\)[[:space:]]*RUN_FIX_REPO=true' "${path}" \
    && pass "${f}: B2 parses --run-fix-repo flag" \
    || fail "${f}: B2 missing --run-fix-repo CLI parse branch"

  grep -q -- '--run-fix-repo' "${path}" \
    && pass "${f}: B3 documents --run-fix-repo in help" \
    || fail "${f}: B3 --run-fix-repo not mentioned in help text"

  # Accept either:  ${RUN_FIX_REPO} && run_fix_repo
  # or:             if ... $RUN_FIX_REPO; then run_fix_repo; fi
  if grep -qE '\$\{?RUN_FIX_REPO\}?[[:space:]]*&&[[:space:]]*run_fix_repo' "${path}" \
       || grep -qE '\$RUN_FIX_REPO[^A-Z_].*run_fix_repo' "${path}"; then
    pass "${f}: B4 dispatches to run_fix_repo when flag set"
  else
    fail "${f}: B4 missing run_fix_repo dispatch gated by \$RUN_FIX_REPO"
  fi

  # B5: function picks the OS-correct script.
  if grep -q 'fix-repo.ps1' "${path}" && grep -q 'fix-repo.sh' "${path}" \
       && grep -qE 'MINGW|MSYS|CYGWIN' "${path}"; then
    pass "${f}: B5 selects fix-repo.{sh,ps1} per OS"
  else
    fail "${f}: B5 missing OS-conditional fix-repo selection"
  fi
done

# ── PowerShell installers (P1–P4) ───────────────────────────────────
for f in "${PS_INSTALLERS[@]}"; do
  path="${ROOT}/${f}"
  if [[ ! -f "${path}" ]]; then
    fail "${f}: file missing"
    continue
  fi

  grep -qE '\[switch\]\$RunFixRepo' "${path}" \
    && pass "${f}: P1 declares [switch]\$RunFixRepo" \
    || fail "${f}: P1 missing [switch]\$RunFixRepo parameter"

  grep -q 'INSTALL_RUN_FIX_REPO' "${path}" \
    && pass "${f}: P2 reads INSTALL_RUN_FIX_REPO env var" \
    || fail "${f}: P2 missing INSTALL_RUN_FIX_REPO env wiring"

  grep -qE 'if[[:space:]]*\(\s*\$RunFixRepo\s*\)\s*\{\s*Invoke-FixRepo' "${path}" \
    && pass "${f}: P3 dispatches to Invoke-FixRepo when -RunFixRepo set" \
    || fail "${f}: P3 missing 'if (\$RunFixRepo) { Invoke-FixRepo }' dispatch"

  grep -q 'fix-repo.ps1' "${path}" \
    && pass "${f}: P4 invokes fix-repo.ps1" \
    || fail "${f}: P4 fix-repo.ps1 not referenced in installer"
done

# ── Generator template (G1) ─────────────────────────────────────────
GEN="${ROOT}/scripts/generate-bundle-installers.mjs"
if [[ -f "${GEN}" ]]; then
  for needle in \
    'INSTALL_RUN_FIX_REPO' \
    '\-\-run-fix-repo' \
    'RUN_FIX_REPO=true' \
    '\${RUN_FIX_REPO} && run_fix_repo' \
    '\[switch\]\$RunFixRepo' \
    'if \(\$RunFixRepo\) \{ Invoke-FixRepo \}' \
    'fix-repo.sh' \
    'fix-repo.ps1'; do
    grep -qE -- "${needle}" "${GEN}" \
      && pass "generator: G1 emits '${needle}'" \
      || fail "generator: G1 missing emission of '${needle}'"
  done
else
  fail "generator: scripts/generate-bundle-installers.mjs missing"
fi

# ── F1: functional happy-path on cli-install.sh (representative) ────
# We don't run the full installer (network heavy). Instead, extract the
# `run_fix_repo` function + its prereq helpers, neutralise side-deps,
# and invoke against a sandbox containing a marker fix-repo.sh.
TMP="$(mktemp -d)"
trap 'rm -rf "${TMP}"' EXIT
TARGET="${TMP}/sandbox"
mkdir -p "${TARGET}"
MARKER="${TARGET}/.fix-repo-was-run"
cat > "${TARGET}/fix-repo.sh" <<EOF
#!/usr/bin/env bash
echo "fix-repo invoked with TARGET=\${PWD}" > "${MARKER}"
exit 0
EOF
chmod +x "${TARGET}/fix-repo.sh"

# Build a minimal driver that stitches the installer's helpers + a
# stand-in environment, then calls run_fix_repo.
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
  # Extract the four helper functions from cli-install.sh.
  awk '
    /^confirm_fix_repo\(\)/      {capture=1}
    /^snapshot_pre_fix_repo\(\)/ {capture=1}
    /^perform_rollback\(\)/      {capture=1}
    /^run_fix_repo\(\)/          {capture=1}
    capture {print}
    capture && /^\}$/ {capture=0; print ""}
  ' "${ROOT}/cli-install.sh"
  echo 'RUN_FIX_REPO=true'
  echo '${RUN_FIX_REPO} && run_fix_repo'
} > "${DRIVER}"

if bash "${DRIVER}" </dev/null > "${TMP}/driver.out" 2> "${TMP}/driver.err"; then
  if [[ -f "${MARKER}" ]]; then
    pass "F1 cli-install.sh run_fix_repo executes target fix-repo.sh"
  else
    fail "F1 driver succeeded but fix-repo.sh marker absent"
    sed 's/^/      /' "${TMP}/driver.out" "${TMP}/driver.err" >&2
  fi
else
  fail "F1 driver exited non-zero"
  sed 's/^/      /' "${TMP}/driver.err" >&2
fi

# ── Summary ─────────────────────────────────────────────────────────
printf '\n  PASS=%d  FAIL=%d\n' "${PASS}" "${FAIL}"
[[ "${RC}" -eq 0 ]] && echo "✅ --run-fix-repo / -RunFixRepo wired correctly across all installers"
exit "${RC}"
