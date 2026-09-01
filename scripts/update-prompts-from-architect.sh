#!/usr/bin/env bash
set -e

# Forwarding wrapper: Prompts are compiled internally from 01-prompts/
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
exec "$SCRIPT_DIR/update-prompts.sh" "$@"
