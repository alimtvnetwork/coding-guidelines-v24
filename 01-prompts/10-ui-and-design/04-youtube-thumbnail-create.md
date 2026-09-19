# YouTube Thumbnail & Banner Design — Visual Identity & Typography Workflow

> **Prompt Version:** 1.1.0
> **Target Environment:** Lovable, Image Generation AI (Flux, Ideogram, Midjourney), Canva, Figma & Design AI Platforms
> **Synchronization:** Main Meta-Repo & Connected Workspaces

This prompt instructs design AI platforms on how to craft high-conversion, professional YouTube thumbnails and channel banners with strict text fidelity, zero hallucination, cinematic photographic shot direction, and a high-contrast visual hierarchy.

---

## Strictly Avoid (Critical Negative Constraints)

> [!CAUTION]
> **ZERO TOLERANCE FOR TEXT HALLUCINATION, UNBLENDED EDGES, CLIP-ART & UNSEQUENCED FILES (AUTO-REJECT)**
> AI image models frequently produce waxy skin, harsh cutout edges, tacky clip-art stickers, and scrambled text. Accuracy and photographic realism are the absolute top priorities.

1. **NO Text Hallucinations or Gibberish:** NEVER render garbled, pseudo-Latin, merged glyphs, or invented words. Every single character in titles, channel names, handles, and badges MUST match the user's exact spelling verbatim.
2. **NO Falsified Titles or Subtitles:** NEVER invent or embellish credentials, names, or quotes not explicitly provided or approved by the user.
3. **NO Unblended Subject Cutouts (MANDATORY BOTTOM GRADIENT FADE):** NEVER paste a subject cutout with hard, abrupt bottom or side edges. The subject MUST blend naturally into the bottom edge of the frame using a soft linear gradient fade or dark vignette so the torso/suit dissolves organically into the canvas.
4. **NO Tacky Clip-Art, Star Badges, or Fake Icons (TOTAL BAN):** NEVER add cheap clip-art shapes, cartoonish star stickers (e.g. `#01` star badges), arbitrary geometric badges, or fake icons. Do NOT invent or create avatars or icons unless explicitly requested by the user. Professional thumbnails rely on typography, authentic subject photography, and atmospheric lighting—NOT clip-art.
5. **NO Chaotic Overlapping Lines Across Subjects:** NEVER run background lines, graph strokes, or grid wires across the subject's face, neck, or body. Elements must stay behind the subject or maintain clean negative space.
6. **NO Unsequenced or Uppercase Filenames:** ALL generated files and images MUST use lowercase kebab-case preceded by a two-digit zero-padded sequence number (e.g. `01-thumbnail-1280x720.png`, `02-thumbnail-1920x1080.png`, `01-prompt.md`, `02-plan.md`). Unsequenced names like `thumbnail.png` are strictly BANNED.
7. **NO Low-Contrast Text:** NEVER place light text over bright, busy backgrounds or dark text over dark shadows without proper contrast treatment (e.g., dropshadows, dark vignettes, or solid badge backdrops).
8. **NO Empty Background Placeholders:** NEVER generate an empty background containing blank colored boxes or disconnected shapes while putting all text in an SVG with broken images. The composition must be fully realized, cohesive, and intentional.
9. **NO Uncontrolled Facial or Hand Distortions:** The human subject in the foreground must have anatomically correct eyes, glasses, hands, and fingers without AI melting or extra digits.
10. **NO HTML or Web Page Code:** Do NOT generate full HTML web pages or application templates. This is strictly a graphic design, thumbnail, and banner workflow.
11. **NO Low-Resolution Exports:** NEVER export or specify standard 72 DPI blurry images. Always specify Full HD (`1920x1080 px` for thumbnails, `2560x1440 px` for banners).

---

## 1. Input Capture & Verification

Before generating any design concepts or prompts, capture and validate the following inputs:

