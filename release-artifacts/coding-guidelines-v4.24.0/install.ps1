<#
.SYNOPSIS
    Download spec/linters/scripts from a GitHub repo with rich install controls.

.DESCRIPTION
    Conforms to: spec/14-update/27-generic-installer-behavior.md

    Power-user flags:
      -Repo owner/repo            Override source repo
      -Branch main                Override branch (ignored if -Version given)
      -Version vX.Y.Z             Install a specific release tag (PINNED MODE, §4)
      -Folders spec,linters       Explicit folder list (subpaths OK: spec/14-update)
      -Dest C:\path               Install destination (default: cwd)
      -ConfigFile my-config.json  Use custom config file
      -Prompt                     Ask before overwriting each existing file (y/n/a/s)
      -Force                      Overwrite all existing files without prompting
      -DryRun                     Show what would change; write nothing
      -ListVersions               List available release tags and exit
      -ListFolders                List available top-level folders for the chosen ref and exit
      -NoProbe (-n,-NoLatest)     Skip the latest-version probe
      -NoDiscovery                Skip V→V+N parallel discovery (spec §5.3)
      -NoMainFallback             Skip main-branch fallback (spec §5.3)
      -Offline (-UseLocalArchive) Skip all network ops; require local archive
      -RunFixRepo                 After verify, execute fix-repo.ps1 so the repo
                                  is patched before the installer exits
                                  (env: INSTALL_RUN_FIX_REPO=1)

    EXIT CODES (spec §8):
      0  success
      1  generic failure
      2  offline mode required a network operation (or handshake mismatch)
      3  pinned release / asset not found (PINNED MODE only)
      4  verification failed (checksum / required-paths)
      5  inner installer / handoff rejected

.EXAMPLE
    .\install.ps1
    .\install.ps1 -Version v1.22.0 -Folders spec/14-update -Prompt
    .\install.ps1 -DryRun
    irm https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v19/main/install.ps1 | iex
#>

param(
    [string]$Repo = "",
    [string]$Branch = "",
    [string]$Version = "",
    [string]$Dest = "",
    [string]$ConfigFile = "install-config.json",
    [string[]]$Folders = @(),
    [switch]$Prompt,
    [switch]$Force,
    [switch]$DryRun,
    [switch]$ListVersions,
    [switch]$ListFolders,
    [Alias('n','NoLatest')]
    [switch]$NoProbe,
    [switch]$NoDiscovery,
    [switch]$NoMainFallback,
    [Alias('UseLocalArchive')]
    [switch]$Offline,
    [switch]$RunFixRepo,
    [Alias('y','AssumeYes')]
    [switch]$Yes,
    [switch]$RollbackOnFixRepoFailure,
    [switch]$FullRollback,
    [string]$LogDir = "",
    [switch]$ShowFixRepoLog,
    [int]$MaxFixRepoLogs = -1,
    [string]$PinnedByReleaseInstall = ""
)

# Env-var equivalent so wrapper scripts can opt in without threading the flag.
if (-not $RunFixRepo) {
    $envFlag = $env:INSTALL_RUN_FIX_REPO
    if ($envFlag -and @("1","true","TRUE","yes","YES") -contains $envFlag) { $RunFixRepo = $true }
}
if (-not $Yes) {
    $envYes = $env:INSTALL_FIX_REPO_YES
    if ($envYes -and @("1","true","TRUE","yes","YES") -contains $envYes) { $Yes = $true }
}
if (-not $RollbackOnFixRepoFailure) {
    $envRb = $env:INSTALL_ROLLBACK_ON_FIX_REPO_FAILURE
    if ($envRb -and @("1","true","TRUE","yes","YES") -contains $envRb) { $RollbackOnFixRepoFailure = $true }
}
if (-not $FullRollback) {
    $envFR = $env:INSTALL_FULL_ROLLBACK
    if ($envFR -and @("1","true","TRUE","yes","YES") -contains $envFR) { $FullRollback = $true }
}
if ($FullRollback) { $RollbackOnFixRepoFailure = $true }
if (-not $LogDir) { $LogDir = $env:INSTALL_LOG_DIR }
if (-not $LogDir) { $LogDir = "" }
if (-not $ShowFixRepoLog) {
    $envShow = $env:INSTALL_SHOW_FIX_REPO_LOG
    if ($envShow -and @("1","true","TRUE","yes","YES") -contains $envShow) { $ShowFixRepoLog = $true }
}
if ($MaxFixRepoLogs -lt 0) {
    $envMax = $env:INSTALL_MAX_FIX_REPO_LOGS
    if ($envMax -and ($envMax -match '^\d+$')) { $MaxFixRepoLogs = [int]$envMax } else { $MaxFixRepoLogs = 0 }
}

