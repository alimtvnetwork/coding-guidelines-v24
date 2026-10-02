# Consistency Report: Security Guidelines (AI Execution Prompt)

> **/goal** Maintain and verify structural consistency, inventory completeness, and cross-reference integrity for Security Guidelines specifications.
> **/learn** Audit numeric sequencing, kebab-case file naming, table completeness, and acceptance criteria synchronization.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Verify 100% presence and integrity of all registered specification files in this directory.
- [ ] `/learn` Ensure zero broken relative markdown links across all module documentation.
- [ ] `/goal` Verify all specification files contain active AI execution headers and testable acceptance criteria.
- [ ] `/learn` Validate zero absolute paths via `python linter-scripts/check-relative-paths.py`.

. **CRITICAL AI INSTRUCTION:** This specification is an active AI execution directive. All code generated or modified must strictly follow the rules below.

**Version:** 3.2.0
**Last Updated:** 2026-04-16

---

## Module Health

| Criterion | Status |
|-----------|--------|
| `readme.md` present | ✅ |
| `99-consistency-report.md` present | ✅ |
| Lowercase kebab-case naming | ✅ |
| Unique numeric sequence prefixes | ✅ |

**Health Score:** 100/100 (A+)

---

## File Inventory

| # | File | Status |
|---|------|--------|
| 00 | `readme.md` | ✅ Present |
| 01 | `01-axios-version-control/` | ✅ Present (subfolder, 4 files) |
| 99 | `99-consistency-report.md` | ✅ Present |

**Total:** 2 files + 1 subfolder

---

## Cross-Reference Validation

All internal links verified valid. ✅

---

## Validation History

| Date | Version | Action |
|------|---------|--------|
| 2026-04-02 | 1.0.0 | Initial module created; moved Axios version control from `02-spec/01-app/` |

---

## Verification & Acceptance Criteria

### AC-CG-CONSISTENCY-SEC: Security Guidelines Consistency Report Conformance

**Given** Specification files under this module directory.
**When** Linters and CI/CD consistency scripts audit the module inventory.
**Then** All registered files are present, filenames strictly lowercase, and 100% relative paths verified with exit code 0.

**Verification command:**
```bash
python linter-scripts/check-relative-paths.py
```
**Expected:** exit 0. Zero absolute paths or file:/// URIs.
