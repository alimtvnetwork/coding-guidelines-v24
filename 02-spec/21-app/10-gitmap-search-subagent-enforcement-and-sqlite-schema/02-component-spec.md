# Component Specification: GitMap Search Subagent Enforcement & Prompt/Skill Hardening

> **Spec ID:** `02-spec/21-app/10-gitmap-search-subagent-enforcement-and-sqlite-schema/02-component-spec.md`  
> **Parent Architecture Spec:** `02-spec/21-app/10-gitmap-search-subagent-enforcement-and-sqlite-schema/01-architecture-spec.md`  
> **Status:** APPROVED  
> **Author:** Lead Component Architect & Spec Subagent 02  
> **Target Subsystems:**
> - `01-prompts/14-execute/13-execute-parent-task-with-n-steps-v6.md`
> - `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`
> - `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`
> - `01-prompts/22-letterly/03-execute-n-steps-letterly.md`
> - `01-prompts/23-cursor-prompts/03-execute-n-steps-letterly-cursor.md`
> - `.agents/skills/letterly-execute-n-steps/skill.md`
> - `.cursor/skills/letterly-execute-n-steps/skill.md`
> - `AGENTS.md` (Global AI Guidelines Section 12)  
> **Companion Subtask Plan:** `.ai-memory/plans/subtasks/10-gitmap-search-subagent-enforcement-and-sqlite-schema/02-subagent-prompts.md`

---

## 1. Executive Summary & Component Objectives

During complex multi-agent coding sessions, background subagents (`TypeName: "research"` and `TypeName: "self"`) dispatched by the orchestrator have routinely defaulted to raw shell search utilities—predominantly `rg` (`ripgrep`), `grep`, and PowerShell `Select-String`—as captured in runtime forensic logs.

This component specification details the exact prompt, skill, and configuration updates required to enforce GitMap search primacy across all agent tiers:

1. **`execute-parent-task-with-n-steps-v6` (.agents, .cursor, and 01-prompts/14-execute/):**
   - Hardens Section 3 (GitMap High-Speed Command Primacy) with an unambiguous, auto-reject ban on `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem -Recurse`, and `findstr`.
   - Replaces generic planning subagent prompt placeholders with complete, turnkey Discovery Brief templates that mandate GitMap search and forbid file writes.
   - Replaces the worker brief template with full GitMap search directives, ripgrep bans, and explicit SQLite action logging commands.
   - Updates the verification checklist to mandate inspecting subagent tool logs for banned search binaries.

2. **`letterly-execute-n-steps` (.agents, .cursor, 01-prompts/22-letterly/, 01-prompts/23-cursor-prompts/):**
   - Injects a mandatory GitMap Search Directive as Actionable Item 2 into all formatted voice dictations.
   - Ensures that any downstream agent prompted via Letterly receives GitMap search primacy as a non-negotiable instruction from the user prompt itself.

3. **`AGENTS.md` Section 12:**
   - Adds a new global meta-repository standard establishing GitMap Search Primacy and the strict prohibition of raw shell search utilities across all agents and subagents.

---

## 2. Component 1: `execute-parent-task-with-n-steps-v6` Prompt & Skill Suite

### 2.1. File Locations & Synchronization Boundary
The V6 execution protocol is maintained across three synchronized mirrors:
- `01-prompts/14-execute/13-execute-parent-task-with-n-steps-v6.md` (Canonical prompt template)
- `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` (Native Antigravity agent skill)
- `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md` (Cursor IDE skill mirror)

All three files must receive identical updates to preserve cross-platform behavior.

### 2.2. Section 3 Hardening: GitMap Command Primacy & Search Tooling

The search protocol in Section 3 is updated to explicitly ban `rg` and list exact GitMap command equivalents:

