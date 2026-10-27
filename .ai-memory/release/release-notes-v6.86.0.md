## [v6.86.0] - 2026-10-27

### Added
- **New Skill `follow-coding-guidelines`:** Authored in `.agents/skills/follow-coding-guidelines/skill.md` and `.cursor/skills/follow-coding-guidelines/skill.md`, establishing the 12 non-negotiable coding pillars, strict lowercase filename enforcement, and `/learn` / `/plan` cognitive workflow loops.
- **GitMap Agent Task Spec Tracking:** Wired relative `--spec-file` (`-s`) and `--spec-folder` parameters into `gitmap task init`, storing relative spec boundaries directly in `ParentTask` table in `agent-task.db`.
- **GitMap Show Log UI Modernization:** Embedded authentic GitMap SVG logo emblem & favicon, added searchable branch combo filter, and two-commit compare diff mode with Beyond Compare / VS Code integration.
- **GitMap Context Menu Separation:** Distinguished folder context menu (`Directory\shell\gitmap`) from folder background context menu (`Directory\Background\shell\gitmap`), and added `Open Installer Manager UI` and `Settings (UI)` actions.

### Changed
- **GitMap Skill Sanitization:** Normalized commit examples to Title Case (`Skills`, `Cloner`, `Docs`) and lowercase, purged deprecated commands (`disk`, `os storage`, `os dev-clean`, `vscode repair`, `starbar client ls`, `cluster worker/status`), and updated Section 14 with a dedicated Heatmap Shorthand vs Direct SQL query matrix.
- **Skills Inventory Reduction:** Moved 8 redundant and prompt-wrapper skills (`gitmap-ssh-cluster-fleet`, `prompt-library`, `python-boolean-and-conditional-enhancement`, `python-constants-and-magic-number-elimination`, `python-dry-architecture-and-caching`, `release-management`, `release-now`, `release-orchestrator`) to `.agents/skills-archived/` and `.cursor/skills-archived/`.
- **GitMap Bare Slug Clone Resolution:** Enhanced `ResolveRepoSlug` to probe remote existence via `probeRemoteGitURL` before committing to constructed URLs, maintaining full suggestion catalog output for invalid bare slugs while enabling instant bare name cloning for existing repositories.

---

### Quick Install One-Liners (v6.86.0)

**PowerShell:**
```powershell
irm https://github.com/alimtvnetwork/coding-guidelines-v24/releases/download/v6.86.0/release-install.ps1 | iex
```

**Bash:**
```bash
curl -fsSL https://github.com/alimtvnetwork/coding-guidelines-v24/releases/download/v6.86.0/release-install.sh | bash
```

### Release Assets Included
- `coding-guidelines-v6.86.0.zip`
- `coding-guidelines-v6.86.0.tar.gz`
- `release-install.ps1`
- `release-install.sh`
- `checksums.txt`
