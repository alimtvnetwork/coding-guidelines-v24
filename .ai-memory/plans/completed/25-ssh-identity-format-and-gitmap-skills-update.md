# Execution Plan: SSH Identity Format & GitMap Skills Fleet Synchronization

> Plan ID: `25-ssh-identity-format-and-gitmap-skills-update`
> Specification: `02-spec/21-app/20-ssh-identity-format-and-gitmap-skills-update/`
> Status: In Progress
> Execution Engine: V6 Parameter-Driven Engine (`N=300`, `A=2`, `H=2`, `C=30`)

---

## 1. Objectives & Scope

1. **SSH Key Identity Format**: Implement `<email>-muse<muse-sequence>-<operator-name>` (e.g. `alex.tone@company.com-muse-10-opRoden` / `alex.john@company.com-muse-10-opRoden`) in Step 2 of `01-muse-master-prompt.md`, `muse-master-prompt` skills, and `03-onboarding-protocol-spec.md`.
2. **Interactive Variable Ingestion**: Mandate interactive prompting for `<email>`, `<muse-sequence>`, and `<operator-name>` if any are unknown.
3. **Git Configuration Command**: Configure `git config --global user.email "<email>"` and `git config --global user.name "<operator-name> (Muse <muse-sequence>)"`.
4. **GitMap Skills Currency**: Update `.agents/skills/gitmap/skill.md` and `.cursor/skills/gitmap/skill.md` with full documentation for `gitmap llm train`, `gitmap llm train muse --days 7`, `gitmap agent`, and new commands.
5. **Absence of Deprecated Muse Prompts**: Ensure all target repositories have 0 stale Muse prompts (`02`, `03`, `04`, `05`).
6. **Fleet Deployment**: Deploy and push changes across all 43 connected repositories.

---

## 2. Checklist & Subtasks

- [x] Subtask 1: Author architecture spec and plan.
- [ ] Subtask 2: Update `01-muse-master-prompt.md`, `muse-master-prompt` skills, and onboarding spec with SSH identity format.
- [ ] Subtask 3: Update `gitmap` skills in `.agents` and `.cursor` with latest commands.
- [ ] Subtask 4: Regenerate sync assets, run linters, and commit/push `coding-guidelines`.
- [ ] Subtask 5: Execute live fleet synchronization across all 43 repositories and verify completion.
