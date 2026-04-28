// AUTO-GENERATED — do not edit by hand.
// Source: spec/01-spec-authoring-guide/17-version-schema.md §6
// Regenerate: npm run gen:role-enum

export const ROLE_VALUES = [
  "PrimaryAuthor",
  "Contributor",
  "Maintainer",
  "Reviewer",
  "Sponsor",
] as const;

export type Role = (typeof ROLE_VALUES)[number];

export const ROLE_DESCRIPTIONS: Record<Role, string> = {
  PrimaryAuthor: "Original creator and primary owner. Exactly one per repo.",
  Contributor: "Made non-trivial contributions (code, spec, docs).",
  Maintainer: "Has merge / release authority on the integration branch.",
  Reviewer: "Has review authority but does not maintain the repo day-to-day.",
  Sponsor: "Funded, hosted, or otherwise enabled the work.",
};

const ROLE_SET: ReadonlySet<string> = new Set(ROLE_VALUES);

export function isRole(value: unknown): value is Role {
  return typeof value === "string" && ROLE_SET.has(value);
}
