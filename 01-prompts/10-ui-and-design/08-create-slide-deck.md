# Create a slide deck from the design system

> **Prompt Version:** 1.0.0

Before you create or edit a presentation, read only these files. Paths are from the repository root. Do not open an external presentation repository.

1. `02-spec/07-design-system/24-slide-presentation-system.md`
2. `02-spec/07-design-system/27-slide-canvas-and-themes.md`
3. `02-spec/07-design-system/28-slide-layouts.md`
4. `02-spec/07-design-system/29-slide-navigation-and-builder.md`

## Rules

- Stage is `1920×1080`. Scale is `min(viewportWidth / 1920, viewportHeight / 1080)` with `transform-origin: center center`.
- Pick one theme id from file 27. Pick one transition family from file 29. Scripted decks cycle `Slide`, `Fade`, `Zoom`, `Flip`, `Rise`. JSON decks use `fade` or `camera-zoom`.
- Every slide layout id is in the closed union in file 28. Reject any other id.
- Step reveals follow file 24: `data-step` stays dimmed at `opacity: 0.15` and `blur(2px)` until `currentStepIndex` reaches that step.
- HUD for a new deck: `top: 32px`, `right: 32px`, height `56px`, radius `9999px`. Active progress mark is `32px` by `8px`.
- Builder key is `E`. Reorder from the pagination pill. Options panel sets layout id and theme id only. That builder is file 29. It is not `02-spec/07-design-system/26-visual-builder.md`.

## Refusal

Do not import White Blue section ids onto a slide. Do not invent a transition duration that file 29 does not state. Do not name a client, a vendor, or a private repository.
