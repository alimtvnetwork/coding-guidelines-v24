# 98 — Changelog

**Spec:** `19-main-worker-service`

---

## v2.10.0 — 2026-05-06 (Phase 12 — Final consolidation)

**Scope:** Closes the Backup System spec arc (Phases 7–11). No new feature surface; this phase is wiring, diagrams, acceptance criteria, and cross-spec stubs. Final version bump to `5.13.0`.

- **Diagrams** — three new `.mmd` files in `diagrams/` (all carry the standard NON-AUTHORITATIVE PROJECTION banner):
  - `erd-backup-tier.mmd` v1.0.0 — projects all 10 Backup-tier App-DB tables (`SyncOpLedger`, `BackupPairing`, `BackupKeyEpoch`, `BackupSyncWatermark`, `BackupOutboxEnvelope`, `BackupApplyIdempotency`, `BackupApplyDeadLetter`, `BackupSnapshotCatalog`, `BackupSnapshotJob`, `BackupRestoreJob`) with PascalCase + INTEGER PKs + Notes/Comments per Rule 11 / Description per Rule 10.
  - `seq-incremental-backup.mmd` v1.0.0 — primary → backup CDC flow: trigger → ledger → outbox seal → BE-1 → 5-stage Apply pipeline (with V7 idempotency branch) → watermark advance + ACK; explicit DLQ note (no silent skips).
  - `seq-backup-restore.mmd` v1.0.0 — operator restore-by-date: BE-3 enqueue (with `MAIN-830-01/02` failure branches) → snapshot decrypt under HKDF `"BackupSnapshot/v1"` → re-seal under current Active KeyEpoch → BE-6 inbox import → watermark realignment.
- **Diagrams index** (`diagrams/readme.md`) bumped to v1.1.0 — three new rows added to both the authoritative-source table and the user-facing tables; ERDs and Sequence Diagrams sections both extended.
- **Acceptance criteria** (`97-acceptance-criteria.md`) bumped to v1.1.0 — new section **"Backup-tier acceptance (Phases 7–11)"** with 13 criteria covering: CDC capture, KeyEpoch enforcement, S2S `421 Misdirected Request` enforcement, V7 idempotency, DLQ-no-silent-skip (CODE RED), `sqlite3_backup_init` integrity, distinct HKDF salts for envelope vs snapshot, forward-secrecy on restore, 30-day retention with never-auto-shorten, watermark realignment after restore, mandatory `PairingId` JWT claim, Rules 10/11/12 compliance, linter rule promotion.
- **Cross-spec stubs** (deferred to Phase 12 by Phases 9–11 changelogs):
  - `05-auth-and-2fa.md` §S2S — note pending: cite `21-backup-endpoints.md` §3 for the `Backup` audience and 5 scopes (`Backup.Diff.Write`, `Backup.Rotate.Write`, `Backup.Restore.Write`, `Backup.Restore.Apply`, `Backup.Read`).
  - `12-jwt-delivery-contract.md` — note pending: document mandatory `PairingId` claim on `Backup`-audience tokens (mismatch → `MAIN-800-04`).
  - `06-core-api-endpoints.md` §2 — note pending: merge BE-1..BE-6 catalogue rows from `21-…` §2 + `23-…` §8 into the canonical endpoint table.
- **Linter promotion** — `96-linter-audit.md` to lift the `BACKUP-*` and `DB-SYNCOP-*` rule families from "draft" to "enforced in CI" (referenced by acceptance criteria; promotion follows the standard linter-scripts cycle per memory rule).
- **Seed promotion** — `AppBackupTrackedTable` seed referenced by acceptance criterion 1 to land via the same migration as `BackupApplyIdempotency` UNIQUE-on-`EnvelopeId` lock (no schema change in this phase).
- **Open questions still pending** (non-blocking, carried into post-5.13.0 maintenance):
  - OQ-23-1 — snapshot dedup pyramid for low-write primaries.
  - OQ-23-2 — partial-table restore.
  - OQ-23-3 — `PinReason` column on `BackupSnapshotCatalog` for the `Pinned` status.
- **Version bump** — `5.13.0-phase11` → **`5.13.0`** (final). Phase suffix removed; the Backup System spec arc is now feature-complete.

**Closes:** Phases 7–11 (`18-…` through `23-…md`). The 19-main-worker-service spec folder now contains the full Backup System contract (24 numbered files: `00-…23` plus `96`/`97`/`98`/`99`).

---

## v2.9.0 — 2026-05-06 (Phase 11 — Snapshot storage + restore flow)

**Scope:** Resolves locked decision **D14** (date-by-date full snapshot storage on backup; main-controlled restore by date). Closes open question **OQ-A4** — snapshot retention adopted at **30 days rolling** (linter floor: 7 days). Final backup-tier spec; only diagrams + acceptance criteria + linter promotion remain (Phase 12).

- New file **`23-snapshot-storage-and-restore.md` v1.0.0** — three-moment lifecycle (Build / Catalogue / Restore), eight-step Build pipeline using SQLite's `sqlite3_backup_init` for transactional consistency, snapshot zip password derived from a separate HKDF salt (`"BackupSnapshot/v1"`) to prevent envelope/snapshot keystream collision, eight-step Restore flow that re-seals the snapshot under the **current Active KeyEpoch** (forward secrecy — never revives a Retired epoch), new `BackupSnapshotCatalog` (entity-ish, Rule 10) + `BackupSnapshotJob` (transactional, Rule 11) + `BackupRestoreJob` (transactional, Rule 11) tables on the backup App tier, retention sweep with `Pinned` status reserved for operator-protected snapshots, never-auto-shorten guarantee under disk pressure.
- New endpoint **BE-6** `POST /API/V1/Backup/RestoreInbox` hosted on the **primary** Worker — symmetric counterpart to BE-1 but flowing backward; uses scope `Backup.Restore.Apply`; bypasses `22-backup-apply-logic.md` Stage-4 dispatch (offline App-tier import).
- Final wiring of the **`Backup` S2S audience** reserved by Phase 9 §9: 5 scopes (`Backup.Diff.Write`, `Backup.Rotate.Write`, `Backup.Restore.Write`, `Backup.Restore.Apply`, `Backup.Read`); mandatory `PairingId` JWT claim; mismatch short-circuits with `MAIN-800-04`.
- `13-error-codes.md` → **v1.5.0**: §2.10 extended with `WORKER-940-01..04` (`SnapshotQuiesceTimeout` 21204, `RestoreImportFailed` 21205, `SnapshotBuildTimeout` 21206, `SnapshotSealFailed` 21207). §3.11 extended with `MAIN-840-02 SnapshotCorrupt` (21192). Reserved-range table refreshed; `MAIN-21193-21199` reserved for future overflow.
- `15-tunable-constants.md` → **v1.10.0**: new §2.15 — `SnapshotRetentionDays=30` (resolves OQ-A4), `Snapshot.BuildHourUtc=3`, `Snapshot.QuiesceTimeoutSeconds=120`, `Snapshot.MaxBuildSeconds=1800` (30 m), `Restore.PrimaryAckTimeoutSeconds=600` (10 m). All Backup-tier tunables now allocated.

