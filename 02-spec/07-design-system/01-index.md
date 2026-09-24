# AI-Adaptable Design System

> **/goal** Master and enforce the architectural standards, specifications, and CI/CD validation rules for 07 Design System.
> **/learn** Read the sequentially ordered specification files in this directory, follow the actionable CI/CD checklist, and apply mandatory rules before generating code.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Read and understand all numbered specifications under `07-design-system/`.
- [ ] `/learn` Adhere strictly to `.ai-memory/folder-structure.md` and `.ai-memory/strictly-avoid.md`.
- [ ] `/goal` Verify zero explicit `true` boolean evaluations and no mixed-polarity conditionals.
- [ ] `/learn` Run all local verification linters via `python 03-ai-scripts/06-cicd-local-runner.py`.

. **CRITICAL AI INSTRUCTION:** This `01-index.md` file is the primary entry point for this directory. AI agents MUST read this file first before exploring other files in this folder.

**Version:** 4.0.0
**Updated:** 2026-09-24
**Status:** Active
**AI Confidence:** Production-Ready
**Ambiguity:** None

---

## Overview

This is the **canonical design system specification** for the project. It defines all visual behavior, interaction patterns, multi-theme architecture, color tokens, motion rules, and component construction guidance in a single, portable reference. Any AI agent or human contributor reading this specification should be able to:

1. **Reproduce** the current visual language on a new website
2. **Extend** the system with new pages and components that remain visually consistent
3. **Re-theme** the entire design by selecting from the standardized multi-theme catalog (Navy & Purple, VS Code themes, Heatmaps, or Warm Editorial)
4. **Migrate** the system to WordPress or any other CMS without rewriting component logic

The design system follows a **variable-driven architecture**: all colors, spacing, borders, and visual tokens are defined as CSS custom properties (HSL and OKLCH formats) in a single root file. Components never use hardcoded color values — they reference semantic tokens. Changing a token propagates to every component that uses it.

All animations and transitions prioritize **GPU-composited CSS3 transforms and opacity** — no heavy runtime libraries gating initial content paints.

---

## Design Philosophy

| Principle | Description |
|-----------|-------------|
| **Variable-First** | Every color, spacing, and visual property comes from a CSS custom property. No hardcoded values in components. |
| **Semantic Tokens** | Colors are named by purpose (`--primary`, `--accent`, `--muted`), not by value (`--purple`, `--pink`). |
| **HSL & OKLCH Color Models** | All colors use HSL space-separated format with OKLCH depth calculations for seamless lightness/saturation adjustments. |
| **4-Plane Depth Hierarchy** | Clear optical elevation: Base (Plane 0), Raised (Plane 1), Surface (Plane 2), and Elevated (Plane 3). |
| **60 / 30 / 10 Balance** | 60% dominant neutral, 30% structural surfaces/text, 10% purposeful accent. |
| **Single-Accent Von Restorff** | Exactly one prominent accent action per viewport to maintain unmistakable conversion focus. |
| **CSS3 Motion & Carousels** | High-performance transforms, controlled multi-card sliding loops, and instant reduced-motion fallbacks. |
| **Dark/Light Parity** | Every token has both light and dark values. Components never branch on theme — tokens handle it. |
| **Anti-AI-Slop Governance** | Concrete rejection rubrics that ban generic 3-card grids, unanchored heroes, and purple gradient soup. |
| **Portability** | Platform-agnostic. Works natively with React, Tailwind v4, static HTML, WordPress, or modern SSR frameworks. |

---

## Scoring

| Metric | Value |
|--------|-------|
| AI Confidence | Production-Ready ✅ |
| Ambiguity | None 🟢 |
| Health Score | 100/100 |

---

## Keywords

`design-system` · `multi-theme` · `navy-purple` · `vscode-themes` · `heatmaps` · `css-variables` · `4-plane-depth` · `sliding-carousels` · `anti-ai-slop` · `60-30-10-rule` · `von-restorff` · `typography` · `motion-system` · `ai-training`

---

## File Inventory

