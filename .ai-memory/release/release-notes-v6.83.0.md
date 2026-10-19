## [v6.83.0] - 2026-10-19

### Changed & Decoupled
- **Decoupled Step Budgets from Letterly Prompts:** Restructured `01-prompts/22-letterly/03-execute-n-steps-letterly.md` and `01-prompts/23-cursor-prompts/03-execute-n-steps-letterly-cursor.md` into generic, reusable prompt formatters. Removed hardcoded 400-step lifecycle budgets and persona specifics, delegating execution parameters and concurrency directly to canonical skills (`execute-parent-task-with-n-steps-v6`).
- **Eliminated Verbal Fillers & Text Compaction:** Streamlined prose and instruction bodies across all 11 formatters in `01-prompts/22-letterly/` and `01-prompts/23-cursor-prompts/`, stripping wordy filler phrases and redundant explanations for compact, high-density prompt generation.
- **Strict Separation of Concerns:** Retained execution engine configurations (step counts, swarm concurrency $A = 3$, Capacity = 6) and Muse-specific capabilities strictly within dedicated skill definitions and `01-prompts/27-muse-prompts/`.
- **Preserved GitMap Heatmap Telemetry:** Maintained standardized GitMap Heatmap telemetry directives (`gitmap heatmap discuss` / `gitmap heatmap log` / `gitmap heatmap live`) across all Letterly and Cursor formatters to avoid write collisions and track in-flight file touches.
- **Repository Health & Manifest Synchronization:** Ran full manifest synchronizer updating `version.json`, `package.json`, `prompt-version.template.json`, `public/health-score.json`, `readme.md`, and `src/data/specTree.json`.

---

### Quick Install One-Liners (v6.83.0)

**PowerShell:**
```powershell
irm https://github.com/alimtvnetwork/coding-guidelines-v24/releases/download/v6.83.0/release-install.ps1 | iex
```

**Bash:**
```bash
curl -fsSL https://github.com/alimtvnetwork/coding-guidelines-v24/releases/download/v6.83.0/release-install.sh | bash
```

### Release Assets Included
- `coding-guidelines-v6.83.0.zip`
- `coding-guidelines-v6.83.0.tar.gz`
- `release-install.ps1`
- `release-install.sh`
- `checksums.txt`
