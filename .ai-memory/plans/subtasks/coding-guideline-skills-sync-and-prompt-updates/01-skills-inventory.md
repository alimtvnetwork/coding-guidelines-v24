# Subtask 01 — gitmap skills inventory (W1)

Owner: W1. Read-only except the inventory section in the parent plan.

## What to do

1. Inventory gitmap skills present in THIS repo: `01-prompts/26-gitmap/`
   (`01-gitmap-core-engine.md` + readme), `.agents/skills/` (full listing),
   `.cursor/skills/` (full listing), and any skill docs.
2. Note versions/dates where the files carry them.
3. Diff against the ingested set: `gitmap`, `gitmap-agent-orchestrator` +
   6 subsystems (scanner-and-cloner, ssh-cluster-fleet, split-db-engine,
   pipeline-and-diagnostics, developer-hygiene-and-agy, macro-automation-engine).
4. Append a "Skills inventory" section to
   `.ai-memory/plans/coding-guideline-skills-sync-and-prompt-updates.md`
   naming: NEW skills not in the ingested set, CHANGED skills (what changed),
   and anything the codebase should adopt.

## Rules

- TOTAL BAN on rg/grep/git grep/Select-String → `gitmap aum search`,
  `gitmap find`, `gitmap cat` (export `TERM=xterm-256color`).
- Do not modify any skill or prompt file; inventory only.

## Acceptance

- [ ] Inventory section in the parent plan lists new/changed items by name with evidence (paths + versions).
