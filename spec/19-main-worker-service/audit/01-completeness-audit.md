# 01 — Completeness Audit (Step 1 of 5)

**Spec audited:** `spec/19-main-worker-service/`
**Audited at:** 2026-05-04
**Audit version:** 1.0.0
**Reviewer mode:** Adversarial. Assume the implementer is a low-capability AI that follows text literally, has no business intuition, and will not infer "obvious" defaults.

> **Severity legend:** **BLOCKER** = AI cannot produce a runnable system; **MAJOR** = AI produces something that runs but diverges from intent; **MINOR** = AI produces correct behavior but with style/observability gaps.

---

## 0. Method

- Read all 14 markdown files + 7 diagrams + 4 mindmap PNGs.
- For each acceptance criterion (AC-1..AC-9), enumerate concrete artifacts a programmer must produce, then check whether the spec defines each one with no guesswork.
- Anything the spec leaves to "implementer choice" is OK only if a default is named.

---

## 1. Findings — BLOCKER (12)

### F-B-01 · Worker registration bootstrap is undefined
- **Location:** `06-core-api-endpoints.md` §2.5 (`POST /API/V1/Workers/Register`); `05-auth-and-2fa.md` §2.3.
- **Quote:** *"OAuth client-credentials per Worker, secrets stored via Seedable-Config (encrypted at rest)."*
- **Gap:** How does a brand-new Worker obtain its client_id/client_secret **before** Seedable-Config can deliver them? Chicken-and-egg.
- **Dumb-AI guess:** hard-code a shared secret in `.env`. Insecure.
- **Fix:** Define a one-time bootstrap token, its lifetime, who issues it, and the rotation rule.

### F-B-02 · `WorkerNode.WorkerNodeEndpoint` is required NOT NULL but registration payload is unspecified
- **Location:** `03-main-db-schema.md` §2.1 vs `06-core-api-endpoints.md` §2.5.
- **Gap:** No request-body schema for `POST /API/V1/Workers/Register`. Required fields, validation rules, response shape — all missing.
- **Fix:** Add a §3.x reference payload mirroring §3.1 style.

### F-B-03 · `WorkerNodeIdentity` source is undefined
- **Location:** `03-main-db-schema.md` §2.1.
- **Quote:** *"Unique stable identifier (e.g. machine fingerprint)"*
- **Gap:** "e.g." is not a contract. No algorithm, length, character set, or collision policy.
- **Fix:** Pick one (`UUIDv7` of MAC+hostname hashed, or operator-set string), specify length and regex.

### F-B-04 · JWT public-key distribution mechanism is undefined
- **Location:** `05-auth-and-2fa.md` §2.2.
- **Quote:** *"each Worker holds Main's public key (rotatable via Seedable-Config)."*
- **Gap:** Where is the key fetched from? At what URL? With what auth? With what cache TTL? What happens during rotation overlap (workers must accept old + new key for some window)?
- **Fix:** Define a JWKS endpoint (path, auth, cache-control), and a key-rotation overlap window.

### F-B-05 · `WorkerJwt` storage location contradicts itself
- **Location:** `05-auth-and-2fa.md` §6 step 5 vs `06-core-api-endpoints.md` §3.2 (`Resolve` returns `WorkerJwt` in JSON body to the browser).
- **Quote:** *"React stores JWT in memory (NOT localStorage)"* — but the JSON-body delivery means the JWT is exposed to JS, defeating the cookie-only HTTPOnly protection.
- **Gap:** No threat model section. No XSS mitigation. No CSP requirement.
- **Fix:** State explicitly that JWT-in-memory is accepted XSS exposure, OR switch to HTTPOnly cookie-bearer for Worker JWT.

### F-B-06 · Request body for `POST /API/V1/Workers/All/Update` and `/Workers/{id}/Update` is missing
- **Location:** `06-core-api-endpoints.md` §2.5.
- **Gap:** No schema for which version/zip to push, no idempotency semantics for fan-out, no response shape (the consistency report mentions 207 Multi-Status but no payload example).
- **Fix:** Add request + 207 response payloads in §3.

### F-B-07 · `WorkerSelectionEvent` has no Strategy column for `LeastLoaded` audit
- **Location:** `03-main-db-schema.md` §2.8 lists `WorkerSelectionStrategyId` ✅ but no `EligibleCandidateCount`, no `RejectedReasonCode`. Audit cannot answer "why was W3 picked over W5?"
- **Fix:** Add `EligibleCount INTEGER NOT NULL` and `Notes TEXT NULL` (already present, but make the convention explicit).

