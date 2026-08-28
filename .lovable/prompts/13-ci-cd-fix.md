# Instruction (must follow): CI/CD Fix Loop (with RCA & End Aliasing)

Trigger Keywords & Aliases: `fix with RCA`, `fix`, `fix, fix`, `CI/CD fix`, `cicd fix`

/goal The user will provide an error log from a failed CI/CD pipeline or build run. Your objective is to ingest recent Root Cause Analysis (RCA) records from `.lovable/cicd-issues/`, `.lovable/issues/`, and `spec/03-error-manage/` before touching code, perform a surgical RCA on the failed build/test, fix the codebase to resolve the issue, record the memory of this failure in `.lovable/cicd-issues/01-<slug>.md` and `.lovable/strictly-avoid.md`, commit using the commit-fix workflow, and repeat the loop until the CI/CD pipeline is 100% green.

/learn Read past Root Cause Analyses (RCAs) from `.lovable/cicd-issues/` and `.lovable/issues/` so that past mistakes, regression patterns, and build traps are never repeated.

## MUST FOLLOW NON-NEGOTIABLE

Listen, past runs of these turns have been sloppy and stupid as fuck: wrong step counts, partial task lists dumped into chat instead of files, plans and session summaries half-filled with "[N]" placeholders, folders skimmed, open ambiguities ignored, CI/CD issues and `plans/subtasks/` forgotten, user commands dropped, coding guidelines bypassed, detailed specs chopped and summarized into useless junk, uppercase README files left uncorrected, `.lovable/memories/` created by accident, `strictly-avoid.md` overwritten, and explicit user instructions softened after being told not to. WTF. How on earth are you reverting to this carelessness, are you stupid?? Stop doing that, you stupid fuck. Read the whole codebase, read every folder in `spec/` and `.lovable/`, confirm root `readme.md` is strictly lowercase, find the root cause in one sentence, capture commands, issues, and pending tasks without omitting a single item, write the spec files and memory files in the right paths, update every index in the same turn, sync `readme.md` with `what-to-read.md`, preserve detailed specs verbatim with zero truncation, run builds and full unit tests, group commits with clear messages, and push everything to git before ending. Going deep IS the job. If you are not going deep, you are not doing the job. Violating this is auto-reject on the same tier as RULE 0. Avoid stupidity and being careless, you stupid fuck. Where is your attention, are you stupid? Tell me. Your stupidity is going on top of my head. Where did you learn this stupidity? If I could find you, I could slap you.

## Rules & Constraints (Non-Negotiable)

1. Analyze First & Read Past RCAs: Do not blindly change code. First read recent RCAs in `.lovable/cicd-issues/` and `.lovable/strictly-avoid.md`. Then trace the provided CI/CD error to the exact file, line, and dependency, performing a complete Root Cause Analysis (RCA).
2. Update Memory & Avoid List: The RCA and solution must be permanently recorded. Write the details to `.lovable/cicd-issues/01-<slug>.md` (sequenced as `01-`, `02-`, etc.) and register it in `.lovable/cicd-issues/index.md` (or `.lovable/cicd-index.md`). If a new forbidden pattern is identified, append it to `.lovable/strictly-avoid.md`.
3. Commit the Fix: Once the code is fixed, invoke the standard commit-fix procedure. Group changes logically with a clean, descriptive commit message (`fix(ci): <description>`).
4. Iterative Looping: If the pipeline fails again after your fix, the user will provide the new error. You must repeat this exact process—RCA, memory update, code fix, verification, commit, push—until the CI/CD run succeeds.
5. No Blind Overwrites: When updating memory, never delete or truncate existing history. Always append.
6. Anti-Hallucination Contract: If the cause is ambiguous or missing from logs, stop and ask clarifying questions instead of guessing.

## Actionable Items & Checklist

### 1. Pre-Flight & Past RCA Ingestion

- [ ] /learn past failure patterns in `.lovable/cicd-issues/` and `.lovable/strictly-avoid.md`.
- [ ] Read the provided CI/CD error log carefully.
- [ ] Identify the exact file, line, function, and dependency causing the failure.
- [ ] Formulate a one-sentence Root Cause Analysis followed by the full causal chain.

### 2. Memory Update (Mandatory)

- [ ] Create a new issue file at `.lovable/cicd-issues/01-<slug>.md`.
- [ ] Document: Error Summary, Root Cause Analysis, Solution Applied, and "What NOT to Repeat".
- [ ] Update `.lovable/cicd-issues/index.md` in the same operation.
- [ ] If a hard rule was broken, append a one-line prohibition to `.lovable/strictly-avoid.md`.

