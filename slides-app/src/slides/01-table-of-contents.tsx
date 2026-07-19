import { SlideLayout } from "@/components/SlideLayout";
import { DECK } from "@/deck";
import { PROGRESS_BLOCKS, useProgressSnapshot } from "@/lib/progress";

/**
 * Table of contents. Deep-links every slide via hash routing (`#/N`).
 * Also surfaces per-slide review progress: three dots (Symptom/Rule/Action)
 * next to every guideline entry, plus an overall completion bar.
 */

const TITLE_ID = "00-title";
const CLOSING_ID = "12-closing";
const SELF_ID = "01-toc";

// Slides whose Symptom/Rule/Action state we track. Kept in sync with
// each rule slide's `slideId` prop passed to ActionPanel.
const TRACKED_SLIDE_IDS: readonly string[] = [
  "01-naming-conventions",
  "02-nested-if-else",
  "03-boolean-prefixes",
  "04-app-error-wrapper",
  "05-structured-logging",
  "06-magic-strings",
  "07-function-and-file-metrics",
  "08-two-operand-max",
  "09-positively-named-guards",
  "10-spec-first-workflow",
  "11-cache-invalidation",
] as const;

function ProgressDots({ entryId }: { entryId: string }) {
  const all = useProgressSnapshot();
  const entry = all[entryId] ?? {};
  return (
    <span style={{ display: "inline-flex", gap: 6 }} aria-label="review progress">
      {PROGRESS_BLOCKS.map((b) => {
        const on = Boolean(entry[b]);
        return (
          <span
            key={b}
            title={`${b} reviewed: ${on ? "yes" : "no"}`}
            style={{
              width: 12,
              height: 12,
              borderRadius: 999,
              background: on ? "hsl(var(--slide-accent))" : "transparent",
              border: "2px solid hsl(var(--slide-accent))",
              opacity: on ? 1 : 0.5,
            }}
          />
        );
      })}
    </span>
  );
}

function OverallProgress() {
  const all = useProgressSnapshot();
  const total = TRACKED_SLIDE_IDS.length * PROGRESS_BLOCKS.length;
  const done = TRACKED_SLIDE_IDS.reduce((acc, id) => {
    const entry = all[id] ?? {};
    return acc + PROGRESS_BLOCKS.filter((b) => entry[b]).length;
  }, 0);
  const pct = total === 0 ? 0 : Math.round((done / total) * 100);
  return (
    <div style={{ display: "flex", alignItems: "center", gap: 16, marginTop: 12 }}>
      <div
        style={{
          flex: 1,
          height: 10,
          borderRadius: 999,
          background: "rgba(148, 163, 184, 0.25)",
          overflow: "hidden",
        }}
        role="progressbar"
        aria-valuemin={0}
        aria-valuemax={100}
        aria-valuenow={pct}
        aria-label="Overall review progress"
      >
        <div
          style={{
            width: `${pct}%`,
            height: "100%",
            background: "hsl(var(--slide-accent))",
            transition: "width 200ms ease",
          }}
        />
      </div>
      <span
        className="slide-chrome"
        style={{ fontVariantNumeric: "tabular-nums", fontWeight: 600, minWidth: 140, textAlign: "right" }}
      >
        {done} / {total} blocks · {pct}%
      </span>
    </div>
  );
}

export default function TableOfContentsSlide() {
  const entries = DECK.map((slide, index) => ({ ...slide, index }));
  const title = entries.find((e) => e.id === TITLE_ID);
  const closing = entries.find((e) => e.id === CLOSING_ID);
  const body = entries.filter(
    (e) => e.id !== SELF_ID && e.id !== TITLE_ID && e.id !== CLOSING_ID,
  );

  return (
    <SlideLayout
      eyebrow="Contents"
      title="Jump to any guideline"
      subtitle="Tick Symptom, Rule and Action on each slide as you review. Progress persists in this browser."
    >
      <OverallProgress />

      <div
        style={{
          display: "grid",
          gridTemplateColumns: "1fr 1fr",
          gap: "14px 40px",
          marginTop: 20,
        }}
      >
        {body.map((entry) => (
          <a
            key={entry.id}
            href={`#/id/${encodeURIComponent(entry.id)}`}
            className="toc-link"
            style={{
              display: "flex",
              alignItems: "center",
              gap: 20,
              padding: "12px 20px",
              borderRadius: 12,
              textDecoration: "none",
              color: "inherit",
              border: "1px solid rgba(148, 163, 184, 0.25)",
              background: "rgba(15, 23, 42, 0.35)",
              transition: "background 120ms ease, border-color 120ms ease",
            }}
          >
            <span
              className="slide-chrome"
              style={{
                minWidth: 56,
                fontVariantNumeric: "tabular-nums",
                color: "hsl(var(--slide-accent))",
                fontWeight: 600,
              }}
            >
              {String(entry.index).padStart(2, "0")}
            </span>
            <span className="slide-body" style={{ flex: 1 }}>
              {entry.title}
            </span>
            <ProgressDots entryId={entry.id} />
          </a>
        ))}
      </div>

      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
          marginTop: 24,
        }}
      >
        {title ? (
          <a href={`#/${title.index}`} className="slide-chrome toc-chrome-link">
            ← {String(title.index).padStart(2, "0")} · {title.title}
          </a>
        ) : (
          <span />
        )}
        {closing ? (
          <a href={`#/${closing.index}`} className="slide-chrome toc-chrome-link">
            {String(closing.index).padStart(2, "0")} · {closing.title} →
          </a>
        ) : (
          <span />
        )}
      </div>
    </SlideLayout>
  );
}
