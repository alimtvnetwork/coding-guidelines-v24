# 38 — Card and Pricing Components

> **/goal** Specify 3-tier modular pricing tables and high-trust editorial floating cards with exact optical elevation, badges, and LESS mixins.
> **/learn** Pricing tier layout, popular card scaling, and floating container hierarchy live here. Curriculum cards and interactive rows live in `23-building-block-components.md`.

**Version:** 1.0.0
**Status:** Active

---

## 0. Anti-hallucination

If a dimension, color, shadow, or tag is not in this file or a file it cites, do not invent it.

---

## 1. Scope

Use this file when constructing 3-tier subscription/service pricing matrices and floating editorial card rows.

Do not use this file for slide deck layouts (see `28-slide-layouts.md` and `34-slide-layout-catalog.md`).

---

## 2. Tokens

| Token | Hex | HSL | Role |
|---|---|---|---|
| `@color-ground-light` | `#FFFFFF` | `hsl(0, 0%, 100%)` | Pure white resting background |
| `@color-border-subtle` | `#E2E8F0` | `hsl(214, 32%, 91%)` | Standard hairline border |
| `@color-popular-bg` | `#0B1329` | `hsl(224, 58%, 10%)` | Dark obsidian ground for popular tier |
| `@color-popular-border` | `#8B5CF6` | `hsl(258, 90%, 66%)` | Electric violet border for highlighted tier |
| `@color-badge-gold` | `#FACC15` | `hsl(48, 96%, 53%)` | Most popular pill badge background |
| `@color-badge-ink` | `#0F172A` | `hsl(222, 47%, 11%)` | Contrast dark text on gold badge |
| `@color-accent-rose` | `#F43F5E` | `hsl(350, 89%, 60%)` | Editorial underline highlight gradient start |
| `@color-accent-pink` | `#EC4899` | `hsl(330, 81%, 60%)` | Editorial underline highlight gradient stop |

---

## 3. Components

### 3.1 PricingTierCard

- **Name and job:** `PricingTierCard`. Represents an individual plan tier in a 3-column pricing matrix.
- **Anatomy:** Card container, tier title, subtitle/deck, price display, billing period label, feature checklist with check icons, CTA action button, and optional top badge.
- **Geometry:** Resting card: padding `2.25rem 1.75rem`, radius `1.25rem` (`20px`), border `1px solid #e2e8f0`. Highlighted tier card: `border: 2px solid #8b5cf6`, `transform: scale(1.04)`, `z-index: 2`. Badge: `top: -12px`, padding `0.25rem 0.875rem`, radius `9999px`.
- **Tokens:** Standard card fill `#ffffff`, text `#0f172a`, muted `#64748b`. Popular card fill `#0b1329`, title `#ffffff`, check icon background `rgba(139, 92, 246, 0.25)`, CTA button fill `#8b5cf6`.
- **States:** Resting: standard shadow `0 4px 6px -1px rgba(15, 23, 42, 0.05)`. Hover: `transform: translateY(-4px)` (or `scale(1.04) translateY(-4px)` on popular), shadow `0 20px 40px -15px rgba(15, 23, 42, 0.12)`.
- **Motion:** `300ms cubic-bezier(0.16, 1, 0.3, 1)` on transform and box-shadow.
- **Reduced motion:** `transform: none`, instant state shift.
- **Accessibility:** Must include semantic lists `<ul>` / `<li>` for features and accessible button labels naming the tier (`Sign up for Pro`).
- **Closed options:** Resting tier, Popular highlighted tier. Anything else does not exist.
- **Do not:** Do not highlight more than one tier in the same 3-column group. Do not omit feature check icons.

### 3.2 EditorialFloatingSection

- **Name and job:** `EditorialFloatingSection`. Centered editorial section with pill eyebrow, gradient underline title, and a 3-card floating container row.
- **Anatomy:** Outer section, centered header block (badge pill, display title with underline span), 3-column floating card grid.
- **Geometry:** Header max-width `680px`. Title `clamp(2rem, 4vw, 3rem)`. Underline: height `3px`, radius `9999px`. Floating cards: padding `2.25rem`, radius `1.5rem` (`24px`), border `1px solid #f1f5f9`.
- **Tokens:** Background `#ffffff`, title `#0f172a`, underline `linear-gradient(90deg, #f43f5e 0%, #ec4899 100%)`, shadow `0 10px 30px -10px rgba(15, 23, 42, 0.04)`.
- **States:** Resting: static elevation. Hover: `transform: translateY(-6px)`, shadow `0 25px 50px -12px rgba(15, 23, 42, 0.10)`.
- **Motion:** `300ms cubic-bezier(0.16, 1, 0.3, 1)`.
- **Reduced motion:** `transform: none`, instant shadow shift.
- **Accessibility:** AA contrast between background and text; heading level hierarchy preserved.
- **Closed options:** 3-column desktop layout collapsing to 1-column below 768px. Anything else does not exist.
- **Do not:** Do not add drop shadows directly to pure white text without dark container backdrop.

---

## 4. Closed Sets

Allowed pricing grid layouts:

1. `three-column-grid` (desktop `>=1024px`: 3 equal columns, gap `1.5rem`)
2. `single-column-stack` (mobile `<1024px`: stacked 1 column, gap `2rem`)

anything else does not exist.

Allowed tier roles:

1. `standard` (resting white surface)
2. `featured` (scaled 1.04, dark obsidian ground, gold badge)

anything else does not exist.

---

## 5. Refusal List

- Do not scale multiple tiers simultaneously in the same pricing table.
- Do not use loose raw gray borders without checking against `@color-border-subtle` (`#e2e8f0`).
- Do not mix dark obsidian navy grounds into resting light tiers.

---

## 6. Sibling References

- Foundational building block components (curriculum cards, syllabus rows): [`23-building-block-components.md`](./23-building-block-components.md)
- Page assembly and band rhythms: [`25-page-assembly.md`](./25-page-assembly.md)
- Slide layout catalog and SaaS pricing slide: [`34-slide-layout-catalog.md`](./34-slide-layout-catalog.md)
