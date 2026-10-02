# Automated GIF Generation Pipeline (AI Execution Prompt)

> **/goal** Automate high-fidelity animated GIF preview generation for all curriculum topics using Remotion compositions and gifski palette optimization.
> **/learn** Eliminate slide-to-preview visual drift by sharing design tokens and content constants between React slides and Remotion frames, enforcing 1280x720@24fps, infinite looping, and 400 KB per-GIF size ceilings.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Maintain Remotion compositions under `slides-app/src/remotion/` importing canonical tokens and slide components.
- [ ] `/learn` Enforce deterministic frame-by-frame rendering with exact 1280×720 resolution and 24 FPS target.
- [ ] `/goal` Execute `scripts/render-gifs.mjs` to output high-quality GIFs via `gifski` into `public/gifs/`.
- [ ] `/learn` Verify every generated GIF remains strictly under the 400 KB budget with infinite looping configured.

. **CRITICAL AI INSTRUCTION:** This specification is an active AI execution directive. All code generated or modified must strictly follow the rules below.

---

**Version:** 1.0.0

---

## Purpose

Generate one short looping `.gif` per major curriculum topic for embedding on
the **main landing page** (`src/components/landing/`). Each GIF shows the
before/after transformation animating, so a visitor immediately understands
what the review guide teaches without opening the deck.

## Pipeline overview

```
slides-app/src/slides/02-nested-if-else.tsx       (source of truth)
              │
              │  (re-implemented as a Remotion composition that uses
              │   the SAME visual tokens, fonts, and timings)
              ▼
slides-app/src/remotion/compositions/NestedIf.tsx
              │
              ▼
slides-app/scripts/render-gifs.mjs                (Remotion render → MP4 → GIF)
              │
              ▼
public/gifs/nested-if.gif                          (embedded on landing page)
```

## Why a separate Remotion pipeline (not capturing the live deck)

- **Determinism:** Remotion renders frame-by-frame. Capturing a live browser
  produces inconsistent FPS and varying timings.
- **Quality:** Remotion outputs at exact 1280×720@30fps with controlled
  compression.
- **Re-rendering:** Update a slide's content → re-run the script → all GIFs
  rebuild reproducibly.

The Remotion compositions live alongside the React slides and **import the
same content constants and design tokens** — so the GIF and the live slide can
never drift.

## Output spec per GIF

| Property | Value |
|----------|-------|
| Resolution | 1280 × 720 |
| FPS | 24 (smaller file than 30 with no perceivable loss) |
| Duration | 4–6 seconds |
| Loop | Infinite |
| Max file size | **400 KB** per GIF (enforced by the script — re-encode at lower quality if larger) |
| Color depth | 256 colors (default GIF) |
| Compression | `gifski` for high-quality palette generation |

## `scripts/render-gifs.mjs` (skeleton)

```js
import { bundle } from '@remotion/bundler';
import { renderMedia, selectComposition } from '@remotion/renderer';
import { execSync } from 'node:child_process';
import { mkdirSync, statSync } from 'node:fs';

const TOPICS = [
  { id: 'NestedIf',          out: 'nested-if.gif' },
  { id: 'Naming',            out: 'naming.gif' },
  { id: 'BooleanPrefixes',   out: 'boolean-prefixes.gif' },
  { id: 'AppError',          out: 'app-error.gif' },
  { id: 'Logging',           out: 'logging.gif' },
  { id: 'MagicStrings',      out: 'magic-strings.gif' },
  { id: 'Metrics',           out: 'metrics.gif' },
  { id: 'TwoOperand',        out: 'two-operand.gif' },
  { id: 'PositiveGuards',    out: 'positive-guards.gif' },
  { id: 'CacheInvalidation', out: 'cache-invalidation.gif' },
];

const OUT_DIR = '../public/gifs';
mkdirSync(OUT_DIR, { recursive: true });

const bundled = await bundle({ entryPoint: './src/remotion/index.ts' });

for (const topic of TOPICS) {
  const composition = await selectComposition({ serveUrl: bundled, id: topic.id });
  const mp4 = `/tmp/${topic.id}.mp4`;
  await renderMedia({
    composition, serveUrl: bundled, codec: 'h264',
    outputLocation: mp4, muted: true, concurrency: 1,
  });
  // mp4 → gif via gifski (best palette quality)
  execSync(`gifski -o ${OUT_DIR}/${topic.out} --fps 24 --width 1280 --quality 85 ${mp4}`);
  const size = statSync(`${OUT_DIR}/${topic.out}`).size;
  if (size > 400 * 1024) console.warn(`⚠️  ${topic.out} = ${(size/1024).toFixed(0)}KB (over 400KB budget)`);
  else console.log(`✅ ${topic.out} = ${(size/1024).toFixed(0)}KB`);
}
```

## Tooling required

| Tool | Install | Purpose |
|------|---------|---------|
| `@remotion/*` packages | already in `slides-app/package.json` | composition rendering |
| `gifski` | `nix run nixpkgs#gifski -- ...` in sandbox; `cargo install gifski` locally | mp4 → high-quality gif |
| `ffmpeg` | already in PATH | needed by Remotion |

## Embedding on the landing page

The main app's landing-page section (probably a new
`src/components/landing/CurriculumPreview.tsx`) renders a responsive grid of
the 10 GIFs. Each GIF is wrapped in a card with the topic name and a "See in
deck" link that points to the deployed `slides-app/dist/index.html#/<n>` (or
the downloaded zip).

Lazy-load every GIF except the first two (`<img loading="lazy" />`).

## Re-render trigger

Run `bun run render-gifs` whenever a slide's content changes. Add a CI step
that re-runs this and commits updated GIFs (or fails the build if the GIFs are
stale — by hashing the source slide files and comparing to a `gifs.lock.json`).

## Cross-references

- Animation primitives (must match between Framer and Remotion):
  [05-animation-primitives.md](./05-animation-primitives.md)
- Curriculum (source of topic IDs): [06-curriculum.md](./06-curriculum.md)
- Lovable's bundled Remotion skill (used to set up the render env): see the
  `skill/remotion-video` context block.

---

## Verification & Acceptance Criteria

### AC-CG-SLIDE-008: Automated GIF Preview and Animation Generation Pipeline

**Given** Remotion compositions and automated GIF rendering scripts for slide curriculum previews.
**When** The rendering pipeline processes topic compositions to export animated GIFs into `public/gifs/`.
**Then** All topic GIFs meet resolution, framerate, and 400 KB size requirements with zero visual drift from the canonical slides and full compliance with relative path standards.

**Verification command:**
```bash
python linter-scripts/check-relative-paths.py
```
**Expected:** exit 0. Zero absolute paths or file:/// URIs.
