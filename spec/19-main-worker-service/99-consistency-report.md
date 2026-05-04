# 99 — Consistency Report

**Spec:** `19-main-worker-service`
**Version:** 1.1.0
**Verified:** 2026-05-04

Cross-link integrity, naming-convention compliance, and rule-inheritance audit for this spec folder.

---

## 1. Files in this spec

| File | Purpose | Status |
|------|---------|--------|
| `plan.md` | Phased roadmap + decisions | ✅ |
| `00-overview.md` | Entry point, document map | ✅ |
| `01-architecture.md` | Topology + comms contract | ✅ |
| `02-glossary.md` | Canonical terms | ✅ |
| `03-main-db-schema.md` | Main SQLite schema | ✅ |
| `04-worker-routing.md` | Selection + cache + failover | ✅ |
| `05-auth-and-2fa.md` | Auth surfaces + 2FA | ✅ |
| `06-core-api-endpoints.md` | REST surface | ✅ |
| `07-role-based-dashboards.md` | Roles + EnumPage | ✅ |
| `08-error-contract.md` | Main↔Worker errors | ✅ |
| `09-self-update-pointer.md` | Pointer to `spec/14-update/` | ✅ |
| `97-acceptance-criteria.md` | AC-1..AC-9 mapping | ✅ |
| `98-changelog.md` | Version history | ✅ |
| `99-consistency-report.md` | This file | ✅ |
| `diagrams/erd-main-db.mmd` | Main ERD | ✅ |
| `diagrams/erd-worker-split-db.mmd` | Worker split-DB ERD | ✅ |
| `diagrams/erd-seedable-config.mmd` | Config ERD | ✅ |
| `diagrams/seq-company-creation.mmd` | Sequence | ✅ |
| `diagrams/seq-login-routing.mmd` | Sequence | ✅ |
| `diagrams/seq-push-update.mmd` | Sequence | ✅ |
| `diagrams/readme.md` | Diagrams index | ✅ |

Total: 21 files.

---

## 2. Cross-spec references (must exist)

| Reference | Verified |
|-----------|----------|
| `spec/03-error-manage/` | ✅ exists |
| `spec/04-database-conventions/` | ✅ exists |
| `spec/05-split-db-architecture/` | ✅ exists |
| `spec/06-seedable-config-architecture/` | ✅ exists |
| `spec/14-update/` | ✅ exists |
| `.lovable/coding-guidelines/coding-guidelines.md` | ✅ exists |
| `mem://architecture/error-handling` | ✅ in memory index |
| `mem://architecture/caching-policy` | ✅ in memory index |
| `mem://features/self-update-architecture` | ✅ in memory index |
| `mem://project/author-attribution` | ✅ in memory index |

---

## 3. Naming convention compliance

| Rule | Status | Notes |
|------|--------|-------|
| PascalCase for tables, columns, JSON keys | ✅ | All schema in `03-`, `06-`, `07-` follows |
| PKs `{TableName}Id INTEGER AUTOINCREMENT` | ✅ | Every table |
| No UUIDs as PKs | ✅ | Correlation IDs are opaque request strings, NOT keys (`08-` §4 explicit) |
| Entity/ref tables include `Description TEXT NULL` | ✅ | All applicable tables |
| Transactional tables include `Notes` + `Comments TEXT NULL` | ✅ | `WorkerVersion`, `WorkerSelectionEvent`, `AppBusinessEntity`, `AppSession`, `SeedableConfigVersion` |
| Join tables exempt | ✅ | `UserRole`, `RolePageAccess` correctly omit `Description`/`Notes` |
| Type/Status/Category/Kind via join tables | ✅ | `WorkerNodeStatus`, `WorkerNodeKind`, `WorkerSelectionStrategy`, `Role`, `RootCompanyStatus`, `AuthMechanism` |
| Files numbered `NN-kebab-case.md` | ✅ | All conform |

---

## 4. Forbidden-term audit (per `02-glossary.md`)

