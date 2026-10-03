## Quick Install v6.67.0

### Windows (PowerShell)

```powershell
Invoke-WebRequest -Uri https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v24/v6.67.0/install.ps1 -OutFile install.ps1; .\install.ps1 -TargetDir ".ai-memory/prompts" -Version "v6.67.0"
```

### Unix / Linux / macOS (Bash)

```bash
curl -sL https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v24/v6.67.0/install.sh | bash -s -- ".ai-memory/prompts" "v6.67.0"
```

---

## What's Changed in v6.67.0

### Added
- Restructure letterly and cursor prompts with ide skill format and plan enqueueing
