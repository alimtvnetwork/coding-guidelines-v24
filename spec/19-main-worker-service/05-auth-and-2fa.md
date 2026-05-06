# 05 — Authentication and 2FA

**Spec:** `19-main-worker-service`
**Version:** 2.0.0

> **v2.0.0 (Phase 3 — Users moved off Main).** Main is now a **credential-blind reverse proxy** for `/Auth/*` traffic. It owns the routing index `UserDirectory` (`03-main-db-schema.md` §2.4) and nothing else. Password hashes, TOTP secrets, backup codes, and role assignments live exclusively on the assigned Worker's split-DB App tier (`AppUser`, `AppUserRole`). The flows in §5–§7 below are rewritten accordingly. The previous v1.0 flow — where Main verified passwords locally — is **removed**.

Auth is a **first-class given** in BOTH Main and Worker tiers. This file defines the contract; implementer chooses Laravel Sanctum / Passport / custom JWT as long as the contract is honored.

---

## 1. Capability matrix (v2.0.0)

| Capability | Main | Worker |
|-----------|------|--------|
| Email + password sign-up (HTTP entry point) | ✅ proxy | ✅ authoritative |
| Email + password sign-in (HTTP entry point) | ✅ proxy | ✅ authoritative |
| Password hashing / verification | ❌ never | ✅ |
| TOTP enroll + verify | ❌ never | ✅ |
| TOTP backup-code storage | ❌ never | ✅ |
| `UserDirectory` routing index | ✅ | ❌ (mirrored read-only via bootstrap) |
| Session management (UI cookie) | ✅ | optional |
| Worker-JWT issuance | ❌ | ✅ (Worker mints; Main forwards) |
| Cookie-based session for React UI | ✅ | optional |
| Password reset (email link) | ✅ proxy | ✅ authoritative |
| Sign-out (single + all sessions) | ✅ | ✅ |

> **Why Main proxies instead of authenticates.** Per locked decision D5, Main MUST be a thin catalog. Storing password hashes on Main would re-introduce the cross-tenant blast radius the split-DB architecture exists to eliminate. The proxy pattern keeps Main credential-blind while preserving a single public entry URL.

---

## 2. Two Authentication Surfaces

### 2.1 User → Main → Worker (UI-facing, credential proxy)

- **Entry:** browser POSTs `/API/V1/Auth/SignIn` to Main with `{ Email, Password, TotpCode? }`.
- **Main step 1 — routing lookup:** Main reads `UserDirectory WHERE UserEmail = LOWER(:email)` to obtain `WorkerNodeId`. **Constant-time response on miss** (no email enumeration): if no row, Main forwards to a synthetic "null Worker" handler that returns the same generic 401 envelope after the same wall-clock budget as a real verification.
- **Main step 2 — forward:** Main proxies the original body verbatim over a mutual-TLS internal channel to `POST {WorkerEndpoint}/API/V1/Auth/InternalSignIn` with headers `X-Forwarded-For-User: <hash(email)>`, `X-Correlation-Id`, and an OAuth client-credentials Bearer (per §2.3). **Main does NOT log the password** (scrubbed per §3).
- **Worker step:** Worker reads its `AppUser` row, runs the password verifier (Argon2id / bcrypt per §3), checks TOTP if enrolled, then mints the Worker-JWT (per `12-jwt-delivery-contract.md`).
- **Main step 3 — session cookie:** On Worker 200, Main stamps a session cookie (HTTPOnly, Secure, SameSite=Lax) bound to the Worker-JWT and returns the JWT body to the browser. On Worker 401, Main returns 401 unchanged.
- **Main never sees the cleartext password after forwarding.** The proxied body buffer MUST be zeroed (`sodium_memzero` or equivalent) immediately after the forward call returns.

### 2.2 React → Worker (data-facing, after sign-in)

- **Mechanism:** short-lived JWT minted by **the Worker** (not Main, as in v1.0), accepted by that same Worker. JWT claims unchanged from v1.0 except `iss = WorkerEndpoint` (was Main URL).
- **JWT claims:**
  - `sub` = `AppUserId` (Worker-local)
  - `cmp` = `CompanyId`
  - `wnk` = `WorkerNodeId` (so Worker rejects misrouted tokens)
  - `iss` = Worker URL
  - `aud` = Worker URL (self-issued)
  - `exp` = issued + `MainWorker.Auth.WorkerJwtTtlSeconds` (default per `15-tunable-constants.md` §2.4 = 15 min)
  - `iat` = epoch
  - `roles` = array of `RoleCode` strings (Worker computed cascading union per Phase 5)
- **Signing:** asymmetric (RS256). Worker holds private key; Main holds the Worker's public key (rotatable via Seedable-Config) so it can validate JWTs on session-cookie refresh.
- **Refresh:** React calls Worker `/API/V1/Auth/RefreshToken` directly (Main bypassed) when JWT is within `MainWorker.Auth.JwtRefreshLeadSeconds` of expiry (default per `15-tunable-constants.md` §2.4 = 60 s).

### 2.3 Main → Worker (orchestration: push-update, registry sync, credential proxy)
- **Mechanism:** OAuth 2.0 client-credentials grant OR pre-shared API key (configurable).
- **Default:** OAuth client-credentials per Worker, secrets stored via Seedable-Config (encrypted at rest).
- The credential-proxy channel (§2.1) reuses this same client-credential token; the Worker enforces that `/API/V1/Auth/InternalSignIn` is callable **only** from Main's IP allowlist + valid Bearer.
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
- **Backup codes:** generate 10 single-use codes at enrollment (stored as bcrypt hashes in `User.UserTotpBackupCodesHash` per `03-main-db-schema.md` §2.4). When the count of unused codes reaches **0**, the user MUST be forced to regenerate at next sign-in: Main returns `Error.SubCode = TotpBackupExhausted` with HTTP 403 and `X-Auth-Action: RegenerateBackupCodes` header (resolves F-A-05). Regeneration invalidates the prior batch. Power Admin override path: `POST /API/V1/Auth/2FA/RegenerateBackupCodes` (audit logged).
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

> ✅ **Open Question OQ-1 — RESOLVED 2026-05-04 (task #39).**
> Per-endpoint auth-mechanism overrides ARE supported. Contract: single-row whole-replace `PATCH /API/V1/Settings/EndpointAuth` keyed by `EndpointPathPattern`, with `AcceptedMechanisms[]` allow-list, 7 validation rules, lock-list for `Workers/*` + `SelfUpdate`, and idempotent re-application. Full schema + semantics in `06-core-api-endpoints.md` §5. Every successful PATCH writes one `EndpointAuthAuditEvent` row (`03-main-db-schema.md` §2.6.4) inside the same transaction (FU-17 — RESOLVED 2026-05-05). Closes audit finding F-M-10.

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
