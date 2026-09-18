# Completed Plan 15: GitMap Pipeline-AI & Dynamic Waiting Protocol in Prompts and Skills

## Header & Task History
- **Parent Task**: `15-gitmap-pipeline-waiting-prompts`
- **Trigger**: User request to integrate GitMap pipeline-ai with dynamic timeout/waiting into execute prompts, fix with RCA, and CI/CD fix prompts/skills to stop burning credits on rapid busy-polling loops.
- **Started At**: 2026-09-18T19:10:00+08:00
- **Completed At**: 2026-09-18T19:20:00+08:00
- **Total Loops / Steps Taken**: 2 Execution Subagents running in parallel across 2 bounded subtasks.

---

## User Request (Verbatim)

```text
I think we need to update the execute prompts, uh, mentioning that if, again, we have to get to the, um, the, uh, pipeline issue to understand if the pipeline is working or not. Uh, or probably if it says like fix with RCA, that should also em-emphasis, emphasize that, uh, it should follow the Git map to get the pipeline. And also every time it should use the, uh, waiting rather than just loop, loop, loop because it would waste the credits. So if, uh, pipeline is running, then I think Git map already has the code. You can check, uh, for AI, it can actually get the pipeline information. And following that, I think fix with RCA or CICD fix that, that should now... These prompts now should be updated to use properly the Gi-Git map, uh, pipeline AI section with timeout and everything. So check that. So that should have the explanation, uh, to use it properly and use the timeout or sleep, um, based on the estimated, uh, completion of the pipelines. Okay? So, so that, uh, there is no waste of the credit. Okay. Mention this very clearly in execute prompts and also the fix with RCA prompt, and also the CICD fix with release and CICD fix prompt. Do you understand this? Can you please update the prompts and skills in this coding guideline and make a commit and push and make a release? Can you please do that?

# Parent Task N-Step Continuous Loop & Multi-Agent Orchestration — Workflow (must follow)

> **Prompt Version:** 2.2.0
> **Synchronization:** Main Meta-Repo & Connected Workspaces

/goal Autonomously orchestrate and execute the parent task by decomposing it into subtasks and running a continuous N-step self-loop until completion without a single failure.

```text
N = 100
```
```

---

## Accomplished Work & Deliverables

### 1. Phase 0 Antigravity Skill Bootstrap
- Created `.agents/skills/gitmap-pipeline-waiting-prompts/skill.md` with standard YAML frontmatter for progressive disclosure.

### 2. Execute Prompts & Orchestration Skills (Subtask 01)
Updated 9 prompts and 7 skills:
- `01-prompts/14-execute/01-execute-pending-tasks.md`
- `01-prompts/14-execute/02-execute-parent-task-with-n-steps.md`
- `01-prompts/14-execute/03-execute-batched-loop.md`
- `01-prompts/14-execute/05-execute-batched-loop-wor.md`
- `01-prompts/14-execute/06-execute-parent-task-with-n-steps-v2.md`
- `01-prompts/14-execute/07-execute-batched-loop-v2.md`
- `01-prompts/05-coding-guidelines/02-execute-coding-guideline-fix.md`
- `01-prompts/09-commit-and-multi-agent-code-fix/02-execute-pending-tasks.md`
- `01-prompts/09-commit-and-multi-agent-code-fix/04-execute-pending-tasks-v2.md`
- `.agents/skills/execute-parent-task/skill.md`
- `.agents/skills/execute-parent-task-with-n-steps/skill.md`
- `.agents/skills/execute-batched-loop/skill.md`
- `.agents/skills/execute-batched-loop-wor/skill.md`
- `.agents/skills/execute-coding-guideline-fix/skill.md`
- `.agents/skills/execute-pending-tasks/skill.md`
- `.agents/skills/parent-task-n-step-loop/skill.md`

**Core Invariants Added:**
- **GitMap Pipeline-AI Authority:** Mandated `gitmap pipeline-ai status --json` (or alias `gitmap pl-ai status -t <sec>`) for checking pipeline state.
- **Anti-Credit-Waste Waiting Mandate:** Strictly banned rapid polling (`gh run view` in tight loops) and mandated dynamic waiting using `-t <seconds>` or sleeping based on estimated completion duration (`etaSeconds`).
- **NO RAPID CI/CD POLLING (TOTAL BAN)** added to banned operations checklist.

### 3. RCA & CI/CD Fix Prompts & Skills (Subtask 02)
Updated:
- `01-prompts/07-bug-fix/01-fix-with-rca.md`
- `.agents/skills/fix-with-rca/skill.md`
- `01-prompts/16-ci-cd/01-ci-cd-fix.md`
- `01-prompts/16-ci-cd/03-fix-ci-cd-and-run-scripts.md`
- `01-prompts/16-ci-cd/04-ci-cd-fix-with-release.md`
- `.agents/skills/ci-cd-fix/skill.md`
- Created `.agents/skills/ci-cd-fix-with-release/skill.md`

**Core Invariants Added:**
- Mandated GitMap pipeline AI integration for pipeline failure extraction (`##[error]`, `FAIL:`, compile failures) directly into 4-part RCA documents.
- Configured dynamic timeout waiting (`-t <seconds>`) and adaptive sleep cadences (ETA > 120s: 20-30s; 60s < ETA <= 120s: 10-20s; ETA <= 60s: 5-10s).
- Fully synchronized `ci-cd-fix-with-release` skill with the release workflow.
