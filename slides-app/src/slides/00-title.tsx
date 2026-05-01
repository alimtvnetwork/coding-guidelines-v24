import { motion } from "framer-motion";

const EASE = [0.22, 1, 0.36, 1] as const;

const HEADLINE_LINE_1 = "10 transformations";
const HEADLINE_LINE_2 = "that make code reviewable";

const wordParent = {
  hidden: {},
  show: { transition: { staggerChildren: 0.07, delayChildren: 0.25 } },
};

const wordChild = {
  hidden: { opacity: 0, y: 36, filter: "blur(10px)" },
  show: { opacity: 1, y: 0, filter: "blur(0px)", transition: { duration: 0.65, ease: EASE } },
};

interface AnimatedHeadlineProps {
  text: string;
  color?: string;
}

function AnimatedHeadline({ text, color }: AnimatedHeadlineProps) {
  const words = text.split(/\s+/);
  return (
    <span style={{ display: "inline-flex", flexWrap: "wrap", gap: "0.25em", justifyContent: "center", color }}>
      {words.map((word, i) => (
        <motion.span
          key={`${word}-${i}`}
          variants={wordChild}
          style={{ display: "inline-block" }}
        >
          {word}
        </motion.span>
      ))}
    </span>
  );
}

export default function TitleSlide() {
  return (
    <div
      style={{
        flex: 1,
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
        textAlign: "center",
        gap: "var(--space-4)",
      }}
    >
      <motion.div
        initial={{ opacity: 0, scale: 0.92, letterSpacing: "0.6em" }}
        animate={{ opacity: 1, scale: 1, letterSpacing: "0.4em" }}
        transition={{ duration: 0.9, ease: EASE }}
        style={{
          fontSize: 24,
          textTransform: "uppercase",
          color: "hsl(var(--primary))",
          fontWeight: 700,
        }}
      >
        Code-Red Review Guide
      </motion.div>
      <motion.h1
        variants={wordParent}
        initial="hidden"
        animate="show"
        style={{
          fontSize: 128,
          fontWeight: 700,
          margin: 0,
          lineHeight: 1.05,
          letterSpacing: "-0.02em",
          display: "flex",
          flexDirection: "column",
          gap: "0.1em",
        }}
      >
        <AnimatedHeadline text={HEADLINE_LINE_1} />
        <AnimatedHeadline text={HEADLINE_LINE_2} color="hsl(var(--accent))" />
      </motion.h1>
      <motion.div
        initial={{ opacity: 0, y: 12 }}
        animate={{ opacity: 1, y: 0 }}
        transition={{ duration: 0.6, delay: 1.4, ease: EASE }}
        style={{
          fontSize: 28,
          color: "hsl(var(--muted-fg))",
          marginTop: "var(--space-4)",
        }}
      >
        Md. Alim Ul Karim · Riseup Asia LLC
      </motion.div>
      <motion.div
        initial={{ opacity: 0 }}
        animate={{ opacity: 0.6 }}
        transition={{ duration: 0.6, delay: 1.8, ease: EASE }}
        style={{
          fontSize: 22,
          color: "hsl(var(--muted-fg))",
          marginTop: "var(--space-5)",
        }}
      >
        Press → to begin
      </motion.div>
    </div>
  );
}
