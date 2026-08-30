# Release Method Blueprint

This document defines exactly where and how versions are tracked and updated in this repository. Any AI performing a release MUST refer to this file to understand the architecture before modifying version strings.

> [!IMPORTANT]
> **Baseline Template Warning**
> The `bump_versions.py` script and this document are shipped out-of-the-box via the prompt architect installer as a baseline. When starting a new project, AIs **MUST** update the `FILES_TO_BUMP` array in the Python script and the list below to explicitly match the files of the new repository.

## Versioning Scheme

The project follows Semantic Versioning (Major.Minor.Patch).
The **canonical** source of truth for the current version is `version.json` (or `package.json`).

## Files Requiring Version Bumps

When a release occurs, the version string (both plain `1.X.X` and prefixed `v1.X.X`) MUST be updated in the following explicit file paths. **Do NOT run global unbounded searches.**

- `version.json`
- `package.json`
- `prompt-version.template.json`
- `readme.md`
- `.lovable/coding-guidelines/coding-guidelines.md`
- `linter-scripts/validate-guidelines.go`
- `linter-scripts/validate-guidelines.py`
- `spec/14-update/28-worker-push-instruction.md`
- `spec/17-consolidated-guidelines/31-compiled-simple-coding-guidelines.md`
- `spec/19-main-worker-service/diagrams/seq-push-update.mmd`
- `spec/19-main-worker-service/10-worker-bootstrap-protocol.md`
- `spec/19-main-worker-service/14-rbac-and-status-seed.md`
- `spec/19-main-worker-service/15-tunable-constants.md`
- `spec/19-main-worker-service/16-update-channels.md`
- `spec/19-main-worker-service/25-inherited-rules.md`

## Automation Script

The python script `.lovable/release/bump_versions.py` automates this process. It has two operating modes:

### Standard Mode

`python .lovable/release/bump_versions.py --type <major|minor|patch>`
Calculates the new version and safely applies regex replacements only to the files listed above. It does not run any git commands.

### Full Release Mode

`python .lovable/release/bump_versions.py --type <major|minor|patch> --create-release`
Performs the version replacement, then automates the entire git release process. It prevents you from getting stranded on a branch by executing this complete loop:

1. Detects your current active branch (e.g., `main`).
2. Creates and checks out a new branch (`release/vX.Y.Z`).
3. Commits the changes.
4. Creates a git tag (`vX.Y.Z`).
5. Pushes the branch and tag to origin.
6. Detects if `gh` (GitHub) or `glab` (GitLab) CLI is installed and automatically publishes the official platform release.
7. **Checks out the original branch (`main`).**
8. **Merges the release branch back into it.**
9. **Pushes the original branch to origin.**
