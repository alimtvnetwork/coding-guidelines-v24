# Current Plan

**Version:** 5.12.0 → planned 5.13.0 (after Phase 12)
**Updated:** 2026-05-06

---

## Active Initiative — Backup Nodes, Cascading Roles, DB Convention Overhaul

**Mode:** SPEC ONLY. No implementation code. Per user verbatim §Important.5, "No implementation right now. Job is to write the spec only."

**Scope folders:**
- `spec/19-main-worker-service/` — primary
- `spec/04-database-conventions/` — DateTime INTEGER, enum tables Id/Code/Label
- `spec/05-split-db-architecture/` — DateTime INTEGER propagation
- `spec/19-main-worker-service/diagrams/` — ER diagram updates

**Stack-agnostic.** Default examples in Laravel/PHP per existing spec.

---

## Locked Decisions (from user verbatim)

| # | Decision |
|---|----------|
| D1 | `EnumPage` is renamed to **`AccessItem`**. Columns include `Label` and `PageUrlSuffix` (matcher). |
| D2 | DateTime columns are **INTEGER (epoch seconds)**, never TEXT. Applies to ALL specs. |
| D3 | Enum-like ref tables use exactly **`Id`, `Code`, `Label`** (drop `{TableName}Code` prefix style for these). |
| D4 | `Company` table uses **`Slug` and `Name`** only. Drop `CompanyName`. |
| D5 | **Main node holds NO Users.** Main holds Company + Company→Worker mapping only. Users live on Worker. |
| D6 | Worker nodes have a **`Sequence` INTEGER** column. |
| D7 | Worker UI label is **"Region"**. Code stays "Worker". |
| D8 | Worker nodes have **`IsBackup` BOOLEAN** and **`BackupOfWorkerNodeId` FK NULL**. |
| D9 | Backup nodes **never serve traffic**. Receive + store only. |
| D10 | Synced rows carry a **`SyncOp` flag** (`Insert` / `Update` / `Delete`) for diff application. |
| D11 | Cascading roles = **union of AccessItems across all assigned roles**. User can hold multiple roles simultaneously. |
| D12 | Add **cache-bin** tables/notes for role management in the ER diagram. |
| D13 | Encryption: shared internal RSA key pair between worker and its backup; main can issue **rotation instruction**. Zip password follows known pattern. |
| D14 | Date-by-date full snapshot storage on backup node. Restore by date is main-controlled. |

---

## Open Questions (carried forward)

- **OQ-A1:** Cascading beyond simple union — propose role inheritance hierarchy (parent role implies child accesses)? Default = simple union.
- **OQ-A2:** Cache-bin technology — in-memory (process-local) vs SQLite memory DB vs Redis? Default proposal = SQLite `:memory:` per-process with TTL invalidation broadcast via Main.
- **OQ-A3:** Zip password "known pattern" — propose `HMAC-SHA256(SharedSecret, BackupTimestampEpoch)` truncated to 32 hex chars. Confirm.
- **OQ-A4:** Snapshot retention policy (days). Default proposal = 30 days rolling.

---

## Phased Plan (12 phases, one per `next`)

### Phase 1 — Rename EnumPage → AccessItem
- Rename across `07-role-based-dashboards.md`, `03-main-db-schema.md`, `14-rbac-and-status-seed.md`, ERD.
- Define `AccessItem` columns: `AccessItemId`, `Code`, `Label`, `PageUrlSuffix`, `Description`.
- Define matcher logic (suffix match against route).
- Update all references; add migration note in `98-changelog.md`.

### Phase 2 — Global DB Convention Updates
- `spec/04-database-conventions/01-naming-conventions.md` — DateTime = INTEGER (epoch seconds, UTC).
- `spec/04-database-conventions/02-schema-design.md` — enum tables shape `Id/Code/Label`.
- `spec/05-split-db-architecture/` — propagate INTEGER DateTime convention.
- `spec/19-main-worker-service/03-main-db-schema.md` — flip all `*At TEXT` → `*At INTEGER`.
- Update `Company` to `(CompanyId, Slug, Name, ...)`.

### Phase 3 — Move Users off Main
- `03-main-db-schema.md` — remove `User`, `UserRole`, TOTP columns from Main.
- Document Users now live on Worker split-DB.
- Update `05-auth-and-2fa.md` — auth lookup flow becomes: Main resolves Company→Worker, Worker authenticates User.
- Update `11-split-db-tier-reconciliation.md` — move User/UserRole to Worker App tier.

