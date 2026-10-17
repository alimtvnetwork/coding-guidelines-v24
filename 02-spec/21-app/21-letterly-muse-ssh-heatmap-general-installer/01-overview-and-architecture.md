# Specification: Letterly Pull-First Sequencing, Muse SSH Delivery Protocol, Heat Map Tracking & Dual Installers

## 1. Overview & Objectives

This specification defines the architectural and operational contracts for six interconnected capabilities:
1. **Letterly Prompts & Skills Re-sequencing (Pull First Mandate)**: Standardize all Letterly prompt formatters and companion skills so that Item 1 is ALWAYS Git pull via `gitmap pull` (fallback `git pull`) with merge conflict resolution. Re-sequence Item 2 as spec authoring, Item 3 as GitMap search primacy, Item 4 as relative paths, followed by sequential directives.
2. **Mobile Prompt Single-Line Output Invariant**: Guarantee that `01-mobile-letterly.md`, `07-mobile-cicd-fix-letterly.md`, and companion skills produce strictly ONE continuous single-line paragraph with zero newlines and zero conversational filler.
3. **Muse SSH Key Generation & Identity Token**: Standardize the SSH identity format `<email>-muse<muse-sequence>-<operator-name>` with automatic account email retrieval (via `git config user.email` or `gh api user --jq .email`), interactive operator prompting for missing variables, and native GitMap command execution (`gitmap ssh create <identity> -y`, `gitmap ssh view`, `gitmap ssh copy`).
4. **Muse Key Delivery Protocol (Format Error Elimination)**: Impose a total ban on pasting SSH public keys, private keys, tokens, or long strings directly into chat text (which causes line wrapping and invalid format errors). Mandate delivering keys as downloadable/disk files (`~/.ssh/<identity>.pub`) or copying via `gitmap ssh copy`, and explicitly declaring key type (**Authentication Key** vs **Signing Key**).
5. **Agent Heat Map & Decision/Reasoning Telemetry**: Embed granular in-flight action logging (`gitmap agent log`) with explicit reasoning into `execute-parent-task-with-n-steps-v6` so that file reads, modifications, and rationale are indexed in the SQLite task database and rendered via `gitmap heatmap`.
6. **Dual Installers (Coding Guidelines vs General Repository)**: Provide two distinct installer paths:
   - Full coding guidelines installer (`install.ps1`, `install-config.json`): specs, linters, scripts.
   - General repository installer (`general-install.ps1`, `general-install.sh`, `general-install-config.json`, `install.ps1 -General`): strictly `.ai-memory/`, `01-prompts/`, `03-ai-scripts/`, `.agents/skills/`, `.cursor/skills/`, `02-spec/01-spec-authoring-guide/`, `02-spec/21-app/`, and `assets/` structure (zero linters).
7. **GitMap General Installation Integration**: Provide `gitmap cg install general`, `gitmap cg general`, and `gitmap cd general` to automatically synchronize or install the general repository assets into any target project.

---

## 2. Architectural Components

### 2.1 Letterly Numbering & Sequencing Architecture
All desktop, execute, plan, release, run, and verification Letterly templates MUST conform to the unified sequence:
- **Item 1**: `Git pull BEFORE starting any work — first try gitmap pull <repo path> (or run gitmap pull in repository folder); if that fails, fall back to regular git pull and resolve any merge conflicts so the repository is completely up to date and clean before starting.`
- **Item 2**: `Write spec under 02-spec/21-app/<slug>/ and enqueue plan task in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/) first`
- **Item 3**: `Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String`
- **Item 4**: `Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well`
- **Item 5..N**: Mode-specific fixed items and extracted technical directives.

### 2.2 Mobile Prompt Single-Line Output Contract
- Clean input text verbatim, stripping filler words.
- Single continuous line starting with `# High Priority Instruction: `.
- Embeds gitmap pull first directive and relative paths mandate.
- Ends with skill trigger `[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)`.
- Zero newlines, zero line gaps, zero leading markdown fences.

### 2.3 Muse Master Prompt & SSH Delivery Architecture
- Identity syntax: `<email>-muse<muse-sequence>-<operator-name>`
- Example: `alex.tone@company.com-muse-10-opRoden`
- Variable discovery:
  1. `<email>`: Query `git config user.email` or `gh api user --jq .email`. If empty, prompt operator.
  2. `<muse-sequence>` & `<operator-name>`: If unknown, prompt operator before generating key.
- Command execution: `gitmap ssh create "<identity>" -y`
- Delivery protocol:
  - DO NOT paste raw key text into chat.
  - Save key to `~/.ssh/<identity>.pub` and `.ai-memory/ssh/<identity>.pub`.
  - Copy to clipboard via `gitmap ssh copy "<identity>"`.
  - State explicitly: **Key Type: Authentication Key** (or Signing Key if configured).

### 2.4 GitMap Heat Map & Action Decision Logging
- In `execute-parent-task-with-n-steps-v6`:
  Before modifying or reading critical files, agents must log:
  `gitmap agent log --subtask <id> --agent "<name>" --action "<read|modify|plan>" --file "<relative-path>" --details "<reasoning for file access>"`
- In GitMap CLI:
  `gitmap heatmap` or `gitmap agent heatmap` queries `agent-task.db` and renders a terminal heat map showing touched files, action counts, and reasoning.

### 2.5 Dual Installers Specification
- **Full Coding Guidelines (`install.ps1`, `install.sh`)**:
  - `folders`: `spec`, `02-spec`, `linters`, `linter-scripts`, `linters-cicd`, `scripts/fix-repo`, `scripts/visibility-change`, `.ai-memory/coding-guidelines`.
  - Includes all 20+ shared specs and linter tooling.
- **General Repository Installer (`general-install.ps1`, `general-install.sh`, `general-install-config.json`, `install.ps1 -General`)**:
  - `folders`:
    - `.ai-memory`: Folder structure, memory indices, plans/subtasks templates.
    - `01-prompts`: Canonical prompt templates.
    - `03-ai-scripts`: Automation and helper scripts.
    - `.agents/skills` & `.cursor/skills`: Native AI skills.
    - `02-spec/01-spec-authoring-guide`: Spec authoring guide.
    - `02-spec/21-app`: Private application domain specs and templates.
    - `assets`: Screenshot and image guidelines directory structure (`assets/screenshots/`, `assets/images/`).
  - Total exclusion of `linters`, `linter-scripts`, `linters-cicd`, `scripts/fix-repo`, `scripts/visibility-change`.

### 2.6 GitMap Command Integration
- `gitmap cg install general` / `gitmap cg general` / `gitmap cg install --general`
- Alias routing from `gitmap cd general` -> dispatches general installer directly.
- Direct synchronization from local `coding-guidelines` clone when detected, falling back to remote `general-install.ps1`.
