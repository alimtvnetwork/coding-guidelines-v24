# 34 — Master Slide Layout Catalog & Pure DOM Typography Specification

> **/goal** Provide the definitive, comprehensive layout catalog for 16:9 presentation slides on the 1920×1080 virtual canvas with pure DOM typography enforcement and zero baked-in text.
> **/learn** Master the coordinate geometries, slot models, typography scales, and visual zones across the 10 core enterprise slide layouts: Title Hero, Executive Persona, Key Player Bio, Before/After Split, USP Strikethrough, SaaS Pricing, Steps Chain Roadmap, Social Proof, Talent Funnel, and 3-Point Master Cards.

**Version:** 4.0.0
**Status:** Active
**AI Confidence:** High
**Ambiguity:** None

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

## 3. The 10 Master Slide Layouts

### 3.1 Layout 1: Title & Hero Slide (`type: "title"`)
Establishes topic authority, enterprise identity, and presentation context.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│ [Top Category Pill]                                 [Transparent Brand Logo]│
│                                                                             │
│   MASSIVE EDITORIAL HEADLINE (78px - 84px)                                  │
│   Secondary Line with Gradient Accent Word                                  │
│                                                                             │
│   Explanatory Subtitle / Mission Statement (26px)                           │
│                                                                             │
│ ┌───────────────────────┐                                                   │
│ │ Presenter Avatar/Bio  │                                                   │
│ │ Date & Session Code   │                                                   │
│ └───────────────────────┘                        [Bottom Organic SVG Wave]  │
└─────────────────────────────────────────────────────────────────────────────┘
```

- **Top Kicker Pill:** Left `140px`, Top `120px`, Font size `16px` mono uppercase, tracking `0.15em`, background `rgba(124, 58, 237, 0.08)`.
- **Top-Right Logo:** Right `120px`, Top `80px`, height `48px`.
- **Main Headline:** Left `140px`, Top `260px`, Width `1200px`, Font `Ubuntu` bold, `78px`, line-height `1.1`.
- **Subtitle:** Left `140px`, Top `460px`, Width `960px`, Font `Poppins`, `26px`, color `#475569`, line-height `1.4`.
- **Presenter Bio Card:** Left `140px`, Top `640px`, Height `96px`, Avatar circle `64×64px`, Name `20px` bold, Title `15px` muted.
- **Bottom Organic Wave:** Bottom `0px`, Left `0px`, Width `1920px`, Height `180px` dual-gradient SVG ribbon.

---

### 3.2 Layout 2: Executive Persona & CEO Slide (`type: "executive-persona"`)
Balances commanding portrait photography with verifiable credentials and impact metrics.

```
┌───────────────────────────────────────┬───────────────────────────────────────────┐
│ 1. Portrait Staging Zone              │ 2. Executive Credentialing Zone           │
│    - Left: -150px, Width: 1500px      │    - Left: 800px, Right: 72px             │
│    - Height: 100% (Anchored to base)  │    - Top: 132px, Bottom: 118px            │
│    - Halftone dot matrix (10px grid)  │    - Kicker: "Executive Leadership"       │
│    - Radial accent blur aura          │    - Hero Name: 104px Ubuntu Bold Italic  │
│    - Drop shadow:                     │    - Character-by-Character Shading       │
│      drop-shadow(0 32px 64px ...)     │    - Role & Specialization (38px Bold)    │
│                                       │    - Interactive LinkedIn & Credential Tag│
│                                       │    - 2-Column Experience & Impact Grid    │
└───────────────────────────────────────┴───────────────────────────────────────────┘
```

- **Hero Name Character Stepping:**
  - Leading character: Brand primary accent (`$S_4$` / `#7C3AED`).
  - Intermediate character: Warm intermediate step (`$S_6$` / `#fdd072`).
  - Terminal characters: Primary ink text (`$S_0$` / `#0F172A`).
- **Interactive Badges:** LinkedIn preview button with floating screenshot hover card, location tag with map pin, specialization tags.

---

### 3.3 Layout 3: Key Player Bio Slide (`type: "key-player"`)
Structured grid introducing core architects, researchers, or executive directors.

- **Header Block:** Left `140px`, Top `120px`, Title `54px`, Subtitle `22px`.
- **Grid Layout:** 3–4 Member columns (`width: 380px` each, gap `40px`, Top `280px`).
- **Member Card:**
  - Photographic portrait: `380×380px`, rounded `20px`, border `1px solid var(--border)`.
  - Name: `24px` Ubuntu bold.
  - Role: `16px` brand secondary.
  - Bio: `14px` Poppins muted, line-height `1.5`.
  - Social Links: GitHub, LinkedIn, Website vector icons.

