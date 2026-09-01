import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 40: Verify both directions before shipping an integration.
 *
 * Source: 02-spec/17/31 line 78. Curl the backend AND inspect the
 * frontend detection logic. One side is not enough.
 */

const BEFORE = `# Author checks ONE side and ships
$ curl -s localhost:3000/api/orders | jq '.[0].id'
"ord_123"
# ✅ "backend works" -> merged

# Meanwhile in the UI:
const orders = await fetch("/api/orders").then(r => r.json());
orders.map(o => o.orderId);   // ❌ backend returns { id }, UI reads .orderId
// Silent empty list in production, no error, no log line.`;

const AFTER = `# 1) Curl the backend, assert the exact contract
$ curl -s -H "x-request-id: verify-1" localhost:3000/api/orders \\
    | jq '{shape: (.data|type), first: .data[0], meta: .meta}'
{
  "shape": "array",
  "first": { "orderId": "ord_123", "amountMinor": 2499 },
  "meta":  { "requestId": "verify-1" }
}

# 2) Inspect the frontend detection logic on the same payload
const orders = await apiCall<Order[]>("/api/orders");
console.assert(orders.length > 0, "orders empty");
console.assert("orderId" in orders[0], "missing orderId");   // ✅ contract check

# 3) Screenshot the UI showing the row, attach curl + assert output to the PR.`;

export default function VerifyBothDirectionsSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 40 · Verify both directions"
      title="Curl the backend. Inspect the frontend. One side is not proof."
      subtitle="A 200 in Postman does not prove the UI reads the field. A stopped spinner does not prove the backend returned real data. Before you say an integration works, run the request AND read the parsing code AND see the rendered pixels, on the same payload."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="bash"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ Backend returns `{ id }`, UI reads `.orderId`. Curl passes, spinner stops, production ships an empty list with no error."
          afterLabel="✅ Curl asserts the envelope shape, frontend asserts the field name on the same payload, UI screenshot attached. Contract mismatch caught before merge."
        />
        <ActionPanel
          slideId="38-verify-both-directions"
          symptom="A ticket says checkout is broken. The backend author says `curl /api/orders` returns 200. The frontend author says the spinner resolves. Both are technically correct, but the backend renamed `id` to `orderId` and the frontend still reads `.id`, so users see an empty list. Nobody ran BOTH sides on the same payload before merging, so the contract mismatch shipped with no error and no log line."
          rule="Before marking any integration ticket done, produce three artifacts on the SAME payload and attach them to the PR: (1) a `curl` command with the actual request id and the raw response body, (2) the frontend parsing code path with a runtime assertion on the specific field names the UI reads, and (3) a screenshot of the rendered UI showing the value. Backend-only verification is not verification. Frontend-only verification is not verification. Per 02-spec/17/31 line 78."
          doThis="Add a PR template checkbox: `Integration verified both directions (curl output + frontend field assertion + UI screenshot attached)`. For each integration PR, paste the curl command with `x-request-id`, paste the raw response, link the exact frontend line that reads each field, and attach the UI screenshot. Reviewers reject PRs that have only one side. For long-lived integrations, keep a tiny `scripts/verify/<name>.sh` that runs the curl and prints the assertion, so the next author can re-verify in one command."
        />
      </div>
    </SlideLayout>
  );
}
