import { useCallback, useEffect, useMemo, useRef, useState } from "react";
import { Maximize, Grid3x3, Presentation, Sun, Moon, HelpCircle, Search, X } from "lucide-react";
import { AnimatePresence, motion } from "framer-motion";
import { ScaledSlide } from "./components/ScaledSlide";
import { RuleBadge, type RuleSeverity } from "./components/RuleBadge";
import { DECK, SECTIONS, groupBySection, type SlideSection } from "./deck";

type View = "deck" | "grid" | "presenter";

/**
 * Hash routing. Two supported forms:
 *   #/id/<slide-id>   (preferred, stable across reorderings)
 *   #/<index>         (legacy numeric, still honored)
 * Unknown or malformed hashes fall back to slide 0.
 */
function readSlideFromHash(): number {
  const raw = window.location.hash;
  const idMatch = raw.match(/^#\/id\/(.+)$/);
  if (idMatch) {
    const found = DECK.findIndex((s) => s.id === decodeURIComponent(idMatch[1]));
    if (found >= 0) return found;
    console.warn(`[slides] unknown slide id in hash: ${idMatch[1]}, falling back to 0`);
    return 0;
  }
  const numMatch = raw.match(/^#\/(\d+)/);
  if (!numMatch) return 0;
  return clampSlide(parseInt(numMatch[1], 10));
}

function clampSlide(n: number): number {
  if (Number.isNaN(n) || n < 0) return 0;
  if (n >= DECK.length) return DECK.length - 1;
  return n;
}

function writeSlideToHash(n: number) {
  const slide = DECK[n];
  if (!slide) {
    window.location.hash = `/${n}`;
    return;
  }
  window.location.hash = `/id/${encodeURIComponent(slide.id)}`;
}

export default function App() {
  const [index, setIndex] = useState(readSlideFromHash);
  const [view, setView] = useState<View>("deck");
  const [helpOpen, setHelpOpen] = useState(false);
  const [paletteOpen, setPaletteOpen] = useState(false);
  const [theme, setTheme] = useState<"dark" | "light">("dark");
  const prevIndexRef = useRef(index);
  const direction = index >= prevIndexRef.current ? 1 : -1;
  useEffect(() => {
    prevIndexRef.current = index;
  }, [index]);

  useEffect(() => {
    document.documentElement.setAttribute("data-theme", theme);
  }, [theme]);

  useEffect(() => {
    function onHash() {
      setIndex(readSlideFromHash());
    }
    window.addEventListener("hashchange", onHash);
    return () => window.removeEventListener("hashchange", onHash);
  }, []);

  const goto = useCallback((n: number) => {
    const next = clampSlide(n);
    writeSlideToHash(next);
    setIndex(next);
  }, []);

  const next = useCallback(() => goto(index + 1), [index, goto]);
  const prev = useCallback(() => goto(index - 1), [index, goto]);

  const toggleFullscreen = useCallback(async () => {
    if (document.fullscreenElement) {
      await document.exitFullscreen();
    } else {
      await document.documentElement.requestFullscreen();
    }
  }, []);

  useEffect(() => {
    function onKey(e: KeyboardEvent) {
      // Cmd/Ctrl+K toggles palette regardless of focus (except when already inside palette input,
      // which handles it locally by not re-firing — global still fine since toggle is idempotent).
      if ((e.metaKey || e.ctrlKey) && (e.key === "k" || e.key === "K")) {
        e.preventDefault();
        setPaletteOpen((v) => !v);
        return;
      }
      if (e.target instanceof HTMLInputElement || e.target instanceof HTMLTextAreaElement) return;
      switch (e.key) {
        case "ArrowRight":
        case " ":
        case "PageDown":
          e.preventDefault();
          next();
          break;
        case "ArrowLeft":
        case "PageUp":
          e.preventDefault();
          prev();
          break;
        case "Home":
          e.preventDefault();
          goto(0);
          break;
        case "End":
          e.preventDefault();
          goto(DECK.length - 1);
          break;
        case "g":
        case "G":
          setView((v) => (v === "grid" ? "deck" : "grid"));
          break;
        case "p":
        case "P":
          setView((v) => (v === "presenter" ? "deck" : "presenter"));
          break;
        case "f":
        case "F":
          toggleFullscreen();
          break;
        case "?":
        case "h":
        case "H":
          e.preventDefault();
          setHelpOpen((v) => !v);
          break;
        case "Escape":
          if (paletteOpen) {
            setPaletteOpen(false);
            break;
          }
          if (helpOpen) {
            setHelpOpen(false);
            break;
          }
          if (view !== "deck") setView("deck");
          break;
      }
    }
    window.addEventListener("keydown", onKey);
    return () => window.removeEventListener("keydown", onKey);
  }, [next, prev, goto, toggleFullscreen, view, helpOpen, paletteOpen]);

  const Current = DECK[index].component;
  const NextSlide = DECK[Math.min(index + 1, DECK.length - 1)].component;

  if (view === "grid") {
    return (
      <GridView
        currentIndex={index}
        onPick={(n) => {
          goto(n);
          setView("deck");
        }}
        onClose={() => setView("deck")}
      />
    );
  }

  if (view === "presenter") {
    return (
      <PresenterView
        Current={Current}
        Next={NextSlide}
        index={index}
        total={DECK.length}
        onClose={() => setView("deck")}
      />
    );
  }

  return (
    <div className="deck-root">
      <div className="scaled-stage">
        <AnimatePresence mode="wait" custom={direction} initial={false}>
          <motion.div
            key={index}
            custom={direction}
            variants={{
              enter: (dir: number) => ({ opacity: 0, x: dir * 80, scale: 0.985 }),
              center: { opacity: 1, x: 0, scale: 1 },
              exit: (dir: number) => ({ opacity: 0, x: dir * -80, scale: 0.985 }),
            }}
            initial="enter"
            animate="center"
            exit="exit"
            transition={{ duration: 0.45, ease: [0.22, 1, 0.36, 1] }}
            style={{ width: "100%", height: "100%" }}
          >
            <ScaledSlide>
              <Current />
            </ScaledSlide>
          </motion.div>
        </AnimatePresence>
      </div>

      <div className="toolbar">
        <button onClick={() => setView("grid")} title="Grid (G)">
          <Grid3x3 size={14} /> Grid
        </button>
        <button onClick={() => setView("presenter")} title="Presenter (P)">
          <Presentation size={14} /> Presenter
        </button>
        <button onClick={toggleFullscreen} title="Fullscreen (F)">
          <Maximize size={14} /> Fullscreen
        </button>
        <button
          onClick={() => setTheme((t) => (t === "dark" ? "light" : "dark"))}
          title="Theme"
        >
          {theme === "dark" ? <Sun size={14} /> : <Moon size={14} />}
        </button>
        <button onClick={() => setHelpOpen((v) => !v)} title="Keyboard shortcuts (?)" aria-label="Keyboard shortcuts">
          <HelpCircle size={14} /> ?
        </button>
      </div>

      <div className="nav-pill">
        <button onClick={prev} aria-label="Previous">←</button>
        <span style={{ fontFamily: "var(--font-mono)" }}>
          {String(index + 1).padStart(2, "0")} / {String(DECK.length).padStart(2, "0")}
        </span>
        <button onClick={next} aria-label="Next">→</button>
      </div>

      {helpOpen ? <HelpOverlay onClose={() => setHelpOpen(false)} /> : null}
    </div>
  );
}

const SHORTCUTS: readonly { keys: string; label: string }[] = [
  { keys: "→ / Space / PgDn", label: "Next slide" },
  { keys: "← / PgUp", label: "Previous slide" },
  { keys: "Home / End", label: "First / last slide" },
  { keys: "G", label: "Toggle grid overview" },
  { keys: "P", label: "Toggle presenter view" },
  { keys: "F", label: "Toggle fullscreen" },
  { keys: "? / H", label: "Toggle this help" },
  { keys: "Esc", label: "Close overlay or return to deck" },
];

function HelpOverlay({ onClose }: { onClose: () => void }) {
  return (
    <div
      role="dialog"
      aria-modal="true"
      aria-label="Keyboard shortcuts"
      onClick={onClose}
      style={{
        position: "fixed",
        inset: 0,
        background: "rgba(0, 0, 0, 0.65)",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        zIndex: 1000,
      }}
    >
      <div
        onClick={(e) => e.stopPropagation()}
        style={{
          background: "hsl(var(--slide-bg, 222 47% 11%))",
          color: "hsl(var(--slide-fg, 210 40% 98%))",
          border: "1px solid rgba(148, 163, 184, 0.35)",
          borderRadius: 16,
          padding: "28px 32px",
          minWidth: 460,
          maxWidth: 560,
          boxShadow: "0 20px 60px rgba(0, 0, 0, 0.5)",
        }}
      >
        <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center", marginBottom: 16 }}>
          <h2 style={{ margin: 0, fontSize: 20, letterSpacing: "-0.01em" }}>Keyboard shortcuts</h2>
          <button
            onClick={onClose}
            aria-label="Close"
            style={{ background: "transparent", border: 0, color: "inherit", cursor: "pointer", padding: 4 }}
          >
            <X size={18} />
          </button>
        </div>
        <ul style={{ listStyle: "none", padding: 0, margin: 0, display: "grid", gap: 10 }}>
          {SHORTCUTS.map((s) => (
            <li
              key={s.keys}
              style={{ display: "flex", justifyContent: "space-between", alignItems: "center", gap: 24 }}
            >
              <span style={{ fontSize: 14, opacity: 0.85 }}>{s.label}</span>
              <kbd
                style={{
                  fontFamily: "var(--font-mono)",
                  fontSize: 12,
                  padding: "4px 10px",
                  borderRadius: 6,
                  background: "rgba(148, 163, 184, 0.15)",
                  border: "1px solid rgba(148, 163, 184, 0.3)",
                  whiteSpace: "nowrap",
                }}
              >
                {s.keys}
              </kbd>
            </li>
          ))}
        </ul>
        <p style={{ marginTop: 18, marginBottom: 0, fontSize: 12, opacity: 0.6 }}>
          Press <kbd style={{ fontFamily: "var(--font-mono)" }}>?</kbd> or <kbd style={{ fontFamily: "var(--font-mono)" }}>Esc</kbd> to close.
        </p>
      </div>
    </div>
  );
}

type SectionFilter = SlideSection | "all";
type SeverityFilter = RuleSeverity | "all";

function GridView({
  currentIndex,
  onPick,
  onClose,
}: {
  currentIndex: number;
  onPick: (n: number) => void;
  onClose: () => void;
}) {
  const [sectionFilter, setSectionFilter] = useState<SectionFilter>("all");
  const [severityFilter, setSeverityFilter] = useState<SeverityFilter>("all");

  const grouped = useMemo(() => {
    const all = groupBySection();
    return all
      .filter((group) => sectionFilter === "all" || group.section.id === sectionFilter)
      .map((group) => ({
        ...group,
        slides: group.slides.filter((s) =>
          severityFilter === "all" ? true : s.severity === severityFilter,
        ),
      }))
      .filter((group) => group.slides.length > 0);
  }, [sectionFilter, severityFilter]);

  const visibleCount = grouped.reduce((sum, g) => sum + g.slides.length, 0);
  const indexOf = (id: string) => DECK.findIndex((s) => s.id === id);

  return (
    <div className="grid-view">
      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          alignItems: "center",
          maxWidth: 1600,
          margin: "0 auto 16px",
          gap: 16,
          flexWrap: "wrap",
        }}
      >
        <h2 style={{ margin: 0, fontSize: 24 }}>
          Slides ({visibleCount} of {DECK.length})
        </h2>
        <div style={{ display: "flex", gap: 8, alignItems: "center", flexWrap: "wrap" }}>
          <FilterGroup
            label="Section"
            value={sectionFilter}
            onChange={(v) => setSectionFilter(v as SectionFilter)}
            options={[
              { value: "all", label: "All" },
              ...SECTIONS.map((s) => ({ value: s.id, label: s.label })),
            ]}
          />
          <FilterGroup
            label="Severity"
            value={severityFilter}
            onChange={(v) => setSeverityFilter(v as SeverityFilter)}
            options={[
              { value: "all", label: "All" },
              { value: "hard", label: "Hard" },
              { value: "warn", label: "Warn" },
              { value: "style", label: "Style" },
            ]}
          />
          <button
            onClick={onClose}
            style={{
              background: "hsl(var(--bg-raised))",
              color: "hsl(var(--fg))",
              border: "1px solid hsl(var(--border))",
              padding: "8px 16px",
              borderRadius: 8,
              cursor: "pointer",
            }}
          >
            Close (Esc)
          </button>
        </div>
      </div>

      {grouped.length === 0 ? (
        <div style={{ textAlign: "center", opacity: 0.7, padding: 48 }}>
          No slides match the current filters.
        </div>
      ) : (
        grouped.map((group) => (
          <section key={group.section.id} style={{ marginBottom: 32 }}>
            <div
              style={{
                maxWidth: 1600,
                margin: "0 auto 12px",
                display: "flex",
                alignItems: "baseline",
                gap: 12,
                borderBottom: "1px solid hsl(var(--border))",
                paddingBottom: 8,
              }}
            >
              <h3 style={{ margin: 0, fontSize: 18, letterSpacing: "0.02em" }}>
                {group.section.label}
              </h3>
              <span style={{ opacity: 0.6, fontSize: 13 }}>
                {group.section.description} · {group.slides.length} slide
                {group.slides.length === 1 ? "" : "s"}
              </span>
            </div>
            <div className="grid-view-grid">
              {group.slides.map((slide) => {
                const i = indexOf(slide.id);
                const Comp = slide.component;
                return (
                  <div
                    key={slide.id}
                    className="grid-thumb"
                    style={{
                      outline: i === currentIndex ? "2px solid hsl(var(--primary))" : "none",
                    }}
                    onClick={() => onPick(i)}
                  >
                    <ScaledSlide>
                      <Comp />
                    </ScaledSlide>
                    <div className="grid-thumb-label">
                      <span>
                        {String(i + 1).padStart(2, "0")} · {slide.title}
                      </span>
                      {slide.severity ? (
                        <RuleBadge severity={slide.severity} ruleId={slide.ruleId} />
                      ) : null}
                    </div>
                  </div>
                );
              })}
            </div>
          </section>
        ))
      )}
    </div>
  );
}

