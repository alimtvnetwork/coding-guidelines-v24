# Blind AI Execution Checklist

**Updated:** 28-Oct-26
**Second pass:** also pass `DS-IMP-101` through `DS-IMP-200` when the shot uses layouts, motion, webcam, quiz chrome, or tree rename.

A weak model follows this list in order. It does not improvise a palette. It does not open a second token file "for inspiration."

## Before painting

- [ ] The job is one of: slide, amber deck, named stage theme, AI SaaS, Sweet Digs, White Blue, tree.
- [ ] The authority file for that job is copied from [01-architecture-and-token-contract.md](01-architecture-and-token-contract.md).
- [ ] No path in the answer uses a drive letter or `file:///`.
- [ ] Archived skills under `.agents/skills-archived/` are not used.

## Slide shot (default)

- [ ] Stage is 1920 by 1080 and scales with the min ratio.
- [ ] `data-theme` is `bright-gold-tech` on `.slide-stage`.
- [ ] The six tokens match the bright-gold list. No second accent.
- [ ] Cover type is Ubuntu 700 at 128px. Body is Poppins. Bullets are 28px.
- [ ] Safe area is 96px horizontal and 80px vertical.
- [ ] Ground is near-black plus one spotlight. No purple-blue wash.
- [ ] Entrance is `gold-rise` 0.75s, stagger 0.22s.
- [ ] Ordinary hover is at most 0.3s and lifts no more than 2px.
- [ ] The recommendation card is the only control allowed to scale to 1.06 and lift 8px.
- [ ] Reduced motion collapses duration to 0.01ms.
- [ ] The title is HTML. The footer index is `{current} / {total}`.
- [ ] Copy contains a number or a concrete claim.
- [ ] Resting shadows are hairlines. Gold glow is hover-only.

## Tree shot

- [ ] Pane defaults to 340px and clamps between 280 and 760.
- [ ] Collapse is a 40px rail.
- [ ] Row type is about 12px. Indent is about 10px. Chevrons swap.
- [ ] Hover is one wash. Selection has four role colors.
- [ ] The right pane follows the leaf.
- [ ] Empty tree is a skeleton. Empty selection is one sentence.
- [ ] A branch longer than 50 is capped unless the scope is one root.

## Second-pass shot

- [ ] Scripted transitions stay inside `Slide`, `Fade`, `Zoom`, `Flip`, and `Rise` at 0.45s.
- [ ] A center stage uses `clamp(40px, 4.2vw, 72px)`, not the 128px cover title.
- [ ] Webcam sizes come from one file. File 24 and file 43 are not mixed.
- [ ] Bare digit `1` follows `DS-IMP-198`.
- [ ] Tree rename is F2 or a double-click on the title. A double-click on the row body opens the inspector.
- [ ] `letsmarknow` is still missing. Do not invent it.

## Reject and redo

Redo the frame when any cheap tell in the architecture file is present.

## Done when

The model can point at one authority file per color it used, and every motion it used appears in the decision table or the tree token table.
