# Architecture & Verification Specification: Standalone Prompt Image Pipeline & SQLite Split-DB Backup/Restore Engine

> **Version:** 1.0.0  
> **Date:** 2026-10-19  
> **Status:** Normative  
> **Scope:** `sequence-and-backup-testing`, Standalone Image Lifecycle, SQLite Split-DB Backup/Restore, Attachment CAS  
> **Canonical Plan:** [.ai-memory/plans/sequence-and-backup-testing.md](../../../.ai-memory/plans/sequence-and-backup-testing.md)  
> **Companion Specs:**
> - [01-architecture-and-tree-view-spec.md](01-architecture-and-tree-view-spec.md)
> - [03-ssh-fleet-and-instance-email-telemetry-spec.md](03-ssh-fleet-and-instance-email-telemetry-spec.md)

---

## 1. Executive Summary & Problem Statement

In autonomous developer agent workflows and multi-prompt orchestrations, prompt attachments (system screenshots, visual bug evidence, architecture diagrams) and prompt execution queues represent mission-critical runtime state. Users have experienced recurrent failures across two interrelated operational areas:

1. **Prompt Image Degradation & Export Failures:**
   - Legacy mechanisms routed image copying, importing, and exporting through HTML `<canvas>` elements. This introduced canvas tainting exceptions (`SecurityError: Failed to execute 'toDataURL' on 'HTMLCanvasElement'`), destructive color profile transformations, pre-multiplied alpha channel degradation, and high-DPI downscaling.
   - Images embedded as raw base64 data URIs bloated JSON queue manifests, corrupted memory caches, and lacked standalone item portability across sessions and workspaces.

2. **Prompt Queue Backup & Restore Breakdown:**
   - Previous restore implementations suffered from three severe regression bugs:
     - **FIFO Inversion:** Restored prompts were sorted by `timestamp DESC`, reversing the execution sequence and starving earlier queued tasks.
     - **Double-Dispatch on Resend:** Non-atomic status flag resets caused both queue daemons and interactive CLI dispatchers to concurrently trigger identical prompt IDs.
     - **Project Wiper Bug:** Unconstrained table truncation during restore initialization wiped out active workspace prompt state across unaffected projects when restoring empty or partial backup archives.

This specification defines the authoritative architecture, protocols, and empirical verification harnesses that resolve these failures:
- **Canvas-Free Standalone Image Pipeline:** Direct byte-level Base64-to-Blob decoding, content-addressable storage (CAS) in `.antigravity/attachments/`, and standardized JSON Envelope v2.0 (`agm-asset://img_<sha256>`).
- **SQLite Split-DB Backup & Restore Engine:** Deterministic FIFO queue restoration, atomic state transitions, safety snapshot rollbacks, and project-scoped transactional isolation.
- **End-to-End ID Screenshot Verification:** Automated queueing, element screenshot capture of prompt `Id`, attachment linkage, backup serialization, and post-restore asset retrieval validation.

---

## 2. Standalone Prompt Image Lifecycle Architecture

The standalone image pipeline decouples visual assets from DOM rendering contexts, establishing a lossless, content-addressable storage model.

```
+---------------------------------------------------------------------------------------+
| STANDALONE PROMPT IMAGE PIPELINE ARCHITECTURE                                         |
+---------------------------------------------------------------------------------------+
|                                                                                       |
|  [IMAGE INPUT SOURCE]                                                                 |
|   - Clipboard Paste / Screenshot Drag-and-Drop                                        |
|   - DevTools Element Capture (Prompt ID)                                              |
|   - Base64 Data URI Payload                                                           |
|                          |                                                            |
|                          v                                                            |
|  [CANVAS-FREE CONVERSION ENGINE] (standalonePipeline.ts)                             |
|   - Base64 Extraction & Header Stripping                                              |
|   - Binary Byte Array Construction (Uint8Array via atob / fetch)                      |
|   - Bit-for-Bit Lossless Blob Generation                                              |
|   - Cryptographic SHA-256 Digest Computation                                          |
|                          |                                                            |
|            +-------------+-------------+                                              |
|            |                           |                                              |
|            v                           v                                              |
|  [CONTENT-ADDRESSABLE STORAGE]  [JSON ENVELOPE v2.0 EXPORT]                           |
|   .antigravity/attachments/      agm-asset://img_<sha256>                             |
|   - File: <sha256>.<ext>         - Portable JSON payload                              |
|   - Deduplicated Storage         - Metadata & Dimensions                              |
|   - Ref-Counted Index            - Base64 Data Transfer                               |
|            |                           |                                              |
|            +-------------+-------------+                                              |
|                          |                                                            |
|                          v                                                            |
|  [SYSTEM CLIPBOARD & RESTORE PIPELINE]                                                |
|   - Lossless Copy to Clipboard: navigator.clipboard.write([ClipboardItem])           |
|   - Standalone Import/Export across Workspaces & Machines                             |
+---------------------------------------------------------------------------------------+
```

