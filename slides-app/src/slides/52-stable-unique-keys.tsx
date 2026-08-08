import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 54: stable, unique keys derived from data, never the array index.
 *
 * Source: spec/17/31 line 106.
 */

const BEFORE = `// Every list uses the array index or a fresh id per render.
export function TaskList({ tasks }: TaskListProps) {
  const sorted = [...tasks].sort(byDueDate);

  return (
    <ul>
      {sorted.map((task, i) => (
        <li key={i}>
          <TaskRow task={task} />
          <NoteInput defaultValue={task.Note} />
        </li>
      ))}
    </ul>
  );
}

export function CommentList({ comments }: CommentListProps) {
  return (
    <ul>
      {comments.map((c) => (
        <li key={crypto.randomUUID()}>{c.Body}</li>
      ))}
    </ul>
  );
}`;

const AFTER = `// Keys come from the row's own identity and stay stable across renders.
export function TaskList({ tasks }: TaskListProps) {
  const sorted = useMemo(() => [...tasks].sort(byDueDate), [tasks]);

  return (
    <ul>
      {sorted.map((task) => (
        <li key={task.TaskId}>
          <TaskRow task={task} />
          <NoteInput defaultValue={task.Note} />
        </li>
      ))}
    </ul>
  );
}

export function CommentList({ comments }: CommentListProps) {
  return (
    <ul>
      {comments.map((c) => (
        <li key={c.CommentId}>{c.Body}</li>
      ))}
    </ul>
  );
}`;

export default function StableUniqueKeysSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 54 · React · stable, unique keys from data"
      title="Keys are React's identity, not decoration. Index keys reset state on every reorder; fresh UUIDs remount every row."
      subtitle="spec/17/31 line 106: lists must have stable, unique `key` props derived from data, never the array index unless the list is truly static. Keys tell React which DOM node maps to which item across renders."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="tsx"
          before={BEFORE}
          after={AFTER}
          beforeLabel="Bad. `key={i}` in `TaskList`: after `.sort(byDueDate)` the task at index 0 changed, but React thinks 'row 0' is still 'row 0' and keeps the old `NoteInput` mounted, so half-typed notes jump onto the wrong task. `key={crypto.randomUUID()}` in `CommentList`: every render mints new keys, so React unmounts and remounts every `li`, wiping focus, cursor position, CSS transitions, and any child state. Both bugs are invisible in dev with 3 items and catastrophic in prod with 300."
          afterLabel="Good. `key={task.TaskId}` binds the row to its persistent database id, so sorting or filtering rewires the DOM the way React was designed to. `key={c.CommentId}` is stable across renders, so `li` state (selection, animations, focus) survives. Sorting happens inside `useMemo` so the derived array itself is stable when input hasn't changed. If a row has no natural id (rare, pre-save drafts), mint one once with `useId()` or a client-side `DraftId` field, not per-render `crypto.randomUUID()`."
        />
        <ActionPanel
          slideId="52-stable-unique-keys"
          symptom="Support ticket: 'I typed a note on Task A, clicked the Due Date column to sort, and my note is now attached to Task B.' Repro: type in a `NoteInput` on any row, sort the list, watch the note follow the row index instead of the task. Separately, the comments panel loses input focus and restarts its fade-in animation on every keystroke because a parent re-renders and every `li` gets a fresh `crypto.randomUUID()` key. Both traced to the same root cause class: keys not tied to row identity."
          rule="spec/17/31 line 106 is absolute: lists must have stable, unique `key` props derived from data, never the array index unless the list is truly static (a hard-coded nav menu, a 4-item settings tab list that never reorders or filters). 'Truly static' means: the array is a literal, the items never insert, delete, sort, or filter, and no child holds state. Everything else uses the row's own identifier: `TaskId`, `CommentId`, `UserId`, the `{TableName}Id` PK from SCHEMA-002. Never `Math.random()`, never `crypto.randomUUID()` inline, never `Date.now()`, never a composite of unstable fields. If two rows can share the same id (grouped views), namespace it: `key={`${groupId}:${task.TaskId}`}`."
          doThis="Enforce mechanically: (1) ESLint `react/no-array-index-key` set to `error`, not `warn`, no per-file disables without a `// list is static because ...` justification comment; (2) custom rule `react/no-unstable-key` that flags any `key={...}` whose expression contains `Math.random`, `crypto.randomUUID`, `Date.now`, `new Date`, `performance.now`, or a call to a function whose name starts with `generate`/`create`/`make`Id; (3) codemod `scripts/codemods/index-key-to-id.ts` that rewrites `map((x, i) => ... key={i})` to `map((x) => ... key={x.Id})` and lists remaining call sites for human review; (4) code review checklist: every `key=` on a `.map` must resolve to a database PK or a `useId()` value assigned once, and reviewers reject anything else. When a row genuinely has no id yet (unsaved draft), add a `DraftId: string` field to the model, mint it once at creation with `crypto.randomUUID()` stored in state, and key on that. When in doubt, `useId()` at the parent and pass it in."
        />
      </div>
    </SlideLayout>
  );
}
