# YouTube Thumbnail & Banner Design — Visual Identity & Typography Workflow

> **Prompt Version:** 1.0.0
> **Target Environment:** Lovable, Image Generation AI (Flux, Ideogram, Midjourney), Canva, Figma & Design AI Platforms
> **Synchronization:** Main Meta-Repo & Connected Workspaces

This prompt instructs design AI platforms on how to craft high-conversion, professional YouTube thumbnails and channel banners with strict text fidelity, zero hallucination, and a high-contrast visual hierarchy.

---

## Strictly Avoid (Critical Negative Constraints)

> [!CAUTION]
> **ZERO TOLERANCE FOR TEXT HALLUCINATION & SPELLING DISTORTION (AUTO-REJECT)**
> AI image models frequently distort, misspell, or fabricate text glyphs. Text accuracy is the absolute #1 priority.

1. **NO Text Hallucinations or Gibberish:** NEVER render garbled, pseudo-Latin, merged glyphs, or invented words. Every single character in titles, channel names, handles, and badges MUST match the user's exact spelling verbatim.
2. **NO Falsified Titles or Subtitles:** NEVER invent or embellish credentials, names, or quotes not explicitly provided or approved by the user.
3. **NO Unblended Subject Cutouts (MANDATORY BOTTOM GRADIENT FADE):** NEVER paste a subject cutout with hard, abrupt bottom or side edges (see Anti-Pattern reference). The subject MUST blend naturally into the bottom edge of the frame using a soft linear gradient fade or dark vignette so the torso/suit dissolves organically into the canvas.
4. **NO Tacky Clip-Art, Star Badges, or Fake Icons (TOTAL BAN):** NEVER add cheap clip-art shapes, cartoonish star stickers (e.g. `#01` star badges), arbitrary geometric badges, or fake icons. Do NOT invent or create avatars or icons unless explicitly requested by the user. Professional thumbnails rely on typography, authentic subject photography, and atmospheric lighting—NOT clip-art.
5. **NO Chaotic Overlapping Lines Across Subjects:** NEVER run background lines, graph strokes, or grid wires across the subject's face, neck, or body. Elements must stay behind the subject or maintain clean negative space.
6. **NO Uppercase Filenames:** NEVER use uppercase letters in folder names or file names (`youtube-thumbnails/`, `thumbnail-1280x720.png`, `prompt.md` are required; `Thumbnails/` is BANNED).
7. **NO Low-Contrast Text:** NEVER place light text over bright, busy backgrounds or dark text over dark shadows without proper contrast treatment (e.g., dropshadows, dark vignettes, or solid badge backdrops).
8. **NO Uncontrolled Facial or Hand Distortions:** The human subject in the foreground must have anatomically correct eyes, glasses, hands, and fingers without AI melting or extra digits.
9. **NO HTML or Web Page Code:** Do NOT generate full HTML web pages or application templates. This is strictly a graphic design, thumbnail, and banner workflow.
10. **NO Low-Resolution Exports:** NEVER export or specify standard 72 DPI blurry images. Always specify Full HD (`1920x1080 px` for thumbnails, `2560x1440 px` for banners).

---

## 1. Input Capture & Verification

Before generating any design concepts or prompts, capture and validate the following inputs:

1. `channel_name` / `host_name`: Exact name of the creator or brand (e.g., "Proloy Hasan").
2. `primary_title`: Main headline or topic (e.g., "Top 1% Podcast").
3. `subtitles_and_credentials`: Professional titles or roles separated by bars (e.g., `Author | Marketer | Trainer | Consultant | Podcaster`).
4. `core_pillars`: 3 key thematic words (e.g., `Ideas | Strategy | Impact`).
5. `tagline`: Supporting tagline (e.g., "Real Stories. Real People. Real Growth.").
6. `quote_or_callout`: Expressive quote or hook (e.g., `"Better Ideas Bigger Impact"`, `Let's Build Something Great`).
7. `youtube_handle`: Exact channel handle (e.g., `/top1percentPodcast`).
8. `youtube_url`: Full channel URL (e.g., `https://youtube.com/@top1percentPodcast`).
9. `website_url`: Creator or brand website URL (e.g., `https://proloyhasan.com`).
10. `email_address`: Contact or business inquiry email address to embed (e.g., `contact@proloyhasan.com`).
11. `qr_code`: Optional scannable QR code destination URL or asset to embed (e.g., linking to newsletter, booking calendar, or channel subscribe link).
12. `channel_icon_or_avatar`: Profile picture, logo mark, or avatar image (optional — ask the user if they have one to provide, or skip if none).
13. `person_photo_or_avatar`: Photo of the creator/host, visual description, or avatar cutout (optional — ask the user if they have an image/photo of the person to include, or skip if they prefer a text/graphics-only design).
14. `color_palette`: Primary brand colors (e.g., Deep Navy Background `#0f172a`, Golden Yellow Highlight `#facc15`, Crisp White `#ffffff`).

