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
fi

echo "  pre-commit: drift-checking committed PNGs vs sources..."
node scripts/render-diagrams.mjs --check
