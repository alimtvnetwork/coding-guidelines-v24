# Reading `version.json`

How application code, build scripts, and shell tooling should consume the
canonical repo manifest defined in
[`spec/01-spec-authoring-guide/17-version-schema.md`](../spec/01-spec-authoring-guide/17-version-schema.md).

> **TL;DR**
> 1. Always read from the repo-root `version.json`. Never hardcode `Version`,
>    `Title`, `RepoSlug`, `RepoUrl`, or `LastCommitSha`.
> 2. Use the **PascalCase** keys (`Version`, `Title`, `RepoSlug`,
>    `LastCommitSha`, …). Treat camelCase keys (`version`, `git`, `stats`,
>    `folders`) as **deprecated transitional** fields — they will be removed.
> 3. If `version.json` is missing or unparseable, **log a warning and fall
>    back to safe defaults**. Never crash the host application.

---

## 1. Canonical PascalCase shape

```json
{
  "Version": "4.24.0",
  "Title": "Coding Guidelines",
  "RepoSlug": "coding-guidelines",
  "RepoUrl": "https://github.com/mahin/coding-guidelines",
  "LastCommitSha": "8dc2161da12441992726ccc0e7788f60527bced1",
  "Description": "Cross-language coding standards…",
  "Authors": [
    { "Name": "…", "Urls": ["…"], "Role": "PrimaryAuthor", "Background": "…" }
  ]
}
```

`Role` is a **closed enum**: `PrimaryAuthor`, `Contributor`, `Maintainer`,
`Reviewer`, `Sponsor`. Anything else is a schema violation and will fail
the `validate-version-json` CI gate.

---

## 2. Safe-fallback contract

When `version.json` is missing or unparseable, every reader MUST:

1. Log a warning identifying the file path and the cause.
2. Substitute safe defaults:
   - `Version` → `"0.0.0"`
   - `Title` → `RepoSlug` (or the package name)
   - `LastCommitSha` → `""`
   - `Authors` → `[]`
3. Keep running. Never crash the host application solely because the
   manifest is missing.

This rule is enforced by AC-VS-008 in the schema spec.

---

## 3. Reading from React / TypeScript

The Vite app imports `version.json` directly so it ships with the bundle —
no fetch, no race, no fallback path needed at runtime:

```ts
// src/hooks/useDashboardData.ts
import versionJson from "../../version.json";
import type { VersionInfo } from "@/types/dashboard";

const version = versionJson as VersionInfo;
console.log(version.Version, version.LastCommitSha);
```

For browser code that loads `version.json` over HTTP (for example a
deployed dashboard), use the safe-fallback pattern:

```ts
const FALLBACK = {
  Version: "0.0.0",
  Title: "Unknown",
  RepoSlug: "unknown",
  RepoUrl: "",
  LastCommitSha: "",
  Description: "",
  Authors: [],
} as const;

async function loadVersion() {
  try {
    const r = await fetch("/version.json", { cache: "no-store" });
    if (r.ok === false) throw new Error(`HTTP ${r.status}`);
    return await r.json();
  } catch (err) {
    console.warn("[version.json] falling back to defaults:", err);
    return FALLBACK;
  }
}
```

---

## 4. Reading from Node scripts

```js
// scripts/example.mjs
import { readFileSync, existsSync } from "node:fs";
import { resolve } from "node:path";

const PATH = resolve(process.cwd(), "version.json");

function loadVersion() {
  if (existsSync(PATH) === false) {
    console.warn(`[version.json] missing at ${PATH} — using defaults`);
    return { Version: "0.0.0", Title: "unknown", LastCommitSha: "" };
  }
  try {
    return JSON.parse(readFileSync(PATH, "utf8"));
  } catch (err) {
    console.warn(`[version.json] unparseable: ${err.message} — using defaults`);
    return { Version: "0.0.0", Title: "unknown", LastCommitSha: "" };
  }
}

const v = loadVersion();
console.log(`v${v.Version} @ ${v.LastCommitSha.slice(0, 7) || "no-sha"}`);
```

---

## 5. Reading from Bash (installers, CI)

`jq` is the preferred reader. Always supply a default with `// "fallback"`
so a missing key never expands to `null`:

```bash
VERSION=$(jq -r '.Version // "0.0.0"'           version.json 2>/dev/null || echo "0.0.0")
SLUG=$(jq    -r '.RepoSlug // "unknown"'        version.json 2>/dev/null || echo "unknown")
SHA=$(jq     -r '.LastCommitSha // ""'          version.json 2>/dev/null || echo "")

echo "Installing ${SLUG} v${VERSION} (${SHA:0:7})"
```

If `jq` is not available, fall back to Python (already used by
`release-install.sh`):

```bash
VERSION=$(python3 -c '
import json, sys
try:
  print(json.load(open("version.json")).get("Version", "0.0.0"))
except Exception:
  print("0.0.0")
')
```

---

## 6. Reading from PowerShell

```powershell
function Get-RepoVersion {
    $Path = Join-Path $PSScriptRoot "version.json"
    if (-not (Test-Path -LiteralPath $Path)) {
        Write-Warning "version.json missing at $Path — using defaults"
        return [pscustomobject]@{ Version = "0.0.0"; Title = "unknown"; LastCommitSha = "" }
    }
    try {
        return Get-Content -LiteralPath $Path -Raw | ConvertFrom-Json
    } catch {
        Write-Warning "version.json unparseable: $($_.Exception.Message)"
        return [pscustomobject]@{ Version = "0.0.0"; Title = "unknown"; LastCommitSha = "" }
    }
}

$V = Get-RepoVersion
Write-Host "v$($V.Version) @ $($V.LastCommitSha.Substring(0, 7))"
```

---

## 7. Do / Don't

| ✅ Do                                                           | ❌ Don't                                                |
|----------------------------------------------------------------|---------------------------------------------------------|
| Read PascalCase keys (`Version`, `RepoSlug`, `LastCommitSha`). | Read deprecated camelCase (`version`, `git.sha`).       |
| Wrap reads in try/catch and warn on failure.                   | Throw or `process.exit` because the file is missing.    |
| Use `jq` `//` defaults or Python `.get(..., default)`.         | Assume keys are always present (legacy repos may lag).  |
| Treat `version.json` as **read-only** in app code.             | Hand-edit `LastCommitSha` — the husky hook owns it.     |
| Run `npm run validate:version` in CI.                          | Skip the schema gate "just for this PR".                |

---

## 8. Related

- Schema spec: [`spec/01-spec-authoring-guide/17-version-schema.md`](../spec/01-spec-authoring-guide/17-version-schema.md)
- Mirror: [`spec/authoring-guideline/VersionSchema.md`](../spec/authoring-guideline/VersionSchema.md)
- Sync script: [`scripts/sync-version.mjs`](../scripts/sync-version.mjs) — owns writes to `version.json`
- Validator: [`scripts/validate-version-json.mjs`](../scripts/validate-version-json.mjs) — `npm run validate:version`
- CI gate: [`.github/workflows/validate-version-json.yml`](../.github/workflows/validate-version-json.yml)
