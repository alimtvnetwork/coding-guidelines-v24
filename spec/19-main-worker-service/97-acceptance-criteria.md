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

*Acceptance criteria v1.0.0 — 2026-05-04*