**Cross-spec impact:**
- `05-auth-and-2fa.md` §S2S — Phase 12 cleanup will add a one-line stub citing `23-…` §9 for the `Backup` audience (no schema change needed; audience names are config).
- `12-jwt-delivery-contract.md` — Phase 12 cleanup will document the mandatory `PairingId` claim on `Backup`-audience tokens.
- `06-core-api-endpoints.md` §2 — Phase 12 cleanup will merge BE-1..BE-6 catalogue rows from `21-…` §2 + `23-…` §8 into the canonical endpoint table.
- ER diagram regen deferred to Phase 12 — Worker ER must show `BackupSnapshotCatalog`, `BackupSnapshotJob`, `BackupRestoreJob`.
- A successful restore (R7) **resets** the incremental watermark by definition — `BackupSyncWatermark.LastAcceptedSyncOpSeq` is realigned to the snapshot's max `SyncOpSeq` so subsequent BE-1 deliveries continue without re-shipping pre-snapshot rows.

**Decisions resolved (this phase):**
- D14 — fully spec'd (date-named files, Main-controlled restore-by-date).
- OQ-A4 — **30 days rolling** with operator override and 7-day compliance floor.

**Open questions still pending:**
- OQ-23-1 (snapshot dedup pyramid for low-write primaries), OQ-23-2 (partial-table restore), OQ-23-3 (`PinReason` column for `Pinned` status) — all logged in `23-…` §14, non-blocking; OQ-23-3 will be picked up by the Phase-12 migration.

---

## v2.8.0 — 2026-05-06 (Phase 10 — Backup apply pipeline)

**Scope:** Server-side processing pipeline that runs on the backup node once BE-1 (`21-backup-endpoints.md` §4) accepts a sealed envelope. Wire is owned by Phase 9, encryption by Phase 8, CDC source-side by Phase 7. Snapshot/restore remains Phase 11.

- New file **`22-backup-apply-logic.md` v1.0.0** — five-stage strictly-sequential pipeline (Decrypt → Open → Validate → Dispatch → Persist ACK), seven validation rules V1–V7, single-TX `BEGIN IMMEDIATE` per envelope with idempotent dispatch (`Insert`/`Update` = upsert, `Delete` = absent-row tolerated), explicit DLQ on any failure (no silent skips per CODE RED), V7 idempotency short-circuit using a `UNIQUE` constraint as the lock (no advisory mutexes). Two new App-tier tables on the backup: `BackupApplyIdempotency` and `BackupApplyDeadLetter`, both with `{TableName}Id` PK + `Notes`/`Comments TEXT NULL` (transactional Rule 11) + INTEGER `*At` (D2). CODE-RED-compliant per-row pseudocode with positively-named guards (`AssertKnownSyncOp`, `AssertKnownTable`, `AssertNonEmptyPk`).
- `13-error-codes.md` → **v1.4.0**: §2.10 extended with four new Worker apply codes `WORKER-930-01..04` opening a fresh overflow window `WORKER-21200-21299` (per §1 Slot-overflow rule, since `WORKER-21095-21099` was fully consumed by Phase 8). §3.11 added with `MAIN-840-01 BackupApplyExhausted` consuming the first slot of the Phase-11-reserved window (`MAIN-21191`); reserved-range table refreshed — `MAIN-21192-21199` now reserved for snapshot/restore.
- `15-tunable-constants.md` → **v1.9.0**: new §2.14 with four apply-pipeline keys — `MaxRetriesPerEnvelope=5`, `TransactionTimeoutSeconds=30`, `DeadLetterRetentionDays=30`, `IdempotencyRowRetentionDays=14`.

**Cross-spec impact:**
- `BackupApplyIdempotency` + `BackupApplyDeadLetter` are App-tier-local on the backup; the cross-tier reconciliation file (`11-…`) does not need an entry.
- BE-1's idempotency short-circuit (V7) tightens the contract referenced in `21-…` §4.4 — replay returns the **stored** `OriginalResponseJson`, not a freshly-recomputed body.
- Tracked-table allowlist (`AppBackupTrackedTable` ref) is reserved for the Phase 12 seed; `BACKUP-APPLY-003` linter will enforce membership.
- `MAIN-840-01` is surfaced via BE-5 Health (`21-…` §8) — no new endpoint surface in Phase 10.
- ER diagram regen deferred to Phase 12 — Worker ER must show `BackupApplyIdempotency` + `BackupApplyDeadLetter`.

**Open questions still pending:**
- **OQ-A4** — Snapshot retention policy (Phase 11).
- OQ-22-1 (per-envelope WAL pragma), OQ-22-2 (DLQ auto-sweep semantics), OQ-22-3 (tracked-table allowlist seeding strategy) logged in `22-…` §12, non-blocking.

---

## v2.7.0 — 2026-05-06 (Phase 9 — Backup endpoints contract)

**Scope:** Wire surface for Phases 6–8. Five S2S OAuth-protected HTTP endpoints hosted on the backup node, all Main-triggered. Apply logic remains Phase 10; snapshot storage / retention remains Phase 11.

