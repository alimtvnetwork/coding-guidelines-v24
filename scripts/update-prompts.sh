#!/usr/bin/env bash
set -e

CONFIG_FILE="scripts/prompt-sync-config.json"

if [ ! -f "$CONFIG_FILE" ]; then
    echo "Error: Config file not found at $CONFIG_FILE"
    echo "Please run this script from the repository root."
    exit 1
fi

echo "Compiling prompts from source category folders to flat structure..."

python3 -c "
import json, os, shutil

with open('$CONFIG_FILE', 'r', encoding='utf-8') as f:
    raw = f.read()

config = json.loads(raw)
vars_map = config.get('variables', {})

for k, v in vars_map.items():
    raw = raw.replace('\${' + k + '}', v)

compiled = json.loads(raw)

count = 0
for m in compiled.get('mappings', []):
    src = m['source']
    tgt = m['target']
    os.makedirs(os.path.dirname(tgt), exist_ok=True)
    if os.path.exists(src):
        shutil.copy2(src, tgt)
        count += 1
    else:
        print(f'Warning: Source prompt not found: {src}')

print(f'Successfully compiled and synced {count} prompts.')
"

echo "Validating prompt registry index..."
export PYTHONIOENCODING="utf-8"
python3 linter-scripts/check-prompts-loaded.py

echo "Done!"
