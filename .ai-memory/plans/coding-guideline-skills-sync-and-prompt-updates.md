# Plan — coding-guideline-skills-sync-and-prompt-updates

- slug: `coding-guideline-skills-sync-and-prompt-updates`
- repo: `coding-guidelines-v24` (github.com/alimtvnetwork/coding-guidelines-v24, private)
- spec: `02-spec/21-app/coding-guideline-skills-sync-and-prompt-updates/01-skills-sync-and-prompt-updates-spec.md`
- orchestrator: coordinator subagent, v6 playbook (A=2 per phase, disjoint file boxes)
- status: in progress

## User Request (Verbatim)

From the user's "Learning and Storytelling Audio Creation" prompt, coding-guidelines
track (lossless capture in parent context): re-read the gitmap skills and bring
new/updated ones into the codebase; build a synchronization mechanism inside
coding-guidelines-v24 for repos that do not need spec (sync prompts + skills,
never spec); update the letterly prompt with a new version whose first
instruction is Git pull (`gitmap pull`, fallback to regular `git pull`); update
the muse prompt similarly (pull-first for any repo, add story-telling-v1 and
other missing repos to the registry, add a do-not-wait-for-approval checklist
item); commit + release coding-guidelines-v24 with green CI; sync the skills and
prompts into this repo; re-read the gitmap skills. Standing: turbo mode, act
without confirmation; never delete.

## Task breakdown

### Task-01: spec + plan + subtasks (lead)
- **State:** `[DONE]`
- **Understood:** `[YES]` — plan-first per v6: SPEC.md, this plan, 4 subtask files.
- **Target Files:** `02-spec/21-app/coding-guideline-skills-sync-and-prompt-updates/`, `.ai-memory/plans/`

### Task-02: gitmap skills inventory (W1)
- **State:** `[DONE]` (2026-10-17: "Skills inventory" section appended below —
  8 gitmap skills, 4 NEW, `gitmap` core CHANGED 2026-10-16, 6 subsystems stable;
  `gitmap-agent-orchestrator` has no repo folder, coverage folded into core `gitmap` skill)
- **Understood:** `[YES]` — inventory new/updated gitmap skills in-repo, diff vs ingested set, record in plan.
- **Target Files:** `01-prompts/26-gitmap/`, `.agents/skills/`, `.cursor/skills/` (read-only) + this plan (inventory section)

### Task-03: prompts+skills sync mechanism (W1)
- **State:** `[DONE]` (2026-10-17: `03-ai-scripts/50-sync-prompts-skills-to-targets.py`
  + `03-ai-scripts/50-sync-prompts-skills-to-targets.md` + registry
  `scripts/prompt-sync-targets.json` (story-telling-v1 example) created;
  `03-ai-scripts/readme.md` index updated; test-run vs `~/workspace/custom-skills`
  — 620 copied run 1, 0 changes run 2 (idempotent); `02-spec/` exclusion
  gate self-checked: all spec paths denied, sync roots allowed)
- **Understood:** `[YES]` — re-runnable script + doc inside this repo; syncs prompts + skills, never `02-spec/`; test vs `~/workspace/custom-skills`.
- **Target Files:** `03-ai-scripts/50-sync-prompts-skills-to-targets.py` (+ doc)

### Task-04: letterly prompt new version (W2)
- **State:** `[DONE]` (2026-10-17: `01-prompts/22-letterly/11-git-pull-first-letterly.md`
  created — pull-first as Actionable Item 1: `gitmap pull` → fallback `git pull`,
  all letterly invariants kept; `22-letterly/readme.md` index row 11 added.
  Verified by lead: file read, pull-first wording confirmed.)
- **Understood:** `[YES]` — `01-prompts/22-letterly/11-<slug>-letterly.md` with pull-first as step 1; readme index updated.
- **Target Files:** `01-prompts/22-letterly/`

### Task-05: muse prompts new versions (W2)
- **State:** `[DONE]` (2026-10-17: `01-prompts/27-muse-prompts/04-muse-master-with-repos-v2.md`
  v1.2.0 — pull-first in operating contract + step 5 of protocol + fallback table,
  registry expanded to 13 repos (story-telling-v1, macro-ahk, scripts-fixture,
  cat-my-v12 "Cat Malaysia"), no-wait checklist; `05-muse-master-prompt-v2.md`
  v6.2.0 — pull-first added (was absent; no-wait already covered, minimal
  19+/4- diff); category `readme.md` catalog updated. Verified by lead via
  `gitmap aum search`: pull-first hits at lines 14/62/72/82, no-wait bullet present.
  `01-`/`03-` originals untouched.)
- **Understood:** `[YES]` — new versions in `01-prompts/27-muse-prompts/` with pull-first rule, expanded registry (story-telling-v1, macro-ahk, scripts-fixture, cat-my-v12 "Cat Malaysia"), no-wait checklist; readme catalog updated.
- **Target Files:** `01-prompts/27-muse-prompts/`

### Task-06: commit + release + CI green (lead)
- **State:** `[IN PROGRESS]`
- **Understood:** `[YES]` — `gitmap cpf`, minor release via `bump_versions.py --create-release`, CI green via `gh`.
- **Target Files:** repo root (`version.json` via script)

### Task-07: final report (lead)
- **State:** `[TODO]`

## Wave plan

- Phase 1 (lead): Task-01 — DONE (this plan).
- Phase 2 (parallel): Task-02 + Task-03 (W1) | Task-04 + Task-05 (W2) — disjoint boxes.
- Phase 3 (lead): Task-06 + Task-07.

