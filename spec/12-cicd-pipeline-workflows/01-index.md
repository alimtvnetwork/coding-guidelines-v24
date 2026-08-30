# CI/CD Pipeline Workflows

> **/goal** Master and enforce the architectural standards, specifications, and CI/CD validation rules for CI/CD Pipeline Workflows.
> **/learn** Read the sequentially ordered specification files in this directory, follow the actionable CI/CD checklist, and apply mandatory rules before generating code.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Read and understand all numbered specifications under `12-cicd-pipeline-workflows/`.
- [ ] `/learn` Adhere strictly to `.lovable/folder-structure.md` and `.lovable/strictly-avoid.md`.
- [ ] `/goal` Verify zero explicit `true` boolean evaluations and no mixed-polarity conditionals.
- [ ] `/learn` Run all local verification linters via `python .lovable/ai-fix-scripts/03-cicd-local-runner.py`.

---

**Version:** 4.0.0  
**Updated:** 2026-08-30  
**AI Confidence:** Production-Ready  
**Ambiguity:** None

---

## Purpose

Central location for all CI/CD pipeline specifications, deployment automation, and related infrastructure-as-code documentation. All pipeline-related content — build pipelines, deployment workflows, environment promotion strategies, and CI/CD tooling configurations — MUST be documented in this folder.

---

## Scope

This module covers two distinct pipeline archetypes, shared conventions, and cross-cutting concerns:

| Archetype | Subfolder | Description |
|-----------|-----------|-------------|
| Browser Extension Deploy | `01-browser-extension-deploy/` | Node.js/pnpm multi-component builds, zip packaging, Chrome Web Store |
| Go Binary Deploy | `02-go-binary-deploy/` | Cross-compiled Go binaries, tar.gz/zip, install scripts, code signing |
| Reusable CI Guards | `03-reusable-ci-guards/` | Language-agnostic baseline diff gating and quality guards |
| Gitmap Extended Pipeline | `04-gitmap-pipeline/` | ❓ Extended multi-architecture workflow definitions (Pending Review) |
| Shared Conventions | Root files | Common patterns used across all pipeline types |

---

## Feature Inventory

### Root (Shared Conventions & Workflows)

| # | File | Description | Status |
|---|------|-------------|--------|
| 02 | [02-ci-pipeline.md](./02-ci-pipeline.md) | Core CI pipeline execution matrix and stages | ✅ Active |
| 03 | [03-shared-conventions.md](./03-shared-conventions.md) | Platform, triggers, concurrency, version resolution, checksums | ✅ Active |
| 04 | [04-github-release-standard.md](./04-github-release-standard.md) | Release body assembly, pre-release detection, asset matrix | ✅ Active |
| 05 | [05-release-pipeline.md](./05-release-pipeline.md) | Release workflow and automated deployment triggers | ✅ Active |
| 06 | [06-vulnerability-scanning.md](./06-vulnerability-scanning.md) | Standalone and in-CI vulnerability scanning patterns | ✅ Active |
| 07 | [07-install-script-generation.md](./07-install-script-generation.md) | Reusable PS1+Bash installer pattern, placeholder strategy, checksum verification | ✅ Active |
| 08 | [08-installation-flow.md](./08-installation-flow.md) | End-to-end install: one-liners, terminal output, upgrade, uninstall | ✅ Active |
| 09 | [09-changelog-integration.md](./09-changelog-integration.md) | Changelog format, CI extraction, release body assembly, terminal display | ✅ Active |
| 10 | [10-code-signing.md](./10-code-signing.md) | SignPath integration, feature-flag gating, signature verification | ✅ Active |
| 11 | [11-self-update-mechanism.md](./11-self-update-mechanism.md) | Generic CLI self-update blueprint: deploy path, rename-first, handoff, cleanup | ✅ Active |
| 12 | [12-version-and-help.md](./12-version-and-help.md) | Version display, help system, command-level docs, CI verification | ✅ Active |
| 13 | [13-environment-variable-setup.md](./13-environment-variable-setup.md) | `env` command: persistent variables, PATH registration, auto-home | ✅ Active |
| 14 | [14-release-body-and-changelog.md](./14-release-body-and-changelog.md) | Changelog extraction, release body template, asset matrix assembly | ✅ Active |
| 15 | [15-terminal-output-standards.md](./15-terminal-output-standards.md) | Output formatting: icons, tables, progress, errors, CI summaries | ✅ Active |
| 16 | [16-binary-icon-branding.md](./16-binary-icon-branding.md) | Windows binary icon embedding via `go-winres`: icon, manifest, version info | ✅ Active |
| 17 | [17-release-pipeline-issues-rca.md](./17-release-pipeline-issues-rca.md) | 🔴 Root-cause analysis ledger of CI/CD failures and standing rules | ✅ Active |
| 18 | [18-known-issues-and-fixes.md](./18-known-issues-and-fixes.md) | ❓ Common pipeline error patterns and remediation paths | ❓ Pending Owner Review |
| 19 | [19-lint-gating-rules.md](./19-lint-gating-rules.md) | ❓ Strict lint gating strategies and baseline diff rules | ❓ Pending Owner Review |
| 20 | [20-ai-release-synchronization.md](./20-ai-release-synchronization.md) | ❓ Multi-repository release synchronization workflows | ❓ Pending Owner Review |
| 21 | [21-changelog-awk-integration.md](./21-changelog-awk-integration.md) | ❓ AWK-free JSON-based changelog parsers in CI | ❓ Pending Owner Review |
| 22 | [22-query-wrapper-python-ts.md](./22-query-wrapper-python-ts.md) | ❓ Query wrapper cross-language CI pipelines | ❓ Pending Owner Review |
| 23 | [23-strict-enum-enforcement.md](./23-strict-enum-enforcement.md) | ❓ Automated enum convention linters in CI | ❓ Pending Owner Review |
| 24 | [24-rca-release-skew.md](./24-rca-release-skew.md) | ❓ RCA on version skew during parallel release jobs | ❓ Pending Owner Review |
| 25 | [25-blue-green-deployment.md](./25-blue-green-deployment.md) | Zero-downtime blue/green deployment strategy | ✅ Active |
| 26 | [26-flaky-test-quarantine.md](./26-flaky-test-quarantine.md) | Flaky test detection and automated quarantine pattern | ✅ Active |
| 27 | [27-contract-testing.md](./27-contract-testing.md) | Microservice and API consumer contract testing | ✅ Active |
| 28 | [28-e2e-testing-pattern.md](./28-e2e-testing-pattern.md) | End-to-end integration and smoke test runner patterns | ✅ Active |
| 99 | [99-consistency-report.md](./99-consistency-report.md) | Consistency validation report for CI/CD workflows | ✅ Active |

---

## Subfolders

* [01-browser-extension-deploy/01-index.md](./01-browser-extension-deploy/01-index.md) — Chrome extension automated build and packaging.
* [02-go-binary-deploy/01-index.md](./02-go-binary-deploy/01-index.md) — Cross-platform Go binary packaging, signing, and release.
* [03-reusable-ci-guards/01-index.md](./03-reusable-ci-guards/01-index.md) — Six modular, language-agnostic CI guards.
* [04-gitmap-pipeline/01-index.md](./04-gitmap-pipeline/01-index.md) — ❓ Gitmap extended workflow references.
