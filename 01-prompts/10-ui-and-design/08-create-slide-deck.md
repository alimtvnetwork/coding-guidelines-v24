# Create Presentation Slide Deck & Live Builder System

> **Prompt Version:** 2.0.0
> **Trigger keywords:** `create-slide-deck`, `presentation-spec`, `slide-builder`, `slide-layout`, `deck-system`

**/goal** Autonomously architect, construct, and style high-authority, declarative, 16:9 presentation slide decks, interactive presenter HUDs, 10-step precision color ramps, and real-time Slide Builder Modes on the 1920×1080 virtual canvas with pure live DOM typography and zero hallucinations.

**/learn** Before writing slide schemas, components, or JSON decks, AI agents MUST read the sequentially numbered specification files listed below in exact order. All relative paths are from the git repository root.

---

## 1. Mandatory Reading Sequence (Strict Relative Paths)

AI agents MUST sequentially ingest these specification files:

1. `02-spec/07-design-system/24-slide-presentation-system.md` — Responsive coordinate transforms (`1920×1080` canvas), webcam PIP overlay, step reveals, dual-screen presenter console.
2. `02-spec/07-design-system/27-slide-canvas-and-themes.md` — Virtual canvas scale formula, theme schema definitions, JSON deck configuration.
3. `02-spec/07-design-system/28-slide-layouts.md` — Base slot model, closed layout union, and field constraints.
4. `02-spec/07-design-system/29-slide-navigation-and-builder.md` — HUD, transitions, hotkeys (`ArrowRight`, `Space`, `B`/`E`, `G`, `S`), and builder chrome.
5. `02-spec/07-design-system/30-slide-palette-type-and-shell.md` — Dark amber palette, typography scales, spotlight radial glow, shell layers.
6. `02-spec/07-design-system/31-slide-controller-buttons.md` — Floating top-right HUD controller pill, circular 40×40px action buttons, tabular numbers counter, deep link share, fullscreen toggle, bottom dot pagination.
7. `02-spec/07-design-system/32-slide-color-options.md` — 10-step gradient precision system ($S_0$ to $S_9$) across 4 flagship palettes, character-by-character shading engine, 7 semantic pill presets, relative luminance formula, and 9-cell text alignment.
8. `02-spec/07-design-system/34-slide-layout-catalog.md` — The Non-Image Text Mandate (pure DOM typography; zero baked-in text) and the 10 master enterprise slide layouts: Title Hero, Executive Persona, Key Player Bio, Before/After Split, USP Strikethrough, SaaS Pricing, Steps Chain Roadmap, Social Proof, Talent Funnel, and 3-Point Master Cards.
9. `02-spec/07-design-system/35-slide-builder-canvas-inspector.md` — Decoupled dual-store architecture (`useDeckStore` persisted vs `useEditStore` ephemeral), 7 visual canvas stacking layers, builder hotkeys (`B`/`E`, `Tab`, `Cmd+Z`, `1`–`4`), bounding box coordinate overrides (`boxes: Record<string, EditBox>`), acoustic audio cue engine with debouncing, and headless Chromium print-ready PDF export.
10. `02-spec/07-design-system/37-image-specifications.md` — Non-image typography mandate, image plates, safe zones, avatar rules, and asset bounds.
11. `02-spec/07-design-system/38-card-and-pricing-components.md` — SaaS pricing card matrix, feature list hierarchy, and highlighted tier geometry.

---

## 2. The Non-Image Text Mandate (Pure DOM Typography)

> [!CRITICAL]
> **TOTAL BAN ON BAKED-IN TEXT:**
> All headlines, subheads, kickers, bullet points, metrics, author names, credentials, and captions MUST be rendered as live, selectable DOM HTML elements (`<h1>`, `<h2>`, `<p>`, `<span>`) styled with CSS typography tokens.
> Under no circumstances should text be flattened into raster images (`.png`, `.jpg`, `.webp`). Raster images are strictly reserved for photographic hero visual plates, author avatars, and partner logos.

