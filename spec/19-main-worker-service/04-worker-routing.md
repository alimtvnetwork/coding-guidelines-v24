# 04 — Worker Routing

**Spec:** `19-main-worker-service`
**Version:** 1.2.0

> **v1.2.0 (Phase 4 — Backup nodes excluded; "Region" UI label):** The eligibility filter (§1.4) gains a positive guard `IsPrimary(node)` which requires `WorkerNode.IsBackup = 0`. Backup nodes (D8/D9) MUST NEVER appear in any selection strategy's candidate set, including `Manual` (a Power Admin attempting to route to a backup gets `WORKER-300-04 BackupNotRoutable`). RoundRobin cursor walks `WorkerNode.Sequence ASC` over `IsBackup = 0` rows only. The dashboard column previously labelled "Worker" is renamed to **"Region"** in UI copy; the underlying field name stays `WorkerNode` everywhere in code and APIs.

How the Main Server picks which Worker handles a new tenant, caches that decision, and recovers when a Worker fails.

The author's mindmap below frames routing as two questions the Main Server keeps asking: **"Give me new"** (which worker is most free, what's the endpoint of communication) and **"Knowledge"** (total workers, machine/website inventory, who is free). The strategies in §1 are the formal answers.

![EndPointService → WorkerPattern](./images/02-endpoint-service-worker-pattern.png)

See also [`images/03-worker-subdomain-routing.png`](./images/03-worker-subdomain-routing.png) for the per-tenant `wN.<domain>` subdomain layout this routing produces.

---

## 1. Selection Strategies

Strategy is configurable via Seedable-Config key `MainWorker.Routing.DefaultStrategy` (canonical default in `15-tunable-constants.md` §2.5). Stored in `WorkerSelectionStrategy` table. Allowed values: `RoundRobin`, `LeastLoaded`, `Manual`. Main MUST refuse to start when the configured value is outside this allow-list (no silent fallback — CODE RED).

### 1.1 `RoundRobin`
- Pick next eligible worker ordered by `WorkerNode.Sequence ASC` (ties broken by `WorkerNodeRegisteredAt ASC`).
- Cursor persisted in main DB (single-row config table or `WorkerSelectionEvent` last-row lookup).
- Pros: trivially predictable, deterministic across restarts. Cons: ignores load.

### 1.2 `LeastLoaded` (**default** — resolves OQ-2)
- Pick `Active` worker with fewest assigned `Company` rows.
- Tiebreaker: oldest `WorkerNodeRegisteredAt`.
- Pros: balances over time. Cons: slightly more expensive query (still O(N) on workers, N is small).

### 1.3 `Manual`
- Power Admin specifies `WorkerNodeId` in the create request.
- Used for testing and reserved-capacity tenants.
- Requires `User has access to EnumPage.PowerAdminPage`.

### 1.4 Eligibility filter (applies to all strategies)
A worker is eligible only if **all** are true (positive guards, per CODE RED):
- `IsPrimary(node)` → `WorkerNode.IsBackup = 0` (Phase 4, D9 — backups never serve traffic).
- `IsWorkerActive(node)` → `WorkerNodeStatusCode = 'Active'`
- `IsWorkerReachable(node)` → last heartbeat within `MainWorker.Heartbeat.IntervalSeconds × MissedThreshold` per `15-tunable-constants.md` §2.3 <!-- TUNABLE-WAIVER: derived product, not a literal -->.
- `HasCapacity(node)` → assigned company count strictly less than `MainWorker.Routing.MaxCompaniesPerWorker`. **NULL = unlimited** (resolves F-A-06; the legacy `0` magic value is rejected — Main MUST refuse to start if the configured value is `0` or negative). When `NULL`, the guard returns `true` unconditionally.

If no eligible worker exists → return `WorkerUnavailable` error per `08-error-contract.md`.

### 1.5 Default selection rationale (resolves OQ-2)

**Decision (Phase 12.3):** `MainWorker.Routing.DefaultStrategy = LeastLoaded`. Pinned in `15-tunable-constants.md` §2.5.

**Why `LeastLoaded` over `RoundRobin`:**

