# Subtask-04 — SQL/Enum Prompt Instruction + Guideline/Error Prompt Fine-Tune

- parent task: `execute-n-steps-letterly-prompt-formatter`
- parent spec: `02-spec/21-app/03-execute-n-steps-letterly-prompt-formatter/04-sql-repository-pattern-and-enum-actions-spec.md`
- parent plan: `.ai-memory/plans/pending/03-execute-n-steps-letterly-prompt-formatter.md` (section 4, 2026-10-23 Dispatch Addendum)
- owner: Coordinator-01 (prompts track)
- status: DONE (2026-10-23) — Task-01 + Task-02 complete, committed, pushed. Coordinator-02's fleet sync (Subtask-03) is ungated.
- date: 2026-10-23

## Box (disjoint)

- Mine: `01-prompts/**` for this dispatch.
- NOT mine: `.agents/skills/**`, `.cursor/skills/**` (Coordinator-02, Subtask-05); lead's spec/plan/index files.

## Deliverables

### Task-01 — A–D instruction block in Letterly execute prompts

- Add a mandatory instruction block carrying spec rules A–D (SQL repository pattern in every language; every action from an enum — switchable, exhaustively checked; enum definition/usage/writing + util-methods checkup; error-management tie-ins) to `01-prompts/22-letterly/03-execute-n-steps-letterly.md`, placed so GENERATED prompts include it as a non-negotiable.
- Mirror the same block into `01-prompts/23-cursor-prompts/03-execute-n-steps-letterly-cursor.md`, respecting the mirror's own conventions (`.cursor/skills` link, slug line, /plan motto wording).
- Each file's existing format and voice matched exactly.

### Task-02 — Fine-tune coding-guideline + error-management prompts (no confirmation round)

- `01-prompts/05-coding-guidelines/01-plan-coding-guideline-audit.md`
- `01-prompts/05-coding-guidelines/02-execute-coding-guideline-fix.md`
- `01-prompts/05-coding-guidelines/03-cg-audit-gap-n-steps.md`
- `01-prompts/15-cg-execute/02-error-management.md` (error-management prompt, located via `gitmap find`)
- `01-prompts/15-cg-execute/19-result-wrapper-and-apperror-returns.md` (error-management companion)
- Only lines the A–D rules actually change; minimal diffs.

## Preflight evidence

- Pull-first: `gitmap pull` → coding-guidelines-v24 up-to-date (ba475f7) [main] (clean), exit 0.
- Heatmap collision control (2026-10-23):
  - `gitmap rc heatmap status` → Total Actions: 0, Uncommitted Actions: 0.
  - `gitmap rc heatmap tree` → empty (no claims).
  - `gitmap rc heatmap live` → unrecognized subcommand (usage lists only sync/status/commit under rc heatmap).
  - `gitmap rc heatmap sync` → synchronized RC heatmap telemetry with repo-cache, exit 0.
  - Fallback `gitmap heatmap discuss` → unknown command (E1000); `gitmap agent heatmap` → "No file claims recorded."
  - Conclusion: zero in-flight claims on any target file; safe to write.

## Worker partition (V6, A=2 per stage)

| Stage | Worker | Box | Mode |
|---|---|---|---|
| Research | R1 | `01-prompts/22-letterly/`, `01-prompts/23-cursor-prompts/` (03 files + siblings) | read-only |
| Research | R2 | `01-prompts/05-coding-guidelines/`, `01-prompts/15-cg-execute/02-error-management.md`, `01-prompts/15-cg-execute/19-result-wrapper-and-apperror-returns.md` | read-only |
| Execution | E1 (dispatched) | the two 03 Letterly files | edit |
| Execution | E2a (dispatched) | `01-prompts/05-coding-guidelines/` (3 files) | edit |
| Execution | E2b (dispatched) | `01-prompts/15-cg-execute/02-error-management.md`, `01-prompts/15-cg-execute/19-result-wrapper-and-apperror-returns.md` | edit |

## Progress log