1. `channel_name` / `host_name`: Exact name of the creator or brand (e.g., "Tech Horizon").
2. `primary_title`: Main headline or topic (e.g., "Full Stack Mastery").
3. `subtitles_and_credentials`: Professional titles or roles separated by bars (e.g., `Software Architect | Cloud Engineer | Podcaster`).
4. `core_pillars`: 3 key thematic words (e.g., `Architecture | Scale | Security`).
5. `tagline`: Supporting tagline (e.g., "Real Engineering. Real Systems.").
6. `quote_or_callout`: Expressive quote or hook (e.g., `"Build Systems That Scale"`, `Subscribe for Weekly Deep Dives`).
7. `youtube_handle`: Exact channel handle (e.g., `@techhorizon`).
8. `youtube_url`: Full channel URL (e.g., `https://youtube.com/@techhorizon`).
9. `website_url`: Creator or brand website URL (e.g., `https://techhorizon.dev`).
10. `email_address`: Contact or business inquiry email address to embed (e.g., `contact@techhorizon.dev`).
11. `qr_code`: Optional scannable QR code destination URL or asset to embed (e.g., linking to newsletter, booking calendar, or channel subscribe link).
12. `channel_icon_or_avatar`: Profile picture, logo mark, or avatar image (optional — ask the user if they have one to provide, or skip if none).
13. `person_photo_or_avatar`: Photo of the creator/host, visual description, or avatar cutout (optional — ask the user if they have an image/photo of the person to include, or skip if they prefer a text/graphics-only design).
14. `font_family`: Primary font family (Defaults to **Ubuntu** across all typography, titles, and overlays).
15. `color_palette`: Primary brand colors (e.g., Deep Charcoal `#0f172a`, Electric Cyan `#06b6d4`, Golden Amber `#f59e0b`, Crisp White `#ffffff`).

> [!IMPORTANT]
> Always ask the user if they have a **person's photo, subject cutout, or avatar image** to include, or if they prefer to **skip** it (for a graphics/typography-focused layout). Remember: do NOT generate unsolicited avatars, clip-art icons, or star badges. Also ask if they want to embed an **email address** or a **scannable QR code**. If `channel_name`, `primary_title`, `youtube_handle`, or `website_url` are missing or not provided, **STOP and ask the user** before proceeding.

---

## 2. Photographic, Lighting & Compositional Specification

Professional YouTube thumbnails require deliberate cinematic direction rather than generic AI generation:

### A. Camera Shot & Framing

- **Lens & Optics:** 85mm prime portrait lens equivalent with a wide aperture (`f/1.8` to `f/2.8`), creating an authentic shallow depth of field where the subject remains razor-sharp while the background dissolves into soft, creamy bokeh.
- **Shot Distance:** Eye-level medium close-up or bust shot (from mid-chest to top of head). Never extreme wide-angle or fisheye distortion that distorts facial proportions.
- **Rule of Thirds Placement:** Position the subject anchored on the right third (or left third) of the 16:9 canvas. This leaves two uninterrupted thirds of clean, dark negative space for high-impact typography and branding elements.
- **Gaze & Expression:** Direct, engaging eye contact with the camera lens. Natural, authentic facial expression (calm authority, intense focus, or expressive reaction matching the video topic).

### B. 3-Point Cinematic Lighting

- **Key Light:** Large, diffused softbox positioned at 45 degrees to the subject, creating gentle, natural facial modeling without harsh nose or chin shadows.
- **Rim / Edge Light (Crucial):** Distinctive, colored rim light matching the primary brand accent (e.g. electric cyan, warm amber, or neon red) striking the subject's hair, shoulders, and jawline from behind. This creates strong edge separation from dark backgrounds.
- **Fill Light:** Subtle, low-intensity ambient fill preserving rich textural details in shadows without flattening the image.

### C. Texture & Environmental Realism

- **Organic Skin Texture:** Visible pores, natural micro-contrast, realistic skin tones, and authentic specular highlights on the forehead and cheekbones. Total ban on plastic, waxy, or airbrushed AI skin.
- **Crisp Textile Weave:** High-fidelity fabric texture on clothing (woven cotton, jacket lapels, collar stitches).
- **Catchlights & Reflections:** Crisp, natural catchlights in the pupils of the eyes and subtle realistic reflections on glasses if worn.
- **Atmospheric Background:** Dark textured studio interior, bokeh city lights at dusk, or ambient tech mesh. The background must feel rich and spatial, not a flat solid color or an empty black box.

### D. Mandatory Bottom & Edge Gradient Fade

