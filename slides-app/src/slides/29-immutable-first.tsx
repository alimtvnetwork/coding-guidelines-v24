import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 31: Immutable-first (Rust-style) assignment.
 *
 * Source: spec/17/31 line 35. Assign every variable once at declaration.
 * Never reassign except loop indices. Prefer const/final/val. Build result
 * objects with spread or copy, never in-place mutation.
 */

const BEFORE = `let totals = { subtotal: 0, tax: 0, total: 0 };
totals.subtotal = sumItems(order.items);
totals.tax = totals.subtotal * 0.2;
totals.total = totals.subtotal + totals.tax;

let items = order.items;
items.push(giftCard);
items.sort((a, b) => a.price - b.price);

let label = "pending";
if (order.isPaid) {
  label = "paid";
}`;

const AFTER = `const subtotal = sumItems(order.items);
const tax = subtotal * 0.2;
const totals = { subtotal, tax, total: subtotal + tax };

const items = [...order.items, giftCard]
  .slice()
  .sort((a, b) => a.price - b.price);

const label = order.isPaid ? "paid" : "pending";`;

export default function ImmutableFirstSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 31 · Immutable-first"
      title="Assign once. Copy, do not mutate."
      subtitle="Prefer const, final, val. Reassignment is a smell; in-place push/splice/sort on shared arrays is a bug waiting to happen. Loop indices are the only sanctioned exception."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="typescript"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ let with staged reassignment, in-place mutation of a shared array, late rebinding"
          afterLabel="✅ const at declaration, spread-copy before sort, ternary picks the final value once"
        />
        <ActionPanel
          slideId="29-immutable-first"
          symptom="A pricing helper reassigns `totals` field by field, then a caller mutates `order.items` in place with `push` and `sort`. Two screens later, an unrelated component re-renders with a re-sorted list because it shared the same array reference. The bug reproduces only after a specific click order and takes an afternoon to trace."
          rule="Assign every variable once at declaration with `const` (TS/JS), `final` (Dart/Java), `val` (Kotlin/Scala), or `let` without `mut` (Rust). Never reassign except loop indices. Build result objects with spread or `structuredClone`; never mutate arrays or objects returned by hooks, props, or shared state. Per spec/17/31 line 35."
          doThis="Compute the final value first, then bind it once. Replace `let x; if (...) x = a; else x = b;` with a ternary or an extracted helper. Before calling `sort`, `reverse`, or `splice`, copy with `[...arr]` or `arr.slice()`. If you genuinely need mutation for perf (very large arrays in a tight loop), isolate it inside one function and comment why."
        />
      </div>
    </SlideLayout>
  );
}
