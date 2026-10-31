# Master Execution Plan: Coding Guideline and Folder Structure Improvement

> **Plan File:** `.ai-memory/plans/pending/29-coding-guideline-and-folder-structure-improvement.md`  
> **Slug:** `coding-guideline-and-folder-structure-improvement`  
> **Status:** Completed on 30-Oct-26. What-to-read is version 5.0.0. The named obsolete `02-spec/21-app` prompt folders are already gone. Later specs stay. `.gitkeep` stays only in empty folders. `syncbridge` exists. Archive folders are excluded from sync. The worker catalog is `02-spec/19-main-worker-service/100-improvement-possibilities.md`. No outside sync was run.  
> **Budget:** 400 Steps | Concurrency: A = 3, H = 2 (Capacity: 6)  
> **Spec Reference:** `02-spec/21-app/29-coding-guideline-and-folder-structure-improvement/`  

---

## User Request (Verbatim)

```markdown
# Coding Guideline and Folder Structure Improvement: high priority instruction, non-negotiable task

In the coding guideline, there are several factors. Those are very wrong. Come with each one of them. One of the factors is that the what to read folder is not very professional or not very good, so you need to work on it. The second factor is that you need to work on this for the coding guideline. So for the coding guideline, the folder structure in the AI memory is not updated in the what to read. And also there are a lot of plan files in the coding guideline, which should be cleared for now. And also for the AI instruction, whenever a folder has a file, it should actually remove that .gitkeep file. So that needs to be mentioned carefully. In the ambiguity section, make sure that it is clear. In coding guideline, we don't need the ambiguity. So in cases you have an ambiguity question, so you ask for some problem issues like console sync, file sync. So whatever the problem that you have in the ambiguity question, can you please try to implement some of those by creating another package on top of it by connecting this sync and stuff? Can you please do that so that our logger and everything is perfect? That's another thing. So what to read, you need to correct and check that recently we should have the latest `gitmap` skills should be added in our skill section and also find outdated skills in the skill section for agents and also cursor. Try to remove the outdated skill and any old files, old prompt, try to move this to archive folder, and do not sync those and make sure those old prompts should be removed from the synced repositories. According to a guideline old file, try to remove that as well. In the what to read, and the folder structure, these are pretty old folder structure. It is not following the current folder structure. Inside the AI memory, you need to correct it very carefully right now. And also some of the repositories I saw that old folder structure remains. You need to correct it and if there is anything that we can improve in the folder structure, first suggest it. Do not sync outside for now. Let's fix inside in the coding guideline for now. Then we can synchronize later on. Rather than brute forcing, we will think carefully and then we will sync. Do you understand? Also, the main worker service that is spec, check that. Find 100 possibilities that could be improved inside that spec. And also the folder 21 contains create 100 prompt improvements. These are the files we don't need. The app folder for the coding guideline should be clean. Let's improve the coding guideline first and clear out the app issues, app Database (Database (Db)). All the index zero one or overview files should be README file, and that needs to be corrected everywhere. I do see still overview files and old prompts moved to archive folder. Do not sync those. And also the change history, do not sync. So these are the things I think we need to correct. Change history, we don't need to sync with other repositories. Remember that. Archive folder do not need to have. So inside the archive, all the old stuff, just put it there. So this should not be synced. Remember that. Can you please follow through?
```

---

## Subtask Schedule

| Subtask ID | Code | Subtask Title | Assigned Owner | Owned Files | Status |
|:---:|:---:|:---|:---:|:---|:---:|
| `sub-161904-task-01` | **Task-01** | Spec Authoring & Subtask Planning | Lead Orchestrator | `02-spec/21-app/29-coding-guideline-and-folder-structure-improvement/*`, `.ai-memory/plans/pending/29-*`, `.ai-memory/plans/subtasks/29-*` | DONE |
| `sub-161913-task-02` | **Task-02** | Professionalize What-to-Read & Modernize Folder Structure | Worker 01 | `.ai-memory/what-to-read.md`, `.ai-memory/folder-structure.md`, `02-spec/01-spec-authoring-guide/02-folder-structure.md` | PENDING |
| `sub-161914-task-03` | **Task-03** | Purge Outdated Plans & Clean App Spec Hierarchy | Worker 02 | `02-spec/21-app/`, `.ai-memory/plans/pending/`, `.ai-memory/plans/subtasks/`, `.ai-memory/plans/readme.md` | PENDING |
| `sub-161914-task-04` | **Task-04** | Remove .gitkeep Files from Populated Folders & Codify Rule | Worker 03 | Repository `.gitkeep` files, `.ai-memory/strictly-avoid.md` | PENDING |
| `sub-161914-task-05` | **Task-05** | Ambiguity Resolution & AppLogger Sync Bridge Package | Worker 01 | `04-code/golang/pkg/applogger/syncbridge/*`, `.ai-memory/ambiguous-questions/` | PENDING |
| `sub-161919-task-06` | **Task-06** | Modernize GitMap Skills & Archive Outdated Skills | Worker 02 | `.agents/skills/`, `.cursor/skills/`, `.agents/skills-archived/`, `.cursor/skills-archived/` | PENDING |
| `sub-161919-task-07` | **Task-07** | Archive Old Files/Prompts & Sync Boundary Isolation | Worker 03 | `01-prompts/`, `06-archive/`, `03-ai-scripts/38-sync-prompts-skills-scripts.py`, `.agents/scripts/38-sync-prompts-skills-scripts.py` | PENDING |
| `sub-161919-task-08` | **Task-08** | 100 Improvement Possibilities for Main Worker Service Spec | Worker 01 | `02-spec/19-main-worker-service/100-improvement-possibilities.md`, `02-spec/19-main-worker-service/readme.md` | PENDING |
| `sub-161919-task-09` | **Task-09** | Convert Overview/Index Files to Standard Readme.md | Worker 02 | Targeted overview/index files across `02-spec/` and `.ai-memory/` | PENDING |
| `sub-161919-task-10` | **Task-10** | Final Verification, Heatmap Closure & Atomic Commit | Lead Orchestrator | Workspace-wide git staging, `.ai-memory/heatmap.db` | PENDING |