| Criterion | `RoundRobin` | `LeastLoaded` (chosen) | `Manual` |
|---|---|---|---|
| Cold-cluster fairness (workers added at different times) | ❌ Newly added workers stay near-empty until the cursor wraps a full cycle | ✅ New workers get traffic immediately because their company-count starts at zero | n/a |
| Recovery after a worker is quarantined and returns | ❌ Returned worker is one cursor-step behind, gets ~1/N of new tenants | ✅ Returned worker is the most-empty by definition, gets the next assignments | n/a |
| Long-running fairness on uneven tenant lifetimes | ❌ Drifts unboundedly when tenant churn is asymmetric | ✅ Self-correcting — each new assignment compensates for prior imbalance | n/a |
| Predictability for tests & runbooks | ✅ Strictly deterministic order | ⚠️ Deterministic given (counts, tiebreaker) but counts shift | ✅ Operator-pinned |
| Query cost per assignment | O(1) cursor read | O(N) `COUNT(*) GROUP BY WorkerNodeId` over a tiny `WorkerNode` table | O(1) explicit pick |
| Behavior when all workers tied (cold start, same age) | ✅ Walks `Sequence` | ✅ Falls back to `WorkerNodeRegisteredAt ASC` tiebreaker — equivalent to `RoundRobin` first pass | n/a |

**When to override the default:**

| Override to | Use case |
|---|---|
| `RoundRobin` | Synthetic test environments where deterministic assignment ordering is required for replay; benchmark suites that must isolate routing variance from load variance. |
| `Manual` | Reserved-capacity tenants (enterprise contracts pinning to a dedicated worker); incident-response routing during a partial outage; canary-tenant pinning to a worker on a new release channel. |

**What is NOT a reason to override:** "we have only one worker" (any strategy works), "we want fairness" (already optimal), "we want speed" (the O(N) cost on a worker registry of typically 5–50 rows is sub-millisecond and dwarfed by the HTTP round trip).

**Migration path away from the default:** flip `MainWorker.Routing.DefaultStrategy` via Seedable-Config; existing `Company → Worker` mappings are NOT rebalanced (per §3.2 — tenant data lives on the assigned worker). Only **new** company creations observe the new strategy.

---

## 2. Caching

Per memory `mem://architecture/caching-policy`: explicit TTL, deterministic keys, invalidate on mutation.

| Cache key | Value | TTL | Invalidate when |
|-----------|-------|-----|-----------------|
| `MainWorker:Company:{CompanyId}:WorkerNodeId` | INTEGER | 15 min <!-- TUNABLE-WAIVER: cache TTL — owned by caching-policy memory, not MainWorker tunables --> | Worker reassignment, worker offline |
| `MainWorker:Registry:Active` | List of `WorkerNode` | 60 s <!-- TUNABLE-WAIVER: cache TTL — owned by caching-policy memory, not MainWorker tunables --> | Worker register/deregister/status change |
| `MainWorker:Session:{SessionId}:RecentCompanyId` | INTEGER | session lifetime | Logout |

Cache backend: Laravel cache driver (file/redis/memcached) — implementer's choice. The contract is the keys and TTLs above.

---

## 3. Failover

### 3.1 Worker becomes unreachable mid-request
1. Main retries per `15-tunable-constants.md` §2.1 (`RetryMaxAttempts`, `RetryBackoffSeconds`, `RetryJitterPct`).
2. On final failure: surface `WorkerUnreachable` to caller. Do NOT silently reroute — the user's data lives on that specific worker.
3. Log event with `X-Correlation-Id`. Per `spec/03-error-manage/`, never swallow.

### 3.2 Worker marked offline
- Background heartbeat checker flips status to `Quarantined` after `MainWorker.Heartbeat.MissedThreshold` consecutive misses (per `15-tunable-constants.md` §2.3); cooldown before re-eligibility = `MainWorker.Heartbeat.QuarantineCooldownSeconds`.
- Existing `Company → Worker` mappings are NOT reassigned automatically. Tenant data is on that worker.
- Power Admin can trigger manual reassignment via `POST /API/V1/Workers/{From}/Migrate/{To}` (deferred — not in initial endpoint set).

### 3.3 Worker comes back online
- On heartbeat resume, status flips to `Active`.
- Existing tenants resume routing automatically.

---

## 4. Migration of Existing Tenants (deferred)

Migrating a `Company` from Worker A → Worker B requires:
1. Quiesce traffic to Worker A for that company.
2. Copy split-DB rows (per `spec/05-split-db-architecture/`).
3. Update `Company.WorkerNodeId` on Main.
4. Invalidate routing cache.
5. Resume traffic.

Out of scope for v1.0 — flagged as deferred work. Only `Manual` strategy lets Power Admin influence assignment for new tenants.

---

## 5. Routing Decision Function (pseudocode)

Compliant with CODE RED (≤15 lines, zero nesting, positive guards, max 2 operands):

```php
public function pickWorker(int $companyId, string $strategyCode): WorkerNode
{
    $eligible = $this->getEligibleWorkers();
    $this->guardAtLeastOneEligible($eligible);
    $strategy = $this->strategyResolver->resolve($strategyCode);
    $worker   = $strategy->pick($eligible);
    $this->recordSelectionEvent($companyId, $worker, $strategyCode);
    return $worker;
}
```

