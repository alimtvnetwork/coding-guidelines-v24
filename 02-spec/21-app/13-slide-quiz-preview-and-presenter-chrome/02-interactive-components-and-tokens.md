# 02 — Interactive Components & Tokens

> **Canonical numbers:** `02-spec/07-design-system/42-slide-quiz-preview-chrome-and-default-shadows.md`. This file explains how quiz UI maps onto slide products.

---

## 1. Presentation option card

The exam product uses class **`presentation-option-card`** on `<button>` elements (single- and multi-select). Selected state sets inline `opacity`, `backgroundColor`, `border`, and optional `boxShadow` glow; hover uses **global CSS** (transform + shadow + gradient).

Implementers MUST NOT change hover physics without updating file 42:

| State | Behavior |
|:---|:---|
| Rest | `opacity: 0.82`, default elevation token |
| Hover | `opacity: 1`, `translate3d(6px, 0, 0)`, primary border at 75% alpha, horizontal primary wash, hover elevation token |
| Selected | Full opacity, active card bg/border from theme map, optional `0 0 16px` border-color glow at 27% alpha |

Badge column (`option-badge`): `36px`–`40px` square, `rounded-xl`, mono letter `A`–`Z`. On hover, badge scales `1.05` and picks up primary tint.

Label column: class **`option-text-shadow`** for default text-shadow (rest) and heavier offset shadow on hover.

---

## 2. Checkbox vs card selection

Two patterns coexist:

| Pattern | Control | When |
|:---|:---|:---|
| **Card button** | Whole row is `<button type="button">` | MCQ and multiselect in focus/presentation runners |
| **Radix checkbox** | `Checkbox` primitive `h-4 w-4`, `rounded-sm`, `border-primary` | Forms admin tables, sortable field cards |

Multiselect in presentation mode still uses **card toggles**, not visible checkboxes, with **`CheckCircle2`** icon on the trailing edge when selected. Do not mix a native `<input type="checkbox">` inside the card button (double activation).

Admin checkbox reference (copy-safe for Gemini):

```tsx
import * as CheckboxPrimitive from "@radix-ui/react-checkbox";
import { Check } from "lucide-react";

export function Checkbox(props: CheckboxPrimitive.CheckboxProps) {
  return (
    <CheckboxPrimitive.Root
      className="peer h-4 w-4 shrink-0 rounded-sm border border-primary ring-offset-background data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
      {...props}
    >
      <CheckboxPrimitive.Indicator className="flex items-center justify-center text-current">
        <Check className="h-4 w-4" />
      </CheckboxPrimitive.Indicator>
    </CheckboxPrimitive.Root>
  );
}
```

---

## 3. Botanical-light theme (improved green)

Legacy measured values used `#16A34A` primary on `#F4F8F5` ground with `#DCFCE7` active fill. File 42 **tightens contrast** and reduces neon glow:

- Primary hue stays **142**; lightness **38%** (was 45%) for AA-friendly text pairs.
- Active fill uses **142 48% 94%** instead of pure mint blocks.
- Hover shadow primary alpha **0.24** (was 0.28) to avoid muddy halos on white cards.
- Background radial accents use **teal 160** at ≤ **4%** alpha (measured pattern from green-choice background gradients).

Deck slides that need the same green MUST register **`botanical-light`** in theme switch data or map it to an existing light theme slot—do not hardcode hex in TSX.

---

## 4. Motion

| Token | Value |
|:---|:---|
| Card transition | `220ms cubic-bezier(0.16, 1, 0.3, 1)` on transform; `220ms ease` on box-shadow |
| Text shadow transition | `200ms ease` |
| Stagger classes | `stagger-1` … `stagger-6` with `0.06s` step (exam `theme.css`) |
| Reduced motion | `@media (prefers-reduced-motion: reduce)` forces `0.01ms` transitions and `transform: none` |

Slide deck transitions remain **`0.45s`** / `cubic-bezier(0.4, 0, 0.2, 1)` per `29-slide-navigation-and-builder.md`.

---

## 5. Images and plates

Any hero image, avatar plate, or media frame on quiz or slide surfaces SHOULD use the same elevation tokens as option cards:

- Rest: `--elevation-rest`
- Hover/focus-visible: `--elevation-hover`
- Optional focus ring: `0 0 0 3px hsl(var(--ring) / 0.35)`

See file 42 section 2 for literal shadow strings.
