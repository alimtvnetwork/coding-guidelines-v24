## [v6.80.0] - 2026-10-19

### Added
- **Generic Muse Master Prompt:** Added `01-prompts/27-muse-prompts/02-muse-generic.md` tailored for any codebase with zero hardcoded repos or secrets.
- **Dynamic Repository Onboarding Gate:** Muse dynamically queries which repositories to clone/link into `~/workspace/repos/`.
- **GitMap Heatmap Integration:** Full support for `gitmap heatmap log` (cognitive telemetry with reasoning) and `gitmap agent heatmap` (file contention and lease hotspot detection).
- **Per-Repository "Pull First" Protocol:** Guaranteed conflict-free, clean trees via `gitmap pull` before file inspection or editing.
- **Companion Automation Scripts:** Cross-platform First-Time Setup (`setup-first-time.sh` / `.ps1`) and Daily Run (`run-every-time.sh` / `.ps1`) scripts.
- **Spec & Canonical Alignment:** Seamless alignment with `02-spec/21-app/<slug>/` and `.ai-memory/plans/<slug>.md`.

---

### Quick Install One-Liners (v6.80.0)

**PowerShell:**
```powershell
irm https://github.com/alimtvnetwork/coding-guidelines-v24/releases/download/v6.80.0/release-install.ps1 | iex
```

**Bash:**
```bash
curl -fsSL https://github.com/alimtvnetwork/coding-guidelines-v24/releases/download/v6.80.0/release-install.sh | bash
```

### Release Assets Included
- `coding-guidelines-v6.80.0.zip`
- `coding-guidelines-v6.80.0.tar.gz`
- `release-install.ps1`
- `release-install.sh`
- `checksums.txt`
