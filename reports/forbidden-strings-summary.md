# Forbidden-Strings Summary Report

- **Generated (UTC):** 2026-05-05T14:26:32.000Z
- **Source:** `python3 linter-scripts/forbidden-strings-summary.py --markdown`
- **Config:** [`linter-scripts/forbidden-strings.toml`](../linter-scripts/forbidden-strings.toml)
- **Scope:** Last local CI invocation of the forbidden-strings linter against the working tree.


## `STALE-REPO-SLUG`

- **Description:** Pre-renumber repo slug (coding-guidelines-v1..v14). Current canonical slug is coding-guidelines-v20.
- **Pattern:** `coding-guidelines-v(1[0-4]|[1-9])(?!\.\d)\b`
- **Canonical replacement:** `coding-guidelines-v20`
- **Status:** ✅ clean (0 findings)

## `STALE-MODULE-PATH`

- **Description:** Stale module path reference (movie-cli-v1). Current canonical namespace is movie-cli-v2.
- **Pattern:** `movie-cli-v1\b`
- **Canonical replacement:** `movie-cli-v2`
- **Status:** ✅ clean (0 findings)

## `LEGACY-CDN-DOMAIN`

- **Description:** Legacy CDN domain (cdn.riseup-asia.com). Current canonical domain is cdn.riseup.asia.
- **Pattern:** `cdn\.riseup-asia\.com`
- **Canonical replacement:** `cdn.riseup.asia`
- **Status:** ✅ clean (0 findings)

## `SPEC19-NO-IMPL`

- **Description:** Forbidden Spec/19 implementation phrases (phase 1 implementation, begin coding, ship the service, starter skeleton, etc.). spec/19-main-worker-service is spec-only.
- **Pattern:** `(?i)\b(phase\s*1\s*implementation|begin\s*coding|ship\s*the\s*service|starter\s*skeleton|begin\s*spec[/\s-]*19\s*phase\s*1|start\s*implementing\s*spec[/\s-]*19)\b`
- **Status:** ✅ clean (0 findings)

✅ All rules clean — no findings.
