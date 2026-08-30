#!/usr/bin/env bash
# Verify codegen determinism wrapper (delegates to cross-platform Python script)
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
exec python3 "$SCRIPT_DIR/verify_codegen_determinism.py" "$@"
