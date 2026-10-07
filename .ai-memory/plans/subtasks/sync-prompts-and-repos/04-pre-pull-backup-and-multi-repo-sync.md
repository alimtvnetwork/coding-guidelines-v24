# Subtask 04: Multi-Repository Pull, Safety Backup & Deployment

> **Subtask Code:** `Task-04`  
> **Parent Plan:** `.ai-memory/plans/sync-prompts-and-repos.md`  
> **Status:** QUEUED  

---

## Deliverables

1. Verify target repository states using `python 03-ai-scripts/38-sync-prompts-skills-scripts.py`.
2. Ensure each repository pulls latest changes from base branch.
3. Ensure each repository creates and pushes a safety backup branch (`backup/sync-<timestamp>`).
4. Mirror canonical prompts (`01-prompts/`), agent skills (`.agents/skills/`, `.cursor/skills/`), and shared specs (`02-spec/01-*` to `02-spec/20-*`).
5. Enforce additive-only AI scripts and bump script protection.
