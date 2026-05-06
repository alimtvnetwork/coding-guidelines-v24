# 11 — Split-DB Tier Reconciliation (Main + Worker)

**Spec:** `19-main-worker-service`
**Version:** 1.0.0
**Created:** 2026-05-04
**Status:** Authoritative
**Resolves:** audit findings F-X-01, F-X-04, F-D-09 (top-10 fix #2). Unblocks AC-2.
**Authority:** This file is the canonical mapping between Main/Worker spec and `spec/05-split-db-architecture/`. On any tier-count or tier-name conflict, **spec/05 wins** and this file translates.

---

## 1. Why this file exists

Earlier drafts of `spec/19-main-worker-service/` referred to a "3-tier split-DB" (Root / App / Session). `spec/05-split-db-architecture/01-fundamentals.md` actually defines **6 tiers** (Root, Settings, App, Session, Cache, Document). The audit (F-X-01) flagged this as a BLOCKER for AC-2.

This file pins the correct mapping for Main and Worker. Cache and Document tiers are intentionally **not used** by either tier in v1.0.

---

## 2. Authoritative tier table

| Tier (spec/05) | Used by Main? | Used by Worker? | Physical location | Purpose in Main/Worker context |
|---|---|---|---|---|
| **Root** | ✅ Yes | ✅ Yes | `data/root.db` | Global registry. Main: tenant→worker map. Worker: company shard registry. |
| **Settings** | ✅ Yes | ✅ Yes | Inside `root.db` (per spec/05 §Tier table) | Seedable + user config. Worker bootstrap config (per `10-worker-bootstrap-protocol.md` §2). |
| **App** | ❌ No | ✅ Yes | `data/{CompanySlug}/app.db` | Per-company business data. Main has no App tier — it owns no business data. |
| **Session** | ✅ Yes | ✅ Yes | Main: `data/sessions/{SessionId}.db`. Worker: `data/{CompanySlug}/sessions/{SessionId}.db` | Per-login session DB. Main owns auth sessions; Worker owns app sessions. |
| **Cache** | ❌ No | ❌ No | — | Reserved by spec/05 for RAG; not used in v1.0. |
| **Document** | ❌ No | ❌ No | — | Reserved by spec/05 for RAG; not used in v1.0. |

> **Tier count for Main:** 3 used (Root, Settings, Session).
> **Tier count for Worker:** 4 used (Root, Settings, App, Session).
> **The "3-vs-4 tier" debate is resolved here:** Main = 3, Worker = 4. Both are subsets of spec/05's 6.

---

## 3. Spec/19 prose corrections (apply during next edit pass)

The following files contain stale "3-tier" or missing-Settings language. They MUST be updated to defer to this file:

| File | Stale phrase | Replace with |
|---|---|---|
| `01-architecture.md` | "3-tier split-DB (Root/App/Session)" | "Per `11-split-db-tier-reconciliation.md` — Main uses 3 tiers (Root/Settings/Session); Worker uses 4 tiers (Root/Settings/App/Session)." |
| `03-main-db-schema.md` | references to "App tier" on Main | Move to Root tier; Main has no App tier. |
| `09-self-update-pointer.md` | (silent on Settings tier) | Add note: Worker bootstrap config lives in Settings tier per `10-worker-bootstrap-protocol.md` §2. |
| `diagrams/erd-worker-split-db.mmd` | shows 3 tiers as boxes | Add Settings tier box; add banner "non-authoritative projection of spec/05 §Tier table." |

These edits are tracked as follow-up tasks (see §8) — this doc establishes the contract; the rewrites happen incrementally.

---

## 4. Per-tier table allocation (Main)

> **v2.1.0 (Phase 3) update.** `User` and `UserRole` are **removed from Main**. Routing is served by the new `UserDirectory` table (Root tier). Authoritative `AppUser` / `AppUserRole` rows now live on the Worker (§5).

| Table | Tier | Source spec |
|---|---|---|
| `Company` | Root | `19/03-main-db-schema.md` §2.3 |
| `WorkerNode` | Root | `19/10-worker-bootstrap-protocol.md` §8 |
| `WorkerNodeStatus` | Root | `19/10-worker-bootstrap-protocol.md` §8 |
| `UserDirectory` (routing index, no secrets) | Root | `19/03-main-db-schema.md` §2.4 (v2.1.0) |
| `Role` (catalog) | Settings | `19/03-main-db-schema.md` §2.6 |
| `AccessItem` (catalog) | Settings | `19/07-role-based-dashboards.md` (seeded via spec/06) |
| `RoleAccessItem` (catalog) | Settings | `19/07-role-based-dashboards.md` |
| `EndpointAuthSetting` | Settings | `19/06-core-api-endpoints.md` §5 |
| `AuthMechanism` | Settings | `19/06-core-api-endpoints.md` §5 |
| `UpdateSchedule` | Settings | `19/06-core-api-endpoints.md` §4 |
| `AuthSession` | Session | `19/05-auth-and-2fa.md` |
| `TwoFactorChallenge` | Session | `19/05-auth-and-2fa.md` (relayed; ephemeral) |
| ~~`User` (auth identity)~~ | ~~Root~~ | **REMOVED v2.1.0 — moved to Worker as `AppUser`.** |
| ~~`UserRole`~~ | ~~Root~~ | **REMOVED v2.1.0 — moved to Worker as `AppUserRole`.** |

---

## 5. Per-tier table allocation (Worker)

> **v2.1.0 (Phase 3) update.** Worker is now the authoritative identity store. `AppUser` carries `PasswordHash`, `PasswordSalt`, `TotpSecret`, `TotpEnrolledAt`, `TotpBackupCodesHash`. `AppUserRole` carries the user-to-role assignments that Main used to hold.

| Table | Tier | Source spec |
|---|---|---|
| `RootCompany` | Root | `19/diagrams/erd-worker-split-db.mmd` |
| `RootCompanyStatus` | Root | same |
| `RootCompanyContact` | Root | same |
| `AppCompanyShard` | Root | (registry of App-tier DBs) |
| `WorkerBootstrapState` | Settings | `19/10-worker-bootstrap-protocol.md` §9 |
| `WorkerUpdateInstruction` | Settings | `spec/14-update/28-worker-push-instruction.md` §7 |
| `AppUser` (authoritative identity, v2.1.0) | App | `19/05-auth-and-2fa.md` §3, `19/diagrams/erd-worker-split-db.mmd` |
| `AppUserRole` (user→role assignments, v2.1.0) | App | `19/14-rbac-and-status-seed.md` §6 |
| `AppBusinessEntity` | App | same |
| `AppSession` | Session | same |

> Note: `WorkerUpdateInstruction` was originally placed in App tier in spec/14-update/28 §7. **Correction:** it belongs in Settings tier (worker-wide, not company-scoped). Update spec/14-update/28 §7 in a follow-up task.

---

## 6. Provisioning order on worker boot

Per `10-worker-bootstrap-protocol.md` §3 step 2 ("Self-test split-DB tiers"), the Worker MUST verify in this exact order:

1. Root tier — open `data/root.db`, verify schema version.
2. Settings tier — verify `WorkerBootstrapState` table exists in Root DB.
3. App tier — for each row in `AppCompanyShard`, verify `data/{CompanySlug}/app.db` exists and migrates.
4. Session tier — directory `data/{CompanySlug}/sessions/` exists and is writable.

Failure at any step → exit `WORKER-000-02 SPLIT_DB_TIER_MISSING` per `10/§6`.

---

## 7. Diagram banner template

Add to every Mermaid file under `spec/19/diagrams/` that touches DB tiers:

```
%% Authority: spec/05-split-db-architecture/ defines tier semantics.
%% Tier mapping for this spec: spec/19-main-worker-service/11-split-db-tier-reconciliation.md
%% This diagram is a non-authoritative projection. On conflict, prose wins.
```

---

## 8. Follow-up tasks generated by this doc

| # | Task | File to touch |
|---|---|---|
| FU-1 | Update "3-tier" prose | `19/01-architecture.md` |
| FU-2 | Move Main `App`-tier refs to Root | `19/03-main-db-schema.md` |
| FU-3 | Add Settings tier note | `19/09-self-update-pointer.md` |
| FU-4 | Add Settings box + banner | `19/diagrams/erd-worker-split-db.mmd` |
| FU-5 | Move `WorkerUpdateInstruction` to Settings tier | `spec/14-update/28-worker-push-instruction.md` §7 |
| FU-6 | Add tier-mapping section reference | `spec/05-split-db-architecture/01-fundamentals.md` (back-link) |

These are catalogued in `audit/05-implementation-pivot-score.md` Top-10 — Fix #9 covers FU-4; FU-1/2/3/5/6 are new and tracked here.

---

## 9. Open Questions (logged, non-blocking)

- **OQ-11-1** Should Main also gain an App tier in v2.0 to host cross-tenant analytics? Inferred: No — Main remains business-data-free per AC-1.
- **OQ-11-2** Cache + Document tiers: gate behind a feature flag for future RAG support? Inferred: defer; spec/05 already gates them.

---

## 10. Cross-references

- `spec/05-split-db-architecture/01-fundamentals.md` — source of truth for all tier semantics.
- `spec/19-main-worker-service/10-worker-bootstrap-protocol.md` — uses §6 provisioning order.
- `spec/19-main-worker-service/03-main-db-schema.md` — affected by FU-2.
- `spec/14-update/28-worker-push-instruction.md` — affected by FU-5.
- `spec/06-seedable-config-architecture/` — seeds Settings-tier rows.

---

*Split-DB tier reconciliation v1.0.0 — 2026-05-04*
