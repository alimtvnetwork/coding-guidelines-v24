#!/usr/bin/env bash
# scripts/lint-ci-verify-present.sh
#
# Mirrors the "Verify linter scripts present" step from
# .github/workflows/ci.yml. Catches missing-script regressions
# locally before they fail in CI. Keep in lockstep with ci.yml.
set -euo pipefail

required=(
  linter-scripts/validate-guidelines.go
  linter-scripts/validate-guidelines.py
  linter-scripts/check-axios-version.sh
  linter-scripts/check-spec-folder-refs.py
  linter-scripts/check-forbidden-spec-paths.sh
  linter-scripts/check-forbidden-strings.py
  linter-scripts/check-readme-install-section.py
  linter-scripts/check-placeholder-comments.py
  linter-scripts/check-prompts-loaded.py
  linter-scripts/check-spec-cross-links.py
  linter-scripts/forbidden-strings.toml
)

missing=0
for f in "${required[@]}"; do
  if [[ ! -f "$f" ]]; then
    echo "::error::$f missing" >&2
    missing=$((missing + 1))
  fi
done
if [[ $missing -gt 0 ]]; then
  echo "lint-ci: $missing required linter file(s) missing" >&2
  exit 1
fi
echo "lint-ci: all ${#required[@]} required linter files present"
