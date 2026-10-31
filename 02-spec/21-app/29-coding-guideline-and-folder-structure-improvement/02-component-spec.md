# Component Specification: Coding Guideline and Folder Structure Improvement

> **Document:** `02-spec/21-app/29-coding-guideline-and-folder-structure-improvement/02-component-spec.md`  
> **Status:** Draft / Active  
> **Scope:** Module Contracts, Data Structures, Package Interfaces, and Registry Schemas  

---

## 1. Component Boundaries

### Component 1: `ai-memory-router` (`.ai-memory/what-to-read.md`)
- **Responsibility:** Single authoritative entrypoint guiding AI agents through required reading order before touching code, planning, or releasing.
- **Contract:**
  - Strict 5-tier layout with rationale for each item.
  - Zero raw `git` commands; instructs high-speed `gitmap` inspection.
  - IDE-detectable skill syntax: `[skill-name](file;.agents/skills/<skill-name>)`.
  - Date updated to current October 2026 milestone.
  - Version: `5.0.0`.

### Component 2: `folder-structure-spec` (`.ai-memory/folder-structure.md`)
- **Responsibility:** Canonical single source of truth for repository directory layout and `.ai-memory/` ecosystem.
- **Contract:**
  - Authoritative directory tree containing all root tiers: `01-prompts/`, `02-spec/`, `03-ai-scripts/`, `04-code/`, `05-changes-history/`, `06-archive/`, `.ai-memory/`, `.agents/`, `.cursor/`, `src/`, `slides-app/`.
  - Elimination of phantom files (`01-overview.md`, `05-plan.md`, `ai-fix-scripts/`).
  - Strict relative path mandate and lowercase file naming.
  - Inclusion of Tier 25 (`02-spec/25-spec-audits/`).

### Component 3: `applogger-syncbridge` (`04-code/golang/pkg/applogger/syncbridge/`)
- **Responsibility:** Asynchronous multi-sink coordination and context-aware flush bridge for `applogger`.
- **Structs & Types:**
  - `BridgeConfig`:
    - `QueueCapacity int`: Buffered channel size (default: 4096).
    - `WorkerCount int`: Background consumer goroutines (default: 2).
    - `FlushTimeout time.Duration`: Max time allowed for multi-sink `Sync()` (default: 2s).
    - `IsDropOnFull bool`: Drop on buffer saturation vs blocking backpressure.
  - `EntryWithContext`:
    - `Context context.Context`: Carries tracing headers (`traceId`, `spanId`).
    - `Entry applogger.LogEntry`: Structured log entry.
    - `Done chan struct{}`: Synchronous acknowledgment channel.
  - `MultiSinkBridge`:
    - `NewMultiSinkBridge(cfg BridgeConfig, sinks ...applogger.LogSink) *MultiSinkBridge`
    - `DispatchContext(ctx context.Context, entry applogger.LogEntry) *appfault.AppError`
    - `Flush(ctx context.Context) *appfault.AppError`
    - `Close() *appfault.AppError`

### Component 4: `skills-catalog-governor` (`.agents/skills/` & `.cursor/skills/`)
- **Responsibility:** Active AI capability registry.
- **Contract:**
  - Active directory retains only canonical, non-deprecated skills.
  - Purge 41 deprecated/stub skills from `.agents/skills/` and `.cursor/skills/` into `.agents/skills-archived/` and `.cursor/skills-archived/`.
  - Ensure all 8 `gitmap` skills are verified and prominent.

### Component 5: `multi-repo-sync-guard` (`03-ai-scripts/38-sync-prompts-skills-scripts.py`)
- **Responsibility:** Cross-repository synchronization engine.
- **Contract:**
  - `EXCLUDE_NAMES` set includes:
    - `"05-changes-history"`
    - `"06-archive"`
    - `"skills-archived"`
    - `"12-old-plan-prompts"`
    - `"18-insults"`
    - `"19-old-execute-prompts"`
  - Mirror script in `.agents/scripts/38-sync-prompts-skills-scripts.py` kept in 100% parity.

### Component 6: `main-worker-improvement-catalog` (`02-spec/19-main-worker-service/100-improvement-possibilities.md`)
- **Responsibility:** Normative catalog of 100 concrete, grounded improvement possibilities for the main worker service.
- **Contract:**
  - 10 functional clusters × 10 items per cluster = 100 items.
  - Each item details: Title, Current State & Deficiency, Concrete Improvement, and Architectural Benefit.
  - Registered in `02-spec/19-main-worker-service/readme.md`.
