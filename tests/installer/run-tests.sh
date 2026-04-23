#!/usr/bin/env bash
# =====================================================================
# run-tests.sh — installer behavior test harness.
#
# Verifies the two normative properties from
# spec/14-update/27-generic-installer-behavior.md:
#
#   T1 (PINNED) — when --version is supplied and the release archive
#                 404s, the installer MUST exit non-zero and MUST NOT
#                 request the main-branch tarball. Only valid with
#                 --no-main-fallback (the strict-pinning gate added in
#                 v3.74.0). Without that flag, current bundle
#                 installers fall back to main by design.
#
#   T2 (IMPLICIT) — V→V+20 cross-repo probe runs in parallel, returns
#                   the highest existing K, and completes inside the
#                   10s deadline.
#
# T2 currently exercises the standalone discover-parallel.sh probe
# (which is the canonical implementation that Phase 2 of
# .lovable/plans/installer-behavior-rollout.md will inline into every
# generated installer). Once Phase 2 lands, T2 will additionally drive
# the bundle installers themselves.
#
# Usage: bash tests/installer/run-tests.sh
# Exit:  0 all green, 1 any failure.
# =====================================================================
set -uo pipefail

HERE="$(cd "$(dirname "$0")" && pwd)"
REPO_ROOT="$(cd "$HERE/../.." && pwd)"
PASS=0; FAIL=0

say()  { printf '  %s\n' "$*"; }
pass() { printf '  ✅ %s\n' "$*"; PASS=$((PASS+1)); }
fail() { printf '  ❌ %s\n' "$*" >&2; FAIL=$((FAIL+1)); }

# ── Build a fixture archive once (used by every shim 200 response). ──
fixture_dir="$(mktemp -d)"
trap 'rm -rf "$fixture_dir"' EXIT
mkdir -p "$fixture_dir/payload/spec-slides" \
         "$fixture_dir/payload/slides-app/dist"
echo "<!doctype html><title>fixture</title>" \
  > "$fixture_dir/payload/slides-app/dist/index.html"
echo "fixture" > "$fixture_dir/payload/spec-slides/README.md"
( cd "$fixture_dir/payload" && tar -czf "$fixture_dir/release.tar.gz" . )

make_shim_dir() {
  local d; d="$(mktemp -d)"
  cp "$HERE/curl-shim.sh" "$d/curl"
  chmod +x "$d/curl"
  # No wget — force the installer down the curl path.
  echo "$d"
}

# =====================================================================
# T1 — PINNED MODE never falls back to main when --no-main-fallback
# =====================================================================
printf '\nT1: pinned mode + --no-main-fallback refuses main fallback\n'

t1_dir="$(mktemp -d)"
t1_shim="$(make_shim_dir)"
t1_log="$t1_dir/curl.log"
t1_manifest="$t1_dir/manifest.tsv"
# Manifest: every release-asset URL 404s. Codeload main-branch URL is
# RULED IN (200) on purpose — the test asserts the installer never
# REQUESTS it. If the installer were to ignore --no-main-fallback,
# the request would land here and we'd see it in the log.
cat > "$t1_manifest" <<EOF
codeload.github.com	200	$fixture_dir/release.tar.gz
EOF

SHIM_MANIFEST="$t1_manifest" SHIM_LOG="$t1_log" \
  PATH="$t1_shim:$PATH" \
  bash "$REPO_ROOT/slides-install.sh" \
    --version v999.0.0 \
    --target "$t1_dir/install" \
    --no-main-fallback \
    --no-open \
    > "$t1_dir/stdout.log" 2> "$t1_dir/stderr.log"
rc=$?

if [[ "$rc" -eq 3 ]]; then
  pass "exit code 3 (pinned tag not found)"
else
  fail "expected exit 3, got $rc"
  say "stderr tail:"; tail -n 5 "$t1_dir/stderr.log" | sed 's/^/      /'
fi

if grep -q 'codeload.github.com' "$t1_log"; then
  fail "installer requested codeload (main-branch fallback) — MUST NOT"
  grep 'codeload' "$t1_log" | sed 's/^/      /'
else
  pass "no request to codeload.github.com (no main fallback)"
fi

if grep -q 'releases/download/v999.0.0' "$t1_log"; then
  pass "release-asset URL was attempted (as expected)"
else
  fail "expected a request to releases/download/v999.0.0/*"
fi

# =====================================================================
# T2 — IMPLICIT MODE: V→V+20 parallel discovery
# =====================================================================
printf '\nT2: V→V+20 parallel cross-repo discovery\n'

t2_dir="$(mktemp -d)"
t2_shim="$(make_shim_dir)"
t2_log="$t2_dir/curl.log"
t2_discover="$t2_dir/discover.log"
t2_manifest="$t2_dir/manifest.tsv"
# Successor repos exist at V+3, V+7, V+19. Highest must win → 19.
# All probes 200 a HEAD; nothing else.
cat > "$t2_manifest" <<'EOF'
coding-guidelines-v18	200	/dev/null
coding-guidelines-v22	200	/dev/null
coding-guidelines-v34	200	/dev/null
EOF

START="$(date +%s)"
highest="$(SHIM_MANIFEST="$t2_manifest" SHIM_LOG="$t2_log" \
           DISCOVER_LOG="$t2_discover" PATH="$t2_shim:$PATH" \
           bash "$HERE/discover-parallel.sh" \
                "alimtvnetwork/coding-guidelines" 15 \
           2> "$t2_dir/stderr.log" || true)"
ELAPSED=$(( $(date +%s) - START ))

if [[ "$highest" == "34" ]]; then
  pass "highest successor wins → v34"
else
  fail "expected highest=34, got '$highest'"
fi

probe_count=$(grep -c '^PROBE' "$t2_discover" 2>/dev/null || echo 0)
if [[ "$probe_count" -eq 20 ]]; then
  pass "fired 20 parallel probes (V+1..V+20)"
else
  fail "expected 20 probes, fired $probe_count"
fi

if [[ "$ELAPSED" -le 10 ]]; then
  pass "completed in ${ELAPSED}s (≤10s deadline)"
else
  fail "took ${ELAPSED}s (exceeds 10s deadline)"
fi

# Parallelism check: if probes were serial @ 5s timeout each, a fully
# missing range would take 100s. Even with all-200 HEAD responses,
# serial 20× would dominate. We assert wall-clock ≪ 20× single-probe.
if [[ "$ELAPSED" -le 5 ]]; then
  pass "wall-clock ${ELAPSED}s ≪ serial worst-case (parallel confirmed)"
else
  fail "wall-clock ${ELAPSED}s — probes may not be running in parallel"
fi

# =====================================================================
printf '\n────────────────────────────────────────────\n'
printf '  PASS: %d   FAIL: %d\n' "$PASS" "$FAIL"
printf '────────────────────────────────────────────\n'

[[ "$FAIL" -eq 0 ]] || exit 1