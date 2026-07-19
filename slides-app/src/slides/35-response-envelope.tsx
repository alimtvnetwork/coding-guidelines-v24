import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 37: Universal response envelope.
 *
 * Source: spec/17/31 line 80. Backend APIs return { data, errors[], meta }.
 * Frontend parses via ONE shared helper, never per-caller.
 */

const BEFORE = `// backend: /api/orders returns raw array
res.json(await db.orders.list());

// backend: /api/checkout returns { error: "...", code: 42 }
res.status(400).json({ error: "invalid card", code: 42 });

// frontend: every caller invents its own shape
const orders = await fetch("/api/orders").then(r => r.json());
// orders.map(...)  ❌ crashes if backend returned { error }

try {
  const r = await fetch("/api/checkout", ...);
  const body = await r.json();
  if (body.error) toast(body.error);          // ❌ ad-hoc key
} catch (e) { /* swallowed */ }               // ❌ silent`;

const AFTER = `// backend: every route returns the same envelope
type Envelope<T> = {
  data: T | null;
  errors: Array<{ code: ErrorCode; message: string; field?: string }>;
  meta: { requestId: string; page?: PageMeta };
};

res.json({ data: orders, errors: [], meta: { requestId } });
res.status(400).json({
  data: null,
  errors: [{ code: ErrorCodes.CheckoutInvalidCard, message: "Invalid card" }],
  meta: { requestId },
});

// frontend: one shared parser, every caller uses it
export async function apiCall<T>(input: RequestInfo, init?: RequestInit): Promise<T> {
  const res  = await fetch(input, init);
  const body: Envelope<T> = await res.json();

  if (body.errors.length > 0) {
    logError("apiCall", { url: input, requestId: body.meta.requestId, errors: body.errors });
    throw new AppError(body.errors[0].code, { requestId: body.meta.requestId });
  }

  return body.data as T;
}

const orders = await apiCall<Order[]>("/api/orders");   // ✅ typed, uniform`;

export default function ResponseEnvelopeSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 37 · Universal response envelope"
      title="One envelope. One parser. Every route, every caller."
      subtitle="Backend routes always return `{ data, errors[], meta }`. The frontend has exactly one `apiCall` helper that reads the envelope, logs the errors with the request id, and throws a typed `AppError`. Callers write a one-liner and get typed data back, or an error they can switch on."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="typescript"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ Two routes, two shapes. Every caller reinvents error handling. Failures get swallowed."
          afterLabel="✅ One `Envelope<T>`, one `apiCall<T>` parser. Callers get typed data or a typed AppError with requestId."
        />
        <ActionPanel
          slideId="35-response-envelope"
          symptom="One route returns a raw array, another returns `{ error, code }`, a third returns `{ ok: true, payload }`. Every frontend caller writes its own try/catch, half of them swallow errors, and support cannot correlate a failing UI to a backend log line because there is no shared request id in the response. New endpoints keep inventing new shapes because there is no reference."
          rule="Every backend response uses the envelope `{ data, errors[], meta }`. `errors` is always an array (empty on success), each item carries a registered `ErrorCode`, a human message, and optional field. `meta` always carries `requestId`. The frontend has exactly one `apiCall` helper that parses the envelope, logs `errors` with `requestId` and `url` context, and throws an `AppError` with the first code so UI code can switch on it. Callers do not read `.errors` directly. Per spec/17/31 line 80 and the error-management digest."
          doThis="Create (or extend) `src/lib/api.ts` with the `apiCall` helper and the `Envelope` type. On the backend, add a middleware or response helper that wraps every handler so success and error paths always produce the same envelope with a `requestId`. Migrate callers one feature at a time: delete per-caller try/catch, replace raw `fetch().then(r => r.json())` with `await apiCall(...)`. Reject any new route or caller that bypasses the envelope or the helper."
        />
      </div>
    </SlideLayout>
  );
}
