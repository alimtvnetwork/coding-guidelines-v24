# Subtask 31.5: GitMap CLI & Show Log UI Enhancements

> **Parent Plan:** `31-coding-guideline-corrections-and-gitmap-skill-updates`  
> **Status:** Completed  
> **Target:** `d:\work\gitmap`  
> **Date:** 2026-10-27  

---

## Objectives & Verified Outcomes
1. **Short Repo Clone Resolution:**
   - Updated `d:\work\gitmap\cli\cmdclone\gh_resolver.go` and `clone.go` to sort matching releases by semantic version, selecting the highest version when bare names are passed, and converting to SSH format when requested.
2. **Show Log Non-Git Root Crash Fix:**
   - Added `newNonGitRepoShowLogError()` in `d:\work\gitmap\cli\cmdlogviewer\server.go` and non-zero git exit code check in `d:\work\gitmap\cli\cmdgit\git_cmd.go`. When executed outside git repositories, prints `✗ Not a git repository: <path>` cleanly without any Go runtime stack traces.
3. **Native Process Kill Command:**
   - Implemented `gitmap proc kill <pid> [--force]` (and `gitmap kill <pid>`) in `d:\work\gitmap\cli\cmdproc\proc.go` with self/caller/system PID protection and two-stage graceful fallback.
4. **Heatmap Query Shortcuts:**
   - Implemented `failures`, `hotfiles`, `recent`, and `tasks` in `d:\work\gitmap\cli\heatmap\heatmap_shortcuts.go` executing sub-millisecond SQLite queries.
   - Added `RootSpecFilePath` and `TaskFolder` columns and records to Heatmap DB schema in `d:\work\gitmap\cli\store\heatmap_schema.go` and `heatmap_rc_schema.go`.
5. **Show Log UI Enhancements:**
   - Embedded SVG favicon in HTML `<head>`.
   - Replaced text buttons with compact SVG icon buttons with `[data-tooltip]` hover tooltips.
   - Added search filter to branch selector dropdown.
   - Added diff tool configuration modal for Beyond Compare and VS Code.
