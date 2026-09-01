$ErrorActionPreference = "Stop"

# Forwarding wrapper: Prompts are now compiled internally from 01-prompts/
& "$PSScriptRoot/update-prompts.ps1" @args
