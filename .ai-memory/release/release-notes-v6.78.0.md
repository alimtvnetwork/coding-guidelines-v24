## Quick Install v6.78.0

### Windows (PowerShell)

```powershell
Invoke-WebRequest -Uri https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v24/v6.78.0/install.ps1 -OutFile install.ps1; .\install.ps1 -TargetDir ".ai-memory/prompts" -Version "v6.78.0"
```

### Unix / Linux / macOS (Bash)

```bash
curl -sL https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v24/v6.78.0/install.sh | bash -s -- ".ai-memory/prompts" "v6.78.0"
```

---

## What's Changed in v6.78.0

### Added
- letterly-mobile - place input text immediately after header colon before checklist
