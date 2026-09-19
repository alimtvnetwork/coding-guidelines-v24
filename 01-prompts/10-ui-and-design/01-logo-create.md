# Logo Design & Branding Generation — Lovable Design Workflow

> **Prompt Version:** 3.1.0
> **Target Environment:** Lovable & Web Design AI Platforms
> **Synchronization:** Main Meta-Repo & Connected Workspaces

This prompt instructs Lovable (and similar web design AI environments) to generate logo, icon, and branding assets for a given product or brand.

---

## Strictly Avoid (Critical Negative Constraints)

> [!CAUTION]
> **TOTAL BAN ON HTML, WEBSITES, PAGES, AND APPLICATION CODE (AUTO-REJECT)**
> The AI MUST strictly avoid generating any web application or frontend/backend code. This prompt is ONLY for branding assets and icons.

1. **NO HTML Files:** NEVER create or output `.html` files (no `index.html`, `app.html`, `test.html`, or any other HTML files).
2. **NO Web Pages or Mockup Apps:** NEVER build landing pages, homepages, dashboard pages, or mock web applications.
3. **NO React / Vue / Svelte Components:** NEVER create `.tsx`, `.jsx`, `.vue`, or `.svelte` components.
4. **NO CSS Stylesheets:** NEVER create external CSS files (`styles.css`, `app.css`). Theme colors must live ONLY in `colors-themes/palette.md`.
5. **NO Server or Script Code:** NEVER create JavaScript/TypeScript scripts, API routes, or backend servers.
6. **NO Tokens.json:** NEVER generate `tokens.json` in `colors-themes/` (it has been completely removed as unnecessary; use `palette.md` only).
7. **NO Solid Backgrounds on PNGs:** NEVER generate PNG icons with solid white, solid black, or solid colored background boxes. All icon PNGs MUST have an alpha-channel transparent background.
8. **NO Uppercase Filenames:** NEVER use uppercase letters in folder names or file names (`README.md`, `Projects/`, `Logo.svg` are BANNED; use `readme.md`, `02-projects/`, `logo.svg`).
9. **NO Base64 Images in SVGs:** NEVER embed raster images or base64 data URLs inside SVG files.

---

## 1. Input Capture & Clarification

Before generating any assets, capture and validate the following inputs:

1. `product_name`: Name of the product, service, or brand.
2. `product_idea`: Brief explanation of the product, its purpose, target audience, and brand voice/tone (e.g., modern, playful, corporate, technical, minimalist).
3. `sample_colors`: Target color hints, hex codes, or preferred palette mood (e.g., "deep navy and electric cyan").
4. `needs_dark_white_variants`: Boolean indicating if dark-mode and white monochrome variants are required (defaults to `true`).
5. `is_animated`: Boolean indicating whether animated assets (`gif-animation.gif`, animated SVG) are requested (defaults to `false` unless explicitly asked).

> [!IMPORTANT]
> If `product_name`, `product_idea`, or `sample_colors` are not provided, **STOP and ask the user** for clarification before proceeding with asset generation.

---

## 2. Directory & File Hierarchy (Strict Lowercase)

All generated assets must follow strict lowercase naming and zero-padded sequence numbers. No uppercase letters are permitted in filenames or folder paths.

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

1. **Root Projects Folder:** All projects live under `02-projects/`.
2. **Project Folder Naming:** `{sequence}-{project-name}` using two-digit zero-padding and kebab-case (e.g., `01-acme-pay`, `02-cloud-sync`).
3. **Never Overwrite:** Never overwrite an existing project folder; always increment the sequence number (`01`, `02`, `03`, ...).
4. **No Tokens.json:** Do not generate `tokens.json`. Color specifications live exclusively inside `colors-themes/palette.md`.
5. **Favicon Placement:** `favicon.ico` and `favicon.png` are placed at the repository root and regenerated per project.
6. **Relative Paths:** All links and image embeds in `readme.md` must use relative paths so GitHub renders them natively.

---

## 3. Asset Specifications

