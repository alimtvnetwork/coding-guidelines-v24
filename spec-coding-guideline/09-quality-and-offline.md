# Offline Quality & Benchmark Standards (AI Execution Prompt)

> **/goal** Establish, enforce, and verify offline operation benchmarks, WCAG AA accessibility, cross-browser compatibility, and performance budgets for slide decks.
> **/learn** Validate zero external network requests, sub-800ms initial load times under `file://`, full keyboard navigation, screen reader live announcements, and automated visual regression testing.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Enforce strict offline isolation: zero `http://` or `https://` URLs in `dist/` and no remote font preconnects.
- [ ] `/learn` Verify keyboard navigation (arrows, space, F, G, P, Esc) and focus-visible rings across all interactive controls.
- [ ] `/goal` Guarantee performance benchmarks: < 800ms initial slide load under `file://` protocol and < 100ms slide transitions.
- [ ] `/learn` Validate WCAG AA contrast ratios (4.5:1), reduced-motion preferences, and screen reader `aria-live` announcements.

. **CRITICAL AI INSTRUCTION:** This specification is an active AI execution directive. All code generated or modified must strictly follow the rules below.

---

**Version:** 1.0.0

---

## Offline guarantees (hard requirements)

The `dist.zip` artifact must satisfy ALL of these. The package script enforces
each one and fails the build if any check fails.

| # | Requirement | Verification |
|---|-------------|--------------|
| 1 | No `http://` or `https://` URL anywhere in `dist/` (HTML, JS, CSS) | `grep -rE 'https?://' dist/ \| grep -v UFL-1.0.txt` returns empty |
| 2 | No `<link rel="preconnect">` or `<link rel="dns-prefetch">` | grep `index.html` |
| 3 | All `<script>` and `<link>` tags use relative paths starting with `./` | grep `index.html` |
| 4 | All `@font-face src:` declarations use `./fonts/...` | grep `dist/assets/*.css` |
| 5 | Opening `dist/index.html` from `file://` shows the title slide within 1s | manual test (or Playwright run with `file://` URL) |
| 6 | Total `dist/` size ≤ 5 MB | `du -sh dist/` |
| 7 | Total `dist.zip` size ≤ 3 MB | `ls -lh dist.zip` |

## Browser support

Target the trainer's likely browsers. Drop legacy support for smaller bundle.

| Browser | Min version | Tested? |
|---------|-------------|---------|
| Chrome / Edge | 110 | required |
| Firefox | 110 | required |
| Safari | 16.4 | required |
| IE / old Edge | — | NOT supported (refuse with a friendly message) |

Set Vite `build.target = 'es2022'` and skip legacy polyfills.

## Accessibility (a11y)

| Rule | How |
|------|-----|
| Keyboard navigation works without a mouse | Arrow keys, Space, F, G, Esc, Home, End |
| Focus ring visible on all interactive controls | `:focus-visible` outline 2px primary |
| Sufficient contrast (WCAG AA) | All text meets 4.5:1 against its background |
| Reduced motion respected | `@media (prefers-reduced-motion: reduce)` zeros out all animation |
| Screen reader announces slide changes | `aria-live="polite"` region announces "Slide X of Y: <title>" |
| Code blocks are real `<pre><code>`, not divs | Shiki output preserves semantic HTML |
| All icons have `aria-label` or `aria-hidden="true"` | enforced via lint |

## Performance budget

| Metric | Budget |
|--------|--------|
| Time to first slide visible (file://) | < 800ms |
| Slide-to-slide transition (cold) | < 100ms |
| Memory footprint after viewing all 13 slides | < 200 MB |
| Initial JS payload (gzipped equivalent) | < 250 KB (no gzip on file://, but a sensible upper bound) |

## Testing strategy

**Unit:** Each slide component renders without throwing (`@testing-library/react`).

**Visual regression:** Playwright opens `dist/index.html` from `file://`,
navigates to each slide, takes a screenshot, diffs against
`tests/__snapshots__/`. Failure on >0.5% pixel diff.

**Offline contract:** A bash script in CI:

```bash
cd slides-app && bun run build && bun run package

# Verify the offline contract

! grep -rE 'https?://' dist/ --include='*.html' --include='*.js' --include='*.css'
[ "$(du -sb dist | cut -f1)" -lt 5242880 ]
[ "$(stat -c%s dist.zip)" -lt 3145728 ]
```

**Manual smoke test:** Before each release, the author unzips `dist.zip` on a
fresh machine WITH NETWORK DISABLED and confirms all 13 slides render with full
typography and animations.

## Documentation deliverables

When implementation lands, these files become required:

- `slides-app/readme.md` — setup, build, package, troubleshoot
- `slides-app/changelog.md` — semver per the main repo's release cadence
- `dist/README.txt` — end-user usage (see
  [07-build-and-zip-pipeline.md](./07-build-and-zip-pipeline.md))

## Cross-references

- Build pipeline (where the offline checks run): [07-build-and-zip-pipeline.md](./07-build-and-zip-pipeline.md)
- Architecture (Vite `base: './'` + bundled fonts): [02-architecture.md](./02-architecture.md)

---

## Verification & Acceptance Criteria

### AC-CG-SLIDE-009: Offline Viewer Quality Benchmarks and Reliability Standards

**Given** The standalone slide deck distribution bundle tested in an isolated, offline environment.
**When** Quality assurance suites, accessibility audits, and CI/CD validation scripts inspect bundle artifacts.
**Then** All offline guarantees are verified (zero external URLs, size < 5 MB), keyboard accessibility meets WCAG AA standards, and relative path linters pass with zero violations.

**Verification command:**
```bash
python linter-scripts/check-relative-paths.py
```
**Expected:** exit 0. Zero absolute paths or file:/// URIs.
