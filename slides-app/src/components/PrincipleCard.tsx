import { motion } from "framer-motion";
import { AlertTriangle, Compass, CheckCircle2 } from "lucide-react";

export interface PrincipleCardProps {
  number: string;
  title: string;
  symptom: string;
  rule: string;
  action: string;
  accent: "destructive" | "primary" | "accent";
  delay: number;
}

const EASE = [0.22, 1, 0.36, 1] as const;

const entrance = {
  hidden: { opacity: 0, y: 18 },
  show: (delay: number) => ({
    opacity: 1,
    y: 0,
    transition: { duration: 0.45, delay, ease: EASE },
  }),
};

interface RowProps {
  icon: React.ReactNode;
  label: string;
  body: string;
  color: string;
}

function Row({ icon, label, body, color }: RowProps) {
  return (
    <div style={{ display: "flex", gap: 12, alignItems: "flex-start" }}>
      <div style={{ color, marginTop: 2, flexShrink: 0 }}>{icon}</div>
      <div style={{ display: "flex", flexDirection: "column", gap: 2 }}>
        <div
          style={{
            fontSize: 16,
            fontWeight: 700,
            letterSpacing: "0.1em",
            textTransform: "uppercase",
            color,
          }}
        >
          {label}
        </div>
        <div style={{ fontSize: 22, lineHeight: 1.32, color: "hsl(var(--fg))" }}>
          {body}
        </div>
      </div>
    </div>
  );
}

export function PrincipleCard({
  number,
  title,
  symptom,
  rule,
  action,
  accent,
  delay,
}: PrincipleCardProps) {
  const color = `hsl(var(--${accent}))`;
  return (
    <motion.div
      variants={entrance}
      initial="hidden"
      animate="show"
      custom={delay}
      style={{
        background: "hsl(var(--bg-raised))",
        border: "1px solid hsl(var(--border))",
        borderLeft: `6px solid ${color}`,
        borderRadius: 16,
        padding: "28px 32px",
        display: "flex",
        flexDirection: "column",
        gap: 18,
        minHeight: 0,
      }}
    >
      <div style={{ display: "flex", alignItems: "baseline", gap: 16 }}>
        <div style={{ fontSize: 30, color, fontWeight: 700, letterSpacing: "0.05em" }}>
          {number}
        </div>
        <div
          style={{
            fontSize: 38,
            fontWeight: 700,
            lineHeight: 1.05,
            fontFamily: "Ubuntu, sans-serif",
          }}
        >
          {title}
        </div>
      </div>
      <div style={{ display: "flex", flexDirection: "column", gap: 14 }}>
        <Row
          icon={<AlertTriangle size={22} strokeWidth={2.5} />}
          label="Symptom"
          body={symptom}
          color="hsl(var(--destructive))"
        />
        <Row
          icon={<Compass size={22} strokeWidth={2.5} />}
          label="Rule"
          body={rule}
          color="hsl(var(--primary))"
        />
        <Row
          icon={<CheckCircle2 size={22} strokeWidth={2.5} />}
          label="Do this next"
          body={action}
          color="hsl(var(--accent))"
        />
      </div>
    </motion.div>
  );
}
