import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 36: Registered error codes.
 *
 * Source: 02-spec/17/31 line 79. Every user-visible error carries a stable
 * code from a central registry. No ad-hoc codes invented at the throw site.
 */

const BEFORE = `// src/features/checkout/pay.ts
throw new AppError("bad-input", { field: "card" });        // ❌ kebab-case
throw new AppError("BAD_INPUT", { field: "cvv" });         // ❌ SCREAM
throw new AppError("payment_failed_2", { retry: false });  // ❌ ad-hoc suffix

// src/features/checkout/refund.ts
throw new AppError("PaymentFailed");                        // ❌ 4th spelling
// UI switch (err.code) { ... } silently falls to "unknown". No log.`;

const AFTER = `// src/errors/registry.ts  (single source of truth)
export const ErrorCodes = {
  CheckoutInvalidCard:    "CheckoutInvalidCard",
  CheckoutInvalidCvv:     "CheckoutInvalidCvv",
  CheckoutPaymentFailed:  "CheckoutPaymentFailed",
  CheckoutRefundRejected: "CheckoutRefundRejected",
} as const;
export type ErrorCode = typeof ErrorCodes[keyof typeof ErrorCodes];

// src/features/checkout/pay.ts
throw new AppError(ErrorCodes.CheckoutInvalidCard, { field: "card" });
throw new AppError(ErrorCodes.CheckoutPaymentFailed, { retry: false });

// UI (exhaustive switch, compiler catches missing branches)
switch (err.code) {
  case ErrorCodes.CheckoutInvalidCard:    return <InvalidCardBanner />;
  case ErrorCodes.CheckoutPaymentFailed:  return <RetryBanner />;
  // ...
}`;

export default function RegisteredErrorCodesSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 36 · Registered error codes"
      title="One registry. One spelling. One switch."
      subtitle="Every user-visible error carries a code that lives in a central `ErrorCodes` registry. No throw site invents a new string. The UI switch is exhaustive, and the compiler catches the day someone adds a code without a matching banner."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="typescript"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ Four spellings of the same error across two files. UI falls through to 'unknown', support tickets can't be searched."
          afterLabel="✅ `ErrorCodes` registry, typed `ErrorCode` union, exhaustive UI switch. Codes are greppable everywhere."
        />
        <ActionPanel
          slideId="34-registered-error-codes"
          symptom="A customer reports 'payment failed', support searches the logs for `payment_failed` and finds nothing because the code was thrown as `PaymentFailed` on Tuesday and `payment_failed_2` on Wednesday. The UI banner is the fallback 'Something went wrong', so the user has no idea whether to retry, edit their card, or contact their bank. Every incident becomes a code archaeology dig."
          rule="Every user-visible error uses a code from a single central registry (for example `src/errors/registry.ts` exporting `ErrorCodes` plus an `ErrorCode` union type). Throw sites reference `ErrorCodes.CheckoutInvalidCard`, never a string literal. UI mapping is an exhaustive `switch` on `ErrorCode` so the TypeScript compiler fails the build when a new code lands without a matching branch. Per 02-spec/17/31 line 79 and the error-management digest."
          doThis="Create `src/errors/registry.ts` (or extend the existing one) and add every new code there with a PascalCase name that includes its feature prefix (`CheckoutInvalidCard`, not `InvalidCard`). Import `ErrorCodes` at the throw site. In the UI, type the switch parameter as `ErrorCode` and add a `never`-typed default branch so missing cases fail type-check. Reject any PR that types an error code as a bare string."
        />
      </div>
    </SlideLayout>
  );
}
