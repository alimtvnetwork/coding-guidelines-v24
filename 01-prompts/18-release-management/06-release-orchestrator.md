# Automated Release Orchestrator & Branch Lifecycle — Release Management (must follow)


> **Prompt Version:** 2.1.0
> **Synchronization:** Main Meta-Repo & Connected Workspaces

Trigger phrases: `orchestrate release`, `release orchestrator`, `bump with release branch`, `release and tag`, `automated release`.

If the user requests a release, version bump with release branching, or automated tag orchestration, execute this prompt.

---

## MUST FOLLOW NON-NEGOTIABLE

Listen, past runs of these turns have been sloppy: wrong step counts, partial task lists dumped into chat instead of files, plans and session summaries half-filled with placeholders, folders skimmed, open ambiguities ignored, CI/CD issues forgotten, user commands dropped, coding guidelines bypassed, detailed specs chopped, uppercase README files left uncorrected, `.lovable/memory/` created by accident, `strictly-avoid.md` overwritten, and explicit user instructions softened after being told not to. Read the whole codebase, confirm root `readme.md` is strictly lowercase, find the root cause in one sentence, capture commands, issues, and pending tasks without omitting a single item, write the spec files and memory files in the right paths, update every index in the same turn, sync `readme.md` with `what-to-read.md`, preserve detailed specs verbatim with zero truncation, run builds and full unit tests, group commits with clear messages, and push everything to git before ending. Going deep IS the job. Violating this is auto-reject on the same tier as RULE 0. Avoid stupidity and being careless.

---

## RULE 0: Version Calculation Standard

1. Read the canonical version source for this repo (`version.json` or `package.json`).
2. Default bump tier is **MINOR**: `MAJOR.MINOR.PATCH` becomes `MAJOR.(MINOR+1).0`. PATCH MUST reset to `0`.
3. Only bump PATCH if user explicitly specified `patch`.
4. Only bump MAJOR if user explicitly specified `major` or breaking change.
5. State previous and new versions explicitly in output before touching files.

---

## Master Architecture: Heavy-Lifting Release Script (`03-ai-scripts/29-release-orchestrator.py`)

All release operations (version bumping, commit creation, release branching, git tagging, and branch reversion) MUST be executed through the centralized heavy-lifting script in the AI scripts directory:

```bash
python 03-ai-scripts/29-release-orchestrator.py --tier <minor|patch|major>
```

### The Reuse-or-Create Protocol

1. **Check for Script:** Inspect if `03-ai-scripts/29-release-orchestrator.py` exists on disk.
2. **If Script Exists:**
   - Run the script directly with the required arguments:
     ```bash
     python 03-ai-scripts/29-release-orchestrator.py --tier minor --scope "<Release summary>"
     ```
3. **If Script Is Missing:**
   - The AI agent MUST immediately construct `03-ai-scripts/29-release-orchestrator.py` inside `03-ai-scripts/` using Python standard libraries (`argparse`, `json`, `os`, `re`, `subprocess`, `sys`, `pathlib`, `datetime`).
   - Register it in `03-ai-scripts/01-index.md` under slot 29.
   - Execute the freshly created script to fulfill the release.

---

## Standalone Bump-Version Recovery

If the repository does not have a project-level bump version script (such as `scripts/bump-version.mjs` or `.lovable/release/bump_versions.py`):
1. **Autonomous Bootstrap:** The orchestrator script must not fail or halt. It MUST autonomously bootstrap `scripts/bump-version.mjs` (or perform in-place updates directly).
2. **In-Place File Updates:** The orchestrator must update:
   - `version.json`: Update `"version"` to `next_version` and `"releaseDate"` to today's UTC date (`YYYY-MM-DD`).
   - `package.json`: Update `"version"` to `next_version`.
   - `readme.md`: Pin the new release version in badges, install snippets, and header versions.
   - `changelog.md`: Prepend the release entry under `# Changelog` with the release version, UTC date, and changelog bullets.

---

## Git Release Lifecycle & Original Branch Invariant

The release orchestrator strictly implements this end-to-end Git workflow:

```
[Start on original_branch (e.g. feature/my-work)]
                   │
                   ▼
1. Detect & Store original_branch (git rev-parse --abbrev-ref HEAD)
                   │
                   ▼
2. Bump SemVer across version.json, package.json, changelog.md, readme.md
                   │
                   ▼
3. Stage & Commit on current branch: `release: vX.Y.Z <scope>`
                   │
                   ▼
4. Create Release Branch: `release/vX.Y.Z` pointing to the commit
                   │
                   ▼
5. Create Annotated Git Tag: `vX.Y.Z` on that release commit
                   │
                   ▼
6. Push: Push `release/vX.Y.Z` and `vX.Y.Z` to remote origin
                   │
                   ▼
7. MANDATORY REVERT: Checkout back to `original_branch`
                   │
                   ▼
[Finish: Active branch is verified to be `original_branch`]
```

### Critical Rules for Branch Reversion:
- **NEVER assume `main` or `master`:** Releases may be triggered from feature branches, bugfix branches, or release candidates. The script MUST record the starting branch and revert to that exact branch.
- **Fail-Safe Restoration:** Even if pushing or downstream steps fail, the script's `finally:` block MUST ensure the working tree is safely checked back out to `original_branch`.

---

## Hard Rules & Pre-Flight Checks

- [ ] Repository must start with a clean working tree (`git status`). Commit or stash uncommitted changes before orchestrating.
- [ ] No explicit boolean checks (`if is_success == True:` is banned; use `if is_success:`).
- [ ] All filenames must be strictly lowercase (e.g. `readme.md`, `changelog.md`).
- [ ] Relative Git paths only (no `file:///` URIs or absolute filesystem paths).
- [ ] The release commit and tag MUST be pushed to Git.
- [ ] The working tree must be on the original branch when the task completes.

---

## Actionable Execution Checklist

- [ ] 1. Identify starting branch: `git rev-parse --abbrev-ref HEAD`.
- [ ] 2. Check for `03-ai-scripts/29-release-orchestrator.py`. If missing, create it.
- [ ] 3. Check for bump version script (`scripts/bump-version.mjs` or `.lovable/release/bump_versions.py`). If missing, bootstrap or let orchestrator handle in-place updates.
- [ ] 4. Run `python 03-ai-scripts/29-release-orchestrator.py --tier <tier> --scope "<scope>"`.
- [ ] 5. Verify that commit `release: vX.Y.Z` exists.
- [ ] 6. Verify that branch `release/vX.Y.Z` points to the release commit.
- [ ] 7. Verify that tag `vX.Y.Z` exists on the release commit.
- [ ] 8. Verify that release branch and tag were pushed to remote.
- [ ] 9. Verify that the active git branch is restored to the starting branch (`git branch --show-current`).
- [ ] 10. Output release summary detailing starting branch, version bump, release branch, tag, and restored active branch.
