# 03 — Main Server DB Schema

**Spec:** `19-main-worker-service`
**Version:** 2.0.0

> **v2.0.0 (Phase 2 — DB convention overhaul):**
> - All `*At` timestamp columns are now `INTEGER` (epoch seconds, UTC) per `spec/04-database-conventions/01-naming-conventions.md` Rule 7.1 v2.
> - All ref / enum-like tables now use the canonical `(Id, Code, Label)` shape per Rule 13. The legacy `{Table}Code` / `{Table}Label` column names are **removed** in this spec; readers MUST migrate.
> - `Company` columns renamed: `CompanySlug` → `Slug`, `CompanyName` → `Name`.
>
> **v1.4.0 carryover (Phase 1):** `EnumPage` → `AccessItem`, `RolePageAccess` → `RoleAccessItem`. New column `AccessItem.PageUrlSuffix` is the route-matcher source of truth. Deprecation aliases for the old names remain accepted through v1.5.0 per `98-changelog.md`.
>
> **DB:** SQLite (default). Same schema portable to PostgreSQL/MySQL.

> **Split-DB tier authority (FU-2):** Main uses only **3 tiers** — Root, Settings, Session — per [`11-split-db-tier-reconciliation.md`](./11-split-db-tier-reconciliation.md) §4. **Main has no App tier** (it owns no business data). Any prior reference placing Main tables in an App tier is a bug; per the reconciliation file, such tables belong in Root or Settings. Tier assignments per table are listed in `11-…` §4.

---

## 1. Principles

- **Thin catalog only.** Main DB stores routing metadata, not business data.
- **PascalCase** for tables, columns, JSON keys (per `spec/04-database-conventions/`).
- **PKs:** `{TableName}Id INTEGER PRIMARY KEY AUTOINCREMENT`. **No UUIDs.**
- `Type` / `Status` / `Category` / `Kind` columns → join tables, never inline strings.
- Entity/ref tables include `Description TEXT NULL`. Transactional tables include `Notes TEXT NULL` and `Comments TEXT NULL`. All nullable, no DEFAULT (memory rules 10/11/12).
- Join tables exempt from the description/notes rule.

ERD: `diagrams/erd-main-db.mmd`.

---

## 2. Tables

### 2.1 `WorkerNode` (entity)

| Column | Type | Null | Notes |
|--------|------|------|-------|
| `WorkerNodeId` | INTEGER | NO | PK, AUTOINCREMENT |
| `WorkerNodeTitle` | TEXT | NO | Human label, may repeat across nodes |
| `WorkerNodeIdentity` | TEXT | NO | Unique stable identifier (e.g. machine fingerprint) |
| `WorkerNodeEndpoint` | TEXT | NO | Base URL, e.g. `https://w1.example.com` |
| `WorkerNodeStatusId` | INTEGER | NO | FK → `WorkerNodeStatus.WorkerNodeStatusId` |
| `WorkerNodeKindId` | INTEGER | NO | FK → `WorkerNodeKind.WorkerNodeKindId` |
| `WorkerNodeRegisteredAt` | INTEGER | NO | Epoch seconds, UTC |
| `WorkerNodeLastSeenAt` | INTEGER | NO | Epoch seconds, UTC; updated on heartbeat |
| `Description` | TEXT | YES | Per Rule 11 |

Unique: `(WorkerNodeIdentity)`.

### 2.2 `WorkerNodeStatus` (ref) and `WorkerNodeKind` (ref)

Both follow the canonical `(Id, Code, Label)` ref shape (Rule 13):

| Column | Type | Null |
|--------|------|------|
| `{TableName}Id` | INTEGER | NO (PK) |
| `Code` | TEXT | NO (unique, e.g. `Active`, `Draining`, `Offline`) |
| `Label` | TEXT | NO (human-readable) |
| `Description` | TEXT | YES |

Seed values via Seedable-Config. Statuses: `Active`, `Draining`, `Offline`, `Quarantined`. Kinds: `Standard`, `HighMemory`, `Reserved` (extensible).

### 2.3 `Company` (entity, MINIMAL identity only)

> ⚠ Full company data lives in the assigned Worker's split-DB. Main stores only what's needed to route.

