# Benchmark: Previous Specs versus This Folder

**Updated:** 28-Oct-26
**Second pass:** `DS-IMP-101` through `DS-IMP-200` were appended. `DS-IMP-001` through `DS-IMP-100` were not rewritten.
**Third pass:** `DS-IMP-201` through `DS-IMP-300` were appended. Earlier ids were not rewritten.
**Method:** read the files named below. No pixel diff was run. No build was run.

## What a weak model could do before

| Gap | Evidence | Typical miss |
| :--- | :--- | :--- |
| Five color contracts | `02-spec/07-design-system/40-theme-switch.md`, `30-slide-palette-type-and-shell.md`, `27-slide-canvas-and-themes.md`, `tokens/design-tokens.json`, `slides-app/src/styles/tokens.css` | Gold slide painted with app blue `217 91% 60%` or amber `#F5A623` |
| Theme count | File 27 lists 10 stage themes. File 40 lists 8 and renames `bright-gold` to `bright-gold-tech` | A ninth theme, or `vscode-dark` on the stage |
| Hover distance | `08-motion-transitions.md` caps lift near `-2px`. `tokens/theme-palette.less` uses `-4px`. Bright-gold recommendation cards use `-8px` plus scale | Every card scales |
| HSL wrapping | File 30 uses bare triplets. File 40 uses full colors | `hsl(hsl(...))` |
| Empty app UI spec | `02-spec/24-app-ui-design-system/readme.md` is a draft stub | The model treats "no spec" as permission to invent |
| Slide skills reduced | `.agents/skills-archived/create-slide-deck/` and `slides-deck-management/` | The model still follows the archived skill |
| No tree contract in this repo | Behavior lived only in `antigravity-manager/src/components/instances/PromptTreeView/` | One selection color, rotating chevrons, pane bound to the root |
| Let's Mark | `letsmarknow` and `letsmarknow-ui` had no UI source | Invented components |

## What this folder forces

| Check | Where | Pass looks like |
| :--- | :--- | :--- |
| One authority per job | `01-architecture-and-token-contract.md` | The model names the file before the hex |
| Default deck | Decision table | `bright-gold-tech` and the six tokens |
| 300 checks, earlier ids kept | `02-one-hundred-design-improvements.md` | Ids `DS-IMP-001` to `DS-IMP-300`. Next append starts at 301 |
| Generic tree | `03-tree-view-and-detail-panel.md` | 340px pane, four selection roles, leaf binding |
| Shot list | `04-blind-ai-execution-checklist.md` | Boxes checked in order |
| Repeat run | `01-prompts/30-improvements/08-design-system-spec-improvements.md` | Next pass appends ids, it does not rewrite 001 |

## Blind-model probe

A model that has only this folder should answer these without guessing. The expected answer is in the architecture file or the tree file.

1. Unnamed slide theme id: `bright-gold-tech`.
2. Canvas hex: `#08080C`.
3. Accent hex: `#FFAE00`.
4. Cover size: 128px Ubuntu 700.
5. Ordinary hover cap: 0.3s and about `-2px`.
6. Reduced-motion duration: `0.01ms`.
7. Tree default width: 340px.
8. Collapsed tree: 40px rail.
9. Detail binds to: the leaf, not the root.
10. Empty tree: skeleton, not a sentence.
11. Let's Mark: missing source, do not invent.
12. Archived `create-slide-deck`: do not restore.

## Second-pass probes

These answers live in `DS-IMP-101` through `DS-IMP-200`. A weak model still answers probes 1 through 12 from the first pass.

13. Scripted transition family: `Slide`, `Fade`, `Zoom`, `Flip`, `Rise`, at 0.45s. Not a sixth name.
14. Swipe threshold: 50px, and the horizontal distance must beat the vertical distance.
15. Center-stage title: `clamp(40px, 4.2vw, 72px)`, not the 128px cover size.
16. Hot pricing tier scale: 1.03. Recommendation cards stay at 1.06. Those are different controls.
17. Webcam sizes: file 24 and file 43 disagree. Follow the file you opened. Do not mix 120px with the `sm`–`xl` table.
18. Digit `1`: camera stage-fill, unless the builder is open or the slide is a profile grid.
19. Tree rename: F2 or a double-click on the title. A double-click on the row body opens the inspector.
20. Let's Mark on the second pass: still missing. Next id was `DS-IMP-201`.

## Third-pass probes

21. AI Join control: height 48px and a 300ms editorial track. Sweet Digs Join stays 42px and 350ms. Do not copy one onto the other.
22. Sweet Digs reduced motion: set animation, transition, and transform to `none`. Do not use the slide `0.01ms` trick in that file.
23. Website featured pricing scale: 1.04. Slide hot tier stays 1.03. Recommendation cards stay 1.06.
24. Definition page title: 72px, and only line two is accent. Cover stays 128px.
25. Image badge gold is `#FFAA00` on canvas `#090D16`. Slide accent stays `#FFAE00` on `#08080C`.
26. Logo: one archetype, reads at 16px, no black plate on bright-gold. Fill from `--pres-accent`.
27. Tree detail: a selected branch with no leaf still offers send. An empty selection is one sentence and no send button. Detail-body double-click opens the inspector, not rename.
28. Let's Mark on this pass: still missing. Next id is `DS-IMP-301`.

## Not claimed

This benchmark does not claim a rendered deck matches a screenshot. It claims a weak model now has a single decision table where it previously had five conflicting files.
