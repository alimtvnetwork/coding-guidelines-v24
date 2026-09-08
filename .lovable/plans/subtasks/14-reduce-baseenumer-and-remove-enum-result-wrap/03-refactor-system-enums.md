# Subtask 03: Refactor System Enums (bytetype, filepermtype, openfiletype)

> **Plan:** `14-reduce-baseenumer-and-remove-enum-result-wrap`  
> **Status:** Completed  
> **Target File:** `04-code/golang/pkg/enum/bytetype/*`, `04-code/golang/pkg/enum/filepermtype/*`, `04-code/golang/pkg/enum/openfiletype/*`  

---

## Intent
1. `bytetype`: Remove `errtype`/`result` imports, change `Parse(s string)` to return `(Variant, bool)` using `basicEnum.ParseLookup` + numeric fallback, add `ParseOrZero`.
2. `filepermtype`: Remove `errtype`/`result` imports, change `Parse(octalStr string)` to return `(Variant, bool)`, add `ParseOrZero`.
3. `openfiletype`: Remove `errtype`/`result` imports, change `Parse(s string)` to `return basicEnum.Parse(s)`, add `ParseOrInvalid`.
4. Update unit tests in each package to verify `(Variant, bool)`.

## Verification
Run:
```bash
go test ./pkg/enum/bytetype -v -count=1
go test ./pkg/enum/filepermtype -v -count=1
go test ./pkg/enum/openfiletype -v -count=1
```