- New file **`21-backup-endpoints.md` v1.0.0** — `BE-1 IncrementalDiff` (multipart upload of sealed Phase-8 envelope; ACKs `LastAcceptedSyncOpSeq` back into `BackupSyncWatermark`), `BE-2 RotateKeys` (steps S3/S6 of the Pair-RSA rotation flow), `BE-3 RestoreByDate` (202-Accepted enqueue, returns `RestoreJobId`), `BE-4 Snapshots` (date-bounded catalogue), `BE-5 Health` (single-call dashboard surface; never throws on degradation). Defence-in-depth `421` re-asserted at proxy. Endpoint↔scope matrix introduces `Backup.Diff.Write`, `Backup.Rotate.Write`, `Backup.Restore.Write`, `Backup.Read` scopes plus a new `Backup` audience to be wired into `05-…` §S2S in Phase 11. CODE-RED handler size budgets pinned per endpoint.
- `13-error-codes.md` → **v1.3.0**: §3.10 added with two new wire-only Main codes — `MAIN-830-01 SnapshotNotFound` (21189, 404) and `MAIN-830-02 RestoreAlreadyInProgress` (21190, 409). Reserved-range table refreshed; `MAIN-21191-21199` now reserved for Phase 11 snapshot/restore overflow.
- `15-tunable-constants.md` → **v1.8.0**: new §2.13 with five backup-endpoint timeouts — `IncrementalDiffTimeoutSeconds=120`, `RotateKeysTimeoutSeconds=30`, `RestoreByDateTimeoutSeconds=60`, `SnapshotsTimeoutSeconds=15`, `HealthTimeoutSeconds=5`.

**Cross-spec impact:**
- `06-core-api-endpoints.md` §2 receives a paste-ready `2.X Backup` table merge in Phase 12 cleanup; this file is the source of truth in the interim.
- `MAIN-830-*` rows are wire-side only here; their storage semantics (filesystem layout, retention sweep) are owned by `22-snapshot-storage-and-restore.md` (Phase 11).
- ER diagram regen deferred to Phase 12 — no schema change in Phase 9 (BE-1 writes are confined to `BackupSyncWatermark` already in `19-…`; BE-3 enqueues a job into the existing worker job table).

**Open questions still pending:**
- **OQ-A4** — Snapshot retention policy (Phase 11).
- OQ-21-1 (streaming vs. multipart for BE-1 at >100 MB envelopes) and OQ-21-2 (BE-5 scope vs. unauth proxy probe) logged in `21-…` §14, non-blocking.

---

## v2.6.0 — 2026-05-06 (Phase 8 — Backup encryption and Pair-RSA key rotation)

**Scope:** Per locked decision **D13** (RSA pair shared between Worker and its Backups; Main issues rotation; zip password follows known pattern). Resolves open question **OQ-A3** (zip password derivation = `HMAC-SHA256(SharedSecret, EnvelopeTimestampEpoch)` truncated to 32 hex chars). Endpoints / apply / restore remain Phases 9–11.

- New file **`20-backup-encryption-and-keys.md` v1.0.0** — three-artefact key inventory (Pair-RSA / Envelope-AES / Zip-Password), envelope sealing pipeline (AES-256-GCM body + RSA-OAEP wrap + RSA-PSS sign + AES-256-ZIP outer), HKDF-derived deterministic zip password resolving OQ-A3, four-state `Pending → Active → Retired → Discarded` rotation state machine, eight-step Main-orchestrated rotation flow with split-brain alerting, `BackupKeyEpoch` table on both primary and backup (Memory: PascalCase + `{TableName}Id` PK + nullable `Description`, INTEGER `*At` per D2), defence-in-depth verification path on the backup (epoch lookup + cipher refusal + signature verify + GCM decrypt).
- `13-error-codes.md` → **v1.2.0**: §2.10 extended with five new Worker decrypt codes `WORKER-920-01..05` (21095-21099 — fully consuming the Worker future-expansion range), §3.9 added with three new Main rotation-orchestration codes `MAIN-820-01..03` (21186-21188). Reserved-range table refreshed; future-expansion `MAIN-21186-21199` narrows to `MAIN-21189-21199`.
- `15-tunable-constants.md` → **v1.7.0**: new §2.12 with five backup-encryption keys — `MaxKeyAgeSeconds=7776000` (90 d), `RotationAckTimeoutSeconds=120`, `RotationActivationDelaySeconds=60`, `RetiredKeyGraceSeconds=86400` (24 h), `RsaKeySizeBits=4096`.

**Cross-spec impact:**
- App-tier mirror: `BackupKeyEpoch` is added on both primary and backup Worker App tiers; the cross-tier reconciliation file (`11-…`) does not need a new entry because App-tier additions are local. Main holds the row too but with `PrivateKeyPem` always NULL (public halves only).
- ER diagram regen deferred to Phase 12 — Worker ER must show `BackupKeyEpoch` with the four-state lifecycle.
- `19-incremental-backup-sync.md` §6 envelope SQLite is now the input artefact to `20-…` §4 step 1 — no schema change.
- Phase 9 (endpoints) will surface `POST /API/V1/Backup/RotateKeys` as the operator-forced rotation trigger named in `20-…` §7.1.

**Open questions still pending:**
- **OQ-A4** — Snapshot retention policy (Phase 11).
- OQ-20-1 (split-brain pager routing) and OQ-20-2 (RSA-4096 vs Ed25519+X25519) logged in `20-…` §14, non-blocking.

---

## v2.5.0 — 2026-05-06 (Phase 7 — Incremental backup sync, CDC)

**Scope:** Per locked decision D10 (`SyncOp` flag on synced rows). Defines the change-data-capture mechanic that lets a primary Worker ship deterministic, replayable diffs to each attached backup. Encryption / wire / apply / restore remain Phases 8–11.

- New file **`19-incremental-backup-sync.md` v1.0.0** — two `SyncOp` shapes (inline column vs. `BackupSyncLog` side table), `SyncOp` ref catalog, per-database monotonic `BackupSyncSequence` allocator, `BackupSyncWatermark` per-attached-backup pointer, CODE-RED-compliant diff-generation driver (resume from `LastAcked`, not `LastShipped`), envelope as a SQLite file with two tables (`Envelope`, `EnvelopeRow`), compaction policies for both shapes with the safety rule "reclaim only past `MIN(LastAckedSyncOpSeq)`", linter hooks `DB-SYNCOP-001/002` queued for Phase 12.
- `13-error-codes.md` — three new Worker codes (`WORKER-910-01..03`, 21092-21094) and one Main code (`MAIN-810-01 BackupCompactionStalled`, 21185). Reserved-range table updated; future-expansion ranges are now `WORKER-21095-21099` and `MAIN-21186-21199`.
- `15-tunable-constants.md` → **v1.6.0**: §2.11 extended with five new keys — `SyncIntervalSeconds=60`, `MaxRowsPerEnvelope=5000`, `TombstoneRetentionSeconds=604800`, `LogRetentionSeconds=604800`, `QuarantineCompactionOverrideSeconds=86400`.

