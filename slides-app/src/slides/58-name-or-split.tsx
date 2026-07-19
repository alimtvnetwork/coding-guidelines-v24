import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 60: name-or-split. If you cannot name the type, you do not
 * understand it yet. Split until you can.
 *
 * Source: spec/17/31 line 112.
 */

const BEFORE = `// One catch-all shape. Nothing is grep-able. "Data" tells you nothing.
export type Data = {
  id: number;
  name: string;
  email: string;
  avatarUrl: string;
  billingAddress: string;
  shippingAddress: string;
  cardLast4: string;
  planId: number;
  planStartedAt: string;
  planEndsAt: string;
  invoices: Array<{ id: number; total: number; paidAt: string | null }>;
  auditEvents: Array<{ at: string; kind: string; actor: string }>;
};

export function useData(userId: number): { data: Data; isLoading: boolean } { ... }`;

const AFTER = `// Split by aggregate. Every type has a name a domain expert would recognise.
export type UserProfile = {
  UserId: number;
  Name: string;
  Email: string;
  AvatarUrl: string;
};

export type BillingSnapshot = {
  BillingAddress: string;
  ShippingAddress: string;
  CardLast4: string;
};

export type SubscriptionSnapshot = {
  PlanId: number;
  StartedAt: string;
  EndsAt: string;
};

export type InvoiceSummary = {
  InvoiceId: number;
  Total: number;
  PaidAt: string | null;
};

export type AuditEvent = {
  At: string;
  Kind: AuditEventKind;
  Actor: string;
};

export type AccountOverview = {
  profile: UserProfile;
  billing: BillingSnapshot;
  subscription: SubscriptionSnapshot;
  invoices: readonly InvoiceSummary[];
  auditEvents: readonly AuditEvent[];
};

export function useAccountOverview(userId: UserId): AccountOverviewQuery { ... }`;

export default function NameOrSplitSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 60 · TypeScript · name it or split it, no `Data`, `Info`, `Config`"
      title="If you cannot invent a domain name for a type, you do not understand it yet. Split until you can."
      subtitle="spec/17/31 line 112: as the author (human or AI), invent the clearest domain name for each type. Naming failure is a comprehension failure. Splitting is the fix, never `Data`, `Info`, `Config`, `Payload`, `Params`, `Options`, or `Thing`."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 16, marginTop: 4 }}>
        <CodeDiff
          language="tsx"
          before={BEFORE}
          after={AFTER}
          beforeLabel="Bad. `Data` is the tell: the author mashed five aggregates (profile, billing, subscription, invoices, audit trail) into one bag because naming any of them properly required understanding the domain, and the shortcut was cheaper. Every consumer now destructures nine unrelated fields out of one object; a change to billing forces a rebuild of every audit view; nothing can be reused; nothing can be tested in isolation; hover shows a 14-field blob."
          afterLabel="Good. Five named aggregates, each with a domain meaning a product manager would recognise. `AccountOverview` composes them by name, so the top-level shape reads as a table of contents. Each aggregate can move to its own module, be fetched independently, cached independently, mocked independently. Renaming a billing field touches billing only. Hover on any variable shows the aggregate name first, fields second."
        />
        <ActionPanel
          slideId="58-name-or-split"
          symptom="A single `Data` type grew to 47 fields across two years. Three failure modes traced to the same root cause: (1) a billing refactor accidentally broke the audit trail because both shared `Data.updatedAt`, and the audit view treated it as event time while billing treated it as invoice-issue time; (2) three teams built three overlapping `useData` hooks because nobody could tell what `Data` meant, and each variant fetched a different subset; (3) a new engineer asked in review chat what `Data` was; no one could answer without opening the file. Root cause: the author could not name it, did not split, and shipped `Data`."
          rule="spec/17/31 line 112 is absolute. Every type gets a domain name a non-engineer stakeholder would recognise. Banned placeholder names: `Data`, `Info`, `Config`, `Options`, `Params`, `Payload`, `Response`, `Request`, `Result`, `Thing`, `Item` (bare), `Value`, `Entry`, `Record`, `Object`, `Details`, `Meta`, `State` (bare), `Props` (bare), `Args`. Any of these on their own is a review-blocking violation. Allowed only when qualified and specific: `LoginRequest`, `InvoiceListResponse`, `RetryOptions`, `PaymentPayload`. If you cannot invent a specific qualified name, you do not yet understand the concept: split the type along its natural seams (aggregate boundaries, transaction boundaries, ownership boundaries) until each piece has a name that survives being read aloud to a product manager. Pairs with NAM-001 (PascalCase domain nouns), REACT-010 (name every composite), REACT-011 (types.ts colocation) as one system."
          doThis="Enforce mechanically: (1) ESLint rule `no-placeholder-type-names` with the banned list above (bare form only), autofix suggests `git blame` the definition and asks the author to rename; (2) code review checklist item `type-names-are-domain-nouns`, PRs failing this label are blocked; (3) design-review reflex: for any new type larger than 5 fields, ask 'what would a product manager call this?' If the answer is a shrug, the type must be split before merge; (4) refactor recipe: identify the aggregates (things with their own lifecycle, ownership, or transactional boundary), extract each into its own named type in `src/types/{domain}.ts`, compose them at the call site; (5) when genuinely uncertain, mark the type `// TODO(NAM-002): rename before ship` and open a ticket; do NOT merge with `Data` in the type tree. Rule of thumb: if you would not put the name in a stakeholder email, it does not belong in the type system."
        />
      </div>
    </SlideLayout>
  );
}
