# CI/CD Issues

> **Version:** 1.2.0  
> **Updated:** 2026-10-27  
> **Purpose:** CI/CD issue tracking, validator findings registry, and pipeline root cause analysis (RCA).

---

## Purpose

Tracks every CI/CD validator finding (`CODE-RED-*`, `STYLE-*`, `BOOL-NEG-*`) encountered during sessions, what fixed it, and what to avoid re-introducing. Source: `linters-cicd/run-all.sh` and `03-ai-scripts/06-cicd-local-runner.py` output.

Active or in-flight CI/CD issues are documented in this directory (`.ai-memory/cicd-issues/`). Once verified and solved, issue post-mortems and RCAs are archived into `.ai-memory/resolved-issues/`.

---

## Resolved CI/CD Issues Index

| # | Title | Status | Rules | Date | Link |
|---|-------|--------|-------|------|------|
| 01 | BOOL-NEG-001 Pipeline Smoke Test Validation | ✅ Solved | BOOL-NEG-001 | 2026-04-25 | [.ai-memory/resolved-issues/01-bool-neg-001-pipeline-untested.md](.ai-memory/resolved-issues/01-bool-neg-001-pipeline-untested.md) |
| 02 | Quickstart docs example: boolean negatives + style violations | ✅ Solved | CODE-RED-023, STYLE-001, STYLE-004 | 2026-04-23 | [.ai-memory/resolved-issues/02-readme-boolean-negatives-and-style.md](.ai-memory/resolved-issues/02-readme-boolean-negatives-and-style.md) |
| 03 | InstallSection.tsx 337-line component split | ✅ Solved | STYLE-005 | 2026-04-23 | [.ai-memory/resolved-issues/03-installsection-component-split.md](.ai-memory/resolved-issues/03-installsection-component-split.md) |
| 04 | fuzzyMatch.ts: function length + negative guards | ✅ Solved | CODE-RED-004, CODE-RED-012, CODE-RED-023 | 2026-04-23 | [.ai-memory/resolved-issues/04-fuzzymatch-function-length-negative-guards.md](.ai-memory/resolved-issues/04-fuzzymatch-function-length-negative-guards.md) |
| 05 | useSearchKeyboard.ts raw `!` operator | ✅ Solved | CODE-RED-023 | 2026-04-23 | [.ai-memory/resolved-issues/05-usesearchkeyboard-raw-not-operator.md](.ai-memory/resolved-issues/05-usesearchkeyboard-raw-not-operator.md) |
| 06 | Bulk STYLE-001 / STYLE-004 newline violations across markdown highlighter | ✅ Solved | STYLE-001, STYLE-003, STYLE-004 | 2026-04-23 | [.ai-memory/resolved-issues/06-markdown-highlighter-newline-violations.md](.ai-memory/resolved-issues/06-markdown-highlighter-newline-violations.md) |
| 07 | Version drift after `package.json` bump (forgot `npm run sync`) | ✅ Solved | version-drift | 2026-04-23 | [.ai-memory/resolved-issues/07-version-drift-after-package-bump.md](.ai-memory/resolved-issues/07-version-drift-after-package-bump.md) |
| 08 | Cross-spec missing-file checker false-positives in `26-spec-outsides` | ✅ Solved | missing-file | 2026-04-23 | [.ai-memory/resolved-issues/08-cross-spec-missing-file-link-checker.md](.ai-memory/resolved-issues/08-cross-spec-missing-file-link-checker.md) |
| 09 | JavaScript Template Literal SyntaxError in `generate-bundle-installers.mjs` | ✅ Solved | build-failure | 2026-08-30 | [.ai-memory/resolved-issues/09-bundle-installers-template-literal-syntax-error.md](.ai-memory/resolved-issues/09-bundle-installers-template-literal-syntax-error.md) |
| 10 | Codegen Determinism Drift After Spec Renumbering & Shell-to-Python Migration | ✅ Solved | codegen-drift | 2026-08-30 | [.ai-memory/resolved-issues/10-codegen-determinism-spec-renumbering-drift.md](.ai-memory/resolved-issues/10-codegen-determinism-spec-renumbering-drift.md) |

---

## Aggregate Stats (Last Clean Run — v3.81.0)

- Violations: **0**
- Files scanned: 612
- Lines scanned: 131,966

---

*CI/CD issues index — v1.2.0*