### A. Vector Icons (`icons-svg/`)

- Pure, clean XML/SVG vector code without editor junk or base64 embedded bitmaps.
- Responsive `viewBox` (e.g., `viewBox="0 0 100 100"`), omitting hardcoded `width` and `height` attributes on the root `<svg>`.
- Standard variants:
  - `logo.svg`: Primary full-color logo.
  - `logo-dark.svg`: Optimized for dark backgrounds with light/vibrant strokes and fills.
  - `logo-white.svg`: Pure white monochrome icon (`#ffffff`) for dark surfaces or overlays.

### B. Transparent Raster Images (`icons-image/`)

- All PNG images MUST have a **transparent alpha-channel background** (no solid white, black, or grey boxes around the icon).
- Mockups (`logo-1024-light.png` and `logo-1024-dark.png`) provide visual contrast tests on clean light/dark canvas backdrops.
- **Specification Listing (consecutive with zero blank lines between entries):**
  - `logo-052.png` — 52x52 px transparent icon
  - `logo-128.png` — 128x128 px transparent icon
  - `logo-256.png` — 256x256 px transparent icon
  - `logo-512.png` — 512x512 px transparent icon
  - `logo-1024-light.png` — 1024x1024 px icon mockup on light surface
  - `logo-1024-dark.png` — 1024x1024 px icon mockup on dark surface

### C. Color Themes (`colors-themes/palette.md`)

- A concise Markdown document listing brand colors with HEX, RGB, and HSL values.
- Includes clear guidance on primary brand color, secondary accent, neutral dark, neutral light, and feedback/glow shades.

### D. Favicons (Repo Root)

- `favicon.ico`: Multi-resolution icon for browser tabs.
- `favicon.png`: High-resolution 32x32 px or 64x64 px PNG favicon.

---

## 4. Prompt Execution Modes

### Version 1: Static Logo & Branding Generation (Default)

Executes standard branding creation:
1. Validates user inputs (product name, idea, tone, colors).
2. Creates the project folder under `02-projects/{seq}-{project-name}/`.
3. Creates `icons-svg/` with `logo.svg`, `logo-dark.svg`, and `logo-white.svg`.
4. Creates `icons-image/` with the 6 consecutive transparent PNG sizes.
5. Creates `colors-themes/palette.md`.
6. Generates `favicon.ico` and `favicon.png` at the repository root.
7. Generates `readme.md` in the project directory displaying all assets in a GitHub-compatible table.

### Version 2: Animated Logo & Branding Generation

Executes static branding creation PLUS animated assets:
1. Completes all steps from Version 1.
2. Generates an animated SVG (`logo-animated.svg` in `icons-svg/`) using smooth CSS `@keyframes` or SMIL for subtle motion (e.g., stroke-dasharray draw effect, pulsing glow, or rotational geometry).
3. Generates `gif-animation.gif` placed at the repo root showcasing the animated icon loop (60–120 frames, seamless loop, 24–30 fps).
4. Embeds the animated preview into `readme.md`.

---

## 5. Concrete Asset Examples

### Example 1: Static SVG Logo (`icons-svg/logo.svg`)

```xml
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100" fill="none" role="img" aria-labelledby="logo-title logo-desc">
  <title id="logo-title">AcmePay Logo</title>
  <desc id="logo-desc">Geometric hexagon with interlocking forward arrows in cyan and navy.</desc>
  <defs>
    <linearGradient id="primary-grad" x1="0%" y1="0%" x2="100%" y2="100%">
      <stop offset="0%" stop-color="#06b6d4" />
      <stop offset="100%" stop-color="#3b82f6" />
    </linearGradient>
  </defs>
  <polygon points="50,5 90,27.5 90,72.5 50,95 10,72.5 10,27.5" stroke="url(#primary-grad)" stroke-width="6" fill="none" stroke-linejoin="round" />
  <path d="M35 50 L50 35 L65 50 M50 35 L50 68" stroke="#06b6d4" stroke-width="6" stroke-linecap="round" stroke-linejoin="round" />
</svg>
```

