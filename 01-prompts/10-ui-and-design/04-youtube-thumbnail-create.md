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
3. **NO Saving Binary Images to Repository:** Do NOT save raw binary image files (`.png`, `.jpg`) into the git repository. All instructions, layouts, and SVG vector overlays must live purely in markdown or code.
4. **NO Low-Contrast Text:** NEVER place light text over bright, busy backgrounds or dark text over dark shadows without proper contrast treatment (e.g., dropshadows, dark vignettes, or solid badge backdrops).
5. **NO Uncontrolled Facial or Hand Distortions:** The human subject in the foreground must have anatomically correct eyes, glasses, hands, and fingers without AI melting or extra digits.
6. **NO HTML or Web Page Code:** Do NOT generate full HTML web pages or application templates. This is strictly a graphic design, thumbnail, and banner workflow.

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
10. `channel_icon_or_avatar`: Profile picture, logo mark, or avatar image description.
11. `subject_description`: Appearance, pose, clothing, and expression of the person in front.
12. `color_palette`: Primary brand colors (e.g., Deep Navy Background `#0f172a`, Golden Yellow Highlight `#facc15`, Crisp White `#ffffff`).

> [!IMPORTANT]
> If `channel_name`, `primary_title`, `youtube_handle`, `website_url`, or `subject_description` / `channel_icon_or_avatar` are missing or not provided, **STOP and ask the user** before proceeding.

---

## 2. Deep Compositional & Visual Breakdown (Reference Model)

The reference banner layout below demonstrates the target standard for high-authority YouTube channel banners and podcast thumbnails.

### Visual Reference (Inline Compressed Thumbnail)

<img src="data:image/jpeg;base64,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" width="400" alt="Reference Layout" />

### Layout Anatomy Breakdown

1. **Center Foreground (Hero Subject):**
   - Professional subject positioned centrally or slightly left-of-center.
   - Clean subject cutout with sharp contrast against the background.
   - Professional, grounded pose (hands clasped, natural posture, glasses, eye contact with the viewer).
   - Crisp lighting highlighting the face, beard, and shoulders.

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

## 5. Ready-to-Use Prompt Templates

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