| Column | Type | Null | Notes |
|--------|------|------|-------|
| `CompanyId` | INTEGER | NO | PK |
| `Slug` | TEXT | NO | Unique, URL-safe (renamed from `CompanySlug` in v2.0.0) |
| `Name` | TEXT | NO | Display name (renamed from `CompanyName` in v2.0.0) |
| `WorkerNodeId` | INTEGER | NO | FK → `WorkerNode` |
| `CompanyAssignedAt` | INTEGER | NO | Epoch seconds, UTC |
| `Description` | TEXT | YES |

Unique: `(Slug)`. Seedable-Config aliases `CompanySlug → Slug` / `CompanyName → Name` accepted through v2.1.0 then removed.

### 2.4 `User` (entity, MINIMAL identity only)

| Column | Type | Null | Notes |
|--------|------|------|-------|
| `UserId` | INTEGER | NO | PK |
| `UserEmail` | TEXT | NO | Unique |
| `UserPasswordHash` | TEXT | NO | Salted, see `05-auth-and-2fa.md` §3 |
| `UserPasswordSalt` | TEXT | NO | |
| `CompanyId` | INTEGER | NO | FK |
| `UserCreatedAt` | INTEGER | NO | Epoch seconds, UTC |
| `UserTotpSecret` | TEXT | YES | Base32-encoded RFC-6238 shared secret. NULL = TOTP not enrolled. Encrypted-at-rest per `05-auth-and-2fa.md` §4. (Resolves F-A-24.) |
| `UserTotpEnrolledAt` | INTEGER | YES | Epoch seconds, UTC; NULL until first successful TOTP verification. |
| `UserTotpBackupCodesHash` | TEXT | YES | JSON array of bcrypt hashes of 10 single-use backup codes per `05-§4`. NULL until enrollment. |
| `Description` | TEXT | YES |

Unique: `(UserEmail)`.

### 2.5 `UserRole` (join, exempt from Description rule)

| Column | Type | Null |
|--------|------|------|
| `UserRoleId` | INTEGER | NO (PK) |
| `UserId` | INTEGER | NO (FK) |
| `RoleId` | INTEGER | NO (FK → `Role`) |

Unique: `(UserId, RoleId)`.

### 2.6 `Role` (ref)

| Column | Type | Null |
|--------|------|------|
| `RoleId` | INTEGER | NO (PK) |
| `Code` | TEXT | NO (unique, e.g. `PowerAdmin`, `AdminUser`, `Member`) |
| `Label` | TEXT | NO |
| `Description` | TEXT | YES |

Seeded via Seedable-Config.

### 2.6.1 `AccessItem` (ref) — renamed from `EnumPage` in v1.4.0

Defines the closed set of access-controlled dashboard surfaces. Single source of truth for what RBAC governs. Replaces the older name `EnumPage`; old name is **deprecated** but accepted for one release as an alias in seed loaders.

| Column | Type | Null | Notes |
|--------|------|------|-------|
| `AccessItemId` | INTEGER | NO | PK, AUTOINCREMENT |
| `Code` | TEXT | NO | Unique PascalCase identifier (e.g. `WorkerRegistry`, `PushUpdate`, `UserManagement`). Used by code as a stable enum value. |
| `Label` | TEXT | NO | Human-readable UI string (e.g. `Worker Registry`). |
| `PageUrlSuffix` | TEXT | YES | Route matcher — the trailing path fragment used to associate a request with this AccessItem (e.g. `/admin/workers`, `/billing`). NULL for non-route capabilities (e.g. background actions). |
| `Description` | TEXT | YES | Per Rule 11. |

Unique: `(Code)`. Indexed on `(PageUrlSuffix)` for matcher lookups.

> **Matcher rule.** A request path matches an `AccessItem` when its normalized path **ends with** `PageUrlSuffix`. Suffix matching keeps the table portable across deployments mounted under different base paths. Multiple AccessItems sharing the same suffix is a seed error.

Seeded via Seedable-Config (9 rows enumerated in `14-rbac-and-status-seed.md`).

### 2.6.2 `RoleAccessItem` (join, exempt from Description rule) — renamed from `RolePageAccess` in v1.4.0

Per-role access grant for each `AccessItem`. (Resolves F-A-23 / F-B-10.)

