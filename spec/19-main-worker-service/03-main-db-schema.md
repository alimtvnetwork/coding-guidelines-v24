# 03 — Main Server DB Schema

**Spec:** `19-main-worker-service`
**Version:** 1.1.0
**DB:** SQLite (default). Same schema portable to PostgreSQL/MySQL.

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
| `WorkerNodeRegisteredAt` | TEXT | NO | ISO-8601 |
| `WorkerNodeLastSeenAt` | TEXT | NO | ISO-8601, updated on heartbeat |
| `Description` | TEXT | YES | Per Rule 11 |

Unique: `(WorkerNodeIdentity)`.

### 2.2 `WorkerNodeStatus` (ref) and `WorkerNodeKind` (ref)

Both follow the same shape:

| Column | Type | Null |
|--------|------|------|
| `{TableName}Id` | INTEGER | NO (PK) |
| `{TableName}Code` | TEXT | NO (unique, e.g. `Active`, `Draining`, `Offline`) |
| `{TableName}Label` | TEXT | NO |
| `Description` | TEXT | YES |

Seed values via Seedable-Config. Statuses: `Active`, `Draining`, `Offline`, `Quarantined`. Kinds: `Standard`, `HighMemory`, `Reserved` (extensible).

### 2.3 `Company` (entity, MINIMAL identity only)

> ⚠ Full company data lives in the assigned Worker's split-DB. Main stores only what's needed to route.

| Column | Type | Null | Notes |
|--------|------|------|-------|
| `CompanyId` | INTEGER | NO | PK |
| `CompanySlug` | TEXT | NO | Unique, URL-safe |
| `CompanyName` | TEXT | NO | Display name |
| `WorkerNodeId` | INTEGER | NO | FK → `WorkerNode` |
| `CompanyAssignedAt` | TEXT | NO | ISO-8601 |
| `Description` | TEXT | YES |

Unique: `(CompanySlug)`.

### 2.4 `User` (entity, MINIMAL identity only)

| Column | Type | Null | Notes |
|--------|------|------|-------|
| `UserId` | INTEGER | NO | PK |
| `UserEmail` | TEXT | NO | Unique |
| `UserPasswordHash` | TEXT | NO | Salted, see `05-auth-and-2fa.md` §3 |
| `UserPasswordSalt` | TEXT | NO | |
| `CompanyId` | INTEGER | NO | FK |
| `UserCreatedAt` | TEXT | NO | |
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
| `RoleCode` | TEXT | NO (unique, e.g. `PowerAdmin`, `AdminUser`, `Member`) |
| `RoleLabel` | TEXT | NO |
| `Description` | TEXT | YES |

Seeded via Seedable-Config.

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

*Main DB schema v1.0.0 — 2026-05-04*
