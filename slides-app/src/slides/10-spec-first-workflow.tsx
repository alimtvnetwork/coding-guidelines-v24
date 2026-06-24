import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { motion } from "framer-motion";
import { ArrowRight } from "lucide-react";

const STEPS = [
  {
    label: "1. Spec",
    detail: "spec/<area>/NN-name.md",
    note: "Write what changes & why",
    color: "primary",
  },
  {
    label: "2. Issue",
    detail: "03-issues/NN-summary.md",
    note: "Track the work item",
    color: "accent",
  },
  {
    label: "3. Code + PR",
    detail: "src/… + tests + lint clean",
    note: "Implement against the spec",
    color: "destructive",
  },
];

export default function SpecFirstWorkflowSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 10 · Process"
      title="Write the spec before you touch the code"
      subtitle="No code change without a spec entry or issue. Humans and AI agents read the same source of truth."
    >
      <div
        style={{
          display: "flex",
          alignItems: "stretch",
          justifyContent: "center",
          gap: "var(--space-3)",
          paddingTop: "var(--space-2)",
        }}
      >
        {STEPS.map((step, i) => (
          <div
            key={step.label}
            style={{ display: "flex", alignItems: "stretch", gap: "var(--space-3)" }}
          >
            <motion.div
              initial={{ opacity: 0, scale: 0.9 }}
              animate={{ opacity: 1, scale: 1 }}
              transition={{
                duration: 0.5,
                delay: 0.2 + i * 0.3,
                ease: [0.22, 1, 0.36, 1],
              }}
              style={{
                background: "hsl(var(--bg-raised))",
                border: `2px solid hsl(var(--${step.color}) / 0.5)`,
                borderRadius: 20,
                padding: "var(--space-3) var(--space-4)",
                minWidth: 320,
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                justifyContent: "center",
                textAlign: "center",
                gap: 6,
              }}
            >
              <div
                style={{
                  fontSize: 44,
                  fontWeight: 700,
                  color: `hsl(var(--${step.color}))`,
                }}
              >
                {step.label}
              </div>
              <div
                style={{
                  fontSize: 20,
                  color: "hsl(var(--muted-fg))",
                  fontFamily: "var(--font-mono)",
                }}
              >
                {step.detail}
              </div>
              <div style={{ fontSize: 20, color: "hsl(var(--fg))" }}>{step.note}</div>
            </motion.div>
            {i < STEPS.length - 1 && (
              <motion.div
                initial={{ opacity: 0, x: -20 }}
                animate={{ opacity: 1, x: 0 }}
                transition={{ duration: 0.4, delay: 0.4 + i * 0.3 }}
                style={{
                  display: "flex",
                  alignItems: "center",
                  color: "hsl(var(--muted-fg))",
                }}
              >
                <ArrowRight size={40} />
              </motion.div>
            )}
          </div>
        ))}
      </div>
      <ActionPanel
        symptom="Code lands without a written commitment. Reviewers and AI agents can't tell what 'done' means."
        rule="Every change starts with a spec file or an issue. The diff implements it, the PR links to it."
        doThis="Before your next commit, write a 5-line spec/<area>/NN-name.md describing what & why. Link it in the PR."
      />
    </SlideLayout>
  );
}
