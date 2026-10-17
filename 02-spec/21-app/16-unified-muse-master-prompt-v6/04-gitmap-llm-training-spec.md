# Specification: GitMap LLM Training & AI Skills Update Architecture (6-Phase Lifecycle)

> Category: `21-app`  
> Status: Active  
> Module: `16-unified-muse-master-prompt-v6`  
> Target: `02-spec/21-app/16-unified-muse-master-prompt-v6/04-gitmap-llm-training-spec.md`  
> Reference: `02-spec/21-app/16-unified-muse-master-prompt-v6/01-overview-and-architecture.md`  
> Engine: GitMap Autonomous Developer Companion (v28+)  

---

## 1. Executive Summary & Purpose

Modern autonomous AI agents (such as Muse, Antigravity, and Cursor) require deterministic, zero-hallucination onboarding and real-time skill acquisition when operating across complex multi-repository fleets. Previously, AI agent onboarding relied on a legacy 4-stage sequential curriculum (Discovery, Refactoring, Verification, Semantic Commit). While effective for basic operations, this model lacked automated git-state hygiene, content-level policy enforcement, persistent cross-session memory re-sync, and non-interactive remediation.

This specification formalizes the **GitMap LLM Training & AI Skills Update Architecture**, evolving the engine into a unified **6-Phase AI Agent Lifecycle**:
1. **Discovery:** Sub-millisecond context gathering and symbol search.
2. **Refactoring:** Policy-grounded, surgical code transformations.
3. **Verification:** Fast local AST syntax and formatting verification.
4. **Semantic Commit:** Hyphen-separated conventional commits with immediate branch push.
5. **Telemetry:** Non-blocking CI/CD workflow observation, dynamic ETA calculation, and RCA log extraction.
6. **Heal & Fix:** In-process git-state remediation (`cmdfix`) and parallel content-fix auditing (`cmdautofix`) with optional recipe application (`--heal-apply`).

Furthermore, this specification defines the `gitmap llm train <ai-name>` command family (specifically `gitmap llm train muse --days 7`), detailing memory digest ingestion, chat-seed reconstitution, flag matrix parameters, CLI aliases, and the exact documentation required to update `.agents/skills/gitmap/SKILL.md` and `.cursor/skills/gitmap/SKILL.md`.

---

## 2. Command Specification: `gitmap llm train <ai-name>`

### 2.1. Overview & Syntax

The command `gitmap llm train` trains AI models on GitMap CLI capabilities, generates or refreshes native skill definitions on disk, and synchronizes memory digests for named autonomous agents:

```bash
gitmap llm train [ai-name] [flags]
```

When an `<ai-name>` positional argument is provided, the command activates specialized training routines for the specified agent.

### 2.2. The `muse` AI Target (`gitmap llm train muse --days 7`)

The primary first-class AI target is `muse`. Autonomous Muse sessions invoke this command during Phase 8 (Muse Memory IO) to re-synchronize working memory and project context:

```bash
gitmap llm train muse [--days N]
```

#### Execution Sequence & Operational Mechanics:
1. **Primary Vault Probe (`repo-cache`):** GitMap checks `repo-cache` (`~/workspace/repos/repo-cache` or active workspace cache) for the `12-muse-ai/` (or `XX-muse-ai/`) directory.
2. **Fallback Vault Probe (`repo-secrets`):** If not found in `repo-cache`, GitMap falls back to inspecting `repo-secrets/12-muse-ai/`.
3. **Windowed Digest Diffing (`--days N`):** Calculates changes and activity across the preceding `N` days (default: 7 days).
4. **Digest Report Generation:** Outputs a structured terminal summary containing:
   - Total files tracked within the memory window.
   - Chronological commit and edit summaries.
   - Status of sequenced digest files (`01-overview.md` through `05-conventions.md`).
