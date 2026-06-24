import { SlideLayout } from "@/components/SlideLayout";
import { CodeDiff } from "@/components/CodeDiff";
import { ActionPanel } from "@/components/ActionPanel";

const BEFORE = `const cache = new Map();

function GetUser(id) {
  if (cache.has(id)) return cache.get(id);
  const u = db.fetch(id);
  cache.set(id, u);
  return u;
}
// UpdateUser never clears cache → stale forever`;

const AFTER = `const cache = new TTLCache({ ttl: 60_000 });

function GetUser(id) {
  return cache.get(id) ?? cache.set(id, db.fetch(id));
}

function UpdateUser(id, patch) {
  db.update(id, patch);
  cache.delete(id);  // ← invalidate
}`;

export default function CacheInvalidationSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 11 · Architecture"
      title="Every cache needs a TTL and a mutation hook"
      subtitle="A cache without invalidation is a bug factory. Pair every read-cache with the write path that clears it."
    >
      <CodeDiff
        language="typescript"
        before={BEFORE}
        after={AFTER}
        beforeLabel="❌ Stale forever"
        afterLabel="✅ TTL + invalidation"
      />
      <ActionPanel
        symptom="Users see stale data after updates. The cache fills, never clears, and bugs are blamed on 'something weird'."
        rule="Caches require: explicit TTL + a documented mutation that calls cache.delete on the same key."
        doThis="Audit one cache.set in your code. Add a TTL, and wire cache.delete into the matching update/delete path."
      />
    </SlideLayout>
  );
}
