import { motion } from "framer-motion";
import type { ReactNode } from "react";

export interface SlideLayoutProps {
  eyebrow?: string;
  title: string;
  subtitle?: string;
  notes?: string;
  footer?: ReactNode;
  children: ReactNode;
}

const EASE = [0.22, 1, 0.36, 1] as const;

const fadeUp = {
  hidden: { opacity: 0, y: 24 },
  show: (delay: number) => ({
    opacity: 1,
    y: 0,
    transition: { duration: 0.5, delay, ease: EASE },
  }),
};

// Per-word reveal so headlines feel composed, not pasted in.
const wordParent = {
  hidden: {},
  show: { transition: { staggerChildren: 0.06, delayChildren: 0.15 } },
};

const wordChild = {
  hidden: { opacity: 0, y: 28, filter: "blur(8px)" },
  show: { opacity: 1, y: 0, filter: "blur(0px)", transition: { duration: 0.55, ease: EASE } },
};

function isStringTruthy(value?: string): boolean {
  return typeof value === "string" && value.length > 0;
}

function splitIntoWords(text: string): string[] {
  return text.split(/\s+/);
}

export function SlideLayout({
  eyebrow,
  title,
  subtitle,
  footer,
  children,
}: SlideLayoutProps) {
  const hasEyebrow = isStringTruthy(eyebrow);
  const hasSubtitle = isStringTruthy(subtitle);
  const words = splitIntoWords(title);
  return (
    <>
      {hasEyebrow && (
        <motion.div
          className="eyebrow"
          variants={fadeUp}
          initial="hidden"
          animate="show"
          custom={0.05}
        >
          {eyebrow}
        </motion.div>
      )}
      <motion.h1
        className="slide-title"
        variants={wordParent}
        initial="hidden"
        animate="show"
        style={{ display: "flex", flexWrap: "wrap", gap: "0.25em" }}
      >
        {words.map((word, i) => (
          <motion.span
            key={`${word}-${i}`}
            variants={wordChild}
            style={{ display: "inline-block" }}
          >
            {word}
          </motion.span>
        ))}
      </motion.h1>
      {hasSubtitle && (
        <motion.div
          className="slide-subtitle"
          variants={fadeUp}
          initial="hidden"
          animate="show"
          custom={0.35}
        >
          {subtitle}
        </motion.div>
      )}
      <motion.div
        className="slide-body"
        variants={fadeUp}
        initial="hidden"
        animate="show"
        custom={0.5}
      >
        {children}
      </motion.div>
      {footer && <div className="slide-footer">{footer}</div>}
    </>
  );
}