### 3. Execution & Code Fix

- [ ] Implement the minimal correct fix in the codebase based strictly on the RCA.
- [ ] Ensure the fix adheres to `spec/02-coding-guidelines/` and `spec/03-error-manage/`.
- [ ] Run local builds, linters, or unit tests if available to verify the fix before committing.

### 4. Verification & Final Checks

- [ ] Run all project tests (e.g., `go test ./...`, `npm test`, or equivalent test suites).
- [ ] If the project uses Go, run the race detector (`go test -race ./...`).
- [ ] Verify zero regressions, no swallowed errors, and no negative boolean anti-patterns.

### 5. Commit, Minor Release & Push

- [ ] Stage changes and commit using the commit-fix workflow (`fix(ci): <issue-slug>`).
- [ ] Tag a minor release or bump version if required by project release rules.
- [ ] Push the commit and tags to the remote repository.
- [ ] Await the next CI/CD result. If green, mark as completed; if failed, loop back to Step 1.

## Core Rules & Non-Negotiable Checklist for AI (Must Verify Before Completing Task)

Before finalizing any code modification, you MUST manually verify the following:

- [ ] Coding Guidelines & Master Consolidated File: I have fully read, checked, and strictly enforced every file in `spec/02-coding-guidelines/`, as well as the master consolidated coding guideline file at `.lovable/coding-guidelines/coding-guidelines.md`.
- [ ] Error Manage Checklist: I have fully read and enforced the error management files at `spec/03-error-manage/`. I understand which files to follow (architecture, response envelopes) and how to follow them (never swallow errors, always wrap with context).
- [ ] Boolean Examples & Fixations: All boolean variables MUST begin with `is`, `has`, `can`, or `should` (e.g., `isReady`, `hasData`). NEVER use negative booleans (e.g., `isNotReady`, `disableCache`). NEVER invert success checks (e.g., `!response.isSuccess` is banned; use `response.isFail`).
- [ ] Anti-Garbage Naming (Non-Negotiable): I have strictly verified that absolutely NO generic garbage variable names (e.g., `comp_100.go`, `temp`, `data`, `obj`, `Input100`, `TestHandleComp100`) were written. All names are highly semantic and domain-specific.
- [ ] Function Signatures (R4, R5, R9): If a function has `> 3 parameters` or the signature is `> 100 chars`, you MUST split it so there is exactly one parameter per line.
- [ ] Error Handling (R7): No silent failures or swallowed errors. Use explicit boolean states (e.g., `isFail`). Never invert success booleans (e.g., avoid `!isSuccess`).
- [ ] Magic Strings/Numbers (R8): Extract all magic strings/numbers into named constants.
- [ ] Enums: TypeScript string unions are banned. All Enums must end with the `Type` suffix.
- [ ] Naming & Casing (R1, R2): PascalCase everywhere. Acronyms (Id, Json, Url) are Pascal case, never all-caps (e.g., `UserId`, not `UserID`).
- [ ] Blank Lines (R13-R20): One blank line before every `return`/`throw`. One blank line after closing `}`. Never two blank lines in a row.

## End of Tunnel Release & Version Bump (Mandatory)

When EVERYTHING is completely finished and fixed (at the very end of the tunnel), you MUST trigger a release.

- [ ] You must bump the MINOR version.
- [ ] You must focus on the `version.json` file as the source of truth for the release.
- [ ] Root README Pinning (FATAL): You MUST pin the latest release version into the root `readme.md` file! Do not skip this! Also, update the changelog according to `version.json` format.
- [ ] If you do not know how to cut a release for this specific repository, or if `version.json` is missing/unclear, you must either search the repository for release instructions or explicitly ask the user for help. Do not guess.
- [ ] You MUST strictly exclude all test files (e.g., `*test*`, `*.spec.*`) from version scanning and modification, as they contain mock data.
- [ ] You must create and maintain `.lovable/memory/release-architecture-map.md` documenting exactly how releases work in the repository. Ensure it is enqueued in `what-to-read.md` and linked in the root `readme.md`.

## Metadata

- slug: cicd-fix
- status: active

## Version Bumping Rule

If the task involves a release or version bump, DO NOT manually search and replace versions across files. Execute `.lovable/release/bump_versions.py --type <major|minor|patch>`. If the script does not exist, trigger the first-time release bootstrapping phase to generate it.
