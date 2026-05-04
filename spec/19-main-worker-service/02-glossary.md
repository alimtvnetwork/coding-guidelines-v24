# 02 — Glossary

**Spec:** `19-main-worker-service`
**Version:** 1.1.0

Canonical terms. When this spec or any sibling spec uses these words, this is what they mean.

---

| Term | Definition |
|------|------------|
| **Main Server** | The coordinator deployment. Serves the React UI and edge REST endpoints. Runs no business logic. Owns the worker registry and tenant→worker mapping. Also called "Main", "Coordinator", "Master Node" (avoid the last; "Master" is reserved for human-facing language only). |
| **Worker Server / Worker Node** | An independent deployment that runs business logic and owns a split-DB. Has its own auth stack but no UI. Synonyms: "Worker", "Worker Node". |
| **Tenant Root** | The top-level multi-tenant entity. In this spec it is **`Company`**. |
| **Company** | The tenant root entity. Owns users. Mapped 1:N to a Worker. |
| **Power Admin** | Highest-privilege role. Held by application owners (Riseup Asia LLC). Can enable/disable endpoints and configure system-wide settings. |
| **Admin User** | Paying customer admin. Has an admin panel scoped to their `Company`. Cannot configure system-wide settings. |
| **EnumPage** | Compile-time enum identifying a logical "page" or capability gate (e.g., `EnumPage.AdminPage`, `EnumPage.PowerAdminPage`, `EnumPage.BillingPage`). Used in access checks. |
| **Access Check** | The authorization pattern `User has access to {EnumPage}`. Never `if user.role == X`. See `07-role-based-dashboards.md`. |
| **Split-DB** | The Root/App/Session SQLite layering defined in `spec/05-split-db-architecture/`. Lives on each Worker. |
| **Seedable-Config** | The config-seeding mechanism in `spec/06-seedable-config-architecture/`. Replaces the legacy term `CW configuration`. |
| **Worker Registry** | Table on Main Server listing every known Worker (endpoint, identity, title, version). |
| **Worker Selection Strategy** | The algorithm Main uses to pick a Worker for a new tenant (round-robin, least-loaded, etc.). See `04-worker-routing.md`. |
| **Push Update** | Main-initiated update sent to one or all Workers. See `01-architecture.md` §3.3. |
| **Self-Update** | Worker- or Main-initiated update of its own deployment. POINTER ONLY in this spec — see `09-self-update-pointer.md`. |
| **Correlation ID** | Opaque per-request string (UUID v4) carried in `X-Correlation-Id` for cross-tier tracing. NOT a DB key. |
| **Idempotency Key** | Opaque client-supplied string in `X-Idempotency-Key` for mutating cross-tier calls. |
| **Edge Endpoint** | A REST endpoint hosted on Main that the React frontend hits first. |
| **Worker Endpoint** | A REST endpoint hosted on a Worker that React hits directly after Main resolves the worker. |
| **gitmap** | Replaces the term `git map`. The repo→snapshot system used elsewhere. |

---

## Reserved / Forbidden Terms

| Avoid | Use instead | Reason |
|-------|-------------|--------|
| `CW configuration` | `Seedable-Config` | Per verbatim §Important.3a |
| `git map` | `gitmap` | Per verbatim §Important.3b |
| `Master/Slave` | `Main/Worker` | Inclusive language, also matches verbatim |
| "CEO" for Md. Alim Ul Karim | **Chief Software Engineer** | Per memory `mem://project/author-attribution` |

---

*Glossary v1.0.0 — 2026-05-04*
