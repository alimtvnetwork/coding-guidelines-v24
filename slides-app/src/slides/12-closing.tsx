import { SlideLayout } from "@/components/SlideLayout";
import { motion } from "framer-motion";

/**
 * Closing recap. Rule-id keyed one-PR actions, one per section of the
 * completed SS-02 catalogue (67 slides across 10 sections). Each item is a
 * single-file, single-PR change; pick the one that helps your current code
 * most, ship today.
 */

const CHECKLIST = [
  { n: "01", rule: "MUST-002", text: "Write the root cause in one sentence before your next fix" },
  { n: "02", rule: "NAM-001", text: "Rename one snake_case symbol to PascalCase" },
  { n: "03", rule: "BOOL-002", text: "Prefix one boolean with Is / Has / Can / Should" },
  { n: "04", rule: "SIZE-001", text: "Split one file above the tier cap (100 tsx, 300 fallback) into focused modules" },
  { n: "05", rule: "CF-001", text: "Flatten one nested if into early-return guards" },
  { n: "06", rule: "ERR-001", text: "Wrap one bare catch with AppError plus context Ids" },
  { n: "07", rule: "LOG-002", text: "Add op, requestId, and key inputs to one log line; strip any secrets" },
  { n: "08", rule: "SCHEMA-002", text: "Convert one UUID PK to {TableName}Id INTEGER AUTOINCREMENT" },
  { n: "09", rule: "REACT-002", text: "Delete one derive-state useEffect; replace with useMemo or a handler" },
  { n: "10", rule: "REACT-011", text: "Move one component's prop types into a sibling types.ts" },
  { n: "11", rule: "A11Y-001", text: "Fix one keyboard trap or missing focus ring flagged by axe" },
  { n: "12", rule: "OPS-002", text: "Add SLO tie, dashboard link, and runbook link to one alert rule" },
];

const EASE = [0.22, 1, 0.36, 1] as const;

export default function ClosingSlide() {
  return (
    <SlideLayout
      eyebrow="Recap · 67 slides, 10 sections · your next PR"
      title="Pick one. Ship it today."
      subtitle="Each item is a single-file, single-PR change keyed to a rule id from the SS-02 catalogue. Don't try to do all 12 at once, pick the one that helps your current code most."
      footer={
        <>
          <span>github.com/alimtvnetwork/coding-guidelines-v24</span>
          <span>Md. Alim Ul Karim · Riseup Asia LLC</span>
        </>
      }

    >
      <div
        style={{
          display: "grid",
          gridTemplateColumns: "1fr 1fr",
          gap: "var(--space-2) var(--space-4)",
          marginTop: "var(--space-2)",
        }}
      >
        {CHECKLIST.map((item, i) => (
          <motion.label
            key={item.n}
            initial={{ opacity: 0, x: -12 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ duration: 0.4, delay: 0.15 + i * 0.04, ease: EASE }}
            style={{
              display: "flex",
              alignItems: "center",
              gap: 14,
              padding: "10px 14px",
              background: "hsl(var(--bg-raised))",
              border: "1px solid hsl(var(--border))",
              borderRadius: 12,
              fontSize: 20,
            }}
          >
            <span
              style={{
                width: 26,
                height: 26,
                borderRadius: 7,
                border: "2px solid hsl(var(--accent))",
                flexShrink: 0,
              }}
              aria-hidden
            />
            <span
              style={{
                fontFamily: "var(--font-mono)",
                color: "hsl(var(--accent))",
                fontWeight: 700,
                minWidth: 32,
              }}
            >
              {item.n}
            </span>
            <span
              style={{
                fontFamily: "var(--font-mono)",
                fontSize: 12,
                padding: "2px 8px",
                borderRadius: 999,
                background: "hsl(var(--bg))",
                border: "1px solid hsl(var(--border))",
                color: "hsl(var(--accent))",
                whiteSpace: "nowrap",
              }}
            >
              {item.rule}
            </span>
            <span style={{ flex: 1 }}>{item.text}</span>
          </motion.label>
        ))}
      </div>
    </SlideLayout>
  );
}
