# LinkedIn Company Page Banner Design — Corporate Branding & Conversion Workflow

> **Prompt Version:** 1.0.0
> **Target Environment:** Lovable, Figma, Canva, Image Generation AI (Flux, Ideogram, Midjourney) & Design AI Platforms
> **Synchronization:** Main Meta-Repo & Connected Workspaces

This prompt instructs design AI platforms on how to craft high-impact, conversion-oriented LinkedIn company page banners with exact ultra-wide dimensional standards, logo collision safe zones, and zero text hallucination.

---

## Strictly Avoid (Critical Negative Constraints)

> [!CAUTION]
> **ZERO TOLERANCE FOR LOGO COLLISION, TEXT HALLUCINATION & WEB CODE (AUTO-REJECT)**
> The AI MUST strictly avoid generating HTML, placing text in the company logo collision zone, or distorting typography.

1. **NO HTML Files or Web Code:** NEVER generate `.html` files, React components (`.tsx`, `.jsx`), or CSS stylesheets. This is strictly a corporate banner design and vector layout workflow.
2. **NO Text in the Company Logo Collision Zone:** NEVER place critical text, company taglines, or URLs in the bottom-left area (x: 0 to 260px). On desktop, the square company profile logo (300x300 px) overlaps the bottom-left of the banner, obscuring anything behind it.
3. **NO Unblended Photo Cutouts (MANDATORY BOTTOM & EDGE BLEND):** Any spokesperson photo, product cutout, or brand element MUST blend smoothly into the banner background using soft gradient feathering or dark vignettes. NEVER paste a cutout with hard, abrupt edges.
4. **NO Tacky Clip-Art, Star Badges, or Fake Icons (TOTAL BAN):** NEVER generate cartoonish star stickers (e.g. `#01` star badges), arbitrary geometric badges, or fake icons. Do NOT invent or create avatars or icons unless explicitly requested by the user. Professional corporate banners rely on clean typography, authentic branding, and subtle abstract graphics.
5. **NO Text Hallucinations or Gibberish:** NEVER render garbled glyphs, pseudo-Latin, or misspelled company names. Every character in the headline, value proposition, and domain URL must match the user's input verbatim.
6. **NO Low-Resolution Exports:** NEVER export or specify standard 72 DPI 1x images that blur upon upload. Always specify **2x or 3x Retina resolution** (`2256 x 382 px` or `3384 x 573 px`).
7. **NO Verbose or Cramped Text:** Due to the extremely narrow 5.91:1 aspect ratio (only 191px high at 1x), NEVER write long paragraphs. Keep headlines under 8 words and subheadings under 12 words.
8. **NO Uppercase Filenames:** NEVER use uppercase letters in folder names or file names (`linkedin-banners/`, `banner-2256x382.png`, `prompt.md` are required; `Banners/` is BANNED).

---

## 1. Dimensional Standards & Safe Zones

LinkedIn company banners use an ultra-wide panoramic aspect ratio (approx 5.91:1) that requires careful vertical and horizontal layout constraints:

### Canvas Dimensions

- **Standard Dimensions (1x):** `1128 x 191 px` (5.91:1 aspect ratio).
- **High-Resolution Retina (2x — Recommended):** `2256 x 382 px` (guarantees crystal-clear typography after upload compression).
- **Ultra High-Resolution (3x):** `3384 x 573 px`.
- **Target File Size & Format:** Under 8MB; export as uncompressed PNG or high-bitrate WebP.

### Safe Zone Architecture

```
+-----------------------------------------------------------------------------+
|                                    TOP SAFE BAND                            |
|                                                                             |
|   [ COMPANY LOGO     ]            CENTER / RIGHT CONTENT ZONE               |
|   [  COLLISION ZONE  ]     - Corporate Value Proposition / Headline         |
|   [  (Square Logo)   ]     - Supporting Tagline                             |
|   [                  ]     - Official Website URL & Call-to-Action          |
|                                                                             |
+-----------------------------------------------------------------------------+
|<--- 0 to 260px ------>|<----------------- 260px to 1128px ----------------->|
```

