# Specification: Muse Prompts Elevation, Concurrency Scaling (A=3, H=2) & Git Repository Troubleshooting Forensics

## 1. Executive Summary & Purpose

This specification defines the architectural enhancements, operational invariants, and concrete runbooks for upgrading the canonical Muse AI prompts and companion workflows:

1. **Concurrency Scaling to 3 Parallel Subagents (`A = 3`, `H = 2`):**
   - Elevate Muse's execution concurrency from 2 to 3 autonomous subagents (`A = 3`, `H = 2`), expanding system concurrency capacity to `3 × 2 = 6` operations.
   - Enforce mandatory 3-stage dispatch: Stage 1 (Planning: 3 concurrent `research` subagents), Stage 2 (Spec Authoring: 3 concurrent `self` subagents), and Stage 3 (Execution: 3 concurrent `self` worker subagents).
   - Mandate immediate dynamic respawn: the exact moment a subagent completes or returns, the Lead Orchestrator immediately claims the next subtask and spawns a replacement subagent, maintaining 100% saturation across all 3 slots.
   - Dynamic Join Fallback: when fewer than 3 pending subtasks remain, available subagents dynamically join running subtasks for parallel reading, syntax linting, and dependency mapping under strict disjoint-ownership rules (Rule: One Owner Per File).

2. **Dedicated Git Repository Troubleshooting & Repo-Cache Forensics (Phase 7.5):**
   - Embed a dedicated section within `01-prompts/27-muse-prompts/01-muse-master-prompt.md`: "Phase 7.5 — Git Repository Troubleshooting & Repo-Cache Forensics".
   - Conditional Execution Directive: Only when troubleshooting or active Git remediation is required does Muse create, update, or write to the troubleshooting archive. Nominal clean runs must never create empty forensic logs or dummy directories.
   - Forensic Archive Schema: All Git remediation actions, diagnostic traces, error messages, and task metadata must be persisted to `repo-cache` under `muse/troubleshooting/<task-slug>/` via `gitmap rc` so Muse and subsequent sessions can inspect, learn from, and reuse past resolutions.

3. **Non-Destructive Git Remediation Invariants:**
   - Strict preservation of published Git history: TOTAL BAN on `git push --force` / `git push -f` and rebasing published commits (protecting Lovable and remote tracking histories).
   - Strict adherence to the Worker Git Ban: Worker subagents report Git anomalies and errors to the Lead Orchestrator via `send_message`; only the Lead Orchestrator inspects and remediates repository Git state.
   - Standardized non-destructive command patterns for branch divergence reconciliation, mechanical conflict resolution, index lock contention clearance, and workspace debris isolation via `gitmap rc`.

---

## 2. Muse 3-Agent Concurrency & Dynamic Respawn Architecture

### 2.1 Concurrency Topology (`A = 3`, `H = 2`, Capacity = 6)

In `01-prompts/27-muse-prompts/01-muse-master-prompt.md`, top-level orchestration parameters are elevated to:

```text
N = 300 (Total self-loop steps budget — editable top-header parameter, default: 300)
A = 3   (MANDATORY number of spawned autonomous subagents running concurrently via invoke_subagent, default: 3)
H = 2   (Operational hands per agent: dual-task batch capacity & parallel tool dispatch, default: 2)
C = 30  (Tool calls per worker before it must report, default: 30)

System Concurrency Capacity = A × H = 3 agents × 2 hands = 6 concurrent subtask operations
PHASE_1_BUDGET = N / 2   (Steps 1 .. 150: Discovery, Architecture Spec, and Lean Subtask Generation)
PHASE_2_BUDGET = N / 2   (Steps 151 .. 300: Mandatory Parallel Subagent Execution, Self-Looping, Targeted Quality Linting)
WAVES = ceil(subtasks / (A x H))
```

### 2.2 Mandatory 3-Stage Dispatch Protocol

Every multi-part execution sequence in Muse must dispatch through three distinct stages, saturating all 3 subagent slots:

```
                      ┌─────────────────────────────────────────┐
                      │      Lead Orchestrator (Root Node)      │
                      │     Capacity Controller (A=3, H=2)      │
                      │    Task Ledger & SQLite DB Custodian    │
                      └────────────────────┬────────────────────┘
                                           │
         ┌─────────────────────────────────┼─────────────────────────────────┐
         ▼                                 ▼                                 ▼
┌─────────────────┐               ┌─────────────────┐               ┌─────────────────┐
│  Worker Slot 01 │               │  Worker Slot 02 │               │  Worker Slot 03 │
│  (A1 Subtask)   │               │  (A2 Subtask)   │               │  (A3 Subtask)   │
└────────┬────────┘               └────────┬────────┘               └────────┬────────┘
         │                                 │                                 │
         ├─────────────────────────────────┼─────────────────────────────────┤
         ▼                                 ▼                                 ▼
[Stage 1: Discovery]              [Stage 1: Discovery]              [Stage 1: Discovery]
3 x research subagents            3 x research subagents            3 x research subagents
         ▼                                 ▼                                 ▼
[Stage 2: Modular Spec]           [Stage 2: Modular Spec]           [Stage 2: Modular Spec]
3 x self spec authors             3 x self spec authors             3 x self spec authors
         ▼                                 ▼                                 ▼
[Stage 3: Bounded Exec]           [Stage 3: Bounded Exec]           [Stage 3: Bounded Exec]
3 x self worker agents            3 x self worker agents            3 x self worker agents
```

1. **Stage 1 — Planning & Exploration (`A = 3`, `TypeName: "research"`):**
   - Lead spawns 3 concurrent read-only research subagents via `invoke_subagent`.
   - Subagents explore codebase architecture, verify symbol usage via `gitmap aum search`, and extract interface signatures.
   - Lead synthesizes findings into the master architecture plan.

2. **Stage 2 — Modular Spec Authoring (`A = 3`, `TypeName: "self"`):**
   - Lead spawns 3 concurrent authoring subagents to produce disjoint specification files under `02-spec/21-app/<slug>/` and subtask files under `.ai-memory/plans/subtasks/<slug>/`.
   - File boundaries are mutually exclusive; zero overlap occurs between authoring targets.

3. **Stage 3 — Bounded Execution (`A = 3`, `TypeName: "self"`):**
   - Lead spawns 3 concurrent worker subagents to apply surgical edits to strictly disjoint file boxes.
   - Subagents operate under the strict Worker Git Ban (edits and local checks only; no Git mutations).

### 2.3 Immediate Dynamic Respawn & Saturation Protocol

To prevent idle execution bubbles and maximize throughput:

1. **Instant Replacement Trigger:**
   The exact instant a subagent sends a completion or status message via `send_message`, the Lead Orchestrator processes the result, updates the SQLite task database (`gitmap agent subtask complete`), and immediately invokes the next pending subtask from the queue.
2. **Duty-Cycle Maintenance:**
   The Lead Orchestrator must never wait for an entire wave of 3 agents to finish before spawning next items. Slots are refilled immediately and independently as each agent completes.
3. **Dynamic Join Fallback:**
   When the remaining pending subtasks are fewer than 3 (`pending_count < A`):
   - Available subagent slots transition to Dynamic Join mode (`gitmap agent subtask join`).
   - The joining subagent assists the active primary worker by performing companion tasks: parallel file analysis, AST linting, reference mapping, or memory log generation.
   - Rule R6 Invariant: The primary worker retains exclusive write locks on target files. The joining subagent performs read-only and analytical assistance without file collisions.

---

## 3. Dedicated Troubleshooting Section in Muse Prompts (Phase 7.5)

### 3.1 Placement & Trigger Invariant

Phase 7.5 is positioned directly between Phase 7 (Verification, Secrets Gate, Atomic Push & CI Telemetry) and Phase 8 (Muse Memory IO) in `01-prompts/27-muse-prompts/01-muse-master-prompt.md`.

> [!IMPORTANT]
> **Conditional Execution Invariant:**
> Phase 7.5 is strictly conditional. It activates **ONLY** when a Git anomaly, merge conflict, repository divergence, index lock collision, or dirty working tree blockage is encountered during execution.
> When all operations are nominal, Muse skips Phase 7.5 entirely and MUST NOT write empty files or create empty directories in `repo-cache`.

### 3.2 Repo-Cache Forensics Archive Specification

When Git troubleshooting is required, Muse must persist comprehensive forensic records into `repo-cache` to enable cross-session learning and rapid automated recovery.

#### Directory Layout:
```text
repo-cache/
└── muse/
    └── troubleshooting/
        └── <task-slug>/
            ├── 01-diagnosis.md           # Error description, divergence metrics & root cause
            ├── 02-remediation-steps.md   # Executed non-destructive Git commands & verification
            └── 03-forensic-log.txt       # Verbatim console output and diff summaries
```

#### Archive Content Requirements:

1. **`01-diagnosis.md`**:
   - Timestamp and active repository name.
   - Nature of failure: branch divergence, index lock, merge conflict, or dirty worktree.
   - Quantitative state: output of `git status --porcelain`, `git rev-list --left-right --count`, or lock timestamp.
   - Root cause classification (concurrent worker edit, upstream remote push, interrupted process).

2. **`02-remediation-steps.md`**:
   - Exact sequence of non-destructive commands executed.
   - Rollback safeguards established before remediation.
   - Final verification status confirming clean working tree and synced branch.

3. **`03-forensic-log.txt`**:
   - Raw verbatim command output, git diff listings, and exit codes.

#### Offloading Mechanism via `gitmap rc`:
Muse uses native GitMap cache commands to persist forensics without polluting workspace worktrees:
```bash
# Offload directory of forensic artifacts
gitmap rc folder "<local-temp-path>/<task-slug>" --repo muse

# Offload individual diagnostic text
gitmap rc text "<diagnostic-summary>" --slug "muse/troubleshooting/<task-slug>/01-diagnosis" --ext .md
```

---

## 4. Non-Destructive Git Remediation Patterns & Runbooks

### 4.1 Strict Safety Invariants

Every Git troubleshooting procedure executed by Muse must comply with three non-negotiable safety rules:

1. **TOTAL BAN on Force Pushing (`git push --force`, `git push -f`):**
   Force-pushing destroys remote commits and causes immediate data loss in connected platforms (such as Lovable). Under no circumstances may Muse ever force-push.
2. **TOTAL BAN on Rebasing Published Commits:**
   Rebasing rewritten history desynchronizes collaborators and connected branches. Branch convergence must always use merge-based or fast-forward strategies.
3. **Strict Worker Git Ban:**
   Subagents spawned by Muse are strictly prohibited from invoking `git add`, `git commit`, `git push`, `git merge`, or `git checkout`. When a worker detects a Git collision, it sends a diagnostic report back to the Lead Orchestrator, which conducts the fix at root level.

---

### 4.2 Pattern 1: Non-Destructive Branch Divergence Reconciliation

**Scenario:** Local branch and remote tracking branch have diverged (`Your branch and 'origin/main' have diverged`).

#### Step-by-Step Runbook:
1. **Inspect Divergence Depth:**
   ```bash
   git rev-list --left-right --count origin/main...HEAD
   ```
   First number represents remote commits ahead; second number represents local unpushed commits.

2. **Inspect Commit History Graph:**
   ```bash
   git log --oneline --graph -n 10 origin/main...HEAD
   ```

3. **Check for Uncommitted Local Changes:**
   ```bash
   git status --porcelain
   ```
   If changes exist, isolate them before merging:
   ```bash
   git stash push -m "muse-pre-merge-divergence-stash"
   ```

4. **Execute Safe Merge Re-Alignment:**
   - Attempt fast-forward first if local has no unique commits:
     ```bash
     git merge --ff-only origin/main
     ```
   - If both branches have commits, execute non-destructive merge:
     ```bash
     git merge --no-ff -m "Merge remote-tracking branch 'origin/main' into HEAD" origin/main
     ```

5. **Re-Apply Stashed Changes (if applicable):**
   ```bash
   git stash pop
   ```

6. **Verify Clean State & Push:**
   ```bash
   git status --porcelain
   gitmap push
   ```

---

### 4.3 Pattern 2: Mechanical Conflict Resolution by Domain Solvers

**Scenario:** Merge operations encounter conflicted files (`CONFLICT (content): Merge conflict in <file>`).

#### Step-by-Step Runbook:
1. **List Conflicted Files:**
   ```bash
   git diff --name-only --diff-filter=U
   ```

2. **Categorize by Domain Type:**
   - **Type A: Additive Indices, Manifests & JSON Registries** (`readme.md`, `package.json`, `index.ts`, `manifest.json`):
     Resolve by union-joining distinct entries, preserving both incoming remote items and local additions.
   - **Type B: Monotonic Sequence Logs & Changelogs** (`changelog.md`, `version.json`):
     Accept higher sequence numbers or concatenate chronological updates.
   - **Type C: Source Code & Core Logic:**
     Inspect conflict blocks markers (`<<<<<<< HEAD`, `=======`, `>>>>>>>`). Retain non-colliding enhancements while strictly matching architectural specifications.

3. **Automated Solver Invocation:**
   Where applicable, invoke the canonical Python reconciliation solver:
   ```bash
   python 03-ai-scripts/47-git-reconcile-and-resolve-conflict.py
   ```