- **Natural Blending:** The subject's torso/suit must NEVER be abruptly sliced at the bottom edge of the thumbnail.
- **Gradient Dissolve:** Apply a soft vertical linear gradient fade (or a dark feathered vignette) across the bottom 20% of the subject cutout, allowing the clothing to melt organically into the dark lower border of the canvas.

### E. Anti-Pattern Breakdown (What NOT to Do)

Avoid these five critical failures that ruin thumbnail quality:
1. **Harsh, Unblended Cutout Bottom:** Abruptly cutting off the subject's suit with a razor-sharp pixel edge at the bottom of the frame.
2. **Tacky Star Sticker Badges:** Placing cartoonish star shapes (e.g. `#01` star stickers), ribbon badges, or clip-art arrows in the corners.
3. **Chaotic Lines Slicing Through Subjects:** Running line graphs, wireframes, or grid lines across the subject's chest, neck, or face.
4. **Unsolicited Avatar/Icon Generation:** Inventing cartoon avatars, arbitrary badges, or random decorative shapes not requested by the user.
5. **Thematic Disconnection:** Placing a formal business-suit cutout against an unrelated gaming or aggressive theme without atmospheric lighting integration.

---

## 3. Typography & Contrast Rules

- **Primary Font Standard (Ubuntu):** All typography MUST default to the **Ubuntu** font family:
  - *Main Headline:* `Ubuntu Bold` or `Ubuntu Bold Condensed` for maximum visual weight and readability at small mobile thumbnail sizes.
  - *Subtitles & Roles:* `Ubuntu Medium` or `Ubuntu Regular`.
  - *Pill Badges & Tags:* `Ubuntu Bold` inside high-contrast rounded containers.
  - *Quotes & Callouts:* `Ubuntu Medium Italic` or authentic clean script.
- **The Accent Color Rule:** Use the vibrant accent color (e.g., Golden Amber `#f59e0b` or Electric Cyan `#06b6d4`) exclusively for 1–2 key focal words. This directs the viewer's eye in under 0.5 seconds.
- **Dark Textured Backdrop:** The background behind typography must be dark navy, deep charcoal, or cinematic dusk (`#0b0f19` to `#1e293b`) with subtle depth of field to guarantee WCAG AAA contrast.
- **Pill Badge Containers:** When text sits over complex textures, enclose it inside a solid high-contrast rounded pill (e.g., amber pill with dark text).

---

## 4. Text Accuracy & Anti-Hallucination Protocol

When generating prompts for image engines (Flux, Midjourney v6, Ideogram) or compositing layers:

1. **Explicit Text Quoting:** Always specify text inside literal quotes in the generation prompt:
   - `with the exact text "TECH" in bold white Ubuntu letters, and "HORIZON" in bold electric-cyan Ubuntu letters`
2. **Character Verification Gate:** Inspect the generated output. If even a single character is warped, merged, or misspelled, the image MUST be rejected or the text layer must be re-rendered as a clean vector overlay.
3. **Hybrid Compositing (Recommended):** For production-grade thumbnails:
   - Use AI to generate the photographic background, cinematic lighting, and subject portrait.
   - Render all typography, badges, and logos as crisp SVG vector layers over the background to guarantee 100% spelling precision.

---

## 5. Directory & File Hierarchy (Strict Lowercase & Two-Digit Sequence)

All generated YouTube thumbnail, banner, and prompt assets must follow strict lowercase naming and two-digit zero-padded sequence numbers. The AI must persist the design plan into `prompts/02-plan.md`, the generation prompt into `prompts/01-prompt.md`, and save all high-resolution images inside `youtube-thumbnails/`:

```
/ (repo root)
└── 02-projects/
    ├── 01-{project-name}/
    │   ├── readme.md (project overview, visual previews, and layout specs)
    │   ├── prompts/
    │   │   ├── 01-prompt.md (the exact prompt, user inputs, and AI parameters used)
    │   │   └── 02-plan.md (design plan, composition strategy, safe zones, and checklist)
    │   └── youtube-thumbnails/
    │       ├── 01-thumbnail-1280x720.png (Standard 16:9 YouTube video thumbnail)
    │       ├── 02-thumbnail-1920x1080.png (Full HD 1080p high-resolution thumbnail)
    │       ├── 03-banner-2560x1440.png (Full YouTube channel banner / TV master)
    │       ├── 04-banner-safe-zone-1546x423.png (Desktop & mobile safe crop banner)
    │       └── 05-vector-overlay.svg (Crisp SVG vector typography, badges, and logo overlay)
    └── 02-{project-name}/
```

