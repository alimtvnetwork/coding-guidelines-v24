# Master Audit Ledger: Coding Guideline Corrections & GitMap Skill Updates (v33)

## Task Metadata
- **Task Slug:** `33-coding-guideline-corrections-and-gitmap-skill-updates`
- **Specification Directory:** `02-spec/21-app/33-coding-guideline-corrections-and-gitmap-skill-updates/`
- **Plan File:** `.ai-memory/plans/pending/33-coding-guideline-corrections-and-gitmap-skill-updates.md`
- **Subtasks Directory:** `.ai-memory/plans/subtasks/33-coding-guideline-corrections-and-gitmap-skill-updates/`
- **Related Spec:** `02-spec/21-app/32-coding-guideline-corrections-and-gitmap-skill-updates/`

## Audit Inventory & Objectives

| Item | Area | Description | Target Files | Status |
| :--- | :--- | :--- | :--- | :--- |
| **01** | GitMap Clone | Fix `gitmap clone <repo>` bare slug resolution fallback to active GitHub owner | `cli/cmdclone/gh_resolver.go`, `cli/cmdclone/gh_resolver_test.go` | Pending |
| **02** | Context Menu | Add Installer UI entry to context menu and graceful non-git directory handling in `show-log` | `cli/cmdinstall/installctxentries.go`, `cli/cmdlogviewer/server.go` | Pending |
| **03** | Show Log UI | Serve PNG favicon/manifest for authentic titlebar icon, compact header controls, and diff tool picker | `cli/cmdlogviewer/favicon.go`, `cli/cmdlogviewer/ui_assets_html.go`, `cli/cmdlogviewer/ui_assets_script.go` | Pending |
| **04** | GitMap Skills | Purge cluster commands, OS dev clean, VS Code repair, and eliminate uppercase commit module tags | `.agents/skills/gitmap/skill.md`, `.cursor/skills/gitmap/skill.md`, `cli/cmd/llm/llm_skill.go` | Pending |
| **05** | Guidelines Skill | Enrich `follow-coding-guidelines` skill with explicit `/learn` directive, grounded checklist, and prompt sequence | `.agents/skills/follow-coding-guidelines/skill.md`, `.cursor/skills/follow-coding-guidelines/skill.md` | Pending |
| **06** | Skills Reduction | Survey all active skills, identify candidate prompt wrappers for deprecation/consolidation, and present reduction matrix | `02-spec/21-app/33-coding-guideline-corrections-and-gitmap-skill-updates/02-skills-reduction-matrix.md` | Pending |
| **07** | E2E & Release | Verify binaries, build tests, run linters, and prepare SemVer release notes | `version.json`, `package.json`, `changelog.md` | Pending |
