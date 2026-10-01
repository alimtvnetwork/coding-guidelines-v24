# Write and Enhance a Design Spec That a Blind AI Can Follow

> **Prompt Version:** 2.0.0
> **Trigger keywords:** `design-spec`, `enhance-design-spec`, `ui-spec`, `slide-spec`, `blind-spec`, `spec-creator`, `spec-enhancer`

**/goal** Autonomously ingest source designs (codebases, specifications, presentation decks, UI components, screenshots) and author or enhance public design-system specifications so comprehensive, mathematically exact, and rigorous that any "blind AI" or human engineer with zero external context can follow them blindly to design websites, presentations, blogs, animations, menus, buttons, builder modes, or images without guessing a single value or hallucinating.

**/learn** A specification is finished ONLY when every component, token, physics parameter, layout slot, and interaction timing in the source has its exact values written down in structured tables, every gap is audited, and the completion gate at the end passes 100%. A file that merely points at another file or leaves vague placeholders does not count as complete.

---

## 1. Required Execution Inputs

Before beginning analysis or writing, establish these parameters:

| Input Parameter | Definition & Constraints | Default Value |
|:---|:---|:---|
| `SOURCE` | Read-only folder paths or files containing reference implementations (code, specs, decks). | User-provided or repository inspection |
| `TARGET` | Spec directory where public design system files are created or enhanced. | `02-spec/07-design-system/` |
| `MODE` | `create` (authoring new specifications) or `enhance` (fixing and expanding existing specs). | `enhance` |
| `DOMAINS` | Any of: `website`, `blog`, `menu`, `button`, `color`, `type`, `spacing`, `motion`, `sections`, `slide-deck`, `slide-builder`, `site-builder`, `image`. | All domains present in source |
| `BANNED_NAMES` | Client, product, company, private repository, and individual person names that MUST NEVER appear in public specifications. | Mandatory anonymization |

---

## 2. Hard Architectural Rules (Zero-Tolerance Violations)

1. **Total Ban on Invented Values:** Every hex, HSL, RGB, OKLCH, pixel, millisecond, cubic-bezier easing, character stagger rate, keycode, and default limit MUST originate from `SOURCE` or grounded mathematics. If the source is silent, explicitly state: `Not specified in source. Do not invent.` Never write "about", "around", "roughly", or "~" in place of a concrete number.
2. **Strict Anonymization Mandate:** Never leak names from `BANNED_NAMES` (no client company names, internal repositories, or personal names) into public specifications, comments, sample text, or image paths. Always use generic enterprise designations (e.g., `{BRAND}`, `{COMPANY}`, `Executive Leadership`, `Enterprise Platform`).
3. **Pure DOM Text Rendering Mandate (Zero Baked-In Text):** In websites, presentations, and UI designs, all headlines, body copy, bullets, metrics, credentials, and captions MUST be rendered as live, selectable DOM HTML elements (`<h1>`, `<p>`, `<span>`) styled with CSS typography tokens. Raster images (`.png`, `.jpg`, `.webp`) are strictly reserved for photographic visual plates, author avatars, and partner logos.
4. **Strict Relative Git Paths Only:** All file paths, markdown links, subtask paths, and citations MUST be strictly relative paths starting from the repository root (e.g., `02-spec/07-design-system/11-button-system.md`). TOTAL BAN on absolute filesystem paths (`C:\...`, `/home/...`) and `file:///` URIs.
5. **Lowercase File Naming Convention:** All generated specification files MUST use strictly lowercase kebab-case naming (e.g., `34-slide-layout-catalog.md`). No uppercase characters, spaces, or underscores in filenames.
6. **Strict Boolean Standards:** Positive booleans MUST ALWAYS be evaluated implicitly (`if isReady`). NEVER evaluate explicitly against true (`if isReady == true` is banned). NEVER combine positive and negative checks in the same condition.
7. **Scoped Domain Palettes:** Marketing websites, presentation decks, and branding graphics each maintain distinct, dedicated color palettes. Never contaminate a marketing page with slide presentation amber or vice-versa.
8. **Closed Sets Rule:** Every selectable options list (button variants, sizes, slide layout types, theme IDs, pill preset colors, transition kinds, hotkeys, builder modes) MUST be specified as a closed set followed by: `Anything else does not exist.`
9. **Values Beside the Rule:** Never simply write "use the primary color". Document the semantic token, HEX value, RGB coordinates, and HSL values in the same row.
10. **Strict File Size Cap:** Specification files MUST remain bounded (at most 300 lines per file). Complex topics must be broken into discrete, single-responsibility files.