# Bookkeeping for rollback.
$Script:InstalledNew = New-Object System.Collections.Generic.List[string]
$Script:InstalledBackups = New-Object System.Collections.Generic.List[object]
$Script:RollbackDir = $null
$Script:PreFixRepoHead = $null

# Offline mode forbids any network operation (spec §5.3, §8 exit 2).
if ($Offline) {
    Write-Host "    ❌ Offline mode is not yet supported by install.ps1. Exit 2 per spec §8." -ForegroundColor Red
    exit 2
}

$ErrorActionPreference = "Stop"
$ProgressPreference = "SilentlyContinue"

$script:Indent = "    "
function Write-Step  { param([string]$Msg) Write-Host "$script:Indent▸ $Msg" -ForegroundColor Cyan }
function Write-OK    { param([string]$Msg) Write-Host "$script:Indent✅ $Msg" -ForegroundColor Green }
function Write-Warn  { param([string]$Msg) Write-Host "$script:Indent⚠️  $Msg" -ForegroundColor Yellow }
function Write-Err   { param([string]$Msg) Write-Host "$script:Indent❌ $Msg" -ForegroundColor Red }
function Write-Dim   { param([string]$Msg) Write-Host "$script:Indent$Msg" -ForegroundColor DarkGray }
function Write-Plain { param([string]$Msg) Write-Host "$script:Indent$Msg" -ForegroundColor White }

function Write-InstallFailure {
    param([System.Management.Automation.ErrorRecord]$ErrorRecord)
    Write-Host ""
    Write-Host "════════════════════════════════════════════════════════" -ForegroundColor Red
    Write-Host "❌ INSTALLER FAILED — diagnostic report" -ForegroundColor Red
    Write-Host "════════════════════════════════════════════════════════" -ForegroundColor Red
    Write-Host "Message  : $($ErrorRecord.Exception.Message)" -ForegroundColor Red
    Write-Host "Type     : $($ErrorRecord.Exception.GetType().FullName)" -ForegroundColor Red
    if ($ErrorRecord.InvocationInfo -and $ErrorRecord.InvocationInfo.PositionMessage) {
        Write-Host "Location :" -ForegroundColor Red
        Write-Host $ErrorRecord.InvocationInfo.PositionMessage -ForegroundColor DarkGray
    }
    if ($ErrorRecord.ScriptStackTrace) {
        Write-Host "Script stack trace:" -ForegroundColor Red
        Write-Host $ErrorRecord.ScriptStackTrace -ForegroundColor DarkGray
    }
}

trap {
    Write-InstallFailure -ErrorRecord $_
    exit 1
}

if ($Prompt -and $Force) {
    Write-Err "-Prompt and -Force are mutually exclusive"
    exit 1
}

# ── Latest-version probe (skipped for -Version / listings / -NoProbe) ──
$script:ProbeOwner   = "alimtvnetwork"
$script:ProbeBase    = "coding-guidelines"
$script:ProbeVersion = 19

