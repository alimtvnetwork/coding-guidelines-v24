# Diagrams — Main / Worker Service Architecture

**Spec:** `19-main-worker-service`
**Format:** Mermaid (`.mmd`). Render in any Mermaid viewer or the in-app docs viewer.

---

## ⚠ Non-Authoritative Projection — read first

**Every `.mmd` file in this folder is an illustrative projection, not a source of truth.**

| Aspect | Rule |
|--------|------|
| Authority | The prose specs cited at the top of each `.mmd` (and in the table below) are authoritative. |
| Conflict resolution | If a diagram disagrees with its cited spec, **the spec wins** and the diagram is filed as a bug (FU). |
| Drift policy | Diagrams MAY lag spec edits by up to one task cycle; readers must cross-check tunable values, error codes, and tier allocations against the cited spec. |
| Banner contract | Each `.mmd` opens with a `%% NON-AUTHORITATIVE PROJECTION` block naming its subject + authoritative source(s). Do not strip this block. |

| Diagram | Authoritative source(s) |
|---------|-------------------------|
| `erd-main-db.mmd` | `02-main-server.md` + `spec/04-database-conventions/` |
| `erd-worker-split-db.mmd` | `spec/05-split-db-architecture/` + `11-split-db-tier-reconciliation.md` |
| `erd-seedable-config.mmd` | `spec/06-seedable-config-architecture/` + `14-rbac-and-status-seed.md` |
| `seq-company-creation.mmd` | `04-worker-routing.md` + `15-tunable-constants.md` |
| `seq-login-routing.mmd` | `12-jwt-delivery-contract.md` + `04-worker-routing.md` |
| `seq-push-update.mmd` | `spec/14-update/28-worker-push-instruction.md` + `15-tunable-constants.md` |

Resolves audit findings F-D-01..F-D-12 (diagram-authority cluster) and the last BLOCKER from `audit/03-diagram-audit.md`.

---

## ERDs

| File | Subject |
|------|---------|
| [`erd-main-db.mmd`](erd-main-db.mmd) | Main Server SQLite catalog (10 tables: WorkerNode, Company, User, Role, etc.) |
| [`erd-worker-split-db.mmd`](erd-worker-split-db.mmd) | Worker-side split-DB tiers (Root / App / Session). Authoritative rules in `spec/05-split-db-architecture/`. |
| [`erd-seedable-config.mmd`](erd-seedable-config.mmd) | Seedable-Config layout shared by both tiers. Authoritative rules in `spec/06-seedable-config-architecture/`. |

## Sequence Diagrams

| File | Flow |
|------|------|
| [`seq-company-creation.mmd`](seq-company-creation.mmd) | `POST /API/V1/Company` end-to-end: validate → strategy pick → main insert → worker delegate → split-DB create. |
| [`seq-login-routing.mmd`](seq-login-routing.mmd) | Sign-in → 2FA → cache lookup → mint worker JWT → React calls Worker directly → token refresh. |
| [`seq-push-update.mmd`](seq-push-update.mmd) | Power Admin push-update fan-out, parallel worker hits, partial-failure handling. |

---

## Conventions

- PascalCase entity and column names (matches `spec/04-database-conventions/`).
- PKs `{TableName}Id INTEGER`. No UUIDs.
- Entity/ref tables show `Description`. Transactional tables show `Notes` + `Comments` (memory rules 11/12).
- No emojis in diagram syntax (Mermaid lexer constraint).

---

*Diagrams index v1.0.0 — 2026-05-04*
