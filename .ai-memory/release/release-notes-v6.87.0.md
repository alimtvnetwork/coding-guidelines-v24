## [v6.87.0] - 2026-10-27

### Added
- **Master 39-Step Coding Guideline Orchestrator (`follow-cg-execute`):** Authored in `.agents/skills/follow-cg-execute/skill.md` and `.cursor/skills/follow-cg-execute/skill.md`. Sequences all 39 coding guideline compliance prompts in order, enforcing bounded 5–8 file fast-forwarding passes with zero intermediate builds/tests (steps 1–38) and terminal full builds, E2E test verification, and automated release ceremony (step 39).
- **GitMap Cognitive Heatmap Query Shorthands:** Wired dedicated CLI aliases (`gitmap heatmap recent`, `gitmap heatmap failed`, `gitmap heatmap hotfiles`, `gitmap heatmap tasks`) and companion `rc heatmap` subcommands, replacing verbose multi-line SQL queries.
- **GitMap Task Spec File & Folder Tracking:** Implemented `HeatmapParentTask` registration with `RootSpecFilePath` and `TaskFolder` in `store/heatmap_store.go` and wired into `gitmap task init` so relative spec boundaries are recorded in both `agent-task.db` and `heatmap.db`.

### Changed
- **PowerShell Completion Script Hardening:** Fixed Cobra completion script generator in `cli/cmd/root_cobra_completion.go` where trailing `+ $Space` was attached to `{ continue }`, resolving the `ScriptBlock op_Addition` exception and restoring permanent tab suggestions.
- **Native Process Termination (`gitmap kill`):** Promoted native compiled `gitmap kill <pid|name|:port> [--force] [--tree]` (alias `terminate`) across all skills and guidelines, codifying total ban on raw shell process termination (`Stop-Process`, `taskkill`).
- **Skills Inventory Consolidation:** Consolidated 13 redundant prompt wrappers into canonical parent skills (`release`, `logo-and-icon-design`, `react-ui-theming-design`, `prompts-and-skills-sync`, `check-pending-improvements`) and moved deprecated wrappers to `.agents/skills-archived/` and `.cursor/skills-archived/`.

---

### Quick Install One-Liners (v6.87.0)

**PowerShell:**
```powershell
irm https://github.com/alimtvnetwork/coding-guidelines-v24/releases/download/v6.87.0/release-install.ps1 | iex
```

**Bash:**
```bash
curl -fsSL https://github.com/alimtvnetwork/coding-guidelines-v24/releases/download/v6.87.0/release-install.sh | bash
```

### Release Assets Included
- `coding-guidelines-v6.87.0.zip`
- `coding-guidelines-v6.87.0.tar.gz`
- `release-install.ps1`
- `release-install.sh`
- `checksums.txt`
