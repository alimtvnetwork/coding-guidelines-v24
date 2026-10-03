# Architecture Specification: Coding Guidelines Actionable Checklists & Testable Acceptance Criteria Standard (AI Execution Prompt)

> **/goal** Transform all coding guidelines across the meta-repository from static, passive reference documentation into authoritative, active AI execution prompts equipped with standardized agent checklists at the top and formal, testable acceptance criteria at the bottom.
> **/learn** Master the structural standards of Prompt Architect specifications, enforce the mandatory top-level prompt block (`> **/goal**` and `> **/learn**`), the actionable CI/CD and agent checklist (`## 🎯 Actionable CI/CD & Agent Checklist`), and testable verification criteria (`Given / When / Then` with automated verification commands).

**Version:** 1.0.0
**Updated:** 2026-10-03
**Status:** Active
**AI Confidence:** Production-Ready
**Ambiguity:** None

---

## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` Verify every coding guideline document begins with the standard AI action prompt header (`> **/goal**` and `> **/learn**`).
- [ ] `/learn` Verify every coding guideline document contains an immediately actionable `## 🎯 Actionable CI/CD & Agent Checklist` covering core rules with `/goal` and `/learn` directives.
- [ ] `/goal` Verify every coding guideline document ends with a testable `## Verification & Acceptance Criteria` section using `Given / When / Then` and verification commands (`python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only`).
- [ ] `/learn` Validate 100% relative paths and zero absolute filesystem paths or `file:///` URIs across all specifications via relative path auditing.

. **CRITICAL AI INSTRUCTION:** This specification defines the mandatory structural contract for all coding guidelines in `02-spec/02-coding-guidelines/`. AI agents authoring or modifying guideline files MUST enforce this exact layout.

---

## 1. User Request (Verbatim)

```text
# High Priority Instruction

Can you please go into the coding guideline and see the coding guideline, how it is written? It's not written as an action or prompt on every file. Try to have from the README.md file inside the coding guideline folder, like a spec and then coding guideline. Try to have every one of the file try to mention as actionable items that the AI agent must follow. Okay. As a checklist. And what is acceptance criteria at the end that needs to be mentioned on each one of the specs for coding guideline inside, like the style guideline, Boolean guideline, naming convention, things like that. Okay. Do you understand? Can you please follow through? Do you have any question and confusion?

# Actionable Items Must Follow Non-Negotiable

1. Review the coding guidelines in the README.md file within the coding guideline folder.
2. Create a specification and coding guideline document.
3. Ensure each file includes actionable items for AI agents as a checklist.
4. Define acceptance criteria for each spec, including style guidelines, Boolean guidelines, and naming conventions.

Must follow and spawn agent using

@[.agents/skills/execute-parent-task-with-n-steps-v6]

## Additional Instructions

learn /learn if you have to learn something and /plan stuff before working please./plan
```

---

## 2. Motivation & Architectural Deficiency

Historically, files in `02-spec/02-coding-guidelines/` were authored as passive reference texts. While rich in explanation and code examples, this structure presented major deficiencies when consumed by autonomous AI agents:

1. **Lack of Immediate Actionability:** An AI agent scanning a file had to read hundreds of lines of prose to deduce the mandatory operational constraints.
2. **Missing Quality Gates:** Individual guideline files lacked unambiguous, testable acceptance criteria at the end, making automated verification subjective and difficult to evaluate in CI/CD.
3. **Inconsistency Across Modules:** While top-level readmes (such as `02-spec/02-coding-guidelines/readme.md`) introduced high-level checklists, individual guideline chapters, subfolder readmes, and changelogs lacked standardized prompt headers and actionable checklists.
4. **Subjective Completion Boundaries:** Without explicit Gherkin scenarios (`Given / When / Then`) and zero-exit command invocations, agents could declare completion prematurely without executing objective verification scripts.

---

## 3. The 4-Part Anatomy of Coding Guidelines

Every specification and guideline file in `02-spec/02-coding-guidelines/` MUST strictly adhere to the following 4-part architectural anatomy:

```
┌────────────────────────────────────────────────────────┐
│ 1. AI Execution Prompt Header                          │
│    - Title with (AI Execution Prompt) suffix           │
│    - > **/goal** [Concrete operational goal]           │
│    - > **/learn** [Key mental models & anti-patterns]  │
├────────────────────────────────────────────────────────┤
│ 2. Actionable CI/CD & Agent Checklist                  │
│    - ## 🎯 Actionable CI/CD & Agent Checklist          │
│    - Checkboxes with /goal and /learn directives       │
│    - . CRITICAL AI INSTRUCTION directive block         │
├────────────────────────────────────────────────────────┤
│ 3. Core Specification Body                             │
│    - Metadata block (Version, Updated, Status, etc.)   │
│    - Numbered rules, architectural explanations        │
│    - ❌ FORBIDDEN vs ✅ REQUIRED code examples        │
├────────────────────────────────────────────────────────┤
│ 4. Verification & Acceptance Criteria                  │
│    - ## Verification & Acceptance Criteria             │
│    - AC-CG-[CAT]-[NNN] with Given / When / Then        │
│    - Verification command and expected exit code 0     │
└────────────────────────────────────────────────────────┘
```

### 3.1 Part 1: AI Execution Prompt Header

Every file must start with a markdown Level 1 heading containing `(AI Execution Prompt)` suffix, immediately followed by the blockquote directives:

```markdown
# [Guideline Title] (AI Execution Prompt)

> **/goal** [Concise actionable goal statement defining the exact standard to enforce]
> **/learn** [Key references, anti-patterns, and mental models required before applying changes]
```

- `> **/goal**`: Directs the AI agent to focus on the immediate objective, files to modify, and standards to enforce.
- `> **/learn**`: Directs the AI agent to internalize anti-patterns, precedence rules, and structural boundaries before modifying code.

### 3.2 Part 2: Actionable CI/CD & Agent Checklist

Immediately after the header (or following the metadata block), the document MUST declare an actionable checklist containing at least 4 checkboxes prefixed with `/goal` and `/learn`:

```markdown
## 🎯 Actionable CI/CD & Agent Checklist

- [ ] `/goal` [Mandatory rule check 1 — core positive requirement]
- [ ] `/learn` [Anti-pattern check 2 — forbidden patterns and mental models]
- [ ] `/goal` [Mandatory rule check 3 — code styling / structural requirement]
- [ ] `/learn` [Verification / compliance check 4 — testable gate / linter verification]

. **CRITICAL AI INSTRUCTION:** This specification is an active AI execution directive. All code generated or modified must strictly follow the rules below.
```

### 3.3 Part 3: Core Specification Body

The body of the specification provides detailed explanations, formal numbered rules (`R1`, `R2`, `R3`...), and paired code examples demonstrating both the anti-pattern and the compliant pattern:

```markdown
### R[N]: [Rule Title]

- [Clear operational description of the rule]
- [Specific edge cases, banned patterns, or mandatory helpers]

#### ❌ FORBIDDEN: [Anti-pattern description]

```language
// Bad code example illustrating the violation
```

#### ✅ REQUIRED: [Compliant pattern description]

```language
// Good code example illustrating compliant implementation
```
```

### 3.4 Part 4: Verification & Acceptance Criteria

Every file MUST conclude with formal, testable acceptance criteria using canonical identifiers and Gherkin formatting:

```markdown
## Verification & Acceptance Criteria

### AC-CG-[CATEGORY]-[NUM]: [Descriptive Title]

**Given** [Precondition or codebase context]
**When** [Target file or component is analyzed by linters or test scripts]
**Then** [Expected deterministic outcome: zero violations, exit code 0]

**Verification command:**
```bash
python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only
```
**Expected:** exit 0. Zero violations detected.
```

---

## 4. Specific Guideline Rules

### 4.1 Style Guidelines

Code style rules enforce readability, visual parsing ease, and compact function boundaries across all languages.

#### Rule R4: Mandatory Blank Line Before `return`

Every `return` statement MUST be preceded by a blank line, unless it is the very first line of a block (or a guard clause immediately following an opening brace).

##### ❌ FORBIDDEN: Return jammed against previous statement

```go
func calculateSubtotal(items []Item) int {
    var total int
    for _, item := range items {
        total += item.Price
    }
    return total // ❌ Violation: missing blank line before return
}
```

##### ✅ REQUIRED: Blank line preceding return

```go
func calculateSubtotal(items []Item) int {
    var total int
    for _, item := range items {
        total += item.Price
    }

    return total
}
```

#### Rule R5: Mandatory Blank Line After Closing Brace `}`

Every closing brace `}` MUST be followed by a blank line, unless it is immediately followed by another closing brace `}` or an `else` block (in languages requiring `} else {`).

##### ❌ FORBIDDEN: Closing brace without vertical gap