5. **Raw AI Traversal Checklist:** Emits an explicit, file-by-file path listing paired with targeted `gitmap cat <filepath>` and `gitmap aum search` command hints. This enables blind LLMs to walk the memory directory systematically without shell traversal commands.
6. **Chat-Seed Reconstruction:** Scans `12-muse-ai/chat-seeds/*.md` to reconstitute active side conversations for repository-specific contexts (`<repo> repo` side chats).
7. **Missing Digest Tolerance:** If no memory digest directory exists, the command terminates gracefully with exit code 0, emitting actionable bootstrapping instructions on how to initialize `12-muse-ai/` in `repo-cache`.

---

## 3. The 6-Phase AI Agent Lifecycle

The GitMap execution curriculum evolves from the legacy 4-stage model to the modern 6-Phase Lifecycle:

```
┌────────────────────────────────────────────────────────────────────────┐
│                    6-PHASE AI AGENT LIFECYCLE                          │
├───────────────┬────────────────────────────────────────────────────────┤
│ 1. Discovery  │ Sub-ms indexed glob, symbol search, DH2D SQLite cache  │
├───────────────┼────────────────────────────────────────────────────────┤
│ 2. Refactor   │ Surgical AST edits, affirmative booleans, AppError     │
├───────────────┼────────────────────────────────────────────────────────┤
│ 3. Verify     │ AST syntax checks, formatting, zero build/test ban     │
├───────────────┼────────────────────────────────────────────────────────┤
│ 4. Commit     │ Hyphen format (gitmap cpf/cpb/cpr), immediate push    │
├───────────────┼────────────────────────────────────────────────────────┤
│ 5. Telemetry  │ Dynamic ETA waiting (gitmap pe -t), 4-part RCA logs    │
├───────────────┼────────────────────────────────────────────────────────┤
│ 6. Heal & Fix │ 6a Git-state heal (cmdfix) + 6b Content audit (autofix)│
└───────────────┴────────────────────────────────────────────────────────┘
```

### 3.1. Phase 1: Discovery

- **Objective:** Rapid context acquisition and symbol discovery without disk traversal or terminal buffer bloat.
- **Commands:**
  - `gitmap find "<pattern>" [-ext <ext>]`: Sub-millisecond glob discovery across 10,000+ files (<10ms).
  - `gitmap find-files <name>` (`gitmap ff`), `gitmap find-files-any <str>` (`gitmap ffa`), `gitmap find-files-startswith <prefix>` (`gitmap ffs`), `gitmap find-files-endswith <suffix>` (`gitmap ffe`).
  - `gitmap list-files [dir]` (`gitmap lf`): Scoped relative file listing.
  - `gitmap search "<query>"`: Instant SQLite indexed symbol & keyword search via DH2D hot cache (<4 KB RAM overhead, 830,000x faster than Python regex grep).
  - `gitmap aum search "<pattern>" [dir] [--ext <ext>]`: Streaming multi-core live search with lazy regex and null-byte binary probe.
  - `gitmap cat <filepath>`: Zero-disk direct stream of file content into process stdout.
  - `gitmap aum locate <tool>`: Instant toolchain locator (<15ms; locates `python`, `vcvarsall.bat`, `msbuild`).
