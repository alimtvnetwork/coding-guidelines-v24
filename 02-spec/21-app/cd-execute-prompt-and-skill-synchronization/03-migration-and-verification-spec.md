# Migration and Verification Specification: CD Execute Prompt & Skill Synchronization

## 1. Migration Strategy
1. **Phase 1: Update Upgrade Engine (`03-ai-scripts/47-upgrade-cg-execute-to-v6.py`):**
   - Correct invalid command references (`gitmap heatmap discuss` -> `gitmap agent heatmap` / `gitmap heatmap status`).
   - Add Step 0 Preflight Heatmap commands (`gitmap rc heatmap tables`, `gitmap heatmap status`).
   - Add Prompt 38 (`38-map-dispatch-and-handler-lookup.md`) to the upgrade generator and mapping tables.
   - Run the upgrade script to upgrade all 38 prompts in `01-prompts/15-cg-execute/`, `01-prompts/v4/15-cg-execute/`, `.agents/skills/`, and `.cursor/skills/`.
2. **Phase 2: Author Canonical Spec & Map Dispatch Guideline:**
   - Write `02-spec/02-coding-guidelines/01-cross-language/35-map-dispatch-and-handler-lookup.md`.
   - Update indices: `02-spec/02-coding-guidelines/01-cross-language/readme.md`, `01-prompts/readme.md`, `.ai-memory/prompts.md`.
3. **Phase 3: Update `execute-parent-task-with-n-steps-v6` & Master `cg-execute` Skills:**
   - Update `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md` and `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`.
   - Update `.agents/skills/cg-execute/skill.md` and `.cursor/skills/cg-execute/skill.md`.
4. **Phase 4: Run Targeted Verification:**
   - Verify all paths are strictly relative.
   - Run relative path linters and prompt loaded linters.
   - Confirm zero broken links and strictly lowercase naming.
5. **Phase 5: Fleet Synchronization:**
   - Run `python 03-ai-scripts/38-sync-prompts-skills-scripts.py` across all connected repositories.
   - Verify clean sync results without touching Spec 21 or bump scripts.
6. **Phase 6: Atomic Commit & Remote Push:**
   - Stage and commit via `gitmap cpf "cg-execute - add map dispatch prompt 38 and fix heatmap rc integration"`.

---

## 2. Verification Protocol

| Verification Check | Target Command | Expected Outcome |
| :--- | :--- | :--- |
| **Heatmap RC Tables** | `gitmap rc heatmap tables` | Exit code 0, lists 8 tables with valid row counts |
| **Heatmap Status** | `gitmap heatmap status` | Exit code 0, displays logged action volume |
| **Agent Heatmap** | `gitmap agent heatmap` | Exit code 0, outputs write-claim rank table |
| **Relative Path Linter** | `python linter-scripts/check-relative-paths.py` | Exit code 0, 0 absolute paths |
| **Prompt Loaded Check** | `python linter-scripts/check-prompts-loaded.py` | Exit code 0, all prompts registered |
| **Fleet Sync Verification** | `python 03-ai-scripts/38-sync-prompts-skills-scripts.py` | All accessible repositories updated with clean ceremony |