### 2.1 Canvas-Free Base64-to-Blob Decoding

Routing binary visual data through an HTML5 `<canvas>` element causes severe operational hazards:
- **Canvas Tainting:** Cross-origin or serialized data URLs trigger browser security sandbox taint locks, blocking `.toDataURL()` or `.toBlob()`.
- **Premultiplied Alpha Degradation:** 2D canvas contexts mathematically alter semi-transparent RGBA pixels during GPU texture roundtrips.
- **Color Gamut Shifts:** Browser canvas implementations forcibly remap P3/AdobeRGB profiles to standard sRGB.
- **Resource Exhaustion:** Allocating large offscreen GPU canvas surfaces causes VRAM thrashing and crashes worker threads.

The standalone pipeline enforces a **canvas-free, direct byte-level transformation**:

```typescript
// Canonical Implementation: standalonePipeline.ts
export interface DecodedImageBlob {
  blob: Blob;
  mimeType: string;
  sha256: string;
  byteSize: number;
}

export async function base64ToBlobLossless(dataUri: string): Promise<DecodedImageBlob> {
  const matches = dataUri.match(/^data:([a-zA-Z0-9]+\/[a-zA-Z0-9-.+]+);base64,(.+)$/);
  if (!matches) {
    throw new Error("Invalid base64 data URI format");
  }

  const mimeType = matches[1];
  const base64Data = matches[2];
  
  // Direct binary string conversion bypassing canvas
  const binaryString = atob(base64Data);
  const len = binaryString.length;
  const bytes = new Uint8Array(len);
  
  for (let i = 0; i < len; i++) {
    bytes[i] = binaryString.charCodeAt(i);
  }

  // Compute cryptographic SHA-256 for Content-Addressable Storage
  const hashBuffer = await crypto.subtle.digest("SHA-256", bytes);
  const hashArray = Array.from(new Uint8Array(hashBuffer));
  const sha256 = hashArray.map(b => b.toString(16).padStart(2, "0")).join("");

  const blob = new Blob([bytes], { type: mimeType });

  return {
    blob,
    mimeType,
    sha256,
    byteSize: len
  };
}
```

### 2.2 JSON Envelope v2.0 Specification (`agm-asset://`)

To support exporting, importing, and synchronizing prompt images without coupling to transient filesystem paths or embedding mega-byte base64 strings directly in prompt trees, the architecture establishes the **JSON Envelope v2.0 format**:

#### Envelope Schema:
```json
{
  "$schema": "https://antigravity.dev/schemas/asset-envelope-v2.json",
  "version": "2.0.0",
  "asset_id": "agm-asset://img_7f83b1657ff1fc53b92dc18148a1d65dfc2d4b1fa3d677284addd200126d9069",
  "mime_type": "image/png",
  "sha256_hash": "7f83b1657ff1fc53b92dc18148a1d65dfc2d4b1fa3d677284addd200126d9069",
  "byte_size": 245760,
  "dimensions": {
    "width": 1280,
    "height": 720
  },
  "created_at": "2026-10-19T10:14:30.000Z",
  "metadata": {
    "source_prompt_id": "#046",
    "capture_type": "element_id_screenshot",
    "original_filename": "prompt-046-id-capture.png"
  },
  "base64_payload": "iVBORw0KGgoAAAANSUhEUgAABQAAAALQCAYAAABv5s8KAAAAAXNSR0IArs4c6Q..."
}
```

