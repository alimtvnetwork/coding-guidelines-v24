# Subtask Plan: Execute V6 Prompt & Skills A=3 Spawning Enforcement & Examples

## Subtask ID: `03-subtask`
- **Parent Plan:** `.ai-memory/plans/gitmap-commands-agent-spawning-heatmap-sync.md`
- **Spec Reference:** `02-spec/21-app/gitmap-commands-agent-spawning-heatmap-sync/02-component-and-examples-spec.md`
- **Primary Owner:** Worker 01 (Phase 2 Execution Wave)
- **Status:** `COMPLETED`

---

## 1. Intent & Technical Deliverables

Upgrade the master execution workflow prompt `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` and its companion skill implementations (`.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` and `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`) to permanently eliminate solo execution escape hatches, enforce exact 3-agent concurrency ($A = 3$), provide exhaustive JSON tool call payloads for all three phases, and integrate GitMap high-speed commands and cognitive heatmap telemetry into the core execution loop.

### Deliverables:
1. **Eliminate "or lead" Escape Hatch:**
   - Modify line 40 of `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` and line 46 of `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` from `2. Spec Step (A = 3 self subagents or lead)` to strictly require spawning exactly 3 `self` authoring subagents with an explicit ban on solo spec authoring.
2. **Add Phase 1 Spec 3-Agent JSON Payload Template:**
   - Insert a complete, copy-pasteable JSON payload in Section 7.1 showing how to spawn 3 `self` authoring subagents (`Spec Author 01`, `Spec Author 02`, `Spec Author 03`) with disjoint file ownership boundaries.
3. **Enforce Single Tool Call Array Schema Validation ($A = 3$):**
   - Mandate that every call to `invoke_subagent` MUST contain an array of exactly 3 subagent objects (`"Subagents": [agent1, agent2, agent3]`). Spawning fewer than 3 subagents in any phase is an auto-reject failure.
4. **Expand Shell Ban in GitMap Search Primacy Section:**
   - Update Section 3 table and rules to ban ALL forms of `Get-ChildItem` (not just `-Recurse`), `dir`, `ls`, `cat`, `type`, and `Get-Content`.
   - Provide concrete GitMap replacement examples (`gitmap lf [dir]`, `gitmap find "*.*" [dir]`, `gitmap cat <filepath>`).
5. **Operationalize Cognitive Heatmap Telemetry:**
   - Embed mandatory `gitmap heatmap log` calls into preflight, Phase 1, Phase 2, and Phase 3 instructions.

---

## 2. Disjoint File Boundaries

### Owned Files (Exclusive Modification Rights for Worker 01):
- `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`
- `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`
- `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`

### Read-Only Context Files:
- `02-spec/21-app/gitmap-commands-agent-spawning-heatmap-sync/01-architecture-spec.md`
- `02-spec/21-app/gitmap-commands-agent-spawning-heatmap-sync/02-component-and-examples-spec.md`
- `AGENTS.md`
- `.ai-memory/plans/gitmap-commands-agent-spawning-heatmap-sync.md`

---

## 3. Negative Boundaries & Invariants

- **TOTAL BAN ON GIT COMMANDS:** Subagents and workers NEVER run `git add`, `git commit`, `git push`, `git checkout`, or any git commands. Only the root orchestrator runs GitMap commit at final completion.
- **TOTAL BAN ON RAW SEARCH TOOLS:** Worker 01 must never run `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem`, `dir`, or `ls`. Use `gitmap aum search` and `gitmap find` exclusively.
- **Strict Relative Git Paths:** Zero absolute paths (`C:\...`, `/home/...`) or `file:///` URIs.
- **Strict Lowercase Filenames:** All referenced files must maintain lowercase naming.
- **Tool Parity:** `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` and `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md` must be kept in 100% sync with the prompt file.

---

## 4. Step-by-Step Implementation Guide for Worker 01

### Step 1: Update Prompt Header & Spawning Gate
In `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`:
1. Navigate to Section `### 🚨 MANDATORY SUBAGENT SPAWNING GATE (A = 3, H = 2, CAPACITY = 6 — TREE SWARM & DYNAMIC JOIN)`.
2. Locate item 2:
   ```text
   2. Spec Step (A = 3 self subagents or lead): Spawn 3 authoring subagents (TypeName: "self") to author modular, disjoint spec files and subtasks (never writing the same file).
   ```
3. Replace with:
   ```text
   2. Spec Step (A = 3 self subagents — ZERO SOLO SPEC AUTHORING): Spawn exactly 3 authoring subagents (TypeName: "self") in a single invoke_subagent tool call to author modular, disjoint spec files and subtasks. The lead orchestrator is strictly forbidden from authoring specs solo.
   ```
4. Add the structural array validation rule:
   ```text
   - MANDATORY ARRAY VALIDATION (EXACTLY 3 SUBAGENTS IN SINGLE CALL): Every call to `invoke_subagent` MUST contain an array of EXACTLY 3 subagent objects: "Subagents": [agent1, agent2, agent3]. Spawning fewer than 3 subagents (or calling invoke_subagent multiple times sequentially with 1 agent) is an auto-reject protocol failure.
   ```

### Step 2: Expand GitMap Primacy Table & Prohibited Shell Commands
In Section 3 ("GitMap High-Speed Command Primacy"):
1. Update the table to include:
   - Directory Inventory: `gitmap lf [dir]` / `gitmap find "*.*" [dir]` (replaces `Get-ChildItem`, `dir`, `ls`).
   - Zero-Disk File Stream: `gitmap cat <filepath>` (replaces `cat`, `type`, `Get-Content`).
   - Cognitive Heatmap Log: `gitmap heatmap log -t <slug> -a <action> -f <file> -r "<reason>"` (replaces unrecorded edits).
2. Explicitly update the Code & Symbol Search Protocol ban line to:
   ```text
   TOTAL BAN (AUTO-REJECT FAILURE): NEVER run `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem` (all variants), `dir`, `ls`, `cat`, `type`, `Get-Content`, or `findstr`.
   ```

### Step 3: Insert Phase 1 Spec JSON Payload Template
In Section 7.1 ("Dispatch Payload Examples"):
1. Insert the complete Phase 1 Spec Payload between Phase 1 Planning and Phase 2 Execution payloads:
   ```json
   #### Phase 1 Spec Authoring Payload (`TypeName: "self"` — Read/Write Specification Authors):
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

### Step 4: Synchronize to Antigravity & Cursor Skills
Apply identical modifications to:
- `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`
- `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`

---

## 5. Verification Evidence Contract

1. **Regex / Substring Search Verification:**
   - Execute: `gitmap aum search "ZERO SOLO SPEC AUTHORING" 01-prompts/14-execute` -> Must return 1 match.
   - Execute: `gitmap aum search "ZERO SOLO SPEC AUTHORING" .agents/skills/execute-parent-task-with-n-steps-v6` -> Must return 1 match.
   - Execute: `gitmap aum search "ZERO SOLO SPEC AUTHORING" .cursor/skills/execute-parent-task-with-n-steps-v6` -> Must return 1 match.
2. **Elimination of "or lead":**
   - Execute: `gitmap aum search "self subagents or lead" 01-prompts/14-execute` -> Must return 0 matches.
   - Execute: `gitmap aum search "self subagents or lead" .agents/skills/execute-parent-task-with-n-steps-v6` -> Must return 0 matches.
3. **Phase 1 Spec Payload Verification:**
   - Execute: `gitmap aum search "Spec Author 01: Architecture" 01-prompts/14-execute` -> Must return 1 match.
