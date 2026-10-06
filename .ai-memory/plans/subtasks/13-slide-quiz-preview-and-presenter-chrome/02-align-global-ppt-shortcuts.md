# Subtask 02 — Align global deck shortcuts & HUD

**Spec:** `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md` section 5  
**Plan:** `.ai-memory/plans/pending/13-slide-quiz-preview-and-presenter-chrome.md`

## Scope

Ensure JSON slide deck exports the same `ShortcutGroup[]` as file 42 and exposes a HUD control that opens the dialog (in addition to `/`).

## Steps

1. GitMap: `gitmap cat src/components/presentation/navigation/shortcuts.ts` in global deck repo.
2. Diff groups against file 42 section 5; add missing **Deck builder** row for `E` and `S` if absent.
3. Wire Cam HUD button to camera pipeline per `31-slide-controller-buttons.md`.
4. Verify AC-SQZ-004 and AC-SQZ-005.

## Out of scope

Quiz option cards (subtask 01).
