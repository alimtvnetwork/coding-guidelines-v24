# 29 — Slide Navigation and Slide Builder

> **/goal** Specify the presenter HUD, the two transition families, the keys, and the slide builder.
> **/learn** Scripted decks cycle five transitions. JSON decks use fade and camera-zoom. The slide builder is not `26-visual-builder.md`.

**Version:** 1.0.0
**Status:** Active

---

## 0. Anti-hallucination

Do not invent a key, a transition, or a chrome size. Website editing (wording, images, menu href) is `26-visual-builder.md` and must not be mounted on a slide stage.

---

## 1. HUD

Top controller pill:

- Canonical offset: `top: 32px` and `right: 32px`
- A scripted deck may already use the class `fixed top-6 right-6 z-50`. On a default `4px` scale that class is `24px`, not `32px`. New decks use `32px`. Do not invent a third offset.
- Height `56px`
- `border-radius: 9999px`
- `z-index` above the stage (`z-50` on the scripted deck)

Bottom progress: the active mark is `32px` wide and `8px` tall, filled with the deck accent (`--pres-accent` or the active theme `hl`). Inactive marks are smaller dots. Do not invent their resting size.

The HUD must stay outside the scaled stage, or it scales with the slide and becomes unreadably small. Mount it on the viewport, not inside the `1920×1080` transform.

---

## 2. Transitions

Pick one family per deck. Do not mix families on adjacent slides.

### 2.1 Scripted deck

Cycle by slide index, in this order only:

1. `Slide`
2. `Fade`
3. `Zoom`
4. `Flip`
5. `Rise`

`Zoom` enters from `scale: 0.85` and `opacity: 0`. Other members of the cycle are named here so an implementation can map them. Do not add a sixth name. Durations for `Slide`, `Fade`, `Flip`, and `Rise` are not fixed in this file. Do not invent milliseconds for them.

### 2.2 JSON deck

Allowed `TransitionKind` values:

```text
fade | camera-zoom | slide
```

The settings UI exposes `fade` and `camera-zoom` only. `slide` may exist on the type and must not appear as a third settings choice unless this file is revised. Default is `fade`.

---

## 3. Keys

| Key | Action |
|---|---|
| `ArrowRight` or `Space` | Next step, then next slide |
| `ArrowLeft` | Previous step, then previous slide |
| `G` | Slide grid |
| `S` | Settings |
| `E` | Toggle the slide builder |

Double-click a counter numeral to jump to that slide index. Do not bind `E` to the website visual builder.

---

## 4. Slide builder

`E` toggles builder chrome on the deck. The audience route does not load it.

- Reorder slides from the pagination pill.
- Edit the current slide in a `SlideOptionsPanel`: layout id from `28-slide-layouts.md`, theme id from `27-slide-canvas-and-themes.md`, and the slots that layout allows.
- Persist on the deck JSON. Do not write a second store for colors or fonts.
- Reject a layout id outside the closed union.
- Reject a theme id outside `27-slide-canvas-and-themes.md`.

The slide builder does not double-click text on the marketing page, does not upload a `2 MB` image through the website overlay, and does not edit a mega-menu href.

---

## 5. Presenter extras

Webcam overlay size, mirror, and drag behavior stay in `24-slide-presentation-system.md`. Do not copy those pixel values into this file. Dual-screen notes and the next-slide preview stay there too.
