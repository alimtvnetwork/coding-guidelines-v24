import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 26: No boolean flag parameters.
 *
 * Source: spec/17/31 line 48.
 *   "No boolean flag parameters on functions. Split into two named functions
 *    instead. `render(true)` is wrong, `renderExpanded()` and
 *    `renderCollapsed()` are right."
 */

const BEFORE = `function render(expanded: boolean) {
  if (expanded) {
    return <Panel rows={rows} />;
  }

  return <Summary rows={rows} />;
}

// Call sites read as riddles:
render(true);
render(false);
render(user.isAdmin);`;

const AFTER = `function renderExpanded() {
  return <Panel rows={rows} />;
}

function renderCollapsed() {
  return <Summary rows={rows} />;
}

// Call sites document intent:
renderExpanded();
renderCollapsed();
user.isAdmin ? renderExpanded() : renderCollapsed();`;

export default function NoBooleanParametersSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 26 · No boolean parameters"
      title="`render(true)` is a riddle. `renderExpanded()` is a name."
      subtitle="Boolean flag parameters hide branching behind an anonymous `true` / `false`. Split into two positively named functions and let the call site say what it wants."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 18, marginTop: 4 }}>
        <CodeDiff
          language="typescript"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ `render(true)` at every call site"
          afterLabel="✅ two named functions, self-documenting calls"
        />
        <ActionPanel
          slideId="24-no-boolean-parameters"
          symptom="You grep for `render(` and every call site says `render(true)` or `render(false)`. Reviewers open the function body just to learn which branch fires. A future third mode forces a second boolean, then a tuple of booleans."
          rule="No boolean flag parameters. Split the function into two (or more) positively named functions, one per branch. `render(true)` is wrong; `renderExpanded()` and `renderCollapsed()` are right. Per spec/17/31 line 48."
          doThis="When you touch a function that takes a `boolean` flag, extract each branch into its own named function in the same diff, then rewrite every call site to pick the named variant. Do not add a second flag to an already-flagged function; refactor first."
        />
      </div>
    </SlideLayout>
  );
}
