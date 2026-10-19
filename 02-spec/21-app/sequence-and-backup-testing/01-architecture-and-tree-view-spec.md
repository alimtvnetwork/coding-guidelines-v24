# Specification: Prompt Tree View Sequencing, Multi-Tier Hierarchy & Fast-Forward Strip

## 1. Overview & Objectives

In prompt management and developer automation environments, sequencing and visibility of prompt execution states are essential to developer flow. Users frequently encounter issues where queued, active, and completed prompts are interleaved, presented out of sequence, or cause user interface flickering during high-frequency status polling.

This specification establishes the canonical architecture and deterministic ordering invariants for the Prompt Tree View and the companion "Now Running + Up Next" fast-forward panel strip:

1. **Three-Tier Multi-Tier Prompt Tree Hierarchy**:
   - **Tier 1 (Apex):** Running prompts positioned at the highest rank (Rank 0) with an emerald badge and animated pulse.
   - **Tier 2 (Middle):** Queued prompts positioned immediately below running tasks (Rank 1) with an amber badge in strict FIFO order (`ORDER BY created_at ASC, id ASC`).
   - **Tier 3 (Base):** Historical / completed / terminated prompts positioned at the base (Rank 2) sorted descending by recency timestamp (`ORDER BY updated_at DESC, id DESC`).
2. **"Now Running + Up Next" Fast-Forward Panel Strip**:
   - High-visibility header strip exposing exactly 1 active running prompt and up to 4 queued prompts.
   - IPC interface `get_running_and_up_next` providing cached, flicker-free rendering with a 30-second TTL constant (`RUNNING_UP_NEXT_TTL_MS = 30000`).
3. **Deterministic Persistence & Retrieval**:
   - Single-query rank tiering against the active prompt store.
   - Optimistic state transitions on prompt submission, cancellation, and reordering.
4. **Strict Architectural Invariants**:
   - Strict relative Git paths across all references and documentation.
   - Strict lowercase filenames across all components.

---

## 2. Multi-Tier Prompt Tree Architecture

The Prompt Tree View displays all prompts belonging to the active workspace organized into three distinct, non-overlapping visual tiers:

```text
+-----------------------------------------------------------------------+
| PROMPT TREE VIEW                                                      |
+-----------------------------------------------------------------------+
| [T1: RUNNING (APEX)]                                                 |
|  [*] #042 - Build Auth Middleware (PID: 8836) [EMERALD PULSE]         |
+-----------------------------------------------------------------------+
| [T2: QUEUED (FIFO)]                                                   |
|  [1] #043 - Update SQLite Split-DB Engine (Created: 10:01:05) [AMBER] |
|  [2] #044 - Refactor Variadic Spread Params (Created: 10:02:18) [AMBER]|
|  [3] #045 - Verify SSH Fleet Reachability (Created: 10:04:30) [AMBER] |
+-----------------------------------------------------------------------+
| [T3: HISTORICAL / COMPLETED (RECENCY DESCENDING)]                     |
|  [v] #041 - Sync Coding Guidelines Prompts (Completed: 09:58:12)      |
|  [v] #040 - Fix RCA Pipeline Annotations (Completed: 09:45:00)        |
|  [x] #039 - Test Docker Container Build (Failed: 09:20:15)            |
+-----------------------------------------------------------------------+
```

### 2.1 Tier 1: Apex Running Prompt (Rank 0)

- **Selection Criteria:** Prompts with state `status = 'running'` or `state = 'in_progress'`.
- **Display Priority:** Rank 0 (strictly pinned to the top of the tree view).
- **Visual Presentation:**
  - Visual Badge: Emerald Badge (`bg-emerald-950/60 text-emerald-400 border border-emerald-500/40`).
  - Activity Indicator: Animated pulse ring (`animate-pulse`) signifying in-flight execution.
  - Telemetry Overlay: Elapsed execution time (`elapsed_ms`), active subagent count, and host identifier.
- **Concurrency Guard:** Exactly 1 prompt per active execution runner lane. If parallel execution lanes are configured, running slots are grouped at Rank 0 and ordered by `started_at ASC`.

### 2.2 Tier 2: Queued Prompts (Rank 1)

