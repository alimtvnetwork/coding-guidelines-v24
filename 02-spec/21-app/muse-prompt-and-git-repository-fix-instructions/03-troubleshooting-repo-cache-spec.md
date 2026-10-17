# Muse Git Troubleshooting & Repo-Cache Forensics: Specification

> **/goal** Specify the architecture, directory structure, canonical 6-artifact diagnostic suite, `gitmap rc` integration, and Muse skill updates for git repository troubleshooting and forensic preservation in `repo-cache`.
> **/learn** Enforce the "Persist Before Proceeding" invariant, eliminate diagnostic scratch script pollution in standard repositories, and upgrade Muse concurrency to `A = 3` with immediate dynamic respawn.

---

## 1. Context & Authority

### 1.1 Authority Mandate
This specification operates under the direct authority of `AGENTS.md` §9 ("Special Default Repositories (`repo-secrets` / `rs` & `repo-cache` / `rc`)"):
- **Zero Scratch Scripts in Standard Repositories:** Whenever temporary scripts (such as PowerShell `.ps1` scripts, diagnostic harnesses, or reusable test fixtures) are created during development, debugging, or troubleshooting, they MUST NOT be left in standard git repositories or project worktrees.
- **Dedicated Repo-Cache Preservation:** Diagnostic harnesses, recovery runbooks, and error autopsies must be offloaded directly to `repo-cache` via `gitmap rc` so that all repositories can reuse them cleanly without polluting git history.
- **Troubleshooting Scope Isolation:** The troubleshooting section in Muse is triggered **only** when troubleshooting or fixing a git repository or pipeline issue, keeping routine operations lean.

### 1.2 Core Problem Statement
In previous autonomous runs, when Muse or subagents encountered git conflicts, diverged histories, locked index files, or pipeline errors:
1. Ad-hoc scratch scripts and test commands were dropped into the working tree, risking accidental commits to production repositories.
2. Troubleshooting knowledge and successful remediation sequences were lost once the chat session ended.
3. Speculative trial-and-error commands were executed without checking whether identical git anomalies had already been analyzed and solved.
4. Concurrency was capped at `A = 2` without immediate dynamic replacement when workers finished early, causing idle cycles.

---

## 2. Repo-Cache Troubleshooting Vault Architecture

### 2.1 Standard Vault Location
All diagnostic data, git state snapshots, repair harnesses, and remediation runbooks generated during troubleshooting must be housed in `repo-cache`:

```text
repo-cache/muse/troubleshooting/<task-slug>/
├── 01-incident-manifest.json
├── 02-error-autopsy.md
├── 03-git-state-snapshot.json
├── 04-harness-scripts/
│   ├── diagnose-conflicts.ps1
│   └── verify-tree.sh
├── 05-remediation-runbook.md
└── 06-verification-log.md
```

- `<task-slug>`: The unique hyphenated task slug corresponding to the active incident (e.g., `git-conflict-sync-engine`, `detached-head-recovery-v28`).
- Directory creation and persistence are handled via native GitMap commands (`gitmap rc`).

---

## 3. Canonical 6-Artifact Diagnostic Suite

Whenever Muse or an authorized worker enters a troubleshooting or fixing state, the agent must construct and persist the following 6 canonical artifacts:

### 3.1 Artifact 01: `01-incident-manifest.json`
Machine-readable metadata capturing the operational context of the incident:

```json
{
  "task_slug": "example-task-slug",
  "target_repo": "gitmap-v28",
  "branch": "main",
  "incident_category": "merge-conflict",
  "severity": "high",
  "status": "investigating",
  "operator_name": "opRoden",
  "created_at": "2026-10-18T05:00:00Z",
  "has_uncommitted_changes": true,
  "is_resolved": false,
  "affected_files_count": 3
}
```

**Field Rules:**
- `incident_category` MUST be one of:
  * `merge-conflict`
  * `diverged-history`
  * `detached-head`
  * `index-lock-collision`
  * `corrupted-ref`
  * `pipeline-failure`
  * `submodule-divergence`
- `severity` MUST be one of: `low`, `medium`, `high`, `critical`.
- `status` MUST be one of: `investigating`, `mitigating`, `verifying`, `resolved`, `escalated`.
- Boolean fields MUST follow positive naming (`has_uncommitted_changes`, `is_resolved`).

