# 98 — Changelog

**Spec:** `19-main-worker-service`

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