- 2026-10-23: spec + plan addendum + gold-standard files read; repo pulled clean; heatmap checks recorded (above); R1 + R2 dispatched.
- 2026-10-23: R1 DONE. Findings: fixed content in this prompt family always lives as numbered items in the Output Format template's `# Actionable Items Must Follow Non-Negotiable` list, declared in formatter instruction 5; `## Additional Instructions` never carries non-negotiables. Placement: A–D become new fixed items 5–8 after existing items 1–4 (placeholder renumbered 5 → 9); instruction 5 updated to "fixed items 1-8 verbatim … items 9..N". Block drafted covering A (item 5), B (item 6), C (item 7), D (item 8); byte-identical in both files (no paths/links in block); each file's own divergences (`.agents` vs `.cursor` links, slug-line note, /plan motto) preserved. E1 dispatched with the exact edit text.
- 2026-10-23: E1 DONE. Both 03 files edited exactly as specified (44 → 48 lines each); E1 verified programmatically that items 5–8 and instruction 5 are identical across both files and that each file's divergences survived. Lead re-verified via `git diff`: instruction 5 updated; items 5–8 inserted verbatim; placeholder renumbered 9; nothing else changed.
- 2026-10-23: Linters on the change set: `gitmap py linter-scripts/check-prompts-loaded.py` → exit 0 (221 prompts on disk, index in sync); `gitmap py linter-scripts/check-forbidden-strings.py` → exit 0 (PASS, no forbidden strings).
- Commit mechanics deviation (disclosed): `gitmap cpf` stages ALL working-tree changes; Coordinator-02's in-flight `.agents/skills/cg-enum-standards/skill.md` edit (their box) is currently dirty in the shared tree, so a cpf commit would sweep it into my Task-01 commit and break both atomicity and the disjoint-box rule. Task-01 is therefore committed with targeted `git add` (relative paths only) + `git commit` using the same hyphen-separated, colon-free message convention + `git push`. Same approach will be used for Task-02 if the tree is still shared-dirty at that point.
- 2026-10-23: Task-01 committed + pushed: `cfe9333b` "prompts - add sql repository pattern and enum actions instruction to letterly execute prompts" (3 files: the two 03 prompts + this ledger; 77 insertions, 4 deletions). Coordinator-02's dirty skill files were left unstaged.
- 2026-10-23: R2 DONE. Audit of the 5 Task-02 files vs A–D: files 1–3 had no repository-pattern/dispatch/enum-util rules (gaps + weak lines identified with exact replacements); file 4 (02-error-management) had D-partial content only; file 5 (19-result-wrapper) already models the repository pattern correctly in its examples but its mandatory wrap examples dropped the repository name (D-inconsistent) — 6 targeted fixes proposed. All SQL/dispatch hits in files 4–5 outside those spots are agent-tooling boilerplate (heatmap/SQLite task DB, invoke_subagent) — confirmed false positives, untouched. E2a (files 1–3) + E2b (files 4–5) dispatched with R2's exact edit text; anchors that miss or multi-match must be skipped and reported, never improvised.
- 2026-10-23: E2b DONE — all 11 edits applied in files 4–5, none skipped (02-error-management 750 → 753 lines; 19-result-wrapper 1084 → 1088 lines). E2a DONE — files 1–3: 93 → 98, 242 → 246, 282 → 287 lines; two edits initially skipped on anchor mismatch (the file's actual text uses `` `*Type` `` where R2 quoted `` `Type` ``: §4 Step 2 item 6 and the severity Major row). Coordinator verified the exact lines via `gitmap aum search` (lines 159 and 107) and applied both with corrected anchors, preserving the file's `` `*Type*` `` style. Lead verified the full Task-02 diff: 5 files, 42 insertions, 21 deletions — only A–D-driven lines changed.
- 2026-10-23: Final linters: `gitmap py linter-scripts/check-prompts-loaded.py` → exit 0 (221 prompts, index in sync); `gitmap py linter-scripts/check-forbidden-strings.py` → exit 0 (PASS). Task-02 committed + pushed with this ledger update (targeted `git add`, same disclosed deviation as Task-01 — the shared tree still carries Coordinator-02's in-flight skill files).

## Follow-ups (for lead)

- Enum-dedicated prompts under `01-prompts/15-cg-execute/` not covered by this subtask's named scope (R2 inventory): `06-constants-and-enums.md`, `16-multi-language-enums-and-traits.md`, `28-python-array-constants-and-dynamic-enums.md`, `27-python-constants-and-magic-number-elimination.md`, plus `38-map-dispatch-and-handler-lookup.md` (B) and `07-data-and-schema.md` (A) — candidates for a follow-up A–D alignment pass.
- `01-prompts/05-coding-guidelines/03-cg-audit-gap-n-steps.md` Tier 3 boolean-prefix line is self-contradictory (R2 noted); not an A–D issue, left untouched.

## Evidence (filled at completion)

- Task-01 diffstat: `01-prompts/22-letterly/03-execute-n-steps-letterly.md` +8/−2 lines net (file 44 → 48 lines); `01-prompts/23-cursor-prompts/03-execute-n-steps-letterly-cursor.md` +8/−2 (44 → 48). Combined `git diff --stat` for the pair: 2 files, 12 insertions(+), 4 deletions(-) plus instruction-5 line rewrite counted within (verified via git diff output 2026-10-23).
- Task-02 diffstat: 5 files changed, 42 insertions(+), 21 deletions(-) — `01-prompts/05-coding-guidelines/01-plan-coding-guideline-audit.md` (93 → 98 lines), `01-prompts/05-coding-guidelines/02-execute-coding-guideline-fix.md` (242 → 246), `01-prompts/05-coding-guidelines/03-cg-audit-gap-n-steps.md` (282 → 287), `01-prompts/15-cg-execute/02-error-management.md` (750 → 753), `01-prompts/15-cg-execute/19-result-wrapper-and-apperror-returns.md` (1084 → 1088). All 29 planned edits applied (27 by E2a/E2b, 2 anchor-corrected by coordinator); zero skipped at completion.
- Linters: check-prompts-loaded.py exit 0; check-forbidden-strings.py exit 0 (verified after Task-01 edits and again after Task-02 edits, final state).
- Commits: Task-01 `cfe9333b` "prompts - add sql repository pattern and enum actions instruction to letterly execute prompts" (pushed). Task-02 commit SHA recorded in the coordinator's final report (this ledger update rides in it).