4. **Verify Marker Removal:**
   Confirm zero unresolved conflict markers remain:
   ```bash
   gitmap aum search "<<<<<<<"
   ```

5. **Stage and Complete Merge:**
   ```bash
   git add <resolved-files>
   git commit -m "Merge: resolve upstream branch conflicts cleanly"
   ```

---

### 4.4 Pattern 3: Git Index Lock Contention Remediation

**Scenario:** Git command aborts with:
`fatal: Unable to create '.git/index.lock': File exists.`

#### Step-by-Step Runbook:
1. **Check for Active Git Processes:**
   Verify whether a background process or file watcher is legitimately executing Git operations:
   ```bash
   # Windows PowerShell:
   Get-Process -Name git -ErrorAction SilentlyContinue

   # Linux / macOS:
   pgrep -l git
   ```

2. **Evaluate Lock Staleness:**
   If no Git processes are running, inspect the creation timestamp of the lock file:
   - On Windows: Check last write time of `.git/index.lock`.
   - On Linux/macOS: Check `ls -l .git/index.lock`.

3. **Surgical Lock Removal (Only When Process Is Inactive):**
   If confirmed that no background process is active:
   ```bash
   # Windows PowerShell:
   Remove-Item -Force .git/index.lock

   # Linux / macOS:
   rm -f .git/index.lock
   ```

4. **Verify Index Integrity:**
   ```bash
   git status
   ```

---

### 4.5 Pattern 4: Workspace Debris & Uncommitted Collision Isolation

**Scenario:** Untracked files, scratch scripts, or temporary test harnesses prevent a clean checkout, branch switch, or pull operation.

#### Step-by-Step Runbook:
1. **Identify Colliding Debris:**
   ```bash
   git status --porcelain
   ```

2. **Archive and Offload to `repo-cache`:**
   Instead of deleting user scripts or loose files (which violates GUARD-02: Never Delete Files Without Asking), offload them into `repo-cache`:
   ```bash
   # Offload temporary file
   gitmap rc file "<debris-file-path>" --repo muse

   # Remove local working copy safely after archive confirmation
   git rm --cached "<debris-file-path>" 2>/dev/null || rm "<debris-file-path>"
   ```

3. **Verify Git Working Tree Is Clean:**
   ```bash
   git status --porcelain
   ```

---

## 5. Master Prompt Updates Specification (`01-muse-master-prompt.md`)

### 5.1 Modifications Required in `01-muse-master-prompt.md`

| Location in File | Current State | Target State |
| :--- | :--- | :--- |
| **Top Parameters (Lines 4–12)** | `A = 2`, `Concurrency Capacity = 4` | `A = 3`, `Concurrency Capacity = 6` |
| **Preamble & Subagent Gate (Lines 61–68)** | `A = 2` across 3 stages | `A = 3` across 3 stages, with Dynamic Respawn and Dynamic Join rules |
| **Actionable Agent Checklist (Lines 55)** | `mandatory multi-agent dispatch (invoke_subagent, A=2, H=2)` | `mandatory multi-agent dispatch (invoke_subagent, A=3, H=2, dynamic respawn)` |
| **Section 5: Phase 6 Step 3 (Line 518)** | `Mandatory Multi-Agent Dispatch (A = 2, H = 2)` | `Mandatory Multi-Agent Dispatch (A = 3, H = 2) & Dynamic Respawn` |
| **Between Phase 7 and Phase 8 (Line 540)** | No troubleshooting section | Insert **Phase 7.5 — Git Repository Troubleshooting & Repo-Cache Forensics** |
| **Phase 9 Master Checklist (Line 569)** | Item V-13 references `A=2, H=2` | Item V-13 references `A=3, H=2`, dynamic respawn, and Phase 7.5 Git troubleshooting |
| **Version Banner (Lines 24 & 574)** | `Version 6.4.0` | `Version 6.5.0` |

### 5.2 Content for New Phase 7.5 in `01-muse-master-prompt.md`

