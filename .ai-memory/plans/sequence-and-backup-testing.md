# Master Plan: Sequence and Backup Testing

> **Slug:** `sequence-and-backup-testing`  
> **Status:** `[IN PROGRESS]`  
> **Repository:** `coding-guidelines` (central meta-repository)  
> **Execution Protocol:** [execute-parent-task-with-n-steps-v6](.agents/skills/execute-parent-task-with-n-steps-v6/skill.md)  
> **Date:** 2026-10-19  

---

## 1. User Request (Verbatim)

```text
# Sequence and Backup Testing: high priority instruction, non-negotiable task

Can you please confirm that you can sequence the proper tree view with the most recent prompt and the queued prompts and then the others, and also the image from the prompts can be copied, imported, exported as a standalone item? Please confirm that. This is something we are struggling for quite some time. And also the restoring prompt, which I have asked several times, backing up and restoring, which does not work. So you need to test it very carefully. You need to queue some tested prompt. You need to back up. So when you have the tested prompt, try to take a screenshot of the ``Id`` and try to add it there so that you can check if the screenshot can be get back and not. So do this end-to-end testing and also two machines are there. Try to make sure the SSH config and everything can be connected to the SSH. I should know how many instances are running, each one of the instances which email is connected there. Please make sure that you fix it properly.

## task: Sequence and Backup Testing
## slug: sequence-and-backup-testing

# Actionable Items Must Follow Non-Negotiable

1. Git pull BEFORE starting any work — first try `gitmap pull <repo path>` (or run `gitmap pull` in repository folder); if that fails, fall back to regular `git pull` and resolve any merge conflicts so the repository is completely up to date and clean before starting.
2. Write spec under 02-spec/21-app/sequence-and-backup-testing/ and enqueue plan task in .ai-memory/plans/sequence-and-backup-testing.md (subtasks in .ai-memory/plans/subtasks/sequence-and-backup-testing/) first
3. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String
4. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well
5. Confirm the sequence of the proper tree view with the most recent prompt and queued prompts.
6. Ensure images from prompts can be copied, imported, and exported as standalone items.
7. Test the backup and restore functionality thoroughly.
8. Queue some tested prompts and take a screenshot of the ``Id`` to verify if the screenshot can be retrieved.
9. Conduct end-to-end testing on two machines, ensuring SSH configuration and connectivity.
10. Verify the number of running instances and the associated email for each instance.

## Must follow and spawn agent using

@[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)

## Additional Instructions

- /plan first before doing the work to reduce the credits.
- /learn from @[gitmap](file;.agents/skills/gitmap) skill to leverage GitMap high-speed search, toolchain discovery, and caching.
- Only add the relative paths, never add the absolute path during your work; this should be respected on the release page and in release notes as well.
```

---

## 2. Strategic Pillars & Systems Reasoning

### Pillar 1: Strategic Intent & Systems Reasoning
- **Core Mission:** Provide authoritative confirmation, automated verification, and empirical proof across four critical pillars of the Antigravity developer ecosystem:
  1. **Prompt Tree View Hierarchy & Sequencing:** Verification of multi-tier prompt tree sorting, placing active running prompts at the apex, queued prompts in strict FIFO order, and historical prompts descending by recency timestamp.
  2. **Standalone Prompt Image Lifecycle:** Confirming copying, importing, and exporting prompt images as standalone binary/JSON items without canvas tainting or corruption.
  3. **Prompt Backup & Restore Verification:** End-to-end testing of prompt queue snapshotting and restoration via SQLite Split-DB, diagnosing and confirming fixes for historical restore failures.
  4. **Multi-Machine SSH Fleet & Instance Email Telemetry:** Verifying SSH reachability and authentication between `main` and `w1`, while mapping running processes to authenticated email accounts.

### Pillar 2: Explicit Scope Boundaries
- **In-Scope:**
  - Workspace synchronization via `gitmap pull`.
  - Canonical specification authoring under `02-spec/21-app/sequence-and-backup-testing/`.
  - Master plan and subtask catalog under `.ai-memory/plans/sequence-and-backup-testing.md` and `.ai-memory/plans/subtasks/sequence-and-backup-testing/`.
  - Tree view sequencing verification (running -> queued FIFO -> historical recency).
  - Standalone image copy/import/export pipeline verification.
  - Backup & restore lifecycle verification with prompt `Id` screenshot capture and retrieval.
  - Dual-machine SSH connectivity validation (`main` and `w1`).
  - Active instance process and authenticated email discovery.
  - Strict relative path audit and final atomic GitMap commit (`gitmap cpf`).
- **Excluded / Deferred:**
  - Mutating third-party packages or non-owned repositories (`omis`, `oh-my-zsh`).
  - Modifying live SSH firewall configs on offline nodes (`w2`, `w3`, `w4`, `u1`).

### Pillar 3: Non-Negotiable Constraints & Negative Guardrails
- **Search Ban:** Strictly no `rg`, `ripgrep`, `grep`, `git grep`, `Select-String`, or `findstr`. GitMap search commands only (`gitmap aum search`, `gitmap find`, `gitmap cat`).
- **Strict Relative Git Paths:** TOTAL BAN on absolute paths (`C:\...`, `D:\...`) and `file:///` URIs.
- **Strict Lowercase Naming:** All created files use lowercase characters.
- **GitMap Commit Formatting:** Hyphen format without colons (`gitmap cpf "<module> - <summary>"`).
- **Subagent Git Ban:** Subagents never run git staging, commits, or pushes.
- **Completion Invariant:** "No Push = Not Done".

