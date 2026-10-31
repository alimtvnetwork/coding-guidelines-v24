# Release Notes: v6.88.0

- **Release Date:** 2026-10-28
- **Release Tag:** `v6.88.0`
- **Specification:** `02-spec/21-app/33-coding-guideline-corrections-and-gitmap-skill-updates/`
- **Plan File:** `.ai-memory/plans/completed/33-coding-guideline-corrections-and-gitmap-skill-updates.md`

## Key Deliverables & Enhancements

### 1. GitMap Clone Bare Slug Resolution Fallback
- **Problem:** Running `gitmap clone <repo-name>` failed if the repository wasn't already in the local cache, requiring the user to provide a full URL.
- **Fix:** Enhanced `ResolveRepoSlug` in `cli/cmdclone/gh_resolver.go` to automatically resolve bare repository names against the active GitHub profile or owner (`resolveActiveOrProfileOwner()`), constructing and verifying `https://github.com/<owner>/<repo-name>` via `probeRemoteGitURL` before falling back to local search.
- **Verification:** Verified with `TestResolveRepoSlug_BareSlugFallback` and live dry-run cloning (`gitmap clone coding-guidelines --dry-run`).

### 2. Show Log Non-Git Folder Resilience
- **Problem:** Launching Show Log from Windows Explorer context menu on a folder that is not a Git repository dumped a fatal Go runtime stack trace (`[E9000:EXECUTION] not a git repository` at `cmdlogviewer/server.go:85`).
- **Fix:** Replaced the unhandled runtime fault with clean diagnostic output that scans child directories, detects active Git repositories, and lists them with clear commands to open them cleanly with zero stack traces and exit status 0.

### 3. Show Log UI Web App Manifest & Authentic Branding
- **Problem:** In Chrome/Edge `--app` mode, the window title bar showed a generic browser earth globe icon instead of the authentic GitMap logo.
- **Fix:** Added `/manifest.json` PWA endpoint and `<link rel="manifest" href="/manifest.json">` with embedded brand icons and theme colors.

### 4. GitMap Skills Uppercase Module Cleanup & Cluster Purge
- Replaced screaming uppercase commit examples (`SKILLS`, `CLONER`, `DOCS`, `RELEASE`) with Title Case / lowercase across `.agents/skills/gitmap/skill.md`, `.cursor/skills/gitmap/skill.md`, and `cli/cmd/llm/llm_skill.go`.
- Purged obsolete cluster commands (`cluster`, `sc`, `ssh`), OS dev clean, and VS Code repair from skills and documentation.

### 5. Follow Coding Guidelines Learning Loop & Skills Reduction Audit
- Enriched `follow-coding-guidelines/skill.md` with explicit `[/learn](slashCommand;learn)` integration, grounded 12-rule audit checklist, and pre-commit self-audit matrix.
- Completed an exhaustive audit of all 108 skills in `02-spec/21-app/33-coding-guideline-corrections-and-gitmap-skill-updates/02-skills-reduction-matrix.md`, establishing a 29-skill reduction proposal consolidating into `follow-cg-execute` and `find-100-improvements`.

---

## Verification & Status
- All unit tests passing (`cmdclone`, `cmdlogviewer`, `cmdinstall`).
- Binary compiled and deployed to active location: `C:\Users\Administrator\AppData\Local\gitmap-cli\gitmap.exe`.
- Git working trees clean and pushed upstream across both repositories.
