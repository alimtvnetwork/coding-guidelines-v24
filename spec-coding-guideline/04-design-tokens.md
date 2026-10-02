# Slide Design Tokens & Typography (AI Execution Prompt)

> **/goal** Enforce strict visual consistency, theme tokenization, typography rules, and 1920×1080 canvas standards across all presentation slide components.
> **/learn** Master the Ubuntu font hierarchy, CSS custom properties with HSL color channels, 6-level spacing scale, and Lucide iconography for offline rendering.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Verify Ubuntu Regular, Bold, and Ubuntu Mono fonts are bundled locally in `slides-app/public/fonts/` with zero Google Fonts CDN calls.
- [ ] `/learn` Enforce HSL component tokens in `tokens.css` for background, foreground, border, brand, and code diff highlighting.
- [ ] `/goal` Maintain fixed 1920×1080 canvas authoring with `--space-6` (96px) outer padding and transform scaling.
- [ ] `/learn` Validate that only tree-shaken Lucide SVG icons are imported without loading runtime icon webfonts.

. **CRITICAL AI INSTRUCTION:** This specification is an active AI execution directive. All code generated or modified must strictly follow the rules below.

---

**Version:** 1.0.0

## Typography — Ubuntu

Ubuntu is the **only** font family. Three TTF files bundled in
`slides-app/public/fonts/`:

| Family | Weight | File |
|--------|--------|------|
| Ubuntu | 400 (Regular) | `Ubuntu-Regular.ttf` |
| Ubuntu | 700 (Bold) | `Ubuntu-Bold.ttf` |
| Ubuntu Mono | 400 | `UbuntuMono-Regular.ttf` |

License: Ubuntu Font License 1.0 (free for redistribution — include
`UFL-1.0.txt` next to the fonts).

Loaded once in `tokens.css`:

```css
@font-face {
  font-family: 'Ubuntu';
  font-weight: 400;
  src: url('./fonts/Ubuntu-Regular.ttf') format('truetype');
  font-display: block;
}
@font-face {
  font-family: 'Ubuntu';
  font-weight: 700;
  src: url('./fonts/Ubuntu-Bold.ttf') format('truetype');
  font-display: block;
}
@font-face {
  font-family: 'Ubuntu Mono';
  font-weight: 400;
  src: url('./fonts/UbuntuMono-Regular.ttf') format('truetype');
  font-display: block;
}

:root {
  --font-sans: 'Ubuntu', system-ui, sans-serif;
  --font-mono: 'Ubuntu Mono', ui-monospace, monospace;
}
```

## Type scale (slide-content scope, applied at 1920×1080)

| Use | Size | Weight |
|-----|------|--------|
| Eyebrow label | 28px | 700 |
| Slide title | 72px | 700 |
| Slide subtitle | 36px | 400 |
| Body text | 28px | 400 |
| Code block | 26px | 400 (mono) |
| Caption / footer | 20px | 400 |

## Color palette

Mirror the main Lovable app's HSL semantic tokens for visual continuity. Defined
in `tokens.css` as CSS custom properties (HSL components, not full strings — so
they compose with opacity).

```css
:root {
  /* surface */
  --bg:            222 47% 6%;      /* deep slate */
  --bg-raised:     222 40% 10%;
  --fg:            210 40% 96%;
  --muted-fg:      215 20% 65%;
  --border:        217 33% 17%;

  /* brand */
  --primary:       217 91% 60%;     /* electric blue */
  --primary-fg:    222 47% 6%;
  --accent:        160 84% 39%;     /* mint — used for ✅ AFTER */
  --destructive:   0 84% 60%;       /* red — used for ❌ BEFORE */

  /* code highlight */
  --code-bg:       222 47% 9%;
  --code-line-add:    160 84% 39% / 0.15;
  --code-line-remove:   0 84% 60% / 0.15;
  --code-line-highlight: 217 91% 60% / 0.18;
}

@media (prefers-color-scheme: light) {
  :root {
    --bg:        210 40% 98%;
    --bg-raised: 0 0% 100%;
    --fg:        222 47% 11%;
    --muted-fg:  215 16% 47%;
    --border:    214 32% 91%;
    --code-bg:   210 40% 96%;
  }
}
```

## Spacing scale

| Token | Value | Use |
|-------|-------|-----|
| `--space-1` | 8px | inline |
| `--space-2` | 16px | between text lines |
| `--space-3` | 24px | between elements in a column |
| `--space-4` | 40px | between code blocks |
| `--space-5` | 64px | between title and body |
| `--space-6` | 96px | slide outer padding |

## Slide canvas

- Fixed authoring resolution: **1920 × 1080**
- Outer padding: `var(--space-6)` on all sides → content area = 1728 × 888
- Background: `hsl(var(--bg))`
- All slides scale via `transform: scale(...)` per the slides-app context (see
  the global `slides-app` context block — `ScaledSlide` component pattern).

## Iconography

Use **Lucide** icons (already familiar to the author). Bundle only the
specific icons used (tree-shaken). No icon font.

## Cross-references

- Slide authoring: [03-slide-authoring.md](./03-slide-authoring.md)
- Animation: [05-animation-primitives.md](./05-animation-primitives.md)

---

## Verification & Acceptance Criteria

### AC-CG-SLIDE-004: Design Tokens, Typography, and Palette Standards

**Given** Design tokens, font definitions, and layout styles in `slides-app/src/styles/tokens.css`.
**When** Style linters and relative path auditors validate the styling assets and font bundle references.
**Then** All font assets resolve locally, color tokens compose via HSL custom properties, and zero absolute paths exist in the spec or style declarations.

**Verification command:**
```bash
python linter-scripts/check-relative-paths.py
```
**Expected:** exit 0. Zero absolute paths or file:/// URIs.
