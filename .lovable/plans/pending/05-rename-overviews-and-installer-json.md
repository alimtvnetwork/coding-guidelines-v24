# Plan: Rename Overviews and Enhance Installers

**Status:** PENDING

## Acceptance Criteria

1. **Rename Spec Overviews:** Rename all 00-overview.md files within spec/ to 01-index.md.
2. **Global Reference Update:** Update every reference from 00-overview.md to 01-index.md across the entire codebase.
3. **AI Reading Instruction:** Add a prominent note in every 01-index.md and the root
eadme.md stating that 01-index.md is the first file any AI should read in a directory.
4. **Installer JSON Flag:** Update install scripts (generate-bundle-installers.mjs) to support a --json (or -Json in PS) flag. When used, the script outputs ONLY the JSON summary to stdout.
5. **Installer Visualization & Tracking:** The installers must properly compute the exact list of files they extract/update, track them in version.json's codingGuideline section, and print a nice console summary of what changed (unless --json is active).
6. **Release Execution:** Bump minor version, update sync indices, commit and tag.

## Subtasks

- 01-rename-and-update-refs.md: Script the renaming of 00-overview.md to 01-index.md and global find/replace. Inject the AI reading instruction.
- 02-enhance-installers.md: Rewrite the post-install logic in generate-bundle-installers.mjs to compute extracted files, handle --json flag, and format console output nicely.
- 03-update-readme.md: Update
eadme.md to document the AI reading rule and the installer --json output logic.
- 04-release-6.30.0.md: Release script execution and readme sync.
