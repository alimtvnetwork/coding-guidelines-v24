# Subtask 01: Author Cross-Language Coding Guideline for Variadic & Spread Parameters

> **Parent Spec:** [`02-spec/21-app/01-variadic-spread-params-and-multi-repo/01-architecture-spec.md`](../../../../02-spec/21-app/01-variadic-spread-params-and-multi-repo/01-architecture-spec.md)  
> **Status:** `PENDING`  
> **Traceability ID:** `Task-01`  
> **Target Files:**
> - `02-spec/02-coding-guidelines/01-cross-language/33-variadic-and-spread-parameters.md` (to create)
> - `02-spec/02-coding-guidelines/01-cross-language/readme.md` (to update registry)
> - `.ai-memory/coding-guidelines.md` (to update reference index)

---

## 1. Subtask Objective & Context

Author the canonical cross-language coding guideline specification:
`02-spec/02-coding-guidelines/01-cross-language/33-variadic-and-spread-parameters.md`.

This specification codifies the rule that functions accepting homogenous collections of elements must provide **variadic / spread parameters** (`...T` in Go, `...items: T[]` in TypeScript, and borrowed slices `&[T]` / generic iterators `impl IntoIterator<Item = T>` in Rust), eliminating call-site wrapper boilerplate (`[]string{id}`, `[id]`, `vec![id]`) and unnecessary heap allocations.

---

## 2. Core Architectural Rules to Formalize

The guideline file must codify the following six core rules (R1 to R6):

### R1: Variadic Parameter Primacy for Homogenous Collections
- Functions whose primary or trailing input is a collection of identical type elements SHOULD declare that parameter as variadic (`...T` in Go, `...items: readonly T[]` in TypeScript).
- Ban declaring rigid slice parameters (`items []string`) for utility, deletion, filtering, registration, and dispatch APIs where callers frequently operate on single elements.

### R2: Spread Operator Forwarding & Transparency
- Functions that receive a variadic parameter and pass it to another variadic function MUST forward using the spread operator (`targetFunc(items...)` in Go, `targetFunc(...items)` in TypeScript).
- Callers that possess an existing slice or array pass it directly using spread without rebuilding a slice.

### R3: Zero Artificial Wrapper Mandate at Call Sites
- Callers passing a single item MUST pass the item directly (`Delete(id)`) without wrapping it in an artificial slice or array literal (`Delete([]string{id})` is banned).
- In tests and seed fixtures, multiple items MUST be passed comma-separated (`Add("a", "b", "c")`), not as a pre-constructed slice unless verifying slice spread behavior.

### R4: Parameter Limit Interoperability (<= 2-3 Parameters Rule)
- A variadic parameter occupies exactly one parameter slot in the function signature.
- Functions must still strictly adhere to the prompt architect constraint banning > 2-3 loose parameters.
- When auxiliary configuration options are needed alongside collection items, place configuration into a dedicated `*Params` struct and keep the variadic parameter at the trailing position (e.g. `func Dispatch(params DispatchParams, targets ...string) *appfault.AppError`).

### R5: Safe Nil & Empty Handling
- Implementations must handle zero variadic arguments safely.
- In Go, `len(items) == 0` when zero arguments are passed. Use affirmative checks: `hasItems := len(items) > 0`.
- In TypeScript, default rest array is `[]`. Check: `const hasItems = items.length > 0`.
- In Rust, borrowed slices `&[T]` check `!items.is_empty()`.

### R6: AppError & Vertical Spacing Conformance
- Go implementations returning failure metadata MUST return `*appfault.AppError`.
- Strict vertical line spacing must be observed: blank line before `if`, blank line after `}`, blank line before `return`.
- Strict boolean conventions: `is` and `has` prefixes only, zero explicit `== true`, and zero mixed polarity.

---

## 3. Concrete Code Patterns & Anti-Patterns to Document

### Go (`...T`)

