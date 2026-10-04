# Component Specification: GitMap AI Training Engine, Core Feature Sequence & Execute Prompts Archival

> **Spec ID:** `02-spec/21-app/11-gitmap-training-and-letterly-golden-prompt-sync/02-component-spec.md`  
> **Parent Architecture Spec:** `02-spec/21-app/11-gitmap-training-and-letterly-golden-prompt-sync/01-architecture-spec.md`  
> **Status:** APPROVED  
> **Author:** Spec Subagent 02  
> **Target Subsystems:**
> - `01-prompts/26-gitmap/01-gitmap-core-engine.md` (New Category 26)
> - `01-prompts/26-gitmap/readme.md` (Category Catalog)
> - `.agents/skills/gitmap/skill.md` (Antigravity Native Skill)
> - `.cursor/skills/gitmap/skill.md` (Cursor IDE Native Skill)
> - `01-prompts/14-execute/` (Resequenced V6-only execution category)
> - `01-prompts/19-old-execute-prompts/` (Archival tier for legacy execute workflows)
> - `.ai-memory/prompts.md` & `01-prompts/readme.md` (Master Prompt Registries)  
> **Companion Subtask Plan:** `.ai-memory/plans/subtasks/11-gitmap-training-and-letterly-golden-prompt-sync/02-gitmap-and-execute.md`

---

## 1. Executive Summary & Component Objectives

This component specification establishes the technical implementation details for three major system capabilities within the Prompt Architect meta-repository:

1. **Category 26 (`01-prompts/26-gitmap/`):** Creation of a dedicated prompt category for GitMap AI training, CLI engine integration, and autonomous developer workflows.
2. **Comprehensive GitMap Feature Sequence & Skill Synchronization:** Exhaustive architectural specification of all 9 GitMap core capabilities (streaming regex search, DH2D SQLite symbol search, file finding, listing, zero-disk file streaming, on-the-fly script execution, script offloading to `repo-cache`, Python toolchain location caching, LLM training curriculum, and atomic hyphen-separated commits). Updates both `.agents/skills/gitmap/skill.md` and `.cursor/skills/gitmap/skill.md`.
3. **Execute Prompts Archival & Resequencing:** Complete segregation of legacy execution prompts into `01-prompts/19-old-execute-prompts/` (slots 03 and 04), leaving strictly the Canonical V6 implementation in `01-prompts/14-execute/`, followed by clean 01-to-07 contiguous sequence normalization across all files and registers.

---

## 2. Component 1: Category 26 (`01-prompts/26-gitmap/`)

### 2.1. Category Placement & Taxonomy
The new category `26-gitmap` is situated directly under `01-prompts/`:

```text
01-prompts/
├── 25-ai-verification/
└── 26-gitmap/
    ├── 01-gitmap-core-engine.md
    └── readme.md
```

Category 26 expands the canonical prompt library from 26 categories (00 through 25) to 27 categories (00 through 26).

### 2.2. Specifications for `01-gitmap-core-engine.md`

#### Purpose & Header
`01-prompts/26-gitmap/01-gitmap-core-engine.md` serves as the primary AI training prompt to educate coding assistants, subagents, and orchestrators on the full scope of GitMap capabilities.

#### Mandatory Content Structure
1. **Interactive Links:** Starts with slash command shortcuts `[/goal](slashCommand;goal)` and `[/learn](slashCommand;learn)` linking directly to `.ai-memory/`.
2. **Top-Instruction Priority Mandate:** Reasserts that user preamble instructions take highest priority.
3. **Core AI Directives:**
   - Explicitly instructs the AI to leverage GitMap CLI commands rather than standard, slow shell commands.
   - Forbids shell search tools (`Select-String`, `Get-ChildItem -Recurse`, `rg`, `ripgrep`, `grep`, `git grep`, `findstr`).
   - Requires caching on-the-fly Python and PowerShell automation scripts via `gitmap rc` for future reuse across repositories.
   - Specifies toolchain location caching (`gitmap aum locate python`) so the environment avoids expensive PATH lookups.
4. **Structured Feature Sequence:** Detailed in Section 3 of this specification.
5. **Atomic Commit Protocol:** Strictly defines `gitmap cpf "<module> - <summary>"` and `gitmap cpb "<module> - <summary>"` where arguments omit leading colons.

### 2.3. Specifications for `01-prompts/26-gitmap/readme.md`

`01-prompts/26-gitmap/readme.md` provides the directory index and catalog:
- Category identification metadata.
- Overview of the GitMap Autonomous Developer Companion CLI.
- Table of Prompts containing `01-gitmap-core-engine.md`.
- Summary of core capabilities with quick commands.

