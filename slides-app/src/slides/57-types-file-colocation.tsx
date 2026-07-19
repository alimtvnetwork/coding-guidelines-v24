import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 59: prop types and event handler types live in a dedicated
 * `types.ts` next to the component (or `src/types/` when shared).
 *
 * Source: spec/17/31 line 111.
 */

const BEFORE = `// src/components/ProfileCard.tsx  --  inline everything, anonymous forever.
export function ProfileCard({
  user,
  onSave,
  onCancel,
}: {
  user: { id: number; name: string; email: string; avatarUrl: string };
  onSave: (u: { id: number; name: string; email: string }) => void;
  onCancel: () => void;
}) {
  const [draft, setDraft] = useState(user);
  return ( ... );
}`;

const AFTER = `// src/components/profile-card/types.ts
import type { User } from "@/types/user";

export type ProfileCardProps = {
  user: User;
  onSave: (next: User) => void;
  onCancel: () => void;
};

export type ProfileDraftChange = {
  field: keyof User;
  value: string;
};

// src/components/profile-card/ProfileCard.tsx
import type { ProfileCardProps } from "./types";
import type { User } from "@/types/user";

export function ProfileCard({ user, onSave, onCancel }: ProfileCardProps) {
  const [draft, setDraft] = useState<User>(user);
  return ( ... );
}

// src/components/profile-card/index.ts
export { ProfileCard } from "./ProfileCard";
export type { ProfileCardProps } from "./types";`;

export default function TypesFileColocationSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 59 · TypeScript · prop and handler types live in a dedicated types.ts"
      title="Never inline an anonymous object type on a component signature. Prop types get a name and a file."
      subtitle="spec/17/31 line 111: prop types and event handler types live in a dedicated `types.ts` next to the component, or in `src/types/` when shared. Extraction keeps components under the 100-line cap and makes prop shapes greppable, importable, and stable across refactors."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="tsx"
          before={BEFORE}
          after={AFTER}
          beforeLabel="Bad. Anonymous prop object literal on the signature means: no name to import for tests, Storybook stories, or wrapper components; IDE hover shows the raw literal instead of `ProfileCardProps`, so the shape reads as noise; the user shape is spelled inline AND drifts from the app-wide `User` type; the save handler accepts a different, narrower object than user, guaranteeing an impedance mismatch when a caller passes the full user back; adding one prop grows the component signature by three to five lines and pushes the file toward the REACT-007 100-line cap for no reason."
          afterLabel="Good. `ProfileCardProps` lives in a sibling `types.ts` and is imported by the component, the tests, the Storybook story, and any wrapper. The user and save-handler payload both use the shared `User` type from `src/types/user.ts`, so there is one source of truth. `index.ts` re-exports the component and the prop type as one public surface. Adding a prop is a one-line change in `types.ts` and the component signature never grows past `ProfileCardProps`."
        />
        <ActionPanel
          slideId="57-types-file-colocation"
          symptom="ProfileCard had 11 props inlined on its signature (34-line function signature, 47-line file). Three ripple failures: the Storybook story redeclared the prop shape with slightly different optionality (`email?` vs `email`), so the story rendered a state that could not exist in the app; a wrapper `EditableProfileCard` retyped the same 11 props from scratch and got two wrong; a rename from `avatarUrl` to `AvatarUrl` for NAM-001 alignment required editing six files instead of one, and one file was missed and shipped a runtime `undefined` in prod. All three trace to one root cause: the prop shape had no name and no home."
          rule="spec/17/31 line 111 is absolute. Prop types and event handler types live in a dedicated `types.ts` next to the component. Shared cross-component types (User, Order, AuditEvent) live in `src/types/<domain>.ts` and are imported. Never inline an anonymous object type on a component signature. The pattern is one folder per non-trivial component: `src/components/profile-card/{ProfileCard.tsx, types.ts, index.ts}` (plus test and stories files when relevant). Trivial single-use components (one prop, one line, zero handlers) may inline, but the moment a second prop or a handler lands, extract. Pairs with REACT-009 (no tuples), REACT-010 (name every composite), and REACT-007 (files under 100 lines) as one system: the types.ts file is where extracted names go so the component file stays small."
          doThis="Enforce mechanically: (1) custom ESLint rule `no-inline-props-type` that flags any React function component whose props parameter has an inline object type annotation with 2+ properties or any function-type property, autofix suggests moving to `./types.ts` and importing; (2) folder convention scaffold `pnpm gen:component ProfileCard` (script under `scripts/gen-component.mjs`) that produces `ProfileCard.tsx`, `types.ts`, `index.ts`, and a Vitest stub in the correct layout; (3) `src/types/` for cross-cutting domain types (`user.ts`, `order.ts`, `audit.ts`), one file per aggregate; (4) code review reflex: any anonymous object type on a component signature is rejected; wrapper components must import the wrapped component's `Props` type, never redeclare; (5) tsconfig `paths` alias `@/types/*` so shared types have a stable import path. When a component grows a fourth prop, split the component (spec/17/31 line 112) before splitting the type."
        />
      </div>
    </SlideLayout>
  );
}