---

## 3. Systematic 5-Phase Workflow

### Phase 1: Comprehensive Source Inventory (Read-Only)
1. Recursively scan `SOURCE` directories, skipping `node_modules`, build artifacts, `.git`, and lockfiles.
2. Read token files, stylesheets, component primitives, hooks, route templates, and existing specs.
3. Construct an exhaustive inventory table recording:
   - **Tokens:** Colors, 10-step gradients, shadows, radii, spacing, gutters, breakpoints, durations, easings.
   - **Typography:** Font families, weights, fluid clamp scales, line heights, tracking, tabular numeral rules.
   - **Components:** Button variants, sizes, sticky glass header, nav links, mega panel, 3D flip card, mobile drawer, footer.
   - **Sections:** Hero layouts, feature grids, drag rails, process stacks, pricing tables, FAQ accordions.
   - **Motion:** Keyframe definitions, reveal thresholds, stagger intervals, reduced-motion fallbacks.
   - **Slides:** 16:9 canvas (`1920×1080`), scale math, 10-step ramps ($S_0$–$S_9$), 10 master layouts, controller HUD, dots, hotkeys, step reveals.
   - **Builders:** Live Slide Builder (`useDeckStore` vs `useEditStore`, 7 layers, bounding boxes) and Website Content Builder (Browse/Edit/Preview, `data-edit-id`, contenteditable, ZIP export).
   - **Images:** Canvas dimensions, safe zones, text caps, aspect ratios.

### Phase 2: Gap & Fidelity Audit
Compare every inventory item against existing specifications in `TARGET`:
- `exact`: Value is explicitly documented and matches source ground truth.
- `wrong`: Value conflicts with source implementation.
- `pointer-only`: Specification merely links to another file without providing the concrete numbers.
- `missing`: Component, token, or behavior is absent from specifications.

### Phase 3: Surgical Specification Authoring & Enhancement
1. Remediate all `wrong` values first, aligning them with ground truth.
2. Author dedicated specification files for all `missing` and `pointer-only` areas.
3. Follow the standardized Spec File Template (Section 5) for every file.

### Phase 4: Bi-Directional Cross-Referencing & Prompt Sync
1. Add every new specification file to the reading checklist in `02-spec/07-design-system/readme.md`.
2. Add every new file to the module inventory table in `readme.md`.
3. Update consuming prompts (`01-prompts/10-ui-and-design/07-follow-ui-ux-design-system.md` and `08-create-slide-deck.md`) so their reading sequences reference all new files.

### Phase 5: Verification & Completion Gate
Execute the automated lint checks and review against the Completion Gate (Section 7).

---

## 4. Domain Completion Standards

A specification domain is considered complete ONLY when all required parameters are explicitly documented:

### 4.1 Colors & 10-Step Gradient Precision System
- Every semantic token name, HEX code, RGB coordinate, and HSL value.
- Mathematical 10-step gradient tables ($S_0$ through $S_9$) for each flagship theme:
  - Linear perceptually uniform lightness formula ($t_i = i/9$).
  - Character-by-character color stepping engine (leading glyph accent, intermediate glyph tint, terminal glyphs primary ink).
- 7 Semantic pill presets (`yellow`, `white`, `purple`, `green`, `blue`, `pink`, `red`) with semantic mappings.
- Relative luminance formula for custom pill ink calculation ($0.299R + 0.587G + 0.114B > 0.6 \implies$ `#0b0b12`, else `#ffffff`).

### 4.2 Typography & Type Scales
- Heading font family (`Ubuntu`), body/UI font family (`Poppins`), and micro-label/eyebrow font family (`JetBrains Mono`).
- Complete fluid type scale with `clamp()` formulas: `text-mega`, `text-h1`, `text-h2`, `text-h3`, `text-h4`, `text-lead`, `text-eyebrow`, `text-stat`, `text-numeral`.
- Line-height ratios, negative tracking on large headings, uppercase tracking on mono eyebrows.
- Prohibition on `@import` in CSS (fonts must load via `<link>` in root).

### 4.3 Button System
- 6 Variants: `primary` (shine-sweep on gradient accent), `solid` (shine-sweep on brand primary), `outline` (pointer-fill on hover), `glass` (gradient-ring, frosted acrylic), `ghost`, `link`.
- 4 Sizing scales: `sm` (36px), `md` (44px), `lg` (52px), `icon` (44×44px).
- Magnetic cursor pull physics (`strength: 0.22`).
- Hardware-accelerated CSS3 interactions: `.shine-sweep` keyframes, `.pointer-fill` radial fills, `.gradient-ring` border masks.
- Presentation controller buttons: 40×40px round action triggers, counter with tabular numbers, deep link share, fullscreen toggle.
- Header shine pill button (`HeaderShineButton` for dynamic shrinking headers).