function Invoke-LatestVersionProbe {
    Write-Step "Detecting installer identity..."
    $sourceUrl = $null
    if ($MyInvocation.ScriptName)            { $sourceUrl = $MyInvocation.ScriptName }
    if (-not $sourceUrl -and $PSCommandPath) { $sourceUrl = $PSCommandPath }
    if (-not $sourceUrl -and $env:INSTALL_PROBE_SOURCE_URL) { $sourceUrl = $env:INSTALL_PROBE_SOURCE_URL }
    $owner = $env:INSTALL_PROBE_OWNER; $base = $env:INSTALL_PROBE_BASE; $cur = $env:INSTALL_PROBE_VERSION
    $urlRegex = '^https?://[^/]+/(?<o>[^/]+)/(?<b>[A-Za-z0-9._-]+?)-v(?<v>\d+)/[^/]+/install\.ps1'
    if ($sourceUrl -and $sourceUrl -match $urlRegex) {
        if (-not $owner) { $owner = $Matches.o }
        if (-not $base)  { $base  = $Matches.b }
        if (-not $cur)   { $cur   = $Matches.v }
    }
    if (-not $owner) { $owner = $script:ProbeOwner }
    if (-not $base)  { $base  = $script:ProbeBase }
    if (-not $cur)   { $cur   = $script:ProbeVersion }
    [int]$current = [int]$cur
    Write-OK "Identity: $owner/$base-v$current  (probing v$($current+1)..v$($current+20))"
    [int]$depth = 0
    if ($env:INSTALL_PROBE_HANDOFF_DEPTH) { [int]::TryParse($env:INSTALL_PROBE_HANDOFF_DEPTH, [ref]$depth) | Out-Null }
    if ($depth -ge 3) { Write-Err "Probe loop guard (depth=$depth)"; exit 1 }
    Write-Step "Probing 20 candidate versions in parallel (timeout 2s, middle-out)..."
    # Middle-out ordering: probe the middle of the range first, then expand
    # outward. With true parallelism this doesn't change correctness, but it
    # makes early-abort heuristics terminate faster when the latest version
    # tends to sit in the middle of the +1..+20 window.
    $low  = $current + 1
    $high = $current + 20
    $mid  = [int][Math]::Floor(($low + $high) / 2)
    $candidates = @($mid)
    for ($offset = 1; $offset -le ($high - $low); $offset++) {
        $upper = $mid + $offset
        $lower = $mid - $offset
        if ($upper -le $high) { $candidates += $upper }
        if ($lower -ge $low)  { $candidates += $lower }
    }
    Add-Type -AssemblyName System.Net.Http -ErrorAction SilentlyContinue
    $handler = [System.Net.Http.HttpClientHandler]::new()
    $client  = [System.Net.Http.HttpClient]::new($handler)
    $client.Timeout = [TimeSpan]::FromSeconds(2)
    $tasks = @{}
    foreach ($n in $candidates) {
        $url = "https://raw.githubusercontent.com/$owner/$base-v$n/main/install.ps1"
        $req = [System.Net.Http.HttpRequestMessage]::new([System.Net.Http.HttpMethod]::Head, $url)
        $tasks[$n] = $client.SendAsync($req)
    }
    $hits = @()
    # Iterate highest → lowest so the first hit we keep is already the winner.
    foreach ($n in ($candidates | Sort-Object -Descending)) {
        try {
            $r = $tasks[$n].GetAwaiter().GetResult()
            if ($r.IsSuccessStatusCode) { $hits += $n }
        } catch { }
    }
    $client.Dispose()
    $hits   = @($hits | Sort-Object -Descending)
    $latest = if ($hits.Count -gt 0) { $hits[0] } else { $current }
    if ($latest -le $current) { Write-OK "Already on latest (v$current)."; return }
    $newerUrl = "https://raw.githubusercontent.com/$owner/$base-v$latest/main/install.ps1"
    Write-OK "Newer version found: v$latest. Handing off..."
    $env:INSTALL_PROBE_HANDOFF_DEPTH = ($depth + 1).ToString()
    $env:INSTALL_PROBE_SOURCE_URL    = $newerUrl
    try { Invoke-RestMethod -Uri $newerUrl | Invoke-Expression; exit $LASTEXITCODE }
    catch { Write-Warn "Hand-off failed: $($_.Exception.Message)." }
}

# Pinning handshake: when invoked by release-install.ps1, the version
# arg MUST agree with the handshake value. Mismatch = exit 2.
if ($PinnedByReleaseInstall) {
    if (-not $Version) {
        $Version = $PinnedByReleaseInstall
    } elseif ($Version -ne $PinnedByReleaseInstall) {
        Write-Err "Pinning handshake mismatch: -Version=$Version vs -PinnedByReleaseInstall=$PinnedByReleaseInstall"
        exit 2
    }
    Write-Step "Pinned by release-install: $PinnedByReleaseInstall (auto-update disabled)"
}