```typescript
export function processBatch(records: readonly RecordItem[]): void {
  for (const record of records) {
    applyTransform(record);
  }
  dispatchNotification("batch completed"); // ❌ Violation: missing blank line after }
}
```

##### ✅ REQUIRED: Vertical blank line after closing brace

```typescript
export function processBatch(records: readonly RecordItem[]): void {
  for (const record of records) {
    applyTransform(record);
  }

  dispatchNotification("batch completed");
}
```

#### Rule R6: Function Length Cap (≤ 15 Lines)

Every function MUST be concise and bounded to at most 15 lines of executable code. Monolithic functions exceeding 15 lines MUST be decomposed into private, focused helper functions.

##### ❌ FORBIDDEN: Monolithic 30-line function doing everything

```go
// ❌ Violation: 30+ lines mixing HTTP parsing, validation, DB query, and formatting
func HandleUserRegistration(w http.ResponseWriter, r *http.Request) {
    // 30 lines of sequential procedural code
}
```

##### ✅ REQUIRED: Decomposed into atomic helpers

```go
func HandleUserRegistration(w http.ResponseWriter, r *http.Request) {
    params, err := parseRegistrationParams(r)
    if err != nil {
        writeError(w, err)
        return
    }

    result, err := registerUserAccount(params)
    if err != nil {
        writeError(w, err)
        return
    }

    writeJSONResponse(w, http.StatusCreated, result)
}
```

#### Rule R7: Parameter Bloat Elimination (Max 3 Parameters)

Functions MUST NOT accept more than 3 loose parameters. When 4 or more arguments are needed, group them into a dedicated parameter struct (`*Params` or `Options`).

##### ❌ FORBIDDEN: Loose parameter explosion (>3 parameters)

```typescript
// ❌ Violation: 5 loose positional parameters
function createUserProfile(
  name: string,
  email: string,
  role: string,
  isActive: boolean,
  tenantId: string
): UserProfile { ... }
```

##### ✅ REQUIRED: Parameter struct encapsulation

```typescript
interface CreateUserProfileParams {
  readonly name: string;
  readonly email: string;
  readonly role: string;
  readonly isActive: boolean;
  readonly tenantId: string;
}

function createUserProfile(params: CreateUserProfileParams): UserProfile {
  return buildUserProfileRecord(params);
}
```

---

### 4.2 Boolean Guidelines

Boolean hygiene rules eradicate logical ambiguity, double negatives, and cognitive overhead.

#### Rule B1: Affirmative `is` and `has` Prefixes Only

All boolean variables, properties, functions, and flags MUST use affirmative prefixes: `is` or `has`. The prefixes `can`, `should`, `was`, `did`, and negative naming (`isNotReady`, `disableFeature`) are **TOTALLY BANNED**.

##### ❌ FORBIDDEN: Negative or non-standard boolean prefixes

```go
var isNotActive bool     // ❌ Violation: negative naming
var shouldProceed bool   // ❌ Violation: banned prefix 'should'
var canExecute bool      // ❌ Violation: banned prefix 'can'
```

##### ✅ REQUIRED: Affirmative `is` / `has` naming

```go
var isActive bool
var hasPermission bool
var isReady bool
```

#### Rule B2: Zero Explicit True Checks (TOTAL BAN on `== true` / `=== true`)

NEVER evaluate a boolean explicitly against `true` (e.g. `if isReady == true` or `if hasToken === true`). Positive booleans MUST ALWAYS be evaluated implicitly.

##### ❌ FORBIDDEN: Explicit true comparison

```go
if isValid == true { // ❌ Violation: explicit comparison to true
    proceed()
}
```

##### ✅ REQUIRED: Implicit boolean evaluation

```go
if isValid {
    proceed()
}
```

#### Rule B3: Zero Mixed Polarity (No `isA && !isB`)

NEVER combine a positive boolean check and a negative boolean check in the same `if` condition (e.g. `if isReady && !isBlocked`). Mixed polarity creates high cognitive load and subtle bugs. Decompose mixed conditions into separate discrete guard clauses or positive composite flags.

##### ❌ FORBIDDEN: Mixed polarity condition

```typescript
// ❌ Violation: positive 'isAvailable' mixed with negative '!isPending'
if (isAvailable && !isPending) {
  processItem();
}
```

##### ✅ REQUIRED: Discrete guard clauses or positive composite helper

```typescript
if (!isAvailable) {
  return;
}

if (isPending) {
  return;
}

processItem();
```