### 3.2 Artifact 02: `02-error-autopsy.md`
Grounded 4-part Root Cause Analysis (RCA) adhering to `AGENTS.md` and `fix-with-rca`:

1. **Symptom:** Verbatim error output, failing command, exit code, and observed repository anomaly.
2. **Trigger:** The exact operation, tool call, concurrent action, or command that immediately triggered the failure.
3. **Root Cause:** Underlying structural, git graph, file lock, or environment reason why the failure occurred.
4. **Mitigation:** Comprehensive remediation plan to restore git cleanliness and the permanent safeguard to prevent recurrence.

### 3.3 Artifact 03: `03-git-state-snapshot.json`
Structured snapshot of the git graph and working tree at the moment of incident detection:

```json
{
  "head_sha": "d3b07384d113edec49eaa6238ad5ff00",
  "remote_tracking_sha": "c1a08273b113edec49eaa6238ad5aa11",
  "ahead_count": 2,
  "behind_count": 5,
  "conflicted_files": [
    "pkg/engine/sync.go"
  ],
  "modified_files": [
    "pkg/engine/types.go"
  ],
  "untracked_files": [],
  "has_detached_head": false,
  "is_rebasing": false,
  "is_merging": true,
  "porcelain_status": "UU pkg/engine/sync.go\nM  pkg/engine/types.go"
}
```

### 3.4 Artifact 04: `04-harness-scripts/`
A dedicated directory containing any reusable diagnostic scripts, test reproduction harnesses, or verification probes:
- Standalone PowerShell (`.ps1`) or Bash (`.sh`) scripts that inspect file locks, verify tree integrity, or parse conflict regions.
- Scripts must be self-contained and free of proprietary credentials or environment-specific hardcoded paths.
- Preserved via `gitmap rc file` or `gitmap rc folder`.

### 3.5 Artifact 05: `05-remediation-runbook.md`
A strictly ordered, reproducible sequence of commands executed to recover the repository:
- Exact commands with parameter flags.
- Pre-conditions required before each command.
- Verification checkpoint after each step to confirm expected state transition.
- Explicit rollback command for each phase if a step fails.

### 3.6 Artifact 06: `06-verification-log.md`
Empirical proof of full recovery:
1. **Porcelain Cleanliness:** Output proving 0 unexpected dirty, staged, or untracked files.
2. **Zero Conflict Markers:** Proof that `gitmap aum search` or string search found zero occurrences of `<<<<<<<`, `=======`, or `>>>>>>>`.
3. **Graph Reconciliation:** Commit graph log showing clean linear history or validated merge commit.
4. **CI/Telemetry Health:** `gitmap pe -t` output demonstrating green pipeline or passing validation gate.

---

## 4. GitMap rc Command Integration

### 4.1 Native `gitmap rc` Operations
All interactions with the troubleshooting vault use native `gitmap rc` commands:

| Action | GitMap Command | Description |
|:---|:---|:---|
| **Navigate to Vault** | `gitmap cd rc` | Navigates directly into the local `repo-cache` root. |
| **Preserve Single File** | `gitmap rc file <path> --repo muse/troubleshooting/<task-slug>` | Copies file into target folder in `repo-cache`, auto-commits and pushes upstream. |
| **Preserve Folder** | `gitmap rc folder <dir> --repo muse/troubleshooting/<task-slug>` | Recursively copies folder into target location in `repo-cache`, auto-commits and pushes. |
| **Write Runbook / Text** | `gitmap rc text "<content>" --slug <slug> --repo muse/troubleshooting/<task-slug>` | Writes content to `<slug>.md`, auto-commits and pushes upstream. |

### 4.2 The "Persist Before Proceeding" Invariant
- **Rule:** An agent is **strictly prohibited** from declaring a git issue resolved, closing a subtask, or completing Turn execution before all 6 artifacts are persisted to `repo-cache` via `gitmap rc`.
- **Enforcement:** If a worker resolves a merge conflict or git graph divergence but leaves the troubleshooting suite uncommitted to `repo-cache`, the subtask is marked incomplete (`FAILED-PERSISTENCE-INVARIANT`).

