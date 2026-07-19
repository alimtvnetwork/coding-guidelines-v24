import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 27: Enum standards across languages.
 *
 * Source: spec/17/31 line 31 (no magic strings), line 168 (Go `type X byte`
 * plus `iota`), line 171 (PHP `->isEqual()`). Cross-language rule: PascalCase
 * enum type + PascalCase members, strict parsing, comparison via typed value.
 */

const BEFORE = `// TypeScript
type Status = "pending" | "ready" | "failed";
if (row.status === "ready") { ship(row); }

// Go
const statusReady = "ready"
if row.Status == statusReady { ship(row) }

// PHP
if ($row->status === 'ready') { ship($row); }`;

const AFTER = `// TypeScript
export enum OrderStatus { Pending, Ready, Failed }
if (row.status === OrderStatus.Ready) { ship(row); }

// Go
type OrderStatus byte
const (
  OrderStatusPending OrderStatus = iota
  OrderStatusReady
  OrderStatusFailed
)
if row.Status == OrderStatusReady { ship(row) }

// PHP
if ($row->status->isEqual(OrderStatus::Ready)) { ship($row); }`;

export default function EnumStandardsSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 27 · Enum standards"
      title="PascalCase enums. Strict parse. Compare against the symbol, not the string."
      subtitle="One enum type per concept, PascalCase members, no string literal comparisons anywhere. Go uses `type X byte` + `iota`. PHP compares via `->isEqual()`, never `===`."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 18, marginTop: 4 }}>
        <CodeDiff
          language="typescript"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ string literals leaking across the codebase"
          afterLabel="✅ typed enum symbols, per-language idiom"
        />
        <ActionPanel
          slideId="25-enum-standards"
          symptom="You grep for the literal `ready` and find it in 14 files across TS, Go, and PHP. A typo like `redy` silently branches to the false path. Renaming the state requires a full-repo find-and-replace and still misses stringly-typed JSON boundaries."
          rule="Every finite set of values is a PascalCase enum type with PascalCase members. TS uses `enum`, Go uses `type X byte` + `iota`, PHP uses backed enums with `->isEqual()`. Parsing external input goes through a strict `ParseX` / `TryFromValue` that returns an error on unknown values. No string literal comparisons at call sites. Per spec/17/31 lines 31, 168, 171."
          doThis="When you touch a stringly-typed status/kind/type field, introduce the enum type in a dedicated file, add a strict parser at the JSON/DB boundary, and replace every string literal comparison with the enum symbol in the same diff. Add a switch with exhaustive default that throws on unknown members."
        />
      </div>
    </SlideLayout>
  );
}
