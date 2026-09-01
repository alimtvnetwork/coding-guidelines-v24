import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 30: Function length 8/15 rule.
 *
 * Source: 02-spec/17/31 line 25. 8 lines preferred, 15 lines hard cap. Skip
 * blank lines and comments when counting. Waiver only via inline comment
 * `// lint-allow: function-length reason="..." max=N`.
 */

const BEFORE = `function processOrder(order: Order) {
  if (!order) return null;
  const items = order.items ?? [];
  let subtotal = 0;
  for (const item of items) {
    subtotal += item.price * item.qty;
  }

  const tax = subtotal * 0.2;
  const shipping = subtotal > 50 ? 0 : 5;
  const discount = order.coupon ? applyCoupon(order.coupon, subtotal) : 0;
  const total = subtotal + tax + shipping - discount;
  logger.info("order.priced", { id: order.id, total });
  const invoice = buildInvoice(order, total);
  sendEmail(order.email, invoice);
  await saveInvoice(invoice);
  await markPaid(order.id);

  return invoice;
}`;

const AFTER = `function processOrder(order: Order) {
  const totals = priceOrder(order);
  const invoice = buildInvoice(order, totals.total);

  return finalizeInvoice(order, invoice);
}

function priceOrder(order: Order): OrderTotals {
  const subtotal = sumItems(order.items ?? []);
  const discount = order.coupon ? applyCoupon(order.coupon, subtotal) : 0;

  return computeTotals(subtotal, discount);
}`;

export default function FunctionLengthSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 30 · Function length"
      title="8 lines preferred. 15 is the hard cap."
      subtitle="Count executable lines only; skip blanks and comments. Past 15 you must extract, or add the exact waiver comment with a reason. `TODO: refactor later` is not a waiver."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="typescript"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ 16 executable lines, 4 responsibilities in one function"
          afterLabel="✅ 4 line orchestrator + 4 line priceOrder, each under the 8 line preferred bar"
        />
        <div
          style={{
            background: "hsl(var(--muted))",
            border: "1px dashed hsl(var(--border))",
            borderRadius: 12,
            padding: "16px 22px",
            fontFamily: "ui-monospace, SFMono-Regular, Menlo, monospace",
            fontSize: 24,
          }}
        >
          // lint-allow: function-length reason="state machine dispatch table" max=22
        </div>
        <ActionPanel
          slideId="28-function-length"
          symptom="A function grows to 40 lines mixing pricing, invoicing, email, and persistence. Reviewers cannot hold it in their head; a bug fix in one branch silently breaks another. Nobody notices the length because the file compiles fine."
          rule="8 lines preferred, 15 lines hard cap, counting only executable statements (skip blank lines and comment lines). Extract cohesive blocks into named helpers. The only escape is an inline waiver on the line above the signature using the exact form `// lint-allow: function-length reason='...' max=N` with a real reason and an explicit max. Per 02-spec/17/31 line 25."
          doThis="When a function passes 15, split it around its verbs: pricing, formatting, side effects. Give each helper an intention-revealing name so the orchestrator reads like a paragraph. Use the waiver only for genuinely irreducible cases (dispatch tables, generated code), never for `it's clearer this way`."
        />
      </div>
    </SlideLayout>
  );
}