```markdown
---

## 6.5 Phase 7.5 — Git Repository Troubleshooting & Repo-Cache Forensics

> [!IMPORTANT]
> **Conditional Activation Directive:**
> Execute this phase ONLY when a Git conflict, branch divergence, index lock collision, or dirty worktree blockage actively impedes repository operations.
> If all operations succeed cleanly, skip Phase 7.5 entirely. Never write empty logs or placeholder files to `repo-cache`.

### 1. Non-Destructive Invariants (Non-Negotiable)
- **TOTAL BAN on Force-Pushing:** NEVER execute `git push --force` or `git push -f`. Published history must be preserved unconditionally.
- **TOTAL BAN on Published Rebase:** NEVER rebase published commits or branch histories.
- **Lovable History Protection:** Ensure remote tracking histories and auto-synced branches remain linear or cleanly merged without rewriting commit SHAs.
- **Worker Git Ban:** Subagents report Git errors to Lead; Lead alone executes Git diagnosis and remediation.

### 2. Git Remediation Runbooks
1. **Branch Divergence:**
   - Detect: `git rev-list --left-right --count origin/<branch>...HEAD`
   - Reconcile: `git merge --no-ff -m "Merge remote-tracking branch 'origin/<branch>' into HEAD" origin/<branch>`
2. **Merge Conflicts:**
   - Detect: `git diff --name-only --diff-filter=U`
   - Resolve: Apply additive domain solvers or `python 03-ai-scripts/47-git-reconcile-and-resolve-conflict.py`, verify zero remaining `<<<<<<<` markers, stage and commit.
3. **Index Lock Contention:**
   - Detect: `fatal: Unable to create '.git/index.lock': File exists`
   - Inspect: Confirm no active git processes via process list.
   - Cleanse: Remove stale `.git/index.lock` only after confirming no active Git process.
4. **Debris Isolation:**
   - Offload loose scripts or untracked debris into `repo-cache` via `gitmap rc file <path> --repo muse` before pulling.

### 3. Repo-Cache Forensic Persistence
Whenever Git troubleshooting is executed, record the diagnosis and remediation steps into `repo-cache` under `muse/troubleshooting/<task-slug>/`:
- `01-diagnosis.md`: Error message, divergence count, and root cause classification.
- `02-remediation-steps.md`: Exact commands run and verification output.
- `03-forensic-log.txt`: Verbatim terminal outputs.
Persist via `gitmap rc folder` or `gitmap rc text`.
```

---

## 6. Catalog Updates Specification (`01-prompts/27-muse-prompts/readme.md`)

In `01-prompts/27-muse-prompts/readme.md`:
1. Bump prompt version to `6.5.0`.
2. Update Overview items to highlight `A = 3, H = 2` concurrency capacity (`6 operations`), immediate dynamic respawn, Phase 7.5 Git repository troubleshooting, and `repo-cache` forensics.
3. Update Directory Index table entry for `01-muse-master-prompt.md`:
   - Version: `6.5.0`
   - Key Capabilities: Add 3-agent swarm concurrency (`A = 3, H = 2`), dynamic respawn, Phase 7.5 non-destructive Git troubleshooting, and `repo-cache` forensics.

---

## 7. Verification & Acceptance Criteria

- [ ] **AC-01 (Concurrency Header):** `01-prompts/27-muse-prompts/01-muse-master-prompt.md` specifies `A = 3`, `H = 2`, and `System Concurrency Capacity = 6`.
- [ ] **AC-02 (3-Stage Dispatch):** `01-prompts/27-muse-prompts/01-muse-master-prompt.md` details 3-stage dispatch for Planning (3 research subagents), Spec Authoring (3 self subagents), and Execution (3 self worker subagents).
- [ ] **AC-03 (Immediate Dynamic Respawn):** Muse prompt mandates immediate replacement subagent spawning the moment any agent returns via `send_message`.
- [ ] **AC-04 (Phase 7.5 Section):** Phase 7.5 is explicitly documented between Phase 7 and Phase 8 in `01-prompts/27-muse-prompts/01-muse-master-prompt.md`.
- [ ] **AC-05 (Conditional Forensics):** Phase 7.5 states that forensics are written to `repo-cache/muse/troubleshooting/<task-slug>/` only when troubleshooting is actively required.
- [ ] **AC-06 (Non-Destructive Invariants):** Strict prohibition on `git push -f`, rebasing published commits, and explicit Worker Git Ban are documented.
- [ ] **AC-07 (Catalog Sync):** `01-prompts/27-muse-prompts/readme.md` is updated to Version `6.5.0` reflecting all new capabilities.
- [ ] **AC-08 (Path Hygiene):** Zero absolute paths (`C:\...`, `/home/...`) and zero `file:///` URIs exist in specs and documentation.
- [ ] **AC-09 (Lowercase Filenames):** All referenced files and directory names are strictly lowercase.
- [ ] **AC-10 (Boolean Hygiene):** All boolean expressions use positive `is*`/`has*` prefixes without explicit `== true`.
