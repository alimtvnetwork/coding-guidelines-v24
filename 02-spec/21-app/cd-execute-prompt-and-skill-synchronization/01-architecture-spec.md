# Architecture Specification: CD Execute Prompt & Skill Synchronization

## 1. Overview & Strategic Intent
This specification formalizes two critical system enhancements within the Prompt Architect meta-repository and connected fleet repositories:
1. **Map Dispatch & Handler Table Pattern (Prompt 38 & Cross-Language Spec 35):** Codifies a canonical coding guideline, execute prompt, and Antigravity/Cursor skills to identify and refactor repetitive token/string/event dispatch branching (chained `if-else if` or verbose `switch-case` matching keys or tokens to invoke discrete functions) into declarative, high-efficiency map dispatch tables (`Record<string, () => void>`, `map[string]func()`, `dict[str, Callable]`, `HashMap<&str, fn()>`).
2. **Heatmap & RC Heatmap Integration & Root Cause Remediation:** Conducts a grounded root cause analysis into why AI agents executing parent tasks fail to execute `heatmap` and `rc heatmap` commands, resolves the command mismatch (`gitmap heatmap discuss` -> `gitmap agent heatmap` / `gitmap heatmap status`), injects mandatory Step 0 Preflight Heatmap & RC commands (`gitmap rc heatmap tables`, `gitmap heatmap status`), updates all 38 CG Execute prompts and companion skills to strict V6 compliance, and synchronizes across all 43 connected repositories.

---

## 2. Root Cause Analysis (RCA): Why AI Does Not Execute Heatmap & RC Commands

### Symptom
During execution of parent tasks under `execute-parent-task-with-n-steps-v6` and CG execute workflows, AI agents do not invoke `gitmap heatmap` or `gitmap rc heatmap` commands.

### Root Cause 1: Invalid Subcommand Reference in Prompts
- The V6 template and CG execute prompts specified:
  `gitmap heatmap discuss [--window 15m] [--threshold 2] [--json]`
- When invoked in the GitMap CLI, `gitmap heatmap discuss` returns:
  `gitmap: [E1000:VALIDATION] heatmap: unknown heatmap command: discuss`
- Agents encountering this failure or scanning tool help abandon or avoid the command.
- The actual CLI command implemented in GitMap for pre-edit contention analysis is `gitmap agent heatmap [--json]` and `gitmap heatmap status`.

### Root Cause 2: Missing Step 0 Preflight & Planning Directives
- In `02-execute-parent-task-with-n-steps-v6.md` Section 5 (Step 0: Preflight), the explicit preflight checklist only mandated:
  - `gitmap agent task init`
  - `gitmap agent task status`
- It completely lacked explicit instructions for:
  - `gitmap rc heatmap tables` (or `gitmap heatmap rc tables`)
  - `gitmap heatmap status`
  - `gitmap agent heatmap`
- Because LLM execution strictly follows sequential numbered preflight steps, unlisted commands are never executed during preflight.

### Root Cause 3: Discovery Agent Context Blindness
- Section 7.1.B (Research Discovery Brief) did not instruct `research` subagents to log reads or check contention.
- It only cited `gitmap aum search`, `gitmap find`, `gitmap lf`, and `gitmap cat`.

### Root Cause 4: Passive Documentation vs. Mandatory Execution Gates
- Heatmap commands were buried as informative references in the worker subtask brief (Section 7.2) rather than being enforced as mandatory preflight gates for the lead agent.

### Root Cause 5: Repository Synchronization Divergence
- Target repositories across `d:\work` lacked the latest V6 updates to CG execute prompts and skills, leaving outdated prompts active in satellite repositories.

---

## 3. Map Dispatch & Handler Lookup Architectural Pattern

### The Anti-Pattern (Chained Branching Dispatch)
```typescript
// BAD: Verbose, repetitive branching dispatching individual functions
const handleKeyDown = (e: KeyboardEvent) => {
  if (e.key === 'Escape') {
    e.preventDefault();
    onClose();
  } else if (e.key === '+' || e.key === '=') {
    e.preventDefault();
    handleZoomIn();
  } else if (e.key === '-') {
    e.preventDefault();
    handleZoomOut();
  } else if (e.key === '0') {
    e.preventDefault();
    handleResetZoom();
  }
};
```

### The Solution (Declarative Map Dispatch Table)
```typescript
// GOOD: Clean, extensible map lookup table with 1-line execution
const KEY_ACTIONS: Record<string, () => void> = {
  Escape: onClose,
  '+': handleZoomIn,
  '=': handleZoomIn,
  '-': handleZoomOut,
  '0': handleResetZoom,
};

const handleKeyDown = (e: KeyboardEvent) => {
  const action = KEY_ACTIONS[e.key];
  if (action) {
    e.preventDefault();
    action();
  }
};
```

### Cross-Language Equivalence
- **Go:** `var handlers = map[string]func(*Context) *appfault.AppError`
- **Python:** `ACTION_HANDLERS: dict[str, Callable[[], None]] = {"key": handler}`
- **Rust:** `let handlers: HashMap<&str, fn(&mut Context)> = ...`
- **PHP:** `$handlers = ['action' => fn() => $this->handleAction()];`

---

## 4. Architectural Boundaries & Non-Negotiables
1. **Spec 21 Boundary:** All application planning lives in `02-spec/21-app/` and is strictly excluded from fleet synchronization.
2. **Shared Guidelines:** The new cross-language guideline is authored in `02-spec/02-coding-guidelines/01-cross-language/35-map-dispatch-and-handler-lookup.md`.
3. **Canonical Prompts:** Authored in `01-prompts/15-cg-execute/` and mirrored in `01-prompts/v4/15-cg-execute/`.
4. **Skills Parity:** Every prompt has a corresponding skill in `.agents/skills/` and `.cursor/skills/`.
5. **Strict Relative Paths:** All documentation, specs, and links use relative paths from git root.
