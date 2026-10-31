# Subtask 30.2 — V6 Execute Prompts & Skills Audit

> **Parent Plan:** `.ai-memory/plans/pending/30-root-cause-analysis-and-fix.md`  
> **Target Package:** `01-prompts/14-execute/`, `.agents/skills/`  
> **Status:** Pending  

---

## 1. Objectives
1. Scan `01-prompts/14-execute/` and `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`.
2. Locate all examples of commit commands such as `gitmap cpf "(MODULE): <summary>"` or `gitmap cpf "<module> - <summary>"`.
3. Ensure no prompt suggests using `commit` as a module name.
4. Update prompt guidelines to explicitly advise using the target subsystem or module (e.g. `Ui`, `Auth`, `Cli`, `Scanner`, `Guidelines`), or omitting the module if not applicable.
