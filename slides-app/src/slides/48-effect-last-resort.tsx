import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 50: `useEffect` is the last resort, guards are positive booleans.
 *
 * Source: spec/17/31 lines 99-101.
 */

const BEFORE = `// Four effects, negative inline guards, effects used to derive state.
export function OrderPanel({ order, user }: OrderPanelProps) {
  const [total, setTotal] = useState(0);
  const [canCheckout, setCanCheckout] = useState(false);
  const [banner, setBanner] = useState<string | null>(null);

  // 1. Derive total from props via effect. Runs on every render.
  useEffect(() => { setTotal(order.items.reduce((s, i) => s + i.price, 0)); }, [order]);

  // 2. Derive a boolean from two props via effect, negative inline guard.
  useEffect(() => {
    if (!user || (!user.isVerified && !user.hasPaymentMethod)) setCanCheckout(false);
    else setCanCheckout(true);
  }, [user]);

  // 3. React to a click by kicking an effect (should be an event handler).
  useEffect(() => { if (total !== 0 && !banner) setBanner("Ready"); }, [total, banner]);

  // 4. Fetch on mount, no cleanup, refetches on every user change.
  useEffect(() => { fetch(\`/api/quote/\${order.id}\`).then(r => r.json()).then(setBanner); });

  return <Shell total={total} canCheckout={canCheckout} banner={banner} />;
}`;

const AFTER = `// Zero effects. Derived values. Positive booleans. Event handler for events.
export function OrderPanel({ order, user }: OrderPanelProps) {
  const total = useMemo(() => sumItemPrices(order.items), [order.items]);   // derived
  const isCheckoutReady = getIsCheckoutReady(user);                          // positive
  const [banner, setBanner] = useState<string | null>(null);

  function handleQuoteClick() {                                              // event, not effect
    apiCall(() => fetchQuote(order.id)).then((quote) => setBanner(quote.label));
  }

  return (
    <Shell total={total} canCheckout={isCheckoutReady} banner={banner} onQuote={handleQuoteClick} />
  );
}

function getIsCheckoutReady(user: User | null): boolean {
  if (!user) return false;
  return user.isVerified && user.hasPaymentMethod;   // positive, extracted, one place
}`;

export default function EffectLastResortSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 50 · React · effects are the last resort"
      title="Default is zero effects. Every remaining guard is a positively named boolean, extracted above the effect."
      subtitle="spec/17/31 lines 99-101: do not use `useEffect` to derive state, transform props, or react to user events. Use derived values, `useMemo`, or event handlers. When an effect is truly required (network, timer, subscription, DOM API), the condition is a positively named boolean, never an inline negative."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="tsx"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ Four effects doing work the render or an event handler should do. Effect 1 mirrors props into state (extra render, stale on the first paint). Effect 2 uses a negative inline guard `!user || (!user.isVerified && !user.hasPaymentMethod)` that reads backwards. Effect 3 sets a banner in response to a click, but reaches for it via an effect on the derived total. Effect 4 fetches with no dependency array and no cleanup, so it refires on every render and races itself."
          afterLabel="✅ Zero effects. `total` is a `useMemo` over the items. `isCheckoutReady` is a positively named boolean returned by `getIsCheckoutReady`, extracted so it is testable and reusable. The click becomes `handleQuoteClick`, a normal event handler wrapped in `apiCall`. No stale closures, no double fetches, no `!x && !y` to parse. If the component later needs to open a websocket, that one effect is added, with a cleanup, and its guard is another positive helper."
        />
        <ActionPanel
          slideId="48-effect-last-resort"
          symptom="`OrderPanel` has four `useEffect`s: one derives `total` from `order.items`, one derives `canCheckout` from `user` behind `!user || (!user.isVerified && !user.hasPaymentMethod)`, one sets a banner in response to a click, one fetches a quote with no dependency array and no cleanup. Symptoms in prod: totals flash the wrong value on first paint, checkout button flickers disabled then enabled, quote endpoint receives 6 requests per mount, and a stale banner from a previous order leaks into the next one. Nobody can tell which effect owns the banner state."
          rule="spec/17/31 lines 99-101 fix this in three parts. (1) Default `useEffect` count is zero. Add one only to synchronize with an external system: network, timer, subscription, DOM API. Deriving state, transforming props, or reacting to user events are NOT external systems, they are `useMemo`, plain expressions, or event handlers. (2) When an effect is truly required, the condition is a positively named boolean like `isReadyToSync` or `hasFreshData`, extracted above the effect. (3) No inline negatives, no nested ternaries, no `!x && y` in the effect body or in its dependency guard. If the natural check is negative, invert into a positive helper and early-return on the positive path."
          doThis="Land the enforcement now: (1) custom ESLint rule `no-derive-state-in-effect` that flags any `useEffect` whose body only calls `setState` from props or other state, autofix suggestion is `useMemo` or a plain const; (2) custom rule `no-negative-effect-guard` that flags `!` at the top of an effect body or a boolean expression containing `!` inside an effect condition, requires the guard to be a named identifier; (3) custom rule `require-effect-cleanup-when-async` that flags `useEffect(() =` bodies calling `fetch`/`setInterval`/`addEventListener` without returning a cleanup; (4) codemod `scripts/extract-effect-guard.mjs` that lifts inline conditions into a `getIs*` helper above the component; (5) CI log line `react.effect.violation` with `{ file, line, rule, effectIndex }` so repeats become measurable. PR checklist: 'Every new `useEffect`: name the external system it syncs with, name the positive guard, show the cleanup.'"
        />
      </div>
    </SlideLayout>
  );
}
