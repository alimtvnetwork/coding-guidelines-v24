# Subtask 07: Terminal E2E & Minor Version Bump Release Ceremony

## Objective
Run full local tests and linters, compile and deploy updated GitMap binary, bump version to `v6.88.0` in manifests, author release notes, commit, push upstream to GitHub, and create Git tag.

## Target Files
- `version.json`
- `package.json`
- `prompt-version.template.json`
- `readme.md`
- `changelog.md`
- `.ai-memory/release/release-notes-v6.88.0.md`

## Steps
1. Run `go test` across modified GitMap packages.
2. Build `gitmap.exe` and swap into active CLI location.
3. Commit and push in `d:/work/gitmap`.
4. Bump version in `coding-guidelines` using `bump_versions.py --type minor`.
5. Update `changelog.md` and write release notes.
6. Commit with `gitmap cpr` and push tag `v6.88.0`.
