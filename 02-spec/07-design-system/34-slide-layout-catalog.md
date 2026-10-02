# 34 — Master Slide Layout Catalog & Pure DOM Typography Specification

> **/goal** Provide the definitive, comprehensive layout catalog for 16:9 presentation slides on the 1920×1080 virtual canvas with pure DOM typography enforcement and zero baked-in text.
> **/learn** Use only the measurements in section 3. Two old catalog types have no component. Do not build those.

**Version:** 4.3.0
**Updated:** 2026-10-02
**Status:** Active
**AI Confidence:** Measured against slide components
**Ambiguity:** Two catalog types have no component. See section 3.

---

## 1. The Non-Image Text Mandate (Pure DOM Typography)

> [!CRITICAL]
> **TOTAL BAN ON BAKED-IN TEXT:**
> Headlines, subtitles, kickers, bullet points, numbered metrics, author bios, and captions MUST ALWAYS be rendered as live, selectable DOM HTML elements (`<h1>`, `<h2>`, `<p>`, `<span>`, `<div>`) styled with CSS typography tokens.
> Under no circumstances should text be flattened into raster images (`.png`, `.jpg`, `.webp`). Raster images are strictly reserved for photographic hero visual plates, author avatars, and partner logos.

---

## 2. 1920×1080 Coordinate Space & Scaling Architecture

All slide layouts operate on a virtual reference coordinate grid of `1920 × 1080` pixels:
- At runtime, `ScaledSlide` detects container dimensions using `ResizeObserver`.
- Computes uniform scale factor: $\text{scale} = \min(\text{width}/1920, \text{height}/1080)$.
- Applies vector scaling: `transform: scale(var(--stage-scale))` with `transform-origin: center center`.
- Container specifies: `contain: layout paint; isolation: isolate; will-change: transform;`.

---

## 3. Layouts measured from slide components

Canvas for every component below is `w-[1920px] h-[1080px]`. Colors that are not written here come from the active theme object. Do not copy a hex from another layout.

A type with no component does not exist. Do not build it.

| Catalog type | Source type | Component | Status |
|---|---|---|---|
| `title` | `title` | `TitleSlide` | measured |
| `executive-persona` | `persona` | `CeoPersonaSlide` | measured; source type is `persona` |
| `key-player` | `key-player` | `KeyPlayerSlide` | measured |
| `before-after` | `before-after` | `BeforeAfterSlide` | measured |
| `usp-strike` | none | none | Not specified in source. Do not invent. |
| `pricing` | `pricing` | `PricingProofSlide` | measured |
| `steps-chain` | `steps-chain` | `StepsChainSlide` | measured |
| `testimonials` | `testimonials` | `TestimonialsSlide` | measured |
| `talent-funnel` | `talent-funnel` | `TalentFunnelSlide` | measured |
| `bullets` | none | none | Not specified in source. Do not invent. |

The source also has `white-master` (`WhiteMasterSlide`), `competitive-edge` (`CompetitiveEdgeSlide`), `tech-stack` (`TechStackSlide`), and `steps` (`StepsSlide`). Those are not the missing `usp-strike` or `bullets` types.

### 3.1 `title` (`TitleSlide`)

Padding `120px`. Kicker `14px`, tracking `0.25em`, uppercase, mono. Logo height `46px`. Headline `82px` Ubuntu, weight black, leading `1.05`, max width `1400px`. Subtitle `26px` Poppins, leading `1.4`, max width `1050px`. Avatar `60×60px`. Name `22px`. Role `16px`. Date `16px` mono. Bottom wave height `140px`.

### 3.2 `persona` (`CeoPersonaSlide`)

Left column width `1100px`, padding `100px`, right padding `60px`. Kicker `13px`, tracking `0.25em`. Name `68px` Ubuntu, weight black, leading none. Role `22px`. Quote `20px` italic, leading `1.4`. Metric value `40px`. Metric label `14px` uppercase. Portrait column width `820px`. Logo `top: 60px`, `right: 80px`, height `44px`.

Not in this component: a `104px` name, a portrait at `left: -150px` with width `1500px`, or character colors `#7C3AED` / `#fdd072`.

### 3.3 `key-player` (`KeyPlayerSlide`)