```markdown
### 🔍 Code & Symbol Search Protocol (TOTAL BAN ON `rg`, `ripgrep`, `Select-String` & `git grep`)
- **Live Streaming Search (Default for discovery, symbol tracking & blast radius):**
  - Search string/symbol: `gitmap aum search "<symbol>" [dir] [-e <.ext>]` (e.g. `gitmap aum search "RunFleetPASCommand" cli -e .go`)
  - Search regex: `gitmap aum search -r "<regex>" [dir] [-e <.ext>]` (e.g. `gitmap aum search -r "(\"pas\"|\"pa\")" cli/cmd -e .go`)
  - Case-insensitive: `gitmap aum search -i "<query>" [dir]`
  - Scoped path search: `gitmap aum search "<pattern>" --path <relative-directory>`
- **Indexed Fast Search:** `gitmap search "<query>"` (Instant SQLite cached keyword/symbol search across indexed repos)
- **Wildcard File Search:** `gitmap find "<pattern>" [-ext <ext>]` (Multi-core glob filename finder)
- **Directory Inventory:** `gitmap list-files [pattern] [-ext <ext>]` or `gitmap lf [pat]`
- **Deterministic File Streaming:** `gitmap cat <filepath>`
- **TOTAL BAN (AUTO-REJECT FAILURE):**
  NEVER run `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem -Recurse`, or `findstr`.
  Executing raw shell search commands violates repository guidelines, bypasses index acceleration, and results in immediate review failure.
```

### 2.3. Section 6.1 Update: Phase 1 Planning Payload with Full Discovery Brief

In Section 6.1 (Planning Step) and Section 7.1 (Dispatch Payload Examples), replace the generic `<Discovery Brief...>` placeholders with complete, turnkey subagent prompts:

```json
{
  "Subagents": [
    {
      "TypeName": "research",
      "Role": "Research 01: Architecture & Blast Radius Discovery",
      "Model": "inherit",
      "Workspace": "inherit",
      "Prompt": "You are Research 01 for task <nn>-<slug>. You have no prior chat context; this brief is your complete specification.\n\n### Core Objective:\nExplore the codebase, map symbol callers, trace dependencies, and discover relevant source files for: <task title and objectives>.\n\n### Tool Capabilities & Strict Read-Only Boundary:\n- You are a READ-ONLY subagent (`TypeName: 'research'`).\n- You have read tools: `run_command`, `view_file`, `search_web`, `read_url_content`.\n- You DO NOT have `write_to_file` or `replace_file_content`. NEVER attempt to create or edit files.\n\n### Mandatory Search Primacy & Total Ban on Raw Grep:\n- Execute ALL searches and symbol discoveries exclusively via GitMap:\n  * Live symbol search: `gitmap aum search \"<symbol>\" [dir] [-e <.ext>]`\n  * Path-scoped search: `gitmap aum search \"<symbol>\" --path <relative-dir>`\n  * Indexed keyword search: `gitmap search \"<query>\"`\n  * File discovery: `gitmap find \"<pattern>\"`\n  * Directory inventory: `gitmap lf [dir]`\n  * View files: `gitmap cat <path>` or native `view_file`\n- TOTAL BAN (AUTO-REJECT FAILURE): NEVER run `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem -Recurse`, or `findstr`.\n- TOTAL BAN ON GIT COMMANDS: NEVER run `git` commands (`git add`, `git commit`, `git status`, `git diff`, etc.).\n\n### Required Findings Report (Send via send_message to Parent):\nWhen your research is complete, send a message to the caller containing:\n1. Key symbol definitions, structures, and entry points.\n2. Call sites and blast radius (files that will be affected by modifications).\n3. Recommended modular file ownership boundaries for worker subtasks.\n4. Verification commands to validate changes."
    },
    {
      "TypeName": "research",
      "Role": "Research 02: Specs & Dependency Mapping",
      "Model": "inherit",
      "Workspace": "inherit",
      "Prompt": "You are Research 02 for task <nn>-<slug>. You have no prior chat context; this brief is your complete specification.\n\n### Core Objective:\nInspect existing specifications under `02-spec/`, coding guidelines under `02-spec/02-coding-guidelines/`, and prior plan tasks for: <task title and objectives>.\n\n### Tool Capabilities & Strict Read-Only Boundary:\n- You are a READ-ONLY subagent (`TypeName: 'research'`).\n- You DO NOT have file authoring tools. NEVER attempt to create or edit files.\n\n### Mandatory Search Primacy & Total Ban on Raw Grep:\n- Execute ALL searches exclusively via GitMap:\n  * Live search in specs: `gitmap aum search \"<term>\" 02-spec`\n  * Indexed search: `gitmap search \"<term>\"`\n  * File discovery: `gitmap find \"*.md\"`\n  * View files: `gitmap cat <file>` or native `view_file`\n- TOTAL BAN: NEVER run `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem`, or `findstr`.\n- TOTAL BAN ON GIT COMMANDS: NEVER run `git` commands.\n\n### Required Findings Report (Send via send_message to Parent):\nSend a message to the caller detailing existing specs, architectural invariants, acceptance criteria, and positive boolean rules."
    }
  ]
}
```

### 2.4. Section 7.2 Update: Self-Contained Worker Brief Template

Update Section 7.2 to enforce explicit ripgrep bans and structured SQLite action logging:

```text
You are Worker <NN> for task <nn>-<slug>. You have no prior chat context; this brief is your complete specification.

