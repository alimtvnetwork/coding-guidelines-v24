# Subtask 03: Heatmap Shorthand Commands & Spec Tracking

- **Parent Plan:** [.ai-memory/plans/pending/32-coding-guideline-corrections-and-gitmap-skill-updates.md](.ai-memory/plans/pending/32-coding-guideline-corrections-and-gitmap-skill-updates.md)
- **Target Repo:** `d:/work/gitmap`
- **Target Files:**
  - `cli/heatmap/heatmap_cmd.go`
  - `cli/heatmap/heatmap_rc_cmd.go`
  - `cli/store/heatmap_store.go`
  - `cli/store/heatmap_rc_store.go`
  - `cli/cmdagent/agent_task_sqlite.go`
- **Status:** Completed

---

## Action Items

1. In `cli/heatmap/heatmap_cmd.go` and `cli/heatmap/heatmap_rc_cmd.go`, wire aliases for shorthand commands:
   - `failed`, `unmet`, `broken` -> `RunHeatmapFailures`
   - `hotfiles`, `hot`, `hotspots` -> `RunHeatmapHotfiles`
   - `recent`, `latest` -> `RunHeatmapRecent`
   - `tasks`, `task-list` -> `RunHeatmapTasks`
2. In `cli/store/heatmap_store.go`, implement `RecordParentTask` with SQLite upsert to write into `HeatmapParentTask` in local `heatmap.db`.
3. In `cli/cmdagent/agent_task_sqlite.go`, wire `gitmap task init` to invoke `RecordParentTask` so `SpecFile` and `SpecFolder` are tracked in both `agent-task.db` and `heatmap.db`.
4. Run unit tests in `cli/heatmap` and `cli/cmdagent` to ensure 100% pass.
