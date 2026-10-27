# Architecture Specification: Coding Guideline Corrections & GitMap CLI/Skill Updates

- **Specification Directory:** `02-spec/21-app/32-coding-guideline-corrections-and-gitmap-skill-updates/`
- **Related Spec:** `02-spec/21-app/31-coding-guideline-corrections-and-gitmap-skill-updates/`
- **Component Subsystems:**
  1. `cli/cmd/root_cobra_completion.go` (PowerShell Script Generation)
  2. `cli/heatmap/` & `cli/store/` (Heatmap Shorthands & Task Spec Tracking)
  3. `.agents/skills/follow-cg-execute/skill.md` (39-Step Orchestrator)
  4. `.agents/skills/gitmap/skill.md` (Native Process Termination & Heatmap Matrix)
  5. `.agents/skills/` & `.cursor/skills/` (Skills Inventory Consolidation)

---

## 1. PowerShell Completion Script Root Cause & Architecture

### 1.1 The Anomaly
When completing commands in PowerShell, users experienced sudden suppression of suggestions.
Inspecting PowerShell runtime logs revealed:
```text
InvalidOperation: Method invocation failed because [System.Management.Automation.ScriptBlock] does not contain a method named 'op_Addition'.
```

### 1.2 Root Cause Analysis
In `d:/work/gitmap/cli/cmd/root_cobra_completion.go`, Cobra's default generator produces lines such as:
```powershell
$CompletionText = $($comp.Name | __gitmap_escapeStringWithSpecialChars) + $Space
```
A previous hardening pass ran:
```go
rawScript = strings.Replace(rawScript,
    "$CompletionText = $($comp.Name | __gitmap_escapeStringWithSpecialChars)",
    "$CompletionText = [string]($comp.Name | __gitmap_escapeStringWithSpecialChars); if (-not $CompletionText) { $CompletionText = [string]$comp.Name }; if (-not $CompletionText) { continue }",
    -1)
```
Because the target string lacked `+ $Space`, the trailing ` + $Space` from Cobra was appended directly after the replacement block, producing:
```powershell
$CompletionText = [string]($comp.Name | __gitmap_escapeStringWithSpecialChars); if (-not $CompletionText) { $CompletionText = [string]$comp.Name }; if (-not $CompletionText) { continue } + $Space
```
In PowerShell, `{ continue } + $Space` treats the script block as an operand in an addition operation, throwing `InvalidOperation`.

### 1.3 Concrete Fix
Replace the with-space variant first, then the without-space variant:
```go
rawScript = strings.Replace(rawScript,
    "$CompletionText = $($comp.Name | __gitmap_escapeStringWithSpecialChars) + $Space",
    "$CompletionText = [string]($comp.Name | __gitmap_escapeStringWithSpecialChars); if (-not $CompletionText) { $CompletionText = [string]$comp.Name }; if (-not $CompletionText) { continue }; $CompletionText = $CompletionText + $Space",
    -1)
rawScript = strings.Replace(rawScript,
    "$CompletionText = $($comp.Name | __gitmap_escapeStringWithSpecialChars)",
    "$CompletionText = [string]($comp.Name | __gitmap_escapeStringWithSpecialChars); if (-not $CompletionText) { $CompletionText = [string]$comp.Name }; if (-not $CompletionText) { continue }",
    -1)
```

---

## 2. Process Termination Architecture (`gitmap kill`)

### 2.1 Mandate
Eliminate raw process management tools (`Stop-Process`, `taskkill /F /PID`, `kill -9`) from AI workflows and skill references.
Mandate the native compiled Go command:
```bash
gitmap kill <pid|name|:port> [flags]
```

### 2.2 Capabilities & Safety Guards
- **Process ID (PID):** `gitmap kill 12345` (graceful termination).
- **Forced Termination:** `gitmap kill 12345 --force` (`-f`).
- **Process Tree Termination:** `gitmap kill 12345 --tree` (`-t`, recursive termination of child processes).
- **Port-Based Termination:** `gitmap kill :8080 -f` (resolves PID listening on TCP port and terminates).
- **Binary Name Termination:** `gitmap kill node.exe -f`.
- **Built-in Self Protection:** Rejects PID 0, 1, 4 (system), current process (`os.Getpid()`), and developer IDEs (`code`, `antigravity`, `pwsh`, `bash`) unless explicitly forced.

