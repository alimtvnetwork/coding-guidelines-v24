# PHP Standards — Changelog (AI Execution Prompt)

> **/goal** Record all architectural revisions, strict typing standards, and static analysis integrations across PHP coding guidelines.
> **/learn** Internalize modern PHP standards: strict types (`declare(strict_types=1)`), typed properties, enum standards, and spacing requirements.

**Version:** 3.2.0
**Last Updated:** 2026-04-16
**Status:** Active
**AI Confidence:** Production-Ready
**Ambiguity:** None

---

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Document updates to PHP coding guidelines, decomposition of reference specs, and code example fixes.
- [ ] `/learn` Ensure all relative links to PHP subfolders (`07-php-standards-reference/`) resolve correctly.
- [ ] `/goal` Verify new version entries conform to project-wide SemVer conventions.
- [ ] `/learn` Audit formatting with `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/04-php --check-only`.

. **CRITICAL AI INSTRUCTION:** Active AI execution record for PHP standards. Log all revisions to PHP guidelines here.

---

All notable changes to the PHP Standards specification are documented here.

---

## v2.1.0 — 2026-03-31

### Changed

- `07-php-standards-reference.md` split into subfolder (5 files, max 252 lines — down from 840)
- Fixed spacing violations in code examples

---

## v2.0.0 — 2026-03-09

### Global Version Bump

Project-wide major version increment (+1.0.0) applied to all specification files in `03-coding-guidelines/04-php`.

#### Changed

- All spec files received a major version bump and date update to 2026-03-09.
- Part of a global effort spanning ~638 files across all 30+ spec folders, establishing a new project-wide versioning baseline.

---

*Keep this file updated when specs change.*

---

## Verification & Acceptance Criteria

_Auto-generated section — see `02-spec/02-coding-guidelines/04-php/97-acceptance-criteria.md` for the full criteria index._

### AC-CG-LOG-PHP: PHP Standards Changelog Conformance

**Given** The PHP standards changelog file in `02-spec/02-coding-guidelines/04-php/98-changelog.md`.
**When** Audited by repository linters and guideline autofixers.
**Then** The file strictly adheres to the 4-part prompt anatomy, accurately logs PHP strict typing, backed enum, and PSR-4 revisions, and passes linter checks with exit code 0.

**Verification command:**
```bash
python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/04-php --check-only
```
**Expected:** exit 0. Zero violations detected.
