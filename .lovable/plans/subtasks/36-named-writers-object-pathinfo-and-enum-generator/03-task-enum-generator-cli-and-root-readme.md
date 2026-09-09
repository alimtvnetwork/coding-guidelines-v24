# Subtask 36.3: Enum Generator CLI and Root Readme

## Status
- **State:** Complete
- **Assigned Files:**
  - `03-ai-scripts/30-enum-generator.py`
  - `readme.md`

## Acceptance Criteria
1. In `03-ai-scripts/30-enum-generator.py`:
   - Add `--out` and `-o` aliases for `--target-dir` so users can specify `--out=04-code/golang/pkg/enum/ordertype`.
   - Ensure the directory is automatically created if it does not already exist.
   - Run gofmt/code-formatter or verify formatting on write.
   - Keep all functions $\le 15$ lines and implicit booleans.
2. In `readme.md`:
   - Add a prominent section in the greeting / quickstart area: `✨ Auto-Generate Production Enums in One Line`.
   - Document the exact one-line command:
     `python 03-ai-scripts/30-enum-generator.py --name=ordertype --type=byte --items="Draft,Placed,Shipped,Delivered,Cancelled" --out=04-code/golang/pkg/enum/ordertype`
   - Summarize the 4 generated files (`variant.go`, `vars.go`, `variant_test.go`, `readme.md`) and how they comply with the repo standards (`baseenumer`, `Min()`, `Max()`, bounded interfaces, JSON unmarshaling, $\le 15$ line functions).
3. Test dry-run and live generation.
