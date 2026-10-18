# Specification: Letterly Prompts Heatmap One-Liner Integration

> [!IMPORTANT]
> **Specification ID:** `02-spec/21-app/letterly-prompts-heatmap-one-liner-integration/01-architecture-spec.md`  
> **Status:** APPROVED & ACTIVE  
> **Scope:** `01-prompts/22-letterly/`, `01-prompts/23-cursor-prompts/`, and companion documentation  
> **Lifecycle Steps Budget:** N = 400  
> **Local Completion Boundary:** STRICT LOCAL ONLY (Total ban on fleet multi-repo sync)

---

## 1. Overview & Objectives

This specification defines the integration of the **GitMap Heatmap concept as a concise one-liner** across all Letterly prompt formatters (`01-prompts/22-letterly/`) and their Cursor IDE mirrors (`01-prompts/23-cursor-prompts/`).

Letterly formatters take raw human input (spoken, mobile-dictated, or desktop-typed) and reformat it into high-precision, non-negotiable execution prompts. Because multi-agent swarms operate with concurrent subagents ($A = 3$, Capacity = 6), pre-edit file contention prevention and cognitive decision telemetry are vital invariants. Adding the Heatmap concept as a clean, standardized one-liner ensures that every generated Letterly prompt instructs autonomous AI agents to leverage the GitMap Heatmap telemetry suite and report transparent database locations.

---

## 2. Standardized Heatmap One-Liner Formulations

### 2.1 Desktop & Multi-Line Letterly Formatters
For desktop formatters (`02-desktop-letterly.md`, `03-execute-n-steps-letterly.md`, `04-plan-letterly.md`, `05-release-letterly.md`, `06-cicd-fix-release-letterly.md`, `08-run-letterly.md`, `09-execute-with-verification-letterly.md`, `10-execute-with-release-letterly.md`, `11-git-pull-first-letterly.md`):

1. **Instruction Rule:**
   ```text
   Heatmap Telemetry Mandate: Track in-flight file touches, prevent write collisions, and query the cognitive activity graph via GitMap Heatmap (`gitmap heatmap discuss` / `gitmap heatmap log` / `gitmap heatmap live`) with transparent 3-tier database location reporting.
   ```

2. **Output Template Directive (under `## Additional Instructions`):**
   ```markdown
   - Track in-flight file touches and prevent write collisions using GitMap Heatmap (`gitmap heatmap discuss` / `gitmap heatmap log` / `gitmap heatmap live`) with transparent 3-tier database location reporting.
   ```

### 2.2 Mobile Single-Line Letterly Formatters
For mobile formatters (`01-mobile-letterly.md`, `07-mobile-cicd-fix-letterly.md`):

1. **Core Directive Inclusion (Step 7):**
   `track file touches and prevent write collisions using GitMap Heatmap (gitmap heatmap discuss / log / live); `

2. **Single-Line Continuous Paragraph Output Format:**
   Embedded between codebase search and relative paths within the single continuous line:
   `...; search codebase exclusively via GitMap; track file touches and prevent write collisions using GitMap Heatmap (gitmap heatmap discuss / log / live); strictly use relative git paths only; ...`

---

## 3. The 3-Tier Database Location Transparency Mandate

When an agent executes a prompt containing the heatmap one-liner, it must transparently report its active database endpoints:

| Tier | Role | Path | Query / Inspection Command |
|:---|:---|:---|:---|
| **Tier 1: Host Root** | Machine-Local Agent Root Database | `~/.gitmap/agent-root.db` | `gitmap heatmap tree` |
| **Tier 2: Repo-Cache** | Companion Repository Database | `repo-cache/.heatmap/<ai-slug>/<repo-slug>/heatmap.db` | `gitmap heatmap rc query "<sql>"` |
| **Tier 3: In-Flight Local** | Local Task Run Database | `.ai-memory/temp-agents/<slug>/agent-task.db` | `gitmap agent task status --json` |

---

## 4. Dual-Folder Parity & Link Syntax Invariants

1. **Antigravity Skills vs Cursor Skills Syntax:**
   - In `01-prompts/22-letterly/`:
     All skill references must use `[<skill-name>](file;.agents/skills/<skill-name>)`.
   - In `01-prompts/23-cursor-prompts/`:
     All skill references must use `[<skill-name>](file;.cursor/skills/<skill-name>)`.
2. **Slash Command Syntax:**
   - Both folders use standard slash commands: `[/plan](slashCommand;plan)`, `[/learn](slashCommand;learn)`, `[/goal](slashCommand;goal)`.
3. **Linter Compliance:**
   - All Letterly files must pass `python linter-scripts/check-letterly-format.py`.
   - All tests in `linter-scripts/tests/test_check_letterly_format.py` must pass with exit code 0.
