import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 33: DRY extract-now.
 *
 * Source: spec/17/31 line 33. Duplicate logic across two sites means
 * extract it now, not later. Two is the trigger, not three.
 */

const BEFORE = `// src/features/checkout/CartSummary.tsx
const subtotal = items.reduce((sum, i) => sum + i.price * i.qty, 0);
const tax = Math.round(subtotal * 0.2 * 100) / 100;
const shipping = subtotal > 50 ? 0 : 5;
const total = subtotal + tax + shipping;

// src/features/orders/InvoicePreview.tsx  (patched last week, drifted)
const subtotal = items.reduce((sum, i) => sum + i.price * i.qty, 0);
const tax = subtotal * 0.2;                       // ❌ rounding forgotten
const shipping = subtotal >= 50 ? 0 : 5;          // ❌ boundary flipped
const total = subtotal + tax + shipping;`;

const AFTER = `// src/features/pricing/computeOrderTotals.ts
export function computeOrderTotals(items: LineItem[]): OrderTotals {
  const subtotal = items.reduce((sum, i) => sum + i.price * i.qty, 0);
  const tax = roundCurrency(subtotal * TAX_RATE);
  const shipping = subtotal > FREE_SHIPPING_THRESHOLD ? 0 : FLAT_SHIPPING;

  return { subtotal, tax, shipping, total: subtotal + tax + shipping };
}

// CartSummary.tsx
const totals = computeOrderTotals(items);

// InvoicePreview.tsx
const totals = computeOrderTotals(items);`;

export default function DryExtractNowSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 33 · DRY extract-now"
      title="Two sites is the trigger. Not three."
      subtitle="The second time you paste the same three lines, stop. Extract to a named helper before landing the PR. Waiting for a third occurrence is how tax rounding drifts and shipping thresholds flip."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="typescript"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ Same math in two files. One was patched, the other drifted (rounding lost, boundary flipped)."
          afterLabel="✅ One computeOrderTotals with named constants. Both call sites call the helper."
        />
        <ActionPanel
          slideId="31-dry-extract-now"
          symptom="A tax fix lands in `CartSummary.tsx` but `InvoicePreview.tsx` still shows the un-rounded value because the same subtotal-plus-tax-plus-shipping block was copy-pasted. QA catches it three sprints later on a customer refund. Nobody remembered the second copy existed."
          rule="Duplicate logic across two sites is the trigger to extract, not three. Extract to a named helper in the correct feature folder, replace both call sites, and delete the inline copies in the same PR. Magic numbers move to named constants at the same time. Per spec/17/31 line 33 and rule 7."
          doThis="On second write, stop and extract. Name the helper by the verb it performs (`computeOrderTotals`, `formatInvoiceLine`), colocate with its feature, and export a named type for the return value. If the two copies have small differences, parameterize; if they diverge in intent, name them separately (`computeCartTotals` vs `computeInvoiceTotals`) so future drift is deliberate."
        />
      </div>
    </SlideLayout>
  );
}