---

### 3.4 Layout 4: Before / After Showcase Slide (`type: "before-after"`)
Demonstrates transformative business and technical value through high-contrast split panels.

```
┌───────────────────────────────────┐ ┌───────────────────────────────────┐
│ BEFORE: FRAGMENTED & MANUAL       │ │ AFTER: UNIFIED & AUTONOMOUS       │
│ (Rose/Slate Muted Theme)          │ │ (Emerald/Violet High-Vibrancy)    │
│                                   │ │                                   │
│ ❌ 14-day manual release cycle    │ │ ✅ 12-minute automated CI/CD push │
│ ❌ High latency database queries  │ │ ✅ Sub-millisecond cached memory  │
│ ❌ 32% user dropoff at signup     │ │ ✅ 89% conversion rate across web │
│ ❌ Siloed departmental data       │ │ ✅ Unified lakehouse data mesh    │
└───────────────────────────────────┘ └───────────────────────────────────┘
  Left Card: Width 790px, Left 140px    Right Card: Width 790px, Left 990px
```

- **Left Card ("Before"):** Background `#FFF5F5`, border `2px solid #FECDD3` (Rose-200), pain point rows with red badge icons.
- **Right Card ("After"):** Background `#FFFFFF`, border `2.5px solid #7C3AED` (or `#10B981`), shadow `0 20px 40px -10px rgba(124, 58, 237, 0.12)`, proof benefit rows with checkmarks and bold metrics.
- **Optional Wipe Mode:** Interactive slider wipe handle overlaying two visual states.

---

### 3.5 Layout 5: USP Strikethrough Strike Slide (`type: "usp-strike"`)
Delivers a powerful differentiator by rejecting industry malpractice and presenting 3 verifiable commitments.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│   Care that comes                                                           │
│   through the door,                                                         │
│   not the inbox.  (with line-through on "the inbox")                        │
│                                                                             │
│ ┌───────────────────────────┬───────────────────────────┬─────────────────┐ │
│ │ [Clock Icon]              │ [User Icon]               │ [MapPin Icon]   │ │
│ │ A TIME                    │ A NAME                    │ A PLACE         │ │
│ │ A visit within 24 hours.  │ The same face every time. │ Anywhere in WA. │ │
│ └───────────────────────────┴───────────────────────────┴─────────────────┘ │
└─────────────────────────────────────────────────────────────────────────────┘
```

- **Headline Typography:** `124px` Ubuntu bold, line-height `1.02`, letter-spacing `-0.03em`.
- **Editorial Strikethrough:** `textDecoration: "line-through"`, `textDecorationColor: "hsl(var(--pres-accent) / 0.7)"`, `textDecorationThickness: 6px`.
- **3-Point Proof Cluster:** Bottom horizontal card strip, Height `180px`, gap `32px`.

---

### 3.6 Layout 6: SaaS Pricing & Metric Proof Slide (`type: "pricing"`)
Transparent commercial engagement tiers with featured plan elevation.

```
┌──────────────────────┐ ┌──────────────────────┐ ┌──────────────────────┐
│ STANDARD SQUAD       │ │ SCALE PARTNER (HOT)  │ │ ENTERPRISE PLATFORM  │
│ $4,500 / month       │ │ $8,900 / month       │ │ Custom Contract      │
│                      │ │                      │ │                      │
│ • 1 Principal Lead   │ │ • 1 Tech Director    │ │ • Full Cross-Func    │
│ • 2 Senior Devs      │ │ • 4 Fullstack Engs   │ │ • Dedicated PM & QA  │
│ • Async standups     │ │ • Daily sync & Slack │ │ • 99.99% SLA Uptime  │
│ [Get Started]        │ │ [Select Scale Tier]  │ │ [Contact Leadership] │
└──────────────────────┘ └──────────────────────┘ └──────────────────────┘
  Left: 160px, W: 500px    Left: 700px, Scale 1.03  Left: 1240px, W: 500px
```

- **Tiers:** 3 Columns, Width `500px` each, gap `40px`, Top `270px`, Height `700px`.
- **Featured Plan ("Hot"):** Border `2.5px solid #7C3AED`, shadow `0 24px 48px -12px rgba(124, 58, 237, 0.18)`, `scale: 1.03`, "MOST POPULAR" gradient ribbon.
- **Price Figures:** `48px` Ubuntu bold.
- **CTA Buttons:** `h-13` (52px), rounded `12px`, full card width.

