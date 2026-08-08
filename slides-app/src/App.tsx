import { useCallback, useEffect, useMemo, useRef, useState } from "react";
import { Maximize, Grid3x3, Presentation, Sun, Moon, HelpCircle, Search, X } from "lucide-react";
import { AnimatePresence, motion } from "framer-motion";
import { ScaledSlide } from "./components/ScaledSlide";
import { RuleBadge, type RuleSeverity } from "./components/RuleBadge";
import { DECK, SECTIONS, groupBySection, type SlideSection } from "./deck";
import { SlideStepContext } from "./lib/step-context";

type View = "deck" | "grid" | "presenter";

interface SlidePosition {
  index: number;
  step: number;
}

/**
 * Hash routing. Supported forms:
 *   #/id/<slide-id>          (preferred, stable across reorderings)
 *   #/id/<slide-id>/<step>   (with sub-step reveal coordinate, step >= 0)
 *   #/<index>                (legacy numeric, still honored)
 *   #/<index>/<step>         (legacy numeric + step)
 * Unknown or malformed hashes fall back to slide 0, step 0.
 */
function readSlideFromHash(): SlidePosition {
  const raw = window.location.hash;
  const idMatch = raw.match(/^#\/id\/([^/]+)(?:\/(\d+))?$/);
  if (idMatch) {
    const found = DECK.findIndex((s) => s.id === decodeURIComponent(idMatch[1]));
    if (found < 0) {
      console.warn(`[slides] unknown slide id in hash: ${idMatch[1]}, falling back to 0`);

      return { index: 0, step: 0 };
    }

    return { index: found, step: clampStep(found, idMatch[2] ? parseInt(idMatch[2], 10) : 0) };
  }

  const numMatch = raw.match(/^#\/(\d+)(?:\/(\d+))?/);
  if (!numMatch) return { index: 0, step: 0 };
  const idx = clampSlide(parseInt(numMatch[1], 10));

  return { index: idx, step: clampStep(idx, numMatch[2] ? parseInt(numMatch[2], 10) : 0) };
}

function clampSlide(n: number): number {
  if (Number.isNaN(n) || n < 0) return 0;
  if (n >= DECK.length) return DECK.length - 1;

  return n;
}

function clampStep(slideIndex: number, step: number): number {
  const max = DECK[slideIndex]?.steps ?? 0;
  if (Number.isNaN(step) || step < 0) return 0;
  if (step > max) return max;

  return step;
}

function writePositionToHash({ index, step }: SlidePosition) {
  const slide = DECK[index];
  const base = slide ? `/id/${encodeURIComponent(slide.id)}` : `/${index}`;
  window.location.hash = step > 0 ? `${base}/${step}` : base;
}

function isPrintMode(): boolean {
  return new URLSearchParams(window.location.search).has("print");
}

export default function App() {
  if (isPrintMode()) return <PrintView />;

  const initial = readSlideFromHash();
  const [index, setIndex] = useState(initial.index);
  const [step, setStep] = useState(initial.step);
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
      const pos = readSlideFromHash();
      setIndex(pos.index);
      setStep(pos.step);
    }

    window.addEventListener("hashchange", onHash);

    return () => window.removeEventListener("hashchange", onHash);
  }, []);

  const goto = useCallback((n: number, s: number = 0) => {
    const nextIndex = clampSlide(n);
    const nextStep = clampStep(nextIndex, s);
    writePositionToHash({ index: nextIndex, step: nextStep });
    setIndex(nextIndex);
    setStep(nextStep);
  }, []);

  const next = useCallback(() => {
    const maxStep = DECK[index]?.steps ?? 0;
    if (step < maxStep) {
      goto(index, step + 1);

      return;
    }

    goto(index + 1, 0);
  }, [index, step, goto]);

  const prev = useCallback(() => {
    if (step > 0) {
      goto(index, step - 1);

      return;
    }

    const prevIndex = clampSlide(index - 1);
    const prevMax = DECK[prevIndex]?.steps ?? 0;
    goto(prevIndex, prevMax);
  }, [index, step, goto]);

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
        onJump={(n) => goto(n)}
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
              <SlideStepContext.Provider value={{ step, maxStep: DECK[index]?.steps ?? 0 }}>
                <Current />
              </SlideStepContext.Provider>
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
        <button onClick={() => setPaletteOpen(true)} title="Command palette (Cmd/Ctrl+K)" aria-label="Command palette">
          <Search size={14} /> Search
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
      {paletteOpen ? (
        <CommandPalette
          onClose={() => setPaletteOpen(false)}
          onPick={(n) => {
            goto(n);
            setPaletteOpen(false);
          }}
        />
      ) : null}
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
  { keys: "Cmd/Ctrl + K", label: "Open command palette" },
  { keys: "? / H", label: "Toggle this help" },
  { keys: "Esc", label: "Close overlay or return to deck" },
];

