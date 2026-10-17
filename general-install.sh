#!/usr/bin/env bash
# ──────────────────────────────────────────────────────────────────────
# general-install.sh — General Repository Asset Installer
# Installs .ai-memory, 01-prompts, 03-ai-scripts, .agents/skills, .cursor/skills,
# 02-spec/01-spec-authoring-guide, 02-spec/21-app, and assets folder structure.
# Total exclusion of all linters and fix-repo scripts.
# ──────────────────────────────────────────────────────────────────────
set -euo pipefail
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
if [[ -f "$DIR/install.sh" ]]; then
  exec "$DIR/install.sh" --general "$@"
else
  curl -fsSL https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v24/main/install.sh | bash -s -- --general "$@"
fi