$skipProbe = $NoProbe -or $Version -or $ListVersions -or $ListFolders -or $env:INSTALL_NO_PROBE -or $PinnedByReleaseInstall
if (-not $skipProbe) {
    try { Invoke-LatestVersionProbe } catch { Write-Warn "Probe error: $($_.Exception.Message)." }
}

# ── Read config (defaults only) ───────────────────────────────────
function Read-InstallConfig {
    param([string]$Path)
    if (-not (Test-Path $Path)) { return $null }
    return (Get-Content -Path $Path -Raw | ConvertFrom-Json)
}

$config = $null
if (Test-Path $ConfigFile) {
    Write-Step "Reading config from $ConfigFile"
    $config = Read-InstallConfig -Path $ConfigFile
}

if ([string]::IsNullOrEmpty($Repo)) {
    $Repo = if ($config -and $config.repo) { $config.repo } else { "alimtvnetwork/coding-guidelines-v19" }
}
if ([string]::IsNullOrEmpty($Branch)) {
    $Branch = if ($config -and $config.branch) { $config.branch } else { "main" }
}
if ([string]::IsNullOrEmpty($Dest)) { $Dest = (Get-Location).Path }
if ($Folders.Count -eq 0) {
    $Folders = if ($config -and $config.folders) { @($config.folders) } else { @("spec", "linters", "linter-scripts", ".lovable/coding-guidelines") }
}

$ref = if ($Version) { $Version } else { $Branch }

# ── Listing modes ─────────────────────────────────────────────────
function Show-ReleaseVersions {
    Write-Step "Fetching releases for $Repo..."
    try {
        $rels = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases?per_page=50" -UseBasicParsing
        Write-Host ""
        $rels | Select-Object -First 50 | ForEach-Object { Write-Plain "  • $($_.tag_name)" }
        Write-Host ""
    } catch {
        Write-Err "Could not fetch releases: $($_.Exception.Message)"
        exit 1
    }
    exit 0
}

function Show-TopFolders {
    Write-Step "Listing folders for $Repo@$ref..."
    try {
        $items = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/contents?ref=$ref" -UseBasicParsing
        Write-Host ""
        $items | Where-Object { $_.type -eq "dir" } | Sort-Object name | ForEach-Object { Write-Plain "  • $($_.name)" }
        Write-Host ""
    } catch {
        Write-Err "Could not list folders: $($_.Exception.Message)"
        exit 1
    }
    exit 0
}

if ($ListVersions) { Show-ReleaseVersions }
if ($ListFolders)  { Show-TopFolders }

# ── Banner (spec §7) ──────────────────────────────────────────────
$installMode = if ($Version) { "pinned" } else { "implicit" }
$sourceKind  = if ($Version) { "tag-tarball" } else { "branch-tarball" }
$versionLabel = if ($Version) { $Version } else { "$Branch (implicit)" }
Write-Host ""
Write-Plain "    📦 Spec & Scripts Installer"
Write-Plain "       mode:    $installMode"
Write-Plain "       repo:    $Repo"
Write-Plain "       version: $versionLabel"
Write-Plain "       source:  $sourceKind"
Write-Plain "       folders: $($Folders -join ', ')"
Write-Plain "       dest:    $Dest"
if ($DryRun)         { Write-Plain "       opts:    DRY-RUN (no writes)" }
if ($Prompt)         { Write-Plain "       opts:    Interactive prompts (y/n/a/s)" }
if ($Force)          { Write-Plain "       opts:    Force overwrite" }
if ($NoDiscovery)    { Write-Plain "       opts:    -NoDiscovery (V→V+N forbidden)" }
if ($NoMainFallback) { Write-Plain "       opts:    -NoMainFallback" }
Write-Host ""

# ── Download archive at ref ───────────────────────────────────────
$tmpDir = Join-Path ([System.IO.Path]::GetTempPath()) ("install-" + [guid]::NewGuid().ToString("N").Substring(0,8))
New-Item -ItemType Directory -Path $tmpDir -Force | Out-Null
$archivePath = Join-Path $tmpDir "repo.zip"

