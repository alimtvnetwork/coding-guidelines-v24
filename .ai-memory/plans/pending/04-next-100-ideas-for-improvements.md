# Master Plan: 04-next-100-ideas-for-improvements

## 1. Task Information
- **Task Title:** Next 100 Ideas for Improvements Prompt, Skills & Fleet Synchronization
- **Task Slug:** `04-next-100-ideas-for-improvements`
- **Status:** `ACTIVE`
- **Preamble Priority Mandate:** Top-Instruction Priority verified and active.
- **Rules Invariants:**
  - Strictly relative Git paths only (`02-spec/...`, `.ai-memory/...`, `01-prompts/...`).
  - Strict lowercase filenames across all files and folders.
  - Symmetrical skills across `.agents/skills/` and `.cursor/skills/`.
  - Zero raw grep/shell searches; GitMap AUM commands exclusively.
  - Workers never execute git commands; index lock collision prevention.
  - Zero builds and test suites ban (Rule R1).
  - Single atomic commit via `gitmap cpf` using hyphen separator without colons.

---

## 2. User Request (Verbatim)

```text
also create next 100 ideas for improvements prompt, what should generate next 100 great ideas for improves for the project where the prompt is running following other ones and also please sync with all repo with skills, clear?
```

---

## 3. Scope & Deliverables

1. **New Canonical Prompt 05:**
   - Path: `01-prompts/30-improvements/05-next-100-ideas-for-improvements.md`
   - Purpose: Autonomous ideation flywheel that discovers, structures, and catalogs the next 100 high-leverage improvement and innovation ideas for any target codebase.
   - Follows the V6 pattern with $N = 400$, $A = 3$, $H = 2$, Capacity = 6, 10 Ideation Clusters (10 ideas each), and 8-field Idea Specification Schema.
   - Enqueues discovered ideas into `.ai-memory/plans/pending/XX-next-100-ideas-for-improvements.md` with subtasks in `.ai-memory/plans/subtasks/XX-next-100-ideas-for-improvements/`.
2. **Companion Skills (Antigravity & Cursor):**
   - `.agents/skills/next-100-ideas-for-improvements/skill.md`
   - `.cursor/skills/next-100-ideas-for-improvements/skill.md`
3. **Master Catalog & Registry Updates:**
   - `01-prompts/30-improvements/readme.md` (Update catalog table to 5 prompts)
   - `01-prompts/readme.md` (Update Category 30 row)
   - `.ai-memory/prompts.md` (Register prompt, slug, and `/next-100-ideas-for-improvements` slash command)
   - `.ai-memory/plans/readme.md` & `02-spec/21-app/readme.md` (Register task 04)
4. **Fleet Synchronization:**
   - Propagate new prompt and skills across all 43 fleet repositories via `03-ai-scripts/38-sync-prompts-skills-scripts.py`.

---

## 4. Subtask Allocation

| Task ID | Title | Assigned Role | Status | Owned Files |
|:---|:---|:---|:---:|:---|
| `Task-01` | Specification Authoring & Subtasks | Spec Authors 01, 02, 03 | `COMPLETED` | `02-spec/21-app/04-next-100-ideas-for-improvements/*`, `.ai-memory/plans/subtasks/04-next-100-ideas-for-improvements/*` |
| `Task-02` | Author Prompt 05 & Update Category Catalog | Worker 01 | `COMPLETED` | `01-prompts/30-improvements/05-next-100-ideas-for-improvements.md`, `01-prompts/30-improvements/readme.md` |
| `Task-03` | Author Symmetric Companion Skills | Worker 02 | `COMPLETED` | `.agents/skills/next-100-ideas-for-improvements/*`, `.cursor/skills/next-100-ideas-for-improvements/*` |
| `Task-04` | Update Catalogs, Registries & Sync Script | Worker 03 | `COMPLETED` | `01-prompts/readme.md`, `.ai-memory/prompts.md`, `.ai-memory/plans/readme.md`, `02-spec/21-app/readme.md` |
| `Task-05` | Fleet Sync & Verification | Lead Orchestrator | `COMPLETED` | `03-ai-scripts/38-sync-prompts-skills-scripts.py`, all 43 repositories |
| `Task-06` | Atomic Commit & Push | Lead Orchestrator | `COMPLETED` | All modified and created files |
