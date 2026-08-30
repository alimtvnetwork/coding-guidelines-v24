# Multi-Language Enums, Traits & Pattern Matching — Coding Guideline (must follow)

Trigger Keywords & Aliases: `cg-enums-traits`, `cg-enums`, `cg-execute enums`, `audit enums`, `php enums traits`, `rust enums`, `golang enums`, `multi-language enums`, `pattern matching audit`


> **Prompt Version:** 2.1.0
> **Synchronization:** Main Meta-Repo & Connected Workspaces

```text
N = 200
```

N = total self-loop steps budget that the agents will perform.

/goal Autonomously scan, plan, refactor, and fix all Enum, Trait, and Pattern Matching architectures across Go, TypeScript, PHP, Rust, and Python codebases, enforcing string-backed enums, `HasEnumHelpers` traits, exhaustive pattern matching, `*Type` suffixes, central `enums/` folder collocation, and strict relative Git paths until 100% green without stopping.

### Master Task Checklist (Atomic Numbered Steps)

1. [ ] /goal Phase 1 (Step A): Deeply scan the target codebase to inventory all architectural violations and anti-patterns.
2. [ ] /goal Phase 1 (Step B): Write the master audit specification in `.lovable/plans/pending/` with an exhaustive Violation Ledger.
3. [ ] /goal Phase 1 (Step C): Decompose the master plan into granular, atomic subtasks in `.lovable/plans/subtasks/`.
4. [ ] /goal Phase 1 (Step D): Verify or create the automated quality linter and register in `.lovable/ai-fix-scripts/index.md`.
5. [ ] /goal Phase 2 (Step A): Open each target file and perform surgical refactoring following authoritative guidelines.
6. [ ] /goal Phase 2 (Step B): Enforce <= 8–15 line function decomposition, single return types, and clean formatting.
7. [ ] /goal Phase 2 (Step C): Execute local linters to verify 0 remaining violations across all modified files.
8. [ ] /goal Phase 2 (Step D): Execute local CI quality gates via `python .lovable/ai-fix-scripts/03-cicd-local-runner.py` with exit code 0 (`exit 0`).
9. [ ] /learn Ingest `.lovable/memory/00-index.md` for project memory index and past learnings.
10. [ ] /learn Ingest `.lovable/strictly-avoid.md` for banned anti-patterns and strict constraints.
11. [ ] /learn Ingest `spec/02-coding-guidelines/00-canonical-size-tier.md` for canonical file and function size tiers.
12. [ ] /learn Ingest `spec/02-coding-guidelines/06-ai-optimization/01-anti-hallucination-rules.md` for hallucination prevention and micro-tasking.
13. [ ] /learn Ingest `spec/02-coding-guidelines/06-ai-optimization/05-citation-requirement.md` for strict relative path citation requirements.
14. [ ] /learn Ingest `spec/02-coding-guidelines/01-cross-language/02-constants-enums.md` for cross-language enum and constant architectures.
15. [ ] /learn Ingest `spec/02-coding-guidelines/05-php/` for PHP 8.1+ backed enums and helper traits.
16. [ ] /learn Ingest `.lovable/coding-guidelines/coding-guidelines.md` for master consolidated coding guidelines.
17. [ ] /goal Create or update agent rules in the repository if missing from agent memory.


```text
PHASE_1_STEPS = N / 2   (Steps 1 .. N/2: Scan Enums & Traits, Build Violation Ledger in .lovable/plans/pending/, Subtasks)
PHASE_2_STEPS = N / 2   (Steps N/2+1 .. N: Refactor Enums & Traits, Match Expressions, Run Typecheckers, Verify Local CI)
```

N, PHASE_1_STEPS, and PHASE_2_STEPS are read-only after initialization. Never modify them mid-execution.

---

## Dedicated Section: Multi-Language Enum & Trait Architecture

Enums represent a finite, closed set of valid domain values. Scattering loose string literals or bare integers destroys type safety. Every language must enforce strongly-typed enums with helper methods and pattern matching.

---

### 1. PHP 8.1+ String-Backed Enums & `HasEnumHelpers` Trait

In PHP 8.1+, all enums must be string-backed (`enum StatusType: string`) and use a standard `HasEnumHelpers` trait to provide `values()`, `names()`, `isValid()`, and `tryFromOrThrow()`.