### Boundaries & Crash Prevention:
- Read any file in the workspace; edit ONLY your Owned Files: <relative paths>.
- TOTAL BAN ON GIT COMMANDS (LOCK COLLISION PREVENTION): NEVER run ANY git commands (`git add`, `git commit`, `git push`, `git status`, `git diff`, `git checkout`). In shared workspaces, worker git calls create `.git/index.lock` collisions that immediately crash parallel agents. Only the lead orchestrator runs git commands after workers complete.
- TOTAL BAN ON COMMITS: Workers NEVER commit, stage, or push. Committing is exclusively reserved for the Lead Agent at Phase 3 via GitMap (`gitmap cpf "<module> - <summary>"` using hyphen `-`; no colon needed in GitMap cpf as colon is already provided).
- Code & Symbol Search (TOTAL BAN ON `rg`, `ripgrep`, `grep`, `Select-String`):
  * Use GitMap exclusively: `gitmap aum search "<pattern>" [dir] [-e <.ext>] [-r] [-i]` or `gitmap search "<pattern>"`.
  * TOTAL BAN: NEVER execute `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem -Recurse`, or `findstr`.
- After C tool calls, stop and report what you have.
- A tool failing twice: reply "STATUS: BLOCKED" with exact error and stop. Never guess paths and never troubleshoot machine.
- Workers that find a secret stop and report "BLOCKED: secret at <file>:<line>". They do not handle it themselves.
- Adhere to R1, R2, and R11 by ID.

### Assigned Subtasks (up to H subtasks):
- Subtask 1: .ai-memory/plans/subtasks/<nn>-<slug>/01-<name>.md
- Subtask 2: .ai-memory/plans/subtasks/<nn>-<slug>/02-<name>.md (if assigned)

### 100% Non-Negotiable Coding Guidelines (AUTO-REJECT ON VIOLATION):
1. Positive booleans ONLY: use `is` and `has` prefixes exclusively. NEVER evaluate explicit `== true`. NEVER combine positive and negative checks in the same condition (`if isA && !isB` is BANNED).
2. Go Structured Errors: return `*appfault.AppError`, never bare `error`.
3. Function Sizing: <= 8 lines preferred, hard cap 15 lines. Extract domain structs and raw generics to `types.go`.
4. Strict Relative Git Paths & Lowercase: zero absolute filesystem paths and zero `file:///` URIs. All new files strictly lowercase.
5. Repo Secrets: if any credentials or private tokens are needed, store them in the `repo-secrets` folder in the default work directory (via `gitmap rs`). Never commit secrets.
6. Zero Builds or Tests: NEVER run `go build`, `npm run build`, `go test`, or `pytest`.
7. Targeted Verification: Run only fast file-scoped linters (e.g. `python 03-ai-scripts/05-guideline-autofixer.py <folder> --check-only`). A check scanning 0 files is a FAIL.
8. GitMap Search Primacy (TOTAL BAN on ripgrep / rg / Select-String): NEVER execute `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem`, or `findstr`. Always use `gitmap aum search "<pattern>" [dir] [-e <.ext>] [-r]` for live symbol/regex discovery.

