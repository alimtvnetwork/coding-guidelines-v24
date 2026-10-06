# 13 — Slide Quiz Preview & Presenter Chrome

> **/goal** Capture the measured quiz preview UX (option cards, shadows, green light theme, checkboxes, center stage) and bind it to the global slide presenter (HUD, camera, shortcuts, builder, transitions).
> **/learn** Canonical tokens and copy-paste CSS live in `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md`. Slide keys and builder toggles stay in `29-slide-navigation-and-builder.md` and `35-slide-builder-canvas-inspector.md`.

**Version:** 1.0.0
**Status:** Active
**Plan:** `.ai-memory/plans/pending/13-slide-quiz-preview-and-presenter-chrome.md`

---

## Contents

| File | Role |
|:---|:---|
| [`01-overview-and-architecture.md`](01-overview-and-architecture.md) | Scope, source repos (GitMap paths), cross-file authority |
| [`02-interactive-components-and-tokens.md`](02-interactive-components-and-tokens.md) | Option cards, shadows, botanical-light palette, checkboxes, motion |
| [`03-presenter-hud-shortcuts-and-layout.md`](03-presenter-hud-shortcuts-and-layout.md) | Camera button, shortcut overlay, center content, responsiveness |
| [`04-acceptance-criteria.md`](04-acceptance-criteria.md) | Testable gates for implementers and blind agents |
| [`05-blind-agent-fault-register.md`](05-blind-agent-fault-register.md) | Fault IDs F-01–F-13 and remediation map to file 42 v2 |

---

## Authority order

When numbers disagree, later wins:

1. `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md`
2. `02-spec/07-design-system/40-theme-switch.md` (deck themes)
3. `02-spec/07-design-system/29-slide-navigation-and-builder.md` (deck keys subset)
4. `02-spec/07-design-system/31-slide-controller-buttons.md` (HUD geometry)
5. `02-spec/07-design-system/35-slide-builder-canvas-inspector.md` (builder mode)
