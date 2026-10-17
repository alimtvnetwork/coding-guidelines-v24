# Plan: Unified Muse Master Onboarding & Execution Prompt (V6 Architecture)

**Version:** 1.0.0
**Status:** Pending
**Parent Plan:** `.ai-memory/plans/pending/19-unified-muse-master-prompt-v6.md`
**Canonical Spec:** `02-spec/21-app/16-unified-muse-master-prompt-v6/`
**Subtasks:** `.ai-memory/plans/subtasks/19-unified-muse-master-prompt-v6/`

---

## 1. Goal

Consolidate all separate Muse prompts in category `27-muse-prompts` (`01-muse-master-prompt.md`, `02-muse-execute-in-a-step.md`, `03-muse-master-with-repos.md`, `04-muse-master-with-repos-v2.md`, `05-muse-master-prompt-v2.md`) into a single unified, self-bootstrapping, execute-N-steps-V6 master prompt.

---

## 2. Key Deliverables & Requirements

1. **V6 Execution Parameters & Checklist:** `N=300`, `A=2`, `H=2`, `C=30`, Top-Instruction Priority Mandate, slash commands, actionable agent checklist, mandatory subagent spawning gate (`invoke_subagent`).
2. **Authentication Bootstrap with SSH Primacy:** Check existing login status; if already authenticated, move immediately to next steps; if not, prefer SSH with browser fallback.
3. **GitMap Installation, Update & Fleet Scan:** Include quick install command (`eval "$(curl -fsSL https://raw.githubusercontent.com/alimtvnetwork/gitmap-v28/main/install-quick.sh)"`), update GitMap first, execute `gitmap scan` to index all workspace repos.
4. **Canonical Starting Point & 14+ Fleet Registry:** Starting point URL `https://github.com/alimtvnetwork/coding-guidelines-v24.git` (stale repo warning), full 14+ repo registry (`gitmap-v28`, `coding-guidelines-v24`, `Antigravity-Manager`, `alim-seo-writing`, `go-email-reader`, `image-generate-v2`, `white-presentation-v1`, `wp-exam-v2`, `cat-my-v12`, `story-telling-v1`, `minecraft-alim-alima`, `macro-ahk`, `scripts-fixture`, `repo-cache`, `repo-secrets`), auto-cloning missing repos.
5. **Pull-All (`gitmap pa`) & Conflict Prevention:** Run `gitmap pa`, ensure clean working trees, zero merge conflicts.
6. **Repo-Cache Integration & Dedicated Side Chats:** `repo-cache/NN-<slug>` sequencing, dedicated side chats named `<repo> repo` and "all summary" chat.
7. **Task Confirmation & Multi-Agent Execution Engine:** Turn 1 confirmed task breakdown, RUNNING + ETA declaration, 5-minute status pings, SQLite task manager, pull first, mandatory subagents (`A=2, H=2`) in disjoint file boxes, pre-commit secrets gate, large-file 15MB 7-Zip split gate, atomic commit via GitMap hyphen format (`gitmap cpf "<module> - <summary>"`), CI telemetry (`gitmap pe -t`), and Muse Memory IO re-sync protocol.
8. **Skills & Catalogs Synchronization:** Update `muse-master-prompt` in `.agents/skills/` and `.cursor/skills/`, update `01-prompts/27-muse-prompts/readme.md`, `01-prompts/readme.md`, `.ai-memory/prompts.md`, and `02-spec/21-app/readme.md`.

---

## 3. Subtasks Breakdown

- [01-spec-and-architecture.md](subtasks/19-unified-muse-master-prompt-v6/01-spec-and-architecture.md): Authored canonical architecture spec and acceptance criteria in `02-spec/21-app/16-unified-muse-master-prompt-v6/`.
- [02-unified-master-prompt.md](subtasks/19-unified-muse-master-prompt-v6/02-unified-master-prompt.md): Author the consolidated master prompt in `01-prompts/27-muse-prompts/01-muse-master-prompt.md` and harmonize `05-muse-master-prompt-v2.md`.
- [03-skills-and-catalogs.md](subtasks/19-unified-muse-master-prompt-v6/03-skills-and-catalogs.md): Update companion skills and sync index catalogs.
