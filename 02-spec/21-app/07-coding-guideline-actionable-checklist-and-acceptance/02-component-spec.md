# Component Specification: Acceptance Criteria Registries & Verification Engine (AI Execution Prompt)

> **/goal** Standardize, unify, and automate the hierarchical acceptance criteria registry architecture and verification engine across all polyglot coding guidelines, integrating domain registries for Python (`12-python`) and Modern C++ (`13-cpp`) into the master registry and linking enforcement directly to `03-ai-scripts/05-guideline-autofixer.py` and `06-cicd-local-runner.py`.
> **/learn** Master the hierarchical criteria taxonomy (`AC-CG-[DOMAIN]-[NUM]`), the strict 1:1 mapping between acceptance criteria and authoritative specifications, the Given/When/Then behavioral verification contract, and the automated quality gate runners that enforce zero-drift compliance across polyglot repositories.

**Version:** 1.0.0
**Updated:** 2026-10-03
**Status:** Active
**AI Confidence:** Production-Ready
**Ambiguity:** None

---

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Verify every coding guideline domain contains a dedicated `97-acceptance-criteria.md` registry.
- [ ] `/learn` Ensure `12-python/97-acceptance-criteria.md` and `13-cpp/97-acceptance-criteria.md` map all domain criteria with full Given/When/Then contracts.
- [ ] `/goal` Verify the master registry (`02-spec/02-coding-guidelines/97-acceptance-criteria.md`) references all 12 domain registries without omission.
- [ ] `/learn` Document how `03-ai-scripts/05-guideline-autofixer.py` executes composite syntax, whitespace, and boolean audits.
- [ ] `/goal` Document how `03-ai-scripts/06-cicd-local-runner.py` coordinates concurrent quality gate checks and prevents regression drift.
- [ ] `/learn` Enforce 100% relative repository paths and zero absolute filesystem paths across all registry specifications.
- [ ] `/goal` Verify compliance by executing `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only` with `Expected: exit 0`.

. **CRITICAL AI INSTRUCTION:** This component specification governs the design, registry contracts, and verification pipelines for all testable acceptance criteria across `02-spec/02-coding-guidelines/`. Autonomous agents authoring or updating criteria MUST follow this schema.

---

## 1. Acceptance Criteria Architecture & System Overview

The coding guidelines specification framework utilizes a two-tier hierarchical registry model to guarantee 100% auditability, machine-verifiability, and traceability across all programming languages, styling rules, and architectural mandates.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│ MASTER ACCEPTANCE CRITERIA REGISTRY                                         │
│ File: 02-spec/02-coding-guidelines/97-acceptance-criteria.md                │
│ Functions as the top-level index, routing, and umbrella verification ledger  │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
                ┌──────────────────────┴──────────────────────┐
                ▼                                             ▼
┌───────────────────────────────┐             ┌───────────────────────────────┐
│ CORE & POLYGLOT REGISTRIES    │             │ SPECIALIZED & APP REGISTRIES  │
├───────────────────────────────┤             ├───────────────────────────────┤
│ • 01-cross-language/97-...    │             │ • 06-ai-optimization/97-...   │
│ • 02-typescript/97-...        │             │ • 06-cicd-integration/97-...  │
│ • 03-golang/97-...            │             │ • 08-fix-installers/97-...    │
│ • 04-php/97-...               │             │ • 08-file-folder-naming/97-...│
│ • 05-rust/97-...              │             │ • 11-security/97-...          │
│ • 07-csharp/97-...            │             │ • 21-app/readme.md            │
│ • 12-python/97-... (New)      │             │ • 22-app-issues/readme.md     │
│ • 13-cpp/97-... (New)         │             │ • 23-app-db/readme.md         │
└───────────────┬───────────────┘             └───────────────┬───────────────┘
                │                                             │
                └──────────────────────┬──────────────────────┘
                                       ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│ AUTOMATED VERIFICATION ENGINE                                               │
│ • 03-ai-scripts/05-guideline-autofixer.py (Whitespace & Boolean Linter)     │
│ • 03-ai-scripts/06-cicd-local-runner.py (Parallel CI Quality Gate Suite)    │
│ • Targeted Linter Scripts (Relative path verification, title parity)        │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 1.1 Master Registry Role (`97-acceptance-criteria.md`)

