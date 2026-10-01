# 33 — Mega Menu Components

> **/goal** Name the menu components and the motion values an implementer must reproduce.
> **/learn** Color tokens stay in `04-white-blue-theme/01-colors-typography-and-tokens.md`. Safe-region close behavior stays in `04-white-blue-theme/02-header-mega-menu-and-footer.md`. This file adds the component list and the entrance numbers that file does not spell out.

**Version:** 1.0.0
**Status:** Active

---

## 0. Anti-hallucination

If a duration is not here or in the header file, do not invent it. Marketing pages use White Blue. They do not use slide amber.

---

## 1. Component list

Build the header from these parts. Do not merge two of them into one control.

| Component | Job |
|---|---|
| `SiteHeader` | Sticky bar, `72px`, `top-0 z-50` |
| `SlideSwapLabel` | Per-character label swap. Nav stagger is `0.04`. The primitive default `0.018` is not the nav value |
| `NavUnderline` | `1px` rule, `scale-x-0` to `scale-x-100`, origin left, over `--dur-fast` (`240ms`) |
| `NavChevron` | `14px` (`size-3.5`). Rotates `180deg` while its panel is open |
| `HeaderCtaPair` | `WhiteBlueButton` `outline` `sm` plus `WhiteBlueButton` `primary` `sm`. Hovering the pair closes the open panel |
| `MegaPanel` | Full-width panel under the header |
| `MegaGroup` | Eyebrow plus a link list |
| `MegaLink` | Label, optional description, left accent rule, trailing arrow |
| `PromoFlipCard` | Right column. Two faces. `min-h: 220px` |
| `HeaderShineButton` | Alternate header pill when the bar is the shrinking pill header in section 5. It is not `WhiteBlueButton` |
| `MobileDrawer` | Below `lg`. Does not set `overflow: hidden` on `body` |
| `StickyCta` | Viewport bar after scroll. Specified in the header file |
| `BackToTop` | Specified in the header file |

Button variants and sizes for `WhiteBlueButton` are `04-white-blue-theme/03-buttons-motion-and-interactions.md`: `primary`, `solid`, `outline`, `glass`, `ghost`, `link`, and sizes `sm` `36px`, `md` `44px`, `lg` `52px`, `icon` `44px`.

---

## 2. Mega panel entrance

Outer shell: `absolute left-0 right-0 top-full z-40 pt-3`.

Open, when motion is allowed:

- From `opacity: 0`, `translateY(-8px)`, `scale(0.985)`
- To `opacity: 1`, `translateY(0)`, `scale(1)`
- Duration `0.26s`, ease `cubic-bezier(0.16, 1, 0.3, 1)`

Close: `opacity: 0`, `translateY(-6px)`, `scale(0.99)`.

Reduced motion: opacity only, duration `0.12s`. No translate. No scale.

Inner card: `rounded-[var(--radius-card)]` (`20px`), `border`, `bg-card`, `shadow-[var(--shadow-lift)]`. Grid is `gap-8 p-8` with the column templates in the header file.

Group entrance: from `opacity: 0`, `translateY(8px)`. Duration `0.28s`. Delay `0.05s + groupIndex * 0.05s`.

Link entrance: from `opacity: 0`, `translateX(-6px)`. Duration `0.26s`. Delay `0.08s + groupIndex * 0.05s + linkIndex * 0.03s`.

Promo entrance: from `opacity: 0`, `translateY(10px)`. Duration `0.3s`. Delay `0.14s`.

---

## 3. Mega link

`rounded-[10px]`, `px-3 py-2`. Hover wash is `color-mix(in oklab, var(--primary) 7%, transparent)`.

Left rule: `width: 1px`, `origin-top`, `scaleY(0)` to `scaleY(1)`, `background-image: var(--gradient-accent)`, duration `--dur-base` (`420ms`), ease `--ease-out`.

Label: `font-display`, `text-sm`, `font-medium`, `SlideSwapLabel`, gap `6px` (`gap-1.5`) before the arrow.

Arrow: `14px` (`size-3.5`). Resting `-translate-x-1 opacity-0`. Hover `translate-x-0 opacity-100`.

Description, when present: `text-xs`, `leading-relaxed`, `text-muted-foreground`, `margin-top: 2px`.

---

## 4. Promo flip card

`perspective: 1400px`. Inner face rotates `rotateY(180deg)` over `820ms` with `--ease-out` and `transform-style: preserve-3d`. Both faces use `backface-visibility: hidden`. The back face starts at `rotateY(180deg)`.

