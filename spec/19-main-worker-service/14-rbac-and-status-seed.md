# 14 — RBAC + Status Seed (Concrete Row Sets)

**Spec:** `19-main-worker-service`
**Version:** 2.0.0

> **v2.0.0 rename (Phase 1):** `EnumPage` → `AccessItem`, `RolePageAccess` → `RoleAccessItem`. New column `AccessItem.PageUrlSuffix` (route matcher). Old table names accepted as aliases for one release per `03-main-db-schema.md` §2.6.1 deprecation notice.
**Created:** 2026-05-04
**Status:** Authoritative
**Resolves:** audit findings F-B-09, F-B-10, F-X-06 (top-10 fix #6). Closes AC-5.
**Mechanism:** `spec/06-seedable-config-architecture/02-features/07-reference-table-seeding.md` (`Tables` block in `config.seed.json`).
**Authority:** This file is the canonical row set for Main-tier RBAC and worker-status reference tables. Implementations MUST copy the JSON in §3 verbatim into `config.seed.json`.

---

## 1. What this seeds

| Table | Tier | Rows | Purpose |
|---|---|---:|---|
| `Role` | Settings | 3 | Top-level roles |
| `EnumPage` | Settings | 9 | Page-capability catalog |
| `RolePageAccess` | Settings | 19 | Role↔Page grants |
| `WorkerNodeStatus` | Root | 4 | Worker lifecycle states |
| `AuthMechanism` | Settings | 4 | Endpoint auth toggles |

All five seeds ship together at SemVer `1.3.0` of `config.seed.json` (FU-12 will bump the seed file).

---

## 2. Schema recap (PascalCase, Code Red Schema Rules 10/11/12)

```sql
CREATE TABLE Role (
    RoleId       INTEGER PRIMARY KEY AUTOINCREMENT,
    RoleCode     TEXT NOT NULL UNIQUE,
    RoleLabel    TEXT NOT NULL,
    Description  TEXT NULL
);

CREATE TABLE EnumPage (
    EnumPageId    INTEGER PRIMARY KEY AUTOINCREMENT,
    EnumPageCode  TEXT NOT NULL UNIQUE,
    EnumPageLabel TEXT NOT NULL,
    Description   TEXT NULL
);

CREATE TABLE RolePageAccess (   -- join table, exempt from Description per Rule 12 carve-out
    RolePageAccessId INTEGER PRIMARY KEY AUTOINCREMENT,
    RoleId           INTEGER NOT NULL REFERENCES Role(RoleId),
    EnumPageCode     TEXT NOT NULL REFERENCES EnumPage(EnumPageCode),
    UNIQUE (RoleId, EnumPageCode)
);

CREATE TABLE WorkerNodeStatus (
    WorkerNodeStatusId    INTEGER PRIMARY KEY AUTOINCREMENT,
    WorkerNodeStatusCode  TEXT NOT NULL UNIQUE,
    WorkerNodeStatusLabel TEXT NOT NULL,
    Description           TEXT NULL
);

CREATE TABLE AuthMechanism (
    AuthMechanismId    INTEGER PRIMARY KEY AUTOINCREMENT,
    AuthMechanismCode  TEXT NOT NULL UNIQUE,
    AuthMechanismLabel TEXT NOT NULL,
    Description        TEXT NULL
);
```

> `RolePageAccess` is a join table per `09-templates/01-error-codes-template.md` style — the Description column is intentionally absent (Code Red Rule 12 join-table carve-out documented in memory).

---

## 3. `Tables` block (paste into `config.seed.json`)

```jsonc
"Tables": {

  "Role": {
    "AddedIn": "1.3.0",
    "Version": "1.3.0",
    "PrimaryKey": ["RoleCode"],
    "MergeStrategy": "UpsertByLogicalKey",
    "Description": "Top-level roles for Main-tier RBAC.",
    "Rows": [
      { "RoleCode": "PowerAdmin", "RoleLabel": "Power Administrator", "Description": "Full cross-tenant control. Issued only to platform operators." },
      { "RoleCode": "AdminUser",  "RoleLabel": "Company Administrator", "Description": "Manages a single company: users, billing, settings." },
      { "RoleCode": "Member",     "RoleLabel": "Member",                "Description": "Standard end-user. Read-only outside personal scope." }
    ]
  },

  "EnumPage": {
    "AddedIn": "1.3.0",
    "Version": "1.3.0",
    "PrimaryKey": ["EnumPageCode"],
    "MergeStrategy": "UpsertByLogicalKey",
    "Description": "Capability catalog enforced by RolePageAccess.",
    "Rows": [
      { "EnumPageCode": "PowerAdminPage",      "EnumPageLabel": "Power Admin",       "Description": "Cross-tenant ops console." },
      { "EnumPageCode": "AdminPage",           "EnumPageLabel": "Admin",             "Description": "Per-company admin home." },
      { "EnumPageCode": "BillingPage",         "EnumPageLabel": "Billing",           "Description": "Invoices, plan, payment methods." },
      { "EnumPageCode": "CompanySettingsPage", "EnumPageLabel": "Company Settings",  "Description": "Company profile + branding." },
      { "EnumPageCode": "UserManagementPage",  "EnumPageLabel": "User Management",   "Description": "Add/remove users, assign roles." },
      { "EnumPageCode": "WorkerRegistryPage",  "EnumPageLabel": "Worker Registry",   "Description": "List + status of worker nodes." },
      { "EnumPageCode": "PushUpdatePage",      "EnumPageLabel": "Push Update",       "Description": "Trigger worker self-updates." },
      { "EnumPageCode": "AuditLogPage",        "EnumPageLabel": "Audit Log",         "Description": "Read-only audit history." },
      { "EnumPageCode": "DashboardPage",       "EnumPageLabel": "Dashboard",         "Description": "Default landing page." }
    ]
  },

  "RolePageAccess": {
    "AddedIn": "1.3.0",
    "Version": "1.3.0",
    "PrimaryKey": ["RoleId", "EnumPageCode"],
    "MergeStrategy": "UpsertByLogicalKey",
    "Description": "Default grants. Operators may add/remove via UI; seed only enforces base baseline.",
    "Rows": [
      { "RoleId": "@Role.PowerAdmin", "EnumPageCode": "PowerAdminPage" },
      { "RoleId": "@Role.PowerAdmin", "EnumPageCode": "AdminPage" },
      { "RoleId": "@Role.PowerAdmin", "EnumPageCode": "BillingPage" },
      { "RoleId": "@Role.PowerAdmin", "EnumPageCode": "CompanySettingsPage" },
      { "RoleId": "@Role.PowerAdmin", "EnumPageCode": "UserManagementPage" },
      { "RoleId": "@Role.PowerAdmin", "EnumPageCode": "WorkerRegistryPage" },
      { "RoleId": "@Role.PowerAdmin", "EnumPageCode": "PushUpdatePage" },
      { "RoleId": "@Role.PowerAdmin", "EnumPageCode": "AuditLogPage" },
      { "RoleId": "@Role.PowerAdmin", "EnumPageCode": "DashboardPage" },

      { "RoleId": "@Role.AdminUser",  "EnumPageCode": "AdminPage" },
      { "RoleId": "@Role.AdminUser",  "EnumPageCode": "BillingPage" },
      { "RoleId": "@Role.AdminUser",  "EnumPageCode": "CompanySettingsPage" },
      { "RoleId": "@Role.AdminUser",  "EnumPageCode": "UserManagementPage" },
      { "RoleId": "@Role.AdminUser",  "EnumPageCode": "AuditLogPage" },
      { "RoleId": "@Role.AdminUser",  "EnumPageCode": "DashboardPage" },

      { "RoleId": "@Role.Member",     "EnumPageCode": "DashboardPage" }
    ]
  },

  "WorkerNodeStatus": {
    "AddedIn": "1.3.0",
    "Version": "1.3.0",
    "PrimaryKey": ["WorkerNodeStatusCode"],
    "MergeStrategy": "UpsertByLogicalKey",
    "Description": "Worker lifecycle states observed by Main.",
    "Rows": [
      { "WorkerNodeStatusCode": "Registering", "WorkerNodeStatusLabel": "Registering", "Description": "Initial state — handshake in progress." },
      { "WorkerNodeStatusCode": "Active",      "WorkerNodeStatusLabel": "Active",      "Description": "Heartbeating; eligible for routing." },
      { "WorkerNodeStatusCode": "Quarantined", "WorkerNodeStatusLabel": "Quarantined", "Description": "Missed >=3 heartbeats; new tenants not routed here." },
      { "WorkerNodeStatusCode": "Retired",     "WorkerNodeStatusLabel": "Retired",     "Description": "Permanently removed; preserved for audit." }
    ]
  },

  "AuthMechanism": {
    "AddedIn": "1.3.0",
    "Version": "1.3.0",
    "PrimaryKey": ["AuthMechanismCode"],
    "MergeStrategy": "UpsertByLogicalKey",
    "Description": "Per-endpoint auth toggle catalog (used by EndpointAuthSetting).",
    "Rows": [
      { "AuthMechanismCode": "Session", "AuthMechanismLabel": "Main session cookie",         "Description": "HTTPOnly cookie set by Main." },
      { "AuthMechanismCode": "Jwt",     "AuthMechanismLabel": "Worker JWT (RS256)",          "Description": "Per spec/19/12-jwt-delivery-contract.md." },
      { "AuthMechanismCode": "OAuth",   "AuthMechanismLabel": "OAuth client-credentials",    "Description": "Worker↔Main service auth." },
      { "AuthMechanismCode": "None",    "AuthMechanismLabel": "Public",                      "Description": "No auth required — explicitly opt-in." }
    ]
  }

}
```

### 3.1 Logical-key references (`@Role.PowerAdmin`)

The `@<Table>.<Code>` syntax tells the seeder to substitute the AUTOINCREMENT id of the row whose unique-code column equals `Code`. Resolved at apply time. If the referenced row is missing, fail with `MWS-21002 SplitDbTierMissing` (closest fit) and abort the merge for that table.

Implementer note: this is a seeder feature (FU-13 — extend `01-fundamentals.md` `mergeSeed` to honor `@`-references). Until shipped, use a two-pass approach: seed `Role` first, then `RolePageAccess` with literal IDs derived from a SELECT.

---

## 4. Verification (post-boot self-check)

After the seeder runs, the worker MUST verify:

```sql
SELECT COUNT(*) FROM Role;             -- expect ≥ 3
SELECT COUNT(*) FROM EnumPage;         -- expect ≥ 9
SELECT COUNT(*) FROM RolePageAccess
  WHERE RoleId = (SELECT RoleId FROM Role WHERE RoleCode='PowerAdmin');
-- expect ≥ 9 (PowerAdmin has all pages)
SELECT COUNT(*) FROM WorkerNodeStatus; -- expect ≥ 4
SELECT COUNT(*) FROM AuthMechanism;    -- expect ≥ 4
```

Any count below expected → exit `MWS-21002 SplitDbTierMissing`.

---

## 5. Re-seeding rules (when authoring `1.4.0`)

- **Adding a new page:** add to `EnumPage.Rows`, add corresponding `RolePageAccess` rows, bump both blocks' `Version` to `1.4.0`. Keep `AddedIn` as the original.
- **Revoking a default grant:** Per spec/06 §SeedWithVersionCheck, `UpsertByLogicalKey` does NOT delete. Operators clear via UI; seed cannot rescind. (Documented as expected behavior — defaults are floors, not ceilings.)
- **Renaming a page:** add new code, keep the old (deprecate via `Description` text), migrate in code over one release.

---

## 6. Cross-references

- `spec/06-seedable-config-architecture/02-features/07-reference-table-seeding.md` — the `Tables` block mechanism.
- `spec/19-main-worker-service/07-role-based-dashboards.md` — consumer of `Role`/`EnumPage`/`RolePageAccess`.
- `spec/19-main-worker-service/10-worker-bootstrap-protocol.md` §8 — consumer of `WorkerNodeStatus`.
- `spec/19-main-worker-service/06-core-api-endpoints.md` §5 — consumer of `AuthMechanism`.
- `spec/19-main-worker-service/11-split-db-tier-reconciliation.md` §4 — tier placement.

---

*RBAC + status seed v1.0.0 — 2026-05-04*