**Cross-spec impact:**
- App-tier tables that participate in backup mirroring will need either Shape A columns (`SyncOpCode`, `SyncOpSeq`, `SyncOpAt`) or a write-side hook into `BackupSyncLog`. The concrete tracked-table list is a Phase-12 follow-up (seed file + `DB-SYNCOP-001` linter).
- `KnownBackupNode.LastSyncWatermark` (Phase 6) is reframed as a denormalized view of `BackupSyncWatermark.LastAckedSyncOpSeq` for human dashboards; the authoritative pointer is the new `BackupSyncWatermark` table.
- ER diagram regen deferred to Phase 12 — Worker ER must show `SyncOp`, `BackupSyncLog`, `BackupSyncWatermark`, `BackupSyncSequence`.

**Open questions still pending:**
- **OQ-A3** — Backup zip password derivation (Phase 8).
- **OQ-A4** — Snapshot retention policy (Phase 11).

---

## v2.4.0 — 2026-05-06 (Phase 6 — Backup nodes concept)

**Scope:** Per locked decisions D8 / D9 / D10 (CDC referenced; defined in Phase 7). Defines what a backup node is, how it registers (extends `10-worker-bootstrap-protocol.md`), how Main propagates the pairing to both ends, and the three independent enforcement points for the "backups never serve traffic" invariant. Wire format / encryption / endpoints / restore are explicitly deferred to Phases 7–11.

- New file **`18-backup-nodes.md` v1.0.0** — Kubernetes-style replica framing, three-tier relationship model (R1/R2/R3 facts), registration request/response additions, Main-side acceptance procedure (CODE RED ≤15 lines), `KnownBackupNode` Worker App-tier mirror table, defence-in-depth `421 Misdirected Request` rule for the no-traffic invariant.
- `13-error-codes.md` — new §3.8 "Backup Lifecycle" series: `MAIN-800-01 BackupChainNotAllowed` (21181, 422), `MAIN-800-02 PrimaryNotFound` (21182, 404), `MAIN-800-03 BackupCapacityExceeded` (21183, 409), `MAIN-800-04 TrafficOnBackupRejected` (21184, 421). Reserved-range table updated; future-expansion ranges narrowed to `MAIN-21172-21180` and `MAIN-21185-21199`.
- `15-tunable-constants.md` → **v1.5.0**: new §2.11 "Backup nodes" with `MainWorker.Backup.MaxBackupsPerPrimary=3`, `MainWorker.Backup.LagWarningSeconds=900`, `MainWorker.Backup.HeartbeatIntervalSeconds=60`.
- `14-rbac-and-status-seed.md` — `WorkerNodeStatus` seed bumped to v1.5.0; row count 4 → 7. Added `Provisioning` (backup just registered, awaiting first diff), `BackupAttached` (healthy backup), `BackupLagging` (backup lag exceeds tunable). Existing primary-only codes annotated as never-assigned-to-backups.

**Cross-spec impact:**
- `WorkerNode` schema (Phase 4) is the structural enabler — no further DB changes in Phase 6.
- `KnownBackupNode` is added to the Worker App tier; the cross-tier reconciliation file (`11-…`) does not need a new entry because App-tier additions are local to the Worker.
- ER diagram regeneration deferred to Phase 12 — Worker ER must show `KnownBackupNode`.

**Open questions still pending:**
- **OQ-A3** — Backup zip password derivation (Phase 8).
- **OQ-A4** — Snapshot retention policy (Phase 11).

---

## v2.3.0 — 2026-05-06 (Phase 5 — Cascading roles + Role-Access cache bin)

**Scope:** Per locked decisions D11 (cascading = union) and D12 (cache-bin in ER). Adopts default proposals for OQ-A1 (simple union, no inheritance) and OQ-A2 (per-process SQLite `:memory:` storage with TTL + Main-broadcast invalidation) until the user overrides.

- New file **`17-cascading-roles-and-cache-bin.md` v1.0.0** — single source of truth for:
  - The union rule for users holding multiple roles (bitwise-OR of `CanRead` / `CanWrite` per AccessItem).
  - Two-tier resolution: catalog stays on Main, per-user resolution + cache live on Worker.
  - Cache-bin schema (`RoleAccessCache`, `RoleCacheCatalogVersion`) in the Worker's in-memory Cache tier.
  - Invalidation broadcast `POST /API/V1/Cache/InvalidateRoleAccess` (idempotent on `CatalogVersion`, retry per §2.1, no rollback on delivery failure — TTL bounds staleness).
  - JWT staleness mitigations: short TTL + `CatalogVersion` stamp + optional `RequireReauthOnCatalogBump`.
- `15-tunable-constants.md` → **v1.4.0**: new §2.10 "Role-access cache bin" with `MainWorker.RoleCache.TtlSeconds` (600 s default) and `MainWorker.RoleCache.RequireReauthOnCatalogBump` (false default).
- `13-error-codes.md`:
  - New §2.10 "Cache Coherence" (Worker): `WORKER-900-01 RoleCacheRecompileFailed` (21090, 500), `WORKER-900-02 EmptyEffectiveAccessSet` (21091, 403).
  - New §3.7 "Cache Coherence" (Main): `MAIN-700-01 CacheInvalidationDeliveryFailed` (21171, 502).
  - Reserved sub-range table updated: 21090-21091 marked consumed; 21171 marked consumed; future-expansion ranges narrowed accordingly.

**Cross-spec impact:**
- Worker JWT mint contract gains `CatalogVersion` claim + read/write AccessItem code arrays. `12-jwt-delivery-contract.md` will need a Phase-12 follow-up entry to document the claim shape (added to the Phase-12 punch list).
- ER diagram regeneration deferred to Phase 12 — Worker ER must show `RoleAccessCache` and `RoleCacheCatalogVersion` (Cache tier, in-memory annotation); Main ER must show the new `RoleAccessInvalidationEvent` audit table once authored in Phase 12.

**Open questions resolved with default proposals (overridable):**
- **OQ-A1** — Cascading semantics → adopted **simple union**.
- **OQ-A2** — Cache-bin tech → adopted **per-process SQLite `:memory:`** behind a swappable contract.

**Open questions still pending (carried into Phase 8 / Phase 11):**
- **OQ-A3** — Backup zip password derivation pattern.
- **OQ-A4** — Snapshot retention policy.

---

## v2.2.0 — 2026-05-06 (Phase 4 — WorkerNode backup & ordering, "Region" UI label)

