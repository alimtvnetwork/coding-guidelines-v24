import { SlideLayout } from "@/components/SlideLayout";

const principles = [
  {
    n: "03",
    title: "Positively Named Guards",
    body: "isReady, hasError, canPublish. Never !isNotReady. The reader should not invert booleans in their head.",
  },
  {
    n: "04",
    title: "Structured Error Wrapping",
    body: "Every error crosses a boundary as AppError with stack trace and context. No silent swallowing. Ever.",
  },
];

export default function CorePrinciples2Slide() {
  return (
    <SlideLayout
      eyebrow="Core Development Principles · 2 of 3"
      title="Make intent unambiguous"
      subtitle="Booleans read like sentences. Errors carry their own evidence."
    >
      <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", gap: 48, marginTop: 24 }}>
        {principles.map((p) => (
          <div
            key={p.n}
            style={{
              background: "hsl(var(--bg-raised))",
              border: "1px solid hsl(var(--border))",
              borderLeft: "6px solid hsl(var(--primary))",
              borderRadius: 16,
              padding: 40,
              display: "flex",
              flexDirection: "column",
              gap: 16,
            }}
          >
            <div style={{ fontSize: 36, color: "hsl(var(--primary))", fontWeight: 700, letterSpacing: "0.05em" }}>{p.n}</div>
            <div style={{ fontSize: 56, fontWeight: 700, lineHeight: 1.1, fontFamily: "Ubuntu, sans-serif" }}>{p.title}</div>
            <div style={{ fontSize: 28, color: "hsl(var(--muted-fg))", lineHeight: 1.4 }}>{p.body}</div>
          </div>
        ))}
      </div>
    </SlideLayout>
  );
}