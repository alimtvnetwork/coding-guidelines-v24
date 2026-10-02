# Architecture Specification: Variadic & Spread Array Parameters, Prompts V6 Alignment & Multi-Repository Guideline Synchronization

> **/goal** Establish the architectural foundation, cross-language contracts, and execution roadmap for variadic and spread array parameters across Go, TypeScript, and Rust, while aligning guideline execution prompts with V6 parameter headers and preparing multi-repository synchronization.
> **/learn** Master the ergonomic superiority of variadic parameters over rigid slices/arrays, zero-wrapper single-element invocations, spread forwarding, language-specific iterator patterns, V6 parameter headers (`N = 300, A = 2, H = 2, C = 30`), and atomic multi-repository rollout pipelines.

**Version:** 1.0.0  
**Updated:** 2026-10-02  
**Status:** Active  
**AI Confidence:** Production-Ready  
**Ambiguity:** None  

---

## 1. Executive Summary & Core Motivation

In polyglot engineering architectures (spanning Go, TypeScript, and Rust), function signatures that accept homogenous collections of items are frequently declared using rigid slice or array types:

```go
// Rigid slice parameter: forces artificial wrappers at every call site
func DeleteUsers(ids []string) *appfault.AppError
```

```typescript
// Rigid array parameter: forces artificial array brackets for single-item invocations
function deleteUsers(ids: string[]): Promise<Result<void>>
```

This rigid parameter pattern introduces systemic friction across codebases:

1. **Call-Site Clutter & Friction:** Single-item invocations—which often constitute over 70% of call sites in practice—are forced to construct artificial wrapper structures (`DeleteUsers([]string{id})` or `deleteUsers([id])`).
2. **Artificial Memory Allocations:** In Go and TypeScript, constructing ad-hoc single-item slice/array literals creates temporary heap or slice-header allocations that escape unnecessarily when passed across boundaries.
3. **Impaired Fluent Composability:** Rigid slice parameters prevent callers from naturally passing comma-separated arguments or chaining variadic builders and options without intermediate list instantiations.
4. **Ergonomic Inversion:** Callers possessing an existing collection can easily pass it to a variadic function using spread syntax (`items...` in Go, `...items` in TypeScript). Conversely, callers possessing a single item cannot easily invoke a rigid slice function without wrapping it.

### The Variadic & Spread Solution

By defining functions with **variadic / spread parameters** (`...T` in Go, `...items: T[]` or union overloads in TypeScript, and zero-cost slice references `&[T]` / `impl IntoIterator<Item = T>` in Rust), APIs achieve optimal ergonomics:

- **Single-Item Calls:** Directly passed without wrapper syntax: `DeleteUsers(id)` / `deleteUsers(id)`.
- **Multi-Item Calls:** Comma-separated without slice declaration: `DeleteUsers(id1, id2)` / `deleteUsers(id1, id2)`.
- **Slice Spread Calls:** Existing slices seamlessly unpacked: `DeleteUsers(allIDs...)` / `deleteUsers(...allIds)`.
- **Empty Invocations:** Zero parameters passed cleanly without `nil` or `[]` literals: `DeleteUsers()`.

---

## 2. Cross-Language Architectural Principles & Concrete Code Examples

### 2.1 Go (`...T` Variadic Slices & Parameter Functions)

In Go, variadic parameters (`...T`) are syntactic sugar for a slice parameter, but with a critical distinction: the Go compiler handles slice construction at the call site automatically. When a caller passes an existing slice using the `...` spread operator, no additional slice allocation occurs.

#### ❌ Anti-Pattern: Rigid Slice Parameters

```go
// ❌ ANTI-PATTERN: Forces single-item callers to allocate and wrap
func RemoveTags(resourceID string, tags []string) *appfault.AppError {
    if len(tags) == 0 {
        return nil
    }

    for _, tag := range tags {
        if isErr := purgeTag(resourceID, tag); isErr != nil {
            return isErr
        }
    }

    return nil
}

// Call site suffering from wrapper boilerplate:
err := RemoveTags("res-101", []string{"obsolete"})
```

#### ✅ Canonical Pattern: Variadic Parameter with Spread Support

