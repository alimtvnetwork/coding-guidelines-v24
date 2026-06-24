import { SlideLayout } from "@/components/SlideLayout";
import { CodeDiff } from "@/components/CodeDiff";
import { ActionPanel } from "@/components/ActionPanel";

const BEFORE = `console.log(
  'failed to charge ' + orderId + ': ' + err.message
);`;

const AFTER = `logger.error('PaymentService.ChargeFailed', {
  File: __filename,
  OrderId,
  Cause: err.message,
});`;

export default function StructuredLoggingSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 05 · Errors"
      title="Log one structured event, never a concatenated string"
      subtitle="Logs must be grep-able and machine-readable. Always emit a named event with file, ids and cause as fields."
    >
      <CodeDiff
        language="typescript"
        before={BEFORE}
        after={AFTER}
        beforeLabel="❌ String concatenation"
        afterLabel="✅ Structured event"
      />
      <ActionPanel
        symptom="Free-text logs can't be filtered, aggregated or alerted on. Search for one order = scrolling forever."
        rule="One logger call per event. Event name in PascalCase + structured fields including File and the relevant Id."
        doThis="Replace every `console.log` in your diff with `logger.info|error` and put dynamic values as fields."
      />
    </SlideLayout>
  );
}
