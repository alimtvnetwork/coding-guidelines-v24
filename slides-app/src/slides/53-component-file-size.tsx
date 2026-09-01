import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 55: React component files under 100 lines.
 *
 * Source: 02-spec/17/31 line 107.
 */

const BEFORE = `// src/pages/Dashboard.tsx  --  312 lines.
export function Dashboard() {
  const [range, setRange] = useState<DateRange>(defaultRange);
  const [filter, setFilter] = useState<Filter>(defaultFilter);
  const { data: orders } = useOrders(range, filter);
  const { data: users } = useUsers();
  const { data: alerts } = useAlerts();

  const kpis = useMemo(() => computeKpis(orders), [orders]);
  const chartData = useMemo(() => toChartSeries(orders, range), [orders, range]);
  const alertRows = useMemo(() => alerts?.filter(isOpen) ?? [], [alerts]);
  // ... 40 more lines of derivations ...

  return (
    <Page>
      <Header title="Dashboard" range={range} onRangeChange={setRange} />
      <FilterBar filter={filter} onChange={setFilter} />
      <KpiStrip>
        {/* 60 lines of inline KpiCard JSX with copy-pasted styles */}
      </KpiStrip>
      <ChartPanel>
        {/* 80 lines of chart config, axes, tooltips inline */}
      </ChartPanel>
      <AlertsTable rows={alertRows}>
        {/* 70 lines of column defs and row renderers */}
      </AlertsTable>
    </Page>
  );
}`;

const AFTER = `// src/pages/Dashboard.tsx  --  38 lines. Composition only.
export function Dashboard() {
  const [range, setRange] = useState<DateRange>(defaultRange);
  const [filter, setFilter] = useState<Filter>(defaultFilter);
  const { orders, users, alerts } = useDashboardData(range, filter);

  return (
    <Page>
      <Header title="Dashboard" range={range} onRangeChange={setRange} />
      <FilterBar filter={filter} onChange={setFilter} />
      <KpiStrip orders={orders} />
      <ChartPanel orders={orders} range={range} />
      <AlertsTable alerts={alerts} />
    </Page>
  );
}

// src/pages/dashboard/useDashboardData.ts     42 lines, one custom hook.
// src/pages/dashboard/KpiStrip.tsx            58 lines, four KpiCard children.
// src/pages/dashboard/KpiCard.tsx             34 lines.
// src/pages/dashboard/ChartPanel.tsx          71 lines.
// src/pages/dashboard/AlertsTable.tsx         74 lines, columns in columns.ts.
// src/pages/dashboard/columns.ts              28 lines, ColumnDef[] only.`;

export default function ComponentFileSizeSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 55 · React · component files under 100 lines"
      title="Extract child components, hooks, and helpers into their own files before the component grows past 100."
      subtitle="02-spec/17/31 line 107: keep component files under 100 lines. 100 is a smoke alarm, not a fire code. It fires before the file becomes unreviewable, unmergeable, and untestable, at exactly the moment extraction is cheap."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="tsx"
          before={BEFORE}
          after={AFTER}
          beforeLabel="Bad. One 312-line `Dashboard.tsx` mixes data fetching, three unrelated `useMemo` derivations, header layout, KPI cards, chart config, and table column defs. PR diffs are unreadable, two people touching different sections conflict on every merge, `React.memo` boundaries are impossible to draw because everything closes over the same locals, and unit tests have to mount the whole page to test one KPI computation. `SIZE-001` (slide 27) would already flag this at 300, but by then the damage is done: extraction requires a rewrite instead of a rename."
          afterLabel="Good. `Dashboard.tsx` is 38 lines of composition, each child is its own file under 100 lines, and the data-fetch trio collapses into `useDashboardData`. Each file is now a review-in-one-sitting unit, `KpiStrip` and `ChartPanel` can be memoized independently, `columns.ts` is testable without React, and two engineers can edit two panels in parallel without conflicts. The 100-line React cap is tighter than the generic 300 (SIZE-001) precisely because JSX + hooks + handlers stack density is roughly 3x plain code."
        />
        <ActionPanel
          slideId="53-component-file-size"
          symptom="`Dashboard.tsx` grew from 90 lines to 312 over six weeks, one 'quick addition' at a time. Symptoms compounding: (1) PR reviewer says 'I can't hold this in my head, LGTM I guess'; (2) two engineers ship conflicting merges every sprint because everyone edits the same file; (3) profiler shows the whole page re-rendering on every keystroke in `FilterBar` because there is no memo boundary to draw; (4) Storybook has no story for `KpiCard` alone because `KpiCard` isn't a component, it's 60 lines of inline JSX; (5) the unit test for `computeKpis` has to import `Dashboard` and mock five hooks. Every symptom shares one root cause: the component was never extracted, so it accumulated."
          rule="02-spec/17/31 line 107 is absolute: React component files stay under 100 lines. This is stricter than the generic 300-line file cap (SIZE-001) because JSX plus hooks plus event handlers packs about 3x the semantic density of plain code, and because React memoization only works at component boundaries. When a file crosses 80 lines, extraction stops being cheap; when it crosses 100, extraction is refactoring. Extract in this order: (1) leaf JSX blocks that render 3+ elements or take 15+ lines into their own child component file; (2) derivation `useMemo`s and data-fetch clusters into a single named custom hook (`useDashboardData`, `useOrderForm`); (3) event handlers longer than 8 lines into helper functions in a sibling file; (4) column definitions, form schemas, chart configs into data files that don't import React at all. The 100-line rule pairs with FUNC-001 (function length 8/15) and DEF-001 (dedicated definitions files) as one system."
          doThis="Enforce mechanically: (1) ESLint `max-lines` with `{ max: 100, skipBlankLines: true, skipComments: true }` scoped to `**/*.tsx` files via override; (2) CI gate `component-size` that counts only `.tsx` under `src/components/**` and `src/pages/**` and blocks the merge with a diff showing which extractions would land under the cap; (3) waiver format identical to `SIZE-001`: inline `// lint-allow: file-length reason='generated route module' max=180` with a real reason and an explicit max, only for generated code and one-shot data tables (never for hand-written components); (4) IDE hint at 80 lines (yellow gutter marker in `.vscode/settings.json` via a Problems provider) so extraction happens before the alarm; (5) code review reflex: any PR that adds 30+ lines to a component over 70 lines is asked 'what's the extracted child called?' before approve. When you cannot name the extracted child, that is the signal you have not understood the component yet (02-spec/17/31 line 112)."
        />
      </div>
    </SlideLayout>
  );
}
