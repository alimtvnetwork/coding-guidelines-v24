import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 51: one effect, one concern, always with cleanup.
 *
 * Source: spec/17/31 lines 102-103.
 */

const BEFORE = `// One effect doing three things, none of them cleaned up.
export function LiveDashboard({ roomId }: LiveDashboardProps) {
  const [messages, setMessages] = useState<Message[]>([]);
  const [tick, setTick] = useState(0);
  const [seed, setSeed] = useState<Seed | null>(null);

  useEffect(() => {
    // Concern A: open a websocket subscription.
    const socket = new WebSocket(\`wss://api/rooms/\${roomId}\`);
    socket.onmessage = (e) => setMessages((prev) => [...prev, JSON.parse(e.data)]);

    // Concern B: start a 1-second timer for the header clock.
    const timer = setInterval(() => setTick((n) => n + 1), 1000);

    // Concern C: one-shot fetch for the initial seed.
    fetch(\`/api/rooms/\${roomId}/seed\`).then((r) => r.json()).then(setSeed);

    // No return. No cleanup. Socket and interval leak on every roomId change.
  }, [roomId]);

  return <Shell messages={messages} tick={tick} seed={seed} />;
}`;

const AFTER = `// Three effects, one concern each, every resource cleaned up.
export function LiveDashboard({ roomId }: LiveDashboardProps) {
  const [messages, setMessages] = useState<Message[]>([]);
  const [tick, setTick] = useState(0);
  const [seed, setSeed] = useState<Seed | null>(null);

  // Concern A: websocket subscription. Cleanup closes it.
  useEffect(() => {
    const socket = new WebSocket(\`wss://api/rooms/\${roomId}\`);
    socket.onmessage = (e) => setMessages((prev) => [...prev, JSON.parse(e.data)]);
    return () => socket.close();
  }, [roomId]);

  // Concern B: header clock. Cleanup clears the interval.
  useEffect(() => {
    const timer = setInterval(() => setTick((n) => n + 1), 1000);
    return () => clearInterval(timer);
  }, []);

  // Concern C: initial seed fetch. Cleanup aborts the request.
  useEffect(() => {
    const controller = new AbortController();
    fetchSeed(roomId, controller.signal).then(setSeed).catch(logIfNotAbort);
    return () => controller.abort();
  }, [roomId]);

  return <Shell messages={messages} tick={tick} seed={seed} />;
}`;

export default function OneEffectOneConcernSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 51 · React · one effect, one concern, always cleaned up"
      title="Split unrelated subscriptions and fetches into separate effects. Every effect that acquires a resource returns a cleanup."
      subtitle="spec/17/31 lines 102-103: if an effect does two unrelated things, split it. Every effect that acquires a resource (socket, timer, subscription, fetch, DOM listener) must return a cleanup function. No exceptions."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="tsx"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ One `useEffect` opens a websocket, starts a 1s interval, and fires a fetch. Zero returns. Every `roomId` change spawns a fresh socket while the old one keeps pushing messages into a stale `setMessages`. Intervals stack: after five room switches the clock ticks 5x per second. The fetch resolves after unmount and triggers a `setState on unmounted component` warning that hides the real crash below it."
          afterLabel="✅ Three effects, one concern each. Socket has `socket.close()` cleanup. Interval has `clearInterval(timer)`. Fetch uses `AbortController` and aborts on unmount or `roomId` change. `logIfNotAbort` swallows only the abort signal and logs everything else with context. Switching rooms 50 times leaves exactly one socket, one interval, one in-flight fetch."
        />
        <ActionPanel
          slideId="49-one-effect-one-concern"
          symptom="`LiveDashboard` has one `useEffect` that (a) opens `new WebSocket(...)`, (b) starts `setInterval(..., 1000)`, and (c) fires `fetch('/api/.../seed')`, all keyed on `roomId`. There is no `return`. Prod symptoms: memory grows 4MB per room switch, the clock speeds up (2x, 3x, 5x) as users navigate, React logs `Warning: Can't perform a React state update on an unmounted component` on every route change, and the server sees duplicate WS connections from the same user. Nobody can tell whether the leak is the socket, the timer, or the fetch because the concerns are tangled."
          rule="spec/17/31 lines 102-103 are two rules that always ship together. (1) One effect, one concern. If an effect does two unrelated things, split it. Never combine unrelated subscriptions or fetches in one effect body, even if the dependency arrays look the same. (2) Every effect that acquires a resource must return a cleanup function. Sockets get `.close()`. Intervals get `clearInterval`. Subscriptions get `unsubscribe`. DOM listeners get `removeEventListener`. Fetches get `AbortController` and abort on unmount. No exceptions, not for 'small' effects, not for 'always-mounted' components, not for mount-only fetches."
          doThis="Land the enforcement now: (1) custom ESLint rule `one-concern-per-effect` that flags any effect body containing two or more of `new WebSocket`, `setInterval`, `setTimeout`, `addEventListener`, `fetch`, `subscribe`, and demands a split; (2) custom ESLint rule `require-effect-cleanup` that flags any effect creating a socket/interval/timeout/listener/AbortController without a matching cleanup in the returned function; (3) a runtime dev-only wrapper `useTrackedEffect` in `src/lib/effects.ts` that counts live resources per component and logs `react.effect.leak` with `{ component, resource, count }` when a component unmounts with a non-zero balance; (4) PR checklist line: 'For every new effect: what one concern, what cleanup, verified in devtools that resource count returns to zero.' Reviewers reject any effect body with two acquisitions or no return."
        />
      </div>
    </SlideLayout>
  );
}
