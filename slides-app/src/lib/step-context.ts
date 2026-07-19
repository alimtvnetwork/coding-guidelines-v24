import { createContext, useContext } from "react";

/**
 * SlideStepContext exposes the current sub-step index to slide components
 * for staged reveals (bullets appearing one at a time, multi-stage diffs).
 *
 * Contract:
 *   - `step` is 0-based. Step 0 = initial render (before any reveals).
 *   - A slide with `steps: 3` in the DECK registry supports steps 0..3
 *     (4 total render states).
 *   - Slides that don't opt in ignore this context entirely; `step` stays 0.
 */
export interface SlideStepValue {
  step: number;
  maxStep: number;
}

export const SlideStepContext = createContext<SlideStepValue>({ step: 0, maxStep: 0 });

/** Hook slides use to read the current reveal step. */
export function useSlideStep(): SlideStepValue {
  return useContext(SlideStepContext);
}

/** Convenience: true when the given step index should be visible now. */
export function isStepVisible(index: number, current: number): boolean {
  return index <= current;
}