The master registry at `02-spec/02-coding-guidelines/97-acceptance-criteria.md` is the authoritative index consolidating all testable quality gates across the repository:
1. **Domain Routing:** Provides high-level summary tables referencing each domain-specific acceptance criteria registry via strictly relative markdown links.
2. **Cross-Domain Auditing:** Enables CI/CD pipelines and autonomous audit agents to verify compliance in a single pass without traversing disjoint directory trees.
3. **Repository Gates:** Defines the overarching repository acceptance gates (`AC-CG-STYLE-002` through `AC-CG-STYLE-005`) that govern global sizing limits, parameter count restrictions, and vertical line spacing.

### 1.2 Domain Registries Inventory

Every major domain within `02-spec/02-coding-guidelines/` maintains an isolated `97-acceptance-criteria.md` file that catalogs every single testable rule within its directory:

| Domain Directory | Registry Path | Scope & Focus |
|:---|:---|:---|
| `01-cross-language` | `01-cross-language/97-acceptance-criteria.md` | Boolean hygiene, code style, naming conventions, type safety, DRY, and complexity limits. |
| `02-typescript` | `02-typescript/97-acceptance-criteria.md` | Status enums, discriminated unions, `any` elimination, and ESLint automation. |
| `03-golang` | `03-golang/97-acceptance-criteria.md` | `appfault.AppError`, `Result[T]` wrappers, positive booleans, enums, and defer safety. |
| `04-php` | `04-php/97-acceptance-criteria.md` | PSR-4 autoloading, WordPress `$wpdb` prepared statements, typed arguments, and escaping. |
| `05-rust` | `05-rust/97-acceptance-criteria.md` | RFC 430 naming, `thiserror`/`anyhow`, Tokio cancellation, and zero unannotated `unsafe`. |
| `06-ai-optimization` | `06-ai-optimization/97-acceptance-criteria.md` | Anti-hallucination rules, source attribution, citation requirements, and memory lifecycles. |
| `06-cicd-integration` | `06-cicd-integration/97-acceptance-criteria.md` | SARIF output contracts, zero-storage CI/CD mandates, and fix repo installer automation. |
| `07-csharp` | `07-csharp/97-acceptance-criteria.md` | PascalCase naming, method parameter limits, structured exception handling, and immutability. |
| `08-file-folder-naming`| `08-file-folder-naming/97-acceptance-criteria.md`| Strictly lowercase kebab-case naming, zero uppercase characters, and sequence numbering. |
| `11-security` | `11-security/97-acceptance-criteria.md` | JWT token lifecycles, HttpOnly cookie storage, AES encryption, OWASP Top 10, and zero secrets. |
| `12-python` | `12-python/97-acceptance-criteria.md` | PEP-8 compliance, Black formatting, static type hints, Pydantic models, and specific exceptions. |
| `13-cpp` | `13-cpp/97-acceptance-criteria.md` | Modern C++20 baseline, concepts, smart pointers (`unique_ptr`/`shared_ptr`), and RAII. |

---

## 2. Standardizing Python & Modern C++ Registries

To complete polyglot coverage and eliminate gaps in the specification audit ledger, `12-python/97-acceptance-criteria.md` and `13-cpp/97-acceptance-criteria.md` establish authoritative registries for Python and C++.

### 2.1 Python Acceptance Criteria Registry (`12-python/97-acceptance-criteria.md`)

The Python registry must catalog all criteria in the `AC-CG-PY-` namespace:

```markdown
| ID | Title | Authoritative Specification | Verification Command |
|:---|:---|:---|:---|
| `AC-CG-PY-001` | Python Guidelines Directory Index Conformance | `12-python/readme.md` | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/12-python --check-only` |
| `AC-CG-PY-002` | Python Coding Standards Conformance | `12-python/02-standards.md` | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/12-python --check-only` |
| `AC-CG-PY-003` | Python Dynamic Enum & Array Constants Standard | `12-python/02-standards.md` | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/12-python --check-only` |
| `AC-CG-PY-004` | Python DRY Architecture & Engine Caching Conformance | `12-python/02-standards.md` | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/12-python --check-only` |
| `AC-CG-PY-REG-001` | Python Acceptance Criteria Registry Conformance | `12-python/97-acceptance-criteria.md` | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/12-python --check-only` |
```

#### Detailed Behavioral Criteria for Python

1. **AC-CG-PY-001: Python Guidelines Directory Index Conformance**
   - **Given** Python specifications and documentation within `12-python/`.
   - **When** Audited against Prompt Architect structure rules and relative path link integrity.
   - **Then** `12-python/readme.md` contains active `/goal` and `/learn` directives, an actionable checklist, accurate file references, and zero absolute paths.

2. **AC-CG-PY-002: Python Coding Standards Conformance**
   - **Given** Python source files and scripts across the codebase.
   - **When** Validated for static type hinting, data validation models, and PEP-8 compliance.
   - **Then** All public functions declare explicit type hints (banning raw `Any`), data schemas employ `pydantic` or `@dataclass`, line length respects the 100-character ceiling, and bare `except:` blocks are strictly absent.

3. **AC-CG-PY-003: Python Dynamic Enum & Array Constants Standard**
   - **Given** Python scripts and modules defining configurable options or state machines.
   - **When** Inspected for hardcoded string constants and Cartesian string permutations.
   - **Then** Identifiers use `Enum` or `StrEnum` types, dynamic array builders eliminate repetitive string combinations, and magic literals are eliminated.

4. **AC-CG-PY-004: Python DRY Architecture & Engine Caching Conformance**
   - **Given** Python CI/CD, automation, or linting scripts in `03-ai-scripts/`.
   - **When** Scanned for duplicate utility implementations or unshared helpers.
   - **Then** Scripts import shared logic from `03-ai-scripts/02-shared-engine.py`, enforce idempotent caching where applicable, and maintain unified exit codes.

5. **AC-CG-PY-REG-001: Python Acceptance Criteria Registry Conformance**
   - **Given** The Python criteria registry at `12-python/97-acceptance-criteria.md`.
   - **When** Evaluated for taxonomic integrity and traceability.
   - **Then** All listed criteria map 1:1 to specification files in `12-python/`, employ strict `Given/When/Then` structures, and specify executable verification commands.

---

### 2.2 Modern C++ Acceptance Criteria Registry (`13-cpp/97-acceptance-criteria.md`)

The C++ registry must catalog all criteria in the `AC-CG-CPP-` namespace:

```markdown
| ID | Title | Authoritative Specification | Verification Command |
|:---|:---|:---|:---|
| `AC-CG-CPP-001` | Modern C++ Guidelines Directory Index Conformance | `13-cpp/readme.md` | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/13-cpp --check-only` |
| `AC-CG-CPP-002` | Modern C++ Standards Conformance | `13-cpp/02-standards.md` | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/13-cpp --check-only` |
| `AC-CG-CPP-003` | C++ Memory Safety & RAII Resource Management | `13-cpp/02-standards.md` | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/13-cpp --check-only` |
| `AC-CG-CPP-004` | C++ FFI Boundary Exception Safety & Standard Types | `13-cpp/02-standards.md` | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/13-cpp --check-only` |
| `AC-CG-CPP-REG-001` | Modern C++ Acceptance Criteria Registry Conformance | `13-cpp/97-acceptance-criteria.md` | `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/13-cpp --check-only` |
```

#### Detailed Behavioral Criteria for Modern C++

1. **AC-CG-CPP-001: Modern C++ Guidelines Directory Index Conformance**
   - **Given** C++ specifications and documentation within `13-cpp/`.
   - **When** Audited against Prompt Architect structure rules and link validity.
   - **Then** `13-cpp/readme.md` features active execution prompts, an actionable agent checklist, accurate relative links, and zero absolute paths.

2. **AC-CG-CPP-002: Modern C++ Standards Conformance**
   - **Given** C++ header and implementation files across repositories.
   - **When** Audited for language dialect baseline and template constraints.
   - **Then** Code conforms to the C++20 standard or newer, concepts are utilized instead of raw SFINAE / `enable_if`, and naming rules (`PascalCase` for types/enums, `snake_case` for functions/variables) are strictly followed.

3. **AC-CG-CPP-003: C++ Memory Safety & RAII Resource Management**
   - **Given** C++ classes, resource handles, and heap-allocated objects.
   - **When** Inspected for memory management and ownership patterns.
   - **Then** Manual `new` and `delete` invocations are strictly absent, ownership is expressed via `std::unique_ptr` or `std::shared_ptr`, and classes adhere to the Rule of Zero (or complete Rule of Five).

4. **AC-CG-CPP-004: C++ FFI Boundary Exception Safety & Standard Types**
   - **Given** C++ modules exposing C-compatible FFI or receiving foreign function calls.
   - **When** Audited for exception leakage and ABI stability.
   - **Then** Zero exceptions escape C++ boundaries into foreign runtime contexts (enforced via `noexcept` and `try/catch` wrapping), and error codes/enums communicate status across FFI layers.

5. **AC-CG-CPP-REG-001: Modern C++ Acceptance Criteria Registry Conformance**
   - **Given** The C++ criteria registry at `13-cpp/97-acceptance-criteria.md`.
   - **When** Evaluated for taxonomic integrity and traceability.
   - **Then** Every entry links 1:1 to authoritative specifications in `13-cpp/`, features explicit `Given/When/Then` definitions, and defines functional verification commands.

---

### 2.3 Master Registry Synchronization Contract

To ensure complete discoverability, `02-spec/02-coding-guidelines/97-acceptance-criteria.md` MUST be synchronized with:

1. **Section `## AC-12: Python Standards Registry`:**
   - Markdown summary table citing `AC-CG-PY-001`, `AC-CG-PY-002`, `AC-CG-PY-003`, `AC-CG-PY-004`, and `AC-CG-PY-REG-001`.
   - Relative pointer to `02-spec/02-coding-guidelines/12-python/97-acceptance-criteria.md`.