```php
<?php

declare(strict_types=1);

namespace App\Enums;

use App\Exceptions\AppException;
use App\Enums\Traits\HasEnumHelpers;

enum OrderStatusType: string
{
    use HasEnumHelpers;

    case Pending    = 'pending';
    case Processing = 'processing';
    case Completed  = 'completed';
    case Cancelled  = 'cancelled';

    public function label(): string
    {
        return match ($this) {
            self::Pending    => 'Pending Payment',
            self::Processing => 'Processing Shipment',
            self::Completed  => 'Order Completed',
            self::Cancelled  => 'Order Cancelled',
        };
    }

    public function isTerminal(): bool
    {
        return match ($this) {
            self::Completed, self::Cancelled => true,
            self::Pending, self::Processing  => false,
        };
    }
}
```

```php
<?php

declare(strict_types=1);

namespace App\Enums\Traits;

use App\Exceptions\AppException;

trait HasEnumHelpers
{
    /** @return list<string> */
    public static function values(): array
    {
        return array_column(self::cases(), 'value');
    }

    /** @return list<string> */
    public static function names(): array
    {
        return array_column(self::cases(), 'name');
    }

    public static function isValid(string $value): bool
    {
        return self::tryFrom($value) !== null;
    }

    public static function fromOrThrow(string $value): static
    {
        $case = self::tryFrom($value);

        if ($case === null) {
            throw AppException::validation("Invalid enum value '$value' for " . static::class);
        }

        return $case;
    }
}
```

---

### 2. Rust Enums, Algebraic Data Types & Exhaustive Matching

In Rust, enums are first-class Algebraic Data Types. Use typed variants with payload data and exhaustive `match` branches:

```rust
// ✅ REQUIRED: Rust ADT Enum with *Type suffix and custom helper methods
#[derive(Debug, Clone, PartialEq, Eq)]
pub enum TaskStatusType {
    Pending,
    Running { progress_pct: u8 },
    Completed { duration_ms: u64 },
    Failed { error_code: String, message: String },
}

impl TaskStatusType {
    pub fn is_finished(&self) -> bool {
        match self {
            Self::Completed { .. } | Self::Failed { .. } => true,
            Self::Pending | Self::Running { .. } => false,
        }
    }

    pub fn status_name(&self) -> &'static str {
        match self {
            Self::Pending => "PENDING",
            Self::Running { .. } => "RUNNING",
            Self::Completed { .. } => "COMPLETED",
            Self::Failed { .. } => "FAILED",
        }
    }
}
```

---

### 3. Go Custom Type Enums & Stringers

In Go, enums are declared with custom named types ending in `Type`, grouped in `const` blocks, and placed in a dedicated `enums/` package:

```go
package enums

// OrderStatusType defines discrete order lifecycle states.
type OrderStatusType string

const (
    OrderStatusPending    OrderStatusType = "pending"
    OrderStatusProcessing OrderStatusType = "processing"
    OrderStatusCompleted  OrderStatusType = "completed"
    OrderStatusCancelled  OrderStatusType = "cancelled"
)

func (s OrderStatusType) IsValid() bool {
    switch s {
    case OrderStatusPending, OrderStatusProcessing, OrderStatusCompleted, OrderStatusCancelled:
        return true
    default:
        return false
    }
}

func (s OrderStatusType) IsTerminal() bool {
    return s == OrderStatusCompleted || s == OrderStatusCancelled
}
```

---

### 4. TypeScript `as const` Object Enums

```typescript
export const OrderStatusType = {
    Pending: 'pending',
    Processing: 'processing',
    Completed: 'completed',
    Cancelled: 'cancelled',
} as const;

export type OrderStatusType = (typeof OrderStatusType)[keyof typeof OrderStatusType];

export function isTerminalStatus(status: OrderStatusType): boolean {
    return status === OrderStatusType.Completed || status === OrderStatusType.Cancelled;
}
```

---

## 5. Phase 1 Violation Ledger Format

In Phase 1, you MUST generate `.lovable/plans/pending/XX-enums-and-traits-audit.md` containing the master inventory table:

```markdown
| Target File | Line | Identifier | Current Pattern | Language | Planned Refactoring | Status |
|---|:---:|---|---|---|---|:---:|
| `app/Models/Order.php` | 24 | `$status` | Loose string literal `'pending'` | PHP | Backed Enum `OrderStatusType` + `HasEnumHelpers` | PENDING |
| `src/task.rs` | 52 | `status_code: u8` | Numeric status code `0, 1, 2` | Rust | ADT Enum `TaskStatusType` with payload | PENDING |
| `pkg/api/order.go` | 18 | `Status string` | Raw unvalidated string | Go | Custom `enums.OrderStatusType` + `.IsValid()` | PENDING |
```

---

---

## Continuous 2-Phase Self-Loop & 2-Agent Concurrency Architecture

To guarantee full execution without stopping after planning mode, the master orchestrator MUST enforce this continuous 2-phase loop:

### 1. 2-Agent Concurrency & Strict `.lovable/` Bounding

