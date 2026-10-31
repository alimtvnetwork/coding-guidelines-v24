# Subtask 03: Show Log UI Favicon, Compact Header Controls, and Diff Picker

## Objective
1. Provide authentic PNG brand icon and Web App Manifest in Show Log UI so Chrome/Edge `--app` mode windows display the authentic GitMap logo instead of generic earth globe.
2. Compact header action buttons (Fullscreen, Refresh, Branch comparison, Exit) into sleek icon buttons with hover tooltips.
3. Add search filter inside branch selector combo.
4. Support Beyond Compare and VS Code diff tool configuration.

## Target Files
- `cli/cmdlogviewer/favicon.go`
- `cli/cmdlogviewer/ui_assets_html.go`
- `cli/cmdlogviewer/ui_assets_script.go`

## Steps
1. In `favicon.go`, add `/manifest.json` and PNG icon serving with valid headers.
2. In `ui_assets_html.go` and `ui_assets_script.go`:
   - Inject `<link rel="manifest" href="/manifest.json">` and PNG `<link rel="icon">`.
   - Update toolbar header with compacted icon buttons with tooltips.
   - Enhance branch selector combo with instant type-ahead filtering.
   - Wire diff triggering to check Beyond Compare or fallback to VS Code.
3. Test compilation and web endpoints.
