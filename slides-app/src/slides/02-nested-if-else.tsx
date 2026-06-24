import { SlideLayout } from "@/components/SlideLayout";
import { CodeDiff } from "@/components/CodeDiff";
import { ActionPanel } from "@/components/ActionPanel";

const BEFORE = `function processOrder(order) {
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

const AFTER = `function processOrder(order) {
  if (!IsValidOrder(order)) return null;
  if (!HasActiveUser(order)) return null;
  if (!HasItems(order))      return null;
  return charge(order);
}`;

export default function NestedIfElseSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 02 · Code Red"
      title="Flatten every if-pyramid into guard clauses"
      subtitle="One indentation level per function. Each precondition exits early with its own positively-named guard."
    >
      <CodeDiff
        language="typescript"
        before={BEFORE}
        after={AFTER}
        beforeLabel="❌ 4-level pyramid"
        afterLabel="✅ Flat guard clauses"
        layout="stacked"
      />
      <ActionPanel
        symptom="Logic is buried 4 levels deep. To read the happy path you scroll past every error case."
        rule="Zero nested if-else. Invert each condition into an early return so the happy path stays flat."
        doThis="Open the deepest function in your diff. Replace its outer if with a `return early` guard."
      />
    </SlideLayout>
  );
}
