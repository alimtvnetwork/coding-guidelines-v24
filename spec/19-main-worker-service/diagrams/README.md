# Diagrams — Main / Worker Service Architecture

**Spec:** `19-main-worker-service`
**Format:** Mermaid (`.mmd`). Render in any Mermaid viewer or the in-app docs viewer.

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
