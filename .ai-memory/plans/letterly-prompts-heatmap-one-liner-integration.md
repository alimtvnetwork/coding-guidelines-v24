# Master Execution Plan: Letterly Prompts Heatmap One-Liner Integration

> **Plan Identifier:** `.ai-memory/plans/letterly-prompts-heatmap-one-liner-integration.md`  
> **Status:** COMPLETED  
> **Specification:** `02-spec/21-app/letterly-prompts-heatmap-one-liner-integration/01-architecture-spec.md`  
> **Total Steps Budget:** N = 400 (Phase 1: 100, Phase 2: 150, Phase 3: 150)  
> **Concurrency Capacity:** A = 3, H = 2, Capacity = 6 (Dynamic Immediate Respawn, 100% Duty Cycle)  
> **Local Completion Boundary:** STRICT LOCAL ONLY (Total ban on fleet multi-repo sync)

---

## User Request (Verbatim)

```text
also in the letterly prompts include the heatmap concept as one liner please
```

---

## Subtask Decomposition

| Task ID | Title | Owned Subtask File | Target Files | Owner Role | Status |
|---|---|---|---|---|---|
| `Task-01` | Letterly Prompts Suite Heatmap One-Liner Update | `.ai-memory/plans/subtasks/letterly-prompts-heatmap-one-liner-integration/01-letterly-prompts-heatmap-update.md` | `01-prompts/22-letterly/*.md` (11 prompts + readme.md) | `LetterlyPromptWorker` | DONE |
| `Task-02` | Cursor Prompts Suite Heatmap One-Liner Mirror | `.ai-memory/plans/subtasks/letterly-prompts-heatmap-one-liner-integration/02-cursor-prompts-heatmap-mirror.md` | `01-prompts/23-cursor-prompts/*.md` (11 prompts + readme.md) | `CursorPromptWorker` | DONE |
| `Task-03` | Quality Verification, Linters & Local Atomic Commit | `.ai-memory/plans/subtasks/letterly-prompts-heatmap-one-liner-integration/03-verification-and-commit.md` | Workspace Root, GitMap Commit Engine | `LeadOrchestrator` | DONE |

---

## Subtask Execution Details

### Task-01: Letterly Prompts Suite Heatmap One-Liner Update
- **Owned Subtask File:** `.ai-memory/plans/subtasks/letterly-prompts-heatmap-one-liner-integration/01-letterly-prompts-heatmap-update.md`
- **Scope:** Update all 11 prompts in `01-prompts/22-letterly/` and `readme.md`:
  - Desktop/standard templates: add Heatmap Mandate in instructions and one-liner in `## Additional Instructions`.
  - Mobile templates (`01`, `07`): add heatmap directive into the single-line continuous paragraph.
  - Maintain `.agents/skills/` link syntax.

### Task-02: Cursor Prompts Suite Heatmap One-Liner Mirror
- **Owned Subtask File:** `.ai-memory/plans/subtasks/letterly-prompts-heatmap-one-liner-integration/02-cursor-prompts-heatmap-mirror.md`
- **Scope:** Mirror all 11 prompts to `01-prompts/23-cursor-prompts/` and `readme.md`:
  - Retain `.cursor/skills/` link syntax.
  - Ensure 100% template alignment with `01-prompts/22-letterly/`.

### Task-03: Quality Verification, Linters & Local Atomic Commit
- **Owned Subtask File:** `.ai-memory/plans/subtasks/letterly-prompts-heatmap-one-liner-integration/03-verification-and-commit.md`
- **Scope:** Run validation linters (`check-relative-paths.py`, `check-forbidden-strings.py`, `check-prompts-loaded.py`, `test_check_letterly_format.py`), stage changes, and commit via `gitmap cpf "letterly - include heatmap telemetry one liner across all prompts and cursor mirrors"`.
- **Key Invariant:** Strict local completion only — zero fleet sync.
