# Architecture Specification: Coding Guideline and Folder Structure Improvement

> **Document:** `02-spec/21-app/29-coding-guideline-and-folder-structure-improvement/01-architecture-spec.md`  
> **Status:** Draft / Active  
> **Scope:** Repository Architecture, AI Memory Alignment, Skills Governance, Main Worker Service Enhancements, App Cleanliness  

---

## 1. Executive Overview & Problem Statement

The Prompt Architect repository maintains the central canonical standard for coding guidelines, multi-agent AI instructions, and toolchains across 43 connected repositories. Over recent development iterations, structural drift and obsolete artifacts have accumulated:

1. **Outdated AI Entrypoint (`.ai-memory/what-to-read.md`):** Stale timestamps (September 2026), deprecated V3/V4 prompt citations, dead links to renamed files (`standards/version-source-of-truth.md` instead of `04-version-source-of-truth.md`), and raw `git` command instructions for subagents.
2. **Canonical Folder Structure Drift (`.ai-memory/folder-structure.md`):** Lists phantom files (`01-overview.md`, `05-plan.md`), deprecated locations (`ai-fix-scripts/` instead of `03-ai-scripts/`, `prompts/` inside `.ai-memory/`), and omits root tiers (`01-prompts/`, `03-ai-scripts/`, `04-code/`, `05-changes-history/`, `06-archive/`).
3. **App Folder & Pending Plans Pollution (`02-spec/21-app/` and `.ai-memory/plans/`):** Obsolete prompt-generator specs (`01-create-prompts-for-100-improvements`, etc.) clutter `02-spec/21-app/`, and orphaned plan files exist in `.ai-memory/plans/pending/`.
4. **Redundant `.gitkeep` Files:** Over 20 `.gitkeep` files persist in non-empty directories that already contain tracked files.
5. **Open Ambiguity in Logging:** Unresolved architectural questions regarding logger sink synchronization (`ConsoleSink` vs `FileSink`) in `.ai-memory/ambiguous-questions/`.
6. **Skills & Prompts Synchronization Leakage:** Deprecated skills linger in active directories (`.agents/skills/` and `.cursor/skills/`), and sync scripts (`03-ai-scripts/38-sync-prompts-skills-scripts.py`) lack explicit exclusions for `05-changes-history`, `06-archive`, and legacy prompt folders, causing deprecated prompts to mirror to all 43 repositories.
7. **Main Worker Service Spec Enhancements:** `02-spec/19-main-worker-service/` requires a grounded catalog of 100 concrete improvement possibilities across 10 functional clusters.

---

## 2. Architectural Pillars

```
+-----------------------------------------------------------------------------------+
|                        CANONICAL REPOSITORY ARCHITECTURE                          |
+-----------------------------------------------------------------------------------+
|  01-prompts/          Active V6 prompts (Letterly, Muse, V6 N-Steps, CG, RCA)     |
|  02-spec/             Normative Specifications (Tiers 01-20 Core, 21-25 Apps)     |
|  03-ai-scripts/       Production Python Automation Suite (Shared Engine, Linters) |
|  04-code/             Enterprise Polyglot Packages (Go: appfault, applogger, etc.)|
|  05-changes-history/  Immutable Transaction Logs (Tasks 01-26, Isolated)          |
|  06-archive/          Historical Deprecated Prompts & Skills (Isolated)           |
|  .ai-memory/          Operational Metadata (Learned Memory, Plans, Heatmap DB)    |
|  .agents/ & .cursor/  AI Skills, Rules, and Scripts (Mirrored & Verified)         |
+-----------------------------------------------------------------------------------+
```

### Pillar 1: 5-Tier AI Reading Hierarchy
Modernize `.ai-memory/what-to-read.md` into a deterministic 5-tier sequence:
- **Tier 1: Safety & Invariants:** `AGENTS.md` and `.ai-memory/strictly-avoid.md`.
- **Tier 2: Identity & Baseline Contracts:** `.ai-memory/readme.md`, `version.json`, `.ai-memory/coding-guidelines.md`.
- **Tier 3: Institutional Memory:** `.ai-memory/memory/` and cross-repo sync boundaries.
- **Tier 4: Active Roadmap & Work State:** GitMap status, `.ai-memory/plans/readme.md`, and active parent plan.
- **Tier 5: Operational Tooling & Specifications:** Canonical V6 execution prompts, AI scripts, and normative specs.

### Pillar 2: Clean Separation of Active vs. Archived Assets
- Active prompts in `01-prompts/` strictly contain current production workflows (V6 loops, Letterly, Muse). Deprecated folders (`12-old-plan-prompts/`, `18-insults/`, `19-old-execute-prompts/`) are segregated into `06-archive/`.
- Active skills in `.agents/skills/` and `.cursor/skills/` strictly contain active workflows. Deprecated stubs are segregated into `.agents/skills-archived/` and `.cursor/skills-archived/`.
- Sync scripts (`38-sync-prompts-skills-scripts.py`) explicitly exclude `05-changes-history/`, `06-archive/`, `skills-archived/`, and old prompt folders from multi-repository mirroring.