- **Invariants:**
  - Strict 500 KB file limit (Rule R19).
  - Binary null-byte probe automatically skips compiled artifacts.
  - TOTAL BAN on `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `dir /s`, and `Get-ChildItem -Recurse`.

### 3.2. Phase 2: Refactoring

- **Objective:** Apply surgical, policy-grounded code transformations preserving business logic.
- **Commands:**
  - `gitmap replace <old> <new>`: Exact literal string replacement with automated audit logging.
  - `gitmap replace-regex <pattern> <replacement>`: Scoped regex refactoring across target directories.
  - `gitmap aum newlines --fix`: Polyglot CRLF-to-LF conversion and trailing whitespace stripper.
  - `gitmap lowercase` (`gitmap lcf`): Safe 2-step case normalization for git-tracked files.
- **Invariants:**
  - Affirmative booleans only (`isReady`, `hasCache`; TOTAL BAN on negative names like `isNotReady` or `disableCache`).
  - No explicit evaluation against `true` (`if isReady == true` is BANNED; use `if isReady`).
  - No mixed polarity in a single condition (`if isReady && !hasCache` is BANNED).
  - Maximum 8–15 lines per function; functions exceeding 15 lines must be decomposed.
  - Return `*appfault.AppError` for structured errors in Go; no loose multi-value error tuples.
  - Mandatory vertical whitespace gaps before `if`, after `}`, before `return`.

### 3.3. Phase 3: Verification

- **Objective:** Rapid syntax, formatting, and AST sanity validation without triggering resource-intensive builds.
- **Commands:**
  - `python linter-scripts/check-coding-guidelines.py`: AST guideline verification.
  - `python go-format-check.py` or `gofmt -l`: Formatting and syntax validation.
- **Invariants:**
  - Rule R1 Zero Build/Test Ban: Never invoke `go test`, `npm test`, `cargo test`, `go build`, or `npm run build` during rapid local micro-batching.
  - All heavy compilation and integration test suites are deferred to the CI/CD pipeline.

### 3.4. Phase 4: Semantic Commit

- **Objective:** Stage, commit, and push changes to remote tracking branches in a single atomic operation.
- **Commands:**
  - `gitmap cpf "<module> - <summary>"`: Feature commit (auto-prefixes `Feature: `).
  - `gitmap cpb "<module> - <summary>"`: Bugfix commit (auto-prefixes `Bug: `).
  - `gitmap cpr "<module> - <summary>"`: Release/chore commit.
  - `gitmap pcp "<module> - <summary>"`: Pre-flight pull, commit, and push.
- **Invariants:**
  - TOTAL BAN on colons inside commit arguments (e.g. `gitmap cpf "Feature: add auth"` is FORBIDDEN; use `gitmap cpf "auth - add token support"`).
  - Pre-commit secrets check: Run `check-forbidden-strings.py` or `gitmap aum search` to prevent accidental credential commits. Offload secrets to `repo-secrets` via `gitmap rs`.

### 3.5. Phase 5: Telemetry

- **Objective:** Non-blocking CI/CD workflow observation, dynamic ETA waiting, and Root Cause Analysis (RCA) log extraction.
- **Commands:**
  - `gitmap pipeline-ai status --json`: Interrogates GitHub Actions workflow status, active commit SHA, and ETA.
  - `gitmap pipeline-ai status -t <eta>` (or `gitmap pe -t`): Dynamically sleeps until the calculated ETA expires, eliminating CPU-spinning poll loops.
  - `gitmap pipeline error-logs` (`gitmap pe`): Extracts failing step logs directly to local file for 4-part RCA.
  - `gitmap pe history-ai [N]`: Dumps historical CI failure dossiers to prevent repeating regressions.
  - `gitmap pipeline purge`: Maintains zero-storage footprint by removing stale workflow artifacts.
- **Invariants:**
  - Total ban on tight polling loops (`while true; sleep 5`).
  - Total ban on `actions/upload-artifact` in workflows (maintains 0.0 GB Actions storage quota).

### 3.6. Phase 6: Heal & Fix

- **Objective:** Multi-repository git-state self-healing and parallel content policy compliance running entirely in-process.
- **Sub-Phase 6a: HEAL (Git-State Remediation — `cmdfix`):**
  - Scans workspace repositories for abnormal git states across four categories:
    1. `dirty-worktree`: Uncommitted modifications or staged changes.
    2. `diverged`: Local branch diverged from upstream tracking branch.
    3. `untracked`: Stale untracked scratch files polluting the worktree.
    4. `other`: Detached HEAD, active rebases, or merge conflict states.
  - Reports remediation summary via `gitmap stash ls`.
  - In default mode (report-only), outputs status counts and recommended remediation commands.
  - With `--heal-apply <stash|wip|discard>`, automatically executes the specified recipe non-interactively across dirty repositories.
- **Sub-Phase 6b: FIX (Parallel Content Audit — `cmdautofix`):**
  - Multi-threaded scanning across all repository files (typically <60ms for 3,500+ files).
  - Audits five rule categories:
    1. `markdown`: Sequence numbering gaps, invalid H1 title headings, broken relative links.
    2. `misspell`: Domain-specific typos and common identifier misspellings.
    3. `naming`: Strictly lowercase filename enforcement.
    4. `newlines`: Polyglot CRLF vs LF detection, trailing newlines, and trailing whitespace.
    5. `paths`: Absolute filesystem paths (`C:\...`, `/home/...`) or `file:///` URIs.
  - In default mode, reports flagged violation counts per category.
  - Fixes are applied automatically when running `gitmap fix all -y`.
