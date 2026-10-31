# Subtask 02: GitMap Kill Process Hygiene

- **Parent Plan:** [.ai-memory/plans/pending/32-coding-guideline-corrections-and-gitmap-skill-updates.md](.ai-memory/plans/pending/32-coding-guideline-corrections-and-gitmap-skill-updates.md)
- **Target Files:**
  - `.agents/skills/gitmap/skill.md`
  - `.cursor/skills/gitmap/skill.md`
  - `02-spec/02-coding-guidelines/`
- **Status:** Completed

---

## Action Items

1. Search for any remaining references to `Stop-Process` or `taskkill` in skills and guidelines.
2. Update the Process Management section to elevate `gitmap kill <pid|name|:port> [--force] [--tree]` as the primary cross-platform tool.
3. Document graceful vs forced termination semantics:
   - `gitmap kill <pid>`: graceful termination.
   - `gitmap kill <pid> -f`: forced immediate termination.
   - `gitmap kill <pid> -t`: process tree termination.
   - `gitmap kill :<port> -f`: terminate process listening on specified TCP port.
4. Codify the total ban on raw shell process termination (`taskkill`, `Stop-Process`, `kill -9`).
