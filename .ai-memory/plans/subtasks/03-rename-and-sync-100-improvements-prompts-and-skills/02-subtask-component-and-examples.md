# Subtask 02: Component Specification, Schemas & Renaming Transformations

> [!IMPORTANT]
> **Subtask Identifier:** `.ai-memory/plans/subtasks/03-rename-and-sync-100-improvements-prompts-and-skills/02-subtask-component-and-examples.md`  
> **Parent Plan:** `.ai-memory/plans/pending/03-rename-and-sync-100-improvements-prompts-and-skills.md`  
> **Specification:** `02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills/02-component-spec.md`  
> **Status:** `[DONE]`  
> **Assigned Role:** Spec Author 02  

---

## 1. Objective & Scope

Author the exhaustive, grounded component specifications, file mapping matrices, YAML frontmatter schemas, IDE skill link invariants, and target repository purge blueprints for the **100 Improvements Prompts and Skills Renaming Suite** (`03-rename-and-sync-100-improvements-prompts-and-skills`).

This subtask delivers `02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills/02-component-spec.md` covering:
1. **Exhaustive Renaming Matrix:** Exact 1:1 mapping of all 4 prompts, 4 Antigravity skills, and 4 Cursor skills from `find-100-*` / `detail-plans-*` to canonical `next-100-*` and `next-detail-out-100-*`.
2. **Prompts Suite Internal Transformation:** Detailed internal content changes across `01-prompts/30-improvements/`:
   - Header titles and slug definitions (`## slug: next-100-ui-improvements`, etc.).
   - Invoke directives and slash commands (`/next-100-ui-improvements`, `/next-100-code-improvements`, etc.).
   - Phase 0 Antigravity and Cursor skill bootstrap directives.
   - Successive wave pending plan targets (`.ai-memory/plans/pending/XX-next-*`).
3. **Symmetric Skills Architecture:** Exact YAML frontmatter (`name`, `description`), header formatting, and IDE-detectable skill link syntax (`[skill](file;.agents/skills/<name>)` vs `[skill](file;.cursor/skills/<name>)`).
4. **Catalog & Documentation Synchronization:** Exact before/after replacements for `01-prompts/30-improvements/readme.md`, `01-prompts/readme.md`, and `.ai-memory/prompts.md`.
5. **Fleet Synchronization & Purge Logic:** Python implementation blueprint for `03-ai-scripts/38-sync-prompts-skills-scripts.py` to automatically purge deprecated prompts and skill directories across the 43 target repositories during sync.
6. **Binary Verification Contract:** Structured validation commands ensuring 100% path hygiene, zero stale leftovers, and zero absolute paths.

---

## 2. Context & Guardrail Invariants

- **Strict Relative Git Paths:** All markdown links, file paths, citations, and subtask paths must strictly use relative paths (`02-spec/...`, `.ai-memory/...`, `01-prompts/...`). TOTAL BAN on absolute filesystem paths (`C:\...`, `/home/...`) and `file:///` URIs.
- **Strict Lowercase Naming:** All created and modified files, directories, and paths must adhere strictly to lowercase naming conventions.
- **Zero Builds & Tests (Rule R1):** No builds, compilation, or full test suites may be executed.
- **No Git Commands:** Subagents must never run `git add`, `git commit`, `git push`, or `git status` to prevent working tree collisions.
- **Zero Placeholders:** Specifications and subtasks must contain fully resolved file paths, exact code samples, and deterministic criteria. Placeholders like `[N]`, `TODO`, or `...` are strictly prohibited.
- **Top-Instruction Priority Mandate:** Top-instruction precedence verified and active.

---

## 3. Action Checklist

- [x] Analyze master plan requirements in `.ai-memory/plans/pending/03-rename-and-sync-100-improvements-prompts-and-skills.md`.
- [x] Inspect existing prompt files in `01-prompts/30-improvements/` (`01-find-100-improvements-in-ui-level.md`, `02-find-100-improvements-in-code-base.md`, `03-find-100-improvements-in-whole-system.md`, `04-detail-plans-and-improvements-for-review.md`).
- [x] Inspect existing skill files in `.agents/skills/` and `.cursor/skills/`.
- [x] Inspect catalog indices in `01-prompts/30-improvements/readme.md` and `01-prompts/readme.md`.
- [x] Inspect fleet synchronization engine and purge mechanisms in `03-ai-scripts/38-sync-prompts-scripts-scripts.py`.
- [x] Document complete 4-part renaming matrix linking old paths to new paths, old slugs to new slugs, and slash commands.
- [x] Define internal content transformation rules for all 4 prompt files (Titles, Slugs, Invocations, Phase 0 Bootstrap, Continuity Plans).
- [x] Define symmetric YAML frontmatter and IDE link syntax rules for all 8 skill files (4 Antigravity + 4 Cursor).
- [x] Document catalog index replacements and companion skill listings for `01-prompts/30-improvements/readme.md`.
- [x] Specify Python purge logic in `03-ai-scripts/38-sync-prompts-skills-scripts.py` to remove legacy prompt files and skill directories in target fleet repositories.
- [x] Author and persist `02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills/02-component-spec.md`.
- [x] Author and persist `.ai-memory/plans/subtasks/03-rename-and-sync-100-improvements-prompts-and-skills/02-subtask-component-and-examples.md`.
- [x] Validate relative path hygiene (zero `file:///` URIs, zero absolute paths, strictly lowercase naming).