## Skills inventory (W1 — recorded 2026-10-17, tree `0c5fad76`)

Baseline reference: ingested set from memory — `gitmap`, `gitmap-agent-orchestrator`
+ 6 subsystems (scanner-and-cloner, ssh-cluster-fleet, split-db-engine,
pipeline-and-diagnostics, developer-hygiene-and-agy, macro-automation-engine).

### Raw counts

- `.agents/skills/`: **176 skill folders** (each carries `skill.md`). `.cursor/skills/`
  is an **exact mirror** — `diff` over the directory listings returns empty.
- gitmap-named skills present: `gitmap`, `gitmap-scanner-and-cloner`,
  `gitmap-ssh-cluster-fleet`, `gitmap-split-db-engine`,
  `gitmap-pipeline-and-diagnostics`, `gitmap-developer-hygiene-and-agy`,
  `gitmap-macro-automation-engine`, `gitmap-pipeline-waiting-prompts` (8 total).
- `01-prompts/26-gitmap/`: `01-gitmap-core-engine.md` (Prompt Version: **6.0.0**,
  "Runtime: Google Antigravity 2.0 (IDE and CLI) & Cursor IDE") + `readme.md`.

### NEW — in this repo, NOT in the ingested set

1. `gitmap-pipeline-waiting-prompts` — `.agents/skills/gitmap-pipeline-waiting-prompts/skill.md`
   (1.3 KB, added 2026-09-18, commit `dd1bdd09`): `gitmap pipeline-ai status -t <eta>`
   dynamic-wait protocol + anti-busy-poll mandate, meant to be baked into
   execute/RCA/CI-CD prompts.
2. `execute-parent-task-with-n-steps-v6` — `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md`
   (59 KB, imported 2026-10-15 from go-email-reader work): v6 parent-orchestrator
   playbook with SQLite task DB + crash forensics.
3. `sync` — `.agents/skills/sync/skill.md` (19.8 KB, updated 2026-10-16): native
   `gitmap sync` + native `gitmap agent` task-engine sync/repo-deployment skill.
   NOTE: its description still says it syncs "shared specs (02-spec/01-* to
   02-spec/20-*)" — this DIVERGES from this program's `02-spec/` never-sync
   stance; see adoption notes.
4. `muse-master-prompt` — `.agents/skills/muse-master-prompt/skill.md` (6.7 KB,
   added 2026-10-09, commit `e9447e96`): companion skill for category
   `01-prompts/27-muse-prompts/`.
5. `prompts-and-skills-sync` (pairs with scripts `38/42/43` sync family).

### CHANGED — since the ingested snapshot

1. `gitmap` core skill — `.agents/skills/gitmap/skill.md` (17.4 KB), updated
   **2026-10-16** (commit `1bd15953`): `gitmap task` → `gitmap agent`
   command migration (`task/subtask/log/diagnose/temp-clear`); legacy
   `46-agent-sqlite-task-manager.py` fallback purged; new §11 "Fleet Telemetry
   & Running-Only Delegation" (`gitmap nro`/`nodes-running-only`, `gitmap os`,
   `gitmap cpu`, `gitmap ram`, `gitmap disk`, `gitmap network`,
   `gitmap nodes tasks`). Also carries post-ingestion features: `mr`/`merge
   write`, `help example <cmd>`, 4 CLI themes (navy-purple default, tokyo-night,
   vscode-dark-plus, monokai-pro), `pipeline-ai` waiting, `login --web`.
2. `sync` skill + `03-ai-scripts/38-sync-prompts-skills-scripts.py` — updated
   **2026-10-17** (commit `0c5fad76`): "keep repository-local skill folders
   during skill synchronization" — a fail-safe directly relevant to WS2's
   new sync script.
3. `execute-parent-task-with-n-steps-v6` — 13+/14- line rework in `1bd15953`
   (same legacy task-manager purge).
4. `muse-master-prompt` — 2-line tweak in `1bd15953`.

### UNCHANGED (stable)

- The 6 subsystem skills: last touched **2026-09-22** (commit `ea3379a7`,
  "feat(skills): add 7 gitmap subsystem and core skills") — predate the
  2026-10-08 ingestion, so the ingested copies are current.

### Missing vs the ingested set

- `gitmap-agent-orchestrator` has **no folder** in this repo and no presence in
  repo git history (checked with `git log --all` + repo-wide
  `gitmap aum search`). Its coverage appears folded into the core `gitmap`
  skill (§11 fleet telemetry + native `gitmap agent` engine) and the `sync`
  skill — so the standalone orchestrator skill name survives only in these
  program docs.

### Adoption notes

- Re-ingest the **core `gitmap` skill** from this repo: the ingested copy
  predates the `gitmap agent` migration and the fleet-telemetry §11 — stale
  copies would teach `gitmap task`/`46-*.py` commands that no longer exist.
- Adopt `gitmap-pipeline-waiting-prompts` + `execute-parent-task-with-n-steps-v6`
  + `sync` + `muse-master-prompt` into the working skill set; they are new
  since ingestion and have no counterpart in the old set.
- The new `03-ai-scripts/50-sync-prompts-skills-to-targets.py` (WS2) keeps
  `02-spec/` out of targets — deliberately stricter than the existing `sync`
  skill's shared-spec allowance; do not "fix" the script to match the skill.
- Because `.cursor/skills/` mirrors `.agents/skills/` exactly, the sync
  mechanism ships both folders as a unit.

## Conflicts

- none.