**Scope:** Per locked decisions D6, D7, D8, D9 — give `WorkerNode` the structural fields needed to express the backup-node concept and the deterministic ordering needed by RoundRobin, and rename the user-facing column to "Region" without touching code identifiers.

- `03-main-db-schema.md` → **v2.2.0**:
  - `WorkerNode` (§2.1) gains `Sequence INTEGER NOT NULL` (RoundRobin order, unique among non-backup peers), `IsBackup INTEGER NOT NULL DEFAULT 0`, `BackupOfWorkerNodeId INTEGER NULL` (self-FK).
  - CHECK constraints: backup-flag and FK move together (`(IsBackup=0 AND BackupOfWorkerNodeId IS NULL) OR (IsBackup=1 AND BackupOfWorkerNodeId IS NOT NULL)`); backup chains forbidden (referenced row MUST have `IsBackup=0`, enforced by trigger).
  - New indexes: `IX_WorkerNode_BackupOf` and partial `IX_WorkerNode_PrimaryEligible (WorkerNodeStatusId, Sequence) WHERE IsBackup = 0`.
- `04-worker-routing.md` → **v1.2.0**: §1.1 RoundRobin walks `Sequence ASC`; §1.4 eligibility filter prefixed with positive guard `IsPrimary(node) → IsBackup = 0`. Manual strategy now rejects backup targets with `WORKER-300-04 BackupNotRoutable`.
- `13-error-codes.md`: added `WORKER-300-04 / 21033 / BackupNotRoutable` (HTTP 409).
- `07-role-based-dashboards.md` → **v2.1.0**: new §9 "UI Labels" — `WorkerNode` renders as **"Region"** in dashboards, forms, and audit views via i18n key `worker_node.label`. Code, API, and DB identifiers unchanged.

**Cross-spec impact:**
- Worker bootstrap (`10-worker-bootstrap-protocol.md`) and self-update pointer (`09-self-update-pointer.md`) are unchanged for primary nodes; backup-node registration / pairing flow is deferred to Phase 6 (`17-backup-nodes.md`).
- ER diagram regeneration deferred to Phase 12 (`diagrams/erd-main-db.mmd`).
- Cache-bin tables for role resolution and the cascading-roles union semantics remain Phase 5 work.

**Open questions carried into Phase 5:** OQ-A1 (cascading semantics — union vs hierarchy), OQ-A2 (cache-bin tech), OQ-A3 (zip password derivation), OQ-A4 (snapshot retention).

---

## v2.1.0 — 2026-05-06 (Phase 3 — Move Users off Main)

**Scope:** Per locked decision D5, Main becomes credential-blind. All identity, password, and 2FA state moves to the assigned Worker's split-DB App tier. Spec-only; no runtime code touched.

- `03-main-db-schema.md` → **v2.1.0**:
  - **REMOVED** `User` table and all auth columns (`UserPasswordHash`, `UserPasswordSalt`, `UserTotpSecret`, `UserTotpEnrolledAt`, `UserTotpBackupCodesHash`).
  - **REMOVED** `UserRole` join table (assignments now live on Worker as `AppUserRole`).
  - **ADDED** `UserDirectory` (§2.4) — routing-only index `(UserDirectoryId, UserEmail, CompanyId, WorkerNodeId, CreatedAt, LastSeenAt, Description)`. Carries no secrets and no PII beyond email.
  - `AccessDenialEvent` (§2.6.3): `UserId` FK replaced by `UserDirectoryId` (nullable) + snapshotted `ActorEmail`. `AccessItemId` FK retained (catalog stays on Main).
  - `EndpointAuthAuditEvent` (§2.6.4): `UpdatedByUserId` FK replaced by `UpdatedByUserDirectoryId` + snapshotted `UpdatedByUserEmail`.
  - Indexes: `IX_User_CompanyId` removed; new `IX_UserDirectory_CompanyId`, `IX_UserDirectory_WorkerNodeId`, `UX_UserDirectory_UserEmail`. `IX_EndpointAuthAuditEvent_Actor_At` re-pointed to `UpdatedByUserDirectoryId`.
  - §4 "What Main DB does NOT store" — added explicit invariant that Main carries no password/TOTP/role-assignment material; grep over Main for `password|totp|secret|hash` MUST return zero column hits.
  - §5 "Migration Notes" — added v2.1.0 forward-only migration script that backfills `UserDirectory`, forwards credentials to each Worker via `MigrateLegacyUsers` bootstrap instruction, and deletes Main `User`/`UserRole` rows only after Worker ACK.
- `05-auth-and-2fa.md` → **v2.0.0**: Main rewritten as credential-blind reverse proxy. New §2.1 (proxy flow with constant-time email-miss handling and post-forward buffer-zero), §2.2 (Worker mints JWT; `iss` flips to Worker URL), §3 (password storage moved to Worker `AppUser`), §4 (TOTP storage moved to Worker), §5–§6 (sign-up/sign-in flows reframed as Main → `Worker /Auth/InternalSignUp` / `/Auth/InternalSignIn` over the credential-proxy channel). `JwtExpiresAt` example flipped to epoch seconds per Rule 7.1 v2.
- `11-split-db-tier-reconciliation.md` → **v1.1.0**: Main §4 — `User` and `UserRole` struck through with the v2.1.0 removal note; `UserDirectory` added to Root tier; `Role`, `AccessItem`, `RoleAccessItem` reaffirmed as Settings-tier **catalogs** (kept on Main, mirrored read-only to each Worker). Worker §5 — `AppUser` annotated as authoritative identity store, `AppUserRole` added as the user→role join.

**Cross-spec impact:**
- Any service reading `MainDB.User.*` MUST switch to either (a) `MainDB.UserDirectory` (routing only) or (b) `WorkerDB.AppUser` (credentials, identity).
- The `/API/V1/Auth/RefreshWorkerToken` endpoint on Main is **deprecated**; React MUST refresh JWTs by calling Worker `/API/V1/Auth/RefreshToken` directly.
- Audit consumers joining `EndpointAuthAuditEvent` on `User.UserId` MUST switch to `UserDirectory.UserDirectoryId` (or fall back to `UpdatedByUserEmail` for hard-deleted directory rows).

**Open questions carried into Phase 4:** OQ-A1 (cascading semantics — union vs hierarchy), OQ-A2 (cache-bin tech), OQ-A3 (zip password derivation), OQ-A4 (snapshot retention).

---

## v2.0.0 — 2026-05-06 (Phase 2 — DB convention overhaul)

