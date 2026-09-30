# Canonical Specification: White Blue Theme Design System, Parent Task N-Steps V3 Prompt & Multi-Repository Synchronization

> **/goal** Define the end-to-end architecture for the `04-white-blue-theme` design system specification, the `10-execute-parent-task-with-n-steps-v3.md` prompt (`N = 300` top header with mandatory subagent spawning), synchronized skills, and the 4-stage multi-repository backup/release/sync/release pipeline across 43 repositories.
> **/learn** Master the 3-format color specification requirement (`HEX`, `RGB`/`RGBA`, `HSL` + `OKLCH`), strict zero-source-brand naming rules, V3 header and `invoke_subagent` enforcement, and atomic multi-repo release orchestration.

**Version:** 1.0.0
**Updated:** 2026-09-30
**Status:** Active
**AI Confidence:** Production-Ready
**Ambiguity:** None

---

## 1. Architectural Overview & Unified Blueprint

This specification establishes three tightly integrated deliverables across the `coding-guidelines` meta-repository and its 43 connected target repositories:

1. **White Blue Theme Design System (`02-spec/07-design-system/04-white-blue-theme/`):**
   - A comprehensive, blind-AI-accessible design system specification capturing a light-first, high-trust enterprise editorial UI ("White Blue Theme" / "White Theme").
   - **Strict Naming Rule:** Zero references to the external reference software name anywhere in `02-spec/`, `01-prompts/`, or skills.
   - **Mandatory 3-Format Color Standard:** Every color token in the palette and semantic tables MUST be explicitly documented in three universal formats: **HEX** (`#RRGGBB`), **RGB / RGBA** (`rgb(r, g, b)` / `rgba(r, g, b, a)`), and **HSL** (`hsl(h, s%, l%)`), alongside the CSS `oklch(...)` runtime value.
   - **Complete Component, Typography, & Motion Blueprint:** Documents the `Ubuntu` + `Poppins` + `JetBrains Mono` typography stack, alternating light/soft section band rhythm (`.band-soft`, `.band-dark`, `.band-void`), hardware-accelerated CSS3 and spring animations (`shine-sweep`, `pointer-fill`, `card-premium`, `row-premium`, `slide-swap`, `spotlight-light`, 3D push-pins, fluted glass scroll-stack), and complete React/Tailwind/LESS code examples.

2. **`Parent Task with N Steps V3` Prompt (`01-prompts/14-execute/10-execute-parent-task-with-n-steps-v3.md`) & Skills:**
   - Combines the V3 header/title and slash-command flavor from `06-old-prompts/v3/14-execute/` with the unified master pipeline of `01-prompts/14-execute/06-execute-parent-task-with-n-steps-v2.md`.
   - Places **`N = 300`** (alongside `A = 2`, `H = 2`) prominently in the top header block so the user can edit `N` immediately at the top of the prompt.
   - Enforces **Mandatory Subagent Spawning (`invoke_subagent`)** as an absolute, non-bypassable requirement in both Phase 1 (parallel discovery and modular spec authoring) and Phase 2 (parallel code execution using `TypeName: "self"`), explicitly banning solo execution when subagent capacity is configured.
   - Synchronizes the corresponding Antigravity skills (`.agents/skills/`) and Cursor skills (`.cursor/skills/`), and indexes the prompt in `.ai-memory/prompts.md` and `01-prompts/readme.md`.

3. **Multi-Repository 4-Stage Pull, Pre-Backup/Release, Guideline/Prompt/Skill Sync & Post-Release Pipeline:**
   - Across the 43 target repositories shown in the user's reference screenshot (`![Target Repositories List](../../assets/screenshots/white-blue-theme-and-v3-sync-01.png)`):
     1. **Stage 1 — Pull:** Checkout the base branch (`main` / `master` / `develop`) and run `git pull origin <base_branch> --no-rebase`.
     2. **Stage 2 — Pre-Sync Backup & Release:** Create and push a timestamped backup branch (`backup/pre-v3-nsteps-sync-<timestamp>`) and ensure a pre-sync release branch (`release/vX.Y.Z`) and annotated tag (`vX.Y.Z`) exist and are pushed to `origin`.
     3. **Stage 3 — Synchronize Guidelines, Prompts, Skills & Scripts:** Mirror `01-prompts/`, `.agents/skills/`, `.cursor/skills/`, `03-ai-scripts/`, `.agents/scripts/`, and coding guideline specs (`02-spec/02-coding-guidelines/`, `02-spec/07-design-system/`, `02-spec/17-consolidated-guidelines/`, `.ai-memory/coding-guidelines.md`, `.ai-memory/prompts.md` where applicable).
     4. **Stage 4 — Post-Sync Release:** Commit and push on the base branch, bump the SemVer patch version, create and push the new release branch (`release/v<next>`) and tag (`v<next>`), and merge back to the base branch.

---

## 2. User Request (Verbatim)

![Target Repositories Screenshot](../../assets/screenshots/white-blue-theme-and-v3-sync-01.png)

