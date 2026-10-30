# Plan: Guideline Prompt and Installer Upgrade

**Status:** Completed on 30-Oct-26. The coding-style checklist is already an AI prompt and states the one-parameter-per-line rule. `02-spec/03-error-manage/02-improvements.md` holds the 50 improvements. Installer subtask 02 was already checked. Release `6.29.0` was not cut. Later releases are v6.84.0 and v6.85.0.

## Acceptance Criteria

1. Restructure the coding style index/overview file into a "prompt" with "must follow instruction" when the AI reads it.
2. Put the rule about "method arguments in each line" in the root of the repo (or root of the coding guideline) formatted for AI.
3. Update the PS and Bash installer generation (generate-bundle-installers.mjs) to handle file removal during upgrades using a manifest in the inner JSON section (codingGuideline), ensuring it NEVER touches the root Version.
4. Ensure the installers emit a JSON summary after running so other AI/CLI tools can parse the installation results.
5. Create a file outlining 50 improvements in the 02-spec/03-error-manage folder.
6. Execute N-Step continuous loop gracefully handling all constraints (MD022, MD032, no [N], no generics).

## Subtasks

- [x] 01-format-overview-prompt.md: `02-spec/02-coding-guidelines/03-coding-style-checklist.md` is the AI prompt. No `00-overview.md` exists.
- [x] 02-upgrade-installer-scripts.md: Modify scripts/generate-bundle-installers.mjs to add deletion logic (based on manifest array), never touch root version, and output .ai-memory/install-summary.json. (✅ Done)
- [x] 03-generate-02-improvements.md: `02-spec/03-error-manage/02-improvements.md`.
- [x] 04-release-6.29.0.md: Not cut. Superseded by v6.84.0 and v6.85.0.
