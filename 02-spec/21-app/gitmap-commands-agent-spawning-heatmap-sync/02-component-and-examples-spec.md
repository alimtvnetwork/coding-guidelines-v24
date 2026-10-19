# Component & Examples Specification: GitMap Commands Primacy, A=3 Subagent Spawning, & Heatmap Telemetry

## 1. Executive Summary & Purpose

This specification governs the architectural enhancements, prompt upgrades, and skill definitions across the repository for three interrelated operational capabilities:
1. **GitMap High-Speed Command Primacy & Shell Tool Ban:** Eliminating all fallbacks to slow shell search, directory traversal, and unbuffered reading tools (`rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `findstr`, `Get-ChildItem` in all variants, `dir`, `ls`, `cat`, `type`, `Get-Content`) by establishing explicit, copy-paste GitMap commands.
2. **Mandatory $A = 3$ Subagent Spawning Enforcement:** Diagnosing and eliminating historical subagent spawning failures (such as the "or lead" loophole and the absence of a Phase 1 Spec JSON payload), and enforcing a strict single-tool-call array schema where `invoke_subagent` must contain exactly 3 subagent objects.
3. **Cognitive Heatmap Telemetry & Lifecycle Integration:** Operationalizing `gitmap heatmap log`, `gitmap heatmap status`, `gitmap heatmap tree`, and companion `repo-cache` (`rc`) heatmap telemetry throughout Step 0 Preflight, Phase 1 Planning/Spec, Phase 2 Execution, and Phase 3 Verification.

---

## 2. Subagent Spawning Failure Analysis ($A = 3$)

### 2.1 Problem Statement & Empirical Failure Modes
In previous executions of the V6 execution workflow (`execute-parent-task-with-n-steps-v6`), AI agents repeatedly failed to spawn 3 subagents concurrently ($A = 3$), particularly during Phase 1 Specification and Phase 2 Execution. Instead, agents either:
1. Executed all spec writing and planning solo within the root orchestrator turn.
2. Printed descriptive chat text (e.g. `Dispatched Spec Author 01, 02, 03...`) without issuing an actual tool call to `invoke_subagent`.
3. Spawned only 1 subagent instead of 3, violating the $A = 3$ concurrency standard.
4. Attempted sequential single-agent calls rather than a unified batch call.

### 2.2 Root Cause 1: The "or lead" Escape Hatch
In `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` (line 40) and `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` (line 46), the workflow text stated:
```text
2. Spec Step (A = 3 self subagents or lead): Spawn 3 authoring subagents (TypeName: "self") to author modular, disjoint spec files and subtasks (never writing the same file).
```
The inclusion of `"or lead"` introduced a semantic loophole. When evaluating whether to incur the token and orchestration cost of spawning 3 authoring subagents, language models interpreted `"or lead"` as an explicit authorization to perform all specification authoring directly in the root orchestrator without dispatching subagents.

### 2.3 Root Cause 2: Missing Phase 1 Spec JSON Payload Template
Section 7.1 of the V6 prompt provided concrete JSON dispatch payloads for:
- **Phase 1 Planning Payload:** `TypeName: "research"` (3 read-only subagents).
- **Phase 2 Execution Payload:** `TypeName: "self"` (3 worker subagents).

However, **no JSON payload template existed for Phase 1 Spec Authoring**. Because the prompt lacked a concrete, copy-paste JSON payload demonstrating how to spawn 3 `self` authoring subagents with disjoint file assignments, the model defaulted to solo authoring or omitted the tool call entirely.

### 2.4 Root Cause 3: Ambiguous Array Cardinality in Validation Rules
While the header parameter specified `A = 3`, the prompt lacked an explicit structural schema validation rule requiring that every call to `invoke_subagent` must supply an array with `len(Subagents) == 3`. Without this constraint, models assumed partial waves ($A = 1$ or $A = 2$) were permissible under low task volume, breaking the Tree Swarm concurrency model.

### 2.5 The Non-Negotiable $A = 3$ Invariant
To eliminate these failures permanently:
1. **Total Elimination of Solo Escape Hatches:** The phrase `"or lead"` is excised from all prompt and skill definitions.
2. **Strict Single-Call Batching:** `invoke_subagent` accepts an array of subagents (`"Subagents": [...]`). All 3 subagents MUST be declared within that single array in one tool invocation.
3. **Hard Failure Criteria:** Spawning fewer than 3 subagents (or failing to call `invoke_subagent`) is classified as a critical protocol violation and an auto-reject failure.

---

## 3. Complete 3-Agent JSON Payload Architecture (All 3 Phases)

### 3.1 Phase 1 Planning Payload (`TypeName: "research"` — Read-Only Discovery)
Used during Steps 1..100 to discover architecture, existing specifications, call graphs, and fleet targets. All 3 agents run concurrently using read-only tools.

```json
{
  "Subagents": [
    {
      "TypeName": "research",
      "Role": "Research 01: Architecture & Blast Radius Discovery",
      "Model": "inherit",
      "Workspace": "inherit",
      "Prompt": "You are Research 01 for task <nn>-<slug>. You have no prior chat context; this brief is your complete specification.\n\n### Core Objective:\nExplore the codebase, map symbol callers, trace dependencies, and discover relevant source files for: <task title and objectives>.\n\n### Tool Capabilities & Strict Read-Only Boundary:\n- You are a READ-ONLY subagent (`TypeName: 'research'`).\n- You have read tools: `run_command`, `view_file`, `search_web`, `read_url_content`.\n- You DO NOT have `write_to_file` or `replace_file_content`. NEVER attempt to create or edit files.\n\n### Mandatory GitMap Search Primacy & Total Ban on Raw Grep/Shell:\n- Execute ALL searches and symbol discoveries exclusively via GitMap:\n  * Live symbol search: `gitmap aum search \"<symbol>\" [dir] [-e <.ext>]`\n  * Path-scoped search: `gitmap aum search \"<symbol>\" --path <relative-dir>`\n  * Indexed keyword search: `gitmap search \"<query>\"`\n  * File discovery: `gitmap find \"<pattern>\"`\n  * Directory inventory: `gitmap lf [dir]`\n  * View files: `gitmap cat <path>` or native `view_file`\n- TOTAL BAN (AUTO-REJECT FAILURE): NEVER run `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem` (any variant), `dir`, `ls`, or `findstr`.\n- TOTAL BAN ON GIT COMMANDS: NEVER run `git` commands (`git add`, `git commit`, `git status`, `git diff`, etc.).\n\n### Required Findings Report (Send via send_message to Parent):\nWhen your research is complete, send a message to the caller containing:\n1. Key symbol definitions, structures, and entry points.\n2. Call sites and blast radius (files that will be affected by modifications).\n3. Recommended modular file ownership boundaries for worker subtasks.\n4. Verification commands to validate changes."
    },
    {
      "TypeName": "research",
      "Role": "Research 02: Specs & Dependency Mapping",
      "Model": "inherit",
      "Workspace": "inherit",
      "Prompt": "You are Research 02 for task <nn>-<slug>. You have no prior chat context; this brief is your complete specification.\n\n### Core Objective:\nInspect existing specifications under `02-spec/`, coding guidelines under `02-spec/02-coding-guidelines/`, and prior plan tasks for: <task title and objectives>.\n\n### Tool Capabilities & Strict Read-Only Boundary:\n- You are a READ-ONLY subagent (`TypeName: 'research'`).\n- You DO NOT have file authoring tools. NEVER attempt to create or edit files.\n\n### Mandatory GitMap Search Primacy & Total Ban on Raw Grep/Shell:\n- Execute ALL searches exclusively via GitMap:\n  * Live search in specs: `gitmap aum search \"<term>\" 02-spec`\n  * Indexed search: `gitmap search \"<term>\"`\n  * File discovery: `gitmap find \"*.md\"`\n  * View files: `gitmap cat <file>` or native `view_file`\n- TOTAL BAN: NEVER run `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem`, `dir`, `ls`, or `findstr`.\n- TOTAL BAN ON GIT COMMANDS: NEVER run `git` commands.\n\n### Required Findings Report (Send via send_message to Parent):\nSend a message to the caller detailing existing specs, architectural invariants, acceptance criteria, and positive boolean rules."
    },
    {
      "TypeName": "research",
      "Role": "Research 03: Repository Dependencies & Fleet Sync Targets",
      "Model": "inherit",
      "Workspace": "inherit",
      "Prompt": "You are Research 03 for task <nn>-<slug>. You have no prior chat context; this brief is your complete specification.\n\n### Core Objective:\nExplore external dependencies, script pipelines (`03-ai-scripts/`), skills (`.agents/skills/` and `.cursor/skills/`), and fleet synchronization manifests for: <task title and objectives>.\n\n### Tool Capabilities & Strict Read-Only Boundary:\n- You are a READ-ONLY subagent (`TypeName: 'research'`).\n- You DO NOT have file authoring tools. NEVER attempt to create or edit files.\n\n### Mandatory GitMap Search Primacy & Total Ban on Raw Grep/Shell:\n- Execute ALL searches exclusively via GitMap (`gitmap aum search`, `gitmap find`, `gitmap lf`, `gitmap cat`).\n- TOTAL BAN: NEVER run `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem`, `dir`, `ls`, or `findstr`.\n- TOTAL BAN ON GIT COMMANDS: NEVER run `git` commands.\n\n### Required Findings Report (Send via send_message to Parent):\nSend a message to the caller detailing script dependencies, skill catalogs, and cross-repo synchronization targets."
    }
  ]
}
```

### 3.2 Phase 1 Spec Authoring Payload (`TypeName: "self"` — Read/Write Specification Authors)
Used during Phase 1 to author modular specifications and decomposed subtask plans across disjoint file boundaries.

```json
{
  "Subagents": [
    {
      "TypeName": "self",
      "Role": "Spec Author 01: Architecture & Data Contracts",
      "Model": "inherit",
      "Workspace": "inherit",
      "Prompt": "You are Spec Author 01 for task <nn>-<slug>. You have no prior chat context; this brief is your complete specification.\n\n### Core Objective:\nAuthor the core architectural specification and subtask plans for architecture and preflight.\n\n### Owned Files (Strictly Disjoint — No Other Agent Touches These):\n- `02-spec/21-app/<slug>/01-architecture-spec.md`\n- `.ai-memory/plans/subtasks/<slug>/01-subtask-git-sync-and-preflight.md`\n- `.ai-memory/plans/subtasks/<slug>/02-subtask-architecture-and-specs.md`\n\n### Invariants & Guardrails:\n- Strict Relative Git Paths ONLY (`02-spec/...`, `.ai-memory/...`). TOTAL BAN on absolute filesystem paths and `file:///` URIs.\n- Strict lowercase filenames across all files.\n- TOTAL BAN ON GIT COMMANDS: Subagents never run git commands (`git add`, `git commit`, `git push`, `git status`, etc.).\n- Use `write_to_file` to author files. Log cognitive actions via `gitmap heatmap log -t <slug> -a \"WRITE\" -f <path> -r \"<reason>\"`.\n\n### Reporting Contract:\nWhen files are written and verified, report completed files and diffstats via `send_message` to parent."
    },
    {
      "TypeName": "self",
      "Role": "Spec Author 02: Component & Examples Specification",
      "Model": "inherit",
      "Workspace": "inherit",
      "Prompt": "You are Spec Author 02 for task <nn>-<slug>. You have no prior chat context; this brief is your complete specification.\n\n### Core Objective:\nAuthor the detailed component and examples specification and subtask plans for prompts and skills.\n\n### Owned Files (Strictly Disjoint — No Other Agent Touches These):\n- `02-spec/21-app/<slug>/02-component-and-examples-spec.md`\n- `.ai-memory/plans/subtasks/<slug>/03-subtask-execute-v6-prompt-a3-and-examples.md`\n- `.ai-memory/plans/subtasks/<slug>/04-subtask-gitmap-skills-update.md`\n\n### Invariants & Guardrails:\n- Strict Relative Git Paths ONLY (`02-spec/...`, `.ai-memory/...`). TOTAL BAN on absolute filesystem paths and `file:///` URIs.\n- Strict lowercase filenames across all files.\n- TOTAL BAN ON GIT COMMANDS: Subagents never run git commands (`git add`, `git commit`, `git push`, `git status`, etc.).\n- Use `write_to_file` to author files. Log cognitive actions via `gitmap heatmap log -t <slug> -a \"WRITE\" -f <path> -r \"<reason>\"`.\n\n### Reporting Contract:\nWhen files are written and verified, report completed files and diffstats via `send_message` to parent."
    },
    {
      "TypeName": "self",
      "Role": "Spec Author 03: Migration, Verification & Heatmap Specification",
      "Model": "inherit",
      "Workspace": "inherit",
      "Prompt": "You are Spec Author 03 for task <nn>-<slug>. You have no prior chat context; this brief is your complete specification.\n\n### Core Objective:\nAuthor the migration and verification specification and subtask plans for fleet sync and validation.\n\n### Owned Files (Strictly Disjoint — No Other Agent Touches These):\n- `02-spec/21-app/<slug>/03-migration-and-verification-spec.md`\n- `.ai-memory/plans/subtasks/<slug>/05-subtask-heatmap-telemetry-and-agents-md.md`\n- `.ai-memory/plans/subtasks/<slug>/06-subtask-multi-repo-sync-and-atomic-commit.md`\n\n### Invariants & Guardrails:\n- Strict Relative Git Paths ONLY (`02-spec/...`, `.ai-memory/...`). TOTAL BAN on absolute filesystem paths and `file:///` URIs.\n- Strict lowercase filenames across all files.\n- TOTAL BAN ON GIT COMMANDS: Subagents never run git commands (`git add`, `git commit`, `git push`, `git status`, etc.).\n- Use `write_to_file` to author files. Log cognitive actions via `gitmap heatmap log -t <slug> -a \"WRITE\" -f <path> -r \"<reason>\"`.\n\n### Reporting Contract:\nWhen files are written and verified, report completed files and diffstats via `send_message` to parent."
    }
  ]
}
```

### 3.3 Phase 2 Execution Payload (`TypeName: "self"` — Read/Write Workers)
Used during Phase 2 to modify code, prompts, and skills in parallel waves.

```json
{
  "Subagents": [
    {
      "TypeName": "self",
      "Role": "Worker 01: Execute V6 Prompts & Skills Upgrades",
      "Model": "inherit",
      "Workspace": "inherit",
      "Prompt": "You are Worker 01 for task <nn>-<slug>. You have no prior chat context; this brief is your complete specification.\n\n### Owned Files (Exclusive Modification Rights):\n- `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`\n- `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`\n- `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`\n\n### Instructions & Tasks:\nFollow `.ai-memory/plans/subtasks/<slug>/03-subtask-execute-v6-prompt-a3-and-examples.md`. Eliminate the 'or lead' loophole, insert complete Phase 1 Spec 3-agent payload template, enforce exact 3-agent array validation, and integrate operational heatmap telemetry.\n\n### Boundaries:\n- TOTAL BAN ON GIT COMMANDS (no git add, commit, push, checkout).\n- Search code exclusively via GitMap (`gitmap aum search`, `gitmap find`). TOTAL BAN on rg, ripgrep, Select-String, Get-ChildItem, dir, ls, grep.\n- Log actions via `gitmap heatmap log`.\n- Report completion and evidence via `send_message` to parent."
    },
    {
      "TypeName": "self",
      "Role": "Worker 02: GitMap Skills & Command Matrix Upgrades",
      "Model": "inherit",
      "Workspace": "inherit",
      "Prompt": "You are Worker 02 for task <nn>-<slug>. You have no prior chat context; this brief is your complete specification.\n\n### Owned Files (Exclusive Modification Rights):\n- `.agents/skills/gitmap/skill.md`\n- `.cursor/skills/gitmap/skill.md`\n\n### Instructions & Tasks:\nFollow `.ai-memory/plans/subtasks/<slug>/04-subtask-gitmap-skills-update.md`. Update command replacement matrix to prohibit all forms of Get-ChildItem, dir, ls, and unbuffered reading tools. Add copy-paste examples for directory inventory, code search, file globbing, streaming, and cognitive heatmap logging.\n\n### Boundaries:\n- TOTAL BAN ON GIT COMMANDS (no git add, commit, push, checkout).\n- Search code exclusively via GitMap. TOTAL BAN on rg, ripgrep, Select-String, Get-ChildItem, dir, ls, grep.\n- Log actions via `gitmap heatmap log`.\n- Report completion and evidence via `send_message` to parent."
    },
    {
      "TypeName": "self",
      "Role": "Worker 03: AGENTS.md Guidelines & Heatmap Primacy",
      "Model": "inherit",
      "Workspace": "inherit",
      "Prompt": "You are Worker 03 for task <nn>-<slug>. You have no prior chat context; this brief is your complete specification.\n\n### Owned Files (Exclusive Modification Rights):\n- `AGENTS.md`\n\n### Instructions & Tasks:\nFollow `.ai-memory/plans/subtasks/<slug>/05-subtask-heatmap-telemetry-and-agents-md.md`. Update Section 12 to prohibit all forms of Get-ChildItem, dir, and ls. Add Section 15 governing Cognitive Heatmap Telemetry Primacy.\n\n### Boundaries:\n- TOTAL BAN ON GIT COMMANDS (no git add, commit, push, checkout).\n- Search code exclusively via GitMap. TOTAL BAN on rg, ripgrep, Select-String, Get-ChildItem, dir, ls, grep.\n- Log actions via `gitmap heatmap log`.\n- Report completion and evidence via `send_message` to parent."
    }
  ]
}
```

---

## 4. Comprehensive GitMap High-Speed Command Matrix & Examples

### 4.1 Prohibited Legacy Tools vs GitMap Mandatory Counterparts

| Prohibited Legacy Tool / Pattern | GitMap Mandatory Command | Technical Rationale & Operational Advantage |
| :--- | :--- | :--- |
| ❌ `Get-ChildItem -Path "..."`, `Get-ChildItem`, `dir`, `ls` | ✅ `gitmap lf [dir]` or `gitmap find "*.*" [dir]` | Prevents slow PowerShell .NET object pipeline traversal, terminal hangs, and directory recursion overhead. Returns clean relative paths in <10ms. |
| ❌ `Get-ChildItem -Recurse`, `find . -name "..."`, `dir /s /b` | ✅ `gitmap find "<pattern>" [-ext <ext>]` | Multi-core indexed glob discovery across 10,000+ files in <15ms without filesystem walk penalty. |
| ❌ `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `findstr` | ✅ `gitmap aum search "<pattern>" [dir] [-e <ext>] [-r] [-i]` | Multi-core streaming live search with binary null-byte probe, 500 KB file cap, and DH2D SQLite hot cache. Prevents buffer bloat and terminal freezing. |
| ❌ Raw unbuffered `cat`, `type`, `Get-Content` | ✅ `gitmap cat <filepath>` | Direct zero-disk stream of file content into stdout for resource-constrained CLI sessions. |
| ❌ `powershell -Command "..."` or `pwsh -Command "..."` | ✅ `gitmap pwsh "<cmd>"` or `gitmap ps "<cmd>"` | Cross-platform PowerShell execution with deterministic `-NoProfile`, UTF-8 encoding, and automatic fallback. |
| ❌ Multi-repo `git status` loops | ✅ `gitmap status --dirty` / `gitmap status --json` | Single-shot multi-repository status matrix displaying dirty counts, ahead/behind branches, and uncommitted changes. |
| ❌ Untracked cognitive reasoning or blind file edits | ✅ `gitmap heatmap log -t <slug> -a <action> -f <file> -r "<reason>"` | Centralized cognitive telemetry into `.ai-memory/heatmap.db` for crash forensics, file contention tracking, and dynamic agent joining. |
| ❌ Blind file editing in multi-agent swarms | ✅ `gitmap agent heatmap` / `gitmap heatmap status` | Pre-edit contention analysis ranking active write leases and collision blast radius across agents. |

### 4.2 Concrete Operational Examples for AI Agents

#### Scenario A: Directory Inventory & File Listing (Replacing `Get-ChildItem`)
```bash
# List all files in a specific directory:
gitmap lf "02-spec/21-app"

# List files matching a glob pattern inside a directory:
gitmap find "*.md" "02-spec/21-app/gitmap-commands-agent-spawning-heatmap-sync"

# List all files across a module:
gitmap find "*.*" "src-tauri/src/modules/repo_db"

# Find files by exact substring:
gitmap find-files-any "heatmap"
```

#### Scenario B: Code & Symbol Search (Replacing `Select-String` / `rg`)
```bash
# Search for a symbol across a target directory with extension filtering:
gitmap aum search "RunFleetPASCommand" cli -e .go

# Case-insensitive regex search:
gitmap aum search -r "(\"pas\"|\"pa\")" cli/cmd -e .go -i

# Search across specifications:
gitmap aum search "Subagents" 01-prompts -e .md
```

#### Scenario C: File Content Streaming (Replacing `cat` / `Get-Content`)
```bash
# Stream file contents directly to stdout without disk buffer bloat:
gitmap cat "02-spec/01-spec-authoring-guide/01-spec-structure.md"

# Stream configuration or prompt header:
gitmap cat "version.json"
```

#### Scenario D: Cognitive Heatmap Telemetry (Lifecycle Integration)
```bash
# 1. Step 0 Preflight Inspection:
gitmap heatmap status
gitmap heatmap tree
gitmap agent heatmap

# 2. Logging a file read before opening:
gitmap heatmap log -t "gitmap-commands-agent-spawning-heatmap-sync" -a "READ" -f "AGENTS.md" -r "Analyze Section 12 shell search ban"

# 3. Logging a file edit after modification:
gitmap heatmap log -t "gitmap-commands-agent-spawning-heatmap-sync" -a "EDIT" -f "AGENTS.md" -r "Add Section 15 Cognitive Heatmap Telemetry and expand shell ban"

# 4. Checking table schemas and action logs:
gitmap heatmap tables
gitmap heatmap query "SELECT HeatmapActionLogId, ActionType, FilePath, ReasoningNotes FROM HeatmapActionLog ORDER BY HeatmapActionLogId DESC LIMIT 5" --json
```

---

## 5. Detailed Specification for File Modifications

### 5.1 Modifications to `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`
1. **Line 40 Update (Excise Escape Hatch):**
   - *Target:* `2. Spec Step (A = 3 self subagents or lead)`
   - *Replacement:* `2. Spec Step (A = 3 self subagents — ZERO SOLO SPEC AUTHORING): Spawn exactly 3 authoring subagents (TypeName: "self") in a single invoke_subagent tool call to author modular, disjoint spec files and subtasks. The lead orchestrator is strictly forbidden from authoring specs solo.`
2. **Section 31 Update (Structural Array Validation):**
   - Add explicit validation rule: `invoke_subagent` MUST contain an array of exactly 3 subagent objects (`Subagents: [agent1, agent2, agent3]`). Spawning fewer than 3 subagents (or issuing solo text without tool invocation) is an auto-reject failure.
3. **Section 3.1 & 3.2 Update (Command Primacy Table & Shell Ban):**
   - Update prohibited search tools to explicitly include `Get-ChildItem` (all variants), `dir`, `ls`, `cat`, `type`, `Get-Content`.
   - Add explicit rows for `gitmap lf [dir]`, `gitmap find "*.*" [dir]`, and `gitmap cat <path>`.
4. **Section 7.1 Update (Phase 1 Spec JSON Payload Template):**
   - Insert complete Phase 1 Spec JSON payload template (Spec Author 01, Spec Author 02, Spec Author 03) alongside Phase 1 Planning and Phase 2 Execution payloads.
5. **Heatmap Lifecycle Telemetry Updates:**
   - Embed operational heatmap logging commands across Step 0, Phase 1, Phase 2, and Phase 3.

### 5.2 Modifications to `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` and `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`
Mirror all changes from `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` with identical line hygiene, maintaining synchronization across both Antigravity and Cursor skill trees.

### 5.3 Modifications to `.agents/skills/gitmap/skill.md` and `.cursor/skills/gitmap/skill.md`
1. **Command Replacement Matrix Expansion:**
   - Update Row 24 to prohibit ALL variants of `Get-ChildItem` (not just `-Recurse`), `dir`, `ls`, `cat`, `type`, `Get-Content`.
   - Add distinct entries for directory inventory (`gitmap lf [dir]`) and glob searching (`gitmap find "<pattern>" [dir]`).
2. **Prominent Copy-Paste Examples:**
   - Add a dedicated subsection under Section 5 ("High-Performance Automation & File Discovery") providing concrete examples for directory listing, substring matching, live regex search, and zero-disk streaming.
3. **Cognitive Heatmap Section Enhancement:**
   - Expand Section 12 ("File Contention Heatmap & Agent Activity Observation") with end-to-end examples of `gitmap heatmap log`, `gitmap heatmap status`, and `gitmap heatmap tree`.

### 5.4 Modifications to `AGENTS.md`
1. **Section 12 Expansion:**
   - Expand shell search ban: Total ban on `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem` (all variants), `dir`, `ls`, and `findstr`.
   - Explicitly cite `gitmap lf` and `gitmap find` for directory and file listing.
2. **Section 15 Addition ("Cognitive Heatmap Telemetry Primacy"):**
   - Mandate that all AI agents record file read/edit decisions with reasoning via `gitmap heatmap log`.
   - Require preflight heatmap inspection (`gitmap heatmap status`, `gitmap heatmap tree`, `gitmap agent heatmap`) prior to touching shared files in multi-agent swarms.

---

## 6. Verification & Conformance Criteria

1. **Schema Conformance:** All JSON payloads must be syntactically valid JSON conforming to the `invoke_subagent` tool schema.
2. **Cardinality Verification:** Every dispatch payload must contain exactly 3 subagent elements in the `Subagents` array.
3. **Relative Path Hygiene:** Zero absolute filesystem paths (`C:\...`, `/home/...`) or `file:///` URIs across all specifications, plans, and instructions.
4. **Tool Parity:** `.agents/skills/` and `.cursor/skills/` must maintain 100% text parity for all updated files.
5. **No Intermediate Commits:** No git commits or branches created during spec authoring or plan generation.
