# 98 — Changelog

**Spec:** `19-main-worker-service`

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
