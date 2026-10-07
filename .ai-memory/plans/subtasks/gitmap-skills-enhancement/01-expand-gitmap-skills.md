# Subtask: Expand GitMap Skills Definition & Matrix

## Metadata
- **Subtask ID:** `01-expand-gitmap-skills`
- **Parent Plan:** `.ai-memory/plans/gitmap-skills-enhancement.md`
- **Spec Reference:** `02-spec/21-app/gitmap-skills-enhancement/01-gitmap-skills-architecture.md`
- **Status:** Pending

## Objective
Author the enhanced GitMap skill file with full command coverage and negative vs positive enforcement matrices. Mirror the skill identically across:
1. `.agents/skills/gitmap/skill.md`
2. `.cursor/skills/gitmap/skill.md`

## Required Sections to Incorporate
1. **Authentication & GitHub Access:**
   - `gitmap login`
   - `gitmap login --web`
   - `gitmap login --token <PAT>`
   - `gitmap login --status`
   - `gitmap token list`
   - `gitmap logout`
2. **Workspace & Repository Status Inspection:**
   - `gitmap status` / `gitmap st`
   - `gitmap status --dirty` (`-d`)
   - `gitmap status --ahead` / `--behind`
   - `gitmap status --json` (`-j`)
   - `gitmap status --table`
   - `gitmap has-any-updates` (`gitmap hau`, `gitmap hac`)
   - `gitmap latest-branch` (`gitmap lb`)
   - `gitmap watch` (`gitmap w`)
3. **High-Speed Script & Shell Runners:**
   - `gitmap py "<code-or-script>"` / `gitmap py -c "<expression>"`
   - `gitmap pwsh "<cmd>"` / `gitmap ps "<cmd>"` / `gitmap ps -c "<cmd>"`
   - `gitmap bash "<cmd>"` / `gitmap sh "<cmd>"` / `gitmap bash -c "<cmd>"`
4. **Secret Offloading & Script Storage:**
   - `gitmap rs file <path>` / `gitmap rs folder <path>` / `gitmap rs text "<secret>"`
   - `gitmap rc file <path>` / `gitmap rc folder <path>` / `gitmap rc text "<content>"`
5. **High-Speed File Discovery & In-Memory Cat:**
   - `gitmap aum search "<pattern>" [dir] [--ext <ext>]`
   - `gitmap search "<query>"`
   - `gitmap find "<pattern>"`
   - `gitmap find-files <name>` (`ff`)
   - `gitmap list-files [dir]` (`lf`)
   - `gitmap cat <filepath>`
6. **Semantic Git Operations:**
   - `gitmap cpf "<module> - <summary>"`
   - `gitmap cpb "<module> - <summary>"`
   - `gitmap cpr "<module> - <summary>"`
   - `gitmap pcp "<module> - <summary>"`
   - `gitmap pull` / `gitmap pull-all`
7. **CI/CD Self-Healing & Diagnostics:**
   - `gitmap pipeline-ai status -t <eta>`
   - `gitmap pipeline error-logs` (`gitmap pe`, `gitmap pe -t`)
   - `gitmap pipeline purge`
8. **Negative vs Positive Substitution Matrix:**
   - Detailed mapping showing strictly forbidden tools (`rg`, `ripgrep`, `grep`, `Select-String`, raw unbuffered `cat`, `python`, `powershell`, colons in commit messages) replaced by GitMap equivalents.

## Acceptance Criteria
- [ ] `.agents/skills/gitmap/skill.md` updated and validated.
- [ ] `.cursor/skills/gitmap/skill.md` updated and synchronized.
- [ ] All paths inside skills are strictly relative paths.
