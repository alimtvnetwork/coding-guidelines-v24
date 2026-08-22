#!/usr/bin/env bash
set -e

TEMP_DIR=$(mktemp -d)
REPO_URL="https://github.com/alimtvnetwork/prompt-architect-v2.git"
OUT_DIR=".lovable/prompts"

echo "Cloning $REPO_URL into temp directory..."
git clone --depth 1 "$REPO_URL" "$TEMP_DIR"

mkdir -p "$OUT_DIR"
rm -f "$OUT_DIR"/*.md

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

copy_prompt "01-unified-ai-prompt-v4.md" "01-general-prompts/01-core-workflow/05-unified-ai-prompt-v4.md"
copy_prompt "02-next-steps.md" "01-general-prompts/01-core-workflow/01-next-steps.md"
copy_prompt "03-pending-tasks.md" "01-general-prompts/01-core-workflow/02-pending-tasks.md"
copy_prompt "04-plan-steps.md" "01-general-prompts/01-core-workflow/03-plan-steps.md"
copy_prompt "05-read-memory-enhanced.md" "01-general-prompts/03-read-write/02-read-memory-enhanced.md"
copy_prompt "06-write-antigravity.md" "01-general-prompts/03-read-write/01-write-antigravity.md"
copy_prompt "07-write-memory.md" "01-general-prompts/03-read-write/03-write-memory.md"
copy_prompt "08-plan-coding-guideline-audit.md" "01-general-prompts/04-coding-guidelines/01-plan-coding-guideline-audit.md"
copy_prompt "09-execute-coding-guideline-fix.md" "01-general-prompts/04-coding-guidelines/02-execute-coding-guideline-fix.md"
copy_prompt "10-execute-batched-loop.md" "01-general-prompts/07-execute/03-execute-batched-loop.md"
copy_prompt "11-inventory-pending-tasks.md" "01-general-prompts/07-execute/04-inventory-pending-tasks.md"
copy_prompt "12-fix-subtask-naming-convention.md" "01-general-prompts/07-execute/04-fix-subtask-naming-convention.md"
copy_prompt "13-ci-cd-fix.md" "01-general-prompts/08-ci-cd/01-ci-cd-fix.md"
copy_prompt "14-cicd-run-ps1.md" "01-general-prompts/08-ci-cd/02-cicd-run-ps1.md"
copy_prompt "15-commit-fix.md" "01-general-prompts/05-commit-and-multi-agent-code-fix/01-commit-fix.md"
copy_prompt "16-commit-fix-v2.md" "01-general-prompts/05-commit-and-multi-agent-code-fix/05-commit-fix-v2.md"
copy_prompt "17-clean-artifacts-and-git-history.md" "01-general-prompts/05-commit-and-multi-agent-code-fix/09-clean-artifacts-and-git-history.md"
copy_prompt "18-release.md" "01-general-prompts/03-release-management/04-release.md"

echo "Cleaning up temp directory..."
rm -rf "$TEMP_DIR"

echo "Checking prompt drift..."
python3 linter-scripts/check-prompts-loaded.py || true
echo "Done!"
