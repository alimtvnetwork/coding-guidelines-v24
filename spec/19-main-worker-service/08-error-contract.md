# 08 — Main↔Worker Error Contract

**Spec:** `19-main-worker-service`
**Version:** 1.1.0

This file defines **only** the error patterns specific to Main↔Worker communication. Generic error rules (catch-log-rethrow, log levels, never-swallow) live in `spec/03-error-manage/` and are inherited verbatim — do NOT duplicate them here.

---

## 1. Inherited Rules (from `spec/03-error-manage/`)

By reference, every handler in this spec MUST:
1. Catch → log → rethrow or handle. Never silent.
2. Log level matches severity (`debug`/`info`/`warn`/`error`/`fatal`).
3. Include explicit file path + operation name in log context (per `mem://architecture/error-handling`).
4. Use the `apperror` package equivalent in the implementer's stack.
5. Never `String(caught)` — pass the raw error object through.

If a rule conflicts with this file, **`spec/03-error-manage/` wins**.

---

## 2. Error Envelope (cross-tier wire format)

Every Main↔Worker error response uses this JSON shape:

```json
{
  "EnvelopeVersion": "1.1.0",
  "Error": {
    "ErrorCode": "WorkerUnreachable",
    "ErrorMessage": "Worker w3.example.com did not respond within 30s",
    "ErrorCategory": "Transport",
    "ErrorSeverity": "Error",
    "CorrelationId": "9f2e1c8a-...",
    "WorkerNodeId": 3,
    "OperationName": "Company.Create",
    "OccurredAt": "2026-05-04T10:22:14Z",
    "Retryable": true,
    "RetryAfterSeconds": 5,
    "OperationId": null,
    "SubCode": null,
    "FieldErrors": null
  }
}
```

Field rules (core):
- `EnvelopeVersion` — SemVer of the envelope schema itself. Bump on additive change. Consumers MUST tolerate unknown fields and MUST refuse to parse a major-version mismatch.
- `ErrorCode` — values from §3 catalog only. PascalCase.
- `ErrorCategory` — `Transport` | `Auth` | `Validation` | `Business` | `Storage` | `Configuration`.
- `ErrorSeverity` — `Warn` | `Error` | `Fatal`.
- `CorrelationId` — echo of inbound `X-Correlation-Id`.
- `Retryable` — boolean. Drives Main's retry decision.
- `RetryAfterSeconds` — present only when `Retryable=true`.

Extension fields (always present in the JSON shape, `null` when not applicable):
- `OperationId` (string, nullable) — set by `IdempotencyConflict` (§3.7) so the caller can reconcile against the original successful response. PascalCase UUID v4 string. (Resolves F-A-12.)
- `SubCode` (string, nullable) — set by `SplitDBWriteFail` (§3.3) and any future error that needs a discriminator under a single `ErrorCode`. PascalCase enum from the per-code sub-code list. (Resolves F-A-15.)
- `FieldErrors` (array, nullable) — set by `ValidationFail` (§3.8); each item `{FieldName: string, FailureReason: string}`. (Resolves F-A-16.)

---

## 3. Failure Taxonomy (the catalog)