Front: `--gradient-accent`, padding `24px` (`p-6`), white title `text-base font-bold`, body `text-sm` at `85%` white. Footer cue is uppercase, `text-xs`, `font-semibold`, tracking `0.14em`, white at `80%`.

Back: `bg-card`, `1px` border, same padding. Title `text-base font-bold`. Body `text-sm text-muted-foreground`.

Back CTA: pill, `rounded-full`, `px-4 py-2.5`, `text-sm font-semibold`, white on `--gradient-accent`, `shadow-[var(--shadow-lift)]`. Hover `scale(1.02)` over `--dur-fast`. The arrow (`16px`) moves `translateX(4px)` over `--dur-base`.

---

## 5. Header shine button

Use this pill on the shrinking header: a transparent bar `1240px` wide that becomes a white pill `900px` wide after scroll. It is a link, not a form submit. Label is caller-supplied, `whitespace-nowrap`, weight `500`.

| | Below `768px` | `768px` and up |
|---|---|---|
| Radius | `12px` | `16px` |
| Padding | `12px` horizontal, `8px` vertical | `16px` horizontal, `12px` vertical |
| Gap | `4px` | `8px` |
| Type | `16px` | `20px` |
| Icon | `16px` | `20px` |

`dark` variant: fill `--night` which is `oklch(0.19 0 0)`, shadow `0 15px 30px -5px var(--night)`, border `white` at `10%`, label white. Shine and sparks are on.

`inverse` variant: fill white, label `--ink`, shadow `0 15px 35px -18px oklch(0 0 0 / 0.6)`. No shine. No sparks.

Shine, dark only: `linear-gradient(45deg, transparent 25%, color-mix(in oklch, var(--brand) 50%, transparent) 50%, transparent 75%)`. Background size `250% 250%`. `@keyframes` moves `background-position` from `200% 0` to `-200% 0`. Duration `3s linear infinite`.

Sparks, dark and `md` and up only: 8 dots, `bg-brand`, `pointer-events-none`. Float is `translateY(5px)` to `translateY(-5px)`, `2.4s ease-in-out infinite alternate`. The dot's parent holds the seeded X and Y so the float does not wipe them.

| # | Size | X | Y | Opacity | Scale | Delay |
|---|---|---|---|---|---|---|
| 1 | 6 | 13 | -13 | 0.80 | 1.00 | 0s |
| 2 | 5 | 16 | -8 | 0.71 | 0.89 | 0.4s |
| 3 | 8 | 9 | -11 | 0.52 | 0.65 | 0.8s |
| 4 | 5 | 18 | -11 | 0.62 | 0.78 | 1.2s |
| 5 | 4 | 3 | -5 | 0.27 | 0.34 | 1.6s |
| 6 | 6 | 9 | -9 | 0.40 | 0.50 | 2.0s |
| 7 | 4 | 17 | -7 | 0.77 | 0.96 | 2.4s |
| 8 | 5 | 4 | -2 | 0.12 | 0.15 | 2.8s |

No hover lift on this button. The shine is the motion. Reduced motion sets `animation: none` on the shine and the sparks. The dots stay at their seeded place.

Focus: `2px` ring in `--brand`, `2px` offset, `focus-visible` only.

---

## 6. Page motion primitives

Compose these. Do not author a new curve.

| Token | Value |
|---|---|
| `--dur-instant` | `120ms` |
| `--dur-fast` | `240ms` |
| `--dur-base` | `420ms` |
| `--dur-slow` | `700ms` |
| `--dur-cine` | `1100ms` |
| `--ease-out` | `cubic-bezier(0.16, 1, 0.3, 1)` |
| `--ease-in-out` | `cubic-bezier(0.65, 0, 0.35, 1)` |

| Primitive | Motion |
|---|---|
| `Reveal` | Fade and rise once. `y` from `28px` to `0`. Compact mobile variant uses `16px` |
| `StaggerGroup` | Children rise in sequence. Default stagger `0.08s` |
| `MaskedHeading` | Per word or per line, `y` from `110%` to `0`. One accent phrase |
| `CountUp` | Number counts once in view |
| `TiltCard` | Pointer tilt, max `8deg`. One per page |
| `SpotlightCard` | Cursor radial highlight |
| `Magnetic` | Pull toward the pointer. Primary CTA only. Strength `0.22` on `WhiteBlueButton` |
| `DragRail` | Horizontal drag, snap, progress |

Budget: at most 3 elements animating at once, 1 pinned section, 1 marquee or drag rail. Hero entrance at most `1.4s`. The H1 is legible within `800ms`. Animate `transform` and `opacity`. Accordion height is the exception.
