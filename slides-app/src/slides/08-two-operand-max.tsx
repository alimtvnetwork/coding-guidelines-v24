import { SlideLayout } from "@/components/SlideLayout";
import { CodeDiff } from "@/components/CodeDiff";
import { ActionPanel } from "@/components/ActionPanel";

const BEFORE = `if (
  user.IsActive &&
  order.HasItems &&
  payment.IsAuthorized &&
  !order.IsLocked
) { /* ship */ }`;

const AFTER = `if (CanShipOrder(user, order, payment)) { /* ship */ }

function CanShipOrder(user, order, payment) {
  if (!user.IsActive)         return false;
  if (!order.HasItems)        return false;
  if (!payment.IsAuthorized)  return false;
  return !order.IsLocked;
}`;

export default function TwoOperandMaxSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 08 · Code Red"
      title="Cap every condition at two operands"
      subtitle="A 4-operand chain hides intent behind precedence. Extract a positively-named helper and let it explain itself."
    >
      <CodeDiff
        language="typescript"
        before={BEFORE}
        after={AFTER}
        beforeLabel="❌ 4-operand chain"
        afterLabel="✅ Named helper"
      />
      <ActionPanel
        symptom="A wall of && and || conditions. The reader has to mentally evaluate the chain to know what 'ship' means."
        rule="No more than 2 operands in a single boolean expression. The third operand is the signal to extract a helper."
        doThis="Find the longest && / || chain in your diff. Extract it into a `Can…/Is…` function that returns a boolean."
      />
    </SlideLayout>
  );
}
