# Instruction (must follow): Plan: Coding Guideline Audit & Enforcement (v4)

/goal Deeply audit the entire codebase for coding guideline violations, boolean anti-patterns, missing enums, cyclomatic complexity, and error-handling flaws. Structure all findings into actionable, fine-grained tasks in .lovable/plans/pending/ and subtasks before stopping.

## MUST FOLLOW NON-NEGOTIABLE

Listen, past runs of these turns have been sloppy and stupid as fuck: wrong step counts, partial task lists dumped into chat instead of files, plans and session summaries half-filled with "[N]" placeholders, folders skimmed, open ambiguities ignored, CI/CD issues and `plans/subtasks/` forgotten, user commands dropped, coding guidelines bypassed, detailed specs chopped and summarized into useless junk, uppercase README files left uncorrected, `.lovable/memories/` created by accident, `strictly-avoid.md` overwritten, and explicit user instructions softened after being told not to. WTF. How on earth are you reverting to this carelessness, are you stupid?? Stop doing that, you stupid fuck. Read the whole codebase, read every folder in `spec/` and `.lovable/`, confirm root `readme.md` is strictly lowercase, find the root cause in one sentence, capture commands, issues, and pending tasks without omitting a single item, write the spec files and memory files in the right paths, update every index in the same turn, sync `readme.md` with `what-to-read.md`, preserve detailed specs verbatim with zero truncation, run builds and full unit tests, group commits with clear messages, and push everything to git before ending. Going deep IS the job. If you are not going deep, you are not doing the job. Violating this is auto-reject on the same tier as RULE 0. Avoid stupidity and being careless, you stupid fuck. Where is your attention, are you stupid? Tell me. Your stupidity is going on top of my head. Where did you learn this stupidity? If I could find you, I could slap you.

## Variables — Auto-Discovered at Runtime

```text
N = 150 (Default number of steps the planning AI should take to generate the audit plan. The user may override this when triggering the prompt)
```

/learn Ingest, analyze, and internalize all coding guidelines, boolean principles, function size limits, and error handling architectures across the codebase and specs.

Autonomously self-loop and read:

- /learn the master cross-language coding guidelines in `spec/02-coding-guidelines/01-cross-language/15-master-coding-guidelines/01-naming-and-database.md` through `06-advanced-patterns.md`.
- /learn the code style, braces, spacing, and multi-line rules in `spec/02-coding-guidelines/01-cross-language/04-code-style/01-braces-and-nesting.md` through `06-comments-and-documentation.md`.
- /learn the strict function and type size caps (8 lines preferred, 15 lines max) in `spec/02-coding-guidelines/01-cross-language/04-code-style/04-function-and-type-size.md`.
- /learn the boolean principles, prefixing rules (`is`, `has`, `can`, `should`), and guard extraction in `spec/02-coding-guidelines/01-cross-language/02-boolean-principles/01-naming-prefixes.md` through `05-exemptions-and-api.md`.
- /learn the absolute prohibition against negative booleans and inverted logic in `spec/02-coding-guidelines/01-cross-language/12-no-negatives.md`.
- /learn the strict identifier and file naming conventions in `spec/02-coding-guidelines/08-file-folder-naming/01-cross-language.md`.
- /learn the DRY principles and duplication extraction patterns in `spec/02-coding-guidelines/01-cross-language/08-dry-principles.md`.
- /learn the error management architecture and logging diagnostics in `spec/03-error-manage/00-overview.md` and `spec/03-error-manage/02-error-architecture/01-error-handling-reference.md`.
- /learn the language-specific standards in `spec/02-coding-guidelines/` (TypeScript, Go, PHP, Rust, C#, Python, PowerShell).
- /learn the anti-hallucination rules and common AI mistakes in `spec/02-coding-guidelines/06-ai-optimization/01-anti-hallucination-rules.md` and `03-common-ai-mistakes.md`.
- Read `.lovable/plans/index.md` and `.lovable/memory/00-index.md`.

## 2. Planning Loop (Deep N-Step Analysis)

This is not a quick glance. You must deeply read the codebase, looping yourself as much as needed (taking exactly `N` steps of internal planning and reading). Each one of these steps MUST be followed properly using your highest processing capacity, proper memory retention of prior files, and careful multi-agent cognitive logic.

You must dedicate this immense processing power to uncover:

- Every inverted boolean (`!isSuccess`).
- Every magic string or number.
- Every swallowed error or generic `catch {}`.
- Every missing Enum or `Type` suffix.
- Every monolithic function exceeding 15 lines.
- Every nested `if` statement.

If there are NO discrepancies, explicitly state: "There are no coding guideline issues or discrepancies." However, assume the codebase is a mess until proven otherwise.

## 3. Root Cause & Fallout Analysis

For every issue found:

- What is the root cause? Why was it written this way?
- How many places does it need to be fixed?
- Fallout Check: If we change this, what else breaks? Will it break the CI/CD pipeline? Will it break tests? Map the entire blast radius.

## 4. Enqueueing Tasks for Sub-Agents

Your final output must be a massively detailed plan stored at `.lovable/plans/pending/01-coding-guideline-fixes.md` and granular subtask files written to `.lovable/plans/subtasks/01-coding-guideline-fixes/01-<subslug>.md`.
The plan must break the work down so granularly (exactly `N` steps) that 3 concurrent sub-agents can be spawned later to safely execute the fixes.

- Step 1..N: Exact file, exact line, exact boolean to rename, exact enum to extract. Keep the writing concise but hyper-specific. Do not write too much fluff.
- Do NOT fix the code in this turn. Your job is ONLY to plan, audit, and enqueue.
- Anti-Hallucination: If referenced guidelines or files are missing, ask clarifying questions rather than guessing.

---

## 5. Coding Guidelines Strict Adherence

/learn You MUST internalize the master coding guidelines located at `01-general-prompts/04-coding-standards/01-coding-guidelines.md`. It contains the ultimate source of truth for Boolean rules, Function limits, Error handling, and language-specific React/Go/Python paradigms. Do not hallucinate rules; enforce exactly what is in that file.

## Metadata

- slug: plan-coding-guideline-audit-v4
- status: active
