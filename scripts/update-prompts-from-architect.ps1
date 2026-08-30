$ErrorActionPreference = "Stop"

# Forwarding wrapper: Prompts are now compiled internally from .lovable/prompts/01-prompts-category/
& "$PSScriptRoot/update-prompts.ps1" @args
