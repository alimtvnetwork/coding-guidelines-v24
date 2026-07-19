import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 58: name every generic parameter and every composite type.
 *
 * Source: spec/17/31 line 110.
 */

const BEFORE = `// Inline composites. Single-letter generics. Nothing to grep, nothing to hover.
export function group‹T, K›(items: T[], key: (x: T) =› K): Map‹K, T[]› {
  const out = new Map‹K, T[]›();
  for (const x of items) {
    const k = key(x);
    const arr = out.get(k) ?? [];
    arr.push(x);
    out.set(k, arr);
  }
  return out;
}

export type UsersById = Map‹string, Array‹{ id: number; name: string; roles: string[] }››;

export function useIndex‹T, U›(
  rows: T[],
  by: (r: T) =› U,
): { data: Map‹U, T[]›; get: (k: U) =› T[] } { ... }`;

const AFTER = `// Named domain types. Named generic parameters. Grep-friendly. Hover-friendly.
export type UserId = string;
export type User = { UserId: number; Name: string; Roles: readonly Role[] };
export type UsersById = ReadonlyMap‹UserId, readonly User[]›;

export function groupBy‹TItem, TKey›(
  items: readonly TItem[],
  toKey: (item: TItem) =› TKey,
): ReadonlyMap‹TKey, readonly TItem[]› {
  const out = new Map‹TKey, TItem[]›();
  for (const item of items) {
    const k = toKey(item);
    const bucket = out.get(k) ?? [];
    bucket.push(item);
    out.set(k, bucket);
  }
  return out;
}

export type IndexResult‹TKey, TItem› = {
  data: ReadonlyMap‹TKey, readonly TItem[]›;
  get: (key: TKey) =› readonly TItem[];
};

export function useIndex‹TItem, TKey›(
  rows: readonly TItem[],
  toKey: (row: TItem) =› TKey,
): IndexResult‹TKey, TItem› { ... }`;

export default function NamedGenericsAndCompositesSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 58 · TypeScript · name every generic parameter and every composite"
      title="`Map‹string, Array‹{id:number}››` inline is wrong. Extract `UsersById`. `T`/`U` are wrong. Use `TItem`/`TKey`."
      subtitle="spec/17/31 line 110: name every generic parameter and every composite type. Bare `T`, `U`, `K`, `V` are laziness in application code. Inline composites hide domain meaning and defeat rename refactors."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="tsx"
          before={BEFORE}
          after={AFTER}
          beforeLabel="Bad. `group‹T, K›` says nothing about intent, IDE hover shows `T` and `K` with no clue what they represent, and the second parameter `key` shadows the builtin `key` React prop when destructured. `Map‹string, Array‹{ id, name, roles }››` is the same shape spelled inline in three places, each drifts independently when one file adds `email`. `useIndex` returns an anonymous object literal (also violates REACT-008 named return) built on top of anonymous composites, so no rename ever propagates."
          afterLabel="Good. `TItem` and `TKey` describe intent, so signatures read like English (`groupBy(items, toKey)`). `UsersById` is a named alias with a domain meaning, defined once, imported everywhere. `IndexResult` is a named composite reused by every call site and any future index-builder. `readonly` markers on inputs and outputs pair with REACT-005 (never mutate) so accidental `.push` is a compile error at the boundary. Rename any of these once and every call site updates."
        />
        <ActionPanel
          slideId="56-named-generics-composites"
          symptom="Two-week refactor to add `Email` to the user record leaked bugs in five places, all traceable to inline composites. (1) `Map‹string, Array‹{ id: number; name: string; roles: string[] }››` was repeated in `useUserIndex`, `useTeamRoster`, `useAuditView`, and `useOrgChart`; adding `email` to one shape did not add it to the other three, so `roster.get(userId)[0].email` type-checked in only one file. (2) `group‹T, K›` had a call site that passed a compound key by accident (`(u) =› [u.OrgId, u.RoleId]`); nothing complained because `K` is inferred as `[number, number]` and tuple-key `Map` lookup silently misses. (3) IDE hover on any variable typed as the inline `Map` shape shows a 200-character blob that engineers stopped reading. Root cause across all three: composite types without names and generic parameters without names."
          rule="spec/17/31 line 110 is absolute. Two parts. (1) Every composite type gets a name. Composite means: any object type with 2+ fields, any `Array of composites` of a composite, any `Map`/`Set`/`Record`/`Promise` of a composite, any union of 2+ non-primitive members, any function type used more than once. `type UsersById = ReadonlyMap‹UserId, readonly User[]›` gets declared once and used everywhere. `Map keyed by string of inline user arrays` inline is banned. (2) Every generic parameter gets a meaningful name in application code. `TItem`, `TKey`, `TResponse`, `TError`, `TProps`, `TState`, `TAction`, `TDeps`, `TResult`. Never bare `T`, `U`, `K`, `V`. Utility library code (fp-ts, ts-toolbelt) is the only exception, and this codebase is not that. The `T` prefix disambiguates generics from concrete types in call-site hovers. Domain aliases pair with NAM-001 (PascalCase for types) and REACT-009 (no anonymous shapes)."
          doThis="Enforce mechanically: (1) ESLint `@typescript-eslint/naming-convention` with a `typeParameter` selector requiring `^T[A-Z]` prefix, so `T` alone or `U` fails; (2) custom rule `no-inline-composite` that flags any inline object type literal with 2+ properties, any inline `Array‹{...}›`, any inline `Map`/`Set`/`Record` whose value or key is a composite, with autofix suggesting an extracted alias name; (3) code review reflex: if a hover tooltip wraps in the IDE tooltip, the type needs a name; (4) place domain aliases in `src/types/‹domain›.ts` (`src/types/user.ts`, `src/types/order.ts`) and import them; component-local types go in a sibling `types.ts` (see REACT-011); (5) when you cannot name a composite, split it until you can (spec/17/31 line 112). Rule of thumb: if you would grep for this shape, it needs a name."
        />
      </div>
    </SlideLayout>
  );
}
