<#
.SYNOPSIS
  Write readme.txt locally (no GitHub raw fetch).

.DESCRIPTION
  Generates ./readme.txt with the line:

      let's start now DD-MMM-YYYY HH:MM:SS AM/PM

  using the current Malaysia (Asia/Kuala_Lumpur) date and 12-hour time.
  Runs entirely from the local repo — does NOT download anything.

.PARAMETER Path
  Output path. Defaults to ./readme.txt in the current directory.

.EXAMPLE
  .\scripts\write-readme-txt.ps1
  .\scripts\write-readme-txt.ps1 -Path .\readme.txt
#>

[CmdletBinding()]
param(
    [string]$Path = (Join-Path (Get-Location) 'readme.txt')
)

$ErrorActionPreference = 'Stop'

function Get-MalaysiaNow {
    $tz  = [System.TimeZoneInfo]::FindSystemTimeZoneById('Singapore Standard Time')
    return [System.TimeZoneInfo]::ConvertTimeFromUtc([DateTime]::UtcNow, $tz)
}

function Format-ReadmeLine {
    param([DateTime]$Now)
    $months = @('Jan','Feb','Mar','Apr','May','Jun','Jul','Aug','Sep','Oct','Nov','Dec')
    $date = '{0:D2}-{1}-{2}' -f $Now.Day, $months[$Now.Month - 1], $Now.Year
    $time = $Now.ToString('hh:mm:ss tt', [System.Globalization.CultureInfo]::InvariantCulture)
    return "let's start now $date $time"
}

$now  = Get-MalaysiaNow
$line = Format-ReadmeLine -Now $now

$dir = Split-Path -Parent $Path
if ($dir -and -not (Test-Path -LiteralPath $dir)) {
    New-Item -ItemType Directory -Force -Path $dir | Out-Null
}

Set-Content -LiteralPath $Path -Value $line -Encoding UTF8 -NoNewline
Add-Content -LiteralPath $Path -Value "`n" -NoNewline -Encoding UTF8

Write-Host "wrote: $Path"
Get-Content -LiteralPath $Path