---

## 3. Cognitive Heatmap Shorthands & Task Spec Tracking

### 3.1 Shorthand Commands vs Direct Queries

| Shorthand Command | Target Subcommand Route | Direct SQL Query Replaced |
| :--- | :--- | :--- |
| `gitmap heatmap failed` (or `failures`, `unmet`) | `RunHeatmapFailures` | `SELECT HeatmapActionLogId, TaskSlug, ActionType, FilePath, ExpectedOutcome, ActualOutcome FROM HeatmapActionLog WHERE IsExpectationMet = 0 ...` |
| `gitmap heatmap hotfiles` (or `hot`, `hotspots`) | `RunHeatmapHotfiles` | `SELECT FilePath, COUNT(*) AS Touches, SUM(...) ... GROUP BY FilePath ...` |
| `gitmap heatmap recent` (or `latest`) | `RunHeatmapRecent` | `SELECT HeatmapActionLogId, TaskSlug, ActionType, FilePath, ReasoningState ... ORDER BY HeatmapActionLogId DESC LIMIT 20` |
| `gitmap heatmap tasks` (or `task-list`) | `RunHeatmapTasks` | `SELECT TaskSlug, TaskName, Status, CreatedAt, UpdatedAt FROM HeatmapParentTask ORDER BY UpdatedAt DESC ...` |

### 3.2 Dual-Tier Task Spec Registration
During task initialization (`gitmap task init --name <title> --spec-file <file> --spec-folder <dir>`):
1. **Tier 2 Database (`agent-task.db`):** Inserts record into `ParentTask` with `SpecFile` and `SpecFolder`.
2. **Heatmap Database (`heatmap.db`):** Calls `RecordParentTask` on `HeatmapDB` to upsert `TaskSlug`, `TaskName`, `RootSpecFilePath`, and `TaskFolder` into `HeatmapParentTask`.
3. **Root Heatmap Database (`root-heatmap.db`):** Calls `RecordParentTask` on `RcHeatmapStore` to track cross-repo progress.

---

## 4. `follow-cg-execute` 39-Step Master Orchestrator

The 39-step execution engine sequences every coding guideline domain sequentially:
- **Steps 1 through 38 (Fast-Forward Refactoring):**
  - Execute bounded micro-batches of 5–8 files.
  - Apply discrete code improvements.
  - Zero intermediate builds, zero intermediate test runs.
  - Intermediate checkpoints saved via `gitmap commit "wip - <scope>"`.
- **Step 39 (Terminal E2E Verification & Release):**
  - Full codebase compilation and build validation.
  - Execution of end-to-end test suites.
  - Verification of 100% green status with zero regressions.
  - Automated release ceremony (`gitmap cpr` and version bump).

---

## 5. Skills Inventory Consolidation Matrix

| Redundant Skills to Archive | Consolidated Canonical Skill | Rationale |
| :--- | :--- | :--- |
| `major-bump`, `minor-bump`, `patch-bump`, `version-bump-docs` | `release` | All 4 are thin prompt wrappers forwarding to `01-prompts/17-release-management/`. The master `release` skill governs all bump levels and release documentation. |
| `logo-create`, `svg-logo` | `logo-and-icon-design` | Both are partial wrappers around logo prompt templates. `logo-and-icon-design` provides the complete specification. |
| `theming-guidelines`, `update-theming` | `react-ui-theming-design` | Wrappers for CSS token updates consolidated into the modern Radix/Tailwind design system skill. |
| `sync`, `sync-codebase`, `sync-other-codebase` | `prompts-and-skills-sync` | Thin wrappers for sync prompts replaced by the programmatic script-driven sync engine. |
| `next-steps`, `pending-tasks` | `check-pending-improvements` | Redundant backlog wrappers replaced by the structured KPI dashboard and next-10 execution queue. |
