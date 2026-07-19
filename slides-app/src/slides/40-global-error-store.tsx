import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 42: Global error store + single top-level modal.
 *
 * Source: spec/17/31 line 80. One store, one modal, fed by apiCall.
 */

const BEFORE = `// Every feature invents its own local error surface.
function CheckoutButton() {
  const [err, setErr] = useState<string | null>(null);
  const onPay = async () => {
    try { await pay(); }
    catch (e) { setErr(String(e)); toast.error(String(e)); }
  };
  return <>{err && <Alert>{err}</Alert>} ...</>;
}

function ProfileForm() {
  const [error, setError] = useState("");   // different name, same job
  // ...
}
// -> Same failure shows as toast + inline alert + modal. Or nothing at all.
// -> Stale errors linger after route change. No single audit surface.`;

const AFTER = `// One store, one modal, fed by apiCall (ERR-005).
// src/lib/errorStore.ts
export const errorStore = create<{ current: AppError | null; push: (e: AppError) => void; clear: () => void }>((set) => ({
  current: null,
  push: (e) => { log.error({ op: "ui.error", code: e.code, requestId: e.requestId }); set({ current: e }); },
  clear: () => set({ current: null }),
}));

// src/lib/apiCall.ts (already exists per ERR-005) pushes on failure:
if (!res.ok) errorStore.getState().push(parseError(body));

// src/app/ErrorModal.tsx  (mounted ONCE at the app root)
export function ErrorModal() {
  const err = errorStore((s) => s.current);
  if (!err) return null;
  return <Modal onClose={() => errorStore.getState().clear()}>{copyFor(err.code)}</Modal>;
}

// Features never render their own error UI. They just await apiCall().`;

export default function GlobalErrorStoreSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 42 · Global error store + single modal"
      title="One store. One modal. Fed by apiCall. Features never render their own error UI."
      subtitle="Local error state is how the same failure ends up shown three times, or not at all. Route away and the stale toast follows you. Centralise it: one store gets the error from apiCall, one modal at the app root renders it, one clear action dismisses it. Every failure is auditable in one place."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="tsx"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ Each feature owns its own `useState` + toast + alert. Same failure surfaces three times or zero times, stale errors survive navigation, and there is no single log or audit of what the user saw."
          afterLabel="✅ `errorStore` receives every failure from `apiCall` (ERR-005) with `code` + `requestId`, a single `<ErrorModal>` at the app root renders it, `clear()` dismisses it, and `log.error` fires with context on every push. One surface. One audit trail."
        />
        <ActionPanel
          slideId="40-global-error-store"
          symptom="A checkout POST fails. The user sees a red toast, then an inline alert under the button, then a modal from the payment provider iframe. They dismiss the toast and route away, but the alert copy is still on screen for two seconds before the next page mounts. Meanwhile the on-call has no idea which of the three surfaces the user actually clicked, because none of them logged the error id."
          rule="Every user-facing error goes through ONE global error store fed by the shared `apiCall` parser (ERR-005). The store holds at most one active error, tagged with `code`, `requestId`, and `at`. A SINGLE ErrorModal component mounted at the app root reads the store and renders it, with copy chosen by `code` from a shared registry (ERR-004). Feature components never call `toast.error`, never keep their own `useState` for errors, and never render their own alerts. On route change the store clears automatically. Every `push` logs at `error` level with the full context per LOG-002. Per spec/17/31 line 80."
          doThis="Create `src/lib/errorStore.ts` with `current`, `push(e)`, and `clear()`. Wire `apiCall` (ERR-005) to call `errorStore.getState().push(parseError(body))` on failure, with `code` from the ERR-004 registry and `requestId` from the response envelope. Mount the ErrorModal component exactly once in `src/app/AppRoot.tsx`, above the router. Add a router subscription that calls `clear()` on route change. Add an ESLint rule (or CI grep) banning `toast.error(` and local error state outside `src/lib/errorStore.ts` and `src/app/ErrorModal.tsx`. Migrate existing feature components in one PR, delete their local alerts, and confirm the single modal receives the failure in DevTools."
        />
      </div>
    </SlideLayout>
  );
}