| Column | Type | Null |
|--------|------|------|
| `RoleAccessItemId` | INTEGER | NO (PK) |
| `RoleId` | INTEGER | NO (FK → `Role`) |
| `AccessItemId` | INTEGER | NO (FK → `AccessItem`) |
| `CanRead` | INTEGER | NO (0/1) |
| `CanWrite` | INTEGER | NO (0/1) |

Unique: `(RoleId, AccessItemId)`. Seeded with 19 rows per `14-rbac-and-status-seed.md`.

### 2.6.3 `AccessDenialEvent` (transactional, audit)

Audit row written by Workers on every 403 returned for an `AccessDenied` envelope (per `08-error-contract.md` §3.5 and `07-§8`). (Resolves F-A-17.)

| Column | Type | Null |
|--------|------|------|
| `AccessDenialEventId` | INTEGER | NO (PK) |
| `UserId` | INTEGER | NO (FK) |
| `AccessItemId` | INTEGER | NO (FK → `AccessItem`) |
| `WorkerNodeId` | INTEGER | YES (FK; NULL when denied at Main edge) |
| `CorrelationId` | TEXT | NO |
| `OccurredAt` | INTEGER | NO (epoch seconds UTC — see Phase 2; legacy TEXT readers must accept either during the v1.4 → v1.5 transition) |
| `Notes` | TEXT | YES |
| `Comments` | TEXT | YES |

Indexed on `(UserId, OccurredAt)` and `(AccessItemId, OccurredAt)` for audit queries.

### 2.6.4 `EndpointAuthAuditEvent` (transactional, audit) — FU-17

Audit row written on every successful `PATCH /API/V1/Settings/EndpointAuth` (per `06-core-api-endpoints.md` §5.6). Sibling shape to `AccessDenialEvent` (§2.6.3). One row per accepted PATCH; idempotent replays (same `X-Idempotency-Key`, same body within `MainWorker.Idempotency.KeyTtlSeconds`) MUST NOT emit a duplicate row. Resolves audit follow-up FU-17.

| Column | Type | Null | Notes |
|--------|------|------|-------|
| `EndpointAuthAuditEventId` | INTEGER | NO | PK, AUTOINCREMENT |
| `EndpointAuthSettingId` | INTEGER | NO | FK → `EndpointAuthSetting.EndpointAuthSettingId` (the row that was written; created or replaced) |
| `EndpointPathPattern` | TEXT | NO | Snapshotted at write time so audit survives later row deletion |
| `HttpMethodMaskOld` | TEXT | YES | NULL when the row was newly created (no prior state) |
| `HttpMethodMaskNew` | TEXT | NO | Post-PATCH value |
| `IsEnabledOld` | INTEGER | YES | NULL on create; `0`/`1` on replace |
| `IsEnabledNew` | INTEGER | NO | Post-PATCH value |
| `OldMechanismsJson` | TEXT | YES | JSON array of prior `AuthMechanismCode[]`; NULL on create |
| `NewMechanismsJson` | TEXT | NO | JSON array of post-PATCH `AuthMechanismCode[]` (sorted ascending for diffability) |
| `ChangeKindId` | INTEGER | NO | FK → `EndpointAuthChangeKind.EndpointAuthChangeKindId` (`Create`, `Replace`, `SoftDisable`, `Reenable`) |
| `UpdatedByUserId` | INTEGER | NO | FK → `User`. Same actor stamped on the parent row's `UpdatedByUserId`. |
| `CorrelationId` | TEXT | NO | Echoes the inbound `X-Correlation-Id` header per `spec/04-database-conventions/06-rest-api-format.md` |
| `IdempotencyKey` | TEXT | NO | The `X-Idempotency-Key` that produced the write. Index supports replay-detection joins. |
| `OccurredAt` | TEXT | NO | ISO-8601 UTC; server-stamped, equals the parent row's `UpdatedAt` |
| `Notes` | TEXT | YES | Per Rule 12 |
| `Comments` | TEXT | YES | Per Rule 12 |

Unique: `(IdempotencyKey)` — guarantees the no-duplicate-on-replay invariant above. Indexed on `(EndpointAuthSettingId, OccurredAt DESC)` and `(UpdatedByUserId, OccurredAt DESC)` for audit queries.

### 2.6.5 `EndpointAuthChangeKind` (ref)