| # | File | Category | Description |
|---|------|----------|-------------|
| 00 | [01-index.md](./01-index.md) | Overview | This file — design system entry point and index |
| 01 | [02-design-principles.md](./02-design-principles.md) | Principles | Visual philosophy, consistency rules, interaction feel |
| 02 | [03-theme-variable-architecture.md](./03-theme-variable-architecture.md) | Theme | Complete CSS custom property registry — the single source of truth |
| 03 | [04-typography.md](./04-typography.md) | Typography | Font stacks, size hierarchy, weight rules, text spacing |
| 04 | [05-spacing-layout.md](./05-spacing-layout.md) | Layout | Spacing scale, container rules, grid/flex patterns, responsive breakpoints |
| 05 | [06-borders-shapes.md](./06-borders-shapes.md) | Borders | Border thickness, radius, color behavior, state changes |
| 06 | [08-motion-transitions.md](./08-motion-transitions.md) | Motion | CSS3 transition durations, easing, keyframe animations, state transforms |
| 07 | [09-code-blocks.md](./09-code-blocks.md) | Components | Code block rendering, language badges, line interaction, fullscreen |
| 08 | [10-header-navigation.md](./10-header-navigation.md) | Components | Header layout, menu structure, hover underlines, icon transitions |
| 09 | [11-button-system.md](./11-button-system.md) | Components | Button variants, slide text animation, highlight styles |
| 10 | [12-sidebar-system.md](./12-sidebar-system.md) | Components | Sidebar tree, active states, expand/collapse, search |
| 11 | [13-section-patterns.md](./13-section-patterns.md) | Patterns | Reusable section templates (hero, feature, team, CTA) |
| 12 | [14-page-creation-rules.md](./14-page-creation-rules.md) | Guide | Rules for building new pages from the design language |
| 13 | [15-wordpress-migration.md](./15-wordpress-migration.md) | Migration | CMS compatibility notes, block theme mapping, admin theming |
| 14 | [16-theme-catalogue-and-palettes.md](./16-theme-catalogue-and-palettes.md) | Multi-Theme | Navy & Purple, VS Code ecosystem, Heatmaps, and Warm Editorial palettes |
| 15 | [17-theme-tokens.json](./17-theme-tokens.json) | Token Data | Machine-readable theme JSON for programmatic consumption |
| 16 | [18-dark-mode-and-materiality.md](./18-dark-mode-and-materiality.md) | Materiality | 4-plane depth hierarchy, hairlines over shadows, progressive blur, grain, 60/30/10 |
| 17 | [19-modern-motion-and-sliding-interactions.md](./19-modern-motion-and-sliding-interactions.md) | Motion | Controlled sliding carousels, entrance grammar, section rhythm, reduced motion |
| 18 | [20-ai-training-and-checklist-guide.md](./20-ai-training-and-checklist-guide.md) | AI Training | Anti-slop rubric, 5-question above-the-fold contract, step-by-step design checklist |
| 19 | [21-css3-animations-and-interactions.md](./21-css3-animations-and-interactions.md) | Motion & Hover | CSS3 keyframes, cubic-bezier easing, darkish hover shades, line hover highlights |
| 20 | [22-native-css-select-and-border-shapes.md](./22-native-css-select-and-border-shapes.md) | Native Controls | Modern base-select, ::picker(select), and organic border-shape geometry |
| 21 | [23-building-block-components.md](./23-building-block-components.md) | Components | Curriculum card anatomy, 4-cell subgrids, standalone SVGs, LESS mixins |
| 22 | [24-slide-presentation-system.md](./24-slide-presentation-system.md) | Presentation | 16:9 virtual canvas, draggable webcam PIP, step reveals, dual-screen console |
| 97 | [97-acceptance-criteria.md](./97-acceptance-criteria.md) | Testing | Testable criteria for design system compliance |
| 99 | [99-consistency-report.md](./99-consistency-report.md) | Meta | Consistency validation report |

---

## Variable Dependency Architecture

```
┌─────────────────────────────────────────┐
│         CSS Custom Properties           │
│  (index.css :root / .dark)              │
│  --primary, --accent, --background...   │
└──────────────┬──────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────┐
│       Tailwind Config Mapping           │
│  (tailwind.config.ts)                   │
│  primary: "hsl(var(--primary))"         │
└──────────────┬──────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────┐
│       Component Token Layer             │
│  Semantic classes: .prose-spec,         │
│  .code-block-wrapper, .checklist-block  │
│  All use hsl(var(--token)) only         │
└──────────────┬──────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────┐
│       Component States                  │
│  :hover, :focus, :active, .dark         │
│  Transform, opacity, box-shadow shifts  │
│  All driven by tokens + CSS3 transitions│
└──────────────┬──────────────────────────┘
               │
               ▼
┌─────────────────────────────────────────┐
│       Page-Level Composition            │
│  Pages assemble components              │
│  No page-specific colors or overrides   │
│  Consistent spacing rhythm              │
└─────────────────────────────────────────┘
```

---

## How to Re-Theme

1. Open `index.css`
2. Change HSL values in `:root { }` and `.dark { }` blocks
3. Every component, page, and interaction automatically updates
4. No component files need editing

**Example — Switch from purple/pink to teal/amber:**

```css
:root {
  --primary: 175 85% 40%;        /* was: 252 85% 60% */
  --accent: 38 92% 50%;          /* was: 330 85% 60% */
  --heading-gradient-from: 175 85% 45%;
  --heading-gradient-to: 38 92% 55%;
}
```

Every heading gradient, link color, button, code block glow, and hover effect updates instantly.

---

## Cross-References

| Reference | Location |
|-----------|----------|
| CSS Variables Source | `src/index.css` |
| Tailwind Config | `tailwind.config.ts` |
| Spec Authoring Guide | `../01-spec-authoring-guide/01-index.md` |
| Docs Viewer UI Spec | `../08-docs-viewer-ui/01-index.md` |
| Visual Rendering Guide | `../08-docs-viewer-ui/02-features/07-visual-rendering-guide.md` |

---

## Verification

_Auto-generated section — see `02-spec/07-design-system/97-acceptance-criteria.md` for the full criteria index._

### AC-DS-001: Design-system conformance: Index

**Given** Scan `src/` for raw color literals, hard-coded spacing, and untokenized typography.
**When** Run the verification command shown below.
**Then** All visual properties resolve to semantic tokens declared in `index.css` / `tailwind.config.ts`; no `text-white`, `bg-#fff`, or hex literals appear in components.

**Verification command:**

```bash
npm run lint
```

**Expected:** exit 0. Any non-zero exit is a hard fail and blocks merge.

_Verification section last updated: 2026-08-30_
