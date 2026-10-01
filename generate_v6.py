import re

with open("01-prompts/14-execute/13-execute-parent-task-with-n-steps-v6.md", "r", encoding="utf-8") as f:
    v6_content = f.read()

precedence_match = re.search(r'(## 1\. Precedence Hierarchy & Scope.*?)(?=---)', v6_content, re.DOTALL)
core_rules_match = re.search(r'(## 2\. Core Operational Rules.*?)(?=---)', v6_content, re.DOTALL)
gitmap_match = re.search(r'(## 3\. GitMap High-Speed Command Primacy.*?)(?=---)', v6_content, re.DOTALL)
preflight_match = re.search(r'(## 5\. Step 0: Preflight, Platform Handshake & Ledger Creation.*?)(?=---)', v6_content, re.DOTALL)
phase1b_match = re.search(r'(## 6\. Phase 1B: Spec, Plan & Lean Subtasks.*?)(?=---)', v6_content, re.DOTALL)
phase2_match = re.search(r'(## 7\. Phase 2: Mandatory Worker Waves & Coding Guidelines Enforcement.*?)(?=---)', v6_content, re.DOTALL)
phase3_match = re.search(r'(## 8\. Phase 3: Consolidation, Evidence Verification & Atomic GitMap Push.*?)(?=---)', v6_content, re.DOTALL)
report_match = re.search(r'(## 9\. Final Report Format.*?)(?=---)', v6_content, re.DOTALL)
verification_match = re.search(r'(## 10\. Targeted Verification Checks.*?)(?=---)', v6_content, re.DOTALL)
discovery_match = re.search(r'(## 11\. Discovery Toolchain.*?)(?=---)', v6_content, re.DOTALL)
rca_match = re.search(r'(## 12\. Issue Destination & RCA Routing.*?)(?=---)', v6_content, re.DOTALL)
must_follow_match = re.search(r'(## MUST FOLLOW NON-NEGOTIABLE.*)', v6_content, re.DOTALL)

