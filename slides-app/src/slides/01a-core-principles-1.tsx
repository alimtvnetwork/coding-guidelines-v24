import { SlideLayout } from "@/components/SlideLayout";

const principles = [
  {
    n: "01",
    title: "Zero-Nesting Discipline",
    body: "No nested if-else. Use early-return guards. One indentation level inside a function body.",
  },
  {
    n: "02",
    title: "Two-Operand Maximum",
    body: "Boolean expressions take at most two operands. Extract the third into a positively named guard.",
  },
];

export default function CorePrinciples1Slide() {
  return (
    <SlideLayout
      eyebrow="Core Development Principles · 1 of 3"
      title="Reviewable code starts here"
      subtitle="Two principles per slide. Internalise them in this order."
    >
      <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", gap: 48, marginTop: 24 }}>
        {principles.map((p) => (
          <div
            key={p.n}
            style={{
              background: "hsl(var(--bg-raised))",
              border: "1px solid hsl(var(--border))",
              borderLeft: "6px solid hsl(var(--accent))",
              borderRadius: 16,
              padding: 40,
              display: "flex",
              flexDirection: "column",
              gap: 16,
            }}
          >
            <div style={{ fontSize: 36, color: "hsl(var(--accent))", fontWeight: 700, letterSpacing: "0.05em" }}>{p.n}</div>
            <div style={{ fontSize: 56, fontWeight: 700, lineHeight: 1.1, fontFamily: "Ubuntu, sans-serif" }}>{p.title}</div>
            <div style={{ fontSize: 28, color: "hsl(var(--muted-fg))", lineHeight: 1.4 }}>{p.body}</div>
          </div>
        ))}
      </div>
    </SlideLayout>
  );
}