## Quick Install v6.47.0

### Windows (PowerShell)

```powershell
Invoke-WebRequest -Uri https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v24/v6.47.0/install.ps1 -OutFile install.ps1; .\install.ps1 -TargetDir ".ai-memory/prompts" -Version "v6.47.0"
```

### Unix / Linux / macOS (Bash)

```bash
curl -sL https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v24/v6.47.0/install.sh | bash -s -- ".ai-memory/prompts" "v6.47.0"
```

---

## What's Changed in v6.47.0

### Added
- upgrade prompts to v1/v2, elevate gitmap aum, sync skills, and fix linux directory permissions
