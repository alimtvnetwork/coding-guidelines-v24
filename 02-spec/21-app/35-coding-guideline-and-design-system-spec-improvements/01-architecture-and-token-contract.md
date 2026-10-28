# Architecture and Token Contract

**Updated:** 28-Oct-26

## User Request (Verbatim)

The user asked for a solid design-system spec so any AI, including a blind or weak one, can make professional presentations with color, animation, and hover in one shot. The spec must follow the presentation design systems, take motion ideas from WP Exam, document the Antigravity tree view generically, list 100 improvements every time the prompt is repeated, and add one repeatable prompt under the improvements folder. Prompts outside that one file stay out of scope. Paths inside repository files stay relative. No absolute paths.

## Why this folder exists

`02-spec/07-design-system/` already holds the slide engine, four website systems, and an eight-theme switcher. Those files disagree with each other and with `slides-app/src/styles/tokens.css`. A weak model mixes them and the deck looks cheap. This folder does not replace file 40. It tells the model which file to obey for which job.

## Decision table

| Job the user named | Authority | Paint these | Forbidden on that job |
| :--- | :--- | :--- | :--- |
| Slides, deck, keynote, presentation, no theme named | `02-spec/07-design-system/40-theme-switch.md` id `bright-gold-tech` plus `02-spec/07-design-system/05-bright-gold-tech/01-tokens-background-and-chrome.md` | Canvas `#08080C`, surface `#111118`, card `#181820`, accent `#FFAE00`, glow `#FFC64D`, ink `#FFFFFF`, muted `#B8BDC6` | Amber deck hex from file 30, blue HSL from `slides-app`, website JSON palettes, `white-blue` on `.slide-stage` |
| User names dark amber or scripted deck | `02-spec/07-design-system/30-slide-palette-type-and-shell.md` only | Canvas `#0B0B0E`, card `#121217`, accent `#F5A623`, hover `#F8B647` | `#FFAE00`, `#08080C` |
| User names one of the other seven stage themes | That row in `02-spec/07-design-system/40-theme-switch.md` | The six tokens in that row only | A ninth theme, `vscode-dark` as a stage id |
| Marketing site, AI SaaS | `02-spec/07-design-system/02-ai-system-design/` | Pink `#FF2D6F`, page `#F7F4F3`, ink `#111827` | Slide-stage gold, purple-blue full-bleed wash |
| Sweet Digs | `02-spec/07-design-system/03-sweet-digs-design-system/` | Emerald `142 71% 45%` unless a named theme in that folder is requested | `scale(1.05)` hover zoom |
| White Blue | `02-spec/07-design-system/04-white-blue-theme/` | Paper `#FFFFFF`, navy `#0D2975`, cobalt `#2563EB` | Inter as the UI font |
| Tree plus right panel | [03-tree-view-and-detail-panel.md](03-tree-view-and-detail-panel.md) | Role tokens in that file | One selection color for every depth |

## Stage geometry (default slide)

1. Author on a 1920 by 1080 stage.
2. Scale with the minimum of viewport width over 1920 and viewport height over 1080.
3. Safe area is 96px left and right, 80px top and bottom. Logo, dots, and the controller sit outside that inset.
4. Display face is Ubuntu 700. Body face is Poppins. Cover title is 128px, line-height 1.05, tracking `-0.02em`. A section title is 88px. A bullet is 28px at line-height 1.5.
5. Components read `var(--canvas)`, `var(--ink)`, `var(--ink-muted)`, `var(--accent)`, `var(--accent-ink)`, and `var(--card)`. They do not hardcode a hex inside a slide component.
6. File 30 stores bare HSL triplets. File 40 stores full `hsl()` or hex. Do not wrap a file-40 value in `hsl()` again.

## Motion contract

| Moment | Duration | Easing | Motion |
| :--- | :--- | :--- | :--- |
| Gold entrance `gold-rise` | 0.75s | `cubic-bezier(0.25, 0.46, 0.45, 0.94)` | Rise. Stagger siblings by 0.22s |
| Theme crossfade | 0.45s | `cubic-bezier(0.4, 0, 0.2, 1)` | Background and color only |
| General UI hover cap | 0.3s maximum | `cubic-bezier(0.4, 0, 0.2, 1)` | `translateY(-2px)` unless a named card exception applies |
| Bright-gold recommendation card | 0.3s | `ease` | `scale(1.06) translateY(-8px)`, accent border at 0.6, glow shadow |
| Choice row (from `wp-exam/src/styles/theme.less`) | 220ms | `cubic-bezier(0.16, 1, 0.3, 1)` | `translate3d(6px, 0, 0)`, not a lift |
| Reduced motion | 0.01ms | — | Kill transform. Do not skip the state change |

Website card lift of `-4px` in `02-spec/07-design-system/tokens/theme-palette.less` is a website rule. It is not the slide hover.

## Cheap tells (reject the frame)

Reject the slide when any of these are true:

- A purple-to-blue full-screen gradient is the ground.
- Two saturated accents sit in one viewport.
- Three equal icon cards are the whole layout.
- The hero is centered, empty, and padded the same on every section.
- A button grows with `scale(1.05)` on hover.
- A dark surface uses a heavy black drop shadow at rest.
- The headline is baked into a picture.
- Copy is adjectives with no number and no concrete claim.
- Inter is used on a White Blue or slide stage.
- `slides-app` blue (`217 91% 60%`) or green (`160 84% 39%`) is painted onto a gold deck.

## What stays untouched

- Do not edit `02-spec/07-design-system/40-theme-switch.md` in the same turn as this spec unless the user asks for a token change.
- Do not bump `version.json`.
- Do not sync `02-spec/21-app/` to other repositories.
