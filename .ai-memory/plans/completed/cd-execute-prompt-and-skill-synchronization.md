# Completed Plan: CD Execute Prompt & Skill Synchronization

## Metadata
- Slug: `cd-execute-prompt-and-skill-synchronization`
- Spec: `02-spec/21-app/cd-execute-prompt-and-skill-synchronization/`
- Status: COMPLETED
- Date: 2026-10-18
- Swarm Topology: Tree Swarm (Lead Orchestrator, Worker 01, Worker 02, Worker 03)
- System Concurrency Capacity: 6 operations (A = 3, H = 2)
- Task DB: `.ai-memory/temp-agents/cd-execute-prompt-and-skill-synchronization/agent-task.db`

---

## User Request (Verbatim)
# CD Execute Prompt and Skill Synchronization: high priority instruction, non-negotiable task

We need a CD execute prompt for this type of code to find and replace it using a map. The map would be replaced by one line rather than using it details. And based on the map finding, it could do another stuff, probably run another function or do another thing, very simplified way. I think, running a function would be better, easier to detect and do the things. So try to have this code smell to be avoided and create a CD execute prompt following the current execute parent task in NS steps V6, and also make sure that all the CD execute prompts actually follow this execute parent task in NS steps V6, and also the skills are updated and skills are synced with all the other known repositories. 

Also using the current execute parent tasks-with-n-steps-, the AI does not use the `heatmap` RC `heatmap` commands. Please check what could be the root cause of it. Did you update the skill? Did you update these skills in all these repositories? Confirm this. What could be the root cause of it? Find it and improve the `CG Execute` prompt and its skills, and also then sync it. First update, then sync.

## task: CD Execute Prompt and Skill Synchronization
## slug: cd-execute-prompt-and-skill-synchronization

# Actionable Items Must Follow Non-Negotiable

1. Git pull BEFORE starting any work — first try `gitmap pull <repo path>` (or run `gitmap pull` in repository folder); if that fails, fall back to regular `git pull` and resolve any merge conflicts so the repository is completely up to date and clean before starting.
2. Write spec under 02-spec/21-app/cd-execute-prompt-and-skill-synchronization/ and enqueue plan task in .ai-memory/plans/cd-execute-prompt-and-skill-synchronization.md (subtasks in .ai-memory/plans/subtasks/cd-execute-prompt-and-skill-synchronization/) first
3. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String
4. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well
5. Create a CD execute prompt to find and replace code using a map, ensuring it runs functions based on map findings.
6. Ensure all CD execute prompts follow the execute parent task in NS steps V6.
7. Update and sync skills across all known repositories.
8. Investigate why the AI does not use `heatmap` RC `heatmap` commands and identify the root cause.
9. Confirm if skills are updated in all repositories and improve the `CG Execute` prompt and its skills.
10. Sync the updated skills after confirming updates.

---

## Root Cause Analysis: Heatmap & RC Command Omission
- **Root Cause 1 (Command Schema Mismatch):** V6 prompt and previous CG Execute prompts instructed `gitmap heatmap discuss [--window 15m]`, which does not exist in the compiled GitMap CLI (`gitmap: [E1000:VALIDATION] heatmap: unknown heatmap command: discuss`). The working native commands are `gitmap agent heatmap` and `gitmap heatmap status`.
- **Root Cause 2 (Missing Step 0 Preflight Directives):** Step 0 in execution prompts instructed `gitmap agent task init` and `gitmap agent task status`, but omitted `gitmap rc heatmap tables` and `gitmap heatmap status`. Because autonomous agents follow Step 0 sequentially, these commands were skipped.
- **Root Cause 3 (Discovery Subagent Context Blindness):** Research subagents briefs lacked explicit directives to inspect heat map databases and record exploration reasoning.
- **Solution:** Replaced `gitmap heatmap discuss` with `gitmap agent heatmap` and `gitmap heatmap status`. Injected mandatory Step 0 Heatmap & RC Preflight checks into `02-execute-parent-task-with-n-steps-v6.md` and all 38 CG Execute prompts and companion skills.

---

## Completed Subtasks Summary

### Task-01: RCA Investigation of Heatmap and RC Commands
- **Agent:** Lead Orchestrator
- **Status:** DONE
- **Evidence:** Documented full 5-part RCA in `02-spec/21-app/cd-execute-prompt-and-skill-synchronization/01-architecture-spec.md`. Validated working commands live (`gitmap rc heatmap tables`, `gitmap heatmap status`, `gitmap agent heatmap`).

### Task-02: Map Dispatch Spec and Prompt 38
- **Agent:** Worker 01
- **Status:** DONE
- **Evidence:**
  - Coding guideline: `02-spec/02-coding-guidelines/01-cross-language/35-map-dispatch-and-handler-lookup.md`
  - Canonical prompt: `01-prompts/15-cg-execute/38-map-dispatch-and-handler-lookup.md`
  - V4 mirror: `01-prompts/v4/15-cg-execute/38-map-dispatch-and-handler-lookup.md`
  - Antigravity skill: `.agents/skills/cg-map-dispatch-and-handler-lookup/skill.md`
  - Cursor skill: `.cursor/skills/cg-map-dispatch-and-handler-lookup/skill.md`
  - Verified catalog registrations in `01-prompts/15-cg-execute/readme.md`, `01-prompts/readme.md`, `.ai-memory/prompts.md`, and `02-spec/02-coding-guidelines/01-cross-language/readme.md`.

### Task-03: Upgrade All CG Execute Prompts and Skills to V6
- **Agent:** Worker 02
- **Status:** DONE
- **Evidence:**
  - Upgraded all 38 prompts in `01-prompts/15-cg-execute/` and `01-prompts/v4/15-cg-execute/` via `03-ai-scripts/47-upgrade-cg-execute-to-v6.py`.
  - Updated `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` and `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md`.
  - Synchronized `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`, `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`, `.agents/skills/cg-execute/skill.md`, and `.cursor/skills/cg-execute/skill.md`.

### Task-04: Fleet Repository Synchronization
- **Agent:** Worker 03
- **Status:** DONE
- **Evidence:** Executed `03-ai-scripts/38-sync-prompts-skills-scripts.py`. Verified dry-run (43/43 OK) and live synchronization across fleet repositories. Released and pushed 41 owned repositories with boundary protections preserved.

---

## Verification Evidence
- Task DB status: `gitmap agent task status -t cd-execute-prompt-and-skill-synchronization` -> `4/4 Subtasks DONE`.
- Cross-language spec verification: `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/01-cross-language --check-only` -> exit 0.
- Relative path verification: `python linter-scripts/check-relative-paths.py` -> exit 0.
- Forbidden strings verification: `python linter-scripts/check-forbidden-strings.py` -> exit 0.
