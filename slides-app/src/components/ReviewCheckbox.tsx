import { Check } from "lucide-react";
import {
  getBlock,
  setBlock,
  useProgressSnapshot,
  type ProgressBlock,
} from "@/lib/progress";

export interface ReviewCheckboxProps {
  slideId: string;
  block: ProgressBlock;
  color: string;
  label?: string;
}

/**
 * Compact "reviewed" tickbox rendered inside each Symptom/Rule/Action card.
 * Persists via `lib/progress` so the TOC and other slides stay in sync.
 */
export function ReviewCheckbox({
  slideId,
  block,
  color,
  label = "Reviewed",
}: ReviewCheckboxProps) {
  // Subscribe so this checkbox re-renders when toggled elsewhere.
  useProgressSnapshot();
  const checked = getBlock(slideId, block);
  const onClick = () => setBlock(slideId, block, !checked);
  return (
    <button
      type="button"
      onClick={onClick}
      aria-pressed={checked}
      aria-label={`Mark ${block} as reviewed`}
      className="slide-chrome"
      style={{
        display: "inline-flex",
        alignItems: "center",
        gap: 8,
        padding: "6px 12px",
        borderRadius: 999,
        border: `1px solid ${color}`,
        background: checked ? color : "transparent",
        color: checked ? "hsl(var(--bg))" : color,
        cursor: "pointer",
        fontWeight: 600,
        transition: "background 120ms ease, color 120ms ease",
      }}
    >
      <span
        aria-hidden
        style={{
          width: 18,
          height: 18,
          borderRadius: 4,
          border: `2px solid ${checked ? "hsl(var(--bg))" : color}`,
          background: checked ? "hsl(var(--bg))" : "transparent",
          display: "inline-flex",
          alignItems: "center",
          justifyContent: "center",
          color,
        }}
      >
        {checked && <Check size={14} strokeWidth={3} />}
      </span>
      {label}
    </button>
  );
}
