#!/usr/bin/env bash
# Regenerate codegen fixtures wrapper (delegates to cross-platform Python script)
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
exec python3 "$SCRIPT_DIR/regen_codegen_fixtures.py" "$@"
