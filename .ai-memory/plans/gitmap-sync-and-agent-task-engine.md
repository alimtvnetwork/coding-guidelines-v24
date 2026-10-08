# Master Execution Plan: GitMap Fleet Sync & Native Agent Task Engine

> **Plan Identifier:** `.ai-memory/plans/gitmap-sync-and-agent-task-engine.md`  
> **Status:** APPROVED & ACTIVE  
> **Specification:** `02-spec/21-app/gitmap-sync-and-agent-task-engine/01-architecture-spec.md`  
> **Total Steps Budget:** N = 300 (Phase 1: 150, Phase 2: 150)  
> **Concurrency Capacity:** A = 2, H = 2

---

## User Request (Verbatim)

```text
What I want is a process inside git-map. We should have a command that will do this synchronization automatically. It should have default project names, see these project names there, but the user can always add their projects. Input its project in JSON. It could be simple as projects URL or JSON with the folder name and for the mode, both should work. Based on that, it can synchronize. Based on that git-map, it will synchronize faster rather than Python script. Think about how to write all this logic regarding this and improve the git-map so that we have functionality to sync files and folders. We need to identify how to implement all of that. The agent part, agent creation using the SQLite, can be improved using the `gitmap` on execution command exploration. The Python script currently creates the SQLite information. Based on that, the git-map should have the command that would deal with this situation smoothly and faster. Try to have these codes implemented first, release first, and then add this in the skills and in the prompt and update everything.
```

---

## Subtask Decomposition

| Task ID | Title | Owned Subtask File | Target Subsystems | Owner Role | Status |
|---|---|---|---|---|---|
| `Task-01` | GitMap Codebase Exploration & Go Command Design | `.ai-memory/plans/subtasks/gitmap-sync-and-agent-task-engine/01-gitmap-codebase-exploration-and-design.md` | `d:\work\gitmap` CLI packages & DB engine | `DiscoverySubagents` | PENDING |
| `Task-02` | Implement `gitmap sync` & `gitmap task` in Go | `.ai-memory/plans/subtasks/gitmap-sync-and-agent-task-engine/02-implement-sync-and-task-in-gitmap.md` | `d:\work\gitmap/cmd/*`, `pkg/*` | `GoEngineWorker` | PENDING |
| `Task-03` | Release GitMap, Update Prompts/Skills & Fleet Sync | `.ai-memory/plans/subtasks/gitmap-sync-and-agent-task-engine/03-release-gitmap-and-update-prompts-skills.md` | `01-prompts/`, `.agents/skills/`, 43 Repos | `FleetSyncWorker` | PENDING |

---

## Non-Negotiable Invariants & Boundaries

1. **Spec 21-25 Exclusion:** Spec 21 is target-repo exclusive. Never overwrite target app specs.
2. **Additive-Only Scripts & Bump Script Protection:** Never clobber target repo bump scripts or modified AI scripts.
3. **Mandatory Completion Invariant:** No Push = Not Done. If code is not committed and pushed upstream to GitHub, the task is strictly considered INCOMPLETE and NOT DONE.
4. **Tool Primacy:** Code exploration exclusively via GitMap (`gitmap aum search`, `gitmap find`, `gitmap cat`, `gitmap ps`); total ban on `rg`, `ripgrep`, `grep`, `Select-String`.
5. **Strict Relative Paths:** No drive letters (`D:\`, `C:\`) or `file:///` URIs inside repository documents, plans, or release notes.
