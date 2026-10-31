# Plan: 24-coding-guideline-and-prompt-execution-improvement

## Task Information
- **Title:** Coding Guideline and Prompt Execution Improvement
- **Slug:** `24-coding-guideline-and-prompt-execution-improvement` (Canonical Slug: `coding-guideline-and-prompt-execution-improvement`)
- **Status:** Completed on 30-Oct-26. "reduce credit" is absent from `01-prompts`. The motto and `XX-<slug>` rule are in `AGENTS.md` and `01-prompts/22-letterly/02-desktop-letterly.md`. `.agents/skills/gitmap/skill.md` already shows command examples. `01-prompts/v4/` is archived, so that path was not edited.
- **Target Files:**
  - `01-prompts/22-letterly/*.md`
  - `01-prompts/23-cursor-prompts/*.md`
  - `01-prompts/27-muse-prompts/02-muse-generic.md`
  - `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md`
  - `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`
  - `linter-scripts/tests/test_check_letterly_format.py`
  - `AGENTS.md`
  - `.agents/skills/gitmap/skill.md`
  - `.agents/skills/gitmap-*/skill.md`
  - `02-spec/21-app/24-coding-guideline-and-prompt-execution-improvement/`
  - `.ai-memory/plans/subtasks/coding-guideline-and-prompt-execution-improvement/`

---

## User Request (Verbatim)

```text
Few of the serious issues in the coding guideline and especially in the latterly execute prompt, execute in n steps. The problem here is that in the prompt, in one step it says, "Try to make the plan to reduce credit." Try to take that line out. We don't need the reduction of the credit. We want the work plan for improved work or less friction. That should be the plan's main motto, not the reduction of the credit. The second thing is that when it writes about the planning, there's no need to include the execute steps. When it talks about the plans in the point number two, it gives the slug, right? So there is one more point. The slug needs to be written as XX hyphen and the slug. The XX would be the sequence of that folder. So that needs to be clarified. One more point about the `gitmap` skills. The `gitmap` skills actually refers to the command, but it does not refer or shows how the command works. So the problem is every time the AI needs to run the help command in order to understand how that command works. So I request you to add some sample so that the AI knows how to use that command effectively. Does this make sense? Can you please improve some of the parts in the `gitmap` skills, latterly prompt, execute in steps, things like that.

## task: Coding Guideline and Prompt Execution Improvement
## slug: coding-guideline-and-prompt-execution-improvement

# Actionable Items Must Follow Non-Negotiable

1. Git pull BEFORE starting any work — first try `gitmap pull <repo path>` (or run `gitmap pull` in repository folder); if that fails, fall back to regular `git pull` and resolve any merge conflicts so the repository is completely up to date and clean before starting.
2. Write spec under 02-spec/21-app/coding-guideline-and-prompt-execution-improvement/ and enqueue plan task in .ai-memory/plans/coding-guideline-and-prompt-execution-improvement.md (subtasks in .ai-memory/plans/subtasks/coding-guideline-and-prompt-execution-improvement/) first
3. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String
4. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well
5. Remove the line "Try to make the plan to reduce credit" from the prompt execution steps.
6. Ensure the plan's main motto is focused on improved work or less friction, not credit reduction.
7. Clarify that the slug should be written as XX hyphen and the slug, where XX is the sequence of that folder.
8. Add examples to `gitmap` skills to demonstrate how commands work, reducing the need for AI to run help commands.
```

---

## Subtask Allocation (Tree Swarm Hierarchy)

| Task ID | Title | Owner | Status | Owned Files |
| :--- | :--- | :--- | :--- | :--- |
| `Task-01` | Prompt Execution Standards & Credit Removal | Worker 01 | DONE | `01-prompts/22-letterly/*`, `01-prompts/23-cursor-prompts/*`, `01-prompts/27-muse-prompts/02-muse-generic.md`, `linter-scripts/tests/test_check_letterly_format.py` |
| `Task-02` | Slug Convention (`XX-<slug>`) & Planning Decoupling | Worker 02 | DONE | `AGENTS.md`, `01-prompts/22-letterly/01-mobile-letterly.md`, `01-prompts/23-cursor-prompts/01-mobile-letterly-cursor.md`, `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md`, `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` |
| `Task-03` | GitMap Skills Examples & Companion Alignment | Worker 03 | DONE | `.agents/skills/gitmap/skill.md`, `.agents/skills/gitmap-developer-hygiene-and-agy/skill.md`, `.agents/skills/gitmap-scanner-and-cloner/skill.md`, `.agents/skills/gitmap-split-db-engine/skill.md`, `.agents/skills/gitmap-ssh-cluster-fleet/skill.md` |
