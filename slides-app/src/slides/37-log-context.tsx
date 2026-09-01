import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 39: Log context requirements.
 *
 * Source: 02-spec/17/31 line 77. Every log carries operation name,
 * request/session id, and key inputs. Never secrets, never PII
 * beyond a user id.
 */

const BEFORE = `// Bare strings, unsearchable in prod
log.error("failed");                              // ❌ no op, no id, no input
log.info("done");                                 // ❌ same
log.warn("retry");                                // ❌ retry of what?

// Full dump, leaks secrets and PII
log.info("login attempt", { ...req.body });       // ❌ leaks password
log.error("charge failed", {
  card: req.body.cardNumber,                      // ❌ PAN in logs
  cvv: req.body.cvv,                              // ❌ CVV in logs
  email: user.email,                              // ❌ PII
  authHeader: req.headers.authorization,          // ❌ bearer token
});`;

const AFTER = `// Named op + ids + key inputs, no secrets, no PII beyond userId
log.info("auth.login.start", {
  op: "auth.login",
  requestId, sessionId,
  userId,                                         // ✅ id only, not email
  method: "password",
});

log.warn("checkout.charge.retry", {
  op: "checkout.charge",
  requestId, orderId, userId,
  attempt: 2, upstream: "stripe",
  code: ErrorCodes.CheckoutUpstreamTimeout,       // ✅ registered code
});

log.error("checkout.charge.failed", {
  op: "checkout.charge",
  requestId, orderId, userId,
  amountMinor: 2499, currency: "USD",             // ✅ safe inputs
  code: ErrorCodes.CheckoutInvalidCard,
  cardLast4: card.last4,                          // ✅ last4 only, never PAN
  // no cardNumber, no cvv, no authHeader, no email
});`;

export default function LogContextSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 39 · Log context requirements"
      title="Every log line: op, requestId, key inputs. Never secrets. Never PII beyond userId."
      subtitle="A bare `log.error('failed')` is invisible in production. A `log.info({ ...req.body })` leaks passwords, card numbers, and bearer tokens into log storage. The middle ground is a named operation, correlation ids, and a hand-picked set of safe inputs."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="typescript"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ Bare strings that support cannot search, or full request dumps that leak PAN, CVV, passwords, bearer tokens, and email into log storage."
          afterLabel="✅ Named `op`, `requestId`, `userId`, hand-picked safe inputs (amount, currency, last4, registered code). No secrets. No PII beyond `userId`."
        />
        <ActionPanel
          slideId="37-log-context"
          symptom="Support gets a ticket referencing a failed checkout at 14:07 UTC. Grep for the user id returns nothing because every log line just says `error: failed`. Meanwhile the compliance scanner flagged 12k log entries containing full card numbers and bearer tokens because someone spread `req.body` and `req.headers` into a log call. Both problems come from the same missing rule: nobody defined what a log line must and must not carry."
          rule="Every log line carries these keys: `op` (dotted operation name like `checkout.charge`), `requestId` (and `sessionId` when present), the primary entity id for the operation (`orderId`, `userId`, `messageId`), and a small set of safe scalar inputs relevant to debugging (amount, currency, attempt, upstream, `code`). Never log passwords, tokens, full card numbers, CVV, full email, full address, or any header that carries an auth token. `userId` is the only PII allowed, and only because it is the correlation key. Per 02-spec/17/31 line 77."
          doThis="Add a `createLogger(op)` factory in `src/lib/logger.ts` that closes over `op` and requires `requestId` on every call, so the shape cannot regress. Add a redaction allowlist: only whitelisted keys pass through, everything else is dropped with a `[redacted]` marker. Add a CI grep that fails the build on `log\\.\\w+\\(.*req\\.(body|headers)` and `log\\.\\w+\\(.*(password|cvv|cardNumber|authorization)`. Sweep existing sites once and re-emit them through the factory."
        />
      </div>
    </SlideLayout>
  );
}
