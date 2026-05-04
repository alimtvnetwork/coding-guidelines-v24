# 98 — Changelog

**Spec:** `19-main-worker-service`

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
- `diagrams/README.md`

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

*Changelog v1.0.0 — 2026-05-04*
