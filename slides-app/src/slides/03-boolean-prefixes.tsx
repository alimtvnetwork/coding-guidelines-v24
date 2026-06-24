import { SlideLayout } from "@/components/SlideLayout";
import { CodeDiff } from "@/components/CodeDiff";
import { ActionPanel } from "@/components/ActionPanel";

const BEFORE = `const active = true;
const items = order.items.length > 0;
function valid(x) {
  return x != null;
}`;

const AFTER = `const IsActive = true;
const HasItems = order.items.length > 0;
function IsValid(x) {
  return x != null;
}`;

export default function BooleanPrefixesSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 03 · Naming"
      title="Prefix every boolean with Is, Has, Can or Should"
      subtitle="A boolean's name should read like a yes/no question. No more guessing whether `active` is a flag or a user."
    >
      <CodeDiff
        language="typescript"
        before={BEFORE}
        after={AFTER}
        beforeLabel="❌ Ambiguous"
        afterLabel="✅ Self-describing"
      />
      <ActionPanel
        symptom="`active`, `valid`, `items` could be flags, strings or arrays — type comes only from reading the assignment."
        rule="Every boolean variable and predicate starts with Is / Has / Can / Should. No exceptions."
        doThis="Grep your file for `const ` followed by a one-word name. Rename any boolean to Is/Has/Can/Should + Noun."
      />
    </SlideLayout>
  );
}
