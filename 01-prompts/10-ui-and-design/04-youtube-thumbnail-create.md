# YouTube Thumbnail & Banner Design — Visual Identity & Typography Workflow

> **Prompt Version:** 2.0.0
> **Target Environment:** Lovable, Image Generation AI (Flux, Ideogram, Midjourney), Canva, Figma & Design AI Platforms
> **Synchronization:** Main Meta-Repo & Connected Workspaces

This prompt instructs design AI platforms on how to craft high-conversion, professional YouTube thumbnails and channel banners with strict text fidelity, zero hallucination, cinematic photographic shot direction, a 5-zone layout grid, and a high-contrast visual hierarchy.

---

## Strictly Avoid (Critical Negative Constraints)

> [!CAUTION]
> **ZERO TOLERANCE FOR TEXT HALLUCINATION, UNBLENDED EDGES, CLIP-ART & UNSEQUENCED FILES (AUTO-REJECT)**
> AI image models frequently produce waxy skin, harsh cutout edges, tacky clip-art stickers, and scrambled text. Accuracy, photographic realism, and spatial balance are the absolute top priorities.

1. **NO Text Hallucinations or Gibberish:** NEVER render garbled, pseudo-Latin, merged glyphs, or invented words. Every single character in titles, channel names, handles, and badges MUST match the user's exact spelling verbatim.
2. **NO Falsified Titles or Subtitles:** NEVER invent or embellish credentials, names, awards, statistics, testimonials, or quotes not explicitly provided or approved by the user.
3. **NO Unblended Subject Cutouts (MANDATORY BOTTOM & EDGE GRADIENT FADE):** NEVER paste a subject cutout with hard, abrupt bottom or side edges. The subject MUST blend naturally into the bottom edge of the frame using a soft linear gradient fade or dark feathered vignette so the torso/suit dissolves organically into the canvas.
4. **NO Tacky Clip-Art, Star Badges, or Fake Icons (TOTAL BAN):** NEVER add cheap clip-art shapes, cartoonish star stickers (e.g. `#01` star badges), arbitrary geometric badges, or fake icons. Do NOT invent or create avatars or icons unless explicitly requested by the user. Professional thumbnails rely on typography, authentic subject photography, and atmospheric lighting—NOT clip-art.
5. **NO Chaotic Overlapping Lines Across Subjects:** NEVER run background lines, graph strokes, or grid wires across the subject's face, neck, or body. Elements must stay behind the subject or maintain clean negative space.
6. **NO Unsequenced or Uppercase Filenames:** ALL generated files and images MUST use lowercase kebab-case preceded by a two-digit zero-padded sequence number (e.g. `01-thumbnail-1280x720.png`, `02-thumbnail-1920x1080.png`, `01-prompt.md`, `02-plan.md`). Unsequenced names like `thumbnail.png` are strictly BANNED.
7. **NO Low-Contrast Text:** NEVER place light text over bright, busy backgrounds or dark text over dark shadows without proper contrast treatment (e.g., dropshadows, dark vignettes, or solid badge backdrops).
8. **NO Empty Background Placeholders:** NEVER generate an empty background containing blank colored boxes or disconnected shapes while putting all text in an SVG with broken images. The composition must be fully realized, cohesive, and intentional.
9. **NO Uncontrolled Facial or Hand Distortions:** The human subject in the foreground must have anatomically correct eyes, glasses, hands, and fingers without AI melting or extra digits.
10. **NO HTML or Web Page Code:** Do NOT generate full HTML web pages or application templates. This is strictly a graphic design, thumbnail, and banner workflow.
11. **NO Low-Resolution Exports:** NEVER export or specify standard 72 DPI blurry images. Always specify Full HD (`1280x720 px` standard, `1920x1080 px` high-res for thumbnails, `2560x1440 px` for banners).
12. **NO Distorted Lighting or Neon Gradients:** Do not use purple, violet, cyan, fluorescent yellow, pastel colors, rainbow effects, or harsh saturated neon gradients unless explicitly requested.

---

