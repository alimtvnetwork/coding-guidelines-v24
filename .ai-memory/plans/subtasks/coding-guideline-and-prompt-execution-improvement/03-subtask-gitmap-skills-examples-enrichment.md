# Subtask Plan: GitMap Skills Examples Enrichment & Companion Skills Alignment

## Subtask ID: `03-subtask`
- **Parent Plan:** `.ai-memory/plans/coding-guideline-and-prompt-execution-improvement.md`
- **Spec Reference:** `02-spec/21-app/24-coding-guideline-and-prompt-execution-improvement/03-migration-and-verification-spec.md`
- **Primary Owner:** Worker 03 (Phase 2 Execution Wave)
- **Status:** `COMPLETED`

---

## 1. Intent & Technical Deliverables

This subtask executes the enrichment of GitMap skill definitions across both Antigravity (`.agents/skills/`) and Cursor (`.cursor/skills/`) ecosystems:
1. **Embed Comprehensive 12-Category Example Catalog:**
   - Add high-visibility, copy-paste real-world examples to `.agents/skills/gitmap/skill.md` covering all 12 operational categories (Authentication, Workspace Status, Runners, Storage/Vaults, AUM File Discovery, CI/CD Healing, Semantic Commits, Cluster/SSH Operations, Fleet Sync & Agent Tasks, Heatmap Telemetry, SQLite DB Suite, and Screenshot Extraction).
2. **Audit & Realign Companion Skills:**
   - Remove all prohibited `--help` instructions (`gitmap cluster --help`, `gitmap sc --help`, `gitmap ssh --help`) and replace with harmless, active management commands (`gitmap cluster status`, `gitmap sc list`, `gitmap ssh health --all`).
   - Correct all concurrency flags, excising hallucinated `--max-concurrency` and standardizing on `--concurrency` (cloner) and `--workers` (sync/search).
   - Correct Antigravity IDE commands in `gitmap-developer-hygiene-and-agy`: replace `gitmap agy clean-cache` with `gitmap agy cache-clear` (`cc`, `ccko`, `cckf`), add `remove-empty-convs` and `running-projects ls`.
   - Correct Chrome profile commands: replace `gitmap chrome profiles --reconcile` and `picker` with canonical `gitmap chrome reconcile`, `gitmap chrome list` (`ls`), `gitmap chrome observe`, and `gitmap chrome open`.
3. **Enforce 1:1 Dual Skill Tree Parity:**
   - Ensure every modification in `.agents/skills/` is identically mirrored to `.cursor/skills/`.

---

## 2. Disjoint File Boundaries

### Owned Files (Exclusive Modification Rights for Worker 03):
- `.agents/skills/gitmap/skill.md`
- `.cursor/skills/gitmap/skill.md`
- `.agents/skills/gitmap-developer-hygiene-and-agy/skill.md`
- `.cursor/skills/gitmap-developer-hygiene-and-agy/skill.md`
- `.agents/skills/gitmap-scanner-and-cloner/skill.md`
- `.cursor/skills/gitmap-scanner-and-cloner/skill.md`
- `.agents/skills/gitmap-ssh-cluster-fleet/skill.md`
- `.cursor/skills/gitmap-ssh-cluster-fleet/skill.md`
- `.agents/skills/gitmap-split-db-engine/skill.md`
- `.cursor/skills/gitmap-split-db-engine/skill.md`
- `.agents/skills/gitmap-macro-automation-engine/skill.md`
- `.cursor/skills/gitmap-macro-automation-engine/skill.md`
- `.agents/skills/gitmap-pipeline-and-diagnostics/skill.md`
- `.cursor/skills/gitmap-pipeline-and-diagnostics/skill.md`

### Read-Only Context Files:
- `02-spec/21-app/24-coding-guideline-and-prompt-execution-improvement/03-migration-and-verification-spec.md`
- `AGENTS.md`
- `.ai-memory/plans/coding-guideline-and-prompt-execution-improvement.md`

---

## 3. Negative Boundaries & Invariants

- **TOTAL BAN ON GIT COMMANDS:** Subagents and workers must NEVER run `git` commands (`git add`, `git commit`, `git push`, `git status`, `git checkout`). Only the lead orchestrator performs atomic commits at the completion of Phase 3.
- **TOTAL BAN ON RAW SHELL SEARCH TOOLS:** Never run `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, `Get-ChildItem` (all variants), `dir`, or `ls`. Use GitMap commands (`gitmap aum search`, `gitmap lf`, `gitmap find`) exclusively.
- **Strict Relative Git Paths:** Zero absolute paths (`C:\...`, `/home/...`) or `file:///` URIs inside any markdown or documentation file.
- **Strict Lowercase Filenames:** All filenames, scripts, and plans must be strictly lowercase.
- **100% Skill Parity:** Both `.agents/skills/` and `.cursor/skills/` must remain in exact 1:1 synchronization.
- **Cognitive Heatmap Tracking:** All file actions must be logged via `gitmap heatmap log` and closed with `gitmap heatmap update`.

