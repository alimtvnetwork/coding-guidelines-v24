import type { ComponentType } from "react";
import Title from "../slides/00-title";
import TableOfContents from "../slides/01-table-of-contents";
import CorePrinciples1 from "../slides/01a-core-principles-1";
import CorePrinciples2 from "../slides/01b-core-principles-2";
import CorePrinciples3 from "../slides/01c-core-principles-3";
import Naming from "../slides/01-naming-conventions";
import NestedIf from "../slides/02-nested-if-else";
import BooleanPrefixes from "../slides/03-boolean-prefixes";
import AppErrorWrapper from "../slides/04-app-error-wrapper";
import StructuredLogging from "../slides/05-structured-logging";
import MagicStrings from "../slides/06-magic-strings";
import Metrics from "../slides/07-function-and-file-metrics";
import TwoOperandMax from "../slides/08-two-operand-max";
import PositiveGuards from "../slides/09-positively-named-guards";
import SpecFirst from "../slides/10-spec-first-workflow";
import CacheInvalidation from "../slides/11-cache-invalidation";
import Closing from "../slides/12-closing";
import type { RuleSeverity } from "../components/RuleBadge";

/**
 * Grouping surfaced in grid view, TOC, and (future) presenter filters.
 * Order here is the canonical section order.
 */
export type SlideSection =
  | "opening"
  | "principles"
  | "naming"
  | "control-flow"
  | "errors"
  | "react"
  | "workflow"
  | "closing";

export interface SlideSectionMeta {
  id: SlideSection;
  label: string;
  /** Short description used as the section divider subtitle in grid view. */
  description: string;
}

export const SECTIONS: readonly SlideSectionMeta[] = [
  { id: "opening", label: "Opening", description: "Title and orientation" },
  { id: "principles", label: "Core Principles", description: "Non-negotiables and mindset" },
  { id: "naming", label: "Naming & Structure", description: "Identifiers, size, shape" },
  { id: "control-flow", label: "Control Flow", description: "Guards, effects, loops" },
  { id: "errors", label: "Error Management", description: "Never swallow, always log" },
  { id: "react", label: "React & TypeScript", description: "Types, hooks, a11y" },
  { id: "workflow", label: "Workflow & Ops", description: "Spec-first, testing, caching" },
  { id: "closing", label: "Closing", description: "Wrap-up and references" },
] as const;

export interface SlideEntry {
  id: string;
  title: string;
  section: SlideSection;
  /** Optional severity; set for rule slides so `RuleBadge` can render it. */
  severity?: RuleSeverity;
  /** Optional rule id from `spec/17/31`, e.g. "NAM-001". */
  ruleId?: string;
  notes?: string;
  component: ComponentType;
}

/**
 * Canonical deck. New rule slides append to their section block; the runtime
 * order is preserved by iteration order below.
 */
export const DECK: readonly SlideEntry[] = [
  { id: "00-title", title: "Title", section: "opening", component: Title },
  { id: "01-toc", title: "Table of Contents", section: "opening", component: TableOfContents },

  { id: "01a-core-principles-1", title: "Core Principles · 1", section: "principles", component: CorePrinciples1 },
  { id: "01b-core-principles-2", title: "Core Principles · 2", section: "principles", component: CorePrinciples2 },
  { id: "01c-core-principles-3", title: "Core Principles · 3", section: "principles", component: CorePrinciples3 },

  { id: "01-naming", title: "Naming conventions", section: "naming", severity: "hard", ruleId: "NAM-001", component: Naming },
  { id: "07-metrics", title: "Function & file metrics", section: "naming", severity: "hard", ruleId: "SIZE-001", component: Metrics },

  { id: "02-nested-if", title: "Nested if-else", section: "control-flow", severity: "hard", ruleId: "CF-001", component: NestedIf },
  { id: "03-boolean-prefixes", title: "Boolean prefixes", section: "control-flow", severity: "warn", ruleId: "BOOL-001", component: BooleanPrefixes },
  { id: "08-two-operand", title: "Two-operand max", section: "control-flow", severity: "hard", ruleId: "CF-002", component: TwoOperandMax },
  { id: "09-positive-guards", title: "Positive guards", section: "control-flow", severity: "warn", ruleId: "CF-003", component: PositiveGuards },

  { id: "04-app-error", title: "AppError wrapper", section: "errors", severity: "hard", ruleId: "ERR-001", component: AppErrorWrapper },
  { id: "05-logging", title: "Structured logging", section: "errors", severity: "warn", ruleId: "ERR-002", component: StructuredLogging },
  { id: "06-magic-strings", title: "Magic strings", section: "errors", severity: "hard", ruleId: "ERR-003", component: MagicStrings },

  { id: "10-spec-first", title: "Spec-first workflow", section: "workflow", severity: "style", ruleId: "WF-001", component: SpecFirst },
  { id: "11-cache-invalidation", title: "Cache invalidation", section: "workflow", severity: "warn", ruleId: "WF-002", component: CacheInvalidation },

  { id: "12-closing", title: "Closing", section: "closing", component: Closing },
] as const;

/** Group deck entries by section, preserving section order defined in `SECTIONS`. */
export function groupBySection(): ReadonlyArray<{ section: SlideSectionMeta; slides: readonly SlideEntry[] }> {
  return SECTIONS.map((section) => ({
    section,
    slides: DECK.filter((slide) => slide.section === section.id),
  }));
}
