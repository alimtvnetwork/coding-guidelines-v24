# Canonical Folder Structure Specification

> **Target:** `.ai-memory/folder-structure.md`
> **Authority:** Single Source of Truth for `.ai-memory/` Ecosystem & Repository Organization
> **Version:** 5.0.0
> **Last Updated:** 2026-10-27T00:00:00Z

This document defines the authoritative canonical structure of the repository, including the top-level repository tree, the `.ai-memory/` AI metadata layer, and the `02-spec/` specification tree. All AI agents MUST strictly adhere to this architecture when reading, indexing, or writing files.

---

## 1. The `XX-<slug>` Deterministic Prefixing System

To maintain deterministic sorting and avoid ambiguity across operating systems, every file and directory within structured modules (`01-prompts/`, `02-spec/`, `03-ai-scripts/`, `04-code/`, `05-changes-history/`, `06-archive/`, `.ai-memory/plans/`) MUST follow the two-digit zero-padded numeric prefix:

- **Format:** `XX-<slug-name>.ext` or `XX-<slug-name>/`
- **Prefix Range:** Starts at `01-` (or `00-` for root configuration indexes) and increments sequentially (`01-`, `02-`, ...).
- **Lowercase Mandate:** Strictly lowercase ASCII (`[0-9a-z_-]`). No uppercase letters, spaces, or camelCase.
- **Rule:** When creating a new file or subtask, identify the highest existing `XX-` prefix in that directory and increment by 1.

---

## 2. Root Repository Architecture

The repository layout consists of 6 top-level numbered categories, key AI integration environments, and application sources:

