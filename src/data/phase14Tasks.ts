/**
 * Phase 14 Kickoff — implementation task index.
 *
 * Source of truth: `03-issues/phase-14-kickoff/PH14-*.md`.
 * This file is a typed mirror for the docs-viewer tracker page.
 *
 * Hard rule: spec/19-main-worker-service is markdown-only. The "link"
 * here is a one-way reference (app -> spec). No implementation code lives
 * inside spec/19. See mem://constraints/spec19-no-implementation.
 */

export type Phase14Status = "Todo" | "InProgress" | "Blocked" | "Done";

export interface Phase14Task {
  Id: string;
  Chapter: string;
  Title: string;
  SpecPath: string;
  IssuePath: string;
  Status: Phase14Status;
  Language: string;
  Repo: string;
}

const SPEC_BASE = "spec/19-main-worker-service";
const ISSUE_BASE = "03-issues/phase-14-kickoff";

interface Phase14Seed {
  Chapter: string;
  Slug: string;
  Title: string;
}

const Seeds: Phase14Seed[] = [
  { Chapter: "00", Slug: "overview", Title: "Overview" },
  { Chapter: "01", Slug: "architecture", Title: "Architecture" },
  { Chapter: "02", Slug: "glossary", Title: "Glossary" },
  { Chapter: "03", Slug: "main-db-schema", Title: "Main DB Schema" },
  { Chapter: "04", Slug: "worker-routing", Title: "Worker Routing" },
  { Chapter: "05", Slug: "auth-and-2fa", Title: "Auth and 2FA" },
  { Chapter: "06", Slug: "core-api-endpoints", Title: "Core API Endpoints" },
  { Chapter: "07", Slug: "role-based-dashboards", Title: "Role Based Dashboards" },
  { Chapter: "08", Slug: "error-contract", Title: "Error Contract" },
  { Chapter: "09", Slug: "self-update-pointer", Title: "Self Update Pointer" },
  { Chapter: "10", Slug: "worker-bootstrap-protocol", Title: "Worker Bootstrap Protocol" },
  { Chapter: "11", Slug: "split-db-tier-reconciliation", Title: "Split DB Tier Reconciliation" },
  { Chapter: "12", Slug: "jwt-delivery-contract", Title: "JWT Delivery Contract" },
  { Chapter: "13", Slug: "error-codes", Title: "Error Codes" },
  { Chapter: "14", Slug: "rbac-and-status-seed", Title: "RBAC and Status Seed" },
  { Chapter: "15", Slug: "tunable-constants", Title: "Tunable Constants" },
  { Chapter: "16", Slug: "update-channels", Title: "Update Channels" },
  { Chapter: "17", Slug: "cascading-roles-and-cache-bin", Title: "Cascading Roles and Cache Bin" },
  { Chapter: "24", Slug: "threat-model", Title: "Threat Model" },
  { Chapter: "25", Slug: "inherited-rules", Title: "Inherited Rules" },
];

function buildTask(seed: Phase14Seed): Phase14Task {
  return {
    Id: `PH14-${seed.Chapter}`,
    Chapter: seed.Chapter,
    Title: seed.Title,
    SpecPath: `${SPEC_BASE}/${seed.Chapter}-${seed.Slug}.md`,
    IssuePath: `${ISSUE_BASE}/PH14-${seed.Chapter}-${seed.Slug}.md`,
    Status: "Todo",
    Language: "LanguageTBD",
    Repo: "RepoTBD",
  };
}

export const Phase14Tasks: Phase14Task[] = Seeds.map(buildTask);

export const Phase14Meta = {
  Phase: 14,
  SurfaceLabel: "spec/19 non-backup (chapters 00–17, 24–25)",
  BackupExcluded: "Chapters 18–23 (backup tier) tracked separately.",
  MutationGate: ">=80%",
  ConstraintRef: "mem://constraints/spec19-no-implementation",
} as const;
