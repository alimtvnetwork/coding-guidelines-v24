# 32 — Slide Color Options

> **/goal** Name every color choice a slide author can pick: theme ids, pill presets, and text position.
> **/learn** Theme tables live with the canvas in `27-slide-canvas-and-themes.md`. This file is the chip and alignment contract those themes did not include.

**Version:** 1.0.0
**Status:** Active

---

## 0. Anti-hallucination

A pill preset is one of the seven names below, or a raw CSS color the author typed. Do not add an eighth name. Do not recolor a chip with a marketing navy.

---

## 1. Theme ids

Pick one id per deck from `27-slide-canvas-and-themes.md`. Allowed JSON ids:

```text
snow | midnight | paper | sunset | print | playbook | noir-gold
```

`playbook` is the warm editorial theme that file 27 now lists (`bg #faf7f3`, highlight `#e8701a`). Default JSON theme is `midnight`.

A theme sets color and font only. Layout stays on the slide.

Light text on a dark ground gets `text-shadow: rgb(0 0 0) 1px 0.7px 0px`. Dark text on a light ground gets `text-shadow: none`.

---

## 2. Highlight versus pill

| Class | Look |
|---|---|
| `.hl` | Inline mark. Radius `0.1em`. Uses `--slide-hl` |
| `.hl-pill` | Chip. `inline-flex`, `line-height: 1.1`, padding `0.14em 0.4em`, weight `900`, radius `0.22em`, `box-shadow: none`, `text-shadow: none` |

Title chips (`.slide-title .hl-pill` and `.slide-title-lg .hl-pill`): padding `0.05em 0.3em 0.2em`, `line-height: 1`, radius `0.16em`, `vertical-align: middle`, `margin-top: 0.06em`.

Default chip fill is `var(--slide-hl)`. Default ink is `var(--slide-hl-ink)`.

---

## 3. Pill color presets

| Name | Background | Ink |
|---|---|---|
| `yellow` | `var(--slide-hl)` | `var(--slide-hl-ink)` |
| `white` | `#ffffff` | `#0b0b12` |
| `purple` | `#a855f7` | `#ffffff` |
| `green` | `#22c55e` | `#0b0b12` |
| `blue` | `#3b82f6` | `#ffffff` |
| `pink` | `#ec4899` | `#ffffff` |
| `red` | `#ef4444` | `#ffffff` |

Ink rule when the author passes a raw CSS color: relative luminance above `0.6` uses `#0b0b12`. Otherwise `#ffffff`. Luminance is `0.299 R + 0.587 G + 0.114 B`, each channel `0` to `255`, divided by `255`. A `var(--slide-hl)` fill keeps the theme ink and skips this rule.

---

## 4. Meaning of each preset

Use these when a pill has no `pillColor` yet. A stored `pillColor` wins. Do not overwrite it.

| Color | Use when the chip text is about |
|---|---|
| `green` | Money, growth, wins, savings, a positive outcome |
| `red` | Risk, cost, loss, warning, negation |
| `blue` | Facts, numbers, time, scope, steps, a percent |
| `purple` | People, a brand, a premium or expert claim |
| `yellow` | Everything else. This is the fallback |

One deck should not be a wall of yellow. Apply green, red, blue, and purple when the words match. Leave unmatched chips on `yellow`.

---

## 5. Nine-cell text position

`align` on a slide is one of:

```text
top-left | top-center | top-right
center-left | center | center-right
bottom-left | bottom-center | bottom-right
```

`center` is the single word `center`, not `center-center`. Changing horizontal alignment keeps the vertical axis. Changing vertical alignment keeps the horizontal axis.

---

## 6. Per-run text style

A highlight may set `style.fontSize` in authoring pixels and `style.color` as a token name or a CSS color. Omit either field to inherit. `plain: true` applies that style and does not draw `.hl` or `.hl-pill`.