new_content = """# [V6] Parent Task N-Step Continuous Loop & Mandatory Multi-Agent Subagent Orchestration — Workflow (must follow)

```text
N = 300 (Total self-loop steps budget — editable top-header parameter, default: 300)
A = 2   (MANDATORY number of spawned autonomous subagents running concurrently via invoke_subagent, default: 2)
H = 2   (Operational hands per agent: dual-task batch capacity & parallel tool dispatch, default: 2)
C = 30  (Tool calls per worker before it must report, default: 30)

System Concurrency Capacity = A × H = 2 agents × 2 hands = 4 concurrent subtask operations
PHASE_1_BUDGET = N / 2   (Steps 1 .. 150: Planning, Parallel Discovery Subagents, Detailed Spec, and Lean Subtask Generation)
PHASE_2_BUDGET = N / 2   (Steps 151 .. 300: Mandatory Parallel Subagent Execution, Self-Looping, Targeted Quality Linting)
WAVES = ceil(subtasks / (A x H))
```

> [!IMPORTANT]
> Prompt Version: 6.0.0
> Runtime: Google Antigravity 2.0 (IDE and CLI)
> Invoke: /execute-parent-task-with-n-steps-v6 <task>
>
> **Top-Instruction Priority Mandate (Above Precedence / Preamble Precedence):**
> Whatever directives, constraints, checklists, or user instructions are given ABOVE this prompt (including in the user preamble, header blocks, or incoming user request above) are HIGHEST PRIORITY and MUST BE FOLLOWED as strictly NON-NEGOTIABLE. They supersede and strictly override any conflicting general advice, default conventions, or lower-level guidelines below. The agent MUST inspect and follow the instructions above with absolute precedence.

[/goal](slashCommand:goal) Autonomously orchestrate and execute the parent task end-to-end: FIRST showcase and list out the given task in visible chat during Turn 1, capture it verbatim, plan it in the repo, spawn autonomous subagents via `invoke_subagent` (A = 2, H = 2; solo execution without calling `invoke_subagent` is an auto-reject failure) in disjoint file boxes using GitMap high-speed commands as primary, prove every single claim with concrete evidence, enforce coding guidelines to 100%, and finish with one atomic GitMap commit that holds strictly this task's files.

[/learn](slashCommand:learn) Enforce the Top-Instruction Priority Mandate: whatever directives, custom requirements, checklists, or user instructions are provided ABOVE this prompt outrank everything below. Turn 1 MUST showcase the given task list in visible chat before any background execution. Each rule is stated once (R1 to R16) and cited by ID. Progress lives in the ledger and in `.ai-memory/plans/`, never only in chat.

### 🚨 MANDATORY SUBAGENT SPAWNING GATE (A = 2, H = 2 — ZERO SOLO EXECUTION ALLOWED)

- **ABSOLUTE, NON-NEGOTIABLE MUST:** Spawning subagents via the `invoke_subagent` tool (`A = 2`, `H = 2`) is an **ABSOLUTE, NON-NEGOTIABLE MUST** in both **Phase 1** (parallel codebase discovery reading and modular spec authoring) and **Phase 2** (parallel subtask code execution with `TypeName: "self"`).
- **SOLO EXECUTION IS AN AUTO-REJECT FAILURE:** The lead orchestrator is **STRICTLY FORBIDDEN** from executing all discovery reads or all subtask code modifications by itself without invoking `invoke_subagent`. Failing to call `invoke_subagent` when `A >= 2` is a critical protocol violation on the same tier as Rule 0.
- **Phase 1 Mandatory Subagent Dispatch:** Immediately after establishing the Confirmed Task Breakdown (Phase 1A) and the single-agent unified blueprint overview (`01-overview.md` or parent plan skeleton), the lead agent MUST call `invoke_subagent` to spawn `A = 2` subagents in parallel for codebase discovery/reading or modular spec sections and yield the turn to await `<SYSTEM_MESSAGE>`.
- **Phase 2 Mandatory Subagent Dispatch (`TypeName: "self"`):** Once subtasks are generated in `.ai-memory/plans/subtasks/xx-<slug>/`, the lead agent MUST call `invoke_subagent` with `TypeName: "self"` to dispatch `A = 2` worker subagents (`H = 2` disjoint subtasks per worker) and yield the turn to await `<SYSTEM_MESSAGE>`.

---

## The Unified Master Pipeline (Atomic Numbered Steps)

Execute this task via a strict 3-Phase pipeline. Do not skip steps.

### Phase 1A: Verbatim Capture, Task Extraction & Chat Output Gate (Step 0)

Before executing any file searches, scans, spec writing, or code changes, you must execute Phase 1A:

1. **Top-Instruction Priority Verification:** Whatever directives, constraints, checklists, or instructions are given before this section or prompt (user preamble, header constraints, prior instructions) must be verified as highest priority and non-negotiable.
2. **Showcase Given Task First (Turn 1 Action):** In your VERY FIRST response turn upon receiving the prompt, you MUST output the confirmed task breakdown directly in visible chat. Never execute tools silently without displaying the task breakdown to the user first!
3. **Lossless Verbatim Capture:** Store incoming prompt losslessly under `## User Request (Verbatim)` in canonical spec and parent plan.
4. **Screenshots & Media:** Decode base64/screenshots immediately into `assets/screenshots/<slug>-<NN>.png`. Reference via relative markdown links (`![Screenshot](assets/screenshots/<slug>-<NN>.png)`).
5. **Discrete Deliverables Extraction:** Break down whatever user requirements were given into discrete, actionable items with ordered traceable IDs (`Task-01`, `Task-02`, ...).
6. **Mandatory Same-Turn Tool Chaining (TOTAL BAN ON TURNING OFF):** Emit the breakdown in chat with clean vertical formatting, and in the **EXACT SAME TURN**, invoke your first tool call (e.g. `write_to_file` to initialize ledger/spec, or run preflight). NEVER emit text alone (which ends the turn prematurely), and never ask "Should I proceed?".

```markdown
### 📋 Confirmed Task Breakdown & Requirement Ingestion

1. **Task-01: [Descriptive Task Title]**
   - **State:** `[IN PROGRESS — EXECUTING IMMEDIATELY]`
   - **Understood:** `[YES]` — [1-2 concise sentences proving understanding of intent, scope, and verified constraints]
   - **Actionable Scope:** [Precise technical deliverable and implementation scope]
   - **Target Files / Area:** `[relative/path/or/module]`

2. **Task-02: [Descriptive Task Title]**
   - **State:** `[QUEUED — EXECUTING NOW WITHOUT USER PROMPT]`
   - **Understood:** `[YES]` — [1-2 concise sentences proving understanding of intent, scope, and verified constraints]
   - **Actionable Scope:** [Precise technical deliverable and implementation scope]
   - **Target Files / Area:** `[relative/path/or/module]`

Proceeding directly to Preflight & Phase 1B Spec Generation (Active Tool Call Running Below).
```

---

"""

new_content += precedence_match.group(1).strip() + "\n\n---\n\n"
new_content += core_rules_match.group(1).strip() + "\n\n---\n\n"
new_content += gitmap_match.group(1).strip() + "\n\n---\n\n"
new_content += preflight_match.group(1).strip() + "\n\n---\n\n"
new_content += phase1b_match.group(1).strip() + "\n\n---\n\n"
new_content += phase2_match.group(1).strip() + "\n\n---\n\n"
new_content += phase3_match.group(1).strip() + "\n\n---\n\n"
new_content += report_match.group(1).strip() + "\n\n---\n\n"
new_content += verification_match.group(1).strip() + "\n\n---\n\n"
new_content += discovery_match.group(1).strip() + "\n\n---\n\n"
new_content += rca_match.group(1).strip() + "\n\n---\n\n"
new_content += must_follow_match.group(1).strip() + "\n"

with open("01-prompts/14-execute/13-execute-parent-task-with-n-steps-v6.md", "w", encoding="utf-8") as f:
    f.write(new_content)