### 3.1 `WorkerUnreachable`
- **Category:** Transport. **Severity:** Error. **Retryable:** true.
- **When:** TCP connect fail, TLS fail, request timeout, DNS fail.
- **Main does:** retry per `15-tunable-constants.md` §2.1 (`RetryMaxAttempts`, `RetryBackoffSeconds`, `RetryJitterPct`). Log each attempt. Final failure → return envelope to caller, mark Worker `LastSeenAt` stale.
- **Worker does:** N/A (Worker never sees it; it's a Main-side observation).

### 3.2 `WorkerVersionMismatch`
- **Category:** Configuration. **Severity:** Error. **Retryable:** false.
- **When:** Worker reports a version Main doesn't expect (e.g. an in-flight push update half-applied).
- **Main does:** stop routing to that Worker, mark `Quarantined`, alert Power Admin. Do NOT retry.
- **Why no retry:** retrying a version-skewed worker risks data corruption.

### 3.3 `SplitDBWriteFail`
- **Category:** Storage. **Severity:** Error. **Retryable:** depends on `SubCode`.
- **Sub-codes** (in `Error.SubCode` extension field):
  - `SplitDBLocked` — retryable (SQLite WAL contention). Backoff 50ms, max 5 attempts <!-- TUNABLE-WAIVER: SQLite-WAL-local micro-retry; distinct from Main↔Worker HTTP retries pinned in 15-tunable-constants.md §2.1 -->.
  - `SplitDBDiskFull` — NOT retryable. Page Power Admin.
  - `SplitDBSchemaDrift` — NOT retryable. Worker is on an older schema; quarantine.
- **Worker does:** wrap the underlying `apperror` per `spec/03-error-manage/`, then map to this envelope at the API boundary. Never expose raw SQLite errors over the wire.

### 3.4 `AuthHandshakeFail`
- **Category:** Auth. **Severity:** Error. **Retryable:** false.
- **When:** JWT signature invalid, claims mismatch, OAuth client-credentials rejected, expired token Main thought was fresh.
- **Worker does:** return 401 with this envelope. NEVER 500.
- **Main does:** if Main initiated the call, refresh credentials and retry **once**. If still fails, surface to caller. If user-initiated, force re-login.

### 3.5 `AccessDenied`
- **Category:** Auth. **Severity:** Warn. **Retryable:** false.
- **When:** JWT valid but `User has access to {EnumPage}` returns false.
- **Worker does:** return 403 with this envelope. Write `AccessDenialEvent` row (audit).
- **Main does:** propagate to caller verbatim. Do NOT retry.

### 3.6 `WorkerOverloaded`
- **Category:** Transport. **Severity:** Warn. **Retryable:** true.
- **When:** Worker rate-limit middleware fires, or Worker reports `WorkerNodeStatus = Draining`.
- **Main does:** for new tenants — re-run worker selection (excluding this worker). For existing tenants — retry with `RetryAfterSeconds` honored.

### 3.7 `IdempotencyConflict`
- **Category:** Validation. **Severity:** Warn. **Retryable:** false.
- **When:** `X-Idempotency-Key` reused with a different request body.
- **Worker does:** return 409 with the original response's `OperationId` so the caller can reconcile.

### 3.8 `ValidationFail`
- **Category:** Validation. **Severity:** Warn. **Retryable:** false.
- **When:** payload fails schema or business validation on Worker.
- **Includes:** `Error.FieldErrors` array with `{FieldName, FailureReason}`.

---

## 4. Correlation ID Propagation

Mandatory chain on every cross-tier hop:

```
React  --X-Correlation-Id: <uuid> -->  Main
Main   --X-Correlation-Id: <same uuid> -->  Worker
Worker --(in response Error.CorrelationId, in logs as cid=...)-->  Main
Main   --(in response Error.CorrelationId)-->  React
```

Rules:
- If inbound request lacks `X-Correlation-Id`, generate a UUID v4 at the edge (Main).
- Workers MUST NOT generate fresh IDs; they propagate Main's.
- Every log line related to the request includes `cid=<uuid>`.
- IDs are opaque — never used as a DB primary key (PKs stay `INTEGER AUTOINCREMENT` per `spec/04-database-conventions/`).

---

## 5. Retry Semantics (single source of truth)

Main applies retries ONLY when the response is `Retryable=true` AND the request is **safe to retry**:

| HTTP method | Retryable by default? |
|-------------|------------------------|
| GET | yes |
| HEAD | yes |
| PUT | yes (idempotent by spec) |
| DELETE | yes (idempotent) |
| POST | only if `X-Idempotency-Key` present |
| PATCH | only if `X-Idempotency-Key` present |

Backoff and attempt budget: see `15-tunable-constants.md` §2.1 (`RetryMaxAttempts`, `RetryBackoffSeconds`, `RetryJitterPct`). Implementations MUST NOT hard-code different defaults.

Pseudocode (CODE RED compliant):

```php
public function callWorkerWithRetry(WorkerCall $call): WorkerResponse
{
    $attempt = 0;
    while ($this->canRetry($attempt, $call)) {
        $response = $this->tryOnce($call);
        if ($this->isSuccessful($response)) { return $response; }
        if ($this->isPermanentFailure($response)) { return $response; }
        $this->sleepBackoff($attempt);
        $attempt++;
    }
    return $this->lastResponse;
}
```

Each helper is its own ≤8-line function. `canRetry`, `isSuccessful`, `isPermanentFailure` are positively named guards (no `!`).

---

## 6. Logging Contract

Every cross-tier error logs:

```
level=error
operation=Company.Create
cid=9f2e1c8a-...
worker_node_id=3
error_code=WorkerUnreachable
attempt=3
elapsed_ms=31204
file=app/Services/WorkerClient.php
caller=callWorkerWithRetry
```

Per `mem://architecture/error-handling`: explicit file path + operation name MANDATORY. No `String(error)`.

---

## 7. What Never Happens

- ❌ Silent retry loop with no log.
- ❌ Returning 500 to React when the cause is a Worker `WorkerUnreachable` (use 502 / 504).
- ❌ Exposing internal exception class names in `ErrorMessage` (info leak).
- ❌ Retrying on `WorkerVersionMismatch`, `AuthHandshakeFail`, `AccessDenied`, `ValidationFail`, `IdempotencyConflict`, `SplitDBDiskFull`, `SplitDBSchemaDrift`.
- ❌ Generating fresh correlation IDs at Worker tier.
- ❌ Catching `Exception` and dropping it on the floor (CODE RED).

---

## 8. Cross-References

- `spec/03-error-manage/` — generic error rules (this file extends, never overrides)
- `01-architecture.md` §4 — Comms contract (transport, auth, headers)
- `04-worker-routing.md` §3 — Failover behavior on `WorkerUnreachable`
- `05-auth-and-2fa.md` §7 — Worker JWT validation (source of `AuthHandshakeFail`)
- `07-role-based-dashboards.md` §8 — `AccessDenialEvent` audit table
- `mem://architecture/error-handling` — `apperror` package + explicit file/path logging

---

*Error contract v1.0.0 — 2026-05-04*
