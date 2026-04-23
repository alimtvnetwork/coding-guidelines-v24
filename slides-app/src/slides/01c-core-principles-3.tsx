import { SlideLayout } from "@/components/SlideLayout";

const principles = [
  {
    n: "05",
    title: "Strict Function & File Metrics",
    body: "Functions 8-15 lines. Files under 300. React components under 100. Hard caps, not suggestions.",
  },
  {
    n: "06",
    title: "Spec-First Workflow",
    body: "Spec the change before writing code. Spec lives in spec/. AI agents and humans read the same source.",
  },
];

export default function CorePrinciples3Slide() {
  return (
    <SlideLayout
      eyebrow="Core Development Principles · 3 of 3"
      title="Constrain the work, free the mind"
      subtitle="Hard limits force small modules. A spec keeps everyone — human or AI — aligned."
    >
      <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", gap: 48, marginTop: 24 }}>
        {principles.map((p) => (
          <div
            key={p.n}
            style={{
              background: "hsl(var(--bg-raised))",
              border: "1px solid hsl(var(--border))",
              borderLeft: "6px solid hsl(var(--destructive))",
              borderRadius: 16,
              padding: 40,
              display: "flex",
              flexDirection: "column",
              gap: 16,
            }}
          >
            <div style={{ fontSize: 36, color: "hsl(var(--destructive))", fontWeight: 700, letterSpacing: "0.05em" }}>{p.n}</div>
            <div style={{ fontSize: 56, fontWeight: 700, lineHeight: 1.1, fontFamily: "Ubuntu, sans-serif" }}>{p.title}</div>
            <div style={{ fontSize: 28, color: "hsl(var(--muted-fg))", lineHeight: 1.4 }}>{p.body}</div>
          </div>
        ))}
      </div>
    </SlideLayout>
  );
}