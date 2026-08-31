# Plan: Lowercase changelog.md

**Status:** PENDING

## Problem Statement
The file changelog.md exists with uppercase characters in the repository, violating the strict lowercase file naming conventions. The user requested all instances of changelog.md be renamed to changelog.md, excluding
ode_modules and .git.

## Acceptance Criteria
1. Find all files named changelog.md (case-sensitive match).
2. Rename them to changelog.md using git mv (handling Windows case-insensitivity quirks using a temp name if necessary).
3. Globally search and replace the string changelog.md with changelog.md in all repository files, avoiding external/vendor folders.
4. Commit the changes using a standard semantic commit (NO RELEASE).

## Subtasks
- 01-rename-files.md: Rename physical files.
- 02-update-references.md: Update string references globally.