> [!IMPORTANT]
> Always ask the user if they have a **person's photo, subject cutout, or avatar image** to include, or if they prefer to **skip** it (for a graphics/typography-focused layout). Remember: do NOT generate unsolicited avatars, clip-art icons, or star badges. Also ask if they want to embed an **email address** or a **scannable QR code**. If `channel_name`, `primary_title`, `youtube_handle`, or `website_url` are missing or not provided, **STOP and ask the user** before proceeding.

---

## 2. Deep Compositional & Visual Breakdown (Reference Model)

The reference banner layout below demonstrates the target standard for high-authority YouTube channel banners and podcast thumbnails.

### Visual Reference (Inline Compressed Thumbnail)

<img src="data:image/jpeg;base64,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" width="400" alt="Reference Layout" />

### Anti-Pattern Reference (What NOT to Do — Unapproved Design)

The thumbnail below illustrates common design failures that MUST be strictly avoided:

<img src="data:image/jpeg;base64,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" width="400" alt="Anti-Pattern Layout - Strictly Avoid" />

#### Why This Output is Rejected (Anti-Pattern Breakdown):

1. **Harsh, Unblended Cutout Edge:** The subject's suit and torso cut off with hard, abrupt edges at the bottom of the canvas. It completely lacks the mandatory soft gradient fade / dark feathered vignette that allows the subject to naturally melt into the frame.
2. **Tacky Star Sticker Badge:** The orange `#01` star badge in the top-right corner looks like cheap clip-art. Professional thumbnails never use cartoon star stickers or arbitrary badges.
3. **Chaotic Lines Slicing Through the Subject:** The jagged orange and cyan line graph cuts directly across the subject's neck and chest, ruining contrast and violating the anti-overlapping rule.
4. **Unsolicited Icon/Avatar Generation:** The AI generated arbitrary star icons and decorative lines that were not requested by the user.
5. **Thematic Disconnection:** A formal corporate suit cutout placed against an aggressive gaming title ("DOOM ETERNAL") with no atmospheric integration, color grading, or environmental lighting harmony.

### Layout Anatomy Breakdown

