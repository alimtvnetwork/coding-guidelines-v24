# 31 — Slide Controller Buttons and Dots

> **/goal** Specify the controller pill, its buttons, the progress bar, and the dot row.
> **/learn** These controls sit on the viewport, never inside the scaled `1920×1080` stage. Keys and the slide builder stay in `29-slide-navigation-and-builder.md`.

**Version:** 1.0.0
**Status:** Active

---

## 0. Anti-hallucination

Colors are the tokens in `30-slide-palette-type-and-shell.md`. Do not recolor a button with a marketing blue.

---

## 1. Pill

| Property | Value |
|---|---|
| Position | `fixed`, `top: 32px`, `right: 32px` |
| Z-index | `50` |
| Height | `56px` |
| Padding | `8px 12px` |
| Radius | `9999px` |
| Background | `hsl(240 8% 8% / 0.85)` plus `backdrop-filter: blur(12px)` |
| Border | `1px solid hsl(var(--border))` |
| Shadow | `0 8px 24px hsl(0 0% 0% / 0.4)` |
| Gap | `4px` between groups |

Three groups, separated by a `1px` by `24px` divider in `hsl(var(--border))`:

1. Previous, counter, next
2. Share
3. Fullscreen

Hover background on every icon button: `hsl(0 0% 100% / 0.08)`, radius `9999px`, `background-color 120ms ease-out`. Active: `hsl(0 0% 100% / 0.14)`. Reduced motion snaps the background with no duration.

---

## 2. Previous and next

| Property | Value |
|---|---|
| Icons | Lucide `ChevronLeft`, `ChevronRight` |
| Icon | `20px`, stroke `2px`, `hsl(var(--foreground))` |
| Hit target | `40×40px`, round |
| Disabled | `opacity: 0.35`, `cursor: not-allowed` |
| Names | "Previous slide", "Next slide" |

Previous is disabled when the index is the first slide. Next is disabled on the last slide. Both write the hash `#slide-{n}`.

---

## 3. Counter

Poppins `500`, `20px`, `hsl(var(--foreground))`. Format is `{current} / {total}` with spaces around the slash. Min width `64px`, centered, `font-variant-numeric: tabular-nums`, `user-select: none`.

---

## 4. Share button

Lucide `Share2`, `20px`, hit target `40×40px`, name "Share current slide".

The link is the current slide, not the deck root. Set `url.hash` to `#slide-{current}` before sharing.

When `navigator.share` exists, share that URL with the deck title and the slide title. Otherwise copy the URL and show a toast "Link to slide {n} copied" that dismisses after `2s`. On mount, read `location.hash` and open that slide.

---

## 5. Fullscreen button

Lucide `Maximize2` when windowed, `Minimize2` when fullscreen. Hit target `40×40px`. Names: "Enter fullscreen" and "Exit fullscreen". Toggle `document.documentElement.requestFullscreen()` and `document.exitFullscreen()`. Swap the icon on `fullscreenchange`. Key `F` toggles. `Escape` exits.

---

## 6. Progress bar

Independent of the pill. `position: fixed; top: 0; left: 0; right: 0; height: 4px`. Track is `hsl(var(--border-subtle))`. Fill width is `current / total * 100%` with the gradient in file 30. Show it on every slide.

---

## 7. Dot row

`fixed`, `bottom: 32px`, centered. Flex row, `gap: 12px`, `z-index: 50`. Max width `min(1600px, calc(100vw - 96px))`.

| State | Shape | Fill |
|---|---|---|
| Inactive | `8×8px` circle | `--foreground-subtle` at `60%` opacity |
| Hover | `8×8px` circle | `--foreground` |
| Active | pill `28×8px`, radius `9999px` | `--primary` |
| Visited, optional | `8×8px` circle | `--foreground-subtle` at `80%` opacity |

Transition: `all 200ms cubic-bezier(0.2, 0.8, 0.2, 1)`. Reduced motion snaps.

Wrap the row in `<nav aria-label="Slide pagination">`. Each dot is a button named "Go to slide {n}: {title}" with `aria-current="true"` on the active dot.

The earlier `32×8` active mark is not this controller. New dark-amber decks use `28×8` active and `8×8` inactive.

---

## 8. Dot tooltip

Appears above the dot. Show delay `120ms`. Hide delay `80ms`. Also on keyboard focus.

| Property | Value |
|---|---|
| Background | `--background-elevated` |
| Border | `1px solid hsl(var(--border))` |
| Radius | `8px` |
| Padding | `8px 14px` |
| Shadow | `0 4px 12px hsl(0 0% 0% / 0.5)` |
| Type | Poppins `500`, `16px`, `--foreground` |
| Number | The `{n}.` prefix is `--primary`. One space, then the title in white |
| Caret | `6px` triangle, same fill, centered on the dot |
| Gap | `8px` between tooltip and dot |

Flip horizontally when the dot is near a viewport edge.

---

## 9. Transitions for this shell

The shell picks the transition. A slide does not animate its own mount.

| Variant | When | Duration | Easing |
|---|---|---|---|
| `crossfade` | Default | `240ms` | `cubic-bezier(0.2, 0.8, 0.2, 1)` |
| `slide-x` | Only when the deck opts in | `320ms` | `cubic-bezier(0.2, 0.8, 0.2, 1)` |
| `none` | `prefers-reduced-motion: reduce` | `0ms` | none |

`crossfade` is opacity `0` to `1`. `slide-x` is opacity `0` and `translateX(48px)` to opacity `1` and `translateX(0)`.

JSON decks still use `fade` and `camera-zoom` from file 29. Do not mix that family with `crossfade` / `slide-x` inside one deck.
