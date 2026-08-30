# Subtask 2: Enhance Installers

1. Edit scripts/generate-bundle-installers.mjs.
2. Bash: Add --json arg parsing. If --json, silence normal echos. After extraction, run find to get actual extracted files. Update version.json, print JSON if requested, otherwise print a clean visual list of files.
3. PowerShell: Add [switch]. Do the same output suppression and file tracking using Get-ChildItem.
4. Generate bundles (
pm run bundles:generate).