---

## 4. Step-by-Step Implementation Guide for Worker 03

### Step 1: Enrich `.agents/skills/gitmap/skill.md` with 12-Category Examples
1. Open `.agents/skills/gitmap/skill.md`.
2. Under Section "## Essential Command Cheat Sheet", embed the comprehensive 12-category real-world example catalog defined in Section 2 of `02-spec/21-app/24-coding-guideline-and-prompt-execution-improvement/03-migration-and-verification-spec.md`.
3. In Section 9 ("Multi-Repo, Cluster & Toolchain Operations"):
   - Replace `gitmap cluster --help` with `gitmap cluster status`.
   - Replace `gitmap sc --help` with `gitmap sc list`.
   - Replace `gitmap ssh --help` with `gitmap ssh health --all`.
4. Mirror all edits identically to `.cursor/skills/gitmap/skill.md`.

### Step 2: Audit and Realign `gitmap-developer-hygiene-and-agy`
1. Open `.agents/skills/gitmap-developer-hygiene-and-agy/skill.md` and `.cursor/skills/gitmap-developer-hygiene-and-agy/skill.md`.
2. In Section 2 ("Essential Commands"):
   - Replace `gitmap agy clean-cache` with `gitmap agy cache-clear` (and note aliases `cc`, `ccko`, `cckf`).
   - Replace `gitmap agy prune --except <conversation-id>` with `gitmap agy remove-empty-convs` and `gitmap agy cache-clear --keep 10`.
   - Add `gitmap agy running-projects ls` (alias `rp`) and `gitmap agy last-active-projects 24` (alias `lap`).
   - Replace `gitmap chrome profiles --reconcile` with `gitmap chrome reconcile` (aliases: `repair`, `sync-state`).
   - Replace `gitmap chrome picker` with `gitmap chrome list` (alias `ls`, `cpl`), `gitmap chrome observe` (alias `tabs`), and `gitmap chrome open [urls]`.
   - Add `gitmap chrome reset [prof] [flags]` and `gitmap chrome diff <prof1> <prof2>`.

### Step 3: Audit and Realign `gitmap-scanner-and-cloner`
1. Open `.agents/skills/gitmap-scanner-and-cloner/skill.md` and `.cursor/skills/gitmap-scanner-and-cloner/skill.md`.
2. In Section 2 ("Essential Commands"):
   - Verify cloning command explicitly uses `--concurrency 8` (e.g. `gitmap clone [manifest.json] --concurrency 8`).
   - Add cautionary note prohibiting hallucinated `--max-concurrency`.
   - Confirm scan flags: `gitmap scan [path]`, `gitmap scan [path] --mode ssh`, `gitmap scan [path] --github-desktop`.

### Step 4: Audit and Realign `gitmap-ssh-cluster-fleet`
1. Open `.agents/skills/gitmap-ssh-cluster-fleet/skill.md` and `.cursor/skills/gitmap-ssh-cluster-fleet/skill.md`.
2. In Section 2 ("Essential Commands"):
   - Confirm commands use active inspection: `gitmap cluster status`, `gitmap sc list`, `gitmap ssh health --all`.
   - Add `gitmap cluster exec "uptime"` and `gitmap sc bash "<cmd>"`.
   - Confirm zero occurrences of `--help`.

### Step 5: Audit and Realign Remaining Companion Skills
1. Review `gitmap-split-db-engine`, `gitmap-macro-automation-engine`, and `gitmap-pipeline-and-diagnostics`.
2. Ensure no `--help` commands or hallucinated flags exist.
3. Validate positive boolean rules and strict relative paths.
4. Mirror all changes to `.cursor/skills/`.

---

## 5. Verification & Conformance Checklist

- [x] **Checklist Item 1 (Examples Catalog):** All 12 categories from the specification are embedded in `.agents/skills/gitmap/skill.md` with concrete syntax and explanations.
- [x] **Checklist Item 2 (Zero `--help` Flags):** Zero occurrences of `--help` in `.agents/skills/gitmap*/` and `.cursor/skills/gitmap*/`.
- [x] **Checklist Item 3 (Concurrency Flag Standard):** All cloning and syncing commands use `--concurrency` or `--workers`; zero occurrences of `--max-concurrency`.
- [x] **Checklist Item 4 (AGY Commands):** `cache-clear`, `ccko`, `cckf`, `remove-empty-convs`, and `running-projects` documented accurately.
- [x] **Checklist Item 5 (Chrome Commands):** `reconcile`, `list`, `observe`, `open`, and `reset` documented accurately.
- [x] **Checklist Item 6 (Skill Tree Parity):** Exact 1:1 line hygiene between `.agents/skills/` and `.cursor/skills/`.
- [x] **Checklist Item 7 (Relative Paths):** Zero absolute paths and zero `file:///` URIs across all modified skills.
- [x] **Checklist Item 8 (Heatmap Closure):** 100% of cognitive heatmap action entries closed with `isExpectationMet: true`.
