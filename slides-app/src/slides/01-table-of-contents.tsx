import { SlideLayout } from "@/components/SlideLayout";
import { DECK } from "@/deck";

/**
 * Table of contents. Deep-links every slide via hash routing (`#/N`).
 * Excludes itself and the title from the numbered rule list, but keeps
 * them clickable as small chrome entries at the top and bottom.
 */
export default function TableOfContentsSlide() {
  const selfId = "01-toc";
  const entries = DECK.map((slide, index) => ({ ...slide, index }));
  const title = entries.find((e) => e.id === "00-title");
  const closing = entries.find((e) => e.id === "12-closing");
  const body = entries.filter(
    (e) => e.id !== selfId && e.id !== "00-title" && e.id !== "12-closing",
  );

  return (
    <SlideLayout
      eyebrow="Contents"
      title="Jump to any guideline"
      subtitle="Every entry deep-links to the slide. Use during review as a lookup index; keys and clicks both work."
    >
      <div
        style={{
          display: "grid",
          gridTemplateColumns: "1fr 1fr",
          gap: "16px 40px",
          marginTop: 24,
        }}
      >
        {body.map((entry) => (
          <a
            key={entry.id}
            href={`#/${entry.index}`}
            className="toc-link"
            style={{
              display: "flex",
              alignItems: "baseline",
              gap: 20,
              padding: "14px 20px",
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
            <span className="slide-body">{entry.title}</span>
          </a>
        ))}
      </div>

      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
          marginTop: 32,
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
