# Slides visual-regression baselines

Playwright stores per-slide PNGs in this folder (auto-created on the
first `--update-snapshots` run). Each file is named `slide-NN.png`
and represents the `.slide-content` element captured on Linux Chromium
at `1280x800` with animations disabled.

## When to regenerate

Regenerate baselines whenever a visual change is intentional (typography,
tokens, layout, component swap, new slide content). Do NOT regenerate to
paper over unexplained diffs; investigate first.

## How to regenerate

Locally (Linux only, matches CI):

```bash
cd slides-app && bun run build && cd ..
bun run test:slides:visual:update
git add slides-app/tests/visual.spec.ts-snapshots
```

Via CI (any OS): run the `slides-visual` workflow with
`update_baselines=true`, download the `updated-visual-baselines`
artifact, and commit its contents here.

macOS/Windows local runs will diff against Linux baselines and fail;
that is expected. Use CI or a Linux container for authoritative runs.
