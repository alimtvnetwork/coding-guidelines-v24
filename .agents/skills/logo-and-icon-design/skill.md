---
name: logo-and-icon-design
description: >-
  Autonomously design, generate, and organize logo, icon, and branding assets following strict lowercase conventions, transparent PNG specs, responsive SVG rules, and dual static/animated modes for Lovable and web design platforms.
---

# Logo & Icon Design Workflow — Asset & Branding Standards

> **Skill Version:** 1.0.0
> **Target Environment:** Lovable, Web Design AI Platforms, & Local Design Workflows
> **Synchronization:** Main Meta-Repo & Connected Workspaces

This skill guides the creation, organization, and validation of production-ready logo, icon, and branding assets.

---

## 1. Strict Boundary & Scope

### TOTAL BAN on Website / HTML Generation

- **Never generate full websites, HTML pages, or application code.**
- This workflow is strictly an asset generator producing vector graphics (SVG), transparent raster icons (PNG), color palettes (Markdown), and preview documentation.

### Input Capture & Validation

Always verify the following inputs before generating assets:
1. `product_name`: Brand/product name.
2. `product_idea`: Purpose, target audience, and brand voice/tone (e.g. minimalist, playful, high-tech).
3. `sample_colors`: Target color palette or mood.
4. `needs_dark_white_variants`: Boolean for dark and white variants (default: `true`).
5. `is_animated`: Boolean for animated assets (default: `false` unless requested).

If any core input is missing, prompt the user for clarification before generating files.

---

## 2. Directory Hierarchy (Strict Lowercase)

All files and folders must follow strict lowercase naming with zero-padded two-digit sequence numbers:

```
/ (repo root)
├── 01-prompts/
└── 02-projects/
    ├── 01-{project-name}/
    │   ├── readme.md
    │   ├── icons-svg/
    │   │   ├── logo.svg
    │   │   ├── logo-dark.svg
    │   │   └── logo-white.svg
    │   ├── icons-image/
    │   │   ├── logo-052.png
    │   │   ├── logo-128.png
    │   │   ├── logo-256.png
    │   │   ├── logo-512.png
    │   │   ├── logo-1024-light.png
    │   │   └── logo-1024-dark.png
    │   └── colors-themes/
    │       └── palette.md
    └── 02-{project-name}/
(repo root assets)
├── favicon.ico
├── favicon.png
└── gif-animation.gif (generated only when animation is requested)
```

### Hierarchy Rules

- Never overwrite an existing project folder; always increment the sequence number (`01`, `02`, ...).
- Do not create `tokens.json`. Color documentation lives exclusively in `colors-themes/palette.md`.
- `readme.md` must use relative links so GitHub renders all assets inline.

---

## 3. Asset Specifications

### A. Vector Graphics (`icons-svg/`)

- Pure, clean XML/SVG vector code.
- Responsive `viewBox` (e.g., `viewBox="0 0 100 100"`), omitting hardcoded `width` and `height` attributes on `<svg>`.
- Accessibility: `<title>` and `<desc>` linked with `aria-labelledby` and `role="img"`.
- Theming: `currentColor` for monochrome icons; clean CSS variables or hex codes for multi-color assets.
- Variants:
  - `logo.svg`: Primary brand logo.
  - `logo-dark.svg`: Optimized for dark surfaces.
  - `logo-white.svg`: Pure white monochrome icon (`#ffffff`).

### B. Transparent Raster Icons (`icons-image/`)

- All PNGs MUST have an **alpha-channel transparent background** (no solid white or black box backgrounds).
- Standard sizes:
  - `logo-052.png` — 52x52 px transparent icon
  - `logo-128.png` — 128x128 px transparent icon
  - `logo-256.png` — 256x256 px transparent icon
  - `logo-512.png` — 512x512 px transparent icon
  - `logo-1024-light.png` — 1024x1024 px mockup on light background
  - `logo-1024-dark.png` — 1024x1024 px mockup on dark background

### C. Color Palette (`colors-themes/palette.md`)

- Document HEX, RGB, and HSL values with visual markdown swatch previews.

### D. Favicons (Repo Root)

- `favicon.ico` and `favicon.png` placed at the repo root and regenerated per project.

---

## 4. Execution Modes

### Mode 1: Static Logo & Branding Generation (Default)

Generates the complete static vector set, transparent PNG sizes, color palette, root favicons, and project `readme.md`.

### Mode 2: Animated Logo & Branding Generation

In addition to static assets, generates:
- Animated SVG (`icons-svg/logo-animated.svg`) using smooth CSS `@keyframes` (pulse, rotation, draw effect).
- Seamless looping GIF (`gif-animation.gif`) placed at repo root.
- Animated asset preview embedded in `readme.md`.

---

## 5. Quality Checklist

- [ ] All file and directory names are strictly lowercase.
- [ ] No HTML pages or website templates were generated.
- [ ] All PNG icons have transparent backgrounds (alpha channel).
- [ ] SVG files have responsive `viewBox` and no hardcoded `width`/`height`.
- [ ] Project `readme.md` uses relative paths to render all assets.
