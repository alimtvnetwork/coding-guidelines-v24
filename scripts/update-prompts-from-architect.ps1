$ErrorActionPreference = "Stop"

$ConfigPath = Join-Path -Path $PSScriptRoot -ChildPath "prompt-sync-config.json"
if (-not (Test-Path -Path $ConfigPath)) {
    Write-Error "Config file not found at $ConfigPath"
    exit 1
}

$Config = Get-Content -Path $ConfigPath | ConvertFrom-Json
$RepoUrl = $Config.repositoryUrl
$OutDir = $Config.outDir

$TempDir = Join-Path -Path $env:TEMP -ChildPath "prompt-architect-v2-$(Get-Random)"

Write-Host "Cloning $RepoUrl into temp directory..."
git clone --depth 1 $RepoUrl $TempDir

if (-not (Test-Path -Path $OutDir)) {
    New-Item -ItemType Directory -Path $OutDir | Out-Null
}
Remove-Item -Path "$OutDir\*.md" -Force -ErrorAction SilentlyContinue

foreach ($Mapping in $Config.mappings) {
    $SourceFile = Join-Path -Path $TempDir -ChildPath $Mapping.source
    $OutFile = Join-Path -Path $OutDir -ChildPath $Mapping.target
    if (Test-Path -Path $SourceFile) {
        Write-Host "Copying $($Mapping.target)..."
        Copy-Item -Path $SourceFile -Destination $OutFile -Force
    } else {
        Write-Warning "Source file not found: $SourceFile"
    }
}

Write-Host "Cleaning up temp directory..."
Remove-Item -Recurse -Force $TempDir

Write-Host "Checking prompt drift..."
$env:PYTHONIOENCODING="utf-8"
python linter-scripts/check-prompts-loaded.py
Write-Host "Done!"
