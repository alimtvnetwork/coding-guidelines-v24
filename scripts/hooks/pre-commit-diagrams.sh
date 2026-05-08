#!/usr/bin/env bash
# ============================================================
# pre-commit-diagrams.sh
#
# Local guard for Mermaid diagrams — mirrors .github/workflows/diagrams-ci.yml.
#
#   1. Always: parse-validate every .mmd against mermaid v11 (cheap).
#   2. If any *.mmd is staged: render them, auto-stage refreshed PNGs.
#   3. Always: drift-check (passes when no PNGs exist or all PNGs are fresh).
#
# Skip with `SKIP_DIAGRAMS_HOOK=1 git commit ...` for emergencies.
# ============================================================

set -euo pipefail

if [ "${SKIP_DIAGRAMS_HOOK:-0}" = "1" ]; then
  echo "  pre-commit: SKIP_DIAGRAMS_HOOK=1 — skipping mermaid checks."
  exit 0
fi

is_command_available() {
  command -v "$1" > /dev/null 2>&1
}

if ! is_command_available node; then
  echo "  pre-commit: node not on PATH — skipping mermaid checks."
  exit 0
fi

STAGED_MMD=$(git diff --cached --name-only --diff-filter=ACMR -- '*.mmd' || true)

# Allow opt-out of the hash cache via NO_DIAGRAM_CACHE=1. Forwarded to
# render-diagrams.mjs as --no-cache so every staged diagram is re-rendered
# from scratch (useful when debugging stale PNGs or cache corruption).
NO_CACHE_FLAG=""
if [ "${NO_DIAGRAM_CACHE:-0}" = "1" ]; then
  NO_CACHE_FLAG="--no-cache"
  echo "  pre-commit: NO_DIAGRAM_CACHE=1 — forcing full re-render (cache bypassed)."
fi

echo "  pre-commit: validating Mermaid sources (mermaid v11 parse)..."
node scripts/validate-mermaid.mjs

if [ -n "$STAGED_MMD" ]; then
  echo "  pre-commit: staged .mmd files detected — rendering PNGs..."
  while IFS= read -r mmd; do
    [ -z "$mmd" ] && continue
    touch "$mmd"
    node scripts/render-diagrams.mjs --only "$mmd"
    png="${mmd%.mmd}.png"
    if [ -f "$png" ]; then
      git add "$png"
      echo "  pre-commit: re-staged $png"
    fi
  done <<< "$STAGED_MMD"

  # Drift-check is scoped to staged-only — keeps pre-commit fast even when
  # the spec/ tree is large. Full-tree drift-check still runs in CI
  # (.github/workflows/diagrams-ci.yml).
  echo "  pre-commit: drift-checking staged diagrams only..."
  node scripts/render-diagrams.mjs --check --staged
else
  echo "  pre-commit: no staged .mmd files — skipping drift-check."
fi