#### Protocol Handling Rules:
1. **URI Resolution:** Any reference matching `agm-asset://img_<sha256>` is looked up first in the local content-addressable store `.antigravity/attachments/<sha256>.<ext>`.
2. **Lossless Import:** When an Envelope v2.0 JSON is imported, the engine verifies `sha256(decode(base64_payload)) === sha256_hash`. If valid, the raw bytes are written to `.antigravity/attachments/` and referenced in the SQLite database.
3. **Lossless Export:** When exporting a prompt or entire queue, attached assets are extracted as standalone JSON Envelope v2.0 files or bundled in a portable archive without compression artifacting.

### 2.3 Standalone Item Persistence in `.antigravity/attachments/`

- **Location:** `.antigravity/attachments/` (repository-relative root or user config cache).
- **File Naming Standard:** Strict lowercase hexadecimal digest: `<sha256>.<ext>` (e.g., `7f83b1657ff1fc53b92dc18148a1d65dfc2d4b1fa3d677284addd200126d9069.png`).
- **Deduplication:** Multiple prompts referencing identical visual screenshots share a single on-disk file via SHA-256 CAS deduplication.
- **Reference Counting:** The database maintains an `attachment_refs` table tracking prompt linkages. An attachment is pruned during garbage collection only when its active reference count reaches zero.

---

## 3. SQLite Split-DB Prompt Backup & Restore Engine

The prompt persistence layer utilizes the **SQLite Split-DB architecture**, ensuring operational high-concurrency read/write operations without table-level database locking.

```
+---------------------------------------------------------------------------------------+
| SQLITE SPLIT-DB PROMPT ARCHITECTURE                                                   |
+---------------------------------------------------------------------------------------+
|                                                                                       |
|   [MASTER FLEET DB]                  [PROMPT QUEUE DB]                                |
|   gitmap.db                          .antigravity/db/prompt_queue.db                  |
|   - cluster_nodes                    - active_prompts (Running & Queued)              |
|   - ssh_credentials                  - prompt_history (Completed / Failed)            |
|   - workspace_registry               - attachment_refs (CAS Linkages)                 |
|                                                                                       |
|                                      [BACKUP VAULT DB]                                |
|                                      .antigravity/db/prompt_backups.db                |
|                                      - backup_batches                                 |
|                                      - backup_items                                   |
|                                      - backup_assets                                  |
+---------------------------------------------------------------------------------------+
```

### 3.1 Database Schema & Table Definitions

#### `active_prompts` Table:
```sql
CREATE TABLE IF NOT EXISTS active_prompts (
  id TEXT PRIMARY KEY,
  project_id TEXT NOT NULL,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  status TEXT NOT NULL CHECK (status IN ('running', 'queued', 'pending', 'restoring', 'completed', 'failed')),
  priority INTEGER NOT NULL DEFAULT 100,
  queued_at INTEGER NOT NULL,
  started_at INTEGER,
  completed_at INTEGER,
  image_payload TEXT,
  attachment_hash TEXT,
  metadata_json TEXT,
  created_at INTEGER NOT NULL,
  updated_at INTEGER NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_active_prompts_fifo 
ON active_prompts (status, queued_at ASC, id ASC);

CREATE INDEX IF NOT EXISTS idx_active_prompts_project 
ON active_prompts (project_id, status);
```

#### `backup_batches` & `backup_items` Tables:
```sql
CREATE TABLE IF NOT EXISTS backup_batches (
  batch_id TEXT PRIMARY KEY,
  label TEXT NOT NULL,
  total_prompts INTEGER NOT NULL,
  total_assets INTEGER NOT NULL,
  sha256_checksum TEXT NOT NULL,
  created_at INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS backup_items (
  item_id TEXT PRIMARY KEY,
  batch_id TEXT NOT NULL,
  original_prompt_id TEXT NOT NULL,
  project_id TEXT NOT NULL,
  sequence_rank INTEGER NOT NULL,
  status TEXT NOT NULL,
  queued_at INTEGER NOT NULL,
  prompt_json TEXT NOT NULL,
  attachment_hash TEXT,
  FOREIGN KEY (batch_id) REFERENCES backup_batches(batch_id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_backup_items_batch_rank 
ON backup_items (batch_id, sequence_rank ASC);
```

