# Subtask 2: Upgrade Installers

1. Edit scripts/generate-bundle-installers.mjs.
2. Enhance the PowerShell and Bash scripts logic for version.json merge:
   - For codingGuideline, calculate installedFiles array (by listing files extracted from the bundle).
   - Before extracting, read codingGuideline.installedFiles from the existing version.json. Any file in the old list that is NOT in the new list should be explicitly deleted.
   - Output an install-summary.json file inside .lovable/ detailing versions updated and files removed.
3. Generate bundles (
pm run bundles:generate).