---

## 4. Execution Trace & Deliverables Summary

### 4.1 Primary Deliverable: `02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills/02-component-spec.md`
- **Section 1: Executive Summary & Renaming Mandate:** Captures user prompt verbatim, articulates systems rationale for the `next-*` nomenclature transition, and specifies non-negotiable boundaries.
- **Section 2: Exhaustive Renaming Matrix:** Complete mapping tables for prompts, Antigravity skills, Cursor skills, slugs, and slash commands, illustrated with a Mermaid architecture diagram.
- **Section 3: Prompts Suite Transformation Specifications:** Comprehensive section-by-section breakdown of required changes for:
  - `01-next-100-ui-improvements.md`
  - `02-next-100-code-improvements.md`
  - `03-next-100-whole-system-improvements.md`
  - `04-next-detail-out-100-improvements.md`
- **Section 4: Antigravity & Cursor Skills Architecture Transformation:** Detailed specifications for the 8 skills across `.agents/skills/` and `.cursor/skills/`, including exact YAML frontmatter, header lines, and platform-specific IDE link syntax (`[skill](file;.agents/skills/<name>)` vs `[skill](file;.cursor/skills/<name>)`).
- **Section 5: Catalog & Documentation Synchronization:** Full replacement blocks for `01-prompts/30-improvements/readme.md`, `01-prompts/readme.md`, and `.ai-memory/prompts.md`.
- **Section 6: Fleet Synchronization & Target Repository Purge Logic:** Concrete Python code blueprint for `03-ai-scripts/38-sync-prompts-skills-scripts.py`, defining pre-sync unlinking of old prompts and `shutil.rmtree` on old skill directories across all 43 target repositories.
- **Section 7: Concrete Examples & Transformation Diff Patterns:** Three concrete before/after unified diff blocks illustrating prompt header migration, skill YAML frontmatter update, and companion skill cross-linking.
- **Section 8: Binary Verification Criteria & Invariants:** 10 testable assertions with exact GitMap commands, target directories, and expected zero-error outcomes.

### 4.2 Secondary Deliverable: `.ai-memory/plans/subtasks/03-rename-and-sync-100-improvements-prompts-and-skills/02-subtask-component-and-examples.md`
- Persisted subtask tracking ledger with complete metadata, action checklist, execution trace, and verification criteria.

---

## 5. Binary Verification Criteria

| Check | Target | Command / Proof | Expected Outcome | Result |
|:---|:---|:---|:---|:---:|
| Spec 02 Exists | `02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills/02-component-spec.md` | `gitmap lf 02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills` | File present, >300 lines | PASS |
| Subtask 02 Exists | `.ai-memory/plans/subtasks/03-rename-and-sync-100-improvements-prompts-and-skills/02-subtask-component-and-examples.md` | `gitmap lf .ai-memory/plans/subtasks/03-rename-and-sync-100-improvements-prompts-and-skills` | File present, non-empty | PASS |
| Absolute Path Ban | Both files | `gitmap aum search "file:///" 02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills` | 0 hits | PASS |
| Absolute Path Ban | Both files | `gitmap aum search "file:///" .ai-memory/plans/subtasks/03-rename-and-sync-100-improvements-prompts-and-skills` | 0 hits | PASS |
| Windows Drive Ban | Both files | `gitmap aum search "C:\\" 02-spec/21-app/03-rename-and-sync-100-improvements-prompts-and-skills` | 0 hits (except descriptive bans) | PASS |
| Lowercase Naming | File paths | Inspect filenames | Strictly lowercase | PASS |
| Zero Placeholders | Both files | Inspect content | 0 instances of `[N]`, `TODO`, or `...` | PASS |
