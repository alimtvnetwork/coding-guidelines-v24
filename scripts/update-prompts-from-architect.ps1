$ErrorActionPreference = "Stop"

$TempDir = Join-Path -Path $env:TEMP -ChildPath "prompt-architect-v2-$(Get-Random)"
$RepoUrl = "https://github.com/alimtvnetwork/prompt-architect-v2.git"
$OutDir = ".lovable/prompts"

Write-Host "Cloning $RepoUrl into temp directory..."
git clone --depth 1 $RepoUrl $TempDir

if (-not (Test-Path -Path $OutDir)) {
    New-Item -ItemType Directory -Path $OutDir | Out-Null
}
Remove-Item -Path "$OutDir\*.md" -Force -ErrorAction SilentlyContinue

function Copy-Prompt($FileName, $RemotePath) {
    $SourceFile = Join-Path -Path $TempDir -ChildPath $RemotePath
    $OutFile = Join-Path -Path $OutDir -ChildPath $FileName
    if (Test-Path -Path $SourceFile) {
        Write-Host "Copying $FileName..."
        Copy-Item -Path $SourceFile -Destination $OutFile -Force
    } else {
        Write-Warning "Source file not found: $SourceFile"
    }
}

Copy-Prompt "01-unified-ai-prompt-v4.md" "01-general-prompts/01-core-workflow/05-unified-ai-prompt-v4.md"
Copy-Prompt "02-next-steps.md" "01-general-prompts/01-core-workflow/01-next-steps.md"
Copy-Prompt "03-pending-tasks.md" "01-general-prompts/01-core-workflow/02-pending-tasks.md"
Copy-Prompt "04-plan-steps.md" "01-general-prompts/01-core-workflow/03-plan-steps.md"
Copy-Prompt "05-read-memory-enhanced.md" "01-general-prompts/03-read-write/02-read-memory-enhanced.md"
Copy-Prompt "06-write-antigravity.md" "01-general-prompts/03-read-write/01-write-antigravity.md"
Copy-Prompt "07-write-memory.md" "01-general-prompts/03-read-write/03-write-memory.md"
Copy-Prompt "08-plan-coding-guideline-audit.md" "01-general-prompts/04-coding-guidelines/01-plan-coding-guideline-audit.md"
Copy-Prompt "09-execute-coding-guideline-fix.md" "01-general-prompts/04-coding-guidelines/02-execute-coding-guideline-fix.md"
Copy-Prompt "10-execute-batched-loop.md" "01-general-prompts/07-execute/03-execute-batched-loop.md"
Copy-Prompt "11-inventory-pending-tasks.md" "01-general-prompts/07-execute/04-inventory-pending-tasks.md"
Copy-Prompt "12-fix-subtask-naming-convention.md" "01-general-prompts/07-execute/04-fix-subtask-naming-convention.md"
Copy-Prompt "13-ci-cd-fix.md" "01-general-prompts/08-ci-cd/01-ci-cd-fix.md"
Copy-Prompt "14-cicd-run-ps1.md" "01-general-prompts/08-ci-cd/02-cicd-run-ps1.md"
Copy-Prompt "15-commit-fix.md" "01-general-prompts/05-commit-and-multi-agent-code-fix/01-commit-fix.md"
Copy-Prompt "16-commit-fix-v2.md" "01-general-prompts/05-commit-and-multi-agent-code-fix/05-commit-fix-v2.md"
Copy-Prompt "17-clean-artifacts-and-git-history.md" "01-general-prompts/05-commit-and-multi-agent-code-fix/09-clean-artifacts-and-git-history.md"
Copy-Prompt "18-release.md" "01-general-prompts/03-release-management/04-release.md"

Write-Host "Cleaning up temp directory..."
Remove-Item -Recurse -Force $TempDir

Write-Host "Checking prompt drift..."
$env:PYTHONIOENCODING="utf-8"
python linter-scripts/check-prompts-loaded.py
Write-Host "Done!"
