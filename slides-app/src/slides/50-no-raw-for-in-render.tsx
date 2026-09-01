import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 52: no raw `for` or `forEach` in render or derived state.
 *
 * Source: 02-spec/17/31 line 104.
 */

const BEFORE = `// Raw for + forEach mutating arrays declared during render.
export function InvoiceTable({ invoices, filters }: InvoiceTableProps) {
  const rows: Row[] = [];
  for (let i = 0; i < invoices.length; i++) {
    const inv = invoices[i];
    if (inv.Status === "draft") continue;
    rows.push({ Id: inv.Id, Total: inv.Total, Label: formatLabel(inv) });
  }

  const totalsByCurrency: Record<string, number> = {};
  invoices.forEach((inv) => {
    if (!totalsByCurrency[inv.Currency]) totalsByCurrency[inv.Currency] = 0;
    totalsByCurrency[inv.Currency] += inv.Total;
  });

  const memoRows = useMemo(() => rows, [rows]);

  return <Grid rows={memoRows} totals={totalsByCurrency} filters={filters} />;
}`;

const AFTER = `// Expression-based iteration. Stable references. Memo actually memoizes.
export function InvoiceTable({ invoices, filters }: InvoiceTableProps) {
  const rows = useMemo<readonly Row[]>(
    () =>
      invoices
        .filter((inv) => inv.Status !== "draft")
        .map((inv) => ({ Id: inv.Id, Total: inv.Total, Label: formatLabel(inv) })),
    [invoices],
  );

  const totalsByCurrency = useMemo<Readonly<Record<string, number>>>(
    () =>
      invoices.reduce<Record<string, number>>((acc, inv) => {
        acc[inv.Currency] = (acc[inv.Currency] ?? 0) + inv.Total;

        return acc;
      }, {}),
    [invoices],
  );

  return <Grid rows={rows} totals={totalsByCurrency} filters={filters} />;
}`;

export default function NoRawForInRenderSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 52 · React · no raw `for` or `forEach` in render"
      title="Use `map`, `filter`, `reduce`, `flatMap`, `Array.from`. Iteration returns an expression, never mutates."
      subtitle="02-spec/17/31 line 104: avoid raw `for` and `forEach` in render or derived state. Iteration must produce a value, not push into an array declared during render. `for` is only acceptable for early-exit performance on very large arrays with a comment explaining why."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="tsx"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ `rows` and `totalsByCurrency` are declared fresh on every render, then mutated by `for` and `forEach`. `useMemo(() => rows, [rows])` is a lie: `rows` is a new array reference every render, so the memo re-runs every render and every downstream `React.memo(Grid)` bails out. A stray `continue` inside the `for` silently drops rows and reviewers cannot see it in a diff. `totalsByCurrency` uses truthy-check `if (!totalsByCurrency[...])` which treats an existing `0` as missing, resetting the accumulator."
          afterLabel="✅ `filter().map()` and `reduce()` are expressions with stable references when `invoices` is stable. `useMemo` now caches correctly across renders. `readonly Row[]` and `Readonly<Record<...>>` block accidental mutation downstream. `acc[inv.Currency] ?? 0` uses nullish coalescing so an existing `0` stays `0`. No mid-render mutation, no hidden `continue`, no reference churn."
        />
        <ActionPanel
          slideId="50-no-raw-for-in-render"
          symptom="`InvoiceTable` renders 60 times per second when the parent hovers, `Grid` (wrapped in `React.memo`) re-renders every one of those, and DevTools Profiler blames 'props changed: rows, totals'. The offending code is a `for` loop that pushes into `const rows: Row[] = []` declared during render, plus a `forEach` that mutates `totalsByCurrency`. A later PR added `if (inv.Status === 'draft') continue;` inside the `for` to hide draft invoices, and QA still cannot figure out why one specific draft still appears (answer: a second `continue` earlier in the loop short-circuited past the filter). The `useMemo` wrapping `rows` gives a false sense of caching because `rows` is a fresh array reference on every render."
          rule="02-spec/17/31 line 104 forbids raw `for` and `forEach` in render or derived state. Every iteration in render, in a `useMemo`, in a selector, or in a derived-state helper must be an expression: `map`, `filter`, `reduce`, `flatMap`, `Array.from`, or `Object.entries().map(...)`. Never declare an array or object during render and mutate it. `for` is only acceptable when (a) the array is very large, (b) you need early-exit performance, and (c) a comment explains both. `forEach` is never acceptable in render because it returns nothing and encourages mutation. This rule pairs with line 105 (never mutate state/props/hook returns) and line 106 (stable keys) as the React-render-correctness trio."
          doThis="Ship the enforcement in one PR: (1) ESLint rule `no-raw-loop-in-render` that flags `ForStatement`, `ForOfStatement`, `ForInStatement`, and `CallExpression[callee.property.name='forEach']` inside any function whose name is PascalCase (component) or starts with `use` (hook) or is the callback of `useMemo`/`useCallback`; (2) auto-fix suggestions to `map`/`filter`/`reduce`; (3) waiver comment `// eslint-disable-next-line no-raw-loop-in-render -- perf: early exit on N over 10k, see BENCH-042` required with a benchmark link; (4) codemod `scripts/codemods/for-to-map.ts` for the initial sweep; (5) DevTools Profiler check on the top 20 hot components, confirm `React.memo` bailouts stop firing on unrelated parent renders. Reviewers reject any new `for` or `forEach` in a `.tsx` file without the waiver comment."
        />
      </div>
    </SlideLayout>
  );
}
