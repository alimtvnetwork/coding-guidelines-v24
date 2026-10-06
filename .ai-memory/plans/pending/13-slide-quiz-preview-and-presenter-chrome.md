# Plan 13 — Slide Quiz Preview & Presenter Chrome

**Status:** In progress (spec authored)  
**Canonical spec:** `02-spec/21-app/13-slide-quiz-preview-and-presenter-chrome/`  
**Design tokens:** `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md`

---

## Goal

Unify exam quiz preview chrome (option cards, shadows, botanical-light green) with global slide presenter behavior (HUD, camera, shortcuts, builder, center stage) so blind agents copy one CSS contract.

---

## Phase 1 — Spec authoring

- [x] Create `02-spec/21-app/13-slide-quiz-preview-and-presenter-chrome/` (overview, components, HUD, AC).
- [x] Add `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md`.
- [x] Register folder in `02-spec/21-app/readme.md` and file 42 in `02-spec/07-design-system/readme.md`.
- [x] Cross-link `29-slide-navigation-and-builder.md` and `21-css3-animations-and-interactions.md`.
- [x] Update `98-confidence-report.md` row for quiz/slide chrome.
- [x] v2 hardening: fault register F-01–F-13, file 42 §5.3 TS export, §5.4 key dispatch, §1.4 aliases, `40` §7 embed skins, `31` §3.2.1 HUD button, contrast script, AC-SQZ-009–012, primary **142 70% 30%**.

---

## Phase 2 — Implementation (subtasks)

| Subtask | Path | Owner |
|:---|:---|:---|
| Port CSS tokens to exam repo | `.ai-memory/plans/subtasks/13-slide-quiz-preview-and-presenter-chrome/01-port-wp-exam-theme-css.md` | App team |
| Align global-ppt shortcuts + HUD | `.ai-memory/plans/subtasks/13-slide-quiz-preview-and-presenter-chrome/02-align-global-ppt-shortcuts.md` | Presentations team |
| Register `botanical-light` in theme switch JSON | `.ai-memory/plans/subtasks/13-slide-quiz-preview-and-presenter-chrome/03-theme-switch-registration.md` | Design system |

---

## Verification

- Gherkin: `02-spec/21-app/13-slide-quiz-preview-and-presenter-chrome/04-acceptance-criteria.md`
- Manual: hover one option card and one image plate; confirm paired shadows change together.

---

## GitMap commands (discovery only)

```bash
gitmap find presentation-option-card
gitmap find SHORTCUTS
gitmap find FocusQuizRunner
```
