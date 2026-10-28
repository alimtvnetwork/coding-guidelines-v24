# 100 Design System Improvements

**Updated:** 28-Oct-26
**Ids:** `DS-IMP-001` through `DS-IMP-200`
**Next run:** append `DS-IMP-201` onward. Never renumber. Never delete a row.

Each row is a check a weak model must pass. Source paths are relative inside this repository unless the path starts with a sibling repo name.

| Id | Do | Do not |
| :--- | :--- | :--- |
| DS-IMP-001 | Name the job before painting: slide, amber deck, website system, or tree. | Start from `slides-app/src/styles/tokens.css` for a gold deck. |
| DS-IMP-002 | Default unnamed slides to `data-theme="bright-gold-tech"`. | Invent a ninth stage theme. |
| DS-IMP-003 | Set the theme attribute on `.slide-stage` only. | Set `white-blue` on the stage. |
| DS-IMP-004 | Use the six bright-gold tokens from file `05-bright-gold-tech/01-tokens-background-and-chrome.md`. | Copy `#0A0A14` or `#FFAD01` into new CSS. |
| DS-IMP-005 | Keep gold as the only accent. | Add ember as a second brand color. |
| DS-IMP-006 | Read stage colors from `var(--canvas)` and its five siblings. | Hardcode hex inside a slide component. |
| DS-IMP-007 | When the user says dark amber, switch the whole deck to file 30. | Mix `#F5A623` onto `#08080C`. |
| DS-IMP-008 | Treat file 27's ten theme ids as catalogue history. | Put `vscode-dark` in the stage switcher. |
| DS-IMP-009 | Store file-40 colors as full `hsl()` or hex. | Wrap those values in `hsl()` a second time. |
| DS-IMP-010 | Keep `botanical-light` as a quiz-embed skin only. | Offer it as a keynote stage theme. |
| DS-IMP-011 | Author coordinates on 1920 by 1080. | Let type wrap because the window changed. |
| DS-IMP-012 | Scale with `min(viewport/1920, viewport/1080)`. | Use a percentage width for the cover title. |
| DS-IMP-013 | Keep the 96px side safe area and 80px vertical safe area. | Park the logo inside the text inset. |
| DS-IMP-014 | Set the cover title at 128px Ubuntu 700, line-height 1.05, tracking `-0.02em`. | Use the website `clamp()` hero on a slide. |
| DS-IMP-015 | Set bullets at 28px Poppins, line-height 1.5. | Drop body copy to 14px on a projected deck. |
| DS-IMP-016 | Use Ubuntu for display and Poppins for body on slides. | Use Inter on a slide or on White Blue. |
| DS-IMP-017 | Keep the headline as live HTML. | Bake the title into an image. |
| DS-IMP-018 | Ground the slide in near-black plus one amber spotlight. | Use a purple-blue full-screen gradient. |
| DS-IMP-019 | Optional grid is 48px at 2.5 percent white. | Draw a loud graph-paper grid. |
| DS-IMP-020 | Footer index is `{current} / {total}` in muted ink. | Center a row of colorful dots as the only chrome. |
| DS-IMP-021 | Entrance is `gold-rise` at 0.75s with the documented curve. | Fade the whole stage in with `opacity` only and no rise. |
| DS-IMP-022 | Stagger siblings by 0.22s. | Animate every block with the same delay. |
| DS-IMP-023 | Crossfade a theme change in 0.45s. | Swap theme colors in one frame. |
| DS-IMP-024 | Cap ordinary hover at 0.3s. | Use a 700ms hover on a button. |
| DS-IMP-025 | Ordinary hover lifts `-2px` and animates transform plus opacity. | Lift website cards `-4px` on a slide. |
| DS-IMP-026 | Recommendation-card hover may use `scale(1.06) translateY(-8px)` at 0.3s. | Apply that scale to every card. |
| DS-IMP-027 | Resting dark surfaces use a hairline, about 8 to 16 percent white. | Use a heavy black drop shadow on a dark card at rest. |
| DS-IMP-028 | Gold glow is for hover and the spotlight. | Leave a gold glow on at rest. |
| DS-IMP-029 | Honor `prefers-reduced-motion` with 0.01ms duration. | Delete the state change when motion is reduced. |
| DS-IMP-030 | Progress width eases in 240ms `ease-out`. | Jump the progress bar. |
| DS-IMP-031 | Glass cards use radius 24px, 1px border at 8 percent accent, blur 20px. | Use a 4px solid gold border on every card. |
| DS-IMP-032 | Accent bar is 4 by 72px, gold to `#FFC64D`. | Stretch a full-width gold bar under the title. |
| DS-IMP-033 | Pills stay outlined on a definition slide. | Fill every pill with gold. |
| DS-IMP-034 | One shouting element per slide. | Bold every line. |
| DS-IMP-035 | Hold 60/30/10: field, support, accent. | Give two saturated accents equal area. |
| DS-IMP-036 | Copy states a number or a concrete claim. | Fill the slide with adjectives. |
| DS-IMP-037 | Decorative icons stay 56 to 96px, stroke 1.5px, opacity 0.08 to 0.12. | Center a 3-icon feature row on a cover. |
| DS-IMP-038 | Keep decorative marks out of the center 900 by 400 on a cover. | Overlay icons on the title. |
| DS-IMP-039 | Logo height is 48px, inset 60px and 36px from the stage edge. | Scale the logo with the title. |
| DS-IMP-040 | Section bands on websites alternate light, soft, and dark. | Paint every band the same white. |
| DS-IMP-041 | AI SaaS pink is `#FF2D6F`, hover `#E02663`, page `#F7F4F3`. | Recolor that system with slide gold. |
| DS-IMP-042 | AI SaaS motion tokens stay 120, 200, 300, and 450ms. | Invent a 1s hover. |
| DS-IMP-043 | AI SaaS card hover is `translateY(-4px)` in 250ms. | Also scale the card. |
| DS-IMP-044 | Pressed controls use `scale(0.98)`. | Use `scale(1.05)` as the hover. |
| DS-IMP-045 | Sweet Digs nav underline grows from 0 to 75 percent in 300ms. | Zoom the link. |
| DS-IMP-046 | Sweet Digs card lift is `translateY(-6px)` with a tinted glow. | Use a generic gray shadow. |
| DS-IMP-047 | Sweet Digs Join control slides text in 350ms. | Swap the label with no motion. |
| DS-IMP-048 | White Blue content max is 1280px. | Stretch lines past that measure. |
| DS-IMP-049 | White Blue section pad is `clamp(72px, 8vw, 120px)`. | Use 24px section padding on a marketing page. |
| DS-IMP-050 | White Blue outline buttons fill from the bottom in 240ms. | Jump the fill. |
| DS-IMP-051 | White Blue press moves `translateY(1px)`. | Bounce the button. |
| DS-IMP-052 | Shine sweep runs 3s and only while hovered. | Autoplay shine on every card. |
| DS-IMP-053 | Website JSON palettes stay in `tokens/design-tokens.json`. | Feed those hex values to the slide stage. |
| DS-IMP-054 | Docs prose type stays in `02-spec/07-design-system/04-typography.md`. | Use the 128px cover scale in a docs article. |
| DS-IMP-055 | Choice rows slide `translate3d(6px, 0, 0)` in 220ms. | Lift quiz options upward. |
| DS-IMP-056 | Choice-row rest opacity starts near 0.82 and clarifies on hover. | Start every option at full contrast and full gold. |
| DS-IMP-057 | Selection rings use `box-shadow: 0 0 0 2px`. | Change border width and shift layout. |
| DS-IMP-058 | Stagger list entrance by about 60ms. | Delay items by 400ms. |
| DS-IMP-059 | Theme apply may dip opacity for 180ms in and 280ms out. | Flash a white frame. |
| DS-IMP-060 | Keep motion easings to editorial `cubic-bezier(0.16, 1, 0.3, 1)` and state `cubic-bezier(0.4, 0, 0.2, 1)`. | Add a third bounce curve. |
| DS-IMP-061 | Tree type is about 12px with 2px vertical padding. | Use 16px rows in a compact tree. |
| DS-IMP-062 | Indent one level by about 10px plus a 1px rule. | Indent 24px per level. |
| DS-IMP-063 | Draw branch glyphs on nested rows. | Rely on indent alone. |
| DS-IMP-064 | Swap chevron icons. Do not rotate them. | Animate `rotate(90deg)` on tree chevrons. |
| DS-IMP-065 | Show a child chevron only when the node has children. | Show a chevron on every row. |
| DS-IMP-066 | Hover is one wash at every depth. | Hover a child with the parent's selected color. |
| DS-IMP-067 | Selection color changes by role: root, branch, leaf, nested agent. | Use one blue bar for every selected row. |
| DS-IMP-068 | Do not put a focus ring on ordinary tree rows. | Outline every row on click. |
| DS-IMP-069 | Nested-agent rows may use a 1px focus ring. | Reuse that ring on branch rows. |
| DS-IMP-070 | Keep a single density. | Add a compact/comfortable toggle unless the user asks. |
| DS-IMP-071 | Default tree pane is 340px, clamped 280 to 760, and persisted. | Fix the pane at 240px. |
| DS-IMP-072 | Collapse to a 40px rail, not width 0. | Hide the tree with `display: none` and no way back. |
| DS-IMP-073 | Bind the right pane to the selected leaf, not the root. | Show root metadata when a leaf is selected. |
| DS-IMP-074 | Selecting a leaf resolves one child record and shows its text. | Leave the panel on the previous leaf. |
| DS-IMP-075 | Empty tree uses a skeleton. | Show the "select an item" sentence when the tree itself is empty. |
| DS-IMP-076 | Empty selection, tree present, uses one sentence that names the leaf. | Use a blank white pane. |
| DS-IMP-077 | Cap long branches (the reference cap is 50) unless scope is one root. | Render thousands of rows with no cap. |
| DS-IMP-078 | Arrow keys walk expanded rows and wrap. | Move focus only by pointer. |
| DS-IMP-079 | Child chevron clicks do not select the parent. | Toggle and select on the same target. |
| DS-IMP-080 | Right pane is `flex-1` beside the tree. | Stack the detail under the tree on a desktop width. |
| DS-IMP-081 | Header of the pane shows identity, state, and the primary actions. | Hide actions in a second modal. |
| DS-IMP-082 | Offer preview, raw, and edit as tabs. | Dump raw and preview in one scrolling column. |
| DS-IMP-083 | Long text previews at about 200 words with a word count. | Paste unbounded text into the header. |
| DS-IMP-084 | Running state is a token dot, reference `#1af18d`, not a banner. | Recolor the whole row green while running. |
| DS-IMP-085 | Map product hex to role tokens. Do not leave `bg-blue-100` as the spec. | Copy Tailwind class strings into a second product unchanged. |
| DS-IMP-086 | Do not train `--ui-surface-hover` as the row hover if the rows do not use it. | Assume the token file and the rows match. |
| DS-IMP-087 | Sidebar width animation, when used, animates width with the editorial ease. | Animate `margin-left` and jump the content. |
| DS-IMP-088 | Feature bands use one hairline, one icon well, 16px title, 12px body. | Use three competing card styles in one band. |
| DS-IMP-089 | Floating tools sit inset from the corner on a blurred surface. | Pin them flush to the viewport edge. |
| DS-IMP-090 | Preview chrome is a 2px primary hairline and ignores pointer events. | Cover the slide with a badge. |
| DS-IMP-091 | Follow `follow-ui-ux-design-system` and `write-and-enhance-design-spec`. | Restore `.agents/skills-archived/create-slide-deck/`. |
| DS-IMP-092 | Leave `02-spec/24-app-ui-design-system/` as a stub that points here. | Duplicate this contract into that empty folder. |
| DS-IMP-093 | Cite `02-spec/07-design-system/20-ai-training-and-checklist-guide.md` for cheap-versus-expensive. | Replace that file's 60/30/10 rule. |
| DS-IMP-094 | When repeating this prompt, append new ids after the last `DS-IMP`. | Rewrite 001 to 100. |
| DS-IMP-095 | Search with `gitmap aum search`, `gitmap find`, and `gitmap cat`. | Search with ripgrep or Select-String. |
| DS-IMP-096 | Write only relative paths. | Write a drive letter or a `file:///` URI. |
| DS-IMP-097 | Record a missing sibling as missing. | Invent Let's Mark components. `letsmarknow` had no UI source on 28-Oct-26. |
| DS-IMP-098 | Keep slide implementation review-first. | Recolor `slides-app` in the same turn as a spec-only request. |
| DS-IMP-099 | Update `02-spec/21-app/readme.md` and `.ai-memory/plans/readme.md` when this folder changes. | Leave the new folder unindexed. |
| DS-IMP-100 | End a spec run with one hyphen GitMap commit. No version bump unless the user asks. | Cut a release because the spec changed. |