| Column | Type | Null |
|--------|------|------|
| `EndpointAuthChangeKindId` | INTEGER | NO (PK) |
| `EndpointAuthChangeKindCode` | TEXT | NO (unique: `Create`, `Replace`, `SoftDisable`, `Reenable`) |
| `EndpointAuthChangeKindLabel` | TEXT | NO |
| `Description` | TEXT | YES |

Seeded via Seedable-Config — 4 rows. Resolution rules:

- `Create` — no prior `EndpointAuthSetting` row for the pattern existed.
- `Replace` — prior row existed AND (`HttpMethodMask` changed OR `AcceptedMechanisms` set changed) AND `IsEnabled` did not transition.
- `SoftDisable` — `IsEnabled` transitioned `1 → 0`.
- `Reenable` — `IsEnabled` transitioned `0 → 1`.

### 2.7 `WorkerVersion` (transactional)

Tracks which version each Worker is currently running.

| Column | Type | Null | Notes |
|--------|------|------|-------|
| `WorkerVersionId` | INTEGER | NO | PK |
| `WorkerNodeId` | INTEGER | NO | FK |
| `WorkerVersionSemver` | TEXT | NO | e.g. `1.4.2` |
| `WorkerVersionRecordedAt` | TEXT | NO | ISO-8601 |
| `Notes` | TEXT | YES | Per Rule 12 |
| `Comments` | TEXT | YES | Per Rule 12 |

### 2.8 `WorkerSelectionEvent` (transactional, audit)

Records every routing decision. Useful for debugging load distribution.

| Column | Type | Null |
|--------|------|------|
| `WorkerSelectionEventId` | INTEGER | NO (PK) |
| `CompanyId` | INTEGER | NO (FK) |
| `WorkerNodeId` | INTEGER | NO (FK) |
| `WorkerSelectionStrategyId` | INTEGER | NO (FK → `WorkerSelectionStrategy`) |
| `WorkerSelectionEventAt` | TEXT | NO |
| `Notes` | TEXT | YES |
| `Comments` | TEXT | YES |

### 2.9 `WorkerSelectionStrategy` (ref)

| Column | Type | Null |
|--------|------|------|
| `WorkerSelectionStrategyId` | INTEGER | NO (PK) |
| `WorkerSelectionStrategyCode` | TEXT | NO (`RoundRobin`, `LeastLoaded`, `Manual`) |
| `WorkerSelectionStrategyLabel` | TEXT | NO |
| `Description` | TEXT | YES |

---

## 3. Indexes

| Index | Columns |
|-------|---------|
| `IX_Company_WorkerNodeId` | `Company(WorkerNodeId)` |
| `IX_User_CompanyId` | `User(CompanyId)` |
| `IX_WorkerVersion_WorkerNodeId_RecordedAt` | `WorkerVersion(WorkerNodeId, WorkerVersionRecordedAt DESC)` |
| `IX_WorkerSelectionEvent_At` | `WorkerSelectionEvent(WorkerSelectionEventAt DESC)` |
| `UX_EndpointAuthAuditEvent_IdempotencyKey` | `EndpointAuthAuditEvent(IdempotencyKey)` UNIQUE |
| `IX_EndpointAuthAuditEvent_Setting_At` | `EndpointAuthAuditEvent(EndpointAuthSettingId, OccurredAt DESC)` |
| `IX_EndpointAuthAuditEvent_Actor_At` | `EndpointAuthAuditEvent(UpdatedByUserId, OccurredAt DESC)` |

---

## 4. What Main DB does NOT store

- Company business fields (address, social media, employee count, etc.)
- User profile data beyond auth
- Any per-tenant business state
- Session bodies (kept in cache or session store, not the catalog)

All of the above belong in the Worker's split-DB per `spec/05-split-db-architecture/`.

---

## 5. Migration Notes

- Use the implementer's standard migration tool (Laravel migrations for the default stack).
- Migrations are idempotent and forward-only.
- Seed data for `Role`, `WorkerNodeStatus`, `WorkerNodeKind`, `WorkerSelectionStrategy` ships via Seedable-Config (`spec/06-seedable-config-architecture/`).

---

*Main DB schema v1.4.0 — 2026-05-06 (Phase 1: AccessItem rename)*
