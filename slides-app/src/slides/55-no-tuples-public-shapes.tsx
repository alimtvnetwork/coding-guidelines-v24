import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 57: no tuples as public shapes. Every hook return, prop bundle,
 * reducer state, action, context value, and function argument bag gets a
 * named type or interface.
 *
 * Source: 02-spec/17/31 line 109.
 */

const BEFORE = `// Tuples everywhere. Positional. Anonymous. Fragile.
export function useUser(): [User | null, boolean, Error | null] { ... }
export function useToggle(initial: boolean): [boolean, () => void, () => void] { ... }

type Action =
  | ["increment", number]
  | ["set", { count: number; label: string }]
  | ["reset"];

export function counterReducer(
  state: [number, string, boolean],
  action: Action,
): [number, string, boolean] { ... }

// Call sites.
const [u, l, e] = useUser();
const [on, open, close] = useToggle(false);
dispatch(["set", { count: 5, label: "hi" }]);`;

const AFTER = `// Every public shape has a name. Every field has a name.
export type UserQueryResult = { user: User | null; isLoading: boolean; error: Error | null };
export function useUser(): UserQueryResult { ... }

export type ToggleControls = { isOn: boolean; open: () => void; close: () => void };
export function useToggle(initial: boolean): ToggleControls { ... }

export type CounterState = { count: number; label: string; isDirty: boolean };
export type CounterAction =
  | { kind: "Increment"; by: number }
  | { kind: "Set"; count: number; label: string }
  | { kind: "Reset" };

export function counterReducer(state: CounterState, action: CounterAction): CounterState { ... }

// Call sites.
const { user, isLoading, error } = useUser();
const { isOn, open, close } = useToggle(false);
dispatch({ kind: "Set", count: 5, label: "hi" });`;

export default function NoTuplesAsPublicShapesSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 57 · TypeScript · no tuples as public shapes"
      title="Tuples are laziness in a trench coat. Every hook return, prop bundle, reducer state, reducer action, context value, and argument bag gets a name."
      subtitle="02-spec/17/31 line 109: no tuples as public shapes. If a value has two or more fields or gets destructured at the call site, it needs a name. `useUser(): [User, boolean, Error]` is wrong; `useUser(): UserQueryResult` with `{ user, isLoading, error }` is right."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="tsx"
          before={BEFORE}
          after={AFTER}
          beforeLabel="Bad. Positional destructure means every rename at the call site is silent (`const [u, l, e]` vs `const [user, loading, err]`) and every reordering of the return tuple compiles but corrupts every caller. Action tuples like `['set', { count, label }]` force a wide discriminated string literal that IDEs can autocomplete only partially and refactors miss. Adding a fourth field to `useToggle` (a `toggle()` fn) forces every call site to grow its destructure and re-order the positions, so a two-line hook change becomes a 40-file PR. Grep for `useUser` returns nothing useful because everyone destructures with different local names."
          afterLabel="Good. Named types are grep targets, IDE quick-info shows `{ user, isLoading, error }` on hover, adding a field is a one-line diff in one file, and renaming a field is a rename refactor across the codebase. Action objects with a `kind` discriminant get exhaustive `switch` checks (`case 'Increment':`), so removing a case is a compile error, not a silent no-op. PascalCase action kinds match NAM-001. `isDirty` in `CounterState` is added without touching a single reducer call site."
        />
        <ActionPanel
          slideId="55-no-tuples-public-shapes"
          symptom="A refactor to add a `refetch` function to `useUser` broke 47 call sites and shipped in prod because 'the types were fine.' They were: every `[u, l, e] = useUser()` compiled, they just now bound `refetch` to `e` and the error handler to nothing. Post-mortem: (1) tuple returns have no field names, so IDE tooltips show `[User | null, boolean, Error | null]` with no clue which is which; (2) call sites use ad-hoc local names (`[u, l, e]`, `[user, loading, err]`, `[data, isFetching, error]`), so grep and refactor tools cannot correlate; (3) action tuples force stringly-typed dispatch like `dispatch(['set', ...])` that skips exhaustive-switch coverage. All three symptoms share one root cause: public shapes without names."
          rule="02-spec/17/31 line 109 is absolute: no tuples as public shapes. 'Public' means anything crossing a module boundary: exported function returns, exported function parameters, exported types, context values, reducer states, reducer actions, prop objects, event payloads, custom hook returns. If a value has two or more fields or gets destructured at the call site, it gets a named `type` or `interface`. React's own the built-in useState pair is the ONE tuple exception in the codebase, tolerated because (a) React owns it, (b) both positions are heavily conventionalized (`[value, setValue]`), and (c) the second slot is always a setter. Custom hooks do not get to piggy-back on that exception, they wrap it. Tuples remain fine strictly for local, non-exported returns of arity 2 where both positions are conventional and named at the destructure site."
          doThis="Enforce mechanically: (1) ESLint `@typescript-eslint/consistent-type-definitions: [error, type]` for consistency, plus a project rule `custom/no-tuple-return` that flags any exported function whose return type is a tuple literal `[A, B, ...]` with an exception for `useState`-style hooks; (2) TypeScript rule `custom/named-hook-return` (pairs with REACT-008) rejects anonymous object returns like an anonymous object return and requires a named alias; (3) reducer actions get a `kind: 'PascalCase'` discriminant, enforced by `custom/action-shape` that walks any type ending in `Action` and requires `kind: string` present; (4) codemod `scripts/codemods/tuple-to-object.ts` walks exported `use*`, `create*`, `make*` returns, infers field names from JSDoc or falls back to `a`/`b`/`c` with a TODO comment; (5) code review reflex: any `[a, b]` in an exported signature is rejected unless it wraps `useState`. When you cannot name a field, the field is doing two things, split it."
        />
      </div>
    </SlideLayout>
  );
}
