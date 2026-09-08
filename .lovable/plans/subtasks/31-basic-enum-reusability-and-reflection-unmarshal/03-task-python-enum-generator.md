# Subtask 31.3: Update Python Enum Generator Scaffolder

## Context
Update `03-ai-scripts/30-enum-generator.py` so that newly scaffolded enums automatically use `basicEnum` with 2-parameter `UnmarshalJSON`.

## Target File
`03-ai-scripts/30-enum-generator.py` [MODIFY]:
1. Update `vars.go` template:
   - Instantiate `var basicEnum = baseenumer.NewBasicInteger(variantLabels[:], Invalid)` or `NewBasicString`.
   - Implement `All()`, `Values()`, `Parse(s)` by delegating directly to `basicEnum`.
2. Update `variant.go` template:
   - `func (v *Variant) UnmarshalJSON(data []byte) error { return basicEnum.UnmarshalJSON(data, v) }`.

## Verification
- Test generator with `--dry-run` on string and integer types.
- Ensure generated code passes linting and function length caps ($\le 15$ lines).
