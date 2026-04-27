<#
.SYNOPSIS  Token-rewrite engine for fix-repo.ps1.
#>

$ErrorActionPreference = 'Stop'

function Get-TargetVersions {
    param([int]$Current, [int]$Span)
    $start = [Math]::Max(1, $Current - $Span)
    $end   = $Current - 1
    if ($end -lt $start) { return @() }
    return $start..$end
}

function Get-RewritePattern {
    param([string]$Base, [int]$N)
    $escaped = [regex]::Escape("$Base-v$N")
    return "$escaped(?!\d)"
}

function Invoke-FileRewrite {
    param(
        [string]$FullPath,
        [string]$Base,
        [int[]]$Targets,
        [int]$Current,
        [bool]$DryRun
    )
    $original = [System.IO.File]::ReadAllText($FullPath)
    $updated  = $original
    $count    = 0
    foreach ($n in $Targets) {
        $pattern = Get-RewritePattern -Base $Base -N $n
        $replaced = [regex]::Replace($updated, $pattern, "$Base-v$Current")
        if ($replaced -ne $updated) {
            $count += ([regex]::Matches($updated, $pattern)).Count
            $updated = $replaced
        }
    }
    if ($count -eq 0) { return 0 }
    if (-not $DryRun) { [System.IO.File]::WriteAllText($FullPath, $updated) }
    return $count
}
