# Plan Tracker

**Updated:** 2026-04-19  
**Version:** 3.5.0

> Mirrors `.lovable/plan.md`. The canonical roadmap lives there.

---

## Completed Plans (recent)

| # | Task | Date |
|---|------|------|
| 23 | Created write-memory prompt v3.3 | 2026-04-16 |
| 24 | FAQ features in code (suppression parsing, baseline flags, TOML) | 2026-04-19 |
| 25 | STYLE-099 SuppressionWithoutReason synthetic finding | 2026-04-19 |
| 26 | Created `99-troubleshooting.md` (CICD) | 2026-04-19 |
| 27 | Performance impl: middle-out walker, `--jobs`, `--check-timeout`, TOOL-TIMEOUT (v3.12.0) | 2026-04-19 |
| 28 | `--version` flag on every check script (v3.13.0) | 2026-04-19 |
| 29 | `01-naming-conventions.md` v3.3.0 — Rule 2 clarification + Rule 9 | 2026-04-19 |
| 30 | `02-schema-design.md` v3.3.0 — §6 Mandatory Descriptive Columns | 2026-04-19 |
| 31 | BOOL-NEG-001 linter (v3.14.0) | 2026-04-19 |
| 32 | Inverted-field codegen tool (Go + PHP + TS) | 2026-04-19 |
| 33 | Cross-linked Rule 9 from boolean-principles + no-negatives (v2.2.0) | 2026-04-19 |
| 34 | `01-naming-conventions.md` v3.4.0 — Rule 8 three-bucket table | 2026-04-19 |
| 35 | `01-naming-conventions.md` v3.5.0 — Rules 10/11/12 (Description/Notes/Comments) | 2026-04-19 |
| 36 | Restructured `.lovable/` to single-file convention; write-memory prompt v1.0.0 | 2026-04-19 |

(For dates 2026-04-02 → 2026-04-16, see `.lovable/plan.md` Completed Plans Historical.)

---

## Pending Plans

| # | Task | Priority | Notes |
|---|------|----------|-------|
| 01 | Smoke-test BOOL-NEG-001 in `run-all.sh` | Medium | Standalone passes |
| 02 | Go-aware BOOL-NEG-001 variant | Medium | `embed.FS` SQL |
| 03 | Unit tests for BOOL-NEG-001 | Medium | |
| 04 | Round-trip tests for codegen inversion table | Medium | |
| 05 | Wire codegen into CI (`git diff --exit-code`) | Medium | Highest-value next |
| 06 | Linter for missing `Description`/`Notes`/`Comments` | Medium | Rules 10–12 |
| 07 | Strengthen BOOL-NEG-001 with replacement hints | Low | |
| 08 | Cross-link link-checker over `spec/` | Low | |
| 09 | Align `02-schema-design.md` §6 wording with naming v3.5.0 | Low | |
| 10 | Publish the app | Medium | |
| 11 | E2E browser testing | Medium | |
| 12 | Mobile responsiveness | Low | |
| 13 | Update `99-consistency-report.md` | Medium | |
| 14 | Expand sub-90% guidelines | Medium | |
| 15 | Bootstrap PHP plugins for CODE-RED-001..004 | Medium | |
| 16 | `--total-timeout` + per-file 2s parse timeouts | Low | |
| 17 | `exclude-paths` glob support | Medium | |

---

*Plan tracker — v3.5.0 — 2026-04-19*
