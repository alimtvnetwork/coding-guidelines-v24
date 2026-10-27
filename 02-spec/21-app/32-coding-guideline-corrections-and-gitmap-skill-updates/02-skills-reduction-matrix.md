# Skills Reduction Matrix: Archival & Consolidation Audit

- **Specification Directory:** `02-spec/21-app/32-coding-guideline-corrections-and-gitmap-skill-updates/`
- **Source Repositories:** `coding-guidelines` and `gitmap`
- **Target Archive Locations:** `.agents/skills-archived/` and `.cursor/skills-archived/`

---

## 1. Inventory of Skills to Consolidate & Archive

| Skill Directory | Original Purpose | Consolidated Canonical Home | Status |
| :--- | :--- | :--- | :--- |
| `major-bump` | Wrapper for `01-prompts/17-release-management/01-major-bump.md` | `release` | Move to `skills-archived/` |
| `minor-bump` | Wrapper for `01-prompts/17-release-management/02-minor-bump.md` | `release` | Move to `skills-archived/` |
| `patch-bump` | Wrapper for `01-prompts/17-release-management/03-patch-bump.md` | `release` | Move to `skills-archived/` |
| `version-bump-docs` | Wrapper for `01-prompts/17-release-management/05-version-bump-docs.md` | `release` | Move to `skills-archived/` |
| `logo-create` | Wrapper for `01-prompts/12-logo-banner-and-mockups/01-logo-create.md` | `logo-and-icon-design` | Move to `skills-archived/` |
| `svg-logo` | Wrapper for `01-prompts/12-logo-banner-and-mockups/02-svg-logo.md` | `logo-and-icon-design` | Move to `skills-archived/` |
| `theming-guidelines` | Wrapper for `01-prompts/13-themes/01-theming-guidelines.md` | `react-ui-theming-design` | Move to `skills-archived/` |
| `update-theming` | Wrapper for `01-prompts/13-themes/02-update-theming.md` | `react-ui-theming-design` | Move to `skills-archived/` |
| `sync` | Wrapper for `01-prompts/18-multi-repo-sync/01-sync.md` | `prompts-and-skills-sync` | Move to `skills-archived/` |
| `sync-codebase` | Wrapper for `01-prompts/18-multi-repo-sync/02-sync-codebase.md` | `prompts-and-skills-sync` | Move to `skills-archived/` |
| `sync-other-codebase` | Wrapper for `01-prompts/18-multi-repo-sync/03-sync-other-codebase.md` | `prompts-and-skills-sync` | Move to `skills-archived/` |
| `next-steps` | Wrapper for `01-prompts/09-next-step-and-pending-plans/01-next-steps.md` | `check-pending-improvements` | Move to `skills-archived/` |
| `pending-tasks` | Wrapper for `01-prompts/09-next-step-and-pending-plans/02-pending-tasks.md` | `check-pending-improvements` | Move to `skills-archived/` |

---

## 2. Preservation Guarantee

1. **Non-Destructive Archival:** No skill content is permanently deleted. Archived skills are placed into `.agents/skills-archived/<skill-name>/` and `.cursor/skills-archived/<skill-name>/` where git history and reference capabilities remain accessible.
2. **Canonical Skill Enhancement:** All actionable instructions, checklists, and non-negotiables from the archived skills are unified into their respective canonical parent skills.
3. **Cross-Repo Parity:** Both `coding-guidelines` and `gitmap` receive identical skill archival updates to preserve fleet symmetry.
