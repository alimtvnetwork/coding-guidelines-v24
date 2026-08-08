import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CalloutQuote } from "@/components/CalloutQuote";

/**
 * SS-02 task 19: "Comments lie, code does not" slide.
 *
 * Reinforces the guideline that documentation drifts while code runs. Uses
 * the Go stdlib `path.Clean` documented example as evidence: the doc example
 * is executable, run by `go test`, and therefore cannot silently rot.
 */

interface EvidenceRow {
  label: string;
  detail: string;
}

const ROWS: readonly EvidenceRow[] = [
  {
    label: "Prose comment",
    detail:
      "// returns the shortest path equivalent — drifts silently when the function changes; nothing fails.",
  },
  {
    label: "Doc example",
    detail:
      "func ExamplePath_Clean() { ... }  Executed by `go test`; output mismatch fails the build.",
  },
  {
    label: "Effect",
    detail:
      "The example IS the spec. Rename an arg, change a return, the CI turns red before the doc ever ships stale.",
  },
];

const GO_EXAMPLE = `func ExamplePath_Clean() {
    paths := []string{
        "path/to/../file",
        "path//double-slash",
        "path/./same-dir/",
    }

    for _, p := range paths {
        fmt.Printf("Clean(%q) = %q\\n", p, path.Clean(p))
    }

    // Output:
    // Clean("path/to/../file")   = "path/file"
    // Clean("path//double-slash")= "path/double-slash"
    // Clean("path/./same-dir/")  = "path/same-dir"
}`;

export default function CommentsLieCodeDoesNotSlide() {
  return (
    <SlideLayout
      eyebrow="Principle · Documentation as code"
      title="Comments lie. Code does not."
      subtitle="Prose drifts silently. Executable examples fail loudly. Prefer the loud signal."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 20, marginTop: 8 }}>
        <ActionPanel
          slideId="16-comments-lie-code-does-not"
          symptom="Doc comments that describe a function's behavior word for word, then drift out of sync the first time the function is refactored, with zero failing signal."
          rule="Write executable documentation. In Go use `ExampleXxx` funcs (run by `go test`); in TS/Rust use doctests or Storybook stories. Free-prose comments only for irreducible intent that the code cannot express."
          doThis="On the next behavior change, replace one prose comment with a runnable example (Go `Example`, TS test, Rust doctest). Delete the prose. Watch CI verify the doc for you."
        />
        <div style={{ display: "grid", gridTemplateColumns: "1.15fr 1fr", gap: 24 }}>
          <div
            style={{
              padding: "18px 22px",
              borderRadius: 14,
              background: "hsl(var(--bg-raised))",
              border: "1px solid hsl(var(--primary) / 0.4)",
              borderTop: "4px solid hsl(var(--primary))",
              display: "flex",
              flexDirection: "column",
              gap: 10,
              minWidth: 0,
            }}
          >
            <div
              style={{
                fontSize: 14,
                fontWeight: 700,
                letterSpacing: "0.12em",
                textTransform: "uppercase",
                color: "hsl(var(--primary))",
              }}
            >
              Go stdlib · path.Clean example
            </div>
            <pre
              style={{
                margin: 0,
                padding: 0,
                fontFamily: "var(--font-mono)",
                fontSize: 13,
                lineHeight: 1.45,
                color: "hsl(var(--fg))",
                whiteSpace: "pre",
                overflow: "hidden",
              }}
            >
              {GO_EXAMPLE}
            </pre>
            <div style={{ fontSize: 12, color: "hsl(var(--fg-muted))" }}>
              Source: go.dev/src/go/doc/example.go
            </div>
          </div>
          <div style={{ display: "flex", flexDirection: "column", gap: 10 }}>
            {ROWS.map((row) => (
              <div
                key={row.label}
                style={{
                  padding: "14px 18px",
                  borderRadius: 12,
                  background: "hsl(var(--bg-raised))",
                  border: "1px solid hsl(var(--border))",
                  display: "flex",
                  flexDirection: "column",
                  gap: 4,
                }}
              >
                <div
                  style={{
                    fontSize: 12,
                    fontWeight: 700,
                    letterSpacing: "0.1em",
                    textTransform: "uppercase",
                    color: "hsl(var(--accent))",
                  }}
                >
                  {row.label}
                </div>
                <div style={{ fontSize: 14, lineHeight: 1.4, color: "hsl(var(--fg))" }}>
                  {row.detail}
                </div>
              </div>
            ))}
            <CalloutQuote
              quote="If the comment and the code disagree, both are wrong."
              attribution="Norm Schryer"
              accent="accent"
            />
          </div>
        </div>
      </div>
    </SlideLayout>
  );
}
