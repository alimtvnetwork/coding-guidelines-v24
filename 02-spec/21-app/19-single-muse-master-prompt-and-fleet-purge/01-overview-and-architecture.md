# Architecture Specification: Single Muse Master Prompt & Fleet-Wide Deprecated Prompt Purge

> Version: 1.0.0
> Date: 2026-10-17
> Status: Normative
> Scope: Category `27-muse-prompts`, Multi-Repo Synchronization & Deprecated File Cleanup

---

## 1. Executive Summary

Having unified all onboarding steps, fleet scanning, repository cloning, memory synchronization, and V6 multi-agent task execution into `01-muse-master-prompt.md`, retaining fragmented or mirror prompt files (`02-muse-execute-in-a-step.md`, `03-muse-master-with-repos.md`, `04-muse-master-with-repos-v2.md`, and `05-muse-master-prompt-v2.md`) created confusion, ambiguity, and directory clutter for developers and AI agents.

This specification formalizes:
1. **Single Prompt Primacy**: `01-muse-master-prompt.md` is the only active prompt file permitted in `01-prompts/27-muse-prompts/`.
2. **Permanent Purge**: Complete deletion of all redundant files (`02`, `03`, `04`, `05`) from version control.
3. **Reference Harmonization**: Rewriting all cross-references in `01-prompts/readme.md`, `.ai-memory/prompts.md`, companion skills (`muse-master-prompt`), and specifications to point exclusively to `01-muse-master-prompt.md`.
4. **Fleet-Wide Purge Loop**: Codifying automated unlinking of the 4 deprecated prompts within `03-ai-scripts/38-sync-prompts-skills-scripts.py` so that all 43 connected repositories are cleanly purged during synchronization without deleting repository-owned assets.

---

## 2. Directory Layout & Target State

### Canonical State in `01-prompts/27-muse-prompts/`

```text
01-prompts/27-muse-prompts/
├── 01-muse-master-prompt.md       # SOLE authoritative unified master onboarding & execution prompt
└── readme.md                      # Directory index & catalog
```

### Purged File Manifest
- `01-prompts/27-muse-prompts/02-muse-execute-in-a-step.md` (DELETED)
- `01-prompts/27-muse-prompts/03-muse-master-with-repos.md` (DELETED)
- `01-prompts/27-muse-prompts/04-muse-master-with-repos-v2.md` (DELETED)
- `01-prompts/27-muse-prompts/05-muse-master-prompt-v2.md` (DELETED)

---

## 3. Fleet Synchronization Mechanics

In `03-ai-scripts/38-sync-prompts-skills-scripts.py` and `.agents/scripts/38-sync-prompts-skills-scripts.py`:
During `sync_repo()`, prior to mirroring directories:
```python
# Purge deprecated Muse prompts consolidated into 01-muse-master-prompt.md
for dep_muse in (
    "01-prompts/27-muse-prompts/02-muse-execute-in-a-step.md",
    "01-prompts/27-muse-prompts/03-muse-master-with-repos.md",
    "01-prompts/27-muse-prompts/04-muse-master-with-repos-v2.md",
    "01-prompts/27-muse-prompts/05-muse-master-prompt-v2.md",
):
    dep_path = target / dep_muse
    if dep_path.exists():
        try:
            dep_path.unlink()
        except Exception:
            pass
```

This guarantees that every repository synchronized will have these deprecated prompt files cleanly unlinked, while preserving all repository-owned files (`-0`).
