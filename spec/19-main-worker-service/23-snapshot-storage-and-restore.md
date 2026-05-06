# 23 — Snapshot Storage and Restore Flow

**Spec:** `19-main-worker-service`
**Version:** 1.1.0
**Created:** 2026-05-06
**Status:** Authoritative (spec-only, per plan §Mode)
**Resolves:** Locked decision **D14** (date-by-date full snapshot storage on backup; main-controlled restore by date). Closes open question **OQ-A4** — default snapshot retention adopted at **30 days rolling**.
**Depends on:** [`18-backup-nodes.md`](./18-backup-nodes.md), [`19-incremental-backup-sync.md`](./19-incremental-backup-sync.md), [`20-backup-encryption-and-keys.md`](./20-backup-encryption-and-keys.md), [`21-backup-endpoints.md`](./21-backup-endpoints.md) §6, [`22-backup-apply-logic.md`](./22-backup-apply-logic.md), [`05-auth-and-2fa.md`](./05-auth-and-2fa.md) §S2S.

---

## Keywords

`snapshot-storage` · `restore-by-date` · `retention-policy` · `backup-audience` · `point-in-time-recovery`

---

## 1. Purpose

Pin the **at-rest** half of the backup story. Phases 6–10 covered registration, change capture, encryption, wire, and apply. This file owns:

1. How a backup node materialises a date-anchored full snapshot from the apply mirror.
2. Where snapshots live on disk, how they are named, and how long they survive.
3. The Main-controlled restore flow that pushes a chosen snapshot back to the primary.
4. The `Backup` S2S audience that authenticates every Phase-9 endpoint plus the new restore-inbox endpoint on the primary.

---

## 2. Snapshot lifecycle

Three lifecycle moments, each owned by a single component.

| Moment | Owner | Trigger | Output |
|---|---|---|---|
| **Build** | Backup node | Daily cron at `MainWorker.Backup.Snapshot.BuildHourUtc` | New file under `var/snapshots/<YYYY-MM-DD>.zip` + new `BackupSnapshotCatalog` row |
| **Catalogue** | Backup node | After successful Build | INSERT into `BackupSnapshotCatalog`; surfaces via BE-4 |
| **Restore** | Main → Backup → Primary | Operator BE-3 call | Backup decompresses + signs + ships to primary's restore-inbox |

Each moment is independent. A failed Build does **not** block the next day's Build (idempotent by date). A failed Restore does not damage the snapshot — Builds are append-only; Restores are read-only.

---

## 3. Build process (on the backup node)

Strictly sequential, CODE-RED ≤15 lines per step.

```
B1  AcquireSnapshotLock(date)                       — UNIQUE row in BackupSnapshotJob
B2  AssertApplyMirrorQuiesced()                     — no in-flight Stage-4 TX (per 22-… §4)
B3  CopyAppTierToStaging(date)                      — sqlite3 .backup INTO staging file
B4  SealSnapshotZip(staging, date)                  — AES-256 zip; password per §4
B5  ComputeSha256(zip)                              — fingerprint for catalogue
B6  PersistCatalogRow(date, size, sha256, KeyEpoch) — BackupSnapshotCatalog INSERT
B7  ReleaseSnapshotLock(date)
B8  RunRetentionSweep()                             — see §6
```

B2 quiescence is **soft** — uses SQLite's `BEGIN IMMEDIATE` to wait for in-flight applies, bounded by `MainWorker.Backup.Snapshot.QuiesceTimeoutSeconds` (default **120 s**). Timeout fails the Build with `WORKER-940-01 SnapshotQuiesceTimeout`; the day's snapshot is skipped, not partial.

B3 uses SQLite's online backup API (`sqlite3_backup_init`) — never a raw file copy — to guarantee a transactionally consistent snapshot even if writes resume mid-copy.

---

## 4. Snapshot zip password

Reuse the Phase-8 derivation pattern with a distinct salt to prevent password collision with envelopes:

```
SnapshotPassword(KeyEpoch, SnapshotDateEpoch) =
    Hex( HMAC-SHA256(
            SharedSecret,                        // per 20-… §5
            BigEndian64(SnapshotDateEpoch)
         ) ).Substring(0, 32)
```

The `SharedSecret` HKDF in `20-…` §5 is rebuilt with `salt = "BackupSnapshot/v1"` (vs. `"BackupZip/v1"` for envelopes) — same primitive, separate keystream, no cross-context leak.

`SnapshotDateEpoch` = midnight-UTC epoch seconds of the snapshot date. Deterministic; receiver reproduces from `BackupSnapshotCatalog` metadata alone.

