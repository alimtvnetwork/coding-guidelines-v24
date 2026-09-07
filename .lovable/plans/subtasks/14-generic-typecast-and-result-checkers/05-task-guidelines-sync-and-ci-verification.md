# Subtask 05: Guidelines Documentation, Sync, and CI Verification

## Objective
Update the core coding guidelines with the new Checker interface conventions and optimized reflection patterns, sync mirrors, execute all CI quality gates, and document performance pointers for `ReflectSetTo`.

## Target Files
- `02-spec/17-consolidated-guidelines/34-compiled-simple-coding-guidelines.md`
- `.lovable/coding-guidelines.md` (mirrored via script)
- `.cursorrules` (mirrored via script)

## Detailed Instructions
1. In `02-spec/17-consolidated-guidelines/34-compiled-simple-coding-guidelines.md`:
   - Document the complete Checker interface hierarchy (`IsSuccessChecker`, `IsFailureChecker`, `IsInvalidChecker`, `IsNullChecker`, `IsEmptyChecker`, `IsDefinedChecker`).
   - Document the fast-path reflection casting architecture.
2. Run `node scripts/sync-guidelines.mjs` to synchronize the mirrors.
3. Run `python 03-ai-scripts/06-cicd-local-runner.py` to ensure all 22 quality gates exit with code 0.
4. Prepare technical performance analysis comparing legacy reflection vs fast-path type-switched reflection.
