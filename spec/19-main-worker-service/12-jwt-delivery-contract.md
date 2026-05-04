# 12 — JWT Delivery Contract (XSS-Safe)

**Spec:** `19-main-worker-service`
**Version:** 1.0.0
**Created:** 2026-05-04
**Status:** Authoritative
**Resolves:** audit findings F-A-12, F-D-04, F-B-05 (top-10 fix #3). Closes AC-4.
**Authority:** Canonical contract for how JWTs travel between Main, the React UI, and Workers. On any conflict with `05-auth-and-2fa.md`, diagram files, or other prose, **this file wins**.

---

## 1. The two tokens (do not confuse)

| Token | Carrier | Audience | Storage |
|---|---|---|---|
| **Main session cookie** | `HTTPOnly`, `Secure`, `SameSite=Lax` cookie set by Main | React UI ↔ Main | Browser cookie jar (no JS access) |
| **Worker JWT (RS256)** | JSON response body field `WorkerJwt` | React UI ↔ Worker | **In-memory only** (React state) |

The Main session cookie is **not a JWT**. The Worker JWT is **not delivered as a cookie**. This separation is intentional — see §3.

---

## 2. Authoritative delivery decision

### Worker JWT delivery: **JSON response body, in-memory storage**

**Decision:** The Worker JWT is returned in the **JSON body** of `GET /API/V1/Company/{CompanySlug}/Resolve` (per `06-core-api-endpoints.md` §3.2). React stores it **in memory only** (closure / `useRef` / Zustand non-persistent store). It is **never** written to `localStorage`, `sessionStorage`, IndexedDB, or any cookie.

**Decision is binding.** Implementers MUST NOT:

- ❌ Place the Worker JWT in a cookie (any flags).
- ❌ Persist it across page reloads.
- ❌ Pass it via URL query string.
- ❌ Store it in `localStorage`/`sessionStorage`.

After page reload, React calls `/API/V1/Company/{CompanySlug}/Resolve` again (cheap; uses Main session cookie) to get a fresh JWT.

---

## 3. Why not cookies for the Worker JWT?

| Concern | Cookie delivery | JSON-body + in-memory delivery (chosen) |
|---|---|---|
| XSS exfiltration | HTTPOnly blocks JS reads ✅ | JS can read it ⚠️ |
| CSRF | Auto-sent by browser ⚠️ | Never auto-sent ✅ |
| Cross-domain (React → Worker on different host) | Requires `SameSite=None; Secure` + CORS credentials ⚠️ | Just an `Authorization` header ✅ |
| Re-issue cost | Sticky until expiry ⚠️ | Trivial (refetch on resolve) ✅ |
| Survives page reload | Yes ⚠️ | No ✅ (forces re-resolve, re-checks tenant→worker mapping) |

**Trade-off accepted:** XSS risk is mitigated by (a) strict CSP per §5, (b) short TTL (15 min default per `15-tunable-constants.md` §2.4 `MainWorker.Auth.WorkerJwtTtlSeconds`), and (c) JWT being scoped to one `WorkerNodeId` + one company. CSRF is the larger systemic risk for cookie-borne API tokens, so we eliminate it.

---

## 4. End-to-end token flow

```
Browser (React)              Main                       Worker
     |                         |                          |
1. POST /Auth/SignIn --------->|                          |
2. <-- 200 + Set-Cookie:       |                          |
       MainSession=...; HTTPOnly; Secure; SameSite=Lax    |
3. GET /Company/{slug}/Resolve >|                          |
   (browser auto-attaches      |                          |
    MainSession cookie)        |                          |
4. <-- 200 JSON                |                          |
   { WorkerEndpoint, WorkerJwt, JwtExpiresAt, ... }       |
5. React stores WorkerJwt in   |                          |
   useRef / Zustand (memory)   |                          |
6. GET /Company/{slug} ----------------------->|          |
   Authorization: Bearer <WorkerJwt>           |          |
   X-Correlation-Id: ...                       |          |
                                               |--verify  |
                                               |  RS256   |
                                               |  + claim |
                                               |  wnk     |
   <-----------------200 JSON-------------------|          |
```

---

## 5. Mandatory Content Security Policy

To make the in-memory storage genuinely safer, both Main and Worker tiers MUST serve a CSP that blocks the common XSS exfiltration vectors:

```
Content-Security-Policy:
  default-src 'self';
  script-src 'self';
  connect-src 'self' https://*.example.com;
  frame-ancestors 'none';
  base-uri 'none';
  object-src 'none';
  form-action 'self';
```

`connect-src` MUST list the worker host pattern explicitly. No `unsafe-inline`, no `unsafe-eval`. Hashes / nonces required for any inline `<script>`.

This is non-negotiable — without CSP, in-memory storage gives no real XSS advantage over `localStorage`.

---

## 6. Refresh & rotation

| Event | React behavior |
|---|---|
| JWT within `MainWorker.Auth.JwtRefreshLeadSeconds` (default 60 s per `15-tunable-constants.md` §2.4) of `JwtExpiresAt` | Call `POST /API/V1/Auth/RefreshWorkerToken` (uses Main session cookie); replace in-memory token. |
| Page reload / new tab | In-memory token is gone; React calls `/Company/{slug}/Resolve` to get a fresh one. |
| Main session cookie expires | All Worker JWTs become unrefreshable; React redirects to `/sign-in`. |
| Worker `kid` mismatch (signing key rotated) | Worker returns `WORKER-100-02 KID_UNKNOWN`; React calls `/Resolve` to get a JWT signed by the new key. |
| User clicks "Sign out" | Call `/Auth/SignOut` (clears Main cookie); React drops in-memory JWT. |
| User clicks "Sign out everywhere" | Call `/Auth/SignOutAll`; same as above + invalidates all sessions server-side. |

---

## 7. Worker-side JWT verification (canonical)

Worker MUST verify, in order:

1. RS256 signature against `JwtPublicKeyPem` from bootstrap (`10-worker-bootstrap-protocol.md` §3.2).
2. `kid` matches stored `JwtSigningKeyId`. Mismatch → `WORKER-100-02 KID_UNKNOWN` (401).
3. `exp` is in the future, allowing 60 s clock skew.
4. `wnk` claim equals this worker's `WorkerNodeId`. Mismatch → `WORKER-100-03 WRONG_WORKER` (403).
5. `iss` claim equals the configured `MainBaseUrl` host.
6. `aud` claim equals `worker`.

Any failure → registered error code (per `spec/03-error-manage/03-error-code-registry/`); never silently fall through.

---

## 8. JWT claim contract (canonical)

```jsonc
{
  "iss": "main.example.com",       // Main host
  "aud": "worker",                  // fixed string
  "sub": 42,                        // UserId
  "wnk": 3,                         // WorkerNodeId
  "cmp": "riseup-asia",             // CompanySlug
  "rls": ["PowerAdmin"],            // role codes
  "iat": 1714824000,                // RFC 7519 epoch
  "exp": 1714824900,                // iat + 900s default
  "jti": "01J...ULID",              // unique per token
  "kid": "main-jwt-2026-q2"         // matches JwtSigningKeyId
}
```

No PII beyond `UserId`. No password hashes. No 2FA secrets.

---

## 9. Test cases (mandatory for CI)

| # | Scenario | Expected |
|---|---|---|
| T-1 | Worker JWT placed in `localStorage` by React | Lint/test fail |
| T-2 | Worker JWT in cookie | Lint/test fail (no `Set-Cookie` from Main carries `WorkerJwt`) |
| T-3 | Page reload retains in-memory JWT | Fail (must be gone) |
| T-4 | CSP missing or contains `unsafe-inline` | Fail |
| T-5 | JWT TTL > 900 s | Fail |
| T-6 | Worker accepts JWT with wrong `wnk` | Fail (must return 403) |
| T-7 | Worker accepts JWT with expired `exp` | Fail (must return 401) |
| T-8 | URL contains `?token=` or `?jwt=` | Fail |
| T-9 | JWT logged in any access log | Fail (must be redacted) |

These tests live in `spec/19-main-worker-service/97-acceptance-criteria.md` AC-4.

---

## 10. Diagram corrections (apply during next pass)

| Diagram | Stale element | Replace with |
|---|---|---|
| `seq-login-routing.mmd` | "JWT in JSON body" with no XSS note | Add caption: "JWT delivered in JSON body; React stores in-memory only per `12-jwt-delivery-contract.md`. CSP per §5 mandatory." |
| `seq-company-creation.mmd` | implicit JWT delivery | Add same caption. |

Tracked as follow-up FU-7 / FU-8 (extends §8 list of `11-split-db-tier-reconciliation.md`).

---

## 11. Open Questions (logged, non-blocking)

- **OQ-12-1** Should refresh use a separate `RefreshToken` cookie (rotation pattern) instead of relying on Main session cookie? Inferred: defer to v2.0; current model is simpler and the Main session cookie already has all the necessary properties.
- **OQ-12-2** Service-worker assisted token storage (off-thread)? Inferred: out of scope for v1.0 — adds attack surface.

---

## 12. Cross-references

- `spec/19-main-worker-service/05-auth-and-2fa.md` — high-level auth contract.
- `spec/19-main-worker-service/06-core-api-endpoints.md` §3.2 — `/Resolve` response shape.
- `spec/19-main-worker-service/10-worker-bootstrap-protocol.md` §3.2 — `JwtPublicKeyPem` + `JwtSigningKeyId` source.
- `spec/03-error-manage/03-error-code-registry/` — MUST register `WORKER-100-02 KID_UNKNOWN`, `WORKER-100-03 WRONG_WORKER`.

---

*JWT delivery contract v1.0.0 — 2026-05-04*