2. **Section `## AC-13: Modern C++ Standards Registry`:**
   - Markdown summary table citing `AC-CG-CPP-001`, `AC-CG-CPP-002`, `AC-CG-CPP-003`, `AC-CG-CPP-004`, and `AC-CG-CPP-REG-001`.
   - Relative pointer to `02-spec/02-coding-guidelines/13-cpp/97-acceptance-criteria.md`.

3. **Cross-References Footer Expansion:**
   - Link to `Python Standards` (`./12-python/readme.md`).
   - Link to `Python Acceptance Criteria Registry` (`./12-python/97-acceptance-criteria.md`).
   - Link to `Modern C++ Standards` (`./13-cpp/readme.md`).
   - Link to `Modern C++ Acceptance Criteria Registry` (`./13-cpp/97-acceptance-criteria.md`).

---

## 3. Automated Verification Engine & Enforcement

The acceptance criteria registries are not passive lists; they are actively checked, validated, and guarded against drift by the repository's automated verification engine.

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                       AUTOMATED VERIFICATION PIPELINE                       │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
                ┌──────────────────────┴──────────────────────┐
                ▼                                             ▼
┌───────────────────────────────┐             ┌───────────────────────────────┐
│ 05-guideline-autofixer.py     │             │ 06-cicd-local-runner.py       │
├───────────────────────────────┤             ├───────────────────────────────┤
│ • Fast composite linter runner│             │ • Parallel multi-threaded     │
│ • 04-newline-fixer (spacing)  │             │   quality gate executor       │
│ • 08-naming-autofixer (bool)  │             │ • Concurrent job matrix (Pool)│
│ • --check-only mode for CI/CD │             │ • Relative path verification  │
│ • Sub-25ms audit latency      │             │ • JSON summary & silent ticks │
└───────────────┬───────────────┘             └───────────────┬───────────────┘
                │                                             │
                └──────────────────────┬──────────────────────┘
                                       ▼
┌─────────────────────────────────────────────────────────────────────────────┐
│ QUALITY GATE OUTCOME: Exit Code 0 (Pass) or Exit Code 1 (Drift Detected)    │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 3.1 Guideline Autofixer (`03-ai-scripts/05-guideline-autofixer.py`)