interface FilterOption {
  value: string;
  label: string;
}

function FilterGroup({
  label,
  value,
  onChange,
  options,
}: {
  label: string;
  value: string;
  onChange: (v: string) => void;
  options: FilterOption[];
}) {
  return (
    <label
      style={{
        display: "inline-flex",
        alignItems: "center",
        gap: 6,
        fontSize: 13,
        color: "hsl(var(--fg))",
      }}
    >
      <span style={{ opacity: 0.7 }}>{label}:</span>
      <select
        value={value}
        onChange={(e) => onChange(e.target.value)}
        style={{
          background: "hsl(var(--bg-raised))",
          color: "hsl(var(--fg))",
          border: "1px solid hsl(var(--border))",
          padding: "6px 10px",
          borderRadius: 6,
          cursor: "pointer",
        }}
      >
        {options.map((o) => (
          <option key={o.value} value={o.value}>
            {o.label}
          </option>
        ))}
      </select>
    </label>
  );
}

function PresenterView({
  Current,
  Next,
  index,
  total,
  onClose,
}: {
  Current: React.ComponentType;
  Next: React.ComponentType;
  index: number;
  total: number;
  onClose: () => void;
}) {
  const [seconds, setSeconds] = useState(0);
  useEffect(() => {
    const t = window.setInterval(() => setSeconds((s) => s + 1), 1000);
    return () => window.clearInterval(t);
  }, []);
  const mm = String(Math.floor(seconds / 60)).padStart(2, "0");
  const ss = String(seconds % 60).padStart(2, "0");
  return (
    <div className="presenter-view">
      <div className="presenter-main">
        <ScaledSlide>
          <Current />
        </ScaledSlide>
      </div>
      <div className="presenter-side">
        <div className="presenter-next">
          <ScaledSlide>
            <Next />
          </ScaledSlide>
        </div>
        <div className="presenter-notes">
          <strong>Slide {index + 1} of {total}</strong>
          {"\n\n"}
          Use ← → to navigate. Press P or Esc to exit presenter view. Press F for fullscreen.
        </div>
        <div className="presenter-timer">⏱ {mm}:{ss}</div>
        <button
          onClick={onClose}
          style={{
            background: "hsl(var(--bg-raised))",
            color: "hsl(var(--fg))",
            border: "1px solid hsl(var(--border))",
            padding: "8px 16px",
            borderRadius: 8,
            cursor: "pointer",
          }}
        >
          Close presenter (Esc)
        </button>
      </div>
    </div>
  );
}
