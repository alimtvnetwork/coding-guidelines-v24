# Completed Plan: Muse Prompt and Git Repository Fix Instructions

## User Request (Verbatim)
# Muse Prompt and Git Repository Fix Instructions: high priority instruction, non-negotiable task

If it is verified that inside the Muse prompt, we have added that if Muse has to fix a Git repo or something like this, any command that should be written inside the repo cache with the Muse folder with the task information and other stuff, so that Muse can also read and understand. If it has to fix on Git or any other commands to fix the stuff, troubleshooting. It makes sure to have Muse in the prompt section. Muse should have, let's say, troubleshooting folder inside it. Only if it is troubleshooting or fixing it should update that. And also one more problem that I do see in the execute in NS steps in V6, it is not showing what it's supposed to do. Previously, in the past, it was supposed to show what it has understood that needs to be written out very extremely. Try to force that and also update the skills. Update this with all the other repositories, and then make sure that tiny bits and pieces are improved. And also the agent needs to run three agents from now on with the V6. Three agents so that parallel work and if one agent is done, it should spawn another agent. So same thing should be mentioned inside the Muse prompt as well. Enhance it, update the skill and sync the skill and prompts with all the repositories that you know of.

## Canonical Specifications
- Architecture Spec: `02-spec/21-app/muse-prompt-and-git-repository-fix-instructions/01-architecture-spec.md`
- Muse Prompts Spec: `02-spec/21-app/muse-prompt-and-git-repository-fix-instructions/02-muse-prompts-spec.md`
- Troubleshooting & Repo-Cache Spec: `02-spec/21-app/muse-prompt-and-git-repository-fix-instructions/03-troubleshooting-repo-cache-spec.md`

## Summary of Accomplishments & Evidence
1. **Extreme "Understood" Turn-1 Showcase in V6 Engine:**
   - Overhauled Phase 1A Step 2 across `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md`, `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`, `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`, and `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md`.
   - Mandated an exhaustive top-level `### 🧠 Deep Comprehension & Understood Mandate` containing 5 mandatory pillars (Strategic Intent, Explicit Scope Boundaries, Non-Negotiable Constraints & Negative Guardrails Ingestion, Blast Radius & Dependency Topology, Multi-Agent Concurrency Topology) before task cards.
   - Upgraded task cards from shallow 1-2 sentence summaries to a 4-part verification schema (Intent & Technical Deliverable, Negative Boundaries & Invariants, Disjoint File Boundaries, Verification Evidence Contract).
   - Codified an auto-reject rule against generic or superficial boilerplate understanding text.

2. **3-Agent Swarm Concurrency Scaling (`A = 3, H = 2`, Capacity 6) & Dynamic Respawn:**
   - Scaled both V6 engine and Muse master prompts/skills from 2 to 3 parallel subagents (`A = 3`), 2 operational hands per agent (`H = 2`), achieving a System Concurrency Capacity of 6 concurrent subtask operations.
   - Codified the Immediate Dynamic Respawn Protocol (Section 7.4): the exact moment any single subagent completes and reports, the Lead Orchestrator immediately spawns a replacement worker to claim the next pending subtask, or commands the worker to dynamically join an ongoing subtask in collaborative assistance mode (parallel reading, companion linters, memory logging) to maintain 100% capacity saturation at 3 active subagents.

3. **Dedicated Phase 7.5: Git Repository Troubleshooting & Repo-Cache Forensics in Muse Master Prompt:**
   - Updated `01-prompts/27-muse-prompts/01-muse-master-prompt.md` to v6.5.0 and updated `01-prompts/27-muse-prompts/readme.md`.
   - Embedded dedicated section "Phase 7.5 — Git Repository Troubleshooting & Repo-Cache Forensics".
   - Mandated that when Muse fixes a Git repository or executes troubleshooting commands, all commands, diagnosis, error autopsy, and task metadata must be persisted to `repo-cache/muse/troubleshooting/<task-slug>/` via `gitmap rc` so Muse can persist, inspect, and reuse them across sessions.
   - Enforced rule: only when troubleshooting or fixing is actively required does Muse update and write to the troubleshooting archive.
   - Codified non-destructive Git remediation patterns: branch divergence reconciliation (`git rev-list`, `git merge --ff-only` or `--no-ff`), mechanical conflict resolution by domain solvers (`47-git-reconcile-and-resolve-conflict.py`), Git index lock contention clearing (`.git/index.lock`), and workspace debris isolation via `gitmap rc`.

4. **Muse Skills Modernization with Repo-Cache Hierarchy:**
   - Upgraded `.agents/skills/muse-master-prompt/skill.md` and `.cursor/skills/muse-master-prompt/skill.md`.
   - Documented the canonical 6-artifact diagnostic suite under `repo-cache/muse/troubleshooting/<task-slug>/`:
     * `01-incident-manifest.json` (metadata: task_slug, target_repo, branch, incident_category, severity, status).
     * `02-error-autopsy.md` (4-part RCA: Symptom, Trigger, Root Cause, Mitigation).
     * `03-git-state-snapshot.json` (head_sha, remote_tracking_sha, ahead/behind counts, conflicted files).
     * `04-harness-scripts/` (reusable diagnostic scripts preserved via `gitmap rc file`).
     * `05-remediation-runbook.md` (reproducible command sequence executed).
     * `06-verification-log.md` (proof of recovery: clean porcelain, 0 conflict markers, passing CI).
   - Documented `gitmap rc` offloading syntax (`gitmap rc file`, `gitmap rc folder`, `gitmap rc text`).
   - Integrated the "Persist Before Proceeding" invariant and Troubleshooting Recall protocol.

5. **Multi-Agent Parallel Execution Wave & Task DB Verification:**
   - 3 parallel subagents (Worker 01, Worker 02, Worker 03) executed concurrently.
   - All 3 subtasks marked `DONE` with exit 0 in `.ai-memory/temp-agents/muse-prompt-and-git-repository-fix-instructions/agent-task.db`.

6. **Quality & Verification Checks:**
   - Relative Git Paths: Verified 100% relative paths from git root (0 absolute filesystem paths and 0 `file:///` URIs).
   - Lowercase Naming: Verified 100% lowercase conventions observed across all new and modified paths.
   - Secrets Gate: Verified 0 private keys, passwords, or forbidden tokens.