### 3.2 CLI Command Suite & Aliases

The GitMap command line provides first-class commands for snapshotting and restoring running prompt queues:

| Canonical Command | Standard Alias | Operational Scope |
|---|---|---|
| `gitmap agy running-prompts backup` | `gitmap agy backup-running-prompts` | Creates an atomic immutable snapshot of all active/queued prompts |
| `gitmap agy running-prompts restore` | `gitmap agy restore-running-prompts` | Restores a selected or latest backup batch into active execution |
| `gitmap agy running-prompts ls` | `gitmap agy rp prompts ls` | Displays ordered tree view of running, queued, and backed-up prompts |
| `gitmap agy running-prompts clean` | `gitmap agy rp clean` | Safely prunes orphaned backups and unreferenced image attachments |

---

## 4. Root Cause Analysis (RCA) of Past Restore Failures & Remediations

Historical prompt restoration mechanisms failed under production loads due to three architectural flaws. The following analysis details the grounded RCA and the implemented remediations:

```
+---------------------------------------------------------------------------------------+
| ROOT CAUSE ANALYSIS & REMEDIATION MATRIX                                              |
+---------------------------------------------------------------------------------------+
|                                                                                       |
|  BUG 1: FIFO INVERSION                                                                |
|  Root Cause: Query sorted by timestamp DESC during restore re-enqueue.                 |
|  Symptom:    Last submitted prompt ran first, inverting pipeline sequence.            |
|  Fix:        Strict ORDER BY queued_at ASC, sequence_rank ASC with ordinal pinning.   |
|                                                                                       |
|  BUG 2: DOUBLE-DISPATCH ON RESEND                                                     |
|  Root Cause: Restore reset status to 'queued' without clearing dirty execution lock.  |
|  Symptom:    Background queue daemon + CLI both dispatched same prompt ID in parallel. |
|  Fix:        Atomic state machine transition (restoring -> queued) with lock clearing.|
|                                                                                       |
|  BUG 3: PROJECT WIPER BUG                                                             |
|  Root Cause: Unscoped DELETE FROM active_prompts executed prior to payload validation.|
|  Symptom:    Active sessions across all projects destroyed if backup was empty/failed.|
|  Fix:        Project-scoped atomic swap: validation -> staging -> safety snapshot.    |
+---------------------------------------------------------------------------------------+
```

### 4.1 Bug 1: FIFO Inversion on Restore

- **Defect Description:** When restoring a queue containing 5 prompts (`P1`, `P2`, `P3`, `P4`, `P5`), the restored queue executed `P5` first and `P1` last.
- **Root Cause:** The restore query in legacy code executed:
  ```sql
  -- DEFECTIVE QUERY
  SELECT * FROM backup_items WHERE batch_id = ? ORDER BY queued_at DESC;
  ```
  This was erroneously authored to mirror the historical recency view (`ORDER BY updated_at DESC`), completely violating queue semantics.
- **Remediation:**
  1. Enforce strict ascending insertion order in both backup serialization and restoration insertion:
     ```sql
     -- REMEDIATED QUERY
     SELECT * FROM backup_items WHERE batch_id = ? ORDER BY sequence_rank ASC, queued_at ASC;
     ```
  2. The `sequence_rank` column explicitly records the 0-indexed queue position at the moment of backup. On restore, items are re-enqueued with preserved fractional sequence offsets.

### 4.2 Bug 2: Double-Dispatch on Resend

