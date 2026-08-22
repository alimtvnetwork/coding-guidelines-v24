$ErrorActionPreference = "Stop"

$TempDir = Join-Path -Path $env:TEMP -ChildPath "prompt-architect-v2-$(Get-Random)"
$RepoUrl = "https://github.com/alimtvnetwork/prompt-architect-v2.git"
$OutDir = ".lovable/prompts"

Write-Host "Cloning $RepoUrl into temp directory..."
git clone --depth 1 $RepoUrl $TempDir

if (-not (Test-Path -Path $OutDir)) {
    New-Item -ItemType Directory -Path $OutDir | Out-Null
}

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

Copy-Prompt "01-read-memory-enhanced.md" "01-general-prompts/03-read-write/02-read-memory-enhanced.md"
Copy-Prompt "02-write-antigravity.md" "01-general-prompts/03-read-write/01-write-antigravity.md"
Copy-Prompt "03-write-memory.md" "01-general-prompts/03-read-write/03-write-memory.md"
Copy-Prompt "04-plan-coding-guideline-audit.md" "01-general-prompts/04-coding-guidelines/01-plan-coding-guideline-audit.md"
Copy-Prompt "05-execute-coding-guideline-fix.md" "01-general-prompts/04-coding-guidelines/02-execute-coding-guideline-fix.md"
Copy-Prompt "06-commit-fix.md" "01-general-prompts/05-commit-and-multi-agent-code-fix/01-commit-fix.md"
Copy-Prompt "07-commit-fix-v2.md" "01-general-prompts/05-commit-and-multi-agent-code-fix/05-commit-fix-v2.md"
Copy-Prompt "08-release.md" "01-general-prompts/03-release-management/04-release.md"
Copy-Prompt "09-clean-artifacts-and-git-history.md" "01-general-prompts/05-commit-and-multi-agent-code-fix/09-clean-artifacts-and-git-history.md"

Write-Host "Cleaning up temp directory..."
Remove-Item -Recurse -Force $TempDir

Write-Host "Checking prompt drift..."
$env:PYTHONIOENCODING="utf-8"
python linter-scripts/check-prompts-loaded.py
Write-Host "Done!"
