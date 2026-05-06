# 97 — Acceptance Criteria

**Spec:** `19-main-worker-service`
**Version:** 1.0.0

Direct mapping from verbatim §Acceptance Criteria 1–9 to spec deliverables. Each criterion lists the file(s) where the contract is defined and the test conditions for compliance.

---

## AC-1 — Main server

| Sub-criterion | Defined in | Test |
|---------------|-----------|------|
| Serves UI and React frontend | `01-architecture.md` §2 | Main process serves static React bundle on `/` |
| Holds no business logic | `01-architecture.md` §2 | Code review: no business code outside `app/Edge/` |
| Routes all data requests to workers | `04-worker-routing.md`, `seq-login-routing.mmd` | Integration: every `/API/V1/Company/{slug}` after resolve hits Worker |
| Tracks workers, versions, tenant→worker mappings in SQLite | `03-main-db-schema.md` §2.1, §2.3, §2.7 | Schema migration creates `WorkerNode`, `Company`, `WorkerVersion` |

## AC-2 — Worker server

| Sub-criterion | Defined in | Test |
|---------------|-----------|------|
| Holds all business logic | `01-architecture.md` §2 | Code review: business code lives in Worker repo |
| Has auth, 2FA, session, sign-up, sign-in, JWT or cookie | `05-auth-and-2fa.md` §1 | Endpoint test: all rows of §1 table return 2xx on happy path |
| Has no UI dependency | `01-architecture.md` §2 | Worker repo has no React/Vue/HTML view templates |
| Uses split DB per `spec/05-split-db-architecture/` | `01-architecture.md` §2, `diagrams/erd-worker-split-db.mmd` | Worker startup creates Root/App/Session DBs |

## AC-3 — Company creation

| Sub-criterion | Defined in | Test |
|---------------|-----------|------|
| `POST /API/V1/Company` works end-to-end | `06-core-api-endpoints.md` §2.2, §3.1, `diagrams/seq-company-creation.mmd` | E2E: POST returns 201, Worker has full row, Main has minimal row |
| Worker selected by load-balanced strategy | `04-worker-routing.md` §1 | Unit: 100 sequential creates with `LeastLoaded` distribute within ±10% |
| Main DB stores only minimal company identification | `03-main-db-schema.md` §2.3 | Schema review: only `CompanyId`, `CompanySlug`, `CompanyName`, `WorkerNodeId`, `CompanyAssignedAt`, `Description` |
| Worker DB stores full company data via split-DB | `diagrams/erd-worker-split-db.mmd` | Schema review: `RootCompany` includes Address, Website, Social fields |

## AC-4 — Login and subsequent requests

| Sub-criterion | Defined in | Test |
|---------------|-----------|------|
| First request resolves the worker via Main | `05-auth-and-2fa.md` §6, `diagrams/seq-login-routing.mmd` | E2E: SignIn response includes `WorkerEndpoint` |
| Cache used when available | `04-worker-routing.md` §2 | Unit: second resolve within 15min <!-- TUNABLE-WAIVER: cache TTL — owned by caching-policy memory --> skips DB |
| All subsequent calls go directly to Worker | `01-architecture.md` §3.2 | Network trace: no Main hop after resolve |
| Worker authentication uses JWT, OAuth, or configured method | `05-auth-and-2fa.md` §2.2 | RS256 JWT validated by Worker per `05-auth-and-2fa.md` §7 |

## AC-5 — Self-update (pointer only)

| Sub-criterion | Defined in | Test |
|---------------|-----------|------|
| Endpoint, JSON instruction download, zip-based update flow documented | `09-self-update-pointer.md` §3 | File exists and references `spec/14-update/` |
| Saved redirect URL fallback rule documented | `09-self-update-pointer.md` §3 step 4–5 | File contains stale-hours rule |

> No runtime test. This AC is satisfied by the existence and accuracy of `09-self-update-pointer.md`.

## AC-6 — Push update

| Sub-criterion | Defined in | Test |
|---------------|-----------|------|
| Main can push to all workers or one worker | `06-core-api-endpoints.md` §2.5 (`/Workers/{id}/Update`, `/Workers/All/Update`) | Endpoint test |
| PowerShell-based zip publish endpoint exists | `06-core-api-endpoints.md` §2.5 (`/Workers/PublishZip`), `09-self-update-pointer.md` §5 | Endpoint test with multipart upload |
| Version tracking per worker stored | `03-main-db-schema.md` §2.7 (`WorkerVersion`) | After push, query `WorkerVersion` returns new row per worker |

## AC-7 — Update schedule

| Sub-criterion | Defined in | Test |
|---------------|-----------|------|
| Configurable: hourly, every N hours, daily, weekly, monthly, yearly, specific time + TZ | `06-core-api-endpoints.md` §4, `09-self-update-pointer.md` §6 | Settings PATCH accepts each `Cadence` value |