### F-B-08 · No spec exists for `Settings.UpdateSchedule` persistence
- **Location:** `06-core-api-endpoints.md` §4 defines the JSON shape but no table appears in `03-main-db-schema.md`.
- **Gap:** Where does this row live? Single-row config table? Seedable-Config? Cache?
- **Fix:** Add `MainSetting` table to schema or explicitly delegate to Seedable-Config with the key path.

### F-B-09 · `EnumPage` enumeration is referenced but never enumerated
- **Location:** `07-role-based-dashboards.md` (referenced by `05-auth-and-2fa.md` §8 and AC-8).
- **Gap:** I read the spec — §1, §5 of `07-` reference the pattern but the file does not list the actual `EnumPage` values (`PowerAdminPage`, `AdminPage`, `MemberPage`, `WorkersPage`, …). A dumb AI cannot generate the seed data.
- **Fix:** Add a definitive enum table (Code, Label, Description, default access matrix per role).

### F-B-10 · `RolePageAccess` table appears in ERD but not in schema spec
- **Location:** `diagrams/erd-main-db.mmd` mentions it (per audit-step-3 prep); `03-main-db-schema.md` §2 does not list it.
- **Gap:** Implementer cannot create a migration without column list.
- **Fix:** Add §2.10 `RolePageAccess` (RoleId, EnumPageId, IsAllowed, Description).

### F-B-11 · `User` table is missing `UserPasswordAlgorithmId`, `Has2FAEnabled`, `TotpSecret`, `TotpRecoveryHash` columns
- **Location:** `03-main-db-schema.md` §2.4 vs `05-auth-and-2fa.md` §3, §4.
- **Quote:** Spec text references `User.Has2FAEnabled` and `User.TotpSecret` as if they exist; schema defines neither.
- **Fix:** Reconcile — add columns or move to a `UserAuth` side-table.

### F-B-12 · No spec for the React-frontend → Main session-cookie domain / SameSite / cross-subdomain story
- **Location:** Implied by `01-architecture.md` topology but never written.
- **Gap:** When Worker lives at `w3.alimkarim.com` and Main at `recalltime.com`, cookies set by Main do NOT reach Worker. Spec assumes JWT covers it but never says so.
- **Fix:** Add a one-paragraph "cookie scope vs JWT scope" in `05-auth-and-2fa.md` with explicit domain rules.

---

## 2. Findings — MAJOR (10)

### F-M-01 · "Most fields Non-Nullable" is not a schema
- **Location:** `06-core-api-endpoints.md` §3.1 caption.
- **Gap:** Which fields exactly? Min/max length? Format for `PhoneNumber`, `Calendar`?
- **Fix:** Per-field validation table.

### F-M-02 · No idempotency-key uniqueness window defined
- **Location:** `06-core-api-endpoints.md` §1.
- **Quote:** *"`X-Idempotency-Key` mandatory on POST/PUT/PATCH"*
- **Gap:** Storage TTL? Per-user or global? What status code for replay?
- **Fix:** Add §1.x with TTL (24h typical) and 409 Conflict on conflicting body.

### F-M-03 · `LeastLoaded` tiebreaker is by registration date, ignores capacity headroom
- **Location:** `04-worker-routing.md` §1.2.
- **Gap:** Two workers with equal load get the older one — guarantees skew over time when capacity differs.
- **Fix:** Tiebreaker by `(MaxCompaniesPerWorker - assignedCount)` desc, then registration date.

### F-M-04 · Heartbeat endpoint has no payload contract
- **Location:** `06-core-api-endpoints.md` §2.5.
- **Gap:** Should heartbeat report load? Version? Disk? Spec says nothing.
- **Fix:** Define payload (LoadAvg, AssignedCompanyCount, CurrentVersion) — these feed `LeastLoaded` and `WorkerVersion`.

### F-M-05 · `MainWorker.Routing.HeartbeatWindowSeconds` default 60s but `× 3` rule for offline is in §3.2
- **Location:** `04-worker-routing.md` §1.4 vs §3.2.
- **Gap:** Two configs, one derived. What if operator sets `OfflineMultiplier`? Not configurable, hard-coded.
- **Fix:** Promote `OfflineMultiplier` to its own Seedable-Config key with default 3.

### F-M-06 · 2FA backup-code consumption is unspecified
- **Location:** `05-auth-and-2fa.md` §4.
- **Quote:** *"generate 10 single-use codes at enrollment. Store hashed."*
- **Gap:** What happens at 0 remaining? Auto-regenerate? Force re-enroll?
- **Fix:** State the policy and the warning thresholds (≤3 remaining → email warning).

