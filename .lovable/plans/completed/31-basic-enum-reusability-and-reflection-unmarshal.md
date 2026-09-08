# Master Plan 31: BasicEnum Reusability & Reflection-Based Type Name Resolution

## Overview
Implement reusable generic enum constructs (`BasicIntegerEnum` and `BasicStringEnum`) in `04-code/golang/pkg/baseenumer`, eliminate manual `typeName string` parameters through reflection-based type resolution (`ResolveTypeName`), reduce `UnmarshalJSON` parameter count to 2 (`data, target`), and update the Python enum generator (`03-ai-scripts/30-enum-generator.py`) and existing enums to use this clean architecture.

## Checklist

- [x] Step 31.1: Add reflection-based `ResolveTypeName[V any](target *V) string` in `04-code/golang/pkg/baseenumer/resolve_type.go` and update `UnmarshalStringJSON` / `UnmarshalIntegerJSON` to eliminate explicit `typeName` parameter.
- [x] Step 31.2: Implement `BasicIntegerEnum[V IntNumber]` and `BasicStringEnum[V ~string]` in `04-code/golang/pkg/baseenumer/basic_enum.go` providing `UnmarshalJSON(data, target)`, `All()`, `Values()`, and `Parse(s)`.
- [x] Step 31.3: Update `03-ai-scripts/30-enum-generator.py` to generate enums using `basicEnum` with 2-parameter `UnmarshalJSON`.
- [x] Step 31.4: Refactor repository enum packages to use `basicEnum` / streamlined unmarshaling.
- [x] Step 31.5: Run full verification suite (all 25 Go packages, code formatter, sequence integrity, 31 CI/CD quality gates, sync-check) and move plan to completed.

## Subtasks
- [01-task-reflection-type-name-and-unmarshal-signature.md](../subtasks/31-basic-enum-reusability-and-reflection-unmarshal/01-task-reflection-type-name-and-unmarshal-signature.md)
- [02-task-basic-enum-structs.md](../subtasks/31-basic-enum-reusability-and-reflection-unmarshal/02-task-basic-enum-structs.md)
- [03-task-python-enum-generator.md](../subtasks/31-basic-enum-reusability-and-reflection-unmarshal/03-task-python-enum-generator.md)
- [04-task-repo-wide-enum-refactor.md](../subtasks/31-basic-enum-reusability-and-reflection-unmarshal/04-task-repo-wide-enum-refactor.md)
- [05-task-quality-gates-and-verification.md](../subtasks/31-basic-enum-reusability-and-reflection-unmarshal/05-task-quality-gates-and-verification.md)
