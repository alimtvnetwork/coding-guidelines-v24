import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 OPS-003: runbooks. Every page-severity alert has an executable runbook.
 * Closes the Ops chapter and the SS-02 rule catalogue.
 */

const BEFORE = `# Confluence page, "Checkout troubleshooting", last updated 14 months ago.
# Free-form prose: "check the logs, restart if needed, ping @alice".
# Alice left the company 8 months ago. The log path in step 2 is wrong
# (service moved to a new host). Step 4 says "run the fix script" but the
# script is not linked and does not exist under that name any more.
# At 3am the on-call reads all four screens, gives up, and pages the CTO.`;

const AFTER = `# ops/runbooks/checkout-api/p99-latency-burn.md
# Every runbook: same shape, every step is either a copy-pastable command or a UI click path.
#
# ---
# alert: CheckoutApiP99LatencyBurn
# owner: team-checkout
# last-verified: 2026-07-14   # updated on every game-day drill
# ---
#
# ## Symptom (what the pager just told you)
# checkout-api p99 latency above 400ms SLO for 10 min. Dashboard link.
#
# ## Blast radius (who is affected right now)
# All checkout traffic. Revenue impact ~ $2k per minute at peak.
#
# ## First 5 minutes (assess, do not fix yet)
# 1. Open dashboard: <link>. Confirm the burn is still active.
# 2. Check deploys: `kubectl rollout history deploy/checkout-api -n prod`.
#    If a deploy landed in the last 30 min, jump to "Recent deploy" branch.
# 3. Check upstream: open the payments-gateway dashboard, look for correlated latency.
#
# ## Mitigations (pick the first that applies)
# - Recent deploy in last 30 min:
#   `kubectl rollout undo deploy/checkout-api -n prod`   # rollback, ~40s
# - Upstream (payments) is the cause:
#   Page team-payments via <pager-link>. Enable degraded-checkout flag: <link>.
# - No deploy, no upstream, box hot (CPU > 85 percent for > 5 min):
#   `kubectl scale deploy/checkout-api --replicas=+2 -n prod`   # add headroom
#
# ## After the fire is out
# 1. File incident ticket with root cause + this runbook's outcome.
# 2. Add or tune whatever the runbook was missing. Bump last-verified.`;

export default function RunbooksSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 66 · Ops · every page-severity alert has an executable runbook"
      title="Runbooks are code: same shape, versioned, drilled quarterly. Every step is a command or a click path, not prose."
      subtitle="A runbook that says 'check the logs and restart if needed' is not a runbook, it is a hope. The on-call at 3am needs: what happened, who is affected, what to check first, and the exact command to run to stop the bleeding. In that order."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="md"
          before={BEFORE}
          after={AFTER}
          beforeLabel="Bad. Free-form Confluence page from 14 months ago. Author left 8 months ago. Log path is stale, referenced script does not exist, no blast-radius statement, no mitigation ranked by likelihood, no rollback command, no last-verified date. The on-call at 3am spends the first 20 minutes verifying the runbook is even accurate before touching production, then pages the CTO because they cannot tell whether it is safe to restart. Root cause: runbooks were prose in a wiki with no owner, no schema, no drill cadence, no verification date."
          afterLabel="Good. Runbook lives in the repo at `ops/runbooks/{service}/{alert-slug}.md`, front-matter declares alert name, owner, and `last-verified` date. Every runbook has the same five sections in the same order: Symptom, Blast radius, First 5 minutes, Mitigations (ranked), After. Every mitigation is a copy-pastable command or an explicit click path, never 'check the logs'. Drilled quarterly in game-days; `last-verified` is updated on every drill."
        />
        <ActionPanel
          slideId="65-runbooks"
          symptom="Q1 pager retrospective: mean time-to-mitigate was 34 minutes, but only 12 minutes for incidents whose alert linked a runbook that was verified within the last 90 days, versus 51 minutes for incidents whose runbook was stale or missing. Three incidents in the last quarter escalated to executive paging because the on-call could not tell whether the runbook's rollback command was still safe (the deployment tool had changed and the runbook was 11 months old). One incident had the on-call run an outdated script that made the outage worse. Root cause: runbooks were unstructured prose, unowned, unversioned, and never drilled."
          rule="Every alert of `severity=page` (OPS-002) MUST have a runbook at `ops/runbooks/{service}/{alert-slug}.md` in the repo. Required front-matter: `alert` (matches OPS-002 alert name exactly), `owner` (matches `services.json`), `last-verified` (ISO date). Required sections in this exact order and no others: (1) Symptom, one sentence, mirrors the pager message; (2) Blast radius, who is affected and rough revenue or user impact; (3) First 5 minutes, assess-only steps, never mutating; (4) Mitigations, ranked by likelihood, each entry is either a copy-pastable command OR an explicit UI click path, never free-form prose like 'check the logs'; (5) After, what to file and what to tune. Every runbook is drilled at least once per quarter in a scheduled game-day; the drill run updates `last-verified` and any step that did not work. Runbooks older than 180 days without a `last-verified` refresh block their owning service from a release (WF-005 gate)."
          doThis="Enforce mechanically: (1) `scripts/validate-runbooks.mjs` in CI parses every `ops/runbooks/**/*.md`, fails on missing front-matter fields, fails on section headers not matching the required order, fails on any mitigation bullet that does not start with a code fence or an explicit 'click' path; (2) cross-check: every `severity=page` alert in `ops/alerts/*.yml` must have a matching runbook file (CI hard-fail on drift); (3) `ops/runbooks/_template.md` is the starting point for new runbooks, copy-and-fill; (4) game-day scheduler: `scripts/schedule-gameday.mjs` picks the oldest-verified runbook every two weeks and files a ticket for the owning team; the drill outcome PR updates `last-verified` in the same commit; (5) `WF-005` release script reads `last-verified` for every runbook owned by the releasing service and blocks the release if any is older than 180 days; (6) post-incident review (RCA-002) has a required question 'was the runbook accurate' and adds an action item if the answer is no; (7) on-call handoff opens the top-3 most-fired runbooks of the past week for a 10-minute read-through."
        />
      </div>
    </SlideLayout>
  );
}
