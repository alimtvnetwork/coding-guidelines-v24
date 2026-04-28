# Docs Index

Deep-dive documentation for the coding-guidelines repository. The root [`readme.md`](../readme.md) stays under 400 lines by design — long-form material lives here.

> **Version:** <!-- STAMP:VERSION -->4.24.0<!-- /STAMP:VERSION ·
> **Updated:** <!-- STAMP:UPDATED -->2026-04-28<!-- /STAMP:UPDATED -->

---

## Core references

| Page | What's inside |
|---|---|
| [`principles.md`](principles.md) | The 9 core development principles, 10 CODE RED rules, the cross-language rule index, and the AI optimization suite. Start here. |
| [`architecture.md`](architecture.md) | Spec authoring conventions, folder structure (`01-20` core / `21+` app), architecture decisions, and the error-management summary. |
| [`author.md`](author.md) | Author bio (Md. Alim Ul Karim), Riseup Asia LLC, neutral AI assessments, FAQ, and the design philosophy behind the spec system. |

## Installers & tooling

| Page | What's inside |
|---|---|
| [`installer-fix-repo-flags.md`](installer-fix-repo-flags.md) | Post-`fix-repo` log retention and rollback flags: `--max-fix-repo-logs`, `INSTALL_MAX_FIX_REPO_LOGS`, `--rollback-on-fix-repo-failure`, `--full-rollback`, plus the interaction matrix. |
| [`slides-installer.md`](slides-installer.md) | The `slides` bundle installer — flags, packaging pipeline, and offline behavior. |
| [`github-repo-metadata.md`](github-repo-metadata.md) | Canonical GitHub repo description, topics, and About-section sourcing rules (manual action required when changed). |

## Process & quality

| Page | What's inside |
|---|---|
| [`spec-author-dx.md`](spec-author-dx.md) | Spec-author developer experience — tooling, workflow ergonomics, and what makes spec authoring fast. |
| [`guidelines-audit.md`](guidelines-audit.md) | Independent audit of the coding guidelines with top-3 quick wins, drift detection, and the remediation log. |

## Refactor case studies

| Page | What's inside |
|---|---|
| [`refactors/payment-banner-hider.md`](refactors/payment-banner-hider.md) | Worked example: refactoring the payment-banner hider to comply with CODE RED metrics (function length, zero-nesting, positive guards). |

---

## Related entry points

- **Live spec tree:** [`../spec/`](../spec/) — 22 numbered folders, the source of truth for all rules.
- **Health dashboard:** [`../spec/health-dashboard.md`](../spec/health-dashboard.md) — coverage, drift, and quality metrics.
- **Consolidated index:** [`../spec/17-consolidated-guidelines/00-overview.md`](../spec/17-consolidated-guidelines/00-overview.md) — single-page reference of every rule.
- **Changelog:** [`../changelog.md`](../changelog.md)

---

<sub>Maintained by <a href="https://alimkarim.com/">Md. Alim Ul Karim</a> · <a href="https://riseup-asia.com/">Riseup Asia LLC</a></sub>
