import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 34: Mermaid-first for 3+ components.
 *
 * Source: spec/17/31 line 42 (rule 10). Any feature with three or
 * more components gets a Mermaid component diagram before code.
 */

const BEFORE = `// Day 1: jumped straight to code. No diagram.
// CheckoutPage.tsx  (renders everything, 380 lines)
//   ├── CartList (owns items state)
//   ├── AddressForm (owns address, also fetches saved addresses)
//   ├── PaymentPanel (owns card state, also owns totals recompute)
//   └── OrderSummary (re-fetches items to recompute totals)
//
// Result: totals live in two places, address fetch races payment mount,
// three components own overlapping slices of the same cart state.`;

const AFTER = `%% docs/checkout/components.mmd  (committed before any .tsx)
flowchart TD
  Page[CheckoutPage<br/>route + data loader]
  Store[(useCheckoutStore<br/>cart, address, totals)]
  Cart[CartList<br/>read-only]
  Addr[AddressForm<br/>read + update]
  Pay[PaymentPanel<br/>read totals only]
  Sum[OrderSummary<br/>read totals only]

  Page --> Store
  Store --> Cart
  Store --> Addr
  Store --> Pay
  Store --> Sum

%% Reviewed in PR #0. Only then do the four .tsx files land.`;

export default function MermaidFirstComponentsSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 34 · Mermaid-first for 3+ components"
      title="Draw the graph. Then write the components."
      subtitle="Three or more components in one feature is the trigger. Commit a Mermaid component diagram to the feature folder before any .tsx lands. The diagram is the contract that says who owns state, who reads it, and who is forbidden from talking to whom."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="text"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ Four components shipped without a diagram. State ownership drifts, totals live in two places, mounts race."
          afterLabel="✅ components.mmd lands first. One store owns state, leaves are read-only. Reviewers approve the graph before the code."
        />
        <ActionPanel
          slideId="32-mermaid-first-components"
          symptom="A checkout PR lands four new components in one folder. Two weeks later, totals disagree between `PaymentPanel` and `OrderSummary` because both recompute from their own slice of state, and nobody can point to a single diagram that says which component owns what. Refactor takes a full sprint because the ownership was never written down."
          rule="Any feature with three or more components requires a Mermaid component diagram committed to the feature folder (for example `docs/checkout/components.mmd`) before any `.tsx` is written. The diagram names each component, marks the state owner, and draws arrows for data flow. PRs that add a third component without updating or adding the diagram fail review. Per spec/17/31 line 42 rule 10."
          doThis="On the third component, stop coding. Open a `.mmd` file in the feature folder, sketch the graph as `flowchart TD`, mark the store or hook that owns state, and draw arrows only for reads and writes that actually exist. Open the PR with the diagram alone first, get review on ownership, then land the components. Re-render the PNG via `npm run diagrams:render` so the diagram stays visible in the readme."
        />
      </div>
    </SlideLayout>
  );
}
