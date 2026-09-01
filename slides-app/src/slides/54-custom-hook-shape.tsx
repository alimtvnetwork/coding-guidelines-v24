import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 56: custom hooks start with `use`, return named object type,
 * never call other hooks conditionally.
 *
 * Source: 02-spec/17/31 line 108.
 */

const BEFORE = `// A "helper" that quietly calls hooks, returns a tuple, and skips the use- prefix.
export function fetchOrder(id: OrderId) {
  const [order, setOrder] = useState<Order | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  if (id) {
    useEffect(() => {
      api.getOrder(id).then(setOrder).finally(() => setLoading(false));
    }, [id]);
  }

  return [order, loading, error] as const;
}

// Call site.
const [o, l, e] = fetchOrder(orderId);   // which one was loading again?`;

const AFTER = `// Prefixed with use-, returns a named object type, hooks called unconditionally.
export type OrderQueryResult = {
  order: Order | null;
  isLoading: boolean;
  error: Error | null;
};

export function useOrder(id: OrderId): OrderQueryResult {
  const [order, setOrder] = useState<Order | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<Error | null>(null);

  useEffect(() => {
    if (!id) return;
    let cancelled = false;
    api
      .getOrder(id)
      .then((next) => !cancelled && setOrder(next))
      .catch((err) => !cancelled && setError(err))
      .finally(() => !cancelled && setIsLoading(false));

    return () => { cancelled = true; };
  }, [id]);

  return { order, isLoading, error };
}

// Call site.
const { order, isLoading, error } = useOrder(orderId);`;

export default function CustomHookShapeSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 56 · React · custom hooks: use- prefix, named return, unconditional"
      title="A function that calls hooks IS a hook. Name it `useX`, return an object with a name, and never gate a hook call on a condition."
      subtitle="02-spec/17/31 line 108: custom hooks start with `use`, return a named object type (never a bare tuple), and never call other hooks conditionally. The `use` prefix is how React and its lint rules recognize hook call sites; the named return is how call sites stay readable when the shape grows."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="tsx"
          before={BEFORE}
          after={AFTER}
          beforeLabel="Bad. `fetchOrder` calls `useState` and `useEffect` but the name has no `use` prefix, so `eslint-plugin-react-hooks` cannot enforce the Rules of Hooks on it and any static analyzer sees a plain function. The `if (id) { useEffect(...) }` is a hook call inside a conditional, which violates the fundamental invariant that hook order stays identical across renders (the moment `id` toggles, React attaches state to the wrong slot and every subsequent hook shifts). The `[order, loading, error]` tuple destructures positionally, so any reorder or insertion at the hook silently breaks every call site with a wrong-type-that-still-compiles bug. Missing cleanup in the effect leaks state writes on unmount."
          afterLabel="Good. Renamed to `useOrder` so the linter enforces Rules of Hooks. Return type `OrderQueryResult` is declared once and reused, so IDE autocomplete works and adding a `refetch` field is one line in one place. Hooks are called unconditionally; the `if (!id) return;` guard is inside the effect body, not around the effect. Cleanup flag prevents state writes after unmount. Call sites destructure by name (`{ order, isLoading, error }`), so reordering the return type is safe and adding fields never breaks anyone."
        />
        <ActionPanel
          slideId="54-custom-hook-shape"
          symptom="Three related bugs traced to one root cause. (1) `TypeError: Cannot read properties of null (reading 'Title')` fires intermittently in prod: `fetchOrder` calls `useEffect` inside `if (id)`, so when `id` first arrives after being falsy React attaches the effect to the wrong hook slot and the `order` state actually lives where `loading` used to live. (2) A junior added a `refetch` field to the return tuple: half the codebase now reads `const [order, loading, refetch, error]` and the other half still reads `[order, loading, error]`, all type-check because tuples are just arrays. (3) `eslint-plugin-react-hooks` never warned because the function is called `fetchOrder`, not `useOrder`. Every symptom is 02-spec/17/31 line 108."
          rule="02-spec/17/31 line 108 is absolute for any function that calls another hook. Three parts, all required: (1) name starts with `use` (`useOrder`, `useDebouncedValue`, `useFocusTrap`), which is how `react-hooks/rules-of-hooks` and `react-hooks/exhaustive-deps` find the call site and how humans know a function is stateful; (2) return a named `type` or `interface` with named fields (`OrderQueryResult`, `DebouncedValue`), never a bare tuple (see REACT-009 for the general rule against tuples as public shapes); (3) hook calls at the top level only, in the same order every render. Never gate a `useState`/`useEffect`/`useMemo`/`useCallback`/`useContext`/`useRef` call on `if`, `&&`, `?:`, `try`, `switch`, `for`, or `while`. The guard belongs inside the hook body (`useEffect that guards with if !id return inside the body`), never around the hook call. Custom hooks compose custom hooks under the same rules."
          doThis="Enforce mechanically: (1) `eslint-plugin-react-hooks` with `rules-of-hooks: error` and `exhaustive-deps: error`; the plugin only fires on names starting with `use`, so add a project rule `custom/hook-name-prefix` that flags any function whose body contains `use[A-Z]` calls but whose own name does not start with `use`; (2) TypeScript project rule: hook return types must be a named type reference, not an anonymous object literal or tuple literal, enforced by `custom/hook-return-named-type` that walks return-type annotations of exported `use*` functions; (3) codegen: `scripts/codemods/hook-tuple-to-object.ts` rewrites tuple returns to `{ a, b, c }` and updates call sites via ts-morph; (4) code review reflex: any PR adding a `use*` function is asked 'what is the return type named?' before approve, and any `[a, b, c]` return in a hook is rejected; (5) pair with REACT-005 (no mutation) and REACT-006 (stable keys) as the three-legged React hygiene stool. When you cannot name the return type, split the hook until you can (02-spec/17/31 line 112)."
        />
      </div>
    </SlideLayout>
  );
}
