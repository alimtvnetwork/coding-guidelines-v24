# Slides Architecture & Isolation (AI Execution Prompt)

> **/goal** Architect and maintain the standalone `slides-app/` Vite + React presentation system strictly isolated from the main application without cross-boundary imports or shared bundle dependencies.
> **/learn** Enforce the isolated directory boundary, relative asset resolution (`base: './'`), bundled local Ubuntu fonts, and packaging to `dist.zip` for double-clickable offline execution.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Verify `slides-app/` has an independent `package.json`, `tsconfig.json`, and Vite 5 configuration isolated from root dependencies.
- [ ] `/learn` Ensure zero imports cross the boundary between `slides-app/` and root `/src/`, `/spec/`, or `/linters/`.
- [ ] `/goal` Guarantee `base: './'` in `slides-app/vite.config.ts` so static HTML and assets resolve correctly under offline viewing.
- [ ] `/learn` Verify root `.gitignore` excludes `slides-app/node_modules/`, `slides-app/dist/`, and `slides-app/dist.zip`.

. **CRITICAL AI INSTRUCTION:** This specification is an active AI execution directive. All code generated or modified must strictly follow the rules below.

---

**Version:** 1.0.0

## Directory layout (repo root)

```
/dev-server/
├── src/                       ← MAIN Lovable app (untouched)
├── spec/                      ← MAIN spec tree
├── spec-slides/               ← THIS spec (the one you are reading)
└── slides-app/                ← NEW — separate Vite + React app
    ├── package.json           ← own deps, own scripts
    ├── vite.config.ts         ← base: "./"  (critical for double-click HTML)
    ├── tsconfig.json
    ├── index.html
    ├── public/
    │   └── fonts/
    │       └── Ubuntu-*.ttf   ← bundled, no Google Fonts CDN at runtime
    ├── src/
    │   ├── main.tsx           ← entry
    │   ├── App.tsx            ← deck shell, keyboard nav, fullscreen
    │   ├── styles/
    │   │   ├── tokens.css     ← Ubuntu + colors + spacing (mirrors main app palette)
    │   │   └── slide.css      ← .slide-content scoping (see slides-app context)
    │   ├── components/
    │   │   ├── SlideLayout.tsx
    │   │   ├── ScaledSlide.tsx
    │   │   ├── CodeDiff.tsx        ← before/after code blocks with highlight animation
    │   │   ├── BeforeAfter.tsx     ← split-view morph component
    │   │   └── Toolbar.tsx
    │   ├── slides/                  ← ONE FILE PER TOPIC (10 files)
    │   │   ├── 00-title.tsx
    │   │   ├── 01-naming-conventions.tsx
    │   │   ├── 02-nested-if-else.tsx
    │   │   ├── 03-boolean-prefixes.tsx
    │   │   ├── 04-app-error-wrapper.tsx
    │   │   ├── 05-structured-logging.tsx
    │   │   ├── 06-magic-strings.tsx
    │   │   ├── 07-function-and-file-metrics.tsx
    │   │   ├── 08-two-operand-max.tsx
    │   │   ├── 09-positively-named-guards.tsx
    │   │   ├── 10-spec-first-workflow.tsx
    │   │   └── 11-cache-invalidation.tsx
    │   └── deck.ts                  ← ordered slide registry
    ├── scripts/
    │   ├── package-zip.mjs          ← zips dist/ → dist.zip
    │   └── render-gifs.mjs          ← Remotion pipeline (separate concern)
    └── dist/                        ← BUILD OUTPUT (gitignored)
        ├── index.html
        ├── assets/*.js
        ├── assets/*.css
        └── fonts/Ubuntu-*.ttf
```

## Isolation rules (hard requirements)

1. **No imports across the boundary.** `slides-app/` must NOT import anything
   from `/src/`, `/spec/`, `/linters/`, etc. It is a self-contained app that
   happens to live in this monorepo.
2. **Independent `package.json`.** Its dependencies are NOT added to the root
   `package.json`. Lovable's build never touches it.
3. **Independent `node_modules/`.** Installed via `cd slides-app && bun install`.
4. **Root `.gitignore`** must add `slides-app/node_modules/`,
   `slides-app/dist/`, `slides-app/dist.zip`.
5. **Root `package.json`** gets two convenience passthrough scripts (optional):
   - `"slides:build": "cd slides-app && bun run build && bun run package"`
   - `"slides:gifs": "cd slides-app && bun run render-gifs"`

## Tech stack (slides-app only)

| Layer | Choice | Why |
|-------|--------|-----|
| Build | Vite 5 | Same as main app, fast, single-command static build |
| Framework | React 18 + TypeScript | Reuse author's mental model; component-per-slide |
| Styling | Plain CSS + CSS custom properties (NO Tailwind) | Smaller bundle, zero config, works offline trivially |
| Animation | Framer Motion (lightweight subset) OR pure CSS keyframes | Code diff highlight + before/after morph need precise control |
| Code highlighting | Shiki (compiled to static HTML at build time, NO runtime highlighter) | Zero runtime cost, perfect colors, offline |
| Fonts | Ubuntu Regular + Bold + Mono, **bundled as TTF in `public/fonts/`** | Offline guarantee — never load from Google Fonts CDN |
| Routing | None (single-page deck, hash for slide index) | `#/3` = slide 3, works from `file://` |
| Packaging | `bun run package` script using `archiver` or `adm-zip` | Produces `dist.zip` |

## Critical Vite config

```ts
// slides-app/vite.config.ts
export default defineConfig({
  base: './',                   // ← REQUIRED: makes assets resolve from file://
  build: {
    outDir: 'dist',
    assetsInlineLimit: 0,       // never inline — easier to inspect & cache
    rollupOptions: {
      output: { format: 'es' }  // modern browsers only; trainers all have Chrome/Edge
    }
  }
});
```

The `base: './'` is non-negotiable — without it the built `index.html` uses
absolute paths like `/assets/index.js` which fail when opened from `file://`.

## Why a separate app (vs. a route in the main app)

| Concern | Separate `slides-app/` | Route in main app |
|---------|------------------------|-------------------|
| Build output is double-clickable HTML | ✅ trivial | ❌ requires hacks (SSG, hash-router) |
| No accidental coupling to main app's tokens, components, providers | ✅ enforced by directory | ❌ requires discipline |
| Bundle size unrelated to main app | ✅ independent | ❌ shared |
| Can be removed/replaced without touching main app | ✅ delete one folder | ❌ refactor |
| Trainer can fork just the slides | ✅ self-contained | ❌ tangled |

Decision: **separate app**.

## Cross-references

- Slide authoring contract: [03-slide-authoring.md](./03-slide-authoring.md)
- Build pipeline: [07-build-and-zip-pipeline.md](./07-build-and-zip-pipeline.md)

---

## Verification & Acceptance Criteria

### AC-CG-SLIDE-002: Slides Architecture and Isolation Boundary

**Given** The standalone `slides-app/` presentation subsystem within the repository monorepo.
**When** Codebases and linters verify the architecture boundaries, module imports, and path references.
**Then** Zero imports cross from `slides-app/` to parent directories, Vite config maintains `base: './'`, and all relative paths conform to zero-storage standards.

**Verification command:**
```bash
python linter-scripts/check-relative-paths.py
```
**Expected:** exit 0. Zero absolute paths or file:/// URIs.
