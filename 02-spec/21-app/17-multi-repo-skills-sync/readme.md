# Multi-Repository Skills Synchronization & Non-Destructive Fleet Deployment

> Category: `21-app`  
> Status: Active  
> Architecture: Non-Destructive Multi-Repository Fleet Synchronization Engine  
> Target Subsystems: `03-ai-scripts/38-sync-prompts-skills-scripts.py`, `.agents/skills/sync/`, `.cursor/skills/sync/`, `.agents/skills/muse-master-prompt/`, `.agents/skills/gitmap/`, `01-prompts/27-muse-prompts/`, and all 43 registered repositories  
> Central Reference Spec: [`02-spec/21-app/17-multi-repo-skills-sync/`](./)  

---

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Verify the non-destructive mirroring invariant: zero deletions (`removed == 0`) across all target repositories.
- [ ] `/learn` Adhere strictly to the 5 Non-Negotiable Boundaries (Spec 21 Exclusion, Additive Scripts, Bump Script Protection, Memory & Plans Protection, Archive Exclusion).
- [ ] `/goal` Verify that pre-flight pull and safety backup branch creation (`backup/sync-<timestamp>`) execute prior to any file changes.
- [ ] `/learn` Enforce atomic commits with hyphen messages (`feat(sync) - ...`) and the mandatory 'No Push = Not Done' completion invariant.
- [ ] `/goal` Validate relative path hygiene via `python linter-scripts/check-relative-paths.py` (0 absolute paths or file protocol URIs).

---

## Overview

This specification category governs the autonomous, enterprise-wide synchronization of canonical prompt architectures, AI agent skills, shared coding guidelines, and automation scripts from the meta-repository `coding-guidelines` across the entire connected fleet of 43 target repositories.

The central pillar of this specification is the **Strict Non-Destructive Invariant**: target repositories frequently create and maintain local tools, custom scripts, localized test fixtures, custom prompts, and domain documentation. The synchronization engine must never delete, unlink, or purge any file created or updated by a target repository on its own. All synchronization operations are strictly additive or updates to canonical files, with guaranteed zero deletions (`removed == 0`).

---

## Documents

| Document | Purpose |
| :--- | :--- |
| [`01-overview-and-architecture.md`](01-overview-and-architecture.md) | High-level system architecture, user request verbatim, non-negotiable boundaries, payload inventory, and execution flow |
| [`02-acceptance-criteria.md`](02-acceptance-criteria.md) | Formal Gherkin acceptance criteria (AC-SYNC-001 through AC-SYNC-010) and full 43-repository verification matrix |
| [`03-non-destructive-sync-spec.md`](03-non-destructive-sync-spec.md) | Non-destructive mirroring technical specification (`mirror_directory` Step 1 bypass), 43 target repositories catalog, and branch detection protocol |

---

## Core Invariants & Boundaries

1. **Strict Non-Destructive Invariant for Repo-Owned Files:** Zero unlinks and zero directory deletions across all target repository trees. Target repositories retain 100% of locally authored files.
2. **Canonical Assets Update & Addition Only:** Only canonical assets originating from `coding-guidelines` are copied (if new) or updated (if changed).
3. **Boundary 1 (Spec 21 Exclusion):** `02-spec/21-*` through `02-spec/25-*` are private to each application and are never touched or synchronized.
4. **Boundary 2 (Additive-Only AI Scripts):** Scripts in `03-ai-scripts/` and `.agents/scripts/` that have local repository modifications are never overwritten. Brand-new scripts are added cleanly.
5. **Boundary 3 (Bump Script Protection):** Target version bump scripts (`bump*`, `version.json`, `bump-version.mjs`) are immutable.
6. **Boundary 4 (Memory & Plans Protection):** Target repository `.ai-memory/memory/` and `.ai-memory/plans/` are isolated and protected.
7. **Boundary 5 (Archive Exclusion):** `06-archive/` and `06-old-prompts/` are never synchronized downstream.
8. **Pre-Change Safety Protocol:** Pull latest changes on `<base_branch>`, create and push remote safety backup branch `backup/sync-<timestamp>`, and return to `<base_branch>`.
9. **Mandatory Completion Invariant ('No Push = Not Done'):** All commits, release tags, and merge commits must be pushed upstream to remote GitHub origins.
10. **Strict Relative Git Paths Mandate:** All file paths must be strictly relative Git paths. Absolute paths and file protocol URIs are totally banned.

---

## Cross-References

| Reference | Location |
| :--- | :--- |
| Application Specs Root Index | [`../readme.md`](../readme.md) |
| Multi-Repo Sync Script | [`../../../03-ai-scripts/38-sync-prompts-skills-scripts.py`](../../../03-ai-scripts/38-sync-prompts-skills-scripts.py) |
| Full-Fleet Sync Prompt | [`../../../01-prompts/24-sync/01-sync.md`](../../../01-prompts/24-sync/01-sync.md) |
| Sync Skill | [`../../../.agents/skills/sync/skill.md`](../../../.agents/skills/sync/skill.md) |
| Unified Muse Master Prompt | [`../../../01-prompts/27-muse-prompts/01-muse-master-prompt.md`](../../../01-prompts/27-muse-prompts/01-muse-master-prompt.md) |
| Muse Master Prompt Skill | [`../../../.agents/skills/muse-master-prompt/skill.md`](../../../.agents/skills/muse-master-prompt/skill.md) |
| Relative Path Verification Linter | [`../../../linter-scripts/check-relative-paths.py`](../../../linter-scripts/check-relative-paths.py) |