```go
// ✅ CANONICAL: Variadic parameter with zero-wrapper single calls and slice spread
func RemoveTags(resourceID string, tags ...string) *appfault.AppError {
    hasTags := len(tags) > 0
    if !hasTags {
        return nil
    }

    for _, tag := range tags {
        appErr := purgeTag(resourceID, tag)
        if appErr != nil {
            return appErr
        }
    }

    return nil
}

// Call site 1 (Single item — zero boilerplate):
err1 := RemoveTags("res-101", "obsolete")

// Call site 2 (Multiple comma-separated items):
err2 := RemoveTags("res-101", "obsolete", "deprecated", "stale")

// Call site 3 (Existing slice — unpacked with spread):
existingTags := []string{"tag-a", "tag-b"}
err3 := RemoveTags("res-101", existingTags...)

// Call site 4 (Zero items — completely valid):
err4 := RemoveTags("res-101")
```

#### Guidelines for Go Parameter Limits & Structs

Under Prompt Architect guidelines, functions must not exceed 2–3 loose parameters. The variadic parameter counts as **one parameter slot** and MUST be placed as the final parameter. When a function requires additional configuration options along with a collection, encapsulate the configuration into a `*Params` struct or utilize the functional options pattern:

```go
// ✅ Compliant with <= 3 parameters rule: 1 struct + 1 variadic slice
func BulkUpdateUsers(params UserUpdateParams, userIDs ...string) *appfault.AppError {
    hasUsers := len(userIDs) > 0
    if !hasUsers {
        return nil
    }

    for _, id := range userIDs {
        appErr := executeUserUpdate(id, params)
        if appErr != nil {
            return appErr
        }
    }

    return nil
}
```

---

### 2.2 TypeScript (`...items: T[]` Rest Parameters & Single-or-Array Unions)

TypeScript provides two idiomatic techniques for flexible collection ingestion:
1. **Rest Parameters (`...items: readonly T[]`):** For functions where items represent the trailing or primary arguments.
2. **Single-or-Array Normalization (`T | readonly T[]`):** For configuration objects, properties, or functions where an options bag is passed.

#### ❌ Anti-Pattern: Rigid Array Ingestion

```typescript
// ❌ ANTI-PATTERN: Rigid array forces array literals everywhere
export async function trackEvents(events: string[]): Promise<Result<void>> {
  if (events.length === 0) {
    return Result.ok(undefined);
  }

  for (const event of events) {
    await emitTelemetry(event);
  }

  return Result.ok(undefined);
}

// Call site: awkward brackets for the most common case
await trackEvents(['user_signup']);
```

#### ✅ Canonical Pattern A: Rest Parameters (`...events`)

```typescript
// ✅ CANONICAL: Rest parameters with readonly array safety
export async function trackEvents(
  ...events: readonly string[]
): Promise<Result<void>> {
  const hasEvents = events.length > 0;
  if (!hasEvents) {
    return Result.ok(undefined);
  }

  for (const event of events) {
    const result = await emitTelemetry(event);
    if (!result.isSuccess) {
      return result;
    }
  }

  return Result.ok(undefined);
}

// Call site 1 (Single item):
await trackEvents('user_signup');

// Call site 2 (Multiple comma-separated items):
await trackEvents('click_nav', 'open_modal', 'submit_form');

// Call site 3 (Spread existing array):
const backlog = ['event_1', 'event_2'];
await trackEvents(...backlog);
```

#### ✅ Canonical Pattern B: Single-or-Array Union Normalizer

When an argument is passed inside a parameter struct or options object, use the `SingleOrArray<T>` union pattern paired with a zero-mutation normalizer:

```typescript
export type SingleOrArray<T> = T | readonly T[];

export function normalizeToArray<T>(input: SingleOrArray<T> | undefined): readonly T[] {
  const isUndefined = input === undefined;
  if (isUndefined) {
    return [];
  }

  const isArray = Array.isArray(input);
  if (isArray) {
    return input as readonly T[];
  }

  return [input as T];
}

// Configuration struct allowing single item or array:
export interface FilterOptions {
  readonly categories?: SingleOrArray<string>;
  readonly limit?: number;
}

export function applyFilters(options: FilterOptions): Result<FilteredView> {
  const categories = normalizeToArray(options.categories);
  // Seamlessly handles either:
  // { categories: "billing" } OR { categories: ["billing", "security"] }
  return Result.ok(buildView(categories, options.limit));
}
```

