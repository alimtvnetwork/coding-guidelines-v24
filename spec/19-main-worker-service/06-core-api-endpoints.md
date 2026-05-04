# 06 — Core API Endpoints

**Spec:** `19-main-worker-service`
**Version:** 1.0.0

Authoritative REST surface for both tiers. All paths are `/API/V1/...`. JSON request/response with PascalCase keys.

---

## 1. Conventions

- **Versioning:** path-based (`/API/V1`).
- **Content-Type:** `application/json` request and response.
- **JSON keys:** PascalCase.
- **Errors:** envelope per `08-error-contract.md`.
- **Headers:** `X-Correlation-Id` mandatory inbound (generated server-side if missing); `X-Idempotency-Key` mandatory on POST/PUT/PATCH.
- **Auth column meanings:** `Session` = cookie; `JWT` = worker JWT (RS256); `OAuth` = client-credentials; `None` = public.

---

## 2. Endpoint Catalog

### 2.1 Auth (Main)

| Method | Path | Auth | Purpose |
|--------|------|------|---------|
| POST | `/API/V1/Auth/SignUp` | None | Create user (and optionally company) |
| POST | `/API/V1/Auth/SignIn` | None | Verify password, start session |
| POST | `/API/V1/Auth/Verify2FA` | None (challenge-bound) | Submit TOTP for active challenge |
| POST | `/API/V1/Auth/SignOut` | Session | End current session |
| POST | `/API/V1/Auth/SignOutAll` | Session | End all sessions for user |
| POST | `/API/V1/Auth/PasswordResetRequest` | None | Email reset link |
| POST | `/API/V1/Auth/PasswordResetConfirm` | None (token-bound) | Set new password |
| POST | `/API/V1/Auth/Enroll2FA` | Session | Begin TOTP enrollment |
| POST | `/API/V1/Auth/Confirm2FA` | Session | Confirm TOTP code, enable 2FA |
| POST | `/API/V1/Auth/Disable2FA` | Session + TOTP | Disable 2FA |
| POST | `/API/V1/Auth/RefreshWorkerToken` | Session | Mint a fresh worker JWT |

### 2.2 Company (Main → routes to Worker)

| Method | Path | Auth | Purpose |
|--------|------|------|---------|
| POST | `/API/V1/Company` | Session | Create company (Main routes to Worker) |
| GET | `/API/V1/Company/{CompanySlug}/Resolve` | Session | Returns `WorkerEndpoint` + worker JWT |
| GET | `/API/V1/Company/{CompanySlug}` | JWT (on Worker) | Read company (after resolve) |
| PATCH | `/API/V1/Company/{CompanySlug}` | JWT (on Worker) | Update company |
| DELETE | `/API/V1/Company/{CompanySlug}` | JWT (on Worker) + PowerAdmin | Hard-delete |

### 2.3 User (Main → routes to Worker)

| Method | Path | Auth | Purpose |
|--------|------|------|---------|
| POST | `/API/V1/Users` | Session + access | Create user under current company |
| GET | `/API/V1/Users/{UserId}` | JWT (on Worker) | Read user |
| PATCH | `/API/V1/Users/{UserId}` | JWT (on Worker) + access | Update user |

### 2.4 Status / Version (both tiers, intentionally minimal)

| Method | Path | Auth | Purpose |
|--------|------|------|---------|
| GET | `/API/V1/Status` | None (default; configurable) | Liveness + readiness |
| GET | `/API/V1/Version` | None (default; configurable) | App name, title, current version, update available flag |

### 2.5 Workers (Main only — Power Admin)

| Method | Path | Auth | Purpose |
|--------|------|------|---------|
| GET | `/API/V1/Workers` | Session + PowerAdmin | List worker registry |
| POST | `/API/V1/Workers/Register` | OAuth | Worker registers itself with Main |
| POST | `/API/V1/Workers/{WorkerNodeId}/Heartbeat` | OAuth | Worker liveness ping |
| POST | `/API/V1/Workers/{WorkerNodeId}/Update` | Session + PowerAdmin | Push update to one worker |
| POST | `/API/V1/Workers/All/Update` | Session + PowerAdmin | Push update to all workers |
| POST | `/API/V1/Workers/PublishZip` | Session + PowerAdmin (multipart) | Upload deployment zip via PowerShell |

### 2.6 Self-Update (both tiers — pointer only)

| Method | Path | Auth | Purpose |
|--------|------|------|---------|
| POST | `/API/V1/SelfUpdate` | OAuth | Trigger self-update workflow (see `09-self-update-pointer.md`) |

### 2.7 Settings (Main — Power Admin)

| Method | Path | Auth | Purpose |
|--------|------|------|---------|
| GET | `/API/V1/Settings/EndpointAuth` | Session + PowerAdmin | Per-endpoint auth toggles |
| PATCH | `/API/V1/Settings/EndpointAuth` | Session + PowerAdmin | Update toggles (see OQ-1 in `05-auth-and-2fa.md`) |
| GET | `/API/V1/Settings/UpdateSchedule` | Session + PowerAdmin | Read schedule |
| PATCH | `/API/V1/Settings/UpdateSchedule` | Session + PowerAdmin | Update schedule (see §4) |

