import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 32: Dedicated definitions file.
 *
 * Source: spec/17/31 line 32. Types, enums, constants, and interfaces get
 * their own file, not inline next to the first use. Applies to every
 * language; the file layout below uses TypeScript as the illustration.
 */

const BEFORE = `// src/features/orders/OrderCard.tsx
type OrderStatus = "pending" | "paid" | "shipped" | "cancelled";

const MAX_LINE_ITEMS = 50;
const FREE_SHIPPING_THRESHOLD = 50;

interface OrderCardProps {
  order: { id: number; status: OrderStatus; total: number };
  onCancel: (id: number) => void;
}

export function OrderCard({ order, onCancel }: OrderCardProps) {
  return <article>{/* ... */}</article>;
}`;

const AFTER = `// src/features/orders/types.ts
export type OrderStatus = "pending" | "paid" | "shipped" | "cancelled";
export interface OrderCardProps {
  order: Order;
  onCancel: (id: OrderId) => void;
}

// src/features/orders/constants.ts
export const MAX_LINE_ITEMS = 50;
export const FREE_SHIPPING_THRESHOLD = 50;

// src/features/orders/OrderCard.tsx
import type { OrderCardProps } from "./types";

export function OrderCard({ order, onCancel }: OrderCardProps) {
  return <article>{/* ... */}</article>;
}`;

export default function DedicatedDefinitionsSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 32 · Dedicated definitions files"
      title="Types, enums, constants live in their own file."
      subtitle="Inline definitions bind a type to the first component that used it, then everyone imports through that component. Extract on day one, not after the second consumer appears."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="typescript"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ Type, props interface, and constants inlined in the component file"
          afterLabel="✅ types.ts + constants.ts alongside the component; component file stays focused on JSX"
        />
        <ActionPanel
          slideId="30-dedicated-definitions"
          symptom="A second screen needs `OrderStatus`, so it imports it from `OrderCard.tsx`. Now the card file cannot be renamed or deleted without breaking unrelated modules. Circular imports appear the moment a util also wants the type. Reviewers stop noticing because 'it works.'"
          rule="Types, enums, interfaces, and typed constants each live in a dedicated file, not inline next to the first use. Colocate as `types.ts` / `constants.ts` next to the feature; promote to `src/types/` or `src/constants/` when shared across features. Per spec/17/31 line 32 and rule 13 for React prop types."
          doThis="Create `types.ts` and `constants.ts` the moment you write the first non-trivial type or literal. Import types with `import type` so bundlers can drop them. When you catch an inline `type Props = {...}` on a component signature during review, request extraction before approving. No 'we will extract it later.'"
        />
      </div>
    </SlideLayout>
  );
}
