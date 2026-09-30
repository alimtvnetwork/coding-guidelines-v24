# Plan 16: Minor Release v6.49.0, then Harden the Execute Folder Prompts

- **Status:** IN PROGRESS
- **Created:** 2026-10-01
- **Builds on:** [15-execute-parent-task-v4-antigravity.md](15-execute-parent-task-v4-antigravity.md) (the V3 audit, defects D1 to D14, and the V4 fixes reused here)
- **Folder in scope:** `01-prompts/14-execute/` and the `.agents/skills/` copies of its prompts
- **Git scope:** this plan is committed to `main` first; then the v6.49.0 release (release branch, tag, merge to `main`, GitHub release); then one commit on `main` with the prompt fixes.

---

## 1. User Request (Verbatim)

```text
Okay. So in similar fashion, can we improve other prompts? So I want you to go inside the CTX tube folder and find more prompts if we can enhance. But before that, I want you to make a minor bump and release, and then you make any changes inside the CT execute prompt. Okay? First, make a plan and also write the plan into the file system. And you do not ask for any questions to me. At the end, you write the summary. At the end, you write if there is any conclusion. Okay? Confirm your findings and improvements. Is it clear? So after you release, you come back to the main branch, make your changes, then finally make a change like this in the code, main language we are using. Is it clear?
```

### Interpretation (no questions allowed, so each assumption is logged here)

1. "CTX tube folder" and "CT execute prompt" are read as the execute folder, `01-prompts/14-execute/`, which is the folder V4 was written for.
2. "In similar fashion" is read as applying the V4 fixes (plan 15, D1 to D14) to the older prompts in that folder, not as writing another full rewrite.
3. "Make a change like this in the code, main language we are using" is read as: commit and push the prompt changes to `main` the same way as the V4 commit (explicit paths, one `Feature:` commit). No source code outside the prompts and their skill copies changes.
4. Order: plan first, then the minor release, then return to `main` and make the prompt changes.

## 2. Findings (Evidence)

A pattern count over `01-prompts/14-execute/` (run 2026-10-01) found the V4 defects in every older prompt:

| Prompt | `"Model"` field | `git add -A` | `go generate` commit | Autofixer on a file | `<SYSTEM_MESSAGE>` | "Override any planning" |
|---|---|---|---|---|---|---|
| 01-execute-pending-tasks | 0 | 3 | 1 | 0 | 0 | 0 |
| 02-execute-parent-task-with-n-steps | 2 | 3 | 0 | 1 | 3 | 1 |
| 03-execute-batched-loop | 0 | 2 | 1 | 0 | 0 | 0 |
| 04-execute-ai-instruction-writer | 0 | 0 | 1 | 0 | 0 | 0 |
| 05-execute-batched-loop-wor | 0 | 2 | 1 | 0 | 0 | 0 |
| 06-execute-parent-task-with-n-steps-v2 | 2 | 3 | 1 | 1 | 3 | 1 |
| 07-execute-batched-loop-v2 | 0 | 2 | 1 | 0 | 0 | 0 |
| 08-excute-parent-old | 0 | 3 | 1 | 0 | 0 | 1 |
| 09-parent-task-in-below-steps | 2 | 3 | 1 | 1 | 3 | 1 |
| 10 (V3) | 2 | 3 | 1 | 1 | 5 | 1 |
| 11 (V4) | 0 | 0 | 0 | 0 (uses `--check-only`) | 0 | 0 |

Two more facts:

- Antigravity runs the `.agents/skills/` copies, not the prompts. Six skills are full copies of execute prompts and carry the same sentences: `execute-pending-tasks`, `execute-batched-loop`, `execute-ai-instruction-writer`, `execute-batched-loop-wor`, `execute-parent-task`, and `parent-task-in-below-steps`. Fixing only the prompts would change nothing for Antigravity runs.
- The defects are the same sentences repeated word for word, so one fixed replacement per sentence template fixes every copy without rewording anything else.

## 3. Fix Templates

Each template replaces one exact sentence. Nothing else in the prompts changes.

1. **T1 Invalid payload field (D1).** Delete the `"Model": "inherit",` lines from the `invoke_subagent` payload examples. `invoke_subagent` entries take only `TypeName`, `Role`, `Prompt`, and an optional `Workspace`.
2. **T2 Whole-tree staging (D4).** Replace every `git add -A` instruction with staging by explicit path (`git add -- <paths this task changed>`). Where `gitmap cpf`, `cpb`, or `cpr` is offered, add that they stage every file, so they are safe only when `git status --porcelain` was clean before the task.
3. **T3 Autofixer check that can pass without checking (D3).** Replace `python 03-ai-scripts/05-guideline-autofixer.py <file>` with the folder form plus `--check-only --ext <.ext>`, and state that 0 files scanned is a FAIL.
4. **T4 Undocumented wakeup tag (D6).** Replace the `<SYSTEM_MESSAGE>` wording with "subagent results arrive as messages", which is what the Antigravity docs describe.
5. **T5 Instruction the model cannot follow (D6).** Replace "Override any planning mode stop directives." with a note that the only allowed pause is the one the Artifact Review Policy (a user setting) imposes.
6. **T6 Generated code (D5).** Replace "run `go generate ./...` and commit the generated files" with: commit regenerated files only if `git ls-files` shows the repo already tracks them (so CI sees no drift); otherwise never commit them (Hard Rule 1). V4's R15 gets the same wording, so all execute prompts agree.
7. **T7 Build and test contradiction (D2).** In prompts that ban builds and tests (06 and 09), the closing block keeps the user's words except "run builds and full unit tests", which becomes "run the targeted checks (builds and full unit tests stay in CI unless the owner asks for them)". This is the same change V4 made.
8. **T8 Pointer to V4.** Prompts 02 and 06 (V1 and V2 of the same workflow) get a one-line note under the title that V4 supersedes them for Antigravity runs.

