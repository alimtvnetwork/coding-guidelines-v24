#!/usr/bin/env bash
set -e

# Use relative path or let python resolve it
CONFIG_FILE="scripts/prompt-sync-config.json"

if [ ! -f "$CONFIG_FILE" ]; then
    echo "Error: Config file not found at $CONFIG_FILE"
    echo "Please run this script from the repository root."
    exit 1
fi

# Extract values using Python to avoid jq dependency
REPO_URL=$(python3 -c "import json; print(json.load(open('$CONFIG_FILE'))['repositoryUrl'])")
OUT_DIR=$(python3 -c "import json; print(json.load(open('$CONFIG_FILE'))['outDir'])")

TEMP_DIR=$(mktemp -d)

echo "Cloning $REPO_URL into temp directory..."
git clone --depth 1 "$REPO_URL" "$TEMP_DIR"

mkdir -p "$OUT_DIR"
rm -f "$OUT_DIR"/*.md

# Read mappings and process via bash to avoid Python /tmp/ path translation issues on Windows
python3 -c "
import json
config = json.load(open('$CONFIG_FILE'))
for m in config['mappings']:
    print(m['source'].strip() + '|' + m['target'].strip())
" | while IFS='|' read -r src target; do
    target=$(echo "$target" | tr -d '\r\n[:space:]')
    src=$(echo "$src" | tr -d '\r\n[:space:]')
    sourcefile="$TEMP_DIR/$src"
    outfile="$OUT_DIR/$target"
    if [ -f "$sourcefile" ]; then
        echo "Copying $target..."
        cp "$sourcefile" "$outfile"
    else
        echo "Warning: Source file not found: $sourcefile"
    fi
done

echo "Cleaning up temp directory..."
rm -rf "$TEMP_DIR"

echo "Checking prompt drift..."
python3 linter-scripts/check-prompts-loaded.py || true
echo "Done!"