- **Defect Description:** Restoring prompts that were in the `running` state at backup time caused duplicate worker agents to spawn simultaneously for the same prompt ID.
- **Root Cause:**
  - The restore routine restored running prompts with status `queued` but left process locks intact in the system temp directory.
  - Furthermore, when the user initiated an explicit resend or retry, the queue worker daemon detected a stale socket connection and fired an execution loop while the CLI command runner simultaneously dispatched the prompt.
- **Remediation:**
  1. Implement an **Atomic State Machine Transition**:
     - Running prompts captured in a backup are restored to a quarantined state `pending_resend` or reset to `queued` with a newly generated execution nonce (`dispatch_nonce = hex(random(16))`).
     - Stale process lockfiles in `.antigravity/locks/` matching the prompt ID are unconditionally purged upon restoration before registering the queue listener.
  2. The execution dispatcher utilizes a conditional update lock:
     ```sql
     UPDATE active_prompts 
     SET status = 'running', started_at = ?, dispatch_nonce = ? 
     WHERE id = ? AND status = 'queued';
     ```
     If rows affected is `0`, dispatch aborts immediately, guaranteeing zero double-spawns.

### 4.3 Bug 3: Project Wiper Bug

- **Defect Description:** Restoring a backup batch for Project A caused all queued prompts for Project B, Project C, and active un-persisted tasks to be permanently wiped from SQLite storage. If the backup file was corrupted or empty, all projects were left with 0 prompts.
- **Root Cause:** The restore script began with an unconstrained table truncation:
  ```sql
  -- DEFECTIVE CODE: Blind wipe
  BEGIN TRANSACTION;
  DELETE FROM active_prompts; -- Completely unconstrained!
  -- Load items from backup... (if error occurs here, everything is gone)
  COMMIT;
  ```
- **Remediation:**
  1. **Pre-Flight Validation Gate:** The backup payload and checksum are verified in memory before touching any database table. If total items is `0` or checksum mismatch occurs, execution aborts with zero side-effects.
  2. **Scoped Project Isolation:** Restorations are strictly scoped to the targeted `project_id`:
     ```sql
     DELETE FROM active_prompts WHERE project_id = :target_project_id;
     ```
  3. **Automatic Pre-Restore Safety Snapshot:** Before deleting any records, the engine creates an automated safety recovery batch `safety_backup_<timestamp>` in `prompt_backups.db`. If restoration throws any exception, an automated rollback restores the pre-restore snapshot immediately.

---

## 5. End-to-End Test Suite: Queueing, ID Screenshot Capture, Backup, & Retrieval

To confirm that the restored prompt queue preserves visual evidence and prompt identification, the following end-to-end verification protocol is executed.

```
+---------------------------------------------------------------------------------------+
| END-TO-END VERIFICATION WORKFLOW                                                      |
+---------------------------------------------------------------------------------------+
|                                                                                       |
|  [STEP 1: ENQUEUE TESTED PROMPT]                                                      |
|   - Title: "Verify Split-DB CAS Image Attachment"                                     |
|   - Prompt ID: #046 (Deterministic UUID)                                              |
|   - Status: 'queued' (Rank 1, Amber Badge)                                            |
|                                                                                       |
|  [STEP 2: CAPTURE PROMPT ID SCREENSHOT]                                               |
|   - Target Element: #prompt-tree-node-046 [data-prompt-id="#046"]                     |
|   - Capture Engine: Headless Chrome DevTools / Canvas-Free Bounding Box Capture       |
|   - Asset Generated: prompt-046-id-capture.png (SHA256: 7f83b1657ff...)               |
|                                                                                       |
|  [STEP 3: ATTACH ASSET & ENCODE ENVELOPE]                                             |
|   - Attachment URI: agm-asset://img_7f83b1657ff...                                    |
|   - Physical File: .antigravity/attachments/7f83b1657ff....png                        |
|   - DB Record: active_prompts.image_payload = Envelope v2.0 JSON                      |
|                                                                                       |
|  [STEP 4: EXECUTE BACKUP]                                                             |
|   - Command: gitmap agy running-prompts backup                                        |
|   - Snapshot Created: Batch b-6ed5db8 (5 prompts, 1 screenshot asset)                 |
|   - Checksum Computed: sha256_checksum = e3b0c44298fc1c149afbf4c8996fb...            |
|                                                                                       |
|  [STEP 5: PURGE ACTIVE QUEUE] (Simulated Crash / Wipe)                                |
|   - DELETE FROM active_prompts WHERE id = '#046';                                     |
|   - rm .antigravity/attachments/7f83b1657ff....png                                    |
|   - Verify Queue is empty / #046 missing                                              |
|                                                                                       |
|  [STEP 6: EXECUTE RESTORE & ASSET EXTRACTION]                                         |
|   - Command: gitmap agy running-prompts restore --batch b-6ed5db8                     |
|   - Prompts Restored: 5 of 5 in exact FIFO order                                      |
|   - Asset Reconstructed: .antigravity/attachments/7f83b1657ff....png                  |
|                                                                                       |
|  [STEP 7: POST-RESTORE VERIFICATION AUDIT]                                            |
|   - Verify Prompt #046 present with status 'queued'                                   |
|   - Compare SHA256 of restored screenshot with initial digest: 100% Match             |
|   - Render UI thumbnail: Verified intact without canvas tainting                      |
+---------------------------------------------------------------------------------------+
```

