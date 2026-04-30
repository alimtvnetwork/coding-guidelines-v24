<#
.SYNOPSIS
  Installer for the fix-repo toolkit.

.DESCRIPTION
  Downloads from github.com/alimtvnetwork/coding-guidelines-v19 (branch: main)
  and installs into the current working directory:
    - fix-repo.sh, fix-repo.ps1
    - fix-repo.config.json
    - scripts/fix-repo/*.sh and *.ps1 helpers
    - fix-repo-contract.md (spec MD at root)

  After install, if CWD is inside a git repo, runs `./fix-repo.ps1 -DryRun`
  as a safe preview. If not a git repo, skips the run and exits 0.

.PARAMETER NoRun
  Install only — do not run the dry-run preview.

.PARAMETER Branch
  Branch to download from. Defaults to 'main'.

.EXAMPLE
  iwr -useb https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v19/main/get-fixrepo.ps1 | iex
  ./get-fixrepo.ps1
  ./get-fixrepo.ps1 -NoRun
  ./get-fixrepo.ps1 -Branch dev
#>

[CmdletBinding()]
param(
    [switch]$NoRun,
    [string]$Branch = 'main'
)

$ErrorActionPreference = 'Stop'

$Repo    = 'alimtvnetwork/coding-guidelines-v19'
$RawBase = "https://raw.githubusercontent.com/$Repo/$Branch"

function Get-Manifest {
    @(
        @{ Remote = 'fix-repo.sh';                                  Local = 'fix-repo.sh' },
        @{ Remote = 'fix-repo.ps1';                                 Local = 'fix-repo.ps1' },
        @{ Remote = 'fix-repo.config.json';                         Local = 'fix-repo.config.json' },
        @{ Remote = 'scripts/fix-repo/repo-identity.sh';            Local = 'scripts/fix-repo/repo-identity.sh' },
        @{ Remote = 'scripts/fix-repo/file-scan.sh';                Local = 'scripts/fix-repo/file-scan.sh' },
        @{ Remote = 'scripts/fix-repo/rewrite.sh';                  Local = 'scripts/fix-repo/rewrite.sh' },
        @{ Remote = 'scripts/fix-repo/config.sh';                   Local = 'scripts/fix-repo/config.sh' },
        @{ Remote = 'scripts/fix-repo/RepoIdentity.ps1';            Local = 'scripts/fix-repo/RepoIdentity.ps1' },
        @{ Remote = 'scripts/fix-repo/FileScan.ps1';                Local = 'scripts/fix-repo/FileScan.ps1' },
        @{ Remote = 'scripts/fix-repo/Rewrite.ps1';                 Local = 'scripts/fix-repo/Rewrite.ps1' },
        @{ Remote = 'scripts/fix-repo/Config.ps1';                  Local = 'scripts/fix-repo/Config.ps1' },
        @{ Remote = 'spec/02-coding-guidelines/06-cicd-integration/08-fix-repo-and-installers/01-fix-repo-contract.md'; Local = 'fix-repo-contract.md' }
    )
}

function New-ParentDir {
    param([string]$Path)
    $dir = Split-Path -Parent $Path
    if (-not $dir) { return }
    if (Test-Path -LiteralPath $dir) { return }
    New-Item -ItemType Directory -Force -Path $dir | Out-Null
}

function Save-RemoteFile {
    param([string]$Remote, [string]$Local)
    New-ParentDir -Path $Local
    $url = "$RawBase/$Remote"
    Invoke-WebRequest -Uri $url -OutFile $Local -UseBasicParsing
}

function Install-AllFiles {
    $count = 0
    foreach ($entry in Get-Manifest) {
        Save-RemoteFile -Remote $entry.Remote -Local $entry.Local
        $count++
    }
    Write-Host "get-fixrepo: installed $count file(s) from $Repo@$Branch"
}

function Test-IsGitRepo {
    $null = & git rev-parse --show-toplevel 2>$null
    return ($LASTEXITCODE -eq 0)
}

function Invoke-DryRunPreview {
    if ($NoRun) {
        Write-Host "get-fixrepo: -NoRun set; skipping preview"
        return
    }
    if (-not (Test-IsGitRepo)) {
        Write-Host "get-fixrepo: not a git repository — files installed; skipping fix-repo run"
        return
    }
    Write-Host ""
    Write-Host "get-fixrepo: running './fix-repo.ps1 -DryRun' as a safe preview…"
    Write-Host ""
    & (Join-Path -Path (Get-Location) -ChildPath 'fix-repo.ps1') -DryRun
}

function Invoke-Main {
    Install-AllFiles
    Invoke-DryRunPreview
}

Invoke-Main
exit 0