## 1. Canvas Architecture & 5-Zone Layout Grid

All thumbnails are engineered on a `1280 × 720 px` canvas (16:9 aspect ratio) with strict spatial zone boundaries:

### Canvas Specifications

- **Dimensions:** `1280 × 720 pixels` (16:9 aspect ratio).
- **Edge Safe Area:** Keep all important faces, text, and focal graphics at least `55 pixels` away from every outside edge.
- **Central Core:** Primary content must remain inside the central 80% of the canvas.
- **Continuous Scene Blending:** Do NOT make sections look like separate rectangular floating cards. Blend them together as one continuous cinematic environment.

### 5-Zone Spatial Breakdown

```
+---------------------------------------------------------------------------------------------------+
|  [ ZONE A: LEFT CONTENT ]   [ ZONE B: PORTRAIT ]   [ ZONE C: NAME & CRED ]  [ ZONE D: ACHIEVEMENT]|
|  x: 25 to 405 px            x: 355 to 735 px       x: 655 to 965 px         x: 930 to 1245 px     |
|  - Handwritten Quote        - Seated Person        - Main Name (Stacked)    - "Top 1%" Badge      |
|  - Book Display / Table     - Natural Overlap      - Professional Titles    - Supporting Statement|
|  - Warm Studio Lights       - 54-58% Canvas Height - Values Line + Underline- Mic / YouTube Line  |
|                                                                                                   |
|                             [ ZONE E: LOWER INFORMATION STRIP ]                                   |
|                             x: 670 to 1280 px | y: 548 to 720 px                                  |
|                             - 4 Professional Role Groups behind lower body                        |
+---------------------------------------------------------------------------------------------------+
```

- **Zone A (Left Content Zone, x: 25 to 405 px, width ~380px):** Contains the handwritten quote, physical book display on a wood table, and warm studio lighting.
- **Zone B (Portrait Zone, x: 355 to 735 px, width ~380px):** Contains the seated featured subject. Overlaps the left and center zones slightly to connect the composition organically.
- **Zone C (Name & Credential Zone, x: 655 to 965 px, width ~310px):** Contains the stacked prominent name, professional credentials with vertical bar separators, and the values line with a hand-painted gold underline.
- **Zone D (Achievement Zone, x: 930 to 1245 px, width ~315px):** Contains the "Top 1%" achievement callout, supporting statement, YouTube identity line, and the podcast microphone entering from the right edge.
- **Zone E (Lower Information Strip, y: 548 to 720 px, x: 670 to 1280 px):** A dark translucent strip containing four professional role groups. It appears behind the subject's lower body where the two areas overlap.

---

## 2. Complete Color System & Distribution Formula

Apply a disciplined, luxury-authority color hierarchy across every element:

### Color Palette

- **Main Background Tones:**
  - Deep midnight navy: `#07111F`
  - Dark blue-black: `#0A1424`
  - Charcoal black: `#15171C`
  - Warm near-black brown: `#211612`
- **Primary Text Tones:**
  - Main white: `#F7F7F4`
  - Soft white: `#ECEBE6`
  - Secondary light gray: `#B8BBC2`
  - Muted gray: `#858B94`
- **Gold & Yellow Accents:**
  - Primary warm gold: `#F5A817`
  - Bright highlight gold: `#FFB51B`
  - Dark gold shadow: `#B96E08`
  - Hand-drawn underline gold: `#E6A51D`
- **Supporting Accents:**
  - Deep burgundy clothing: `#651E2B` (shadow: `#361019`)
  - YouTube red: `#FF0000` (strictly for the official icon)
  - Dark wood table: `#4B281D` (highlight: `#92553B`, shadow: `#27130E`)
  - Icon white: `#F5F4F0`
  - Divider gray: `#A5A5A5` at 60% opacity

### Color Distribution Formula

- **65%:** Deep navy, charcoal, and black (atmospheric background and shadow separation).
- **15%:** Warm brown and burgundy (studio warmth, clothing, and table textures).
- **12%:** White and light gray (high-contrast typography and readable headlines).
- **7%:** Gold accents (focal names, "1%", and hand-painted underlines).
- **<= 1%:** YouTube red (official logo badge only).

