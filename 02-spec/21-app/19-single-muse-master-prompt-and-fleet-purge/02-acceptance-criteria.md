# Acceptance Criteria: Single Muse Master Prompt & Fleet-Wide Deprecated Prompt Purge

> Version: 1.0.0
> Date: 2026-10-17
> Status: Normative

---

## Acceptance Criteria Matrix

### Scenario 1: Canonical Prompt Category Cleanliness
- **Given** The `01-prompts/27-muse-prompts/` directory in `coding-guidelines`
- **When** Auditing all files present in the folder
- **Then** Exactly two files exist: `01-muse-master-prompt.md` and `readme.md`
- **And** `02-muse-execute-in-a-step.md`, `03-muse-master-with-repos.md`, `04-muse-master-with-repos-v2.md`, and `05-muse-master-prompt-v2.md` do NOT exist

### Scenario 2: Cross-Reference & Skill Integrity
- **Given** Canonical prompt indexes and companion skills
- **When** Inspecting `01-prompts/readme.md`, `.ai-memory/prompts.md`, `.agents/skills/muse-master-prompt/skill.md`, and `.cursor/skills/muse-master-prompt/skill.md`
- **Then** No references to `05-muse-master-prompt-v2.md` or deleted files exist
- **And** Step 5 strictly references `01-muse-master-prompt.md`

### Scenario 3: Automated Fleet Unlinking
- **Given** The 43 connected target repositories
- **When** `03-ai-scripts/38-sync-prompts-skills-scripts.py` executes across the fleet
- **Then** Any stale Muse prompt files (`02`, `03`, `04`, `05`) are automatically unlinked
- **And** Zero repository-owned files are deleted
- **And** All target repositories receive the updated `01-muse-master-prompt.md`, `readme.md`, and companion skills

### Scenario 4: Upstream Git Completion
- **Given** The local working tree in `coding-guidelines` and all target repositories
- **When** Changes are finalized
- **Then** Working trees are completely clean
- **And** Commits are pushed upstream to remote tracking branches satisfying the "No Push = Not Done" invariant
