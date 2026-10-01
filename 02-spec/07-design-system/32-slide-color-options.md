# 32 — Slide Color Options & 10-Step Gradient Precision System

> **/goal** Provide the definitive mathematical 10-step gradient and shade ramp system ($S_0$ to $S_9$), 7 semantic pill presets, relative luminance formula, character-by-character shading engine, and 9-cell text alignment coordinates.
> **/learn** Master the exact HSL, RGB, and HEX coordinates across the 4 flagship themes (Royal Violet, Corporate Gold, Enterprise Blue, Clinical Emerald), calculate text contrast via relative luminance ($0.299R + 0.587G + 0.114B > 0.6 \implies$ dark ink `#0b0b12`, else `#ffffff`), and bind 9-cell positioning.

**Version:** 4.0.0
**Status:** Active
**AI Confidence:** High
**Ambiguity:** None

---

## 1. Executive System Overview

To eliminate rendering hallucinations and ensure that AI models and human engineers reproduce identical visual authority, color transitions across presentation decks must never rely on ambiguous interpolation. This specification codifies:
1. **Mathematical 10-Step Gradient System:** Discrete steps ($S_0$ through $S_9$) in HSL, RGB, and HEX across 4 flagship palettes.
2. **Character-by-Character Shading Engine:** Per-glyph color progression on hero names and title keywords.
3. **7 Semantic Pill Presets:** Deterministic badge and chip fills with mathematical ink calculation.
4. **9-Cell Text Position System:** Standardized 2D stage alignment coordinates.

---

## 2. Mathematical 10-Step Ramp Formula

For any two anchor colors $C_{\text{start}} (S_0)$ and $C_{\text{end}} (S_9)$, intermediate steps $S_i$ ($i \in [0, 9]$) are calculated via linear perceptually uniform lightness and hue interpolation:

$$t_i = \frac{i}{9}, \quad i \in \{0, 1, 2, 3, 4, 5, 6, 7, 8, 9\}$$
$$\text{Hue}_i = \text{Hue}_{\text{start}} + t_i \cdot (\text{Hue}_{\text{end}} - \text{Hue}_{\text{start}})$$
$$\text{Sat}_i = \text{Sat}_{\text{start}} + t_i \cdot (\text{Sat}_{\text{end}} - \text{Sat}_{\text{start}})$$
$$\text{Light}_i = \text{Light}_{\text{start}} + t_i \cdot (\text{Light}_{\text{end}} - \text{Light}_{\text{start}})$$

---

## 3. Flagship 10-Step Theme Gradient Tables

### 3.1 Theme 1: Royal Violet (Sample White Presentation System)
Transitions from Deep Midnight Indigo ($S_0$) through Vibrant Electric Violet ($S_5$) to Soft Glowing Lilac ($S_9$):

| Step | Level / Role | HSL Coordinate | RGB Coordinate | HEX Value | Primary Architectural Application |
|:---:|:---|:---|:---|:---:|:---|
| **$S_0$** | Anchor Dark | `hsl(243, 75%, 25%)` | `rgb(16, 16, 112)` | `#101070` | Wave base layer, bottom contour |
| **$S_1$** | Deep Shadow | `hsl(255, 70%, 32%)` | `rgb(49, 25, 138)` | `#31198A` | Deep container borders, text shadow |
| **$S_2$** | Base Violet | `hsl(265, 68%, 38%)` | `rgb(77, 31, 163)` | `#4D1FA3` | Secondary icons, bullet line tails |
| **$S_3$** | Core Brand Deep | `hsl(270, 72%, 44%)` | `rgb(107, 33, 168)`| `#6B21A8` | Circular bullet badges, button bg |
| **$S_4$** | Primary Accent | `hsl(262, 83%, 58%)` | `rgb(124, 58, 237)`| `#7C3AED` | Pill marker, headline accent words |
| **$S_5$** | Electric Bright | `hsl(260, 88%, 66%)` | `rgb(147, 91, 245)`| `#935BF5` | Neon heart inner stroke, hover glow |
| **$S_6$** | Soft Lavender | `hsl(265, 90%, 75%)` | `rgb(180, 134, 248)`| `#B486F8` | Vertical dividing line, pill border |
| **$S_7$** | Light Lilac | `hsl(268, 92%, 84%)` | `rgb(216, 185, 252)`| `#D8B9FC` | Ambient watermark arcs, glow aura |
| **$S_8$** | Ultra-Light Tint| `hsl(270, 94%, 92%)` | `rgb(238, 222, 254)`| `#EEDEFE` | Card hover fill, soft table zebra |
| **$S_9$** | Atmospheric Fog| `hsl(270, 95%, 97%)` | `rgb(250, 245, 255)`| `#FAF5FF` | Subtle canvas background wash |

