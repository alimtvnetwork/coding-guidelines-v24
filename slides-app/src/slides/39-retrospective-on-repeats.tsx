import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 41: Retrospective on repeats.
 *
 * Source: spec/17/31 line 79. Same bug class twice, stop and write
 * an RCA before the third fix.
 */

const BEFORE = `# Sprint 12
BUG-812  Checkout crashes when API returns { error: "..." }
FIX      try { ... } catch { toast("Something went wrong") }   // ship it

# Sprint 14
BUG-901  Profile page crashes when API returns { error: "..." }
FIX      try { ... } catch { toast("Something went wrong") }   // ship it

# Sprint 16
BUG-988  Settings page crashes when API returns { error: "..." }
FIX      try { ... } catch { toast("Something went wrong") }   // ship it
# Three fixes, same root cause, still no shared envelope parser.`;

const AFTER = `# Sprint 14, second occurrence detected -> STOP, write RCA before fix #3.

docs/rca/2026-07-19-envelope-parse-repeat.md
  Symptom      3 pages crash on { error } responses (BUG-812, BUG-901).
  Root cause   No shared apiCall parser. Each caller re-implements
               happy-path JSON access, ignores { errors[] } branch.
  Class fix    Introduce apiCall<T>() (ERR-005), migrate all 14 fetch
               sites in one PR, add ESLint ban on raw fetch().json().
  Prevention   CI grep for /\\.json\\(\\)/ outside src/lib/apiCall.ts.
  Owners       @frontend-platform, due next sprint, tracked in P0 board.

# Sprint 16 BUG-988 does not exist. The class is closed.`;

export default function RetrospectiveOnRepeatsSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 41 · Retrospective on repeats"
      title="Same bug twice? Stop. Write the RCA before the third fix."
      subtitle="One occurrence is a bug. Two is a pattern. Patching the third instance the same way you patched the first two is not engineering, it is bookkeeping. Stop the individual fix, name the class, ship a class-level fix, and add a check that prevents the next occurrence."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="text"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ Three sprints, three tickets, three copy-pasted try/catch toasts. Same root cause (no shared envelope parser) reappears every sprint in a new module. Symptom hidden, root cause untouched, rework compounds."
          afterLabel="✅ On the second occurrence, work stops. An RCA names the class, prescribes a shared fix (apiCall<T>() per ERR-005), migrates every site in one PR, and adds a lint rule that fails CI on the next raw fetch().json(). The third bug never files."
        />
        <ActionPanel
          slideId="39-retrospective-on-repeats"
          symptom="Three sprints in a row a different page crashes on `{ error: '...' }` responses. Each ticket gets the same one-line fix: wrap the caller in try/catch and show a generic toast. Nobody stops to ask why the same shape keeps landing, so the fourth page ships next month with the same bug, and the shared parser that would kill the class is still not written."
          rule="When the same bug class lands a second time (same root cause in a different file, same shape from a different endpoint, same race in a different reducer), STOP before writing fix #2. Open a short RCA under `docs/rca/YYYY-MM-DD-<slug>.md` naming symptom, root cause, class fix, prevention (lint/CI/test), and owners. Ship the class fix in one PR that migrates every known site, plus the prevention check that fails CI on the next occurrence. Only reopen individual-fix mode after the RCA is merged. Per spec/17/31 line 79."
          doThis="Add a triage question to every bug: `Have we fixed this class before?` If yes, label the ticket `rca-required`, freeze other work on that class, and assign an owner to write the RCA using the template in `docs/rca/_template.md` (symptom, root cause in ONE sentence, class fix, prevention check, owners, due date). The class-fix PR must migrate all known sites in one atomic change and add a CI rule (lint, grep, or test) that fails on the next occurrence. Review the `rca-required` label in weekly retro; anything older than one sprint escalates."
        />
      </div>
    </SlideLayout>
  );
}