### Concurrency-Safe SQLite Action Logging (CRASH FORENSICS MANDATE):
- Worker subtasks are tracked in the run database: `<databasePath>`.
- Step 1: Claim assigned subtask atomically:
  `python 03-ai-scripts/46-agent-sqlite-task-manager.py claim --db <databasePath> --agent "Worker <NN>"`
- Step 2: BEFORE touching or modifying any owned file, you MUST log your in-flight action:
  `python 03-ai-scripts/46-agent-sqlite-task-manager.py log-action --db <databasePath> --subtask-id <id> --agent "Worker <NN>" --action "write_to_file" --file "<path>" --details "<action description>"`
  *(Note: This guarantees that if a tool execution crashes or the session is interrupted, the database permanently records the exact file you were touching and what caused the crash!)*
- Step 3: Run targeted validation on modified files only.
- Step 4: When your subtask passes targeted checks, mark completion in the database:
  `python 03-ai-scripts/46-agent-sqlite-task-manager.py complete --db <databasePath> --subtask-id <id> --agent "Worker <NN>" --evidence "PASS exit 0, <files>"`
- Step 5: If blocked or failing, record the failure:
  `python 03-ai-scripts/46-agent-sqlite-task-manager.py fail --db <databasePath> --subtask-id <id> --agent "Worker <NN>" --reason "<reason>"`

### Output Contract:
Write your subtask output to .ai-memory/plans/subtasks/<nn>-<slug>/01-<name>.json and reply with this JSON block, once per subtask, then stop:
{
  "task": "Task-01",
  "status": "DONE",
  "filesChanged": ["<path1>", "<path2>"],
  "checks": "<command> -> exit <code>, <files scanned>",
  "acceptance": { "ac1": "PASS <evidence>" },
  "assumptions": [],
  "blockers": []
}
```

### 2.5. Section 10 Checklist Addition: Search Primacy Gate
Add an explicit audit gate to the Phase 3 checklist:
```markdown
- [ ] GITMAP SEARCH PRIMACY (TOTAL BAN ON RG/GREP): Verified that all searches across lead and subagents were executed via GitMap (`gitmap aum search`, `gitmap search`, `gitmap find`). Confirmed that neither the lead nor any subagent executed `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, or `findstr`.
```

---

## 3. Component 2: `letterly-execute-n-steps` Prompt & Skill Suite

### 3.1. File Locations & Synchronization Boundary
The Letterly execution formatter is maintained across four locations:
- `01-prompts/22-letterly/03-execute-n-steps-letterly.md` (Antigravity voice formatter)
- `01-prompts/23-cursor-prompts/03-execute-n-steps-letterly-cursor.md` (Cursor voice formatter)
- `.agents/skills/letterly-execute-n-steps/skill.md` (Antigravity skill)
- `.cursor/skills/letterly-execute-n-steps/skill.md` (Cursor skill)

### 3.2. Actionable Items Specification Hardening
In all four files, Rule 4 of the instructions is updated to require that Item 2 of `# Actionable Items Must Follow Non-Negotiable` is ALWAYS the GitMap Search Directive:

```markdown
4. Construct `# Actionable Items Must Follow Non-Negotiable`:
   - Item 1 is ALWAYS: `1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/) first`
   - Item 2 is ALWAYS: `2. Execute all symbol and text searches strictly via GitMap (gitmap aum search, gitmap search, gitmap find); ripgrep (rg), grep, and PowerShell Select-String are strictly forbidden across lead and all subagents`
   - Item 3..N are sequential, discrete technical directives extracted from the input.
   - Final Item is ALWAYS: `Run retrospective AI verification prompt/script (01-retrospective-ai-verification.md / 03-ai-scripts/47-retrospective-ai-verification.py) or skill [<skill-link>] to audit specs, touched files, code quality, and CI/CD status upon task completion`
```

### 3.3. Output Format Example Update

The example block in each file is updated to reflect the new Item 2:

```markdown
Output Format:

[/plan](slashCommand;plan)

# High Priority Instruction

${Input Text Verbatim}

# Actionable Items Must Follow Non-Negotiable