### 3.2 Theme 2: Midnight Corporate Gold (Executive Persona & Keynote Dark)
Transitions from Deep Obsidian Slate ($S_0$) to Warm Radiant Gold ($S_9$):

| Step | Level / Role | HSL Coordinate | RGB Coordinate | HEX Value | Primary Architectural Application |
|:---:|:---|:---|:---|:---:|:---|
| **$S_0$** | Pure Obsidian | `hsl(222, 47%, 11%)` | `rgb(15, 23, 42)` | `#0F172A` | Primary slide dark background |
| **$S_1$** | Midnight Navy | `hsl(217, 33%, 17%)` | `rgb(29, 41, 59)` | `#1D293B` | Glass card surfaces, panel fills |
| **$S_2$** | Charcoal Border| `hsl(215, 25%, 27%)` | `rgb(51, 65, 85)` | `#334155` | Card dividing rules, inactive tabs |
| **$S_3$** | Muted Slate Ink| `hsl(215, 16%, 47%)` | `rgb(100, 116, 139)`| `#64748B` | Subtitle text, presenter captions |
| **$S_4$** | Intermediate Tint| `hsl(38, 45%, 52%)` | `rgb(187, 142, 78)` | `#BB8E4E` | Intermediate glyph gradient step |
| **$S_5$** | Muted Gold | `hsl(42, 75%, 58%)` | `rgb(228, 172, 68)` | `#E4AC44` | Metric highlights, timeline pins |
| **$S_6$** | Warm Amber | `hsl(45, 88%, 65%)` | `rgb(243, 194, 88)` | `#F3C258` | Character gradient accent step |
| **$S_7$** | Radiant Amber | `hsl(40, 96%, 72%)` | `rgb(253, 208, 114)`| `#FDD072` | Main character accent highlight |
| **$S_8$** | Champagne Gold | `hsl(43, 94%, 82%)` | `rgb(254, 231, 164)`| `#FEE7A4` | Hover aura, radial backdrop glow |
| **$S_9$** | White Gold Tip | `hsl(45, 92%, 94%)` | `rgb(255, 250, 224)`| `#FFFAE0` | Highlight badge text, glint points |

### 3.3 Theme 3: Enterprise Slate Blue (High-Authority SaaS & Engineering)
Transitions from Deep Slate Ink ($S_0$) to Crisp Pure White ($S_9$):