---

## 3. 3-Zone Cinematic Background Transition

The background must never be a flat solid color or an artificial neon gradient. Construct a seamless left-to-right cinematic transition:

1. **Left Background (Warm Studio & Bookshelves):**
   - Base color: dark charcoal brown `#211612`.
   - Add heavily blurred bookshelves and warm office ambient lights in soft amber `#C47A32`.
   - Low-contrast, creamy bokeh with zero sharp edges.
2. **Center Background (Contrast Separation Void):**
   - Base color: dark blue-black `#0A1424`.
   - The darkest area of the background sits directly behind the person's face and shoulders to create maximum separation and pop.
   - Subtle cool edge light in desaturated blue-gray `#6F8193`.
   - No bright objects directly behind the head.
3. **Right Background (Distant Architectural Silhouette):**
   - Base color: midnight navy `#07111F`.
   - Add a subtle, distant professional city or architectural skyline near the lower-right in muted gray-blue `#536170` at 20–30% opacity.
   - Very subtle warm horizon glow in dusty amber `#B47646` near the bottom.
   - Background remains darker than all foreground text and subject layers.

---

## 4. Input Capture & Verification

Before generating any design concepts or prompts, capture and validate the following inputs:

1. `channel_name` / `host_name`: Exact name of the creator or brand (e.g., "MD ALIM UL KARIM").
2. `primary_title`: Main headline or topic (e.g., "Better Ideas, Bigger Impact").
3. `subtitles_and_credentials`: Professional titles or roles separated by vertical bars (e.g., `Author | Marketer | Trainer | Consultant | Podcaster`).
4. `core_pillars`: 3 key thematic words (e.g., `Ideas | Strategy | Impact`).
5. `achievement_callout`: Main authority badge or metric (e.g., `"Top 1%"`).
6. `supporting_statement`: 1–2 line hook (e.g., `"Real Stories, Real People, Real Growth."`).
7. `youtube_handle`: Exact channel handle (e.g., `@alimulkarim`).
8. `youtube_url`: Full channel URL (e.g., `https://youtube.com/@alimulkarim`).
9. `website_url`: Creator or brand website URL (e.g., `https://alimulkarim.com`).
10. `email_address`: Contact or business inquiry email address to embed (e.g., `contact@alimulkarim.com`).
11. `qr_code`: Optional scannable QR code destination URL or asset to embed (optional).
12. `channel_icon_or_avatar`: Profile picture, logo mark, or avatar image (optional — ask user if available, or skip).
13. `person_photo_or_avatar`: High-resolution photo of the person, visual description, or subject cutout (optional — ask user if available, or skip).
14. `font_family`: Primary font family (Defaults to **Ubuntu** across all typography, titles, and overlays, with brush-script style for quotes).
15. `color_palette`: Primary brand colors (defaults to the 65/15/12/7/1 system above).

> [!IMPORTANT]
> Always ask the user if they have a **person's photo, subject cutout, or avatar image** to include, or if they prefer to **skip** it (for a graphics/typography-focused layout). Remember: do NOT generate unsolicited avatars, clip-art icons, or star badges. Also ask if they want to embed an **email address** or a **scannable QR code**. If `channel_name`, `primary_title`, `youtube_handle`, or `website_url` are missing or not provided, **STOP and ask the user** before proceeding.

---

## 5. Photographic, Lighting & Subject Specification

Professional YouTube thumbnails require deliberate cinematic direction rather than generic AI generation:

### A. Camera Shot & Framing

- **Lens & Optics:** 85mm prime portrait lens equivalent with a wide aperture (`f/1.8` to `f/2.8`), creating an authentic shallow depth of field where the subject remains razor-sharp while the background dissolves into soft, creamy bokeh.
- **Shot Distance:** Seated medium bust or torso shot (from waist/hands to head). The subject occupies approximately `54–58%` of canvas height from head to seated hands, with lower body continuing to the bottom edge.
- **Placement (Zone B):** Center of face at approximately `x = 555, y = 150`, portrait width `360–390 px`. Overlaps Zone A and Zone C slightly.
- **Pose & Expression:** Seated, facing forward, direct eye contact with the camera. Calm, confident, approachable expression with hands naturally clasped. Body turned no more than 5 degrees away.

