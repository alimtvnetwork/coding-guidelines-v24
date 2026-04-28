# Docs Index

Deep-dive documentation for the coding-guidelines repository. The root [`readme.md`](../readme.md) stays under 400 lines by design — long-form material lives here.

> **Version:** <!-- STAMP:VERSION -->4.24.0<!-- /STAMP:VERSION ·
> **Updated:** <!-- STAMP:UPDATED -->2026-04-28<!-- /STAMP:UPDATED -->

---

## Table of contents

- [Core references](#core-references)
  - [principles](#principles)
  - [architecture](#architecture)
  - [author](#author)
- [Installers & tooling](#installers--tooling)
  - [installer fix-repo flags](#installer-fix-repo-flags)
  - [slides installer](#slides-installer)
  - [github repo metadata](#github-repo-metadata)
- [Process & quality](#process--quality)
  - [spec-author DX](#spec-author-dx)
  - [guidelines audit](#guidelines-audit)
- [Refactor case studies](#refactor-case-studies)
  - [payment banner hider](#payment-banner-hider)
- [Related entry points](#related-entry-points)

---

## Core references

### principles

- **Path:** [`principles.md`](principles.md)
- **Summary:** The 9 core development principles, 10 CODE RED rules, the cross-language rule index, and the AI optimization suite.
- **Start here if:** you're new to the repo and want the foundational rules.

### architecture

- **Path:** [`architecture.md`](architecture.md)
- **Summary:** Spec authoring conventions, folder structure (`01-20` core / `21+` app), architecture decisions, and the error-management summary.
- **Start here if:** you need to understand how specs and folders are organized.

### author

- **Path:** [`author.md`](author.md)
- **Summary:** Author bio (Md. Alim Ul Karim), Riseup Asia LLC, neutral AI assessments, FAQ, and the design philosophy behind the spec system.
- **Start here if:** you want context on who maintains the repo and why.

---

## Installers & tooling

### installer fix-repo flags

- **Path:** [`installer-fix-repo-flags.md`](installer-fix-repo-flags.md)
- **Summary:** Post-`fix-repo` log retention and rollback flags: `--max-fix-repo-logs`, `INSTALL_MAX_FIX_REPO_LOGS`, `--rollback-on-fix-repo-failure`, `--full-rollback`, plus the interaction matrix.
- **Start here if:** you're tuning installer log retention or rollback behavior.

### slides installer

- **Path:** [`slides-installer.md`](slides-installer.md)
- **Summary:** The `slides` bundle installer — flags, packaging pipeline, and offline behavior.
- **Start here if:** you're packaging or installing the slides bundle.

### github repo metadata

- **Path:** [`github-repo-metadata.md`](github-repo-metadata.md)
- **Summary:** Canonical GitHub repo description, topics, and About-section sourcing rules (manual action required when changed).
- **Start here if:** you're updating the GitHub About section or topics.

---

## Process & quality

### spec-author DX

- **Path:** [`spec-author-dx.md`](spec-author-dx.md)
- **Summary:** Spec-author developer experience — tooling, workflow ergonomics, and what makes spec authoring fast.
- **Start here if:** you author or maintain specs day-to-day.

### guidelines audit

- **Path:** [`guidelines-audit.md`](guidelines-audit.md)
- **Summary:** Independent audit of the coding guidelines with top-3 quick wins, drift detection, and the remediation log.
- **Start here if:** you want a critical view of guideline coverage and gaps.

---

## Refactor case studies

### payment banner hider

- **Path:** [`refactors/payment-banner-hider.md`](refactors/payment-banner-hider.md)
- **Summary:** Worked example: refactoring the payment-banner hider to comply with CODE RED metrics (function length, zero-nesting, positive guards).
- **Start here if:** you want a concrete CODE RED refactor walkthrough.

---

## Related entry points

- **Live spec tree:** [`../spec/`](../spec/) — 22 numbered folders, the source of truth for all rules.
- **Health dashboard:** [`../spec/health-dashboard.md`](../spec/health-dashboard.md) — coverage, drift, and quality metrics.
- **Consolidated index:** [`../spec/17-consolidated-guidelines/00-overview.md`](../spec/17-consolidated-guidelines/00-overview.md) — single-page reference of every rule.
- **Changelog:** [`../changelog.md`](../changelog.md)

---

<sub>Maintained by <a href="https://alimkarim.com/">Md. Alim Ul Karim</a> · <a href="https://riseup-asia.com/">Riseup Asia LLC</a></sub>