---

## 3. 1920×1080 Virtual Canvas & Adaptive Uniform Scaling

All slide components author geometry on an authoritative reference canvas of `1920 × 1080` pixels:

1. **Runtime Scaling Formula:**
   $$\text{scale} = \min\left(\frac{\text{viewportWidth}}{1920}, \frac{\text{viewportHeight}}{1080}\right)$$
2. **Container CSS Vector Transform:**
   ```css
   .slide-stage {
     width: 1920px;
     height: 1080px;
     transform-origin: center center;
     transform: scale(var(--stage-scale));
     contain: layout paint;
     isolation: isolate;
     will-change: transform;
   }
   ```
3. **Stage Positioning:** Centered both vertically and horizontally in the browser viewport with letterboxing on non-16:9 displays.

---

## 4. Floating Controller HUD & Pagination Rules

The presenter HUD is mounted on the global viewport above all slides (`z-index: 50`), never inside the scaled stage:

```
┌─────────────────────────────────────────────────────────────┐
│  [‹] Prev    5 / 37    [›] Next  │  [⌁] Share  │  [⤢] Full  │
└─────────────────────────────────────────────────────────────┘
  ▲─── Fixed Top-32px Right-32px, Height 56px, Radius 9999px ──▲
```

1. **HUD Geometry:** `fixed top-8 right-8 h-14 rounded-full px-3 py-2 bg-[hsl(240_8%_8%/0.85)] backdrop-blur-md border border-border shadow-[0_8px_24px_rgba(0,0,0,0.4)]`.
2. **Three Control Groups (Divided by `1×24px` borders):**
   - **Navigation Group:** Circular 40×40px `ChevronLeft` + Counter + Circular 40×40px `ChevronRight`. Counter format: `{current} / {total}` in Poppins 500, `20px`, tabular numbers, min width `64px`.
   - **Share Group:** 40×40px `Share2` button. Deep links to `#slide-{current}`. Triggers `navigator.share` or clipboard copy + Sonner toast.
   - **Fullscreen Group:** 40×40px `Maximize2` / `Minimize2` toggle button.
3. **Top Progress Bar:** Fixed `top: 0 left: 0 right: 0 h-1` (4px). Track `hsl(var(--border-subtle))`, fill width `(current / total) * 100%` using the theme gradient.
4. **Bottom Pagination Dots:** Fixed `bottom-8` centered flex row. Inactive dots are `8×8px` circles; the active dot expands into a `28×8px` pill with transition `200ms cubic-bezier(0.2, 0.8, 0.2, 1)`.

---

## 5. 10-Step Precision Gradients & Character Shading

1. **Theme Ramps:** Every deck binds to one of the 4 flagship 10-step tables ($S_0$ through $S_9$) in `32-slide-color-options.md`:
   - `white-violet`: Royal Violet White ($S_0$: Deep Indigo `#101070` to $S_9$: Lilac Fog `#FAF5FF`).
   - `noir-gold`: Corporate Gold ($S_0$: Pure Obsidian `#0F172A` to $S_9$: White Gold `#FFFAE0`).
   - `enterprise-blue`: Slate Blue ($S_0$: Dark Ink Slate `#020817` to $S_9$: Pure White `#FFFFFF`).
   - `clinical-emerald`: Emerald Green ($S_0$: Pine Forest `#063729` to $S_9$: Pure White `#F5FEFA`).
2. **Character-by-Character Shading Engine:** Prominent hero names use per-glyph color stepping:
   - Leading character ($C_0$): Primary accent ($S_4$/$S_5$).
   - Intermediate character ($C_1$): Warm intermediate step ($S_6$/$S_7$).
   - Terminal characters ($C_2..C_n$): Primary ink text ($S_0$/$S_1$).
