# 28 — Slide Layouts

> **/goal** Close the set of slide layouts and name the slots each id actually has.
> **/learn** One layout per slide. Colors come from `27-slide-canvas-and-themes.md` and `32-slide-color-options.md`. Step dimming stays in `24-slide-presentation-system.md`.

**Version:** 1.1.0
**Status:** Active

---

## 0. Shared fields

Every layout extends the same base. Do not invent a field.

| Field | Rule |
|---|---|
| `id` | Stable string |
| `type` | One id from section 1 |
| `title` | Required string |
| `align` | One of the 9 cells in `32-slide-color-options.md`. Omit uses the type default |
| `padding` | Authoring pixels. Default `120` |
| `themeId` | Overrides the deck theme for this slide only |
| `background` | CSS color or image URL. Omit uses the theme |
| `gradient` | `linear` (default angle `135`) or `radial`, with 2 to 4 stops. Overrides `background` when set |
| `transition` | Overrides the deck transition for this slide only |
| `icons` | Floating marks. Center `x`/`y` on the `1920×1080` stage. Default size `120`, default opacity `0.14`. Behaviors: `float`, `drift`, `orbit`, `sway`, `pulse` |
| `images` | Extra placed images. Distinct from an `image` layout |
| `enabled` | `false` skips the slide in navigation and the dot count. Default is on |

`RichText` is a list of strings and highlight chips. A chip is `{ text, pill?, plain?, pillColor?, style? }`. `plain` keeps size and color and does not draw a chip. Pill colors are in `32-slide-color-options.md`.

Step reveals: `opacity: 0.15` and `filter: blur(2px)` until `currentStepIndex` reaches that step, then `cubic-bezier(0.16, 1, 0.3, 1)`. That rule stays in file 24.

---

## 1. Closed union

```text
left | center | steps | timeline | process | quote | bullets | image
poll | qa | embed | reveal-grid | counter-stat | typewriter
compare | priority | depth-stack
```

---

## 2. Slots

### `left`

`heading` (RichText, required), optional `kicker`, optional `body`, optional `media` (`src` + `alt`). Title and body stay on the left. Media, when present, is the right half.

### `center`

`heading` (required), optional `subhead`, optional `display`. Both lines are centered. No side column. No card grid.

### `steps`

`heading` (string) and `steps[]`. Each step has `label`, `detail` (RichText), optional `title`, optional `media` (`src`, `alt`, `caption`, `fit` of `cover` or `contain`). One step is primary. The rest stay in the dimmed step style.

### `timeline`

Optional `heading`. `items[]` with `label`, optional `title`, optional `detail`. One rail. Nodes share one accent.

### `process`

Optional `heading` and `subhead`. `stages[]`. Each stage has `title`, optional `label` (auto-numbered when omitted), optional `bullets` (1 to 3 RichText lines), optional `icon`, optional `color`. Stages sit in one row or one column of circles. No nested cards.

### `quote`

`quote` (RichText) and optional `attribution`. No bullet list.

### `bullets`

`heading` (RichText), optional `kicker`, `bullets` (RichText list). One list. Bullets may step.

### `image`

`src` required. Optional `alt`, `caption`, `heading`. `fit` is `cover`, `contain`, or `split`. `split` places text beside the image. `cover` and `contain` do not add a paragraph column beyond the caption.

### `poll`

`question` (string) and `options` (string list). No open text field.

### `qa`

Optional `prompt`. One question surface. Not a list of many questions.

### `embed`

`url` required. Optional `heading`, `caption`. `allow` defaults to `fullscreen`. The frame stays inside the stage and does not cover the HUD in file 29.

### `reveal-grid`

Optional `heading`. `items` is 2 to 6 cells. Each cell has `title`, optional `detail`, optional `icon`. One cell per step. Cells share one size.

### `counter-stat`

Optional `heading`. `stats` is 1 to 4 figures. Each figure has `value` (number), `label` (RichText), optional `prefix`, optional `suffix`. `durationMs` defaults to `1200`. Numerals use tabular figures.

### `typewriter`

Optional `heading`. `lines` is 1 to 6 strings, one line per step. `richLines`, when set, replaces `lines` for render and step count and may contain pill chips. Optional `lineStyles`. `cps` defaults to `28`. `typingSound` defaults to off.

### `compare`

Optional `heading`. `before` and `after`, each `{ src, alt?, label? }`. This is a before/after image wipe, not two text columns. `durationMs` defaults to `1000`.

### `priority`

`quote` (RichText, contains one pill chip) and optional `attribution`. Not a numbered rank list.

### `depth-stack`

Optional `heading`. `sentences` is 1 to 6 RichText lines, one per step. The newest sentence comes to the front. Older sentences tilt back, shrink, and fade. `perspective` defaults to `1200`. Optional `sentenceStyles`. This is not the marketing `scroll-stack`.

---

## 3. Forbidden on every layout

- A second title role besides the slots above.
- White Blue section ids.
- A mega menu.
- A layout id outside section 1.
- Copy that names a client or a private repository.
