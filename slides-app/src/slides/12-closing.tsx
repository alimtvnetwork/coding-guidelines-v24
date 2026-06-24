import { SlideLayout } from "@/components/SlideLayout";
import { motion } from "framer-motion";

const CHECKLIST = [
  { n: "01", text: "Rename one snake_case symbol to PascalCase" },
  { n: "02", text: "Flatten one nested if into early-return guards" },
  { n: "03", text: "Prefix one boolean with Is / Has / Can / Should" },
  { n: "04", text: "Wrap one catch with AppError + Ids" },
  { n: "05", text: "Replace one console.log with a structured logger call" },
  { n: "06", text: "Promote one repeated string literal to an enum" },
  { n: "07", text: "Split one >300-line file into focused modules" },
  { n: "08", text: "Extract one long && chain into a Can…/Is… helper" },
  { n: "09", text: "Rename one negative predicate to its positive form" },
  { n: "10", text: "Write a 5-line spec before your next commit" },
  { n: "11", text: "Add TTL + invalidation to one existing cache" },
];

const EASE = [0.22, 1, 0.36, 1] as const;

export default function ClosingSlide() {
  return (
    <SlideLayout
      eyebrow="Recap · Your next PR"
      title="Pick one. Ship it today."
      subtitle="Each item is a single-file, single-PR change. Don't try to do all 11 at once — pick the one that helps your current code most."
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
            transition={{ duration: 0.4, delay: 0.2 + i * 0.05, ease: EASE }}
            style={{
              display: "flex",
              alignItems: "center",
              gap: 16,
              padding: "12px 16px",
              background: "hsl(var(--bg-raised))",
              border: "1px solid hsl(var(--border))",
              borderRadius: 12,
              fontSize: 24,
            }}
          >
            <span
              style={{
                width: 32,
                height: 32,
                borderRadius: 8,
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
                minWidth: 36,
              }}
            >
              {item.n}
            </span>
            <span>{item.text}</span>
          </motion.label>
        ))}
      </div>
    </SlideLayout>
  );
}
