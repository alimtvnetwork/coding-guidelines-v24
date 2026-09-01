import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 49: React & TypeScript chapter opener.
 *
 * Source: 02-spec/17/31 lines 97-112. Effects are last resort, types are named,
 * lists are immutable, components stay small.
 */

const BEFORE = `// A "typical" React component: effect-driven, tuple-returning, inline-typed.
export function ProfileCard({ user, onSave }: {
  user: { id: string; name: string };
  onSave: (u: { id: string; name: string }) => void;
}) {
  const [name, setName] = useState(user.name);
  const [dirty, setDirty] = useState(false);

  // useEffect used to derive state from props. Runs on every keystroke.
  useEffect(() => {
    if (!user || (user && user.name && name !== user.name)) setDirty(true);
    else setDirty(false);
  }, [user, name]);

  // useEffect used to fetch on mount, no cleanup, mutates array in place.
  useEffect(() => { fetch("/api/audit").then(r => r.json()).then(list => list.push({ user })); }, []);

  return <div>{/* 180 lines below */}</div>;
}`;

const AFTER = `// Named types, zero effects, derived state, positive guard, small file.
// See slides 50-62 for each rule enforced individually.
export type ProfileCardProps = { user: User; onSave: (next: User) => void };

export function ProfileCard({ user, onSave }: ProfileCardProps) {
  const [name, setName] = useState(user.name);
  const isDirty = name !== user.name;           // derived, not an effect
  const canSave = isDirty && name.length > 0;   // positive boolean

  return (
    <ProfileCardShell>                          {/* extracted, <100 lines */}
      <NameField value={name} onChange={setName} />
      <SaveButton disabled={!canSave} onClick={() => onSave({ ...user, name })} />
    </ProfileCardShell>
  );
}`;

export default function ReactChapterOpenerSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 49 · React & TypeScript · chapter opener"
      title="React & TypeScript: effects are the last resort, types are always named, files stay small."
      subtitle="The next 13 slides enforce 02-spec/17/31 lines 97-112: guard effects behind positive booleans, ban tuple public shapes, keep components under 100 lines, and give every generic and prop bag a real name. This slide is the map: it lists what changes and why."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="tsx"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ Effects used to derive state and to fetch (no cleanup, mutates in place). Inline anonymous prop type. Negative guard with double-check `!user || (user && ...)`. Component grows past 100 lines because nothing is extracted. Every future bug fix touches the same god-component."
          afterLabel="✅ `isDirty` is a derived value, not an effect. `canSave` is a positive boolean. `ProfileCardProps` is a named type. The component is a shell that composes `NameField` and `SaveButton`, each in its own file under 100 lines. No effects means no cleanup bugs, no double-fetches, no stale closures."
        />
        <ActionPanel
          slideId="47-react-chapter-opener"
          symptom="A `ProfileCard` grows to 380 lines with four `useEffect`s: one to derive `isDirty` from props, one to fetch audit data on mount with no cleanup, one to sync local state back to the parent, and one nobody remembers writing. The prop bag is typed inline as an anonymous object literal so every consumer redeclares it slightly differently. A refactor that renames `User.name` to `User.fullName` misses three call sites because TypeScript can't cross-reference an anonymous shape."
          rule="React & TypeScript work under seven grouped rules from 02-spec/17/31 lines 97-112: (1) `useEffect` is the last resort, default is zero, add one only to sync with an external system; (2) every effect uses positive guards extracted above the effect, never inline negatives; (3) one effect, one concern, always with a cleanup; (4) never mutate state/props/hook returns, build a new value; (5) lists key on stable data, never index; (6) component files stay under 100 lines, custom hooks start with `use` and return a named object, never a tuple; (7) every prop bag, hook return, reducer state/action, context value, and generic parameter has a named `type` or `interface` in a dedicated `types.ts`. The next 13 slides drill into each rule with a Symptom/Rule/Action."
          doThis="Land the enforcement stack alongside this chapter: (1) ESLint rules `react-hooks/exhaustive-deps` (error), `react/jsx-key` (error), `@typescript-eslint/no-explicit-any` (error), plus custom rules `no-inline-prop-types`, `no-tuple-hook-return`, `no-negative-effect-guard`, `max-lines-per-component: 100`; (2) a CI job `react-shape-audit` that logs `react.shape.violation` with `{ file, rule, line }` for every hit and blocks the merge on any hard rule; (3) a codemod script `scripts/extract-inline-prop-types.mjs` that lifts inline object types into a sibling `types.ts` on demand. PR template line: 'Any new `useEffect`? Justify the external system. Any new prop type? Confirm it lives in `types.ts`.' Reviewers use slides 50-62 as the checklist."
        />
      </div>
    </SlideLayout>
  );
}