1. **Center Foreground (Hero Subject / Person's Photo or Avatar):**
   - **When Provided (Mandatory Bottom Fade):** Professional subject cutout positioned centrally or slightly left-of-center with sharp contrast against the background, crisp lighting on face/shoulders, and natural posture. **CRITICAL:** The bottom edge of the subject MUST have a soft gradient fade (feathered mask / dark bottom vignette) so the torso/suit naturally dissolves into the lower frame without sharp, abrupt edges.
   - **When Skipped:** Adapt the composition to a graphics, typography, or iconography-focused layout; center the primary headline and name, or place a high-contrast 3D icon/badge or thematic illustration in place of the human subject.
   - **NO Unsolicited Avatars/Icons:** Do NOT generate cartoon avatars or arbitrary icon badges unless explicitly provided or requested.

2. **Center-Right (Hero Typography & Name):**
   - **Primary Name:** Giant high-contrast lettering:
     - First Name (`Proloy`): Bold, clean, modern sans-serif in **pure white** (`#ffffff`).
     - Last Name (`Hasan`): Bold, clean, modern sans-serif in **electric/golden yellow** (`#facc15`).
   - **Professional Roles:** Italicized/serif subtitles separated by vertical bars:
     - `Author | Marketer | Trainer | Consultant | Podcaster`
   - **Pillar Keywords:** Underlined thematic foundation:
     - `Ideas | Strategy | Impact` with a subtle golden accent line.

3. **Right Flank (Podcast & Channel Branding):**
   - Vertical dividing rule providing visual separation.
   - **Podcast Identity:** Stylized microphone icon + bold white `Top 1%` text.
   - **Pill Badge:** Bold black text inside a solid rounded **yellow pill badge** (`Podcast`).
   - **Tagline:** Clean secondary text: `Real Stories. Real People. Real Growth.`
   - **Social Proof / Handle:** Red YouTube play button icon followed by `/top1percentPodcast`.
   - **Atmospheric Elements:** Studio condenser microphone in the foreground; dusk/night cityscape in the background with warm bokeh lights.

4. **Left Flank (Authority & Social Proof):**
   - Display of physical books with legible cover titles resting on a warm wood surface.
   - Inspiring headline quote above the books: `"Better Ideas Bigger Impact"` in expressive white script with a golden yellow underline.
   - Inset black-and-white documentary photo showing the creator live on stage speaking to an audience.

5. **Lower Third (Credential Badges):**
   - 4 minimalist white line-art icons paired with clean labels:
     - `[Open Book Icon]` — Bestselling Author
     - `[Graduation Cap Icon]` — Hard-skill Trainer
     - `[Megaphone Icon]` — Brand Marketing Professional
     - `[Briefcase Icon]` — Business Consultant
   - **Closing Call to Action (Far Right):** Dynamic handwritten script in golden yellow: `Let's Build Something Great` with a stylized underline.

---

## 3. Typography & Contrast Rules

- **The Golden Yellow Accent Rule:** Use vibrant golden yellow (`#facc15` or `#eab308`) exclusively for key focal words (e.g. surname, badge containers, underline strokes, call-to-action scripts). This draws the eye immediately.
- **Dark Textured Backdrop:** The background behind the text must be dark navy, deep charcoal, or cinematic dusk (`#0b0f19` to `#1e293b`) with subtle depth of field/bokeh to make white and yellow typography pop without visual interference.
- **Font Pairing Harmony:**
  - *Headline & Names:* Heavy geometric sans-serif (e.g., Montserrat Bold, Poppins, Inter Black).
  - *Subtitles & Roles:* Clean medium sans-serif or refined editorial serif.
  - *Quotes & Callouts:* Organic, authentic brush/script font (e.g., Caveat, Permanent Marker, or clean calligraphy).
- **Badge Containers:** When text sits over complex areas, enclose it inside a solid high-contrast rounded pill (e.g., yellow pill with black text).

---

## 4. Text Accuracy & Anti-Hallucination Protocol

When generating prompts for image engines (Flux, Midjourney v6, Ideogram) or compositing layers:

1. **Explicit Text Quoting:** Always specify text inside literal quotes in the generation prompt:
   - `with the exact text "Proloy" in bold white sans-serif letters, and "Hasan" in bold golden-yellow sans-serif letters`
2. **Character Verification Gate:** Inspect the generated output. If even a single character is warped, merged, or misspelled (e.g. "Hassan" instead of "Hasan", or garbled Bengali book titles), the image MUST be rejected or the text layer must be re-rendered as a clean vector overlay.
3. **Hybrid Compositing (Recommended):** For mission-critical production banners:
   - Use AI to generate the background, lighting, and subject cutout.
   - Render all typography, badges, and logos as crisp SVG or vector layers over the background to guarantee 100% spelling precision.

---

## 5. Directory & File Hierarchy (Strict Lowercase)

All generated YouTube thumbnail, banner, and prompt assets must follow strict lowercase naming and zero-padded sequence numbers. The AI must persist the exact prompt used into `prompts/prompt.md` and save all high-resolution thumbnail and banner images inside `youtube-thumbnails/`:

```
/ (repo root)
└── 02-projects/
    ├── 01-{project-name}/
    │   ├── readme.md (project overview, visual previews, and layout specs)
    │   ├── prompts/
    │   │   └── prompt.md (the exact prompt, user inputs, and AI parameters used)
    │   └── youtube-thumbnails/
    │       ├── thumbnail-1280x720.png (Standard 16:9 YouTube video thumbnail)
    │       ├── thumbnail-1920x1080.png (Full HD 1080p high-resolution thumbnail)
    │       ├── banner-2560x1440.png (Full YouTube channel banner / TV master)
    │       ├── banner-safe-zone-1546x423.png (Desktop & mobile safe crop banner)
    │       └── vector-overlay.svg (Crisp SVG vector typography, badges, and logo overlay)
    └── 02-{project-name}/
```

### Hierarchy Rules

1. **Root Projects Folder:** All projects live under `02-projects/`.
2. **Project Folder Naming:** `{sequence}-{project-name}` using two-digit zero-padding and kebab-case (e.g., `01-tech-podcast`, `02-coding-insights`).
3. **Prompt Preservation:** The exact prompt given to the generation engine, user inputs, and model parameters MUST be saved in `prompts/prompt.md` so designs can be reproduced, audited, and iterated on.
4. **Asset Organization:** All thumbnail and banner raster images MUST be stored inside `youtube-thumbnails/`.
5. **Relative Paths:** All links and image embeds in `readme.md` must use relative paths (e.g., `![Thumbnail](youtube-thumbnails/thumbnail-1280x720.png)`).

---

## 6. Asset Specifications

### A. YouTube Thumbnails (`youtube-thumbnails/`)

- `thumbnail-1280x720.png`: Standard YouTube video thumbnail (16:9 aspect ratio, under 2MB).
- `thumbnail-1920x1080.png`: Full HD high-resolution thumbnail for pristine visual quality on high-DPI displays.
- `banner-2560x1440.png`: Full YouTube channel banner (TV master dimension).
- `banner-safe-zone-1546x423.png`: Centered safe zone banner crop ensuring logos and text are fully visible on desktop and mobile.
- `vector-overlay.svg`: Crisp SVG vector layer containing all typography, badges, URLs, and icons for hybrid compositing.

### B. Generation Prompt Archive (`prompts/prompt.md`)

- Contains the full generation prompt, model parameters (aspect ratio, style, negative prompts), and exact text strings used for the generation run.

---

## 7. Ready-to-Use Prompt Templates

### Template 1: Midjourney / Flux Photorealistic Generation Prompt

```text
Wide cinematic 16:9 YouTube banner layout. In the center foreground, a professional South Asian male host wearing black-rimmed glasses, neat beard, maroon overshirt over a white t-shirt, seated with hands clasped, warm studio lighting. Dark textured office and bookshelf background on the left with warm amber lamps; dark architectural cityscape at dusk on the right with bokeh lighting and a studio microphone on a boom arm. Crisp negative space for typography. Deep navy, charcoal, and warm amber color grading, 8k resolution, photorealistic commercial photography --ar 16:9 --style raw
```

### Template 2: Typography & Vector Overlay Specification (Figma / SVG)

```markdown
# Typography Specification
- Primary Heading: "Proloy" (Font: Montserrat Bold, Color: #FFFFFF, Size: 72pt)
- Secondary Heading: "Hasan" (Font: Montserrat Bold, Color: #FACC15, Size: 72pt)
- Roles: "Author | Marketer | Trainer | Consultant | Podcaster" (Font: Inter Medium Italic, Color: #E2E8F0, Size: 22pt)
- Pillars: "Ideas | Strategy | Impact" (Font: Inter SemiBold, Color: #94A3B8, Underline: #FACC15, Size: 18pt)
- Podcast Badge: "Top 1%" (#FFFFFF, Bold) + "Podcast" (#0F172A, Bold, Background: #FACC15 pill)
- Quote: ""Better Ideas Bigger Impact"" (Font: Caveat / Handwritten Script, Color: #FFFFFF, Underline: #FACC15, Size: 36pt)
- Bottom Badges: 4 icons (Book, Cap, Megaphone, Briefcase) with white text #FFFFFF, 14pt
- Call to Action: "Let's Build Something Great" (Font: Handwritten Script, Color: #FACC15, Size: 28pt)
```
