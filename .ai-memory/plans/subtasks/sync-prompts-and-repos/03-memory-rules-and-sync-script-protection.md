# Subtask 03: Memory Rules & Synchronization Invariants

> **Subtask Code:** `Task-03`  
> **Parent Plan:** `.ai-memory/plans/sync-prompts-and-repos.md`  
> **Status:** QUEUED  

---

## Deliverables

1. Update `.ai-memory/memory/learned/18-cross-repository-sync-rules.md` documenting the 5 Non-Negotiable Boundaries:
   - Spec 21 Exclusion: Never sync application-specific specs (`02-spec/21-*` through `25-*`).
   - Additive-Only AI Scripts: Never overwrite scripts that the target repository modified.
   - Bump Script Protection: Never overwrite version bump scripts.
   - Memory & Plans Protection: Never overwrite `.ai-memory/memory/` or `.ai-memory/plans/`.
   - Pre-Flight Pull & Safety Backup Branching (`backup/sync-<timestamp>`).
2. Sync `.ai-memory/what-to-read.md` with the updated rules.