---

## 3. Component 2: GitMap Comprehensive Feature Sequence & AI Training Specification

Any AI agent interacting with GitMap must master 9 core capabilities. These features are sequenced systematically below:

### 3.1. Feature 1: Multi-Core Streaming Regex Search (`gitmap aum search`)
- **CLI Invocations:**
  - Live search: `gitmap aum search "<pattern>" [dir] [-e <.ext>] [-r] [-i]`
  - Alias: `gitmap aum grep`
- **Technical Mechanism:** Multi-threaded parallel disk scanner with lazy regular expression compilation, streaming results directly to stdout as they match.
- **Auto-Reject Enforcements:**
  - Mandatory target directory scoping: AI agents must pass `[dir]` (e.g. `cli`, `02-spec`, `pkg`) whenever possible.
  - Mandatory extension scoping: AI agents must pass `-e <.ext>` (e.g. `-e .go`, `-e .ts`, `-e .md`) to avoid reading irrelevant assets.
  - TOTAL BAN on shell search utilities: `Select-String`, `rg`, `ripgrep`, `grep`, `git grep`, `Get-ChildItem -Recurse`, `findstr`.

### 3.2. Feature 2: Indexed Global Symbol Search (`gitmap search`)
- **CLI Invocations:**
  - Symbol query: `gitmap search "<query>" [--limit <n>]`
- **Technical Mechanism:** Sub-millisecond SQLite hot cache query using the DH2D split-db index.
- **Usage Recommendation:** AI agents should query `gitmap search` first for known symbols, function names, types, and constants before initiating full disk scans.

### 3.3. Feature 3: Rapid File Finding (`gitmap find` / `gitmap ff` / `gitmap ffa`)
- **CLI Invocations:**
  - Wildcard pattern matching: `gitmap find "<pattern>" [-ext <ext>]`
  - Exact file match: `gitmap find-files <name>` (alias: `gitmap ff <name>`)
  - Substring file match: `gitmap find-files-any <str>` (alias: `gitmap ffa <str>`)
  - Prefix file match: `gitmap find-files-startswith <prefix>` (alias: `gitmap ffs <prefix>`)
  - Suffix file match: `gitmap find-files-endswith <suffix>` (alias: `gitmap ffe <suffix>`)
- **Performance Characteristics:** Resolves paths in under 10 milliseconds across codebases with more than 10,000 files.

### 3.4. Feature 4: File Inventory & Listing (`gitmap lf` / `gitmap list-files`)
- **CLI Invocations:**
  - File inventory: `gitmap list-files [pattern] [-ext <ext>]`
  - Shorthand alias: `gitmap lf [pat]`
- **Technical Mechanism:** Emits relative file paths matching the directory pattern. Replaces noisy directory traversal commands.

### 3.5. Feature 5: Deterministic Terminal File Streaming (`gitmap cat`)
- **CLI Invocations:**
  - Stream file content: `gitmap cat <filepath>`
- **Technical Mechanism:** Streams file contents directly into the process stdout without buffer bloat. Enables quick inspection of file contents in resource-constrained environments.

### 3.6. Feature 6: On-The-Fly Script Execution & Script Offloading (`gitmap py` / `gitmap ps` / `gitmap rc`)
- **CLI Invocations:**
  - Python execution: `gitmap py "<code-or-script>"`
  - PowerShell execution: `gitmap pwsh "<cmd>"` or `gitmap ps "<cmd>"` (runs with `-NoProfile` and automatic fallback)
  - Bash execution: `gitmap bash "<cmd>"` or `gitmap sh "<cmd>"`
  - Script caching: `gitmap rc <file|folder|text>`
    - `gitmap rc file <script.ps1> [--repo <repo-name>]`
    - `gitmap rc folder <folderpath> [--repo <repo-name>]`
    - `gitmap rc text "<script-content>" --slug <slug> --ext <.ps1|.py>`
- **Technical Rationale for AI:**
  - When an AI creates diagnostic scripts, test harnesses, or data extractors on-the-fly, it must not leave them scattered across the project working tree.
  - By using `gitmap rc`, the script is saved to `repo-cache` (`repo-storage`), committed automatically, and made reusable for future tasks and other repositories without repository pollution.

### 3.7. Feature 7: Python Toolchain Locator & Cache Backup (`gitmap aum locate python`)
- **CLI Invocations:**
  - Tool locator: `gitmap aum locate python`
  - Status check: `gitmap aum cache status`
