# Subtask 04: GitMap Skill Uppercase Cleanup and Cluster Removal

## Objective
Purge cluster commands (`cluster`, `sc`, `ssh`), OS dev clean, VS Code repair, and eliminate uppercase commit module tags across all GitMap documentation, skills, and in-engine CLI text.

## Target Files
- `.agents/skills/gitmap/skill.md`
- `.cursor/skills/gitmap/skill.md`
- `d:/work/gitmap/cli/cmd/llm/llm_skill.go`

## Steps
1. In `gitmap/skill.md` (both `.agents` and `.cursor`):
   - Replace uppercase commit module tags with Title Case or lowercase (`Skills`, `Cloner`, `Docs`, `Release`).
   - Remove cluster-related commands from learn and examples.
   - Remove disk OS storage dev clean, port remove, and VS Code repair.
   - Ensure `gitmap kill` is the only recommended process termination command.
2. In `d:/work/gitmap/cli/cmd/llm/llm_skill.go`:
   - Replace uppercase `FEAT(MODULE)` / `FIX(MODULE)` examples with Title Case hyphenated syntax.
   - Remove cluster section.
3. Verify with GitMap searches for residual uppercase commit examples.
