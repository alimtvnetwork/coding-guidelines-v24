# 28 — Slide Layouts

> **/goal** Close the set of slide layouts. A deck may use only these ids.
> **/learn** One-screen anatomy per id. Step reveals stay in `24-slide-presentation-system.md`.

**Version:** 1.0.0
**Status:** Active

---

## 0. Anti-hallucination

Do not add a layout id. Do not combine two ids on one slide. Colors and type come from `27-slide-canvas-and-themes.md`. If this file does not name a slot, the slot does not exist.

Step reveals: elements with `data-step="1"`, `data-step="2"`, and so on stay at `opacity: 0.15` and `filter: blur(2px)` until `currentStepIndex` reaches that step, then ease with `cubic-bezier(0.16, 1, 0.3, 1)`. Arrow keys advance the step before they advance the slide. That rule lives in `24-slide-presentation-system.md` and is not restated with new numbers here.

---

## 1. Closed union

```text
left | center | steps | timeline | process | quote | bullets | image
poll | qa | embed | reveal-grid | counter-stat | typewriter
compare | priority | depth-stack
```

---

## 2. Anatomy

Each layout is one `1920×1080` screen. Slots are required unless marked optional.

### `left`

Title and body on the left half. Optional media on the right half. No second title. No centered hero stack.

### `center`

One title and one supporting line, both centered. No side column. No card grid.

### `steps`

Ordered steps, one primary step visible at the current `data-step`. Remaining steps stay in the dimmed step style from section 0. No free-form bullets beside the step list.

### `timeline`

A single horizontal or vertical rail with dated or numbered nodes. Nodes share one accent. No second rail.

### `process`

Three to six labeled stages in one row or one column. Each stage is a short label plus one line. No nested cards.

### `quote`

One quotation and one attribution. No bullet list. No image required. Optional one mark using `--slide-hl` or the theme `hl`.

### `bullets`

One title and a single bullet list. Bullets may use `data-step`. No second list.

### `image`

One full-bleed or half-bleed image with optional one-line caption. No paragraph column longer than that caption.

### `poll`

One question and two to four choices. No open text field on the slide.

### `qa`

One question as the title and one answer block. Not a list of many questions.

### `embed`

One framed region for a diagram or external frame, plus a one-line caption. The frame stays inside the stage. It does not cover the HUD in `29-slide-navigation-and-builder.md`.

### `reveal-grid`

A grid of cells that reveal by `data-step`. All cells share one size. No mixed card heights.

### `counter-stat`

One to three numerals with one label each. Numerals use tabular figures. No paragraph under each numeral beyond the label.

### `typewriter`

One line revealed as typed text. A second line is optional and static. No bullet list.

### `compare`

Exactly two columns. Each column has a heading and a short list. No third column.

### `priority`

A ranked list, rank shown as `01`, `02`, `03`. One rank per row. No ties, and no unnumbered row.

### `depth-stack`

Stacked panels, one in front. Back panels are visible as edges only. Do not turn this into the marketing `scroll-stack` from `25-page-assembly.md`.

---

## 3. Forbidden on every layout

- A second H1.
- White Blue section ids (`capability-stack`, `solutions-grid`, `workflow-board`, `scroll-stack`, `capability-tabs`, `pricing`, `lead-form`).
- A mega menu.
- A layout id outside section 1.
- Copy that names a client or a private repository.