```text
reference-white-blue-ui

Okay. So I want you to read the reference White Blue UI from this work directory. And mostly what I want you to do is understand its theming, component, coloring, everything. So create a white theme inside your design spec folder. That means inside the spec folder, there is this design system folder, okay? Where you can create a theme called 04 white theme [source-brand], and then do not mention [source-brand], just try to mention white theme. Okay? White blue theme. And in this white blue theme, I want you to add everything for an AI to understand and follow through the aspects of the, let's say, components, coloring, color theme. So understand how these other themes are written, how the other design systems are mentioned. Okay? So based on that, I want you to create your own component base, color base, and every time the color needs to be in three formats. If there is animation, I want you to mention this animation. So take some code parts as well, some examples, so that any AI can understand it very well. Okay. I really like the cleanness of that software, how the website is displayed. Cleanness, color, fonts, everything, the component-wise. It's really amazing. I want you to follow everything so that an AI who is actually blind, it can follow through and understand and fix it. Also, at the same time, I want you to create a prompt. There is a prompt, right? So parent task with N steps V2. I want you to create V3. So V3 would be similar to the previous one. If you know, we have V3 folder, right? So that has a title and how the prompts were there. So we will have some of the flavor from this, and the N will be on top. N will be on the header where we can change it. Usually from now on, the N will be 300 steps, and the agent spawning would be must. Currently, the agents are not spawning, sub agents. That is a big mistake I think you have done. So I can fix that and just create a new prompt. Parent task with N steps V3. Okay. Keep it in the same folder. Try to go deep and then create that prompt, and then from there, try to create and sync the skills. Okay? I hope you understand. And also after that, I want you to sync with other code bases so that the coding guideline is sync, coding guideline and skills. Okay? And before you sync that, try to do a pull on those repositories that I'm mentioning. Okay? So in those repositories, we're going to take a backup first. Okay? Backup and release first, and then you are going to put the new prompts, new skills, and then make another release. Do you understand? Yes.
```

---

## 3. Extracted Actionable Task List

| Task ID | Deliverable | Target Scope |
|:---|:---|:---|
| `Task-01` | Ingest uploaded screenshot (`assets/screenshots/white-blue-theme-and-v3-sync-01.png`) and reverse-engineer the reference UI tokens, typography, animations, and components | `assets/screenshots/white-blue-theme-and-v3-sync-01.png`, `02-spec/07-design-system/` |
| `Task-02` | Author `02-spec/07-design-system/04-white-blue-theme/` (`readme.md` + `01`..`04` spec files) with 3-format colors (`HEX`, `RGB`/`RGBA`, `HSL`), fonts, animations, code examples, and zero mention of the source brand name; update `02-spec/07-design-system/readme.md`, `99-consistency-report.md`, and `02-spec/17-consolidated-guidelines/10-design-system.md` | `02-spec/07-design-system/04-white-blue-theme/`, `02-spec/07-design-system/readme.md`, `02-spec/17-consolidated-guidelines/10-design-system.md` |
| `Task-03` | Author `01-prompts/14-execute/10-execute-parent-task-with-n-steps-v3.md` with V3 title/header flavor, `N = 300` prominently at the top header, and mandatory `invoke_subagent` (`A = 2, H = 2`) enforcement | `01-prompts/14-execute/10-execute-parent-task-with-n-steps-v3.md`, `01-prompts/readme.md`, `.ai-memory/prompts.md` |
| `Task-04` | Create and synchronize Antigravity and Cursor skills for V3 (`execute-parent-task-with-n-steps-v3`, `execute-parent-task-with-n-steps`, `parent-task-n-step-loop`) | `.agents/skills/`, `.cursor/skills/` |
| `Task-05` | Upgrade `03-ai-scripts/38-sync-prompts-skills-scripts.py` and execute the 4-stage `pull -> backup & pre-release -> sync guidelines/prompts/skills -> post-release` workflow across all 43 target repositories | `03-ai-scripts/38-sync-prompts-skills-scripts.py`, 43 target repositories |
| `Task-06` | Consolidate execution plan into `.ai-memory/plans/completed/`, verify all linters (`check-prompts-loaded.py`, `check-relative-paths.py`, `21-sequence-integrity-linter.py`), and release/push `coding-guidelines-v24` | `.ai-memory/plans/completed/14-white-blue-theme-v3-prompt-and-multi-repo-sync.md` |

---

## 4. Verification & Acceptance Criteria

### AC-APP-014: White Blue Theme, V3 N-Steps Prompt & Multi-Repo Sync Verification

**Given** The `04-white-blue-theme` specification suite, `10-execute-parent-task-with-n-steps-v3.md` prompt, skills, and multi-repo sync script are authored.
**When** Running repository linters and multi-repo sync verification.
**Then** All color tables provide `HEX`, `RGB`/`RGBA`, and `HSL` formats; zero references to the source project name exist in `02-spec/07-design-system/04-white-blue-theme/`; `10-execute-parent-task-with-n-steps-v3.md` defines `N = 300` in the top header and enforces `invoke_subagent`; and all 43 target repositories are pulled, backed up, pre-released, synced, and post-released cleanly.

**Verification command:**

```bash
python linter-scripts/check-prompts-loaded.py && python linter-scripts/check-relative-paths.py && python 03-ai-scripts/21-sequence-integrity-linter.py 02-spec/07-design-system
```

**Expected:** exit 0.
