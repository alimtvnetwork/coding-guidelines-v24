import { SlideLayout } from "@/components/SlideLayout";
import { PrincipleCard } from "@/components/PrincipleCard";

export default function CorePrinciples3Slide() {
  return (
    <SlideLayout
      eyebrow="Core Development Principles · 3 of 3"
      title="Constrain the work, free the mind"
      subtitle="Hard limits force small modules. A shared spec keeps every author, human or AI, aligned."
    >
      <div style={{ display: "grid", gridTemplateColumns: "1fr 1fr", gap: 36, marginTop: 16 }}>
        <PrincipleCard
          number="05"
          progressId="01c-core-principles-3#05"
          title="Strict Function & File Metrics"
          symptom="200-line functions and 800-line files hide bugs and defeat every reviewer after the first screen."
          rule="Functions 8-15 lines. Files under 300. React components under 100. Hard caps, not suggestions."
          action="Run `npm run lint`. Split any file or function the `max-lines` rule flags before pushing."
          accent="destructive"
          delay={0.5}
        />
        <PrincipleCard
          number="06"
          progressId="01c-core-principles-3#06"
          title="Spec-First Workflow"
          symptom="Code lands before the spec, so intent drifts across authors and AI agents rebuild the wrong thing."
          rule="Spec the change under `spec/` before writing code. Humans and AI agents read the same source of truth."
          action="Open `spec/` and add a numbered file for your next feature. Get it reviewed before opening a code PR."
          accent="destructive"
          delay={0.65}
        />
      </div>
    </SlideLayout>
  );
}
