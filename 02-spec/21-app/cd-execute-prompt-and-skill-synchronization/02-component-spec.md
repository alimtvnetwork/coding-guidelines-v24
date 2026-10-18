# Component Specification: CD Execute Prompt & Skill Synchronization

## 1. Deliverables Inventory

| Component File | Role & Target |
| :--- | :--- |
| `02-spec/02-coding-guidelines/01-cross-language/35-map-dispatch-and-handler-lookup.md` | Core coding guideline defining the anti-pattern, transformation rules, type safety, and multi-language examples |
| `01-prompts/15-cg-execute/38-map-dispatch-and-handler-lookup.md` | Canonical V6 execution prompt for Map Dispatch refactoring |
| `01-prompts/v4/15-cg-execute/38-map-dispatch-and-handler-lookup.md` | V4 mirror of prompt 38 |
| `.agents/skills/cg-map-dispatch-and-handler-lookup/skill.md` | Antigravity skill for Map Dispatch |
| `.cursor/skills/cg-map-dispatch-and-handler-lookup/skill.md` | Cursor IDE skill mirror for Map Dispatch |
| `01-prompts/15-cg-execute/readme.md` | Updated index containing Prompt 38 |
| `01-prompts/v4/15-cg-execute/readme.md` | Updated V4 index containing Prompt 38 |
| `.agents/skills/cg-execute/skill.md` | Updated master index skill with Step 0 Heatmap preflight and Prompt 38 |
| `.cursor/skills/cg-execute/skill.md` | Updated Cursor master index skill |
| `01-prompts/v4/14-execute/02-execute-parent-task-with-n-steps-v6.md` | Corrected Heatmap command references and added mandatory Step 0 Heatmap preflight |
| `.agents/skills/execute-parent-task-with-n-steps-v6/skill.md` | Updated Antigravity V6 skill |
| `.cursor/skills/execute-parent-task-with-n-steps-v6/skill.md` | Updated Cursor V6 skill |
| `03-ai-scripts/47-upgrade-cg-execute-to-v6.py` | Enhanced upgrade engine to include Prompt 38, fix heatmap commands, and inject Step 0 Preflight |
| All 37 existing CG Execute prompts & skills | Standardized to strict V6 with fixed Heatmap & RC commands |

---

## 2. Heatmap & RC Command Contract

### Step 0 Mandatory Preflight Heatmap Commands:
1. `gitmap agent task init --name "<task name>" --budget 400`
2. `gitmap agent task status --json`
3. `gitmap rc heatmap tables` (inspect companion database schema and active tables)
4. `gitmap heatmap status` (verify repository local heatmap baseline)
5. `gitmap agent heatmap` (inspect active write leases and contention hotspots)

### In-Flight Cognitive Logging:
- Touch/Exploration: `gitmap heatmap log -t <slug> -a "READ" -f "<path>" -r "Reasoning: <rationale>" -s "EXPLORATION"`
- Modification: `gitmap heatmap log -t <slug> -a "EDIT" -f "<path>" -r "Decision: <rationale>" -s "EXECUTION"`
- Companion RC Logging: `gitmap rc heatmap log -t <slug> -a "<READ|EDIT>" -f "<path>" -r "<rationale>"`

### Tri-Alias Parity Guarantee:
- `gitmap heatmap rc tables` == `gitmap rc heatmap tables` == `gitmap heatmap rc-heatmap tables`
- `gitmap heatmap rc view <tbl>` == `gitmap rc heatmap view <tbl>` == `gitmap heatmap rc-heatmap view <tbl>`
- `gitmap heatmap rc query "<sql>"` == `gitmap rc heatmap query "<sql>"` == `gitmap heatmap rc-heatmap query "<sql>"`

---

## 3. Map Dispatch Coding Standard Rules
1. **Lookup Primacy:** Any chained `if-else if-else` or `switch` matching 3 or more literal keys/strings/tokens to execute discrete functions MUST be refactored into a map lookup table.
2. **Immutability:** Dispatch maps should be declared as module-level constants or component-level immutable records (e.g. `const KEY_ACTIONS: Record<KeyType, () => void>`).
3. **Discrete Handlers:** Handlers mapped inside the table MUST be named functions or concise arrows that invoke dedicated handlers, maintaining single-responsibility and <= 8–15 line function limits.
4. **Positive Booleans:** Lookup presence checks must use affirmative checks: `const handler = MAP[key]; if (handler) { ... }`.
5. **Structured Error Handling:** When returning results or errors in Go, map handlers must return monadic `Result[T]` or `*appfault.AppError`.