**Scope:** Apply the global DB convention upgrades from `spec/04-database-conventions/` v2 to the Main schema. Spec-only; no runtime code touched.

> **Clarification (post-edit):** Naming **Rule 1** is universal and is **not** relaxed by Rule 13 — every PK on every table is still `{TableName}Id` (e.g. `WorkerNodeStatusId`, `RoleId`, `EndpointAuthChangeKindId`, `WorkerSelectionStrategyId`). Rule 13's "simplification" applies **only** to the descriptive columns `Code`, `Label`, and `Description`, which drop the `{Table}` prefix because those columns never travel as FKs. `02-schema-design.md` §6.5 was rewritten to make this explicit, and the `WorkerNodeStatus` / `WorkerNodeKind` example in `03-main-db-schema.md` §2.2 was expanded to show the full PK names rather than a `{TableName}Id` placeholder.

- `03-main-db-schema.md` → **v2.0.0**:
  - All `*At` columns flipped from `TEXT` (ISO-8601) to `INTEGER` (epoch seconds, UTC) per Naming Rule 7.1 v2: `WorkerNodeRegisteredAt`, `WorkerNodeLastSeenAt`, `CompanyAssignedAt`, `UserCreatedAt`, `UserTotpEnrolledAt`, `AccessDenialEvent.OccurredAt`, `EndpointAuthAuditEvent.OccurredAt`, `WorkerVersionRecordedAt`, `WorkerSelectionEventAt`. (Removes the temporary "TEXT or INTEGER" wording introduced in v1.4.0 on `AccessDenialEvent.OccurredAt`.)
  - All ref / enum-like tables flattened to canonical `(Id, Code, Label)` per Rule 13: `WorkerNodeStatus`, `WorkerNodeKind`, `Role`, `EndpointAuthChangeKind`, `WorkerSelectionStrategy`. Old `{Table}Code` / `{Table}Label` column names are removed in this spec.
  - `Company.CompanySlug` → `Company.Slug`; `Company.CompanyName` → `Company.Name`. Unique index updated to `(Slug)`. Seedable-Config inbound-name aliases accepted through v2.1.0 then removed.
  - Added Phase-3 banner over §2.4 `User`: `User`, `UserRole`, and TOTP columns will move off Main entirely in v2.1.0 (D5).
- `spec/04-database-conventions/01-naming-conventions.md` → **Rule 7.1 rewritten as v2** ("Epoch-INTEGER Timestamp"). Old TEXT/ISO-8601 storage rule deprecated and forbidden for new schemas. Examples table and "Complete Example" code block updated to `INTEGER NOT NULL DEFAULT (unixepoch())`.
- `spec/04-database-conventions/02-schema-design.md` → §6.4 examples updated to epoch defaults; new **§6.5 Rule 13 — Enum / Lookup Table Canonical Shape `(Id, Code, Label)`** with column table, rationale, lookup pattern, and forbidden alternatives. Template row in §5 updated.
- `spec/05-split-db-architecture/01-fundamentals.md` → **v3.4.0**: convention-propagation banner added stating that every tier (Root / Settings / App / Session / Cache / Document) inherits Rule 7.1 v2 + Rule 13.

**Cross-spec impact:** Any consumer reading `WorkerNodeStatusCode` / `RoleCode` / `RoleLabel` / `CompanySlug` / `CompanyName` / ISO-8601 `*At` strings MUST migrate. Suggested migration: `unixepoch(<OldName>)` for backfill, then drop the old columns in the next minor.

Linter status: column renames are structural; existing R2 / R3 waivers in `13-error-codes.md` unaffected.

---

## v1.4.0 — 2026-05-06 (Phase 1 — `EnumPage` → `AccessItem` rename)

**Scope:** Schema + seed + dashboard rename only. No runtime code touched (per memory rule "Spec/19 is SPEC-ONLY").

- `03-main-db-schema.md` → **v1.4.0**: §2.6.1 renamed `EnumPage` → `AccessItem`; columns flattened from `EnumPageId/EnumPageCode/EnumPageLabel/Description` to `AccessItemId/Code/Label/PageUrlSuffix/Description`. New `PageUrlSuffix TEXT NULL` column is the route matcher (suffix match against normalized request path). §2.6.2 renamed `RolePageAccess` → `RoleAccessItem` with FK column `AccessItemId`. §2.6.3 `AccessDenialEvent.EnumPageId` → `AccessItemId`; `OccurredAt` flagged for INTEGER conversion in Phase 2.
- `14-rbac-and-status-seed.md` → **v2.0.0**: full seed JSON rewritten for `AccessItem` + `RoleAccessItem`. Each AccessItem row carries `Code`, `Label`, `PageUrlSuffix` (e.g. `/admin`, `/billing`, `/regions`). 19 `RoleAccessItem` grant rows now include explicit `CanRead`/`CanWrite`. Verification SQL counts updated.
- `07-role-based-dashboards.md` → **v2.0.0**: PHP `enum AccessItem` cases shortened to bare codes (`PowerAdmin`, `Admin`, `Billing`, …) — no `Page` suffix. Access-check function renamed `userHasAccessToPage` → `userHasAccessToItem`. Middleware param `$pageCode` → `$accessItemCode`. §4 deduplicated (no longer redefines columns; refers to `03-…` §2.6).
- **Deprecation contract:** Old names `EnumPage` / `RolePageAccess` accepted as seed-loader aliases through v1.4.x; removal scheduled for v1.5.0.
- **Cross-spec impact:** None outside `19-…`. Phase 2 will propagate INTEGER DateTime convention which removes the temporary "TEXT or INTEGER" wording on `AccessDenialEvent.OccurredAt`.

Linter status: structural rename only — seed `Tables` block validates against `06-seedable-config-architecture/02-features/07-reference-table-seeding.md`. No error-code changes.

---

## v1.3.0 — 2026-05-05 (FU-18 EndpointAuthLocked error code)