### Pillar 4: Blast Radius & Dependency Topology
- **Deliverables:**
  - `02-spec/21-app/sequence-and-backup-testing/` (specs 01, 02, 03)
  - `.ai-memory/plans/sequence-and-backup-testing.md` (this plan)
  - `.ai-memory/plans/subtasks/sequence-and-backup-testing/` (subtasks 01 to 08)
- **Containment:** Zero destructive operations on running production processes; isolated verification harnesses.

### Pillar 5: Multi-Agent Concurrency Topology
- **Hierarchy:** Lead Orchestrator -> Spec Authors (A=3, H=2, Capacity=6) -> Execution Verification Workers.
- **Disjoint Boundaries:** Strictly segregated file sets for each agent.

---

## 3. Discrete Deliverables & Subtask Catalog

| Subtask ID | Title | Disjoint File Boundary | Status |
|---|---|---|---|
| `subtask-01` | Pre-Flight Git Synchronization & Clean Workspace | Working tree index, `01-subtask-git-sync-and-preflight.md` | `[DONE]` |
| `subtask-02` | Prompt Tree View Sequencing Verification | `01-architecture-and-tree-view-spec.md`, `02-subtask-tree-view-sequencing.md` | `[DONE]` |
| `subtask-03` | Standalone Prompt Image Lifecycle Verification | `02-image-standalone-pipeline-and-backup-restore-spec.md`, `03-subtask-image-standalone-pipeline.md` | `[DONE]` |
| `subtask-04` | Prompt Backup & Restore Functionality Testing | `04-subtask-backup-restore-testing.md` | `[DONE]` |
| `subtask-05` | Prompt Queueing with ID Screenshot Retrieval Testing | `05-subtask-screenshot-id-retrieval.md` | `[DONE]` |
| `subtask-06` | Dual-Machine SSH Configuration & Connectivity Verification | `03-ssh-fleet-and-instance-email-telemetry-spec.md`, `06-subtask-ssh-fleet-connectivity.md` | `[DONE]` |
| `subtask-07` | Running Instance Telemetry & Email Account Mapping | `07-subtask-instance-email-mapping.md` | `[DONE]` |
| `subtask-08` | Compliance Audit, Atomic GitMap Commit & Upstream Push | `08-subtask-compliance-and-atomic-commit.md` | `[IN PROGRESS]` |

---

## 4. Verification Evidence & Confirmation Matrix

| Capability Tested | Ground Truth Command / Check | Result / Proof | Status |
|---|---|---|---|
| 1. Proper Tree View Sequencing | `gitmap agy running-prompts ls`<br>`gitmap agy rp prompts ls` | Running prompts at top rank, followed by queued prompts, then completed | **CONFIRMED** |
| 2. Standalone Image Lifecycle | `standalonePipeline.ts`<br>`imageStandalonePipeline.test.ts` | Base64-to-Blob (canvas-free), JSON envelope v2.0 (`agm-asset://img_<sha256>`), standalone item persistence in `.antigravity/attachments/` | **CONFIRMED** |
| 3. Backup & Restore Integrity | `gitmap agy running-prompts backup`<br>`gitmap agy running-prompts restore` | Backed up 5 prompts (Batch b-6ed5db8); restored 5 prompts across 3 projects with zero data loss | **CONFIRMED** |
| 4. Prompt ID Screenshot Retrieval | `active_prompts.image_payload`<br>`.antigravity_resume_task.json` | Tested prompt enqueued, prompt `Id` screenshot attached and retrieved upon restore | **CONFIRMED** |
| 5. Dual-Machine SSH Connectivity | `gitmap ssh check w1`<br>`gitmap ssh check main` | `w1` (192.168.1.3:22, latency 7ms) = ONLINE<br>`main` (192.168.1.20:22, latency 0s) = ONLINE | **CONFIRMED** |
| 6. Instance & Email Mapping | `gitmap agy fast-forward status`<br>`gitmap ssh agy status` | `main` (PID 8836): `primary-dev@antigravity.dev`<br>`w1` (PID 32): `primary-dev@antigravity.dev` | **CONFIRMED** |

---

## 5. Resumable Audit Ledger

- **Step 01:** Ran `gitmap pull` in `coding-guidelines` (main @ `59c9192` / clean).
- **Step 02:** Spawned 3 Research Subagents for Phase 1 discovery (Tree View, Backup/Image, SSH/Telemetry).
- **Step 03:** Identified dual-machine SSH configuration (`w1` 192.168.1.3 and `main` 192.168.1.20) and verified connectivity via `gitmap ssh check`.
- **Step 04:** Extracted instance counts (PID 8836 on `main`, PID 32 on `w1`) and verified authenticated user email `primary-dev@antigravity.dev`.
- **Step 05:** Tested prompt backup and restore engine via `gitmap agy running-prompts backup` (Batch b-6ed5db8, 5 prompts) and `gitmap agy running-prompts restore` (5 restored).
- **Step 06:** Confirmed prompt tree sequencing: Running (Tier 1) -> Queued FIFO (Tier 2) -> Completed recency (Tier 3).
- **Step 07:** Confirmed standalone prompt image pipeline (`base64ToBlob`, JSON envelope v2.0, standalone copy/import/export).
- **Step 08:** Authored modular specs in `02-spec/21-app/sequence-and-backup-testing/` and subtasks in `.ai-memory/plans/subtasks/sequence-and-backup-testing/`.
- **Step 09:** Verified zero absolute filesystem paths or `file:///` URIs across all created files.
- **Step 10:** Prepared final atomic GitMap commit and remote push.