## AC-8 — Roles

| Sub-criterion | Defined in | Test |
|---------------|-----------|------|
| Power Admin, Admin User, additional roles supported | `07-role-based-dashboards.md` §2 | Seed data inserts both, additional via Seedable-Config |
| Access checked via `User has access to {EnumPage}` pattern | `07-role-based-dashboards.md` §1, §5 | Code grep: zero occurrences of `role === 'PowerAdmin'`, `is_admin` checks, etc. |

## AC-9 — Security

| Sub-criterion | Defined in | Test |
|---------------|-----------|------|
| Passwords salted and strongly encrypted | `05-auth-and-2fa.md` §3 | Argon2id or bcrypt cost ≥12, per-user salt stored |
| Endpoints protected by default; settings allow per-endpoint configuration | `05-auth-and-2fa.md` §8, `06-core-api-endpoints.md` §5 | Default-deny middleware; `EndpointAuthSetting` table exists |

---

## Cross-cutting acceptance

| Concern | Defined in | Test |
|---------|-----------|------|
| Main↔Worker errors use the envelope contract | `08-error-contract.md` §2, §3 | Wire test: every error response matches schema |
| Correlation ID propagated end-to-end | `08-error-contract.md` §4 | Trace test: same `cid` in React, Main, Worker logs |
| Stack flexibility honored | `00-overview.md` §3 | Spec is stack-agnostic; default Laravel called out |
| Replacers applied | `02-glossary.md` §Reserved | grep: zero `CW configuration`, zero `git map` strings in this spec folder |

---

## Backup-tier acceptance (Phases 7–11)

| Criterion | Defined in | Test / failure condition |
|-----------|-----------|--------------------------|
| CDC ledger captures every Insert/Update/Delete on tracked tables | `19-incremental-backup-sync.md` §3 + `AppBackupTrackedTable` seed | Trigger test: row mutation must produce exactly one `SyncOpLedger` row with monotonic `SyncOpSeq` |
| Envelopes are sealed under the Active `KeyEpoch` only | `20-backup-encryption-and-keys.md` §4 | Negative test: sealing under Retired epoch must fail with `WORKER-21090-…` |
| BE-1..BE-6 reject non-S2S traffic at the proxy | `21-backup-endpoints.md` §3 | Wire test: any request without S2S OAuth + `PairingId` claim returns `421 Misdirected Request` |
| Apply pipeline is idempotent on envelope replay | `22-backup-apply-logic.md` §5 (V7) | Replay test: same `EnvelopeId` twice → second call returns 200 with unchanged `LastAcceptedSyncOpSeq`; no duplicate row writes |
| Apply failures land in DLQ — never silently dropped | `22-backup-apply-logic.md` §6 | Fault-injection test: each failure mode produces a `BackupApplyDeadLetter` row with `FailureCode` populated (CODE RED enforcement) |
| Daily snapshot built via `sqlite3_backup_init` | `23-snapshot-storage-and-restore.md` §3 | Build test: snapshot file is openable as SQLite DB and passes `PRAGMA integrity_check` |
| Snapshot zip uses HKDF salt `"BackupSnapshot/v1"` (distinct from envelope salt) | `23-snapshot-storage-and-restore.md` §4 | Crypto test: derived password differs from envelope keystream for the same `KeyEpoch` |
| Restore re-seals under current Active KeyEpoch (forward secrecy) | `23-snapshot-storage-and-restore.md` §6 | Restore test: a restore performed after a key rotation MUST NOT decrypt with the retired epoch |
| Snapshot retention defaults to 30 days rolling, never auto-shortened under disk pressure | `23-…` §7 + `15-tunable-constants.md` §2.15 | Retention sweep test: snapshots older than 30 days deleted unless `Status=Pinned`; disk-pressure path raises alert instead of trimming |
| Successful restore realigns `BackupSyncWatermark` to snapshot's max `SyncOpSeq` | `23-…` §6 + `seq-backup-restore.mmd` | Integration test: BE-1 deliveries after restore must not re-ship pre-snapshot rows |
| `Backup` audience tokens carry mandatory `PairingId` claim | `21-backup-endpoints.md` §3 + `12-jwt-delivery-contract.md` (Phase-12 stub) | Wire test: missing or mismatched `PairingId` → `MAIN-800-04` |
| All Backup-tier tables comply with DB Schema Rules 10/11/12 | `04-database-conventions/` + `erd-backup-tier.mmd` | Linter test: `MISSING-DESC-001` and `DB-FREETEXT-001` pass against the new tables |
| Linter rules `BACKUP-*` and `DB-SYNCOP-*` enforced in CI | `96-linter-audit.md` (Phase-12 promotion) | CI test: linter-scripts run as part of standards enforcement; failure blocks merge |

---

*Acceptance criteria v1.1.0 — 2026-05-06 (Phase 12 — Backup-tier acceptance added).*