### Hierarchy Rules

1. **Root Projects Folder:** All projects live under `02-projects/`.
2. **Project Folder Naming:** `{sequence}-{project-name}` using two-digit zero-padding and kebab-case (e.g., `01-tech-podcast`, `02-coding-insights`).
3. **Prompt & Plan Preservation (Mandatory):** When formulating the thumbnail/banner design, every detail of the design plan (visual composition, color grading, typography pairing, bottom gradient fade strategy, safe zone mapping, and anti-pattern checklist) MUST be saved directly to the file system at `prompts/02-plan.md`. The exact prompt given to the generation engine, user inputs, and model parameters MUST be saved in `prompts/01-prompt.md` so designs can be reproduced, audited, and iterated on.
4. **Asset Organization:** All thumbnail and banner raster images MUST be stored inside `youtube-thumbnails/` with zero-padded sequence prefixes (`01-`, `02-`, etc.).
5. **Relative Paths:** All links and image embeds in `readme.md` must use relative paths (e.g., `![Thumbnail](youtube-thumbnails/01-thumbnail-1280x720.png)`).

---

## 6. Asset Specifications

### A. YouTube Thumbnails (`youtube-thumbnails/`)

- `01-thumbnail-1280x720.png`: Standard YouTube video thumbnail (16:9 aspect ratio, under 2MB).
- `02-thumbnail-1920x1080.png`: Full HD high-resolution thumbnail for pristine visual quality on high-DPI displays.
- `03-banner-2560x1440.png`: Full YouTube channel banner (TV master dimension).
- `04-banner-safe-zone-1546x423.png`: Centered safe zone banner crop ensuring logos and text are fully visible on desktop and mobile.
- `05-vector-overlay.svg`: Crisp SVG vector layer containing all Ubuntu typography, badges, URLs, and icons for hybrid compositing.

### B. Generation Prompt & Plan Archive (`prompts/`)

- `01-prompt.md`: Contains the full generation prompt, model parameters (aspect ratio, style, negative prompts), and exact text strings used for the generation run.
- `02-plan.md`: Contains the full design plan, composition strategy, color palette, typography specification, safe zone mapping, and anti-pattern quality verification checklist.

---

## 7. Ready-to-Use Prompt Templates

### Template 1: Midjourney / Flux Photorealistic Generation Prompt

```text
Cinematic 16:9 YouTube thumbnail portrait. On the right third of the frame, a professional subject captured in a medium close-up with an 85mm portrait lens at f/1.8, razor-sharp focus on face and eyes, natural skin texture with visible pores, authentic expression, modern dark jacket. Dramatic 3-point studio lighting with a soft diffused key light and vibrant electric cyan rim light tracing the hair and shoulders. The lower torso dissolves seamlessly into a soft dark gradient fade at the bottom edge. Dark atmospheric studio background with deep charcoal tones and subtle warm amber bokeh lights on the left two-thirds. Ample clean negative space for typography overlay on the left. High commercial photography quality, 8k resolution --ar 16:9 --style raw
```

### Template 2: Typography & Vector Overlay Specification (Figma / SVG)

```markdown
# Typography Specification (Ubuntu Font Family)
- Primary Heading: "{first_title}" (Font: Ubuntu Bold, Color: #FFFFFF, Size: 72pt)
- Secondary Accent Heading: "{second_title}" (Font: Ubuntu Bold, Color: #06B6D4, Size: 72pt)
- Subtitle / Roles: "{subtitles_and_roles}" (Font: Ubuntu Medium, Color: #E2E8F0, Size: 22pt)
- Pillars: "{core_pillars}" (Font: Ubuntu Medium, Color: #94A3B8, Underline: #06B6D4, Size: 18pt)
- Badge Container: "EPISODE" (#FFFFFF, Bold) + "01" (#0F172A, Bold, Background: #06B6D4 pill)
- Contact / URL: "{website_url}" (Font: Ubuntu Regular, Color: #E2E8F0, Size: 18pt)
```
