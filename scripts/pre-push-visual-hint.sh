#!/usr/bin/env bash
# scripts/pre-push-visual-hint.sh
#
# Prints the "how to bake missing visual baselines" hint block used by
# .husky/pre-push when scripts/validate-visual-baselines.mjs --strict
# fails. Extracted from the git hook so both branches (sandbox vs
# off-sandbox) can be tested by scripts/tests/pre-push-visual-hint.test.sh.
#
# Environment overrides (test-only):
#   HINT_ENV_OVERRIDE=sandbox      force sandbox branch
#   HINT_ENV_OVERRIDE=host         force off-sandbox branch
# Default probe: presence of /nix/store (matches
# scripts/bake-baselines-sandbox.mjs isSandbox() heuristic).
#
# Exit code is always 0. The caller decides whether to abort the push.

set -eu

detect_env() {
  if [ -n "${HINT_ENV_OVERRIDE:-}" ]; then
    echo "$HINT_ENV_OVERRIDE"
    return
  fi
  if [ -d /nix/store ]; then
    echo "sandbox"
  else
    echo "host"
  fi
}

env_kind=$(detect_env)

case "$env_kind" in
  sandbox)
    bake_cmd="npm run slides:bake-baselines:sandbox"
    bake_ctx="Lovable sandbox: resolves nix-store libs for Chromium"
    ;;
  host)
    bake_cmd="npm run slides:bake-baselines"
    bake_ctx="local dev / CI runner with system Chromium libs"
    ;;
  *)
    echo "  pre-push-visual-hint: unknown HINT_ENV_OVERRIDE='$env_kind' (expected 'sandbox' or 'host')" >&2
    exit 2
    ;;
esac

echo "  pre-push: visual baseline drift, deck grew without a matching baseline."
echo "  Bake locally with: $bake_cmd  ($bake_ctx)"
echo "  Or dispatch .github/workflows/slides-visual.yml with update_baselines=true and commit the artifact."
