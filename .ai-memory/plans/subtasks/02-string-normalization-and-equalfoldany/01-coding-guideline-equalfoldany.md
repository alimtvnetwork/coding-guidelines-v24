# Subtask 01: Author Cross-Language Coding Guideline for String Normalization & EqualFoldAny

> **Parent Spec:** [`02-spec/21-app/02-string-normalization-and-equalfoldany/01-architecture-spec.md`](../../../../02-spec/21-app/02-string-normalization-and-equalfoldany/01-architecture-spec.md)  
> **Status:** `PENDING`  
> **Traceability ID:** `Task-02-Subtask-01`  
> **Target Files:**
> - `02-spec/02-coding-guidelines/01-cross-language/34-string-normalization-and-equalfoldany.md` (to create)
> - `02-spec/02-coding-guidelines/01-cross-language/readme.md` (to update registry)
> - `.ai-memory/coding-guidelines.md` (to update reference index)

---

## 1. Subtask Objective & Context

Author the canonical cross-language coding guideline specification:  
`02-spec/02-coding-guidelines/01-cross-language/34-string-normalization-and-equalfoldany.md`.

This specification codifies the rule that codebases must eliminate repeated string manipulation (`ToLower`, `TrimSpace`) and chained equality checks (`EqualFold(s, "y") || EqualFold(s, "yes")`), replacing them with a centralized, variadic string utility function (`EqualFoldAnyTrim` / `EqualFoldAny`). Furthermore, it enforces the **Search First Protocol**, requiring AI agents and developers to inspect canonical utility locations (`pkg/strutil/strutil.go`, `src/lib/strutil.ts`, `src/util/strutil.rs`, `pkg/strutil/strutil.py`) before writing any string helpers, eradicating duplicate, unshared local functions.

---

## 2. Core Architectural Rules to Formalize

The guideline file must formalize the following six core rules (R1 to R6):

### R1: Centralized String Utility Primacy (Banned Ad-Hoc Chained Comparisons)
- Application code, CLI commands, HTTP handlers, and services MUST NOT evaluate multiple string candidate matches using inline chained OR (`||`) expressions with repeated equality checks or transformations (e.g. `strings.EqualFold(s, "a") || strings.EqualFold(s, "b")`).
- Callers MUST invoke the centralized string utility function (`strutil.EqualFoldAnyTrim` or `strutil.EqualFoldAny`).
- Banned patterns include chaining `EqualFold`, repeating `TrimSpace()`, or combining `.trim().toLowerCase() === ...` chains.

### R2: Target-First Variadic Candidate Signature
- The string utility function MUST place the target string being tested as the first parameter.
- Candidate matching strings MUST be accepted as trailing variadic arguments:
  - Go: `EqualFoldAnyTrim(target string, candidates ...string) bool`
  - TypeScript: `equalFoldAnyTrim(target: string, ...candidates: readonly string[]): boolean`
  - Rust: `equal_fold_any_trim(target: &str, candidates: &[&str]) -> bool`
  - Python: `equal_fold_any_trim(target: str, *candidates: str) -> bool`
- Callers can pass individual string candidates as comma-separated values or unpack an existing slice/array using the language's native spread operator (`candidates...`, `...candidates`).

### R3: Pre-Normalization & Short-Circuit Optimization
- `EqualFoldAnyTrim` MUST perform whitespace trimming on the target string exactly once upfront prior to testing candidate values.
- Candidate iteration MUST short-circuit and return `true` immediately upon finding the first match, avoiding redundant checks or allocations.
- If zero candidates are provided, the utility MUST safely return `false` without panicking or throwing errors.

### R4: "Search First" Canonical Location Mandate
- Before authoring or proposing any string comparison helper, AI agents and engineers MUST inspect the repository's canonical utility package:
  - Go: `pkg/strutil/strutil.go` (or `internal/strutil/strutil.go`)
  - TypeScript: `src/lib/strutil.ts` (or `src/utils/strutil.ts`)
  - Rust: `src/util/strutil.rs` (or `src/strutil/mod.rs`)
  - Python: `pkg/strutil/strutil.py` (or `src/strutil.py`)
