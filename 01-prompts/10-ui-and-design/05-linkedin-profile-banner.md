# LinkedIn Profile Banner Design — Visual Identity & Authority Workflow

> **Prompt Version:** 1.0.0
> **Target Environment:** Lovable, Figma, Canva, Image Generation AI (Flux, Ideogram, Midjourney) & Design AI Platforms
> **Synchronization:** Main Meta-Repo & Connected Workspaces

This prompt instructs design AI platforms on how to craft high-authority, conversion-focused personal LinkedIn profile banners with exact dimensional standards, profile picture collision safe zones, and zero text hallucination.

---

## Strictly Avoid (Critical Negative Constraints)

> [!CAUTION]
> **ZERO TOLERANCE FOR AVATAR COLLISION, TEXT HALLUCINATION & WEB CODE (AUTO-REJECT)**
> The AI MUST strictly avoid generating HTML, placing text in the avatar collision zone, or distorting typography.

1. **NO HTML Files or Web Code:** NEVER generate `.html` files, React components (`.tsx`, `.jsx`), or CSS stylesheets. This is strictly a banner design and vector layout workflow.
2. **NO Text in the Avatar Collision Zone:** NEVER place critical text, headlines, contact info, faces, or logos in the bottom-left area (x: 0 to 380px). On both desktop and mobile, LinkedIn's circular profile picture will obscure anything placed there.
3. **NO Text Hallucinations or Gibberish:** NEVER render garbled glyphs, pseudo-Latin, or misspelled words. Every character in the headline, website URL, and credentials must match the user's input verbatim.
4. **NO Low-Resolution Exports:** NEVER export or specify standard 72 DPI 1x images that blur upon upload. Always specify **2x or 3x Retina resolution** (`3168 x 792 px` or `4752 x 1188 px`).
5. **NO Low-Contrast Text:** NEVER place light text over bright, busy backgrounds or dark text over dark shadows without proper contrast treatment (e.g. solid badge backdrops, dark vignettes, or subtle dropshadows).
6. **NO Saving Binary Images to Repository:** Do NOT save raw binary image files (`.png`, `.jpg`) into the git repository. All instructions, layouts, and SVG vector overlays must live purely in markdown or code.

---

## 1. Dimensional Standards & Safe Zones

LinkedIn applies aggressive cropping and compression to profile banners across devices. All banners must be designed against these precise dimensional specifications:

### Canvas Dimensions

