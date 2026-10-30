# Three-theme repo improvement audit

**Status:** Published 30-Oct-26
**Plan:** `.ai-memory/plans/pending/02-slides-system-overhaul.md`

## Theme 1 — Guideline source of truth

`02-spec/17-consolidated-guidelines/34-compiled-simple-coding-guidelines.md` stays the rule source. Slides teach those rules. They do not invent a second palette or a second rule id.

## Theme 2 — Slides as the teaching surface

`slides-app/src/deck/registry.ts` is the shipped deck. It includes the control-flow, error, schema, React, accessibility, workflow, and operations slides, and visual snapshots exist through `slide-69`. The July note that said 13 of 70 were shipped is stale.

## Theme 3 — Release and CI

Steps 3, 4, and 5 of the parent plan were already shipped: the release zip, `scripts/sync-guidelines.mjs`, smoke, axe-core, and visual specs. Recoloring `slides-app` to bright-gold stays a separate review. This audit does not change tokens.
