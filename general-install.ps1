<#
.SYNOPSIS
    General Repository Asset Installer
    Installs .ai-memory, 01-prompts, 03-ai-scripts, .agents/skills, .cursor/skills,
    02-spec/01-spec-authoring-guide, 02-spec/21-app, and assets folder structure.
    Total exclusion of all linters and fix-repo scripts.
#>
param(
    [string]$Repo = "",
    [string]$Branch = "",
    [string]$Version = "",
    [string]$Dest = "",
    [string[]]$Folders = @(),
    [switch]$Prompt,
    [switch]$Force,
    [switch]$DryRun,
    [switch]$NoProbe
)

$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$installScript = Join-Path $scriptDir "install.ps1"

if (Test-Path $installScript) {
    & $installScript -General -Repo $Repo -Branch $Branch -Version $Version -Dest $Dest -Folders $Folders -Prompt:$Prompt -Force:$Force -DryRun:$DryRun -NoProbe:$NoProbe
} else {
    $webScript = "https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v24/main/install.ps1"
    $code = (Invoke-RestMethod -Uri $webScript -UseBasicParsing)
    $sb = [ScriptBlock]::Create($code)
    & $sb -General -Repo $Repo -Branch $Branch -Version $Version -Dest $Dest -Folders $Folders -Prompt:$Prompt -Force:$Force -DryRun:$DryRun -NoProbe:$NoProbe
}