---

## 5. New table: `BackupSnapshotCatalog` (App tier on backup, mirrored to Main as a view)

```
CREATE TABLE BackupSnapshotCatalog (
    BackupSnapshotCatalogId  INTEGER PRIMARY KEY AUTOINCREMENT,
    BackupWorkerNodeId       INTEGER NOT NULL,
    PrimaryWorkerNodeId      INTEGER NOT NULL,
    SnapshotDate             TEXT    NOT NULL,        -- 'YYYY-MM-DD' UTC
    SnapshotDateEpoch        INTEGER NOT NULL,        -- midnight-UTC epoch seconds (D2)
    KeyEpoch                 INTEGER NOT NULL,        -- pinned at Build time
    SizeBytes                INTEGER NOT NULL,
    Sha256Hex                TEXT    NOT NULL,
    BuiltAtEpoch             INTEGER NOT NULL,
    StoragePath              TEXT    NOT NULL,        -- relative under var/snapshots/
    Status                   TEXT    NOT NULL,        -- Available | Pinned | Reaped | Corrupt
    ReapedAtEpoch            INTEGER NULL,
    PinReason                TEXT    NULL,            -- audit trail for Status='Pinned' (Phase 12.2)
    PinnedAtEpoch            INTEGER NULL,            -- when the pin was applied (D2)
    PinnedByActor            TEXT    NULL,            -- operator identity that pinned (S2S sub or PowerAdmin UserId)
    Description              TEXT    NULL,            -- entity-ish ref → Rule 10 (still nullable)
    UNIQUE (BackupWorkerNodeId, PrimaryWorkerNodeId, SnapshotDate)
);
```

Sibling lock table:

```
CREATE TABLE BackupSnapshotJob (
    BackupSnapshotJobId      INTEGER PRIMARY KEY AUTOINCREMENT,
    BackupWorkerNodeId       INTEGER NOT NULL,
    SnapshotDate             TEXT    NOT NULL,
    Status                   TEXT    NOT NULL,        -- Building | Built | Failed
    StartedAtEpoch           INTEGER NOT NULL,
    EndedAtEpoch             INTEGER NULL,
    FailureCode              TEXT    NULL,
    Notes                    TEXT    NULL,            -- transactional → Rule 11
    Comments                 TEXT    NULL,
    UNIQUE (BackupWorkerNodeId, SnapshotDate)
);
```

Memory compliance: PascalCase + `{TableName}Id` PK; entity-ish catalogue carries `Description NULL` (Rule 10); transactional job table carries `Notes`+`Comments NULL` (Rule 11); INTEGER `*At` (D2).

---

## 6. Retention (resolves OQ-A4)

**Default:** `MainWorker.Backup.SnapshotRetentionDays = 30` — rolling.

Sweep semantics (`B8 RunRetentionSweep`):

```
S1  Compute cutoff = today_utc - SnapshotRetentionDays
S2  For each BackupSnapshotCatalog row with SnapshotDate < cutoff AND Status='Available':
        Delete file at StoragePath
        UPDATE row SET Status='Reaped', ReapedAtEpoch=now
S3  Never reap rows with Status='Corrupt' — operator must inspect first
S4  Emit one summary log line per sweep with (reaped_count, freed_bytes, oldest_remaining_date)
```

Operator overrides:

| Override | Mechanism | Effect |
|---|---|---|
| Per-pairing retention | Seedable-Config `MainWorker.Backup.SnapshotRetentionDays` | Lifts/lowers the global default. Linter `BACKUP-SNAP-002` enforces ≥ 7 days minimum (compliance floor). |
| Pin a snapshot | BE-3 sub-route `POST /API/V1/Backup/Snapshot/Pin` (per `21-backup-endpoints.md`) — see §6.1 below for the column contract | Sets `Status='Pinned'` and stamps `PinReason`, `PinnedAtEpoch`, `PinnedByActor`. Pinned rows survive every retention sweep until explicitly unpinned. |

**No auto-shortening.** The sweep never deletes a snapshot whose date is ≥ cutoff regardless of disk pressure — disk pressure is an operational alert, not a data-loss trigger.

---

### 6.1 Pin / unpin protocol (resolves OQ-23-3)

`Status='Pinned'` is the only retention-bypass mechanism. Because pinning blocks the sweep indefinitely, every pin MUST carry an audit trail — silent pins are a CODE RED violation (no swallowed reasons).

**Required column contract on every pin transition (`Available` → `Pinned`):**