### B. 3-Point Cinematic Lighting

- **Warm Key Light:** Upper-left softbox in warm amber `#F0B078`, creating natural facial modeling with gentle tonal transitions.
- **Soft Cool Fill:** Front-right ambient fill in muted blue-gray `#7D91A8`, preserving shadow detail on cheekbones and clothing.
- **Thin Warm Rim Light (Crucial):** Rim light in warm gold `#D99145` tracing the left shoulder, hair, and jawline for clean separation from dark backgrounds.
- **Shadow Quality:** Soft natural contact shadow beneath the chin and behind the body. No harsh cutout edges, no glowing halo outlines, and no white sticker borders.

### C. Texture & Environmental Realism

- **Organic Skin Texture:** Visible skin pores, natural micro-contrast, realistic skin tones, and authentic specular highlights. Total ban on plastic, waxy, or airbrushed AI skin.
- **Crisp Textile Weave:** High-fidelity fabric texture on clothing (e.g., deep burgundy overshirt `#651E2B`, white inner T-shirt `#F1F0EC`, dark navy trousers `#12213A`).
- **Glasses & Eyes:** Clear eyes behind lenses with controlled, minimal specular reflections.
- **Mandatory Bottom Gradient Fade:** Apply a soft vertical linear gradient fade across the bottom 20% of the subject cutout, allowing the lower body to dissolve organically into the canvas.

---

## 6. Component & Typography Specifications

### A. Left-Side Handwritten Quote (Zone A)

- **Exact Text:** e.g., `“Better Ideas, Bigger Impact”`
- **Position:** `x = 90, y = 55`, max width `285 px`, max height `125 px`. Keep at least `30 px` from books.
- **Font Style:** Bold handwritten brush-script (e.g., Caveat Brush, Kalam Bold, or Ubuntu Bold Script). Feels personally painted with visible brush variation.
- **Color:** Soft white `#F7F7F4` with subtle shadow (`#05070A` at 55% opacity, 3px offset).
- **Size:** Approximately `48–54 px`.
- **Underline:** Curved hand-painted brush stroke beneath the second line in warm gold `#E6A51D` (darker edge `#B96E08`), thickness `7–12 px`, rising slightly toward the right.

### B. Book Display on Wood Table (Zone A)

- **Position:** `x = 25 to 405 px`, `y = 225 to 475 px`, width `350–375 px`. Keep at least `25 px` from the subject's torso.
- **Four Upright Books:**
  1. Book 1: Mustard yellow cover `#F4BC24` with blue/white artwork.
  2. Book 2: Deep emerald green cover `#087158` with mint details.
  3. Book 3: Royal blue cover `#0758A4` with white details.
  4. Book 4: Warm off-white cover `#ECE8DF` with charcoal title and yellow accent.
- **Preserve Real Covers:** Preserve exact titles, typography, and Bengali or English lettering from supplied references. Do NOT invent fake book covers or garbled symbols.
- **Wood Table:** Dark wood `#4B281D`, highlight `#92553B`, shadow `#27130E` with realistic horizontal wood grain and soft warm reflections beneath the books.

### C. Main Name: Primary Title (Zone C)

- **Exact Stacked Wording:** e.g.,
  ```
  MD ALIM UL
  KARIM
  ```
- **Position:** Left edge `x = 660`, top `y = 85`, max width `300 px`. Keep `30–45 px` clear space from subject head, and `18–26 px` from Zone D. Both lines share the same left edge.
- **Font Family (Ubuntu Standard):** `Ubuntu Bold` or extra-bold geometric sans-serif (weight 800–900), tightly stacked with line height 0.82–0.9.
- **First Name Color:** Soft white `#F7F7F4`, lower shading `#D7D9DC`, dark shadow `#02060D` (65% opacity, 5px offset).
- **Last Name Color:** Warm golden yellow `#F5A817`, bright highlight `#FFB51B`, dark shadow `#B96E08`. Must be the strongest color accent in the center.

