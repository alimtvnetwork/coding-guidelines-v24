import type { CSSProperties } from "react";

/**
 * Severity of a coding-guideline rule, mirrored from
 * `spec/17-consolidated-guidelines/31-compiled-simple-coding-guidelines.md`.
 *
 * - `hard`  : CI-enforced, blocks merge.
 * - `warn`  : linter warning, expected to be fixed same-PR.
 * - `style` : style/consistency preference, non-blocking.
 */
export type RuleSeverityType = "hard" | "warn" | "style";

export interface RuleBadgeProps {
  severity: RuleSeverityType;
  /** Optional rule id, e.g. "NAM-001". Rendered after the severity label. */
  ruleId?: string;
}

interface SeverityStyle {
  label: string;
  accentToken: string;
}

const SEVERITY_STYLES: Record<RuleSeverityType, SeverityStyle> = {
  hard: { label: "Hard", accentToken: "destructive" },
  warn: { label: "Warn", accentToken: "primary" },
  style: { label: "Style", accentToken: "accent" },
};

/**
 * Compact pill that surfaces rule severity in slide headers.
 * Uses `.slide-badge` semantics from tokens.css so it stays 20px chrome
 * (nowrap, padded proportionally) instead of scaling with body text.
 */
export function RuleBadge({ severity, ruleId }: RuleBadgeProps) {
  const style = SEVERITY_STYLES[severity];
  const pill: CSSProperties = {
    display: "inline-flex",
    alignItems: "center",
    gap: 8,
    padding: "6px 14px",
    borderRadius: 999,
    background: `hsl(var(--${style.accentToken}) / 0.14)`,
    border: `1px solid hsl(var(--${style.accentToken}) / 0.55)`,
    color: `hsl(var(--${style.accentToken}))`,
    fontWeight: 600,
  };
  const dot: CSSProperties = {
    width: 8,
    height: 8,
    borderRadius: 999,
    background: `hsl(var(--${style.accentToken}))`,
  };

  return (
    <span className="slide-badge" style={pill} aria-label={`Severity: ${style.label}`}>
      <span style={dot} aria-hidden="true" />
      <span>{style.label}</span>
      {ruleId ? <span style={{ opacity: 0.75 }}>· {ruleId}</span> : null}
    </span>
  );
}
