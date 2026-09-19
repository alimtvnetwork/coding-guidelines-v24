# SVG Icon & Vector Graphic Creation — Design Workflow

> **Prompt Version:** 3.2.0
> **Target Environment:** Lovable & Web Design AI Platforms
> **Synchronization:** Main Meta-Repo & Connected Workspaces

This prompt instructs Lovable (and similar web design AI environments) on the exact technical standards for generating scalable, responsive, and accessible SVG vector icons and graphics.

---

## Strictly Avoid (Critical Negative Constraints)

> [!CAUTION]
> **TOTAL BAN ON HTML WRAPPERS, EMBEDDED BITMAPS, AND SOLID BACKGROUNDS (AUTO-REJECT)**
> The AI MUST strictly avoid generating HTML, web pages, solid background shapes, or embedding raster bitmaps inside SVGs.

1. **NO HTML Files or Wrappers:** NEVER create `.html` files and NEVER wrap SVG code inside `<html>`, `<body>`, `<div>`, or any HTML container. Output ONLY the raw SVG code inside an `xml` or `svg` code block.
2. **NO Web Pages or UI Templates:** NEVER generate a website, landing page, mock application, or UI component when asked for an SVG logo/icon.
3. **NO React / Vue Components:** NEVER generate `.tsx`, `.jsx`, or component files unless explicitly requested.
4. **NO Base64 Raster Images:** NEVER embed base64-encoded bitmap images (`<image href="data:image/png;base64,...">`). All graphics MUST be pure mathematical vector paths, polygons, and curves.
5. **NO Hardcoded Fixed Dimensions on Root:** NEVER include hardcoded `width="..."` and `height="..."` attributes on the root `<svg>` tag that prevent responsive scaling; use `viewBox` instead.
6. **NO Solid Background Elements:** NEVER insert a background `<rect>` or container shape (e.g. `<rect width="100%" height="100%" fill="#000"/>` or `<rect fill="#fff"/>`). The canvas MUST remain 100% transparent. "Dark mode" SVGs invert the strokes/fills to lighter colors, but the background is ALWAYS transparent.
7. **NO Editor Bloat:** NEVER include third-party editor metadata (e.g., `xmlns:inkscape`, `sodipodi:docname`, `adobe:ns`).
8. **NO Unstyled Elements in Monochrome:** NEVER use hardcoded black/white hex fills on monochrome icons; use `currentColor` so the icon inherits text color dynamically.

---

## 1. Clean SVG Architecture

- **Valid XML & SVG Syntax:** Ensure the output is fully valid XML and well-formed SVG markup.
- **Responsive `viewBox`:** Always define an appropriate `viewBox` (e.g., `viewBox="0 0 24 24"` for UI icons, `viewBox="0 0 100 100"` for logos/illustrations).
- **Omit Hardcoded Dimensions:** Remove hardcoded `width` and `height` attributes on the root `<svg>` element so the icon scales responsively within its CSS container.
- **Minimal Grouping:** Do not wrap elements in redundant `<g>` tags unless needed for shared transforms, styling, or animations.
- **No Editor Metadata:** Do not include Adobe Illustrator, Inkscape, or Figma metadata/namespaces (`xmlns:inkscape`, `sodipodi:docname`, etc.).
- **TOTAL BAN on Base64 Images:** NEVER embed base64-encoded raster images (`<image href="data:image/png;base64,...">`). All artwork must be pure vector paths, polygons, circles, and curves.
- **No HTML Wrapping:** Provide only the raw SVG code inside an `xml` or `svg` code block unless explicitly requested.

---

## 2. Accessibility & Semantics

- Always include a `<title>` and `<desc>` element for screen readers.
- Assign matching `id` attributes to `<title>` and `<desc>` and link them via `aria-labelledby` on the root `<svg>`.
- Add `role="img"` to the root `<svg>`.

---

## 3. Theming & Color Strategy

- **Monochrome Icons:** Use `fill="currentColor"` or `stroke="currentColor"` so the icon seamlessly inherits text color from parent CSS.
- **Multi-Color Logos:** Use semantic CSS variables (`var(--primary)`, `var(--accent)`) or clean standard hex codes.
- **Gradients:** Place `<linearGradient>` and `<radialGradient>` definitions inside `<defs>` with semantic, unique IDs.
- **Variants (All 100% Transparent Canvas):** When generating brand sets, provide:
  - Default full-color vector on a transparent canvas.
  - Dark mode variant (`logo-dark.svg`): Inverted, lighter, or vibrant strokes/fills designed for dark themes—on a **100% transparent canvas** (NO dark background `<rect>`).
  - Pure white monochrome variant (`logo-white.svg`): Pure white (`#ffffff`) strokes/fills on a **100% transparent canvas**.

---

## 4. Execution Modes & Concrete Examples

### Version 1: Static Vector SVG Icon (Default)

Generates a clean, production-grade vector icon with full responsive scaling, accessibility tags, and `currentColor` support.

```xml
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" role="img" aria-labelledby="shield-check-title shield-check-desc">
  <title id="shield-check-title">Security Verified</title>
  <desc id="shield-check-desc">A shield icon containing a checkmark indicating verified security status.</desc>
  <path d="M12 22s8-4 8-10V5l-8-3-8 3v7c0 6 8 10 8 10z" />
  <path d="m9 12 2 2 4-4" />
</svg>
```

### Version 2: Animated Vector SVG Icon

Generates an interactive or looping animated SVG using self-contained CSS `@keyframes` embedded within a `<style>` block. Ideal for loading states, hero branding, or interactive micro-animations.

```xml
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100" fill="none" role="img" aria-labelledby="sync-title sync-desc">
  <title id="sync-title">Syncing Data</title>
  <desc id="sync-desc">Two rotating curved arrows indicating real-time data synchronization.</desc>
  <style>
    @keyframes spin-clockwise {
      0% { transform: rotate(0deg); }
      100% { transform: rotate(360deg); }
    }
    @keyframes pulse-stroke {
      0%, 100% { stroke-opacity: 0.5; }
      50% { stroke-opacity: 1; }
    }
    .rotating-group {
      transform-origin: 50px 50px;
      animation: spin-clockwise 3s linear infinite;
    }
    .pulsing-arrow {
      animation: pulse-stroke 1.5s ease-in-out infinite;
    }
  </style>
  <g class="rotating-group">
    <path class="pulsing-arrow" d="M50 15 A35 35 0 0 1 85 50 L75 50 L90 65 L95 50 L85 50" stroke="#06b6d4" stroke-width="5" stroke-linecap="round" stroke-linejoin="round" fill="none" />
    <path class="pulsing-arrow" d="M50 85 A35 35 0 0 1 15 50 L25 50 L10 35 L5 50 L15 50" stroke="#3b82f6" stroke-width="5" stroke-linecap="round" stroke-linejoin="round" fill="none" />
  </g>
  <circle cx="50" cy="50" r="8" fill="#06b6d4" />
</svg>
```

---

## 5. Output Format

- Provide strictly the raw SVG code inside a fenced code block with language identifier `xml` or `svg`.
- Do not output HTML wrappers or surrounding page containers.
