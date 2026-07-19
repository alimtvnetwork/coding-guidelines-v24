import type { ComponentType } from "react";
import Title from "../slides/00-title";
import TableOfContents from "../slides/01-table-of-contents";
import MustFollow from "../slides/13-must-follow";
import RootCauseFirst from "../slides/14-root-cause-first";
import SpecFirstVsCodeFirst from "../slides/15-spec-first-vs-code-first";
import CommentsLieCodeDoesNot from "../slides/16-comments-lie-code-does-not";
import MethodDocDecisionTree from "../slides/17-method-doc-decision-tree";
import VersionBumpMythBuster from "../slides/18-version-bump-myth-buster";
import TrustBoundariesTeaser from "../slides/19-trust-boundaries-teaser";
import BackupTierFreeze from "../slides/20-backup-tier-freeze";
import DbSchemaNaming from "../slides/21-db-schema-naming";
import ZeroUnderscoreAcronyms from "../slides/22-zero-underscore-acronyms";
import BooleanNaming from "../slides/23-boolean-naming";
import NoBooleanParameters from "../slides/24-no-boolean-parameters";
import EnumStandards from "../slides/25-enum-standards";
import LineGapDiscipline from "../slides/26-line-gap-discipline";
import FileSizeTiers from "../slides/27-file-size-tiers";
import FunctionLength from "../slides/28-function-length";
import ImmutableFirst from "../slides/29-immutable-first";
import DedicatedDefinitions from "../slides/30-dedicated-definitions";
import DryExtractNow from "../slides/31-dry-extract-now";
import MermaidFirstComponents from "../slides/32-mermaid-first-components";
import AssetsFolderConvention from "../slides/33-assets-folder-convention";
import RegisteredErrorCodes from "../slides/34-registered-error-codes";
import ResponseEnvelope from "../slides/35-response-envelope";
import LogLevelSeverity from "../slides/36-log-level-severity";
import LogContext from "../slides/37-log-context";
import VerifyBothDirections from "../slides/38-verify-both-directions";
import RetrospectiveOnRepeats from "../slides/39-retrospective-on-repeats";
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
  /**
   * Optional max sub-step index for staged reveals. When set to N, the slide
   * supports steps 0..N (N+1 total render states). Consumed by App.tsx nav
   * logic and by slide components via `useSlideStep()`. Absent = 0 (no reveals).
   */
  steps?: number;
  /**
   * Optional free-form search tags for the command palette. Match on concept
   * keywords (e.g. "guard", "effect", "wcag") that are not in the title,
   * rule id, or section. Lowercase; kept short (1-3 words each).
   */
  tags?: readonly string[];
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

  { id: "13-must-follow", title: "Must Follow · Five non-negotiables", section: "principles", severity: "hard", ruleId: "MUST-001", tags: ["must-follow", "non-negotiable", "root cause", "verify", "changelog"], component: MustFollow },
  { id: "14-root-cause-first", title: "Root cause before fix", section: "principles", severity: "hard", ruleId: "MUST-002", tags: ["root cause", "rca", "debugging", "symptom patch", "one sentence"], component: RootCauseFirst },
  { id: "15-spec-first-vs-code-first", title: "Spec-first vs code-first · PR contrast", section: "principles", severity: "hard", ruleId: "MUST-003", tags: ["spec-first", "pull request", "review", "contract", "before after"], component: SpecFirstVsCodeFirst },
  { id: "16-comments-lie-code-does-not", title: "Comments lie, code does not", section: "principles", severity: "warn", ruleId: "MUST-004", tags: ["comments", "docs", "example", "go", "doctest", "path.clean"], component: CommentsLieCodeDoesNot },
  { id: "17-method-doc-decision-tree", title: "Method-doc decision tree", section: "principles", severity: "warn", ruleId: "MUST-005", tags: ["docs", "decision tree", "checklist", "rename", "split", "godoc"], component: MethodDocDecisionTree },
  { id: "18-version-bump-myth-buster", title: "Version bump myth-buster", section: "principles", severity: "hard", ruleId: "MUST-006", tags: ["version", "semver", "changelog", "release", "v1.2", "myth"], component: VersionBumpMythBuster },
  { id: "19-trust-boundaries-teaser", title: "Trust boundaries · Main > Worker > Backup > Git", section: "principles", severity: "hard", ruleId: "MUST-007", tags: ["trust", "isolation", "spec/19", "backup", "git", "one-way"], component: TrustBoundariesTeaser },
  { id: "20-backup-tier-freeze", title: "Backup-tier freeze · defaults today, contract at Phase 12", section: "principles", severity: "warn", ruleId: "MUST-008", tags: ["backup", "freeze", "spec/19", "phase 12", "MAIN-900-01", "defer"], component: BackupTierFreeze },
  { id: "01a-core-principles-1", title: "Core Principles · 1", section: "principles", tags: ["mindset", "must-follow", "non-negotiable"], component: CorePrinciples1 },
  { id: "01b-core-principles-2", title: "Core Principles · 2", section: "principles", tags: ["mindset", "review"], component: CorePrinciples2 },
  { id: "01c-core-principles-3", title: "Core Principles · 3", section: "principles", tags: ["mindset", "ownership"], component: CorePrinciples3 },

  { id: "01-naming", title: "Naming conventions", section: "naming", severity: "hard", ruleId: "NAM-001", tags: ["pascalcase", "camelcase", "identifiers", "acronyms"], component: Naming },
  { id: "21-db-schema-naming", title: "DB schema naming · PascalCase entities, camelCase fields, {Table}Id PKs", section: "naming", severity: "hard", ruleId: "NAM-002", tags: ["database", "schema", "pascalcase", "camelcase", "primary key", "uuid"], component: DbSchemaNaming },
  { id: "22-zero-underscore-acronyms", title: "Zero underscores · acronyms stay UPPERCASE", section: "naming", severity: "hard", ruleId: "NAM-003", tags: ["underscore", "acronym", "uppercase", "http", "url", "json", "api"], component: ZeroUnderscoreAcronyms },
  { id: "23-boolean-naming", title: "Boolean naming · is/has/can/should, positive only", section: "naming", severity: "hard", ruleId: "BOOL-002", tags: ["boolean", "is", "has", "can", "should", "positive framing", "negation"], component: BooleanNaming },
  { id: "24-no-boolean-parameters", title: "No boolean parameters · split into two named functions", section: "naming", severity: "hard", ruleId: "BOOL-003", tags: ["boolean", "parameter", "flag", "refactor", "render", "named function"], component: NoBooleanParameters },
  { id: "25-enum-standards", title: "Enum standards · PascalCase, strict parse, per-language idiom", section: "naming", severity: "hard", ruleId: "ENUM-001", tags: ["enum", "pascalcase", "iota", "isequal", "magic string", "typescript", "go", "php"], component: EnumStandards },
  { id: "26-line-gap-discipline", title: "Line-gap discipline · grouped imports, blank before return, no double blanks", section: "structure", severity: "hard", ruleId: "STYLE-001", tags: ["whitespace", "blank line", "imports", "return", "structure", "readability"], component: LineGapDiscipline },
  { id: "27-file-size-tiers", title: "File-size tiers · 100 tsx, 120 class, 300 fallback", section: "structure", severity: "hard", ruleId: "SIZE-001", tags: ["file size", "300 lines", "100 lines", "component", "class", "refactor", "extract"], component: FileSizeTiers },
  { id: "28-function-length", title: "Function length · 8 preferred, 15 hard cap, lint-allow waiver", section: "structure", severity: "hard", ruleId: "FUNC-001", tags: ["function length", "8 lines", "15 lines", "waiver", "lint-allow", "extract", "refactor"], component: FunctionLength },
  { id: "29-immutable-first", title: "Immutable-first · assign once, spread-copy, no in-place mutation", section: "structure", severity: "hard", ruleId: "IMMUT-001", tags: ["immutable", "const", "let", "mutation", "spread", "rust", "reassignment"], component: ImmutableFirst },
  { id: "30-dedicated-definitions", title: "Dedicated definitions files · types, enums, constants get their own file", section: "structure", severity: "hard", ruleId: "DEF-001", tags: ["types", "enums", "constants", "interfaces", "extract", "colocation", "types.ts"], component: DedicatedDefinitions },
  { id: "31-dry-extract-now", title: "DRY extract-now · two sites is the trigger, not three", section: "structure", severity: "hard", ruleId: "DRY-001", tags: ["dry", "duplication", "extract", "helper", "drift", "copy-paste"], component: DryExtractNow },
  { id: "32-mermaid-first-components", title: "Mermaid-first · diagram before three components land", section: "structure", severity: "hard", ruleId: "COMP-001", tags: ["mermaid", "diagram", "components", "architecture", "state ownership", "flowchart"], component: MermaidFirstComponents },
  { id: "33-assets-folder-convention", title: "Assets folder · NN-folder / NN-file, no -final, no -v2", section: "structure", severity: "hard", ruleId: "ASSET-001", tags: ["assets", "folder", "naming", "sequence prefix", "icons", "logos", "convention"], component: AssetsFolderConvention },
  { id: "34-registered-error-codes", title: "Registered error codes · one registry, typed union, exhaustive switch", section: "errors", severity: "hard", ruleId: "ERR-004", tags: ["error codes", "registry", "apperror", "exhaustive switch", "typed union", "support"], component: RegisteredErrorCodes },
  { id: "35-response-envelope", title: "Response envelope · { data, errors[], meta } and one apiCall parser", section: "errors", severity: "hard", ruleId: "ERR-005", tags: ["envelope", "api", "apiCall", "requestId", "typed errors", "shared parser"], component: ResponseEnvelope },
  { id: "36-log-level-severity", title: "Log-level severity map · debug/info/warn/error/fatal, one meaning each", section: "errors", severity: "hard", ruleId: "LOG-001", tags: ["logging", "severity", "debug", "info", "warn", "error", "fatal", "on-call"], component: LogLevelSeverity },
  { id: "37-log-context", title: "Log context · op, requestId, key inputs, never secrets or PII", section: "errors", severity: "hard", ruleId: "LOG-002", tags: ["logging", "context", "requestId", "redaction", "pii", "secrets", "grep"], component: LogContext },
  { id: "38-verify-both-directions", title: "Verify both directions · curl the backend AND inspect the frontend on the same payload", section: "errors", severity: "hard", ruleId: "INT-001", tags: ["integration", "verification", "curl", "contract", "frontend", "backend", "pr-checklist"], component: VerifyBothDirections },
  { id: "07-metrics", title: "Function & file metrics", section: "naming", severity: "hard", ruleId: "SIZE-001", tags: ["size", "cognitive complexity", "function length", "file length"], component: Metrics },

  { id: "02-nested-if", title: "Nested if-else", section: "control-flow", severity: "hard", ruleId: "CF-001", tags: ["guard clause", "early return", "pyramid", "no-nested-if"], component: NestedIf },
  { id: "03-boolean-prefixes", title: "Boolean prefixes", section: "control-flow", severity: "warn", ruleId: "BOOL-001", tags: ["is", "has", "can", "should", "boolean naming"], component: BooleanPrefixes },
  { id: "08-two-operand", title: "Two-operand max", section: "control-flow", severity: "hard", ruleId: "CF-002", tags: ["boolean expression", "extract intent", "readability"], component: TwoOperandMax },
  { id: "09-positive-guards", title: "Positive guards", section: "control-flow", severity: "warn", ruleId: "CF-003", tags: ["positive condition", "negation", "guard"], component: PositiveGuards },

  { id: "04-app-error", title: "AppError wrapper", section: "errors", severity: "hard", ruleId: "ERR-001", tags: ["error handling", "wrap", "context", "apperror"], component: AppErrorWrapper },
  { id: "05-logging", title: "Structured logging", section: "errors", severity: "warn", ruleId: "ERR-002", tags: ["observability", "log", "context", "surface errors"], component: StructuredLogging },
  { id: "06-magic-strings", title: "Magic strings", section: "errors", severity: "hard", ruleId: "ERR-003", tags: ["constants", "enum", "literals"], component: MagicStrings },

  { id: "10-spec-first", title: "Spec-first workflow", section: "workflow", severity: "style", ruleId: "WF-001", tags: ["spec", "process", "planning"], component: SpecFirst },
  { id: "11-cache-invalidation", title: "Cache invalidation", section: "workflow", severity: "warn", ruleId: "WF-002", tags: ["cache", "ttl", "invalidate", "keys"], component: CacheInvalidation },

  { id: "12-closing", title: "Closing", section: "closing", tags: ["wrap-up", "references", "q&a"], component: Closing },
] as const;

/** Group deck entries by section, preserving section order defined in `SECTIONS`. */
export function groupBySection(): ReadonlyArray<{ section: SlideSectionMeta; slides: readonly SlideEntry[] }> {
  return SECTIONS.map((section) => ({
    section,
    slides: DECK.filter((slide) => slide.section === section.id),
  }));
}