### D. Professional Credentials & Values Line (Zone C)

- **Credentials Block:** Directly beneath the name (`y = 305`, max width `320 px`).
  - Text: e.g., `Author | Marketer | Trainer` / `Consultant | Podcaster`
  - Font: `Ubuntu Medium` or clean brush script (size `26–31 px`, color `#ECEBE6`, separators `#C8C5BE`).
- **Values Line:** Below credentials (`y = 400`, max width `300 px`).
  - Text: `Ideas  |  Strategy  |  Impact`
  - Font: `Ubuntu Medium` (size `21–25 px`, color `#B8BBC2`, separators `#858B94`).
  - Underline: Thin hand-painted gold line `#E6A51D` (thickness `4–7 px`, rising 5–8 degrees toward the right).

### E. Achievement Area & Supporting Statement (Zone D)

- **"Top 1%" Callout:** Horizontal start `x = 945`, top `y = 165`.
  - "Top": Soft white `#F7F7F4`, size `78–90 px`.
  - "1%": Warm gold `#F5A817`, highlight `#FFB51B`, size `78–94 px`.
  - Both on the same line with precisely aligned baseline.
- **Supporting Statement:** Directly beneath `Top 1%` (`y = 305`).
  - Text: e.g., `Real Stories, Real People,` / `Real Growth.`
  - Font: `Ubuntu Medium` (size `22–27 px`, color `#B8BBC2`).
- **YouTube Channel Line:** Below supporting statement (`y = 385`).
  - Official YouTube play button icon (red `#FF0000`, white triangle, 30–34px wide).
  - Channel name or handle in `Ubuntu Bold` (`18–22 px`, color `#ECEBE6`).

### F. Podcast Microphone (Zone D)

- **Position:** Head centered near `x = 1220, y = 235`, entering from far-right edge at `25–35 degrees` downward-left angle.
- **Color & Finish:** Dark metal body `#171A20`, mesh highlights `#5B616B`, subtle burgundy reflection `#651E2B`, tiny warm highlight `#D99145`. No bright chrome or silver.
- **Depth:** Realistic shallow depth of field, slightly less sharp than the subject's face.

### G. Lower Professional Role Strip (Zone E)

- **Position:** `x = 675 to 1280 px`, `y = 550 to 720 px`. Positioned behind the subject where they overlap.
- **Background:** Base `#211612` at 88% opacity, upper edge `#4B281D` at 50%, lower edge `#0D0908` at 95%. Subtle 1–2px top separation line in `#7F5A45` at 35% opacity.
- **Four Role Groups (width ~140–150 px each, separated by 60% opacity dividers `#A5A5A5`):**
  1. Open-book line icon + `Bestselling Author`
  2. Graduation-cap line icon + `Hard-skill Trainer`
  3. Megaphone line icon + `Brand Marketing Professional`
  4. Briefcase line icon + `Business Consultant`
- **Icon & Label Typography:** Clean white line icons (`#F5F4F0`, 42–50px), labels in `Ubuntu Medium` (`18–23 px`, `#F5F4F0`, centered).

---

## 7. Spacing, Visual Separation & Required Reading Hierarchy

### Strict Spacing Rules

- Canvas edge to important text: minimum `55 pixels`
- Person's face to name: `30–45 pixels`
- Name to "Top 1%" area: `18–26 pixels`
- First name to last name gap: `0–8 pixels`
- Name to credential block: `18–25 pixels`
- Credentials to values line: `16–22 pixels`
- Values line to gold underline: `8–12 pixels`
- "Top 1%" to supporting statement: `18–26 pixels`
- Supporting statement to YouTube line: `15–20 pixels`
- Books to portrait: minimum `25 pixels`
- Quote to books: minimum `30 pixels`
- Microphone to achievement text: minimum `20 pixels`