### Example 2: Animated SVG Logo (`icons-svg/logo-animated.svg`)

```xml
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100" fill="none" role="img" aria-labelledby="anim-title anim-desc">
  <title id="anim-title">AcmePay Animated Logo</title>
  <desc id="anim-desc">Animated geometric hexagon with glowing pulse and rotating core.</desc>
  <style>
    @keyframes pulse-glow {
      0%, 100% { stroke-opacity: 0.6; filter: drop-shadow(0 0 2px #06b6d4); }
      50% { stroke-opacity: 1; filter: drop-shadow(0 0 8px #06b6d4); }
    }
    @keyframes rotate-core {
      0% { transform: rotate(0deg); }
      100% { transform: rotate(360deg); }
    }
    .hex-pulse {
      animation: pulse-glow 3s ease-in-out infinite;
    }
    .core-spin {
      transform-origin: 50px 50px;
      animation: rotate-core 12s linear infinite;
    }
  </style>
  <polygon class="hex-pulse" points="50,5 90,27.5 90,72.5 50,95 10,72.5 10,27.5" stroke="#06b6d4" stroke-width="6" fill="none" stroke-linejoin="round" />
  <g class="core-spin">
    <circle cx="50" cy="50" r="16" stroke="#3b82f6" stroke-width="4" stroke-dasharray="6 4" fill="none" />
    <circle cx="50" cy="34" r="4" fill="#06b6d4" />
  </g>
</svg>
```

### Example 3: Color Palette (`colors-themes/palette.md`)

```markdown
# Brand Color Palette: AcmePay

## Primary Colors
| Role | Name | Swatch | HEX | RGB | HSL |
|---|---|---|---|---|---|
| Primary | Cyan Glow | ![#06b6d4](https://placehold.co/15x15/06b6d4/06b6d4.png) | `#06b6d4` | `rgb(6, 182, 212)` | `hsl(189, 94%, 43%)` |
| Accent | Electric Blue | ![#3b82f6](https://placehold.co/15x15/3b82f6/3b82f6.png) | `#3b82f6` | `rgb(59, 130, 246)` | `hsl(217, 91%, 60%)` |

## Neutral & Surface Colors
| Role | Name | Swatch | HEX | RGB | HSL |
|---|---|---|---|---|---|
| Background Dark | Midnight Navy | ![#0f172a](https://placehold.co/15x15/0f172a/0f172a.png) | `#0f172a` | `rgb(15, 23, 42)` | `hsl(222, 47%, 11%)` |
| Surface Light | Pure White | ![#ffffff](https://placehold.co/15x15/ffffff/ffffff.png) | `#ffffff` | `rgb(255, 255, 255)` | `hsl(0, 0%, 100%)` |
```

### Example 4: Project README Display (`02-projects/01-acme-pay/readme.md`)

```markdown
# AcmePay — Branding & Logo Assets

## Overview
AcmePay is a modern financial platform designed for effortless developer billing.

## Vector Logos
| Primary Logo | Dark Mode | Monochrome White |
|:---:|:---:|:---:|
| ![Primary](icons-svg/logo.svg) | ![Dark](icons-svg/logo-dark.svg) | ![White](icons-svg/logo-white.svg) |

## Transparent Icon Sizes
| 52px | 128px | 256px | 512px |
|:---:|:---:|:---:|:---:|
| <img src="icons-image/logo-052.png" width="52" /> | <img src="icons-image/logo-128.png" width="128" /> | <img src="icons-image/logo-256.png" width="256" /> | <img src="icons-image/logo-512.png" width="256" /> |

## Contrast Mockups
| Light Surface (1024px) | Dark Surface (1024px) |
|:---:|:---:|
| <img src="icons-image/logo-1024-light.png" width="300" /> | <img src="icons-image/logo-1024-dark.png" width="300" /> |

## Color Palette
See [Color Palette](colors-themes/palette.md) for hex codes, swatches, and usage rules.
```
