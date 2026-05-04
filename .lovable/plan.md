# Current Plan

**Version:** 5.10.0
**Updated:** 2026-05-04

---

## v5.10.0 — Main↔Worker Service Architecture Spec (2026-05-04)

**Scope:** New spec folder `spec/19-main-worker-service/` documenting the Main orchestrator + Worker node architecture. Spec/docs only — no runtime code touched.

### Done — Phase 1 (Foundations)
- `00-overview.md`, `01-architecture.md`, `02-glossary.md`, `plan.md`.

### Done — Phase 2 (Core Specs)
- `03-main-db-schema.md` — thin catalog (WorkerNode, Company, User, WorkerSelection).
- `04-worker-routing.md` — RoundRobin / LeastLoaded / Manual + failover + TTL cache.
- `05-auth-and-2fa.md` — Main cookies, RS256 React→Worker JWTs, S2S OAuth, TOTP, Argon2id.
- `06-core-api-endpoints.md` — `/API/V1/` Auth, Worker, Company, UpdateSchedule.
- `07-role-based-dashboards.md` — capability gating via `EnumPage`.

### Done — Phase 3 (Diagrams)
- `diagrams/erd-main-db.mmd`, `erd-worker-split-db.mmd`, `erd-seedable-config.mmd`.
- `diagrams/seq-company-creation.mmd`, `seq-login-routing.mmd`, `seq-push-update.mmd`.
- `diagrams/README.md` index.

### Done — Phase 4 (Closeout)
- `08-error-contract.md` — JSON envelope, ErrorCategory, CorrelationId, retry rules, idempotency.
- `09-self-update-pointer.md` — pointer to `spec/14-update/` distinguishing push vs self-update.
- `97-acceptance-criteria.md`, `98-changelog.md`, `99-consistency-report.md`.

### Done — Phase 5 (Sync)
- `package.json` bumped 5.9.0 → 5.10.0.
- `scripts/sync-version.mjs` → `version.json` v5.10.0.
- `scripts/sync-spec-tree.mjs` → `src/data/specTree.json` regenerated (651 files / 92 folders).

### Locks (do not regress)
1. PascalCase across schema/JSON/types; PKs `{Table}Id INTEGER AUTOINCREMENT`; no UUIDs.
2. Rule 10/11/12 — `Description`/`Notes`/`Comments` columns nullable, no DEFAULT.
3. Worker routing strategies confined to RoundRobin / LeastLoaded / Manual.
4. Auth tiers: Main = cookie session; React→Worker = RS256 JWT; S2S = OAuth client-credentials.
5. Push update fan-out returns 207 Multi-Status on partial failure.

---