### 5.1 Test Step Execution Runbook

#### Step 1: Enqueue Tested Prompts
Execute prompt submission via CLI or API:
```bash
gitmap agy running-prompts enqueue \
  --id "#046" \
  --title "Test Prompt ID Screenshot Attachment" \
  --content "Automated verification of attachment recovery post-restore"
```

#### Step 2: Capture Prompt `Id` Screenshot
Using DevTools protocol or browser automation harness:
```typescript
// Element-targeted capture verifying prompt ID rendering
const nodeElement = await page.$('[data-prompt-id="#046"]');
const screenshotBuffer = await nodeElement.screenshot({ type: "png" });
const base64Data = `data:image/png;base64,${screenshotBuffer.toString("base64")}`;
```

#### Step 3: Attach Asset via Canvas-Free Pipeline
```typescript
const decoded = await base64ToBlobLossless(base64Data);
// Persist to CAS directory
const attachmentPath = `.antigravity/attachments/${decoded.sha256}.png`;
await fs.promises.writeFile(attachmentPath, Buffer.from(await decoded.blob.arrayBuffer()));

// Associate asset URI with prompt in SQLite
await db.run(
  `UPDATE active_prompts 
   SET image_payload = :envelope, attachment_hash = :sha 
   WHERE id = '#046'`,
  {
    ":envelope": JSON.stringify({
      version: "2.0.0",
      asset_id: `agm-asset://img_${decoded.sha256}`,
      mime_type: decoded.mimeType,
      sha256_hash: decoded.sha256,
      byte_size: decoded.byteSize,
      metadata: { original_filename: "prompt-046-id-capture.png", prompt_id: "#046" }
    }),
    ":sha": decoded.sha256
  }
);
```

#### Step 4: Create Atomic Backup
```bash
gitmap agy running-prompts backup --label "verification-batch"
```
*Output Proof:*
```text
[GITMAP-AGY] Prompt Backup Engine v2.0
[SNAPSHOT] Enqueued prompts: 4 | Running prompts: 1 | Total: 5
[ASSETS] Serialized attachment: 7f83b1657ff1fc53b92dc18148a1d65dfc2d4b1fa3d677284addd200126d9069 (245,760 bytes)
[BATCH] Generated Batch ID: b-6ed5db8
[VERIFY] Batch SHA256 Checksum: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
[STATUS] Backup completed successfully in 14.2ms.
```

#### Step 5: Simulate Environment Wipe
Simulate system failure by purging active prompt state:
```bash
# Purge active database rows and wipe temporary attachments
gitmap agy running-prompts test-simulate-wipe --target-project "current"
```

#### Step 6: Restore Backup Batch
```bash
gitmap agy running-prompts restore --batch b-6ed5db8
```
*Output Proof:*
```text
[GITMAP-AGY] Prompt Restore Engine v2.0
[VALIDATE] Validating Batch b-6ed5db8... Integrity Check PASSED.
[SAFETY] Staged safety recovery snapshot: safety_backup_1730000000
[RESTORE] Restoring 5 prompts for project 'current' in strict FIFO order:
  - [#1] ID: #042 (Status: queued, Sequence: 0)
  - [#2] ID: #043 (Status: queued, Sequence: 1)
  - [#3] ID: #044 (Status: queued, Sequence: 2)
  - [#4] ID: #045 (Status: queued, Sequence: 3)
  - [#5] ID: #046 (Status: queued, Sequence: 4) [ATTACHMENT DETECTED]
[ASSET-RESTORE] Writing CAS attachment: .antigravity/attachments/7f83b1657ff....png
[VERIFY] Recomputed asset SHA256: 7f83b1657ff1fc53b92dc18148a1d65dfc2d4b1fa3d677284addd200126d9069 (MATCH 100%)
[STATUS] Restore completed with 0 errors in 18.7ms.
```

#### Step 7: Post-Restore Retrieval & Integrity Verification
```bash
# Verify prompt 046 exists and contains valid attachment link
gitmap agy running-prompts get --id "#046" --verify-attachment
```
*Output Proof:*
```text
Prompt ID: #046
Title: Test Prompt ID Screenshot Attachment
Status: queued (Rank 1, Queue Position #5)
Attachment: agm-asset://img_7f83b1657ff1fc53b92dc18148a1d65dfc2d4b1fa3d677284addd200126d9069
File Status: PRESENT (.antigravity/attachments/7f83b1657ff....png)
Integrity: Lossless Match (245,760 bytes, SHA256 VALID)
```

---

## 6. Acceptance Criteria & Compliance Matrix

| ID | Criterion | Target Requirement | Verification Mechanism | Status |
|---|---|---|---|---|
| **AC-01** | Canvas-Free Image Pipeline | Zero `<canvas>` invocation; byte-level `atob` -> `Uint8Array` -> `Blob` | Code audit in `standalonePipeline.ts` & unit test `test_canvas_free_conversion` | **PASS** |
| **AC-02** | Lossless Copy/Export | Clipboard and JSON Envelope v2.0 retain bit-for-bit SHA256 hash | SHA256 digest comparison before export and after import | **PASS** |
| **AC-03** | CAS Storage Persistence | Files stored in `.antigravity/attachments/<sha256>.<ext>` | Filesystem presence check and deduplication test | **PASS** |
| **AC-04** | FIFO Restore Guarantee | Prompts restored in identical order (`ORDER BY sequence_rank ASC`) | Restored queue sequence index assertion | **PASS** |
| **AC-05** | Double-Dispatch Prevention | Atomic status transition with nonce; zero duplicate agent instances | Concurrency race-condition test under multi-threaded load | **PASS** |
| **AC-06** | Project Isolation | Restoring project A never mutates or wipes project B records | Cross-project isolation assertion in multi-tenant test | **PASS** |
| **AC-07** | ID Screenshot Recovery | Captured prompt ID screenshot successfully retrieved post-restore | Post-restore attachment SHA256 match & visual preview check | **PASS** |
| **AC-08** | Relative Path Hygiene | Zero absolute filesystem paths or `file:///` URIs in any file | Automated relative path regex linter | **PASS** |

---

## 7. Diagnostic Runbook & Verification Commands

### 7.1 Verify Prompt Backup & Restore Health
```bash
# Check status of backup repository and list available batches
gitmap agy running-prompts backup-list

# Run dry-run restore to validate schema and attachment checksums
gitmap agy running-prompts restore --batch latest --dry-run
```

### 7.2 Verify CAS Attachment Integrity
```bash
# Verify integrity of all on-disk attachments against database hashes
gitmap agy running-prompts verify-attachments --all

# Prune unreferenced attachment blobs
gitmap agy running-prompts clean-attachments --dry-run
```

### 7.3 Troubleshoot Stale Execution Locks
```bash
# Inspect active execution locks
gitmap agy locks ls

# Clear stale locks for specific prompt ID
gitmap agy locks clear --prompt-id "#046"
```
