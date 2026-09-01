import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 53: never mutate state, props, or arrays/objects returned by hooks.
 *
 * Source: 02-spec/17/31 line 105.
 */

const BEFORE = `// Mutating state, props, and a React Query cache entry in place.
export function TaskBoard({ board }: TaskBoardProps) {
  const [tasks, setTasks] = useState<Task[]>(initialTasks);
  const { data: users } = useUsers();

  function handleAdd(next: Task) {
    tasks.push(next);          // mutates state array
    setTasks(tasks);           // same reference, React bails out
  }

  function handleRename(id: TaskId, title: string) {
    const t = tasks.find((x) => x.Id === id);
    if (t) t.Title = title;    // mutates state object
    setTasks([...tasks]);      // hack: new outer ref hides the real bug
  }

  function handleSortUsers() {
    users?.sort((a, b) => a.Name.localeCompare(b.Name)); // mutates cached array
  }

  board.Meta.LastViewedAt = Date.now(); // mutates prop

  return <Grid tasks={tasks} users={users} board={board} onAdd={handleAdd} />;
}`;

const AFTER = `// Build new values. Never touch what you did not create.
export function TaskBoard({ board }: TaskBoardProps) {
  const [tasks, setTasks] = useState<readonly Task[]>(initialTasks);
  const { data: users } = useUsers();

  function handleAdd(next: Task) {
    setTasks((prev) => [...prev, next]);
  }

  function handleRename(id: TaskId, title: string) {
    setTasks((prev) => prev.map((t) => (t.Id === id ? { ...t, Title: title } : t)));
  }

  const sortedUsers = useMemo(
    () => (users ? [...users].sort((a, b) => a.Name.localeCompare(b.Name)) : []),
    [users],
  );

  const boardWithView = useMemo<Board>(
    () => ({ ...board, Meta: { ...board.Meta, LastViewedAt: Date.now() } }),
    [board],
  );

  return <Grid tasks={tasks} users={sortedUsers} board={boardWithView} onAdd={handleAdd} />;
}`;

export default function NoMutateStatePropsHookReturnsSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 53 · React · never mutate state, props, or hook returns"
      title="Build a new value with spread or `structuredClone`. Reference equality is React's contract, not yours."
      subtitle="02-spec/17/31 line 105: never mutate state, props, or arrays/objects returned by hooks. React, React Query, Zustand, and every memo boundary rely on reference equality to decide what re-renders and what is cached. Mutation silently corrupts the cache and skips renders."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="tsx"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ `tasks.push(next)` mutates state; `setTasks(tasks)` passes the same reference so React skips the render (bailout on `Object.is`). Renaming mutates the object inside the array and then hides the mutation with `setTasks([...tasks])`, which fools the outer bailout but leaves the inner reference identical, so any child memoized on `task` still sees the old title. `users?.sort(...)` mutates the array owned by React Query's cache: every component reading `useUsers()` now sees the sorted order, forever, until the query is invalidated. `board.Meta.LastViewedAt = Date.now()` mutates a prop, which is undefined behavior and breaks time-travel debugging."
          afterLabel="✅ Every write returns a new value. `setTasks((prev) => [...prev, next])` and `prev.map(...)` produce fresh references so React re-renders and memo boundaries see prop changes. `[...users].sort(...)` copies before sorting, leaving the React Query cache untouched. `{ ...board, Meta: { ...board.Meta, LastViewedAt: ... } }` builds a new prop shape instead of mutating the incoming one. `readonly Task[]` makes accidental `.push` a type error at the boundary."
        />
        <ActionPanel
          slideId="51-no-mutate-state-props-hook-returns"
          symptom="`TaskBoard` has three interlocking bugs. (1) Adding a task does nothing on screen because `tasks.push(next); setTasks(tasks)` passes the same reference and React bails out; a junior 'fixed' it with `setTasks([...tasks])`, which now works for add but hides the deeper issue. (2) Renaming a task updates the list header but the row itself keeps the old title until the page reloads, because the task object was mutated in place and the row is `React.memo`-wrapped on the task reference. (3) Users on the sidebar are alphabetized after visiting the board even when other pages want insertion order, because `users?.sort(...)` mutated the React Query cache entry. Bonus: `board.Meta.LastViewedAt = Date.now()` on a prop occasionally throws in strict-mode dev builds and never in prod, so it hides in staging."
          rule="02-spec/17/31 line 105 is absolute: never mutate state, props, or arrays/objects returned by hooks. Build a new value with spread, `map`, `filter`, `structuredClone`, or an immutable helper (Immer produce, `use-immer`). This covers `useState` values, `useReducer` state, function props, component props, `useQuery` data, `useSWR` data, `useContext` values, and anything returned by a custom hook. The rule pairs with line 104 (no raw `for`/`forEach` in render, which encourages mutation) and line 106 (stable keys, which relies on stable references). Reference equality is React's contract for bailouts, `React.memo`, `useMemo`, `useCallback`, and dep arrays; mutation invalidates that contract silently and the resulting bugs cost days to diagnose."
          doThis="Enforce it mechanically: (1) mark every state and hook return `readonly` at the type level (`useState<readonly Task[]>`, `type UsersQueryResult = { data: readonly User[] | undefined }`) so mutation is a compile error; (2) enable ESLint `functional/immutable-data` or `no-param-reassign` with `{ props: true }` scoped to `.tsx` files; (3) custom ESLint rule `no-mutate-hook-return` that flags `.push`, `.pop`, `.shift`, `.splice`, `.sort`, `.reverse`, and direct assignment on identifiers that came out of a hook (`useState`, `useQuery`, `useContext`, or any `use*`); (4) codemod pass `scripts/codemods/mutation-to-spread.ts` for the initial sweep, followed by human review; (5) React Query `defaultOptions.queries.structuralSharing: true` stays on, and a runtime dev-only `Object.freeze` on query results in development to make mutation throw loudly. Reviewers reject any `.push`/`.sort`/direct-assignment on state, props, or hook returns without a spread-copy first."
        />
      </div>
    </SlideLayout>
  );
}