- **Selection Criteria:** Prompts with state `status = 'queued'` or `status = 'pending'`.
- **Display Priority:** Rank 1 (positioned immediately beneath the running prompt tier).
- **Ordering Guarantee:** Strict First-In, First-Out (FIFO) ordering:
  ```sql
  ORDER BY created_at ASC, id ASC
  ```
- **Visual Presentation:**
  - Visual Badge: Amber Badge (`bg-amber-950/60 text-amber-400 border border-amber-500/40`).
  - Queue Position Badge: Discrete sequence index number `[#1]`, `[#2]`, `[#3]`.
  - Timestamp Display: Human-readable submission time and relative wait duration.
- **Reordering Protocol:** Manual drag-and-drop reordering updates an explicit `queue_priority` index while defaulting un-prioritized items to insertion order.

### 2.3 Tier 3: Historical & Completed Prompts (Rank 2)

- **Selection Criteria:** Prompts with terminal states: `status IN ('completed', 'failed', 'cancelled', 'aborted')`.
- **Display Priority:** Rank 2 (positioned at the base of the tree view).
- **Ordering Guarantee:** Strict reverse-chronological order by recency:
  ```sql
  ORDER BY COALESCE(completed_at, updated_at) DESC, id DESC
  ```
- **Visual Presentation:**
  - Status Icons: Green checkmark for `completed`, red cross for `failed`, slate slash for `cancelled`.
  - Grouping: Collapsible accordion panels grouped by completion date (e.g., "Today", "Yesterday", "Older").

---

## 3. "Now Running + Up Next" Panel Strip

The "Now Running + Up Next" strip is a persistent, compact, horizontal overview widget located above the main workspace view.

### 3.1 Strip Layout & Content

- **Slot 0 (Running Slot):**
  - Displays the active prompt title, prompt ID, and animated emerald pulse.
  - If no prompt is currently running, displays `Idle - Ready for Prompts` with a dimmed slate indicator.
- **Slots 1 to 4 (Up Next Pipeline):**
  - Displays up to 4 queued prompts in strict FIFO order from Tier 2.
  - Each item shows queue index (`#1` through `#4`), title snippet, and amber pill badge.
- **Queue Overflow Counter:**
  - If total queued prompts exceed 4, an overflow badge displays `+N more queued` with a click action navigating to the full Prompt Tree View.

### 3.2 IPC Interface Contract: `get_running_and_up_next`

Frontend components fetch the strip state via the typed IPC call `get_running_and_up_next`:

```typescript
export interface PromptSummaryItem {
  id: string;
  title: string;
  slug: string;
  status: 'running' | 'queued' | 'completed' | 'failed' | 'cancelled';
  createdAt: number;
  startedAt?: number;
  completedAt?: number;
  rankTier: number;
  queuePosition?: number;
}

export interface RunningAndUpNextPayload {
  running: PromptSummaryItem | null;
  upNext: PromptSummaryItem[]; // Maximum 4 items, sorted FIFO
  totalQueuedCount: number;
  lastSyncTimestamp: number;
  cacheTtlMs: number;
}
```

### 3.3 Zero-Flicker Caching Strategy (`RUNNING_UP_NEXT_TTL_MS`)

To prevent micro-stutter and layout shifts during state transitions or background telemetry polling:

- **TTL Constant:** `RUNNING_UP_NEXT_TTL_MS = 30000` (30 seconds client-side cache validity window).
- **Stale-While-Revalidate (SWR):** UI immediately renders from cached state while asynchronously verifying queue modifications.
- **Event-Driven Invalidation:** The cache is explicitly invalidated ONLY upon state-altering events:
  1. `prompt:enqueued` (new prompt submitted)
  2. `prompt:started` (prompt transitioned from queued to running)
  3. `prompt:completed` (prompt transitioned from running to historical)
  4. `prompt:reordered` (user manually reordered queue)
  5. `prompt:cancelled` (prompt removed from queue)
- **DOM Stability:** Stable React keys (`key={item.id}`) and hardware-accelerated CSS opacity transitions eliminate layout recalculation and prevent Cumulative Layout Shift (CLS < 0.01).

---

## 4. Database Query Architecture & Schema Invariants

The Split-DB prompt engine executes a unified query utilizing explicit `rank_tier` projection:

