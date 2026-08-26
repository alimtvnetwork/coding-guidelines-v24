---
name: Version Source of Truth Standard
description: version.json is the single source of truth for repository versioning. Changing version in any repo is done solely in version.json and propagated via sync.
type: standard
---

# Version Source of Truth Standard (ersion.json)

**Canonical spec:** spec/01-spec-authoring-guide/17-version-schema.md
**Rule:** Hard Rule 19 in spec/17-consolidated-guidelines/31-compiled-simple-coding-guidelines.md

## 1. Single Source of Truth

Every repository MUST maintain a root-level ersion.json file.
This file is the single, canonical source of truth for:
- Repository Version (Version)
- Repository Name / Slug (RepoSlug)
- Repository URL (RepoUrl)
- Last Commit SHA (LastCommitSha)
- Author Attribution (Authors)
- Integrated Specs & Prompt Architect version tracking (codingGuideline, promptArchitectByRiseupAsia)

## 2. Self-Explaining JSON Structure

The root ersion.json file contains explicit _purpose and _instructions metadata at the top:
`json
{
  "_purpose": "version.json is the single canonical source of truth for repository versioning, releases, metadata, and cross-repo dependencies. All tools, packages, scripts, CI/CD pipelines, manifests, and AI agents must read this file for version information. To bump or change the version in this repository, update the 'Version' field in this file and run the project sync script (e.g. 'npm run sync' or 'npm run sync:version').",
  "_instructions": {
    "versionSourceOfTruth": "Every codebase must use version.json at the repository root as the single source of truth for version numbers.",
    "howToChangeVersion": "To change the version of the repository, edit the 'Version' field in version.json, then run 'npm run sync' (or the project's sync command) to propagate it across all files, docs, badges, and manifests.",
    "prohibition": "Never hardcode version numbers in code, scripts, or documentation. Always read from or sync with version.json."
  },
  "Version": "6.19.0",
  ...
}
`

## 3. How to Change or Bump the Version

To release or bump the version of ANY repository:
1. Edit the "Version" (and "version") string in ersion.json at the root of the repository.
2. Run the repository's synchronization script:
   `ash
   npm run sync
   # or node scripts/sync-version.mjs
   `
3. The sync script automatically propagates the new version to:
   - package.json
   - Manifests & metadata files
   - Documentation & Markdown headers / stamps
   - Health score & README badges
4. Commit the changes and tag the release:
   `ash
   git commit -m "release: vX.Y.Z"
   git tag vX.Y.Z
   git push && git push --tags
   `

## 4. Strict Prohibitions (Non-Negotiable)

- ❌ **NEVER** hardcode version numbers across random source files, constants, or markdown bodies.
- ❌ **NEVER** create competing version files.
- ❌ **NEVER** allow version drift across files; ersion.json always wins.

## 5. Install & Update Transfer

During bundle install and update phases, the installer automatically transfers .lovable/memory/ into target repositories so AI agents, developers, and CI pipelines immediately inherit this standard without manual configuration.