| Step | Level / Role | HSL Coordinate | RGB Coordinate | HEX Value | Primary Architectural Application |
|:---:|:---|:---|:---|:---:|:---|
| **$S_0$** | Dark Ink Slate | `hsl(222, 84%, 5%)` | `rgb(2, 8, 23)` | `#020817` | High-contrast headline text |
| **$S_1$** | Core Navy Brand | `hsl(222, 47%, 11%)` | `rgb(15, 23, 42)` | `#0F172A` | Primary buttons, active state tabs |
| **$S_2$** | Dark Blue Gray | `hsl(217, 33%, 25%)` | `rgb(43, 61, 88)` | `#2B3D58` | Secondary button hover, dark card |
| **$S_3$** | Medium Slate | `hsl(215, 22%, 38%)` | `rgb(76, 96, 122)` | `#4C607A` | Metadata text, secondary headers |
| **$S_4$** | Muted Slate | `hsl(215, 16%, 47%)` | `rgb(100, 116, 139)`| `#64748B` | Paragraph body copy, icons |
| **$S_5$** | Sky Accent Dark | `hsl(201, 80%, 45%)` | `rgb(23, 142, 207)` | `#178ECF` | Interactive link buttons, focus ring|
| **$S_6$** | Vivid Azure Sky | `hsl(199, 89%, 48%)` | `rgb(14, 165, 233)` | `#0EA5E9` | Dynamic stats numbers, active pills|
| **$S_7$** | Soft Slate Border| `hsl(214, 32%, 91%)` | `rgb(226, 232, 240)`| `#E2E8F0` | Container outlines, vertical rules |
| **$S_8$** | Light Muted Gray | `hsl(210, 40%, 96%)` | `rgb(241, 245, 249)`| `#F1F5F9` | Table alternate rows, badge backdrops|
| **$S_9$** | Crisp Pure White | `hsl(0, 0%, 100%)` | `rgb(255, 255, 255)`| `#FFFFFF` | Slide background, button text |

### 3.4 Theme 4: Emerald Clinical Green (Healthcare & Evidence Metrics)
Transitions from Deep Pine ($S_0$) to Clinical Pure White ($S_9$):

| Step | Level / Role | HSL Coordinate | RGB Coordinate | HEX Value | Primary Architectural Application |
|:---:|:---|:---|:---|:---:|:---|
| **$S_0$** | Deep Pine Forest| `hsl(165, 80%, 12%)` | `rgb(6, 55, 41)` | `#063729` | Base dark contrast ground |
| **$S_1$** | Dark Evergreen | `hsl(162, 75%, 20%)` | `rgb(13, 89, 67)` | `#0D5943` | Clinical headline text |
| **$S_2$** | Clinical Green | `hsl(160, 72%, 28%)` | `rgb(20, 123, 93)`| `#147B5D` | Subtitles, certified headers |
| **$S_3$** | Deep Mint Green | `hsl(158, 68%, 35%)` | `rgb(29, 150, 114)`| `#1D9672` | Card borders, secondary buttons |
| **$S_4$** | Core Brand Mint| `hsl(155, 75%, 44%)` | `rgb(28, 196, 145)`| `#1CC491` | Verification checkmarks, badges |
| **$S_5$** | Electric Emerald| `hsl(150, 84%, 55%)` | `rgb(46, 235, 163)`| `#2EEBA3` | Active status badges, metric glow |
| **$S_6$** | Mint Highlight | `hsl(145, 86%, 68%)` | `rgb(103, 244, 187)`| `#67F4BB` | Character gradient accent steps |
| **$S_7$** | Soft Mint Aura | `hsl(140, 88%, 80%)` | `rgb(160, 248, 211)`| `#A0F8D3` | Background medical wave ribbons |
| **$S_8$** | Pale Clinical Tint| `hsl(135, 85%, 92%)` | `rgb(215, 252, 236)`| `#D7FCEC` | Card background wash |
| **$S_9$** | Clinical White | `hsl(130, 80%, 98%)` | `rgb(245, 254, 250)`| `#F5FEFA` | Clean clinical canvas background |

---

## 4. Character-by-Character Shading Engine

In executive persona and hero title slides, hero names utilize character-level color stepping to produce a luxurious, organic transition rather than a flat fill:

```tsx
<motion.h1 className="font-heading italic" style={{ fontSize: 104, fontWeight: 700 }}>
  {/* Leading Character: High-Contrast Brand Primary Accent */}
  <span style={{ color: "hsl(var(--pres-accent))" }}>A</span>
  {/* Intermediate Character: Warm Intermediate Step (S4 / S6) */}
  <span style={{ color: "#fdd072" }}>l</span>
  {/* Terminal Characters: Crisp Primary Ink Text */}
  <span style={{ color: "hsl(var(--pres-text))" }}>ex</span>{" "}
  {/* Second Word */}
  <span style={{ color: "hsl(var(--pres-accent))" }}>M</span>
  <span style={{ color: "#fdd072" }}>o</span>
  <span style={{ color: "hsl(var(--pres-text))" }}>rgan</span>
</motion.h1>
```

- **Leading Glyph ($C_0$):** High-contrast brand primary accent (`$S_4$` or `$S_5$`).
- **Intermediate Glyph ($C_1$):** Warm intermediate tint (`$S_6$` or `$S_7$`).
- **Terminal Glyphs ($C_2..C_n$):** Primary ink text color (`$S_0$` or `$S_1$`).

---

## 5. 7 Semantic Pill Presets & Contrast Formula

Chips and pill badges use a closed union of 7 semantic color tokens:

| Preset Name | Background Fill | Semantic Application | Default Ink | Contrast Ratio |
|:---|:---|:---|:---:|:---:|
| **`yellow`** | `var(--slide-hl, #FBBF24)` | General emphasis, default fallback | `#0b0b12` | 12.8:1 |
| **`white`** | `#FFFFFF` | Inverted chips on dark backgrounds | `#0b0b12` | 19.5:1 |
| **`purple`** | `#A855F7` | People, founders, executive claims | `#FFFFFF` | 4.8:1 |
| **`green`** | `#22C55E` | Money, revenue growth, positive outcomes | `#0b0b12` | 9.6:1 |
| **`blue`** | `#3B82F6` | Facts, metrics, durations, timeline steps | `#FFFFFF` | 4.6:1 |
| **`pink`** | `#EC4899` | Creative, marketing, consumer delight | `#FFFFFF` | 4.5:1 |
| **`red`** | `#EF4444` | Risks, pain points, losses, warnings | `#FFFFFF` | 4.7:1 |

### 5.1 Relative Luminance Ink Formula
When an author or AI passes a custom raw CSS color for `pillColor`, the text ink color MUST be deterministically calculated via ITU-R BT.709 relative luminance:

$$L = 0.299 \cdot R_{\text{norm}} + 0.587 \cdot G_{\text{norm}} + 0.114 \cdot B_{\text{norm}}$$

- If $L > 0.6$: Ink MUST be dark `#0b0b12`.
- If $L \le 0.6$: Ink MUST be pure white `#FFFFFF`.

---

## 6. 9-Cell Text Position Coordinate Matrix

The `align` field on any slide layout is bound to a strict 9-cell coordinate enum:

```
┌─────────────────┬─────────────────┬─────────────────┐
│  top-left       │  top-center     │  top-right      │
├─────────────────┼─────────────────┼─────────────────┤
│  center-left    │  center         │  center-right   │
├─────────────────┼─────────────────┼─────────────────┤
│  bottom-left    │  bottom-center  │  bottom-right   │
└─────────────────┴─────────────────┴─────────────────┘
```

> **Note:** The center cell is strictly named `center`, never `center-center`. Changing the horizontal axis preserves vertical alignment; changing the vertical axis preserves horizontal alignment.

---

## 7. Anti-Hallucination Checklist

- [ ] Every slide gradient is mapped to one of the 4 standard 10-step tables ($S_0$ through $S_9$).
- [ ] Character shading strictly applies accent to leading glyph, intermediate tint to second glyph, and primary ink to remainder.
- [ ] Pill presets only select from the 7 defined names (`yellow`, `white`, `purple`, `green`, `blue`, `pink`, `red`).
- [ ] Custom pill colors enforce relative luminance threshold ($L > 0.6 \implies$ `#0b0b12`, else `#FFFFFF`).
- [ ] Text alignment rejects any value outside the 9 defined coordinate strings.