| Column | Required value | Validation |
|---|---|---|
| `Status` | `'Pinned'` | Enum check (linter `BACKUP-SNAP-005`). |
| `PinReason` | NON-empty TEXT, ≤ 500 chars | Free-text per Rule 12 (TEXT NULL on the table — but the pin operation refuses NULL/empty). |
| `PinnedAtEpoch` | `now()` epoch seconds (D2) | Set server-side; client-supplied values rejected. |
| `PinnedByActor` | One of: `"S2S:<PairingId>"`, `"User:<UserId>"` | Derived from the calling token; never client-supplied. |

**Unpin transition (`Pinned` → `Available`):** clears all four pin columns in a single transaction. The previous values are preserved in the `EndpointAuthAuditEvent` row written for the pin/unpin call (per `06-core-api-endpoints.md` §5.6) — so the table itself is the **current** state and the audit log is the **history**.

**Forbidden transitions:**

- `Pinned` → `Reaped` directly. The sweep MUST skip pinned rows; explicit unpin is required first.
- `Pinned` row with NULL `PinReason`. Linter `BACKUP-SNAP-005` fails the build.
- Pin via raw SQL `UPDATE` outside the documented endpoint. Operational runbooks MUST route through BE-3's sub-route so `EndpointAuthAuditEvent` is written.

**Why three columns instead of just `PinReason`:** "who" and "when" are the audit questions operators always ask after the fact. Capturing them on the row itself (denormalised against the audit log) makes incident response a single SELECT, not a JOIN against a multi-million-row audit table.

---

## 7. Restore flow (Main → Backup → Primary)

Initiated by BE-3 (`21-backup-endpoints.md` §6). Eight steps, three nodes, one job-id.

```
R1  Operator calls BE-3 on Backup with SnapshotDate + TargetPrimaryWorkerNodeId.
R2  Backup creates BackupRestoreJob row (Status='Accepted'), returns RestoreJobId, 202.
R3  Backup async: Look up BackupSnapshotCatalog by SnapshotDate.
       Missing  → MAIN-830-01 SnapshotNotFound (job → Failed).
       Corrupt  → MAIN-840-02 SnapshotCorrupt (Phase 11 reserved; see §9).
R4  Backup decrypts snapshot zip (per §4 password) into staging.
R5  Backup re-seals as a Phase-8 envelope variant ("RestoreEnvelope") signed under the
    CURRENT Active KeyEpoch — NOT the snapshot's original KeyEpoch — so the primary
    decrypts with the in-force epoch.
R6  Backup POSTs to the primary's NEW restore-inbox endpoint:
       POST /API/V1/Backup/RestoreInbox  (hosted on PRIMARY worker)
       Body: multipart (Metadata + RestoreEnvelope), scope `Backup.Restore.Apply`.
R7  Primary verifies signature, replaces App tier via offline import (out of normal apply
    path — no SyncOp dispatch), bumps a new "RestoredFromSnapshot" audit row.
R8  Primary ACKs to Backup; Backup updates BackupRestoreJob.Status='Completed'; Main
    polls BE-5 Health to surface completion.
```