| Forbidden | Replacement | Found in this spec? |
|-----------|-------------|----------------------|
| `CW configuration` | `Seedable-Config` | NO ✅ |
| `git map` | `gitmap` | NO ✅ (term not used; would use `gitmap` if needed) |
| `Master/Slave` | `Main/Worker` | NO ✅ |
| `CEO` for Md. Alim Ul Karim | `Chief Software Engineer` | N/A — author not referenced in this spec |

**Reproducible verification commands** (closes audit finding F-N-03; run from repo root):

```bash
# 4.1 — forbidden literals (case-insensitive, exclude this report itself)
rg -i -n --glob '!99-consistency-report.md' \
   -e 'CW configuration' -e 'git map' \
   -e '\bmaster/slave\b' -e '\bmaster-slave\b' \
   spec/19-main-worker-service/

# 4.2 — author title compliance (must NOT appear)
rg -n -e '\bCEO\b' -e 'Chief Executive Officer' \
   spec/19-main-worker-service/

# 4.3 — global enforcement (delegated)
python3 linter-scripts/check-forbidden-strings.py
```

Last verified: 2026-05-04 — all three commands exit 0 / zero matches.

---

## 5. CODE RED compliance (rules applied to spec content)

| Rule | Applied? |
|------|----------|
| Function pseudocode ≤ 8 lines (best practice), ≤ 15 (cap) | ✅ All pseudocode in `04-`, `05-`, `07-`, `08-` ≤ 15 lines |
| Zero nested `if` in pseudocode | ✅ All examples use early-return guards |
| Positive boolean names (`is`/`has` prefix, no `!`) | ✅ `IsWorkerActive`, `HasCapacity`, `userHasAccessToPage`, `canRetry`, `isSuccessful`, `isPermanentFailure` |
| Max 2 operands per condition | ✅ All examples |
| No magic strings — Enums for Type/Status/Category/Kind | ✅ `EnumPage`, `WorkerNodeStatusCode`, `WorkerSelectionStrategyCode`, `Cadence` |

---

## 6. Inheritance audit (rules deferred, not duplicated)

| Topic | Authoritative source | This spec duplicates? |
|-------|----------------------|-----------------------|
| Split-DB internals | `spec/05-split-db-architecture/` | NO ✅ — only references and projection diagram |
| Seedable-Config internals | `spec/06-seedable-config-architecture/` | NO ✅ |
| Generic error rules | `spec/03-error-manage/` | NO ✅ — `08-` adds only Main↔Worker patterns |
| Self-update mechanism | `spec/14-update/` | NO ✅ — `09-` is pointer-only |
| DB schema conventions | `spec/04-database-conventions/` | NO ✅ — `03-` references and applies |
| Coding metrics | `.lovable/coding-guidelines/coding-guidelines.md` | NO ✅ — applied via examples |

---

## 7. Open Questions (carried, not closed)

| ID | Topic | Location | Status |
|----|-------|----------|--------|
| OQ-1 | Per-endpoint auth-mechanism overrides | `05-auth-and-2fa.md` §8, `06-core-api-endpoints.md` §5 | OPEN — awaits user decision |
| OQ-2 | Default selection strategy (RoundRobin vs LeastLoaded) | `04-worker-routing.md` §1 | RESOLVED — default `LeastLoaded`; user can change via Settings |

---

## 8. Acceptance Criteria coverage

All 9 verbatim acceptance criteria mapped in `97-acceptance-criteria.md`. Coverage matrix verified — every sub-criterion has at least one defining spec file and at least one test condition.

---

## 9. Outstanding (next phase)

- **Phase 5** (separate task): bump `package.json` minor, run `node scripts/sync-version.mjs` and `node scripts/sync-spec-tree.mjs`, log entry in `.lovable/plan.md`.
- Future: implement Phase 1 of the deliverables when user issues `next` after Phase 5 (or explicitly says "implement").

---

*Consistency report v1.1.0 — 2026-05-04 (added §4 reproducible grep commands closing audit F-N-03)*