- **Heal & Fix Roll-Up:** Emits a final consolidated status table certifying workspace health before session conclusion or handoff.

---

## 4. Supported Flags & Options Matrix

The `gitmap llm train` command accepts the following flags:

| Flag | Type | Default | Description & Operational Behavior |
| :--- | :--- | :--- | :--- |
| `--days <N>` | `int` | `7` | Sets the memory digest inspection window in days for agent targets (e.g. `gitmap llm train muse --days 7`). |
| `--heal-apply <recipe>` | `string` | `""` | Non-interactively applies git-state remediation recipes during Phase 6a. Accepted values: `stash`, `wip`, `discard`. Default is empty (report-only). |
| `--loop` | `bool` | `false` | Triggers the continuous autonomous 6-phase AI self-looping execution cycle. |
| `--self-loop <N>` | `int` | `0` | Specifies the exact number of consecutive iterations for the AI self-loop. |
| `--skill-path <path>` | `string` | `.agents/skills/gitmap/SKILL.md` | Target filesystem path for writing the generated Antigravity skill markdown file. |
| `--text-only` | `bool` | `false` | Emits the complete training curriculum text to stdout without writing files to disk. Useful for read-only environments. |
| `--url` | `bool` | `false` | Outputs the raw public URL to the authoritative `llm.md` specification (`https://raw.githubusercontent.com/alimtvnetwork/gitmap-v28/main/llm.md`). |
| `--urls` | `bool` | `false` | Outputs the complete 7-document Authoritative Public Documentation Registry for LLM memory ingestion. |
| `-j`, `--json` | `bool` | `false` | Emits structured, machine-readable JSON containing author, sponsor, efficiency factors, phase breakdown, and document links. |

---

## 5. Aliases & Invocation Shortcuts

To support diverse shell environments, muscle memory, and concise agent scripting, GitMap exposes the following first-class command aliases:

| Command Invocations | Type | Description |
| :--- | :--- | :--- |
| `gitmap llm train` | Canonical | Full 6-phase chained curriculum execution and skill generation. |
| `gitmap train` | Primary Alias | Shortest invocation alias for `gitmap llm train`. |
| `gitmap llm-train` | Hyphenated Alias | Single-token hyphenated form of `gitmap llm train`. |
| `gitmap train-muse` | Direct Target Alias | Direct shortcut for `gitmap llm train muse`. |
| `gitmap llm-train-muse` | Hyphenated Target Alias | Direct hyphenated shortcut for `gitmap llm train muse`. |
| `gitmap llm chain` | Chained Legacy Alias | Backward-compatible alias representing the chained curriculum flow. |

All aliases support the identical flag set (`--days`, `--loop`, `--self-loop`, `--heal-apply`, `--text-only`, `--urls`, `--url`, `--json`, `--skill-path`).

---

## 6. Authoritative Public Documentation Registry (7 Mandatory Links)

When invoked with `--urls` (or inspected via `--json`), GitMap outputs the 7 mandatory authoritative documents required for LLM memory ingestion:

| # | Document Title | Category | Canonical Public URL | Description |
| :---: | :--- | :--- | :--- | :--- |
| 1 | Core LLM Specification & AI Guidelines | Specification | `https://raw.githubusercontent.com/alimtvnetwork/gitmap-v28/main/llm.md` | Authoritative 6-phase execution lifecycle, CLI tool reference, and pair programming standards. |
| 2 | GitMap Native Antigravity Skill | Skill | `https://raw.githubusercontent.com/alimtvnetwork/gitmap-v28/main/.agents/skills/gitmap/SKILL.md` | Comprehensive cheatsheet, regex search rules, replacement matrices, and toolchain commands. |
| 3 | What-To-Read Memory Manifest | Memory | `https://raw.githubusercontent.com/alimtvnetwork/gitmap-v28/main/what-to-read.md` | Canonical reading order, code-red principles, active specifications, and persistent ledger. |
| 4 | AI Memory Ledger & Issue History | Memory | `https://raw.githubusercontent.com/alimtvnetwork/gitmap-v28/main/.ai-memory/what-to-read.md` | Historical context, resolved issues, regression prevention dossiers, and plan indices. |
| 5 | Anti-Pattern Command Replacement Matrix | Guideline | `https://raw.githubusercontent.com/alimtvnetwork/gitmap-v28/main/02-spec/02-coding-guidelines/06-ai-optimization/10-anti-pattern-replacements.md` | Mandatory substitutions banning Select-String, Get-ChildItem -Filter, git grep, and findstr. |
| 6 | Anti-Hallucination Rules & Verification | Guideline | `https://raw.githubusercontent.com/alimtvnetwork/gitmap-v28/main/02-spec/02-coding-guidelines/06-ai-optimization/02-anti-hallucination-rules.md` | Strict zero-assumption policy, real ground truth verification, and relative path hygiene. |
| 7 | Common AI Mistakes & Guardrails | Guideline | `https://raw.githubusercontent.com/alimtvnetwork/gitmap-v28/main/02-spec/02-coding-guidelines/06-ai-optimization/04-common-ai-mistakes.md` | Catalog of recurring AI failure modes and compiler/runtime crash prevention guardrails. |

---

## 7. Exact Documentation Required for GitMap Skill Files

Both `.agents/skills/gitmap/SKILL.md` and `.cursor/skills/gitmap/SKILL.md` must accurately document the 6-phase lifecycle, supported flags, aliases, and training targets. The exact markdown text required for Section 8 and Operational Invariant 1 is defined below.

### 7.1. Section 8: Autonomous Agent Onboarding & Curriculum (LLM)

The following block defines the exact contents required in **Section 8** of both skill files:

```markdown
### 8. Autonomous Agent Onboarding & Curriculum (LLM)
- `gitmap llm train [ai-name] [flags]` (aliases: `gitmap train`, `gitmap llm-train`, `gitmap llm chain`) — Full 6-phase chained onboarding curriculum (Discovery, Refactoring, Verification, Semantic Commit, Telemetry, Heal & Fix), auto-generates Antigravity skill, author/sponsor attribution.
- `gitmap llm train muse [--days N]` (aliases: `gitmap train-muse`, `gitmap llm-train-muse`) — Syncs the Muse 7-day memory digest from `repo-cache` (fallback: `repo-secrets`), diffs last N days, prints raw AI traversal checklist, and reconstitutes `chat-seeds`.
- `gitmap llm train --loop` / `gitmap llm train --self-loop <N>` — Executes continuous autonomous 6-phase AI self-looping execution cycles.
- `gitmap llm train --heal-apply <stash|wip|discard>` — Non-interactively applies git-state remediation recipes during Phase 6a.
- `gitmap llm train --text-only` — Outputs curriculum text to stdout without writing skill file to disk.
- `gitmap llm train --urls` — Outputs the 7 authoritative public documentation registry links for LLM ingestion.
- `gitmap llm train --url` — Outputs raw public URL to `llm.md` specification.
- `gitmap llm train -j` / `gitmap llm train --json` — Emits structured machine-readable JSON curriculum and document metadata.
- `gitmap llm-docs` (alias: `gitmap ld`) — Consolidated markdown command matrix reference for LLMs.
- `gitmap llm` — Displays full LLM specification and operational guidelines.
```

### 7.2. Operational Guardrails & Non-Negotiable Invariants — Item 1