### Required Reading Hierarchy (Mobile & Desktop)

1. The featured person's face & eye contact
2. The primary stacked name (white first name + gold last name)
3. The "Top 1%" achievement callout
4. The professional credentials & roles
5. The four colorful books on the table
6. The handwritten quote ("Better Ideas, Bigger Impact")
7. The lower professional role strip

---

## 8. Text Accuracy & Anti-Hallucination Protocol

When generating prompts for image engines (Flux, Midjourney v6, Ideogram) or compositing layers:

1. **Explicit Text Quoting:** Always specify text inside literal quotes in the generation prompt:
   - `with the exact text "MD ALIM UL" in bold white Ubuntu letters, and "KARIM" in bold warm-gold Ubuntu letters`
2. **Character Verification Gate:** Inspect the generated output. If even a single character is warped, merged, or misspelled, the image MUST be rejected or the text layer must be re-rendered as a clean vector overlay.
3. **Hybrid Compositing (Recommended):** For production-grade thumbnails:
   - Use AI to generate the photographic background, cinematic lighting, wood table, and subject portrait.
   - Render all typography, badges, book covers, and icons as crisp SVG vector layers over the background to guarantee 100% spelling precision.

---

## 9. Directory & File Hierarchy (Strict Lowercase & Two-Digit Sequence)

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

## 10. Asset Specifications

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

## 11. Master Production Sample Prompt & Templates

### Template 1: Master Production Personal-Brand Thumbnail Prompt

Use this complete reference prompt when generating a cinematic, authority-grade YouTube thumbnail:

```text
Create a highly polished, cinematic personal-brand YouTube thumbnail at exactly 1280 × 720 pixels, 16:9 aspect ratio.

Recreate a cinematic personal-brand composition with a 5-zone layout grid:
1. Left Content Zone (x: 25 to 405 px): In the upper-left, the handwritten quote “Better Ideas, Bigger Impact” in soft white #F7F7F4 brush lettering with a curved hand-painted gold underline in #E6A51D. Below the quote, four upright colorful books standing on a dark polished wood table (#4B281D) with realistic horizontal grain and soft warm reflection.
2. Portrait Zone (x: 355 to 735 px): Seated professional subject facing forward with direct eye contact, relaxed confident expression, hands clasped in lap, wearing a deep burgundy overshirt #651E2B over a clean white T-shirt, black glasses. Captured with an 85mm portrait lens at f/1.8, razor-sharp focus on face and eyes, organic skin texture with visible pores. Warm key light from upper-left #F0B078, soft cool fill from front-right #7D91A8, and a thin warm gold rim light #D99145 tracing the left shoulder. The lower torso dissolves seamlessly into a soft dark gradient fade at the bottom edge.
3. Name & Credential Zone (x: 655 to 965 px): Immediately right of the subject's head, prominently display the stacked name "MD ALIM UL" in bold white Ubuntu letters (#F7F7F4) and "KARIM" in bold warm-gold Ubuntu letters (#F5A817). Beneath the name, professional credentials "Author | Marketer | Trainer | Consultant | Podcaster" in soft white Ubuntu Medium (#ECEBE6), followed by the values line "Ideas | Strategy | Impact" in light gray (#B8BBC2) with a thin hand-painted gold underline #E6A51D.
4. Achievement Zone (x: 930 to 1245 px): On the right side, "Top 1%" with "Top" in white #F7F7F4 and "1%" in warm gold #F5A817. Beneath it, "Real Stories, Real People, Real Growth." in clean gray #B8BBC2, and a YouTube play button with verified channel handle. A dark professional podcast microphone enters from the far-right edge angled downward at 30 degrees.
5. Lower Role Strip (x: 670 to 1280 px, y: 550 to 720 px): A dark translucent strip in #211612 behind the lower body with four evenly spaced professional role groups separated by thin 60% opacity gray dividers: Bestselling Author, Hard-skill Trainer, Brand Marketing Professional, and Business Consultant, each with a minimalist white line icon.

Background: Seamless transition from dark charcoal brown with blurred warm office lights on the left, to dark blue-black directly behind the subject's head for strong separation, to midnight navy with a subtle architectural silhouette at 20% opacity on the right. High commercial photography quality, 8k resolution, zero text hallucination, no clip-art stickers --ar 16:9 --style raw
```

