import { motion } from "framer-motion";
import { Quote } from "lucide-react";

export interface CalloutQuoteProps {
  /** The aphorism / quote body. Keep it short: 1-2 lines at most. */
  quote: string;
  /** Optional attribution (author, source, or context). */
  attribution?: string;
  /** Visual accent. Defaults to "primary". */
  accent?: "primary" | "accent" | "destructive";
  /** Optional entrance delay, in seconds. */
  delay?: number;
}

const EASE = [0.22, 1, 0.36, 1] as const;

/**
 * CalloutQuote renders an aphorism panel used across guideline slides
 * (e.g. "comments lie, code does not"). Symptom -> Rule -> Action slides
 * use this to reinforce the rule with a memorable one-liner.
 */
export function CalloutQuote({
  quote,
  attribution,
  accent = "primary",
  delay = 0,
}: CalloutQuoteProps) {
  const accentColor = `hsl(var(--${accent}))`;
  return (
    <motion.figure
      initial={{ opacity: 0, y: 12 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.5, delay, ease: EASE }}
      style={{
        margin: 0,
        padding: "28px 32px",
        borderRadius: 14,
        background: "hsl(var(--bg-raised))",
        border: `1px solid ${accentColor}`,
        borderLeft: `6px solid ${accentColor}`,
        display: "flex",
        gap: 20,
        alignItems: "flex-start",
      }}
    >
      <Quote
        size={36}
        strokeWidth={2.5}
        style={{ color: accentColor, flexShrink: 0, marginTop: 4 }}
        aria-hidden
      />
      <div style={{ display: "flex", flexDirection: "column", gap: 10 }}>
        <blockquote
          style={{
            margin: 0,
            fontSize: 34,
            lineHeight: 1.3,
            fontWeight: 600,
            color: "hsl(var(--fg))",
            fontStyle: "italic",
          }}
        >
          &ldquo;{quote}&rdquo;
        </blockquote>
        {attribution && (
          <figcaption
            style={{
              fontSize: 20,
              letterSpacing: "0.08em",
              textTransform: "uppercase",
              color: accentColor,
              fontWeight: 700,
            }}
          >
            {attribution}
          </figcaption>
        )}
      </div>
    </motion.figure>
  );
}
