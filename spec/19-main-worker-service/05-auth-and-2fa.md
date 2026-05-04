# 05 — Authentication and 2FA

**Spec:** `19-main-worker-service`
**Version:** 1.0.0

Auth is a **first-class given** in BOTH Main and Worker tiers. This file defines the contract; implementer chooses Laravel Sanctum / Passport / custom JWT as long as the contract is honored.

---

## 1. What both tiers MUST ship with

| Capability | Main | Worker |
|-----------|------|--------|
| Email + password sign-up | ✅ | ✅ |
| Email + password sign-in | ✅ | ✅ |
| 2FA (TOTP) enroll + verify | ✅ | ✅ |
| Session management | ✅ | ✅ |
| JWT issuance | ✅ | ✅ |
| Cookie-based session | ✅ (for React UI) | optional |
| Password reset (email link) | ✅ | ✅ |
| Sign-out (single + all sessions) | ✅ | ✅ |

Worker has these even though it has no UI — they're required for service-to-service auth and direct-from-React calls after Main resolves the worker.

---

## 2. Two Authentication Surfaces

### 2.1 User → Main (UI-facing)
- **Mechanism:** session cookie (HTTPOnly, Secure, SameSite=Lax) issued by Main.
- **2FA gate:** if `User.Has2FAEnabled = true`, Main blocks issuance until TOTP code verified.
- **Token:** opaque session ID, mapped server-side. NOT a JWT (cookies don't need to carry payload).

### 2.2 React → Worker (data-facing, after Main resolves the worker)
- **Mechanism:** short-lived JWT minted by Main, accepted by the resolved Worker.
- **JWT claims:**
  - `sub` = `UserId`
  - `cmp` = `CompanyId`
  - `wnk` = `WorkerNodeId` (so Worker rejects misrouted tokens)
  - `iss` = Main's URL
  - `aud` = Worker's URL
  - `exp` = issued + `MainWorker.Auth.WorkerJwtTtlSeconds` (default per `15-tunable-constants.md` §2.4 = 15 min)
  - `iat` = epoch
  - `roles` = array of `RoleCode` strings
- **Signing:** asymmetric (RS256). Main holds private key; each Worker holds Main's public key (rotatable via Seedable-Config).
- **Refresh:** React calls Main `/API/V1/Auth/RefreshWorkerToken` when JWT is within `MainWorker.Auth.JwtRefreshLeadSeconds` of expiry (default per `15-tunable-constants.md` §2.4 = 60 s).

### 2.3 Main → Worker (orchestration: push-update, registry sync)
- **Mechanism:** OAuth 2.0 client-credentials grant OR pre-shared API key (configurable).
- **Default:** OAuth client-credentials per Worker, secrets stored via Seedable-Config (encrypted at rest).
- Per-endpoint flexibility per verbatim §Main Server Concept 3c — see Open Question OQ-1 below.

---

## 3. Password Storage

Per verbatim §Login 3:

- **Hash:** Argon2id (preferred) or bcrypt. Bcrypt cost is **environment-pinned** to remove ambiguity (resolves F-A-03): `MainWorker.Auth.BcryptCost` defaults to `12` when `Env=dev|test`, `14` when `Env=prod|staging`. Implementations MUST refuse cost < 12 and MUST NOT exceed 14 unless overridden by Power Admin via Seedable-Config (caps at 16).
- **Salt:** unique per user, stored alongside hash in `UserPasswordSalt`. The chosen hash function may also embed a salt; storing it explicitly keeps the contract stack-portable.
- **Pepper:** global pepper from Seedable-Config secret `MainWorker.Auth.PasswordPepper`. **MUST be set when `Env=prod|staging`** (resolves F-A-04 — was previously "optional"). In `dev|test` MAY be empty; if empty, Main MUST log a one-shot WARN at startup so drift is visible. When present, mixed in before hashing.
- **No plaintext anywhere.** Logs MUST scrub `password`, `confirmPassword`, `currentPassword`.
- **No retrieval.** Reset is replace-only.
- **Breach check:** Implementations MUST verify against a HIBP-style API on sign-up and password change when `MainWorker.Auth.EnableBreachCheck=true` (default `true` in prod, `false` in dev).

---

## 4. Two-Factor Authentication (2FA)

- **Standard:** TOTP (RFC 6238), 30s window <!-- TUNABLE-WAIVER: RFC 6238 mandates 30s; not a MainWorker tunable -->, 6 digits.
- **Enrollment:** Main shows QR (otpauth URI), user scans, submits one TOTP to confirm. On success, store `User.TotpSecret` (encrypted at rest with key from Seedable-Config).
- **Backup codes:** generate 10 single-use codes at enrollment. Store hashed.
- **Verification points:** sign-in, password change, 2FA disable, role escalation.
- **Recovery:** Power Admin can reset 2FA for any user (audit logged).

---

## 5. Sign-Up Flow (Main)

1. POST `/API/V1/Auth/SignUp` with email + password + (optional) `CompanySlug`.
2. Main runs guards: `IsEmailWellFormed`, `IsPasswordStrongEnough`, `IsCompanySlugAvailable`.
3. Main creates `User` row in main DB (auth fields only).
4. If new company: Main runs worker selection (`04-worker-routing.md`) and forwards full company payload to chosen Worker.
5. Main returns 201 + sets session cookie (or returns "verify email" status if email-confirm flag is on).

---

## 6. Sign-In Flow (Main)

1. POST `/API/V1/Auth/SignIn` with email + password.
2. Main verifies hash, checks `User.Has2FAEnabled`.
3. If 2FA on: return 200 with `{ "Step": "AwaitTotp", "ChallengeId": "..." }`. Client POSTs `/API/V1/Auth/Verify2FA` with the TOTP code + ChallengeId. Main grants session on success.
4. Main resolves `User → Company → WorkerNode`, mints worker JWT, returns:
   ```json
   {
     "WorkerEndpoint": "https://w3.example.com",
     "WorkerJwt": "<RS256 token>",
     "JwtExpiresAt": "2026-05-04T12:15:00Z"
   }
   ```
5. React stores JWT in memory (NOT localStorage), uses it for direct Worker calls.

---

## 7. Worker JWT Validation (Worker side)

Every Worker request validates:
1. Signature against Main's public key.
2. `exp` not passed.
3. `aud` matches this Worker's URL.
4. `wnk` claim matches this Worker's `WorkerNodeId`.
5. `cmp` claim matches the `CompanyId` resolved from the requested resource.

Failure → 401 with `08-error-contract.md` envelope. NEVER 500 on auth failure.

---

## 8. Endpoint Authentication Defaults

| Endpoint pattern | Default auth | Configurable? |
|------------------|--------------|----------------|
| `/API/V1/Auth/*` | None (sign-up/in) or session (sign-out) | No |
| `/API/V1/Status` | None | Yes (admin can require auth) |
| `/API/V1/Version` | None | Yes |
| `/API/V1/Company/*` | Session (Main) or Worker JWT (Worker) | Per-endpoint via Seedable-Config |
| `/API/V1/Workers/*` | Session + `User has access to EnumPage.PowerAdminPage` | No (always protected) |
| `/API/V1/SelfUpdate` | OAuth client-credentials | No (always protected) |

> ❓ **Open Question OQ-1 (from verbatim §Main Server Concept 3c):**
> Whether and how to allow per-endpoint auth-mechanism overrides (e.g., let one endpoint accept basic auth while another requires JWT). Settings table sketch is in `06-core-api-endpoints.md` §Settings; final design deferred to a follow-up spec revision.

---

## 9. Anti-patterns (CODE RED)

- ❌ `if user.role === 'admin'` — use `User has access to EnumPage.X` (`07-role-based-dashboards.md`).
- ❌ Storing JWTs in `localStorage` (XSS exposure).
- ❌ Long-lived worker JWTs (> 1 hour) <!-- TUNABLE-WAIVER: anti-pattern threshold, not a tunable; canonical TTL is MainWorker.Auth.WorkerJwtTtlSeconds in 15-tunable-constants.md §2.4 -->.
- ❌ Symmetric JWT signing across tiers (key sharing risk).
- ❌ Returning 500 on bad credentials (info leak; use 401 + neutral message).
- ❌ `if (!isAuthenticated)` — invert to `if (isAuthenticated)` and use early-return guards.

---

*Auth and 2FA v1.0.0 — 2026-05-04*