```
coding-guidelines/
├── 01-prompts/                          # Master prompt library & categorized execution prompts
│   ├── 01-general/                      # Conversational entry points and base system prompts
│   ├── 02-core-workflow/                # Canonical autonomous workflows and execution protocols
│   ├── 03-read-write/                   # Memory ingestion, persistence, and sync prompts
│   ├── 04-coding-standards/             # Static analysis and guideline compliance prompts
│   ├── 05-coding-guidelines/            # Cross-language refactoring and repair prompts
│   ├── 06-testing-and-qa/               # Autonomous test runners and quality gates
│   ├── 07-bug-fix/                      # 4-part RCA and diagnostic debugging prompts
│   ├── 08-dry-code/                     # DRY analysis, utility extraction, and pattern unification
│   ├── 09-commit-and-multi-agent/       # Commit hygiene and multi-agent coordination
│   ├── 13-plan-audit/                   # Application specification and plan auditing prompts
│   ├── 14-execute/                      # Execution loops (including V6 tree swarm engine)
│   ├── 15-cg-execute/                   # Coding guideline batch refactoring executors
│   ├── 17-ci-cd/                        # CI/CD repair and zero-storage workflow prompts
│   ├── 18-release-management/           # SemVer release blueprints and version bumpers
│   ├── 20-memory-consolidate/           # Institutional memory consolidation prompts
│   ├── 22-letterly/                     # Letterly dictation formatters (mobile, desktop, muse)
│   ├── 23-cursor-prompts/               # Cursor IDE mirrored instruction templates
│   └── 27-muse-prompts/                 # Muse fleet orchestration and master prompts
│
├── 02-spec/                             # Single Source of Truth for system specifications
│   ├── 01-spec-authoring-guide/         # Specification formatting and authoring rules
│   ├── 02-coding-guidelines/            # Cross-language standards and boolean principles
│   ├── 03-error-manage/                 # Structured error handling (*appfault.AppError)
│   ├── 04-database-conventions/         # Schema standards, naming, ORM, and REST APIs
│   ├── 05-split-db-architecture/        # Split SQLite operational and configuration DB pattern
│   ├── 06-seedable-config-architecture/ # Versioned seedable configuration systems
│   ├── 07-design-system/                # Color palettes, typography, spacing, and UI tokens
│   ├── 08-docs-viewer-ui/               # Documentation viewer interface specifications
│   ├── 09-code-block-system/            # Syntax highlighting and interactive code block engine
│   ├── 10-research/                     # Comparative benchmarks and technology evaluations
│   ├── 11-powershell-integration/       # Cross-platform PowerShell automation standards
│   ├── 12-cicd-pipeline-workflows/      # CI/CD pipeline automation and zero-storage workflows
│   ├── 13-generic-cli/                  # Terminal UX, ANSI colors, flags, and help contracts
│   ├── 14-update/                       # CLI self-update architecture and distribution
│   ├── 15-distribution-and-runner/      # Cross-platform installer contracts and runners
│   ├── 16-generic-release/              # Universal release pipeline blueprint and asset matrices
│   ├── 17-consolidated-guidelines/      # AI-readable summaries of all major spec modules
│   ├── 18-wp-plugin-how-to/             # WordPress plugin architecture and REST API patterns
│   ├── 19-main-worker-service/          # Split-tier service architecture and proxy nodes
│   ├── 21-app/                          # App-specific feature specifications and workflows
│   ├── 22-app-issues/                   # App-specific bug analysis, root causes, and fixes
│   ├── 23-app-db/                       # App-specific data models and migration strategies
│   ├── 24-app-ui-design-system/         # App-specific UI tokens and layout patterns
│   └── 25-spec-audits/                  # Specification scorecards and benchmark audits
│
├── 03-ai-scripts/                       # High-speed Python automation toolchain
│   ├── readme.md                        # Master script catalog, usage syntax, and CLI flags
│   ├── 02-shared-engine.py              # Central engine: constants, regex registry, locks, cache
│   ├── 03-file-manipulator.py           # Case conversion, re-sequencing, and UTF-8 LF fixer
│   ├── 04-newline-fixer.py              # Newline normalizer and trailing whitespace stripper
│   ├── 05-guideline-autofixer.py        # Composite guideline runner (newlines + booleans)
│   ├── 06-cicd-local-runner.py          # Parallel local CI runner for 18 quality gates
│   ├── 07-relative-path-fixer.py        # Absolute filesystem path and URI sanitizer
│   ├── 08-naming-autofixer.py           # Lowercase and boolean convention linter
│   ├── 14-version-sync-checker.py       # Version parity guard across manifest files
│   ├── 38-sync-prompts-skills-scripts.py# Canonical asset multi-repo sync engine
│   └── 47-upgrade-cg-execute-to-v6.py   # Automated prompt and skill upgrade engine
│
├── 04-code/                             # Reusable production source code & reference packages
│   └── golang/                          # Enterprise Go packages (pkg/appfault, pkg/streamwriter, pkg/fileutil)
│
├── 05-changes-history/                  # Transaction logs and architectural revision records
│   ├── 01-gitmap-ai-scripts-and-spec-sync/
│   ├── 16-boolean-prefixes-id-standard-and-writer-context/
│   ├── 25-regex-harvesting-dbengine-and-isdefined-standard/
│   └── 26-write-prompts-git-audit-and-recent-tasks-register/
│
├── 06-archive/                          # Superseded prompts, legacy plans, and deprecations
│   ├── v1/                              # Legacy v1 prompts and workflows
│   ├── v3/                              # Legacy v3 prompts and execution templates
│   └── v4/                              # Archived v4 prompts and cg-execute modules
│
├── .ai-memory/                          # Institutional memory, roadmap, and state layer
│   ├── what-to-read.md                  # Authoritative 5-Tier AI reading order
│   ├── folder-structure.md              # Authoritative canonical folder structure specification
│   ├── strictly-avoid.md                # Universal hard prohibitions & CODE RED constraints
│   ├── coding-guidelines.md             # Baseline coding standards summary
│   ├── suggestions.md                   # Pending architectural suggestions & enhancements
│   ├── plans/                           # Micro-task execution center (pending, subtasks, completed)
│   ├── memory/                          # Institutional knowledge base (learned, standards, constraints)
│   ├── release/                         # Release automation, version mutators, and release notes
│   ├── ambiguous-questions/             # Active and resolved ambiguity tracking
│   ├── suggestions/                     # Granular suggestion proposals and history
│   ├── cicd-issues/                     # CI/CD pipeline failure diagnostics and RCAs
│   └── assets/                          # User-provided mockups, diagrams, and reference media
│
├── .agents/                             # Antigravity agent configuration and canonical skills
│   ├── rules/                           # Behavioral rules and coding guidelines
│   ├── scripts/                         # Exact mirror of 03-ai-scripts/ for Antigravity runners
│   └── skills/                          # Reusable skill definitions (SKILL.md)
│
├── .cursor/                             # Cursor IDE rules, prompts, and mirror skills
│   └── skills/                          # Cursor-compatible skill mirrors
│
├── src/                                 # Documentation viewer React web application source
│   ├── components/                      # UI primitives, layout components, and markdown renderers
│   ├── hooks/                           # Custom React hooks (collapsible, search, dashboard)
│   └── pages/                           # Application views (DocsViewer, Dashboard, LandingPage)
│
└── slides-app/                          # Interactive slide deck presentation application
    ├── docs/                            # Visual baseline recipes and architecture guides
    └── tests/                           # Visual regression and snapshot tests
```

