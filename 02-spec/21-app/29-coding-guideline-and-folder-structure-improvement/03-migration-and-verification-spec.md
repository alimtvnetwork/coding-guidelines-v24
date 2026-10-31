# Migration and Verification Specification: Coding Guideline and Folder Structure Improvement

> **Document:** `02-spec/21-app/29-coding-guideline-and-folder-structure-improvement/03-migration-and-verification-spec.md`  
> **Status:** Draft / Active  
> **Scope:** Migration Protocols, Invariant Verification, Automated Checks, and Rollback Procedures  

---

## 1. Migration Strategy (Phased Safe Execution)

The remediation is executed in bounded phases to ensure stability without affecting ongoing builds or running processes:

```
Phase 1: Planning, Research & Spec Authoring (Active - Exit 0)
Phase 2: Archival, Sanitization & Cleanups (Skills, Prompts, App Folders, .gitkeep)
Phase 3: Implementation & Invariant Verification (SyncBridge, What-To-Read, 100 Improvements, Sync Script)
Phase 4: Telemetry Closure & Atomic GitMap Commit
```

### 1.1 App Spec Cleanup Protocol
1. Remove obsolete 100-improvement prompt generator directories:
   - `02-spec/21-app/01-create-prompts-for-100-improvements/`
   - `02-spec/21-app/02-detail-plans-and-improvements-for-review/`
   - `02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills/`
   - `02-spec/21-app/04-next-100-ideas-for-improvements/`
   - `02-spec/21-app/05-improvements-folder-prompts/`
2. Remove corresponding pending plans:
   - `.ai-memory/plans/pending/01-create-prompts-for-100-improvements.md`
   - `.ai-memory/plans/pending/02-detail-plans-and-improvements-for-review.md`
   - `.ai-memory/plans/pending/03-rename-and-sync-100-improvements-prompts-and-skills.md`
   - `.ai-memory/plans/pending/04-next-100-ideas-for-improvements.md`
   - `.ai-memory/plans/pending/05-improvements-folder-prompts.md`
3. Remove subtask directories under `.ai-memory/plans/subtasks/` matching those slugs.
4. Clean temporary agent run folders in `.ai-memory/temp-agents/` matching those slugs.

### 1.2 Skills Archival Protocol
1. Audit all 41 identified legacy skills.
2. Ensure copy exists in `.agents/skills-archived/` and `.cursor/skills-archived/`.
3. Safely remove legacy skills from `.agents/skills/` and `.cursor/skills/`.
4. Validate that active skills directory contains only current production skills and all 8 `gitmap` skills.

### 1.3 `.gitkeep` Removal Protocol
1. Remove all 23 identified redundant `.gitkeep` files from non-empty folders.
2. Retain 5 legitimate `.gitkeep` files in truly empty placeholder folders.
3. Update `.ai-memory/strictly-avoid.md` to formally prohibit `.gitkeep` in populated folders.

---

## 2. Verification Evidence Contracts

| Check | Target | Command / Check Method | Expected Result |
|:---:|:---|:---|:---|
| **V1** | What-to-Read Freshness | `gitmap cat .ai-memory/what-to-read.md` | Contains 5-tier layout, current October 2026 dates, zero broken links |
| **V2** | Folder Structure Parity | `gitmap cat .ai-memory/folder-structure.md` | Contains all 6 top-level numbered directories, zero phantom files |
| **V3** | App Folder Cleanliness | `gitmap lf 02-spec/21-app` | Zero stale `create 100` directories |
| **V4** | Pending Plans Cleanliness | `gitmap lf .ai-memory/plans/pending` | Stale plans removed |
| **V5** | `.gitkeep` Redundancy | `gitmap find ".gitkeep"` | Only 5 legitimate `.gitkeep` files remain |
| **V6** | `syncbridge` Package | `gitmap lf 04-code/golang/pkg/applogger/syncbridge` | Package files present with valid Go types |
| **V7** | Ambiguity Resolution | `gitmap lf .ai-memory/ambiguous-questions/02-ambiguity-resolved` | Resolution documented |
| **V8** | Skills Inventory | `gitmap lf .agents/skills` | Outdated skills purged; 8 GitMap skills present |
| **V9** | Sync Script Exclusion | `gitmap cat 03-ai-scripts/38-sync-prompts-skills-scripts.py` | `05-changes-history`, `06-archive`, and old prompts in `EXCLUDE_NAMES` |
| **V10** | 100 Improvements Doc | `gitmap cat 02-spec/19-main-worker-service/100-improvement-possibilities.md` | Exactly 100 grounded items across 10 clusters |
| **V11** | Zero Colon Commits | GitMap commit validation | `gitmap cpc "(GUIDELINE) - <summary>"` format |
