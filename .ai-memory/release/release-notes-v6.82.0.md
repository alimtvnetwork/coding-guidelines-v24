## [v6.82.0] - 2026-10-19

### Added & Upgraded
- **Muse Prompt 400-Step Upgrade:** Upgraded `01-prompts/27-muse-prompts/01-muse-master-prompt.md` and `01-prompts/27-muse-prompts/02-muse-generic.md` to default to an aggressive 400-step lifecycle budget (`N = 400`).
- **Structured Phase Budgeting:** Systematic partitioning across execution phases:
  - Phase 1 (Steps 1..100): Discovery, prompt capture, architecture spec authoring (`02-spec/21-app/<slug>/`), and lean subtasks setup (`.ai-memory/plans/subtasks/<slug>/`).
  - Phase 2 (Steps 101..250): Intermediate archival, repo-cache setup, prerequisite refactoring, and file staging.
  - Phase 3 (Steps 251..400): Worker swarm parallel execution, write-and-execute implementation, targeted linter checks, and atomic commit push.
- **Powerful Step Execution Directive (Write-and-Execute Invariant):** Mandated across Muse prompts and skill mirrors that every step and turn must be high-impact, writing files directly to disk and executing validations immediately without conversational delays or permission-seeking.
- **Prompts-to-Skills Migration Audit & Parity:** Comprehensive repository audit confirming 100% parity across all 128 skills in `.agents/skills/` and `.cursor/skills/`, with valid YAML frontmatter and proper relative markdown references.
- **Multi-Repository Fleet Synchronization:** Fleet-wide synchronization across all connected repositories using compiled Go engine `gitmap sync`.

---

### Quick Install One-Liners (v6.82.0)

**PowerShell:**
```powershell
irm https://github.com/alimtvnetwork/coding-guidelines-v24/releases/download/v6.82.0/release-install.ps1 | iex
```

**Bash:**
```bash
curl -fsSL https://github.com/alimtvnetwork/coding-guidelines-v24/releases/download/v6.82.0/release-install.sh | bash
```

### Release Assets Included
- `coding-guidelines-v6.82.0.zip`
- `coding-guidelines-v6.82.0.tar.gz`
- `release-install.ps1`
- `release-install.sh`
- `checksums.txt`
