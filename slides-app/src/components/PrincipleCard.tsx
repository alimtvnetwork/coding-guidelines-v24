import { motion } from "framer-motion";
import { AlertTriangle, Compass, CheckCircle2 } from "lucide-react";
import { ReviewCheckbox } from "@/components/ReviewCheckbox";
import type { ProgressBlock } from "@/lib/progress";

export interface PrincipleCardProps {
  number: string;
  title: string;
  symptom: string;
  rule: string;
  action: string;
  accent: "destructive" | "primary" | "accent";
  delay: number;
  /** Enables per-block review checkboxes when set. */
  progressId?: string;
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
  progressId?: string;
  block?: ProgressBlock;
}

function Row({ icon, label, body, color, progressId, block }: RowProps) {
  return (
    <div style={{ display: "flex", gap: 12, alignItems: "flex-start" }}>
      <div style={{ color, marginTop: 2, flexShrink: 0 }}>{icon}</div>
      <div style={{ display: "flex", flexDirection: "column", gap: 6, flex: 1 }}>
        <div
          style={{
            display: "flex",
            alignItems: "center",
            justifyContent: "space-between",
            gap: 12,
          }}
        >
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
          {progressId && block && (
            <ReviewCheckbox
              slideId={progressId}
              block={block}
              color={color}
              label=""
            />
          )}
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
  progressId,
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
          progressId={progressId}
          block="symptom"
        />
        <Row
          icon={<Compass size={22} strokeWidth={2.5} />}
          label="Rule"
          body={rule}
          color="hsl(var(--primary))"
          progressId={progressId}
          block="rule"
        />
        <Row
          icon={<CheckCircle2 size={22} strokeWidth={2.5} />}
          label="Do this next"
          body={action}
          color="hsl(var(--accent))"
          progressId={progressId}
          block="action"
        />
      </div>
    </motion.div>
  );
}