- **2-Agent Limit (Max 2 Threads Each):** When dispatching work, spawn **at most 2 sub-agents concurrently**, with **no more than 2 threads per agent**.
- **Strict Folder Bounding (`.lovable/`):** Subagents can ONLY write planning files, subtasks, status reports, and logs inside `.lovable/` (`.lovable/plans/`, `.lovable/temp/active-locks.json`, `.lovable/memory/issues/`).
- **Context Diet:** Provide subagents with minimal instructions (e.g. "Read subtask file `.lovable/plans/subtasks/XX/01-task.md` and execute it"). Do not paste huge files into agent prompts.

### 2. Phase 1: Planning Mode & Subtask Generation (Steps 1 .. N/2)

- Spawn 2 planning subagents to scan the codebase for target guideline violations.
- Write the master architectural specification in `.lovable/plans/pending/XX-audit.md` with an exhaustive Violation Ledger table.
- Decompose the master plan into granular subtasks in `.lovable/plans/subtasks/XX/01-task.md`, `02-task.md`, etc.
- **MANDATORY AUTO-LOOP (DO NOT STOP):** Once Phase 1 planning completes, the master orchestrator **MUST NOT STOP or ask the user for confirmation**. It MUST immediately self-loop and transition directly into Phase 2 execution mode.

### 3. Phase 2: Execution Mode & Parallel Refactoring (Steps N/2+1 .. N)

- Spawn 2 execution subagents (max 2 threads each) to execute subtasks in parallel on disjoint files.
- Subagents refactor code following all coding guidelines (<= 8–15 line functions, single return types, universal `*AppError` wrapping, Unix LF line endings).
- Move completed subtasks from `.lovable/plans/subtasks/` to `.lovable/plans/completed/` and update `.lovable/plans/index.md`.
- **Failure Memory & Feedback Loop:** If a subagent fails:
  - Rollback dirty working tree and log error details to `.lovable/plans/last-failure.md` and `.lovable/memory/issues/XX-failure.md`.
  - The next subagent spawned MUST read the previous failure log first, record it as a pending memory task, and implement the necessary fix.
- Execute local linters and `python .lovable/ai-fix-scripts/03-cicd-local-runner.py` ensuring `exit 0` before concluding.

## Strict In-Repository Execution & `.lovable/` Bounding Mandate

> [!IMPORTANT]
> **STRICT IN-REPOSITORY EXECUTION & `.lovable/` STORAGE CONTRACT:**
>
> 1. **In-Codebase Execution Only:** Whenever a Python script (runner, autofixer, linter, test aggregator) is executed or created, it MUST be executed **strictly within the repository root** (current working directory), NEVER outside the codebase or against external arbitrary directories.
> 2. **Strict Folder Bounding (`.lovable/`):** All AI scripts, local runners, autofixers, helper utilities, memory issue logs, and planning files MUST be created inside the `.lovable/` folder:
>    - Python AI Scripts: `.lovable/ai-fix-scripts/` (e.g. `01-file-manipulator.py`, `02-guideline-autofixer.py`, `03-cicd-local-runner.py`, `04-relative-path-fixer.py`, `05-naming-autofixer.py`, `06-cli-help-auditor.py`).
>    - RCA & Issue Logs: `.lovable/memory/issues/` and `.lovable/cicd-issues/`.
>    - Execution Plans & Subtasks: `.lovable/plans/pending/`, `.lovable/plans/subtasks/`.
>    - Coding Guidelines Mirror: `.lovable/coding-guidelines/`.
> 3. **Worker Pool & Log Aggregation Architecture:** All local runners and test orchestrators must use a concurrent worker pool (2–3 workers via `ThreadPoolExecutor`), announce enqueued tasks upfront, show real-time progress, handle failures gracefully without cancelling sibling workers, and print a consolidated final summary with full stdout/stderr error logs for failed jobs.
> 4. **`force` Keyword Support:** If the user wrote `force`, `force rebuild`, or `force create` on top of the prompt or trigger: **ALWAYS recreate/regenerate the Python runner script from scratch**, regardless of whether the file already exists on disk.
> 5. **No External or Random File Creation:** NEVER write scripts, temporary test scripts, or scratch files to root, `/tmp`, global system paths, or outside the repository boundary.

---

## AI Fix Scripts Memory (Reusable Tooling)