---

## 3. Reference Payloads

### 3.1 `POST /API/V1/Company` request

```json
{
  "CompanyName": "Riseup Asia LLC",
  "CompanyWebsite": "https://riseup-asia.com",
  "CompanySlug": "riseup-asia",
  "Address": "Kuala Lumpur, Malaysia",
  "PhoneNumber": "+60-...",
  "NumberOfPeople": 25,
  "Calendar": "https://cal.com/...",
  "WhatsApp": "+60-...",
  "Facebook": "https://facebook.com/...",
  "LinkedIn": "https://linkedin.com/company/...",
  "PreferredWorkerNodeId": null
}
```

Field nullability (resolves F-A-01 — replaces prior "Most fields Non-Nullable"):

| Field | Nullable | Notes |
|-------|----------|-------|
| `CompanyName` | NO | |
| `CompanyWebsite` | YES | |
| `CompanySlug` | NO | Unique. |
| `Address` | YES | |
| `PhoneNumber` | NO | E.164 preferred. |
| `NumberOfPeople` | YES | Integer ≥ 1 when set. |
| `Calendar` | YES | |
| `WhatsApp` | YES | |
| `Facebook` | YES | |
| `LinkedIn` | YES | |
| `PreferredWorkerNodeId` | YES | Only honored when worker-selection strategy is `Manual` AND caller has `PowerAdmin` access; ignored otherwise (no error). |

### 3.2 `GET /API/V1/Company/{CompanySlug}/Resolve` response

```json
{
  "CompanySlug": "riseup-asia",
  "CompanyId": 42,
  "WorkerNodeId": 3,
  "WorkerEndpoint": "https://w3.example.com",
  "WorkerJwt": "<RS256 token>",
  "JwtExpiresAt": "2026-05-04T12:15:00Z"
}
```

### 3.3 `GET /API/V1/Version` response

```json
{
  "ApplicationName": "MainServer",
  "ApplicationTitle": "Coordinator",
  "CurrentVersion": "1.0.0",
  "UpdateAvailable": false,
  "LatestVersion": null
}
```

---

## 4. Update Schedule Settings (per verbatim §Update Schedule Settings)

`Settings.UpdateSchedule` shape:

```json
{
  "Cadence": "Weekly",
  "EveryNHours": null,
  "SpecificTimeOfDay": "04:00",
  "TimeZone": "Asia/Kuala_Lumpur",
  "Enabled": true
}
```

`Cadence` enum: `Hourly`, `EveryNHours`, `Daily`, `Weekly`, `Monthly`, `Yearly`. When `EveryNHours`, populate `EveryNHours` (allowed values 5, 6, 12, 24). `SpecificTimeOfDay` is null unless `Cadence` ∈ {Daily, Weekly, Monthly, Yearly}. Default: `Weekly @ 04:00 Asia/Kuala_Lumpur`.

---

## 5. Settings Schema Sketch (for OQ-1)

```sql
CREATE TABLE EndpointAuthSetting (
    EndpointAuthSettingId INTEGER PRIMARY KEY AUTOINCREMENT,
    EndpointPathPattern   TEXT NOT NULL,
    AuthMechanismId       INTEGER NOT NULL REFERENCES AuthMechanism(AuthMechanismId),
    IsEnabled             INTEGER NOT NULL,
    Description           TEXT NULL
);

CREATE TABLE AuthMechanism (
    AuthMechanismId   INTEGER PRIMARY KEY AUTOINCREMENT,
    AuthMechanismCode TEXT NOT NULL UNIQUE,  -- Session | Jwt | OAuth | None
    AuthMechanismLabel TEXT NOT NULL,
    Description       TEXT NULL
);
```

Final design awaits OQ-1 resolution.

---

## 6. Rate Limiting (defaults — MUST apply unless overridden via Seedable-Config)

Defaults below are MANDATORY out-of-the-box. Implementations MUST apply them unless explicitly overridden via Seedable-Config keys per `15-tunable-constants.md` §2.6 (resolves F-A-02 — replaces prior "recommended defaults" softening).

| Endpoint group | Default | Seedable-Config key |
|----------------|---------|---------------------|
| `/API/V1/Auth/*` | 10 / minute / IP | `MainWorker.RateAuthPerMinutePerIp` |
| `/API/V1/Workers/*` | 60 / minute / token | `MainWorker.RateWorkerPerMinutePerToken` |
| Other authenticated | 600 / minute / user | `MainWorker.RateOtherPerMinutePerUser` |

Implementer uses framework-native middleware (e.g. Laravel `throttle`). On limit-exceeded, return `WorkerOverloaded` envelope per `08-error-contract.md` §3.6 with HTTP 429 + `Retry-After` header.

---

*Core API endpoints v1.0.0 — 2026-05-04*
