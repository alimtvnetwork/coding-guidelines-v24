import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 OPS-002: alert rules. Every alert has SLO, dashboard, runbook, owner.
 */

const BEFORE = `# Prometheus rule, ad-hoc, no context.
- alert: HighLatency
  expr: latency > 1
  for: 5m
  annotations:
    summary: "latency high"
# Pager fires at 3am. On-call reads "latency high", opens no dashboard (there is no link),
# has no runbook (there is no link), has no idea whose service it is (there is no owner label),
# and has no idea if 1 what (seconds? ms? p50? p99?). Twenty-two minutes lost to orientation.`;

const AFTER = `# Every alert declares: SLO tie, dashboard link, runbook link, owner, severity, layer.
- alert: CheckoutApiP99LatencyBurn
  expr: |
    histogram_quantile(0.99,
      sum by (le) (rate(http_request_duration_seconds_bucket{service="checkout-api"}[5m]))
    ) > 0.400
  for: 10m
  labels:
    severity: page          # page | ticket | log-only
    layer: api              # api | worker | data | edge
    owner: team-checkout    # matches services.json owner
    slo: checkout-api-latency-p99-400ms
  annotations:
    summary: "checkout-api p99 latency above 400ms SLO for 10 min"
    description: "Route breakdown: {{ $labels.route }}. Error budget burn rate: fast."
    dashboard: "https://grafana/dash/checkout-api?viewPanel=3"
    runbook: "https://runbooks/checkout-api#p99-latency-burn"`;

export default function AlertRulesSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 65 · Ops · every alert has SLO, dashboard, runbook, owner"
      title="Alerts fire on SLO burn, not raw thresholds. Every rule links dashboard, runbook, owner, severity, layer."
      subtitle="Alerts without runbooks are hazing. Alerts on raw thresholds (`cpu at 80 percent`) page for weather, not weather damage. Tie every alert to an SLO burn, and make the pager message carry every link the on-call needs in the first 30 seconds."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="yaml"
          before={BEFORE}
          after={AFTER}
          beforeLabel="Bad. `latency > 1` fires at 3am with no units, no dashboard link, no runbook link, no owner. On-call spends 22 minutes finding the right service, guessing at units, and paging a second person to ask 'is this normal for this hour'. Two of the last four incidents had the wrong on-call paged because the alert had no `owner` label; both took over 30 minutes to route correctly. Root cause: alert rules were free-form YAML with no required schema."
          afterLabel="Good. Every alert declares SLO tie, severity (page vs ticket vs log-only), layer, owner, dashboard URL, runbook URL. Alerts fire on SLO burn rate (multi-window, multi-burn), not on raw threshold noise. The pager message renders as 'checkout-api p99 latency above 400ms SLO for 10 min. Dashboard: ... Runbook: ...' so the on-call opens two tabs and is oriented before the second sip of coffee."
        />
        <ActionPanel
          slideId="64-alert-rules"
          symptom="Six-month pager audit: 43 percent of pages were false positives (raw threshold crossed but no user impact), 28 percent had no runbook link, 19 percent paged the wrong team on the first fire and had to be re-routed, 4 pages fired for the same underlying cause within one hour because there was no dedup and no SLO grouping. On-call satisfaction score dropped from 4.1 to 2.6 across the year. Root cause: alerts were written per-incident by whoever was firefighting, with no required schema, no SLO discipline, and no ownership metadata."
          rule="Every alert rule lives in `ops/alerts/{service}.yml`, reviewed in PRs like source code, and MUST declare: (1) `severity` from the fixed set `page`, `ticket`, `log-only` (page = wake a human, ticket = business hours, log-only = trend awareness); (2) `layer` from `api`, `worker`, `data`, `edge`; (3) `owner` matching a team id in `services.json`; (4) `slo` naming the SLO document it defends; (5) `annotations.summary` in the shape `{service} {signal} {condition} for {duration}`; (6) `annotations.dashboard` deep link to the exact panel in the OPS-001 dashboard; (7) `annotations.runbook` deep link to the OPS-003 runbook step. Page-severity alerts fire on SLO burn rate (multi-window, multi-burn: fast burn = 2 percent of monthly budget in 1 hour, slow burn = 10 percent in 6 hours), not on raw threshold crossings. Raw-threshold alerts are permitted only at `log-only` severity for trend awareness. Dedup: alerts group by `service` + `slo` so one root cause produces one page, not one page per instance."
          doThis="Enforce mechanically: (1) `scripts/validate-alerts.mjs` in CI parses every `ops/alerts/*.yml`, fails on any missing required label or annotation, fails on `severity=page` without SLO burn expression, fails on `owner` not present in `services.json`; (2) `ops/alerts/_schema.json` is the canonical schema, referenced from every alert file; (3) PR template checkbox (WF-004) 'alert rule has SLO tie, dashboard, runbook, owner'; (4) synthetic pager test in CI: for each `severity=page` alert, fire it in a staging pager and assert the message rendered contains the dashboard and runbook URLs (catches broken templates); (5) monthly pager-quality retro (RCA-002) reviews false-positive rate per service; any alert over 20 percent false-positive rate for two months is either tuned or downgraded to `ticket`; (6) new services blocked from production release (WF-005) until at least the four golden-signals alerts exist with page severity for Errors and Latency, ticket for Traffic dropout and Saturation."
        />
      </div>
    </SlideLayout>
  );
}
