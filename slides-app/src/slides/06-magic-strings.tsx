import { SlideLayout } from "@/components/SlideLayout";
import { CodeDiff } from "@/components/CodeDiff";
import { ActionPanel } from "@/components/ActionPanel";

const BEFORE = `if (order.status === 'shipped') { /* ... */ }
if (order.status === 'shippped') { /* typo */ }
// silent bug — string typo never caught`;

const AFTER = `enum OrderStatus {
  Pending = 'Pending',
  Shipped = 'Shipped',
  Delivered = 'Delivered',
}
if (order.Status === OrderStatus.Shipped) { /* ... */ }
// compiler catches typos at build time`;

export default function MagicStringsSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 06 · Code Red"
      title="Promote repeated literals to an enum or const"
      subtitle="If a string literal appears in two places, the compiler can no longer protect you. Promote it once and reuse."
    >
      <CodeDiff
        language="typescript"
        before={BEFORE}
        after={AFTER}
        beforeLabel="❌ Repeated literals"
        afterLabel="✅ Enum-checked"
      />
      <ActionPanel
        symptom="A one-character typo in a status string ships to production and silently breaks a branch."
        rule="Never compare against the same string literal twice. The second occurrence is the trigger to extract an enum."
        doThis="Pick the most-compared string in your diff (status, kind, type) and replace usages with an enum value."
      />
    </SlideLayout>
  );
}