Each helper (`getEligibleWorkers`, `guardAtLeastOneEligible`, `recordSelectionEvent`) is its own ≤8-line function. `guardAtLeastOneEligible` throws `WorkerUnavailable` when the list is empty.

### 5.1 Required interfaces (resolves F-A-33)

The pseudocode above relies on two types that implementations MUST define so the contract is portable across PHP, Go, Rust, and TypeScript stacks:

```php
interface WorkerSelectionStrategyInterface {
    /** Returns one WorkerNode from the supplied eligible candidates.
     *  MUST be deterministic given identical input. MUST throw
     *  `WorkerUnavailable` if the input is empty. */
    public function pick(array $candidates): WorkerNode;

    /** Returns the PascalCase code this strategy is registered under
     *  in `WorkerSelectionStrategy` (ref table, see `03-` §2.9). */
    public function code(): string;
}

interface WorkerSelectionStrategyResolverInterface {
    /** Resolves a strategy code to its concrete strategy. MUST throw
     *  `ConfigurationError` if the code is unknown. */
    public function resolve(string $strategyCode): WorkerSelectionStrategyInterface;
}
```

Allowed concrete strategies in v1.0 (one class each, registered by code): `RoundRobin`, `LeastLoaded`, `Manual`. Adding a new strategy = new class + new ref-table row + Seedable-Config bump (per `15-tunable-constants.md` rules); no existing class changes.

---

## 6. Observability

Every selection writes one row to `WorkerSelectionEvent`. Operators can query distribution:

```sql
SELECT WorkerNodeId, COUNT(*) AS Picked
FROM   WorkerSelectionEvent
WHERE  WorkerSelectionEventAt > datetime('now', '-7 days')
GROUP  BY WorkerNodeId
ORDER  BY Picked DESC;
```

---

## 7. Endpoint Schema (authoritative)

All endpoints below MUST honour the cross-spec header rules in `spec/04-database-conventions/06-rest-api-format.md` (`X-Correlation-Id`, `X-Idempotency-Key`, `X-Auth-Action`) and the tunables in `15-tunable-constants.md`.

### 7.1 Common envelope

| Field | Rule |
|-------|------|
| Request body | `application/json; charset=utf-8`. PascalCase keys. |
| Response body | `{ "Data": <T> | null, "Error": <ErrorEnvelope> | null, "Meta": { "CorrelationId": "<ULID>", "IdempotencyReplay": <bool> } }`. Exactly one of `Data` / `Error` is non-null. |
| Error envelope | Per `13-error-codes.md` §3 (`Code`, `Message`, `HttpStatus`, `Details`, `CorrelationId`). |
| Auth | Bearer JWT per `12-jwt-delivery-contract.md`. Unauthenticated endpoints explicitly noted. |
| Idempotency | `POST` / `PUT` / `PATCH` REQUIRE `X-Idempotency-Key` (ULID, 26 chars). Cached for `IdempotencyKeyTtlSeconds = 86400` per `15-tunable-constants.md`. Replay returns the original response with `X-Idempotency-Replay: true` and `Meta.IdempotencyReplay = true`. |

### 7.2 Endpoint catalogue