`03-ai-scripts/05-guideline-autofixer.py` serves as the primary tactical gate for coding guideline compliance:
- **Composite Execution:** Seamlessly invokes both `04-newline-fixer.py` and `08-naming-autofixer.py` using shared types from `02-shared-engine.py`.
- **Newline & Vertical Spacing Rules:** Enforces mandatory blank lines before `if` statements, after closing braces `}`, before `return` statements, and around parameter struct initializations.
- **Boolean Integrity:** Verifies affirmative boolean naming (`is*`, `has*` only) and bans explicit equality evaluations against `true`/`false` (`if is_valid == True:`).
- **Execution Modes:**
  - *Fix Mode (Default):* Modifies target files directly to bring them into immediate compliance.
  - *Audit Mode (`--check-only`):* Scans the target folder without modifying files. Returns exit code `0` if all files comply; returns a non-zero exit code if violations exist.

### 3.2 CI/CD Local Runner (`03-ai-scripts/06-cicd-local-runner.py`)

`03-ai-scripts/06-cicd-local-runner.py` is the comprehensive quality gate orchestrator:
- **Parallel Worker Pool:** Utilizes Python's `ThreadPoolExecutor` (capped at CPU cores or 8 workers) to execute 20+ CI quality gates concurrently in under 3 seconds.
- **Drift Prevention:** Executes cross-cutting verifications including:
  1. `check-relative-paths.py`: Guarantees zero absolute filesystem paths or `file:///` URIs across specs and code.
  2. `check-sequence-titles.py`: Validates sequence numbers, kebab-case naming, and header alignment.
  3. `guideline-autofixer.py --check-only`: Verifies that no styling or boolean violations were introduced.
- **Diagnostic Transparency:** On failure, captures stderr, exit code, and execution duration, outputting structured diagnostics without uploading routine artifacts to CI storage.

---

## 4. Verification & Acceptance Criteria

### AC-CG-REG-MASTER-001: Master Registry Completeness & Polyglot Parity

**Given** The master acceptance criteria registry at `02-spec/02-coding-guidelines/97-acceptance-criteria.md`.
**When** Audited for coverage across all guideline subdirectories in `02-spec/02-coding-guidelines/`.
**Then** All twelve domain registries (Cross-Language, TypeScript, Golang, PHP, Rust, AI Optimization, CI/CD Integration, C#, File Naming, Security, Python, and C++) are indexed with accurate relative links and summary tables.

---

### AC-CG-REG-PY-001: Python Acceptance Criteria Registry Standardization

**Given** The Python guideline directory at `02-spec/02-coding-guidelines/12-python/`.
**When** Checked for registry presence and criterion completeness.
**Then** `12-python/97-acceptance-criteria.md` exists, catalogs criteria `AC-CG-PY-001` through `AC-CG-PY-REG-001` with explicit `Given/When/Then` blocks, and defines runnable verification commands.

---

### AC-CG-REG-CPP-001: Modern C++ Acceptance Criteria Registry Standardization

**Given** The C++ guideline directory at `02-spec/02-coding-guidelines/13-cpp/`.
**When** Checked for registry presence and criterion completeness.
**Then** `13-cpp/97-acceptance-criteria.md` exists, catalogs criteria `AC-CG-CPP-001` through `AC-CG-CPP-REG-001` with explicit `Given/When/Then` blocks, and defines runnable verification commands.

---

### AC-CG-REG-TOOL-001: Automated Guideline Autofixer & CI/CD Runner Verification

**Given** The complete set of specification files across `02-spec/02-coding-guidelines/`.
**When** The automated verification engine is executed in audit mode.
**Then** The guideline autofixer passes with zero errors, zero formatting violations, and exit code 0.

**Verification command:**
```bash
python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only
```
**Expected:** exit 0. Zero violations detected.

---

## 5. Related Specifications & References

- [Master Acceptance Criteria Registry](../../02-coding-guidelines/97-acceptance-criteria.md) — Master quality gate index
- [Coding Guidelines Overview](../../02-coding-guidelines/readme.md) — Core coding guideline directory
- [Python Coding Guidelines](../../02-coding-guidelines/12-python/readme.md) — Python standards and models
- [Modern C++ Coding Guidelines](../../02-coding-guidelines/13-cpp/readme.md) — C++20 and RAII standards
- [Subtask Plan: Harmonize Polyglot Registries](../../../.ai-memory/plans/subtasks/07-coding-guideline-actionable-checklist-and-acceptance/02-harmonize-polyglot-registries.md) — Step-by-step rollout plan
