import { SlideLayout } from "@/components/SlideLayout";
import { CodeDiff } from "@/components/CodeDiff";
import { ActionPanel } from "@/components/ActionPanel";

const BEFORE = `if (!isInvalid(input) && !isMissing(token)) {
  proceed();
}
// double-negation — readers must mentally flip every check`;

const AFTER = `if (IsValid(input) && IsPresent(token)) {
  proceed();
}
// reads naturally — every check states what we want`;

export default function PositivelyNamedGuardsSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 09 · Naming"
      title="Name guards for the positive case, negate at the call site"
      subtitle="`!isInvalid` is a riddle. `IsValid` reads like English. Let the caller decide whether to flip it."
    >
      <CodeDiff
        language="typescript"
        before={BEFORE}
        after={AFTER}
        beforeLabel="❌ Double-negation"
        afterLabel="✅ Positive guard"
      />
      <ActionPanel
        symptom="`!isInvalid(x) && !isMissing(y)` — readers run a NOT-NOT loop in their head before they understand the branch."
        rule="Predicates are always positive: IsValid, HasItems, CanShip. The caller is free to write `if (!IsValid)`."
        doThis="Rename one `isInvalid/isMissing/isBad`-style function in your diff to its positive form, update callers."
      />
    </SlideLayout>
  );
}
