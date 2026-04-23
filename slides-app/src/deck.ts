import type { ComponentType } from "react";
import Title from "./slides/00-title";
import CorePrinciples1 from "./slides/01a-core-principles-1";
import CorePrinciples2 from "./slides/01b-core-principles-2";
import CorePrinciples3 from "./slides/01c-core-principles-3";
import Naming from "./slides/01-naming-conventions";
import NestedIf from "./slides/02-nested-if-else";
import BooleanPrefixes from "./slides/03-boolean-prefixes";
import AppErrorWrapper from "./slides/04-app-error-wrapper";
import StructuredLogging from "./slides/05-structured-logging";
import MagicStrings from "./slides/06-magic-strings";
import Metrics from "./slides/07-function-and-file-metrics";
import TwoOperandMax from "./slides/08-two-operand-max";
import PositiveGuards from "./slides/09-positively-named-guards";
import SpecFirst from "./slides/10-spec-first-workflow";
import CacheInvalidation from "./slides/11-cache-invalidation";
import Closing from "./slides/12-closing";

export interface SlideEntry {
  id: string;
  title: string;
  notes?: string;
  component: ComponentType;
}

export const DECK: readonly SlideEntry[] = [
  { id: "00-title", title: "Title", component: Title },
  { id: "01a-core-principles-1", title: "Core Principles · 1", component: CorePrinciples1 },
  { id: "01b-core-principles-2", title: "Core Principles · 2", component: CorePrinciples2 },
  { id: "01c-core-principles-3", title: "Core Principles · 3", component: CorePrinciples3 },
  { id: "01-naming", title: "Naming conventions", component: Naming },
  { id: "02-nested-if", title: "Nested if-else", component: NestedIf },
  { id: "03-boolean-prefixes", title: "Boolean prefixes", component: BooleanPrefixes },
  { id: "04-app-error", title: "AppError wrapper", component: AppErrorWrapper },
  { id: "05-logging", title: "Structured logging", component: StructuredLogging },
  { id: "06-magic-strings", title: "Magic strings", component: MagicStrings },
  { id: "07-metrics", title: "Function & file metrics", component: Metrics },
  { id: "08-two-operand", title: "Two-operand max", component: TwoOperandMax },
  { id: "09-positive-guards", title: "Positive guards", component: PositiveGuards },
  { id: "10-spec-first", title: "Spec-first workflow", component: SpecFirst },
  { id: "11-cache-invalidation", title: "Cache invalidation", component: CacheInvalidation },
  { id: "12-closing", title: "Closing", component: Closing },
] as const;