- **Technical Mechanism:** GitMap scans known toolchain paths in under 15 milliseconds, caches the exact executable path in the local SQLite engine (`installation.db`), and maintains a backup locator profile.
- **AI Operational Benefit:** Once located, subsequent script invocations execute directly without PATH traversal overhead or environment ambiguity.

### 3.8. Feature 8: LLM Training Curriculum & Pipeline Diagnostics (`gitmap llm train` / `gitmap ld` / `gitmap pe`)
- **CLI Invocations:**
  - Full curriculum: `gitmap llm train` (alias: `gitmap llm chain`)
  - Text-only curriculum: `gitmap llm train --text-only`
  - Markdown reference matrix: `gitmap llm-docs` (alias: `gitmap ld`)
  - Pipeline error diagnostics: `gitmap pipeline errors` (alias: `gitmap pe`)
  - Pipeline history analysis: `gitmap pe history-ai`
- **AI Operational Benefit:** Enables blind AI agents to onboard immediately onto GitMap capabilities and extract pipeline error traces for 4-part Root Cause Analysis.

### 3.9. Feature 9: Semantic Hyphen-Separated Atomic Commits (`gitmap cpf` / `gitmap cpb` / `gitmap cpr`)
- **CLI Invocations:**
  - Feature commit: `gitmap cpf "<module> - <summary>"`
  - Bugfix commit: `gitmap cpb "<module> - <summary>"`
  - Release / chore commit: `gitmap cpr "<module> - <summary>"`
- **Strict Invariant on Message Formatting:**
  - GitMap automatically adds `Feature: ` or `Bug: ` prefixes to commit messages.
  - AI agents MUST NOT pass colons inside the message argument (e.g. `gitmap cpf "Feature: my summary"` is FORBIDDEN).
  - Format MUST use hyphens: `gitmap cpf "module - my summary"`.

---

## 4. Component 3: GitMap Skills Specification (`.agents/skills/gitmap/` & `.cursor/skills/gitmap/`)

Both native skills `.agents/skills/gitmap/skill.md` and `.cursor/skills/gitmap/skill.md` must be synchronized to reflect the complete 9-feature sequence:

```yaml
---
name: gitmap
description: Autonomous developer companion and CLI for ultra-fast repository scanning, polyglot automation (AUM), cluster/SSH delegation, pipeline self-healing, and coding guideline enforcement.
---
```

### Key Sections Required in Skill Files:
1. **Overview & Authorship:** MD ALIM UL KARIM (alimtvnetwork), sponsored by RISEUP ASIA LLC.
2. **High-Performance Automation (AUM):** Streaming search, indexed symbol search, boundary guard (500 KB limit), sequence fix, newlines fix, tool locator (`gitmap aum locate python`).
3. **Autonomous Agent Onboarding & Curriculum (LLM):** `gitmap llm train`, `gitmap ld`, `gitmap pe history-ai`.
4. **CI/CD Self-Healing (Pipeline AI):** Status, dynamic waiting (`-t <eta>`), error logs extraction, storage purge.
5. **Fast File Discovery & Terminal Streaming:** `gitmap find`, `gitmap ff`, `gitmap ffa`, `gitmap ffs`, `gitmap ffe`, `gitmap cat`, `gitmap lf`.
6. **Script Execution & Cache Offloading:** `gitmap py`, `gitmap ps` / `gitmap pwsh`, `gitmap bash` / `gitmap sh`, `gitmap rc`.
7. **Semantic Commit & Push:** `gitmap cpf`, `gitmap cpb`, `gitmap cpr` with hyphen-separated syntax.
8. **Operational Guardrails:** Mandatory pre-flight pull, scoped search, 500 KB limit, positive booleans, relative paths.

---

## 5. Component 4: Execute Prompt Archival & Resequencing Architecture

### 5.1. Problem & Archival Motivation
Currently, `01-prompts/14-execute/` contains legacy or redundant prompts:
- `02-execute-parent-task-with-n-steps.md` (redundant with V6)
- `09-parent-task-in-below-steps.md` (superseded by modern below-steps patterns)

Additionally, `13-execute-parent-task-with-n-steps-v6.md` is the canonical V6 implementation, but sits at index 13 with gaps in the numbering (missing 06, 08, 10, 11, 12).

The user mandate requires:
- Moving all other execute parent tasks in N steps to the old section (`01-prompts/19-old-execute-prompts/`).
- Keeping ONLY V6 in `14-execute/`.
- Resequencing the remaining prompts in `14-execute/` into a clean, gap-free contiguous sequence.

### 5.2. Archival Mapping to `01-prompts/19-old-execute-prompts/`

