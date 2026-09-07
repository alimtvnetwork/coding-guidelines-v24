# Subtask 27.3: PayloadConv Direct Result Type

## Objective
Define `type BytesResult = result.Wrap[[]byte]` in `pkg/payloadconv/` and update `ToBytes(payload any) BytesResult`.

## Target Files
1. `04-code/golang/pkg/payloadconv/converter.go` (and test)

## Acceptance Criteria
- `ToBytes` returns `BytesResult`.
- `pkg/payloadconv` tests pass 100%.
