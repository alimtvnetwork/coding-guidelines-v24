# Source ledger rows 1

Updated: 2026-10-02

| Status | Spec file | Line | Value | Source symbol | Spec line |
|---|---|---|---|---|---|
| unverified | `24-slide-presentation-system.md` | 1 | `24` | No source row yet. Do not treat this number as measured. | # 24 — Slide Presentation System & Presenter Engine Architecture |
| unverified | `24-slide-presentation-system.md` | 3 | `27` | No source row yet. Do not treat this number as measured. | > **/goal** Specify webcam picture-in-picture, step reveals, and the dual-screen presenter console. Canvas, themes, layouts, HUD, and the sl |
| unverified | `24-slide-presentation-system.md` | 3 | `28` | No source row yet. Do not treat this number as measured. | > **/goal** Specify webcam picture-in-picture, step reveals, and the dual-screen presenter console. Canvas, themes, layouts, HUD, and the sl |
| unverified | `24-slide-presentation-system.md` | 3 | `29` | No source row yet. Do not treat this number as measured. | > **/goal** Specify webcam picture-in-picture, step reveals, and the dual-screen presenter console. Canvas, themes, layouts, HUD, and the sl |
| unverified | `24-slide-presentation-system.md` | 4 | `1920` | No source row yet. Do not treat this number as measured. | > **/learn** Master the responsive coordinate transforms (`1920x1080` canvas), `PresenterWebcamOverlay` webcam stream integration, `stepMoti |
| unverified | `24-slide-presentation-system.md` | 6 | `1.0` | No source row yet. Do not treat this number as measured. | **Version:** 1.0.0 |
| unverified | `24-slide-presentation-system.md` | 7 | `2026` | No source row yet. Do not treat this number as measured. | **Updated:** 2026-09-24 |
| unverified | `24-slide-presentation-system.md` | 7 | `09` | No source row yet. Do not treat this number as measured. | **Updated:** 2026-09-24 |
| unverified | `24-slide-presentation-system.md` | 7 | `24` | No source row yet. Do not treat this number as measured. | **Updated:** 2026-09-24 |
| unverified | `24-slide-presentation-system.md` | 14 | `1` | No source row yet. Do not treat this number as measured. | ## 1. System Overview & Presentation Repositories Synthesis |
| unverified | `24-slide-presentation-system.md` | 16 | `1920` | No source row yet. Do not treat this number as measured. | Web-based decks use one `1920×1080` stage, scaled as specified in `27-slide-canvas-and-themes.md`. This file keeps five presenter pillars. I |
| unverified | `24-slide-presentation-system.md` | 16 | `1080` | No source row yet. Do not treat this number as measured. | Web-based decks use one `1920×1080` stage, scaled as specified in `27-slide-canvas-and-themes.md`. This file keeps five presenter pillars. I |
| unverified | `24-slide-presentation-system.md` | 16 | `27` | No source row yet. Do not treat this number as measured. | Web-based decks use one `1920×1080` stage, scaled as specified in `27-slide-canvas-and-themes.md`. This file keeps five presenter pillars. I |
| unverified | `24-slide-presentation-system.md` | 18 | `1` | No source row yet. Do not treat this number as measured. | 1. **Fixed-Aspect Ratio Virtual Canvas:** Standard `1920x1080` coordinate space scaled down dynamically using CSS `transform: scale(min(w/19 |
| unverified | `24-slide-presentation-system.md` | 18 | `1920` | No source row yet. Do not treat this number as measured. | 1. **Fixed-Aspect Ratio Virtual Canvas:** Standard `1920x1080` coordinate space scaled down dynamically using CSS `transform: scale(min(w/19 |
| unverified | `24-slide-presentation-system.md` | 19 | `2` | No source row yet. Do not treat this number as measured. | 2. **Floating Draggable Webcam PIP (`PresenterWebcamOverlay`):** Video stream with circular boundary, glowing accent perimeter, draggable co |
| unverified | `24-slide-presentation-system.md` | 20 | `3` | No source row yet. Do not treat this number as measured. | 3. **Step-by-Step Progressive Reveals:** Sub-slide animations where individual bullets, code blocks, or diagram nodes reveal on subsequent a |
| unverified | `24-slide-presentation-system.md` | 21 | `4` | No source row yet. Do not treat this number as measured. | 4. **Presenter Dual-Screen Console:** Broadcast channel sync communicating between audience display and presenter control view (notes, elaps |
| unverified | `24-slide-presentation-system.md` | 22 | `5` | No source row yet. Do not treat this number as measured. | 5. **Standalone Vector Layouts:** System illustrated via standalone SVG diagrams in `01-svg/slide-layout.svg`. |
| unverified | `24-slide-presentation-system.md` | 22 | `01` | No source row yet. Do not treat this number as measured. | 5. **Standalone Vector Layouts:** System illustrated via standalone SVG diagrams in `01-svg/slide-layout.svg`. |
| unverified | `24-slide-presentation-system.md` | 26 | `2` | No source row yet. Do not treat this number as measured. | ## 2. 16:9 Canvas Virtual Coordinate Architecture |
| unverified | `24-slide-presentation-system.md` | 26 | `16` | No source row yet. Do not treat this number as measured. | ## 2. 16:9 Canvas Virtual Coordinate Architecture |
| unverified | `24-slide-presentation-system.md` | 26 | `9` | No source row yet. Do not treat this number as measured. | ## 2. 16:9 Canvas Virtual Coordinate Architecture |
| unverified | `24-slide-presentation-system.md` | 35 | `1920px` | No source row yet. Do not treat this number as measured. | /   / VIRTUAL SLIDE CANVAS (Fixed 1920px x 1080px)                          /   / |
| unverified | `24-slide-presentation-system.md` | 35 | `1080px` | No source row yet. Do not treat this number as measured. | /   / VIRTUAL SLIDE CANVAS (Fixed 1920px x 1080px)                          /   / |
| unverified | `24-slide-presentation-system.md` | 38 | `1920` | No source row yet. Do not treat this number as measured. | /   / transform: scale(min(viewportWidth / 1920, viewportHeight / 1080));  /   / |
| unverified | `24-slide-presentation-system.md` | 38 | `1080` | No source row yet. Do not treat this number as measured. | /   / transform: scale(min(viewportWidth / 1920, viewportHeight / 1080));  /   / |
| unverified | `24-slide-presentation-system.md` | 42 | `160px` | No source row yet. Do not treat this number as measured. | /   /                             [Draggable Webcam PIP: 160px x 160px]     /   / |
| unverified | `24-slide-presentation-system.md` | 42 | `160px` | No source row yet. Do not treat this number as measured. | /   /                             [Draggable Webcam PIP: 160px x 160px]     /   / |
| unverified | `24-slide-presentation-system.md` | 48 | `2.1` | No source row yet. Do not treat this number as measured. | ### 2.1 Responsive Canvas Scaling (TypeScript) |
| unverified | `24-slide-presentation-system.md` | 60 | `1920` | No source row yet. Do not treat this number as measured. | baseWidth = 1920, |
| unverified | `24-slide-presentation-system.md` | 61 | `1080` | No source row yet. Do not treat this number as measured. | baseHeight = 1080 |
| unverified | `24-slide-presentation-system.md` | 64 | `2` | No source row yet. Do not treat this number as measured. | const offsetX = (containerWidth - baseWidth * scale) / 2; |
| unverified | `24-slide-presentation-system.md` | 65 | `2` | No source row yet. Do not treat this number as measured. | const offsetY = (containerHeight - baseHeight * scale) / 2; |
| unverified | `24-slide-presentation-system.md` | 73 | `3` | No source row yet. Do not treat this number as measured. | ## 3. Draggable Webcam PIP Overlay (`PresenterWebcamOverlay`) |
| unverified | `24-slide-presentation-system.md` | 77 | `3.1` | No source row yet. Do not treat this number as measured. | ### 3.1 Component Features |
| unverified | `24-slide-presentation-system.md` | 78 | `9999px` | No source row yet. Do not treat this number as measured. | - Circular or squircle geometry with customizable corner radius (`border-radius: 9999px` or `2rem`). |
| unverified | `24-slide-presentation-system.md` | 78 | `2rem` | No source row yet. Do not treat this number as measured. | - Circular or squircle geometry with customizable corner radius (`border-radius: 9999px` or `2rem`). |
| unverified | `24-slide-presentation-system.md` | 81 | `1` | No source row yet. Do not treat this number as measured. | - Video flip/mirror control (`transform: scaleX(-1)`). |
| unverified | `24-slide-presentation-system.md` | 82 | `120px` | No source row yet. Do not treat this number as measured. | - Quick keyboard shortcut (`C` key) to cycle size presets (Small 120px, Medium 180px, Large 240px, Hidden). |
| unverified | `24-slide-presentation-system.md` | 82 | `180px` | No source row yet. Do not treat this number as measured. | - Quick keyboard shortcut (`C` key) to cycle size presets (Small 120px, Medium 180px, Large 240px, Hidden). |
| unverified | `24-slide-presentation-system.md` | 82 | `240px` | No source row yet. Do not treat this number as measured. | - Quick keyboard shortcut (`C` key) to cycle size presets (Small 120px, Medium 180px, Large 240px, Hidden). |
| unverified | `24-slide-presentation-system.md` | 84 | `3.2` | No source row yet. Do not treat this number as measured. | ### 3.2 LESS Styling (Preferred) |
| unverified | `24-slide-presentation-system.md` | 93 | `9999` | No source row yet. Do not treat this number as measured. | z-index: 9999; |
| unverified | `24-slide-presentation-system.md` | 94 | `180px` | No source row yet. Do not treat this number as measured. | width: 180px; |
| unverified | `24-slide-presentation-system.md` | 95 | `180px` | No source row yet. Do not treat this number as measured. | height: 180px; |
| unverified | `24-slide-presentation-system.md` | 96 | `9999px` | No source row yet. Do not treat this number as measured. | border-radius: 9999px; |
| unverified | `24-slide-presentation-system.md` | 98 | `3px` | No source row yet. Do not treat this number as measured. | border: 3px solid #8b5cf6; |
| unverified | `24-slide-presentation-system.md` | 98 | `8` | No source row yet. Do not treat this number as measured. | border: 3px solid #8b5cf6; |
| unverified | `24-slide-presentation-system.md` | 99 | `0` | No source row yet. Do not treat this number as measured. | box-shadow: 0 10px 25px -5px rgba(139, 92, 246, 0.5), |
| unverified | `24-slide-presentation-system.md` | 99 | `10px` | No source row yet. Do not treat this number as measured. | box-shadow: 0 10px 25px -5px rgba(139, 92, 246, 0.5), |
| unverified | `24-slide-presentation-system.md` | 99 | `25px` | No source row yet. Do not treat this number as measured. | box-shadow: 0 10px 25px -5px rgba(139, 92, 246, 0.5), |
| unverified | `24-slide-presentation-system.md` | 99 | `5px` | No source row yet. Do not treat this number as measured. | box-shadow: 0 10px 25px -5px rgba(139, 92, 246, 0.5), |
| unverified | `24-slide-presentation-system.md` | 99 | `139` | No source row yet. Do not treat this number as measured. | box-shadow: 0 10px 25px -5px rgba(139, 92, 246, 0.5), |
| unverified | `24-slide-presentation-system.md` | 99 | `92` | No source row yet. Do not treat this number as measured. | box-shadow: 0 10px 25px -5px rgba(139, 92, 246, 0.5), |
| unverified | `24-slide-presentation-system.md` | 99 | `246` | No source row yet. Do not treat this number as measured. | box-shadow: 0 10px 25px -5px rgba(139, 92, 246, 0.5), |
| unverified | `24-slide-presentation-system.md` | 99 | `0.5` | No source row yet. Do not treat this number as measured. | box-shadow: 0 10px 25px -5px rgba(139, 92, 246, 0.5), |
| unverified | `24-slide-presentation-system.md` | 100 | `0` | No source row yet. Do not treat this number as measured. | 0 0 0 1px rgba(255, 255, 255, 0.1); |
| unverified | `24-slide-presentation-system.md` | 100 | `0` | No source row yet. Do not treat this number as measured. | 0 0 0 1px rgba(255, 255, 255, 0.1); |
| unverified | `24-slide-presentation-system.md` | 100 | `0` | No source row yet. Do not treat this number as measured. | 0 0 0 1px rgba(255, 255, 255, 0.1); |
| unverified | `24-slide-presentation-system.md` | 100 | `1px` | No source row yet. Do not treat this number as measured. | 0 0 0 1px rgba(255, 255, 255, 0.1); |
| unverified | `24-slide-presentation-system.md` | 100 | `255` | No source row yet. Do not treat this number as measured. | 0 0 0 1px rgba(255, 255, 255, 0.1); |
| unverified | `24-slide-presentation-system.md` | 100 | `255` | No source row yet. Do not treat this number as measured. | 0 0 0 1px rgba(255, 255, 255, 0.1); |
| unverified | `24-slide-presentation-system.md` | 100 | `255` | No source row yet. Do not treat this number as measured. | 0 0 0 1px rgba(255, 255, 255, 0.1); |
| unverified | `24-slide-presentation-system.md` | 100 | `0.1` | No source row yet. Do not treat this number as measured. | 0 0 0 1px rgba(255, 255, 255, 0.1); |
| unverified | `24-slide-presentation-system.md` | 104 | `200ms` | No source row yet. Do not treat this number as measured. | transition: box-shadow 200ms ease, border-color 200ms ease; |
| unverified | `24-slide-presentation-system.md` | 104 | `200ms` | No source row yet. Do not treat this number as measured. | transition: box-shadow 200ms ease, border-color 200ms ease; |
| unverified | `24-slide-presentation-system.md` | 108 | `0` | No source row yet. Do not treat this number as measured. | box-shadow: 0 15px 35px -5px rgba(139, 92, 246, 0.7); |
| unverified | `24-slide-presentation-system.md` | 108 | `15px` | No source row yet. Do not treat this number as measured. | box-shadow: 0 15px 35px -5px rgba(139, 92, 246, 0.7); |
| unverified | `24-slide-presentation-system.md` | 108 | `35px` | No source row yet. Do not treat this number as measured. | box-shadow: 0 15px 35px -5px rgba(139, 92, 246, 0.7); |
| unverified | `24-slide-presentation-system.md` | 108 | `5px` | No source row yet. Do not treat this number as measured. | box-shadow: 0 15px 35px -5px rgba(139, 92, 246, 0.7); |
| unverified | `24-slide-presentation-system.md` | 108 | `139` | No source row yet. Do not treat this number as measured. | box-shadow: 0 15px 35px -5px rgba(139, 92, 246, 0.7); |
| unverified | `24-slide-presentation-system.md` | 108 | `92` | No source row yet. Do not treat this number as measured. | box-shadow: 0 15px 35px -5px rgba(139, 92, 246, 0.7); |
| unverified | `24-slide-presentation-system.md` | 108 | `246` | No source row yet. Do not treat this number as measured. | box-shadow: 0 15px 35px -5px rgba(139, 92, 246, 0.7); |
| unverified | `24-slide-presentation-system.md` | 108 | `0.7` | No source row yet. Do not treat this number as measured. | box-shadow: 0 15px 35px -5px rgba(139, 92, 246, 0.7); |
| unverified | `24-slide-presentation-system.md` | 112 | `100%` | No source row yet. Do not treat this number as measured. | width: 100%; |
| unverified | `24-slide-presentation-system.md` | 113 | `100%` | No source row yet. Do not treat this number as measured. | height: 100%; |
| unverified | `24-slide-presentation-system.md` | 118 | `1` | No source row yet. Do not treat this number as measured. | transform: scaleX(-1); |
| unverified | `24-slide-presentation-system.md` | 124 | `12px` | No source row yet. Do not treat this number as measured. | top: 12px; |
| unverified | `24-slide-presentation-system.md` | 125 | `12px` | No source row yet. Do not treat this number as measured. | right: 12px; |
| unverified | `24-slide-presentation-system.md` | 126 | `10px` | No source row yet. Do not treat this number as measured. | width: 10px; |
| unverified | `24-slide-presentation-system.md` | 127 | `10px` | No source row yet. Do not treat this number as measured. | height: 10px; |
| unverified | `24-slide-presentation-system.md` | 128 | `9999px` | No source row yet. Do not treat this number as measured. | border-radius: 9999px; |
| unverified | `24-slide-presentation-system.md` | 129 | `10` | No source row yet. Do not treat this number as measured. | background-color: #10b981; |
| unverified | `24-slide-presentation-system.md` | 130 | `2px` | No source row yet. Do not treat this number as measured. | border: 2px solid #0b1329; |
| unverified | `24-slide-presentation-system.md` | 130 | `0` | No source row yet. Do not treat this number as measured. | border: 2px solid #0b1329; |
| unverified | `24-slide-presentation-system.md` | 137 | `4` | No source row yet. Do not treat this number as measured. | ## 4. Incremental Step Reveal Architecture |
| unverified | `24-slide-presentation-system.md` | 150 | `4.1` | No source row yet. Do not treat this number as measured. | ### 4.1 Step Motion Execution |
| unverified | `24-slide-presentation-system.md` | 151 | `1` | No source row yet. Do not treat this number as measured. | Elements with `data-step="1"`, `data-step="2"` stay hidden or dimmed (`opacity: 0.15; filter: blur(2px);`) until the presenter reaches their |
| unverified | `24-slide-presentation-system.md` | 151 | `2` | No source row yet. Do not treat this number as measured. | Elements with `data-step="1"`, `data-step="2"` stay hidden or dimmed (`opacity: 0.15; filter: blur(2px);`) until the presenter reaches their |
| unverified | `24-slide-presentation-system.md` | 151 | `0.15` | No source row yet. Do not treat this number as measured. | Elements with `data-step="1"`, `data-step="2"` stay hidden or dimmed (`opacity: 0.15; filter: blur(2px);`) until the presenter reaches their |
| unverified | `24-slide-presentation-system.md` | 151 | `2px` | No source row yet. Do not treat this number as measured. | Elements with `data-step="1"`, `data-step="2"` stay hidden or dimmed (`opacity: 0.15; filter: blur(2px);`) until the presenter reaches their |
| unverified | `24-slide-presentation-system.md` | 151 | `0.16` | No source row yet. Do not treat this number as measured. | Elements with `data-step="1"`, `data-step="2"` stay hidden or dimmed (`opacity: 0.15; filter: blur(2px);`) until the presenter reaches their |
| unverified | `24-slide-presentation-system.md` | 151 | `1` | No source row yet. Do not treat this number as measured. | Elements with `data-step="1"`, `data-step="2"` stay hidden or dimmed (`opacity: 0.15; filter: blur(2px);`) until the presenter reaches their |
| unverified | `24-slide-presentation-system.md` | 151 | `0.3` | No source row yet. Do not treat this number as measured. | Elements with `data-step="1"`, `data-step="2"` stay hidden or dimmed (`opacity: 0.15; filter: blur(2px);`) until the presenter reaches their |
| unverified | `24-slide-presentation-system.md` | 151 | `1` | No source row yet. Do not treat this number as measured. | Elements with `data-step="1"`, `data-step="2"` stay hidden or dimmed (`opacity: 0.15; filter: blur(2px);`) until the presenter reaches their |
| unverified | `24-slide-presentation-system.md` | 155 | `5` | No source row yet. Do not treat this number as measured. | ## 5. Deck Registry & URL Routing |
| unverified | `24-slide-presentation-system.md` | 163 | `01` | No source row yet. Do not treat this number as measured. | │   │   ├── 01-agentic-ai/ |
| unverified | `24-slide-presentation-system.md` | 164 | `01` | No source row yet. Do not treat this number as measured. | │   │   │   ├── 01-intro.tsx |
| unverified | `24-slide-presentation-system.md` | 165 | `02` | No source row yet. Do not treat this number as measured. | │   │   │   ├── 02-architecture.tsx |
| unverified | `24-slide-presentation-system.md` | 167 | `02` | No source row yet. Do not treat this number as measured. | │   │   └── 02-coding-guidelines/ |
| unverified | `24-slide-presentation-system.md` | 168 | `01` | No source row yet. Do not treat this number as measured. | │   │       ├── 01-booleans.tsx |
| unverified | `24-slide-presentation-system.md` | 169 | `02` | No source row yet. Do not treat this number as measured. | │   │       ├── 02-error-handling.tsx |
| unverified | `24-slide-presentation-system.md` | 179 | `4` | No source row yet. Do not treat this number as measured. | - Audience View: `/deck/01-agentic-ai/#/4` (Slide 4) |
| unverified | `24-slide-presentation-system.md` | 180 | `4` | No source row yet. Do not treat this number as measured. | - Presenter Console: `/deck/01-agentic-ai/presenter#/4` (Slide 4 with notes & timer) |
| unverified | `24-slide-presentation-system.md` | 184 | `6` | No source row yet. Do not treat this number as measured. | ## 6. Where the rest of the slide system lives |
| unverified | `24-slide-presentation-system.md` | 190 | `27` | No source row yet. Do not treat this number as measured. | / Canvas scale, JSON themes, noir-gold / `27-slide-canvas-and-themes.md` / |
| unverified | `24-slide-presentation-system.md` | 191 | `28` | No source row yet. Do not treat this number as measured. | / Closed layout union and real slots / `28-slide-layouts.md` / |
| unverified | `24-slide-presentation-system.md` | 192 | `29` | No source row yet. Do not treat this number as measured. | / HUD, keys, slide builder / `29-slide-navigation-and-builder.md` / |
| unverified | `24-slide-presentation-system.md` | 193 | `30` | No source row yet. Do not treat this number as measured. | / Dark amber palette, type scale, shell layers / `30-slide-palette-type-and-shell.md` / |
| unverified | `24-slide-presentation-system.md` | 194 | `31` | No source row yet. Do not treat this number as measured. | / Controller buttons, dots, crossfade / `31-slide-controller-buttons.md` / |
| unverified | `24-slide-presentation-system.md` | 195 | `32` | No source row yet. Do not treat this number as measured. | / Pill colors and nine-cell align / `32-slide-color-options.md` / |
| unverified | `24-slide-presentation-system.md` | 197 | `25` | No source row yet. Do not treat this number as measured. | A marketing page or a blog uses `25-page-assembly.md` and `04-white-blue-theme/`. It does not use this canvas. |
| unverified | `24-slide-presentation-system.md` | 197 | `04` | No source row yet. Do not treat this number as measured. | A marketing page or a blog uses `25-page-assembly.md` and `04-white-blue-theme/`. It does not use this canvas. |
| unverified | `25-page-assembly.md` | 1 | `25` | No source row yet. Do not treat this number as measured. | # 25 — Page Assembly (Sites and Blogs) |
| unverified | `25-page-assembly.md` | 6 | `1.0` | No source row yet. Do not treat this number as measured. | **Version:** 1.0.0 |
| unverified | `25-page-assembly.md` | 11 | `0` | No source row yet. Do not treat this number as measured. | ## 0. Anti-hallucination |
| unverified | `25-page-assembly.md` | 13 | `27` | No source row yet. Do not treat this number as measured. | If a length, color, duration, or section name is not in this file or in the files it cites, do not invent it. Marketing pages use the White  |
| unverified | `25-page-assembly.md` | 17 | `1` | No source row yet. Do not treat this number as measured. | 1. `02-spec/07-design-system/04-white-blue-theme/readme.md` |
| unverified | `25-page-assembly.md` | 17 | `02` | No source row yet. Do not treat this number as measured. | 1. `02-spec/07-design-system/04-white-blue-theme/readme.md` |
| unverified | `25-page-assembly.md` | 18 | `2` | No source row yet. Do not treat this number as measured. | 2. `02-spec/07-design-system/04-white-blue-theme/01-colors-typography-and-tokens.md` |
| unverified | `25-page-assembly.md` | 18 | `02` | No source row yet. Do not treat this number as measured. | 2. `02-spec/07-design-system/04-white-blue-theme/01-colors-typography-and-tokens.md` |
| unverified | `25-page-assembly.md` | 19 | `3` | No source row yet. Do not treat this number as measured. | 3. `02-spec/07-design-system/04-white-blue-theme/02-header-mega-menu-and-footer.md` |
| unverified | `25-page-assembly.md` | 19 | `02` | No source row yet. Do not treat this number as measured. | 3. `02-spec/07-design-system/04-white-blue-theme/02-header-mega-menu-and-footer.md` |
| unverified | `25-page-assembly.md` | 20 | `4` | No source row yet. Do not treat this number as measured. | 4. `02-spec/07-design-system/04-white-blue-theme/03-buttons-motion-and-interactions.md` |
| unverified | `25-page-assembly.md` | 20 | `02` | No source row yet. Do not treat this number as measured. | 4. `02-spec/07-design-system/04-white-blue-theme/03-buttons-motion-and-interactions.md` |
| unverified | `25-page-assembly.md` | 21 | `5` | No source row yet. Do not treat this number as measured. | 5. `02-spec/07-design-system/04-white-blue-theme/04-cards-heroes-and-section-library.md` |
| unverified | `25-page-assembly.md` | 21 | `02` | No source row yet. Do not treat this number as measured. | 5. `02-spec/07-design-system/04-white-blue-theme/04-cards-heroes-and-section-library.md` |
| unverified | `25-page-assembly.md` | 23 | `3` | No source row yet. Do not treat this number as measured. | Copy caps for a body section: eyebrow at most 3 words, H2 at most 8 words, lead at most 28 words, card body at most 32 words. The H1 is at m |
| unverified | `25-page-assembly.md` | 23 | `8` | No source row yet. Do not treat this number as measured. | Copy caps for a body section: eyebrow at most 3 words, H2 at most 8 words, lead at most 28 words, card body at most 32 words. The H1 is at m |
| unverified | `25-page-assembly.md` | 23 | `28` | No source row yet. Do not treat this number as measured. | Copy caps for a body section: eyebrow at most 3 words, H2 at most 8 words, lead at most 28 words, card body at most 32 words. The H1 is at m |
| unverified | `25-page-assembly.md` | 23 | `32` | No source row yet. Do not treat this number as measured. | Copy caps for a body section: eyebrow at most 3 words, H2 at most 8 words, lead at most 28 words, card body at most 32 words. The H1 is at m |
| unverified | `25-page-assembly.md` | 23 | `9` | No source row yet. Do not treat this number as measured. | Copy caps for a body section: eyebrow at most 3 words, H2 at most 8 words, lead at most 28 words, card body at most 32 words. The H1 is at m |
| unverified | `25-page-assembly.md` | 23 | `20` | No source row yet. Do not treat this number as measured. | Copy caps for a body section: eyebrow at most 3 words, H2 at most 8 words, lead at most 28 words, card body at most 32 words. The H1 is at m |
| unverified | `25-page-assembly.md` | 23 | `13` | No source row yet. Do not treat this number as measured. | Copy caps for a body section: eyebrow at most 3 words, H2 at most 8 words, lead at most 28 words, card body at most 32 words. The H1 is at m |
| unverified | `25-page-assembly.md` | 25 | `33` | No source row yet. Do not treat this number as measured. | Menu components and their entrance timings are `33-mega-menu-components.md`. Buttons are `04-white-blue-theme/03-buttons-motion-and-interact |
| unverified | `25-page-assembly.md` | 25 | `04` | No source row yet. Do not treat this number as measured. | Menu components and their entrance timings are `33-mega-menu-components.md`. Buttons are `04-white-blue-theme/03-buttons-motion-and-interact |
| unverified | `25-page-assembly.md` | 25 | `33` | No source row yet. Do not treat this number as measured. | Menu components and their entrance timings are `33-mega-menu-components.md`. Buttons are `04-white-blue-theme/03-buttons-motion-and-interact |
| unverified | `25-page-assembly.md` | 29 | `1` | No source row yet. Do not treat this number as measured. | ## 1. Shell |
| unverified | `25-page-assembly.md` | 33 | `1` | No source row yet. Do not treat this number as measured. | 1. `SiteHeader` from `04-white-blue-theme/02-header-mega-menu-and-footer.md` |
| unverified | `25-page-assembly.md` | 33 | `04` | No source row yet. Do not treat this number as measured. | 1. `SiteHeader` from `04-white-blue-theme/02-header-mega-menu-and-footer.md` |
| unverified | `25-page-assembly.md` | 34 | `2` | No source row yet. Do not treat this number as measured. | 2. Exactly one hero. Default is the light-first split hero (`CapabilityStack`) from `04-cards-heroes-and-section-library.md` section 3. |
| unverified | `25-page-assembly.md` | 34 | `04` | No source row yet. Do not treat this number as measured. | 2. Exactly one hero. Default is the light-first split hero (`CapabilityStack`) from `04-cards-heroes-and-section-library.md` section 3. |
| unverified | `25-page-assembly.md` | 34 | `3` | No source row yet. Do not treat this number as measured. | 2. Exactly one hero. Default is the light-first split hero (`CapabilityStack`) from `04-cards-heroes-and-section-library.md` section 3. |
| unverified | `25-page-assembly.md` | 35 | `3` | No source row yet. Do not treat this number as measured. | 3. Three to seven body sections, chosen only from section 3 of this file. |
| unverified | `25-page-assembly.md` | 35 | `3` | No source row yet. Do not treat this number as measured. | 3. Three to seven body sections, chosen only from section 3 of this file. |
| unverified | `25-page-assembly.md` | 36 | `4` | No source row yet. Do not treat this number as measured. | 4. One closer. Default is the lead form from `04-cards-heroes-and-section-library.md` section 7.3, or a final band that reuses an allowed se |
| unverified | `25-page-assembly.md` | 36 | `04` | No source row yet. Do not treat this number as measured. | 4. One closer. Default is the lead form from `04-cards-heroes-and-section-library.md` section 7.3, or a final band that reuses an allowed se |
| unverified | `25-page-assembly.md` | 36 | `7.3` | No source row yet. Do not treat this number as measured. | 4. One closer. Default is the lead form from `04-cards-heroes-and-section-library.md` section 7.3, or a final band that reuses an allowed se |
| unverified | `25-page-assembly.md` | 37 | `5` | No source row yet. Do not treat this number as measured. | 5. Footer from `04-white-blue-theme/02-header-mega-menu-and-footer.md` |
| unverified | `25-page-assembly.md` | 37 | `04` | No source row yet. Do not treat this number as measured. | 5. Footer from `04-white-blue-theme/02-header-mega-menu-and-footer.md` |
| unverified | `25-page-assembly.md` | 43 | `2` | No source row yet. Do not treat this number as measured. | ## 2. Band rhythm |
| unverified | `25-page-assembly.md` | 49 | `3` | No source row yet. Do not treat this number as measured. | ## 3. Allowed body sections |
| unverified | `25-page-assembly.md` | 55 | `04` | No source row yet. Do not treat this number as measured. | / `capability-stack` / Split hero with a self-assembling stack / `04-cards-heroes-and-section-library.md` section 3 / |
| unverified | `25-page-assembly.md` | 55 | `3` | No source row yet. Do not treat this number as measured. | / `capability-stack` / Split hero with a self-assembling stack / `04-cards-heroes-and-section-library.md` section 3 / |
| unverified | `25-page-assembly.md` | 56 | `4` | No source row yet. Do not treat this number as measured. | / `solutions-grid` / Flagship card grid under a sticky editorial header / same file, section 4 / |
| unverified | `25-page-assembly.md` | 57 | `5` | No source row yet. Do not treat this number as measured. | / `workflow-board` / Pinned sticky-note process / same file, section 5 / |
| unverified | `25-page-assembly.md` | 58 | `6` | No source row yet. Do not treat this number as measured. | / `scroll-stack` / Fluted glass sticky card deck / same file, section 6 / |
| unverified | `25-page-assembly.md` | 59 | `7.1` | No source row yet. Do not treat this number as measured. | / `capability-tabs` / Tabbed capability board on `.band-soft` / same file, section 7.1 / |
| unverified | `25-page-assembly.md` | 60 | `7.2` | No source row yet. Do not treat this number as measured. | / `pricing` / Scope-builder pricing / same file, section 7.2 / |
| unverified | `25-page-assembly.md` | 61 | `7.3` | No source row yet. Do not treat this number as measured. | / `lead-form` / Lead form card / same file, section 7.3 / |
| unverified | `25-page-assembly.md` | 62 | `3` | No source row yet. Do not treat this number as measured. | / `hero-page` / Breadcrumb, eyebrow, one accent word, lead, two CTAs, optional 3-item proof / same hero type scale / |
| unverified | `25-page-assembly.md` | 63 | `33` | No source row yet. Do not treat this number as measured. | / `hero-split` / Copy column plus one media card (`TiltCard`) / file 33, max tilt `8deg` / |
| unverified | `25-page-assembly.md` | 63 | `8` | No source row yet. Do not treat this number as measured. | / `hero-split` / Copy column plus one media card (`TiltCard`) / file 33, max tilt `8deg` / |
| unverified | `25-page-assembly.md` | 67 | `33` | No source row yet. Do not treat this number as measured. | / `band-stats` / Copy, CTA, `CountUp` grid / file 33 / |
| unverified | `25-page-assembly.md` | 72 | `2` | No source row yet. Do not treat this number as measured. | / `grid-cards` / 2 or 3 columns, icon, title, body / `card-premium` / |
| unverified | `25-page-assembly.md` | 72 | `3` | No source row yet. Do not treat this number as measured. | / `grid-cards` / 2 or 3 columns, icon, title, body / `card-premium` / |
| unverified | `25-page-assembly.md` | 75 | `720px` | No source row yet. Do not treat this number as measured. | / `prose-body` / Long form on a `720px` measure / blog article / |
| unverified | `25-page-assembly.md` | 78 | `2px` | No source row yet. Do not treat this number as measured. | Card primitives allowed inside those sections: `SurfaceCard`, `GlassCard`, `NeuCard`, `Pill`, `card-premium`, `row-premium`. Hover on `card- |
| unverified | `25-page-assembly.md` | 78 | `0` | No source row yet. Do not treat this number as measured. | Card primitives allowed inside those sections: `SurfaceCard`, `GlassCard`, `NeuCard`, `Pill`, `card-premium`, `row-premium`. Hover on `card- |
| unverified | `25-page-assembly.md` | 78 | `1` | No source row yet. Do not treat this number as measured. | Card primitives allowed inside those sections: `SurfaceCard`, `GlassCard`, `NeuCard`, `Pill`, `card-premium`, `row-premium`. Hover on `card- |
| unverified | `25-page-assembly.md` | 78 | `1` | No source row yet. Do not treat this number as measured. | Card primitives allowed inside those sections: `SurfaceCard`, `GlassCard`, `NeuCard`, `Pill`, `card-premium`, `row-premium`. Hover on `card- |
| unverified | `25-page-assembly.md` | 82 | `4` | No source row yet. Do not treat this number as measured. | ## 4. Page kinds |
| unverified | `25-page-assembly.md` | 84 | `4.1` | No source row yet. Do not treat this number as measured. | ### 4.1 Marketing page |
| unverified | `25-page-assembly.md` | 86 | `33` | No source row yet. Do not treat this number as measured. | `SiteHeader` built from `33-mega-menu-components.md`, one hero (`capability-stack`, `hero-page`, `hero-split`, or `hero-form`), then three t |
| unverified | `25-page-assembly.md` | 86 | `3` | No source row yet. Do not treat this number as measured. | `SiteHeader` built from `33-mega-menu-components.md`, one hero (`capability-stack`, `hero-page`, `hero-split`, or `hero-form`), then three t |
| unverified | `25-page-assembly.md` | 88 | `4.2` | No source row yet. Do not treat this number as measured. | ### 4.2 Blog index |
| unverified | `25-page-assembly.md` | 90 | `02` | No source row yet. Do not treat this number as measured. | `SiteHeader`, a short hero that reuses the split-hero type scale (one `text-h1`, one lead), then a `solutions-grid` of article cards (title, |
| unverified | `25-page-assembly.md` | 92 | `4.3` | No source row yet. Do not treat this number as measured. | ### 4.3 Blog article |
| unverified | `25-page-assembly.md` | 94 | `9` | No source row yet. Do not treat this number as measured. | `SiteHeader`, one editorial hero (title at most 9 words, eyebrow in `text-eyebrow` at `12px` mono), then a prose column on `--paper` using b |
| unverified | `25-page-assembly.md` | 94 | `12px` | No source row yet. Do not treat this number as measured. | `SiteHeader`, one editorial hero (title at most 9 words, eyebrow in `text-eyebrow` at `12px` mono), then a prose column on `--paper` using b |
| unverified | `25-page-assembly.md` | 94 | `01` | No source row yet. Do not treat this number as measured. | `SiteHeader`, one editorial hero (title at most 9 words, eyebrow in `text-eyebrow` at `12px` mono), then a prose column on `--paper` using b |
| unverified | `25-page-assembly.md` | 98 | `5` | No source row yet. Do not treat this number as measured. | ## 5. Menu groups |
| unverified | `25-page-assembly.md` | 102 | `1` | No source row yet. Do not treat this number as measured. | 1. Product |
| unverified | `25-page-assembly.md` | 103 | `2` | No source row yet. Do not treat this number as measured. | 2. Industries |
| unverified | `25-page-assembly.md` | 104 | `3` | No source row yet. Do not treat this number as measured. | 3. Services |
| unverified | `25-page-assembly.md` | 105 | `4` | No source row yet. Do not treat this number as measured. | 4. About |
| unverified | `25-page-assembly.md` | 106 | `5` | No source row yet. Do not treat this number as measured. | 5. Resources |
| unverified | `25-page-assembly.md` | 108 | `02` | No source row yet. Do not treat this number as measured. | Contact and legal links live in the footer and in the header CTA pair, not as extra mega columns. Each mega panel follows `02-header-mega-me |
| unverified | `25-page-assembly.md` | 108 | `0.04` | No source row yet. Do not treat this number as measured. | Contact and legal links live in the footer and in the header CTA pair, not as extra mega columns. Each mega panel follows `02-header-mega-me |
| unverified | `25-page-assembly.md` | 108 | `0` | No source row yet. Do not treat this number as measured. | Contact and legal links live in the footer and in the header CTA pair, not as extra mega columns. Each mega panel follows `02-header-mega-me |
| unverified | `25-page-assembly.md` | 108 | `100` | No source row yet. Do not treat this number as measured. | Contact and legal links live in the footer and in the header CTA pair, not as extra mega columns. Each mega panel follows `02-header-mega-me |
| unverified | `25-page-assembly.md` | 108 | `14` | No source row yet. Do not treat this number as measured. | Contact and legal links live in the footer and in the header CTA pair, not as extra mega columns. Each mega panel follows `02-header-mega-me |
| unverified | `25-page-assembly.md` | 108 | `110ms` | No source row yet. Do not treat this number as measured. | Contact and legal links live in the footer and in the header CTA pair, not as extra mega columns. Each mega panel follows `02-header-mega-me |
| unverified | `25-page-assembly.md` | 108 | `220ms` | No source row yet. Do not treat this number as measured. | Contact and legal links live in the footer and in the header CTA pair, not as extra mega columns. Each mega panel follows `02-header-mega-me |
| unverified | `25-page-assembly.md` | 108 | `180` | No source row yet. Do not treat this number as measured. | Contact and legal links live in the footer and in the header CTA pair, not as extra mega columns. Each mega panel follows `02-header-mega-me |
| unverified | `25-page-assembly.md` | 108 | `820ms` | No source row yet. Do not treat this number as measured. | Contact and legal links live in the footer and in the header CTA pair, not as extra mega columns. Each mega panel follows `02-header-mega-me |
| unverified | `25-page-assembly.md` | 108 | `1400px` | No source row yet. Do not treat this number as measured. | Contact and legal links live in the footer and in the header CTA pair, not as extra mega columns. Each mega panel follows `02-header-mega-me |
| unverified | `25-page-assembly.md` | 112 | `3` | No source row yet. Do not treat this number as measured. | - 3 or more groups: `lg:grid-cols-[1fr_1fr_1fr_0.9fr]` |
| unverified | `25-page-assembly.md` | 112 | `1` | No source row yet. Do not treat this number as measured. | - 3 or more groups: `lg:grid-cols-[1fr_1fr_1fr_0.9fr]` |
| unverified | `25-page-assembly.md` | 113 | `2` | No source row yet. Do not treat this number as measured. | - 2 groups: `lg:grid-cols-[1fr_1fr_1.1fr]` |
| unverified | `25-page-assembly.md` | 113 | `1` | No source row yet. Do not treat this number as measured. | - 2 groups: `lg:grid-cols-[1fr_1fr_1.1fr]` |
| unverified | `25-page-assembly.md` | 114 | `1` | No source row yet. Do not treat this number as measured. | - 1 group: `lg:grid-cols-[1.6fr_1fr]` |
| unverified | `25-page-assembly.md` | 114 | `1.6` | No source row yet. Do not treat this number as measured. | - 1 group: `lg:grid-cols-[1.6fr_1fr]` |
| unverified | `25-page-assembly.md` | 118 | `6` | No source row yet. Do not treat this number as measured. | ## 6. Sibling References |
| unverified | `25-page-assembly.md` | 120 | `37` | No source row yet. Do not treat this number as measured. | - Standalone marketing images, banners, and thumbnails: [`37-image-specifications.md`](./37-image-specifications.md) |
| unverified | `25-page-assembly.md` | 121 | `3` | No source row yet. Do not treat this number as measured. | - Detailed 3-tier pricing table and floating editorial card components: [`38-card-and-pricing-components.md`](./38-card-and-pricing-componen |
| unverified | `25-page-assembly.md` | 121 | `38` | No source row yet. Do not treat this number as measured. | - Detailed 3-tier pricing table and floating editorial card components: [`38-card-and-pricing-components.md`](./38-card-and-pricing-componen |
| unverified | `25-page-assembly.md` | 125 | `7` | No source row yet. Do not treat this number as measured. | ## 7. Refusal list |
| unverified | `25-page-assembly.md` | 127 | `3` | No source row yet. Do not treat this number as measured. | - Do not add a section id that is not in section 3. |
| unverified | `25-page-assembly.md` | 128 | `33` | No source row yet. Do not treat this number as measured. | - Menu parts that are not in `33-mega-menu-components.md` do not exist. |
| unverified | `25-page-assembly.md` | 130 | `72px` | No source row yet. Do not treat this number as measured. | - Do not change header height (`72px`), scroll threshold (`12px`), or menu motion. |
| unverified | `25-page-assembly.md` | 130 | `12px` | No source row yet. Do not treat this number as measured. | - Do not change header height (`72px`), scroll threshold (`12px`), or menu motion. |
| unverified | `26-visual-builder.md` | 1 | `26` | No source row yet. Do not treat this number as measured. | # 26 — Visual Builder Overlay |
| unverified | `26-visual-builder.md` | 4 | `29` | No source row yet. Do not treat this number as measured. | > **/learn** Gate, modes, sanitizer, storage bucket, and export. This file is not the slide builder. Slides use `29-slide-navigation-and-bui |
| unverified | `26-visual-builder.md` | 6 | `4.3` | No source row yet. Do not treat this number as measured. | **Version:** 4.3.0 |
| unverified | `26-visual-builder.md` | 7 | `2026` | No source row yet. Do not treat this number as measured. | **Updated:** 2026-10-02 |
| unverified | `26-visual-builder.md` | 7 | `10` | No source row yet. Do not treat this number as measured. | **Updated:** 2026-10-02 |
| unverified | `26-visual-builder.md` | 7 | `02` | No source row yet. Do not treat this number as measured. | **Updated:** 2026-10-02 |
| unverified | `26-visual-builder.md` | 12 | `0` | No source row yet. Do not treat this number as measured. | ## 0. Anti-hallucination |
| unverified | `26-visual-builder.md` | 20 | `1` | No source row yet. Do not treat this number as measured. | ## 1. Gate |
| unverified | `26-visual-builder.md` | 25 | `1` | No source row yet. Do not treat this number as measured. | {SITE_DOMAIN}/?builder=1&email={OWNER_EMAIL} |
| unverified | `26-visual-builder.md` | 32 | `2` | No source row yet. Do not treat this number as measured. | ## 2. Modes |
| unverified | `26-visual-builder.md` | 50 | `3` | No source row yet. Do not treat this number as measured. | ## 3. Text |
| verified | `26-visual-builder.md` | 68 | `80` | builder kit 05: slugify truncation | `slugify` is lowercase, NFKD, non-alphanumerics collapsed to `-`, trimmed, truncated to 80 characters. Do not change 80 once ids are stored. |
| verified | `26-visual-builder.md` | 68 | `80` | builder kit 05: slugify truncation | `slugify` is lowercase, NFKD, non-alphanumerics collapsed to `-`, trimmed, truncated to 80 characters. Do not change 80 once ids are stored. |
| verified | `26-visual-builder.md` | 72 | `10px` | builder kit 05: touch move threshold | A touch counts as a tap only when the finger moves less than `10px` and lifts within `600ms`. |
| verified | `26-visual-builder.md` | 72 | `600ms` | builder kit 05: touch lift threshold | A touch counts as a tap only when the finger moves less than `10px` and lifts within `600ms`. |
| unverified | `26-visual-builder.md` | 76 | `4` | No source row yet. Do not treat this number as measured. | ## 4. Images and icons |
| verified | `26-visual-builder.md` | 78 | `2` | builder kit 06: image cap | Accepted types, exactly: `image/png`, `image/jpeg`, `image/webp`, `image/svg+xml`. Reject AVIF. Hard-reject any file over `2 MB`. Do not war |
| verified | `26-visual-builder.md` | 85 | `200` | builder kit 06: custom SVG cap | - Custom SVG max `200 KB`. Parse with `DOMParser` in the SVG namespace. Keep `viewBox`, or fall back to `0 0 24 24`. Recreate children with  |
| unverified | `26-visual-builder.md` | 85 | `0` | No source row yet. Do not treat this number as measured. | - Custom SVG max `200 KB`. Parse with `DOMParser` in the SVG namespace. Keep `viewBox`, or fall back to `0 0 24 24`. Recreate children with  |
| unverified | `26-visual-builder.md` | 85 | `0` | No source row yet. Do not treat this number as measured. | - Custom SVG max `200 KB`. Parse with `DOMParser` in the SVG namespace. Keep `viewBox`, or fall back to `0 0 24 24`. Recreate children with  |
| verified | `26-visual-builder.md` | 85 | `24` | builder kit 07: token floor | - Custom SVG max `200 KB`. Parse with `DOMParser` in the SVG namespace. Keep `viewBox`, or fall back to `0 0 24 24`. Recreate children with  |
| verified | `26-visual-builder.md` | 85 | `24` | builder kit 07: token floor | - Custom SVG max `200 KB`. Parse with `DOMParser` in the SVG namespace. Keep `viewBox`, or fall back to `0 0 24 24`. Recreate children with  |
| unverified | `26-visual-builder.md` | 90 | `5` | No source row yet. Do not treat this number as measured. | ## 5. Menu |
| unverified | `26-visual-builder.md` | 110 | `6` | No source row yet. Do not treat this number as measured. | ## 6. Layout |
| unverified | `26-visual-builder.md` | 118 | `01` | No source row yet. Do not treat this number as measured. | - After each move, recompute index, inset, and `z-index` on stacked slots, and renumber visible `01` / `02` / `03` labels in the same frame. |
| unverified | `26-visual-builder.md` | 118 | `02` | No source row yet. Do not treat this number as measured. | - After each move, recompute index, inset, and `z-index` on stacked slots, and renumber visible `01` / `02` / `03` labels in the same frame. |
| unverified | `26-visual-builder.md` | 118 | `03` | No source row yet. Do not treat this number as measured. | - After each move, recompute index, inset, and `z-index` on stacked slots, and renumber visible `01` / `02` / `03` labels in the same frame. |
