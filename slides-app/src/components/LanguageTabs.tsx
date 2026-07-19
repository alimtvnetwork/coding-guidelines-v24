import { useState, type ReactNode } from "react";
import { motion, AnimatePresence } from "framer-motion";

export type LanguageId = "go" | "ts" | "php" | "rust" | (string & {});

export interface LanguageTab {
  /** Stable identifier used for selection state and deep-link hooks. */
  id: LanguageId;
  /** Short display label shown on the tab pill (e.g. "Go", "TypeScript"). */
  label: string;
  /** Rendered content for this tab (usually a `CodeDiff` or `<pre>` block). */
  content: ReactNode;
}

export interface LanguageTabsProps {
  tabs: LanguageTab[];
  /** Which tab id is active by default. Defaults to the first tab. */
  defaultTabId?: LanguageId;
  /** Optional label announced to screen readers (defaults to "Language"). */
  ariaLabel?: string;
}

const EASE = [0.22, 1, 0.36, 1] as const;

/**
 * LanguageTabs lets a single rule slide show equivalent snippets across
 * languages (Go / TS / PHP / Rust) without duplicating slides. It is a
 * content-agnostic tab shell: callers pass any ReactNode per tab (typically
 * a `CodeDiff` or a static `<pre>` block).
 *
 * Highlighter bundling: the shared `CodeDiff` currently ships only the
 * TypeScript grammar to respect the 8 MB offline-contract ceiling. Adding
 * more grammars is tracked separately; LanguageTabs itself is grammar-free.
 */
export function LanguageTabs({
  tabs,
  defaultTabId,
  ariaLabel = "Language",
}: LanguageTabsProps) {
  const initialId = defaultTabId ?? tabs[0]?.id;
  const [activeId, setActiveId] = useState<LanguageId | undefined>(initialId);
  const active = tabs.find((tab) => tab.id === activeId) ?? tabs[0];

  return (
    <div style={{ display: "flex", flexDirection: "column", gap: 16 }}>
      <div
        role="tablist"
        aria-label={ariaLabel}
        style={{
          display: "flex",
          gap: 10,
          padding: 6,
          background: "hsl(var(--bg-raised))",
          border: "1px solid hsl(var(--border))",
          borderRadius: 12,
          alignSelf: "flex-start",
        }}
      >
        {tabs.map((tab) => {
          const isActive = tab.id === active?.id;
          return (
            <button
              key={tab.id}
              type="button"
              role="tab"
              aria-selected={isActive}
              aria-controls={`lang-panel-${tab.id}`}
              id={`lang-tab-${tab.id}`}
              onClick={() => setActiveId(tab.id)}
              style={{
                position: "relative",
                padding: "10px 22px",
                borderRadius: 8,
                border: "none",
                cursor: "pointer",
                background: "transparent",
                color: isActive ? "hsl(var(--primary-fg))" : "hsl(var(--fg-muted))",
                fontSize: 22,
                fontWeight: 700,
                letterSpacing: "0.04em",
              }}
            >
              {isActive && (
                <motion.span
                  layoutId="lang-tab-pill"
                  transition={{ duration: 0.25, ease: EASE }}
                  style={{
                    position: "absolute",
                    inset: 0,
                    background: "hsl(var(--primary))",
                    borderRadius: 8,
                    zIndex: 0,
                  }}
                />
              )}
              <span style={{ position: "relative", zIndex: 1 }}>{tab.label}</span>
            </button>
          );
        })}
      </div>

      <AnimatePresence mode="wait">
        <motion.div
          key={active?.id}
          role="tabpanel"
          id={`lang-panel-${active?.id}`}
          aria-labelledby={`lang-tab-${active?.id}`}
          initial={{ opacity: 0, y: 8 }}
          animate={{ opacity: 1, y: 0 }}
          exit={{ opacity: 0, y: -8 }}
          transition={{ duration: 0.22, ease: EASE }}
        >
          {active?.content}
        </motion.div>
      </AnimatePresence>
    </div>
  );
}