- **Dead Zone (Bottom-Left):** From `x: 0` to `x: 260px`, the square company profile logo overlaps the banner on desktop. Keep this area completely free of text and focal graphics.
- **Master Safe Zone:** Position all primary headlines, taglines, and callouts between `x: 300px` and `x: 1080px`, and between `y: 20px` and `y: 170px` (or at 2x: `x: 600px` to `x: 2160px`, `y: 40px` to `y: 340px`).

---

## 2. Input Capture & Verification

Before generating any design concepts or layout prompts, capture and validate the following inputs:

1. `company_name`: Exact legal or brand name (e.g., "Acme Cloud").
2. `value_proposition`: 1 concise, powerful headline explaining what the company does (e.g., "Automating Enterprise Security at Global Scale").
3. `core_offerings` / `industry`: Product categories or industry keywords (e.g., `Cloud Compliance • Zero-Trust Access • Real-Time Audit`).
4. `website_url`: Official company domain (e.g., `https://www.acmecloud.com`).
5. `email_address`: Contact or inquiries email address to embed on the banner (e.g., `contact@acmecloud.com` or `sales@acmecloud.com`).
6. `qr_code`: Optional scannable QR code destination URL or asset to embed (e.g., linking to a product demo, mobile app download, or newsletter signup).
7. `call_to_action`: Primary business action (e.g., "Explore the Platform", "Join Our Team", "Start Free 14-Day Trial").
8. `company_logo_or_avatar`: Description or vector icon of the brand mark, avatar, or spokesperson photo to place on the banner (optional — ask the user if they have one to provide, or skip if they prefer an abstract/typographic layout).
9. `brand_colors`: Corporate color palette (e.g., Deep Slate `#0f172a`, Electric Blue `#2563eb`, Pure White `#ffffff`).

> [!IMPORTANT]
> Always ask the user if they have a **company logo, brand avatar, or spokesperson photo** to include, or if they prefer to **skip** it (using a clean typographic or abstract layout). Remember: do NOT generate unsolicited avatars, clip-art icons, or star badges. Also ask if they want to embed an **email address** or a **scannable QR code** (e.g. for product demo, app download, or contact inquiries). If `company_name`, `value_proposition`, `website_url`, or `brand_colors` are missing or not provided, **STOP and ask the user** before proceeding.

---

## 3. Visual Layout Versions (Multiple Modes)

### Version 1: Value Proposition & Product Impact (Default)

Ideal for B2B tech, SaaS, and professional services:
- **Background:** Clean corporate gradient (slate navy to deep indigo) with subtle abstract tech wave or network mesh on the right.
- **Left Flank:** Negative space accommodating the square logo overlap.
- **Center-Left:** Bold, punchy headline in white and brand accent, followed by the supporting value proposition.
- **Right Flank:** High-contrast call-to-action pill and clean website URL.

### Version 2: Employer Branding / Talent Acquisition

Ideal for high-growth companies recruiting top talent:
- **Background:** Warm, authentic office environment with soft focus, or dynamic geometric gradient.
- **Center:** Bold headline: "We Are Hiring" or "Build the Future With Us".
- **Right Flank:** Highlighting open roles (e.g. "Engineering • Design • Sales") + `careers.acmecloud.com`.

### Version 3: Product UI / Tech Showcase

Ideal for developer platforms and software products:
- **Background:** Dark sleek aesthetic with ambient glows (`#0b0f19`).
- **Center-Left:** Problem-solution hook with clean typography.
- **Right Flank:** Angled isometric 3D product interface or dashboard preview showcasing the software UI.

---

## 4. Typography & Ultra-Wide Layout Constraints

- **Headline Brevity:** The canvas is only 191px high (at 1x). Headlines must be short, punchy, and limited to 1 or 2 lines maximum.
- **Font Hierarchy:**
  - *Main Headline:* Bold sans-serif (e.g., Montserrat Bold, Inter Bold, Poppins) at large scale.
  - *Tagline / Value Prop:* Clean medium sans-serif with generous letter spacing.
  - *URL & CTA:* High-contrast pill container with clean bold text for maximum legibility.
- **Verbatim Text Enforcement:** Company names, domain URLs, and metrics must be verified character-by-character to prevent AI spelling errors.