### Template 2: Flux / Midjourney Photorealistic Scene Generation Prompt

```text
Cinematic 16:9 YouTube thumbnail scene. Seated professional subject captured in a medium bust portrait with an 85mm portrait lens at f/1.8, razor-sharp focus on face and eyes, natural skin texture with visible pores, authentic confident expression, wearing a modern deep burgundy overshirt over a white inner shirt. Dramatic 3-point studio lighting with a warm soft key light from the upper-left and vibrant gold rim light tracing the hair and shoulder. The lower torso dissolves seamlessly into a soft dark gradient fade at the bottom edge. On the left, four colorful published books on a polished dark mahogany table with soft reflections. Dark atmospheric studio background with deep charcoal tones on the left and midnight navy on the right with subtle distant architectural bokeh. Ample clean negative space in the center and right for typography overlay. High commercial photography quality, 8k resolution --ar 16:9 --style raw
```

### Template 3: Figma / SVG Typography & Vector Overlay Specification (Ubuntu Font Family)

```markdown
# Canvas: 1280 x 720 px (16:9, Ubuntu Font Family)

## Zone A: Quote (x: 90, y: 55)
- Line 1: "“Better Ideas" (Font: Caveat Brush / Ubuntu Bold Script, Color: #F7F7F4, Size: 50pt)
- Line 2: "Bigger Impact”" (Font: Caveat Brush / Ubuntu Bold Script, Color: #F7F7F4, Size: 50pt)
- Underline: Hand-painted curve (Color: #E6A51D, Width: 220px, Height: 8px, y: 165px)

## Zone C: Main Name (x: 660, y: 85)
- First Line: "MD ALIM UL" (Font: Ubuntu Bold, Color: #F7F7F4, Size: 86pt, Shadow: #02060D 65% 5px down-right)
- Second Line: "KARIM" (Font: Ubuntu Bold, Color: #F5A817, Size: 86pt, Shadow: #02060D 70% 5px down-right)
- Credentials (y: 305): "Author | Marketer | Trainer | Consultant | Podcaster" (Font: Ubuntu Medium, Color: #ECEBE6, Size: 28pt)
- Values (y: 400): "Ideas  |  Strategy  |  Impact" (Font: Ubuntu Medium, Color: #B8BBC2, Size: 23pt)
- Values Underline: Hand-painted gold stroke (Color: #E6A51D, Width: 200px, Height: 5px, y: 435px)

## Zone D: Achievement (x: 945, y: 165)
- Achievement: "Top" (#F7F7F4) + " 1%" (#F5A817) (Font: Ubuntu Bold, Size: 84pt)
- Hook (y: 305): "Real Stories, Real People,\nReal Growth." (Font: Ubuntu Medium, Color: #B8BBC2, Size: 24pt)
- YouTube Line (y: 385): YouTube Play Icon (32px, #FF0000) + "{youtube_handle}" (Font: Ubuntu Bold, Color: #ECEBE6, Size: 20pt)

## Zone E: Lower Role Strip (x: 675, y: 550 to 720)
- Background: #211612 at 88% opacity, top line #7F5A45 at 35%
- Group 1: Book Icon + "Bestselling Author" (Font: Ubuntu Medium, Color: #F5F4F0, Size: 19pt)
- Group 2: Cap Icon + "Hard-skill Trainer" (Font: Ubuntu Medium, Color: #F5F4F0, Size: 19pt)
- Group 3: Megaphone Icon + "Brand Marketing Professional" (Font: Ubuntu Medium, Color: #F5F4F0, Size: 19pt)
- Group 4: Briefcase Icon + "Business Consultant" (Font: Ubuntu Medium, Color: #F5F4F0, Size: 19pt)
```