---

### 2.3 Rust (Zero-Cost Slices `&[T]` & Generic Iterators `impl IntoIterator<Item = T>`)

Rust does not support variadic arguments in safe function signatures (outside declarative macros or C FFI). However, Rust achieves identical or superior ergonomics and zero runtime cost through two canonical patterns:

1. **Borrowed Slice References (`&[T]`):** Zero-allocation borrowing over arrays, vectors, or single-element arrays (`&[item]`).
2. **Generic Collection Trait (`impl IntoIterator<Item = T>` or `impl AsRef<[T]>`):** Permits passing arrays, slices, vectors, or single-item iterators (`std::iter::once`).

#### ❌ Anti-Pattern: Owned Heap Vector Parameter

```rust
// ❌ ANTI-PATTERN: Forces caller to allocate a Vec on the heap for every call
pub fn process_identifiers(ids: Vec<String>) -> Result<(), AppError> {
    if ids.is_empty() {
        return Ok(());
    }

    for id in ids {
        execute_step(&id)?;
    }

    Ok(())
}

// Call site forced to allocate heap buffer for a single ID:
process_identifiers(vec!["id-100".to_string()])?;
```

#### ✅ Canonical Pattern: Borrowed Slice or Generic IntoIterator

```rust
// ✅ CANONICAL PATTERN 1: Borrowed slice reference (zero allocation)
pub fn process_identifiers(ids: &[&str]) -> Result<(), AppError> {
    let has_items = !ids.is_empty();
    if !has_items {
        return Ok(());
    }

    for id in ids {
        execute_step(id)?;
    }

    Ok(())
}

// Call site 1 (Single item borrowed directly from stack array):
process_identifiers(&["id-100"])?;

// Call site 2 (Multiple comma-separated items in stack array):
process_identifiers(&["id-100", "id-101", "id-102"])?;

// Call site 3 (Borrowed existing Vec):
let my_vec = vec!["a", "b"];
process_identifiers(&my_vec)?;
```

```rust
// ✅ CANONICAL PATTERN 2: Generic iterator for ultimate flexibility
pub fn ingest_records<I, S>(records: I) -> Result<(), AppError>
where
    I: IntoIterator<Item = S>,
    S: AsRef<str>,
{
    for record in records {
        dispatch_record(record.as_ref())?;
    }

    Ok(())
}

// Call site: accepts standard array, vector, or std::iter::once without heap wrappers:
ingest_records(["alpha", "beta"])?;
ingest_records(std::iter::once("single-entry"))?;
```

---

## 3. Comprehensive Good vs Bad Contrast Matrix

| Scenario | ❌ Rigid Slice / Array Pattern | ✅ Variadic / Spread Pattern | Ergonomic & Performance Gain |
|:---|:---|:---|:---|
| **Single item invocation (Go)** | `fn([]string{"id"})` | `fn("id")` | Zero slice wrapper syntax; compiler escapes eliminated. |
| **Multiple items invocation (Go)** | `fn([]string{"a", "b", "c"})` | `fn("a", "b", "c")` | Clean comma-delimited call; no type annotation at call site. |
| **Existing slice invocation (Go)** | `fn(existingSlice)` | `fn(existingSlice...)` | Explicit spread operator `...`; identical performance. |
| **Empty collection invocation (Go)** | `fn(nil)` or `fn([]string{})` | `fn()` | Clean zero-argument invocation; slice initialized as empty. |
| **Single item invocation (TS)** | `fn(["user_1"])` | `fn("user_1")` | No array brackets; natural function invocation. |
| **Multiple items invocation (TS)** | `fn(["a", "b", "c"])` | `fn("a", "b", "c")` | Comma-separated arguments; matches standard JS rest conventions. |
| **Spread existing array (TS)** | `fn(items)` | `fn(...items)` | Unpacks collection cleanly with native JS spread operator. |
| **Single item invocation (Rust)** | `fn(vec!["id"])` (heap alloc) | `fn(&["id"])` or `fn(once("id"))` | Zero-allocation stack array borrow; no heap churn. |
| **Readability & Intent** | Emphasizes collection container over domain values | Emphasizes domain values; container is an implementation detail | Lower cognitive load; higher signal-to-noise ratio in code. |