---

## 5. Directory & File Hierarchy (Strict Lowercase)

All generated LinkedIn company page banner assets must follow strict lowercase naming and zero-padded sequence numbers. The AI must persist the exact prompt used into `prompts/prompt.md` and save all high-resolution banner images inside `linkedin-banners/`:

```
/ (repo root)
└── 02-projects/
    ├── 01-{project-name}/
    │   ├── readme.md (project overview, banner preview, and layout specs)
    │   ├── prompts/
    │   │   ├── prompt.md (the exact prompt, user inputs, and AI parameters used)
    │   │   └── plan.md (design plan, corporate messaging, safe zones, and checklist)
    │   └── linkedin-banners/
    │       ├── banner-1128x191.png (Standard 1x company page banner)
    │       ├── banner-2256x382.png (High-Resolution 2x Retina banner — Recommended)
    │       ├── banner-3384x573.png (Ultra High-Resolution 3x banner)
    │       └── vector-overlay.svg (Crisp SVG vector typography, badges, and logo overlay)
    └── 02-{project-name}/
```

### Hierarchy Rules

1. **Root Projects Folder:** All projects live under `02-projects/`.
2. **Project Folder Naming:** `{sequence}-{project-name}` using two-digit zero-padding and kebab-case (e.g., `01-acme-cloud`, `02-enterprise-saas`).
3. **Prompt & Plan Preservation (Mandatory):** When formulating the company page banner design, every detail of the design plan (corporate value proposition, safe zone architecture, color palette, bottom gradient fade strategy, and verification checklist) MUST be saved directly to the file system at `prompts/plan.md`. The exact prompt given to the generation engine, user inputs, and model parameters MUST be saved in `prompts/prompt.md` so designs can be reproduced, audited, and iterated on.
4. **Asset Organization:** All banner raster images and vector overlays MUST be stored inside `linkedin-banners/`.
5. **Relative Paths:** All links and image embeds in `readme.md` must use relative paths (e.g., `![Banner](linkedin-banners/banner-2256x382.png)`).

---

## 6. Asset Specifications

### A. Company Page Banners (`linkedin-banners/`)

- `banner-1128x191.png`: Standard LinkedIn company page banner (5.91:1 aspect ratio, under 8MB).
- `banner-2256x382.png`: 2x Retina high-resolution banner (guarantees crystal-clear typography against compression).
- `banner-3384x573.png`: 3x Ultra high-resolution banner for maximum fidelity.
- `vector-overlay.svg`: Crisp SVG vector layer containing all typography, badges, URLs, and brand marks for hybrid compositing.

### B. Generation Prompt Archive (`prompts/prompt.md`)

- Contains the full generation prompt, model parameters (aspect ratio, style, negative prompts), and exact text strings used for the generation run.

---

## 7. Ready-to-Use Templates

### Template 1: Midjourney / Flux Background Generation Prompt

```text
Ultra-wide panoramic 6:1 banner background for LinkedIn company page, dimensions 2256x382. Sleek dark navy and deep blue corporate technology gradient with subtle abstract glowing network lines on the far right, smooth clean negative space on the left and center for corporate typography overlay, high-end enterprise B2B software aesthetic, minimal noise, 8k resolution --ar 6:1 --style raw
```

### Template 2: Figma / SVG Typography Overlay Specification

```markdown
# Canvas: 2256 x 382 px (2x Retina)
- Company Logo Collision Zone: x: 0 to 520 px (Keep empty)
- Primary Headline: "Automating Enterprise Security at Scale" (Font: Montserrat Bold, Color: #FFFFFF, Size: 52pt, x: 580px, y: 150px)
- Supporting Tagline: "Zero-Trust Infrastructure for Modern DevOps Teams" (Font: Inter Medium, Color: #94A3B8, Size: 26pt, x: 580px, y: 220px)
- Website URL & CTA: "acmecloud.com" (Font: Inter SemiBold, Color: #38BDF8, Size: 24pt, x: 580px, y: 280px)
- CTA Badge: "Start Free Trial" (Background: #2563EB, Text: #FFFFFF Bold, Size: 20pt, x: 1800px, y: 190px)
```