---

### 4.3 Naming Conventions

#### Rule N1: PascalCase for Types, Schemas, and Models

All struct names, interface names, type aliases, database table schemas, and model classes MUST strictly use `PascalCase`.

##### ❌ FORBIDDEN: snake_case or camelCase type declarations

```go
type user_profile struct { // ❌ Violation: snake_case struct
    account_id string
}
```

##### ✅ REQUIRED: PascalCase type declarations

```go
type UserProfile struct {
    AccountID string
}
```

#### Rule N2: Semantic, Intention-Revealing Identifiers

Variable and parameter names MUST reveal their semantic intent and domain units. Suffix time-based variables with their unit (`TimeoutSeconds`, `DurationMs`).

##### ❌ FORBIDDEN: Vague and unitless names

```go
var t int      // ❌ Violation: single letter
var timeout int // ❌ Violation: unit is ambiguous (ms? seconds? minutes?)
```

##### ✅ REQUIRED: Intention-revealing with units

```go
var timeoutSeconds int
var requestDurationMs int64
```

#### Rule N3: Zero Generic Garbage Identifiers (TOTAL BAN)

Identifiers like `data`, `info`, `temp`, `obj`, `res`, `item`, `val`, `foo`, `bar` are **TOTALLY BANNED**. Always use specific domain terms (`userPayload`, `auditRecord`, `fileDescriptor`).

##### ❌ FORBIDDEN: Generic garbage variable names

```go
func parseData(data map[string]any) (any, error) { // ❌ Violation: 'data' and generic return
    item := data["info"]                          // ❌ Violation: 'item' and 'info'
    return item, nil
}
```

##### ✅ REQUIRED: Explicit domain identifiers

```go
func parseUserConfiguration(rawConfigMap map[string]any) (UserConfig, error) {
    userSettingValue, hasSetting := rawConfigMap["accountSetting"]
    if !hasSetting {
        return UserConfig{}, errMissingSetting
    }

    return buildUserConfig(userSettingValue)
}
```

---

## 5. Acceptance Criteria Taxonomy & Categorization

To maintain full traceability across all guidelines, acceptance criteria are categorized with standard prefixes:

| Category Prefix | Domain / Scope | Target Rule Focus |
|:---|:---|:---|
| `AC-CG-STYLE-` | Code style, spacing, sizing, braces | Blank line before return (R4), after `}` (R5), function length ≤15 lines, max 3 params |
| `AC-CG-BOOL-` | Boolean logic, polarity, naming | Affirmative `is`/`has` prefixes, zero `== true`, zero mixed polarity |
| `AC-CG-NAME-` | Naming conventions, semantic clarity | PascalCase types/schemas, semantic unit-tagged names, zero generic garbage |
| `AC-CG-TYPE-` | Type safety, immutability, param structs | Parameter structs, immutability, casting elimination, Result wrappers |
| `AC-CG-ARCH-` | Architecture, DRY, complexity | Cyclomatic complexity ≤ 10, DRY extraction, static analysis |

---

## Verification & Acceptance Criteria

### AC-CG-SPEC-001: Coding Guidelines Architecture Specification Conformance

**Given** The coding guidelines standard specification in `02-spec/21-app/07-coding-guideline-actionable-checklist-and-acceptance/01-architecture-spec.md`.
**When** Audited by repository linters and guideline autofixers.
**Then** The file strictly includes the AI Execution Prompt header (`> **/goal**` and `> **/learn**`), the actionable checklist (`## 🎯 Actionable CI/CD & Agent Checklist`), the verbatim user prompt, complete 4-part anatomy, detailed style, boolean, and naming rules, and concludes with automated testable criteria.

**Verification command:**
```bash
python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines --check-only
```
**Expected:** exit 0. Zero violations detected.

---

## 7. Related Specifications

- [`02-spec/02-coding-guidelines/readme.md`](../../02-coding-guidelines/readme.md) — Root coding guidelines index
- [`02-spec/02-coding-guidelines/01-cross-language/readme.md`](../../02-coding-guidelines/01-cross-language/readme.md) — Cross-language coding guidelines index
- [`.ai-memory/plans/subtasks/07-coding-guideline-actionable-checklist-and-acceptance/01-standardize-changelog-checklists.md`](../../../.ai-memory/plans/subtasks/07-coding-guideline-actionable-checklist-and-acceptance/01-standardize-changelog-checklists.md) — Subtask 01: Standardize Changelog Checklists