---

## 4. Implementation Blueprint for Meta-Repository Deliverables

To implement and enforce this standard across the meta-repository and connected repositories, the following four deliverables must be executed:

### Deliverable 1: Cross-Language Coding Guideline Specification
- **Target File:** `02-spec/02-coding-guidelines/01-cross-language/33-variadic-and-spread-parameters.md`
- **Scope:**
  - Define rules R1 through R6 for variadic parameters, spread forwarders, single-or-array normalizers, and iterator abstractions.
  - Detail concrete Go, TypeScript, and Rust transformation recipes.
  - Enforce vertical blank line rules, positive booleans, and `*appfault.AppError` returns in all examples.
  - Update `02-spec/02-coding-guidelines/01-cross-language/readme.md` to register spec #33.

### Deliverable 2: Coding Guideline Execution Prompt #36
- **Target File:** `01-prompts/15-cg-execute/36-variadic-and-spread-parameters.md`
- **Scope:**
  - Sequence as prompt #36 in `01-prompts/15-cg-execute/` catalog.
  - Implement full V6 parameter header format (`N = 300, A = 2, H = 2, C = 30`).
  - Mandatory subagent spawning gate (`invoke_subagent` for Phase 1 discovery and Phase 2 execution).
  - Include 3-Phase lifecycle (Verbatim capture, Spec & Subtasks, Disjoint file refactoring, GitMap atomic push).
  - Zero-build and zero-test routine execution ban (linters only).
  - Update `01-prompts/15-cg-execute/readme.md` and `01-prompts/readme.md`.

### Deliverable 3: Native Antigravity Skill
- **Target File:** `.agents/skills/cg-variadic-and-spread-parameters/skill.md`
- **Scope:**
  - YAML frontmatter with `name: cg-variadic-and-spread-parameters` and trigger descriptions.
  - Provide actionable instructions for AST scanning, finding functions taking rigid single-slice parameters, and converting to variadic / spread syntax.
  - Provide call-site refactoring transformations across Go, TypeScript, and Rust.

### Deliverable 4: Align Prompts 31 & 32 with V6 Parameter Headers
- **Target Files:**
  - `01-prompts/15-cg-execute/31-cg-execute-in-below-steps.md`
  - `01-prompts/15-cg-execute/32-cg-follow-other-prompts.md`
- **Scope:**
  - Standardize top header blocks to V6 parameters:
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
  - Standardize GitMap atomic commit instructions to hyphen format (`gitmap cpf "<module> - <summary>"` / `gitmap cpb "<module> - <summary>"`).
  - Align with mandatory subagent spawning gate and zero-test / zero-build rules.

### Deliverable 5: Multi-Repository Synchronization & Deployment
- **Target Script:** `03-ai-scripts/38-sync-prompts-skills-scripts.py`
- **Scope:**
  - Ensure prompt 36 and the new skill are included in sync manifests.
  - Synchronize updated prompts, skills, and coding guidelines across all connected repositories.
  - Execute standard 4-stage pipeline: `pull` -> `pre-sync backup & release` -> `sync` -> `post-sync release`.

---

## 5. Verification & Acceptance Criteria

### AC-APP-015: Variadic Parameter Specification & Prompt Alignment Verification

**Given** The architecture spec, subtask plans, coding guideline #33, prompt #36, skill `cg-variadic-and-spread-parameters`, and aligned prompts #31 and #32 are authored.  
**When** Running repository sequence integrity, relative path, and prompt load linters.  
**Then** All file paths are strictly relative; filenames strictly lowercase; prompts contain V6 parameter headers; and no broken links or sequence gaps exist.

**Verification commands:**

```bash
python linter-scripts/check-prompts-loaded.py
python linter-scripts/check-relative-paths.py
python 03-ai-scripts/21-sequence-integrity-linter.py 02-spec/02-coding-guidelines/01-cross-language
python 03-ai-scripts/21-sequence-integrity-linter.py 01-prompts/15-cg-execute
```

**Expected Result:** Exit code 0 across all linters.