Re-sealing in **R5** under the Active epoch (not the snapshot's pinned epoch) ensures forward secrecy — a long-retired KeyEpoch never has to be revived on the primary just to restore a 25-day-old snapshot.

`BackupRestoreJob` table:

```
CREATE TABLE BackupRestoreJob (
    BackupRestoreJobId       INTEGER PRIMARY KEY AUTOINCREMENT,
    BackupWorkerNodeId       INTEGER NOT NULL,
    PrimaryWorkerNodeId      INTEGER NOT NULL,
    SnapshotDate             TEXT    NOT NULL,
    Reason                   TEXT    NOT NULL,        -- per BE-3 §6.1
    ForceOverwrite           INTEGER NOT NULL,        -- bool
    Status                   TEXT    NOT NULL,        -- Accepted | Building | Shipping | Completed | Failed
    AcceptedAtEpoch          INTEGER NOT NULL,
    CompletedAtEpoch         INTEGER NULL,
    FailureCode              TEXT    NULL,
    PrimaryAckPayloadJson    TEXT    NULL,
    Notes                    TEXT    NULL,
    Comments                 TEXT    NULL,
    UNIQUE (BackupWorkerNodeId, PrimaryWorkerNodeId, SnapshotDate, AcceptedAtEpoch)
);
```

`UNIQUE` includes `AcceptedAtEpoch` so the same date can be restored twice intentionally — the `MAIN-830-02 RestoreAlreadyInProgress` guard from Phase 9 handles concurrent attempts.

---

## 8. New endpoint: `POST /API/V1/Backup/RestoreInbox` (hosted on **primary** Worker)

Symmetric counterpart to BE-1 but flowing **backward**. Catalogued as **BE-6** for Phase-12 merge into `06-core-api-endpoints.md`.

| # | Method | Path | Hosted on | Auth | Purpose |
|---|---|---|---|---|---|
| BE-6 | POST | `/API/V1/Backup/RestoreInbox` | **Primary** Worker | S2S OAuth + scope `Backup.Restore.Apply` | Receive a re-sealed full snapshot from a paired backup; replace App tier offline. |

Request shape mirrors BE-1 §4.1 (multipart with `Metadata` + `Envelope`) — see `21-backup-endpoints.md` §4.1 for the schema. Differences:

- `Metadata.EnvelopeKind = "RestoreEnvelope"` (vs. `"IncrementalDiff"`).
- `Metadata.SnapshotDate` present.
- Apply path is **R7 offline import** — does NOT pass through `22-backup-apply-logic.md` Stage 4.

Response (`200 OK`):

```jsonc
{
  "RestoreJobId":         "01J...ULID",
  "AcceptedAtEpoch":      1746528900,
  "AppTierBytesReplaced": 4194304,
  "PreviousAppTierSha256":"...",
  "NewAppTierSha256":     "..."
}
```

Failure codes:

| Trigger | Code |
|---|---|
| Caller scope missing `Backup.Restore.Apply` | `MAIN-100-01 AuthHandshakeFail` |
| `Metadata.SnapshotDate` not parseable | `WORKER-300-03 RequestBodyInvalid` |
| Signature verify failure | `WORKER-920-05 EnvelopeSignatureInvalid` |
| Offline import write fails | `WORKER-940-02 RestoreImportFailed` |

---

## 9. `Backup` S2S audience

Phase 9 §9 reserved the `Backup` audience for this file. Final wiring:

| Audience | Issuer | Subject | Allowed scopes | Consumer |
|---|---|---|---|---|
| `Backup` | Main | A specific Backup or Primary identity | `Backup.Diff.Write`, `Backup.Rotate.Write`, `Backup.Restore.Write`, `Backup.Restore.Apply`, `Backup.Read` | BE-1..BE-6 |

Token claims (per `12-jwt-delivery-contract.md`):

```jsonc
{
  "iss": "main",
  "aud": "Backup",
  "sub": "WorkerNode/27",                  // backup or primary
  "scope": "Backup.Diff.Write Backup.Read",
  "exp": <now + WorkerJwtTtlSeconds>,
  "PairingId": "12-27"                     // PrimaryWorkerNodeId-BackupWorkerNodeId
}
```

`PairingId` is **mandatory** on every `Backup`-audience token — the receiver short-circuits with `MAIN-800-04 TrafficOnBackupRejected` on mismatch with the local `KnownBackupNode` row.

Add a stub note to `05-auth-and-2fa.md` §S2S in Phase 12: "Backup audience defined in `23-snapshot-storage-and-restore.md` §9."

---

## 10. Tunables introduced (mirrored verbatim into `15-tunable-constants.md` §2.15)

| Key | Default | Unit | Used by | Notes |
|---|---:|---|---|---|
| `MainWorker.Backup.SnapshotRetentionDays` | **30** | days | §6 sweep | **Resolves OQ-A4.** Linter `BACKUP-SNAP-002` enforces ≥ 7 (compliance floor). |
| `MainWorker.Backup.Snapshot.BuildHourUtc` | **3** | hour-of-day (0-23) | §3 cron | Off-peak default; overridable per node. |
| `MainWorker.Backup.Snapshot.QuiesceTimeoutSeconds` | **120** | seconds | §3 B2 | Below `MaxRetriesPerEnvelope × TransactionTimeoutSeconds` (5×30=150) so quiesce never out-races a single envelope. |
| `MainWorker.Backup.Snapshot.MaxBuildSeconds` | **1800** | seconds (30 m) | §3 B3 | Hard ceiling on `sqlite3_backup_step` total elapsed; exceeded → `WORKER-940-03 SnapshotBuildTimeout`. |
| `MainWorker.Backup.Restore.PrimaryAckTimeoutSeconds` | **600** | seconds (10 m) | §7 R8 | Backup waits this long for the primary's BE-6 200 before flipping `BackupRestoreJob.Status='Failed'`. |

---

## 11. Errors introduced (mirrored verbatim into `13-error-codes.md`)

Worker tier (snapshot/restore failures, range `WORKER-940-*` — slots `21204-21207` in the freshly-allocated overflow window from Phase 10):

| Prefixed | Flat | HTTP | Name | Meaning |
|---|---:|---:|---|---|
| `WORKER-940-01` | `21204` | 504 | `SnapshotQuiesceTimeout` | B2 quiesce wait exceeded `QuiesceTimeoutSeconds`. |
| `WORKER-940-02` | `21205` | 500 | `RestoreImportFailed` | BE-6 R7 offline import write failed. |
| `WORKER-940-03` | `21206` | 504 | `SnapshotBuildTimeout` | B3 `sqlite3_backup_step` exceeded `MaxBuildSeconds`. |
| `WORKER-940-04` | `21207` | 500 | `SnapshotSealFailed` | B4 AES-256-zip seal raised IO/crypto error. |

Main tier (snapshot integrity, slot `21192` — first free in Phase-11-reserved window per §9.2 of Phase 10):

| Prefixed | Flat | HTTP | Name | Meaning |
|---|---:|---:|---|---|
| `MAIN-840-02` | `21192` | n/a | `SnapshotCorrupt` | R3 SHA-256 mismatch against `BackupSnapshotCatalog.Sha256Hex`; surfaced via BE-5. |

§4 (range table) of `13-error-codes.md` updates: `WORKER-21204-21207` consumed; `MAIN-21192` consumed; `MAIN-21193-21199` reserved for future.

---

## 12. Linter hooks queued for Phase 12

| ID | Rule |
|---|---|
| `BACKUP-SNAP-001` | Every `BackupSnapshotCatalog.Status='Available'` row MUST have `StoragePath` resolving to a real file under `var/snapshots/` matching `Sha256Hex`. |
| `BACKUP-SNAP-002` | `MainWorker.Backup.SnapshotRetentionDays` MUST be ≥ 7. Compliance floor. |
| `BACKUP-SNAP-003` | `BackupSnapshotJob` rows older than 90 days with `Status='Failed'` MUST be either reaped or have a paired `BackupApplyDeadLetter` row — orphaned failures fail CI. |
| `BACKUP-SNAP-004` | Build pipeline functions MUST be ≤15 lines, zero nested `if`, ≤2 operands per boolean (CODE RED). |
| `BACKUP-SNAP-005` | Every `BackupSnapshotCatalog.Status='Pinned'` row MUST have NON-empty `PinReason`, NON-NULL `PinnedAtEpoch`, and NON-NULL `PinnedByActor` matching `^(S2S:|User:).+`. Silent pins fail CI. |

---

## 13. Cross-references

- Decision register: `.lovable/plan.md` D14, OQ-A4.
- `18-backup-nodes.md` §6 `KnownBackupNode` — pairing assertion source for BE-6.
- `19-incremental-backup-sync.md` — incremental path co-existing with snapshot path; restores reset the watermark per R7.
- `20-backup-encryption-and-keys.md` §5 — HKDF salt convention reused with `"BackupSnapshot/v1"`.
- `21-backup-endpoints.md` §6 — BE-3 wire (caller of restore flow).
- `22-backup-apply-logic.md` §4 — incremental Stage-4 dispatch; restore path **bypasses** this.
- `13-error-codes.md` §2.10 / §3.11 / §4 — code allocations + range table refresh.
- `15-tunable-constants.md` §2.15 — mirrored tunables.
- `05-auth-and-2fa.md` §S2S — `Backup` audience added per §9.
- `12-jwt-delivery-contract.md` — mandatory `PairingId` claim per §9.

---

## 14. Open Questions (logged, non-blocking)

- **OQ-23-1** Should snapshots be deduplicated across days (rolling diff zip pyramid) to save disk for low-write-rate primaries? Inferred: defer to v2.0 — flat date-named files are dumb-AI friendliest and disk is cheap at 30-day default.
- **OQ-23-2** Should the restore flow support partial-table restore (e.g., one tenant)? Inferred: no — out of scope; full App-tier replacement only. Tenant-level recovery is owned by the application data model, not the backup tier.
- **OQ-23-3** ✅ **Resolved Phase 12.2.** `Pinned` status now requires `PinReason TEXT NULL` (NOT NULL on insert), `PinnedAtEpoch INTEGER NULL`, and `PinnedByActor TEXT NULL`. Pin/unpin contract codified in §6.1; enforced by linter `BACKUP-SNAP-005`.

---

*Snapshot storage + restore flow v1.1.0 — 2026-05-06 (Phase 12.2 — OQ-23-3 resolved: `PinReason` + `PinnedAtEpoch` + `PinnedByActor` columns + §6.1 pin/unpin protocol + `BACKUP-SNAP-005`).*