1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/) first
2. Execute all symbol and text searches strictly via GitMap (gitmap aum search, gitmap search, gitmap find); ripgrep (rg), grep, and PowerShell Select-String are strictly forbidden across lead and all subagents
3. [Second actionable technical directive extracted from input]
4. [Third actionable technical directive extracted from input]
5. Run retrospective AI verification prompt/script (01-retrospective-ai-verification.md / 03-ai-scripts/47-retrospective-ai-verification.py) or skill [<skill-link>] to audit specs, touched files, code quality, and CI/CD status upon task completion

Must follow and spawn agent using

[execute-parent-task-with-n-steps-v6](<skill-path>)

## Additional Instructions

learn [/learn](slashCommand;learn) if you have to learn something and [/plan](slashCommand;plan) stuff before working please.
```

---

## 4. Component 3: `AGENTS.md` Global AI Guidelines

### 4.1. Section 12 Addition: GitMap Search Primacy & Strict Shell Search Ban
To establish an unassailable meta-rule across all repositories and AI sessions, add Section 12 to `AGENTS.md`:

```markdown
## 12. GitMap Search Primacy & Shell Search Bans (Total Ban on rg, grep, Select-String)

- **Mandatory GitMap Search Primacy:** All agents (lead orchestrators, planning agents, and worker subagents) MUST execute all file discovery, code searches, symbol tracking, and content streaming exclusively through GitMap CLI commands:
  - **Live Code & Regex Search:** `gitmap aum search "<pattern>" [dir] [-e <.ext>] [-r] [-i]`
  - **Path-Scoped Search:** `gitmap aum search "<pattern>" --path <relative-directory>`
  - **Fast Indexed Search:** `gitmap search "<query>"`
  - **File Discovery & Globbing:** `gitmap find "<glob>"`
  - **Directory Inventory:** `gitmap list-files [dir]` or `gitmap lf [dir]`
  - **Streaming File Content:** `gitmap cat <filepath>`
- **Total Ban on Raw Grep & Shell Search Tools (AUTO-REJECT FAILURE):**
  - AI agents and subagents are **STRICTLY FORBIDDEN** from invoking:
    * `rg` or `ripgrep`
    * `grep` or `git grep`
    * `Select-String` (PowerShell)
    * `Get-ChildItem -Recurse` for file content searches
    * `findstr` (Windows Command Prompt)
  - Executing any banned search tool violates repository performance and token economy standards.
- **Subagent Enforcement Mandate:**
  - Lead orchestrators must inject the GitMap search directive and the explicit ripgrep ban into every subagent prompt brief dispatched via `invoke_subagent`.
  - Research subagents (`TypeName: "research"`) are read-only and must never attempt to create or write files.
  - Worker subagents (`TypeName: "self"`) must log in-flight actions to the SQLite task database (`03-ai-scripts/46-agent-sqlite-task-manager.py log-action`) before modifying files.
```

---

## 5. Architectural Invariants & Quality Verification

| Invariant ID | Target Component | Requirement | Verification Method |
| :--- | :--- | :--- | :--- |
| `INV-SEARCH-01` | All Prompts & Skills | Explicit ban on `rg`, `ripgrep`, `grep`, `Select-String` | Automated regex check across prompt files |
| `INV-SEARCH-02` | Subagent Dispatch | Research discovery brief contains read-only constraint & GitMap commands | Inspection of `Subagents` array in V6 prompt |
| `INV-SEARCH-03` | Worker Dispatch | Worker brief contains SQLite `claim`, `log-action`, `complete`, `fail` commands | Inspection of Section 7.2 in V6 prompt |
| `INV-LETTERLY-01` | Letterly Formatters | Item 2 of Actionable Items is ALWAYS the GitMap search mandate | Inspection of 4 Letterly prompt/skill files |
| `INV-META-01` | `AGENTS.md` | Section 12 details GitMap search primacy and shell search bans | Inspection of `AGENTS.md` |
| `INV-PATHS-01` | All Files | Strictly relative git paths; zero absolute paths or `file:///` URIs | Automated pattern scan for `file:///` and drive letters |
| `INV-BOOL-01` | SQLite & Scripts | Positive boolean naming (`is`, `has`) and implicit condition evaluation | Linter scan with `05-guideline-autofixer.py` |
