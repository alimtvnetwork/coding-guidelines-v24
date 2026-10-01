# 27 — Slide Canvas and Themes

> **/goal** Fix the 16:9 slide stage and the only theme tokens a deck may use.
> **/learn** Canvas scale, scripted-deck amber, JSON-deck themes, and noir-gold. Marketing pages do not use this file. They use `04-white-blue-theme/`.

**Version:** 1.0.0
**Status:** Active

---

## 0. Anti-hallucination

If a hex, HSL, or size is not in this file, do not invent it. Components read CSS variables. A hardcoded hex in a slide component is a defect, except where this file lists a token's reference hex.

Webcam picture-in-picture and dual-screen notes stay in `24-slide-presentation-system.md`.

---

## 1. Canvas

The virtual stage is `1920` by `1080`.

```text
scale = min(viewportWidth / 1920, viewportHeight / 1080)
transform-origin: center center
transform: scale(scale)
```

Layout is authored in that coordinate space. Do not use viewport units for type or gaps inside the stage.

---

## 2. Scripted deck tokens

Use these for a deck of fixed slide components (not a JSON theme picker).

| Token | HSL | Reference hex | Role |
|---|---|---|---|
| `--pres-bg` | `240 20% 4%` | `#0a0a14` | Stage |
| Stage alt | — | `#0B0B0E` | Allowed alternate ground |
| `--pres-accent` | `41 100% 50%` | `#ffae00` | Accent |
| Controller accent | `38 91% 55%` | `#F5A623` | Controller and pill chrome only |
| `--slide-hl` | — | `#FFD83A` | Inline mark on a dark JSON slide (`midnight`) |

Do not use `#FFD83A` as a large fill. It is a mark.

Headings: Ubuntu. Body: Poppins. On the `1920×1080` stage, heading size is `88px` to `140px`. Body size is `28px` to `40px`. Do not pick a size outside that range for those roles.

---

## 3. JSON theme ids

A theme sets color and font only. Layout is per slide (`28-slide-layouts.md`). Allowed ids:

| Id | `bg` | `fg` | `muted` | `hl` | `hlInk` |
|---|---|---|---|---|---|
| `snow` | `#000000` | `#ffffff` | `#b8b8b8` | `#ffffff` | `#000000` |
| `midnight` | `#101010` | `#ffffff` | `#b8b8b8` | `#ffd83a` | `#1a1100` |
| `paper` | `#f5f0e6` | `#1a1a1a` | `#615a4f` | `#1d4ed8` | `#f5f0e6` |
| `sunset` | `#1b0d1f` | `#ffeaf0` | `#c89aa6` | `#ff7a59` | `#1b0d1f` |
| `print` | `#ffffff` | `#000000` | `#444444` | `#000000` | `#ffffff` |
| `playbook` | `#faf7f3` | `#141414` | `#6b6b6b` | `#e8701a` | `#1a1a1a` |

Fonts for every row: heading and display `"Ubuntu", system-ui, sans-serif`. Body `"Poppins", system-ui, sans-serif`.

Pick one id per deck. Default is `midnight`. Do not mix `hl` from one id with `bg` from another. Pill presets and the nine-cell `align` grid are in `32-slide-color-options.md`. The dark amber shell (type scale, spotlight, controller colors) is `30-slide-palette-type-and-shell.md`.

---

## 4. Noir-gold theme

Id: `noir-gold`. Do not write these hex values in components. Map them to variables.

| Token | HSL | Reference hex | Use |
|---|---|---|---|
| `--background` | `0 0% 5%` | `#0D0D0D` | Stage |
| `--foreground` | `0 0% 100%` | `#FFFFFF` | Default text |
| `--gold` | `45 56% 54%` | `#C9A84C` | Accent, eyebrows, connectors |
| `--gold-glow` | `45 73% 67%` | `#E8C77E` | Hover and glow |
| `--cream` | `42 79% 75%` | `#F0D78C` | Title when `titleStyle` is `cream` |
| `--ember` | `13 79% 56%` | `#E85D3A` | Secondary accent, at most one per slide |
| `--ink` | `0 0% 8%` | `#141414` | Text on cream or gold fills |
| `--border` | `0 0% 18%` | `#2E2E2E` | Hairline |
| `--muted-foreground` | `0 0% 65%` | `#A6A6A6` | Secondary text |

`--primary` and `--ring` alias `--gold`.

`--gradient-noir` is `linear-gradient(180deg, #0D0D0D, #1A1A1A)`.

Bright-gold overrides, same id family, only these three:

| Token | `noir-gold` | Bright gold |
|---|---|---|
| `--gold` | `#C9A84C` | `#F3A502` |
| `--gold-glow` | `#E8C77E` | `#FFC547` |
| `--cream` | `#F0D78C` | `#FFF1D6` |

Background, ember, foreground, and type stay shared. Titles use Ubuntu Bold. Body on this theme may use Inter only when the deck id is `noir-gold` or bright gold. Every other theme in this file uses Poppins for body. Do not use Inter on a White Blue marketing page.

---

## 5. Separation from marketing UI

White Blue navy `#0D2975`, cobalt `#2563EB`, and violet `#822EE8` are page tokens. They are not slide accents. A slide must not import `04-white-blue-theme` color ramps to fill the stage.
 
---
 
## 6. Sibling References
 
- Standalone marketing images, social cards, and thumbnail specs: [`37-image-specifications.md`](./37-image-specifications.md)
- Slide layouts and constraints: [`28-slide-layouts.md`](./28-slide-layouts.md)
- 10 Master slide layouts: [`34-slide-layout-catalog.md`](./34-slide-layout-catalog.md)
