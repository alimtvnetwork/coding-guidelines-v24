# Completed Plan: Coding Guideline and Skill Optimization

## User Request (Verbatim)
# Coding Guideline and Skill Optimization: high priority instruction, non-negotiable task

For the coding guideline, try to improve the icon prompt as much as you can. Also, update the skill. And also for the literally section, we should not create any skill into the agent or also the cursor, because it runs from another tool. And also try to optimize some of the skills which we don't need as a prompt or no need for any skill to be there. So some of the skills you can find repetitive and try to find those by analyzing the `gitmap` and `Antigravity Manager (AGM)` which you see in its spec and other places. What are the skills that has been used so far? Skills which are not very effective, try to archive it at the first level to a backup folder. That is number one, and commit, and then we remove it later on. Update the image prompt or icon generate prompt, and then you start with this optimization. So do you remember the flat commit segment? I don't see this in the help text or also in the UI. Can this update and push for this? Start cleaning up the skills which is not necessary. So remember to have execute parent task in the steps. The steps make it 400. The first 100 step will be plan, then 150 for... Then 300 for the implementing of the task, I think. Or you can divide the count better way that you think. So increase the count on the steps. So check if we can improve the prompt for literally. And reading is good, then try to do a `gitmap` PN, then wait, and then synchronize the new skills and prompts with all the known repositories that you already know and previously how we have synced them to. But during this process, do not remove any file which is created by that repository, not by us. Synchronization is also an important factor. Do that.

## Canonical Specifications
- Architecture Spec: `02-spec/21-app/coding-guideline-and-skill-optimization/01-architecture-spec.md`
- Implementation Spec: `02-spec/21-app/coding-guideline-and-skill-optimization/02-implementation-spec.md`

## Summary of Accomplishments & Evidence
1. **Icon Generation Modernization:**
   - Updated `01-prompts/10-ui-and-design/01-logo-create.md`, `01-prompts/10-ui-and-design/03-svg-logo.md`, `.agents/skills/logo-and-icon-design/skill.md`, `.agents/skills/logo-create/skill.md`, `.agents/skills/svg-logo/skill.md`, and `.cursor/skills/` mirrors.
   - Enforced 100% transparent canvas (`alpha = 0`) across all SVGs and PNG exports; banned dark `<rect>` container boxes.
   - Implemented dual-mode inverted artwork: Dark Mode means inverted high-luminance strokes ($L > 0.70$) on a transparent canvas; Light Mode uses deep saturated tones ($L < 0.40$).
   - Standardized quiet zone (8-10% inner padding), 2-decimal rounded coordinates, semantic `<defs><linearGradient>`, accessibility attributes, and design token integration (`--pres-accent`, `--pres-text`, `--primary`, `--foreground`).
2. **Flat Commit Segment Documentation:**
   - Updated `AGENTS.md` §10 with a dedicated "GitMap Semantic & Flat Commit Suite" section.
   - Documented `gitmap commit "<message>"` / `gitmap cm "<message>"` as 1-step auto-stage (`git add -A`) local commit without prefix and without remote push.
   - Documented `gitmap cpc "<module> - <summary>"` (chore).
   - Contrasted flat local commits against semantic auto-pushing commits (`gitmap cpf`, `gitmap cpb`, `gitmap cpr`), reiterating that task completion requires push.
   - Updated `.agents/skills/gitmap/skill.md` and `.cursor/skills/gitmap/skill.md`.
3. **Letterly Skill Purge & Prompt Elevation:**
   - Purged all 11 external `letterly-*` skills from `.agents/skills/` and `.cursor/skills/` (22 directories removed) because Letterly runs externally.
   - Modernized all 11 Letterly prompts in `01-prompts/22-letterly/` and 11 Cursor prompt mirrors in `01-prompts/23-cursor-prompts/` to use the 400-step lifecycle budget and reference canonical skills.
   - Verified that `linter-scripts/tests/test_check_letterly_format.py` passes 8/8 tests with exit code 0.
4. **Skill Catalog Optimization & First-Level Archival:**
   - Archived 39 deprecated, duplicate, and ineffective skills into `.agents/skills-archived/` and `.cursor/skills-archived/`.
   - Created comprehensive archive catalog `readme.md` files mapping archived skills to their canonical `cg-*` replacements.
5. **Execution Prompt 400-Step Budget Alignment:**
   - Updated `execute-parent-task-with-n-steps-v6` across `.agents/skills/`, `.cursor/skills/`, and `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md` to establish $N = 400$ steps (100 plan, 150 intermediate/archival, 150 implementation/sync).
6. **Cross-Repository Fleet Synchronization:**
   - Executed `python 03-ai-scripts/38-sync-prompts-skills-scripts.py --workers 6` across 43 connected repositories.
   - Strictly enforced the non-destructive rule: zero deletions of downstream repository-created files (`+N/-0` across all repositories).
   - Successfully committed, tagged, and pushed updates across 41 repositories with automated pre-change backup branches (`backup/sync-*`).
7. **Verification & Quality Gates:**
   - Relative Git Paths: Verified `python linter-scripts/check-relative-paths.py` passed with 0 violations across 3,723 tracked files.
   - Secrets Gate: Verified `python linter-scripts/check-forbidden-strings.py` passed with 0 forbidden strings or secrets.
   - Lowercase Naming: Verified 100% lowercase conventions observed across all new and modified paths.