```go
// ❌ WRONG: Rigid slice forces callers to wrap single items
func InvalidateCacheKeys(keys []string) *appfault.AppError

// ❌ Call site friction:
err := InvalidateCacheKeys([]string{sessionKey})

// ✅ REQUIRED: Variadic parameter
func InvalidateCacheKeys(keys ...string) *appfault.AppError {
    hasKeys := len(keys) > 0
    if !hasKeys {
        return nil
    }

    for _, key := range keys {
        appErr := purgeKey(key)
        if appErr != nil {
            return appErr
        }
    }

    return nil
}

// ✅ Clean call site:
err := InvalidateCacheKeys(sessionKey)
```

### TypeScript (`...items: readonly T[]` & `SingleOrArray<T>`)

```typescript
// ❌ WRONG: Rigid array forces array literal for single item
function emitMetrics(metrics: MetricEvent[]): Promise<Result<void>>

// ✅ REQUIRED: Rest parameter
function emitMetrics(
  ...metrics: readonly MetricEvent[]
): Promise<Result<void>> {
  const hasMetrics = metrics.length > 0;
  if (!hasMetrics) {
    return Result.ok(undefined);
  }

  for (const metric of metrics) {
    recordMetric(metric);
  }

  return Result.ok(undefined);
}

// Call site:
await emitMetrics(singleMetric);
await emitMetrics(metricA, metricB);
await emitMetrics(...metricBatch);
```

### Rust (`&[T]` & `impl IntoIterator<Item = T>`)

```rust
// ❌ WRONG: Owned Vec parameter causes unnecessary heap allocation for single item
pub fn register_subscribers(subscribers: Vec<SubscriberId>) -> Result<(), AppError>

// ✅ REQUIRED: Borrowed slice reference allows stack-borrowed single item
pub fn register_subscribers(subscribers: &[SubscriberId]) -> Result<(), AppError> {
    let has_subscribers = !subscribers.is_empty();
    if !has_subscribers {
        return Ok(());
    }

    for sub in subscribers {
        persist_subscriber(sub)?;
    }

    Ok(())
}

// Call site: zero allocation stack slice
register_subscribers(&[active_subscriber])?;
```

---

## 4. Execution Step-by-Step Checklist

- [ ] **Step 1:** Create `02-spec/02-coding-guidelines/01-cross-language/33-variadic-and-spread-parameters.md` adhering to the standard template (Metadata, Motivation, Rules R1–R6, Code Examples for Go/TS/Rust, Anti-Patterns vs Canonical Patterns, Acceptance Criteria).
- [ ] **Step 2:** Update `02-spec/02-coding-guidelines/01-cross-language/readme.md` table to include entry `33-variadic-and-spread-parameters.md` with active status and cross-references.
- [ ] **Step 3:** Update `.ai-memory/coding-guidelines.md` section to reference rule #33.
- [ ] **Step 4:** Run sequence integrity linter `python 03-ai-scripts/21-sequence-integrity-linter.py 02-spec/02-coding-guidelines/01-cross-language`.
- [ ] **Step 5:** Run relative path checker `python linter-scripts/check-relative-paths.py`.

---

## 5. Verification & Acceptance Criteria

- **AC-CG-033-A:** Spec file `02-spec/02-coding-guidelines/01-cross-language/33-variadic-and-spread-parameters.md` exists and contains 100% relative paths.
- **AC-CG-033-B:** Registry in `02-spec/02-coding-guidelines/01-cross-language/readme.md` contains sequence #33 without gap.
- **AC-CG-033-C:** All code examples pass guideline checks (positive booleans, `*appfault.AppError`, vertical line gaps).

**Verification commands:**

```bash
python linter-scripts/check-relative-paths.py
python 03-ai-scripts/21-sequence-integrity-linter.py 02-spec/02-coding-guidelines/01-cross-language
```

**Expected Result:** Exit code 0.