Padding `100px`. Title `52px`. Subtitle `18px`, max width `1000px`. Logo height `42px`. Portrait `420×480px`, radius `24px` (`rounded-3xl`). Name `24px`. Role `15px`. Pillar title `20px`. Pillar body `16px`.

Not in this component: a row of `380×380px` portraits.

### 3.4 `before-after` (`BeforeAfterSlide`)

Padding `100px`. Title `52px`. Subtitle `18px`. Two columns, `gap: 40px` (`gap-10`). Panel title `28px`. Row text `17px`. Badge `12px` mono. The before badge uses `text-rose-300` and `bg-rose-900/60`. The after badge uses `theme.accentColor`.

Not in this component: cards of width `790px` at `left: 140px` and `left: 990px`, or a wipe handle.

### 3.5 `usp-strike`

Not specified in source. Do not invent. The comparison component is `competitive-edge`: padding `100px`, title `52px`, a 12-column table (`col-span-5`, `col-span-3`, `col-span-4`), row text `16px`, cell padding `32px` horizontal and `20px` vertical.

### 3.6 `pricing` (`PricingProofSlide`)

Padding `100px`. Title `52px`. Subtitle `18px`. Three columns, `gap: 32px` (`gap-8`), max width `1640px`. Tier name `24px`. Price `52px`. Cadence `16px`. Feature `16px`. Button `py-4` (`16px`), `text-[16px]`, `rounded-xl`. Featured badge is `11px` mono, `absolute -top-3.5`.

Not in this component: column width `500px`, `top: 270px`, height `700px`, or `scale: 1.03`.

### 3.7 `steps-chain` (`StepsChainSlide`)

Padding `100px`. Kicker `13px`. Logo height `42px`. Statement column width `480px`. Title `50px`. Subtitle `18px`. Step title `20px`. Step body `15px`. Step index `11px` mono.

Not in this component: a horizon line at `top: 364px`, or cards of `370×440px`.

`steps` (`StepsSlide`) is a different layout: grid `560px` plus the remainder, `gap: 56px` (`gap-14`), heading `48px` and `52px`, body `19px`.

### 3.8 `testimonials` (`TestimonialsSlide`)

Padding `100px`. Title `52px`. Two columns, `gap: 40px`, max width `1640px`. Quote `22px` italic, leading `1.5`. Name `20px`. Role `15px`. Logo height `42px` in the header.

Not in this component: cards at `top: 280px` sized `790×460px`, a `60px` avatar, or a logo bar at `top: 820px`.

### 3.9 `talent-funnel` (`TalentFunnelSlide`)

Padding `100px`. Title `52px`. Stack max width `1400px`, `gap: 16px` (`gap-4`). Stage width is `100 - index * 15` percent of that stack. Padding `24px` (`p-6`). Index box `48×48px` (`w-12 h-12`). Title `22px`. Description `15px`. Metric `24px`. Rate `13px`. Final stage text is `#ffffff`. Hover scale is `1.01`.

Not in this component: bands of `1640 / 1380 / 1120 / 860` by `110px`.

### 3.10 `bullets`

Not specified in source. Do not invent. The split statement-and-plate component is `white-master` (`WhiteMasterSlide`): logo `top: 60px`, `right: 100px`, height `46px`. Text block `top: 120px`, `left: 140px`, width `920px`. Kicker `13px`, tracking `0.25em`. Headline `68px`, leading `1.08`. Subtitle `24px`, leading `1.45`, max width `840px`. Icon `54×54px`. Point title `22px`. Point body `17px`. Plate `top: 0`, `right: 0`, `960×1080px`. Wave height `180px`.

---

## 4. Anti-Hallucination & Quality Verification Checklist

- [ ] Every used type is one of `title`, `persona`, `key-player`, `before-after`, `pricing`, `steps-chain`, `testimonials`, `talent-funnel`, `white-master`, `competitive-edge`, `tech-stack`, `steps`. Anything else does not exist.
- [ ] `usp-strike` and `bullets` are not built.
- [ ] Text is DOM, not baked into images.
- [ ] Headline utility is `font-ubuntu`. Body utility is `font-poppins`. Meta utility is `font-mono`. The family behind `font-mono` is not specified in these components.
- [ ] A number in an older copy of this file that is listed as "Not in this component" is not used.
