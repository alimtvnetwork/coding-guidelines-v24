# Acceptance Criteria: Mobile Letterly Formatter Modernization, Tree Swarm Hierarchy & Dynamic Task Join

## 1. Traceable Acceptance Matrix

| Criterion ID | Target Component | Verifiable Requirement | Verification Method |
| :--- | :--- | :--- | :--- |
| **AC-01** | `01-mobile-letterly.md` & `01-mobile-letterly-cursor.md` | Mobile formatters begin with `[/goal] [/learn] from [gitmap] [/plan]`, use chained header `# <task title> - high priority instruction, non-negotiable task (task: <task title>, slug: <task-slug>): `, blend pull-first, spec authoring, gitmap search, relative paths, verbatim input, and extracted discrete understandable tasks into strictly ONE continuous line with ZERO newlines. | Inspection & regex line count check (`wc -l == 1` for output format). |
| **AC-02** | `07-mobile-cicd-fix-letterly.md` & `07-mobile-cicd-fix-letterly-cursor.md` | Mobile CI/CD fix formatters conform to the exact chained header, `[/goal] [/learn] from [gitmap] [/plan]` prefix, live pipeline diagnosis via `gitmap pe -t`, 4-part RCA, atomic commit, minor release ceremony, and extracted task list in ONE continuous line with ZERO newlines. | Inspection & line count check. |
| **AC-03** | `.agents/skills/letterly-mobile/` & `.cursor/skills/letterly-mobile/` | Companion Antigravity and Cursor skills mirror the professional chaining, skill train, and single-line format verbatim. | File diff & inspection. |
| **AC-04** | `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md` | V6 execution prompt integrates: (1) Tree Swarm Hierarchy (Root Orchestrator -> Branch Coordinators -> Leaf Workers & Assistants); (2) Phase 1 Brain & Comprehensive Roadmap creation up front; (3) Dynamic Agent Join / Collaborative Assistance protocol whenever an agent finishes early and uncompleted tasks remain. | Inspection of Sections 1, 6, 7, and R5/R6. |
| **AC-05** | `.agents/skills/execute-parent-task-with-n-steps-v6/` & `.cursor/` | Antigravity and Cursor V6 companion skills mirror the Tree Swarm, Brain Roadmap, and Dynamic Join protocols. | File diff & inspection. |
| **AC-06** | `.agents/skills/gitmap/` & `.cursor/skills/gitmap/` | GitMap companion skills document `gitmap agent subtask join` and `gitmap agent task tree` alongside split-DB and heat map operations. | Inspection of command table & agent guide. |
| **AC-07** | Relative Paths Mandate | Zero absolute filesystem paths (`C:\...`, `/home/...`) and zero `file:///` URIs in any modified or newly authored files. Only relative git paths used. | `python linter-scripts/check-relative-paths.py .` |
| **AC-08** | Lowercase Filename Hygiene | All modified, generated, and referenced file paths strictly lowercase without exception. | Workspace directory inspection. |
| **AC-09** | Mechanical Linter Verification | `check-letterly-format.py` tests pass (`python -m unittest linter-scripts/tests/test_check_letterly_format.py`), `check-prompts-loaded.py` passes cleanly. | Script execution exit code 0. |
| **AC-10** | Mandatory Completion Invariant | Working tree clean, changes staged, committed atomically via GitMap (`gitmap cpf "<module> - <summary>"`), pushed upstream to `main`, and minor version released. | `git status --porcelain` empty & git push confirmed. |

---

## 2. Evidence Verification Contract

Each criterion must be verified with concrete terminal output and exit code 0. No assumptions or assertions without evidence.