- Authoring private, unshared helper functions inside individual feature or command files (such as `isYes()`, `checkConfirm()`, or `matchesOption()`) is **STRICTLY PROHIBITED**.
- If the utility does not yet exist in the repository, it MUST be added to the canonical `strutil` location rather than introduced as an isolated private helper.

### R5: Prompt Architect Boolean Hygiene
- Positive booleans MUST ALWAYS be evaluated implicitly: `if isMatch { ... }`.
- Evaluating booleans explicitly against `true` (e.g. `if isMatch == true`) is **TOTALLY BANNED**.
- Mixed polarity within a single condition (e.g. `if isReady && !isBlocked`) is **TOTALLY BANNED**; split into separate discrete guard clauses.
- Boolean variables and helper return values MUST use `is` or `has` prefixes (e.g. `isMatch`, `hasCandidate`).

### R6: Strict Vertical Line Spacing
- Maintain mandatory blank line spacing across all code examples and implementations:
  - Blank line before every `if` statement.
  - Blank line after every closing brace `}`.
  - Blank line before every `return` statement.
  - Blank lines around multiline struct initializations and function parameters.

---

## 3. Concrete Code Patterns & Anti-Patterns to Document

### Go (`pkg/strutil/strutil.go` & `cli/cmd/releaseundo.go`)

```go
// ❌ WRONG: Chained EqualFold with manual TrimSpace (releaseundo.go anti-pattern)
func confirmUndoRelease(tag string) bool {
    fmt.Printf("Delete %s locally and on origin? [y/N]: ", tag)
    var reply string
    _, _ = fmt.Scanln(&reply)
    trimmed := strings.TrimSpace(reply)

    return strings.EqualFold(trimmed, "y") || strings.EqualFold(trimmed, "yes")
}

// ✅ REQUIRED: Centralized strutil.EqualFoldAnyTrim invocation
func confirmUndoRelease(tag string) bool {
    fmt.Printf("Delete %s locally and on origin? [y/N]: ", tag)
    var reply string
    _, _ = fmt.Scanln(&reply)

    return strutil.EqualFoldAnyTrim(reply, "y", "yes")
}

// ✅ Canonical package implementation in pkg/strutil/strutil.go:
package strutil

import "strings"

// EqualFoldAny reports whether target matches any candidate under Unicode case-folding.
func EqualFoldAny(target string, candidates ...string) bool {
    for _, candidate := range candidates {
        isMatch := strings.EqualFold(target, candidate)
        if isMatch {
            return true
        }
    }

    return false
}

// EqualFoldAnyTrim reports whether target (after trimming whitespace) matches any candidate.
func EqualFoldAnyTrim(target string, candidates ...string) bool {
    trimmed := strings.TrimSpace(target)

    return EqualFoldAny(trimmed, candidates...)
}
```

### TypeScript (`src/lib/strutil.ts`)

```typescript
// ❌ WRONG: Inefficient repeated chaining and array inclusion
function isAffirmative(input: string): boolean {
  const clean = input.trim().toLowerCase();
  return clean === 'y' || clean === 'yes' || clean === 'true';
}

// ✅ REQUIRED: Centralized equalFoldAnyTrim
import { equalFoldAnyTrim } from '@/lib/strutil';

const isConfirmed = equalFoldAnyTrim(userInput, 'y', 'yes', 'true');

// ✅ Canonical implementation in src/lib/strutil.ts:
export function equalFoldAny(
  target: string,
  ...candidates: readonly string[]
): boolean {
  const normalizedTarget = target.toLowerCase();

  for (const candidate of candidates) {
    const isMatch = normalizedTarget === candidate.toLowerCase();
    if (isMatch) {
      return true;
    }
  }

  return false;
}

export function equalFoldAnyTrim(
  target: string,
  ...candidates: readonly string[]
): boolean {
  const trimmed = target.trim();

  return equalFoldAny(trimmed, ...candidates);
}
```

