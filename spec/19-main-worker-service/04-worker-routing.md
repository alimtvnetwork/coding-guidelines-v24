# 04 — Worker Routing

**Spec:** `19-main-worker-service`
**Version:** 1.0.0

How the Main Server picks which Worker handles a new tenant, caches that decision, and recovers when a Worker fails.

---

## 1. Selection Strategies

Strategy is configurable via Seedable-Config key `MainWorker.Routing.DefaultStrategy`. Stored in `WorkerSelectionStrategy` table.

### 1.1 `RoundRobin`
- Pick next `Active` worker in registry order.
- Cursor persisted in main DB (single-row config table or `WorkerSelectionEvent` last-row lookup).
- Pros: trivially predictable. Cons: ignores load.

### 1.2 `LeastLoaded` (recommended default)
- Pick `Active` worker with fewest assigned `Company` rows.
- Tiebreaker: oldest `WorkerNodeRegisteredAt`.
- Pros: balances over time. Cons: slightly more expensive query (still O(N) on workers, N is small).

### 1.3 `Manual`
- Power Admin specifies `WorkerNodeId` in the create request.
- Used for testing and reserved-capacity tenants.
- Requires `User has access to EnumPage.PowerAdminPage`.

### 1.4 Eligibility filter (applies to all strategies)
A worker is eligible only if **all** are true (positive guards, per CODE RED):
- `IsWorkerActive(node)` → `WorkerNodeStatusCode = 'Active'`
- `IsWorkerReachable(node)` → last heartbeat within `MainWorker.Routing.HeartbeatWindowSeconds` (default 60s)
- `HasCapacity(node)` → assigned company count below `MainWorker.Routing.MaxCompaniesPerWorker` (0 = unlimited)

If no eligible worker exists → return `WorkerUnavailable` error per `08-error-contract.md`.

---

## 2. Caching

Per memory `mem://architecture/caching-policy`: explicit TTL, deterministic keys, invalidate on mutation.

| Cache key | Value | TTL | Invalidate when |
|-----------|-------|-----|-----------------|
| `MainWorker:Company:{CompanyId}:WorkerNodeId` | INTEGER | 15 min | Worker reassignment, worker offline |
| `MainWorker:Registry:Active` | List of `WorkerNode` | 60 s | Worker register/deregister/status change |
| `MainWorker:Session:{SessionId}:RecentCompanyId` | INTEGER | session lifetime | Logout |

Cache backend: Laravel cache driver (file/redis/memcached) — implementer's choice. The contract is the keys and TTLs above.

---

## 3. Failover

### 3.1 Worker becomes unreachable mid-request
1. Main retries per `01-architecture.md` §4 (max 3 attempts, exponential backoff).
2. On final failure: surface `WorkerUnreachable` to caller. Do NOT silently reroute — the user's data lives on that specific worker.
3. Log event with `X-Correlation-Id`. Per `spec/03-error-manage/`, never swallow.

### 3.2 Worker marked offline
- Background heartbeat checker flips status to `Offline` after `HeartbeatWindowSeconds × 3`.
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
    $worker  = $this->strategyResolver->resolve($strategyCode)->pick($eligible);
    $this->recordSelectionEvent($companyId, $worker, $strategyCode);
    return $worker;
}
```

Each helper (`getEligibleWorkers`, `guardAtLeastOneEligible`, `recordSelectionEvent`) is its own ≤8-line function. `guardAtLeastOneEligible` throws `WorkerUnavailable` when the list is empty.

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

*Worker routing v1.0.0 — 2026-05-04*
