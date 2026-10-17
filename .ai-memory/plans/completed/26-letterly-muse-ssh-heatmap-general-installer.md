# Execution Plan: Letterly Pull-First Sequencing, Muse SSH Delivery Protocol, Heat Map Tracking & Dual Installers

> Plan ID: `26-letterly-muse-ssh-heatmap-general-installer`
> Specification: `02-spec/21-app/21-letterly-muse-ssh-heatmap-general-installer/`
> Status: Completed
> Concurrency Engine: Parameter-Driven V6 Engine (`N=300`, `A=2`, `H=2`, `C=30`)

---

## 1. Objectives & Scope

1. **Letterly Pull-First Sequencing**: Update all Letterly formatters (`22-letterly`, `23-cursor-prompts`) and companion skills so Item 1 is Git pull via `gitmap pull` (fallback `git pull`) and conflict resolution. Re-sequence subsequent items (spec authoring -> GitMap search -> relative paths).
2. **Mobile Prompt Single-Line Output**: Ensure mobile Letterly prompts output strictly ONE continuous single-line paragraph with zero newlines.
3. **Muse SSH Key Generation & Identity Token**: Implement `<email>-muse<muse-sequence>-<operator-name>` identity, account email ingestion (or prompt fallback), and `gitmap ssh create` command.
4. **Muse Key Delivery Protocol (Fix Format Error)**: Total ban on pasting keys in chat; require downloadable file delivery (`~/.ssh/<identity>.pub`) and explicit key type designation (Authentication vs Signing).
5. **Agent Heat Map & Action Decision Logging**: Mandate `gitmap agent log` in `execute-parent-task-with-n-steps-v6` and enhance `gitmap heatmap` command in `gitmap`.
6. **Dual Installers**: Create `general-install-config.json`, `general-install.ps1`, and `general-install.sh`, and update `install.ps1` with `-General` support.
7. **GitMap General Installation Command**: Implement `gitmap cg install general` / `gitmap cd general` in `gitmap`.

---

## 2. Checklist & Subtasks

- [x] Subtask 1: Letterly prompts and skills re-sequencing (Item 1 pull first, Item 2 spec, Item 3 gitmap search, Item 4 relative paths), create `letterly-git-pull-first` skill, enforce mobile one-liner, and update `check-letterly-format.py`.
- [x] Subtask 2: Muse master prompt and onboarding spec updates (SSH identity format, account email auto-ingestion, gitmap ssh create commands, key delivery file protocol, key type designation).
- [x] Subtask 3: Heat map action logging in `execute-parent-task-with-n-steps-v6` and `gitmap heatmap` CLI command in `gitmap`.
- [x] Subtask 4: Dual installers: `general-install-config.json`, `general-install.ps1`, `general-install.sh`, and `install.ps1` `-General` flag.
- [x] Subtask 5: GitMap CLI integration for `gitmap cg install general` / `gitmap cd general`.
- [x] Subtask 6: Linter validation, tests, consolidation, and release ceremony.