### 4.4 Avant-Garde Navigation & Mega Menu System
- Header bar: `72px` height, sticky `top-0 z-50`, `bg-background/90`, `backdrop-blur-xl`.
- Scroll threshold: Viewport scroll > `12px` activates `border-b border-border shadow-[var(--shadow-card)]` easing over `420ms [0.16, 1, 0.3, 1]`.
- Alternative shrinking pill header: `1240px` transparent container to `900px` white pill after scroll.
- `SlideSwapLabel` nav links: Per-character vertical slide swap on hover with `stagger: 0.04s` for nav (`0.018s` for buttons), top glyph `-110%`, duplicate bottom glyph `0%` over `520ms var(--ease-out)`. Collapses to plain span under reduced motion.
- Underline indicator: `1px` rule, `origin-left scale-x-0 group-hover:scale-x-100` over `240ms`.
- Chevron indicator: `14px` (`size-3.5`), rotates `180deg` on panel open over `240ms`.
- Pointer Safe Region: `pad = 14px` buffer around header and panel rects; pointer within cancels close; pointer leaving schedules close with `220ms` debounce. Esc key closes immediately.
- MegaPanel dropdown: `top-full pt-3 z-40`, enters from `y: -8, scale: 0.985` over `260ms [0.16, 1, 0.3, 1]`. Multi-column grid templates (3+ groups, 2 groups, 1 group). Group entrance `y: 8 -> 0` (`0.05s + gi*0.05s`). Link entrance `x: -6 -> 0` (`0.08s + gi*0.05s + li*0.03s`).
- Growing left hairline: `1px`, `scaleY(0) -> scaleY(1)` with gradient accent over `420ms`. Trailing arrow `size-3.5` transitions `-translate-x-1 opacity-0 -> translate-x-0 opacity-100`.
- 3D Flip Promo Card: `min-h-[220px]`, `perspective: 1400px`, `transform-style: preserve-3d`, flips `rotateY(180deg)` over `820ms`.
- Sticky-safe mobile drawer: Wheel and touchmove suppression on window **without** setting `overflow: hidden` on `body`.

### 4.5 Slide Presentation Engine & 10 Master Layouts
- Pure DOM text mandate: Absolute ban on baked-in raster text.
- Virtual canvas: `1920 × 1080` (16:9), vector scale $\min(\text{w}/1920, \text{h}/1080)$, `transform-origin: center center`.
- Floating Controller HUD: Fixed `top: 32px, right: 32px`, height `56px`, radius `9999px`, 3 divided groups, top 4px progress bar, bottom dot row (active `28×8px`, inactive `8×8px`).
- 10 Master Layout Models:
  1. `title`: 78px headline, category pill, subtitle, presenter bio card (64×64 avatar), bottom organic SVG dual wave ribbon.
  2. `executive-persona`: Asymmetric portrait staging, halftone matrix, character-shaded hero name (104px), LinkedIn card, location tag.
  3. `key-player`: 3–4 member grid, 380×380 portraits, roles, bios, social badges.
  4. `before-after`: Rose-200 negative card vs Violet/Emerald positive card, pain points vs metrics, image wipe.
  5. `usp-strike`: 124px statement with 6px strikethrough rejecting industry practice, paired with 3-point proof cluster.
  6. `pricing`: 3-tier card grid, featured Hot tier with `scale: 1.03` and gradient border, price figures 48px, full-width CTA.
  7. `steps-chain`: 4-phase horizontal roadmap, numbered 48×48px step badges, progress horizon line, duration pills.
  8. `testimonials`: Dual quote cards in 26px italic Poppins, author avatars, bottom partner logo bar with grayscale hover.
  9. `talent-funnel`: 4 progressively narrowing capability bands (1640px down to 860px).
  10. `bullets`: Ground-truth 3 bullet cards with 48×48px icon containers paired with right-side photographic plate.

### 4.6 Slide Builder Mode & Canvas Inspector
- Dual-store separation: `useDeckStore` (persisted to localStorage `deck-v1`) vs `useEditStore` (ephemeral editor state: selected element, undo/redo stacks).
- 7 Visual layers: Base -> Watermarks -> Media -> DOM Typography -> Ink Annotations -> Selection Overlays -> Inspector HUD.
- Hotkeys: `B`/`E` toggle builder mode, `Tab` cycle elements, `Cmd/Ctrl+Z` undo, `Cmd/Ctrl+Shift+Z` redo, `Escape` deselect, `1`–`4` theme switch.
- Bounding box overrides: `boxes: Record<string, EditBox>`.
- Audio cue engine: Swoosh (120ms debounce), click (80ms debounce), tap (45ms debounce).
- Headless Chromium print-ready PDF export.

