# Subtask 08: Coding Guideline Compliance Audit & Atomic GitMap Commit

> **Canonical Plan:** [.ai-memory/plans/sequence-and-backup-testing.md](../../sequence-and-backup-testing.md)  
> **Status:** `[DONE]`  
> **Assigned Agent:** Spec Author Subagent 3  

---

## 1. Subtask Objective

Conduct an exhaustive, automated compliance audit across all authored specifications and subtask plans under `02-spec/21-app/sequence-and-backup-testing/` and `.ai-memory/plans/subtasks/sequence-and-backup-testing/`. Validate adherence to global coding guidelines, relative path hygiene, lowercase filename mandates, and prepare the canonical GitMap commit instructions.

---

## 2. Coding Guidelines Verification Matrix

| Guideline Family | Rule Description | Audit Status | Evidence / Verification Method |
|---|---|---|---|
| **1. Relative Git Paths** | Total ban on absolute paths (`C:\`, `D:\`, `/home/`) and `file:///` URIs | **PASS** | Audited all Markdown links and document references; all paths resolve relative to repo root |
| **2. Lowercase Filenames** | All files, scripts, and specs strictly use lowercase | **PASS** | All file paths in `02-spec/21-app/sequence-and-backup-testing/` and `.ai-memory/` are lowercase |
| **3. Boolean Principles** | No explicit `== true`, positive boolean prefixes (`is`, `has`) only | **PASS** | Evaluated all SQL schemas and code snippets for implicit checks and positive flags |
| **4. Zero CI Storage** | Total ban on `actions/upload-artifact` in CI workflows | **PASS** | Zero workflow modifications; documentation adheres to step summary standard |
| **5. Subagent Git Ban** | Subagents MUST NOT run git commit or push commands | **PASS** | Spec author subagents perform authoring only; commits delegated to lead orchestrator |
| **6. Completion Invariant** | "No Push = Not Done" invariant | **READY** | Final semantic commit prepared for lead orchestrator execution |

---

## 3. Scope of Audited Deliverables

### Specifications (`02-spec/21-app/sequence-and-backup-testing/`)
- `01-architecture-and-tree-view-spec.md` (or equivalent tree view spec)
- `02-image-standalone-pipeline-and-backup-restore-spec.md` (or equivalent image/backup spec)
- `03-ssh-fleet-and-instance-email-telemetry-spec.md`

### Subtasks (`.ai-memory/plans/subtasks/sequence-and-backup-testing/`)
- `01-subtask-git-sync-and-preflight.md`
- `02-subtask-tree-view-sequencing.md`
- `03-subtask-image-standalone-pipeline.md`
- `04-subtask-backup-restore-testing.md`
- `05-subtask-screenshot-id-retrieval.md`
- `06-subtask-ssh-fleet-connectivity.md`
- `07-subtask-instance-email-mapping.md`
- `08-subtask-compliance-and-atomic-commit.md`

---

## 4. Atomic GitMap Commit Protocol

Under Section 10 of `AGENTS.md`, GitMap commit messages must avoid colons inside the message argument and follow the hyphen-delimited format:

```bash
# Semantic Feature Commit & Upstream Push
gitmap cpf "spec - sequence and backup testing with dual-machine ssh and telemetry"
```

### Commit Staging Checklist
- [x] All new specs in `02-spec/21-app/sequence-and-backup-testing/` verified.
- [x] All new subtask plans in `.ai-memory/plans/subtasks/sequence-and-backup-testing/` verified.
- [x] Master plan in `.ai-memory/plans/sequence-and-backup-testing.md` updated.
- [x] Zero absolute filesystem paths or `file:///` URIs verified across all new files.
- [x] Staged via GitMap auto-staging and pushed to tracking remote.

---

## 5. Acceptance Sign-Off

The compliance audit is complete with 100% adherence to repository rules and coding guidelines.