### 4.3 Troubleshooting Recall Protocol (Prior Art First)
Before executing speculative git repair commands (`git reset`, `git rebase`, `git merge --abort`, `git checkout -b`):
1. **Search Prior Runbooks:** The agent must inspect `repo-cache/muse/troubleshooting/` using `gitmap find` and `gitmap cat`.
2. **Match Incident Category:** If an existing runbook with the same `incident_category` exists, review its `05-remediation-runbook.md` and `02-error-autopsy.md`.
3. **Execute Proven Playbook:** Apply the verified remediation sequence rather than guessing.

---

## 5. Muse Skill Architecture Updates

The skills `.agents/skills/muse-master-prompt/skill.md` and `.cursor/skills/muse-master-prompt/skill.md` must be updated with the following core architectural changes:

### 5.1 Concurrency Upgrade (`A = 3`, `A × H = 6`)
- **Worker Concurrency `A`:** Raised from 2 to 3.
- **Hands per Agent `H`:** Kept at 2.
- **System Concurrency Capacity:** `A × H = 3 agents × 2 hands = 6 concurrent subtask operations`.
- **Subagent Spawning Gate:** Mandatory invocation of `A = 3` subagents at every phase:
  1. Planning Phase: 3 `research` subagents (`TypeName: "research"`).
  2. Spec Phase: 3 `self` subagents (`TypeName: "self"`).
  3. Execution Phase: 3 `self` worker subagents (`TypeName: "self"`).
- **Solo Execution Ban:** Lead orchestrator solo execution of multi-part work remains an auto-reject failure.

### 5.2 Immediate Dynamic Respawn Engine
- **Continuous Saturation:** The orchestrator must not wait for all 3 workers in a wave to finish before spawning the next wave.
- **Dynamic Slot Re-fill:** The exact millisecond any single worker reports back via `send_message` or finishes its bounded task, the lead orchestrator immediately spawns a replacement subagent to claim the next pending subtask from `.ai-memory/temp-agents/<slug>/agent-task.db`.
- **Dynamic Join Fallback:** When pending unassigned subtasks drops below 3, available subagents dynamically join in-flight tasks to assist with parallel reading, pre-verification, companion lint checks, or edge case testing, ensuring 0 idle agent slots while respecting the strict disjoint file box invariant.

### 5.3 Dedicated Troubleshooting Section Insertion
Add Section 4 ("Git Repository Troubleshooting & Repo-Cache Forensics Protocol") to both Muse skills:
- Activation rule: Triggered ONLY when troubleshooting or fixing a git repository.
- Workflow execution:
  1. Detect incident and classify into `incident_category`.
  2. Recall prior runbooks from `repo-cache/muse/troubleshooting/`.
  3. Capture initial git snapshot (`03-git-state-snapshot.json`) and manifest (`01-incident-manifest.json`).
  4. Perform 4-part RCA (`02-error-autopsy.md`).
  5. Execute remediation steps and log to `05-remediation-runbook.md`.
  6. Verify recovery and document in `06-verification-log.md`.
  7. Push entire suite to `repo-cache` using `gitmap rc folder`.

### 5.4 Subagent Git Mutation Ban Reminder
- Subagents (`TypeName: "self"`) are strictly forbidden from executing git state mutation commands (`git add`, `git commit`, `git push`, `git checkout`, `git reset`, `git rebase`).
- All git repository troubleshooting commands and `gitmap rc` operations are orchestrated by the Lead Agent or dedicated troubleshooting workers under strict isolation.

---

## 6. Acceptance Criteria

- [ ] Directory architecture specifies `repo-cache/muse/troubleshooting/<task-slug>/`.
- [ ] 6 canonical artifacts detailed with full schemas and positive boolean standards (`is_resolved`, `has_uncommitted_changes`).
- [ ] `gitmap rc` commands (`file`, `folder`, `text`, `cd rc`) clearly mapped.
- [ ] "Persist Before Proceeding" invariant defined with zero tolerance for dirty repos.
- [ ] Troubleshooting recall protocol mandates checking `repo-cache` before speculative git execution.
- [ ] Muse skills update plan details `A = 3`, `A x H = 6`, immediate dynamic respawn, dynamic join, and Section 4 addition.
- [ ] All file paths in this spec are strictly relative to git root with zero absolute paths and zero `file:///` URIs.
- [ ] Strictly lowercase filenames and paths used throughout.
