# Subtask 31.3: GitMap Skill Sanitization & Overhaul

> **Parent Plan:** `31-coding-guideline-corrections-and-gitmap-skill-updates`  
> **Status:** Completed  
> **Target:** `.agents/skills/gitmap/skill.md` & `.cursor/skills/gitmap/skill.md`  
> **Date:** 2026-10-27  

---

## Objectives
1. Sanitize lines 715–740 in `gitmap/skill.md`:
   - Replace ALL-CAPS module tags (`SKILLS`, `CLONER`, `DOCS`, `RELEASE`) with Title Case (`Skills`, `Cloner`, `Docs`) or lowercase (`skills`, `cloner`, `docs`).
   - Fix release commit example from `gitmap cpr "RELEASE - ..."` to `gitmap cpr "vX.Y.Z - <scope>"`.
2. Eliminate duplicate commit tables and repetitive descriptions across the file.
3. Purge bloat:
   - Remove `gitmap disk`, `gitmap os storage`, `gitmap os dev-clean`.
   - Remove `gitmap vscode repair`.
   - Remove `gitmap starbar client ls` / `client list`.
   - Remove `gitmap cluster worker`, `gitmap cluster status`.
4. Keep and format `gitmap cargo` commands (`cargo build`, `cargo test`, `cargo run`, `cargo install`).
5. Document native process termination (`gitmap proc kill <pid> [--force]`).
6. Mirror updates to `.cursor/skills/gitmap/skill.md`.