### 4.7 Website Content Builder Mode
- Zero-server, client-side visual editor.
- 3 Operating modes: Browse, Edit, Preview.
- Element tagging: `data-edit-id="[page]-[section]-[element]"` and `data-source-file="path/to/content.ts"`.
- Text editing: `contenteditable` with caret preservation, HTML sanitization, floating `[[accent]]` gradient toolbar.
- Media replacement: Upload/drop modal, alt text editor, SEO filename slugification, IndexedDB blob storage.
- Navigation editing: Edit nav labels, URLs, and dropdown descriptions inline.
- Deterministic export: Downloadable ZIP archive containing `changes.md` and `assets/` subfolder.

---

## 5. Standardized Spec File Template

Every newly created or enhanced specification file MUST follow this structure:

```markdown
# NN — [Topic Title]

> **/goal** One sentence defining what an AI or human engineer can build using this specification.
> **/learn** One sentence summarizing key tokens, formulas, and references to sibling specifications.

**Version:** 4.0.0
**Status:** Active
**AI Confidence:** High
**Ambiguity:** None

---

## 1. System Overview & Scope
[When to use this specification, architectural pillars, and anti-hallucination boundaries]

---

## 2. Token Registry & Exact Mathematics
[Tables of tokens with Hex, HSL, RGB, OKLCH, dimensions, and mathematical formulas]

---

## 3. Component Taxonomy & Architectural Blueprint
[Component anatomy, geometric dimensions, states, motion curves, and TypeScript interfaces]

---

## 4. Closed Sets & Valid Options
[The complete enumerated list of allowed IDs, variants, or modes, followed by:]
*Anything else does not exist.*

---

## 5. Reference Implementation
[Tested, framework-agnostic code snippet or component blueprint]

---

## 6. Anti-Hallucination & Quality Verification Checklist
[Exhaustive checkboxes guaranteeing that every required number, property, and rule is satisfied]
```

---

## 6. Automated Verification Commands

Run these checks from the repository root:

```powershell
# 1. Verify no absolute paths or file:/// URIs
Get-ChildItem -Path "02-spec/07-design-system", "01-prompts/10-ui-and-design" -Filter "*.md" | Select-String -Pattern "([A-Za-z]:\\|file:///|/home/|/Users/)"

# 2. Verify no vague numbers (about, roughly, approx)
Get-ChildItem -Path "02-spec/07-design-system", "01-prompts/10-ui-and-design" -Filter "*.md" | Select-String -Pattern "\b(about|around|roughly|approx\.?)\s+\d|~\s?\d"

# 3. Check line counts (must be <= 300 lines per file)
Get-ChildItem -Path "02-spec/07-design-system/*.md" | ForEach-Object { "{0,4} lines: {1}" -f (Get-Content $_.FullName).Count, $_.Name }
```

---

## 7. Completion Gate

An AI agent MUST NOT report completion until every gate is verified:

- [ ] Every inventory row from the source has been mapped to an exact specification with concrete numbers.
- [ ] NO `pointer-only` stubs remain; all referenced values are fully declared.
- [ ] Zero baked-in text: Pure DOM typography mandate is enforced across websites and slide decks.
- [ ] The Avant-Garde navigation system includes 72px sticky header, 12px scroll threshold, safe-region math, `SlideSwapLabel`, growing left hairlines, and 3D flip card.
- [ ] Universal button system includes 6 variants, 4 sizes, magnetic cursor physics, `.shine-sweep`, `.pointer-fill`, and controller buttons.
- [ ] Slide presentation engine includes 16:9 canvas scaling, 10-step gradient tables ($S_0$–$S_9$), 10 master layout models, and Live Slide Builder Mode.
- [ ] Website Content Builder Mode includes 3 modes, `data-edit-id`, contenteditable caret preservation, and Markdown+ZIP export.
- [ ] No banned or private company names exist in public specification files.
- [ ] All file paths are strictly relative paths from the git root.
- [ ] All new files are indexed in `02-spec/07-design-system/readme.md`.
- [ ] Consuming prompts (`07-follow-ui-ux-design-system.md` and `08-create-slide-deck.md`) list all new specification files in their reading sequences.
