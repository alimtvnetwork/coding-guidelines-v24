import { SlideLayout } from "@/components/SlideLayout";
import { CodeDiff } from "@/components/CodeDiff";
import { ActionPanel } from "@/components/ActionPanel";

const BEFORE = `try {
  await charge(order);
} catch (err) {
  throw err;
  // → "Error: ECONNREFUSED" — no idea where it came from
}`;

const AFTER = `try {
  await charge(order);
} catch (err) {
  throw AppError.wrap(err, 'PaymentService.Charge', { OrderId });
  // → { Code, File, Line, Message, Cause, OrderId }
}`;

export default function AppErrorWrapperSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 04 · Errors"
      title="Wrap every catch with AppError + file:line context"
      subtitle="A naked re-throw loses the trail. AppError.wrap pins the failure to a code location and preserves the cause."
    >
      <CodeDiff
        language="typescript"
        before={BEFORE}
        after={AFTER}
        beforeLabel="❌ Naked re-throw"
        afterLabel="✅ Wrapped with context"
      />
      <ActionPanel
        symptom="Production logs show `ECONNREFUSED` with no idea which call, which order, or which file raised it."
        rule="Every catch block re-throws via AppError.wrap(err, 'Component.Action', { context }). Never `throw err;` alone."
        doThis="Find one `throw err` or empty catch in your service. Replace with AppError.wrap including the relevant Id."
      />
    </SlideLayout>
  );
}
