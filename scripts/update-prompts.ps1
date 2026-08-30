$ErrorActionPreference = "Stop"

$RepoRoot = Resolve-Path (Join-Path -Path $PSScriptRoot -ChildPath "..")
$ConfigPath = Join-Path -Path $PSScriptRoot -ChildPath "prompt-sync-config.json"

if (-not (Test-Path -Path $ConfigPath)) {
    Write-Error "Config file not found at $ConfigPath"
    exit 1
}

$RawConfig = Get-Content -Path $ConfigPath -Raw
$Config = $RawConfig | ConvertFrom-Json

# Compile variables before execution
$CompiledJson = $RawConfig
if ($Config.variables) {
    foreach ($Prop in $Config.variables.PSObject.Properties) {
        $VarName = "\$\{$($Prop.Name)\}"
        $VarVal = $Prop.Value
        $CompiledJson = $CompiledJson -replace $VarName, $VarVal
    }
}

$CompiledConfig = $CompiledJson | ConvertFrom-Json

Write-Host "Compiling prompts from source category folders to flat structure..."

$SyncCount = 0
foreach ($Mapping in $CompiledConfig.mappings) {
    $SourcePath = Join-Path -Path $RepoRoot -ChildPath $Mapping.source
    $TargetPath = Join-Path -Path $RepoRoot -ChildPath $Mapping.target

    $TargetDir = Split-Path -Path $TargetPath -Parent
    if (-not (Test-Path -Path $TargetDir)) {
        New-Item -ItemType Directory -Path $TargetDir -Force | Out-Null
    }

    if (Test-Path -Path $SourcePath) {
        Copy-Item -Path $SourcePath -Destination $TargetPath -Force
        $SyncCount++
    } else {
        Write-Warning "Source prompt not found: $SourcePath"
    }
}

Write-Host "Successfully compiled and synced $SyncCount prompts."

Write-Host "Validating prompt registry index..."
$env:PYTHONIOENCODING="utf-8"
python linter-scripts/check-prompts-loaded.py

Write-Host "Done!"