`01-prompts/19-old-execute-prompts/` currently contains:
- `01-execute-robust-loop.md`
- `02-fix-subtask-naming-convention.md`

The moved files take slots 03 and 04:
- `01-prompts/14-execute/02-execute-parent-task-with-n-steps.md` -> `01-prompts/19-old-execute-prompts/03-execute-parent-task-with-n-steps.md`
- `01-prompts/14-execute/09-parent-task-in-below-steps.md` -> `01-prompts/19-old-execute-prompts/04-parent-task-in-below-steps.md`

### 5.3. Resequencing Mapping for `01-prompts/14-execute/`

With `02` and `09` moved to `19-old-execute-prompts/`, and `13-execute-parent-task-with-n-steps-v6.md` established as the sole parent orchestrator in slot 02:

| Old Path in `14-execute/` | New Resequenced Path | Role / Scope |
| :--- | :--- | :--- |
| `01-execute-pending-tasks.md` | `01-execute-pending-tasks.md` | Continuous loop executing pending task queue |
| `13-execute-parent-task-with-n-steps-v6.md` | `02-execute-parent-task-with-n-steps-v6.md` | **Canonical V6 Master Orchestrator** (N=300, A=2, H=2, C=30) |
| `03-execute-batched-loop.md` | `03-execute-batched-loop.md` | Batched multi-agent loop with file collision matrix |
| `04-execute-ai-instruction-writer.md` | `04-execute-ai-instruction-writer.md` | Subagent instruction writer and spec generator |
| `05-execute-batched-loop-wor.md` | `05-execute-batched-loop-wor.md` | Batched multi-agent loop without release |
| `07-execute-batched-loop-v2.md` | `06-execute-batched-loop-v2.md` | Streamlined batched loop execution engine V2 |
| `14-run.md` | `07-run.md` | Autonomous project runner script orchestration |

Total active prompts in `01-prompts/14-execute/`: exactly 7 files (`01` through `07`), zero gaps, zero legacy duplicates.

### 5.4. Registries & References Update Matrix

#### 1. `01-prompts/14-execute/readme.md`
- Updates table of files to reflect the 7 resequenced files.
- Highlights `02-execute-parent-task-with-n-steps-v6.md` as the Canonical V6 Master Orchestrator.
- Notes the archival of previous versions into `01-prompts/19-old-execute-prompts/` and `06-archive/execute/`.

#### 2. `01-prompts/19-old-execute-prompts/readme.md`
- Creates or updates the catalog in `01-prompts/19-old-execute-prompts/readme.md` listing files 01 through 04.

#### 3. `.ai-memory/prompts.md`
- Updates entries under category `14-execute` to files 01 through 07.
- Adds entries under category `19-old-execute-prompts` for files 03 and 04.
- Adds entry for new category `26-gitmap`: `01-prompts/26-gitmap/01-gitmap-core-engine.md` and `readme.md`.

#### 4. `01-prompts/readme.md`
- Updates Directory Index to list Category 26: `26-gitmap/`.
- Updates Table of Categories to include Category 26 (`01-gitmap-core-engine.md`).
- Updates Table row for Category 14 (`14-execute/`) to list `02-execute-parent-task-with-n-steps-v6.md` and `07-run.md`.
- Updates Table row for Category 19 (`19-old-execute-prompts/`) to include `03-execute-parent-task-with-n-steps.md` and `04-parent-task-in-below-steps.md`.

#### 5. Skill Files Referencing V6 Execution
- All references across `.agents/skills/` and `.cursor/skills/` to `01-prompts/14-execute/13-execute-parent-task-with-n-steps-v6.md` or `01-prompts/14-execute/02-execute-parent-task-with-n-steps.md` must be updated to reference `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`.

---

## 6. Implementation Invariants & Quality Acceptance Gates

1. **Relative Paths Invariant (TOTAL BAN on Absolute Paths):**
   Every markdown link, file reference, and citation must use strict relative paths from the repository root (e.g. `01-prompts/26-gitmap/01-gitmap-core-engine.md`). Never use `file:///` URIs or filesystem absolute paths (`C:\...`, `/home/...`).
2. **Positive Booleans Invariant:**
   All boolean logic, options, and documentation descriptions must use positive naming (`isReady`, `hasCache`, `isValid`). Explicit equality against `true` or `false` is prohibited.
3. **No Git Commands in Subagents:**
   Subagents must never run raw `git` commands directly. All commit actions are executed through GitMap commands (`gitmap cpf`, `gitmap cpb`).
4. **Registry Parity Gate:**
   Running `python linter-scripts/check-prompts-loaded.py` must pass with zero missing or orphaned prompt files.