$archiveUrl = if ($Version) {
    "https://codeload.github.com/$Repo/zip/refs/tags/$Version"
} else {
    "https://codeload.github.com/$Repo/zip/refs/heads/$Branch"
}

try {
    Write-Step "Downloading $Repo@$ref..."
    try { Invoke-WebRequest -Uri $archiveUrl -OutFile $archivePath -UseBasicParsing }
    catch {
        $archiveUrl = "https://codeload.github.com/$Repo/zip/$ref"
        Invoke-WebRequest -Uri $archiveUrl -OutFile $archivePath -UseBasicParsing
    }

    Write-Step "Extracting..."
    $extractDir = Join-Path $tmpDir "extracted"
    Expand-Archive -Path $archivePath -DestinationPath $extractDir -Force
    $archiveRoot = Get-ChildItem -Path $extractDir -Directory | Select-Object -First 1
    if (-not $archiveRoot) { Write-Err "Failed to find archive root"; exit 1 }

    # ── Merge ─────────────────────────────────────────────────────
    $script:PromptAll = $false
    $script:PromptSkipAll = $false
    $copied = 0; $skippedFolders = 0; $wroteNew = 0; $overwrote = 0; $skippedFiles = 0

    function Test-ShouldOverwrite {
        param([string]$Target)
        if ($script:PromptAll)     { return $true }
        if ($script:PromptSkipAll) { return $false }
        if ($Force)                { return $true }
        if (-not $Prompt)          { return $true }
        while ($true) {
            $rel = $Target.Replace($Dest, '').TrimStart('\','/')
            $ans = Read-Host "? Overwrite $rel ? [y]es/[n]o/[a]ll/[s]kip-all"
            switch ($ans.ToLower()) {
                'y' { return $true }
                'n' { return $false }
                'a' { $script:PromptAll = $true; return $true }
                's' { $script:PromptSkipAll = $true; return $false }
                default { Write-Host "  enter y, n, a, or s" }
            }
        }
    }

    function Initialize-RollbackDir {
        if ($Script:RollbackDir) { return }
        $ts = (Get-Date).ToUniversalTime().ToString("yyyyMMddTHHmmssZ")
        $Script:RollbackDir = Join-Path $Dest ".install-rollback/$ts"
        New-Item -ItemType Directory -Path (Join-Path $Script:RollbackDir "backups") -Force | Out-Null
    }

    function Save-OverwriteBackup {
        param([string]$Target)
        Initialize-RollbackDir
        $rel = $Target.Replace($Dest, '').TrimStart('\','/')
        $backup = Join-Path $Script:RollbackDir "backups/$rel"
        New-Item -ItemType Directory -Path (Split-Path $backup -Parent) -Force | Out-Null
        Copy-Item -LiteralPath $Target -Destination $backup -Force
        $Script:InstalledBackups.Add([pscustomobject]@{ Target = $Target; Backup = $backup })
    }

    function Add-NewPath {
        param([string]$Target)
        Initialize-RollbackDir
        $Script:InstalledNew.Add($Target)
        Add-Content -LiteralPath (Join-Path $Script:RollbackDir "new-paths.txt") -Value $Target
    }

    function Merge-File {
        param([string]$Src, [string]$Target)
        $targetDir = Split-Path $Target -Parent
        $rel = $Target.Replace($Dest, '').TrimStart('\','/')
        if (Test-Path $Target) {
            if (Test-ShouldOverwrite -Target $Target) {
                if ($DryRun) { Write-Dim "  ~ would overwrite $rel" }
                else {
                    if ($FullRollback) { Save-OverwriteBackup -Target $Target }
                    New-Item -ItemType Directory -Path $targetDir -Force | Out-Null
                    Copy-Item $Src $Target -Force
                }
                $script:overwrote++
            } else {
                Write-Dim "  - skip $rel"; $script:skippedFiles++
            }
        } else {
            if ($DryRun) { Write-Dim "  + would create $rel" }
            else {
                New-Item -ItemType Directory -Path $targetDir -Force | Out-Null
                Copy-Item $Src $Target -Force
                if ($FullRollback) { Add-NewPath -Target $Target }
            }
            $script:wroteNew++
        }
    }

    foreach ($folder in $Folders) {
        $srcPath = Join-Path $archiveRoot.FullName $folder
        if (-not (Test-Path $srcPath)) {
            Write-Warn "Folder '$folder' not found in $Repo@$ref — skipping"
            $skippedFolders++; continue
        }
        Write-Step "Merging: $folder"
        Get-ChildItem -Path $srcPath -Recurse -File | ForEach-Object {
            $relativePath = $_.FullName.Substring($srcPath.Length).TrimStart('\','/')
            $targetFile = Join-Path (Join-Path $Dest $folder) $relativePath
            Merge-File -Src $_.FullName -Target $targetFile
        }
        $copied++
    }

    # Top-level files: copy each from archive root into Dest. Missing files
    # are warned (not fatal) so installer remains forward-compatible.
    $topLevelFiles = @("fix-repo.sh", "fix-repo.ps1", "visibility-change.sh", "visibility-change.ps1")
    foreach ($tlf in $topLevelFiles) {
        $srcFile = Join-Path $archiveRoot.FullName $tlf
        if (-not (Test-Path $srcFile)) {
            Write-Warn "Top-level file '$tlf' not found in $Repo@$ref — skipping"
            continue
        }
        Write-Step "Merging file: $tlf"
        Merge-File -Src $srcFile -Target (Join-Path $Dest $tlf)
    }

    # ── Verify required files (spec §8: exit 4 on missing required path) ──
    # Required files MUST exist in Dest after install. Skipped under -DryRun.
    if (-not $DryRun) {
        $requiredFiles = @("fix-repo.sh", "fix-repo.ps1")
        $missing = @()
        foreach ($rf in $requiredFiles) {
            if (-not (Test-Path (Join-Path $Dest $rf))) { $missing += $rf }
        }
        if ($missing.Count -gt 0) {
            Write-Err ("Install verification FAILED — {0} required file(s) missing in {1}" -f $missing.Count, $Dest)
            foreach ($m in $missing) { Write-Host "     • $m" -ForegroundColor Red }
            Write-Host ""
            Write-Host "   The archive was downloaded but did NOT contain the expected" -ForegroundColor Red
            Write-Host "   top-level scripts. Re-run without -Version to fetch main, or" -ForegroundColor Red
            Write-Host "   pin to a release that includes fix-repo.{sh,ps1}." -ForegroundColor Red
            exit 4
        }
        Write-OK "Verified $($requiredFiles.Count) required file(s) present"
    }

    # ── Optional: auto-run fix-repo.ps1 so the repo is patched before exit ──
    # Gated by -RunFixRepo or INSTALL_RUN_FIX_REPO=1. Skipped under -DryRun.
    # Failures propagate as exit 5 per spec §8.
    if ((-not $DryRun) -and $RunFixRepo) {
        $fixScript = Join-Path $Dest "fix-repo.ps1"
        if (-not (Test-Path -LiteralPath $fixScript -PathType Leaf)) {
            Write-Err "-RunFixRepo: $fixScript not found after install."
            exit 5
        }
        if ($Yes) {
            Write-Host "  ▸ Auto-confirmed (-Yes / INSTALL_FIX_REPO_YES=1)" -ForegroundColor DarkGray
        } elseif ([Environment]::UserInteractive -and -not [Console]::IsInputRedirected) {
            Write-Host ""
            Write-Host "⚠️  About to run $fixScript" -ForegroundColor Yellow
            Write-Host "   This will rewrite versioned-repo-name tokens across tracked text files." -ForegroundColor Yellow
            $reply = Read-Host "Proceed? [y/N]"
            if ($reply -notmatch '^(y|Y|yes|YES)$') {
                Write-Host "fix-repo skipped by user — exiting with code 5." -ForegroundColor Yellow
                exit 5
            }
        } else {
            Write-Err "-RunFixRepo requires confirmation but session is non-interactive."
            Write-Err "   Re-run with -Yes (or INSTALL_FIX_REPO_YES=1) to bypass the prompt."
            exit 5
        }
        if ($RollbackOnFixRepoFailure) {
            $isGitRepo = $false
            try { & git -C $Dest rev-parse --git-dir 2>$null | Out-Null; if ($LASTEXITCODE -eq 0) { $isGitRepo = $true } } catch {}
            if (-not $isGitRepo) {
                Write-Warn "-RollbackOnFixRepoFailure: $Dest is not a git repo; rollback disabled."
                $RollbackOnFixRepoFailure = $false
            } else {
                $Script:PreFixRepoHead = (& git -C $Dest rev-parse HEAD 2>$null).Trim()
                Write-Step ("Rollback armed: HEAD={0}{1}" -f $Script:PreFixRepoHead, $(if ($FullRollback) { ', full-rollback=on' } else { '' }))
            }
        }
        if ($LogDir) {
            if ([System.IO.Path]::IsPathRooted($LogDir)) { $logDir = $LogDir }
            else { $logDir = Join-Path $Dest $LogDir }
        } else {
            $logDir = Join-Path $Dest ".install-logs"
        }
        if (-not (Test-Path -LiteralPath $logDir)) { New-Item -ItemType Directory -Path $logDir -Force | Out-Null }
        $ts = (Get-Date).ToUniversalTime().ToString("yyyyMMddTHHmmssZ")
        $logFile = Join-Path $logDir "fix-repo-$ts.log"
        @(
            "# fix-repo log",
            "# started:  $((Get-Date).ToUniversalTime().ToString('yyyy-MM-ddTHH:mm:ssZ'))",
            "# script:   $fixScript",
            "# dest:     $Dest",
            "# os:       $([System.Runtime.InteropServices.RuntimeInformation]::OSDescription)",
            "# shell:    PowerShell $($PSVersionTable.PSEdition) $($PSVersionTable.PSVersion)",
            "# uname:    $([System.Runtime.InteropServices.RuntimeInformation]::OSDescription) / $([System.Runtime.InteropServices.RuntimeInformation]::OSArchitecture)",
            "# cwd:      $((Get-Location).Path)",
            "# rollback: on-failure=$RollbackOnFixRepoFailure  full=$FullRollback",
            "# ──────────────────────────────────────────────────────────"
        ) | Set-Content -LiteralPath $logFile -Encoding UTF8
        Write-Host ""
        Write-Step "Running fix-repo: $fixScript"
        Write-Step "Log: $logFile"
        & $fixScript 2>&1 | Tee-Object -FilePath $logFile -Append
        $rc = $LASTEXITCODE
        Add-Content -LiteralPath $logFile -Value "# exit: $rc  finished: $((Get-Date).ToUniversalTime().ToString('yyyy-MM-ddTHH:mm:ssZ'))"
        if ($MaxFixRepoLogs -lt 0 -or -not ($MaxFixRepoLogs -is [int])) {
            Write-Step "Log pruning: SKIPPED (--max-fix-repo-logs=$MaxFixRepoLogs is not a non-negative integer)"
        } elseif ($MaxFixRepoLogs -eq 0) {
            Write-Step "Log pruning: DISABLED (--max-fix-repo-logs=0)"
        } else {
            $allLogs = @(Get-ChildItem -LiteralPath $logDir -Filter 'fix-repo-*.log' -File -ErrorAction SilentlyContinue |
                Sort-Object LastWriteTime -Descending)
            $stale = $allLogs | Select-Object -Skip $MaxFixRepoLogs
            $removedCount = 0
            foreach ($s in $stale) { Remove-Item -LiteralPath $s.FullName -Force -ErrorAction SilentlyContinue; $removedCount++ }
            $kept = $allLogs.Count - $removedCount
            Write-Step "Log pruning: --max-fix-repo-logs=$MaxFixRepoLogs | found=$($allLogs.Count) kept=$kept pruned=$removedCount dir=$logDir"
        }
        if ($ShowFixRepoLog) {
            Write-Host ""
            Write-Host "─── fix-repo log: $logFile ─────────────────────────────"
            Get-Content -LiteralPath $logFile | ForEach-Object { Write-Host $_ }
            Write-Host "─── end of log ──────────────────────────────────────────"
        }
        if ($rc -ne 0) {
            Write-Err "fix-repo.ps1 failed (exit $rc) — see $logFile"
            if ($RollbackOnFixRepoFailure) {
                Write-Host ""
                Write-Warn "═══ ROLLBACK TRIGGERED (fix-repo failed) ═══"
                Write-Warn "Rollback flags: -RollbackOnFixRepoFailure=$RollbackOnFixRepoFailure  -FullRollback=$FullRollback"
                Write-Step "Restoring tracked files from HEAD..."
                & git -C $Dest checkout -- . 2>&1 | Tee-Object -FilePath $logFile -Append
                Write-OK "fix-repo edits reverted"
                if ($FullRollback) {
                    Write-Step "Removing files created by this install run..."
                    $removed = 0
                    foreach ($p in $Script:InstalledNew) {
                        if (Test-Path -LiteralPath $p) { Remove-Item -LiteralPath $p -Force -ErrorAction SilentlyContinue; $removed++ }
                    }
                    Write-Step "Restoring overwritten files from backups..."
                    $restored = 0
                    foreach ($b in $Script:InstalledBackups) {
                        if (Test-Path -LiteralPath $b.Backup) { Copy-Item -LiteralPath $b.Backup -Destination $b.Target -Force; $restored++ }
                    }
                    Write-OK "Removed $removed new file(s); restored $restored overwritten file(s)"
                }
                Write-Warn "Rollback complete. Snapshot kept at: $($Script:RollbackDir)"
            } else {
                Write-Warn "Rollback: NOT TRIGGERED (-RollbackOnFixRepoFailure=$RollbackOnFixRepoFailure  -FullRollback=$FullRollback)"
            }
            exit 5
        }
        Write-Step "Rollback: not needed (fix-repo succeeded; flags: -RollbackOnFixRepoFailure=$RollbackOnFixRepoFailure -FullRollback=$FullRollback)"
        Write-OK "fix-repo completed (log: $logFile)"
    }

    # ── Summary ───────────────────────────────────────────────────
    Write-Host ""
    Write-Plain "════════════════════════════════════════════════════════"
    if ($copied -gt 0)         { Write-OK "$copied folder(s) processed" }
    if ($wroteNew -gt 0)       { Write-OK "$wroteNew new file(s)" }
    if ($overwrote -gt 0)      { Write-OK "$overwrote file(s) overwritten" }
    if ($skippedFiles -gt 0)   { Write-Warn "$skippedFiles file(s) skipped" }
    if ($skippedFolders -gt 0) { Write-Warn "$skippedFolders folder(s) missing in source" }
    if ($DryRun)               { Write-Warn "DRY-RUN — no changes written" }
    Write-Host ""
    Write-Plain "  Source:      $Repo @ $ref"
    Write-Plain "  Destination: $Dest"
    Write-Plain "  Folders:     $($Folders -join ', ')"
    Write-Host ""
    Write-Plain "════════════════════════════════════════════════════════"
}
catch {
    Write-InstallFailure -ErrorRecord $_
    if ($_.Exception.StackTrace) {
        Write-Host ""
        Write-Host ".NET stack trace:" -ForegroundColor Red
        Write-Host $_.Exception.StackTrace -ForegroundColor DarkGray
    }
    if ($_.Exception.InnerException) {
        Write-Host ""
        Write-Host "Inner exception: $($_.Exception.InnerException.Message)" -ForegroundColor Red
    }
    Write-Host ""
    Write-Host "Context:" -ForegroundColor Yellow
    Write-Host "  Repo    : $Repo" -ForegroundColor DarkGray
    Write-Host "  Ref     : $ref" -ForegroundColor DarkGray
    Write-Host "  Dest    : $Dest" -ForegroundColor DarkGray
    Write-Host "  Folders : $($Folders -join ', ')" -ForegroundColor DarkGray
    Write-Host "  TmpDir  : $tmpDir" -ForegroundColor DarkGray
    Write-Host ""
    exit 1
}
finally {
    if (Test-Path $tmpDir) {
        Remove-Item -Path $tmpDir -Recurse -Force -ErrorAction SilentlyContinue
    }
    if (-not (Test-Path $tmpDir)) {
        Write-OK "Temp cleaned: $tmpDir"
    } else {
        Write-Warn "Temp NOT fully removed: $tmpDir (manual cleanup recommended)"
    }
}