---

### 3.7 Layout 7: Steps Chain & Process Roadmap Slide (`type: "steps-chain"`)
Maps temporal deployment sequences and engineering milestones.

```
  (1) Discovery  ─────► (2) Architecture ─────► (3) Development ─────► (4) Launch
  ┌──────────────┐     ┌──────────────┐     ┌──────────────┐ ┌─────────────┐
  │ 2 Weeks      │     │ 3 Weeks      │     │ 8 Weeks      │ │ Continuous  │
  │ • Auditing   │     │ • Schema DDL │     │ • Front/Back │ │ • Automated │
  │ • User Flows │     │ • Tech Stack │     │ • CI/CD Pipe │ │ • Security  │
  └──────────────┘     └──────────────┘     └──────────────┘ └─────────────┘
```

- **Horizon Line:** Top `364px`, Left `140px`, Width `1640px`, Height `3px`, background `#E2E8F0` with gradient progress fill.
- **Step Badges:** `48×48px` circle centered on the line, background `#7C3AED`, white bold numeral.
- **Cards:** Top `410px`, Height `440px`, Width `370px` each, horizontal gap `53px`, padding `28px`.
- **Duration Pills:** Background `#F1F5F9`, text `#475569`, `14px` bold.

---

### 3.8 Layout 8: Social Proof & Testimonials Slide (`type: "testimonials"`)
Consolidates enterprise trust and executive endorsements.

- **Dual Testimonial Cards:** Top `280px`, Height `460px`, Width `790px` each, gap `40px`, Left `140px`.
- **Quote Typography:** `26px` Poppins italic, color `#1E293B`, line-height `1.5`.
- **Author Block:** Avatar circle `60×60px`, Name `20px` bold, Title `15px` muted.
- **Bottom Partner Logo Bar:** Top `820px`, Left `140px`, Width `1640px`, Height `100px`, flex row, justify space-around, grayscale opacity `0.6` hover `1.0`.

---

### 3.9 Layout 9: Talent Funnel & Capability Stack Slide (`type: "talent-funnel"`)
Visualizes selective vetting pipelines or multi-tiered architectural capability layers.

- **Funnel Geometry:** 4 Progressively narrowing horizontal bands (Top `280px` to `780px`).
- **Stage 1 (Top / Widest):** Width `1640px`, Height `110px`, "Top 3% Global Engineering Talent Pool".
- **Stage 2:** Width `1380px`, Height `110px`, "Algorithmic & Architecture Vetting".
- **Stage 3:** Width `1120px`, Height `110px`, "Production Code Simulation & Codebase Audit".
- **Stage 4 (Bottom / Narrowest):** Width `860px`, Height `110px`, "High-Velocity Client Deployment".

---

### 3.10 Layout 10: 3-Point Master Cards Slide (`type: "bullets"`)
The ground-truth sample layout pairing an editorial statement, 3 structured bullet cards with icon badges, and a right-hand focal visual plate.

- **Left Column (Width: 900px, Left: 140px):**
  - Pill Marker: Left `140px`, Top `180px`, `52×5px` violet pill.
  - Headline: Left `140px`, Top `210px`, `56px` Ubuntu bold.
  - Subtitle: Left `140px`, Top `320px`, `24px` Poppins.
  - 3 Bullet Cards: Top `420px`, `530px`, `640px`. Height `88px` each, icon container `48×48px` circle, text `20px` Poppins medium.
- **Right Column (Width: 700px, Right: 140px):**
  - Photographic hero image with feathered gradient mask on left edge.
  - Ambient radial glow aura (`rgba(124, 58, 237, 0.15)`).

---

## 4. Anti-Hallucination & Quality Verification Checklist

- [ ] All layout dimensions are explicitly declared on the `1920×1080` canvas.
- [ ] ZERO text is flattened or baked into images.
- [ ] Headings strictly use `Ubuntu`, body text uses `Poppins`, and metadata uses `JetBrains Mono`.
- [ ] Step reveals preserve dimmed state (`opacity: 0.15; filter: blur(2px)`) until activated.
- [ ] Pricing Hot tiers use `scale: 1.03` with a gradient border and elevated shadow.
- [ ] Before/After split uses high-contrast rose muted vs violet/emerald vibrant styling.