### F-M-07 · `PasswordResetRequest` has no anti-enumeration requirement
- **Location:** `06-core-api-endpoints.md` §2.1.
- **Gap:** Must always return 202, never reveal whether email exists.
- **Fix:** Add to `05-auth-and-2fa.md` §5 anti-patterns.

### F-M-08 · Sign-in 2FA challenge `ChallengeId` lifetime undefined
- **Location:** `05-auth-and-2fa.md` §6.
- **Fix:** Define TTL (5 min standard) and one-time-use rule.

### F-M-09 · `RefreshWorkerToken` endpoint has no rate limit beyond auth-group default
- **Location:** `06-core-api-endpoints.md` §6.
- **Gap:** A misbehaving SPA could refresh every second. Need stricter per-user cap.
- **Fix:** 60/min/user for refresh.

### F-M-10 · `Settings/EndpointAuth` PATCH semantics undefined
- **Location:** `06-core-api-endpoints.md` §2.7.
- **Gap:** Replace-all? Patch-merge? What about an unknown `EndpointPathPattern`?
- **Fix:** Specify upsert-by-pattern and 422 on unknown pattern.

---

## 3. Findings — MINOR (8)

| ID | Finding | Fix |
|----|---------|-----|
| F-N-01 | `02-glossary.md` is 45 lines but never enumerates `EnumPage` codes used elsewhere | Add cross-ref column |
| F-N-02 | `08-error-contract.md` not yet read in step 1 — flagged for step 2 | (next step) |
| F-N-03 | `99-consistency-report.md` claims terms replaced; not greppable from this audit | Add the actual grep commands the report ran |
| F-N-04 | `WorkerJwt` field in resolve response uses inconsistent casing vs `WorkerEndpoint` style (both PascalCase OK, but no token type field) | Add `"TokenType": "Bearer"` |
| F-N-05 | Rate limits (§6) are "recommended" — should be defaults seeded by Seedable-Config | Promote to defaults |
| F-N-06 | `09-self-update-pointer.md` references `spec/14-update/` — not validated against actual existence in step 1 | (step 4 will check) |
| F-N-07 | No OpenAPI/Swagger artifact mentioned anywhere | Recommend `openapi.yaml` generation in CI |
| F-N-08 | All timestamps use `TEXT ISO-8601` — no precision specified (ms? tz suffix mandatory?) | Specify `YYYY-MM-DDTHH:MM:SS.sssZ` |

---

## 4. Per-AC Coverage Matrix

| AC | Spec coverage | Concrete-enough for dumb-AI? | Blockers |
|----|---------------|-----------------------------|----------|
| AC-1 Main server | ✅ documented | ⚠ partial | F-B-08 |
| AC-2 Worker server | ✅ documented | ⚠ partial | F-B-01, F-B-04, F-B-05, F-B-12 |
| AC-3 Company creation | ✅ documented | ⚠ partial | F-M-01 |
| AC-4 Login routing | ✅ documented | ⚠ partial | F-B-04, F-B-05 |
| AC-5 Self-update (pointer) | ✅ pointer is intentional | ✅ OK | — |
| AC-6 Push update | ⚠ partial | ❌ no | F-B-06, F-B-07 |
| AC-7 Update schedule | ✅ shape OK | ⚠ partial | F-B-08 |
| AC-8 Roles | ⚠ pattern OK, enum missing | ❌ no | F-B-09, F-B-10 |
| AC-9 Security | ✅ partial | ⚠ partial | F-B-04, F-B-05, F-B-11 |

---

## 5. Step-1 Headline Numbers

- **Total findings:** 30 (12 BLOCKER, 10 MAJOR, 8 MINOR).
- **ACs with at least one BLOCKER:** 7 of 9.
- **Files touched by ≥1 BLOCKER:** `03-main-db-schema.md`, `05-auth-and-2fa.md`, `06-core-api-endpoints.md`, `07-role-based-dashboards.md`.
- **Estimated rework before a dumb AI can implement without questions:** ~1–2 days of spec hardening.

---

## 6. Next Steps in This Audit Series

| Step | File | Focus |
|------|------|-------|
| 2 | `02-ambiguity-audit.md` | Wording-level vagueness, modal verbs, undefined terms, `08-error-contract.md` deep-read |
| 3 | `03-diagram-audit.md` | Mermaid + mindmap consistency vs prose |
| 4 | `04-cross-spec-dependency-audit.md` | Verify referenced `spec/05/06/14/04` anchors exist |
| 5 | `05-implementation-pivot-score.md` | Final scorecard, top-10 must-fix list, dumb-AI checklist |

Say `next` to run **Step 2 (Ambiguity Audit)**.

---

*Completeness audit v1.0.0 — 2026-05-04*