### Rust (`src/util/strutil.rs`)

```rust
// ❌ WRONG: Manual trimming and chained eq_ignore_ascii_case
fn is_positive_response(input: &str) -> bool {
    let trimmed = input.trim();
    trimmed.eq_ignore_ascii_case("y") || trimmed.eq_ignore_ascii_case("yes")
}

// ✅ REQUIRED: Canonical equal_fold_any_trim
use crate::util::strutil::equal_fold_any_trim;

let is_positive = equal_fold_any_trim(user_input, &["y", "yes"]);

// ✅ Canonical implementation in src/util/strutil.rs:
pub fn equal_fold_any(target: &str, candidates: &[&str]) -> bool {
    for candidate in candidates {
        let is_match = target.eq_ignore_ascii_case(candidate);
        if is_match {
            return true;
        }
    }

    false
}

pub fn equal_fold_any_trim(target: &str, candidates: &[&str]) -> bool {
    let trimmed = target.trim();

    equal_fold_any(trimmed, candidates)
}
```

### Python (`pkg/strutil/strutil.py`)

```python
# ❌ WRONG: Chained lower/strip checks scattered across scripts
if s.strip().lower() == "y" or s.strip().lower() == "yes":
    process()

# ✅ REQUIRED: Canonical equal_fold_any_trim
from pkg.strutil import equal_fold_any_trim

if equal_fold_any_trim(user_input, "y", "yes"):
    process()

# ✅ Canonical implementation in pkg/strutil/strutil.py:
def equal_fold_any(target: str, *candidates: str) -> bool:
    """Report whether target matches any candidate case-insensitively."""
    normalized_target = target.casefold()

    for candidate in candidates:
        is_match = normalized_target == candidate.casefold()
        if is_match:
            return True

    return False


def equal_fold_any_trim(target: str, *candidates: str) -> bool:
    """Report whether trimmed target matches any candidate case-insensitively."""
    trimmed = target.strip()

    return equal_fold_any(trimmed, *candidates)
```

---

## 4. Execution Step-by-Step Checklist

- [ ] **Step 1:** Create `02-spec/02-coding-guidelines/01-cross-language/34-string-normalization-and-equalfoldany.md` adhering to the standard template (Metadata, Motivation, Rules R1–R6, Code Examples for Go/TS/Rust/Python, Anti-Patterns vs Canonical Patterns, Acceptance Criteria).
- [ ] **Step 2:** Update `02-spec/02-coding-guidelines/01-cross-language/readme.md` table to register entry `34-string-normalization-and-equalfoldany.md` with active status and cross-references.
- [ ] **Step 3:** Update `.ai-memory/coding-guidelines.md` index to include rule #34 and summary of `EqualFoldAnyTrim` standard.
- [ ] **Step 4:** Run sequence integrity linter `python 03-ai-scripts/21-sequence-integrity-linter.py 02-spec/02-coding-guidelines/01-cross-language`.
- [ ] **Step 5:** Run relative path checker `python linter-scripts/check-relative-paths.py`.

---

## 5. Verification & Acceptance Criteria

- **AC-CG-034-A:** Spec file `02-spec/02-coding-guidelines/01-cross-language/34-string-normalization-and-equalfoldany.md` exists and contains 100% relative paths.
- **AC-CG-034-B:** Registry in `02-spec/02-coding-guidelines/01-cross-language/readme.md` contains sequence #34 without gap.
- **AC-CG-034-C:** All code examples pass guideline checks (positive booleans, zero `== true`, zero mixed polarity, vertical line gaps preserved).
- **AC-CG-034-D:** Includes verbatim documentation of the `releaseundo.go` case study.

**Verification commands:**

```bash
python linter-scripts/check-relative-paths.py
python 03-ai-scripts/21-sequence-integrity-linter.py 02-spec/02-coding-guidelines/01-cross-language
```

**Expected Result:** Exit code 0.