The following block defines the exact replacement required for **Operational Invariant 1** in both skill files:

```markdown
1. **Learning & Skill Acquisition:** Run `gitmap llm train` (or `gitmap train`) to initialize or update GitMap skills across the 6-phase AI lifecycle (Discovery, Refactoring, Verification, Semantic Commit, Telemetry, Heal & Fix), or `gitmap train muse --days 7` to sync memory digests. Never run broad keyword searches like `gitmap aum search "train"` to discover how commands work.
```

---

## 8. Verification & Acceptance Criteria

### AC-TRAIN-001: 6-Phase Lifecycle Formalization
- **Given** An AI agent inspects GitMap training specifications.
- **When** Reviewing the lifecycle phases.
- **Then** The lifecycle explicitly contains 6 distinct phases: (1) Discovery, (2) Refactoring, (3) Verification, (4) Semantic Commit, (5) Telemetry, and (6) Heal & Fix, fully superseding the legacy 4-stage curriculum.

### AC-TRAIN-002: Dedicated Heal & Fix Sub-Phases
- **Given** Phase 6 execution details.
- **When** Evaluating git-state and content-fix responsibilities.
- **Then** Phase 6a is defined as workspace git-state remediation (`cmdfix` / `gitmap stash ls`) with `--heal-apply` recipe automation, and Phase 6b is defined as multi-threaded content auditing (`cmdautofix` / `gitmap fix all -y`).

### AC-TRAIN-003: Muse Target Specification
- **Given** The command `gitmap llm train muse --days 7` (or aliases `gitmap train-muse`, `gitmap llm-train-muse`).
- **When** Running against workspace storage.
- **Then** The spec details the primary probe to `repo-cache` for `12-muse-ai/`, fallback probe to `repo-secrets`, window calculation across N days, emission of the raw AI traversal checklist, chat-seed reconstitution, and graceful handling of missing directories.

### AC-TRAIN-004: Supported Flags Matrix
- **Given** The flags table in Section 4.
- **When** Checking against CLI capabilities.
- **Then** The flags `--days <N>`, `--loop`, `--self-loop <N>`, `--heal-apply <stash|wip|discard>`, `--skill-path`, `--text-only`, `--url`, `--urls`, and `-j/--json` are all comprehensively documented with types, defaults, and operational semantics.

### AC-TRAIN-005: Aliases Inventory
- **Given** CLI invocation alternatives.
- **When** Checking supported shortcuts.
- **Then** The aliases `gitmap train`, `gitmap llm-train`, `gitmap train-muse`, `gitmap llm-train-muse`, and `gitmap llm chain` are cataloged.

### AC-TRAIN-006: Skill Documentation Text
- **Given** Documentation requirements for `.agents/skills/gitmap/SKILL.md` and `.cursor/skills/gitmap/SKILL.md`.
- **When** Reviewing Section 7 of this specification.
- **Then** Exact, verbatim markdown replacement blocks are provided for Section 8 (Autonomous Agent Onboarding & Curriculum) and Operational Invariant 1 (Learning & Skill Acquisition).

### AC-TRAIN-007: Authoritative Registry Verification
- **Given** The authoritative public documentation matrix.
- **When** Reviewing URLs emitted by `--urls` and `--json`.
- **Then** All 7 mandatory documents (`llm.md`, `SKILL.md`, `what-to-read.md`, `.ai-memory/what-to-read.md`, `10-anti-pattern-replacements.md`, `02-anti-hallucination-rules.md`, `04-common-ai-mistakes.md`) are documented with their categories and descriptions.

### AC-TRAIN-008: Strict Relative Git Paths & Lowercase Hygiene
- **Given** The generated specification file `02-spec/21-app/16-unified-muse-master-prompt-v6/04-gitmap-llm-training-spec.md`.
- **When** Running path and naming linters.
- **Then** Zero absolute filesystem paths or `file:///` URIs exist, all repository paths are strictly relative to git root, and all filenames are strictly lowercase.
