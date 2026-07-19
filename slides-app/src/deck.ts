/**
 * Backward-compatible re-export. The canonical registry now lives in
 * `./deck/registry.ts` with `SlideSection` grouping and rule severity metadata.
 * Existing imports from `@/deck` keep working.
 */
export { DECK, SECTIONS, groupBySection } from "./deck/registry";
export type { SlideEntry, SlideSection, SlideSectionMeta } from "./deck/registry";
