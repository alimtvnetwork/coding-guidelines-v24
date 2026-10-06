# Subtask 01 — Port theme CSS (exam product)

**Spec:** `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md`  
**Plan:** `.ai-memory/plans/pending/13-slide-quiz-preview-and-presenter-chrome.md`

## Scope

Replace ad hoc `--option-text-shadow-*` naming with semantic `--text-shadow-*` and `--elevation-*` aliases (keep legacy aliases as one-line forwards for one release).

## Steps

1. GitMap: `gitmap cat src/styles/theme.css` in exam repo.
2. Merge section 1–2 from file 42; map `[data-theme="green-choice"]` to `[data-theme="botanical-light"]` or duplicate selectors.
3. Apply botanical-light HSL from file 42 section 3; remove redundant `#16A34A` literals in TSX where `hsl(var(--primary))` suffices.
4. Verify AC-SQZ-001, AC-SQZ-002, AC-SQZ-007.

## Out of scope

Slide deck HUD (subtask 02).
