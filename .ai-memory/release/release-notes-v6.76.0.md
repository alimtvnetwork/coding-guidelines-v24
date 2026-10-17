## Quick Install v6.76.0

### Windows (PowerShell)

```powershell
Invoke-WebRequest -Uri https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v24/v6.76.0/install.ps1 -OutFile install.ps1; .\install.ps1 -TargetDir ".ai-memory/prompts" -Version "v6.76.0"
```

### Unix / Linux / macOS (Bash)

```bash
curl -sL https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v24/v6.76.0/install.sh | bash -s -- ".ai-memory/prompts" "v6.76.0"
```

---

## What's Changed in v6.76.0

### Added
- letterly-mobile - add task name and slug subheaders and heatmap commands
