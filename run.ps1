<#
.SYNOPSIS
    Root-level convenience runner. Dispatches to lint or slides sub-commands.

.DESCRIPTION
    Sub-commands (positional first arg):
      (none)        → lint  (legacy default — git pull + Go validator on src/)
      lint          → same as no-args, but explicit
      slides        → git pull → build slides-app/ → preview → open in browser
      help          → print this table

    Spec: spec/15-distribution-and-runner/02-runner-contract.md

.EXAMPLE
    .\run.ps1
    .\run.ps1 lint -Path cmd -MaxLines 20
    .\run.ps1 slides
    .\run.ps1 help
#>

param(
    [Parameter(Position = 0)]
    [string]$Command = "",

    [string]$Path = "src",
    [switch]$Json,
    [int]$MaxLines = 15,
    [switch]$d
)

$ErrorActionPreference = "Stop"

function Show-Help {
    $helpFile = Join-Path $PSScriptRoot "scripts" "runner-help.ps.txt"
    Get-Content -LiteralPath $helpFile | ForEach-Object { Write-Host $_ }
}

function Invoke-Lint {
    $inner = Join-Path $PSScriptRoot "linter-scripts" "run.ps1"
    if (-not (Test-Path $inner)) {
        Write-Host "❌ Cannot find $inner" -ForegroundColor Red
        exit 1
    }
    $splat = @{ Path = $Path; MaxLines = $MaxLines }
    if ($Json) { $splat["Json"] = $true }
    if ($d)    { $splat["d"]    = $true }
    & $inner @splat
    exit $LASTEXITCODE
}

function Test-Command {
    param([string]$Name)
    return [bool](Get-Command $Name -ErrorAction SilentlyContinue)
}

function Assert-SlidesDir {
    param([string]$Dir)
    if (Test-Path $Dir) { return }
    Write-Host "❌ slides-app/ not found at $Dir" -ForegroundColor Red
    Write-Host "   See spec-slides/00-overview.md for the slides spec." -ForegroundColor Yellow
    exit 1
}

function Resolve-SlidesRunner {
    if (Test-Command "bun")  { return "bun" }
    if (Test-Command "pnpm") { return "pnpm" }
    Write-Host "❌ Need 'bun' or 'pnpm' on PATH to build slides-app." -ForegroundColor Red
    Write-Host "   Install bun:  irm bun.sh/install.ps1 | iex" -ForegroundColor Yellow
    exit 1
}

function Invoke-SlidesBuild {
    param([string]$Runner)
    Write-Host "▸ install dependencies..." -ForegroundColor Cyan
    & $Runner install
    if ($LASTEXITCODE -ne 0) { Write-Host "❌ install failed" -ForegroundColor Red; exit 1 }
    Write-Host "▸ build..." -ForegroundColor Cyan
    & $Runner run build
    if ($LASTEXITCODE -ne 0) { Write-Host "❌ build failed" -ForegroundColor Red; exit 1 }
}

function Wait-SlidesReady {
    param([string]$Url)
    for ($i = 0; $i -lt 20; $i++) {
        Start-Sleep -Milliseconds 500
        try {
            $r = Invoke-WebRequest -Uri $Url -UseBasicParsing -TimeoutSec 1 -ErrorAction Stop
            if ($r.StatusCode -lt 500) { return $true }
        } catch { }
    }
    return $false
}

function Start-SlidesPreview {
    param([string]$Runner)
    $url = "http://localhost:4173/"
    Write-Host "▸ start preview server (background)..." -ForegroundColor Cyan
    $preview = Start-Process -FilePath $Runner -ArgumentList @("run", "preview") -PassThru -NoNewWindow
    $ready = Wait-SlidesReady -Url $url
    if (-not $ready) { Write-Host "⚠️  preview not reachable at $url — opening browser anyway" -ForegroundColor Yellow }
    Write-Host "▸ opening $url" -ForegroundColor Cyan; Start-Process $url
    Write-Host ""; Write-Host "▸ slides — preview running. Press Ctrl-C to stop." -ForegroundColor Green; Write-Host ""
    Wait-Process -Id $preview.Id
}

function Invoke-Slides {
    Write-Host ""; Write-Host "▸ slides — building offline deck and opening in browser" -ForegroundColor Cyan; Write-Host ""
    $slidesDir = Join-Path $PSScriptRoot "slides-app"
    Assert-SlidesDir -Dir $slidesDir
    Write-Host "▸ git pull (best effort)..." -ForegroundColor Cyan
    try { git pull | Out-Host } catch { Write-Host "⚠️  git pull failed — continuing with local state" -ForegroundColor Yellow }
    $runner = Resolve-SlidesRunner
    Write-Host "▸ using package runner: $runner" -ForegroundColor Cyan
    Push-Location $slidesDir
    try { Invoke-SlidesBuild -Runner $runner; Start-SlidesPreview -Runner $runner }
    finally { Pop-Location }
}

function Invoke-Visibility {
    $inner = Join-Path $PSScriptRoot "visibility-change.ps1"
    if (-not (Test-Path $inner)) {
        Write-Host "❌ Cannot find $inner" -ForegroundColor Red
        exit 1
    }
    # Forward all remaining args verbatim (drop the leading 'visibility' token)
    $forward = @()
    if ($args.Count -gt 0) { $forward = $args }
    & $inner @forward
    exit $LASTEXITCODE
}

function Invoke-FixRepo {
    $inner = Join-Path $PSScriptRoot "fix-repo.ps1"
    if (-not (Test-Path $inner)) {
        Write-Host "❌ Cannot find $inner" -ForegroundColor Red
        exit 1
    }
    $forward = @()
    if ($args.Count -gt 0) { $forward = $args }
    & $inner @forward
    exit $LASTEXITCODE
}

switch ($Command.ToLower()) {
    ""           { Invoke-Lint }
    "lint"       { Invoke-Lint }
    "slides"     { Invoke-Slides }
    "visibility" { Invoke-Visibility @args }
    "fix-repo"   { Invoke-FixRepo @args }
    "help"       { Show-Help; exit 0 }
    "-h"      { Show-Help; exit 0 }
    "--help"  { Show-Help; exit 0 }
    "-?"      { Show-Help; exit 0 }
    default   {
        if ($Command.StartsWith("-")) {
            # Treat as a lint flag — re-route through lint with $Command as -Path-style fallback
            Invoke-Lint
        } else {
            Write-Host "❌ Unknown command: $Command" -ForegroundColor Red
            Show-Help
            exit 2
        }
    }
}
