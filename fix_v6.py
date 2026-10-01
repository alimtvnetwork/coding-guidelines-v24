import re

file_path = "01-prompts/14-execute/13-execute-parent-task-with-n-steps-v6.md"

with open(file_path, "r", encoding="utf-8") as f:
    content = f.read()

# Replace the MANDATORY SUBAGENT SPAWNING GATE block
new_gate = """### 🚨 MANDATORY SUBAGENT SPAWNING GATE (A = 2, H = 2 — ZERO SOLO EXECUTION ALLOWED)

- **ACTUAL TOOL CALL REQUIRED:** You must ACTUALLY CALL the `invoke_subagent` tool via your tool-calling API. Do NOT just print the text "Dispatched Worker..." and stop. If you only print text, the agents will not spawn and the task will fail! You must execute the `invoke_subagent` JSON tool payload.
- **3-STAGE MANDATORY DISPATCH:** You must invoke `A = 2` agents (`invoke_subagent`) at EVERY stage of the workflow:
  1. **Planning Step:** Spawn 2 subagents to research the codebase and author the step-by-step execution plan.
  2. **Spec Step:** Spawn 2 subagents to write the detailed architectural spec.
  3. **Execution Step:** Spawn 2 worker subagents to execute the actual code modifications.
- **SOLO EXECUTION IS AN AUTO-REJECT FAILURE:** The lead orchestrator is **STRICTLY FORBIDDEN** from executing planning, spec writing, or code changes by itself without calling the `invoke_subagent` tool. Failing to call the actual tool is a critical protocol violation.

---"""

content = re.sub(
    r'### \U0001F6A8 MANDATORY SUBAGENT SPAWNING GATE.*?---\n',
    new_gate + "\n",
    content,
    flags=re.DOTALL
)

# Update Phase 1B to be Planning Step
phase_1_new = """## 6. Phase 1: Planning Step & Spec Step (Steps 1 .. PHASE_1_BUDGET)

You must use `invoke_subagent` to delegate both planning and spec writing.

1. **Planning Step (A = 2 Agents):** Lead agent calls `invoke_subagent` to spawn 2 subagents (e.g., `Research 01`, `Research 02`). Their prompt MUST instruct them to research the codebase, define the boundaries, and write the Execution Plan (`.ai-memory/plans/pending/NN-<slug>.md`) and the Root Task JSON Manifest.
   - *Tool Call:* Lead must actually execute the `invoke_subagent` tool, then print `Dispatched Planning Agents`, and then END TURN to wait for `<SYSTEM_MESSAGE>`.
2. **Spec Step (A = 2 Agents):** Once planning is done, the lead agent calls `invoke_subagent` again to spawn 2 subagents. Their prompt MUST instruct them to write the Canonical Spec (`02-spec/21-app/NN-<slug>.md` or folder) and decompose subtasks into `.ai-memory/plans/subtasks/`.
   - *Tool Call:* Lead must actually execute the `invoke_subagent` tool, then print `Dispatched Spec Agents`, and then END TURN to wait for `<SYSTEM_MESSAGE>`.
3. **Readiness Gate:** Complete Phase 1 planning and spec authoring within `PHASE_1_BUDGET` steps, then proceed **UNCONDITIONALLY** into Phase 2."""

content = re.sub(
    r'## 6\. Phase 1B: Spec, Plan & Lean Subtasks.*?---',
    phase_1_new + "\n\n---",
    content,
    flags=re.DOTALL
)

# Update Phase 2 instructions to emphasize ACTUAL tool call
phase_2_new = """## 7. Phase 2: Execution Step (Worker Waves) (Steps (PHASE_1_BUDGET + 1) .. N)

> [!CRITICAL]
> **MANDATORY `invoke_subagent` DISPATCH (ZERO SOLO EXECUTION):**
> You MUST ACTUALLY CALL the `invoke_subagent` tool to spawn A workers (`TypeName: "self"`, up to H subtasks per worker) in parallel. Executing all subtasks solo in main agent without the tool call is an immediate auto-reject failure.

### 7.1 Dispatch Payload (`invoke_subagent`)

The `invoke_subagent` payload holds A entries (`Worker 01 .. Worker <A>`):

```json
{
  "Subagents": [
    {
      "TypeName": "self",
      "Role": "Worker 01: [Assigned Subtask]",
      "Model": "inherit",
      "Workspace": "inherit",
      "Prompt": "<Worker Brief Below>"
    },
    {
      "TypeName": "self",
      "Role": "Worker 02: [Assigned Subtask]",
      "Model": "inherit",
      "Workspace": "inherit",
      "Prompt": "<Worker Brief Below>"
    }
  ]
}
```"""

content = re.sub(
    r'## 7\. Phase 2: Mandatory Worker Waves & Coding Guidelines Enforcement.*?(?=### 7\.2)',
    phase_2_new + "\n\n",
    content,
    flags=re.DOTALL
)

turn_yielding = """### 7.3 Turn-Yielding & Verification Protocol

1. **Invoke & Yield:** You must ACTUALLY CALL the `invoke_subagent` tool in your turn. Print the progress line (`Dispatched Worker 01 .. Worker <A> (wave k / WAVES); waiting for their results.`) and **STOP CALLING TOOLS** to end your turn.
2. **Verify Worker Reports Independently:** Confirm `git diff --stat -- <owned files>` matches `filesChanged`, no files outside owned files modified, re-run targeted checks for `exit 0` on non-zero files.
3. **Reject Violations:** Send failures via `send_message`. On `BLOCKED`, lead does work and logs `LEAD_FALLBACK: <reason>`. After two failed rounds, mark `FAILED`, write RCA, continue (R13).
4. **Update Ledger:** Record status, evidence, changed paths in `ledger.md` via `replace_file_content`.
5. **Loop:** Dispatch subsequent waves via `invoke_subagent` until all subtasks are `DONE` or `FAILED`."""

content = re.sub(
    r'### 7\.3 Turn-Yielding & Verification Protocol.*?---',
    turn_yielding + "\n\n---",
    content,
    flags=re.DOTALL
)

with open(file_path, "w", encoding="utf-8") as f:
    f.write(content)
print("Updated V6 prompt successfully.")
