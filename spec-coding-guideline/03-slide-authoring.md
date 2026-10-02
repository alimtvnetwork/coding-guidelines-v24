# Slide Authoring Guidelines (AI Execution Prompt)

> **/goal** Standardize the authoring of interactive presentation slides with component contracts, mandatory before/after visual contrast, strict code sizing limits, and offline-safe local assets.
> **/learn** Enforce the one-file-per-slide model, `<SlideLayout>` and `<CodeDiff>` component usage, 12-line code sample maximums, and zero external CDN dependencies.

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Verify each slide is authored as an isolated component under `slides-app/src/slides/NN-topic-name.tsx` and registered in `deck.ts`.
- [ ] `/learn` Enforce visual contrast (before vs. after) on every technical slide; pure-text training slides are strictly forbidden.
- [ ] `/goal` Cap code blocks at a maximum of 12 lines per `before` or `after` block to guarantee readability on 1920×1080 canvas.
- [ ] `/learn` Verify zero external CDN URLs or remote asset calls; all images, fonts, and dependencies must resolve locally from `public/`.

. **CRITICAL AI INSTRUCTION:** This specification is an active AI execution directive. All code generated or modified must strictly follow the rules below.

---

**Version:** 1.0.0

## One file = one slide

Each slide is a default-exported React component under `slides-app/src/slides/`,
named with the convention `NN-topic-name.tsx` (numeric prefix sets order).

```tsx
// slides-app/src/slides/02-nested-if-else.tsx
import { SlideLayout } from '@/components/SlideLayout';
import { CodeDiff } from '@/components/CodeDiff';

export default function NestedIfElseSlide() {
  return (
    <SlideLayout
      eyebrow="Code Red"
      title="Nested if-else → zero nesting"
      subtitle="Replace pyramid logic with positively named guards"
    >
      <CodeDiff
        language="typescript"
        before={BEFORE_CODE}
        after={AFTER_CODE}
        beforeLabel="❌ Pyramid (avoid)"
        afterLabel="✅ Guard clauses"
        highlightLines={{ before: [3, 4, 5, 6], after: [2, 3, 4] }}
      />
    </SlideLayout>
  );
}

const BEFORE_CODE = `function processOrder(order) {
  if (order) {
    if (order.user) {
      if (order.user.isActive) {
        if (order.items.length > 0) {
          return charge(order);
        }
      }
    }
  }
  return null;
}`;

const AFTER_CODE = `function processOrder(order) {
  if (!isValidOrder(order)) return null;
  if (!hasActiveUser(order)) return null;
  if (!hasItems(order)) return null;
  return charge(order);
}`;
```

## Required props for `SlideLayout`

| Prop | Type | Required | Purpose |
|------|------|----------|---------|
| `eyebrow` | `string` | optional | Small uppercase label above title (e.g. "Code Red", "Naming") |
| `title` | `string` | **yes** | Main slide title |
| `subtitle` | `string` | optional | One-line context |
| `notes` | `string` | optional | Trainer notes (hidden in normal view, visible in presenter view) |
| `children` | `ReactNode` | **yes** | Slide body |

## Slide registration

```ts
// slides-app/src/deck.ts
import Title from './slides/00-title';
import Naming from './slides/01-naming-conventions';
import NestedIf from './slides/02-nested-if-else';
// ...
export const DECK = [
  { id: '00-title', component: Title },
  { id: '01-naming', component: Naming },
  { id: '02-nested-if', component: NestedIf },
  // ...
] as const;
```

The deck shell reads `DECK`, renders the current slide based on `location.hash`
(`#/3` → index 3), and handles keyboard navigation.

## Slide body building blocks

The library exposes three primary content components — every slide should use at
least one:

### 1. `<CodeDiff>` — side-by-side or stacked code comparison

```tsx
<CodeDiff
  language="typescript"
  before={CODE_BEFORE}
  after={CODE_AFTER}
  beforeLabel="❌ Before"
  afterLabel="✅ After"
  layout="side-by-side"   // or "stacked"
  highlightLines={{ before: [3,4], after: [2] }}
  morphOnReveal={true}    // animates the before → after transition
/>
```

### 2. `<BeforeAfter>` — generic split-screen with morph

For non-code comparisons (e.g. error trace before/after AppError wrapping).

```tsx
<BeforeAfter
  beforeLabel="Without AppError"
  afterLabel="With AppError"
  before={<RawStackTrace />}
  after={<StructuredTrace />}
/>
```

### 3. `<BulletList>` — animated staggered bullets

```tsx
<BulletList items={[
  'Errors carry file:line context',
  'Wrapping preserves stack',
  'Logging is one line, structured',
]} />
```

## Hard rules

1. **Every slide must have a before/after or a clear contrast** — this deck is a
   training tool. Pure-text slides are forbidden.
2. **Code samples must be syntactically valid** in the declared language —
   Shiki will render them at build time and will warn on invalid syntax.
3. **No external network calls** — no `<img src="https://...">`, no fonts from
   CDN, no embedded YouTube. Everything ships in `public/`.
4. **Max 12 lines of code per `before` or `after` block** — longer samples won't
   fit at 1920×1080 with readable type.
5. **One concept per slide** — split if you have two.

## Cross-references

- Animation primitives: [05-animation-primitives.md](./05-animation-primitives.md)
- Curriculum: [06-curriculum.md](./06-curriculum.md)

---

## Verification & Acceptance Criteria

### AC-CG-SLIDE-003: Slide Authoring Standards and Component Contracts

**Given** Interactive presentation slides created for the coding guideline curriculum in `slides-app/src/slides/`.
**When** Linters and component authoring checks inspect slide implementations, prop definitions, and asset links.
**Then** All slides adhere to the one-file-per-slide convention, utilize valid `<SlideLayout>` contracts, embed local assets, and maintain 100% relative repository paths.

**Verification command:**
```bash
python linter-scripts/check-relative-paths.py
```
**Expected:** exit 0. Zero absolute paths or file:/// URIs.
