# Subtask 04: Dual Installers & PowerShell Automation

## Objective
Implement two separate installers: (1) full coding guidelines installer (`install.ps1`, `install-config.json`), and (2) lightweight general repository installer (`general-install.ps1`, `general-install.sh`, `general-install-config.json`, `install.ps1 -General`).

## Targets
- `install.ps1`
- `install.sh`
- `general-install.ps1`
- `general-install.sh`
- `general-install-config.json`

## Acceptance Criteria
- `general-install-config.json` specifies strictly `.ai-memory`, `01-prompts`, `03-ai-scripts`, `.agents/skills`, `.cursor/skills`, `02-spec/01-spec-authoring-guide`, `02-spec/21-app`, and `assets` folder structure.
- Total exclusion of all linters from general installation.
- `install.ps1 -General` and `general-install.ps1` run successfully and install only general assets.