- **Standard Dimensions (1x):** `1584 x 396 px` (4:1 aspect ratio).
- **High-Resolution Retina (2x — Recommended):** `3168 x 792 px` (guarantees razor-sharp text after LinkedIn's upload compression).
- **Ultra High-Resolution (3x):** `4752 x 1188 px`.
- **Target File Size & Format:** Under 8MB; export as uncompressed PNG or high-bitrate WebP.

### Safe Zone Architecture

```
+-----------------------------------------------------------------------------+
|                                    TOP SAFE BAND                            |
|                                                                             |
|   [ AVATAR COLLISION ]            CENTER / RIGHT CONTENT ZONE               |
|   [  DEAD ZONE       ]     - Headline & Value Proposition                   |
|   [  (Profile Pic)   ]     - Core Pillars / Credential Badges               |
|   [                  ]     - Website URL & Call-to-Action                   |
|                                                                             |
+-----------------------------------------------------------------------------+
|<--- 0 to 380px ------>|<----------------- 380px to 1584px ----------------->|
```

- **Dead Zone (Bottom-Left):** From `x: 0` to `x: 380px` and `y: 180px` to `y: 396px`, the circular avatar covers the canvas on desktop and shifts further inward on mobile. Keep this area free of text.
- **Mobile Crop Margin:** Mobile devices crop approximately `100px` from the left and right edges. Keep all critical messaging within the central safe zone (`x: 400px` to `x: 1450px`).

---

## 2. Input Capture & Verification

Before generating any design concepts or layout prompts, capture and validate the following inputs:

1. `full_name`: Person's exact name (e.g., "Jane Doe").
2. `professional_headline`: Core value proposition (e.g., "Fractional CMO | Scaling B2B SaaS from $1M to $10M ARR").
3. `current_company` / `venture`: Current organization, agency, or brand.
4. `core_pillars`: 3–4 key focus areas (e.g., `GTM Strategy • Demand Generation • Team Leadership`).
5. `website_url`: Personal or business website URL (e.g., `https://janedoe.com`).
6. `email_address`: Professional contact email address to embed on the banner (e.g., `jane@janedoe.com`).
7. `qr_code`: Optional scannable QR code destination URL or asset to embed (e.g., linking directly to a calendar booking page, digital vCard, or portfolio).
8. `contact_or_cta`: Explicit call to action (e.g., "Book a strategy audit at janedoe.com", "DM for speaking inquiries").
9. `portrait_photo_or_subject`: Visual description of the person or whether a cutout photo will be placed on the right flank.
10. `brand_colors`: Primary brand palette (e.g., Midnight Navy `#0b0f19`, Vibrant Cyan `#06b6d4`, Pure White `#ffffff`).

> [!IMPORTANT]
> Always ask the user if they want to embed an **email address** or a **scannable QR code** (e.g. for booking or vCard). If `full_name`, `professional_headline`, `website_url`, or `brand_colors` are missing or not provided, **STOP and ask the user** before proceeding.

---

## 3. Visual Layout Versions (Multiple Modes)

### Version 1: Authority & Social Proof Layout (Default)

Ideal for consultants, founders, and executives:
- **Background:** Deep navy or charcoal gradient with subtle geometric or architectural depth of field.
- **Left Flank:** Clean, ambient negative space accommodating the profile picture.
- **Center:** Bold two-line headline in high-contrast white and gold/cyan accent, followed by 3 core pillars separated by bullet dots.
- **Right Flank:** High-resolution subject portrait cutout or authority proof (e.g. "Featured in Forbes, TechCrunch" or book cover mockup).
- **Lower Band:** Clear website URL and rounded pill call-to-action badge.

### Version 2: Minimalist Executive Layout

Ideal for enterprise leaders, investors, and board members:
- **Background:** Monochromatic dark slate or textured matte finish with subtle linear accent lighting.
- **Left Flank:** Clean dark negative space.
- **Center-Right:** Single, powerful positioning statement in elegant modern typography.
- **Right Flank:** Minimalist brand logo mark and official corporate website URL.

### Version 3: Visual Showcase / Product Showcase Layout

Ideal for creators, authors, and product builders:
- **Background:** Studio gradient with warm ambient lighting.
- **Center:** Creator name and hook headline.
- **Right Flank:** 3D product mockup (e.g. SaaS dashboard screen, published book, or podcast badge).
- **Lower Third:** Credential badges with clean line-art icons and exact website URL.

---

## 4. Typography, Contrast & Anti-Hallucination Rules

- **Verbatim Text Enforcement:** All names, titles, and URLs must be explicitly quoted and verified character-by-character.
- **Contrast Ratios:** Text must meet WCAG AAA standards against the banner backdrop. Use darkened vignettes behind text layers or solid badge containers for URLs.
- **Font Hierarchy:**
  - *Headline:* Bold geometric sans-serif (e.g. Montserrat Bold, Inter ExtraBold, Poppins) at large scale.
  - *Subtitles:* Clean medium grotesque sans-serif (e.g. Inter Medium, Roboto).
  - *Accents / CTAs:* Subtle vibrant accent color (e.g. `#facc15` Golden Yellow, `#06b6d4` Electric Cyan) on a high-contrast container.

---

## 5. Ready-to-Use Templates

### Template 1: Midjourney / Flux Background Generation Prompt

```text
Wide panoramic 4:1 banner background for LinkedIn profile, dimensions 3168x792. Deep navy blue and charcoal dark modern architectural interior with subtle amber bokeh lights, clean minimalist negative space on the left and center for typography overlay, cinematic soft directional lighting from the right, commercial photography, premium executive aesthetic, 8k resolution --ar 4:1 --style raw
```

### Template 2: Figma / SVG Typography Overlay Specification

```markdown
# Canvas: 3168 x 792 px (2x Retina)
- Avatar Dead Zone: x: 0 to 760 px (Keep empty)
- Primary Headline: "Scaling B2B SaaS to $10M ARR" (Font: Montserrat Bold, Color: #FFFFFF, Size: 64pt, x: 820px, y: 260px)
- Pillars: "GTM Strategy  •  Demand Gen  •  Executive Advisory" (Font: Inter SemiBold, Color: #06B6D4, Size: 32pt, x: 820px, y: 380px)
- Website URL: "janedoe.com" (Font: Inter Medium, Color: #E2E8F0, Size: 28pt, x: 820px, y: 500px)
- Call-to-Action Pill: "Book Strategy Call" (Background: #06B6D4, Text: #0B0F19 Bold, Size: 24pt, x: 820px, y: 580px)
```
