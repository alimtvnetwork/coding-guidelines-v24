import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 25: Boolean naming prefixes + positive framing.
 *
 * Source: 02-spec/17/31 §"Boolean Naming" lines 40-49.
 *   - Prefixes: is / has / can / should / was / will / did / must.
 *   - Positive only: invert `isNotReady` to `isReady` and flip the check.
 *   - No bare adjectives (`enabled`), no `flag` / `bool` / `check`.
 */

const BEFORE = `const disabled = !user.active;
const notReady = status !== "ready";
const check = items.length > 0;
if (!disabled && !notReady && check) {
  submit();
}`;

const AFTER = `const isEnabled = user.active;
const isReady = status === "ready";
const hasItems = items.length > 0;
if (isEnabled && isReady && hasItems) {
  submit();
}`;

export default function BooleanNamingSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 25 · Boolean naming"
      title="`is` / `has` / `can` / `should`. Positive framing only."
      subtitle="Every boolean starts with an approved prefix. If the natural name is negative, invert it and flip the check site."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 18, marginTop: 4 }}>
        <CodeDiff
          language="typescript"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ negatives + bare adjectives + `check`"
          afterLabel="✅ positive `is*` / `has*`, readable guard"
        />
        <ActionPanel
          slideId="23-boolean-naming"
          symptom="A guard reads `if (!disabled && !notReady && check)`. Readers stall parsing double negatives; the next diff adds a fourth negation and the intent is lost."
          rule="Every boolean starts with `is`, `has`, `can`, `should`, `was`, `will`, `did`, or `must`. Positive framing only: no `isNotReady`, no `hasNoAccess`. No bare adjectives (`enabled`), no `flag` / `bool` / `check` names. Per 02-spec/17/31 §Boolean Naming lines 40-49."
          doThis="When you touch a negative or bare-adjective boolean, invert its definition to positive, rename with an approved prefix, and flip every call site in the same diff. Do not leave mixed conventions in the file."
        />
      </div>
    </SlideLayout>
  );
}
