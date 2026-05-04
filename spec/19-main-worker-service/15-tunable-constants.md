# 15 — Tunable Constants (Single-Value Pins)

**Spec:** `19-main-worker-service`
**Version:** 1.0.0
**Created:** 2026-05-04
**Status:** Authoritative
**Resolves:** audit findings F-A-15, F-A-16, F-B-12 (top-10 fix #7). Closes AC-7, partially AC-6.
**Authority:** Single source of truth for every numeric tunable referenced anywhere in `spec/19-main-worker-service/` or `spec/14-update/28-worker-push-instruction.md`. On any conflict, **this file wins**. All values seedable via `spec/06-seedable-config-architecture/` (key/value `Categories` block, not `Tables`).

---

## 1. Why this file exists

Earlier drafts of spec/19 mentioned "3 retries" in one place and "5 retries" in another (audit F-A-15). Idempotency-key TTL was implied but never named. Heartbeat interval appeared as `30s` in `10-worker-bootstrap-protocol.md` §3.2 but was unmentioned in the routing doc — leaving the dumb-AI implementer free to invent. This file pins every such constant.

Each row below is **the** value. Implementations MAY override via Seedable-Config at install time but MUST start from this default and MUST NOT hard-code different defaults in source.

---

## 2. Canonical constants

### 2.1 Retry / backoff

| Key | Default | Unit | Used by | Notes |
|---|---:|---|---|---|
| `MainWorker.Retry.MaxAttempts` | **3** | count | All Main↔Worker HTTP calls | Resolves F-A-15 (3-vs-5 contradiction). Includes initial attempt. |
| `MainWorker.Retry.BackoffSeconds` | `[2, 8, 30]` | seconds[] | Same | Length MUST equal `MaxAttempts - 1`. Exponential-ish, ceiling 30s. |
| `MainWorker.Retry.JitterPct` | **20** | percent | Same | ±20% applied per attempt. |
| `WorkerPushUpdate.MaxRetries` | **3** | count | `spec/14-update/28` §3.1 `OnFailure.MaxRetries` | Mirrors above; `RetryBackoffSeconds` defaults to `[30, 120, 300]` per §3 of that file. |

### 2.2 Idempotency

| Key | Default | Unit | Used by | Notes |
|---|---:|---|---|---|
| `MainWorker.Idempotency.KeyTtlSeconds` | **86400** (24h) | seconds | Worker + Main idempotency stores | Single canonical TTL. Replays beyond this MAY re-execute. |
| `MainWorker.Idempotency.KeyMaxLength` | **64** | chars | Header validation | ULID = 26; reservation buffer for future formats. |
| `MainWorker.Idempotency.StoreCleanupSeconds` | **3600** (1h) | seconds | Background sweeper | Deletes rows older than `KeyTtlSeconds`. |

### 2.3 Heartbeat & quarantine

| Key | Default | Unit | Used by | Notes |
|---|---:|---|---|---|
| `MainWorker.Heartbeat.IntervalSeconds` | **30** | seconds | `10-worker-bootstrap-protocol.md` §3.2 / §7 | Worker→Main ping cadence. |
| `MainWorker.Heartbeat.MissedThreshold` | **3** | count | `10` §7, `04-worker-routing.md` | Consecutive misses → quarantine. |
| `MainWorker.Heartbeat.QuarantineCooldownSeconds` | **300** (5m) | seconds | `04-worker-routing.md` | Quarantined worker eligible for re-eval after this. |
| `MainWorker.Heartbeat.GraceWindowSeconds` | **5** | seconds | Main-side scheduler | Tolerated jitter before counting a miss. |

### 2.4 Auth & JWT

| Key | Default | Unit | Used by | Notes |
|---|---:|---|---|---|
| `MainWorker.Auth.WorkerJwtTtlSeconds` | **900** (15m) | seconds | `12-jwt-delivery-contract.md` §6 | Already pinned there; mirrored here for the single-table view. |
| `MainWorker.Auth.JwtRefreshLeadSeconds` | **60** | seconds | `12` §6 | React refreshes when within this window of `exp`. |
| `MainWorker.Auth.MainSessionTtlSeconds` | **28800** (8h) | seconds | Main session cookie | Sliding window: each request extends. |
| `MainWorker.Auth.ClockSkewToleranceSeconds` | **60** | seconds | `12` §7, `10` §3 | Same value across both contexts. |

### 2.5 Routing & pool

| Key | Default | Unit | Used by | Notes |
|---|---:|---|---|---|
| `MainWorker.Routing.HttpTimeoutSeconds` | **15** | seconds | Main → Worker request timeout | Distinct from retry attempt budget. |
| `MainWorker.Routing.HttpHandshakeTimeoutSeconds` | **30** | seconds | Push-update POST handshake (per `spec/14-update/28` §2) | Already pinned there. |
| `MainWorker.Routing.MaxConcurrentPerWorker` | **64** | count | Connection pool cap | Prevents stampede. |

### 2.6 Rate limiting (mirrors `06-core-api-endpoints.md` §6)

| Key | Default | Unit | Notes |
|---|---:|---|---|
| `MainWorker.RateLimit.AuthEndpointsPerMinutePerIp` | **10** | rpm | Already pinned in `06` §6. |
| `MainWorker.RateLimit.WorkerEndpointsPerMinutePerToken` | **60** | rpm | Same. |
| `MainWorker.RateLimit.OtherAuthenticatedPerMinutePerUser` | **600** | rpm | Same. |

### 2.7 Self-update execution window (mirrors `spec/14-update/28` §3)

| Key | Default | Unit | Notes |
|---|---:|---|---|
| `WorkerPushUpdate.MaxRunDurationSeconds` | **600** (10m) | seconds | Hard cap before self-abort + rollback. |
| `WorkerPushUpdate.HandoffTimeoutSeconds` | **60** | seconds | Per `spec/14-update/28` §5 step 8. |
| `WorkerPushUpdate.InstructionRetentionDays` | **14** | days | Per `spec/14-update/28` §7. |
| `WorkerPushUpdate.IssuedSkewSeconds` | **300** (5m) | seconds | Per `spec/14-update/28` §3 — Worker rejects an instruction whose `IssuedAtUtc` is older than this. Distinct from JWT clock-skew (§2.4). Bumped to v1.1.0 to retire the §28 line-82 TUNABLE-WAIVER. |

### 2.8 Self-update pointer (consumer: `09-self-update-pointer.md`)

| Key | Default | Unit | Used by | Notes |
|---|---:|---|---|---|
| `MainWorker.SelfUpdate.RedirectStaleHours` | **36** | hours | `09-self-update-pointer.md` §1 step 5 | If the cached redirect URL is older than this OR unreachable, re-resolve via the original endpoint. Bumped to v1.1.0 to retire the §09 line-41 TUNABLE-WAIVER. |

### 2.9 Worker bootstrap (consumer: `10-worker-bootstrap-protocol.md`)

| Key | Default | Unit | Used by | Notes |
|---|---:|---|---|---|
| `MainWorker.Bootstrap.RetryBackoffSeconds` | `[10, 30, 90, 300]` | seconds[] | `10-worker-bootstrap-protocol.md` §6 (`WORKER-100-01 OAUTH_HANDSHAKE_FAIL`) | Cold-bootstrap retry ladder for Worker→Main OAuth handshake; distinct from the steady-state `MainWorker.Retry.BackoffSeconds` of §2.1. After exhaustion the Worker exits and is restarted by its supervisor. Bumped to v1.1.0 to retire the §10 line-137 TUNABLE-WAIVER. |
| `MainWorker.Bootstrap.RetryMaxAttempts` | **4** | count | Same | MUST equal `len(RetryBackoffSeconds)`. |

---

## 3. Single-value rule (for the dumb AI)

| Rule | Why |
|---|---|
| Every numeric tunable referenced in spec/19 or spec/14-update/28 MUST appear in §2. | Eliminates "guess what value the author meant." |
| If a value appears in two prose locations, both MUST cite this file by anchor (e.g. `15-tunable-constants.md` §2.3). | Forces single source of truth. |
| Updating a default = bump `config.seed.json` `Categories.MainWorker` `Version`, edit ONE row in §2, run linter. | No multi-file scavenger hunt. |
| Override at runtime = Seedable-Config Categories block (NOT `Tables`). | Per `spec/06-seedable-config-architecture/01-fundamentals.md`. |

---

## 4. `config.seed.json` Categories binding (paste-ready)

Add (or merge with) the following category at SemVer `1.3.0` of `config.seed.json` (folded into FU-12 from task 12):

```jsonc
"MainWorker": {
  "DisplayName": "Main / Worker tunables",
  "Description": "Single-value tunables for spec/19. See 15-tunable-constants.md.",
  "Version": "1.3.0",
  "AddedIn":  "1.3.0",
  "Settings": {

    "RetryMaxAttempts":            { "Type": "number",  "Default": 3,         "Min": 1,    "Max": 10 },
    "RetryBackoffSeconds":         { "Type": "string",  "Default": "2,8,30",  "Description": "Comma-separated; length = RetryMaxAttempts - 1." },
    "RetryJitterPct":              { "Type": "number",  "Default": 20,        "Min": 0,    "Max": 100 },

    "IdempotencyKeyTtlSeconds":    { "Type": "number",  "Default": 86400,     "Min": 60 },
    "IdempotencyKeyMaxLength":     { "Type": "number",  "Default": 64,        "Min": 26,   "Max": 256 },
    "IdempotencyStoreCleanupSec":  { "Type": "number",  "Default": 3600,      "Min": 60 },

    "HeartbeatIntervalSeconds":    { "Type": "number",  "Default": 30,        "Min": 5,    "Max": 600 },
    "HeartbeatMissedThreshold":    { "Type": "number",  "Default": 3,         "Min": 1,    "Max": 20 },
    "HeartbeatQuarantineCooldown": { "Type": "number",  "Default": 300,       "Min": 30 },
    "HeartbeatGraceWindowSeconds": { "Type": "number",  "Default": 5,         "Min": 0 },

    "WorkerJwtTtlSeconds":         { "Type": "number",  "Default": 900,       "Min": 60,   "Max": 3600 },
    "JwtRefreshLeadSeconds":       { "Type": "number",  "Default": 60,        "Min": 10 },
    "MainSessionTtlSeconds":       { "Type": "number",  "Default": 28800,     "Min": 300 },
    "ClockSkewToleranceSeconds":   { "Type": "number",  "Default": 60,        "Min": 0,    "Max": 300 },

    "RoutingHttpTimeoutSeconds":   { "Type": "number",  "Default": 15,        "Min": 1 },
    "RoutingHandshakeTimeoutSec":  { "Type": "number",  "Default": 30,        "Min": 1 },
    "RoutingMaxConcurrentPerNode": { "Type": "number",  "Default": 64,        "Min": 1 },

    "RateAuthPerMinutePerIp":      { "Type": "number",  "Default": 10 },
    "RateWorkerPerMinutePerToken": { "Type": "number",  "Default": 60 },
    "RateOtherPerMinutePerUser":   { "Type": "number",  "Default": 600 },

    "PushUpdateMaxRunSeconds":     { "Type": "number",  "Default": 600,       "Min": 30 },
    "PushUpdateHandoffTimeoutSec": { "Type": "number",  "Default": 60,        "Min": 5 },
    "PushUpdateRetentionDays":     { "Type": "number",  "Default": 14,        "Min": 1 },
    "PushUpdateIssuedSkewSec":     { "Type": "number",  "Default": 300,       "Min": 30,   "Max": 3600 },

    "SelfUpdateRedirectStaleHours":{ "Type": "number",  "Default": 36,        "Min": 1,    "Max": 720 },

    "BootstrapRetryBackoffSec":    { "Type": "string",  "Default": "10,30,90,300", "Description": "Comma-separated; length = BootstrapRetryMaxAttempts." },
    "BootstrapRetryMaxAttempts":   { "Type": "number",  "Default": 4,         "Min": 1,    "Max": 10 }

  }
}
```

---

## 5. Stale-prose corrections (follow-ups)

The following docs cite tunables inline. Each MUST be edited to cite this file instead. Tracked as FU-14:

| File | Stale phrase | Replace with |
|---|---|---|
| `04-worker-routing.md` | "retry up to 5 times" / "3 retries" | "Retries per `15-tunable-constants.md` §2.1." |
| `08-error-contract.md` | implicit retry budget | Add cite to `15` §2.1. |
| `10-worker-bootstrap-protocol.md` §3.2 | "30 s" hard-coded literal | Replace with: "default `30 s` per `15-tunable-constants.md` §2.3." |
| `spec/14-update/28-worker-push-instruction.md` §3.1 `MaxRetries: 3` | already correct | Add cross-cite to `15` §2.1. |
| `spec/14-update/28-worker-push-instruction.md` §7 `WorkerUpdateInstructionRetentionDays` | scattered | Add cross-cite to `15` §2.7. |

---

## 6. Linter assertion (FU-15)

`linter-scripts/check-tunable-constants.py` MUST verify:

1. Every numeric literal in spec/19 prose that ends with `s`, `sec`, `seconds`, `min`, `minutes`, `h`, `hours`, `attempts`, `times`, or `retries` is either
   (a) named in §2 of this file, or
   (b) explicitly waivered via `<!-- TUNABLE-WAIVER: rationale -->` comment.
2. No two §2 rows share the same Key.
3. `config.seed.json` `Categories.MainWorker.Settings.*.Default` matches §4 verbatim.

Failure = build break.

---

## 7. Open Questions (logged, non-blocking)

- **OQ-15-1** Should retry backoff be true exponential (`base^n`) instead of explicit array? Inferred: explicit array is dumb-AI-friendlier and bounds the worst case.
- **OQ-15-2** Should `MainSessionTtlSeconds` be sliding or absolute? Inferred: sliding, matches Laravel Sanctum default; absolute reserved for v2.0 if compliance demands.

---

## 8. Cross-references

- `spec/06-seedable-config-architecture/01-fundamentals.md` — `Categories` block schema (§4 conforms).
- `spec/19-main-worker-service/04-worker-routing.md` — consumer (FU-14).
- `spec/19-main-worker-service/10-worker-bootstrap-protocol.md` §3.2/§7 — consumer (FU-14).
- `spec/19-main-worker-service/12-jwt-delivery-contract.md` §6 — consumer (already cites tunables but pinned here).
- `spec/14-update/28-worker-push-instruction.md` §3/§7 — consumer (FU-14).
- `spec/19-main-worker-service/13-error-codes.md` — `WORKER-300-01` triggered when `IdempotencyKeyTtlSeconds` exceeded.

---

*Tunable constants v1.0.0 — 2026-05-04*
