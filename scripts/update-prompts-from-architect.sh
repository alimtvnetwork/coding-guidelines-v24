#!/usr/bin/env bash
set -e

TEMP_DIR=$(mktemp -d)
REPO_URL="https://github.com/alimtvnetwork/prompt-architect-v2.git"
OUT_DIR=".lovable/prompts"

echo "Cloning $REPO_URL into temp directory..."
git clone --depth 1 "$REPO_URL" "$TEMP_DIR"

mkdir -p "$OUT_DIR"

copy_prompt() {
  local filename="$1"
  local remotepath="$2"
  local sourcefile="$TEMP_DIR/$remotepath"
  local outfile="$OUT_DIR/$filename"
  
  if [ -f "$sourcefile" ]; then
    echo "Copying $filename..."
    cp "$sourcefile" "$outfile"
  else
    echo "Warning: Source file not found: $sourcefile"
  fi
}

copy_prompt "01-read-memory-enhanced.md" "01-general-prompts/03-read-write/02-read-memory-enhanced.md"
copy_prompt "02-write-antigravity.md" "01-general-prompts/03-read-write/01-write-antigravity.md"
copy_prompt "03-write-memory.md" "01-general-prompts/03-read-write/03-write-memory.md"
copy_prompt "04-plan-coding-guideline-audit.md" "01-general-prompts/04-coding-guidelines/01-plan-coding-guideline-audit.md"
copy_prompt "05-execute-coding-guideline-fix.md" "01-general-prompts/04-coding-guidelines/02-execute-coding-guideline-fix.md"
copy_prompt "06-commit-fix.md" "01-general-prompts/05-commit-and-multi-agent-code-fix/01-commit-fix.md"
copy_prompt "07-commit-fix-v2.md" "01-general-prompts/05-commit-and-multi-agent-code-fix/05-commit-fix-v2.md"
copy_prompt "08-release.md" "01-general-prompts/03-release-management/04-release.md"
copy_prompt "09-clean-artifacts-and-git-history.md" "01-general-prompts/05-commit-and-multi-agent-code-fix/09-clean-artifacts-and-git-history.md"

echo "Cleaning up temp directory..."
rm -rf "$TEMP_DIR"

echo "Checking prompt drift..."
python3 linter-scripts/check-prompts-loaded.py
echo "Done!"
