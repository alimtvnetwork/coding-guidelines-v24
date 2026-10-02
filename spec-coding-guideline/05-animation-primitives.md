# Slide Animation Primitives (AI Execution Prompt)

> **/goal** Define and govern purposeful, comprehension-driven animation primitives, timing budgets, and reduced-motion fallbacks for slide deck transitions.
> **/learn** Master the 5 animation primitives (entrance stagger, line highlight, before/after morph, strikethrough, bullet stagger), Framer Motion integration, and total auto-motion budgets (≤2.2s).

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Ensure every animation serves instructional comprehension rather than decorative motion.
- [ ] `/learn` Enforce the 2.2-second maximum auto-revealed motion budget per slide so presenters are never stalled.
- [ ] `/goal` Verify layout transitions and morph animations maintain parity between Framer Motion and Remotion GIF pipelines.
- [ ] `/learn` Enforce mandatory `@media (prefers-reduced-motion: reduce)` accessibility overrides to instantly render final states.

. **CRITICAL AI INSTRUCTION:** This specification is an active AI execution directive. All code generated or modified must strictly follow the rules below.

---

**Version:** 1.0.0

## Philosophy

Every animation in this deck must serve **comprehension**, not decoration. Each
primitive maps to a learning moment:

| Primitive | Teaches |
|-----------|---------|
| Entrance stagger | "Read in this order" |
| Code line highlight | "These specific lines are the issue" |
| Before → After morph | "The fix is a transformation, not a rewrite" |
| Diff strikethrough | "This goes away" |

## Library choice

Use **Framer Motion** (`framer-motion@^11`) for the deck UI. It tree-shakes
well, handles enter/exit cleanly, and supports layout animations (needed for
the before/after morph).

For frame-perfect GIF rendering (separate pipeline — see
[07-gif-generation.md](./07-gif-generation.md)) we use Remotion. The **same
visual result** must be achievable in both — so all motion is defined as plain
keyframe values, not Framer-specific physics.

## Primitive 1 — Entrance stagger

Each top-level child of `<SlideLayout>` fades up with a 60ms stagger.

```tsx
// SlideLayout internals
<motion.div
  initial={{ opacity: 0, y: 24 }}
  animate={{ opacity: 1, y: 0 }}
  transition={{ duration: 0.4, delay: index * 0.06, ease: [0.16, 1, 0.3, 1] }}
/>
```

## Primitive 2 — Code line highlight

`<CodeDiff>` accepts `highlightLines={{ before: [3,4], after: [2] }}`. After the
slide enters, the highlighted lines pulse with a colored background:

- Before block: `hsl(var(--destructive) / 0.15)` glow
- After block: `hsl(var(--accent) / 0.15)` glow

Pulse animation: 1.5s, 2 iterations, then settle to a static 0.08-opacity tint.

```css
@keyframes pulse-line {
  0%   { background: hsl(var(--code-line-highlight)); }
  50%  { background: hsl(var(--code-line-highlight) / 2); }
  100% { background: hsl(var(--code-line-highlight) / 0.4); }
}
```

## Primitive 3 — Before → After morph

`<CodeDiff>` with `morphOnReveal={true}` shows only the BEFORE block, then on
spacebar (or after 4s if `autoAdvance`) crossfades to AFTER while the labels
swap. Optional: lines that are unchanged stay in place using Framer's `layoutId`
to give a "lines reorganize" effect.

Two layouts:

- **side-by-side** (default for ≤6-line samples) — both visible, AFTER fades in
  from `opacity: 0.3` to `1` while BEFORE dims to `0.5`
- **stacked** with morph (for ≥7-line samples) — cleaner on small screens

## Primitive 4 — Diff strikethrough

Lines marked as `removed` get a red strikethrough animation that draws
left-to-right over 0.4s before the line itself collapses (height → 0) over
0.3s.

## Primitive 5 — Bullet stagger

`<BulletList>` items fade in from `x: -16` with 80ms stagger. A small
`CheckCircle` icon (Lucide, mint color) pops in via spring (`scale: 0 → 1`,
`stiffness: 300, damping: 20`) 200ms after each bullet.

## Timing budget per slide

| Phase | Duration |
|-------|----------|
| Entrance stagger (eyebrow → title → subtitle → body) | 0–0.7s |
| Code highlight pulse | 0.7s–2.2s |
| (Optional) Before → After morph on user advance | +0.6s |

Total auto-revealed motion ≤ 2.2s — trainer should not have to wait.

## Reduced motion

`@media (prefers-reduced-motion: reduce)` disables all transitions; everything
appears in its final state instantly. This is mandatory for accessibility.

## Cross-references

- GIF generation reuses these timings: [08-gif-generation.md](./08-gif-generation.md)
- Slide authoring: [03-slide-authoring.md](./03-slide-authoring.md)

---

## Verification & Acceptance Criteria

### AC-CG-SLIDE-005: Frame Animation Primitives and Motion Standards

**Given** Slide animation definitions and transitions implemented using Framer Motion or CSS keyframes.
**When** Animation parameters, timing budgets, and accessibility queries are audited across slide specs and components.
**Then** Auto-revealed motion stays within the ≤2.2s budget, reduced motion is respected, and all relative paths conform to repository standards.

**Verification command:**
```bash
python linter-scripts/check-relative-paths.py
```
**Expected:** exit 0. Zero absolute paths or file:/// URIs.
