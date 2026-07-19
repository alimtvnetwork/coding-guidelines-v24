import { SlideLayout } from "@/components/SlideLayout";
import { PrincipleCard } from "@/components/PrincipleCard";

export default function CorePrinciples1Slide() {
  return (
    <SlideLayout
      eyebrow="Core Development Principles · 1 of 3"
      title="Reviewable code starts here"
      subtitle="Every rule slide in this deck follows the same shape: Symptom, Rule, Do this next. These two set the shape."
    >
      <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", gap: 36, marginTop: 16 }}>
        <PrincipleCard
          number="01"
          title="Zero-Nesting Discipline"
          symptom="Logic buried 3+ levels deep. Reviewers scroll past every error case to find the happy path."
          rule="No nested if-else. Each precondition exits early with its own guard. One indentation level per body."
          action="Pick the deepest function in your next diff. Invert its outer if into an early return."
          accent="accent"
          delay={0.5}
        />
        <PrincipleCard
          number="02"
          title="Two-Operand Maximum"
          symptom="Boolean expressions with 3+ operands force reviewers to mentally evaluate the truth table."
          rule="At most two operands per boolean. Extract the third into a positively named guard function."
          action="Grep your PR for `&& .* &&` or `|| .* ||`. Extract each hit into a named predicate."
          accent="accent"
          delay={0.65}
        />
      </div>
    </SlideLayout>
  );
}
