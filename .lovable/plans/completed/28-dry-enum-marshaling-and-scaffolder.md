# Plan 28: DRY Enum Marshaling in `pkg/baseenumer` & Smart Multi-File Enum Scaffolder

## Overview
This plan implements a DRY, generic JSON marshaling architecture in `04-code/golang/pkg/baseenumer` to eliminate repetitive `MarshalJSON`, `UnmarshalJSON`, `unmarshalData`, and `unmarshalString` boilerplate across all 11 enum packages in the repository. Additionally, it rebuilds `03-ai-scripts/30-enum-generator.py` into a smart, multi-file Go enum scaffolder supporting byte-, int-, and string-backed enums with dual CLI comma and JSON input modes.

## Checklist

- [x] Step 28.1: Implement generic JSON marshaling in `04-code/golang/pkg/baseenumer/json_marshaling.go` and tests in `json_marshaling_test.go`.
- [x] Step 28.2: Refactor file I/O enums (`fileoptype`, `filewritemodetype`, `filepermtype`, `openfiletype`) to use DRY marshaling.
- [x] Step 28.3: Refactor diagnostic, process, and error enums (`logleveltype`, `processstatetype`, `prioritytype`, `severitytype`, `bytetype`, `pkg/errtype/processstatetype`, `pkg/errtype/logleveltype`) to use DRY marshaling.
- [x] Step 28.4: Overhaul `03-ai-scripts/30-enum-generator.py` to support 4-file package generation (`variant.go`, `vars.go`, `variant_test.go`, `readme.md`), byte/int/string backing types, CLI comma-separated input, and JSON config input.
- [x] Step 28.5: Run full verification suite (Go tests, code formatter, CI/CD local runner, sync-check) and move plan to completed.

## Subtasks
- [01-baseenumer-generic-marshaling.md](../subtasks/28-dry-enum-marshaling-and-scaffolder/01-baseenumer-generic-marshaling.md)
- [02-refactor-file-io-enums.md](../subtasks/28-dry-enum-marshaling-and-scaffolder/02-refactor-file-io-enums.md)
- [03-refactor-diagnostic-and-errtype-enums.md](../subtasks/28-dry-enum-marshaling-and-scaffolder/03-refactor-diagnostic-and-errtype-enums.md)
- [04-smart-enum-scaffolder-script.md](../subtasks/28-dry-enum-marshaling-and-scaffolder/04-smart-enum-scaffolder-script.md)
- [05-quality-gates-and-verification.md](../subtasks/28-dry-enum-marshaling-and-scaffolder/05-quality-gates-and-verification.md)