## Second pass — DS-IMP-101 through DS-IMP-200

These rows do not replace 001–100. Where two source files disagree, the row says which file wins. Do not average the numbers.

| Id | Do | Do not | Source |
| :--- | :--- | :--- | :--- |
| DS-IMP-101 | Place eyebrow, title, and content with `justify-content` on the outer section. Tighten the header with `mb-6`. | Push the group with `pt-*` or `mt-*` on the children, or put `flex-1` on the grid. | `02-spec/07-design-system/28-slide-layouts.md` |
| DS-IMP-102 | Set horizontal brand alignment with `paddingLeft` and `paddingRight` of `var(--brand-inset-x)`. | Align the slide to the logo with Tailwind `px-24`. | `02-spec/07-design-system/28-slide-layouts.md` |
| DS-IMP-103 | When `compact` is true, start cards at the top and set those grids to `grid-auto-rows: min-content` with a `1rem` row gap. | Leave `grid-auto-rows: 1fr` so two cards each take half the height. | `02-spec/07-design-system/28-slide-layouts.md` |
| DS-IMP-104 | On the two-column content slide, start the copy at x 96, cap it at 1100px, use an 88px headline, and draw the chart at x 1180, y 280, 640 by 520, bars 72px wide, gap 48px, radius `8px 8px 0 0`. | Drop the chart origin or the 22px cover subtitle. | `02-spec/07-design-system/28-slide-layouts.md` |
| DS-IMP-105 | Stay inside the closed 17 layout ids: `left`, `center`, `steps`, `timeline`, `process`, `quote`, `bullets`, `image`, `poll`, `qa`, `embed`, `reveal-grid`, `counter-stat`, `typewriter`, `compare`, `priority`, `depth-stack`. | Invent an 18th id, or put a private client name in slide copy. | `02-spec/07-design-system/28-slide-layouts.md` |
| DS-IMP-106 | Mount the progress HUD on the viewport at top 32px, right 32px, height 56px, radius 9999px. Active mark is 28 by 8px. Inactive mark is 8 by 8px. File 29 sets `z-index: 50`. | Use `top-6 right-6`, or mount the HUD inside the 1920 by 1080 transform. | `02-spec/07-design-system/29-slide-navigation-and-builder.md` |
| DS-IMP-107 | Cycle one scripted family: `Slide`, `Fade`, `Zoom`, `Flip`, `Rise`. Duration 0.45s, ease `cubic-bezier(0.4, 0, 0.2, 1)`. Slide uses x 100% to -100%. Zoom uses scale 0.85 to 1.15. Flip uses rotateY ±45deg at perspective 1200px. Rise uses y ±80px. Backward flips the signs. | Add a sixth name, or mix this family with the JSON family on the next slide. | `02-spec/07-design-system/29-slide-navigation-and-builder.md` |
| DS-IMP-108 | Limit JSON `TransitionKind` to `fade`, `camera-zoom`, and `slide`. The settings UI shows `fade` and `camera-zoom` only. Default is `fade`. | Show `slide` as a third settings choice unless that file is revised. | `02-spec/07-design-system/29-slide-navigation-and-builder.md` |
| DS-IMP-109 | Advance on a swipe only after 50px. `ArrowRight` or `Space` steps, then advances. `G` opens the grid. `S` opens settings. `/` opens shortcuts. Double-click the counter to jump. | Advance on a shorter drag, or invent a sound file so `M` has something to mute. | `02-spec/07-design-system/29-slide-navigation-and-builder.md` |
| DS-IMP-110 | Capture at scale 1 or 2. JPEG quality 0.92. PNG quality 0.95. Fail a frame when the mean is at or under 14 and the deviation is at or under 8. Settle 950ms, poll 50ms, give up at 5000ms, 3 retries. Read `var(--canvas)` from the active theme. | Ship the blank 64 by 36 sample, or paint the capture with a fixed hex. | `02-spec/07-design-system/29-slide-navigation-and-builder.md` |
| DS-IMP-111 | Draw the controller pill at height 56px, padding 8px 12px, blur 12px, shadow `0 8px 24px hsl(0 0% 0% / 0.40)`, and portal it to `document.body`. File 31 sets this pill at `z-index: 60`. | Leave the pill inside a stage that has `overflow: hidden`. | `02-spec/07-design-system/31-slide-controller-buttons.md` |
| DS-IMP-112 | Persist `ControllerPosition` in `ctrl.position.v1` across the eight anchors. `TopRight` uses 32px plus the safe-area inset. The other seven use 24px plus the matching inset. Bottom anchors open tooltips upward. | Drop `env(safe-area-inset-*)`, or open a bottom menu downward. | `02-spec/07-design-system/31-slide-controller-buttons.md` |
| DS-IMP-113 | After 2500ms idle, fade the pill to opacity 0.15, or to 0 in presenter mode. Restore opacity 1 in 150ms `ease-out` on move, key, or hover. Suppress auto-hide while a popover is open. | Keep the pill fully opaque through an idle talk. | `02-spec/07-design-system/31-slide-controller-buttons.md` |
| DS-IMP-114 | Use a 40 by 40px hit target. Hover wash is `hsl(0 0% 100% / 0.08)`. Active wash is `0.14`. Disabled chevrons are opacity 0.35. The counter is Poppins 500 at 20px, tabular, min-width 64px, format `{current} / {total}`. Share writes `#slide-{current}` and a 2-second toast. | Shrink the hit target, or omit the slide hash. | `02-spec/07-design-system/31-slide-controller-buttons.md` |
| DS-IMP-115 | Synthesize chimes in Web Audio. Tick is 800Hz sine, gain 0.08, decay 30ms. Chord is 523.25Hz and 659.25Hz, gain 0.06, 120ms. First-run key `ctrl.onboarded.v1` stores `"1"`. Reduced motion snaps dot size with no interpolation. | Download an external sound file for the tick. | `02-spec/07-design-system/31-slide-controller-buttons.md` |
| DS-IMP-116 | Build every ramp as `S0` through `S9` with `t = i/9`. Royal Violet anchors are `S0` `#101070`, `S4` `#7C3AED`, `S5` `#935BF5`, `S9` `#FAF5FF`. | Interpolate a gradient with an unnamed in-between stop. | `02-spec/07-design-system/32-slide-color-options.md` |
| DS-IMP-117 | Keep Corporate Gold `S5` `#E4AC44`, Enterprise Blue `S6` `#0EA5E9`, and Emerald `S4` `#1CC491` on their listed anchors. | Recolor those three ramps off the listed hex. | `02-spec/07-design-system/32-slide-color-options.md` |
| DS-IMP-118 | Keep Sunset accent `#F97316`, Monochrome `S0` `#09090B` to `S9` `#FAFAFA`, and Cyan accent `#06B6D4`. | Swap Sunset or Cyan onto the bright-gold stage accent. | `02-spec/07-design-system/32-slide-color-options.md` |
| DS-IMP-119 | Store `slide.gradient` as `linear` or `radial`, angle default 135, and 2 to 4 stops. Shade a hero name at 104px weight 700: first glyph accent, second glyph `#fdd072`, remainder the text token. | Use more than 4 stops, or flat-fill the hero name with one color. | `02-spec/07-design-system/32-slide-color-options.md` |
| DS-IMP-120 | Use only the seven pills `yellow`, `white`, `purple`, `green`, `blue`, `pink`, `red` with their listed contrast ratios. Custom ink uses `L = 0.299R + 0.587G + 0.114B`. `L > 0.6` selects `#0b0b12`. The middle align cell is `center`. | Name the middle cell `center-center`, or add an eighth pill. | `02-spec/07-design-system/32-slide-color-options.md` |
| DS-IMP-121 | On a title slide, place the kicker at left 140 top 120, 14px mono, tracking `0.25em`. Headline is 82px Ubuntu at top 260, max 1400px. Subtitle is 26px Poppins at top 460. Avatar is 60 by 60. Wave height is 140 to 180px. | Bake that text into a png, jpg, or webp. | `02-spec/07-design-system/34-slide-layout-catalog.md` |
| DS-IMP-122 | On `usp-strike`, set the headline to 124px, tracking `-0.03em`, strike thickness 6px at accent 0.7. On `pricing`, the hot tier border is 2.5px `#7C3AED` and the scale is 1.03. Price is 48 to 52px. CTA height is 52px, radius 12px. | Use a 1px strike, or scale the hot tier to the recommendation-card 1.06. | `02-spec/07-design-system/34-slide-layout-catalog.md` |
| DS-IMP-123 | On `before-after`, each card is 790px. Before is `#FFF5F5` with a rose border. After uses a 2.5px violet or emerald border. On `steps-chain`, the horizon is 3px at top 364, badges are 48 by 48, cards are 370px wide. | Paint both halves the same white, or drop the horizon. | `02-spec/07-design-system/34-slide-layout-catalog.md` |
| DS-IMP-124 | On `quote`, the mark is 120px at accent 0.25 and the line is 60px. On `depth-stack`, perspective is 1200px and the back layers are scale 0.92 at opacity 0.5, then 0.96 at 0.8. Partner logos rest at opacity 0.6 and hover to 1. | Flatten the stack to one scale, or leave partner logos fully opaque at rest. | `02-spec/07-design-system/34-slide-layout-catalog.md` |
| DS-IMP-125 | On `embed`, the frame is 1640 by 760, radius 24px. On `poll` and `qa`, the QR card is 280 by 280 at right 140, and vote bars are 1100 by 72. | Shrink the embed, or replace the QR card with a room code only. | `02-spec/07-design-system/34-slide-layout-catalog.md` |
| DS-IMP-126 | On `key-player`, portraits stay 380 to 420px and the bio is live 14px text. On `left`, the narrative is 800px and the plate is 760 by 720 at radius 20px. On `bullets`, cards are 88px tall with a 48px icon well. | Bake bios or quotes into the portrait. | `02-spec/07-design-system/34-slide-layout-catalog.md` |
| DS-IMP-127 | Persist the deck in `deck-v1`. Keep selection, drag, and undo in the edit store, and clear that store on refresh. Draft key is `deck.draft.v1`. Mutate only through `applyEdit()`. | Write colors into a second store, or leave edit mode mounted on the audience route. | `02-spec/07-design-system/35-slide-builder-canvas-inspector.md` |
| DS-IMP-128 | Stack the builder as base 0, watermarks 10, media 20, DOM type 30, ink 40, selection 50, inspector 60. The selection box is 2px `#0ea5e9` with 8px handles. Persist drags as `{ x, y, width, height }`. | Put ink above the selection box, or change border width in a way that shifts layout. | `02-spec/07-design-system/35-slide-builder-canvas-inspector.md` |
| DS-IMP-129 | Size the builder panel to 264px wide, default left 16 top 16, and clamp it inside the viewport. Snap guides within 5px. Preview stagger is 0.08s. Camera zoom pad runs from 1.0 to 2.5. | Let the panel drag past the viewport, or snap farther than 5px. | `02-spec/07-design-system/35-slide-builder-canvas-inspector.md` |
| DS-IMP-130 | While the builder is on, `B` or `E` toggles it, and that chord is suppressed inside a text field. Undo is `Cmd/Ctrl+Z`. Redo is `Cmd/Ctrl+Shift+Z`. Escape deselects. Digits 1 through 7 select the seven flagship themes. | Fire the builder toggle while a text field is focused, or drop the title when the type changes. | `02-spec/07-design-system/35-slide-builder-canvas-inspector.md` |
| DS-IMP-131 | Print at `/slides/print?theme=light&reducedMotion=true`, one slide per page, zero animation delay. PPTX text stays editable. | Animate the print route, or flatten PPTX text to images. | `02-spec/07-design-system/35-slide-builder-canvas-inspector.md` |
| DS-IMP-132 | Open the presenter console at 1280 by 800. Sync on `BroadcastChannel("deck-sync")` with `NAVIGATE`, `STEP`, `THEME_CHANGE`, and `RELOAD`. Speaker notes are 20px. | Run presenter notes inside the audience window. | `02-spec/07-design-system/24-slide-presentation-system.md` |
| DS-IMP-133 | Follow the webcam size table in the file you opened. File 24 cycles Hidden, 120, 180, then 240 by 160. File 43 uses `sm` 240 by 135 through `xl` 640 by 360, with key `1` filling the stage. | Mix the 120px cycle with the 240px preset table in one component. | `02-spec/07-design-system/43-slide-webcam-overlay.md` |
| DS-IMP-134 | Rest unrevealed steps at opacity 0.15, blur 2px, and `translateY(8px)`. Reveal them over 300ms `cubic-bezier(0.16, 1, 0.3, 1)`. | Advance the slide while steps remain, or show unrevealed items at full opacity. | `02-spec/07-design-system/24-slide-presentation-system.md` |
| DS-IMP-135 | On `G`, use gap 24px, padding 48px, and thumbnails 280 by 157. On `J`, anchor the jumper at top 24px. `Enter` goes to `#slide-{n}`. Suppress shortcuts while an input is focused. | Move the overview by pointer only. | `02-spec/07-design-system/24-slide-presentation-system.md` |
| DS-IMP-136 | At 1024px and up, show 3 cards and advance one card width plus the gutter with `translate3d`. | Enable native scroll-snap on that animated track, or wipe more than one card. | `02-spec/07-design-system/19-modern-motion-and-sliding-interactions.md` |
| DS-IMP-137 | At 768px and below, show one active card plus a 24px peek of the next. | Show the desktop 3-up row at that width. | `02-spec/07-design-system/19-modern-motion-and-sliding-interactions.md` |
| DS-IMP-138 | Autoplay every 5.2s, and pause on hover, focus-within, touch, off-screen, and reduced motion. | Leave the 5.2s timer running through hover or reduced motion. | `02-spec/07-design-system/19-modern-motion-and-sliding-interactions.md` |
| DS-IMP-139 | On a clone edge, wait for `transitionend`, set `transition: none`, reset the index, force reflow, then restore the transition. | Change the index while the slide transition is still on. | `02-spec/07-design-system/19-modern-motion-and-sliding-interactions.md` |
| DS-IMP-140 | On a carousel card, brighten the border and scale the image with `scale(1.02)`. | Add a `translateY` hover lift on a card inside the horizontal track. | `02-spec/07-design-system/19-modern-motion-and-sliding-interactions.md` |
| DS-IMP-141 | Run `rise` as opacity plus `translate3d(0, 18px, 0)` in 520ms, then stagger children by 70ms, once in reading order. | Replay that entrance when the user scrolls back up. | `02-spec/07-design-system/19-modern-motion-and-sliding-interactions.md` |
| DS-IMP-142 | Stagger hero lines by 80ms. Reveal the visual from scale 0.96 in 650ms. Drift a logo track for 40s `linear` and pause it on hover. | Ease an infinite ticker with the editorial curve. | `02-spec/07-design-system/19-modern-motion-and-sliding-interactions.md` |
| DS-IMP-143 | Tint a light row with `rgba(15, 23, 42, 0.05)` and text `#0f172a`. | Hover that row with a gray that disappears on white. | `02-spec/07-design-system/21-css3-animations-and-interactions.md` |
| DS-IMP-144 | Draw the row pill at 3.5px, from `scaleY(0.3)` to `scaleY(1)` in 250ms on the editorial curve, and shift padding-left to `1.75rem`. | Use a static full-height bar. | `02-spec/07-design-system/21-css3-animations-and-interactions.md` |
| DS-IMP-145 | Dismiss overlays with `cubic-bezier(0.4, 0, 1, 1)` in 150 to 200ms. | Close them on the 520ms entrance duration. | `02-spec/07-design-system/21-css3-animations-and-interactions.md` |
| DS-IMP-146 | Rest a light web card at `0 4px 6px -1px rgba(15, 23, 42, 0.04)` and hover to `0 20px 40px -15px rgba(15, 23, 42, 0.12)` in 300ms. | Scale that card, or paste this slate stack onto a dark slide. | `02-spec/07-design-system/21-css3-animations-and-interactions.md` |
| DS-IMP-147 | Run a marquee for 35s `linear` with a 10 to 90 percent mask, and pause it on hover. | Keep the ribbon moving under the pointer. | `02-spec/07-design-system/21-css3-animations-and-interactions.md` |
| DS-IMP-148 | Open an accordion with `grid-template-rows` from `0fr` to `1fr` in 300ms, and turn the chevron from 45deg to -135deg in 250ms. | Measure panel height in JavaScript. | `02-spec/07-design-system/21-css3-animations-and-interactions.md` |
| DS-IMP-149 | On light quiz surfaces use `--text-shadow-rest: 0 1px 2px rgba(0, 0, 0, 0.06)` and `--elevation-hover` at primary alpha 0.24. | Invent a second shadow variable in feature CSS. | `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md` |
| DS-IMP-150 | On dark option decks set hover primary alpha to 0.28 and `--elevation-selected` to `0 0 16px hsl(var(--primary) / 0.27)`. | Use the 0.28 alpha on light surfaces. | `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md` |
| DS-IMP-151 | Rest a dark `.modern-quiz-card` on `0 10px 25px -5px rgba(0, 0, 0, 0.4)` plus `0 8px 10px -6px rgba(0, 0, 0, 0.3)`, and transition the shadow in 250ms. | Reuse the light option elevation on that plate. | `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md` |
| DS-IMP-152 | Run `slide-up-fade` for 420ms from `translate3d(0, 12px, 0)`, with delays stepping by 0.06s. | Delay each option by 400ms. | `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md` |
| DS-IMP-153 | Dispatch presenter keys first-match: form focus, then dialog Escape, then builder `B` or `E` and digits 1 through 7, then a pending jump, then slide-specific 1 through 4, then camera keys, then deck navigation. | Let a later group steal an earlier match. | `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md` |
| DS-IMP-154 | Pad a center stage with `96px 120px`. Cap the headline at 920px and `clamp(40px, 4.2vw, 72px)`, line-height 1.08. Kicker tracking is `0.12em`. | Set this layout’s title at the 128px cover size. | `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md` |
| DS-IMP-155 | Mirror `green-choice` to botanical-light (`--primary: 142 70% 30%`) and require the contrast script to exit 0. Keep the review hex twins out of TSX. | Paste `#F3F7F5` or `#166534` into a component. | `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md` |
| DS-IMP-156 | Give `StepTimelineSlide` a 360px left column and a 720px description. Cards are 84px tall, radius 16px, rest opacity 0.70. Focus uses a 2px primary border and `scale(1.02)`. | Put the description column on the left. | `02-spec/07-design-system/42-slide-step-and-sound-system.md` |
| DS-IMP-157 | Keep the closed sound set: `whoosh` 350ms, `click` 180ms, `fadeClick` 180ms, `zoom` 500ms, `fadeZoom` 450ms, `pop` 120ms, with the listed gains. | Add a seventh sound kind. | `02-spec/07-design-system/42-slide-step-and-sound-system.md` |
| DS-IMP-158 | Debounce playback at 60ms, skip when `document.hidden`, store mute at `deck.sound.muted`, and start `AudioContext` on the first pointer or key. | Start a second play inside that 60ms gate. | `02-spec/07-design-system/42-slide-step-and-sound-system.md` |
| DS-IMP-159 | Hotspot diameter is 32px. Callout width is 280px. `pulse-ring` period is 2.4s. `ArrowLeft` steps backward before it changes slides. | Change slides on the first step key while steps remain. | `02-spec/07-design-system/42-slide-step-and-sound-system.md` |
| DS-IMP-160 | Mount the file-43 webcam at `z-index: 55` and honor `--brand-inset-x` at 218px. Default shape for `md` is bottom-right. | Cover the top 80px brand band or the bottom 80px controller zone. | `02-spec/07-design-system/43-slide-webcam-overlay.md` |
| DS-IMP-161 | In the webcam file, `C` toggles the stream, `M` mirrors with `scaleX(-1)`, `F` toggles auto-frame, `+` and `-` step zoom by 0.2 between 1.0 and 3.0, and `Esc` exits. Persist `deck.webcam.settings.v1`. Stop the track on unmount. | Read `M` as minimize or `F` as deck fullscreen from this file. | `02-spec/07-design-system/43-slide-webcam-overlay.md` |
| DS-IMP-162 | Mask a squircle with `radial-gradient(circle at center, black 65%, transparent 70%)` and radius `28% / 28%`. Rim gradient runs `#fdd072`, `#bb8e4e`, `#f3c258`. Auto-frame damps at k = 0.12 over 300ms. | Snap the face pan in one frame. | `02-spec/07-design-system/43-slide-webcam-overlay.md` |
| DS-IMP-163 | Split the inspector 60 percent live and 40 percent next. `I` opens it. `R` resets the clock to `00:00`. Within 30s of target use amber semibold. Past target use red bold pulse. On-target band is ±0:45. Do not reset a stored `deck.inspector.startedAt` on refresh. | Reset the clock on refresh when `startedAt` is already stored. | `02-spec/07-design-system/44-slide-presenter-inspector-and-handouts.md` |
| DS-IMP-164 | Each handout page holds exactly 3 slides at `@page { size: 1920px 1080px landscape; margin: 0 }`, padding 48px 64px. A short last page keeps the empty row. | Stretch two rows to fill a 3-row page. | `02-spec/07-design-system/44-slide-presenter-inspector-and-handouts.md` |
| DS-IMP-165 | Ruled notes use a repeating gradient every 28px at height 84px. `?auto=1` calls print after 800ms. Print hides the controller, the webcam, and the pagination dots. | Print the webcam bubble or the controller pill. | `02-spec/07-design-system/44-slide-presenter-inspector-and-handouts.md` |
| DS-IMP-166 | A 1200 by 630 social card has at most 5 text zones. Hero numerals are Ubuntu 700 at 190 to 220px. Headlines are 54 to 62px, two lines max. | Add a sixth zone, or pair a 300px numeral with 24px support. | `02-spec/07-design-system/45-ai-visual-whitespace-discipline.md` |
| DS-IMP-167 | Keep social gaps at or above 60px, side margins at 80px, and a 55px perimeter. Segments use radius 16 to 18px and at most 3 per row. | Shrink a duplicate kicker to make it fit. | `02-spec/07-design-system/45-ai-visual-whitespace-discipline.md` |
| DS-IMP-168 | Step dark planes about 3 to 5 percent. Plane 0 is near `#0A0F1D` or warm `#0B0A09`. `#F8FAFC` on plane 0 is 16.4:1. `#64748B` on plane 2 is 3.5:1 and is only for large or bold text. | Use canvas `#000000`, or seat three sections on the same plane. | `02-spec/07-design-system/18-dark-mode-and-materiality.md` |
| DS-IMP-169 | Rest a dark card on a 1px border at white 0.08, an inset hairline at 0.12, and shadow `0 8px 24px -12px rgba(0, 0, 0, 0.6)`, in 0.2s. Grain opacity is 0.035 and stays off long prose and dense tables. | Cut a photo in with no scrim. | `02-spec/07-design-system/18-dark-mode-and-materiality.md` |
| DS-IMP-170 | Tint a quiz preview with `box-shadow: inset 0 0 0 9999px hsl(var(--primary) / 0.05)`. | Paint that 5 percent wash with `background-color`. | `wp-exam/src/styles/quiz-preview.css` |
| DS-IMP-171 | Mark preview mode with a fixed 2px primary top edge, `z-index: 60`, `pointer-events: none`, and a 10px mono pill. | Use a solid banner across the canvas. | `wp-exam/src/styles/quiz-preview.css` |
| DS-IMP-172 | On Riseup, ring the preview pill with `#e8c547` and keep the fill transparent. On Dracula, set the pill text to `#bac7e8`. | Fill the Riseup pill with gold. | `wp-exam/src/styles/quiz-preview.css` |
| DS-IMP-173 | Enter the quiz header with `translateY(10px)` over 0.45s, then options over 0.4s starting at 140ms and adding 70ms, capping the 6th and later at 490ms. | Start the first option at 0ms. | `wp-exam/src/styles/quiz-preview.css` |
| DS-IMP-174 | Under reduced motion, collapse that quiz cascade to a 120ms opacity fade. | Leave the 10px rise running. | `wp-exam/src/styles/quiz-preview.css` |
| DS-IMP-175 | Transition option cards as separate properties: transform and opacity 0.3s, border and background 0.2s, box-shadow 0.22s. | Use `transition-all`. | `wp-exam/src/styles/quiz-preview.css` |
| DS-IMP-176 | On option hover, scale only the badge to 1.05 and fill it at primary 0.22. Set the label to weight 600. Primary color is for badges, icons, borders, and filled buttons. | Paint paragraph text in the primary color. | `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md` |
| DS-IMP-177 | Keep `--option-text-shadow-rest` and `--option-text-shadow-hover` as aliases for one release. | Delete the aliases while call sites still read them. | `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md` |
| DS-IMP-178 | Re-declare `--text-shadow-*` and `--elevation-*` on `.quiz-embed[data-theme="botanical-light"]`. | Assume root tokens survive inside the embed. | `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md` |
| DS-IMP-179 | While the camera is fullscreen, still honor Right, Down, Enter, Space, PageUp, and PageDown. `[` exits. `]` cycles the cinematic state. | Drop Down and PageUp while the camera covers the stage. | `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md` |
| DS-IMP-180 | Open the shortcut map on unmodified `/` only. | Bind `/` while focus is an input, textarea, or contentEditable. | `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md` |
| DS-IMP-181 | Start a tree rename with F2 on the selected row, or by double-clicking the title span. Enter saves. Blur saves. Escape restores the previous title. | Treat a double-click on the row body as rename. That click opens the inspector. | `antigravity-manager/src/components/instances/PromptTreeView/ConversationNode.tsx` |
| DS-IMP-182 | Trim the rename draft. If it is empty or unchanged, abort. Apply the title, then roll it back if the rename call fails. | Leave the failed title on the row. | `antigravity-manager/src/components/instances/PromptTreeView/ConversationNode.tsx` |
| DS-IMP-183 | Let ArrowUp and ArrowDown walk rows in expanded branches only, wrapping at the ends. | Include rows inside a collapsed branch, or steal the keys from a text field. | `antigravity-manager/src/components/instances/PromptTreeView/TreePanel.tsx` |
| DS-IMP-184 | Treat search, status, category, scope, and the media filter as one set. Reset clears all five. | Clear the search and leave the media filter on. | `antigravity-manager/src/components/instances/PromptTreeView/TreePanel.tsx` |
| DS-IMP-185 | Keep running and queued rows visible while automated rows are hidden. | Hide a running row because it is classified as automated. | `antigravity-manager/src/components/instances/PromptTreeView/ProjectNode.tsx` |
| DS-IMP-186 | When nothing matches a query, name the query and offer a clear action. When filters match nothing, say so and offer reset. | Use one empty sentence for search, filters, and a truly empty tree. | `antigravity-manager/src/components/instances/PromptTreeView/TreePanel.tsx` |
| DS-IMP-187 | Park untitled idle rows under an archived group with a dashed rail, toggled per branch. Running or queued rows are never stale. | Mix those rows into the live list. | `antigravity-manager/src/components/instances/PromptTreeView/conversationGuards.ts` |
| DS-IMP-188 | Close the tree on Escape only when no context menu and no details dialog is open. | Let Escape dismiss the tree while a nested dialog is up. | `antigravity-manager/src/components/instances/PromptTreeView/PromptTreeViewModal.tsx` |
| DS-IMP-189 | Persist pane width as a product-local pattern: read on open, clamp, write on mouseup, reset on splitter double-click, ignore storage failures. The reference key `agm_prompt_tree_panel_width` is an example. | Require that exact key in another product. | `antigravity-manager/src/components/instances/PromptTreeView/PromptTreeViewModal.tsx` |
| DS-IMP-190 | Hide the running strip when it is idle and error-free. On a failed poll, keep the last good rows and show a stale mark. | Re-sort or drop the last good rows inside the strip. | `antigravity-manager/src/components/instances/PromptTreeView/NowRunningUpNext.tsx` |
| DS-IMP-191 | On a recommendation card, scale the avatar to 1.12 and move the arrow `translateX(4px)` to opacity 1. Recolor the name to the accent. | Scale the quote. | `presentations-repos/bsrm-presentation-hiltrax/spec/_shared/design-system/26-hover-interactions.md` |
| DS-IMP-192 | Hover HUD icon buttons to accent text on `hsl(0 0% 100% / 0.1)` in 150ms. The active 32 by 8px dot does not change on hover. | Grow the active pill on hover. | `presentations-repos/bsrm-presentation-hiltrax/spec/_shared/design-system/26-hover-interactions.md` |
| DS-IMP-193 | Use CSS for color and opacity hovers. Hiltrax file 26 allows a spring at stiffness 300 and damping 20 only for a `whileHover` scale of 1.05. | Use that 1.05 spring as the default slide hover, or put a spring on a color change. | `presentations-repos/bsrm-presentation-hiltrax/spec/_shared/design-system/26-hover-interactions.md` |
| DS-IMP-194 | Advance on a horizontal swipe only when it exceeds 50px and exceeds the vertical distance. Disable deck keys while the counter field is focused. | Treat a mostly vertical drag as next. | `presentations-repos/bsrm-presentation-hiltrax/spec/_shared/design-system/28-navigation-hud.md` |
| DS-IMP-195 | Keep share toasts at 2000ms with the stock toaster. The share menu is 224px wide, radius 12px, blur 16px. | Theme the toaster with slide gold. | `presentations-repos/bsrm-presentation-hiltrax/spec/_shared/design-system/15-toast-system.md` |
| DS-IMP-196 | Draw section accent bars 4px tall, radius 2px, fading to the glow token. Rest a metric number in a two-layer accent text-shadow, `0 0 30px` at 0.4 and `0 0 60px` at 0.15. | Force that horizontal bar to 72px wide. | `presentations-repos/bsrm-presentation-hiltrax/spec/_shared/design-system/27-glass-card-system.md` |
| DS-IMP-197 | When file 29 says HUD `z-index` 50 and file 31 says 60, follow the file you are editing. | Average the two indexes. | `02-spec/07-design-system/31-slide-controller-buttons.md` |
| DS-IMP-198 | Bare digit `1` is camera stage-fill, except while the builder is open (themes 1 through 7) and except on a profile-grid slide (profiles 1 through 4). | Bind stage-fill and “open profile A” to the same key on that slide. | `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md` |
| DS-IMP-199 | The builder hotkey is `B` or `E`, and it is suppressed in text fields. | Bind `E` to the website visual builder. | `02-spec/07-design-system/29-slide-navigation-and-builder.md` |
| DS-IMP-200 | Record `letsmarknow` and `letsmarknow-ui` as still missing on this pass. The next append starts at `DS-IMP-201`. | Invent a Let's Mark component, or renumber 001 through 200. | `01-prompts/30-improvements/08-design-system-spec-improvements.md` |