- `13-error-codes.md` → **v1.1.0**: +§3.4 row `MAIN-400-10 EndpointAuthLocked` / flat `21170` / HTTP 403, message "Endpoint pattern matches the lock-list (`/API/V1/Workers/*` or `/API/V1/SelfUpdate`) and cannot be reconfigured via `PATCH /API/V1/Settings/EndpointAuth`." Source: `06-core-api-endpoints.md` §5.4 R-5 + `05-auth-and-2fa.md` §8. Added §1 *Slot-overflow rule* documenting the first allocation that breaks strict `211{YY}` mapping (4xx routing flats `21140-21149` were exhausted by tasks #32 + #39, so the new code took `21170` from the `MAIN-21170-21199` reserved range). §4 reserved-range table refreshed: `21170` marked consumed, residual reserve narrowed to `MAIN-21171-21199` plus a new `MAIN-21162-21169` external-services band.
- `error-codes.json` → **v1.2.0**: +entry for `MAIN-400-10` with all 8 fields (Code/Flat/Name/HttpStatus/Tier/Message/Source/Retryable=false). `TotalCodes` 48 → 49. `Generated` 2026-05-04 → 2026-05-05.
- `06-core-api-endpoints.md` §5.4 R-5 + §5.7 cross-refs: dropped "to be catalogued / to be assigned" hedging; both now cite the assigned `MAIN-400-10` / `21170` slot directly. (No version bump — text-only refinement to v1.2.0 of the same file.)

Linter verification (4/4 green): `check-mws-error-codes` (R1-R4 — 52 codes verified, 21 R2 waivers loaded; new code has 2 source references so no waiver needed), `check-spec-cross-links`, `check-spec-folder-refs`, `check-tunable-constants`. Closes FU-18.

---

## v1.2.0 — 2026-05-05 (FU-17 audit-trail wiring)

---

## v1.1.0 — 2026-05-04 (spec hardening; tasks #07–35)

26 spec-hardening tasks executed against the 5-step audit suite (`audit/01..05`). Headline: **all 26 BLOCKERs → 0**, **all 27 MAJORs → 0** (1 deferred to OQ-1), 76 MINORs → small residual. No breaking schema or contract changes; all additions are clarifications or codifications of previously implicit rules.

### Added — new spec files

- `10-worker-bootstrap-protocol.md` (v1.0.0) — 8-step deterministic boot, `/Workers/Register` contract, JWT public-key fetch (no `/jwks` — static URL + cache), version pinning, `WorkerNode` + `WorkerBootstrapState` schemas, 9 `WORKER-*` error codes. Closes audit F-B-01/02/03, F-X-08. Unblocks AC-1, AC-3, AC-4.
- `11-split-db-tier-reconciliation.md` (v1.0.0) — Pins Main = 3 tiers (Root/Settings/Session), Worker = 4 tiers (Root/Settings/App/Session) per spec/05's 6-tier model. Per-tier table allocation. Closes F-X-01/04, F-D-09. Unblocks AC-2.
- `12-jwt-delivery-contract.md` (v1.0.0) — Worker JWT pinned to JSON-body + in-memory storage (NOT cookie/localStorage), mandatory CSP, claim contract, 9 CI test cases. Closes F-A-12, F-D-04, F-B-05. Closes AC-4.
- `13-error-codes.md` (v1.1.0) — 30 codes (22 `WORKER-*` + 8 `MAIN-*`) catalogued with prefixed↔flat mapping; MWS prefix range `21000-21199` registered in `spec/03-error-manage/03-error-code-registry/01-registry.md`; `error-codes.json` mirror generated. Closes F-X-08, F-A-21, F-B-08. Unblocks AC-6.
- `14-rbac-and-status-seed.md` (v1.0.0) — 3 Roles + 9 EnumPages + 19 RolePageAccess + 4 WorkerNodeStatus + 4 AuthMechanism rows; `@Role.Code` logical-key syntax. Closes F-B-09/10, F-X-06. Closes AC-5.
- `15-tunable-constants.md` (v1.1.0) — 30 numeric tunables (retry, `IdempotencyKeyTtlSeconds=86400`, heartbeat 30s/3-miss, JWT 900s, routing timeouts, rate limits, push-update windows, bootstrap retry, IssuedSkew, SelfUpdate-RedirectStaleHours). `config.seed.json` `MainWorker` category included verbatim. Closes F-A-15, F-A-16, F-B-12, F-M-02/05/08/09, F-N-05. Closes AC-7.
- `96-linter-audit.md` (v1.0.0) — Linter pipeline reference.
- `error-codes.json` — Machine-readable mirror of §13.

### Bumped — root spec files

- `02-glossary.md` → **v1.1.0**: +5 entries (Quarantined, Draining, Seedable-Config superset row, apperror package, Power Admin↔PowerAdmin distinction). Closes F-A-36..40.
- `03-main-db-schema.md` → **v1.2.0**: +`User.UserTotpSecret/UserTotpEnrolledAt/UserTotpBackupCodesHash` (F-A-24/F-B-11); +`EnumPage` (§2.6.1), `RolePageAccess` (§2.6.2), `AccessDenialEvent` (§2.6.3) (F-A-23/F-B-10/F-A-17); +`MainSetting` (F-B-08); +`WorkerSelectionEvent` audit cols (F-B-07).
- `04-worker-routing.md` → **v1.1.0**: §1.2 LeastLoaded tiebreaker by capacity-headroom (F-M-03); §1.4 HasCapacity guard rejects `0`-magic (F-A-06); §5.1 strategy interfaces (F-A-33); inline tunable literals replaced with §15 citations.
- `05-auth-and-2fa.md` → **v1.1.0**: §3 bcrypt-cost env pinning (F-A-03), pepper MUST in prod (F-A-04), breach-check MUST when enabled; §4 backup-codes-at-zero policy + `X-Auth-Action: RegenerateBackupCodes` (F-A-05/F-M-06); §5 `PasswordResetRequest` always-202 anti-enumeration (F-M-07); §6 cookie-scope vs JWT-scope paragraph (F-B-12).
- `06-core-api-endpoints.md` → **v1.1.0**: §3.1 11-row Nullable validation table (F-M-01/F-A-01); §6 rate limits promoted to MANDATORY defaults (F-A-02); §2.5 `/Workers/Register` payload (F-B-02); `/Workers/.../Update` request body (F-B-06).
- `07-role-based-dashboards.md` → **v1.1.0**: §5 stack-agnostic 3-step access-guard contract above the Laravel example + Express equivalent (F-A-34).
- `08-error-contract.md` → **v1.1.0**: §2 envelope +`EnvelopeVersion`/`OperationId`/`SubCode`/`FieldErrors` (F-A-12/15/16/28); §3.4 `X-Auth-Action: Reauthenticate` header (F-A-26); §5 `lastResponse` initialised via `makeNullResponse(call)` (F-A-35); §8 ErrorCode→HTTP-status mapping (F-A-31); §9 Worker→Main envelope + 3 new ErrorCodes `WorkerRegisterRejected/WorkerHeartbeatRejected/WorkerPushAckUnknownJid` (F-A-32); §10 audit-closure log.
- `09-self-update-pointer.md` → **v1.2.0**: bounded sunset (3-way expiry: spec/19 v2.0.0 OR prod-green-14d OR 2026-12-31); §9 deletion checklist (F-A-09); inline tunables replaced with §15 citations.
- `00-overview.md`, `01-architecture.md` → **v1.1.0**: bumped for image-import + tunable citations.

### Cross-spec contributions

- `spec/03-error-manage/03-error-code-registry/01-registry.md` — Registered MWS prefix `21000-21199`.
- `spec/04-database-conventions/01-naming-conventions.md` — Added Rule 7.1 (ISO-8601 precision: `YYYY-MM-DDTHH:MM:SS.sssZ`, mandatory ms + UTC `Z`). Closes F-N-08.
- `spec/04-database-conventions/06-rest-api-format.md` — Promoted `X-Correlation-Id` / `X-Idempotency-Key` / `X-Auth-Action` to authoritative section. Closes F-X-10, F-A-22.
- `spec/06-seedable-config-architecture/02-features/07-reference-table-seeding.md` (new) — Tables-block seed schema with `UpsertByLogicalKey`/`AppendOnly` strategies, `TableSeedMeta`+`TableSeedChangelog` bookkeeping. Closes top-10 fix #6.
- `spec/14-update/28-worker-push-instruction.md` (new) — JID schema, transport, RenameFirst flow, error codes, worker-side `WorkerUpdateInstruction` table. Closes F-X-14/15/17 (top-10 fix #5). Pins `MaxRetries=3`.

### Diagrams

- All 6 diagrams in `diagrams/` carry banner v1.0.0 **NON-AUTHORITATIVE PROJECTION** with citation to authoritative source(s). `diagrams/readme.md` rewritten with conflict-resolution rule + per-file authority table. Closes F-D-01..F-D-12.
- `diagrams/erd-main-db.mmd` → banner v1.1.0: synced to schema v1.2.0 (+`EnumPage`, +`AccessDenialEvent`, +User TOTP triple, `RolePageAccess` upgraded to FK with `CanRead`/`CanWrite`).

### Linters

- New: `linter-scripts/check-tunable-constants.py` (T1 presence + waiver, T2 unique keys, T3 §4↔§2 default parity).
- New: `linter-scripts/check-mws-error-codes.py` (R2 no-orphan).
- `linter-scripts/run.sh` and `run.ps1` rewrote Step 3 — runs all 15 spec/docs linters with `--skip-linters` / `--linters-only` toggles. Pipeline 15/15 green.

### Audit closure

- `audit/01-completeness-audit.md` — re-triaged in §7 (v1.1.0); **30/30 findings closed** (28 fixed + 1 deferred to OQ-1 + 1 deferred post-v1.0).
- `audit/04-cross-spec-dependency-audit.md` — anchor sweep verified clean (task #33).
- `audit/02`, `audit/03`, `audit/05` — partial closure pending re-triage.

### Deferred (post-v1.1.0)

- OQ-1: per-endpoint auth-mechanism overrides (F-M-10) — design awaits user resolution.
- OQ-15-1 / OQ-15-2: ✅ resolved in task #37 (`15-tunable-constants.md` v1.2.0).
- `seq-login-routing.mmd` sync for `X-Auth-Action: Reauthenticate` and `X-Auth-Action: RegenerateBackupCodes` signals: ✅ resolved in task #38 (banner v1.1.0).
- F-N-07: OpenAPI/Swagger artifact generation.

---

## v1.0.0 — 2026-05-04


Initial authoring. Phases 1–4 of the spec roadmap complete.

### Added
- `plan.md` — phased roadmap, locked decisions (Q1–Q5), open questions (OQ-1, OQ-2)
- `00-overview.md` — purpose, scope, stack flexibility, document map
- `01-architecture.md` — topology, request lifecycle, comms contract, caching
- `02-glossary.md` — canonical terms + forbidden-term replacements (`CW configuration` → `Seedable-Config`, `git map` → `gitmap`)
- `03-main-db-schema.md` — 9 tables (WorkerNode, WorkerNodeStatus/Kind, Company, User, UserRole, Role, WorkerVersion, WorkerSelectionEvent/Strategy)
- `04-worker-routing.md` — RoundRobin / LeastLoaded / Manual strategies, eligibility filter, caching, failover
- `05-auth-and-2fa.md` — three auth surfaces (cookie / RS256 JWT / OAuth), Argon2id, TOTP 2FA, OQ-1 flagged
- `06-core-api-endpoints.md` — full REST surface, payloads, update schedule, settings
- `07-role-based-dashboards.md` — `EnumPage` pattern, `RolePageAccess`, three default dashboards, `<RequiresAccess>` wrapper
- `08-error-contract.md` — Main↔Worker envelope, 8-entry failure taxonomy, retry semantics, correlation-ID propagation
- `09-self-update-pointer.md` — pointer-only doc; defers to `spec/14-update/`
- `97-acceptance-criteria.md` — verbatim AC-1..AC-9 mapped to deliverables
- `diagrams/erd-main-db.mmd`, `erd-worker-split-db.mmd`, `erd-seedable-config.mmd`
- `diagrams/seq-company-creation.mmd`, `seq-login-routing.mmd`, `seq-push-update.mmd`
- `diagrams/readme.md`

### Decisions locked
- Tenant root: **Company-as-root** (multi-tenant; user-as-root is degenerate 1:1).
- Spec slot: `spec/19-main-worker-service/` (slots 19–20 free).
- Diagrams home: in-spec `diagrams/` subfolder.
- Error-manage integration: inline contract + reference, no duplication.
- Default stack: Laravel; spec is stack-agnostic (.NET / Go / Python / WordPress also explicitly supported).
- Default worker selection: `LeastLoaded`.
- Worker JWT: RS256, 15-min TTL.
- Password hash: Argon2id (preferred) / bcrypt cost ≥12.

### Deferred
- Self-update implementation (pointer only; lives in `spec/14-update/`).
- Tenant migration between workers (sketched in `04-worker-routing.md` §4, not v1.0).
- OQ-1: per-endpoint auth-mechanism overrides — schema sketched in `06-core-api-endpoints.md` §5; final design awaits user resolution.

---

*Changelog v1.1.0 — 2026-05-04*