- [ ] `/goal` **Reuse First:** I have rigorously scanned and `/learn`ed `.lovable/ai-fix-scripts/index.md` to check if a helper script already exists before writing any new temporary code.
- [ ] **Strict In-Repository Execution:** All Python scripts (`.lovable/ai-fix-scripts/*.py`) MUST be executed strictly within the codebase repository root, NEVER outside the codebase.
- [ ] **Strict .lovable/ Folder Storage:** All AI scripts, local runners, autofixers, and helper utilities MUST be created inside `.lovable/ai-fix-scripts/`. NEVER create scripts in root or external paths.
- [ ] **Automated Naming & Style Fixers:** Use `python .lovable/ai-fix-scripts/05-naming-autofixer.py` and `02-guideline-autofixer.py` to audit boolean prefixes and newlines.
- [ ] **Go Generate Sync:** If you modify Go constants, enums, or stringers, you MUST run `go generate ./...` in the relevant directory (e.g., `cd gitmap && go generate ./...`) and commit the resulting generated files to prevent CI drift.
- [ ] **Commit & Track:** All new helper scripts were written strictly to `.lovable/ai-fix-scripts/` and committed to Git for future reuse.
- [ ] **Index Documentation:** I have updated `.lovable/ai-fix-scripts/index.md` using sequential script naming. For every script, I have included a `<details>` collapsible tag explaining exactly why the script is there and what it does.

---

## Pre-Reply / Loop Checklist (Must Verify Every Loop Iteration)

- [ ] Git working tree is clean before new code changes.
- [ ] Sub-agents are actively assigned disjoint files verified against `.lovable/temp/active-locks.json`.
- [ ] Completed tasks were `mv`'d to `plans/completed/` and `plans/index.md` was updated.
- [ ] 3-strike rule respected: failed tasks cleanly rolled back and logged to `last-failure.md`.
- [ ] **Strict Relative Git Paths:** All file paths, markdown links, citations, and subtask references in plans, specs, and memory logs are strictly relative to the git repository root. Zero absolute paths (`D:\...`, `C:\...`) or `file:///` URIs.
- [ ] **`*Type` Suffix:** All enum type definitions end with `Type` (e.g., `OrderStatusType`).
- [ ] **Exhaustive Matching:** All `match`/`switch` expressions cover 100% of enum cases without unhandled branches.
- [ ] **PHP Backed Enums:** All PHP enums use string-backing and `HasEnumHelpers` trait.
- [ ] **Rust ADT Enums:** All Rust variants implement proper `match` arms and derive macros.
- [ ] **LF Line Endings (`\n`):** All files use Unix LF line endings. Zero CRLF (`\r\n`).
- [ ] **UTF-8 Encoding (No BOM):** All files encoded in UTF-8 without BOM.
- [ ] **Single Trailing Newline:** Every file ends with exactly one terminating newline (`\n`).
- [ ] **Blank Line Before `if`:** Exactly one blank line precedes every `if` statement (unless at the top of a block).
- [ ] **Blank Line After `}`:** Exactly one blank line follows every closing brace `}` (unless closing the enclosing block).
- [ ] **Blank Line Before `return`:** Exactly one blank line precedes `return` / `throw` in multi-line blocks.
- [ ] **Zero Nested `if`:** All conditionals flattened to depth 0 using guard clauses and early returns.
- [ ] **Function Sizing:** All functions <= 8 lines preferred (hard cap 15 lines).
- [ ] Local CI runner `python .lovable/ai-fix-scripts/03-cicd-local-runner.py` exited with code 0.

---

## Non-Negotiable Coding Guidelines Checklist (Auto-Reject on Violation)

/goal You MUST verify every item on this checklist before committing any code. If a subagent violated one of these rules, you must reject their work.

- [ ] Strict Relative Git Paths: All file paths, markdown links, citations, and subtask references in plans, specs, and memory logs are strictly relative to the git repository root. Zero absolute paths or `file:///` URIs.
- [ ] Master Guidelines: I have fully read and strictly enforced `spec/02-coding-guidelines/01-cross-language/02-constants-enums.md` and `.lovable/coding-guidelines/coding-guidelines.md`.
- [ ] Enum Suffix: Enforced `*Type` naming across all languages.
- [ ] LF Line Endings & UTF-8 (No BOM): Verified Unix LF and UTF-8 across all files.
- [ ] Blank Line Before `if`: Verified blank line before every `if` statement across all modified files.
- [ ] Blank Line After `}`: Verified blank line after every closing brace `}` followed by code.
- [ ] Blank Line Before `return`: Verified blank line before every `return`/`throw` in multi-line blocks.
- [ ] Zero Nested `if`: Zero nested `if` statements (depth > 1).


1. [ ] /learn the section as a /goal [AI Fix Scripts Memory](#ai-fix-scripts-memory)

- [ ] Action Summary: I have output a detailed `- [x]` checklist summarizing exactly what I accomplished this turn to prove I did not hallucinate.