### Phase 4 — Worker Node Field Additions
- Add `Sequence INTEGER NOT NULL` to `WorkerNode`.
- Add `IsBackup INTEGER NOT NULL DEFAULT 0` (boolean).
- Add `BackupOfWorkerNodeId INTEGER NULL` FK self-ref.
- Add UI mapping note: "Worker" → "Region" (frontend label).
- Update `04-worker-routing.md` — backup nodes excluded from selection pool.

### Phase 5 — Role / AccessItem N-M + Cache Bin + Cascading
- Document `UserRole` (N-M), `RoleAccessItem` (N-M).
- Define cache-bin tables: `RoleAccessCache` (per-role compiled access set, TTL).
- Cascading rules section: union semantics, multi-role behavior, examples (Admin alone, Editor alone, Admin+Editor).
- Cache invalidation protocol from Main.

### Phase 6 — Backup Node Concept Spec
- New file: `17-backup-nodes.md`.
- Sections: relationship model, registration flow, propagation Main→Worker→Worker-DB, "no serving traffic" invariant.
- ER additions to `WorkerNode` table (already in Phase 4).
- Worker-side mirror table `KnownBackupNode` (in Worker DB).

### Phase 7 — Incremental Diff Generator
- New file: `18-incremental-backup-sync.md`.
- Sections: cron schedule, last-sync watermark table, per-table walk, change-data-capture using `SyncOp` flag column on each app row.
- Watermark table schema in Worker DB.

### Phase 8 — Encryption + Zip Pipeline
- New file: `19-backup-encryption-and-keys.md`.
- RSA key pair shared between worker and its backups.
- Zip password derivation pattern.
- Key rotation instruction issued from Main: endpoint contract + state machine (Pending / Active / Retired).

### Phase 9 — Backup Endpoints Contract
- New file: `20-backup-endpoints.md`.
- Endpoints:
  1. `POST /API/V1/Backup/IncrementalDiff` (on backup node)
  2. `POST /API/V1/Backup/RotateKeys` (on backup node, Main-triggered)
  3. `POST /API/V1/Backup/RestoreByDate` (on backup node, Main-triggered)
  4. `GET  /API/V1/Backup/Snapshots` (list available dates)
  5. `GET  /API/V1/Backup/Health` (status)
- Auth: S2S OAuth (per `05-auth-and-2fa.md`).
- Add error codes in `13-error-codes.md`.

### Phase 10 — Backup Apply Logic
- Append to `20-backup-endpoints.md` or new `21-backup-apply-logic.md`.
- Decrypt → unzip → open SQLite diff → iterate rows → apply by `SyncOp` flag → idempotency via `(SourceTable, SourceRowId, SyncOpSeq)`.

### Phase 11 — Snapshot Storage + Restore Flow
- New file: `22-snapshot-storage-and-restore.md`.
- Date-named full DB zips on backup node filesystem.
- Retention policy (default 30 days).
- Restore flow: Main → Backup `RestoreByDate` → Backup decompresses → Worker pulls or Backup pushes to Worker.

### Phase 12 — ER Diagram + Acceptance Criteria + Sync
- Update `diagrams/erd-main-db.mmd`: AccessItem rename, Users removed, WorkerNode new fields, cache-bin notation, INTEGER DateTime.
- New diagram: `diagrams/erd-worker-db.mmd` (Users + KnownBackupNode + watermark).
- New diagram: `diagrams/seq-backup-incremental.mmd`.
- New diagram: `diagrams/seq-backup-restore.mmd`.
- Update `97-acceptance-criteria.md` with one AC per phase.
- Bump `package.json` 5.12.0 → 5.13.0.
- Run `npm run sync`.

---

## Status

- [x] **Phase 0** — Plan written (this commit). Awaiting `next`.
- [ ] Phases 1–12 — pending.

---

## Cross-spec impact summary

| Spec | Impacted by Phases |
|------|-------------------|
| `04-database-conventions/` | 2 |
| `05-split-db-architecture/` | 2, 3 |
| `19-main-worker-service/03-main-db-schema.md` | 1, 2, 3, 4, 5 |
| `19-main-worker-service/05-auth-and-2fa.md` | 3 |
| `19-main-worker-service/07-role-based-dashboards.md` | 1, 5 |
| `19-main-worker-service/11-split-db-tier-reconciliation.md` | 3 |
| `19-main-worker-service/13-error-codes.md` | 9 |
| `19-main-worker-service/14-rbac-and-status-seed.md` | 1, 5 |
| `19-main-worker-service/diagrams/` | 12 |
| New files: `17-…` through `22-…` | 6–11 |

---

*Plan v2.0.0 — 2026-05-06. Type `next` to begin Phase 1.*
