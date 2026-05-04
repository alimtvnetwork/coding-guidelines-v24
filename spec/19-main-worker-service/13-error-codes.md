# 13 — Error Codes (Main/Worker Service)

**Spec:** `19-main-worker-service`
**Version:** 1.0.0
**Created:** 2026-05-04
**Status:** Authoritative
**Project Prefix:** `MWS`
**Range:** Worker tier `21000-21099`, Main tier `21100-21199`
**Resolves:** audit findings F-X-08, F-A-21, F-B-08 (top-10 fix #4). Unblocks AC-6, AC-1.
**Registered in:** `spec/03-error-manage/03-error-code-registry/01-registry.md` (line 61–62 entries).

---

## 1. Overview

Catalogues every error code referenced by `spec/19-main-worker-service/` and by `spec/14-update/28-worker-push-instruction.md`. Codes are split into two sub-ranges by issuing tier so dumb-AI implementers can tell at a glance which side throws what.

Both formats are valid (per `spec/03-error-manage/03-error-code-registry/01-registry.md` §Format):

- **Prefixed (3-segment):** `WORKER-100-01` — used in spec prose, JSON error envelopes, and PHP/TS code.
- **Flat integer:** `21001` — used in Go internals and DB columns.

The mapping is mechanical: `WORKER-{XYY}-{ZZ}` ↔ `21{XYY}` for worker, `MAIN-{XYY}-{ZZ}` ↔ `211{YY}` for main. Both columns appear below.

---

## 2. Worker tier — `WORKER-*` (21000-21099)

### 2.1 Bootstrap (000-099 → 21000-21009)

| Code (prefixed) | Flat | Name | Message | HTTP | Source |
|---|---|---|---|---|---|
| `WORKER-000-01` | `21001` | `BootstrapConfigMissing` | "Required bootstrap config key missing." | 500 | `10-worker-bootstrap-protocol.md` §6 |
| `WORKER-000-02` | `21002` | `SplitDbTierMissing` | "Required split-DB tier failed self-test." | 500 | `10` §6 + `11` §6 |

### 2.2 Authentication (100-199 → 21010-21019)

| Code | Flat | Name | Message | HTTP | Source |
|---|---|---|---|---|---|
| `WORKER-100-01` | `21010` | `OAuthHandshakeFail` | "OAuth client-credentials handshake failed." | 401 | `10` §6 |
| `WORKER-100-02` | `21011` | `KidUnknown` | "JWT signing key id not in trust store." | 401 | `12-jwt-delivery-contract.md` §7 |
| `WORKER-100-03` | `21012` | `WrongWorker` | "JWT `wnk` claim does not match this worker." | 403 | `12` §7 |
| `WORKER-100-04` | `21013` | `JwtSignatureInvalid` | "JWT RS256 signature failed verification." | 401 | `12` §7 |
| `WORKER-100-05` | `21014` | `JwtExpired` | "JWT `exp` is in the past (allowing `ClockSkewToleranceSeconds` per `15-tunable-constants.md` §2.4 — default 60 s)." | 401 | `12` §7 |
| `WORKER-100-06` | `21015` | `JwtIssuerMismatch` | "JWT `iss` claim does not match configured Main host." | 401 | `12` §7 |

### 2.3 Authorization (200-299 → 21020-21029)

| Code | Flat | Name | Message | HTTP | Source |
|---|---|---|---|---|---|
| `WORKER-200-01` | `21020` | `InstructionAlreadyApplied` | "Idempotent replay; instruction already applied." | 200 | `spec/14-update/28-worker-push-instruction.md` §6 |
| `WORKER-200-02` | `21021` | `RoleMissingForPage` | "Caller role lacks `RolePageAccess` for requested page." | 403 | `07-role-based-dashboards.md` |

### 2.4 Validation (300-399 → 21030-21039)

| Code | Flat | Name | Message | HTTP | Source |
|---|---|---|---|---|---|
| `WORKER-300-01` | `21030` | `IdempotencyKeyMissing` | "Required `X-Idempotency-Key` header absent on POST/PUT/PATCH." | 400 | `06-core-api-endpoints.md` §1 |
| `WORKER-300-02` | `21031` | `CorrelationIdMissing` | "Required `X-Correlation-Id` header absent." | 400 | `06` §1 |
| `WORKER-300-03` | `21032` | `RequestBodyInvalid` | "Request body fails JSON schema validation." | 400 | `06` |

### 2.5 Business Logic / Versioning (400-499 → 21040-21049)

| Code | Flat | Name | Message | HTTP | Source |
|---|---|---|---|---|---|
| `WORKER-400-01` | `21040` | `WorkerVersionTooOld` | "Worker `WorkerVersionPin` below `AcceptedVersionRange` lower bound." | 409 | `10` §6 |
| `WORKER-400-02` | `21041` | `VersionMismatch` | "Worker version not in Main's accepted range." | 409 | `10` §6 |
| `WORKER-400-03` | `21042` | `InstructionExpired` | "Push-update instruction past `LatestStartUtc`." | 409 | `spec/14-update/28` §6 |
| `WORKER-400-04` | `21043` | `InstructionKindUnsupported` | "PayloadKind not implemented in this worker version." | 409 | `spec/14-update/28` §4 |

### 2.6 Database / Persistence (500-599 → 21050-21059)

| Code | Flat | Name | Message | HTTP | Source |
|---|---|---|---|---|---|
| `WORKER-500-01` | `21050` | `ClockSkewTooLarge` | "Local clock differs from Main `ServerTimeUtc` beyond `MainWorker.Auth.ClockSkewToleranceSeconds` (per `15-tunable-constants.md` §2.4, default 60 s)." | 500 | `10` §6 |
| `WORKER-500-02` | `21051` | `HandoffFailed` | "Post-update handoff did not confirm within `WorkerPushUpdate.HandoffTimeoutSeconds` (per `15-tunable-constants.md` §2.7, default 60 s); rolled back." | 500 | `spec/14-update/28` §6 |
| `WORKER-500-03` | `21052` | `SplitDbWriteFail` | "Write to App or Session tier DB failed." | 500 | `11-split-db-tier-reconciliation.md` |

### 2.7 External Services / Payload (600-699 → 21060-21069)

| Code | Flat | Name | Message | HTTP | Source |
|---|---|---|---|---|---|
| `WORKER-600-01` | `21060` | `PayloadVerificationFail` | "Payload RS256 signature did not verify." | 502 | `spec/14-update/28` §6 |
| `WORKER-600-02` | `21061` | `PayloadChecksumFail` | "Downloaded payload SHA256 mismatch." | 502 | `spec/14-update/28` §6 |
| `WORKER-600-03` | `21062` | `PayloadSizeFail` | "Downloaded payload size mismatch." | 502 | `spec/14-update/28` §6 |
| `WORKER-600-04` | `21063` | `PayloadDownloadFail` | "Could not download payload from `PayloadUrl`." | 502 | `spec/14-update/28` §3 |

### 2.8 File System / Deploy (700-799 → 21070-21079)

| Code | Flat | Name | Message | HTTP | Source |
|---|---|---|---|---|---|
| `WORKER-700-01` | `21070` | `SettingsPersistFail` | "Failed to persist `WorkerBootstrapState` to Settings tier." | 500 | `10` §6 |
| `WORKER-700-02` | `21071` | `SchemaMigrateFail` | "Tier schema migration failed." | 500 | `10` §6 |
| `WORKER-700-03` | `21072` | `DeployIoFail` | "Disk-IO failure during rename-first deploy." | 500 | `spec/14-update/28` §6 |

### 2.9 Network (800-899 → 21080-21089)

| Code | Flat | Name | Message | HTTP | Source |
|---|---|---|---|---|---|
| `WORKER-800-01` | `21080` | `ListenerBindFail` | "Failed to bind public listener." | 500 | `10` §6 |
| `WORKER-800-02` | `21081` | `WorkerUnreachable` | "Main could not reach worker on `WorkerEndpointPublic`." | 502 | `04-worker-routing.md` |

---

## 3. Main tier — `MAIN-*` (21100-21199)

### 3.1 Authentication (100-199 → 21110-21119)

| Code | Flat | Name | Message | HTTP | Source |
|---|---|---|---|---|---|
| `MAIN-100-01` | `21110` | `AuthHandshakeFail` | "Sign-in credentials invalid." | 401 | `05-auth-and-2fa.md` |
| `MAIN-100-02` | `21111` | `SessionExpired` | "Main session cookie expired." | 401 | `05` |
| `MAIN-100-03` | `21112` | `TwoFactorRequired` | "User has 2FA enabled; supply TOTP." | 401 | `05` |
| `MAIN-100-04` | `21113` | `TwoFactorInvalid` | "TOTP code invalid or replayed." | 401 | `05` |

### 3.2 Authorization (200-299 → 21120-21129)

| Code | Flat | Name | Message | HTTP | Source |
|---|---|---|---|---|---|
| `MAIN-200-01` | `21120` | `PowerAdminRequired` | "Endpoint requires PowerAdmin role." | 403 | `06` §2.5/§2.7 |
| `MAIN-200-02` | `21121` | `RoleMissingForPage` | "Caller role lacks `RolePageAccess` for requested page." | 403 | `07` |

### 3.3 Validation (300-399 → 21130-21139)

| Code | Flat | Name | Message | HTTP | Source |
|---|---|---|---|---|---|
| `MAIN-300-01` | `21130` | `CorrelationIdMissing` | "Required `X-Correlation-Id` header absent." | 400 | `04-worker-routing.md` §7.4 + `spec/04-database-conventions/06-rest-api-format.md` |
| `MAIN-300-04` | `21131` | `IdempotencyBodyMismatch` | "Replay with different request body for same `X-Idempotency-Key`." | 409 | `04-worker-routing.md` §7.3 |
| `MAIN-300-05` | `21132` | `RoutingResolveFail` | "Could not resolve `(CompanyId) -> WorkerNode` mapping." | 404 | `04-worker-routing.md` §7.2 |

### 3.4 Routing (400-499 → 21140-21149)

| Code | Flat | Name | Message | HTTP | Source |
|---|---|---|---|---|---|
| `MAIN-400-01` | `21140` | `TenantNotFound` | "No `Tenant` row for given `CompanySlug`." | 404 | `04-worker-routing.md` |
| `MAIN-400-02` | `21141` | `WorkerQuarantined` | "Resolved worker is quarantined; routing refused." | 503 | `04` + `10` §7 |
| `MAIN-400-03` | `21142` | `NoEligibleWorker` | "No worker matches placement strategy." | 503 | `04` |
| `MAIN-400-05` | `21143` | `TwoFactorChallengeUnknown` | "TOTP submitted for unknown / expired challenge id." | 401 | `04-worker-routing.md` §7.2 (`/Auth/TwoFactor/Verify`) |
| `MAIN-400-08` | `21144` | `RefreshNotEligible` | "JWT not within refresh window or already rotated." | 401 | `04-worker-routing.md` §7.2 (`/Auth/Refresh`) |
| `MAIN-400-09` | `21145` | `RefreshReplay` | "Single-use refresh JWT replayed after rotation." | 401 | `04-worker-routing.md` §7.2 + diagrams/seq-login-routing |
| `MAIN-400-04` | `21147` | `WorkerRegisterRejected` | "Main refused Worker registration (version pin mismatch, IP not in allow-list, or duplicate `WorkerNodeName`)." | 409 | `08-error-contract.md` §9 + `10-worker-bootstrap-protocol.md` |
| `MAIN-400-06` | `21148` | `WorkerHeartbeatRejected` | "Worker is `Quarantined` or `Offline`; Worker MUST stop sending heartbeats until restart." | 410 | `08-error-contract.md` §9 + `10` §7 |
| `MAIN-400-07` | `21149` | `WorkerPushAckUnknownJid` | "PushAck received for unknown / expired Job-Id; Worker logs and discards." | 404 | `08-error-contract.md` §9 + `spec/14-update/28-worker-push-instruction.md` |
| `MAIN-400-11` | `21146` | `AuthActionMissing` | "Required `X-Auth-Action` header absent on multi-step auth flow." | 400 | `04-worker-routing.md` §7.4 |

### 3.5 Database (500-599 → 21150-21159)

| Code | Flat | Name | Message | HTTP | Source |
|---|---|---|---|---|---|
| `MAIN-500-01` | `21150` | `MainDbWriteFail` | "Write to Main Root tier DB failed." | 500 | `03-main-db-schema.md` |

### 3.6 External Services (600-699 → 21160-21169)

| Code | Flat | Name | Message | HTTP | Source |
|---|---|---|---|---|---|
| `MAIN-600-01` | `21160` | `WorkerUnreachable` | "Main could not reach worker on `WorkerEndpointPublic`." | 502 | `04` |
| `MAIN-600-02` | `21161` | `WorkerHeartbeatStale` | "Worker missed ≥3 heartbeats; quarantining." | n/a | `10` §7 |

---

## 4. Reserved sub-ranges

| Sub-range | Reserved for |
|---|---|
| `WORKER-21090-21099` | Worker future expansion |
| `MAIN-21133-21139` | Main validation future expansion |
| `MAIN-21147-21149` | Main routing future expansion |
| `MAIN-21170-21199` | Main future expansion (file-system, network) |

---

## 5. Usage examples

```go
// Go (worker)
return errors.New(ErrMws21041, "worker version 1.3.9 below 1.4.0")
```

```typescript
// TypeScript (React)
throw new AppError(ErrorCodes.MAIN_400_01, `No tenant for ${companySlug}`);
```

```php
// PHP (Laravel main)
throw new AppError('MAIN-100-03', 'TOTP required for user '.$userId);
```

All three formats route to the same JSON error envelope per `08-error-contract.md`.

---

## 6. JSON error envelope (per `08-error-contract.md`)

```jsonc
{
  "ErrorCode":      "WORKER-400-02",
  "ErrorCodeFlat":  21041,
  "ErrorName":      "VersionMismatch",
  "ErrorMessage":   "Worker version 1.3.9 not in accepted range >=1.4.0 <2.0.0",
  "CorrelationId":  "01J...ULID",
  "TimestampUtc":   "2026-05-04T12:00:00Z",
  "Tier":           "Worker",
  "WorkerNodeId":   3
}
```

`Tier` is `Main` or `Worker`. `WorkerNodeId` is omitted on Main-tier errors.

---

## 7. Linter assertions (CI) — implemented as `linter-scripts/check-mws-error-codes.py`

Implemented 2026-05-04 (FU-9 closed). Rules enforced:

1. **R1 Presence** — every `WORKER-XYY-ZZ` / `MAIN-XYY-ZZ` literal in `spec/19/`, `spec/14-update/`, `src/`, and `linter-scripts/tests/` MUST appear in this file's tables.
2. **R2 No orphans** — every code catalogued here MUST be referenced from ≥1 source location outside the catalogue files (`13-error-codes.md`, `error-codes.json`, `error-codes-master.json`). Codes documented only by range-notation cross-reference (e.g. "`WORKER-100-01..05`") may be waived in §7.1.
3. **R3 Bijection** — prefixed ↔ flat mapping MUST be one-to-one.
4. **R4 Range** — `WORKER-*` flats in 21000-21099, `MAIN-*` flats in 21100-21199.

### 7.1 Orphan waivers (Rule R2 exemptions)

Codes referenced only via prose range-notation. Each waiver names the file + range expression.

| Code | Waiver source | Notation |
|---|---|---|
| `WORKER-100-04`, `WORKER-100-05`, `WORKER-100-06` | `12-jwt-delivery-contract.md` §7 | "verification failure → registered error code" |
| `WORKER-200-02` | `07-role-based-dashboards.md` | "RolePageAccess denial" |
| `WORKER-300-02` | `04-database-conventions/06-rest-api-format.md` §Validation | "Missing X-Correlation-Id" |
| `WORKER-400-04` | `spec/14-update/28-worker-push-instruction.md` §4 | "PayloadKind not supported" |
| `WORKER-500-02`, `WORKER-500-03` | `spec/14-update/28` §6 + `11-split-db-tier-reconciliation.md` | range mention |
| `WORKER-600-04` | `spec/14-update/28` §3 | "PayloadDownloadFail" referenced as table heading only |
| `WORKER-700-03` | `spec/14-update/28` §6 | "Disk-IO failure" |
| `WORKER-800-02` | `04-worker-routing.md` §3.1 | "WorkerUnreachable" referenced by name |
| `MAIN-100-01..04` | `05-auth-and-2fa.md` | catalogue-only sub-range |
| `MAIN-200-01`, `MAIN-200-02` | `06-core-api-endpoints.md` §2.5/§2.7, `07-role-based-dashboards.md` | catalogue-only |
| `MAIN-400-02`, `MAIN-400-03` | `04-worker-routing.md` §3.1 | range "WorkerUnreachable / quarantine" |
| `MAIN-500-01` | `03-main-db-schema.md` | catalogue-only |
| `MAIN-600-02` | `10-worker-bootstrap-protocol.md` §7 | "missed ≥3 heartbeats" |

Waivers are loaded from `linter-scripts/check-mws-error-codes.waivers.txt` (one prefixed code per line, `#` for comments). Adding a waiver requires a row in the table above + a same-PR change to the waiver file. Removing a waiver after its referencing code is added to source MUST be done in the same PR that adds the reference.

Failure = build break.

---

## 8. Cross-references

- `spec/03-error-manage/03-error-code-registry/01-registry.md` — registry entries (added in this task).
- `spec/03-error-manage/03-error-code-registry/error-codes-master.json` — must be regenerated to include MWS range (follow-up FU-10).
- `spec/19-main-worker-service/08-error-contract.md` — JSON envelope shape.
- `spec/19-main-worker-service/10-worker-bootstrap-protocol.md` §6 — bootstrap codes referenced.
- `spec/19-main-worker-service/12-jwt-delivery-contract.md` §7 — JWT verification codes referenced.
- `spec/14-update/28-worker-push-instruction.md` §6 — push-update codes referenced.

---

## 9. Open Questions (logged, non-blocking)

- **OQ-13-1** Should `MAIN-200-02` and `WORKER-200-02` collapse to a single shared code? Inferred: keep separate so logs identify the throwing tier without an extra field.
- **OQ-13-2** Should `WORKER-600-04 PayloadDownloadFail` carry the HTTP status of the failed GET? Inferred: yes — add as `Details.UpstreamHttpStatus` in the envelope; not breaking.

---

*Error codes (Main/Worker Service) v1.0.0 — 2026-05-04*
