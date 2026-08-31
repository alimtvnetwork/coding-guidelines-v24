# Subtask 2: Release

1. Verify
eadme.md is lowercase.
2. Update .lovable/what-to-read.md with any new spec files (if created).
3. Run
px markdownlint "**/*.md" --ignore node_modules --ignore .lovable/prompts --fix (if available) or manual check.
4. Run
ode scripts/release.mjs --tier minor --scope "Implement apperror.New creator namespace" --skip-slides.
5. Run
ode scripts/sync-check.mjs --fix.
6. Run
pm run bundles:generate.
7. Update versions in version.template.json, .lovable/memory/standards/version-source-of-truth.md, and .lovable/memory/01-index.md.
8. Commit, tag v6.27.0, and push.
