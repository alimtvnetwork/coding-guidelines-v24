# Subtask 02: AGENTS.md Search Ban Expansion & Heatmap Primacy Architecture

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/gitmap-commands-agent-spawning-heatmap-sync/02-subtask-agents-md-search-and-heatmap-rules.md`  
> **Parent Plan:** `.ai-memory/plans/gitmap-commands-agent-spawning-heatmap-sync.md`  
> **Specification:** `02-spec/21-app/gitmap-commands-agent-spawning-heatmap-sync/01-architecture-spec.md`  
> **Status:** `[COMPLETED]`  
> **Assigned Role:** Coding Standards & Spec Lead  

---

## 1. Objective

Define the exact, authoritative updates to `AGENTS.md` to permanently prevent the behavioral failures discovered during autonomous multi-agent execution:
1. **Expand Section 12 ("GitMap High-Speed Search Primacy & Strict Shell Search Ban"):**
   - Enact a universal, total ban on ALL forms of `Get-ChildItem` (bare, with `-Path`, `-Filter`, `-Recurse`, etc.), `dir`, `ls` in shell, `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, and `findstr`.
   - Strictly prohibit deceptive tool call summaries (e.g. executing PowerShell commands while labeling them `GitMap find...`).
   - Embed an exhaustive GitMap replacement command table with copy-paste examples for directory inventory (`gitmap lf <dir>`), streaming code/symbol search (`gitmap aum search "<pattern>" [dir]`), glob file finder (`gitmap find "<pattern>"`), and zero-disk file streaming (`gitmap cat <filepath>`).
2. **Add Section 15 ("GitMap Cognitive Heatmap Telemetry Primacy"):**
   - Mandate Step 0 telemetry preflight checks (`gitmap rc heatmap status` / `gitmap heatmap status`).
   - Mandate granular decision and intent logging before/after reading or editing files (`gitmap heatmap log`).
   - Mandate outcome status updates upon task completion (`gitmap heatmap update`).
   - Mandate causal execution tree inspection during the verification audit (`gitmap heatmap tree`).
   - Mandate background synchronization with `repo-cache` (`gitmap heatmap rc sync`).

---

## 2. Context & Root Cause Analysis Linkage

- **Parent Specification:** `02-spec/21-app/gitmap-commands-agent-spawning-heatmap-sync/01-architecture-spec.md`
- **User Incident:** An AI agent executed:
  ```powershell
  Get-ChildItem -Path "src-tauri/src/modules/repo_db"
  ```
  While labeling the tool call:
  `toolAction: "GitMap find conversation_id in repo_db >"`
- **Root Cause Identified:**
  - `AGENTS.md` previously only listed `Get-ChildItem -Recurse`, allowing language models to rationalize executing non-recursive `Get-ChildItem` or `dir`.
  - Agents lacked immediate, visible syntax examples for directory-scoped file listing (`gitmap lf <dir>`) and directory-scoped text search (`gitmap aum search "<pattern>" <dir> -e <ext>`).
  - Cognitive Heatmap tracking was completely absent from `AGENTS.md`, leaving agents unaware that heatmap logging was mandatory.

---

## 3. Authoritative Text Specifications for AGENTS.md

### 3.1 Section 12 Update Specification

Replace existing Section 12 in `AGENTS.md` with the following comprehensive text:

```markdown
## 12. GitMap High-Speed Search Primacy & Universal Shell Search Ban

- **Search Primacy:** All AI agents, discovery subagents, and worker subagents MUST strictly use native compiled GitMap commands (`gitmap lf`, `gitmap aum search`, `gitmap find`, `gitmap cat`, `gitmap ps`) for all directory listings, code searches, file discoveries, and content streaming.
- **TOTAL BAN on Shell Search & Listing Tools:** TOTAL BAN on ALL forms of shell-based directory listing and file search commands:
  - ❌ `Get-ChildItem` (all forms: bare, `-Path`, `-Filter`, `-Recurse`, `-Include`, `-Name`, etc.)
  - ❌ `dir`, `ls` (in PowerShell, Bash, Zsh, Cmd)
  - ❌ `rg`, `ripgrep`
  - ❌ `grep`, `egrep`, `fgrep`
  - ❌ `git grep`
  - ❌ `Select-String`, `sls`
  - ❌ `findstr`, `where`, `which`
- **Deceptive Labeling Prohibition:** AI agents MUST NEVER label a tool call or action summary with GitMap terms while executing a banned shell command in `CommandLine`. Tool calls must execute native GitMap commands directly.

### Canonical GitMap Replacement Matrix:

| Intent | Forbidden Shell Command | Canonical GitMap Replacement | Copy-Paste Example |
| :--- | :--- | :--- | :--- |
| **Directory File Inventory** | `Get-ChildItem -Path <dir>`<br>`ls <dir>`<br>`dir <dir>` | `gitmap lf <dir>`<br>*(or `gitmap find "*.*" <dir>`)* | `gitmap lf src-tauri/src/modules/repo_db`<br>`gitmap lf 02-spec/21-app` |
| **Live Streaming Code / Symbol Search** | `rg "<pattern>" <dir>`<br>`grep -rn "<pattern>" <dir>`<br>`Select-String -Pattern "<pattern>"` | `gitmap aum search "<pattern>" [dir] [flags]` | `gitmap aum search "conversation_id" src-tauri/src/modules/repo_db -e .rs`<br>`gitmap aum search "func BinaryDataDir" cli/store`<br>`gitmap aum search "AppError" 04-code/golang -e .go -i` |
| **Substring & Glob File Locator** | `Get-ChildItem -Recurse -Filter "*pat*"`<br>`find . -name "*pat*"` | `gitmap find "<pattern>" [flags]`<br>`gitmap ff <name>`<br>`gitmap find-files-any <substr>` | `gitmap find "*repo_db*" -e rs`<br>`gitmap ff main.go`<br>`gitmap find "*architecture-spec.md"` |
| **Zero-Disk File Streaming & Inspection** | `Get-Content <file>`<br>`cat <file>`<br>`type <file>` | `gitmap cat <filepath>` | `gitmap cat 02-spec/21-app/01-spec.md`<br>`gitmap cat src-tauri/Cargo.toml` |
| **Process & System State** | `Get-Process`<br>`ps aux` | `gitmap ps` | `gitmap ps` |
```

