import { motion } from "framer-motion";
import { AlertTriangle, Compass, CheckCircle2 } from "lucide-react";
import type { ReactNode } from "react";

export interface ActionPanelProps {
  symptom: string;
  rule: string;
  doThis: string;
}

const EASE = [0.22, 1, 0.36, 1] as const;

const cardEntrance = {
  hidden: { opacity: 0, y: 16 },
  show: (delay: number) => ({
    opacity: 1,
    y: 0,
    transition: { duration: 0.45, delay, ease: EASE },
  }),
};

interface CardProps {
  icon: ReactNode;
  kicker: string;
  body: string;
  accent: "destructive" | "primary" | "accent";
  delay: number;
}

function Card({ icon, kicker, body, accent, delay }: CardProps) {
  return (
    <motion.div
      variants={cardEntrance}
      initial="hidden"
      animate="show"
      custom={delay}
      style={{
        background: "hsl(var(--bg-raised))",
        border: `1px solid hsl(var(--${accent}) / 0.45)`,
        borderTop: `4px solid hsl(var(--${accent}))`,
        borderRadius: 14,
        padding: "20px 24px",
        display: "flex",
        flexDirection: "column",
        gap: 10,
        minHeight: 0,
      }}
    >
      <div
        style={{
          display: "flex",
          alignItems: "center",
          gap: 12,
          color: `hsl(var(--${accent}))`,
          fontSize: 22,
          fontWeight: 700,
          letterSpacing: "0.06em",
          textTransform: "uppercase",
        }}
      >
        {icon}
        <span>{kicker}</span>
      </div>
      <div style={{ fontSize: 26, lineHeight: 1.35, color: "hsl(var(--fg))" }}>
        {body}
      </div>
    </motion.div>
  );
}

export function ActionPanel({ symptom, rule, doThis }: ActionPanelProps) {
  return (
    <div
      style={{
        display: "grid",
        gridTemplateColumns: "1fr 1fr 1fr",
        gap: "var(--space-3)",
        marginTop: "var(--space-3)",
      }}
    >
      <Card
        icon={<AlertTriangle size={26} strokeWidth={2.5} />}
        kicker="Symptom"
        body={symptom}
        accent="destructive"
        delay={1.3}
      />
      <Card
        icon={<Compass size={26} strokeWidth={2.5} />}
        kicker="The Rule"
        body={rule}
        accent="primary"
        delay={1.45}
      />
      <Card
        icon={<CheckCircle2 size={26} strokeWidth={2.5} />}
        kicker="Do this in your next PR"
        body={doThis}
        accent="accent"
        delay={1.6}
      />
    </div>
  );
}
