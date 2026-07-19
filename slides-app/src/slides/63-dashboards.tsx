import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 OPS-001: dashboards. Every service ships a golden-signals dashboard.
 * Chapter J (Ops) opener.
 */

const BEFORE = `// We had "monitoring" but no dashboard.
// Investigating an incident meant SSH-ing to the box and tailing logs by hand,
// while a customer waited on Zoom. Nobody could answer "is latency up right now?"
// because there was no chart of latency, only 40 GB of raw JSON logs.`;

const AFTER = `// Every service ships a golden-signals dashboard under docs/dashboards/<service>.json,
// version-controlled, provisioned by CI. Four panels, always in the same order:
//
//   1. Traffic       (requests per second, by route)
//   2. Errors        (5xx per second and 5xx as percent of total, by route)
//   3. Latency       (p50, p95, p99 in ms, by route)
//   4. Saturation    (CPU percent, memory percent, connection-pool usage percent)
//
// Plus one service-specific "domain" panel (queue depth, cache hit rate,
// or the single KPI that defines "is this service healthy").
//
// Every panel links to (a) the alert that fires on it (OPS-002) and
// (b) the runbook step that handles it (OPS-003).`;

export default function DashboardsSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 64 · Ops · every service ships a golden-signals dashboard"
      title="Traffic, Errors, Latency, Saturation. One dashboard per service, provisioned from code, linked to alerts and runbooks."
      subtitle="If you cannot answer 'is this service healthy right now?' in under 10 seconds by opening one URL, you do not have a dashboard, you have log files. Golden signals (Google SRE) are the non-negotiable four; add one domain panel for the KPI that defines the service."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="ts"
          before={BEFORE}
          after={AFTER}
          beforeLabel="Bad. 'Monitoring' meant logs. During the March incident the on-call spent 22 minutes correlating log timestamps with customer reports before anyone could say 'yes p99 latency is 12x normal'. There was no chart, no baseline, no way for a second responder to catch up without re-reading the same logs. Root cause: nobody had defined what 'healthy' looks like as numbers on a screen."
          afterLabel="Good. Every service has `docs/dashboards/<service>.json` (Grafana provisioning format), committed to the repo, applied by CI on merge to main. Four golden-signals panels always in the same order and same units, plus one domain panel. On-call opens one URL, reads four panels, knows in 10 seconds whether the alert is real, which route is affected, and whether the box is out of headroom. Every panel has a footer link to its alert (OPS-002) and its runbook step (OPS-003)."
        />
        <ActionPanel
          slideId="63-dashboards"
          symptom="Post-incident review of the last six production incidents: median time-to-first-hypothesis was 18 minutes, and 5 of 6 incidents had the on-call ask 'do we have a chart of X?' at least once. Two incidents had the wrong service investigated for the first 25 minutes because there was no traffic panel to show that the spike was in the API gateway, not the worker. Root cause: dashboards were either missing, inconsistent (each team invented its own panel order and units), or existed only in one engineer's browser bookmarks."
          rule="Every deployable service (api, worker, scheduler, gateway) ships a golden-signals dashboard at `docs/dashboards/<service>.json` in the repo, provisioned automatically by CI to the observability stack. The dashboard has exactly five panels, always in this order and with these units: (1) Traffic, requests per second, broken down by route; (2) Errors, 5xx per second AND 5xx as percent of total, broken down by route; (3) Latency, p50 + p95 + p99 in milliseconds, broken down by route; (4) Saturation, CPU percent + memory percent + primary-pool utilisation percent; (5) one Domain panel, the single KPI that defines 'is this service doing its job' (queue depth for a worker, cache hit rate for a cache, checkout completion rate for checkout). Every panel has a description field naming the SLI, the alert that watches it (OPS-002), and the runbook step that responds to it (OPS-003). Dashboards are code, reviewed in PRs, versioned in git; changes go through the same PR template (WF-004) as source code."
          doThis="Enforce mechanically: (1) `scripts/validate-dashboards.mjs` in CI checks every service listed in `services.json` has a matching `docs/dashboards/<service>.json` and that each dashboard has all five panels in the required order with the required units; (2) `docs/dashboards/_template.json` is the starting point for new services, copy-and-fill; (3) provisioning runs on merge to main via `.github/workflows/dashboards-provision.yml` and posts a diff comment on the PR showing what will change; (4) each service's readme has a required 'Dashboard' section linking to the provisioned URL; (5) on-call handoff checklist (in the runbook, OPS-003) starts with 'open the five service dashboards, confirm all-green'; (6) any new deployable service is blocked from merging until its dashboard file exists (CI hard-fail); (7) drift check: nightly job compares live panel config to the JSON in repo and opens an issue on drift, so no human 'quick fix' in the Grafana UI survives past 24 hours."
        />
      </div>
    </SlideLayout>
  );
}