function CommandPalette({
  onClose,
  onPick,
}: {
  onClose: () => void;
  onPick: (index: number) => void;
}) {
  const [query, setQuery] = useState("");
  const [active, setActive] = useState(0);
  const inputRef = useRef<HTMLInputElement | null>(null);

  useEffect(() => {
    inputRef.current?.focus();
  }, []);

  type Match =
    | { kind: "slide"; slide: (typeof DECK)[number]; index: number }
    | { kind: "section"; section: (typeof SECTIONS)[number]; index: number; count: number };

  const matches = useMemo<Match[]>(() => {
    const raw = query.trim().toLowerCase();
    // Section-jump prefixes: `s:<name>` or `#<name>` filter to sections only.
    const sectionOnly = raw.startsWith("s:") || raw.startsWith("#");
    const stripped = sectionOnly ? raw.replace(/^(s:|#)/, "").trim() : raw;

    const sectionHits: Match[] = SECTIONS
      .map((section) => {
        const idx = DECK.findIndex((s) => s.section === section.id);

        return { section, idx };
      })
      .filter(({ section, idx }) => {
        if (idx < 0) return false;
        if (!stripped) return sectionOnly;
        const hay = `${section.id} ${section.label} ${section.description}`.toLowerCase();

        return hay.includes(stripped);
      })
      .map(({ section, idx }) => ({
        kind: "section" as const,
        section,
        index: idx,
        count: DECK.filter((s) => s.section === section.id).length,
      }));

    if (sectionOnly) return sectionHits;

    const slideAll = DECK.map((slide, index) => ({ kind: "slide" as const, slide, index }));
    if (!stripped) return [...sectionHits, ...slideAll];
    const slideHits = slideAll.filter(({ slide }) => {
      const tagText = (slide.tags ?? []).join(" ");
      const hay = `${slide.id} ${slide.title} ${slide.ruleId ?? ""} ${slide.section} ${tagText}`.toLowerCase();

      return hay.includes(stripped);
    });

    return [...sectionHits, ...slideHits];
  }, [query]);

  useEffect(() => {
    setActive(0);
  }, [query]);

  function handleKey(e: React.KeyboardEvent<HTMLInputElement>) {
    if (e.key === "ArrowDown") {
      e.preventDefault();
      setActive((a) => Math.min(a + 1, matches.length - 1));
    } else if (e.key === "ArrowUp") {
      e.preventDefault();
      setActive((a) => Math.max(a - 1, 0));
    } else if (e.key === "Enter") {
      e.preventDefault();
      const hit = matches[active];
      if (hit) onPick(hit.index);
    } else if (e.key === "Escape") {
      e.preventDefault();
      onClose();
    }
  }

  return (
    <div
      role="dialog"
      aria-modal="true"
      aria-label="Command palette"
      onClick={onClose}
      style={{
        position: "fixed",
        inset: 0,
        background: "rgba(0, 0, 0, 0.65)",
        display: "flex",
        alignItems: "flex-start",
        justifyContent: "center",
        paddingTop: "10vh",
        zIndex: 1001,
      }}
    >
      <div
        onClick={(e) => e.stopPropagation()}
        style={{
          background: "hsl(var(--slide-bg, 222 47% 11%))",
          color: "hsl(var(--slide-fg, 210 40% 98%))",
          border: "1px solid rgba(148, 163, 184, 0.35)",
          borderRadius: 14,
          width: "min(640px, 92vw)",
          boxShadow: "0 20px 60px rgba(0, 0, 0, 0.5)",
          overflow: "hidden",
        }}
      >
        <div style={{ display: "flex", alignItems: "center", gap: 10, padding: "14px 16px", borderBottom: "1px solid rgba(148,163,184,0.2)" }}>
          <Search size={16} style={{ opacity: 0.6 }} />
          <input
            ref={inputRef}
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            onKeyDown={handleKey}
            placeholder="Search title, rule id, tag... or `s:react` / `#ops` to jump to a section"
            aria-label="Search slides"
            style={{
              flex: 1,
              background: "transparent",
              border: 0,
              outline: "none",
              color: "inherit",
              fontSize: 15,
              fontFamily: "inherit",
            }}
          />
          <kbd style={{ fontFamily: "var(--font-mono)", fontSize: 11, opacity: 0.6 }}>Esc</kbd>
        </div>
        <ul
          role="listbox"
          aria-label="Slide results"
          style={{ listStyle: "none", margin: 0, padding: 6, maxHeight: "50vh", overflowY: "auto" }}
        >
          {matches.length === 0 ? (
            <li style={{ padding: "18px 12px", opacity: 0.6, fontSize: 13 }}>No matches.</li>
          ) : (
            matches.map((match, i) => {
              const isActive = i === active;
              const commonRow = {
                display: "flex",
                alignItems: "center",
                gap: 10,
                padding: "10px 12px",
                borderRadius: 8,
                cursor: "pointer",
                background: isActive ? "rgba(148,163,184,0.15)" : "transparent",
              } as const;

              if (match.kind === "section") {
                return (
                  <li
                    key={`section:${match.section.id}`}
                    role="option"
                    aria-selected={isActive}
                    onMouseEnter={() => setActive(i)}
                    onClick={() => onPick(match.index)}
                    style={{ ...commonRow, borderLeft: "3px solid hsl(var(--accent, 217 91% 60%))" }}
                  >
                    <span style={{ fontFamily: "var(--font-mono)", fontSize: 11, opacity: 0.55, minWidth: 32 }}>§</span>
                    <span style={{ fontSize: 14, flex: 1, fontWeight: 600 }}>
                      Jump to section: {match.section.label}
                    </span>
                    <span style={{ fontSize: 11, opacity: 0.6 }}>{match.count} slides</span>
                    <span style={{ fontFamily: "var(--font-mono)", fontSize: 11, opacity: 0.5 }}>
                      → {String(match.index + 1).padStart(2, "0")}
                    </span>
                  </li>
                );
              }

              const { slide, index } = match;

              return (
                <li
                  key={slide.id}
                  role="option"
                  aria-selected={isActive}
                  onMouseEnter={() => setActive(i)}
                  onClick={() => onPick(index)}
                  style={commonRow}
                >
                  <span style={{ fontFamily: "var(--font-mono)", fontSize: 11, opacity: 0.55, minWidth: 32 }}>
                    {String(index + 1).padStart(2, "0")}
                  </span>
                  <span style={{ fontSize: 14, flex: 1 }}>{slide.title}</span>
                  {slide.tags && slide.tags.length > 0 ? (
                    <span style={{ display: "flex", gap: 4, flexWrap: "wrap", maxWidth: 220, justifyContent: "flex-end" }}>
                      {slide.tags.slice(0, 3).map((t) => (
                        <span
                          key={t}
                          style={{
                            fontSize: 10,
                            padding: "2px 6px",
                            borderRadius: 999,
                            background: "rgba(148,163,184,0.15)",
                            opacity: 0.85,
                            whiteSpace: "nowrap",
                          }}
                        >
                          {t}
                        </span>
                      ))}
                    </span>
                  ) : null}
                  {slide.ruleId ? (
                    <span style={{ fontFamily: "var(--font-mono)", fontSize: 11, opacity: 0.6 }}>{slide.ruleId}</span>
                  ) : null}
                  {slide.severity ? <RuleBadge severity={slide.severity} /> : null}
                  <span style={{ fontSize: 11, opacity: 0.5 }}>{slide.section}</span>
                </li>
              );
            })
          )}
        </ul>
      </div>
    </div>
  );
}

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
  onJump,
  onClose,
}: {
  Current: React.ComponentType;
  Next: React.ComponentType;
  index: number;
  total: number;
  onJump: (n: number) => void;
  onClose: () => void;
}) {
  const [seconds, setSeconds] = useState(0);
  useEffect(() => {
    const t = window.setInterval(() => setSeconds((s) => s + 1), 1000);

    return () => window.clearInterval(t);
  }, []);
  const mm = String(Math.floor(seconds / 60)).padStart(2, "0");
  const ss = String(seconds % 60).padStart(2, "0");
  const currentSection = DECK[index]?.section;
  const sectionEntries = useMemo(
    () =>
      SECTIONS.map((sec) => {
        const first = DECK.findIndex((s) => s.section === sec.id);
        const count = DECK.filter((s) => s.section === sec.id).length;

        return { sec, first, count };
      }).filter((e) => e.first >= 0),
    [],
  );

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
        <div className="presenter-chips" role="toolbar" aria-label="Jump to section">
          {sectionEntries.map(({ sec, first, count }) => {
            const active = sec.id === currentSection;

            return (
              <button
                key={sec.id}
                onClick={() => onJump(first)}
                aria-pressed={active}
                title={`${sec.label} · ${count} slide${count === 1 ? "" : "s"}`}
                style={{
                  fontSize: 12,
                  padding: "5px 10px",
                  borderRadius: 999,
                  cursor: "pointer",
                  border: `1px solid ${active ? "hsl(var(--primary))" : "hsl(var(--border))"}`,
                  background: active ? "hsl(var(--primary))" : "hsl(var(--bg-raised))",
                  color: active ? "hsl(var(--primary-fg, var(--bg)))" : "hsl(var(--fg))",
                  whiteSpace: "nowrap",
                }}
              >
                {sec.label} <span style={{ opacity: 0.65 }}>· {count}</span>
              </button>
            );
          })}
        </div>
        <div className="presenter-notes">
          <strong>Slide {index + 1} of {total}</strong>
          {"\n\n"}
          Use ← → to navigate. Click a chip above to jump to a section. Press P or Esc to exit, F for fullscreen.
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

/**
 * Print handout: renders every slide at native 1920x1080, stacked vertically.
 * Trigger via `?print`. Combined with the @page rule in slide.css, Cmd/Ctrl+P →
 * Save as PDF produces a landscape handout that matches the on-screen design.
 * See slides-app skill §10 (Print Mode).
 */
function PrintView() {
  useEffect(() => {
    document.title = `Handout · ${DECK.length} slides`;
  }, []);

  return (
    <div className="print-root">
      {DECK.map((slide, i) => {
        const S = slide.component;

        return (
          <section key={slide.id} className="print-page" aria-label={`Slide ${i + 1}: ${slide.title}`}>
            <div className="slide-content">
              <S />
            </div>
            <div className="print-page-stamp" aria-hidden="true">
              {String(i + 1).padStart(2, "0")} / {DECK.length}
              {slide.ruleId ? ` · ${slide.ruleId}` : ""} · {slide.title}
            </div>
          </section>
        );
      })}
    </div>
  );
}
