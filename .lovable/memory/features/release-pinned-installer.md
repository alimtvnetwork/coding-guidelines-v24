---
name: Release-Pinned Installer Scripts
description: Separate release-install.ps1/sh additive to install.ps1/sh; pins exact version, never queries latest, never crosses v1/v2/v3 repo boundaries.
type: feature
---

# Release-Pinned Installer

**Files:** `release-install.ps1` and `release-install.sh` at repo root (additive — does NOT replace existing `install.ps1` / `install.sh`).

**Purpose:** Linked from each GitHub Release page. Installs exactly the version that page represents. No latest lookup, no auto-upgrade, no cross-repo drift.

**Version source priority:**
1. `-Version` (PS) / `--version` (sh) CLI override
2. `VERSION_PLACEHOLDER` baked in at release time via `sed`
3. None resolved → fatal exit 1 (no fallback to latest)

**Hard MUST-NOT list:**
- No call to `api.github.com/.../releases/latest`
- No `--repo` / `-Repo` flag (prevents v1/v2/v3 drift)
- No upgrade nags / "newer version available" notices
- No fallback to latest on 404 — fail loud

**Flag:** `-NoUpdate` (PS) / `--no-update` (sh), defaults to `true`. Reserved; toggling currently a no-op since release script never queries latest.

**Repo pinning:** `REPO_PLACEHOLDER` baked in at release time. If unreplaced at runtime → exit 1.

**Pipeline:** Same `sed` replacement step as existing install scripts; published as release asset on every tag.

**Spec:** `spec/14-update/25-release-pinned-installer.md` v1.0.0
