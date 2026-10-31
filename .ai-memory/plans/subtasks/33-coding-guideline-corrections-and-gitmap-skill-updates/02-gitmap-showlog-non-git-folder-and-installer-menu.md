# Subtask 02: GitMap Show Log Non-Git Folder Resilience & Context Menu Installer UI

## Objective
1. When `gitmap show-log` is invoked on a folder that is not a Git repository (e.g. `D:\work`), eliminate the fatal Go runtime stack trace and provide a clean diagnostic with redirection/suggestion to run `gitmap ui` or pick a detected repository.
2. Ensure folder context menu directly includes Installer Manager UI.

## Target Files
- `cli/cmdlogviewer/server.go`
- `cli/cmdinstall/installctxentries.go`

## Steps
1. In `cli/cmdlogviewer/server.go`, refine `newNonGitRepoShowLogError` and `ResolveTargetRepoAndPath`:
   - Avoid printing stack traces for directory detection.
   - If child directories have git repos, suggest `gitmap ui` or opening a specific child repository.
2. In `cli/cmdinstall/installctxentries.go`:
   - Verify `03_installer_ui` is cleanly registered in folder background menus.
3. Test locally with `go test ./cli/cmdlogviewer/...`.