```sql
SELECT
  id,
  title,
  slug,
  prompt_text,
  status,
  created_at,
  updated_at,
  started_at,
  completed_at,
  CASE
    WHEN status = 'running' THEN 0
    WHEN status IN ('queued', 'pending') THEN 1
    ELSE 2
  END AS rank_tier,
  ROW_NUMBER() OVER (
    PARTITION BY (CASE WHEN status IN ('queued', 'pending') THEN 1 ELSE 0 END)
    ORDER BY created_at ASC, id ASC
  ) AS queue_seq
FROM active_prompts
WHERE workspace_id = :workspace_id
ORDER BY
  rank_tier ASC,
  CASE WHEN rank_tier = 1 THEN created_at END ASC,
  CASE WHEN rank_tier = 2 THEN COALESCE(completed_at, updated_at) END DESC;
```

---

## 5. Verification Test Plan & Step-by-Step Validation Commands

To empirically verify the multi-tier ordering, FIFO queue sequencing, and flicker-free strip retrieval, execute the following commands:

### Step 1: Query Baseline Prompt Tree State
```bash
gitmap agy running-prompts ls
gitmap agy rp prompts ls
```
- **Assertion:** Verify output groups prompts cleanly by status (running at rank 0, queued at rank 1, completed at rank 2).

### Step 2: Enqueue Sequential Test Prompts
```bash
gitmap agy running-prompts add --title "Test Prompt Alpha" --slug "test-alpha"
gitmap agy running-prompts add --title "Test Prompt Beta" --slug "test-beta"
gitmap agy running-prompts add --title "Test Prompt Gamma" --slug "test-gamma"
```
- **Assertion:** All 3 test prompts appear in Tier 2 with amber badges in exact FIFO order: Alpha (`#1`), Beta (`#2`), Gamma (`#3`).

### Step 3: Verify "Now Running + Up Next" IPC Payload
```bash
gitmap agy rp get-running-and-up-next
```
- **Assertion:** Verify JSON output returns:
  - `running`: null or current active task.
  - `upNext`: array containing Alpha, Beta, Gamma in order.
  - `totalQueuedCount`: >= 3.
  - `cacheTtlMs`: 30000.

### Step 4: Promote First Queued Prompt to Running
```bash
gitmap agy running-prompts start --slug "test-alpha"
```
- **Assertion:**
  - Alpha immediately elevates to Tier 1 (Rank 0) with emerald badge and active pulse.
  - Beta promotes to queue position `#1` in Tier 2.
  - Gamma promotes to queue position `#2` in Tier 2.

### Step 5: Complete Running Prompt & Observe Historical Placement
```bash
gitmap agy running-prompts complete --slug "test-alpha"
```
- **Assertion:**
  - Alpha transitions to Tier 3 (Rank 2) positioned at the top of completed items (most recent completed timestamp).
  - Next queued prompt (Beta) automatically promotes to running or remains at head of queue.

---

## 6. Acceptance Criteria Matrix

| Criteria ID | Requirement | Verification Method |
|---|---|---|
| AC-TREE-001 | Running prompts strictly appear at the apex (Rank 0) with emerald badge and animated pulse | Inspect DOM tree order and CSS class `animate-pulse` |
| AC-TREE-002 | Queued prompts appear in Tier 2 in strict FIFO order (`created_at ASC, id ASC`) | Validate sequence index vs `created_at` timestamp |
| AC-TREE-003 | Completed prompts appear in Tier 3 sorted descending by recency timestamp | Verify descending timestamp ordering in Tier 3 |
| AC-TREE-004 | "Now Running + Up Next" panel strip exposes 1 running + up to 4 queued prompts | Inspect `get_running_and_up_next` JSON return payload |
| AC-TREE-005 | Cache TTL constant `RUNNING_UP_NEXT_TTL_MS` equals 30000 ms with event-driven invalidation | Verify configuration constant and SWR behavior |
| AC-TREE-006 | Zero UI flicker or layout shifts (CLS < 0.01) during queue updates | Performance audit and stable key verification |
| AC-TREE-007 | All repository links and citations use strictly relative Git paths | Linter check with zero absolute paths or `file:///` URIs |
| AC-TREE-008 | All authored specification and plan files adhere to strict lowercase naming | File system and path verification |
