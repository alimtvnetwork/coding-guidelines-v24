# Plan: Normalize Encoding and Line Endings

**Status:** COMPLETED

## Acceptance Criteria
1. Fix all .md (and other text) file encodings to UTF-8 without BOM.
2. Standardize line endings to use only \n (LF) everywhere.
3. Update the 02-spec/02-coding-guidelines/03-03-coding-style-checklist.md to explicitly state these encoding and line ending rules.
4. Execute release and sync processes smoothly without failing.

## Subtasks
- 01-normalize-encoding.md: Script the BOM removal and line ending normalization across the repo.
- 02-update-checklist.md: Append the encoding rules to 03-coding-style-checklist.md.
- 03-release-6.31.0.md: Release minor bump v6.31.0.