---

## 3. Canonical `02-spec/` Specification Hierarchy

The `02-spec/` directory is organized into numbered tiers:

- **Tiers 01–20 (Core Fundamentals):** Reusable standards, architecture blueprints, database conventions, and design systems.
- **Tiers 21+ (App-Specific Modules & Audits):** Application business logic, features, domain-specific schemas, and specification audits.

| Tier | Module Directory | Focus Area |
|:---:|---|---|
| **01** | `02-spec/01-spec-authoring-guide/` | Meta-guide on writing and maintaining specifications |
| **02** | `02-spec/02-coding-guidelines/` | Cross-language standards (TypeScript, Go, PHP, Rust, C#) |
| **03** | `02-spec/03-error-manage/` | Error envelopes, application errors, and error codes |
| **04** | `02-spec/04-database-conventions/` | Database schema conventions, ORM mappings, views |
| **05** | `02-spec/05-split-db-architecture/` | Split database pattern (operational vs. config DBs) |
| **06** | `02-spec/06-seedable-config-architecture/` | Seedable configuration with versioned changelogs |
| **07** | `02-spec/07-design-system/` | Design tokens, component hierarchies, dark mode |
| **08** | `02-spec/08-docs-viewer-ui/` | Documentation viewer UI and slide presentation system |
| **09** | `02-spec/09-code-block-system/` | Syntax highlighting pipeline and code block copy/run |
| **10** | `02-spec/10-research/` | Technology evaluations and benchmark studies |
| **11** | `02-spec/11-powershell-integration/` | Cross-platform PowerShell automation standards |
| **12** | `02-spec/12-cicd-pipeline-workflows/` | CI/CD workflows, GitHub Actions, and quality gates |
| **13** | `02-spec/13-generic-cli/` | Generic CLI UX, exit codes, and flag semantics |
| **14** | `02-spec/14-update/` | CLI self-update architecture and release delivery |
| **15** | `02-spec/15-distribution-and-runner/` | Cross-platform installers, runners, and forwarding |
| **16** | `02-spec/16-generic-release/` | Universal release pipeline blueprint and asset matrices |
| **17** | `02-spec/17-consolidated-guidelines/` | Single-file AI summaries of all major spec modules |
| **18** | `02-spec/18-wp-plugin-how-to/` | WordPress plugin architecture and REST APIs |
| **19** | `02-spec/19-main-worker-service/` | Split-tier worker service and proxy nodes |
| **20** | `02-spec/20-*` *(Reserved)* | Reserved for future core fundamentals |
| **21** | `02-spec/21-app/` | App-specific workflows, features, and domain specifications |
| **22** | `02-spec/22-app-issues/` | App bug analysis, root cause analyses, and fix documentation |
| **23** | `02-spec/23-app-db/` | App-specific data models, table designs, and migration paths |
| **24** | `02-spec/24-app-ui-design-system/` | App-specific UI components, theming tokens, and design patterns |
| **25** | `02-spec/25-spec-audits/` | Specification audits, blind-AI readiness scorecards, and matrices |

---

## 4. Governance & Synchronization Invariants

1. **Strict Relative Paths:** All internal markdown links, citations, and subtask paths MUST use relative paths starting from the repository root (e.g. `.ai-memory/plans/pending/01-task.md`, `02-spec/02-coding-guidelines/readme.md`). Absolute filesystem paths (`C:\...`, `/home/...`, `file:///...`) are strictly prohibited.
2. **Strict Lowercase:** All generated files and directories MUST be lowercase.
3. **Plurality Invariants:** Always `.ai-memory/memory/` (singular), `.ai-memory/plans/` (plural), `01-prompts/` (plural), `.ai-memory/suggestions/` (plural), `03-ai-scripts/` (plural).
4. **Mirror Parity:** Every script in `03-ai-scripts/` MUST have an exact mirror in `.agents/scripts/`.
5. **Quality Gates:** Before concluding any engineering task, execute `python 03-ai-scripts/06-cicd-local-runner.py` to confirm all 18 quality checks pass.