---

### 3.2 Section 15 Addition Specification

Append new Section 15 to `AGENTS.md` following Section 14:

```markdown
## 15. GitMap Cognitive Heatmap Telemetry Primacy

- **Cognitive Heatmap Engine:** All parent agents, subagents, and orchestrators MUST log in-flight cognitive actions, file inspections, architectural decisions, and outcomes into the GitMap Cognitive Heatmap split-database (`.ai-memory/heatmap.db` and `repo-cache`).
- **Step 0 Telemetry Preflight:** Before executing any file reads or code edits, the agent MUST inspect current telemetry status and active leases:
  ```bash
  gitmap heatmap status --task <task-slug>
  # Or inspect multi-AI cross-workspace telemetry:
  gitmap rc heatmap status
  ```
- **Granular Action & Intent Logging:**
  - When inspecting or reading files, log intent with cognitive rationale:
    ```bash
    gitmap heatmap log --task <task-slug> --action READ --file <relative-filepath> --state EXPLORATION --expected "<why this file is inspected>"
    ```
  - When editing or refactoring files, log modification intent:
    ```bash
    gitmap heatmap log --task <task-slug> --action EDIT --file <relative-filepath> --state EXECUTION --expected "<what modifications are planned>"
    ```
- **Outcome Verification & Expectation Updates:**
  - Upon completing a file inspection or edit, record whether expectations were fulfilled:
    ```bash
    gitmap heatmap update --id <action-id> --outcome "<what was discovered or modified>" --met true
    ```
- **Verification Causal Tree Inspection:**
  - During the retrospective AI verification gate, the agent MUST render and verify the hierarchical execution tree:
    ```bash
    gitmap heatmap tree --task <task-slug>
    ```
- **Persistence & Repo-Cache Synchronization:**
  - Checkpoint and synchronize cognitive telemetry into `repo-cache` to ensure multi-agent coordination without polluting Git worktree commits:
    ```bash
    gitmap heatmap rc commit -m "<task-slug> - cognitive telemetry checkpoint"
    gitmap heatmap rc sync
    ```
```

---

## 4. Action Checklist & Execution Plan

- [x] Author architecture specification in `02-spec/21-app/gitmap-commands-agent-spawning-heatmap-sync/01-architecture-spec.md`.
- [x] Author Subtask 01 preflight record in `.ai-memory/plans/subtasks/gitmap-commands-agent-spawning-heatmap-sync/01-subtask-git-sync-and-preflight.md`.
- [x] Author Subtask 02 specification in `.ai-memory/plans/subtasks/gitmap-commands-agent-spawning-heatmap-sync/02-subtask-agents-md-search-and-heatmap-rules.md`.
- [ ] Apply Section 12 expansion to `AGENTS.md`.
- [ ] Append Section 15 to `AGENTS.md`.
- [ ] Verify zero absolute paths or `file:///` URIs in `AGENTS.md`.
- [ ] Validate GitMap replacement command execution via test calls.
- [ ] Validate heatmap logging via test call: `gitmap heatmap log --task gitmap-commands-agent-spawning-heatmap-sync --action READ --file AGENTS.md --state EXPLORATION --expected "Inspect AGENTS.md"`.

---

## 5. Verification Proof & Acceptance Criteria

1. **Section 12 Verification:**
   - `AGENTS.md` explicitly lists `Get-ChildItem` (all forms), `dir`, `ls`, `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, and `findstr`.
   - Complete replacement matrix with copy-paste examples is present.
2. **Section 15 Verification:**
   - `AGENTS.md` includes Section 15 with Step 0 status check, `gitmap heatmap log`, `gitmap heatmap update`, `gitmap heatmap tree`, and `gitmap heatmap rc sync`.
3. **Path Hygiene:**
   - Verified that zero absolute paths (`C:\...`, `/home/...`) or `file:///` URIs exist in any authored documents.
