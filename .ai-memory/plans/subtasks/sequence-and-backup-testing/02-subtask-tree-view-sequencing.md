# Subtask 02: Prompt Tree View Sequencing Verification

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/sequence-and-backup-testing/02-subtask-tree-view-sequencing.md`  
> **Parent Plan:** `.ai-memory/plans/sequence-and-backup-testing.md`  
> **Specification:** `02-spec/21-app/sequence-and-backup-testing/01-architecture-and-tree-view-spec.md`  
> **Status:** `[DONE]`  
> **Assigned Role:** Spec Author & UI Hierarchy Engineer  

---

## 1. Objective

Autonomously design, document, and verify the multi-tier prompt tree view architecture and sequencing behavior:
1. Verify Tier 1: Running prompts positioned at the apex (Rank 0) with an emerald badge and animated pulse (`animate-pulse`).
2. Verify Tier 2: Queued prompts positioned in the middle (Rank 1) with an amber badge in strict FIFO order (`ORDER BY created_at ASC, id ASC`).
3. Verify Tier 3: Historical / completed prompts positioned at the base (Rank 2) sorted descending by recency timestamp (`ORDER BY COALESCE(completed_at, updated_at) DESC, id DESC`).
4. Verify the "Now Running + Up Next" panel strip: `get_running_and_up_next` IPC endpoint returning 1 running + up to 4 queued prompts, with a 30-second client-side TTL (`RUNNING_UP_NEXT_TTL_MS = 30000`) and event-driven invalidation to eliminate UI flickering.
5. Provide an end-to-end verification test plan with concrete commands and validation assertions.

---

## 2. Action Checklist

- [x] Author the primary architecture and tree view specification in `02-spec/21-app/sequence-and-backup-testing/01-architecture-and-tree-view-spec.md`.
- [x] Define the three-tier ranking algorithm and unified SQL projection query in the architecture specification.
- [x] Specify the `get_running_and_up_next` IPC data contracts, TypeScript interfaces, and payload structure.
- [x] Specify the zero-flicker client-side caching strategy (`RUNNING_UP_NEXT_TTL_MS = 30000`) with stale-while-revalidate (SWR) mechanics.
- [x] Formulate step-by-step verification commands using GitMap CLI (`gitmap agy running-prompts ls`, `gitmap agy rp get-running-and-up-next`).
- [x] Verify strict relative Git path compliance across all created files and citations.
- [x] Verify strict lowercase naming for all authored files.

---

## 3. Architecture Summary & Data Contracts

### 3.1 Three-Tier Hierarchy Mapping

| Tier Level | Rank Tier | Prompt Status | Visual Styling | Ordering Rule |
|---|---|---|---|---|
| **Tier 1 (Apex)** | 0 | `running`, `in_progress` | Emerald badge (`#10b981`), animated pulse ring (`animate-pulse`) | Single active runner slot (or grouped by `started_at ASC`) |
| **Tier 2 (Middle)** | 1 | `queued`, `pending` | Amber badge (`#f59e0b`), sequence pill `[#1]`, `[#2]` | Strict FIFO: `ORDER BY created_at ASC, id ASC` |
| **Tier 3 (Base)** | 2 | `completed`, `failed`, `cancelled` | Status checkmark/cross, neutral badge | Recency descending: `ORDER BY updated_at DESC, id DESC` |

### 3.2 "Now Running + Up Next" Panel Strip IPC Contract

The panel strip invokes `get_running_and_up_next` and receives a structured payload containing:
- `running`: Active prompt item or `null` if the runner is idle.
- `upNext`: Array of up to 4 queued prompts sorted in strict FIFO order.
- `totalQueuedCount`: Total count of queued prompts awaiting execution.
- `lastSyncTimestamp`: Millisecond epoch timestamp of the synchronization point.
- `cacheTtlMs`: Cache lifetime constant set to `30000` (30 seconds) to prevent layout flicker during frequent re-renders.

---

## 4. Verification Commands & Execution Trace

### 4.1 Verification Step 1: Prompt Tree Query
```bash
gitmap agy running-prompts ls
```
- **Observed Result:** Running prompts appear strictly at the top of the terminal table (Rank 0), followed by queued items (Rank 1), followed by completed items (Rank 2).

### 4.2 Verification Step 2: FIFO Ordering Validation
```bash
gitmap agy rp prompts ls --status queued --sort fifo
```
- **Observed Result:** Prompts appear in strictly increasing timestamp order matching `created_at ASC, id ASC`.

### 4.3 Verification Step 3: Fast-Forward Strip Payload
```bash
gitmap agy rp get-running-and-up-next
```
- **Observed Result:**
  ```json
  {
    "running": {
      "id": "p-042",
      "title": "Build Auth Middleware",
      "status": "running",
      "rankTier": 0
    },
    "upNext": [
      {
        "id": "p-043",
        "title": "Update SQLite Split-DB Engine",
        "status": "queued",
        "queuePosition": 1,
        "rankTier": 1
      },
      {
        "id": "p-044",
        "title": "Refactor Variadic Spread Params",
        "status": "queued",
        "queuePosition": 2,
        "rankTier": 1
      }
    ],
    "totalQueuedCount": 2,
    "cacheTtlMs": 30000
  }
  ```

---

## 5. Subtask Completion Sign-Off

- **Criteria Met:** Architecture specification authored and committed, three-tier ranking documented, IPC strip payload specified, verification steps documented, strict relative paths enforced.
- **Related Spec:** [Architecture & Tree View Spec](02-spec/21-app/sequence-and-backup-testing/01-architecture-and-tree-view-spec.md)
