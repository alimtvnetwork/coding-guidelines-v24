## [v6.81.0] - 2026-10-19

### Added
- **GitMap RC Heatmap Integration:** Full integration of `gitmap rc heatmap <status|tables|query|view|tree|live|sync>` across all execution prompts (`01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`, `01-prompts/15-cg-execute/01-execute-coding-guideline-fix.md`, `01-prompts/22-letterly/`, `01-prompts/23-cursor-prompts/`, and `01-prompts/27-muse-prompts/`).
- **Automatic RC Vault Creation:** Automated provisioning of split-database telemetry in `repo-cache` (`rc`) or `.ai-memory/heatmap.db` linked to the active mission and workspace folder if not already existing.
- **Roadmap-Connected Heatmap for Dynamic Join:** Roadmap linkage connecting `.ai-memory/plans/<slug>.md` to `gitmap rc heatmap tree` and `gitmap rc heatmap status`, allowing dynamically joining subagents (`gitmap agent subtask join`) to discover previous read/edit traces and resume without redundant reading steps.
- **Agent Task Existence Detection & Prompt MD5 Hash Verification:** Automated detection in `gitmap agent task init` reporting existing task ID, active subagents, subtasks status (`gitmap agent subtask ls`), and touched files. Ingestion and verification of prompt MD5 hash code to seamlessly resume identical tasks or update roadmaps for evolved directives.
- **Skills Synchronization:** Complete synchronization of `.agents/skills/` and `.cursor/skills/` for `execute-parent-task-with-n-steps-v6`, `gitmap`, and `muse-master-prompt`.

---

### Quick Install One-Liners (v6.81.0)

**PowerShell:**
```powershell
irm https://github.com/alimtvnetwork/coding-guidelines-v24/releases/download/v6.81.0/release-install.ps1 | iex
```

**Bash:**
```bash
curl -fsSL https://github.com/alimtvnetwork/coding-guidelines-v24/releases/download/v6.81.0/release-install.sh | bash
```

### Release Assets Included
- `coding-guidelines-v6.81.0.zip`
- `coding-guidelines-v6.81.0.tar.gz`
- `release-install.ps1`
- `release-install.sh`
- `checksums.txt`