| # | Method + Path | Purpose | Auth | Idempotent | Request body | Success status | Error codes (subset) |
|---|---------------|---------|------|------------|--------------|----------------|----------------------|
| 1 | `POST /API/V1/Workers/Register` | Worker boot — register node, fetch JWT pubkey. Per `10-worker-bootstrap-protocol.md`. | mTLS or shared bootstrap secret | Yes (key = `WorkerNodeCode`) | `{ WorkerNodeCode, WorkerNodeName, AdvertisedHost, AdvertisedPort, WorkerVersionCode, Capabilities[] }` | `201 Created` → `{ WorkerNodeId, JwtPublicKeyPem, AssignedSubdomain, HeartbeatIntervalSeconds }` | `WORKER-100-01..05` |
| 2 | `POST /API/V1/Workers/{WorkerNodeId}/Heartbeat` | Liveness ping. | Worker JWT | Yes (natural — last-write-wins) | `{ LoadFactor (0..1), AssignedCompanyCount, FreeDiskBytes, UptimeSeconds }` | `200 OK` → `{ NextHeartbeatInSeconds }` | `WORKER-200-01..03` |
| 3 | `POST /API/V1/Companies` | Create tenant; routing picks worker per §1. | User JWT (role: PowerAdmin or CompanyAdmin) | Yes (key = client ULID) | `{ CompanyCode, CompanyName, OwnerUserId, RequestedWorkerNodeId? (Manual only), Notes?, Comments? }` | `201 Created` → `{ CompanyId, WorkerNodeId, AssignedSubdomain, SplitDbProvisionState }` | `MAIN-300-01..04`, `WORKER-300-01..02` |
| 4 | `GET /API/V1/Companies/{CompanyId}/Routing` | Resolve worker for tenant (cache-first). | User JWT | n/a (GET) | — | `200 OK` → `{ WorkerNodeId, AdvertisedHost, AdvertisedPort, CacheHit (bool), TtlRemainingSeconds }` | `MAIN-300-05`, `WORKER-300-03` |
| 5 | `POST /API/V1/Auth/Login` | Email + password → user session + worker JWT. Per `12-jwt-delivery-contract.md`. | none | Yes (key = client ULID; replay returns same JWT only if not yet expired) | `{ Email, Password, X-Auth-Action: "Login" header }` | `200 OK` → `{ UserId, WorkerJwt, WorkerJwtExpiresAt, RequiresTwoFactor (bool) }` | `MAIN-400-01..04` |
| 6 | `POST /API/V1/Auth/TwoFactor/Verify` | Complete 2FA, mint final worker JWT. | partial-auth JWT from §5 | Yes (key = client ULID) | `{ Code, X-Auth-Action: "TwoFactorVerify" header }` | `200 OK` → `{ WorkerJwt, WorkerJwtExpiresAt }` | `MAIN-400-05..07` |
| 7 | `POST /API/V1/Auth/Refresh` | Rotate worker JWT before expiry. | expiring worker JWT | No (rotation is single-use; replay → `MAIN-400-09`) | `{}` (auth from header) | `200 OK` → `{ WorkerJwt, WorkerJwtExpiresAt }` | `MAIN-400-08..10` |
| 8 | `POST /API/V1/Workers/{WorkerNodeId}/PushUpdate` | Power-Admin push-update fan-out. Per `spec/14-update/28-worker-push-instruction.md`. | User JWT (role: PowerAdmin) | Yes (key = `WorkerUpdateInstructionId`) | `{ TargetVersionCode, RolloutMode ("Immediate" | "Scheduled"), ScheduledAt?, Notes }` | `202 Accepted` → `{ WorkerUpdateInstructionId, EstimatedCompletionAt }` | `WORKER-500-01..06` |
| 9 | `GET /API/V1/Workers` | List registry (cached, see §2). | User JWT (role: PowerAdmin) | n/a | — | `200 OK` → `{ Workers[]: { WorkerNodeId, Status, LoadFactor, AssignedCompanyCount, LastHeartbeatAt } }` | `MAIN-600-01` |
| 10 | `POST /API/V1/Workers/{From}/Migrate/{To}` | **Deferred** (see §4). Documented for contract stability. | User JWT (role: PowerAdmin) | Yes (key = client ULID) | `{ CompanyIds[], Reason }` | `202 Accepted` → `{ MigrationId }` | `WORKER-700-01..03` |

### 7.3 Idempotency rules (normative)

1. **Key scope.** Cached per `(EndpointPath, AuthenticatedSubjectId, IdempotencyKey)`. Reusing a key across users is a `400 IdempotencyKeyConflict`.
2. **Window.** `IdempotencyKeyTtlSeconds = 86400` (`15-tunable-constants.md`). Late replays past the window are treated as fresh requests.
3. **Body invariance.** Replay with a different request body for the same key → `409 IdempotencyBodyMismatch` (`MAIN-300-04`).
4. **Replay surface.** Replays return the original `HttpStatus`, body, `X-Correlation-Id` of the original request, AND `X-Idempotency-Replay: true`.
5. **Failure caching.** Only `2xx` and deterministic `4xx` (validation) responses are cached. `5xx` and `429` MUST NOT be cached.
6. **GET / non-mutating.** `GET` endpoints ignore `X-Idempotency-Key` (no caching contract).

### 7.4 Validation contract

- Missing `X-Correlation-Id` on any request → `400 BadRequest` (`MAIN-300-01`).
- Missing `X-Idempotency-Key` on `POST`/`PUT`/`PATCH` → `400 BadRequest` (`WORKER-300-01`).
- Missing `X-Auth-Action` on multi-step auth (§5, §6) → `400 BadRequest` (`MAIN-400-11`).
- Body schema violations → `422 UnprocessableEntity` with per-field `Details[]`.

Resolves audit findings F-X-11..F-X-15 (endpoint contract gaps) and closes the original top-10 fix list.

---

*Worker routing v1.1.0 — 2026-05-04 (added §7 endpoint schema; audit fix #10)*