3. **7 Semantic Pill Presets:** `yellow` (fallback), `white` (dark grounds), `purple` (people/execs), `green` (money/growth), `blue` (facts/steps), `pink` (creative), `red` (risks/costs). Ink follows ITU-R BT.709 relative luminance ($L > 0.6 \implies$ `#0b0b12`, else `#ffffff`).

---

## 6. The 10 Master Slide Layouts

Every slide MUST implement one of the closed layout types from `34-slide-layout-catalog.md`:

1. **`title`:** 78px headline, category pill kicker, subtitle, presenter bio card (avatar 64×64), bottom organic SVG dual wave ribbon.
2. **`executive-persona`:** Asymmetric portrait staging, halftone dot matrix, character gradient hero name (104px), interactive LinkedIn preview card, location tag.
3. **`key-player`:** 3–4 member grid, 380×380px portraits, roles, bios, social links.
4. **`before-after`:** High-contrast split cards: Left ("Before" negative, Rose-200 border, pain points) vs Right ("After" positive, Violet/Emerald border, proof metrics, image wipe).
5. **`usp-strike`:** 124px statement with 6px editorial strikethrough rejecting industry malpractice, paired with horizontal 3-point proof cluster.
6. **`pricing`:** 3-tier card grid, featured "Hot" tier with `scale: 1.03` and gradient border, price figures 48px, full-width CTA buttons.
7. **`steps-chain`:** 4-phase horizontal roadmap, numbered 48×48px step badges, connecting horizon progress line, duration pills, deliverable lists.
8. **`testimonials`:** Dual quote cards in 26px italic Poppins, author avatars, bottom partner logo bar with grayscale hover reveal.
9. **`talent-funnel`:** 4 progressively narrowing capability bands (1640px down to 860px).
10. **`bullets`:** Ground-truth 3-point bullet cards with 48×48px icon containers paired with right-side photographic plate.

---

## 7. Slide Builder Mode & Live Canvas Inspector

1. **Dual-Store Separation:**
   - `useDeckStore` (persisted to localStorage `deck-v1`) holds the authoritative presentation JSON tree.
   - `useEditStore` (ephemeral) manages `selectedElementId`, bounding box coordinates, and undo/redo stacks.
2. **7 Canvas Stacking Layers:** Layer 0 (Base) -> Layer 1 (Watermarks) -> Layer 2 (Media) -> Layer 3 (DOM Typography) -> Layer 4 (Ink) -> Layer 5 (Selection Overlays) -> Layer 6 (Inspector HUD).
3. **Hotkeys:** `B` / `E` toggles builder mode, `Tab` cycles editable elements, `Cmd/Ctrl+Z` undos, `Cmd/Ctrl+Shift+Z` redos, `Escape` clears selection, `1`–`4` quick-switches theme palettes.
4. **Audio Cue Engine:** Swoosh on slide transitions (`120ms` debounce), click on sub-step reveals (`80ms` debounce), tap on typewriter reveals (`45ms` debounce).

---

## 8. Blind-AI Execution Checklist

Before outputting code or completing your task, verify every item:

- [ ] Canvas is strictly `1920×1080` with uniform `min(w/1920, h/1080)` vector scaling.
- [ ] ZERO text is flattened into raster images; all text is live DOM elements.
- [ ] Controller HUD is anchored to fixed `top-8 right-8` on the viewport with 3 divided groups.
- [ ] Counter displays `{current} / {total}` with tabular figures.
- [ ] Dot pagination expands the active dot to `28×8px` and inactive to `8×8px`.
- [ ] Slide layout strictly matches one of the 10 models in `34-slide-layout-catalog.md`.
- [ ] Gradients adhere strictly to the 10-step tables ($S_0$ through $S_9$) in `32-slide-color-options.md`.
- [ ] Character-by-character color shading applies to hero names and title keywords.
- [ ] Builder mode maintains dual-store separation (`useDeckStore` vs `useEditStore`) and 7 visual layers.
- [ ] All relative paths referenced start from the git repository root.
- [ ] NO private or company names are present in public files or examples.
