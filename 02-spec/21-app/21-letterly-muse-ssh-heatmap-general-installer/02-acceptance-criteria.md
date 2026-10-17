# Acceptance Criteria: Letterly Pull-First Sequencing, Muse SSH Delivery Protocol, Heat Map Tracking & Dual Installers

## 1. Letterly Prompts & Skills Re-sequencing
- [ ] **AC-01**: In `01-prompts/22-letterly/` (`02-desktop`, `03-execute-n-steps`, `04-plan`, `05-release`, `06-cicd-fix-release`, `08-run`, `09-execute-with-verification`, `10-execute-with-release`, `11-git-pull-first`), Actionable Item 1 is ALWAYS Git pull via `gitmap pull` (fallback `git pull`) with conflict resolution.
- [ ] **AC-02**: Actionable Item 2 is ALWAYS writing spec under `02-spec/21-app/<slug>/` and enqueuing plan under `.ai-memory/plans/<slug>.md`.
- [ ] **AC-03**: Actionable Item 3 is ALWAYS searching exclusively via GitMap (`gitmap aum search`, `gitmap find`, `gitmap cat`, `gitmap ps`, `gitmap py`, `gitmap llm train`) with total ban on `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`.
- [ ] **AC-04**: Actionable Item 4 is ALWAYS strictly relative Git paths (`02-spec/...`, `.ai-memory/...`, `cmd/...`) with total ban on absolute paths or `file:///` URIs.
- [ ] **AC-05**: Companion Cursor prompts in `01-prompts/23-cursor-prompts/` mirror this exact 1-to-4 sequence using `.cursor/skills/` skill links.
- [ ] **AC-06**: `letterly-git-pull-first` skill exists in both `.agents/skills/letterly-git-pull-first/skill.md` and `.cursor/skills/letterly-git-pull-first/skill.md`.
- [ ] **AC-07**: `linter-scripts/check-letterly-format.py` and `linter-scripts/tests/test_check_letterly_format.py` pass cleanly with the updated fixed items.

## 2. Mobile Prompt Single-Line Output Invariant
- [ ] **AC-08**: `01-prompts/22-letterly/01-mobile-letterly.md` and `.agents/skills/letterly-mobile/skill.md` mandate outputting strictly ONE continuous single-line paragraph with ZERO newlines and ZERO line breaks.
- [ ] **AC-09**: Output includes gitmap pull first, relative paths mandate, and skill invocation suffix.
- [ ] **AC-10**: Cursor equivalent in `01-prompts/23-cursor-prompts/01-mobile-letterly-cursor.md` and `.cursor/skills/letterly-mobile/skill.md` satisfies the same one-liner invariant.

## 3. Muse Master Prompt & SSH Delivery Protocol
- [ ] **AC-11**: `01-prompts/27-muse-prompts/01-muse-master-prompt.md`, `.agents/skills/muse-master-prompt/skill.md`, and `.cursor/skills/muse-master-prompt/skill.md` mandate `<email>-muse<muse-sequence>-<operator-name>` identity format.
- [ ] **AC-12**: Muse automatically inspects `git config user.email` or `gh api user --jq .email` for account email; prompts operator interactively for missing variables before generation.
- [ ] **AC-13**: Key generation uses native GitMap command: `gitmap ssh create "<identity>" -y`.
- [ ] **AC-14**: TOTAL BAN on pasting keys into chat as raw text. Keys must be saved as a downloadable disk file (`~/.ssh/<identity>.pub`, `.ai-memory/ssh/<identity>.pub`) or copied via `gitmap ssh copy "<identity>"`.
- [ ] **AC-15**: Key delivery explicitly identifies the key type: **Authentication Key** vs **Signing Key**.

## 4. Heat Map & In-Flight Action Decision Logging
- [ ] **AC-16**: `01-prompts/14-execute/02-execute-parent-task-with-n-steps-v6.md`, `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`, and `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md` mandate logging file reads and modifications with reasoning via `gitmap agent log`.
- [ ] **AC-17**: `gitmap heatmap` command in `gitmap` renders touched files, action counts, and agent reasoning.

## 5. Dual Installers (Coding Guidelines vs General Repository)
- [ ] **AC-18**: `install.ps1` and `install.sh` support `-General` / `--general` switch.
- [ ] **AC-19**: `general-install-config.json` specifies strictly `.ai-memory`, `01-prompts`, `03-ai-scripts`, `.agents/skills`, `.cursor/skills`, `02-spec/01-spec-authoring-guide`, `02-spec/21-app`, and `assets` folder structure.
- [ ] **AC-20**: General installation excludes all linters, linter-scripts, and cicd test runners.
- [ ] **AC-21**: Standalone `general-install.ps1` and `general-install.sh` scripts are created and operational.

## 6. GitMap General Installation Integration
- [ ] **AC-22**: `gitmap cg install general`, `gitmap cg general`, `gitmap cg install --general`, and `gitmap cd general` trigger general repository installation.
- [ ] **AC-23**: GitMap synchronizes from local `coding-guidelines` repo if present, or executes the general installer.
