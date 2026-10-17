## Quick Install v6.75.0

### Windows (PowerShell)

```powershell
Invoke-WebRequest -Uri https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v24/v6.75.0/install.ps1 -OutFile install.ps1; .\install.ps1 -TargetDir ".ai-memory/prompts" -Version "v6.75.0"
```

### Unix / Linux / macOS (Bash)

```bash
curl -sL https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v24/v6.75.0/install.sh | bash -s -- ".ai-memory/prompts" "v6.75.0"
```

---

## What's Changed in v6.75.0

### Added
- letterly-muse - git pull first, ssh file delivery, heatmap logging, dual installers
