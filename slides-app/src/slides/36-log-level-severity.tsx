import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 38: Log-level severity map.
 *
 * Source: 02-spec/17/31 line 76. Five levels, strict semantics.
 * debug = trace, info = lifecycle, warn = recoverable,
 * error = user-visible failure, fatal = process exit only.
 */

const BEFORE = `// Everything is "error" because it feels safe
log.error("retrying upstream call", { attempt: 2 });   // ❌ recoverable
log.error("user logged in", { userId });                // ❌ lifecycle
log.error("cache miss on product page", { productId }); // ❌ trace

// Lifecycle noise at warn drowns real problems
log.warn("app started on port 8080");                   // ❌ info
log.warn("db pool initialized, size=20");               // ❌ info

// Fatal used for a caught exception, process keeps running
try { await charge(order); }
catch (e) { log.fatal("charge failed", e); }            // ❌ recoverable, not fatal`;

const AFTER = `// debug: trace, high-volume, off in prod by default
log.debug("cache miss on product page", { productId, requestId });

// info: lifecycle you would tell a new engineer about
log.info("app started", { port: 8080, commit });
log.info("user logged in", { userId, requestId });

// warn: recoverable, retry worked, degraded but serving
log.warn("upstream retry succeeded", { attempt: 2, url, requestId });

// error: user saw a failure or a request did not complete
log.error("charge failed after retries", {
  orderId, requestId, code: ErrorCodes.CheckoutInvalidCard,
});

// fatal: process is about to exit, on-call gets paged
log.fatal("cannot bind to port, exiting", { port, cause });
process.exit(1);`;

export default function LogLevelSeveritySlide() {
  return (
    <SlideLayout
      eyebrow="Rule 38 · Log-level severity map"
      title="Five levels. One meaning each. No sliding scale."
      subtitle="`debug` for trace, `info` for lifecycle, `warn` for recoverable, `error` for user-visible failure, `fatal` only when the process is about to exit. Calibrate the level to what an on-call human should do about it, not to how nervous the author feels."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="typescript"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ Everything at error, lifecycle at warn, caught exception at fatal. Dashboards go red, on-call ignores alerts, real fatals get missed."
          afterLabel="✅ Each level maps to a distinct human action. Debug is silent in prod, info is lifecycle, warn is recovered, error is user-visible, fatal pages on-call."
        />
        <ActionPanel
          slideId="36-log-level-severity"
          symptom="The error dashboard shows 40k events per hour, 95 percent are successful retries or lifecycle chatter. On-call has muted the channel, so when the payment worker actually died last Tuesday nobody saw the fatal for 90 minutes. New engineers copy whatever level the nearest line used, so calibration drifts further every sprint."
          rule="Pick the level by asking: what should a human do about this log line? Nothing, that is `debug`. Just know it happened, that is `info`. Retry worked but keep an eye on it, that is `warn`. A user saw a failure or a request did not complete, that is `error`. The process is about to exit and on-call should be paged, that is `fatal`. Fatal is followed by process termination, always. Per 02-spec/17/31 line 76."
          doThis="Add the five-line severity map to `src/lib/logger.ts` as a top-of-file comment so every author sees it. In code review, reject any `log.error` where a retry succeeded, any `log.warn` for startup or shutdown, any `log.fatal` that is not immediately followed by `process.exit` or an unrecoverable crash. Sweep the current codebase once with `rg 'log\\.(warn|error|fatal)'` and re-calibrate the top 20 noisiest sites first."
        />
      </div>
    </SlideLayout>
  );
}
