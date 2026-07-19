import { SlideLayout } from "@/components/SlideLayout";
import { PrincipleCard } from "@/components/PrincipleCard";

export default function CorePrinciples2Slide() {
  return (
    <SlideLayout
      eyebrow="Core Development Principles · 2 of 3"
      title="Make intent unambiguous"
      subtitle="Booleans read like sentences. Errors carry their own evidence."
    >
      <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", gap: 36, marginTop: 16 }}>
        <PrincipleCard
          number="03"
          title="Positively Named Guards"
          symptom="Names like `!isNotReady` or `!hasNoError` force the reader to invert booleans in their head."
          rule="Guards are affirmative: `isReady`, `hasError`, `canPublish`. Never a negated negative."
          action="Rename every `!isNot*` / `!has No*` in your diff. Flip the condition to match the new name."
          accent="primary"
          delay={0.5}
        />
        <PrincipleCard
          number="04"
          title="Structured Error Wrapping"
          symptom="Errors get swallowed by bare `catch {}` or re-thrown as strings, losing stack and context."
          rule="Every error that crosses a boundary is an `AppError` with stack trace and structured context."
          action="Search your PR for `catch` blocks with no `AppError`. Wrap each with context and rethrow."
          accent="primary"
          delay={0.65}
        />
      </div>
    </SlideLayout>
  );
}