## 4. Out of Scope (and Why)

- **V3 (`10-execute-parent-task-with-n-steps-v3.md`) and its three skill copies** stay unchanged. They are the payload of `03-ai-scripts/42-sync-v3-prompts-and-skills.py` for the sibling repos, and V4 is their fix.
- **Skills outside the execute folder** that share some of these sentences (`cg-*`, `ci-cd-fix`, `ci-cd-fix-with-release`, `release-*`, `coding-guidelines`, `temp-end-to-end-tests`, `execute-coding-guideline-fix`) are a follow-up.
- **Deeper rewrites** (the zero-question mandate, repeated rules, step accounting) are what V4 already does; the older prompts only get the mechanical fixes.
- **Renaming `08-excute-parent-old.md`** (typo) would break the prompt index and pointer skills; it is a follow-up.
- **`03-ai-scripts/29-release-orchestrator.py`** stages with `git add -A`. It is safe here because the tree is clean when the release starts; changing the script is a follow-up.

## 5. Release Approach (v6.48.0 to v6.49.0, minor)

1. Commit this plan to `main` and push, so the release starts from a clean, synced tree.
2. Run `python 03-ai-scripts/29-release-orchestrator.py --tier minor --scope "<real changes since v6.48.0>" --no-push`. It runs the pre-release gate (`06-cicd-local-runner.py --run-tests`), creates `release/v6.49.0`, bumps the version through `03-ai-scripts/37-bump-version.py` (which runs `npm run sync`), writes `.ai-memory/release/release-notes-vX.Y.Z.md` for the new version, commits, tags `v6.49.0`, merges into `main`, and returns to `main`.
3. Inspect the release commit, then push `main`, `release/v6.49.0`, and the tag, and publish the GitHub release with `gh release create v6.49.0 --title v6.49.0 --notes-file .ai-memory/release/release-notes-v6.49.0.md --generate-notes`.
4. Any failed or flagged step is logged under `.ai-memory/release/issues/`.

## 6. Verification (Prompt Changes)

- `python linter-scripts/check-prompts-loaded.py`, `check-prompt-and-spec-paths.py`, `check-sequence-integrity.py`, and `check-forbidden-strings.py` all exit 0.
- A pattern search over prompts 01 to 09 and the six skill copies finds no `"Model"`, no `git add -A`, no autofixer call without `--check-only`, no `<SYSTEM_MESSAGE>`, and no "Override any planning".
- `git diff --stat` shows only the files listed in the subtask ledger.

## 7. Subtask Ledger

### S1: Write and commit this plan

- **Owned files:** `.ai-memory/plans/pending/16-execute-folder-hardening-and-v6-49-release.md`, `.ai-memory/plans/readme.md`
- **Status:** DONE. Evidence: `check-sequence-integrity.py` and `check-forbidden-strings.py` exit 0; committed to `main` in a `Docs:` commit before the release.

### S2: Minor release v6.49.0

- **Owned files:** the files `37-bump-version.py` and `npm run sync` change, plus the release notes file (`.ai-memory/release/release-notes-vX.Y.Z.md`)
- **Acceptance:** tag `v6.49.0` and branch `release/v6.49.0` exist on origin, `main` contains the release commit, and the GitHub release has the install one-liners.
- **Status:** PENDING

### S3: Fix prompts 01 to 09 (T1 to T8)

- **Owned files:** `01-prompts/14-execute/01-execute-pending-tasks.md` through `01-prompts/14-execute/09-parent-task-in-below-steps.md`
- **Status:** PENDING

### S4: Fix the six skill copies (same templates)

- **Owned files:** `skill.md` in `.agents/skills/execute-pending-tasks/`, `execute-batched-loop/`, `execute-ai-instruction-writer/`, `execute-batched-loop-wor/`, `execute-parent-task/`, and `parent-task-in-below-steps/`
- **Status:** PENDING

### S5: Align V4 R15 with T6

- **Owned files:** `01-prompts/14-execute/11-execute-parent-task-with-n-steps-v4.md`
- **Status:** PENDING

### S6: Indexes, verification, commit, push

- **Owned files:** `.ai-memory/what-to-read.md` (changelog line), this plan (statuses)
- **Acceptance:** section 6 passes; one `Feature:` commit on `main`, pushed.
- **Status:** PENDING

## 8. Follow-ups

1. Apply T2 to T6 to the non-execute skills listed in section 4.
2. Change `03-ai-scripts/29-release-orchestrator.py` to stage the release files by path.
3. Rename `08-excute-parent-old.md` and update its index entries and pointer skills.
4. Everything still open in plan 15 section 10 (live Antigravity run, alias skills, sibling-repo rollout).
