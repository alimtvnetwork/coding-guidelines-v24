#!/usr/bin/env bash
set -e

# Forwarding wrapper: Prompts are compiled internally from .lovable/prompts/01-prompts-category/
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
exec "$SCRIPT_DIR/update-prompts.sh" "$@"