### Pillar 3: `.gitkeep` Zero-Redundancy Invariant
A `.gitkeep` file exists solely to preserve otherwise empty directory trees in Git. Once any tracked file (markdown, source code, config, asset) exists within that directory, the `.gitkeep` file is redundant and MUST be removed.

### Pillar 4: Logger Multi-Sink Synchronization (`syncbridge`)
To resolve logger sink synchronization ambiguity, implement `04-code/golang/pkg/applogger/syncbridge/`:
- **Decoupled Asynchronous Workers:** Dedicated ring buffer queue with configurable worker pool, preventing slow I/O sinks (`FileSink`, `SQLiteSink`) from blocking fast sinks (`ConsoleSink`) or calling goroutines.
- **Distributed Context Tracing:** Native ingestion of `context.Context` extracting `traceId` and `spanId` into structured log metadata.
- **Non-Blocking Synchronized Flush:** Concurrent `Sync()` fan-out across all registered sinks with context deadline timeouts.
- **Leaf Error Sovereignty:** Strict `*appfault.AppError` returns eliminating unhandled errors.

### Pillar 5: Main Worker Service Spec Quality Matrix (100 Improvements)
A grounded inventory of 100 architectural, operational, security, and schema improvements for `02-spec/19-main-worker-service/` organized across 10 functional clusters:
- Cluster 01: Core Worker Architecture & Topology Hardening
- Cluster 02: Split-DB Schema, Multi-Tenant Isolation & Migration Integrity
- Cluster 03: Authentication, Session Tokens & 2FA/TOTP Resilience
- Cluster 04: Structured Error Handling, Leaf Error Sovereignty & AppError Integration
- Cluster 05: Backup Replication, Key Rotation & Disaster Recovery
- Cluster 06: Worker Registration, Heartbeat Leases & Health Checking
- Cluster 07: Subdomain Routing, Edge Proxying & Reverse Tunneling
- Cluster 08: Observability, Cognitive Heatmap Telemetry & Structured Logging
- Cluster 09: Cross-Language Coding Guidelines, Positive Booleans & Monadic Results
- Cluster 10: Chaos Engineering, Test Isolation & E2E Validation Boundaries

---

## 3. User Request (Verbatim)

```markdown
# Coding Guideline and Folder Structure Improvement: high priority instruction, non-negotiable task

In the coding guideline, there are several factors. Those are very wrong. Come with each one of them. One of the factors is that the what to read folder is not very professional or not very good, so you need to work on it. The second factor is that you need to work on this for the coding guideline. So for the coding guideline, the folder structure in the AI memory is not updated in the what to read. And also there are a lot of plan files in the coding guideline, which should be cleared for now. And also for the AI instruction, whenever a folder has a file, it should actually remove that .gitkeep file. So that needs to be mentioned carefully. In the ambiguity section, make sure that it is clear. In coding guideline, we don't need the ambiguity. So in cases you have an ambiguity question, so you ask for some problem issues like console sync, file sync. So whatever the problem that you have in the ambiguity question, can you please try to implement some of those by creating another package on top of it by connecting this sync and stuff? Can you please do that so that our logger and everything is perfect? That's another thing. So what to read, you need to correct and check that recently we should have the latest `gitmap` skills should be added in our skill section and also find outdated skills in the skill section for agents and also cursor. Try to remove the outdated skill and any old files, old prompt, try to move this to archive folder, and do not sync those and make sure those old prompts should be removed from the synced repositories. According to a guideline old file, try to remove that as well. In the what to read, and the folder structure, these are pretty old folder structure. It is not following the current folder structure. Inside the AI memory, you need to correct it very carefully right now. And also some of the repositories I saw that old folder structure remains. You need to correct it and if there is anything that we can improve in the folder structure, first suggest it. Do not sync outside for now. Let's fix inside in the coding guideline for now. Then we can synchronize later on. Rather than brute forcing, we will think carefully and then we will sync. Do you understand? Also, the main worker service that is spec, check that. Find 100 possibilities that could be improved inside that spec. And also the folder 21 contains create 100 prompt improvements. These are the files we don't need. The app folder for the coding guideline should be clean. Let's improve the coding guideline first and clear out the app issues, app Database (Database (Db)). All the index zero one or overview files should be README file, and that needs to be corrected everywhere. I do see still overview files and old prompts moved to archive folder. Do not sync those. And also the change history, do not sync. So these are the things I think we need to correct. Change history, we don't need to sync with other repositories. Remember that. Archive folder do not need to have. So inside the archive, all the old stuff, just put it there. So this should not be synced. Remember that. Can you please follow through?
```
