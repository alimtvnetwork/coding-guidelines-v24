# 15 — Prompt Architect Version Tracking in version.json

## Objective

Standardize how the Prompt Architect installation is tracked inside any downstream
repository's `version.json`, via a required `promptArchitectByRiseupAsia` block.

## Status: DONE

---

## Deliverables & Checklist

### D1 — Template file

- [x] Create `prompt-version.template.json` at repo root
- Fields required:
  - `promptArchitectByRiseupAsia.author` — name, title, url
  - `promptArchitectByRiseupAsia.sourceRepository` — url
  - `promptArchitectByRiseupAsia.installedAt` — ISO-8601 string (set by installer)
  - `promptArchitectByRiseupAsia.version` — semver string (from version.json at install time)
  - `promptArchitectByRiseupAsia.lastCommit` — sha
  - `promptArchitectByRiseupAsia.fileMapping` — array of { source, target } objects

### D2 — Installer smart-merge update (generate-bundle-installers.mjs)

- [x] Bash installer: inject `promptArchitectByRiseupAsia` block when copying `.lovable/prompts/`
- [x] PowerShell installer: same
- [x] Both must read `prompt-sync-config.json` mappings to build `fileMapping` array
- [x] Must NOT overwrite existing `promptArchitectByRiseupAsia` block — only add/update

### D3 — Memory documentation

- [x] Append entry to `.lovable/memory/index.md` explaining this tracking block

### D4 — Release

- [x] All changes committed, `npm run sync:version`, push

---

## Author Attribution (non-negotiable exact values)

```json
"author": {
  "name": "Md. Alim Ul Karim",
  "title": "Chief Software Engineer",
  "url": "https://github.com/aukgit/alim.karim.profile"
}
```

## Source Repository

```
https://github.com/alimtvnetwork/prompt-architect-v2
```

## Coding Rules (non-negotiable)

- JSON keys: camelCase
- No magic strings — all values read from config files
- No temp scripts committed
- All changes pushed before marking DONE
