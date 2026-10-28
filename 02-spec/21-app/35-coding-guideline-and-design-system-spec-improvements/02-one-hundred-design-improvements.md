# 100 Design System Improvements

**Updated:** 28-Oct-26
**Ids:** `DS-IMP-001` through `DS-IMP-100`
**Next run:** append `DS-IMP-101` onward. Never renumber. Never delete a row.

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
